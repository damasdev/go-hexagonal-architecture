package consumer

type Config struct {
	exchangeKind       string
	exchangeDurable    bool
	exchangeAutoDelete bool
	exchangeInternal   bool
	exchangeNoWait     bool

	queueDurable    bool
	queueAutoDelete bool
	queueExclusive  bool
	queueNoWait     bool

	prefetchCount  int
	prefetchSize   int
	prefetchGlobal bool

	consumeAutoAck   bool
	consumeExclusive bool
	consumeNoLocal   bool
	consumeNoWait    bool
}

func LoadDefaultConfig() Config {
	return Config{
		exchangeKind:       "direct",
		exchangeDurable:    true,
		exchangeAutoDelete: false,
		exchangeInternal:   false,
		exchangeNoWait:     false,

		queueDurable:    true,
		queueAutoDelete: false,
		queueExclusive:  false,
		queueNoWait:     false,

		prefetchCount:  1,
		prefetchSize:   0,
		prefetchGlobal: false,

		consumeAutoAck:   false,
		consumeExclusive: false,
		consumeNoLocal:   false,
		consumeNoWait:    false,
	}
}
