// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Language Programming language
type Language string

// List of Language.
const (
	LANGUAGE_PYTHON     Language = "PYTHON"
	LANGUAGE_JAVASCRIPT Language = "JAVASCRIPT"
	LANGUAGE_TYPESCRIPT Language = "TYPESCRIPT"
	LANGUAGE_JAVA       Language = "JAVA"
	LANGUAGE_GO         Language = "GO"
	LANGUAGE_YAML       Language = "YAML"
	LANGUAGE_RUBY       Language = "RUBY"
	LANGUAGE_CSHARP     Language = "CSHARP"
	LANGUAGE_PHP        Language = "PHP"
	LANGUAGE_KOTLIN     Language = "KOTLIN"
	LANGUAGE_SWIFT      Language = "SWIFT"
)

var allowedLanguageEnumValues = []Language{
	LANGUAGE_PYTHON,
	LANGUAGE_JAVASCRIPT,
	LANGUAGE_TYPESCRIPT,
	LANGUAGE_JAVA,
	LANGUAGE_GO,
	LANGUAGE_YAML,
	LANGUAGE_RUBY,
	LANGUAGE_CSHARP,
	LANGUAGE_PHP,
	LANGUAGE_KOTLIN,
	LANGUAGE_SWIFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Language) GetAllowedValues() []Language {
	return allowedLanguageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Language) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Language(value)
	return nil
}

// NewLanguageFromValue returns a pointer to a valid Language
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLanguageFromValue(v string) (*Language, error) {
	ev := Language(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Language: valid values are %v", v, allowedLanguageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Language) IsValid() bool {
	for _, existing := range allowedLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Language value.
func (v Language) Ptr() *Language {
	return &v
}
