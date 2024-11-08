/*
DealsService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DealUnitFlowUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealUnitFlowUpdateDto{}

// DealUnitFlowUpdateDto struct for DealUnitFlowUpdateDto
type DealUnitFlowUpdateDto struct {
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ParentBusinessProcessId NullableString `json:"parentBusinessProcessId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	TenantEnrollmentId NullableString `json:"tenantEnrollmentId,omitempty"`
}

// NewDealUnitFlowUpdateDto instantiates a new DealUnitFlowUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealUnitFlowUpdateDto() *DealUnitFlowUpdateDto {
	this := DealUnitFlowUpdateDto{}
	return &this
}

// NewDealUnitFlowUpdateDtoWithDefaults instantiates a new DealUnitFlowUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealUnitFlowUpdateDtoWithDefaults() *DealUnitFlowUpdateDto {
	this := DealUnitFlowUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DealUnitFlowUpdateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DealUnitFlowUpdateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DealUnitFlowUpdateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DealUnitFlowUpdateDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DealUnitFlowUpdateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DealUnitFlowUpdateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DealUnitFlowUpdateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DealUnitFlowUpdateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetParentBusinessProcessId returns the ParentBusinessProcessId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowUpdateDto) GetParentBusinessProcessId() string {
	if o == nil || IsNil(o.ParentBusinessProcessId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentBusinessProcessId.Get()
}

// GetParentBusinessProcessIdOk returns a tuple with the ParentBusinessProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowUpdateDto) GetParentBusinessProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentBusinessProcessId.Get(), o.ParentBusinessProcessId.IsSet()
}

// HasParentBusinessProcessId returns a boolean if a field has been set.
func (o *DealUnitFlowUpdateDto) HasParentBusinessProcessId() bool {
	if o != nil && o.ParentBusinessProcessId.IsSet() {
		return true
	}

	return false
}

// SetParentBusinessProcessId gets a reference to the given NullableString and assigns it to the ParentBusinessProcessId field.
func (o *DealUnitFlowUpdateDto) SetParentBusinessProcessId(v string) {
	o.ParentBusinessProcessId.Set(&v)
}
// SetParentBusinessProcessIdNil sets the value for ParentBusinessProcessId to be an explicit nil
func (o *DealUnitFlowUpdateDto) SetParentBusinessProcessIdNil() {
	o.ParentBusinessProcessId.Set(nil)
}

// UnsetParentBusinessProcessId ensures that no value is present for ParentBusinessProcessId, not even an explicit nil
func (o *DealUnitFlowUpdateDto) UnsetParentBusinessProcessId() {
	o.ParentBusinessProcessId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowUpdateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowUpdateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DealUnitFlowUpdateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DealUnitFlowUpdateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DealUnitFlowUpdateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DealUnitFlowUpdateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTenantEnrollmentId returns the TenantEnrollmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowUpdateDto) GetTenantEnrollmentId() string {
	if o == nil || IsNil(o.TenantEnrollmentId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantEnrollmentId.Get()
}

// GetTenantEnrollmentIdOk returns a tuple with the TenantEnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowUpdateDto) GetTenantEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantEnrollmentId.Get(), o.TenantEnrollmentId.IsSet()
}

// HasTenantEnrollmentId returns a boolean if a field has been set.
func (o *DealUnitFlowUpdateDto) HasTenantEnrollmentId() bool {
	if o != nil && o.TenantEnrollmentId.IsSet() {
		return true
	}

	return false
}

// SetTenantEnrollmentId gets a reference to the given NullableString and assigns it to the TenantEnrollmentId field.
func (o *DealUnitFlowUpdateDto) SetTenantEnrollmentId(v string) {
	o.TenantEnrollmentId.Set(&v)
}
// SetTenantEnrollmentIdNil sets the value for TenantEnrollmentId to be an explicit nil
func (o *DealUnitFlowUpdateDto) SetTenantEnrollmentIdNil() {
	o.TenantEnrollmentId.Set(nil)
}

// UnsetTenantEnrollmentId ensures that no value is present for TenantEnrollmentId, not even an explicit nil
func (o *DealUnitFlowUpdateDto) UnsetTenantEnrollmentId() {
	o.TenantEnrollmentId.Unset()
}

func (o DealUnitFlowUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealUnitFlowUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ParentBusinessProcessId.IsSet() {
		toSerialize["parentBusinessProcessId"] = o.ParentBusinessProcessId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.TenantEnrollmentId.IsSet() {
		toSerialize["tenantEnrollmentId"] = o.TenantEnrollmentId.Get()
	}
	return toSerialize, nil
}

type NullableDealUnitFlowUpdateDto struct {
	value *DealUnitFlowUpdateDto
	isSet bool
}

func (v NullableDealUnitFlowUpdateDto) Get() *DealUnitFlowUpdateDto {
	return v.value
}

func (v *NullableDealUnitFlowUpdateDto) Set(val *DealUnitFlowUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDealUnitFlowUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDealUnitFlowUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealUnitFlowUpdateDto(val *DealUnitFlowUpdateDto) *NullableDealUnitFlowUpdateDto {
	return &NullableDealUnitFlowUpdateDto{value: val, isSet: true}
}

func (v NullableDealUnitFlowUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealUnitFlowUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


