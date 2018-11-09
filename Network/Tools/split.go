package tools

import (
	"bytes"
	"encoding/binary"
	"github.com/chenjiajun1207/GameServer/Log"
)

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF {
		if len(data) > 8 {
			var lenData int32
			err := binary.Read(bytes.NewBuffer(data[4:8]), binary.LittleEndian, &lenData)

			if err != nil {
				log15.Error(err.Error())
				return 0, nil, err
			}
			if int(lenData)+8 <= len(data) {
				log15.Debug("data", "data=", data) //Only work for Lvl equals 4
				return int(lenData) + 8, data[0 : int(lenData)+8], nil
			}
		}
	}
	return
}
