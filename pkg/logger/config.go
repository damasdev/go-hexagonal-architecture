package logger

type configs struct {
	name  *string
	level *LogLevel
}

type config func(opts *configs)

func WithLevel(level LogLevel) config {
	return func(opts *configs) {
		opts.level = &level
	}
}

func WithName(name string) config {
	return func(configs *configs) {
		configs.name = &name
	}
}

func (opts *configs) getName() *string {
	return opts.name
}

func (opts *configs) getLevel() *LogLevel {
	return opts.level
}
