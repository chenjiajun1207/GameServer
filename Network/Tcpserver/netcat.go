// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"bufio"
	"lib/Crazy/Log"
	"lib/Crazy/Network/Tools"
	"lib/Crazy/Network/Usedata"
	"net"
	"os"
	"strings"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log15.Error(err.Error())
		return
	}
	go func() {
		scanner := bufio.NewScanner(conn)

		scanner.Split(tools.Split)

		for scanner.Scan() {
			pkg, _ := tools.Decode(scanner.Bytes())
			log15.Info(string(pkg.Data))
		}

		if err := scanner.Err(); err != nil {
			log15.Error(string(scanner.Bytes()))
			return
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	for {

		result, err := reader.ReadString('\n')
		if err != nil {
			log15.Error(err.Error())
			return
		}
		result = strings.Replace(result, "\n", "", -1)
		log15.Debug(string(result))
		usedata := usedata.NewCodecChat(1, []byte(result))
		temp, err := usedata.Encode()
		log15.Debug("safdsafd", "temp", temp)
		if err != nil {
			log15.Error(err.Error())
			return
		}
		conn.Write(temp)
	}
	conn.Close()
}
