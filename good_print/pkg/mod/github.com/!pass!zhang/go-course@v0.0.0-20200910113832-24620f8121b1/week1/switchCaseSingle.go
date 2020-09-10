package main

import "fmt"

func main() {
	// 单值
	operate := "delete"
	switch operate {
        case "add":
            fmt.Println("添加")
        case "delete":
            fmt.Println("删除")
        case "modify":
            fmt.Println("修改")
        case "query":
            fmt.Println("查询")
        default:
            fmt.Println("操作错误")
    }
    input := "y"
    switch input {
    case "yes", "y":
        fmt.Println("确认")
    case "no", "n":
        fmt.Println("取消")
    default:
        fmt.Println("操作错误")
    }
}