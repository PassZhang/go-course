package main

import "fmt"

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName（接收者为Dog的指针对象）

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func (dog *Dog) PCall() {
	fmt.Printf("%s： 汪汪\n", dog.name)

}

// 定义User结构体方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee 结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}

}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)

}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {

	dog := Dog{"豆豆"}

	dog2 := &Dog{"大毛"}

	// 方法值
	methodv01 := dog.PSetName
	methodv02 := dog.Call
	methodv03 := dog2.PSetName
	methodv04 := dog2.Call

	fmt.Printf("%T, %T, %T, %T\n", methodv01, methodv02, methodv03, methodv04)

	// 通过方法值调用方法
	fmt.Printf("%p, %p\n", &dog, dog2)

	methodv01("乐乐1")
	methodv02() // ???

	methodv03("乐乐2")
	methodv04() // ???

	fmt.Println(dog.name, dog2.name)

	// // 为结构体Dog定义方法PCall（接收者为Dog的指针对象）

	// methodv05 := dog.PCall(小明)
	// methodv06 := dog2.PCall(小红)

	// methodv05()
	// methodv06()

	// 方式表达式
	methode01 := (*Dog).PSetName
	methode02 := Dog.Call

	fmt.Printf("%T, %T\n", methode01, methode02)

	// 通过方法表达式指定对象调用方法
	methode01(&dog, "贝贝1") // (*Dog).PSetName(&dog, "贝贝1")
	methode01(dog2, "贝贝2") // (*Dog).PSetName(dog2, "贝贝2")

	methode02(dog)   //Dog.Call(dog)
	methode02(*dog2) //Dog.Call(*dog2)

	// 取引用
	// methode03 := Dog.PSetName
	methode04 := (*Dog).Call

	fmt.Printf("%T\n", methode04)

	methode04(&dog)
	methode04(dog2)
}
