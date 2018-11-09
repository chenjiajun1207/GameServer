package acceptor

import (
	"bufio"
	"github.com/chenjiajun1207/GameServer/Log"
	"github.com/chenjiajun1207/GameServer/Network/Tools"
	"net"
)

type TransportState int

const (
	closed TransportState = iota
)

type Acceptor struct {
	conn             net.Conn
	messageProcesser func([]byte)
	done             chan struct{}
	transportState   TransportState
}

func Newsocket(conn net.Conn, messageprocess func([]byte)) *Acceptor {
	return &Acceptor{
		conn:             conn,
		messageProcesser: messageprocess,
	}
}

func (this *Acceptor) Start() {
	go this.doWork()
}

func (this *Acceptor) doWork() {
	results := make(chan []byte, 600)
	processdata := func(done <-chan struct{}) <-chan []byte {
		go func() {
			defer close(results)
			for {
				select {
				case result := <-results:
					this.messageProcesser(result)
				case <-done:
					return
				case <-this.done:
					return
				}
			}
		}()
		return results
	}
	terminated := make(chan struct{})
	processdata(terminated)

	scanner := bufio.NewScanner(this.conn)

	scanner.Split(tools.Split)

	for scanner.Scan() {
		results <- scanner.Bytes()
		log15.Debug("scanner Bytes", "scanner", scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		terminated <- struct{}{}
		log15.Error(err.Error())
		return
	}
}

func (this *Acceptor) Close() {
	if this.transportState != closed {
		this.transportState = closed
		this.done <- struct{}{}
	}
}
