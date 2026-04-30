// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalUpdateAttributes Attributes for updating the triage state or assignee of a security signal.
type SecurityMonitoringSignalUpdateAttributes struct {
	// Optional comment to display on archived signals.
	ArchiveComment *string `json:"archive_comment,omitempty"`
	// Reason a signal is archived.
	ArchiveReason *SecurityMonitoringSignalArchiveReason `json:"archive_reason,omitempty"`
	// Object representing a given user entity.
	Assignee *SecurityMonitoringTriageUser `json:"assignee,omitempty"`
	// The new triage state of the signal.
	State *SecurityMonitoringSignalState `json:"state,omitempty"`
	// Version of the updated signal. If server side version is higher, update will be rejected.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalUpdateAttributes instantiates a new SecurityMonitoringSignalUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalUpdateAttributes() *SecurityMonitoringSignalUpdateAttributes {
	this := SecurityMonitoringSignalUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringSignalUpdateAttributesWithDefaults instantiates a new SecurityMonitoringSignalUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalUpdateAttributesWithDefaults() *SecurityMonitoringSignalUpdateAttributes {
	this := SecurityMonitoringSignalUpdateAttributes{}
	return &this
}

// GetArchiveComment returns the ArchiveComment field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalUpdateAttributes) GetArchiveComment() string {
	if o == nil || o.ArchiveComment == nil {
		var ret string
		return ret
	}
	return *o.ArchiveComment
}

// GetArchiveCommentOk returns a tuple with the ArchiveComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) GetArchiveCommentOk() (*string, bool) {
	if o == nil || o.ArchiveComment == nil {
		return nil, false
	}
	return o.ArchiveComment, true
}

// HasArchiveComment returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) HasArchiveComment() bool {
	return o != nil && o.ArchiveComment != nil
}

// SetArchiveComment gets a reference to the given string and assigns it to the ArchiveComment field.
func (o *SecurityMonitoringSignalUpdateAttributes) SetArchiveComment(v string) {
	o.ArchiveComment = &v
}

// GetArchiveReason returns the ArchiveReason field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalUpdateAttributes) GetArchiveReason() SecurityMonitoringSignalArchiveReason {
	if o == nil || o.ArchiveReason == nil {
		var ret SecurityMonitoringSignalArchiveReason
		return ret
	}
	return *o.ArchiveReason
}

// GetArchiveReasonOk returns a tuple with the ArchiveReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) GetArchiveReasonOk() (*SecurityMonitoringSignalArchiveReason, bool) {
	if o == nil || o.ArchiveReason == nil {
		return nil, false
	}
	return o.ArchiveReason, true
}

// HasArchiveReason returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) HasArchiveReason() bool {
	return o != nil && o.ArchiveReason != nil
}

// SetArchiveReason gets a reference to the given SecurityMonitoringSignalArchiveReason and assigns it to the ArchiveReason field.
func (o *SecurityMonitoringSignalUpdateAttributes) SetArchiveReason(v SecurityMonitoringSignalArchiveReason) {
	o.ArchiveReason = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalUpdateAttributes) GetAssignee() SecurityMonitoringTriageUser {
	if o == nil || o.Assignee == nil {
		var ret SecurityMonitoringTriageUser
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) GetAssigneeOk() (*SecurityMonitoringTriageUser, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) HasAssignee() bool {
	return o != nil && o.Assignee != nil
}

// SetAssignee gets a reference to the given SecurityMonitoringTriageUser and assigns it to the Assignee field.
func (o *SecurityMonitoringSignalUpdateAttributes) SetAssignee(v SecurityMonitoringTriageUser) {
	o.Assignee = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalUpdateAttributes) GetState() SecurityMonitoringSignalState {
	if o == nil || o.State == nil {
		var ret SecurityMonitoringSignalState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) GetStateOk() (*SecurityMonitoringSignalState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given SecurityMonitoringSignalState and assigns it to the State field.
func (o *SecurityMonitoringSignalUpdateAttributes) SetState(v SecurityMonitoringSignalState) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalUpdateAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalUpdateAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SecurityMonitoringSignalUpdateAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchiveComment != nil {
		toSerialize["archive_comment"] = o.ArchiveComment
	}
	if o.ArchiveReason != nil {
		toSerialize["archive_reason"] = o.ArchiveReason
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchiveComment *string                                `json:"archive_comment,omitempty"`
		ArchiveReason  *SecurityMonitoringSignalArchiveReason `json:"archive_reason,omitempty"`
		Assignee       *SecurityMonitoringTriageUser          `json:"assignee,omitempty"`
		State          *SecurityMonitoringSignalState         `json:"state,omitempty"`
		Version        *int64                                 `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archive_comment", "archive_reason", "assignee", "state", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchiveComment = all.ArchiveComment
	if all.ArchiveReason != nil && !all.ArchiveReason.IsValid() {
		hasInvalidField = true
	} else {
		o.ArchiveReason = all.ArchiveReason
	}
	if all.Assignee != nil && all.Assignee.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Assignee = all.Assignee
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
