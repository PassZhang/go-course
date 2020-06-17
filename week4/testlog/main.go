package main

import (
	"log"
	"os"
)

func main() {
	// 设置格式的
	// flags
	log.SetFlags(log.Flags() | log.Ldate | log.Lshortfile)
	// prefix
	log.SetPrefix("main:")

	log.Println("我是第一条println日志")
	// log.Fatalln("我是一个Fatal日志")
	// log.Panicln("我是一条Panicln日志")
	log.Println("我是第二条println日志")

	// debug , info , wanrning, error

	logger := log.New(os.Stdout, "logger:", log.Flags()|log.Lshortfile|log.Lmicroseconds|log.LstdFlags)
	logger2 := log.New(os.Stdout, "logger2:", log.Flags()|log.Lshortfile|log.LUTC)

	logger.Println("我是一个logger日志")
	logger2.Println("我是一个logger2日志")

	// logger.SetOutput()

	// 标准输入/输出 fmt.Scan fmt.Println
	// 标准输入on.Stdin,
	// 标准输出os.Stdout,
	// 标准错误os.Stderr

}
