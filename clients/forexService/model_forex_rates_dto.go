/*
ForexService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ForexRatesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForexRatesDto{}

// ForexRatesDto struct for ForexRatesDto
type ForexRatesDto struct {
	Success *bool `json:"success,omitempty"`
	Date NullableString `json:"date,omitempty"`
	Base NullableString `json:"base,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	RequestTimestamp *time.Time `json:"requestTimestamp,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

// NewForexRatesDto instantiates a new ForexRatesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForexRatesDto() *ForexRatesDto {
	this := ForexRatesDto{}
	return &this
}

// NewForexRatesDtoWithDefaults instantiates a new ForexRatesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForexRatesDtoWithDefaults() *ForexRatesDto {
	this := ForexRatesDto{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ForexRatesDto) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexRatesDto) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ForexRatesDto) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ForexRatesDto) SetSuccess(v bool) {
	o.Success = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForexRatesDto) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForexRatesDto) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ForexRatesDto) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *ForexRatesDto) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ForexRatesDto) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ForexRatesDto) UnsetDate() {
	o.Date.Unset()
}

// GetBase returns the Base field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForexRatesDto) GetBase() string {
	if o == nil || IsNil(o.Base.Get()) {
		var ret string
		return ret
	}
	return *o.Base.Get()
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForexRatesDto) GetBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Base.Get(), o.Base.IsSet()
}

// HasBase returns a boolean if a field has been set.
func (o *ForexRatesDto) HasBase() bool {
	if o != nil && o.Base.IsSet() {
		return true
	}

	return false
}

// SetBase gets a reference to the given NullableString and assigns it to the Base field.
func (o *ForexRatesDto) SetBase(v string) {
	o.Base.Set(&v)
}
// SetBaseNil sets the value for Base to be an explicit nil
func (o *ForexRatesDto) SetBaseNil() {
	o.Base.Set(nil)
}

// UnsetBase ensures that no value is present for Base, not even an explicit nil
func (o *ForexRatesDto) UnsetBase() {
	o.Base.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ForexRatesDto) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexRatesDto) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ForexRatesDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ForexRatesDto) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *ForexRatesDto) GetRequestTimestamp() time.Time {
	if o == nil || IsNil(o.RequestTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexRatesDto) GetRequestTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestTimestamp) {
		return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *ForexRatesDto) HasRequestTimestamp() bool {
	if o != nil && !IsNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given time.Time and assigns it to the RequestTimestamp field.
func (o *ForexRatesDto) SetRequestTimestamp(v time.Time) {
	o.RequestTimestamp = &v
}

// GetRates returns the Rates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForexRatesDto) GetRates() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForexRatesDto) GetRatesOk() (*map[string]float64, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return &o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *ForexRatesDto) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given map[string]float64 and assigns it to the Rates field.
func (o *ForexRatesDto) SetRates(v map[string]float64) {
	o.Rates = v
}

func (o ForexRatesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForexRatesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Base.IsSet() {
		toSerialize["base"] = o.Base.Get()
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.RequestTimestamp) {
		toSerialize["requestTimestamp"] = o.RequestTimestamp
	}
	if o.Rates != nil {
		toSerialize["rates"] = o.Rates
	}
	return toSerialize, nil
}

type NullableForexRatesDto struct {
	value *ForexRatesDto
	isSet bool
}

func (v NullableForexRatesDto) Get() *ForexRatesDto {
	return v.value
}

func (v *NullableForexRatesDto) Set(val *ForexRatesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableForexRatesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableForexRatesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForexRatesDto(val *ForexRatesDto) *NullableForexRatesDto {
	return &NullableForexRatesDto{value: val, isSet: true}
}

func (v NullableForexRatesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForexRatesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


