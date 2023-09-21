// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/gcp"

import (
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> main

	"cloud.google.com/go/compute/metadata"
	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	localMetadata "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/gcp/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "gcp"
)

// NewDetector returns a detector which can detect resource attributes on:
// * Google Compute Engine (GCE).
// * Google Kubernetes Engine (GKE).
// * Google App Engine (GAE).
// * Cloud Run.
// * Cloud Functions.
func NewDetector(set processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &detector{
<<<<<<< HEAD
		logger:             set.Logger,
		detector:           gcp.NewDetector(),
		resourceAttributes: cfg.ResourceAttributes,
=======
		logger:   set.Logger,
		detector: gcp.NewDetector(),
		rb:       localMetadata.NewResourceBuilder(cfg.ResourceAttributes),
>>>>>>> main
	}, nil
}

type detector struct {
<<<<<<< HEAD
	logger             *zap.Logger
	detector           gcpDetector
	resourceAttributes localMetadata.ResourceAttributesConfig
}

func (d *detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	res := pcommon.NewResource()
	if !metadata.OnGCE() {
		return res, "", nil
	}
	b := &resourceBuilder{logger: d.logger, attrs: res.Attributes()}
	if d.resourceAttributes.CloudProvider.Enabled {
		b.attrs.PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderGCP)
	}
	if d.resourceAttributes.CloudAccountID.Enabled {
		b.add(conventions.AttributeCloudAccountID, d.detector.ProjectID)
	}

	switch d.detector.CloudPlatform() {
	case gcp.GKE:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPKubernetesEngine)
		}
		if d.resourceAttributes.CloudAvailabilityZone.Enabled {
			b.addZoneOrRegion(d.detector.GKEAvailabilityZoneOrRegion, d.resourceAttributes)
		}
		if d.resourceAttributes.K8sClusterName.Enabled {
			b.add(conventions.AttributeK8SClusterName, d.detector.GKEClusterName)
		}
		if d.resourceAttributes.HostID.Enabled {
			b.add(conventions.AttributeHostID, d.detector.GKEHostID)
		}
		// GCEHostname is fallible on GKE, since it's not available when using workload identity.
		if d.resourceAttributes.HostName.Enabled {
			b.addFallible(conventions.AttributeHostName, d.detector.GCEHostName)
		}
	case gcp.CloudRun:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPCloudRun)
		}
		if d.resourceAttributes.FaasName.Enabled {
			b.add(conventions.AttributeFaaSName, d.detector.FaaSName)
		}
		if d.resourceAttributes.FaasVersion.Enabled {
			b.add(conventions.AttributeFaaSVersion, d.detector.FaaSVersion)
		}
		if d.resourceAttributes.FaasID.Enabled {
			b.add(conventions.AttributeFaaSID, d.detector.FaaSID)
		}
		if d.resourceAttributes.CloudRegion.Enabled {
			b.add(conventions.AttributeCloudRegion, d.detector.FaaSCloudRegion)
		}
	case gcp.CloudFunctions:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPCloudFunctions)
		}
		if d.resourceAttributes.FaasName.Enabled {
			b.add(conventions.AttributeFaaSName, d.detector.FaaSName)
		}
		if d.resourceAttributes.FaasVersion.Enabled {
			b.add(conventions.AttributeFaaSVersion, d.detector.FaaSVersion)
		}
		if d.resourceAttributes.FaasID.Enabled {
			b.add(conventions.AttributeFaaSID, d.detector.FaaSID)
		}
		if d.resourceAttributes.CloudRegion.Enabled {
			b.add(conventions.AttributeCloudRegion, d.detector.FaaSCloudRegion)
		}
	case gcp.AppEngineFlex:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPAppEngine)
		}
		b.addZoneAndRegion(d.detector.AppEngineFlexAvailabilityZoneAndRegion, d.resourceAttributes)
		if d.resourceAttributes.FaasName.Enabled {
			b.add(conventions.AttributeFaaSName, d.detector.AppEngineServiceName)
		}
		if d.resourceAttributes.FaasVersion.Enabled {
			b.add(conventions.AttributeFaaSVersion, d.detector.AppEngineServiceVersion)
		}
		if d.resourceAttributes.FaasID.Enabled {
			b.add(conventions.AttributeFaaSID, d.detector.AppEngineServiceInstance)
		}
	case gcp.AppEngineStandard:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPAppEngine)
		}
		if d.resourceAttributes.FaasName.Enabled {
			b.add(conventions.AttributeFaaSName, d.detector.AppEngineServiceName)
		}
		if d.resourceAttributes.FaasVersion.Enabled {
			b.add(conventions.AttributeFaaSVersion, d.detector.AppEngineServiceVersion)
		}
		if d.resourceAttributes.FaasID.Enabled {
			b.add(conventions.AttributeFaaSID, d.detector.AppEngineServiceInstance)
		}
		if d.resourceAttributes.CloudAvailabilityZone.Enabled {
			b.add(conventions.AttributeCloudAvailabilityZone, d.detector.AppEngineStandardAvailabilityZone)
		}
		if d.resourceAttributes.CloudRegion.Enabled {
			b.add(conventions.AttributeCloudRegion, d.detector.AppEngineStandardCloudRegion)
		}
	case gcp.GCE:
		if d.resourceAttributes.CloudPlatform.Enabled {
			b.attrs.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformGCPComputeEngine)
		}
		b.addZoneAndRegion(d.detector.GCEAvailabilityZoneAndRegion, d.resourceAttributes)
		if d.resourceAttributes.HostType.Enabled {
			b.add(conventions.AttributeHostType, d.detector.GCEHostType)
		}
		if d.resourceAttributes.HostID.Enabled {
			b.add(conventions.AttributeHostID, d.detector.GCEHostID)
		}
		if d.resourceAttributes.HostName.Enabled {
			b.add(conventions.AttributeHostName, d.detector.GCEHostName)
		}
	default:
		// We don't support this platform yet, so just return with what we have
	}
	return res, conventions.SchemaURL, multierr.Combine(b.errs...)
}

