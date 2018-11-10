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
	temp, err := tools.Decode(buf)
	if err != nil {
		this.Close()
	}
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
		this.Close()
		log15.Error("pkg encode err", "err", err)
		return
	}
	_, err = this.conn.Write(buf) // why we don't care about the number
	if err != nil {
		this.Close()
		log15.Error("Write", "err", err)
	}
}

func (this *TcpConnection) processdata(buf []byte) {
	_, err := this.conn.Write(buf)
	if err != nil {
		this.Close()
	}
}
