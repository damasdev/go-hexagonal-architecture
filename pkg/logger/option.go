package logger

type options struct {
	data *map[string]interface{}
	err  *error
}

type option func(opts *options)

func WithData(data map[string]interface{}) option {
	return func(opts *options) {
		opts.data = &data
	}
}

func WithError(err error) option {
	return func(options *options) {
		options.err = &err
	}
}

func (opts *options) getData() *map[string]interface{} {
	return opts.data
}

func (opts *options) getError() *error {
	return opts.err
}
