package main

type Api interface {
	Do(val string) string
}

// 简单工厂方法
// golang没有构造方法, 一直通过NewXXX函数初始化
// NewXXX返回的是interface时 就是简单工厂模式
func NewSimpleFac(ty string) Api {
	switch ty {
	case "t1":
		return new(T1)
	case "t2":
		return new(T2)
	}
	return nil
}

type T1 struct {
}

func (t1 *T1) Do(val string) string {
	return "t1: " + val
}

type T2 struct {
}

func (t2 *T2) Do(val string) string {
	return "t2: " + val
}
