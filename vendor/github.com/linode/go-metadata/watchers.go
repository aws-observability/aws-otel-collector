package metadata

import (
	"context"
)

// InstanceWatcher watches for any changes that are reflected
// in the Client.GetInstance(...) function result.
type InstanceWatcher struct {
	Updates chan *InstanceData
	Errors  chan error

	watcher *genericWatcher[*InstanceData]
}

// Start starts the watcher.
// NOTE: Start should only be called once per-watcher.
func (watcher *InstanceWatcher) Start(ctx context.Context) {
	watcher.watcher.Start(ctx)
}

// Close closes the watcher and all related channels.
// If applicable, close will also cancel the poller for this watcher.
func (watcher *InstanceWatcher) Close() {
	watcher.watcher.Close()
}

// NewInstanceWatcher creates a new InstanceWatcher for monitoring
// changes to the current Linode instance.
func (c *Client) NewInstanceWatcher(opts ...WatcherOption) *InstanceWatcher {
	coreWatcher := newGenericWatcher(
		func(ctx context.Context) (*InstanceData, error) {
			return c.GetInstance(ctx)
		},
		opts...,
	)

	return &InstanceWatcher{
		Updates: coreWatcher.Updates,
		Errors:  coreWatcher.Errors,
		watcher: coreWatcher,
	}
}

// NetworkWatcher watches for any changes that are reflected
// in the Client.GetNetwork(...) function result.
type NetworkWatcher struct {
	Updates chan *NetworkData
	Errors  chan error

	watcher *genericWatcher[*NetworkData]
}

// Start starts the watcher.
// NOTE: Start should only be called once per-watcher.
func (watcher *NetworkWatcher) Start(ctx context.Context) {
	watcher.watcher.Start(ctx)
}

// Close closes the watcher and all related channels.
// If applicable, close will also cancel the poller for this watcher.
func (watcher *NetworkWatcher) Close() {
	watcher.watcher.Close()
}

// NewNetworkWatcher creates a new NetworkWatcher for monitoring
// changes to the current Linode instance.
func (c *Client) NewNetworkWatcher(opts ...WatcherOption) *NetworkWatcher {
	coreWatcher := newGenericWatcher(
		func(ctx context.Context) (*NetworkData, error) {
			return c.GetNetwork(ctx)
		},
		opts...,
	)

	return &NetworkWatcher{
		Updates: coreWatcher.Updates,
		Errors:  coreWatcher.Errors,
		watcher: coreWatcher,
	}
}
