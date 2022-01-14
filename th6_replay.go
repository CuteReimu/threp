package threp

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"math"
)

func DecodeTh6Replay(fin io.Reader) (*OldRepInfo, error) {
	dat := make([]byte, 0x30)
	n, err := fin.Read(dat)
	if err != nil {
		return nil, err
	}
	if n < 0x30 {
		return nil, errors.New("not a replay")
	}
	if string(dat[:4]) != "T6RP" {
		return nil, errors.New("not a th06 replay")
	}
	// Decryption
	dat2 := make([]byte, 0, len(dat))
	mask := dat[0x0e]
	for i := 0; i < 0x0f; i++ {
		dat2 = append(dat2, dat[i])
	}
	for i := 0x0f; i < len(dat); i++ {
		dat2 = append(dat2, byte(uint16(dat[i])+0x100-uint16(mask)))
		mask = byte((uint16(mask) + 0x07) & 0xff)
	}

	// check
	if dat2[0x06] > 0x04 || dat2[0x07] > 0x05 {
		return nil, errors.New("decrypt th6 replay failed")
	}

	// replay info
	date := string(dat2[0x10 : 0x10+8])
	name := string(dat2[0x19 : 0x19+8])
	char := dat2[0x06]
	rank := dat2[0x07]
	score := binary.LittleEndian.Uint32(dat2[0x24 : 0x24+4])
	drop := math.Float32frombits(binary.LittleEndian.Uint32(dat2[0x2c : 0x2c+4]))
	return &OldRepInfo{
		Game:    "6",
		Date:    "20" + date[6:8] + "/" + date[:2] + "/" + date[3:5],
		Player:  trimNull(name),
		Char:    safeIndex([]string{"ReimuA", "ReimuB", "MarisaA", "MarisaB"}, char),
		Score:   int64(score),
		Rank:    safeIndex([]string{"Easy", "Normal", "Hard", "Lunatic", "Extra"}, rank),
		Version: "",
		Drop:    drop,
	}, nil
}
