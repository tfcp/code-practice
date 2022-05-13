package strategy

import "fmt"

var (
	// 代码核心(策略变量列表)
	StrategyList = map[string]Strategy{
		"one": &CaseOne{},
		"two": &CaseTwo{},
	}
)

// 代码核心(策略请求上下文, 封装struct方便后续维护)
type StrategyContext struct {
	// 变量
	VarOne, VarTwo string
	VarThree       int
}

type StrategyBusiness struct {
	ctx      *StrategyContext
	strategy Strategy
}

func (this *StrategyBusiness) Do() {
	this.strategy.Do(this.ctx)
}

type Strategy interface {
	Do(*StrategyContext)
}

type CaseOne struct {
}

func (this *CaseOne) Do(ctx *StrategyContext) {
	ret := fmt.Sprintf("caseOne logic v1:%s | v2: %s | v3: %d",
		ctx.VarOne, ctx.VarTwo, ctx.VarThree,
	)
	fmt.Println(ret)
}

type CaseTwo struct {
}

func (this *CaseTwo) Do(ctx *StrategyContext) {
	ret := fmt.Sprintf("caseTwo logic v1:%s | v2: %s | v3: %d",
		ctx.VarOne, ctx.VarTwo, ctx.VarThree,
	)
	fmt.Println(ret)
}

// 代码核心(生成)
func NewStrategy(varOne, varTwo string, varThree int, strategy Strategy) *StrategyBusiness {
	return &StrategyBusiness{
		ctx: &StrategyContext{
			VarOne:   varOne,
			VarTwo:   varTwo,
			VarThree: varThree,
		},
		strategy: strategy,
	}
}
