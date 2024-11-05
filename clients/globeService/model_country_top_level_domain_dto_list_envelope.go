/*
GlobeService

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

// checks if the CountryTopLevelDomainDtoListEnvelope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryTopLevelDomainDtoListEnvelope{}

// CountryTopLevelDomainDtoListEnvelope struct for CountryTopLevelDomainDtoListEnvelope
type CountryTopLevelDomainDtoListEnvelope struct {
	IsSuccess *bool `json:"isSuccess,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	CorrelationId NullableString `json:"correlationId,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	ActivityId NullableString `json:"activityId,omitempty"`
	Result []CountryTopLevelDomainDto `json:"result,omitempty"`
}

// NewCountryTopLevelDomainDtoListEnvelope instantiates a new CountryTopLevelDomainDtoListEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryTopLevelDomainDtoListEnvelope() *CountryTopLevelDomainDtoListEnvelope {
	this := CountryTopLevelDomainDtoListEnvelope{}
	return &this
}

// NewCountryTopLevelDomainDtoListEnvelopeWithDefaults instantiates a new CountryTopLevelDomainDtoListEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryTopLevelDomainDtoListEnvelopeWithDefaults() *CountryTopLevelDomainDtoListEnvelope {
	this := CountryTopLevelDomainDtoListEnvelope{}
	return &this
}

// GetIsSuccess returns the IsSuccess field value if set, zero value otherwise.
func (o *CountryTopLevelDomainDtoListEnvelope) GetIsSuccess() bool {
	if o == nil || IsNil(o.IsSuccess) {
		var ret bool
		return ret
	}
	return *o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) GetIsSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuccess) {
		return nil, false
	}
	return o.IsSuccess, true
}

// HasIsSuccess returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasIsSuccess() bool {
	if o != nil && !IsNil(o.IsSuccess) {
		return true
	}

	return false
}

// SetIsSuccess gets a reference to the given bool and assigns it to the IsSuccess field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetIsSuccess(v bool) {
	o.IsSuccess = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryTopLevelDomainDtoListEnvelope) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryTopLevelDomainDtoListEnvelope) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryTopLevelDomainDtoListEnvelope) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId.Get()) {
		var ret string
		return ret
	}
	return *o.CorrelationId.Get()
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryTopLevelDomainDtoListEnvelope) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrelationId.Get(), o.CorrelationId.IsSet()
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasCorrelationId() bool {
	if o != nil && o.CorrelationId.IsSet() {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given NullableString and assigns it to the CorrelationId field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetCorrelationId(v string) {
	o.CorrelationId.Set(&v)
}
// SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) SetCorrelationIdNil() {
	o.CorrelationId.Set(nil)
}

// UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) UnsetCorrelationId() {
	o.CorrelationId.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CountryTopLevelDomainDtoListEnvelope) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryTopLevelDomainDtoListEnvelope) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryTopLevelDomainDtoListEnvelope) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *CountryTopLevelDomainDtoListEnvelope) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryTopLevelDomainDtoListEnvelope) GetResult() []CountryTopLevelDomainDto {
	if o == nil {
		var ret []CountryTopLevelDomainDto
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryTopLevelDomainDtoListEnvelope) GetResultOk() ([]CountryTopLevelDomainDto, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CountryTopLevelDomainDtoListEnvelope) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []CountryTopLevelDomainDto and assigns it to the Result field.
func (o *CountryTopLevelDomainDtoListEnvelope) SetResult(v []CountryTopLevelDomainDto) {
	o.Result = v
}

func (o CountryTopLevelDomainDtoListEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryTopLevelDomainDtoListEnvelope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSuccess) {
		toSerialize["isSuccess"] = o.IsSuccess
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.CorrelationId.IsSet() {
		toSerialize["correlationId"] = o.CorrelationId.Get()
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCountryTopLevelDomainDtoListEnvelope struct {
	value *CountryTopLevelDomainDtoListEnvelope
	isSet bool
}

func (v NullableCountryTopLevelDomainDtoListEnvelope) Get() *CountryTopLevelDomainDtoListEnvelope {
	return v.value
}

func (v *NullableCountryTopLevelDomainDtoListEnvelope) Set(val *CountryTopLevelDomainDtoListEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryTopLevelDomainDtoListEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryTopLevelDomainDtoListEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryTopLevelDomainDtoListEnvelope(val *CountryTopLevelDomainDtoListEnvelope) *NullableCountryTopLevelDomainDtoListEnvelope {
	return &NullableCountryTopLevelDomainDtoListEnvelope{value: val, isSet: true}
}

func (v NullableCountryTopLevelDomainDtoListEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryTopLevelDomainDtoListEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


