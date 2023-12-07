// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package hostmap

import (
	"fmt"
	"sync"

	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
	"go.uber.org/multierr"

	"github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata/gohai"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata/payload"
)

// HostMap maps from hostnames to host metadata payloads.
type HostMap struct {
	// mu is the mutex for the host map and updater.
	mu sync.Mutex
	// hosts map
	hosts map[string]payload.HostMetadata
}

// New creates a new HostMap.
func New() *HostMap {
	return &HostMap{
		hosts: make(map[string]payload.HostMetadata),
	}
}

// strField gets a field as string from a resource attribute map.
// It can handle fields of type "Str" and "Int".
// It returns:
// - The field's value, if available
// - Whether the field was present in the map
// - Any errors found in the process
func strField(m pcommon.Map, key string) (string, bool, error) {
	val, ok := m.Get(key)
	if !ok {
		// Field not available, don't update but don't fail either
		return "", false, nil
	}

	var value string
	switch val.Type() {
	case pcommon.ValueTypeStr:
		value = val.Str()
	case pcommon.ValueTypeInt:
		value = val.AsString()
	default:
		return "", false, fmt.Errorf("%q has type %q, expected type \"Str\" instead", key, val.Type())
	}

	return value, true, nil
}

// isAWS checks if a resource attribute map
// is coming from an AWS VM.
func isAWS(m pcommon.Map) (bool, error) {
	cloudProvider, ok, err := strField(m, conventions.AttributeCloudProvider)
	if err != nil {
		return false, err
	} else if !ok {
		// no cloud provider field
		return false, nil
	}
	return cloudProvider == conventions.AttributeCloudProviderAWS, nil
}

// instanceID gets the AWS EC2 instance ID from a resource attribute map.
// It returns:
// - The EC2 instance id if available
// - Whether the instance id was found
// - Any errors found retrieving the ID
func instanceID(m pcommon.Map) (string, bool, error) {
	if onAWS, err := isAWS(m); err != nil || !onAWS {
		return "", onAWS, err
	}
	return strField(m, conventions.AttributeHostID)
}

// ec2Hostname gets the AWS EC2 OS hostname from a resource attribute map.
// It returns:
// - The EC2 OS hostname if available
// - Whether the EC2 OS hostname was found
// - Any errors found retrieving the ID
func ec2Hostname(m pcommon.Map) (string, bool, error) {
	if onAWS, err := isAWS(m); err != nil || !onAWS {
		return "", onAWS, err
	}
	return strField(m, conventions.AttributeHostName)
}

// Set a hardcoded host metadata payload.
func (m *HostMap) Set(md payload.HostMetadata) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hosts[md.Meta.Hostname] = md
	return nil
}

// Update the information about a given host by providing a resource.
// The function reports:
//   - Whether the information about the `host` has changed
//   - The host metadata payload stored
//   - Any non-fatal errors that may have occurred during the update
//
// Non-fatal errors are local to the specific field where they happened
// and do not change the other fields. If when filling a field a non-fatal
// error is raised, the error will be reported, the field will be left
// empty and further fields will still be filled.
//
// The order in which resource attributes are read does not affect the final
// host metadata payload, even if non-fatal errors are raised during execution.
func (m *HostMap) Update(host string, res pcommon.Resource) (changed bool, md payload.HostMetadata, err error) {
	md = payload.HostMetadata{
		Flavor:  "otelcol-contrib",
		Meta:    &payload.Meta{},
		Tags:    &payload.HostTags{},
		Payload: gohai.NewEmpty(),
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	var found bool
	if old, ok := m.hosts[host]; ok {
		found = true
		md = old
	}

	md.InternalHostname = host
	md.Meta.Hostname = host

	// InstanceID field
	if iid, ok, err2 := instanceID(res.Attributes()); err2 != nil {
		err = multierr.Append(err, err2)
	} else if ok {
		old := md.Meta.InstanceID
		changed = changed || old != iid
		md.Meta.InstanceID = iid
	}

	// EC2Hostname field
	if ec2Host, ok, err2 := ec2Hostname(res.Attributes()); err2 != nil {
		err = multierr.Append(err, err2)
	} else if ok {
		old := md.Meta.EC2Hostname
		changed = changed || old != ec2Host
		md.Meta.EC2Hostname = ec2Host
	}

	// Gohai - Platform
	md.Platform()["hostname"] = host
	for field, attribute := range platformAttributesMap {
		strVal, ok, fieldErr := strField(res.Attributes(), attribute)
		if fieldErr != nil {
			err = multierr.Append(err, fieldErr)
		} else if ok {
			old := md.Platform()[field]
			changed = changed || old != strVal
			md.Platform()[field] = strVal
		}
	}

	// Gohai - CPU
	for field, attribute := range cpuAttributesMap {
		strVal, ok, fieldErr := strField(res.Attributes(), attribute)
		if fieldErr != nil {
			err = multierr.Append(err, fieldErr)
		} else if ok {
			old := md.CPU()[field]
			changed = changed || old != strVal
			md.CPU()[field] = strVal
		}
	}

	m.hosts[host] = md
	changed = changed && found
	return
}

// Flush all the host metadata payloads and clear them from the HostMap.
func (m *HostMap) Flush() map[string]payload.HostMetadata {
	m.mu.Lock()
	defer m.mu.Unlock()
	hosts := m.hosts
	m.hosts = make(map[string]payload.HostMetadata)
	return hosts
}
