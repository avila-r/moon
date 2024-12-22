package network

import (
	"log"
	"net"
)

type Server struct {
	ID           string
	TCPTransport *TCPTransport

	Peers     chan *Peer
	PeerChain map[net.Addr]*Peer

	RPC chan RPC

	TargetNodes []string

	Break chan any
}

func (s *Server) Start() {
	if err := s.TCPTransport.Start(); err != nil {
		log.Fatalf("unable to start server %v: %v", s.ID, err.Error())
	}

	debugger.Logf("server %v started", s.ID)

	for _, address := range s.TargetNodes {
		debugger.Logf("trying to connect to %v", address)

		go func(address string) {
			conn, err := net.Dial("tcp", address)
			if err != nil {
				debugger.Logf("could not connect to %+v\n", conn)
				return
			}

			s.Peers <- &Peer{
				Connection: conn,
			}
		}(address)
	}

	debugger.Logf("server %v accepting TCP connections on %v", s.ID, s.TCPTransport.TargetAddress)

server:
	for {
		select {
		case peer := <-s.Peers:
			s.PeerChain[peer.Connection.RemoteAddr()] = peer

			go peer.Listen(s.RPC)

		case <-s.Break:
			break server
		}
	}

	debugger.Logf("server %v is shutting down", s.ID)
}
