/*
SecurityService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SecurityPermissionUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPermissionUpdateDto{}

// SecurityPermissionUpdateDto struct for SecurityPermissionUpdateDto
type SecurityPermissionUpdateDto struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
}

type _SecurityPermissionUpdateDto SecurityPermissionUpdateDto

// NewSecurityPermissionUpdateDto instantiates a new SecurityPermissionUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPermissionUpdateDto(name string) *SecurityPermissionUpdateDto {
	this := SecurityPermissionUpdateDto{}
	this.Name = name
	return &this
}

// NewSecurityPermissionUpdateDtoWithDefaults instantiates a new SecurityPermissionUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPermissionUpdateDtoWithDefaults() *SecurityPermissionUpdateDto {
	this := SecurityPermissionUpdateDto{}
	return &this
}

// GetName returns the Name field value
func (o *SecurityPermissionUpdateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityPermissionUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityPermissionUpdateDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityPermissionUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityPermissionUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityPermissionUpdateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecurityPermissionUpdateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecurityPermissionUpdateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecurityPermissionUpdateDto) UnsetDescription() {
	o.Description.Unset()
}

func (o SecurityPermissionUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityPermissionUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *SecurityPermissionUpdateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSecurityPermissionUpdateDto := _SecurityPermissionUpdateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecurityPermissionUpdateDto)

	if err != nil {
		return err
	}

	*o = SecurityPermissionUpdateDto(varSecurityPermissionUpdateDto)

	return err
}

type NullableSecurityPermissionUpdateDto struct {
	value *SecurityPermissionUpdateDto
	isSet bool
}

func (v NullableSecurityPermissionUpdateDto) Get() *SecurityPermissionUpdateDto {
	return v.value
}

func (v *NullableSecurityPermissionUpdateDto) Set(val *SecurityPermissionUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPermissionUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPermissionUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPermissionUpdateDto(val *SecurityPermissionUpdateDto) *NullableSecurityPermissionUpdateDto {
	return &NullableSecurityPermissionUpdateDto{value: val, isSet: true}
}

func (v NullableSecurityPermissionUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPermissionUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


