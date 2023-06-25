package log

type Options struct {
	data *map[string]interface{}
	err  *error
}

type Option func(opts *Options)

func WithData(data map[string]interface{}) Option {
	return func(opts *Options) {
		opts.data = &data
	}
}

func WithError(err error) Option {
	return func(opts *Options) {
		opts.err = &err
	}
}

func (options *Options) GetData() *map[string]interface{} {
	return options.data
}

func (options *Options) GetError() *error {
	return options.err
}
