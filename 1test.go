package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type CDATA struct {
	Content string `xml: ",cdata"` // 定义数据使用CDATA包含
}

type Comment struct {
	Content string `xml:",comment"` // 定义数据为注释
}

type Address struct {
	City   string `xml:"address>city"`
	Street string `xml:"address>street"`
	No     string `xml:"address>no"`
}

// 定义User结构体，并使用标签声明json属性
type User struct {
	ID           int       `xml:"id,attr"`             // xml节点属性，且属性名为id
	Name         string    `xml:"name"`                // xml节点键重命名为name
	Password     string    `xml:"-"`                   // xml序列化忽略该字段
	gender       bool      `xml:"gender"`              // 非公开字段不进行序列化
	Online       bool      `xml:"is_online,omitempty"` // xml节点重命名为is_online
	RegisterTime time.Time `xml:"register_time"`       // xml节点重命名为register_time
	Status       int       `xml:",omitempty"`          // xml节点使用属性名Status，当为0值是不进行序列化
	Address
	Desc    CDATA   `xml:"desc"`
	Comment Comment `xml:"comment"`
	// Extend string `xml:",innerxml"` // 定义数据不使用xml转义
}

type Root struct {
	XMLName xml.Name `xml:"root"` // 定义根节点名称
	User    []User   `xml:"User"` //
}




func main() {

	// 定义用户类型切片
	users := make([]User, 10, 20)

	// 初始化切片
	for i := 0; i < 10; i++ {
		users[i] = User{
			ID: i,
			Name: fmt.Sprintf("user.%d", i),
			Password: fmt.Sprintf("password.%d", i),
			gender: i%2 == 0,
			Online: i%3 == 0,
			RegisterTime: time.Now(),
			Status: i % 5,
			Address: Address{
				City: fmt.Sprintf("City.%d", i),
				Street: fmt.Sprintf("Street.%d", i),
				No: fmt.Sprintf("No.%d", i),
			},
			Desc: CDATA{fmt.Sprintf("Desc.%d", i)},
			Comment: Comment{fmt.Sprintf("Comment.%d", i)},
			// Ex
		}
	}
}
