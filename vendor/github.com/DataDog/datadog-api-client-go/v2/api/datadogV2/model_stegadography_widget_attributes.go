// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StegadographyWidgetAttributes Attributes of a watermarked widget recovered from an image.
type StegadographyWidgetAttributes struct {
	// Horizontal pixel coordinate where the watermark was found in the image.
	Locationx int64 `json:"locationx"`
	// Vertical pixel coordinate where the watermark was found in the image.
	Locationy int64 `json:"locationy"`
	// JSON-encoded string representing the widget state.
	RawData string `json:"rawData"`
	// Hex-encoded watermark string identifying the widget.
	Watermark string `json:"watermark"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStegadographyWidgetAttributes instantiates a new StegadographyWidgetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStegadographyWidgetAttributes(locationx int64, locationy int64, rawData string, watermark string) *StegadographyWidgetAttributes {
	this := StegadographyWidgetAttributes{}
	this.Locationx = locationx
	this.Locationy = locationy
	this.RawData = rawData
	this.Watermark = watermark
	return &this
}

// NewStegadographyWidgetAttributesWithDefaults instantiates a new StegadographyWidgetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStegadographyWidgetAttributesWithDefaults() *StegadographyWidgetAttributes {
	this := StegadographyWidgetAttributes{}
	return &this
}

// GetLocationx returns the Locationx field value.
func (o *StegadographyWidgetAttributes) GetLocationx() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Locationx
}

// GetLocationxOk returns a tuple with the Locationx field value
// and a boolean to check if the value has been set.
func (o *StegadographyWidgetAttributes) GetLocationxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locationx, true
}

// SetLocationx sets field value.
func (o *StegadographyWidgetAttributes) SetLocationx(v int64) {
	o.Locationx = v
}

// GetLocationy returns the Locationy field value.
func (o *StegadographyWidgetAttributes) GetLocationy() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Locationy
}

// GetLocationyOk returns a tuple with the Locationy field value
// and a boolean to check if the value has been set.
func (o *StegadographyWidgetAttributes) GetLocationyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locationy, true
}

// SetLocationy sets field value.
func (o *StegadographyWidgetAttributes) SetLocationy(v int64) {
	o.Locationy = v
}

// GetRawData returns the RawData field value.
func (o *StegadographyWidgetAttributes) GetRawData() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value
// and a boolean to check if the value has been set.
func (o *StegadographyWidgetAttributes) GetRawDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawData, true
}

// SetRawData sets field value.
func (o *StegadographyWidgetAttributes) SetRawData(v string) {
	o.RawData = v
}

// GetWatermark returns the Watermark field value.
func (o *StegadographyWidgetAttributes) GetWatermark() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value
// and a boolean to check if the value has been set.
func (o *StegadographyWidgetAttributes) GetWatermarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Watermark, true
}

// SetWatermark sets field value.
func (o *StegadographyWidgetAttributes) SetWatermark(v string) {
	o.Watermark = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StegadographyWidgetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["locationx"] = o.Locationx
	toSerialize["locationy"] = o.Locationy
	toSerialize["rawData"] = o.RawData
	toSerialize["watermark"] = o.Watermark

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StegadographyWidgetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Locationx *int64  `json:"locationx"`
		Locationy *int64  `json:"locationy"`
		RawData   *string `json:"rawData"`
		Watermark *string `json:"watermark"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Locationx == nil {
		return fmt.Errorf("required field locationx missing")
	}
	if all.Locationy == nil {
		return fmt.Errorf("required field locationy missing")
	}
	if all.RawData == nil {
		return fmt.Errorf("required field rawData missing")
	}
	if all.Watermark == nil {
		return fmt.Errorf("required field watermark missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"locationx", "locationy", "rawData", "watermark"})
	} else {
		return err
	}
	o.Locationx = *all.Locationx
	o.Locationy = *all.Locationy
	o.RawData = *all.RawData
	o.Watermark = *all.Watermark

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
