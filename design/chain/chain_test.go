package chain

import (
	"fmt"
	"testing"
)

func Test_Chain(t *testing.T) {
	chain := new(FilterChain)
	chain.Add(&FilterA{})
	chain.Add(&FilterB{})
	fmt.Println(chain.Filter("test"))
}
