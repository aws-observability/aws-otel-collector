// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package hostmap

import (
	"fmt"
	"sort"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/collector/semconv/v1.21.0"
	"go.uber.org/multierr"
)

const hostTagPrefix = "datadog.host.tag."

var hostTagMapping = map[string]string{
	conventions.AttributeDeploymentEnvironment: "env",
	conventions.AttributeK8SClusterName:        "cluster_name",
	conventions.AttributeCloudProvider:         "cloud_provider",
	conventions.AttributeCloudRegion:           "region",
	conventions.AttributeCloudAvailabilityZone: "zone",
}

// assertStringValue returns the string value of the given value, or an error if the value is not a string.
func assertStringValue(name string, v pcommon.Value) (string, error) {
	if v.Type() != pcommon.ValueTypeStr {
		return "", mismatchErr(name, v.Type(), pcommon.ValueTypeStr)
	}
	return v.Str(), nil
}

// getHostTags returns a slice of tags from the given map.
func getHostTags(m pcommon.Map) ([]string, error) {
	var tags []string
	var err error
	m.Range(func(k string, v pcommon.Value) bool {
		if strings.HasPrefix(k, hostTagPrefix) { // User-defined tags
			if str, err2 := assertStringValue(k, v); err2 != nil {
				err = multierr.Append(err, err2)
			} else if str == "" {
				err = multierr.Append(err, fmt.Errorf("attribute %q has empty string value, expected non-empty string", k))
			} else {
				tags = append(tags, k[len(hostTagPrefix):]+":"+str)
			}
		} else if mappedKey, ok := hostTagMapping[k]; ok { // Well-known tags
			if str, err2 := assertStringValue(k, v); err2 != nil {
				err = multierr.Append(err, err2)
			} else {
				tags = append(tags, mappedKey+":"+str)
			}
		}
		return true
	})

	// Allow for comparison of tags
	sort.Strings(tags)
	return tags, err
}
