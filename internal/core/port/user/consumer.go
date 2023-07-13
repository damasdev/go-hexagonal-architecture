package user

import "context"

type UserConsumer interface {
	Subscribe(ctx context.Context, name string, handler func(message []byte) error)
}
