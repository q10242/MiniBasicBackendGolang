package interfaces

import (
	"log"
	"net/http"
)

// Middleware 定義中介層的接口
type Middleware interface {
	Process(next http.Handler) http.Handler
}

// LoggingMiddleware 實現一個基礎的請求日誌記錄功能
type LoggingMiddleware struct{}

func (l *LoggingMiddleware) Process(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware 實現 API Key 驗證
type AuthMiddleware struct{}

func (a *AuthMiddleware) Process(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "Missing X-API-KEY header", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
