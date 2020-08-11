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

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "BladeLauncher/1.0.17")
	req.Header.Add("Authorization", "Bearer "+AccessToken)

	res, err := client.Do(req)
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

func main() {

	//Executing config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic("config.toml doesn't exist")
	}

	authdata := GetAuthData()
	distroinfo := GetDistroInfo(authdata.AccessToken)

	Updater(&distroinfo)

	fmt.Println("Текущая версия клиента: " + distroinfo.Servers[0].Versions[0].ID)

	os.Setenv("LOGIN", authdata.Profile.Login)
	os.Setenv("TOKEN", authdata.AccessToken)

	cmd := exec.Command(UserHomeDir()+"\\AppData\\Roaming\\.nblade\\instances\\"+distroinfo.Servers[0].Versions[0].ID+"\\bin\\nblade.exe", os.Getenv("LOGIN"), os.Getenv("TOKEN"))
	cmd.Start()
	fmt.Println("Клиент запущен, приятной игры!")

	client.Login("742666702298546207")
	now := time.Now()
	err := client.SetActivity(client.Activity{
		State:      conf.Guild,
		Details:    conf.Nickname,
		LargeImage: "sealcircle_photos_v2_x4",
		LargeText:  "Северный Клинок",
		SmallImage: "",
		SmallText:  "",
		Party: &client.Party{
			ID:         "-1",
			Players:    1,
			MaxPlayers: 5,
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
