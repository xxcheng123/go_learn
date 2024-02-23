package geecache

import (
	"fmt"
	"gee"
	"net/http"
)

const defaultBasePath = "/hello"

type HttpPool struct {
	addr     string
	basePath string
	engine   *gee.Engine
}

func (h *HttpPool) Run() error {
	return h.engine.Run(h.addr)
}

func NewHttpPool(addr string) *HttpPool {
	h := &HttpPool{
		addr:     addr,
		basePath: defaultBasePath,
		engine:   gee.New(),
	}
	h.engine.GET(fmt.Sprintf("%s/:groupName/:key", h.basePath), func(ctx *gee.Context) {
		groupName := ctx.Param("groupName")
		key := ctx.Param("key")
		fmt.Printf("groupName:%s,key:%s\n", groupName, key)
		if group, ok := groups[groupName]; ok {
			if v, err := group.Get(key); err == nil {
				ctx.SetHeader("Content-Type", "application/octet-stream")
				ctx.Bytes(http.StatusOK, v.ByteSlice())
				return
			}
		}
		ctx.String(http.StatusOK, "not found")
	})
	return h
}
