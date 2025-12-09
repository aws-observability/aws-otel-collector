// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteSharedDashboardResponse Response containing token of deleted shared dashboard.
type DeleteSharedDashboardResponse struct {
	// Token associated with the shared dashboard that was revoked.
	DeletedPublicDashboardToken *string `json:"deleted_public_dashboard_token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteSharedDashboardResponse instantiates a new DeleteSharedDashboardResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteSharedDashboardResponse() *DeleteSharedDashboardResponse {
	this := DeleteSharedDashboardResponse{}
	return &this
}

// NewDeleteSharedDashboardResponseWithDefaults instantiates a new DeleteSharedDashboardResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteSharedDashboardResponseWithDefaults() *DeleteSharedDashboardResponse {
	this := DeleteSharedDashboardResponse{}
	return &this
}

// GetDeletedPublicDashboardToken returns the DeletedPublicDashboardToken field value if set, zero value otherwise.
func (o *DeleteSharedDashboardResponse) GetDeletedPublicDashboardToken() string {
	if o == nil || o.DeletedPublicDashboardToken == nil {
		var ret string
		return ret
	}
	return *o.DeletedPublicDashboardToken
}

// GetDeletedPublicDashboardTokenOk returns a tuple with the DeletedPublicDashboardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSharedDashboardResponse) GetDeletedPublicDashboardTokenOk() (*string, bool) {
	if o == nil || o.DeletedPublicDashboardToken == nil {
		return nil, false
	}
	return o.DeletedPublicDashboardToken, true
}

// HasDeletedPublicDashboardToken returns a boolean if a field has been set.
func (o *DeleteSharedDashboardResponse) HasDeletedPublicDashboardToken() bool {
	return o != nil && o.DeletedPublicDashboardToken != nil
}

// SetDeletedPublicDashboardToken gets a reference to the given string and assigns it to the DeletedPublicDashboardToken field.
func (o *DeleteSharedDashboardResponse) SetDeletedPublicDashboardToken(v string) {
	o.DeletedPublicDashboardToken = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteSharedDashboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedPublicDashboardToken != nil {
		toSerialize["deleted_public_dashboard_token"] = o.DeletedPublicDashboardToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteSharedDashboardResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedPublicDashboardToken *string `json:"deleted_public_dashboard_token,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_public_dashboard_token"})
	} else {
		return err
	}
	o.DeletedPublicDashboardToken = all.DeletedPublicDashboardToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
