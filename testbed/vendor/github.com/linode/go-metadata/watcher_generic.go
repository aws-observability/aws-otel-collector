package metadata

import (
	"context"
	"reflect"
	"time"
)

const DefaultWatcherInterval = 5 * time.Minute

// WatcherOption represents an option that can be used to configure
// a watcher.
type WatcherOption func(options *watcherConfig)

type watcherConfig struct {
	Interval time.Duration
}

type resourceFetchFunc[T any] func(ctx context.Context) (T, error)

// genericWatcher is a resource-agnostic watcher that allows
// us to re-use polling logic across various watcher implementations.
type genericWatcher[T any] struct {
	Updates chan T
	Errors  chan error
	cancel  chan bool

	fetchResource resourceFetchFunc[T]

	interval time.Duration
}

// Start starts the watcher.
func (watcher *genericWatcher[T]) Start(ctx context.Context) {
	var oldData T

	ticker := time.NewTicker(watcher.interval)
	defer ticker.Stop()

	defer func() {
		close(watcher.Updates)
		close(watcher.Errors)
		close(watcher.cancel)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-watcher.cancel:
			return
		case <-ticker.C:
			data, err := watcher.fetchResource(ctx)
			if err != nil {
				watcher.Errors <- err
			}

			if !reflect.DeepEqual(data, oldData) {
				watcher.Updates <- data
				oldData = data
			}
		}
	}
}

// Close closes the watcher and all related channels.
func (watcher *genericWatcher[T]) Close() {
	// Send a signal to cancel the poller.
	// All channels wil be implicitly cleaned up in
	// the start goroutine, else they will be
	// cleaned up by the garbage collector.
	watcher.cancel <- true
}

// newGenericWatcher creates a new instance of an endpoint-agnostic watcher.
func newGenericWatcher[T any](
	fetchResource resourceFetchFunc[T],
	opts ...WatcherOption,
) *genericWatcher[T] {
	watcherOpts := watcherConfig{
		Interval: DefaultWatcherInterval,
	}

	for _, opt := range opts {
		opt(&watcherOpts)
	}

	return &genericWatcher[T]{
		Updates:       make(chan T),
		Errors:        make(chan error),
		cancel:        make(chan bool, 1),
		interval:      watcherOpts.Interval,
		fetchResource: fetchResource,
	}
}

// WatcherWithInterval configures the interval at which
// a watcher should poll for changes.
// Default: 5 minutes
func WatcherWithInterval(duration time.Duration) WatcherOption {
	return func(options *watcherConfig) {
		options.Interval = duration
	}
}
