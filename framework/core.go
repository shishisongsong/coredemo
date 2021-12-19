package framework

import "net/http"

// Core is a core handler
type Core struct{}

// NewCore is the constructor for Core
func NewCore() *Core {
	return &Core{}
}

// ServeHTTP implements http.Handler interface
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
