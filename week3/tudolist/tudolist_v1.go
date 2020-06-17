package main 

import (
	"fmt"
	"strconv"
	"strings"
)

// 做一个命令行的任务管理器
// 用户管理

// 1. 考察函数，输入&输出， 复合数据结构，基本数据类型
// 2. 了解数据流程（对数据的操作流程，增、删、改、查）

/*
	1. 任务的输入（添加任务）
	2. 任务列表（查询任务列表）
	3. 任务修改
	4. 任务删除
	5. 详情
*/

// 任务
// id, 任务名称，开始时间， 结束时间，状态，负责人
// ID, name, start_time, end_time, status, user
// 使用map数据结构，[]map[string][string]

var todos = []map[string]string {
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "4", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}
const (
	id = "id"
	name = "name"
	startTime = "start_time"
	endTime = "end_time"
	status = "status"
	user = "user"
)

const (
	statusNew = "未执行"
	statusCompele = "完成"
)

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}

func printTask(task map[string]string)  {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID:", task[id])
	fmt.Println("任务名：", task[name])
	fmt.Println("开始时间：", task[startTime])
	fmt.Println("完成时间：", task[endTime])
	fmt.Println(strings.Repeat("-", 20))
}

func newTask() map[string]string  {
	// id生成（用于todos中最大的ID+1）
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func add()  {
	task := newTask()

	fmt.Println("请输入任务信息：")
	task[name] = input("任务名：")
	task[startTime] = input("开始时间：")
	task[user] = input("负责人：")
	todos = append(todos, task)
	fmt.Println("任务创建成功！")
}

func query()  {
	// all 显示全部
	q := input("请输入查询信息：")

	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

func main()  {
	EXIT:
	for {
		switch (input("请输入操作（add/query/exit/....）:")) {
		case "add":
			add()
		case "query":
			query()
		case "modify":
		case "delete":
		case "exit":
			// return 使用也可以退出
			break EXIT
		default:
			fmt.Println("输入指令错误")
		}
	}
}