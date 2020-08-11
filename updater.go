package main

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
	"unsafe"

	"github.com/cavaliercoder/grab"
	"github.com/cheggaaa/pb"
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

func GetLatestGameVersion(distroinfo *DistroJSON) LatestVersionInfo {

	url := "http://downloads.n-blade.ru/dist/versions/" + distroinfo.Servers[0].Versions[0].ID + ".json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))

	latestversioninfo := LatestVersionInfo{}
	json.Unmarshal(body, &latestversioninfo)
	return latestversioninfo
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func Downloader(path string, link string, sha512sum string) {

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
		fmt.Println(path + " up-to-date!")
		return
	}

	client := grab.NewClient()
	req, _ := grab.NewRequest(path, link)
	// start download
	fmt.Printf("Скачивается %v...\n", req.URL())
	resp := client.Do(req)

	t := time.NewTicker(200 * time.Millisecond)
	defer t.Stop()

	bar := pb.StartNew(int(resp.Size()) / 1024 / 1024)
	bar.Start()

Loop:
	for {
		select {
		case <-t.C:
			bar.Set64(resp.BytesComplete() / 1024 / 1024)
		case <-resp.Done:
			// download is complete
			bar.Finish()
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при загрузке: %v\n", err)
		os.Exit(1)
	}

}

func Updater(distroinfo *DistroJSON) {

	latestversioninfo := GetLatestGameVersion(distroinfo)
	path := UserHomeDir() + "/AppData/Roaming/.nblade/instances/" + distroinfo.Servers[0].Versions[0].ID + "/"

	/* Game files */
	Downloader(path+latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFevSeq.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFdp.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtFev.Artifact.Path, latestversioninfo.Downloads.ResAudioFtFev.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtFev.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank02Fsb.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank02FsbSeq.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank04Fsb.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Path, latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Urls[0], latestversioninfo.Downloads.ResAudioFtBank04FsbSeq.Artifact.Checksum)

	Downloader(path+latestversioninfo.Downloads.PacksBwZpk.Artifact.Path, latestversioninfo.Downloads.PacksBwZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksBwZpk.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.PacksMapsZpk.Artifact.Path, latestversioninfo.Downloads.PacksMapsZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksMapsZpk.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.PacksResZpk.Artifact.Path, latestversioninfo.Downloads.PacksResZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksResZpk.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Path, latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Urls[0], latestversioninfo.Downloads.PacksSpacesZpk.Artifact.Checksum)

	Downloader(path+latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Path, latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3Dx931Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX940Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX941Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Path, latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinD3DX942Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Path, latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Urls[0], latestversioninfo.Downloads.BinFmodEventNetDll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinFmodexDll.Artifact.Path, latestversioninfo.Downloads.BinFmodexDll.Artifact.Urls[0], latestversioninfo.Downloads.BinFmodexDll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Path, latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinLibeay32Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMfc80Dll.Artifact.Path, latestversioninfo.Downloads.BinMfc80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfc80Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMfc80UDll.Artifact.Path, latestversioninfo.Downloads.BinMfc80UDll.Artifact.Urls[0], latestversioninfo.Downloads.BinMfc80UDll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Path, latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Urls[0], latestversioninfo.Downloads.BinMicrosoftVC80CRTManifest.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Path, latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Urls[0], latestversioninfo.Downloads.BinMicrosoftVC80MFCManifest.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Path, latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMsvcm80Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Path, latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinMsvcp80Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCP90DLL.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCR80DLL.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Path, latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Urls[0], latestversioninfo.Downloads.BinMSVCR90DLL.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinNbladeExe.Artifact.Path, latestversioninfo.Downloads.BinNbladeExe.Artifact.Urls[0], latestversioninfo.Downloads.BinNbladeExe.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinSplashBmp.Artifact.Path, latestversioninfo.Downloads.BinSplashBmp.Artifact.Urls[0], latestversioninfo.Downloads.BinSplashBmp.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Path, latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinSsleay32Dll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinVoipDll.Artifact.Path, latestversioninfo.Downloads.BinVoipDll.Artifact.Urls[0], latestversioninfo.Downloads.BinVoipDll.Artifact.Checksum)
	Downloader(path+latestversioninfo.Downloads.BinZlib1Dll.Artifact.Path, latestversioninfo.Downloads.BinZlib1Dll.Artifact.Urls[0], latestversioninfo.Downloads.BinZlib1Dll.Artifact.Checksum)
}
