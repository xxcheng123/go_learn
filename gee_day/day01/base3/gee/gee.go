package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	if method == "" {
		panic("method can not be empty")
	}
	if handler == nil {
		panic("handler can not be nil")
	}
	e.router[fmt.Sprintf("%s-%s", method, pattern)] = handler
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
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
