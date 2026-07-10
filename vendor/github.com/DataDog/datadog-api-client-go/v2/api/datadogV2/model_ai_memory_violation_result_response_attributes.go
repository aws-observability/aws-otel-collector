// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiMemoryViolationResultResponseAttributes Response attributes of an AI memory violation result.
type AiMemoryViolationResultResponseAttributes struct {
	// The creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user who created the result.
	CreatedBy string `json:"created_by"`
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

// NewAiMemoryViolationResultResponseAttributes instantiates a new AiMemoryViolationResultResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiMemoryViolationResultResponseAttributes(createdAt time.Time, createdBy string, line int64, message string, name string, repositoryId string, rule string, sha string, typeVar AiMemoryViolationType) *AiMemoryViolationResultResponseAttributes {
	this := AiMemoryViolationResultResponseAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Line = line
	this.Message = message
	this.Name = name
	this.RepositoryId = repositoryId
	this.Rule = rule
	this.Sha = sha
	this.Type = typeVar
	return &this
}

// NewAiMemoryViolationResultResponseAttributesWithDefaults instantiates a new AiMemoryViolationResultResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiMemoryViolationResultResponseAttributesWithDefaults() *AiMemoryViolationResultResponseAttributes {
	this := AiMemoryViolationResultResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiMemoryViolationResultResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *AiMemoryViolationResultResponseAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLine returns the Line field value.
func (o *AiMemoryViolationResultResponseAttributes) GetLine() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetLineOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetLine(v int64) {
	o.Line = v
}

// GetMessage returns the Message field value.
func (o *AiMemoryViolationResultResponseAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *AiMemoryViolationResultResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *AiMemoryViolationResultResponseAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// GetRule returns the Rule field value.
func (o *AiMemoryViolationResultResponseAttributes) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetRule(v string) {
	o.Rule = v
}

// GetSha returns the Sha field value.
func (o *AiMemoryViolationResultResponseAttributes) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetSha(v string) {
	o.Sha = v
}

// GetType returns the Type field value.
func (o *AiMemoryViolationResultResponseAttributes) GetType() AiMemoryViolationType {
	if o == nil {
		var ret AiMemoryViolationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiMemoryViolationResultResponseAttributes) GetTypeOk() (*AiMemoryViolationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiMemoryViolationResultResponseAttributes) SetType(v AiMemoryViolationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiMemoryViolationResultResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
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
func (o *AiMemoryViolationResultResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time             `json:"created_at"`
		CreatedBy    *string                `json:"created_by"`
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
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "line", "message", "name", "repository_id", "rule", "sha", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
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
