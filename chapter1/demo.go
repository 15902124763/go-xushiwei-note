package chapter1

import "fmt"

// 顺序编程

// 2.1 变量:申请内存存储
// 2.1.1 变量声明
func varDemo()  {
	var v1 int
	var v2 string
	var v3 [3]int
	var v4 []int // 数组切片

	var v5 struct{
		f int
	}

	var v6  *int // 指针
	var v7  map[int] string // map

	v1 = 1
	v2 = "2"

	fmt.Print("\n",v1)
	fmt.Print("\n", v2)
	fmt.Print(v3)
	fmt.Print(v4)
	fmt.Print(v5.f)
	fmt.Print(v6)
	fmt.Print(v7)
	print()
}

// 2.1.4 匿名变量
func GetName()(firstName, secondName string)  {
	return "allen","july"
}

// 2.2常量：编译期就一知且不可改变的值，常量可以使数值类型（包括整型、浮点类型和复数类型）、布尔类型、字符串类型等


func Open()  {
	// 变量声明
	varDemo()
	// 匿名变量
	name, secondName := GetName()
	fmt.Print("\n",name, secondName)
}
