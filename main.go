package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log_collection/util/cron"
	"time"
)

func main() {
	cron.UploadToGit()

	filename := "~/demo/go/my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tails, filename)
			time.Sleep(100 * time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
