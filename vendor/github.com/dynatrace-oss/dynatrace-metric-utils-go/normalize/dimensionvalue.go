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
// limitations under the License

package normalize

import (
	"regexp"
)

var (
	reDvControlCharacters  = regexp.MustCompile("\\p{C}+")
	reDvToEscapeCharacters = regexp.MustCompile(`([= ,\\"])`)

	// This regex checks if there is an odd number of trailing backslashes in the string. It can be
	// read as: {not a slash}{any number of 2-slash pairs}{one slash}{end line}.
	reDvHasOddNumberOfTrailingBackslashes = regexp.MustCompile(`[^\\](?:\\\\)*\\$`)
)

const (
	dimensionValueMaxLength = 250
)

// DimensionValue returns a string without control characters and escaped characters.
func DimensionValue(value string) string {
	if len(value) > dimensionValueMaxLength {
		value = value[:dimensionValueMaxLength]
	}

	value = removeControlCharacters(value)
	value = escapeCharacters(value)

	return value
}

func removeControlCharacters(s string) string {
	return reDvControlCharacters.ReplaceAllString(s, "_")
}

func escapeCharacters(s string) string {
	escaped := reDvToEscapeCharacters.ReplaceAllString(s, "\\$1")
	if len(escaped) > dimensionValueMaxLength {
		escaped = escaped[:dimensionValueMaxLength]

		if reDvHasOddNumberOfTrailingBackslashes.MatchString(escaped) {
			escaped = escaped[:dimensionValueMaxLength-1]
		}
	}

	return escaped
}
