package threp

type OldRepInfo struct {
	Date    string
	Player  string
	Char    string
	Score   int64
	Rank    string
	Version string
	Drop    float32 // 处理落率（百分比）
}

type TH8RepInfo struct {
	Player   string
	PlayTime string
	Char     string
	Score    int64
	Level    string
	Stage    string
	Miss     int
	Bomb     int
	Drop     float32 // 处理落率（百分比）
	Human    float32 // 人妖率（百分比）
	Version  string
}
