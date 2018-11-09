package tools

import (
	"bytes"
	"encoding/binary"
	"lib/Crazy/Log"
	"lib/Crazy/Network/Protocol"
	"lib/Crazy/Network/Usedata"
)

func Decode(buf []byte) (*usedata.Usedata, error) {
	log15.Debug("debug buf", "buf", buf)
	temp := bytes.NewBuffer(buf)
	Tag, err := temp.ReadByte()
	if err != nil {
		log15.Error("read err", "err", err)
		return nil, err
	}
	log15.Debug("debug buf", "Tag", Tag)
	temp.Next(3)
	var length int32
	err = binary.Read(temp, binary.LittleEndian, &length)
	if err != nil {
		log15.Error("read err", "err", err)
		return nil, err
	}
	log15.Debug("debug buf", "length", length)
	usedata := usedata.Usedata{
		Tag:     protocol.PackageType(Tag),
		Datalen: int(length),
		Data:    make([]byte, length),
	}
	err = binary.Read(temp, binary.LittleEndian, usedata.Data)
	if err != nil {
		log15.Error("read err", "err", err)
		return nil, err
	}
	log15.Debug("debug buf", "usedata", usedata)
	return &usedata, nil
}
