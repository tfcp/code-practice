package single

import (
	"fmt"
	"sync"
)

// 懒汉模式(加载慢)

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

type LazySingleton struct {
}

func GetLazyInstance() *LazySingleton {
	if lazySingleton == nil {
		once.Do(func() {
			fmt.Println("once lazy")
			lazySingleton = new(LazySingleton)
		})
	}
	return lazySingleton
}
