// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardListItem A dashboard within a list.
type DashboardListItem struct {
	// Creator of the object.
	Author *Creator `json:"author,omitempty"`
	// Date of creation of the dashboard.
	Created *time.Time `json:"created,omitempty"`
	// URL to the icon of the dashboard.
	Icon datadog.NullableString `json:"icon,omitempty"`
	// ID of the dashboard.
	Id string `json:"id"`
	// The short name of the integration.
	IntegrationId datadog.NullableString `json:"integration_id,omitempty"`
	// Whether or not the dashboard is in the favorites.
	IsFavorite *bool `json:"is_favorite,omitempty"`
	// Whether or not the dashboard is read only.
	IsReadOnly *bool `json:"is_read_only,omitempty"`
	// Whether the dashboard is publicly shared or not.
	IsShared *bool `json:"is_shared,omitempty"`
	// Date of last edition of the dashboard.
	Modified *time.Time `json:"modified,omitempty"`
	// Popularity of the dashboard.
	Popularity *int32 `json:"popularity,omitempty"`
	// List of team names representing ownership of a dashboard.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// Title of the dashboard.
	Title *string `json:"title,omitempty"`
	// The type of the dashboard.
	Type DashboardType `json:"type"`
	// URL path to the dashboard.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardListItem instantiates a new DashboardListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardListItem(id string, typeVar DashboardType) *DashboardListItem {
	this := DashboardListItem{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDashboardListItemWithDefaults instantiates a new DashboardListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardListItemWithDefaults() *DashboardListItem {
	this := DashboardListItem{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *DashboardListItem) GetAuthor() Creator {
	if o == nil || o.Author == nil {
		var ret Creator
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetAuthorOk() (*Creator, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *DashboardListItem) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given Creator and assigns it to the Author field.
func (o *DashboardListItem) SetAuthor(v Creator) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DashboardListItem) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DashboardListItem) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DashboardListItem) SetCreated(v time.Time) {
	o.Created = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardListItem) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardListItem) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *DashboardListItem) HasIcon() bool {
	return o != nil && o.Icon.IsSet()
}

// SetIcon gets a reference to the given datadog.NullableString and assigns it to the Icon field.
func (o *DashboardListItem) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil.
func (o *DashboardListItem) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil.
func (o *DashboardListItem) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value.
func (o *DashboardListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DashboardListItem) SetId(v string) {
	o.Id = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardListItem) GetIntegrationId() string {
	if o == nil || o.IntegrationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId.Get()
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardListItem) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationId.Get(), o.IntegrationId.IsSet()
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *DashboardListItem) HasIntegrationId() bool {
	return o != nil && o.IntegrationId.IsSet()
}

// SetIntegrationId gets a reference to the given datadog.NullableString and assigns it to the IntegrationId field.
func (o *DashboardListItem) SetIntegrationId(v string) {
	o.IntegrationId.Set(&v)
}

// SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil.
func (o *DashboardListItem) SetIntegrationIdNil() {
	o.IntegrationId.Set(nil)
}

// UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil.
func (o *DashboardListItem) UnsetIntegrationId() {
	o.IntegrationId.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *DashboardListItem) GetIsFavorite() bool {
	if o == nil || o.IsFavorite == nil {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || o.IsFavorite == nil {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *DashboardListItem) HasIsFavorite() bool {
	return o != nil && o.IsFavorite != nil
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *DashboardListItem) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *DashboardListItem) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *DashboardListItem) HasIsReadOnly() bool {
	return o != nil && o.IsReadOnly != nil
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *DashboardListItem) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *DashboardListItem) GetIsShared() bool {
	if o == nil || o.IsShared == nil {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetIsSharedOk() (*bool, bool) {
	if o == nil || o.IsShared == nil {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *DashboardListItem) HasIsShared() bool {
	return o != nil && o.IsShared != nil
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *DashboardListItem) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *DashboardListItem) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *DashboardListItem) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *DashboardListItem) SetModified(v time.Time) {
	o.Modified = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *DashboardListItem) GetPopularity() int32 {
	if o == nil || o.Popularity == nil {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetPopularityOk() (*int32, bool) {
	if o == nil || o.Popularity == nil {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *DashboardListItem) HasPopularity() bool {
	return o != nil && o.Popularity != nil
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *DashboardListItem) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardListItem) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardListItem) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *DashboardListItem) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *DashboardListItem) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *DashboardListItem) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *DashboardListItem) UnsetTags() {
	o.Tags.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DashboardListItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DashboardListItem) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DashboardListItem) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *DashboardListItem) GetType() DashboardType {
	if o == nil {
		var ret DashboardType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetTypeOk() (*DashboardType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardListItem) SetType(v DashboardType) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DashboardListItem) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItem) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DashboardListItem) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DashboardListItem) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["id"] = o.Id
	if o.IntegrationId.IsSet() {
		toSerialize["integration_id"] = o.IntegrationId.Get()
	}
	if o.IsFavorite != nil {
		toSerialize["is_favorite"] = o.IsFavorite
	}
	if o.IsReadOnly != nil {
		toSerialize["is_read_only"] = o.IsReadOnly
	}
	if o.IsShared != nil {
		toSerialize["is_shared"] = o.IsShared
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Popularity != nil {
		toSerialize["popularity"] = o.Popularity
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author        *Creator                     `json:"author,omitempty"`
		Created       *time.Time                   `json:"created,omitempty"`
		Icon          datadog.NullableString       `json:"icon,omitempty"`
		Id            *string                      `json:"id"`
		IntegrationId datadog.NullableString       `json:"integration_id,omitempty"`
		IsFavorite    *bool                        `json:"is_favorite,omitempty"`
		IsReadOnly    *bool                        `json:"is_read_only,omitempty"`
		IsShared      *bool                        `json:"is_shared,omitempty"`
		Modified      *time.Time                   `json:"modified,omitempty"`
		Popularity    *int32                       `json:"popularity,omitempty"`
		Tags          datadog.NullableList[string] `json:"tags,omitempty"`
		Title         *string                      `json:"title,omitempty"`
		Type          *DashboardType               `json:"type"`
		Url           *string                      `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created", "icon", "id", "integration_id", "is_favorite", "is_read_only", "is_shared", "modified", "popularity", "tags", "title", "type", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.Created = all.Created
	o.Icon = all.Icon
	o.Id = *all.Id
	o.IntegrationId = all.IntegrationId
	o.IsFavorite = all.IsFavorite
	o.IsReadOnly = all.IsReadOnly
	o.IsShared = all.IsShared
	o.Modified = all.Modified
	o.Popularity = all.Popularity
	o.Tags = all.Tags
	o.Title = all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
