package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T){
	strategyType := "one"

	caseStra := NewStrategy("CaseOneVar1","CaseOneVar2",1,StrategyList[strategyType])
	caseStra.Do()

}