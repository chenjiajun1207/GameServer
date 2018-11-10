//coder with chat server

package tcpserver

import (
	"bufio"
	"flag"
	"github.com/chenjiajun1207/GameServer/Log"
	"github.com/chenjiajun1207/GameServer/Network/Acceptor"
	"github.com/chenjiajun1207/GameServer/Network/Tcpconnection"
	"github.com/chenjiajun1207/GameServer/Network/Tools"
	"github.com/chenjiajun1207/GameServer/Network/Usedata"
	"net"
)

type TcpServer struct {
	connecturl string
	netListen  net.Listener //why not use defer net.Listener.Close()
	tcpconn    *tcpconnection.TcpConnection
}

func (this *TcpServer) Init() {
	var err error
	this.netListen, err = net.Listen("tcp", this.connecturl)
	if err != nil {
		log15.Error("Listen err", "err", err)
		return
	}
	for {
		if this.acceptConn(this.netListen) {
			return
		}
	}
}

func (this *TcpServer) acceptConn(listener net.Listener) bool {
	conn, err := listener.Accept()
	if err != nil {
		log15.Error("Accept err", "err", err)
		return true
	}
	tcpconn := NewTcpConnection(conn)
	tcpconn.Start()
	return false
}
