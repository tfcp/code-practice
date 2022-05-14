package chain

// 责任链应用场景 敏感词过滤场景

type ErrFilter interface {
	Filter(content string) bool
}

type FilterA struct {
}

func (this *FilterA) Filter(content string) bool {
	if content == "test" {
		return false
	}
	return true
}

type FilterB struct {
}

func (this *FilterB) Filter(content string) bool {
	if content == "test" {
		return false
	}
	return true
}

// 核心业务责任链
type FilterChain struct {
	filters []ErrFilter
}

func (this *FilterChain) Add(f ErrFilter) {
	this.filters = append(this.filters, f)
}

func (this *FilterChain) Filter(content string) bool {
	if len(this.filters) > 0 {
		for _, filter := range this.filters {
			if !filter.Filter(content) {
				return false
			}
		}
	}
	return true
}
