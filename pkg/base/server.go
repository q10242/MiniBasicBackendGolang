package base

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/q10242/MiniBasicBackendGolang/pkg/interfaces"
)

type Server struct {
	Router *mux.Router
	Addr   string
}

// NewServer creates a new HTTP server instance
func NewServer(addr string) *Server {
	return &Server{
		Router: mux.NewRouter(),
		Addr:   addr,
	}
}

// AddMiddleware adds global middlewares to the router
func (s *Server) AddMiddleware(middleware ...mux.MiddlewareFunc) {
	for _, m := range middleware {
		s.Router.Use(m)
	}
}

// RegisterCustomRoutes allows custom routes to be registered
func (s *Server) RegisterCustomRoutes(register interfaces.RouterRegister) {
	register.RegisterRoutes(s.Router)
}

// Run starts the HTTP server
func (s *Server) Run() error {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         s.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is running at %s\n", s.Addr)
	return srv.ListenAndServe()
}
