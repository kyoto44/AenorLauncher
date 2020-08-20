package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
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

func GetAuthData(username *widget.Entry, password *widget.Entry) (Authdata, int) {
	auth, err := json.Marshal(LoginJSON{
		Agent:       "minecraftAgent",
		Username:    username.Text,
		RequestUser: true,
		Password:    GetMD5Hash(password.Text),
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
	/*if res.StatusCode != http.StatusOK {
		log.Warn("Неверный логин или пароль.")
		time.Sleep(7 * time.Second)
		panic(err)
	} else {
		log.Info("Успешная авторизация!")
	}
	*/
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	authdata := Authdata{}
	json.Unmarshal(body, &authdata)

	return authdata, res.StatusCode
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

func createLoginForm() *widget.Form {
	userNameEntry := widget.NewEntry()
	userNameEntry.Resize(fyne.NewSize(300, 200))
	userNameEntry.KeyDown(&fyne.KeyEvent{
		Name: fyne.KeyDown,
	})
	userNameFormItem := &widget.FormItem{
		Text:   "Логин",
		Widget: userNameEntry,
	}

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.Resize(fyne.NewSize(300, 200))
	passwordEntry.KeyDown(&fyne.KeyEvent{
		Name: fyne.KeyDown,
	})
	passwordFormItem := &widget.FormItem{
		Text:   "Пароль",
		Widget: passwordEntry,
	}

	nicknameEntry := widget.NewEntry()
	nicknameEntry.Resize(fyne.NewSize(300, 200))
	nicknameEntry.KeyDown(&fyne.KeyEvent{
		Name: fyne.KeyDown,
	})
	nicknameFormItem := &widget.FormItem{
		Text:   "Имя персонажа",
		Widget: nicknameEntry,
	}

	guildEntry := widget.NewEntry()
	guildEntry.Resize(fyne.NewSize(300, 200))
	guildFormItem := &widget.FormItem{
		Text:   "Гильдия",
		Widget: guildEntry,
	}

	progressBar := widget.NewProgressBar()
	progressBarItem := &widget.FormItem{
		Widget: progressBar,
	}

	currentStatus := widget.NewLabel("")
	currentStatusItem := &widget.FormItem{
		Widget: currentStatus,
	}

	loginForm := widget.NewForm(userNameFormItem, passwordFormItem, nicknameFormItem, guildFormItem, currentStatusItem, progressBarItem)
	loginForm.OnSubmit = func() {
		authdata, loginstatus := GetAuthData(userNameEntry, passwordEntry)

		if loginstatus != 200 {
			currentStatus.SetText("Неверный логин или пароль!")

		} else {
			os.Setenv("LOGIN", authdata.Profile.Login)
			os.Setenv("TOKEN", authdata.AccessToken)
			distroinfo = GetDistroInfo(authdata.AccessToken)
			BackupSettings(&distroinfo)
			SetDiscordStatus(nicknameEntry.Text, guildEntry.Text)

			Updater(&distroinfo, progressBar, currentStatus)

			cmd := exec.Command("")
			if runtime.GOOS == "windows" {
				gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
				cmd = exec.Command(gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
			} else if runtime.GOOS == "linux" {
				gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
				cmd = exec.Command("wine", gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
			}
			cmd.Start()
			//loginForm.Hide()
		}

	}
	loginForm.SubmitText = "Войти в игру"
	//loginForm.OnCancel = Cancel
	return loginForm
}

func SetDiscordStatus(nickname string, guild string) {

	var guildimage string
	switch guild {
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
		State:      guild,
		Details:    nickname,
		LargeImage: guildimage,
		LargeText:  guild,
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
}

var authdata Authdata
var username string
var distroinfo DistroJSON
var gamepath string

func CreateRequiredDirs() {

	if runtime.GOOS == "windows" {
		if _, err := os.Stat(UserHomeDir() + "/AppData/Roaming/.nblade/instances/version.txt"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/AppData/Roaming/.nblade/instances/version.txt")
		}

		if _, err := os.Stat(UserHomeDir() + "/AppData/Roaming/.nblade/instances/common/config/temp/"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/AppData/Roaming/.nblade/instances/common/config/temp/")
		}
	} else if runtime.GOOS == "linux" {
		if _, err := os.Stat(UserHomeDir() + "/Northern Blade/version.txt"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/Northern Blade/version.txt")
		}

		if _, err := os.Stat(UserHomeDir() + "/common/config/temp/"); err == nil {
		} else if os.IsNotExist(err) {
			create(UserHomeDir() + "/common/config/temp/")
		}
	}
}

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	LauncherApp := app.New()
	LauncherWindow := LauncherApp.NewWindow("Aenor Launcher")

	loginForm := createLoginForm()
	loginForm.Show()

	/*
		loginWidget := widget.NewVBox(
			widget.NewButton("Login", func() {
				widget.NewModalPopUp(loginForm, LauncherWindow.Canvas())
			}),
		) */

	LauncherWindow.CenterOnScreen()
	LauncherWindow.SetContent(loginForm)
	LauncherWindow.SetFixedSize(true)
	LauncherWindow.Resize(fyne.NewSize(840, 240))

	/*progressBar := widget.NewProgressBar()
	gameStarter := widget.NewVBox(
		loginForm,
		progressBar,
		widget.NewButton("Играть", func() {
			loginForm.Hide()
			Updater(&distroinfo, progressBar)

			cmd := exec.Command("")
			if runtime.GOOS == "windows" {
				gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
				cmd = exec.Command(gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
			} else if runtime.GOOS == "linux" {
				gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
				cmd = exec.Command("wine", gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
			}
			cmd.Start()
		}),
	)
	LauncherWindow.SetContent(gameStarter)
	gameStarter.Hide() */
	LauncherWindow.ShowAndRun()

	log.Info("Текущая версия клиента: " + distroinfo.Servers[0].Versions[0].ID)

	/* //Starting game
	 */
	log.Info("Клиент запущен, приятной игры!")

	//time.Sleep(time.Duration(math.MaxInt64))
}
