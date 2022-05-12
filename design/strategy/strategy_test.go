package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T){
	case1 := NewStrategy("CaseOneVar1","CaseOneVar2",1,&CaseOne{})
	case1.Do()

	case2 := NewStrategy("CaseTwoVar1","CaseTwoVar2",2,&CaseTwo{})
	case2.Do()
}