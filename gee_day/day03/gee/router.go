package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(ctx *Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func newRouter() *router {
	roots := make(map[string]*node, 8)
	httpMethods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
	for _, method := range httpMethods {
		roots[method] = new(node)
	}
	return &router{
		roots:    roots,
		handlers: make(map[string]HandleFunc),
	}
}
func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	vs := strings.Split(pattern, "/")
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}
func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + pattern
	r.roots[method].insert(pattern, parsePattern(pattern), 0)
	r.handlers[key] = handler
}
func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)
	params := make(map[string]string)
	n := r.roots[method].search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
func (r *router) GET(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodGet, pattern, handler)
}
func (r *router) POST(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodPost, pattern, handler)
}

func (r *router) handle(c *Context) {
	route, params := r.getRoute(c.Method, c.Path)
	if route == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Params = params
	if handler, ok := r.handlers[c.Method+route.pattern]; ok {
		handler(c)
	} else {
		c.Status(http.StatusNotFound)
	}
}
