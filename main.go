package main

import (
	"os"
	"os/exec"
	"runtime"

	app "github.com/kyoto44/AenorLauncher/backend"
	"github.com/leaanthony/mewn"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

// var (
// 	authdata   Authdata
// 	username   string
// 	distroinfo DistroJSON
// 	gamepath   string
// )

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:  980,
		Height: 552,
		Title:  "Aenor Launcher",
		JS:     js,
		CSS:    css,
		Colour: "#326fa8",
	})

	frontend.Bind(&app.WalletApplication{})
	err := frontend.Run()

	if err != nil {
		panic(err)
	}

	//authdata, loginstatus := GetAuthData(userNameEntry, passwordEntry)
	os.Setenv("LOGIN", authdata.Profile.Login)
	os.Setenv("TOKEN", authdata.AccessToken)
	distroinfo = GetDistroInfo(authdata.AccessToken)
	BackupSettings(&distroinfo)
	//SetDiscordStatus(nicknameEntry.Text, guildEntry.Text)

	//Updater(&distroinfo, progressBar, currentStatus)

	cmd := exec.Command("")
	if runtime.GOOS == "windows" {
		gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
		cmd = exec.Command(gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
	} else if runtime.GOOS == "linux" {
		gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/bin/nblade.exe"
		cmd = exec.Command("wine", gamepath, os.Getenv("LOGIN"), os.Getenv("TOKEN"))
	}
	cmd.Start()

	log.Info("Текущая версия клиента: " + distroinfo.Servers[0].Versions[0].ID)

	log.Info("Клиент запущен, приятной игры!")

	//time.Sleep(time.Duration(math.MaxInt64))
}
