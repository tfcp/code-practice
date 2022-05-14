package decorator

import "fmt"

// 给我们的屎山代码新增逻辑

type ShitMountain interface {
	BigLogic()
}

type ShitMountainA struct {
}

func (this *ShitMountainA) BigLogic() {
	fmt.Println("this is A logic")
}

type ShitMountainB struct {
}

func (this *ShitMountainB) BigLogic() {
	fmt.Println("this is B logic")
}

// 核心点(别的屎山代码 咱们可没动, 咱们就用咱们自己的屎山代码 井水不犯河水)
type MyShitMountain struct {
	sm    ShitMountain
	myVal string
}

func NewMyShitMountain(s ShitMountain, myVal string) MyShitMountain {
	return MyShitMountain{
		sm:    s,
		myVal: myVal,
	}
}

func (this *MyShitMountain) BigLogic() {
	this.sm.BigLogic()
	fmt.Println("my logic: ", this.myVal)
}
