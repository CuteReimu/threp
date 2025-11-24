package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/CuteReimu/threp"
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
	ret, err := threp.DecodeReplay(f)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ret.String())
}
