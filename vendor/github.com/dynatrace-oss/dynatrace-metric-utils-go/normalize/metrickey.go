// Copyright 2021 Dynatrace LLC
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package normalize

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	reMkIdentifierFirstSectionStart = regexp.MustCompile("^[^a-zA-Z_]+")
	reMkIdentifierSectionStart      = regexp.MustCompile("^[^a-zA-Z0-9_]+")
	reMkInvalidCharacters           = regexp.MustCompile("[^a-zA-Z0-9_-]+")
)

const (
	metricKeyMaxLength = 250
)

// MetricKey creates a valid metric key from any string passed to this function
// or returns an error if the resulting key is invalid.
func MetricKey(key string) (string, error) {
	// trim down long keys
	if len(key) > metricKeyMaxLength {
		key = key[:metricKeyMaxLength]
	}

	var sb strings.Builder
	splitKey := strings.Split(key, ".")

	for i, keySection := range splitKey {
		if i == 0 {
			// key is invalid if the first section is empty
			if keySection == "" {
				return "", fmt.Errorf("first key section is empty (%s)", key)
			}
			// as all invalid characters are replaced with an underscore, this can never be empty
			sb.WriteString(normalizeMetricKeyFirstSection(keySection))
		} else {
			// other key sections that are empty are ignored
			if keySection != "" {
				sb.WriteString(".")
				// key sections that are not empty before normalizing are always non-empty after normalizing.
				sb.WriteString(normalizeMetricKeyLaterSection(keySection))
			}
		}
	}

	// can just return the builder.String here, the key is definitely valid, otherwise the error above would have been returned
	return sb.String(), nil
}

// normalizeMetricKeyCommon is used by both of the other internal normalize functions.
// It replaces trailing and enclosed invalid characters with an underscore
func normalizeMetricKeySectionCommon(section string) string {
	// replace intermediate and trailing ranges of invalid characters with a single underscore.
	section = reMkInvalidCharacters.ReplaceAllString(section, "_")

	return section
}

// normalizeMetricKeyLaterSection is used for all sections except the first
func normalizeMetricKeyLaterSection(section string) string {
	// replace leading invalid characters
	section = reMkIdentifierSectionStart.ReplaceAllString(section, "_")
	return normalizeMetricKeySectionCommon(section)
}

// normalizeMetricKeyFirstSection is only used for the first section of the metric key,
// since the requirements are slightly different from later key sections.
func normalizeMetricKeyFirstSection(section string) string {
	// replace leading invalid chars for first section
	section = reMkIdentifierFirstSectionStart.ReplaceAllString(section, "_")
	return normalizeMetricKeySectionCommon(section)
}
