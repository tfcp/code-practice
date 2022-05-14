package decorator

import (
	"fmt"
	"testing"
)

func Test_Decorator(t *testing.T) {
	d1 := NewMyShitMountain(&ShitMountainA{}, "my shit logic A")
	d1.BigLogic()
	fmt.Println("---------")
	d2 := NewMyShitMountain(&ShitMountainB{}, "my shit logic B")
	d2.BigLogic()
}
