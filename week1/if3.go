// 判断语句if
// package main 

// import "fmt"

// func main()  {
// 	// 条件表达式没有括号
// 	// 支持一个初始化表达式（可以是多变量初始化语句）
// 	// 左大括号必须和条件语句同一行

// 	if number := 7; number < 1 {
// 		fmt.Println(1)
// 	} else if number >= 1 && number <= 10 {
// 		fmt.Println(2)
// 	} else {
// 		fmt.Println(3)
// 	}
// }

// 优雅处理error
package main 
import (
	"net/http"
	"fmt"
)
func main()  {
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
	}
}