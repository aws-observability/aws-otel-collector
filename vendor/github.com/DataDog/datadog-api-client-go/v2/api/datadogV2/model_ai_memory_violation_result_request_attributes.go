// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiMemoryViolationResultRequestAttributes Attributes for creating an AI memory violation result.
type AiMemoryViolationResultRequestAttributes struct {
	// The line number where the violation was found.
	Line int64 `json:"line"`
	// A message explaining the violation result.
	Message string `json:"message"`
	// The file path where the violation was found.
	Name string `json:"name"`
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// The rule identifier in the format ruleset/rule.
	Rule string `json:"rule"`
	// The git commit SHA where the violation was found.
	Sha string `json:"sha"`
	// The type of AI memory violation result indicating whether it is a true positive or false positive.
	Type AiMemoryViolationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiMemoryViolationResultRequestAttributes instantiates a new AiMemoryViolationResultRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiMemoryViolationResultRequestAttributes(line int64, message string, name string, repositoryId string, rule string, sha string, typeVar AiMemoryViolationType) *AiMemoryViolationResultRequestAttributes {
	this := AiMemoryViolationResultRequestAttributes{}
	this.Line = line
	this.Message = message
	this.Name = name
	this.RepositoryId = repositoryId
	this.Rule = rule
	this.Sha = sha
	this.Type = typeVar
	return &this
}

// NewAiMemoryViolationResultRequestAttributesWithDefaults instantiates a new AiMemoryViolationResultRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiMemoryViolationResultRequestAttributesWithDefaults() *AiMemoryViolationResultRequestAttributes {
	this := AiMemoryViolationResultRequestAttributes{}
	return &this
}

// GetLine returns the Line field value.
func (o *AiMemoryViolationResultRequestAttributes) GetLine() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetLineOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetLine(v int64) {
	o.Line = v
}

// GetMessage returns the Message field value.
func (o *AiMemoryViolationResultRequestAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *AiMemoryViolationResultRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *AiMemoryViolationResultRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// GetRule returns the Rule field value.
func (o *AiMemoryViolationResultRequestAttributes) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetRule(v string) {
	o.Rule = v
}

// GetSha returns the Sha field value.
func (o *AiMemoryViolationResultRequestAttributes) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetSha(v string) {
	o.Sha = v
}

// GetType returns the Type field value.
func (o *AiMemoryViolationResultRequestAttributes) GetType() AiMemoryViolationType {
	if o == nil {
		var ret AiMemoryViolationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultRequestAttributes) GetTypeOk() (*AiMemoryViolationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiMemoryViolationResultRequestAttributes) SetType(v AiMemoryViolationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiMemoryViolationResultRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["line"] = o.Line
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["rule"] = o.Rule
	toSerialize["sha"] = o.Sha
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiMemoryViolationResultRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Line         *int64                 `json:"line"`
		Message      *string                `json:"message"`
		Name         *string                `json:"name"`
		RepositoryId *string                `json:"repository_id"`
		Rule         *string                `json:"rule"`
		Sha          *string                `json:"sha"`
		Type         *AiMemoryViolationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Line == nil {
		return fmt.Errorf("required field line missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	if all.Sha == nil {
		return fmt.Errorf("required field sha missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"line", "message", "name", "repository_id", "rule", "sha", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Line = *all.Line
	o.Message = *all.Message
	o.Name = *all.Name
	o.RepositoryId = *all.RepositoryId
	o.Rule = *all.Rule
	o.Sha = *all.Sha
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
