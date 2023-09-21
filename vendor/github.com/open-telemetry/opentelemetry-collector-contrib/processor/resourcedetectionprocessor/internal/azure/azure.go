// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "azure"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is an Azure metadata detector
type Detector struct {
<<<<<<< HEAD
	provider           azure.Provider
	logger             *zap.Logger
	resourceAttributes metadata.ResourceAttributesConfig
=======
	provider azure.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
>>>>>>> main
}

// NewDetector creates a new Azure metadata detector
func NewDetector(p processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &Detector{
<<<<<<< HEAD
		provider:           azure.NewProvider(),
		logger:             p.Logger,
		resourceAttributes: cfg.ResourceAttributes,
=======
		provider: azure.NewProvider(),
		logger:   p.Logger,
		rb:       metadata.NewResourceBuilder(cfg.ResourceAttributes),
>>>>>>> main
	}, nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
<<<<<<< HEAD
	res := pcommon.NewResource()
	attrs := res.Attributes()

=======
>>>>>>> main
	compute, err := d.provider.Metadata(ctx)
	if err != nil {
		d.logger.Debug("Azure detector metadata retrieval failed", zap.Error(err))
		// return an empty Resource and no error
<<<<<<< HEAD
		return res, "", nil
	}
	if d.resourceAttributes.CloudProvider.Enabled {
		attrs.PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAzure)
	}
	if d.resourceAttributes.CloudPlatform.Enabled {
		attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformAzureVM)
	}
	if d.resourceAttributes.HostName.Enabled {
		attrs.PutStr(conventions.AttributeHostName, compute.Name)
	}
	if d.resourceAttributes.CloudRegion.Enabled {
		attrs.PutStr(conventions.AttributeCloudRegion, compute.Location)
	}
	if d.resourceAttributes.HostID.Enabled {
		attrs.PutStr(conventions.AttributeHostID, compute.VMID)
	}
	if d.resourceAttributes.CloudAccountID.Enabled {
		attrs.PutStr(conventions.AttributeCloudAccountID, compute.SubscriptionID)
	}
	// Also save compute.Name in "azure.vm.name" as host.id (AttributeHostName) is
	// used by system detector.
	attrs.PutStr("azure.vm.name", compute.Name)
	attrs.PutStr("azure.vm.size", compute.VMSize)
	attrs.PutStr("azure.vm.scaleset.name", compute.VMScaleSetName)
	attrs.PutStr("azure.resourcegroup.name", compute.ResourceGroupName)

	return res, conventions.SchemaURL, nil
=======
		return pcommon.NewResource(), "", nil
	}

	d.rb.SetCloudProvider(conventions.AttributeCloudProviderAzure)
	d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformAzureVM)
	d.rb.SetHostName(compute.Name)
	d.rb.SetCloudRegion(compute.Location)
	d.rb.SetHostID(compute.VMID)
	d.rb.SetCloudAccountID(compute.SubscriptionID)

	// Also save compute.Name in "azure.vm.name" as host.id (AttributeHostName) is
	// used by system detector.
	d.rb.SetAzureVMName(compute.Name)
	d.rb.SetAzureVMSize(compute.VMSize)
	d.rb.SetAzureVMScalesetName(compute.VMScaleSetName)
	d.rb.SetAzureResourcegroupName(compute.ResourceGroupName)

	return d.rb.Emit(), conventions.SchemaURL, nil
>>>>>>> main
}
