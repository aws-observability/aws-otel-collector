// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagDescriptionAttributes Human-readable description and metadata attached to a Cloud Cost Management tag key, optionally scoped to a single cloud provider.
type CostTagDescriptionAttributes struct {
	// Cloud provider this description applies to (for example, `aws`). Empty when the description is the cross-cloud default for the tag key.
	Cloud string `json:"cloud"`
	// Timestamp when the description was created, in RFC 3339 format.
	CreatedAt string `json:"created_at"`
	// The human-readable description for the tag key.
	Description string `json:"description"`
	// Origin of the description. `human` indicates the description was written by a user, `ai_generated` was produced by AI, and `datadog` is a default supplied by Datadog.
	Source CostTagDescriptionSource `json:"source"`
	// The tag key this description applies to.
	TagKey string `json:"tag_key"`
	// Timestamp when the description was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagDescriptionAttributes instantiates a new CostTagDescriptionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagDescriptionAttributes(cloud string, createdAt string, description string, source CostTagDescriptionSource, tagKey string, updatedAt string) *CostTagDescriptionAttributes {
	this := CostTagDescriptionAttributes{}
	this.Cloud = cloud
	this.CreatedAt = createdAt
	this.Description = description
	this.Source = source
	this.TagKey = tagKey
	this.UpdatedAt = updatedAt
	return &this
}

// NewCostTagDescriptionAttributesWithDefaults instantiates a new CostTagDescriptionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagDescriptionAttributesWithDefaults() *CostTagDescriptionAttributes {
	this := CostTagDescriptionAttributes{}
	return &this
}

// GetCloud returns the Cloud field value.
func (o *CostTagDescriptionAttributes) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value.
func (o *CostTagDescriptionAttributes) SetCloud(v string) {
	o.Cloud = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CostTagDescriptionAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CostTagDescriptionAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *CostTagDescriptionAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CostTagDescriptionAttributes) SetDescription(v string) {
	o.Description = v
}

// GetSource returns the Source field value.
func (o *CostTagDescriptionAttributes) GetSource() CostTagDescriptionSource {
	if o == nil {
		var ret CostTagDescriptionSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetSourceOk() (*CostTagDescriptionSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *CostTagDescriptionAttributes) SetSource(v CostTagDescriptionSource) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *CostTagDescriptionAttributes) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *CostTagDescriptionAttributes) SetTagKey(v string) {
	o.TagKey = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *CostTagDescriptionAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *CostTagDescriptionAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagDescriptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cloud"] = o.Cloud
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["source"] = o.Source
	toSerialize["tag_key"] = o.TagKey
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagDescriptionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cloud       *string                   `json:"cloud"`
		CreatedAt   *string                   `json:"created_at"`
		Description *string                   `json:"description"`
		Source      *CostTagDescriptionSource `json:"source"`
		TagKey      *string                   `json:"tag_key"`
		UpdatedAt   *string                   `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cloud == nil {
		return fmt.Errorf("required field cloud missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud", "created_at", "description", "source", "tag_key", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cloud = *all.Cloud
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.TagKey = *all.TagKey
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
