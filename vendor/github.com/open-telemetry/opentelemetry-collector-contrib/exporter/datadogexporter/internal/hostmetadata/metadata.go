// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package metadata is responsible for collecting host metadata from different providers
// such as EC2, ECS, AWS, etc and pushing it to Datadog.
package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/hostmetadata"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

<<<<<<< HEAD
=======
	"github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata/payload"
>>>>>>> main
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes"
	ec2Attributes "github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes/ec2"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes/gcp"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes/source"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/clientutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/hostmetadata/internal/ec2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/hostmetadata/internal/gohai"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/hostmetadata/internal/system"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/scrub"
)

<<<<<<< HEAD
// HostMetadata includes metadata about the host tags,
// host aliases and identifies the host as an OpenTelemetry host
type HostMetadata struct {
	// Meta includes metadata about the host.
	Meta *Meta `json:"meta"`

	// InternalHostname is the canonical hostname
	InternalHostname string `json:"internalHostname"`

	// Version is the OpenTelemetry Collector version.
	// This is used for correctly identifying the Collector in the backend,
	// and for telemetry purposes.
	Version string `json:"otel_version"`

	// Flavor is always set to "opentelemetry-collector".
	// It is used for telemetry purposes in the backend.
	Flavor string `json:"agent-flavor"`

	// Tags includes the host tags
	Tags *HostTags `json:"host-tags"`

	// Payload contains inventory of system information provided by gohai
	// this is embedded because of special serialization requirements
	// the field `gohai` is JSON-formatted string
	gohai.Payload

	// Processes contains the process payload devired by gohai
	// Because of legacy reasons this is called resources in datadog intake
	Processes *gohai.ProcessesPayload `json:"resources"`
}

// HostTags are the host tags.
// Currently only system (configuration) tags are considered.
type HostTags struct {
	// OTel are host tags set in the configuration
	OTel []string `json:"otel,omitempty"`

	// GCP are Google Cloud Platform tags
	GCP []string `json:"google cloud platform,omitempty"`
}

// Meta includes metadata about the host aliases
type Meta struct {
	// InstanceID is the EC2 instance id the Collector is running on, if available
	InstanceID string `json:"instance-id,omitempty"`

	// EC2Hostname is the hostname from the EC2 metadata API
	EC2Hostname string `json:"ec2-hostname,omitempty"`

	// Hostname is the canonical hostname
	Hostname string `json:"hostname"`

	// SocketHostname is the OS hostname
	SocketHostname string `json:"socket-hostname,omitempty"`

	// SocketFqdn is the FQDN hostname
	SocketFqdn string `json:"socket-fqdn,omitempty"`

	// HostAliases are other available host names
	HostAliases []string `json:"host_aliases,omitempty"`
}

