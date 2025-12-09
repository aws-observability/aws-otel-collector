/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package arrow

import "errors"

var (
	ErrBuilderAlreadyReleased = errors.New("builder already released")
	ErrInvalidResourceID      = errors.New("invalid resource ID")
	ErrInvalidScopeID         = errors.New("invalid scope ID")
)
