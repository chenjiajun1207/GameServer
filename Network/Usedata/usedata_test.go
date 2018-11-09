// +build ignore
package usedata_test

import (
	. "gopkg.in/check.v1"
	"lib/Crazy/Network/Usedata"
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
