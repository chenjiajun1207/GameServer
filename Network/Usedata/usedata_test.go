// +build ignore
package usedata_test

import (
	"github.com/chenjiajun1207/GameServer/Network/Usedata"
	. "gopkg.in/check.v1"
	"testing"
)

type PostTestSuite struct{}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) TestRun111(c *C) {
	temp := make([]byte, 10)
	for index, _ := range temp {
		temp[index] = byte(index)
	}
	c.Log(temp)
	codec := usedata.NewCodecChat(1, temp)
	c.Log(codec)
	buf, err := codec.Encode()
	c.Assert(err, IsNil)
	c.Log(buf)
}
