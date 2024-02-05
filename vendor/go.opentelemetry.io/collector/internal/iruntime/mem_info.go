// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package iruntime // import "go.opentelemetry.io/collector/internal/iruntime"

import (
	"github.com/shirou/gopsutil/v3/mem"
)

// readMemInfo returns the total memory
// supports in linux, darwin and windows
func readMemInfo() (uint64, error) {
	vmStat, err := mem.VirtualMemory()
	return vmStat.Total, err
}