// resourceBuilder simplifies constructing resources using GCP detection
// library functions.
type resourceBuilder struct {
	logger *zap.Logger
	errs   []error
	attrs  pcommon.Map
}

func (r *resourceBuilder) add(key string, detect func() (string, error)) {
	v, err := detect()
	if err != nil {
		r.errs = append(r.errs, err)
		return
	}
	r.attrs.PutStr(key, v)
}

// addFallible adds a detect function whose failures should be ignored
func (r *resourceBuilder) addFallible(key string, detect func() (string, error)) {
	v, err := detect()
	if err != nil {
		r.logger.Info("Fallible detector failed. This attribute will not be available.", zap.String("key", key), zap.Error(err))
		return
	}
	r.attrs.PutStr(key, v)
}

// zoneAndRegion functions are expected to return zone, region, err.
func (r *resourceBuilder) addZoneAndRegion(detect func() (string, string, error), resourceAttributes localMetadata.ResourceAttributesConfig) {
	zone, region, err := detect()
	if err != nil {
		r.errs = append(r.errs, err)
		return
	}
	if resourceAttributes.CloudAvailabilityZone.Enabled {
		r.attrs.PutStr(conventions.AttributeCloudAvailabilityZone, zone)
	}
	if resourceAttributes.CloudRegion.Enabled {
		r.attrs.PutStr(conventions.AttributeCloudRegion, region)
	}
}

func (r *resourceBuilder) addZoneOrRegion(detect func() (string, gcp.LocationType, error), resourceAttributes localMetadata.ResourceAttributesConfig) {
	v, locType, err := detect()
	if err != nil {
		r.errs = append(r.errs, err)
		return
	}

	switch locType {
	case gcp.Zone:
		if resourceAttributes.CloudAvailabilityZone.Enabled {
			r.attrs.PutStr(conventions.AttributeCloudAvailabilityZone, v)
		}
	case gcp.Region:
		if resourceAttributes.CloudRegion.Enabled {
			r.attrs.PutStr(conventions.AttributeCloudRegion, v)
		}
	default:
		r.errs = append(r.errs, fmt.Errorf("location must be zone or region. Got %v", locType))
	}
=======
	logger   *zap.Logger
	detector gcpDetector
	rb       *localMetadata.ResourceBuilder
}

