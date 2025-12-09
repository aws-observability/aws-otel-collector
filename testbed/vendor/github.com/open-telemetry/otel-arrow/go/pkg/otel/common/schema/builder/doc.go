/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

// Package builder decorates the Apache Arrow record/array builders with
// additional functionality to support the concept of adaptive schema and
// transformation nodes.
//
// Everytime that a new data added affects the schema (e.g. a field marked as
// optional becomes required or a dictionary field overflows its index type),
// the schema is updated and the system must be able to re-inject the data into
// the `RecordBuilderExt` in order to avoid data loss.

package builder
