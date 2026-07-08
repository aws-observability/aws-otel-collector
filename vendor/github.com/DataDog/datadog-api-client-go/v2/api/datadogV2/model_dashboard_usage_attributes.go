// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardUsageAttributes Usage statistics for a dashboard. The `viewer` field and all view-count fields (`total_views`, `viewed_at`, `total_views_by_type`) are populated only when Real User Monitoring (RUM) is active for the org.
type DashboardUsageAttributes struct {
	// A user referenced from a dashboard usage record (author or viewer).
	Author NullableDashboardUsageUser `json:"author,omitempty"`
	// When the dashboard was created.
	CreatedAt datadog.NullableTime `json:"created_at,omitempty"`
	// The dashboard quality score, or `null` when no score is available.
	DashboardQualityScore datadog.NullableFloat64 `json:"dashboard_quality_score,omitempty"`
	// When the dashboard was most recently edited.
	EditedAt datadog.NullableTime `json:"edited_at,omitempty"`
	// The Datadog organization that owns the dashboard.
	OrgId int64 `json:"org_id"`
	// Teams the dashboard is tagged with.
	Teams datadog.NullableList[string] `json:"teams,omitempty"`
	// The dashboard title.
	Title *string `json:"title,omitempty"`
	// Total view count for the dashboard. Counts only views captured by Real User Monitoring (RUM); `0` in orgs without RUM.
	TotalViews *int64 `json:"total_views,omitempty"`
	// View counts keyed by view type (`in_app`, `embed`, `public`, `shared`, `api`, `unknown`). Counts only views captured by Real User Monitoring (RUM); empty in orgs without RUM.
	TotalViewsByType map[string]int64 `json:"total_views_by_type,omitempty"`
	// When the dashboard was most recently viewed. Populated only when Real User Monitoring (RUM) is active for the org; `null` in orgs without RUM.
	ViewedAt datadog.NullableTime `json:"viewed_at,omitempty"`
	// A user referenced from a dashboard usage record (author or viewer).
	Viewer NullableDashboardUsageUser `json:"viewer,omitempty"`
	// The total number of widgets on the dashboard.
	WidgetCount datadog.NullableInt64 `json:"widget_count,omitempty"`
	// Widget counts keyed by widget type. The map includes group widgets and widgets without requests.
	WidgetCountByType map[string]int64 `json:"widget_count_by_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardUsageAttributes instantiates a new DashboardUsageAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardUsageAttributes(orgId int64) *DashboardUsageAttributes {
	this := DashboardUsageAttributes{}
	this.OrgId = orgId
	return &this
}

// NewDashboardUsageAttributesWithDefaults instantiates a new DashboardUsageAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardUsageAttributesWithDefaults() *DashboardUsageAttributes {
	this := DashboardUsageAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetAuthor() DashboardUsageUser {
	if o == nil || o.Author.Get() == nil {
		var ret DashboardUsageUser
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetAuthorOk() (*DashboardUsageUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasAuthor() bool {
	return o != nil && o.Author.IsSet()
}

// SetAuthor gets a reference to the given NullableDashboardUsageUser and assigns it to the Author field.
func (o *DashboardUsageAttributes) SetAuthor(v DashboardUsageUser) {
	o.Author.Set(&v)
}

// SetAuthorNil sets the value for Author to be an explicit nil.
func (o *DashboardUsageAttributes) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetAuthor() {
	o.Author.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt.IsSet()
}

// SetCreatedAt gets a reference to the given datadog.NullableTime and assigns it to the CreatedAt field.
func (o *DashboardUsageAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil.
func (o *DashboardUsageAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDashboardQualityScore returns the DashboardQualityScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetDashboardQualityScore() float64 {
	if o == nil || o.DashboardQualityScore.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DashboardQualityScore.Get()
}

// GetDashboardQualityScoreOk returns a tuple with the DashboardQualityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetDashboardQualityScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardQualityScore.Get(), o.DashboardQualityScore.IsSet()
}

// HasDashboardQualityScore returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasDashboardQualityScore() bool {
	return o != nil && o.DashboardQualityScore.IsSet()
}

// SetDashboardQualityScore gets a reference to the given datadog.NullableFloat64 and assigns it to the DashboardQualityScore field.
func (o *DashboardUsageAttributes) SetDashboardQualityScore(v float64) {
	o.DashboardQualityScore.Set(&v)
}

// SetDashboardQualityScoreNil sets the value for DashboardQualityScore to be an explicit nil.
func (o *DashboardUsageAttributes) SetDashboardQualityScoreNil() {
	o.DashboardQualityScore.Set(nil)
}

// UnsetDashboardQualityScore ensures that no value is present for DashboardQualityScore, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetDashboardQualityScore() {
	o.DashboardQualityScore.Unset()
}

// GetEditedAt returns the EditedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetEditedAt() time.Time {
	if o == nil || o.EditedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EditedAt.Get()
}

// GetEditedAtOk returns a tuple with the EditedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetEditedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditedAt.Get(), o.EditedAt.IsSet()
}

// HasEditedAt returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasEditedAt() bool {
	return o != nil && o.EditedAt.IsSet()
}

// SetEditedAt gets a reference to the given datadog.NullableTime and assigns it to the EditedAt field.
func (o *DashboardUsageAttributes) SetEditedAt(v time.Time) {
	o.EditedAt.Set(&v)
}

// SetEditedAtNil sets the value for EditedAt to be an explicit nil.
func (o *DashboardUsageAttributes) SetEditedAtNil() {
	o.EditedAt.Set(nil)
}

// UnsetEditedAt ensures that no value is present for EditedAt, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetEditedAt() {
	o.EditedAt.Unset()
}

// GetOrgId returns the OrgId field value.
func (o *DashboardUsageAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DashboardUsageAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *DashboardUsageAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetTeams() []string {
	if o == nil || o.Teams.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Teams.Get()
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams.Get(), o.Teams.IsSet()
}

// HasTeams returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasTeams() bool {
	return o != nil && o.Teams.IsSet()
}

// SetTeams gets a reference to the given datadog.NullableList[string] and assigns it to the Teams field.
func (o *DashboardUsageAttributes) SetTeams(v []string) {
	o.Teams.Set(&v)
}

// SetTeamsNil sets the value for Teams to be an explicit nil.
func (o *DashboardUsageAttributes) SetTeamsNil() {
	o.Teams.Set(nil)
}

// UnsetTeams ensures that no value is present for Teams, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetTeams() {
	o.Teams.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DashboardUsageAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DashboardUsageAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetTotalViews returns the TotalViews field value if set, zero value otherwise.
func (o *DashboardUsageAttributes) GetTotalViews() int64 {
	if o == nil || o.TotalViews == nil {
		var ret int64
		return ret
	}
	return *o.TotalViews
}

// GetTotalViewsOk returns a tuple with the TotalViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUsageAttributes) GetTotalViewsOk() (*int64, bool) {
	if o == nil || o.TotalViews == nil {
		return nil, false
	}
	return o.TotalViews, true
}

// HasTotalViews returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasTotalViews() bool {
	return o != nil && o.TotalViews != nil
}

// SetTotalViews gets a reference to the given int64 and assigns it to the TotalViews field.
func (o *DashboardUsageAttributes) SetTotalViews(v int64) {
	o.TotalViews = &v
}

// GetTotalViewsByType returns the TotalViewsByType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetTotalViewsByType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.TotalViewsByType
}

// GetTotalViewsByTypeOk returns a tuple with the TotalViewsByType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetTotalViewsByTypeOk() (*map[string]int64, bool) {
	if o == nil || o.TotalViewsByType == nil {
		return nil, false
	}
	return &o.TotalViewsByType, true
}

// HasTotalViewsByType returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasTotalViewsByType() bool {
	return o != nil && o.TotalViewsByType != nil
}

// SetTotalViewsByType gets a reference to the given map[string]int64 and assigns it to the TotalViewsByType field.
func (o *DashboardUsageAttributes) SetTotalViewsByType(v map[string]int64) {
	o.TotalViewsByType = v
}

// GetViewedAt returns the ViewedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetViewedAt() time.Time {
	if o == nil || o.ViewedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ViewedAt.Get()
}

// GetViewedAtOk returns a tuple with the ViewedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetViewedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewedAt.Get(), o.ViewedAt.IsSet()
}

// HasViewedAt returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasViewedAt() bool {
	return o != nil && o.ViewedAt.IsSet()
}

// SetViewedAt gets a reference to the given datadog.NullableTime and assigns it to the ViewedAt field.
func (o *DashboardUsageAttributes) SetViewedAt(v time.Time) {
	o.ViewedAt.Set(&v)
}

// SetViewedAtNil sets the value for ViewedAt to be an explicit nil.
func (o *DashboardUsageAttributes) SetViewedAtNil() {
	o.ViewedAt.Set(nil)
}

// UnsetViewedAt ensures that no value is present for ViewedAt, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetViewedAt() {
	o.ViewedAt.Unset()
}

// GetViewer returns the Viewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetViewer() DashboardUsageUser {
	if o == nil || o.Viewer.Get() == nil {
		var ret DashboardUsageUser
		return ret
	}
	return *o.Viewer.Get()
}

// GetViewerOk returns a tuple with the Viewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetViewerOk() (*DashboardUsageUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Viewer.Get(), o.Viewer.IsSet()
}

// HasViewer returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasViewer() bool {
	return o != nil && o.Viewer.IsSet()
}

// SetViewer gets a reference to the given NullableDashboardUsageUser and assigns it to the Viewer field.
func (o *DashboardUsageAttributes) SetViewer(v DashboardUsageUser) {
	o.Viewer.Set(&v)
}

// SetViewerNil sets the value for Viewer to be an explicit nil.
func (o *DashboardUsageAttributes) SetViewerNil() {
	o.Viewer.Set(nil)
}

// UnsetViewer ensures that no value is present for Viewer, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetViewer() {
	o.Viewer.Unset()
}

// GetWidgetCount returns the WidgetCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetWidgetCount() int64 {
	if o == nil || o.WidgetCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.WidgetCount.Get()
}

// GetWidgetCountOk returns a tuple with the WidgetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetWidgetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WidgetCount.Get(), o.WidgetCount.IsSet()
}

// HasWidgetCount returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasWidgetCount() bool {
	return o != nil && o.WidgetCount.IsSet()
}

// SetWidgetCount gets a reference to the given datadog.NullableInt64 and assigns it to the WidgetCount field.
func (o *DashboardUsageAttributes) SetWidgetCount(v int64) {
	o.WidgetCount.Set(&v)
}

// SetWidgetCountNil sets the value for WidgetCount to be an explicit nil.
func (o *DashboardUsageAttributes) SetWidgetCountNil() {
	o.WidgetCount.Set(nil)
}

// UnsetWidgetCount ensures that no value is present for WidgetCount, not even an explicit nil.
func (o *DashboardUsageAttributes) UnsetWidgetCount() {
	o.WidgetCount.Unset()
}

// GetWidgetCountByType returns the WidgetCountByType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardUsageAttributes) GetWidgetCountByType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.WidgetCountByType
}

// GetWidgetCountByTypeOk returns a tuple with the WidgetCountByType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardUsageAttributes) GetWidgetCountByTypeOk() (*map[string]int64, bool) {
	if o == nil || o.WidgetCountByType == nil {
		return nil, false
	}
	return &o.WidgetCountByType, true
}

// HasWidgetCountByType returns a boolean if a field has been set.
func (o *DashboardUsageAttributes) HasWidgetCountByType() bool {
	return o != nil && o.WidgetCountByType != nil
}

// SetWidgetCountByType gets a reference to the given map[string]int64 and assigns it to the WidgetCountByType field.
func (o *DashboardUsageAttributes) SetWidgetCountByType(v map[string]int64) {
	o.WidgetCountByType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardUsageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DashboardQualityScore.IsSet() {
		toSerialize["dashboard_quality_score"] = o.DashboardQualityScore.Get()
	}
	if o.EditedAt.IsSet() {
		toSerialize["edited_at"] = o.EditedAt.Get()
	}
	toSerialize["org_id"] = o.OrgId
	if o.Teams.IsSet() {
		toSerialize["teams"] = o.Teams.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TotalViews != nil {
		toSerialize["total_views"] = o.TotalViews
	}
	if o.TotalViewsByType != nil {
		toSerialize["total_views_by_type"] = o.TotalViewsByType
	}
	if o.ViewedAt.IsSet() {
		toSerialize["viewed_at"] = o.ViewedAt.Get()
	}
	if o.Viewer.IsSet() {
		toSerialize["viewer"] = o.Viewer.Get()
	}
	if o.WidgetCount.IsSet() {
		toSerialize["widget_count"] = o.WidgetCount.Get()
	}
	if o.WidgetCountByType != nil {
		toSerialize["widget_count_by_type"] = o.WidgetCountByType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardUsageAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                NullableDashboardUsageUser   `json:"author,omitempty"`
		CreatedAt             datadog.NullableTime         `json:"created_at,omitempty"`
		DashboardQualityScore datadog.NullableFloat64      `json:"dashboard_quality_score,omitempty"`
		EditedAt              datadog.NullableTime         `json:"edited_at,omitempty"`
		OrgId                 *int64                       `json:"org_id"`
		Teams                 datadog.NullableList[string] `json:"teams,omitempty"`
		Title                 *string                      `json:"title,omitempty"`
		TotalViews            *int64                       `json:"total_views,omitempty"`
		TotalViewsByType      map[string]int64             `json:"total_views_by_type,omitempty"`
		ViewedAt              datadog.NullableTime         `json:"viewed_at,omitempty"`
		Viewer                NullableDashboardUsageUser   `json:"viewer,omitempty"`
		WidgetCount           datadog.NullableInt64        `json:"widget_count,omitempty"`
		WidgetCountByType     map[string]int64             `json:"widget_count_by_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "dashboard_quality_score", "edited_at", "org_id", "teams", "title", "total_views", "total_views_by_type", "viewed_at", "viewer", "widget_count", "widget_count_by_type"})
	} else {
		return err
	}
	o.Author = all.Author
	o.CreatedAt = all.CreatedAt
	o.DashboardQualityScore = all.DashboardQualityScore
	o.EditedAt = all.EditedAt
	o.OrgId = *all.OrgId
	o.Teams = all.Teams
	o.Title = all.Title
	o.TotalViews = all.TotalViews
	o.TotalViewsByType = all.TotalViewsByType
	o.ViewedAt = all.ViewedAt
	o.Viewer = all.Viewer
	o.WidgetCount = all.WidgetCount
	o.WidgetCountByType = all.WidgetCountByType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