// metadataFromAttributes gets metadata info from attributes following
// OpenTelemetry semantic conventions
func metadataFromAttributes(attrs pcommon.Map) *HostMetadata {
	hm := &HostMetadata{Meta: &Meta{}, Tags: &HostTags{}}
=======
// metadataFromAttributes gets metadata info from attributes following
// OpenTelemetry semantic conventions
func metadataFromAttributes(attrs pcommon.Map) payload.HostMetadata {
	hm := payload.HostMetadata{Meta: &payload.Meta{}, Tags: &payload.HostTags{}}
>>>>>>> main

	if src, ok := attributes.SourceFromAttrs(attrs); ok && src.Kind == source.HostnameKind {
		hm.InternalHostname = src.Identifier
		hm.Meta.Hostname = src.Identifier
	}

	// AWS EC2 resource metadata
	cloudProvider, ok := attrs.Get(conventions.AttributeCloudProvider)
	switch {
	case ok && cloudProvider.Str() == conventions.AttributeCloudProviderAWS:
		ec2HostInfo := ec2Attributes.HostInfoFromAttributes(attrs)
		hm.Meta.InstanceID = ec2HostInfo.InstanceID
		hm.Meta.EC2Hostname = ec2HostInfo.EC2Hostname
		hm.Tags.OTel = append(hm.Tags.OTel, ec2HostInfo.EC2Tags...)
	case ok && cloudProvider.Str() == conventions.AttributeCloudProviderGCP:
		gcpHostInfo := gcp.HostInfoFromAttrs(attrs)
		hm.Tags.GCP = gcpHostInfo.GCPTags
		hm.Meta.HostAliases = append(hm.Meta.HostAliases, gcpHostInfo.HostAliases...)
	}

	return hm
}

<<<<<<< HEAD
func fillHostMetadata(params exporter.CreateSettings, pcfg PusherConfig, p source.Provider, hm *HostMetadata) {
=======
func fillHostMetadata(params exporter.CreateSettings, pcfg PusherConfig, p source.Provider, hm *payload.HostMetadata) {
>>>>>>> main
	// Could not get hostname from attributes
	if hm.InternalHostname == "" {
		if src, err := p.Source(context.TODO()); err == nil && src.Kind == source.HostnameKind {
			hm.InternalHostname = src.Identifier
			hm.Meta.Hostname = src.Identifier
		}
	}

	// This information always gets filled in here
	// since it does not come from OTEL conventions
	hm.Flavor = params.BuildInfo.Command
	hm.Version = params.BuildInfo.Version
	hm.Tags.OTel = append(hm.Tags.OTel, pcfg.ConfigTags...)
	hm.Payload = gohai.NewPayload(params.Logger)
	hm.Processes = gohai.NewProcessesPayload(hm.Meta.Hostname, params.Logger)
	// EC2 data was not set from attributes
	if hm.Meta.EC2Hostname == "" {
		ec2HostInfo := ec2.GetHostInfo(params.Logger)
		hm.Meta.EC2Hostname = ec2HostInfo.EC2Hostname
		hm.Meta.InstanceID = ec2HostInfo.InstanceID
	}

	// System data was not set from attributes
	if hm.Meta.SocketHostname == "" {
		systemHostInfo := system.GetHostInfo(params.Logger)
		hm.Meta.SocketHostname = systemHostInfo.OS
		hm.Meta.SocketFqdn = systemHostInfo.FQDN
	}
}

<<<<<<< HEAD
func pushMetadata(pcfg PusherConfig, params exporter.CreateSettings, metadata *HostMetadata) error {
	if metadata.Meta.Hostname == "" {
		// if the hostname is empty, don't send metadata; we don't need it.
		params.Logger.Debug("Skipping host metadata since the hostname is empty")
		return nil
	}

	path := pcfg.MetricsEndpoint + "/intake"
	buf, _ := json.Marshal(metadata)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(buf))
	clientutil.SetDDHeaders(req.Header, params.BuildInfo, pcfg.APIKey)
	clientutil.SetExtraHeaders(req.Header, clientutil.JSONHeaders)
	client := clientutil.NewHTTPClient(pcfg.TimeoutSettings, pcfg.InsecureSkipVerify)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

=======
func (p *pusher) pushMetadata(hm payload.HostMetadata) error {
	path := p.pcfg.MetricsEndpoint + "/intake"
	buf, _ := json.Marshal(hm)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(buf))
	clientutil.SetDDHeaders(req.Header, p.params.BuildInfo, p.pcfg.APIKey)
	clientutil.SetExtraHeaders(req.Header, clientutil.JSONHeaders)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return err
	}
>>>>>>> main
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf(
<<<<<<< HEAD
			"'%s' error when sending metadata payload to %s",
=======
			"%q error when sending metadata payload to %s",
>>>>>>> main
			resp.Status,
			path,
		)
	}

	return nil
}

<<<<<<< HEAD
func pushMetadataWithRetry(retrier *clientutil.Retrier, params exporter.CreateSettings, pcfg PusherConfig, hostMetadata *HostMetadata) {
	params.Logger.Debug("Sending host metadata payload", zap.Any("payload", hostMetadata))

	_, err := retrier.DoWithRetries(context.Background(), func(context.Context) error {
		return pushMetadata(pcfg, params, hostMetadata)
	})

	if err != nil {
		params.Logger.Warn("Sending host metadata failed", zap.Error(err))
	} else {
		params.Logger.Info("Sent host metadata")
	}

}

