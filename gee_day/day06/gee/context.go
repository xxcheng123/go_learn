package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	Path        string
	Method      string
	StatusCode  int
	Params      map[string]string
	middlewares []HandleFunc
	index       int
	engine      *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Writer: w, Request: r, Path: r.URL.Path, Method: r.Method, index: -1}
}

// 查询POST Form
// 查询Query
// 设置statusCode
// 设置响应header
// 设置响应body string/[]byte/json

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context) String(code int, format string, args ...any) {
	c.SetHeader("Content-Type", "text/plain;charset=utf-8")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, args...)))
}
func (c *Context) Bytes(code int, data []byte) {
	c.SetHeader("Content-Type", "text/plain;charset=utf-8")
	c.Status(code)
	c.Writer.Write(data)
}
func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json;charset=utf-8")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}
}
func (c *Context) HTML(code int, data string) {
	c.SetHeader("Content-Type", "text/html;charset=utf-8")
	c.Status(code)
	c.Writer.Write([]byte(data))
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.middlewares); c.index++ {
		c.middlewares[c.index](c)
	}
}
func (c *Context) Template(code int, name string, data any) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
