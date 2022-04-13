package main

import (
	"fmt"
	"sync"
)

// 单例
type Singleton struct {
	val int
}

var (
	singleton *Singleton
	once sync.Once
)

func GetInstance () *Singleton{
	once.Do(func() {
		fmt.Println("single start")
		singleton = &Singleton{}
		singleton.val = 101
	})
	return singleton
}
