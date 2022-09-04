package shutdown

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	globalShutdown GracefulShutdown

	timeout = time.Second * 5
)

func Timeout() time.Duration {
	return timeout
}

func init() {
	globalShutdown = GracefulShutdown{
		mu:        new(sync.RWMutex),
		callbacks: make([]CallbackFunc, 0),
	}

	osCTX, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel()

		<-osCTX.Done()
		globalShutdown.ForceShutdown()
	}()
}

func Add(fn CallbackFunc) {
	globalShutdown.Add(fn)
}

func Wait() {
	globalShutdown.Wait()
}
