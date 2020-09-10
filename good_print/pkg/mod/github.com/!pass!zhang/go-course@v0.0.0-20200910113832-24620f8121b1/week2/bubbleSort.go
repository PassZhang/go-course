package main

import "fmt"

func main() {
	// [168 180 166 176 169]
	// heigh := []int{168, 180, 166, 176, 169}

	// a := 1
	// b := 2
	// // a, b := 1, 2

	// a, b = b, a
	// fmt.Println(a, b)

	// tmp := a // 这时a=2，b=1， tmp赋值a的值，tmp=2
	// a = b    // a 赋值 b的值，a = 1
	// b = tmp  // b 赋值 tmp的值，b = 2
	// fmt.Println(a, b)

	heigh := []int{168, 180, 166, 176, 169}

	// for i, v := range heigh {
	// 	fmt.Println(i, v)
	// }
	//从小到大
	for j := 0; j < len(heigh); j++ {
		for i := 0; i < len(heigh)-1; i++ {
			if heigh[i] > heigh[i+1] {
				heigh[i], heigh[i+1] = heigh[i+1], heigh[i]
			}
			fmt.Println(i, heigh)
		}
	}
	fmt.Println(heigh)
	// 如果heigh[i]的值大于 heigh[i+1]，这时第一轮也就是变量的第一个值和第二个值进行比较，如果第一个值比二个值大就执行下面的赋值，把第一个值赋值给第二个值，第二个值赋值给第一个值。
	//这时已经赋值成功，小一点的数值已经放到了数组的最前面。
	// for 进行递增，这时再进行第二个值和第三个值的比较，如果第二个值比第三个值大，然后再进行赋值第二个值和第三个值的赋值，以此算法进行循环排序。
	// 最终将最小的值放到最前面，最大的值放到最后面

	//从大到小
	for j := 0; j < len(heigh); j++ {
		for i := 0; i < len(heigh)-1; i++ {
			if heigh[i] < heigh[i+1] {
				heigh[i], heigh[i+1] = heigh[i+1], heigh[i]
			}
			fmt.Println(i, heigh)
		}
	}
	fmt.Println(heigh)
}
