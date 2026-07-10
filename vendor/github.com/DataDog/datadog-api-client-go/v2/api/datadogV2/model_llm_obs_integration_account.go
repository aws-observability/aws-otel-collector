// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsIntegrationAccount A configured account for an LLM provider integration.
type LLMObsIntegrationAccount struct {
	// Provider-specific account identifier.
	AccountId string `json:"account_id"`
	// Human-readable name for the integration account.
	AccountName string `json:"account_name"`
	// Provider region associated with the account, if applicable.
	AccountRegion *string `json:"account_region,omitempty"`
	// Azure OpenAI-specific metadata for an integration account or inference request.
	AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata `json:"azure_openai_metadata,omitempty"`
	// Unique identifier for the integration account.
	Id string `json:"id"`
	// The name of the LLM provider integration.
	Integration string `json:"integration"`
	// Vertex AI-specific metadata for an integration account or inference request.
	VertexAiMetadata *LLMObsVertexAIMetadata `json:"vertex_ai_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsIntegrationAccount instantiates a new LLMObsIntegrationAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsIntegrationAccount(accountId string, accountName string, id string, integration string) *LLMObsIntegrationAccount {
	this := LLMObsIntegrationAccount{}
	this.AccountId = accountId
	this.AccountName = accountName
	this.Id = id
	this.Integration = integration
	return &this
}

// NewLLMObsIntegrationAccountWithDefaults instantiates a new LLMObsIntegrationAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsIntegrationAccountWithDefaults() *LLMObsIntegrationAccount {
	this := LLMObsIntegrationAccount{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *LLMObsIntegrationAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *LLMObsIntegrationAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountName returns the AccountName field value.
func (o *LLMObsIntegrationAccount) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value.
func (o *LLMObsIntegrationAccount) SetAccountName(v string) {
	o.AccountName = v
}

// GetAccountRegion returns the AccountRegion field value if set, zero value otherwise.
func (o *LLMObsIntegrationAccount) GetAccountRegion() string {
	if o == nil || o.AccountRegion == nil {
		var ret string
		return ret
	}
	return *o.AccountRegion
}

// GetAccountRegionOk returns a tuple with the AccountRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetAccountRegionOk() (*string, bool) {
	if o == nil || o.AccountRegion == nil {
		return nil, false
	}
	return o.AccountRegion, true
}

// HasAccountRegion returns a boolean if a field has been set.
func (o *LLMObsIntegrationAccount) HasAccountRegion() bool {
	return o != nil && o.AccountRegion != nil
}

// SetAccountRegion gets a reference to the given string and assigns it to the AccountRegion field.
func (o *LLMObsIntegrationAccount) SetAccountRegion(v string) {
	o.AccountRegion = &v
}

// GetAzureOpenaiMetadata returns the AzureOpenaiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationAccount) GetAzureOpenaiMetadata() LLMObsAzureOpenAIMetadata {
	if o == nil || o.AzureOpenaiMetadata == nil {
		var ret LLMObsAzureOpenAIMetadata
		return ret
	}
	return *o.AzureOpenaiMetadata
}

// GetAzureOpenaiMetadataOk returns a tuple with the AzureOpenaiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetAzureOpenaiMetadataOk() (*LLMObsAzureOpenAIMetadata, bool) {
	if o == nil || o.AzureOpenaiMetadata == nil {
		return nil, false
	}
	return o.AzureOpenaiMetadata, true
}

// HasAzureOpenaiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationAccount) HasAzureOpenaiMetadata() bool {
	return o != nil && o.AzureOpenaiMetadata != nil
}

// SetAzureOpenaiMetadata gets a reference to the given LLMObsAzureOpenAIMetadata and assigns it to the AzureOpenaiMetadata field.
func (o *LLMObsIntegrationAccount) SetAzureOpenaiMetadata(v LLMObsAzureOpenAIMetadata) {
	o.AzureOpenaiMetadata = &v
}

// GetId returns the Id field value.
func (o *LLMObsIntegrationAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsIntegrationAccount) SetId(v string) {
	o.Id = v
}

// GetIntegration returns the Integration field value.
func (o *LLMObsIntegrationAccount) GetIntegration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *LLMObsIntegrationAccount) SetIntegration(v string) {
	o.Integration = v
}

// GetVertexAiMetadata returns the VertexAiMetadata field value if set, zero value otherwise.
func (o *LLMObsIntegrationAccount) GetVertexAiMetadata() LLMObsVertexAIMetadata {
	if o == nil || o.VertexAiMetadata == nil {
		var ret LLMObsVertexAIMetadata
		return ret
	}
	return *o.VertexAiMetadata
}

// GetVertexAiMetadataOk returns a tuple with the VertexAiMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationAccount) GetVertexAiMetadataOk() (*LLMObsVertexAIMetadata, bool) {
	if o == nil || o.VertexAiMetadata == nil {
		return nil, false
	}
	return o.VertexAiMetadata, true
}

// HasVertexAiMetadata returns a boolean if a field has been set.
func (o *LLMObsIntegrationAccount) HasVertexAiMetadata() bool {
	return o != nil && o.VertexAiMetadata != nil
}

// SetVertexAiMetadata gets a reference to the given LLMObsVertexAIMetadata and assigns it to the VertexAiMetadata field.
func (o *LLMObsIntegrationAccount) SetVertexAiMetadata(v LLMObsVertexAIMetadata) {
	o.VertexAiMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsIntegrationAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["account_name"] = o.AccountName
	if o.AccountRegion != nil {
		toSerialize["account_region"] = o.AccountRegion
	}
	if o.AzureOpenaiMetadata != nil {
		toSerialize["azure_openai_metadata"] = o.AzureOpenaiMetadata
	}
	toSerialize["id"] = o.Id
	toSerialize["integration"] = o.Integration
	if o.VertexAiMetadata != nil {
		toSerialize["vertex_ai_metadata"] = o.VertexAiMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsIntegrationAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string                    `json:"account_id"`
		AccountName         *string                    `json:"account_name"`
		AccountRegion       *string                    `json:"account_region,omitempty"`
		AzureOpenaiMetadata *LLMObsAzureOpenAIMetadata `json:"azure_openai_metadata,omitempty"`
		Id                  *string                    `json:"id"`
		Integration         *string                    `json:"integration"`
		VertexAiMetadata    *LLMObsVertexAIMetadata    `json:"vertex_ai_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.AccountName == nil {
		return fmt.Errorf("required field account_name missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "account_name", "account_region", "azure_openai_metadata", "id", "integration", "vertex_ai_metadata"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	o.AccountName = *all.AccountName
	o.AccountRegion = all.AccountRegion
	if all.AzureOpenaiMetadata != nil && all.AzureOpenaiMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AzureOpenaiMetadata = all.AzureOpenaiMetadata
	o.Id = *all.Id
	o.Integration = *all.Integration
	if all.VertexAiMetadata != nil && all.VertexAiMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.VertexAiMetadata = all.VertexAiMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
