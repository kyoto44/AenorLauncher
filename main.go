package main

import (
	"github.com/leaanthony/mewn"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"

	app "github.com/kyoto44/aenorlauncher/backend"
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

	frontend.Bind(&app.LauncherApplication{})
	err := frontend.Run()

	if err != nil {
		panic(err)
	}

	//authdata, loginstatus := GetAuthData(userNameEntry, passwordEntry)

	//BackupSettings(&distroinfo)
	//SetDiscordStatus(nicknameEntry.Text, guildEntry.Text)

	//Updater(&distroinfo, progressBar, currentStatus)

	// log.Info("Текущая версия клиента: " + distroinfo.Servers[0].Versions[0].ID)

	// log.Info("Клиент запущен, приятной игры!")

	//time.Sleep(time.Duration(math.MaxInt64))
}