func (d *detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	if !metadata.OnGCE() {
		return pcommon.NewResource(), "", nil
	}

	d.rb.SetCloudProvider(conventions.AttributeCloudProviderGCP)
	errs := d.rb.SetFromCallable(d.rb.SetCloudAccountID, d.detector.ProjectID)

	switch d.detector.CloudPlatform() {
	case gcp.GKE:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPKubernetesEngine)
		errs = multierr.Combine(errs,
			d.rb.SetZoneOrRegion(d.detector.GKEAvailabilityZoneOrRegion),
			d.rb.SetFromCallable(d.rb.SetK8sClusterName, d.detector.GKEClusterName),
			d.rb.SetFromCallable(d.rb.SetHostID, d.detector.GKEHostID),
		)
		// GCEHostname is fallible on GKE, since it's not available when using workload identity.
		if v, err := d.detector.GCEHostName(); err == nil {
			d.rb.SetHostName(v)
		} else {
			d.logger.Info("Fallible detector failed. This attribute will not be available.",
				zap.String("key", conventions.AttributeHostName), zap.Error(err))
		}
	case gcp.CloudRun:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPCloudRun)
		errs = multierr.Combine(errs,
			d.rb.SetFromCallable(d.rb.SetFaasName, d.detector.FaaSName),
			d.rb.SetFromCallable(d.rb.SetFaasVersion, d.detector.FaaSVersion),
			d.rb.SetFromCallable(d.rb.SetFaasID, d.detector.FaaSID),
			d.rb.SetFromCallable(d.rb.SetCloudRegion, d.detector.FaaSCloudRegion),
		)
	case gcp.CloudRunJob:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPCloudRun)
		errs = multierr.Combine(errs,
			d.rb.SetFromCallable(d.rb.SetFaasName, d.detector.FaaSName),
			d.rb.SetFromCallable(d.rb.SetCloudRegion, d.detector.FaaSCloudRegion),
			d.rb.SetFromCallable(d.rb.SetFaasID, d.detector.FaaSID),
			d.rb.SetFromCallable(d.rb.SetGcpCloudRunJobExecution, d.detector.CloudRunJobExecution),
			d.rb.SetFromCallable(d.rb.SetGcpCloudRunJobTaskIndex, d.detector.CloudRunJobTaskIndex),
		)
	case gcp.CloudFunctions:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPCloudFunctions)
		errs = multierr.Combine(errs,
			d.rb.SetFromCallable(d.rb.SetFaasName, d.detector.FaaSName),
			d.rb.SetFromCallable(d.rb.SetFaasVersion, d.detector.FaaSVersion),
			d.rb.SetFromCallable(d.rb.SetFaasID, d.detector.FaaSID),
			d.rb.SetFromCallable(d.rb.SetCloudRegion, d.detector.FaaSCloudRegion),
		)
	case gcp.AppEngineFlex:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPAppEngine)
		errs = multierr.Combine(errs,
			d.rb.SetZoneAndRegion(d.detector.AppEngineFlexAvailabilityZoneAndRegion),
			d.rb.SetFromCallable(d.rb.SetFaasName, d.detector.AppEngineServiceName),
			d.rb.SetFromCallable(d.rb.SetFaasVersion, d.detector.AppEngineServiceVersion),
			d.rb.SetFromCallable(d.rb.SetFaasID, d.detector.AppEngineServiceInstance),
		)
	case gcp.AppEngineStandard:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPAppEngine)
		errs = multierr.Combine(errs,
			d.rb.SetFromCallable(d.rb.SetFaasName, d.detector.AppEngineServiceName),
			d.rb.SetFromCallable(d.rb.SetFaasVersion, d.detector.AppEngineServiceVersion),
			d.rb.SetFromCallable(d.rb.SetFaasID, d.detector.AppEngineServiceInstance),
			d.rb.SetFromCallable(d.rb.SetCloudAvailabilityZone, d.detector.AppEngineStandardAvailabilityZone),
			d.rb.SetFromCallable(d.rb.SetCloudRegion, d.detector.AppEngineStandardCloudRegion),
		)
	case gcp.GCE:
		d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformGCPComputeEngine)
		errs = multierr.Combine(errs,
			d.rb.SetZoneAndRegion(d.detector.GCEAvailabilityZoneAndRegion),
			d.rb.SetFromCallable(d.rb.SetHostType, d.detector.GCEHostType),
			d.rb.SetFromCallable(d.rb.SetHostID, d.detector.GCEHostID),
			d.rb.SetFromCallable(d.rb.SetHostName, d.detector.GCEHostName),
			d.rb.SetFromCallable(d.rb.SetGcpGceInstanceHostname, d.detector.GCEInstanceHostname),
			d.rb.SetFromCallable(d.rb.SetGcpGceInstanceName, d.detector.GCEInstanceName),
		)
	default:
		// We don't support this platform yet, so just return with what we have
	}
	return d.rb.Emit(), conventions.SchemaURL, errs
>>>>>>> main
}
