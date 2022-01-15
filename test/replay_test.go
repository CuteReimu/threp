package test

import (
	"bytes"
	"github.com/CuteReimu/threp"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"testing"
)

func TestTh6(t *testing.T) {
	fin := bytes.NewReader([]byte("T6RP\x02\x01\x00\x03\xaa\x45\x34\x3f\xb5\xca\x6f\x9f\xa6\xae\xb3\xbb\xc4\xc8\xd2\xd9\xae\xfd\x05\x10\x0b\x1b\x21\x2d\x26\xed\xc1\x5c\x96\xab\xa7\x20\x07\xcc\xf0\x72\x3a\x70\x1d\x8d\xc5\xc5\x94"))
	ret, err := threp.DecodeTh6Replay(fin)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, &threp.OldRepInfo{
		Game:   "6",
		Date:   "2022/01/02",
		Player: "HIMAJIN@",
		Char:   "ReimuA",
		Score:  160932500,
		Rank:   "Lunatic",
		Drop:   0.4163742,
	}, ret)
}

func TestTh8(t *testing.T) {
	fin := &bytes.Buffer{}
	_, err := transform.NewWriter(fin, japanese.ShiftJIS.NewEncoder()).Write([]byte("T8RP\x06\x00\x00\x01\x00\x00\x00\x00\x04\x00\x00\x00プレイヤー名\tDavid Lu\r\nプレイ時刻\t2021/01/14 02:58:23\r\nキャラ名\t妖夢＆幽々子　　\r\nスコア\t\t1304924700\r\n難易度\t\tLunatic\r\n最終ステージ\tStage 6-Kaguya\r\nミス回数\t10\r\nボム回数\t35\r\n処理落ち率\t0.000000%\r\n人間率\t\t49.05％\r\nゲームのバージョン\t1.00d\r\n\x00"))
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeTh8Replay(bytes.NewReader(fin.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, &threp.TH8RepInfo{
		NewRepInfo: threp.NewRepInfo{
			OldRepInfo: threp.OldRepInfo{
				Game:    "8",
				Date:    "2021/01/14 02:58:23",
				Player:  "David Lu",
				Char:    "妖夢＆幽々子",
				Score:   1304924700,
				Rank:    "Lunatic",
				Version: "1.00d",
				Drop:    0.00,
			},
			Stage: "Stage 6-Kaguya",
		},
		Human: 49.05,
		Miss:  10,
		Bomb:  35,
	}, ret)
}

const userLine = "\x00\x00\x00\x00\x93\x8c\x95\xfb\x93\xf8\x97\xb4\x93\xb4\x20\x83\x8a\x83\x76\x83\x8c\x83\x43\x83\x74\x83\x40\x83\x43\x83\x8b\x8f\xee\x95\xf1\r\n"

func TestTh18(t *testing.T) {
	fin := bytes.NewReader([]byte("t18r\x06\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00" + userLine + "Version 1.00\r\nName David Lu\r\nDate 21/08/14 00:22\r\nChara Sanae \r\nRank Lunatic\r\nStage All Clear\r\nScore 106260644\r\nSlow Rate 0.10\r\n\x00\x00USER \x00\x00\x00\x01\x00\x00\x00\x83R\x83\x81\x83\x93\x83g\x82\xF0\x8F\x91\x82\xAF\x82\xB7\x00\x00"))
	ret, err := threp.DecodeNewReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, &threp.NewRepInfo{
		OldRepInfo: threp.OldRepInfo{
			Game:    "18",
			Date:    "2021/08/14 00:22",
			Player:  "David Lu",
			Char:    "Sanae",
			Score:   1062606440,
			Rank:    "Lunatic",
			Version: "1.00",
			Drop:    0.1,
		},
		Stage: "All Clear",
	}, ret)
}
