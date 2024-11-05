/*
WalletsService

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

// checks if the TenantEnrolmentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantEnrolmentDto{}

// TenantEnrolmentDto struct for TenantEnrolmentDto
type TenantEnrolmentDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	UserId NullableString `json:"userId,omitempty"`
	IsRoot *bool `json:"isRoot,omitempty"`
	IsOwner *bool `json:"isOwner,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
}

// NewTenantEnrolmentDto instantiates a new TenantEnrolmentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantEnrolmentDto() *TenantEnrolmentDto {
	this := TenantEnrolmentDto{}
	return &this
}

// NewTenantEnrolmentDtoWithDefaults instantiates a new TenantEnrolmentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantEnrolmentDtoWithDefaults() *TenantEnrolmentDto {
	this := TenantEnrolmentDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantEnrolmentDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantEnrolmentDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TenantEnrolmentDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TenantEnrolmentDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TenantEnrolmentDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantEnrolmentDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantEnrolmentDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *TenantEnrolmentDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *TenantEnrolmentDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *TenantEnrolmentDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantEnrolmentDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantEnrolmentDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *TenantEnrolmentDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *TenantEnrolmentDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *TenantEnrolmentDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantEnrolmentDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantEnrolmentDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *TenantEnrolmentDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *TenantEnrolmentDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *TenantEnrolmentDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetIsRoot returns the IsRoot field value if set, zero value otherwise.
func (o *TenantEnrolmentDto) GetIsRoot() bool {
	if o == nil || IsNil(o.IsRoot) {
		var ret bool
		return ret
	}
	return *o.IsRoot
}

// GetIsRootOk returns a tuple with the IsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantEnrolmentDto) GetIsRootOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoot) {
		return nil, false
	}
	return o.IsRoot, true
}

// HasIsRoot returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasIsRoot() bool {
	if o != nil && !IsNil(o.IsRoot) {
		return true
	}

	return false
}

// SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.
func (o *TenantEnrolmentDto) SetIsRoot(v bool) {
	o.IsRoot = &v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *TenantEnrolmentDto) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantEnrolmentDto) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasIsOwner() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *TenantEnrolmentDto) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *TenantEnrolmentDto) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantEnrolmentDto) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *TenantEnrolmentDto) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *TenantEnrolmentDto) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantEnrolmentDto) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *TenantEnrolmentDto) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *TenantEnrolmentDto) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

func (o TenantEnrolmentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantEnrolmentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if !IsNil(o.IsRoot) {
		toSerialize["isRoot"] = o.IsRoot
	}
	if !IsNil(o.IsOwner) {
		toSerialize["isOwner"] = o.IsOwner
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	return toSerialize, nil
}

type NullableTenantEnrolmentDto struct {
	value *TenantEnrolmentDto
	isSet bool
}

func (v NullableTenantEnrolmentDto) Get() *TenantEnrolmentDto {
	return v.value
}

func (v *NullableTenantEnrolmentDto) Set(val *TenantEnrolmentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantEnrolmentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantEnrolmentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantEnrolmentDto(val *TenantEnrolmentDto) *NullableTenantEnrolmentDto {
	return &NullableTenantEnrolmentDto{value: val, isSet: true}
}

func (v NullableTenantEnrolmentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantEnrolmentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

