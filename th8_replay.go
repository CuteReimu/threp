package threp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strconv"
	"strings"
)

func DecodeTh8Replay(fin io.Reader) (*TH8RepInfo, error) {
	dat := make([]byte, 4)
	n, err := fin.Read(dat)
	if err != nil {
		return nil, err
	}
	if n < 4 {
		return nil, errors.New("not a replay")
	}
	if string(dat) != "T8RP" {
		return nil, errors.New("not a th08 replay")
	}
	return decodeTh8Replay(fin)
}

func decodeTh8Replay(fin io.Reader) (*TH8RepInfo, error) {
	// read data size
	buf := make([]byte, 8)
	n, err := fin.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != 8 {
		return nil, errors.New("decompress failed")
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
	err = seek(fin, int64(binary.LittleEndian.Uint32(buf)-4-8-4))
	if err != nil {
		return nil, err
	}
	// cut USER????????
	buf = make([]byte, 12)
	_, err = fin.Read(buf)
	if err != nil {
		return nil, err
	}
	buf, err = io.ReadAll(fin)
	if err != nil {
		return nil, err
	}
	ret, ok, err := parseJP(buf)
	if ok {
		return ret, err
	}
	ret, ok, err = parseCN(buf)
	if ok {
		return ret, err
	}
	return nil, errors.New("parse failed")
}

func parseJP(buf []byte) (*TH8RepInfo, bool, error) {
	m, err := linesToMap(transform.NewReader(bytes.NewReader(buf), japanese.ShiftJIS.NewDecoder()), "\t")
	if err != nil {
		return nil, false, err
	}
	if _, ok := m["プレイヤー名"]; !ok {
		return nil, false, nil
	}
	ret := &TH8RepInfo{}
	ret.game = "8"
	// line1: プレイヤー名\t(name)\r\n
	ret.Player = m["プレイヤー名"]
	// line2: プレイ時刻\t(yyyy/mm/dd hh24:mi:ss)\r\n
	ret.Date = m["プレイ時刻"]
	// line3: キャラ名\t(char)\r\n
	ret.Char = m["キャラ名"]
	// line4: スコア\t\t(score)\r\n
	score, ok := m["スコア"]
	if ok {
		ret.Score, err = strconv.ParseInt(score, 10, 64)
		if err != nil {
			return nil, true, err
		}
	}
	// line5: 難易度\t\t(level)\r\n
	ret.Rank = m["難易度"]
	// line6: 最終ステージ\t(stage)\r\n
	ret.Stage = m["最終ステージ"]
	// line7: ミス回数\t(miss)\r\n
	miss, ok := m["ミス回数"]
	if ok {
		ret.Miss, err = strconv.Atoi(miss)
		if err != nil {
			return nil, true, err
		}
	}
	// line8: ボム回数\t(bomb)\r\n
	bomb, ok := m["ボム回数"]
	if ok {
		ret.Bomb, err = strconv.Atoi(bomb)
		if err != nil {
			return nil, true, err
		}
	}
	// line9: 処理落ち率\t(slow rate)\r\n
	drop, ok := m["処理落ち率"]
	if ok {
		_, err = fmt.Sscanf(drop, "%f%%", &ret.Drop)
		if err != nil {
			return nil, true, err
		}
	}
	// line10: 人間率\t\t(human rate)\r\n
	human, ok := m["人間率"]
	if ok {
		_, err = fmt.Sscanf(strings.ReplaceAll(human, "％", "%"), "%f%%", &ret.Human)
		if err != nil {
			return nil, true, err
		}
	}
	// line11: ゲームのバージョン\t(version)\r\n
	ret.Version = m["ゲームのバージョン"]
	return ret, true, nil
}

func parseCN(buf []byte) (*TH8RepInfo, bool, error) {
	m, err := linesToMap(transform.NewReader(bytes.NewBuffer(buf), simplifiedchinese.GBK.NewDecoder()), "\t")
	if err != nil {
		return nil, false, err
	}
	if _, ok := m["玩家名"]; !ok {
		return nil, false, nil
	}
	ret := &TH8RepInfo{}
	ret.game = "8"
	// line1: プレイヤー名\t(name)\r\n
	ret.Player = m["玩家名"]
	// line2: プレイ時刻\t(yyyy/mm/dd hh24:mi:ss)\r\n
	ret.Date = m["游戏时间"]
	// line3: キャラ名\t(char)\r\n
	ret.Char = m["角色名"]
	// line4: スコア\t\t(score)\r\n
	score, ok := m["分数"]
	if ok {
		ret.Score, err = strconv.ParseInt(score, 10, 64)
		if err != nil {
			return nil, true, err
		}
	}
	// line5: 難易度\t\t(level)\r\n
	ret.Rank = m["难易度"]
	// line6: 最終ステージ\t(stage)\r\n
	ret.Stage = m["最终面"]
	// line7: ミス回数\t(miss)\r\n
	miss, ok := m["miss回数"]
	if ok {
		ret.Miss, err = strconv.Atoi(miss)
		if err != nil {
			return nil, true, err
		}
	}
	// line8: ボム回数\t(bomb)\r\n
	bomb, ok := m["Bomb回数"]
	if ok {
		ret.Bomb, err = strconv.Atoi(bomb)
		if err != nil {
			return nil, true, err
		}
	}
	// line9: 処理落ち率\t(slow rate)\r\n
	drop, ok := m["処理落率"]
	if ok {
		_, err = fmt.Sscanf(drop, "%f%%", &ret.Drop)
		if err != nil {
			return nil, true, err
		}
	}
	// line10: 人間率\t\t(human rate)\r\n
	human, ok := m["人间率"]
	if ok {
		_, err = fmt.Sscanf(strings.ReplaceAll(human, "％", "%"), "%f%%", &ret.Human)
		if err != nil {
			return nil, true, err
		}
	}
	// line11: ゲームのバージョン\t(version)\r\n
	ret.Version = m["游戏版本"]
	return ret, true, nil
}
