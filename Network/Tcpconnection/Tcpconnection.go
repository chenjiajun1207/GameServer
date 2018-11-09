package tcpconnection

import (
	"github.com/chenjiajun1207/GameServer/Log"
	"github.com/chenjiajun1207/GameServer/Network/Acceptor"
	"github.com/chenjiajun1207/GameServer/Network/Tools"
	"github.com/chenjiajun1207/GameServer/Network/Usedata"
	"net"
)

const (
	PkgHeartbeat = iota
	PKGDATA
)

type TcpConnection struct {
	conn     *TcpSocket
	acceptor *acceptor.Acceptor
}

func NewTcpConnection(conn net.Conn) *TcpConnection {
	NewConnection := &TcpConnection{
		conn: &TcpSocket{
			conn: conn,
		},
	}
	NewConnection.acceptor = acceptor.Newsocket(conn, NewConnection.processmessage)
	return NewConnection
}

func (this *TcpConnection) Start() {
	this.acceptor.Start()
}

func (this *TcpConnection) Close() {
	this.conn.Close()
	this.acceptor.Close()
}

func (this *TcpConnection) processmessage(buf []byte) {
	temp, _ := tools.Decode(buf)
	switch temp.Tag {
	case PkgHeartbeat:
		this.heartbeat()
	case PKGDATA:
		this.processdata(buf)
	default:
	}
}

func (this *TcpConnection) heartbeat() {
	pkg := usedata.NewUseData(PkgHeartbeat, []byte{})
	buf, err := pkg.Encode()
	if err != nil {
		log15.Error("pkg encode err", "err", err)
		return
	}
	this.conn.Write(buf)
}

func (this *TcpConnection) processdata(buf []byte) {
	this.conn.Write(buf)
}
