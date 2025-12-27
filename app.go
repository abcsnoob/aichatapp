package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp tạo instance mới
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CheckConnection: Go sẽ kiểm tra mạng thay vì để JS làm (nhanh và chính xác hơn)
func (a *App) CheckConnection(url string) bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	_, err := client.Get(url)
	if err != nil {
		return false
	}
	return true
}

// Greet: Ví dụ hàm Go trả về lời chào cho Frontend
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Chào ông giáo %s, hệ thống Noob AI đã sẵn sàng!", name)
}