// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsContentBlock A single content block rendered inside a `display_block` interaction.
// `type` discriminates which other fields are meaningful:
//
// - `markdown` / `text`: `content` must be a string.
// - `header`: `content` must be a string; `level`, when set, must be one of `sm`, `md`, `lg`, `xl`.
// - `json`: `content` must be a well-formed JSON value (object, array, or scalar).
// - `image`: `url` is required.
// - `widget`: `tileDef` is required (any well-formed JSON; the frontend owns the renderable schema).
// - `llmobs_trace`: `traceId` is required; `interactionType`, when set, must be `trace` or `experiment_trace`.
//
// `height`, when set, must be positive.
type LLMObsContentBlock struct {
	// Alternative text for an `image` block.
	Alt *string `json:"alt,omitempty"`
	// Block payload. A string for `markdown`, `header`, and `text`; an
	// arbitrary JSON value (object, array, or scalar) for `json`. Omitted
	// for `image`, `widget`, and `llmobs_trace`.
	Content interface{} `json:"content,omitempty"`
	// Optional rendered height. Must be positive when set.
	Height *int64 `json:"height,omitempty"`
	// Upstream interaction type referenced by an `llmobs_trace` block.
	// Restricted to `trace` or `experiment_trace`.
	InteractionType *LLMObsContentBlockLLMObsTraceInteractionType `json:"interactionType,omitempty"`
	// Optional label rendered alongside the block.
	Label *string `json:"label,omitempty"`
	// Visual size for a `header` block.
	Level *LLMObsContentBlockHeaderLevel `json:"level,omitempty"`
	// Tile definition for a `widget` block. Required for `widget`. The
	// schema is owned by the frontend renderer.
	TileDef interface{} `json:"tileDef,omitempty"`
	// Unix-millis time range used by chart blocks.
	TimeFrame *LLMObsContentBlockTimeFrame `json:"timeFrame,omitempty"`
	// Trace identifier. Required for `llmobs_trace` blocks.
	TraceId *string `json:"traceId,omitempty"`
	// Discriminator for a single `display_block` content block. Adding a
	// variant requires coordinated changes in the frontend renderer.
	Type LLMObsContentBlockType `json:"type"`
	// URL of the image. Required for `image` blocks.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsContentBlock instantiates a new LLMObsContentBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsContentBlock(typeVar LLMObsContentBlockType) *LLMObsContentBlock {
	this := LLMObsContentBlock{}
	this.Type = typeVar
	return &this
}

