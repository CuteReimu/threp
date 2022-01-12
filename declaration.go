package threp

type OldRepInfo struct {
	Date    string
	Player  string
	Char    string
	Score   uint32
	Rank    string
	Version string
	Drop    float32 // 处理落率（百分比）
}
