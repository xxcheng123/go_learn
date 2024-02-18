package gee

import "net/http"

type HandleFunc func(ctx *Context)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + pattern
	r.handlers[key] = handler
}

func (r *router) GET(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodGet, pattern, handler)
}
func (r *router) POST(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodPost, pattern, handler)
}

func (r *router) handle(c *Context) {
	if handler, ok := r.handlers[c.Method+c.Path]; ok {
		handler(c)
	} else {
		c.Status(http.StatusNotFound)
	}
}
