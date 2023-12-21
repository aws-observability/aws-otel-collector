// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package hostmap

import (
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

// Platform related OpenTelemetry Semantic Conventions for resource attributes.
// These are NOT in the specification and will be submitted for approval.
const (
	attributeKernelName    = "os.kernel.name"
	attributeKernelRelease = "os.kernel.release"
	attributeKernelVersion = "os.kernel.version"
)

// This set of constants represent fields in the Gohai payload's Platform field.
const (
	fieldPlatformHostname         = "hostname"
	fieldPlatformOS               = "os"
	fieldPlatformProcessor        = "processor"
	fieldPlatformMachine          = "machine"
	fieldPlatformHardwarePlatform = "hardware_platform"
	fieldPlatformKernelName       = "kernel_name"
	fieldPlatformKernelRelease    = "kernel_release"
	fieldPlatformKernelVersion    = "kernel_version"
)

// platformAttributesMap defines the mapping between Gohai fieldPlatform fields
// and resource attribute names (semantic conventions or not).
var platformAttributesMap map[string]string = map[string]string{
	fieldPlatformOS:               conventions.AttributeOSDescription,
	fieldPlatformProcessor:        conventions.AttributeHostArch,
	fieldPlatformMachine:          conventions.AttributeHostArch,
	fieldPlatformHardwarePlatform: conventions.AttributeHostArch,
	fieldPlatformKernelName:       attributeKernelName,
	fieldPlatformKernelRelease:    attributeKernelRelease,
	fieldPlatformKernelVersion:    attributeKernelVersion,
}

// CPU related OpenTelemetry Semantic Conventions for resource attributes.
// TODO: Replace by conventions constants once available.
const (
	attributeHostCPUVendorID    = "host.cpu.vendor.id"
	attributeHostCPUModelName   = "host.cpu.model.name"
	attributeHostCPUFamily      = "host.cpu.family"
	attributeHostCPUModelID     = "host.cpu.model.id"
	attributeHostCPUStepping    = "host.cpu.stepping"
	attributeHostCPUCacheL2Size = "host.cpu.cache.l2.size"
)

// This set of constants represent fields in the Gohai payload's CPU field.
const (
	fieldCPUVendorID  = "vendor_id"
	fieldCPUModelName = "model_name"
	fieldCPUCacheSize = "cache_size"
	fieldCPUFamily    = "family"
	fieldCPUModel     = "model"
	fieldCPUStepping  = "stepping"
)

// cpuAttributesMap defines the mapping between Gohai fieldCPU fields
// and resource attribute names (semantic conventions or not).
var cpuAttributesMap map[string]string = map[string]string{
	fieldCPUVendorID:  attributeHostCPUVendorID,
	fieldCPUModelName: attributeHostCPUModelName,
	fieldCPUCacheSize: attributeHostCPUCacheL2Size,
	fieldCPUFamily:    attributeHostCPUFamily,
	fieldCPUModel:     attributeHostCPUModelID,
	fieldCPUStepping:  attributeHostCPUStepping,
}
