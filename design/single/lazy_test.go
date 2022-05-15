package single

import (
	"fmt"
	"testing"
)

func Test_Lazy(t *testing.T) {
	a := GetLazyInstance()
	b := GetLazyInstance()
	if a == b {
		fmt.Println("this is lazy")
	}
}
