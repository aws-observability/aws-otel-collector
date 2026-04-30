// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HamrOrgConnectionAttributesRequest Attributes for a HAMR organization connection request.
type HamrOrgConnectionAttributesRequest struct {
	// Status of the HAMR connection:
	// - 0: UNSPECIFIED - Connection status not specified
	// - 1: ONBOARDING - Initial setup of HAMR connection
	// - 2: PASSIVE - Secondary organization in passive standby mode
	// - 3: FAILOVER - Liminal status between PASSIVE and ACTIVE
	// - 4: ACTIVE - Organization is an active failover
	// - 5: RECOVERY - Recovery operation in progress
	HamrStatus HamrOrgConnectionStatus `json:"hamr_status"`
	// Indicates whether this organization is the primary organization in the HAMR relationship.
	// If true, this is the primary organization. If false, this is the secondary/backup organization.
	IsPrimary bool `json:"is_primary"`
	// Username or identifier of the user who last modified this HAMR connection.
	ModifiedBy string `json:"modified_by"`
	// Datacenter location of the target organization (e.g., us1, eu1, us5).
	TargetOrgDatacenter string `json:"target_org_datacenter"`
	// Name of the target organization in the HAMR relationship.
	TargetOrgName string `json:"target_org_name"`
	// UUID of the target organization in the HAMR relationship.
	TargetOrgUuid string `json:"target_org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHamrOrgConnectionAttributesRequest instantiates a new HamrOrgConnectionAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHamrOrgConnectionAttributesRequest(hamrStatus HamrOrgConnectionStatus, isPrimary bool, modifiedBy string, targetOrgDatacenter string, targetOrgName string, targetOrgUuid string) *HamrOrgConnectionAttributesRequest {
	this := HamrOrgConnectionAttributesRequest{}
	this.HamrStatus = hamrStatus
	this.IsPrimary = isPrimary
	this.ModifiedBy = modifiedBy
	this.TargetOrgDatacenter = targetOrgDatacenter
	this.TargetOrgName = targetOrgName
	this.TargetOrgUuid = targetOrgUuid
	return &this
}

// NewHamrOrgConnectionAttributesRequestWithDefaults instantiates a new HamrOrgConnectionAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHamrOrgConnectionAttributesRequestWithDefaults() *HamrOrgConnectionAttributesRequest {
	this := HamrOrgConnectionAttributesRequest{}
	return &this
}

// GetHamrStatus returns the HamrStatus field value.
func (o *HamrOrgConnectionAttributesRequest) GetHamrStatus() HamrOrgConnectionStatus {
	if o == nil {
		var ret HamrOrgConnectionStatus
		return ret
	}
	return o.HamrStatus
}

// GetHamrStatusOk returns a tuple with the HamrStatus field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetHamrStatusOk() (*HamrOrgConnectionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HamrStatus, true
}

// SetHamrStatus sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetHamrStatus(v HamrOrgConnectionStatus) {
	o.HamrStatus = v
}

// GetIsPrimary returns the IsPrimary field value.
func (o *HamrOrgConnectionAttributesRequest) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *HamrOrgConnectionAttributesRequest) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetTargetOrgDatacenter returns the TargetOrgDatacenter field value.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgDatacenter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetOrgDatacenter
}

// GetTargetOrgDatacenterOk returns a tuple with the TargetOrgDatacenter field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgDatacenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOrgDatacenter, true
}

// SetTargetOrgDatacenter sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetTargetOrgDatacenter(v string) {
	o.TargetOrgDatacenter = v
}

// GetTargetOrgName returns the TargetOrgName field value.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetOrgName
}

// GetTargetOrgNameOk returns a tuple with the TargetOrgName field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOrgName, true
}

// SetTargetOrgName sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetTargetOrgName(v string) {
	o.TargetOrgName = v
}

// GetTargetOrgUuid returns the TargetOrgUuid field value.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetOrgUuid
}

// GetTargetOrgUuidOk returns a tuple with the TargetOrgUuid field value
// and a boolean to check if the value has been set.
func (o *HamrOrgConnectionAttributesRequest) GetTargetOrgUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOrgUuid, true
}

// SetTargetOrgUuid sets field value.
func (o *HamrOrgConnectionAttributesRequest) SetTargetOrgUuid(v string) {
	o.TargetOrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HamrOrgConnectionAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hamr_status"] = o.HamrStatus
	toSerialize["is_primary"] = o.IsPrimary
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["target_org_datacenter"] = o.TargetOrgDatacenter
	toSerialize["target_org_name"] = o.TargetOrgName
	toSerialize["target_org_uuid"] = o.TargetOrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HamrOrgConnectionAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HamrStatus          *HamrOrgConnectionStatus `json:"hamr_status"`
		IsPrimary           *bool                    `json:"is_primary"`
		ModifiedBy          *string                  `json:"modified_by"`
		TargetOrgDatacenter *string                  `json:"target_org_datacenter"`
		TargetOrgName       *string                  `json:"target_org_name"`
		TargetOrgUuid       *string                  `json:"target_org_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HamrStatus == nil {
		return fmt.Errorf("required field hamr_status missing")
	}
	if all.IsPrimary == nil {
		return fmt.Errorf("required field is_primary missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.TargetOrgDatacenter == nil {
		return fmt.Errorf("required field target_org_datacenter missing")
	}
	if all.TargetOrgName == nil {
		return fmt.Errorf("required field target_org_name missing")
	}
	if all.TargetOrgUuid == nil {
		return fmt.Errorf("required field target_org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hamr_status", "is_primary", "modified_by", "target_org_datacenter", "target_org_name", "target_org_uuid"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.HamrStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.HamrStatus = *all.HamrStatus
	}
	o.IsPrimary = *all.IsPrimary
	o.ModifiedBy = *all.ModifiedBy
	o.TargetOrgDatacenter = *all.TargetOrgDatacenter
	o.TargetOrgName = *all.TargetOrgName
	o.TargetOrgUuid = *all.TargetOrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
