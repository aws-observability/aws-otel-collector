// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgAttributes Attributes of an organization associated with the authenticated user.
type GlobalOrgAttributes struct {
	// Organization information for a global organization association.
	Org GlobalOrg `json:"org"`
	// The login URL used to switch into the organization, if available.
	RedirectUrl datadog.NullableString `json:"redirect_url,omitempty"`
	// The source region of the organization.
	SourceRegion string `json:"source_region"`
	// User information for a global organization association.
	User GlobalOrgUser `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalOrgAttributes instantiates a new GlobalOrgAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalOrgAttributes(org GlobalOrg, sourceRegion string, user GlobalOrgUser) *GlobalOrgAttributes {
	this := GlobalOrgAttributes{}
	this.Org = org
	this.SourceRegion = sourceRegion
	this.User = user
	return &this
}

// NewGlobalOrgAttributesWithDefaults instantiates a new GlobalOrgAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalOrgAttributesWithDefaults() *GlobalOrgAttributes {
	this := GlobalOrgAttributes{}
	return &this
}

// GetOrg returns the Org field value.
func (o *GlobalOrgAttributes) GetOrg() GlobalOrg {
	if o == nil {
		var ret GlobalOrg
		return ret
	}
	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *GlobalOrgAttributes) GetOrgOk() (*GlobalOrg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value.
func (o *GlobalOrgAttributes) SetOrg(v GlobalOrg) {
	o.Org = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrgAttributes) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrgAttributes) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *GlobalOrgAttributes) HasRedirectUrl() bool {
	return o != nil && o.RedirectUrl.IsSet()
}

// SetRedirectUrl gets a reference to the given datadog.NullableString and assigns it to the RedirectUrl field.
func (o *GlobalOrgAttributes) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}

// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil.
func (o *GlobalOrgAttributes) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil.
func (o *GlobalOrgAttributes) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetSourceRegion returns the SourceRegion field value.
func (o *GlobalOrgAttributes) GetSourceRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceRegion
}

// GetSourceRegionOk returns a tuple with the SourceRegion field value
// and a boolean to check if the value has been set.
func (o *GlobalOrgAttributes) GetSourceRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceRegion, true
}

// SetSourceRegion sets field value.
func (o *GlobalOrgAttributes) SetSourceRegion(v string) {
	o.SourceRegion = v
}

// GetUser returns the User field value.
func (o *GlobalOrgAttributes) GetUser() GlobalOrgUser {
	if o == nil {
		var ret GlobalOrgUser
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GlobalOrgAttributes) GetUserOk() (*GlobalOrgUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *GlobalOrgAttributes) SetUser(v GlobalOrgUser) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalOrgAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["org"] = o.Org
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	toSerialize["source_region"] = o.SourceRegion
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalOrgAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Org          *GlobalOrg             `json:"org"`
		RedirectUrl  datadog.NullableString `json:"redirect_url,omitempty"`
		SourceRegion *string                `json:"source_region"`
		User         *GlobalOrgUser         `json:"user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Org == nil {
		return fmt.Errorf("required field org missing")
	}
	if all.SourceRegion == nil {
		return fmt.Errorf("required field source_region missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org", "redirect_url", "source_region", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Org.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Org = *all.Org
	o.RedirectUrl = all.RedirectUrl
	o.SourceRegion = *all.SourceRegion
	if all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
