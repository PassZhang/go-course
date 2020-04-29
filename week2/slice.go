/*
切片(slice)
切片是长度可变的数组，（具有相同数据类型的数据项组成的一组长度可变的序列），切片有三部分组成
- 指针： 执行切片的第一个元素指向的数组元素的地址
- 长度： 切片元素的数量
- 容量： 切片开始到结束位置元素的数量

*/


package main 

import "fmt"

func main() {
	/*
	声明
	切片声明需要制定组成元素的类型，但不需要制定存储元素的数量（长度）。在切片声明后，会被初始化为nil，表示暂不存在的切片
	*/
	var names []string
	fmt.Printf("%T, %t, %v\n", names, names == nil, names)

	/*
	初始化
	1. 使用字面量初始化： []type{v1, v2, .. , vn}
	2. 使用字面量初始化空切片： []type{}
	3. 指定长度和容器字面量初始化：[]type{im:vm, in:vn, ilength:vlength}
	4. 使用make函数初始化make([]type, len)/make([]type, len, cap), 通过make函数创建长度为len， 容量为cap的切片，len 必须小于等于cap
	5. 使用数组切片操作初始化：array[start:end]/array[start:end:cap](end<=cap<=len)
	*/
	fmt.Printf("--------------------------------\n")

	scores := []int{60, 68, 70, 80, 95, 85}
	var scores00 []int = []int{80, 90, 95}
		scores01 := []int{}
		scores02 := []int{0:80, 1:85, 2:95}
		scores03 := make([]int, 10)
		scores04 := make([]int, 3, 10)
		scores05 := scores[1:3]
		scores06 := scores[1:3:4]


	fmt.Printf("%T, %T, %T, %T, %T\n", scores00, scores01, scores02, scores03, scores05)

	fmt.Println(scores00)
	fmt.Println(scores01)
	fmt.Println(scores02)
	fmt.Println(scores03)
	fmt.Println(scores04)
	fmt.Println(scores05)
	fmt.Println(scores06)
	
	fmt.Printf("--------------------------------\n")

	// 操作：
	// 1. 获取切片的长度
	// 使用len函数可获取切片的长度，使用cap函数可以获取切片的容量
	students := make([]string, 4, 10)
	fmt.Println(len(students), cap(students))
	fmt.Printf("%q\n", students)


	fmt.Printf("--------------------------------\n")
	// 2. 访问和修改
	// 通过对编号对切片元素进行访问和修改，元素的编号从左到右依次为：0， 1， 2， ..., n(n为切片长度-1)
	fmt.Printf("%q, %q\n", students[0], students[1])

	students[0] = "KK"
	students[1] = "woniu"
	students[2] = "ada"
	students[3] = "xiaohong"
	fmt.Printf("%q, %q, %q\n", students[0], students[1], students[2])
	fmt.Printf("%q\n", students)

	fmt.Printf("--------------------------------\n")
	// 3. 切片创建切片： slice[start:end]用于创建一个新的切片，end <= src_cap
	teachers := [...]string{"kk", "woniu", "ada", "xuegao", "xiaohong"}
	teachers00 := teachers[:]
	teachers01 := teachers[0:3]
	teachers02 := teachers[1:4]
	teachers03 := teachers00[1:4]

	fmt.Printf("%q\n", teachers)
	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)

	fmt.Printf("--------------------------------\n")
	// 新创建切片长度和容量计算： len: end-start， cap： src_cap-start
	// 切片共享底层数组，若某个切片元素发生变量，则数组其他有共享元素的切片也会发生变化
	teachers01[2] = "小林"
	fmt.Printf("%q\n", teachers)
	fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)

	fmt.Printf("--------------------------------\n")
	// slice[start:end:cap]可用于限制新切片的容量值，end<=cap<=src_cap
	teachers04 := teachers[1:4:4]
	teachers05 := teachers00[1:3:3]

	fmt.Printf("%q\n", teachers)
	fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
	fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)
	fmt.Printf("%d, %d, %q\n", len(teachers04), cap(teachers04), teachers04)
	fmt.Printf("%d, %d, %q\n", len(teachers05), cap(teachers05), teachers05)
    // 新创建切片长度和容量计算： len：end-start， cap：cap-start
	fmt.Printf("--------------------------------\n")

	// 遍历
	// 可以通过for+len+访问方式或者for-range方式对切片中元素进行遍历
	langs := []string{"go", "python", "c#", "java", "f#", "c", "c++", "lua", "lisp", "php", "rust" }
	for i := 0; i < len(langs); i++ {
		fmt.Printf("%d, %q\n", i, langs[i])
	}

	fmt.Printf("--------------------------------\n")
	for i, v := range langs {
		fmt.Printf("%d, %q\n", i, v)
	}
	// 使用for-range遍历切片，range返回两个元素分别为切片元素索引和值

	// 增加元素
	// 使用append对切片增加一个或者多个元素并返回修改后切片，当长度在容量范围内时只增加长度，容量和底层数组不变。当长度超过容量范围则会创建一个新的底层数组并对容量进行智能运算（元素数量 < 1024时，约按原容量1倍增加，>1024 时约按原容量的0.25倍增加）
	fmt.Printf("--------------------------------\n")
	nums := []int{1, 2, 3, 4, 5}

	nums2 := nums[:]
	fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

	nums2 = append(nums2, 6, 7, 8, 9, 10, 11)
	fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

	nums2[0] = 0
	nums2[1] = 1
	fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
	fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

	nums2 = append(nums2, 0, 1, 3, 4, 5)
	fmt.Printf("%d, %d\n", len(nums2), cap(nums2))
	fmt.Printf("--------------------------------\n")

	// 复制切片到另一个切片
	// 复制元素数量为src元素数量和dest元素数量的最小值
	users01 := []string{"00", "01"}
	users02 := []string{"10", "11", "12"}
	users03 := []string{"20", "21", "22", "23"}

	fmt.Printf("%q, %q, %q\n", users01, users02, users03)
	copy(users01, users02)		// 拷贝users02 复制给users01
	fmt.Printf("%q, %q, %q\n", users01, users02, users03)
	copy(users03, users02)		// 拷贝users02 复制给users03
	fmt.Printf("%q, %q, %q\n", users01, users02, users03)

	// 使用
	// 移除元素
	fmt.Printf("--------------------------------\n")
	elements := []int{0, 1, 2, 3, 4, 5}
	copy(elements[3:], elements[4:]) // 拷贝
	fmt.Println(elements[:len(elements) - 1])
	fmt.Printf("--------------------------------\n")

	// 队列
	// 先进先出
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 3)
	queue = append(queue, 2)

	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue[0])
	fmt.Printf("--------------------------------\n")

	// 堆栈
	// 先进后出
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)

	fmt.Println(stack[len(stack) - 1])
	stack = stack[:len(stack) - 1]
	fmt.Println(stack[len(stack) - 1])
	stack = stack[:len(stack) - 1]
	fmt.Println(stack[len(stack) - 1])
}