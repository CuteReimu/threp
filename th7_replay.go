package threp

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"math"
)

func DecodeTh7Replay(fin io.Reader) (*OldRepInfo, error) {
	dat := make([]byte, 4)
	n, err := fin.Read(dat)
	if err != nil {
		return nil, err
	}
	if n < 4 {
		return nil, errors.New("not a replay")
	}
	if string(dat) != "T7RP" {
		return nil, errors.New("not a th07 replay")
	}
	return decodeTh7Replay(fin)
}

func decodeTh7Replay(fin io.Reader) (*OldRepInfo, error) {
	dat, err := io.ReadAll(fin)
	if err != nil {
		return nil, err
	}
	dat = append([]byte("T7RP"), dat...)
	dat2 := make([]byte, len(dat))
	mask := dat[0x0d]
	for i := 0; i < 0x10; i++ {
		dat2[i] = dat[i]
	}
	for i := 0x10; i < len(dat); i++ {
		dat2[i] = (dat[i] - mask) & 0xff
		mask = (mask + 0x07) & 0xff
	}

	dat = dat2
	// decompress
	v04 := uint32(0)
	v1c := uint32(0)
	v30 := uint32(0)
	v28 := uint32(0)
	v34 := uint32(1)
	v11 := uint32(0x80)
	v20 := uint32(0)

	for i := 0; i < 4; i++ {
		v20 = v20*0x100 + uint32(dat[0x17-i])
	}
	v4b := make([]byte, 0x16c80)

	repLength := binary.LittleEndian.Uint32(dat[0x18:])
	dat2 = make([]byte, repLength+0x54)
	for i := 0; i < 0x54; i++ {
		dat2[i] = dat[i]
	}
	index := uint32(0x54)
	i := uint32(0x54)
	for index < repLength {
		flStopDoLoop := 0
		for index < repLength {
			flFirstRun := 1
			tmpFirst := true
			for v30 != 0 || tmpFirst {
				tmpFirst = false
				if v11 == 0x80 {
					v04 = uint32(dat[i])
					if i-0x54 < v20 {
						i += 1
					} else {
						v04 = 0
					}
					v28 += v04
				}
				if flFirstRun == 1 {
					v1c = v04 & v11
					v11 = v11 >> 1
					if v11 == 0 {
						v11 = 0x80
					}
					if v1c == 0 {
						flStopDoLoop = 1
						break
					}
					v30 = 0x80
					v1c = 0
					flFirstRun = 0
				} else {
					if (v11 & v04) != 0 {
						v1c = v1c | v30
					}
					v30 = v30 >> 1
					v11 = v11 >> 1
					if v11 == 0 {
						v11 = 0x80
					}
				}
			}
			if flStopDoLoop == 1 {
				break
			}
			dat2[index] = byte(v1c)
			index += 1
			v4b[v34] = byte(v1c & 0xff)
			v34 = (v34 + 1) & 0x1fff
		}
		if index > repLength {
			break
		}
		v30 = 0x1000
		v1c = 0
		for v30 != 0 {
			if v11 == 0x80 {
				v04 = uint32(dat[i])
				if i-0x54 < v20 {
					i += 1
				} else {
					v04 = 0
				}
				v28 += v04
			}
			if (v11 & v04) != 0 {
				v1c = v1c | v30
			}
			v30 = v30 >> 1
			v11 = v11 >> 1
			if v11 == 0 {
				v11 = 0x80
			}
		}
		v0c := v1c
		if v0c == 0 {
			break
		}
		v30 = 8
		v1c = 0
		for v30 != 0 {
			if v11 == 0x80 {
				v04 = uint32(dat[i])
				if i-0x54 < v20 {
					i += 1
				} else {
					v04 = 0
				}
				v28 += v04
			}
			if (v11 & v04) != 0 {
				v1c = v1c | v30
			}
			v30 = v30 >> 1
			v11 = v11 >> 1
			if v11 == 0 {
				v11 = 0x80
			}
		}
		v24 := v1c + 2
		v10 := uint32(0)
		for v10 <= v24 && index < repLength {
			v2c := v4b[(v0c+v10)&0x1fff]
			dat2[index] = v2c
			index += 1
			v4b[v34] = v2c & 0xff
			v34 = (v34 + 1) & 0x1fff
			v10 += 1
		}
	}

	// replay info
	date := string(dat2[0x58 : 0x58+5])
	name := string(dat2[0x5e : 0x5e+8])
	char := dat2[0x56]
	rank := dat2[0x57]
	myver := dat2[0xe0 : 0xe0+5] // '0100a'
	myver = append(myver[0:2:2], append([]byte{'.'}, myver[2:]...)...)
	if myver[0] == '0' {
		myver = myver[1:]
	}
	score := int64(binary.LittleEndian.Uint32(dat2[0x6c:0x6c+4])) * 10
	drop := math.Float32frombits(binary.LittleEndian.Uint32(dat2[0xcc : 0xcc+4]))
	return &OldRepInfo{
		game:    "7",
		Date:    date,
		Player:  trimNull(name),
		Char:    safeIndex([]string{"ReimuA", "ReimuB", "MarisaA", "MarisaB", "SakuyaA", "SakuyaB"}, char),
		Score:   score,
		Rank:    safeIndex([]string{"Easy", "Normal", "Hard", "Lunatic", "Extra", "Phantasm"}, rank),
		Version: string(myver),
		Drop:    drop,
	}, nil
}
