package threp

type OldRepInfo struct {
	Game    string
	Date    string
	Player  string
	Char    string
	Score   int64
	Rank    string
	Version string
	Drop    float32 // 处理落率（百分比）
	Comment string
}

type TH8RepInfo struct {
	NewRepInfo
	Miss  int
	Bomb  int
	Human float32 // 人妖率（百分比）
}

type NewRepInfo struct {
	OldRepInfo
	Stage string
}
