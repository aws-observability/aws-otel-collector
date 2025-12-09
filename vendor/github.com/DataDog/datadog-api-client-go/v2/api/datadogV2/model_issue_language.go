// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueLanguage Programming language associated with the issue.
type IssueLanguage string

// List of IssueLanguage.
const (
	ISSUELANGUAGE_BRIGHTSCRIPT IssueLanguage = "BRIGHTSCRIPT"
	ISSUELANGUAGE_C            IssueLanguage = "C"
	ISSUELANGUAGE_C_PLUS_PLUS  IssueLanguage = "C_PLUS_PLUS"
	ISSUELANGUAGE_C_SHARP      IssueLanguage = "C_SHARP"
	ISSUELANGUAGE_CLOJURE      IssueLanguage = "CLOJURE"
	ISSUELANGUAGE_DOT_NET      IssueLanguage = "DOT_NET"
	ISSUELANGUAGE_ELIXIR       IssueLanguage = "ELIXIR"
	ISSUELANGUAGE_ERLANG       IssueLanguage = "ERLANG"
	ISSUELANGUAGE_GO           IssueLanguage = "GO"
	ISSUELANGUAGE_GROOVY       IssueLanguage = "GROOVY"
	ISSUELANGUAGE_HASKELL      IssueLanguage = "HASKELL"
	ISSUELANGUAGE_HCL          IssueLanguage = "HCL"
	ISSUELANGUAGE_JAVA         IssueLanguage = "JAVA"
	ISSUELANGUAGE_JAVASCRIPT   IssueLanguage = "JAVASCRIPT"
	ISSUELANGUAGE_JVM          IssueLanguage = "JVM"
	ISSUELANGUAGE_KOTLIN       IssueLanguage = "KOTLIN"
	ISSUELANGUAGE_OBJECTIVE_C  IssueLanguage = "OBJECTIVE_C"
	ISSUELANGUAGE_PERL         IssueLanguage = "PERL"
	ISSUELANGUAGE_PHP          IssueLanguage = "PHP"
	ISSUELANGUAGE_PYTHON       IssueLanguage = "PYTHON"
	ISSUELANGUAGE_RUBY         IssueLanguage = "RUBY"
	ISSUELANGUAGE_RUST         IssueLanguage = "RUST"
	ISSUELANGUAGE_SCALA        IssueLanguage = "SCALA"
	ISSUELANGUAGE_SWIFT        IssueLanguage = "SWIFT"
	ISSUELANGUAGE_TERRAFORM    IssueLanguage = "TERRAFORM"
	ISSUELANGUAGE_TYPESCRIPT   IssueLanguage = "TYPESCRIPT"
	ISSUELANGUAGE_UNKNOWN      IssueLanguage = "UNKNOWN"
)

var allowedIssueLanguageEnumValues = []IssueLanguage{
	ISSUELANGUAGE_BRIGHTSCRIPT,
	ISSUELANGUAGE_C,
	ISSUELANGUAGE_C_PLUS_PLUS,
	ISSUELANGUAGE_C_SHARP,
	ISSUELANGUAGE_CLOJURE,
	ISSUELANGUAGE_DOT_NET,
	ISSUELANGUAGE_ELIXIR,
	ISSUELANGUAGE_ERLANG,
	ISSUELANGUAGE_GO,
	ISSUELANGUAGE_GROOVY,
	ISSUELANGUAGE_HASKELL,
	ISSUELANGUAGE_HCL,
	ISSUELANGUAGE_JAVA,
	ISSUELANGUAGE_JAVASCRIPT,
	ISSUELANGUAGE_JVM,
	ISSUELANGUAGE_KOTLIN,
	ISSUELANGUAGE_OBJECTIVE_C,
	ISSUELANGUAGE_PERL,
	ISSUELANGUAGE_PHP,
	ISSUELANGUAGE_PYTHON,
	ISSUELANGUAGE_RUBY,
	ISSUELANGUAGE_RUST,
	ISSUELANGUAGE_SCALA,
	ISSUELANGUAGE_SWIFT,
	ISSUELANGUAGE_TERRAFORM,
	ISSUELANGUAGE_TYPESCRIPT,
	ISSUELANGUAGE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueLanguage) GetAllowedValues() []IssueLanguage {
	return allowedIssueLanguageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueLanguage(value)
	return nil
}

// NewIssueLanguageFromValue returns a pointer to a valid IssueLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueLanguageFromValue(v string) (*IssueLanguage, error) {
	ev := IssueLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueLanguage: valid values are %v", v, allowedIssueLanguageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueLanguage) IsValid() bool {
	for _, existing := range allowedIssueLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueLanguage value.
func (v IssueLanguage) Ptr() *IssueLanguage {
	return &v
}
