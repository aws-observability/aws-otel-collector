// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consul // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/consul"

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/consul"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/consul/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "consul"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
	provider           consul.Provider
	logger             *zap.Logger
	resourceAttributes metadata.ResourceAttributesConfig
}

// NewDetector creates a new system metadata detector
func NewDetector(p processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	userCfg := dcfg.(Config)
	cfg := api.DefaultConfig()

	if userCfg.Address != "" {
		cfg.Address = userCfg.Address
	}
	if userCfg.Datacenter != "" {
		cfg.Datacenter = userCfg.Datacenter
	}
	if userCfg.Namespace != "" {
		cfg.Namespace = userCfg.Namespace
	}
	if userCfg.Token != "" {
		cfg.Token = string(userCfg.Token)
	}
	if userCfg.TokenFile != "" {
		cfg.Token = userCfg.TokenFile
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed creating consul client: %w", err)
	}

	provider := consul.NewProvider(client, userCfg.MetaLabels)
	return &Detector{provider: provider, logger: p.Logger, resourceAttributes: userCfg.ResourceAttributes}, nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	res := pcommon.NewResource()
	attrs := res.Attributes()

	metadata, err := d.provider.Metadata(ctx)
	if err != nil {
		return res, "", fmt.Errorf("failed to get consul metadata: %w", err)
	}

	for key, element := range metadata.HostMetadata {
		attrs.PutStr(key, element)
	}

	if d.resourceAttributes.HostName.Enabled {
		attrs.PutStr(conventions.AttributeHostName, metadata.Hostname)
	}
	if d.resourceAttributes.CloudRegion.Enabled {
		attrs.PutStr(conventions.AttributeCloudRegion, metadata.Datacenter)
	}
	if d.resourceAttributes.HostID.Enabled {
		attrs.PutStr(conventions.AttributeHostID, metadata.NodeID)
	}

	return res, conventions.SchemaURL, nil
}
