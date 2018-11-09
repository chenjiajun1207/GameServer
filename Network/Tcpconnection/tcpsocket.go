package tcpconnection

import (
	"net"
	"sync"
)

type TcpSocket struct {
	conn  net.Conn
	mutex sync.Mutex
}

func (this *TcpSocket) Close() {
	this.conn.Close()
}

func (this *TcpSocket) Write(b []byte) (n int, err error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.conn.Write(b)
}

func (this *TcpSocket) RemoteAddr() string {
	return this.conn.RemoteAddr().String()
}
