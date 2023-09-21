// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/docker"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/docker"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/docker/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "docker"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
<<<<<<< HEAD
	provider           docker.Provider
	logger             *zap.Logger
	resourceAttributes metadata.ResourceAttributesConfig
}

// NewDetector creates a new system metadata detector
func NewDetector(p processor.CreateSettings, _ internal.DetectorConfig) (internal.Detector, error) {
=======
	provider docker.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
}

// NewDetector creates a new system metadata detector
func NewDetector(p processor.CreateSettings, cfg internal.DetectorConfig) (internal.Detector, error) {
>>>>>>> main
	dockerProvider, err := docker.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed creating detector: %w", err)
	}

<<<<<<< HEAD
	return &Detector{provider: dockerProvider, logger: p.Logger}, nil
=======
	return &Detector{
		provider: dockerProvider,
		logger:   p.Logger,
		rb:       metadata.NewResourceBuilder(cfg.(Config).ResourceAttributes),
	}, nil
>>>>>>> main
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
<<<<<<< HEAD
	res := pcommon.NewResource()
	attrs := res.Attributes()

	osType, err := d.provider.OSType(ctx)
	if err != nil {
		return res, "", fmt.Errorf("failed getting OS type: %w", err)
=======
	osType, err := d.provider.OSType(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting OS type: %w", err)
>>>>>>> main
	}

	hostname, err := d.provider.Hostname(ctx)
	if err != nil {
<<<<<<< HEAD
		return res, "", fmt.Errorf("failed getting OS hostname: %w", err)
	}

	if d.resourceAttributes.HostName.Enabled {
		attrs.PutStr(conventions.AttributeHostName, hostname)
	}
	if d.resourceAttributes.OsType.Enabled {
		attrs.PutStr(conventions.AttributeOSType, osType)
	}

	return res, conventions.SchemaURL, nil
=======
		return pcommon.NewResource(), "", fmt.Errorf("failed getting OS hostname: %w", err)
	}

	d.rb.SetHostName(hostname)
	d.rb.SetOsType(osType)

	return d.rb.Emit(), conventions.SchemaURL, nil
>>>>>>> main
}
