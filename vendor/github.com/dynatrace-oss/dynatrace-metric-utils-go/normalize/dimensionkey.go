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
	"errors"
	"regexp"
	"strings"
)

var (
	reDkSectionStart      = regexp.MustCompile("^[^a-z_]+")
	reDkInvalidCharacters = regexp.MustCompile("[^a-z0-9_:-]+")
)

const (
	dimensionKeyMaxLength = 100
)

// DimensionKey returns a sanitized dimension key that is valid for metrics ingestion.
func DimensionKey(key string) (string, error) {
	if len(key) > dimensionKeyMaxLength {
		key = key[:dimensionKeyMaxLength]
	}

	var sb strings.Builder
	splitKey := strings.Split(key, ".")

	firstSection := true

	for _, keySection := range splitKey {
		if keySection != "" {
			if !firstSection {
				sb.WriteString(".")
			} else {
				firstSection = false
			}
			sb.WriteString(normalizeDimensionKeySection(keySection))
		}
	}

	normalizedKey := sb.String()
	if normalizedKey == "" {
		return "", errors.New("normalized key does not contain any characters")
	}
	return normalizedKey, nil
}

func normalizeDimensionKeySection(section string) string {
	section = strings.ToLower(section)
	section = reDkSectionStart.ReplaceAllString(section, "_")
	section = reDkInvalidCharacters.ReplaceAllString(section, "_")
	return section
}
