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
	"time"
	"bytes"
	"fmt"
)

// checks if the SecurityRoleCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityRoleCreateDto{}

// SecurityRoleCreateDto struct for SecurityRoleCreateDto
type SecurityRoleCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Name string `json:"name"`
	TenantId string `json:"tenantId"`
	Description NullableString `json:"description,omitempty"`
}

type _SecurityRoleCreateDto SecurityRoleCreateDto

// NewSecurityRoleCreateDto instantiates a new SecurityRoleCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityRoleCreateDto(name string, tenantId string) *SecurityRoleCreateDto {
	this := SecurityRoleCreateDto{}
	this.Name = name
	this.TenantId = tenantId
	return &this
}

// NewSecurityRoleCreateDtoWithDefaults instantiates a new SecurityRoleCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityRoleCreateDtoWithDefaults() *SecurityRoleCreateDto {
	this := SecurityRoleCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityRoleCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityRoleCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityRoleCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SecurityRoleCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SecurityRoleCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SecurityRoleCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetName returns the Name field value
func (o *SecurityRoleCreateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityRoleCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityRoleCreateDto) SetName(v string) {
	o.Name = v
}

// GetTenantId returns the TenantId field value
func (o *SecurityRoleCreateDto) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *SecurityRoleCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *SecurityRoleCreateDto) SetTenantId(v string) {
	o.TenantId = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityRoleCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityRoleCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityRoleCreateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecurityRoleCreateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecurityRoleCreateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecurityRoleCreateDto) UnsetDescription() {
	o.Description.Unset()
}

func (o SecurityRoleCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityRoleCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["name"] = o.Name
	toSerialize["tenantId"] = o.TenantId
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *SecurityRoleCreateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"tenantId",
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

	varSecurityRoleCreateDto := _SecurityRoleCreateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecurityRoleCreateDto)

	if err != nil {
		return err
	}

	*o = SecurityRoleCreateDto(varSecurityRoleCreateDto)

	return err
}

type NullableSecurityRoleCreateDto struct {
	value *SecurityRoleCreateDto
	isSet bool
}

func (v NullableSecurityRoleCreateDto) Get() *SecurityRoleCreateDto {
	return v.value
}

func (v *NullableSecurityRoleCreateDto) Set(val *SecurityRoleCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityRoleCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityRoleCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityRoleCreateDto(val *SecurityRoleCreateDto) *NullableSecurityRoleCreateDto {
	return &NullableSecurityRoleCreateDto{value: val, isSet: true}
}

func (v NullableSecurityRoleCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityRoleCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


