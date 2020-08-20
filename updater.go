package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"fyne.io/fyne/widget"
	"github.com/melbahja/got"
	log "github.com/sirupsen/logrus"
	"lukechampine.com/blake3"
)

type LatestVersionInfo struct {
	Resources struct {
		Files []struct {
			Path     string `json:"path"`
			URL      string `json:"url"`
			Checksum string `json:"checksum"`
		} `json:"files"`
	} `json:"resources"`
}

func GetLatestGameVersion(distroinfo *DistroJSON) LatestVersionInfo {

	url := "https://raw.githubusercontent.com/kyoto44/AenorLauncher/master/assets/latest.json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	latestversioninfo := LatestVersionInfo{}
	json.Unmarshal(body, &latestversioninfo)
	return latestversioninfo
}

func Downloader(path string, url string, blake3sum string, s *widget.Label) {

	if _, err := os.Stat(path); err == nil {

		s.SetText("Проверяется файл " + path)
		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		hashsum := blake3.Sum256(file)

		if blake3sum == hex.EncodeToString(hashsum[:]) {
			log.Info(path + " не устарел!")
			return
		}

		d := got.Download{
			URL:         url,
			Dest:        path,
			ChunkSize:   uint64(1048576),
			Interval:    100,
			Concurrency: 10,
		}

		if err := d.Init(); err != nil {
			log.Fatal(err)
		}
		if err := d.Start(); err != nil {
			log.Fatal(err)
		}
		s.SetText("Cкачивается файл " + url)

	} else if os.IsNotExist(err) {
		create(path)
		d := got.Download{
			URL:         url,
			Dest:        path,
			ChunkSize:   uint64(1048576),
			Interval:    100,
			Concurrency: 10,
		}

		if err := d.Init(); err != nil {
			log.Fatal(err)
		}

		if err := d.Start(); err != nil {
			log.Fatal(err)
		}
		s.SetText("Cкачивается файл " + url)
	}
}

func Updater(distroinfo *DistroJSON, p *widget.ProgressBar, s *widget.Label) {

	latestversioninfo := GetLatestGameVersion(distroinfo)

	if runtime.GOOS == "windows" {
		gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/"
	} else if runtime.GOOS == "linux" {
		gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/"
	}

	log.Info("Проверка целостности игровых файлов")

	for i := 0; i < len(latestversioninfo.Resources.Files); i++ {
		Downloader(gamepath+latestversioninfo.Resources.Files[i].Path, latestversioninfo.Resources.Files[i].URL, latestversioninfo.Resources.Files[i].Checksum, s)
		currentvalue := float64(i+1) / float64(len(latestversioninfo.Resources.Files))
		p.SetValue(currentvalue)
	}

	fmt.Println()
	log.Info("Все файлы прошли проверку!")
}
