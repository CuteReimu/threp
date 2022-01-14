package test

import (
	"bytes"
	"github.com/CuteReimu/threp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTh6(t *testing.T) {
	fin := bytes.NewReader([]byte("\x54\x36\x52\x50\x02\x01\x00\x03\xaa\x45\x34\x3f\xb5\xca\x6f\x9f\xa6\xae\xb3\xbb\xc4\xc8\xd2\xd9\xae\xfd\x05\x10\x0b\x1b\x21\x2d\x26\xed\xc1\x5c\x96\xab\xa7\x20\x07\xcc\xf0\x72\x3a\x70\x1d\x8d\xc5\xc5\x94"))
	ret, err := threp.DecodeTh6Replay(fin)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ret, &threp.OldRepInfo{
		Game:   "6",
		Date:   "2022/01/02",
		Player: "HIMAJIN@",
		Char:   "ReimuA",
		Score:  160932500,
		Rank:   "Lunatic",
		Drop:   0.4163742,
	})
}
