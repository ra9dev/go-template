package shutdown

import (
	"context"
	"sync"
	"time"
)

type (
	CallbackFunc func(ctx context.Context)

	GracefulShutdown struct {
		mu        *sync.RWMutex
		callbacks []CallbackFunc
	}
)

func (s *GracefulShutdown) Add(fn CallbackFunc) {
	s.mu.Lock()
	s.callbacks = append(s.callbacks, fn)
	s.mu.Unlock()
}

func (s *GracefulShutdown) ForceShutdown() {
	shutdownCTX, shutdownCancel := context.WithTimeout(context.Background(), Timeout)
	defer shutdownCancel()

	s.mu.RLock()
	callbacks := s.callbacks
	s.mu.RUnlock()

	wg := new(sync.WaitGroup)
	wg.Add(len(callbacks))

	for _, callback := range callbacks {
		threadSafeCallback := callback

		go func() {
			defer wg.Done()

			threadSafeCallback(shutdownCTX)
		}()
	}

	wg.Wait()
}

func (s GracefulShutdown) Wait() {
	time.Sleep(Timeout)
}
