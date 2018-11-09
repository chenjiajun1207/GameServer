package main

import (
	"github.com/chenjiajun1207/GameServer/Log"
	"github.com/chenjiajun1207/GameServer/Network/Protocol"
)

func main() {
	log15.Info("test")
	log15.Info("test", "test", protocol.PkgHeartbeat)
}
