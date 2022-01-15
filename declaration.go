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
	return fmt.Sprintf("TH%s %s %s\n%s %s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Char, strings.TrimSpace(o.Player), o.Date, formatScore(o.Score), o.Drop)
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
	return fmt.Sprintf("TH%s %s %s %s\n%s %s\n%s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Stage, o.Char, strings.TrimSpace(o.Player), o.Date, missBomb, formatScore(o.Score), o.Drop)
}

type NewRepInfo struct {
	OldRepInfo
	Stage string
}

func (o *NewRepInfo) String() string {
	return fmt.Sprintf("TH%s %s %s %s\n%s %s\n分数：%s\n处理落率：%.2f%%", o.game, o.Rank, o.Stage, o.Char, strings.TrimSpace(o.Player), o.Date, formatScore(o.Score), o.Drop)
}

func formatScore(score int64) string {
	var s string
	if score >= 100000000 {
		s += strconv.FormatInt(score/100000000, 10) + "亿"
		score %= 100000000
	}
	if score >= 10000 {
		if len(s) > 0 {
			s += fmt.Sprintf("%04d万", score/10000)
		} else {
			s += strconv.FormatInt(score/10000, 10) + "万"
		}
		score %= 10000
	}
	if len(s) > 0 {
		s += fmt.Sprintf("%04d", score)
	} else {
		s += strconv.FormatInt(score, 10)
	}
	return s
}
