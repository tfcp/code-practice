package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T) {
	strategyType := "one"
	caseStrategy := NewStrategy("CaseOneVar1", "CaseOneVar2", 1, StrategyList[strategyType])
	caseStrategy.Do()
}
