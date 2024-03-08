package test

import (
	"embed"
	"github.com/CuteReimu/threp"
	"testing"
)

//go:embed *.rpy
var fs embed.FS

func TestTh6(t *testing.T) {
	fin, err := fs.Open("th6_01.rpy")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TH6 Lunatic ReimuA\n机签：HIMAJIN@\n分数：1.61亿\n处理落率：0.42%"
	actual := ret.String()
	if expected != actual {
		t.Error(expected)
		t.Error(actual)
	}
}

func TestTh7(t *testing.T) {
	fin, err := fs.Open("th7_02.rpy")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TH7 Lunatic SakuyaA\n机签：HDZ LNB\n分数：7.12亿\n处理落率：0.15%"
	actual := ret.String()
	if expected != actual {
		t.Error(expected)
		t.Error(actual)
	}
}

func TestTh8(t *testing.T) {
	fin, err := fs.Open("th8_01.rpy")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TH8 Lunatic Stage 6-Kaguya 妖夢＆幽々子\n机签：David Lu\n10 Miss 35 Bomb\n分数：13.05亿\n处理落率：0.00%"
	actual := ret.String()
	if expected != actual {
		t.Error(expected)
		t.Error(actual)
	}
}

func TestTh8CN(t *testing.T) {
	fin, err := fs.Open("th8_02.rpy")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TH8 Lunatic Stage 6-Kaguya 妖夢＆幽々子\n机签：David Lu\n10 Miss 35 Bomb\n分数：13.05亿\n处理落率：0.00%"
	actual := ret.String()
	if expected != actual {
		t.Error(expected)
		t.Error(actual)
	}
}

func TestTh18(t *testing.T) {
	fin, err := fs.Open("th18_01.rpy")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := threp.DecodeReplay(fin)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TH18 Lunatic All Clear Sanae\n机签：David Lu\n分数：10.63亿\n处理落率：0.10%"
	actual := ret.String()
	if expected != actual {
		t.Error(expected)
		t.Error(actual)
	}
}
