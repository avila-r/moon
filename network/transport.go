package network

import (
	"log"
	"net"
)

type (
	Transport interface {
		Connect(Transport) error
		Send(net.Addr, []byte) error
		Broadcast([]byte) error
		Address() net.Addr
	}

	TCPTransport struct {
		Listener      net.Listener
		TargetAddress string
		Peers         chan *Peer
	}
)

func (t *TCPTransport) Start() error {
	listener, err := net.Listen("tcp", t.TargetAddress)
	if err != nil {
		return err
	}

	t.Listener = listener

	go func() {
		for {
			conn, err := t.Listener.Accept()
			if err != nil {
				log.Printf("received error from %+v\n", conn)
				continue
			}

			t.Peers <- &Peer{
				Connection: conn,
			}
		}
	}()

	return nil
}
