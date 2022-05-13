## go编程内功之策略模式

### 前言
go语言推崇的是函数式编程, 这样的好处在于开发快速, 上手快, 对新手友好, 对于一些短平快的项目(例如组件类服务)开发非常友好。
相对的坏处就是过于自由, 项目代码的下限很低, 对于一些大型项目的维护性可能会变得很差(这一点和php很像)。
本人就在这里记录一些非常实用的用于go大型项目治理的编码方式, 今天首先来介绍策略模式的应用。


### 策略模式介绍
策略模式是面向对象设计模式中的其中一种模式, 专门用于解决海量if/else, switch/case的场景, 
这个场景在项目中是非常常见的, 所以策略模式也是非常实用的设计模式, 问题编码场景如下:
````
package main

func main() {
	val := "1"
	Logic(val)
}

// 业务方法
func Logic(val string) {
	if val == "1" {
		// logic 1
	} else if val == "2" {
		// logic 2
	} else if val == "..." {
		// 海量 if else 
        // logic ...
	} else {
		// logic x
	}
}
````
这种一大堆的if/else, 看起来头疼, 改起来更容易出问题, 是项目中最常见的问题

### 策略模式核心代码演示
策略应用代码, 如下
````
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

// 不同case不同struct
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
````

业务使用, 如下
````
package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T) {
	strategyType := "one"
    // 可以看到业务代码简单一行即可
	caseStrategy := NewStrategy("CaseOneVar1", "CaseOneVar2", 1, StrategyList[strategyType])
	caseStrategy.Do()
}
````

### 后记
    策略模式非常好的反映了面向对象规范中的开闭原则(对修改关闭, 对新增开放), 每次修改代码, 不用再去污染if/else, 新增场景对象即可, 极大降低代码开发风险, 提升项目维护性


