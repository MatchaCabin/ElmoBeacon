package main

import (
	"ElmoBeacon/db"
	"ElmoBeacon/handler"
	"ElmoBeacon/logger"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.InitLogger()
	db.InitDB()
	// Create an instance of the app structure
	app := handler.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "ElmoBeacon",
		MinWidth:  1600,
		MinHeight: 900,
		Width:     1600,
		Height:    900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
