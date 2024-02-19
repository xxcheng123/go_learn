package gee

import (
	"net/http"
	"strings"
)

type Engine struct {
	router *router
	groups []*RouterGroup
	*RouterGroup
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
		groups: make([]*RouterGroup, 0),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

type RouterGroup struct {
	prefix      string
	engine      *Engine
	parent      *RouterGroup
	middlewares []HandleFunc
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	newRouterGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		engine: r.engine,
		parent: r,
	}
	r.engine.groups = append(r.engine.groups, newRouterGroup)
	return newRouterGroup
}
func (r *RouterGroup) Use(middlewares ...HandleFunc) *RouterGroup {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

//	func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
//		if method == "" {
//			panic("method can not be empty")
//		}
//		if handler == nil {
//			panic("handler can not be nil")
//		}
//		e.router.addRoute(method, pattern, handler)
//	}
func (r *RouterGroup) addRoute(method string, pattern string, handler HandleFunc) {
	if method == "" {
		panic("method can not be empty")
	}
	if handler == nil {
		panic("handler can not be nil")
	}
	//fmt.Println(r.prefix + pattern)
	r.engine.router.addRoute(method, r.prefix+pattern, handler)
}
func (r *RouterGroup) GET(pattern string, handler HandleFunc) {
	r.addRoute("GET", pattern, handler)
}
func (r *RouterGroup) POST(pattern string, handler HandleFunc) {
	r.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	middlewares := make([]HandleFunc, 0)
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c.middlewares = middlewares
	e.router.handle(c)
}

//func (e *Engine) GET(pattern string, handler HandleFunc) {
//	e.addRoute("GET", pattern, handler)
//}
//func (e *Engine) POST(pattern string, handler HandleFunc) {
//	e.addRoute("POST", pattern, handler)
//}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
