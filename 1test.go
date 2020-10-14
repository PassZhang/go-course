package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 定义User结构体，并为每个属性定义标签

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

// 定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}", user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip

}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

// // 打印变量类型信息
// func displayType(t reflect.Type, tab string) {
// 	// 处理类型为nil
// 	if t == nil {
// 		fmt.Printf("<nil>")
// 		return
// 	}
// 	// 获取类型对应的枚举值使用选择语句分别处理每种类型
// 	switch t.Kind() {
// 	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
// 		// 针对基本数据类型显示类型名
// 		fmt.Printf("%s%s", tab, t.Name())
// 	case reflect.Array, reflect.Slice:
// 		// 针对数组和切片，直接打印Type对象
// 		fmt.Printf("%s%s", tab, t)
// 	case reflect.Map:
// 		// 对于映射类型打印键和值的Type对象
// 		fmt.Printf("%smap{\n", tab)
// 		fmt.Printf("%s\tKey:", tab)
// 		fmt.Printf("%s%s", tab, t.Key()) // 获取键的Type对象
// 		fmt.Println()
// 		fmt.Printf("%s\tValue: ", tab)
// 		fmt.Printf("%s%s", tab, t.Elem()) // 获取值的Type对象
// 		fmt.Println()
// 		fmt.Printf("%s}", tab)
// 	case reflect.Func:
// 		// 针对函数打印参数和返回值，对应可变参数在最后一个参数添加之后添加。。。
// 		fmt.Printf("%sfunc (", tab)

// 		// 打印参数信息
// 		// 获取参数数量并遍历
// 		for i := 0; i < t.NumIn(); i++ {
// 			fmt.Printf("%s", t.In(i)) // 根据索引获取第i个参数的Type 对象
// 			if i != t.NumIn()-1 {
// 				fmt.Printf(", ")
// 			}
// 		}
// 		if t.IsVariadic() {
// 			fmt.Printf("...")
// 		}
// 		fmt.Printf(") ")

// 		// 打印返回值信息
// 		if t.NumOut() > 0 {
// 			fmt.Printf("( ")
// 			for i := 0; i < t.NumOut(); i++ {
// 				fmt.Printf("%s", t.Out(i)) // 根据索引获取第i个返回值的Type对象
// 				if i != t.NumOut()-1 {
// 					fmt.Printf(", ")
// 				}
// 			}
// 			fmt.Printf(" )")
// 		}
// 		fmt.Printf("{}")
// 	case reflect.Struct:
// 		// 针对结构体显示结构体属性和方法
// 		fmt.Printf("%stype %s struct {\n", tab, t.Name())

// 		// 获取属性数量并遍历
// 		fmt.Printf("%s\tFields(%d):\n", tab, t.NumField())
// 		for i := 0; i < t.NumField(); i++ {
// 			field := t.Field(i)                                                      // 根据索引获取第i个属性的StructField对象
// 			fmt.Printf("%s\t\t%s\t%s\t`%s`", tab, field.Name, field.Type, field.Tag) // 显示属性名，属性的Type对象和标签
// 			fmt.Printf(",\n")
// 		}

// 		// 获取方法数量并遍历
// 		fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
// 		for i := 0; i < t.NumMethod(); i++ {
// 			// 打印方法信息
// 			displayMethod(t.Method(i), tab+"\t\t")
// 			fmt.Printf(",\n")
// 		}
// 		fmt.Printf("%s}", tab)
// 	case reflect.Ptr:
// 		// 对于指针类型，递归分析其引用值
// 		fmt.Printf("%s*{\n", tab)
// 		displayType(t.Elem(), tab+"\t") // 获取指针的引用值，并递归调用displayType进行分析显示
// 		// 打印指针变量的方法
// 		if t.NumMethod() > 0 { // 获取方法数量
// 			fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
// 			for i := 0; i < t.NumMethod(); i++ {
// 				// 打印方法信息
// 				displayMethod(t.Method(i), tab+"\t\t")
// 				if i != t.NumMethod()-1 {
// 					fmt.Printf(",\n")
// 				}
// 			}
// 		}
// 		fmt.Printf("\n%s", tab)
// 	default:
// 		fmt.Printf("%sUnkonw[%s]", tab, t)
// 	}
// }

