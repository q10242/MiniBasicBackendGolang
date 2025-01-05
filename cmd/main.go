package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/q10242/MiniBasicBackendGolang/lib"
	"github.com/q10242/MiniBasicBackendGolang/pkg/base"
)

// CustomRouter implements RouterRegister for custom routes
type CustomRouter struct{}
type GenericResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// RegisterRoutes defines the routes for the application
func (c *CustomRouter) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		lib.JsonResponse(w, http.StatusOK, "Health Check OK", "OK")
	}).Methods(http.MethodGet)
}

func main() {
	// 初始化伺服器
	srv := base.NewServer(":8080")

	// 添加預設中介軟體
	srv.AddMiddleware(base.LoggingMiddleware)

	// 註冊客製化路由
	customRouter := &CustomRouter{}
	srv.RegisterCustomRoutes(customRouter)

	// 啟動伺服器
	log.Fatal(srv.Run())
}
