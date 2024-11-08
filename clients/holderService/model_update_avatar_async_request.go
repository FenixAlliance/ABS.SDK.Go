/*
HolderService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// checks if the UpdateAvatarAsyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAvatarAsyncRequest{}

// UpdateAvatarAsyncRequest struct for UpdateAvatarAsyncRequest
type UpdateAvatarAsyncRequest struct {
	Avatar **os.File `json:"avatar,omitempty"`
}

// NewUpdateAvatarAsyncRequest instantiates a new UpdateAvatarAsyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAvatarAsyncRequest() *UpdateAvatarAsyncRequest {
	this := UpdateAvatarAsyncRequest{}
	return &this
}

// NewUpdateAvatarAsyncRequestWithDefaults instantiates a new UpdateAvatarAsyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAvatarAsyncRequestWithDefaults() *UpdateAvatarAsyncRequest {
	this := UpdateAvatarAsyncRequest{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UpdateAvatarAsyncRequest) GetAvatar() *os.File {
	if o == nil || IsNil(o.Avatar) {
		var ret *os.File
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarAsyncRequest) GetAvatarOk() (**os.File, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UpdateAvatarAsyncRequest) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given *os.File and assigns it to the Avatar field.
func (o *UpdateAvatarAsyncRequest) SetAvatar(v *os.File) {
	o.Avatar = &v
}

func (o UpdateAvatarAsyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAvatarAsyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	return toSerialize, nil
}

type NullableUpdateAvatarAsyncRequest struct {
	value *UpdateAvatarAsyncRequest
	isSet bool
}

func (v NullableUpdateAvatarAsyncRequest) Get() *UpdateAvatarAsyncRequest {
	return v.value
}

func (v *NullableUpdateAvatarAsyncRequest) Set(val *UpdateAvatarAsyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAvatarAsyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAvatarAsyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAvatarAsyncRequest(val *UpdateAvatarAsyncRequest) *NullableUpdateAvatarAsyncRequest {
	return &NullableUpdateAvatarAsyncRequest{value: val, isSet: true}
}

func (v NullableUpdateAvatarAsyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAvatarAsyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


