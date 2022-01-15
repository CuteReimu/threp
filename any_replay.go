package threp

import (
	"github.com/pkg/errors"
	"io"
	"log"
)

func DecodeReplay(fin io.Reader) (RepInfo, error) {
	buf := make([]byte, 4)
	n, err := fin.Read(buf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if n != 4 {
		return nil, errors.New("not a replay")
	}
	switch string(buf) {
	case "T6RP":
		return decodeTh6Replay(fin)
	case "T7RP":
		return decodeTh7Replay(fin)
	case "T8RP":
		return decodeTh8Replay(fin)
	default:
		game := getNewReplayGame(string(buf))
		if len(game) == 0 {
			return nil, errors.New("not a replay")
		}
		return decodeNewReplay(fin, game)
	}
}
