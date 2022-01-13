package threp

import (
	"bufio"
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
)

func DecodeTh10Replay(fin io.ReadSeeker) (*NewRepInfo, error) {
	buf := make([]byte, 4)
	n, err := fin.Read(buf)
	if err != nil {
		return nil, err
	}
	// replay format check
	if n != 4 || string(buf) != "t10r" {
		return nil, errors.New("not a th10 replay file")
	}

	// read data size
	buf = make([]byte, 8)
	_, err = fin.Read(buf)
	if err != nil {
		return nil, err
	}
	buf = buf[:4]
	n, err = fin.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, errors.New("decompress failed")
	}

	// move to fileinfo block.
	_, err = fin.Seek(int64(binary.LittleEndian.Uint32(buf)-4), io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(fin)
	ret := &NewRepInfo{}
	// retrieve replay info
	// line1: USER????????東方風神録 リプレイファイル情報\r\n
	_, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	// line2: Version (version)\r\n
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Version = getValue("Version ", string(line))

	// line3: Name (name)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Player = getValue("Name ", string(line))

	// line4: Date yy/mm/dd hh24:mi
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Date = "20" + getValue("Date ", string(line))

	// char, rank and stage are only available in 1.00a or higher.
	if strings.Compare(ret.Version, "1.00a") >= 0 {
		// line5: Chara (char)\r\n
		line, _, err = reader.ReadLine()
		if err != nil {
			return nil, err
		}
		ret.Char = getValue("Chara ", string(line))

		// line6: ;Rank (rank)\r\n
		line, _, err = reader.ReadLine()
		if err != nil {
			return nil, err
		}
		ret.Rank = getValue("Rank ", string(line))

		// line7: Extra|Stage (stage)
		line, _, err = reader.ReadLine()
		if err != nil {
			return nil, err
		}
		ret.Stage = getValue("Stage ", string(line))
		if len(ret.Stage) == 0 {
			ret.Stage = getValue("Extra ", string(line))
		}
	}

	// line8: Score (score)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Score, err = strconv.ParseInt(getValue("Score ", string(line)), 10, 64)
	if err != nil {
		return nil, err
	}
	ret.Score *= 10

	// line9: Slow Rate (rate)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	drop, err := strconv.ParseFloat(getValue("Slow Rate ", string(line)), 32)
	if err != nil {
		return nil, err
	}
	ret.Drop = float32(drop)

	// line10 and after:
	// \0USER????????(comment...)\0
	// read to eof. limit 1024bytes.
	buf = make([]byte, 1024)
	n, _ = fin.Read(buf)
	if n > 0 {
		s := trim(string(buf))
		if s[:4] == "USER" && len(s) >= 12 {
			// cut USER????????
			ret.Comment = trim(s[12:])
		}
	}
	return ret, nil
}
