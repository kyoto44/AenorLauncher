package main

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/melbahja/got"
	log "github.com/sirupsen/logrus"
)

type LatestVersionInfo struct {
	Modifiers []struct {
		Rules []struct {
			Type   string `json:"type"`
			Ensure string `json:"ensure"`
		} `json:"rules"`
		Path string `json:"path"`
	} `json:"modifiers"`
	Downloads struct {
		ResAudioFtFevSeq struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft.fev.seq"`
		ResAudioFtFdp struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft.fdp"`
		BinMfc80Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/mfc80.dll"`
		ResAudioFtBank02Fsb struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft_bank02.fsb"`
		ResAudioFtBank02FsbSeq struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft_bank02.fsb.seq"`
		ResAudioFtBank04Fsb struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft_bank04.fsb"`
		BinLibeay32Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/libeay32.dll"`
		BinD3DX941Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/D3DX9_41.dll"`
		ResAudioFtFev struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft.fev"`
		ResAudioFtBank04FsbSeq struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"res/audio/ft_bank04.fsb.seq"`
		BinSsleay32Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/ssleay32.dll"`
		BinD3Dx931Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/d3dx9_31.dll"`
		BinMfcm80UDll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/mfcm80u.dll"`
		BinMSVCR80DLL struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/MSVCR80.DLL"`
		PacksResZpk struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"packs/res.zpk"`
		BinMSVCR90DLL struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/MSVCR90.DLL"`
		BinNbladeExe struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/nblade.exe"`
		BinMicrosoftVC80CRTManifest struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/Microsoft.VC80.CRT.manifest"`
		BinMsvcm80Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/msvcm80.dll"`
		BinVoipDll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/voip.dll"`
		BinD3DX940Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/D3DX9_40.dll"`
		BinMfcm80Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/mfcm80.dll"`
		BinMsvcp80Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/msvcp80.dll"`
		BinMSVCP90DLL struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/MSVCP90.DLL"`
		BinFmodEventNetDll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/fmod_event_net.dll"`
		BinSplashBmp struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/splash.bmp"`
		PacksBwZpk struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"packs/bw.zpk"`
		BinPathsXMLEjs struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/paths.xml.ejs"`
		BinMicrosoftVC80MFCManifest struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/Microsoft.VC80.MFC.manifest"`
		PacksSpacesZpk struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"packs/spaces.zpk"`
		PacksMapsZpk struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"packs/maps.zpk"`
		BinD3DX942Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/D3DX9_42.dll"`
		BinZlib1Dll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/zlib1.dll"`
		BinMfc80UDll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/mfc80u.dll"`
		BinFmodexDll struct {
			Type     string `json:"type"`
			Artifact struct {
				Path     string   `json:"path"`
				Size     int      `json:"size"`
				Urls     []string `json:"urls"`
				Checksum string   `json:"checksum"`
			} `json:"artifact"`
		} `json:"bin/fmodex.dll"`
	} `json:"downloads"`
	Manifest struct {
		Game struct {
			LaunchModuleID string        `json:"launchModuleId"`
			Arguments      []interface{} `json:"arguments"`
		} `json:"game"`
	} `json:"manifest"`
	MinimumLauncherVersion int    `json:"minimumLauncherVersion"`
	Type                   string `json:"type"`
	ID                     string `json:"id"`
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func GetLatestGameVersion(distroinfo *DistroJSON) LatestVersionInfo {

	url := "http://downloads.n-blade.ru/dist/versions/" + distroinfo.Servers[0].Versions[0].ID + ".json"
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

func Downloader(path string, url string, sha512sum string) {

	if _, err := os.Stat(path); err == nil {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}

		hashsum := h.Sum(nil)
		sha512bytes := []byte(sha512sum[7:])

		if string(sha512bytes) == hex.EncodeToString(hashsum[:]) {
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

		d.Progress.ProgressFunc = func(p *got.Progress, d *got.Download) {
			fmt.Printf(
				"\r\r\bЗагружается %s | Размер: %s | Загружено: %s | Скорость: %s/s",
				url,
				humanize.Bytes(uint64(p.TotalSize)),
				humanize.Bytes(uint64(p.Size)),
				humanize.Bytes(p.Speed()),
			)
		}

		if err := d.Start(); err != nil {
			log.Fatal(err)
		}

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

		d.Progress.ProgressFunc = func(p *got.Progress, d *got.Download) {
			fmt.Printf(
				"\r\r\bЗагружается %s | Размер: %s | Загружено: %s | Скорость: %s/s",
				url,
				humanize.Bytes(uint64(p.TotalSize)),
				humanize.Bytes(uint64(p.Size)),
				humanize.Bytes(p.Speed()),
			)
		}
		if err := d.Start(); err != nil {
			log.Fatal(err)
		}
	}
}

func Updater(distroinfo *DistroJSON) {

	latestversioninfo := GetLatestGameVersion(distroinfo)

	if runtime.GOOS == "windows" {
		gamepath = UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/"
	} else if runtime.GOOS == "linux" {
		gamepath = UserHomeDir() + "/Northern Blade/" + distroinfo.Servers[0].Versions[0].ID + "/"
	}

	log.Info("Проверка целостности игровых файлов")
	/*TODO: Refactor is needed, thanks to Northern Blade devs...*/
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtFev.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFev.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFev.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Checksum)

	Downloader(gamepath+latestversioninfo.Downloads.PacksBwZpk.Artifact.Path, latestversioninfo.Downloads.PacksBwZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksBwZpk.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.PacksMapsZpk.Artifact.Path, latestversioninfo.Downloads.PacksMapsZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksMapsZpk.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.PacksResZpk.Artifact.Path, latestversioninfo.Downloads.PacksResZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksResZpk.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Path, latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Checksum)

	Downloader(gamepath+latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Path, latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Path, latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Urls[0], latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinFmodexDll.Artifact.Path, latestversioninfo.Downloads.BinFmodexDll.Artifact.Urls[0], latestversioninfo.Downloads.BinFmodexDll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Path, latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMfc80Dll.Artifact.Path, latestversioninfo.Downloads.BinMfc80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfc80Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMfc80UDll.Artifact.Path, latestversioninfo.Downloads.BinMfc80UDll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfc80UDll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMfcm80Dll.Artifact.Path, latestversioninfo.Downloads.BinMfcm80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfcm80Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMfcm80UDll.Artifact.Path, latestversioninfo.Downloads.BinMfcm80UDll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfcm80UDll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Path, latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Urls[0], latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Path, latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Urls[0], latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Path, latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Path, latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinNbladeExe.Artifact.Path, latestversioninfo.Downloads.BinNbladeExe.Artifact.Urls[0], latestversioninfo.Downloads.BinNbladeExe.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinSplashBmp.Artifact.Path, latestversioninfo.Downloads.BinSplashBmp.Artifact.Urls[0], latestversioninfo.Downloads.BinSplashBmp.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Path, latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinVoipDll.Artifact.Path, latestversioninfo.Downloads.BinVoipDll.Artifact.Urls[0], latestversioninfo.Downloads.BinVoipDll.Artifact.Checksum)
	Downloader(gamepath+latestversioninfo.Downloads.BinZlib1Dll.Artifact.Path, latestversioninfo.Downloads.BinZlib1Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinZlib1Dll.Artifact.Checksum)
	Downloader(gamepath+"bin/paths.xml", "https://kyoto44.com/paths.xml", "sha512:c6020b7b7d1930ca359bf460c621567f33f94b99ff6294fee5192a08a6f690ab8facb990db084158ae3ba432f684cb348d940d74748cab93197adedf72b60efb")

	fmt.Println()
	log.Info("Все файлы прошли проверку!")
}
