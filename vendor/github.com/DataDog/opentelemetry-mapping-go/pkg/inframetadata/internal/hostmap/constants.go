// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package hostmap

import (
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

// Custom attributes not in the OpenTelemetry specification
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