// Pusher pushes host metadata payloads periodically to Datadog intake
func Pusher(ctx context.Context, params exporter.CreateSettings, pcfg PusherConfig, p source.Provider, attrs pcommon.Map) {
=======
func (p *pusher) Push(_ context.Context, hm payload.HostMetadata) error {
	if hm.Meta.Hostname == "" {
		// if the hostname is empty, don't send metadata; we don't need it.
		p.params.Logger.Debug("Skipping host metadata since the hostname is empty")
		return nil
	}

	p.params.Logger.Debug("Sending host metadata payload", zap.Any("payload", hm))

	_, err := p.retrier.DoWithRetries(context.Background(), func(context.Context) error {
		return p.pushMetadata(hm)
	})

	return err
}

var _ inframetadata.Pusher = (*pusher)(nil)

type pusher struct {
	params     exporter.CreateSettings
	pcfg       PusherConfig
	retrier    *clientutil.Retrier
	httpClient *http.Client
}

// NewPusher creates a new inframetadata.Pusher that pushes metadata payloads
func NewPusher(params exporter.CreateSettings, pcfg PusherConfig) inframetadata.Pusher {
	return &pusher{
		params:     params,
		pcfg:       pcfg,
		retrier:    clientutil.NewRetrier(params.Logger, pcfg.RetrySettings, scrub.NewScrubber()),
		httpClient: clientutil.NewHTTPClient(pcfg.TimeoutSettings, pcfg.InsecureSkipVerify),
	}
}

// RunPusher to push host metadata payloads from the host where the Collector is running periodically to Datadog intake.
// This function is blocking and it is meant to be run on a goroutine.
func RunPusher(ctx context.Context, params exporter.CreateSettings, pcfg PusherConfig, p source.Provider, attrs pcommon.Map, reporter *inframetadata.Reporter) {
>>>>>>> main
	// Push metadata every 30 minutes
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()
	defer params.Logger.Debug("Shut down host metadata routine")
<<<<<<< HEAD
	retrier := clientutil.NewRetrier(params.Logger, pcfg.RetrySettings, scrub.NewScrubber())
=======
>>>>>>> main

	// Get host metadata from resources and fill missing info using our exporter.
	// Currently we only retrieve it once but still send the same payload
	// every 30 minutes for consistency with the Datadog Agent behavior.
	//
	// All fields that are being filled in by our exporter
	// do not change over time. If this ever changes `hostMetadata`
	// *must* be deep copied before calling `fillHostMetadata`.
<<<<<<< HEAD
	hostMetadata := &HostMetadata{Meta: &Meta{}, Tags: &HostTags{}}
	if pcfg.UseResourceMetadata {
		hostMetadata = metadataFromAttributes(attrs)
	}
	fillHostMetadata(params, pcfg, p, hostMetadata)

	// Run one first time at startup
	pushMetadataWithRetry(retrier, params, pcfg, hostMetadata)
=======
	hostMetadata := payload.HostMetadata{Meta: &payload.Meta{}, Tags: &payload.HostTags{}}
	if pcfg.UseResourceMetadata {
		hostMetadata = metadataFromAttributes(attrs)
	}
	fillHostMetadata(params, pcfg, p, &hostMetadata)
	// Consume one first time
	if err := reporter.ConsumeHostMetadata(hostMetadata); err != nil {
		params.Logger.Warn("Failed to consume host metadata", zap.Any("payload", hostMetadata))
	}
>>>>>>> main

	for {
		select {
		case <-ctx.Done():
			return
<<<<<<< HEAD
		case <-ticker.C: // Send host metadata
			pushMetadataWithRetry(retrier, params, pcfg, hostMetadata)
=======
		case <-ticker.C:
			if err := reporter.ConsumeHostMetadata(hostMetadata); err != nil {
				params.Logger.Warn("Failed to consume host metadata", zap.Any("payload", hostMetadata))
			}
>>>>>>> main
		}
	}
}
