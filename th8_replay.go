package threp

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"strconv"
	"strings"
)

func DecodeTh8Replay(fin io.Reader) (*TH8RepInfo, error) {
	buf := make([]byte, 4)
	n, err := fin.Read(buf)
	if err != nil {
		return nil, err
	}
	// replay format check
	if n != 4 || string(buf) != "T8RP" {
		return nil, errors.New("not a th08 replay")
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
	err = seek(fin, int64(binary.LittleEndian.Uint32(buf)-4))
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(transform.NewReader(fin, japanese.ShiftJIS.NewDecoder()))
	ret := &TH8RepInfo{}
	ret.Game = "8"
	// line1: プレイヤー名\t(name)\r\n
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Player = getValue("プレイヤー名\t", string(line))

	// line2: プレイ時刻\t(yyyy/mm/dd hh24:mi:ss)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Date = getValue("プレイ時刻\t", string(line))

	// line3: キャラ名\t(char)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Char = getValue("キャラ名\t", string(line))

	// line4: スコア\t\t(score)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Score, err = strconv.ParseInt(getValue("スコア\t\t", string(line)), 10, 64)
	if err != nil {
		return nil, err
	}

	// line5: 難易度\t\t(level)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Rank = getValue("難易度\t\t", string(line))

	// line6: 最終ステージ\t(stage)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Stage = getValue("最終ステージ\t", string(line))

	// line7: ミス回数\t(miss)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Miss, err = strconv.Atoi(getValue("ミス回数\t", string(line)))
	if err != nil {
		return nil, err
	}

	// line8: ボム回数\t(bomb)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Bomb, err = strconv.Atoi(getValue("ボム回数\t", string(line)))
	if err != nil {
		return nil, err
	}

	// line9: 処理落ち率\t(slow rate)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(getValue("処理落ち率\t", string(line)), "%f%%", &ret.Drop)
	if err != nil {
		return nil, err
	}

	// line10: 人間率\t\t(human rate)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(strings.ReplaceAll(getValue("人間率\t\t", string(line)), "％", "%"), "%f%%", &ret.Drop)
	if err != nil {
		return nil, err
	}

	// line11: ゲームのバージョン\t(human rate)\r\n
	line, _, err = reader.ReadLine()
	if err != nil {
		return nil, err
	}
	ret.Version = getValue("ゲームのバージョン\t", string(line))

	return ret, nil
}
