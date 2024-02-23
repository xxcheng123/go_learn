package geecache

import (
	"fmt"
	"gcache/consistenthash"
	gee "gcache/gee"
	"io"
	"net/http"
	"sync"
)

const defaultBasePath = "/hello"

type HttpPool struct {
	addr        string
	basePath    string
	engine      *gee.Engine
	mu          sync.Mutex
	peers       *consistenthash.Map
	httpGetters map[string]*HttpGetter
}

func (h *HttpPool) Run() error {
	return h.engine.Run(h.addr)
}

func (h *HttpPool) Set(peers ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, peer := range peers {
		h.peers.Add(peer)
		h.httpGetters[peer] = &HttpGetter{baseURL: peer}
	}
}

func (h *HttpPool) PickPeer(key string) (PeerGetter, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.peers.Get(key); peer != "" && peer != h.addr {
		return h.httpGetters[peer], true
	}
	return nil, false
}

func NewHttpPool(addr string) *HttpPool {
	h := &HttpPool{
		addr:        addr,
		basePath:    defaultBasePath,
		engine:      gee.New(),
		httpGetters: make(map[string]*HttpGetter),
		peers:       consistenthash.New(10, nil),
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

type HttpGetter struct {
	baseURL string
}

func (h *HttpGetter) Get(group string, key string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/%s", h.baseURL, group, key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
