package main

import (
	"flag"

	"github.com/liushuchun/wechatcmd/wechat"
)

func main() {
	flag.Parse()
	maxChanSize := 10000

	//log.SetLevel(log.DebugLevel)
	msgIn := make(chan wechat.Message, maxChanSize)
	textOut := make(chan string, maxChanSize)
	initList := []string{"test1", "test2", "test4"}
	userList := initList
	chatIn := make(chan wechat.Message, maxChanSize)

	layout := NewLayout(initList, userList, chatIn, msgIn, textOut)
	layout.Init()
}
