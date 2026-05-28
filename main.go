package main

import (
	"embed"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	// Cấu hình Middleware để chèn Header tùy chỉnh
	// Handler này sẽ bọc file server mặc định của Wails
	assetsHandler := assetserver.NewAssetsHandler(assets)
	
	customHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Thêm header tùy chỉnh vào mọi response từ AssetServer
		w.Header().Set("X-Abcsnoob-webview", "true")
		
		// Tiếp tục phục vụ file từ assets
		assetsHandler.ServeHTTP(w, r)
	})

	err := wails.Run(&options.App{
		Title:  "Abc's Noob Social",
		Width:  1280,
		Height: 800,

		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: customHandler, // Áp dụng handler đã tùy chỉnh
		},

		BackgroundColour: &options.RGBA{R: 11, G: 14, B: 17, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},

		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Mica,
			WebviewUserDataPath:  "appdata",
		},
	})

	if err != nil {
		println("Lỗi rồi ông giáo ơi:", err.Error())
	}
}