// NewLLMObsContentBlockWithDefaults instantiates a new LLMObsContentBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsContentBlockWithDefaults() *LLMObsContentBlock {
	this := LLMObsContentBlock{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetAlt() string {
	if o == nil || o.Alt == nil {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetAltOk() (*string, bool) {
	if o == nil || o.Alt == nil {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasAlt() bool {
	return o != nil && o.Alt != nil
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *LLMObsContentBlock) SetAlt(v string) {
	o.Alt = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetContent() interface{} {
	if o == nil || o.Content == nil {
		var ret interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetContentOk() (*interface{}, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given interface{} and assigns it to the Content field.
func (o *LLMObsContentBlock) SetContent(v interface{}) {
	o.Content = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *LLMObsContentBlock) SetHeight(v int64) {
	o.Height = &v
}

// GetInteractionType returns the InteractionType field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetInteractionType() LLMObsContentBlockLLMObsTraceInteractionType {
	if o == nil || o.InteractionType == nil {
		var ret LLMObsContentBlockLLMObsTraceInteractionType
		return ret
	}
	return *o.InteractionType
}

// GetInteractionTypeOk returns a tuple with the InteractionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetInteractionTypeOk() (*LLMObsContentBlockLLMObsTraceInteractionType, bool) {
	if o == nil || o.InteractionType == nil {
		return nil, false
	}
	return o.InteractionType, true
}

// HasInteractionType returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasInteractionType() bool {
	return o != nil && o.InteractionType != nil
}

// SetInteractionType gets a reference to the given LLMObsContentBlockLLMObsTraceInteractionType and assigns it to the InteractionType field.
func (o *LLMObsContentBlock) SetInteractionType(v LLMObsContentBlockLLMObsTraceInteractionType) {
	o.InteractionType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *LLMObsContentBlock) SetLabel(v string) {
	o.Label = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetLevel() LLMObsContentBlockHeaderLevel {
	if o == nil || o.Level == nil {
		var ret LLMObsContentBlockHeaderLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetLevelOk() (*LLMObsContentBlockHeaderLevel, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given LLMObsContentBlockHeaderLevel and assigns it to the Level field.
func (o *LLMObsContentBlock) SetLevel(v LLMObsContentBlockHeaderLevel) {
	o.Level = &v
}

// GetTileDef returns the TileDef field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetTileDef() interface{} {
	if o == nil || o.TileDef == nil {
		var ret interface{}
		return ret
	}
	return o.TileDef
}

// GetTileDefOk returns a tuple with the TileDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetTileDefOk() (*interface{}, bool) {
	if o == nil || o.TileDef == nil {
		return nil, false
	}
	return &o.TileDef, true
}

// HasTileDef returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasTileDef() bool {
	return o != nil && o.TileDef != nil
}

// SetTileDef gets a reference to the given interface{} and assigns it to the TileDef field.
func (o *LLMObsContentBlock) SetTileDef(v interface{}) {
	o.TileDef = v
}

// GetTimeFrame returns the TimeFrame field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetTimeFrame() LLMObsContentBlockTimeFrame {
	if o == nil || o.TimeFrame == nil {
		var ret LLMObsContentBlockTimeFrame
		return ret
	}
	return *o.TimeFrame
}

// GetTimeFrameOk returns a tuple with the TimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetTimeFrameOk() (*LLMObsContentBlockTimeFrame, bool) {
	if o == nil || o.TimeFrame == nil {
		return nil, false
	}
	return o.TimeFrame, true
}

// HasTimeFrame returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasTimeFrame() bool {
	return o != nil && o.TimeFrame != nil
}

// SetTimeFrame gets a reference to the given LLMObsContentBlockTimeFrame and assigns it to the TimeFrame field.
func (o *LLMObsContentBlock) SetTimeFrame(v LLMObsContentBlockTimeFrame) {
	o.TimeFrame = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *LLMObsContentBlock) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value.
func (o *LLMObsContentBlock) GetType() LLMObsContentBlockType {
	if o == nil {
		var ret LLMObsContentBlockType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetTypeOk() (*LLMObsContentBlockType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsContentBlock) SetType(v LLMObsContentBlockType) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LLMObsContentBlock) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsContentBlock) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LLMObsContentBlock) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LLMObsContentBlock) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsContentBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alt != nil {
		toSerialize["alt"] = o.Alt
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.InteractionType != nil {
		toSerialize["interactionType"] = o.InteractionType
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.TileDef != nil {
		toSerialize["tileDef"] = o.TileDef
	}
	if o.TimeFrame != nil {
		toSerialize["timeFrame"] = o.TimeFrame
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
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
func (o *LLMObsContentBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alt             *string                                       `json:"alt,omitempty"`
		Content         interface{}                                   `json:"content,omitempty"`
		Height          *int64                                        `json:"height,omitempty"`
		InteractionType *LLMObsContentBlockLLMObsTraceInteractionType `json:"interactionType,omitempty"`
		Label           *string                                       `json:"label,omitempty"`
		Level           *LLMObsContentBlockHeaderLevel                `json:"level,omitempty"`
		TileDef         interface{}                                   `json:"tileDef,omitempty"`
		TimeFrame       *LLMObsContentBlockTimeFrame                  `json:"timeFrame,omitempty"`
		TraceId         *string                                       `json:"traceId,omitempty"`
		Type            *LLMObsContentBlockType                       `json:"type"`
		Url             *string                                       `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alt", "content", "height", "interactionType", "label", "level", "tileDef", "timeFrame", "traceId", "type", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alt = all.Alt
	o.Content = all.Content
	o.Height = all.Height
	if all.InteractionType != nil && !all.InteractionType.IsValid() {
		hasInvalidField = true
	} else {
		o.InteractionType = all.InteractionType
	}
	o.Label = all.Label
	if all.Level != nil && !all.Level.IsValid() {
		hasInvalidField = true
	} else {
		o.Level = all.Level
	}
	o.TileDef = all.TileDef
	if all.TimeFrame != nil && all.TimeFrame.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeFrame = all.TimeFrame
	o.TraceId = all.TraceId
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
