package usedata

import (
	"bytes"
	"encoding/binary"
	"github.com/chenjiajun1207/GameServer/Network/Protocol"
)

const (
	HeadLen = 8
)

/*
T -----------在用标志位
U -----------预留
L -----------包长度

   T  |  U  |  U  |  U  |  L  |  L  |  L  |  L  |

*/

type Usedata struct {
	Tag     protocol.PackageType
	Datalen int
	Data    []byte
}

func NewUseData(tag protocol.PackageType, buf []byte) *Usedata {
	return &Usedata{
		Tag:     tag,
		Datalen: len(buf),
		Data:    buf,
	}
}

func (this *Usedata) Encode() ([]byte, error) {
	var pkg *bytes.Buffer = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, []byte{byte(this.Tag), 0, 0, 0})
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, int32(this.Datalen))
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, this.Data)
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
