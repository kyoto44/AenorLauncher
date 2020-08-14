package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/hugolgst/rich-go/client"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Username string
	Password string
	Nickname string
	Guild    string
}

type LoginJSON struct {
	Agent       string `json:"agent"`
	Username    string `json:"username"`
	RequestUser bool   `json:"requestUser"`
	Password    string `json:"password"`
}

type SelectedProfile struct {
	Id        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type Authdata struct {
	AccessToken string          `json:"accessToken"`
	ClientToken string          `json:"clientToken"`
	Profile     SelectedProfile `json:"selectedProfile"`
}

type DistroJSON struct {
	Version string `json:"version"`
	Discord struct {
		ClientID       string `json:"clientId"`
		SmallImageText string `json:"smallImageText"`
		SmallImageKey  string `json:"smallImageKey"`
	} `json:"discord"`
	Rss     string `json:"rss"`
	Servers []struct {
		ID   string `json:"id"`
		Name struct {
			EnUS string `json:"en_US"`
		} `json:"name"`
		Description struct {
			EnUS string `json:"en_US"`
		} `json:"description"`
		Icon     string `json:"icon"`
		Versions []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"versions"`
		Address string `json:"address"`
		Discord struct {
			ShortID        string `json:"shortId"`
			LargeImageText string `json:"largeImageText"`
			LargeImageKey  string `json:"largeImageKey"`
		} `json:"discord"`
		MainServer  bool `json:"mainServer"`
		Autoconnect bool `json:"autoconnect"`
	} `json:"servers"`
}

func GetDistroInfo(AccessToken string) DistroJSON {

	url := "https://www.northernblade.ru/api/distribution/"
	method := "GET"

	payload := strings.NewReader("")

	clienthttp := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "BladeLauncher/1.0.17")
	req.Header.Add("Authorization", "Bearer "+AccessToken)

	res, err := clienthttp.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	distrojson := DistroJSON{}
	json.Unmarshal(body, &distrojson)

	//fmt.Println(distrojson)
	return distrojson

}

func GetAuthData() Authdata {
	auth, err := json.Marshal(LoginJSON{
		Agent:       "minecraftAgent",
		Username:    conf.Username,
		RequestUser: true,
		Password:    GetMD5Hash(conf.Password),
	})

	if err != nil {
		panic(err)
	}

	url := "https://www.northernblade.ru/api/authenticate"
	method := "POST"
	clienthttp := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(auth))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := clienthttp.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Warn("Неверный логин или пароль.")
		time.Sleep(7 * time.Second)
		panic(err)
	} else {
		log.Info("Успешная авторизация!")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	authdata := Authdata{}
	json.Unmarshal(body, &authdata)

	return authdata
}

var conf Config

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func BackupSettings(distroinfo *DistroJSON) {

	var versioninfofilepath string
	if runtime.GOOS == "windows" {
		versioninfofilepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/version.txt"
	} else if runtime.GOOS == "linux" {
		versioninfofilepath = UserHomeDir() + "/Northern Blade/version.txt"
	}

	if _, err := os.Stat(versioninfofilepath); err == nil {
		oldversion, err := ioutil.ReadFile(versioninfofilepath)
		if err != nil {
			panic(err)
		}

		if string(oldversion) == "" {
			ioutil.WriteFile(versioninfofilepath, []byte(distroinfo.Servers[0].Versions[0].ID), 0644)
		} else if string(oldversion) != distroinfo.Servers[0].Versions[0].ID {
			log.Info("Обнаружено обновление, сохраняем старые настройки...")
			var settings []byte
			if runtime.GOOS == "windows" {
				settings, err = ioutil.ReadFile(UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + string(oldversion) + "/profiles/preferences.xml")
			} else if runtime.GOOS == "linux" {
				settings, err = ioutil.ReadFile(UserHomeDir() + "/Northern Blade/" + string(oldversion) + "/profiles/preferences.xml")
			}
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(versioninfofilepath, []byte(distroinfo.Servers[0].Versions[0].ID), 0644)
			if runtime.GOOS == "windows" {
				ioutil.WriteFile(UserHomeDir()+"/AppData/Roaming/.nblade/instances/"+distroinfo.Servers[0].Versions[0].ID+"/profiles/preferences.xml", settings, 0644)
			} else if runtime.GOOS == "linux" {
				ioutil.WriteFile(UserHomeDir()+"/Northern Blade/"+distroinfo.Servers[0].Versions[0].ID+"/profiles/preferences.xml", settings, 0644)
			}
			log.Info("Настройки успешно сохранены")
		} else if os.IsNotExist(err) {
			return
		}
	}
}

var gamepath string

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	//Executing config
	if _, err := toml.DecodeFile("config.txt", &conf); err != nil {
		log.Warn("Конфигурационный файл config.txt не обнаружен!")
		time.Sleep(7 * time.Second)
		panic(err)
	}

	authdata := GetAuthData()
	distroinfo := GetDistroInfo(authdata.AccessToken)

	if runtime.GOOS == "windows" {
		if _, err := os.Stat(UserHomeDir() + "/AppData/Roaming/.nblade/instances/version.txt"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/AppData/Roaming/.nblade/instances/version.txt")
		}
	} else if runtime.GOOS == "linux" {
		if _, err := os.Stat(UserHomeDir() + "/Northern Blade/version.txt"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/Northern Blade/version.txt")
		}
	}

	BackupSettings(&distroinfo)

	Updater(&distroinfo)

	log.Info("Текущая версия клиента: " + distroinfo.Servers[0].Versions[0].ID)

	os.Setenv("LOGIN", authdata.Profile.Login)
	os.Setenv("TOKEN", authdata.AccessToken)

	//Starting game
	cmd := exec.Command("")
	if runtime.GOOS == "windows" {
		gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
		cmd = exec.Command(gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
	} else if runtime.GOOS == "linux" {
		gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
		cmd = exec.Command("wine "+gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
	}

	cmd.Start()
	log.Info("Клиент запущен, приятной игры!")

	var guildimage string
	switch conf.Guild {
	case "КАЭР МОРХЕН":
		guildimage = "kaermorhen"
	case "ИМПЕРИЯ":
		guildimage = "empire"
	case "FORCE":
		guildimage = "force"
	case "НЕМЕЗИДА":
		guildimage = "nemezida"
	case "МЕДВЕДИ":
		guildimage = "medvedi"
	case "ВЕТЕРАНЫ":
		guildimage = "veterani"
	default:
		guildimage = "sealcircle_photos_v2_x4"
	}

	client.Login("742666702298546207")
	now := time.Now()
	err := client.SetActivity(client.Activity{
		State:      conf.Guild,
		Details:    conf.Nickname,
		LargeImage: guildimage,
		LargeText:  conf.Guild,
		SmallImage: "",
		SmallText:  "",
		Party: &client.Party{
			ID:         "-1",
			Players:    1,
			MaxPlayers: 1,
		},
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(math.MaxInt64))
}
