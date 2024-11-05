/*
SystemService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LicenseKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseKey{}

// LicenseKey struct for LicenseKey
type LicenseKey struct {
	Key NullableString `json:"key,omitempty"`
}

// NewLicenseKey instantiates a new LicenseKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseKey() *LicenseKey {
	this := LicenseKey{}
	return &this
}

// NewLicenseKeyWithDefaults instantiates a new LicenseKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseKeyWithDefaults() *LicenseKey {
	this := LicenseKey{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseKey) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *LicenseKey) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *LicenseKey) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *LicenseKey) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *LicenseKey) UnsetKey() {
	o.Key.Unset()
}

func (o LicenseKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableLicenseKey struct {
	value *LicenseKey
	isSet bool
}

func (v NullableLicenseKey) Get() *LicenseKey {
	return v.value
}

func (v *NullableLicenseKey) Set(val *LicenseKey) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseKey) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseKey(val *LicenseKey) *NullableLicenseKey {
	return &NullableLicenseKey{value: val, isSet: true}
}

func (v NullableLicenseKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

