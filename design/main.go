package main

import "fmt"

func main() {
	// 单例
	//s1 := GetInstance()
	//s2 := GetInstance()
	//s3 := GetInstance()
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(s3)

	// 简单工厂
	t1 := NewSimpleFac("t1")
	fmt.Println(t1.Do("test"))
	t2 := NewSimpleFac("t2")
	fmt.Println(t2.Do("test"))
}
