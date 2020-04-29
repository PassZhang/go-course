package main

import "fmt"
import "math/rand"
import "time"

const MAX_GUESS	int = 5 

func main() {
	var answer int

	rand.Seed(time.Now().Unix()) // 使用当前时间谁知随机数字
	userStr := "" 					//初始化一个错误的值，让逻辑进入重新输入环节。也用来防止用户输入非数字
	isQuit := false					//是否退出游戏
	result := rand.Int() % 100 //生成0-100的随机数
LEBEL:
	for i := 1; i <= MAX_GUESS; i++ {
		fmt.Print("请输入你猜的数字：")
			if _, err := fmt.Scanln(&answer); err !=  nil {
				fmt.Println("请重新输入")
				continue
			}

			if answer > result {
				fmt.Printf("猜大了，还有%d次机会\n", MAX_GUESS-i)
			} else if answer < result {
				fmt.Printf("猜小了，还有%d次机会\n", MAX_GUESS-i)
			} else {
				fmt.Println("猜对了，太棒了")
				break
			}
	}
	//猜对了之后，询问用户是否继续游戏
	fmt.Print("[2]:再玩一局？yes/no: ")
	fmt.Scan(&userStr)
	if userStr != "y" && userStr != "Y" && userStr != "yes" { //当用户选择退出时，修改退出变量为True
		isQuit = true
	} 
	LEBEL
}