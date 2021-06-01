package main

import (
	"io"
	"log"
	"os"
)

//知识点一： 设置log的方法：修改log包中的前缀，设置输出路径及设置输出的信息
//func init() {
//	f, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		fmt.Println("打开文件失败")
//	}
//	log.SetPrefix("这是prefix前缀 ")
//	log.SetOutput(f)
//}
//func main() {
//	log.Fatalln("这是测试")
//}

//知识点二: 设置不同等级的输出
var Trace *log.Logger
var Info *log.Logger
var Warning *log.Logger
var Error *log.Logger

func init() {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("打开文件错误")
	}
	//io.Discard即抛弃所有，不写
	Trace = log.New(io.Discard, "Trace ", log.Ltime|log.LUTC)
	//os.Stdout标准输出 os.Stdin标准输出 os.Stderr标准错误
	Info = log.New(os.Stdout, "info ", log.Ltime|log.LUTC)
	Warning = log.New(os.Stdout, "Warning ", log.Ltime|log.LUTC)
	//你也可以用同时输出到终端和日志
	//Warning = log.New(io.MultiWriter(os.Stderr, file), "Warning ", log.Ltime|log.LUTC)
	//io.MultiWriter 将多个io.Write return成一个io.Write输出 即同时在终端和日志文件中显示
	Error = log.New(io.MultiWriter(os.Stderr, file), "Error ", log.Ltime|log.LUTC)
}
func main() {
	Trace.Println("鸡毛蒜皮的一些小事")
	Info.Println("一些特别的信息")
	Warning.Println("这是一个警告")
	Error.Println("出现了故障")
}
