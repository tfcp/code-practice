package single

// 饿汉模式

type HungrySingleton struct {
}

var hungrySingleton *HungrySingleton

func init() {
	hungrySingleton = &HungrySingleton{}
}

func GetHungryInstance() *HungrySingleton {
	return hungrySingleton
}
