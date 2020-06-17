package task

import "fmt"

var Name = "controllers task"
var Version string // 需要再运行启动之前进行初始化

func Call() {
	fmt.Println("Controller Call")
}

// 变量属性问题
// 包外可见
// 包外是否可以修改 = > 需要在包外不能修改值
var version string

// 提供对外修改的函数（读，不可写）
func GetVersion() string {
	return version
}

// 包在使用的时候自动对数据进行初始化

func init() {
	Version = "Controller v1.0"
	fmt.Println("Controller init")
}
