package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	app := NewApp()

	// Thiết lập đường dẫn dữ liệu dựa trên thư mục chạy ứng dụng
	// Bạn có thể thay đổi "appdata" thành tên ứng dụng của bạn để tránh xung đột
	userDataPath := "appdata"

	err := wails.Run(&options.App{
		Title:            "Abc's Noob Social",
		Width:            1280,
		Height:           800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 11, G: 14, B: 17, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		// Cấu hình riêng cho từng OS
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Mica,
			WebviewUserDataPath:  userDataPath,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
			},
			WebviewUserDataPath: userDataPath,
		},
		Linux: &linux.Options{
			WebviewUserDataPath: userDataPath,
		},
	})

	if err != nil {
		println("Lỗi rồi ông giáo ơi:", err.Error())
	}
}
