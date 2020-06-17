package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 时间
	now := time.Now()
	fmt.Printf("%T, %#v\n", now, now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// unix 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// => 字符串 2020-04-25 15:50:xx?
	// Printf 格式化字符串 占位符
	// 2006 4位数字的年
	// 01 2位数字的月
	// 02 2位数字的天
	// 03 12进制的小时
	// 15 24进制的小时
	// 04 2位数字的分钟
	// 05 2位数字的秒

	fmt.Println(now.Format("2006-01-02 03:04:05"))

	fmt.Println(now.Format("2006年01月02日 03:04:05"))

	fmt.Println(now.Format("2006年01月02日 15:04:05"))

	fmt.Println(now.Format("15:04:05 2006年01月02日"))

	fmt.Println(now.Format("15:04:05 02/01/2006"))
}
