package gee

import (
	"net/http"
)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	if method == "" {
		panic("method can not be empty")
	}
	if handler == nil {
		panic("handler can not be nil")
	}
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
