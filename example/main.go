package main

import (
	"encoding/json"
	"fmt"
	"github.com/CuteReimu/threp"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	run()
	fmt.Println("Press enter to continue...")
	var s string
	_, _ = fmt.Scanln(&s)
}

func run() {
	log.SetFlags(0)
	if len(os.Args) != 2 || len(strings.TrimSpace(os.Args[1])) == 0 {
		fmt.Println("Args error. Just drag a rpy file onto this program!")
		return
	}
	fileName := strings.TrimSpace(os.Args[1])
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() { _ = f.Close() }()
	buf := make([]byte, 4)
	n, err := f.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	if n != 4 {
		log.Println("not a replay")
		return
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
		return
	}
	var ret interface{}
	switch string(buf) {
	case "T6RP":
		ret, err = threp.DecodeTh6Replay(f)
	case "T7RP":
		ret, err = threp.DecodeTh7Replay(f)
	case "T8RP":
		ret, err = threp.DecodeTh8Replay(f)
	default:
		ret, err = threp.DecodeNewReplay(f)
	}
	if err != nil {
		log.Println(err)
		return
	}
	buf, err = json.MarshalIndent(ret, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buf))
}
