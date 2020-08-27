package backend

import (
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

type LauncherApplication struct {
	Version string
	RT      *wails.Runtime
	log     *logrus.Logger

	authdata   Authdata
	username   string
	distroinfo DistroJSON
	gamepath   string
	status     string
	pbvalue    int

	killSignal chan struct{}

	/*
		Network struct {
			URL     string
			Handles struct {
				Send        string // Takes TX Object, returns TX Hash (200)
				Transaction string // Takes TX Object, returns TX Hash (200)
				Balance     string // Polls the wallets available balance
			}
			BlockExplorer struct {
				URL     string
				Handles struct {
					Transactions string // Takes TX Hash, returns TX info
					Checkpoints  string // Takes Checkpoint block hash, returns checkpoint block info
					Snapshots    string // Takes SnapshotHash, returns info
					CollectTX    string // Takes DAG address, returns tx objects

				}
			}
		}
		paths struct {
			HomeDir      string
			DAGDir       string
			TMPDir       string
			EncryptedDir string
			EmptyTXFile  string
			PrevTXFile   string
			LastTXFile   string
			AddressFile  string
			ImageDir     string
			Java         string
		}
	*/
}

// func (a *LauncherApplication) WailsShutdown() {
// 	a.wallet = models.Wallet{}
// 	close(a.killSignal) // Kills the Go Routines
// 	a.DB.Close()
// }

// WailsInit initializes the Client and Server side bindings
func (a *LauncherApplication) WailsInit(runtime *wails.Runtime) error {
	//var err error

	a.log = logrus.New()
	a.RT = runtime
	a.pbvalue = 0
	// err = a.initDirectoryStructure()
	// if err != nil {
	// 	a.log.Errorln("Unable to set up directory structure. Reason: ", err)
	// }

	// a.initLogger()

	// err = api.InitRPCServer()
	// if err != nil {
	// 	a.log.Panicf("Unable to initialize RPC Server. Reason: %v", err)
	// }
	// a.log.Infoln("RPC Server initialized.")

	// a.UserLoggedIn = false
	// a.NewUser = false
	// a.TransactionFinished = true
	//
	// a.killSignal = make(chan struct{}) // Used to kill go routines and hand back system resources
	// a.wallet.Currency = "USD"          // Set default currency
	// a.WalletCLI.URL = "https://github.com/Constellation-Labs/constellation/releases/download"
	// a.WalletCLI.Version = "2.6.0"
	// a.Version = "1.2.0"

	// a.DB, err = gorm.Open("sqlite3", a.paths.DAGDir+"/store.db")
	// if err != nil {
	// 	a.log.Panicln("failed to connect database", err)
	// }
	// // Migrate the schema
	// a.DB.AutoMigrate(&models.Wallet{}, &models.TXHistory{}, &models.Path{})
	// a.detectJavaPath()
	// a.initMainnetConnection()
	// a.newReleaseAvailable()

	return nil
}
