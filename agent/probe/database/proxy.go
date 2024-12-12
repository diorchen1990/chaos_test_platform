package database

import (
	"fmt"
	"log"
	"net"
)

type DBProxy struct {
	listener net.Listener
	target   string
	rules    []FaultRule
}

func (p *DBProxy) Start() error {
	// 启动代理服务
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))
	if err != nil {
		return err
	}
	p.listener = l
	
	go p.acceptConnections()
	return nil
}

func (p *DBProxy) acceptConnections() {
	for {
		conn, err := p.listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go p.handleConnection(conn)
	}
} 