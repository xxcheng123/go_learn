package geecache

// PeerPicker 根据key选择哪个节点
type PeerPicker interface {
	PickPeer(key string) (PeerGetter, ok bool)
}

// PeerGetter 获取数据
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
