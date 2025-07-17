// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerSamplings Sampling configurations for the Scanning Group.
type SensitiveDataScannerSamplings struct {
	// Datadog product onto which Sensitive Data Scanner can be activated.
	Product *SensitiveDataScannerProduct `json:"product,omitempty"`
	// Rate at which data in product type will be scanned, as a percentage.
	Rate *float64 `json:"rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerSamplings instantiates a new SensitiveDataScannerSamplings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerSamplings() *SensitiveDataScannerSamplings {
	this := SensitiveDataScannerSamplings{}
	var product SensitiveDataScannerProduct = SENSITIVEDATASCANNERPRODUCT_LOGS
	this.Product = &product
	return &this
}

// NewSensitiveDataScannerSamplingsWithDefaults instantiates a new SensitiveDataScannerSamplings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerSamplingsWithDefaults() *SensitiveDataScannerSamplings {
	this := SensitiveDataScannerSamplings{}
	var product SensitiveDataScannerProduct = SENSITIVEDATASCANNERPRODUCT_LOGS
	this.Product = &product
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SensitiveDataScannerSamplings) GetProduct() SensitiveDataScannerProduct {
	if o == nil || o.Product == nil {
		var ret SensitiveDataScannerProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerSamplings) GetProductOk() (*SensitiveDataScannerProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SensitiveDataScannerSamplings) HasProduct() bool {
	return o != nil && o.Product != nil
}

// SetProduct gets a reference to the given SensitiveDataScannerProduct and assigns it to the Product field.
func (o *SensitiveDataScannerSamplings) SetProduct(v SensitiveDataScannerProduct) {
	o.Product = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *SensitiveDataScannerSamplings) GetRate() float64 {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerSamplings) GetRateOk() (*float64, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *SensitiveDataScannerSamplings) HasRate() bool {
	return o != nil && o.Rate != nil
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *SensitiveDataScannerSamplings) SetRate(v float64) {
	o.Rate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerSamplings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerSamplings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Product *SensitiveDataScannerProduct `json:"product,omitempty"`
		Rate    *float64                     `json:"rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"product", "rate"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Product != nil && !all.Product.IsValid() {
		hasInvalidField = true
	} else {
		o.Product = all.Product
	}
	o.Rate = all.Rate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
