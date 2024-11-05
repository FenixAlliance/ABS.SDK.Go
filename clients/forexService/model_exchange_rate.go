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
)

// checks if the ExchangeRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeRate{}

// ExchangeRate struct for ExchangeRate
type ExchangeRate struct {
	Source *Money `json:"source,omitempty"`
	Target *Money `json:"target,omitempty"`
	Rate *Money `json:"rate,omitempty"`
}

// NewExchangeRate instantiates a new ExchangeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRate() *ExchangeRate {
	this := ExchangeRate{}
	return &this
}

// NewExchangeRateWithDefaults instantiates a new ExchangeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRateWithDefaults() *ExchangeRate {
	this := ExchangeRate{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ExchangeRate) GetSource() Money {
	if o == nil || IsNil(o.Source) {
		var ret Money
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRate) GetSourceOk() (*Money, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ExchangeRate) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Money and assigns it to the Source field.
func (o *ExchangeRate) SetSource(v Money) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ExchangeRate) GetTarget() Money {
	if o == nil || IsNil(o.Target) {
		var ret Money
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRate) GetTargetOk() (*Money, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ExchangeRate) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Money and assigns it to the Target field.
func (o *ExchangeRate) SetTarget(v Money) {
	o.Target = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ExchangeRate) GetRate() Money {
	if o == nil || IsNil(o.Rate) {
		var ret Money
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRate) GetRateOk() (*Money, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ExchangeRate) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given Money and assigns it to the Rate field.
func (o *ExchangeRate) SetRate(v Money) {
	o.Rate = &v
}

func (o ExchangeRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableExchangeRate struct {
	value *ExchangeRate
	isSet bool
}

func (v NullableExchangeRate) Get() *ExchangeRate {
	return v.value
}

func (v *NullableExchangeRate) Set(val *ExchangeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRate(val *ExchangeRate) *NullableExchangeRate {
	return &NullableExchangeRate{value: val, isSet: true}
}

func (v NullableExchangeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

