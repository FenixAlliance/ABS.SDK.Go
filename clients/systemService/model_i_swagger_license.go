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

// checks if the ISwaggerLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISwaggerLicense{}

// ISwaggerLicense struct for ISwaggerLicense
type ISwaggerLicense struct {
	Name NullableString `json:"name,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewISwaggerLicense instantiates a new ISwaggerLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISwaggerLicense() *ISwaggerLicense {
	this := ISwaggerLicense{}
	return &this
}

// NewISwaggerLicenseWithDefaults instantiates a new ISwaggerLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISwaggerLicenseWithDefaults() *ISwaggerLicense {
	this := ISwaggerLicense{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerLicense) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerLicense) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ISwaggerLicense) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ISwaggerLicense) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ISwaggerLicense) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ISwaggerLicense) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerLicense) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerLicense) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ISwaggerLicense) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ISwaggerLicense) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ISwaggerLicense) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ISwaggerLicense) UnsetUrl() {
	o.Url.Unset()
}

func (o ISwaggerLicense) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISwaggerLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableISwaggerLicense struct {
	value *ISwaggerLicense
	isSet bool
}

func (v NullableISwaggerLicense) Get() *ISwaggerLicense {
	return v.value
}

func (v *NullableISwaggerLicense) Set(val *ISwaggerLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableISwaggerLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableISwaggerLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISwaggerLicense(val *ISwaggerLicense) *NullableISwaggerLicense {
	return &NullableISwaggerLicense{value: val, isSet: true}
}

func (v NullableISwaggerLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISwaggerLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


