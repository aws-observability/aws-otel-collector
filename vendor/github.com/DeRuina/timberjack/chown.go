//go:build !linux
// +build !linux

// Stub chown implementation for non-Linux systems.
// This file is excluded on Linux, where chown behavior is handled natively.

package timberjack

import (
	"os"
)

var chown = func(_ string, _ os.FileInfo) error {
	return nil
}
