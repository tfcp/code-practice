package single

import (
	"fmt"
	"testing"
)

func Test_Hungry(t *testing.T) {
	a := GetHungryInstance()
	b := GetHungryInstance()
	if a == b {
		fmt.Println("lazySingle")
	}
}
