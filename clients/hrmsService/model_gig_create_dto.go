/*
HrmsService

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

// checks if the GigCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GigCreateDto{}

// GigCreateDto struct for GigCreateDto
type GigCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewGigCreateDto instantiates a new GigCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGigCreateDto() *GigCreateDto {
	this := GigCreateDto{}
	return &this
}

// NewGigCreateDtoWithDefaults instantiates a new GigCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGigCreateDtoWithDefaults() *GigCreateDto {
	this := GigCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GigCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GigCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GigCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GigCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GigCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GigCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GigCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *GigCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o GigCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GigCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableGigCreateDto struct {
	value *GigCreateDto
	isSet bool
}

func (v NullableGigCreateDto) Get() *GigCreateDto {
	return v.value
}

func (v *NullableGigCreateDto) Set(val *GigCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGigCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGigCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGigCreateDto(val *GigCreateDto) *NullableGigCreateDto {
	return &NullableGigCreateDto{value: val, isSet: true}
}

func (v NullableGigCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGigCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


