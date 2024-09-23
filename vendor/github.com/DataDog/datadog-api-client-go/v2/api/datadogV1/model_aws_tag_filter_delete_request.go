// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSTagFilterDeleteRequest The objects used to delete an AWS tag filter entry.
type AWSTagFilterDeleteRequest struct {
	// The unique identifier of your AWS account.
	AccountId *string `json:"account_id,omitempty"`
	// The namespace associated with the tag filter entry.
	Namespace *AWSNamespace `json:"namespace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSTagFilterDeleteRequest instantiates a new AWSTagFilterDeleteRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSTagFilterDeleteRequest() *AWSTagFilterDeleteRequest {
	this := AWSTagFilterDeleteRequest{}
	return &this
}

// NewAWSTagFilterDeleteRequestWithDefaults instantiates a new AWSTagFilterDeleteRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSTagFilterDeleteRequestWithDefaults() *AWSTagFilterDeleteRequest {
	this := AWSTagFilterDeleteRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSTagFilterDeleteRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterDeleteRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSTagFilterDeleteRequest) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSTagFilterDeleteRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AWSTagFilterDeleteRequest) GetNamespace() AWSNamespace {
	if o == nil || o.Namespace == nil {
		var ret AWSNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterDeleteRequest) GetNamespaceOk() (*AWSNamespace, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AWSTagFilterDeleteRequest) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given AWSNamespace and assigns it to the Namespace field.
func (o *AWSTagFilterDeleteRequest) SetNamespace(v AWSNamespace) {
	o.Namespace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSTagFilterDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSTagFilterDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string       `json:"account_id,omitempty"`
		Namespace *AWSNamespace `json:"namespace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "namespace"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	if all.Namespace != nil && !all.Namespace.IsValid() {
		hasInvalidField = true
	} else {
		o.Namespace = all.Namespace
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
