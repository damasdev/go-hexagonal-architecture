package log

type Option struct {
	data *map[string]interface{}
	err  *error

	skip int
}

type OptionFunc func(opt *Option)

func WithData(data map[string]interface{}) OptionFunc {
	return func(opt *Option) {
		opt.data = &data
	}
}

func WithError(err error) OptionFunc {
	return func(opt *Option) {
		opt.err = &err
	}
}

func WithSkip(skip int) OptionFunc {
	return func(opt *Option) {
		opt.skip = skip
	}
}

func (opt *Option) GetData() *map[string]interface{} {
	return opt.data
}

func (opt *Option) GetError() *error {
	return opt.err
}

func (opt *Option) GetSkip() int {
	return opt.skip
}
