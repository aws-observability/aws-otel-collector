// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system"

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/system"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "system"
)

var hostnameSourcesMap = map[string]func(*Detector) (string, error){
	"os":     getHostname,
	"dns":    getFQDN,
	"cname":  lookupCNAME,
	"lookup": reverseLookupHost,
}

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
<<<<<<< HEAD
	provider           system.Provider
	logger             *zap.Logger
	hostnameSources    []string
	resourceAttributes metadata.ResourceAttributesConfig
=======
	provider system.Provider
	logger   *zap.Logger
	cfg      Config
	rb       *metadata.ResourceBuilder
>>>>>>> main
}

// NewDetector creates a new system metadata detector
func NewDetector(p processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	if len(cfg.HostnameSources) == 0 {
		cfg.HostnameSources = []string{"dns", "os"}
	}

<<<<<<< HEAD
	return &Detector{provider: system.NewProvider(), logger: p.Logger, hostnameSources: cfg.HostnameSources, resourceAttributes: cfg.ResourceAttributes}, nil
=======
	return &Detector{
		provider: system.NewProvider(),
		logger:   p.Logger,
		cfg:      cfg,
		rb:       metadata.NewResourceBuilder(cfg.ResourceAttributes),
	}, nil
>>>>>>> main
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	var hostname string

<<<<<<< HEAD
	res := pcommon.NewResource()
	attrs := res.Attributes()

	osType, err := d.provider.OSType()
	if err != nil {
		return res, "", fmt.Errorf("failed getting OS type: %w", err)
	}

	hostID, err := d.provider.HostID(ctx)
	if err != nil {
		return res, "", fmt.Errorf("failed getting host ID: %w", err)
	}

	for _, source := range d.hostnameSources {
		getHostFromSource := hostnameSourcesMap[source]
		hostname, err = getHostFromSource(d)
		if err == nil {
			if d.resourceAttributes.HostName.Enabled {
				attrs.PutStr(conventions.AttributeHostName, hostname)
			}
			if d.resourceAttributes.OsType.Enabled {
				attrs.PutStr(conventions.AttributeOSType, osType)
			}
			if d.resourceAttributes.HostID.Enabled {
				attrs.PutStr(conventions.AttributeHostID, hostID)
			}

			return res, conventions.SchemaURL, nil
=======
	osType, err := d.provider.OSType()
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting OS type: %w", err)
	}

	hostArch, err := d.provider.HostArch()
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting host architecture: %w", err)
	}

	osDescription, err := d.provider.OSDescription(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting OS description: %w", err)
	}

	for _, source := range d.cfg.HostnameSources {
		getHostFromSource := hostnameSourcesMap[source]
		hostname, err = getHostFromSource(d)
		if err == nil {
			d.rb.SetHostName(hostname)
			d.rb.SetOsType(osType)
			if d.cfg.ResourceAttributes.HostID.Enabled {
				if hostID, hostIDErr := d.provider.HostID(ctx); hostIDErr == nil {
					d.rb.SetHostID(hostID)
				} else {
					d.logger.Warn("failed to get host ID", zap.Error(hostIDErr))
				}
			}
			d.rb.SetHostArch(hostArch)
			d.rb.SetOsDescription(osDescription)
			return d.rb.Emit(), conventions.SchemaURL, nil
>>>>>>> main
		}
		d.logger.Debug(err.Error())
	}

<<<<<<< HEAD
	return res, "", errors.New("all hostname sources failed to get hostname")
=======
	return pcommon.NewResource(), "", errors.New("all hostname sources failed to get hostname")
>>>>>>> main
}

// getHostname returns OS hostname
func getHostname(d *Detector) (string, error) {
	hostname, err := d.provider.Hostname()
	if err != nil {
		return "", fmt.Errorf("failed getting OS hostname: %w", err)
	}
	return hostname, nil
}

// getFQDN returns FQDN of the host
func getFQDN(d *Detector) (string, error) {
	hostname, err := d.provider.FQDN()
	if err != nil {
		return "", fmt.Errorf("getFQDN failed getting FQDN: %w", err)
	}
	return hostname, nil
}

func lookupCNAME(d *Detector) (string, error) {
	cname, err := d.provider.LookupCNAME()
	if err != nil {
		return "", fmt.Errorf("lookupCNAME failed to get CNAME: %w", err)
	}
	return cname, nil
}

func reverseLookupHost(d *Detector) (string, error) {
	hostname, err := d.provider.ReverseLookupHost()
	if err != nil {
		return "", fmt.Errorf("reverseLookupHost failed to lookup host: %w", err)
	}
	return hostname, nil
}