// // 打印reflect.Method类型变量method的信息
// func displayMethod(method reflect.Method, tab string) {
// 	// 获取方法接收者
// 	t := method.Type
// 	fmt.Printf("%sfunc %s(", tab, method.Name) // 显示方法名

// 	// 打印参数信息
// 	// 获取参数数量并遍历
// 	for i := 0; i < t.NumIn(); i++ {
// 		fmt.Printf("%s", t.In(i)) // 根据索引获取第i个参数的Type对象
// 		if i != t.NumIn()-1 {
// 			fmt.Printf(", ")
// 		}
// 	}

// 	// 打印可变参数信息
// 	if t.IsVariadic() {
// 		fmt.Printf("...")
// 	}

// 	// 打印返回值信息
// 	fmt.Printf(") ")

// 	if t.NumOut() > 0 {
// 		fmt.Printf("(")
// 		// 获取返回值数量并遍历
// 		for i := 0; i < t.NumOut(); i++ {
// 			fmt.Printf("%s", t.Out(i)) // 根据索引获取第i个返回值的Type对象
// 			if i != t.NumOut()-1 {
// 				fmt.Printf(", ")
// 			}
// 		}
// 		fmt.Printf(") ")
// 	}
// 	fmt.Printf("{}")
// }

// 打印reflect.Value类型变量v的信息
func displayValue(value reflect.Value, tab string) {

	// 获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		// 针对整数类型，获取值对应的Type对象并使用Int函数转换为对应的基本类型，并使用strconv将转换为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatInt(value.Int(), 10))
	case reflect.Float32:
		// 针对浮点类型，获取值对应的Type对象并使用Float函数转换为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatFloat(value.Float(), 'E', -1, 64))
	case reflect.Bool:
		// 针对布尔类型，获取值对应的Type对象并使用Bool函数转换为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatBool(value.Bool()))
	case reflect.String:
		// 针对字符集类型，获取值对应的Type对象并打印值
		fmt.Printf("%s[%s] %s", tab, value.Type(), value)
	case reflect.Array:
		// 针对数组类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())

		// 获取数组的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") // 根据索引获取数组的每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Slice:
		// 针对切片类型，获取值对应的Type对象，长度和容量
		fmt.Printf("%s[%s](%d:%d) {\n", tab, value.Type(), value.Len(), value.Cap())

		// 获取切片的长度
		for i := 0; i < value.len(); i++ {
			displayValue(value.Index(i), tab+"\t") // 根据索引获取数组的每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Map:
		// 针对映射类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())
		// 获取映射迭代对象遍历键值对
		iter := value.MapRange()
		for iter.Next() { // 判断迭代对象是否到末尾
			displayValue(iter.Key(), tab+"\t") // 根据从迭代对象获取当前键，并调用displayValue递归显示
			fmt.Printf(" : ")
			displayValue(iter.Value(), "") // 根据从迭代对象获取当前值，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Struct:
		// 针对结构体类型，获取值对应的Type对象
		structType := value.Type()
		fmt.Printf("%s[%s] {\n", tab, structType)

		// 获取属性数量并遍历
		for i := 0; i < value.NumField(); i++ {
			structField := structType.Field(i)           // 根据索引获取属性的类型
			field := value.Field(i)                      // 根据索引获取属性的值
			fmt.Printf("%s\t%s:", tab, structField.Name) // 打印类型名
			displayValue(field, tab+"\t")                // 调用displayValue递归显示
			fmt.Printf(", \n")

		}
		fmt.Printf("%s}", tab)

	case reflect.Ptr:
		// 针对指针类型，获取值对应的Type对象
		fmt.Printf("%s[%s] (\n", tab, value.Type())

		// 获取指针的引用值，并递归调用displayValue函数递归分析显示
		displayValue(value.Elem(), tab+"\t")
		fmt.Printf("\n%s", tab)
	default:
		fmt.Printf("%sUnkonw[%#v]", tab, value)
	}
}

func main() {
	vars := make([]interface{}, 0, 20)

	var intV int = 1
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身，为人谋而不忠乎？与朋友交而不信呼？传不习呼？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var userV *User = NewUser(1, "kk", "12000000000", 1.68, "少年经不得顺境，中间经不得闲境，晚年经不得逆境", 72)
	var CloserV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, *userV, userV, CloserV)
	for _, v := range vars {
		displayValue(reflect.ValueOf(v), "")
		fmt.Println()
	}
}
