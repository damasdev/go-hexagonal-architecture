package user

import "context"

type UserConsumer interface {
	Consume(ctx context.Context)
}
