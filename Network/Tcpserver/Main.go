//coder with chat server

package main

import (
	"bufio"
	"flag"
	"lib/Crazy/Log"
	"lib/Crazy/Network/Protocol"
	"lib/Crazy/Network/Tools"
	// "lib/Crazy/Network/Usedata"
	"net"
)

type Tcpserver struct {
}

func DoWork(conn net.Conn) {
	defer conn.Close()

	receive := func() <-chan []byte {
		results := make(chan []byte, 600)
		go func() {
			defer close(results)

			scanner := bufio.NewScanner(conn)

			scanner.Split(tools.Split)

			for scanner.Scan() {
				results <- scanner.Bytes()
				log15.Debug("scanner Bytes", "scanner", scanner.Bytes())
			}

			if err := scanner.Err(); err != nil {
				log15.Error(err.Error())
				return
			}
		}()
		return results
	}
	processdata := func(results <-chan []byte) {
		for result := range results {
			temp, _ := tools.Decode(result) //从包中返回usedata,缺无需在main中import
			switch temp.Tag {
			case protocol.PKG_HANDSHAKE:
			case protocol.PKG_HANDSHAKE_ACK:
			case protocol.PkgHeartbeat:

			case protocol.PKG_DATA:
				buf, _ := temp.Encode()
				conn.Write(buf)
			default:
			}

		}
	}

	results := receive()
	processdata(results)
}

func main() {
	loglevel := flag.Int("lvl", 3, "log level")
	flag.Parse()
	root := log15.Root()
	root.SetHandler(log15.CallerFileHandler(log15.LvlFilterHandler(log15.Lvl(*loglevel), log15.StdoutHandler)))
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log15.Error(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log15.Error(err.Error())
			continue
		}
		go DoWork(conn)
	}
}
