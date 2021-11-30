package httpHandler

import (
	"log"
	"net/http"
)

type Router struct {
	hfs map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{hfs: map[string]HandlerFunc{}}
}

func (r *Router) addRoute(method string, pattern string, hf HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.hfs[key] = hf
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if hf, ok := r.hfs[key]; ok {
		hf(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}