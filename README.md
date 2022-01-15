# threp

![](https://img.shields.io/github/languages/top/CuteReimu/threp "语言")
[![](https://img.shields.io/github/workflow/status/CuteReimu/threp/Go)](https://github.com/CuteReimu/threp/actions/workflows/go.yml "代码分析")
[![](https://img.shields.io/github/contributors/CuteReimu/threp)](https://github.com/CuteReimu/threp/graphs/contributors "贡献者")
[![](https://img.shields.io/badge/License-BY--NC--SA%203.0-lightgrey)](https://github.com/CuteReimu/threp/blob/master/LICENSE "许可协议")

解析东方Project的replay文件的Go语言版本

这个仓库是在score.royalflare.net即将关站前夕，为纪念它而写的。现将该网站上开源的python解析replay文件的代码转化为Go语言版本。

## 安装方法

```bash
go get github.com/CuteReimu/threp
```

## 使用方法

```go
package main

import (
	"fmt"
	"github.com/CuteReimu/threp"
	"os"
)

func main() {
	f, _ := os.Open(`th12_ud0001.rpy`)
	defer f.Close()
	result, _ := threp.DecodeNewReplay(f)
	fmt.Println(result.String())
}
```

## 函数说明

- `DecodeTh6Replay`、`DecodeTh7Replay`、`DecodeTh8Replay`分别用来解析th6、th7、th8三作的rpy文件。
- `DecodeNewReplay`用来解析th95及以后各作的rpy文件
- `DecodeReplay`用来解析任意一作的rpy文件，返回`RepInfo`接口

## 进度

- [x] 东方红魔乡
- [x] 东方妖妖梦
- [x] 东方永夜抄
- [x] 东方文花帖
- [x] 东方风神录
- [x] 东方地灵殿
- [x] 东方星莲船
- [x] 东方文花帖DS
- [x] 妖精大战争
- [x] 东方神灵庙
- [x] 东方辉针城
- [x] 东方绀珠传
- [x] 东方天空璋
- [x] 秘封噩梦日记
- [x] 东方鬼形兽
- [x] 东方虹龙洞
- [x] 黄昏酒场

## 许可协议

按照score.royalflare.net提供的许可协议，本仓库也按照BY-NC-SA协议进行开源。
