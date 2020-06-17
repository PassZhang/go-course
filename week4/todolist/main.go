package main

import (
	"fmt"
	"todolist/controllers/task"
	mtask "todolist/models/task" // 别名导入，用于多个包重名的情况

	"github.com/passzhang/strutil"
	/*
		1. 导入mod init name + 目录结构
		2. 包名与所在目录名称保持一致
		3. 调用时包用包名.Var(包名不是文件名)
	*/)

func main() {
	fmt.Println(mtask.Name)
	fmt.Println(task.Name)
	task.Call()
	fmt.Println(task.Version)
	task.GetVersion()
	fmt.Println(strutil.RandString())
}

// 变量首字母大小写问题：首字母大写的时候一般是包外可见，首字母小写主要是在包内可见
