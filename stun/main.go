package main

import (
	"fmt"
	"net"

	"github.com/pion/stun"
	"github.com/pion/transport/v2/stdnet"
)

type Net struct {
	Port int
	*stdnet.Net
}

func newNet(port int) (*Net, error) {
	m, e := stdnet.NewNet()
	if e != nil {
		return nil, e
	}
	return &Net{Port: port, Net: m}, nil
}

func (n *Net) Dial(network, address string) (net.Conn, error) {
	dialer := net.Dialer{
		LocalAddr: &net.TCPAddr{
			Port: n.Port,
		},
	}
	return dialer.Dial(network, address)
}

func main() {
	// Parse a STUN URI
	u, err := stun.ParseURI("turn:stun.nextcloud.com:3478?transport=tcp")
	if err != nil {
		panic(err)
	}
	n, err := newNet(4635)
	if err != nil {
		panic(err)
	}
	c, err := stun.DialURI(u, &stun.DialConfig{
		Net: n,
	})
	if err != nil {
		panic(err)
	}
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	fn(c, message)
}

func fn(c *stun.Client, message *stun.Message) {
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		var xorAddr stun.XORMappedAddress

		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		fmt.Println("your IP is", xorAddr.IP, xorAddr.Port)
	}); err != nil {
		panic(err)
	}
}
