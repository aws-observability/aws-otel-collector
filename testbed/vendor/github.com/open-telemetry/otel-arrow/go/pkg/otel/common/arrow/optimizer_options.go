/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package arrow

type Options struct {
	Sort  bool
	Stats bool
}

func WithSort() func(*Options) {
	return func(o *Options) {
		o.Sort = true
	}
}

func WithStats() func(*Options) {
	return func(o *Options) {
		o.Stats = true
	}
}
