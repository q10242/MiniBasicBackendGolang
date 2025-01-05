package interfaces

import "github.com/gorilla/mux"

// RouterRegister defines an interface for custom route registration
type RouterRegister interface {
	RegisterRoutes(router *mux.Router)
}
