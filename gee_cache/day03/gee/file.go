package gee

import (
	"net/http"
	"path"
)

func (r *RouterGroup) Static(relativePath, root string) {
	fileSystem := http.Dir(root)
	fileServer := http.FileServer(fileSystem)
	fileServer = http.StripPrefix(path.Join(r.prefix, relativePath), http.FileServer(fileSystem))
	r.GET(relativePath+"/*filepath", func(c *Context) {
		file := c.Param("filepath")
		if _, err := fileSystem.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
}
