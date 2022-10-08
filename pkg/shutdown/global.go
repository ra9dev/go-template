package shutdown

import (
	"context"
	"time"
)

var (
	globalShutdown = NewGracefulShutdown()

	timeout = time.Second * 5
)

func Timeout() time.Duration {
	return timeout
}

func Add(fn CallbackFunc) {
	globalShutdown.Add(fn)
}

func Wait() error {
	return globalShutdown.Wait()
}

func Context() context.Context {
	return globalShutdown.Context()
}
