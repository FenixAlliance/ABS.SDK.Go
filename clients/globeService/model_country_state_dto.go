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

// checks if the CountryStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryStateDto{}

// CountryStateDto struct for CountryStateDto
type CountryStateDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	CountryID NullableString `json:"countryID,omitempty"`
}

// NewCountryStateDto instantiates a new CountryStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryStateDto() *CountryStateDto {
	this := CountryStateDto{}
	return &this
}

// NewCountryStateDtoWithDefaults instantiates a new CountryStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryStateDtoWithDefaults() *CountryStateDto {
	this := CountryStateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryStateDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryStateDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CountryStateDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CountryStateDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CountryStateDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CountryStateDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryStateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryStateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CountryStateDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *CountryStateDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *CountryStateDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *CountryStateDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryStateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryStateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CountryStateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CountryStateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CountryStateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CountryStateDto) UnsetName() {
	o.Name.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryStateDto) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryStateDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *CountryStateDto) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *CountryStateDto) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *CountryStateDto) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *CountryStateDto) UnsetCode() {
	o.Code.Unset()
}

// GetCountryID returns the CountryID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryStateDto) GetCountryID() string {
	if o == nil || IsNil(o.CountryID.Get()) {
		var ret string
		return ret
	}
	return *o.CountryID.Get()
}

// GetCountryIDOk returns a tuple with the CountryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryStateDto) GetCountryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryID.Get(), o.CountryID.IsSet()
}

// HasCountryID returns a boolean if a field has been set.
func (o *CountryStateDto) HasCountryID() bool {
	if o != nil && o.CountryID.IsSet() {
		return true
	}

	return false
}

// SetCountryID gets a reference to the given NullableString and assigns it to the CountryID field.
func (o *CountryStateDto) SetCountryID(v string) {
	o.CountryID.Set(&v)
}
// SetCountryIDNil sets the value for CountryID to be an explicit nil
func (o *CountryStateDto) SetCountryIDNil() {
	o.CountryID.Set(nil)
}

// UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
func (o *CountryStateDto) UnsetCountryID() {
	o.CountryID.Unset()
}

func (o CountryStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.CountryID.IsSet() {
		toSerialize["countryID"] = o.CountryID.Get()
	}
	return toSerialize, nil
}

type NullableCountryStateDto struct {
	value *CountryStateDto
	isSet bool
}

func (v NullableCountryStateDto) Get() *CountryStateDto {
	return v.value
}

func (v *NullableCountryStateDto) Set(val *CountryStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryStateDto(val *CountryStateDto) *NullableCountryStateDto {
	return &NullableCountryStateDto{value: val, isSet: true}
}

func (v NullableCountryStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

