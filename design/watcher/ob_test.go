package watcher

import "testing"

func Test_Ob(t *testing.T) {
	bus := Business{}
	l := &Log{}  // 初始化 log对象
	u := &User{} // 初始化 user对象
	bus.Register(l)
	bus.Register(u)
	bus.Notice("test")
	bus.Remove(u)
	bus.Notice("test2222")
}
