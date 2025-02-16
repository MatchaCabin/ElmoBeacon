package main

import (
	"ElmoBeacon/db"
	"ElmoBeacon/handler"
	"ElmoBeacon/logger"
	"embed"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	"path/filepath"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.InitLogger()
	db.InitDB()
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Panic().Err(err).Msg("Could not get user config dir")
	}
	// Create an instance of the app structure
	app := handler.NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "ElmoBeacon",
		Width:     1600,
		Height:    900,
		Frameless: true,
		MinWidth:  1600,
		MinHeight: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewUserDataPath: filepath.Join(userConfigDir, "MccWiki", "ElmoBeacon"),
		},
	})

	if err != nil {
		log.Panic().Err(err).Msg("Could not start app")
	}
}
