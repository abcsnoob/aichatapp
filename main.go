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
	// Khởi tạo ứng dụng
	app := NewApp()

	// Khởi tạo Wails
	err := wails.Run(&options.App{
		Title:  "Abc's Noob Social",
		Width:  1280,
		Height: 800,

		AssetServer: &assetserver.Options{
			Assets: assets,
			// Sử dụng Middleware để chèn Header tùy chỉnh vào mọi response
			Middleware: func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Thêm header tùy chỉnh
					w.Header().Set("X-Abcsnoob-webview", "true")
					
					// Chuyển tiếp request đến handler mặc định của Wails
					next.ServeHTTP(w, r)
				})
			},
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
