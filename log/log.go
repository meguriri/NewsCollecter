package log

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	File string
)

func InitLog() {
	//set logfile Stdout
	logFile, logErr := os.OpenFile(File, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
	}
	log.SetOutput(logFile)
	// 默认会有log.Ldate | log.Ltime（日期 时间），这里重写为 日 时 文件名
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //2015/04/22 11:28:41 test.go:29: content

}

func GetLog() []string {
	logger := []string{}
	f, err := os.Open("./log.log")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	r := bufio.NewReader(f)
	for {
		//遇到'\n'结束读取
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		logger = append(logger, string(buf))
	}
	return logger
}
