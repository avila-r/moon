package network

import (
	"bytes"
	"net"
)

type Peer struct {
	Connection net.Conn
	Outgoing   bool
}

func (p *Peer) Listen(rpc chan RPC) {
	buffer := make([]byte, 4096)

	for {
		payload, err := p.Connection.Read(buffer)
		if err != nil {
			continue // TODO: Log
		}

		rpc <- RPC{
			From:    p.Connection.RemoteAddr(),
			Payload: bytes.NewReader(buffer[:payload]),
		}
	}
}
