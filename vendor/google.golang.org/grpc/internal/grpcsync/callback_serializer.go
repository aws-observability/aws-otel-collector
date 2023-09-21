/*
 *
 * Copyright 2022 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcsync

import (
	"context"
	"sync"

	"google.golang.org/grpc/internal/buffer"
)

// CallbackSerializer provides a mechanism to schedule callbacks in a
// synchronized manner. It provides a FIFO guarantee on the order of execution
// of scheduled callbacks. New callbacks can be scheduled by invoking the
// Schedule() method.
//
// This type is safe for concurrent access.
type CallbackSerializer struct {
<<<<<<< HEAD
	// Done is closed once the serializer is shut down completely, i.e all
	// scheduled callbacks are executed and the serializer has deallocated all
	// its resources.
	Done chan struct{}
=======
	// done is closed once the serializer is shut down completely, i.e all
	// scheduled callbacks are executed and the serializer has deallocated all
	// its resources.
	done chan struct{}
>>>>>>> main

	callbacks *buffer.Unbounded
	closedMu  sync.Mutex
	closed    bool
}

// NewCallbackSerializer returns a new CallbackSerializer instance. The provided
// context will be passed to the scheduled callbacks. Users should cancel the
// provided context to shutdown the CallbackSerializer. It is guaranteed that no
// callbacks will be added once this context is canceled, and any pending un-run
// callbacks will be executed before the serializer is shut down.
func NewCallbackSerializer(ctx context.Context) *CallbackSerializer {
<<<<<<< HEAD
	t := &CallbackSerializer{
		Done:      make(chan struct{}),
		callbacks: buffer.NewUnbounded(),
	}
	go t.run(ctx)
	return t
=======
	cs := &CallbackSerializer{
		done:      make(chan struct{}),
		callbacks: buffer.NewUnbounded(),
	}
	go cs.run(ctx)
	return cs
>>>>>>> main
}

// Schedule adds a callback to be scheduled after existing callbacks are run.
//
// Callbacks are expected to honor the context when performing any blocking
// operations, and should return early when the context is canceled.
//
// Return value indicates if the callback was successfully added to the list of
// callbacks to be executed by the serializer. It is not possible to add
// callbacks once the context passed to NewCallbackSerializer is cancelled.
<<<<<<< HEAD
func (t *CallbackSerializer) Schedule(f func(ctx context.Context)) bool {
	t.closedMu.Lock()
	defer t.closedMu.Unlock()

	if t.closed {
		return false
	}
	t.callbacks.Put(f)
	return true
}

func (t *CallbackSerializer) run(ctx context.Context) {
	var backlog []func(context.Context)

	defer close(t.Done)
=======
func (cs *CallbackSerializer) Schedule(f func(ctx context.Context)) bool {
	cs.closedMu.Lock()
	defer cs.closedMu.Unlock()

	if cs.closed {
		return false
	}
	cs.callbacks.Put(f)
	return true
}

func (cs *CallbackSerializer) run(ctx context.Context) {
	var backlog []func(context.Context)

	defer close(cs.done)
>>>>>>> main
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			// Do nothing here. Next iteration of the for loop will not happen,
			// since ctx.Err() would be non-nil.
<<<<<<< HEAD
		case callback, ok := <-t.callbacks.Get():
			if !ok {
				return
			}
			t.callbacks.Load()
=======
		case callback, ok := <-cs.callbacks.Get():
			if !ok {
				return
			}
			cs.callbacks.Load()
>>>>>>> main
			callback.(func(ctx context.Context))(ctx)
		}
	}

	// Fetch pending callbacks if any, and execute them before returning from
<<<<<<< HEAD
	// this method and closing t.Done.
	t.closedMu.Lock()
	t.closed = true
	backlog = t.fetchPendingCallbacks()
	t.callbacks.Close()
	t.closedMu.Unlock()
=======
	// this method and closing cs.done.
	cs.closedMu.Lock()
	cs.closed = true
	backlog = cs.fetchPendingCallbacks()
	cs.callbacks.Close()
	cs.closedMu.Unlock()
>>>>>>> main
	for _, b := range backlog {
		b(ctx)
	}
}

<<<<<<< HEAD
func (t *CallbackSerializer) fetchPendingCallbacks() []func(context.Context) {
	var backlog []func(context.Context)
	for {
		select {
		case b := <-t.callbacks.Get():
			backlog = append(backlog, b.(func(context.Context)))
			t.callbacks.Load()
=======
func (cs *CallbackSerializer) fetchPendingCallbacks() []func(context.Context) {
	var backlog []func(context.Context)
	for {
		select {
		case b := <-cs.callbacks.Get():
			backlog = append(backlog, b.(func(context.Context)))
			cs.callbacks.Load()
>>>>>>> main
		default:
			return backlog
		}
	}
}
<<<<<<< HEAD
=======

// Done returns a channel that is closed after the context passed to
// NewCallbackSerializer is canceled and all callbacks have been executed.
func (cs *CallbackSerializer) Done() <-chan struct{} {
	return cs.done
}
>>>>>>> main
