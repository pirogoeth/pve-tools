package utils

import (
	"context"
)

type OnDoneCallback func()

func OnDone(ctx context.Context, cb OnDoneCallback) {
	<-ctx.Done()
	cb()
}
