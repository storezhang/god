package mengpo

var _ option = (*optionBefore)(nil)

type optionBefore struct {
	before beforeFunc
}

// Before 配置生命周期前的操作
func Before(before beforeFunc) *optionBefore {
	return &optionBefore{
		before: before,
	}
}

func (b *optionBefore) apply(options *options) {
	options.before = b.before
}
