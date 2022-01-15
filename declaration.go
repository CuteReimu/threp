package threp

import (
	"fmt"
	"strconv"
	"strings"
)

type RepInfo interface {
	Game() string
	String() string
}

type OldRepInfo struct {
	game    string
	Date    string
	Player  string
	Char    string
	Score   int64
	Rank    string
	Version string
	Drop    float32 // 处理落率（百分比）
	Comment string
}

func (o *OldRepInfo) Game() string {
	return o.game
}

func (o *OldRepInfo) String() string {
	return fmt.Sprintf("TH%s %s %s\n机签：%s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Char, strings.TrimSpace(o.Player), formatScore(o.Score), o.Drop)
}

type TH8RepInfo struct {
	NewRepInfo
	Miss  int
	Bomb  int
	Human float32 // 人妖率（百分比）
}

func (o *TH8RepInfo) String() string {
	var missBomb string
	if o.Miss == 0 {
		missBomb += "No Miss "
	} else {
		missBomb += strconv.Itoa(o.Miss) + " Miss "
	}
	if o.Bomb == 0 {
		missBomb += "No Bomb"
	} else {
		missBomb += strconv.Itoa(o.Bomb) + " Bomb"
	}
	return fmt.Sprintf("TH%s %s %s %s\n机签：%s\n%s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Stage, o.Char, strings.TrimSpace(o.Player), missBomb, formatScore(o.Score), o.Drop)
}

type NewRepInfo struct {
	OldRepInfo
	Stage string
}

func (o *NewRepInfo) String() string {
	return fmt.Sprintf("TH%s %s %s %s\n机签：%s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Stage, o.Char, strings.TrimSpace(o.Player), formatScore(o.Score), o.Drop)
}

func formatScore(score int64) string {
	if score >= 100000000 {
		return fmt.Sprintf("%.2f亿", float64(score)/100000000)
	} else if score >= 10000 {
		return fmt.Sprintf("%.1f万", float64(score)/10000)
	}
	return strconv.FormatInt(score, 10)
}
