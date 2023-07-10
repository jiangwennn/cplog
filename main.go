package main

import (
	"context"
	myapp "cplog/app"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	storeDir := app.storeDir()
	config := myapp.NewConfig(storeDir)
	clipboard := myapp.NewClipboard(storeDir, config)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "cplog",
		MinWidth: 1024,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Mac: &mac.Options{
			WindowIsTranslucent: true,
			WebviewIsTransparent: true,
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title: "cplog",
				Message: "",
			},
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 8},
		OnStartup:  func(ctx context.Context) {
			app.startup(ctx)
			clipboard.StartUp(ctx)
			config.StartUp(ctx)
		},
		Bind: []interface{}{
			app,
			clipboard,
			config,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
