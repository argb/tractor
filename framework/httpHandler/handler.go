package httpHandler

import (
	"net/http"
)

type HashMap map[string]interface{}
type HandlerFunc func(ctx *Context)

type Handler struct {
	router *Router
}

func New() *Handler {
	return &Handler{router: NewRouter()}
}

func (h *Handler) addRoute(method string, pattern string, hf HandlerFunc) {
	h.router.addRoute(method, pattern, hf)
}

func (h *Handler) GET(pattern string, hf HandlerFunc) {
	h.addRoute("GET", pattern, hf)
}

func (h *Handler) POST(pattern string, hf HandlerFunc) {
	h.addRoute("POST", pattern, hf)
}

func (h *Handler) Run(addr string) error {
	return http.ListenAndServe(addr, h)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	h.router.handle(c)
}
