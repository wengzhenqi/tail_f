package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log"
	"strings"
	//"go_code/cconfig"
	"github.com/go-basic/ipv4"
)
func Myip()string{
	ip := ipv4.LocalIP()
	return ip
}
func Tail_f()  {
	fileName := logfile
	tailfs,err := tail.TailFile(fileName,tail.Config{
		Location: &tail.SeekInfo{		//如果程序出现异常，保存上次读取的位置，避免重新读取新的管道
			Offset: 0,
			Whence: 2,
		},
		ReOpen:      true,			//文件被移除或被打包，需要重新打开
		MustExist:   false,			// 如果文件不存在，是否退出程序，false是不退出
		Poll:        false,
		Follow:      true,			// 实时跟踪
	})
	if err != nil {
		fmt.Println("tail failed error : ",err)
		return
	}
	var msg  *tail.Line
	var ok bool

	for true{
		msg,ok = <- tailfs.Lines			// ok是判断管道是否被关闭,如果关闭就是文件被重制了，需要重新读取新的管道
		if !ok{
			fmt.Println("tailf failed close reopen, filename: ",fileName)
			continue
		}
		//fmt.Println("读取到的内容",msg.Text)				// msg.text 里存的是读取到的当前内容
		log.Println("读取到的内容",msg.Text)
		//str := "value8"
		str := InitConfig("/root/key_config.txt")
		bb := strings.Contains(msg.Text,str)
		if bb {
			//DingTalk.DingToInfo("f"+"读取到了value8")
			DingToInfo("error: " + Myip() + "     " +msg.Text)
		}
	}
	fmt.Println("tail end")
}
