// +build ignore
package tools_test

import (
	"bytes"
	"encoding/binary"
	. "gopkg.in/check.v1"
	"lib/Crazy/Network/Tools"
	"testing"
)

type PostTestSuite struct{}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) TestRunString(c *C) {
	var buf []byte
	temp := bytes.NewBuffer(buf)
	binary.Write(temp, binary.LittleEndian, byte(1))
	binary.Write(temp, binary.LittleEndian, []byte{0, 0, 0})
	binary.Write(temp, binary.LittleEndian, int32(9))
	binary.Write(temp, binary.LittleEndian, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	pkg, _ := tools.Decode(temp.Bytes())
	c.Log(pkg)
	// buf := make([]byte, 5)
	// for k, _ := range buf {
	// buf[k] = byte(k)
	// }
	// gg := bytes.NewBuffer(buf)
	// gg.Next(2)
	// tt, _ := gg.ReadByte()
	// tt, _ = gg.ReadByte()
	// tt, _ = gg.ReadByte()

	// c.Log(tt)
}
