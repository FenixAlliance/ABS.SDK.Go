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
	"time"
)

// checks if the DealUnitFlowStageCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealUnitFlowStageCreateDto{}

// DealUnitFlowStageCreateDto struct for DealUnitFlowStageCreateDto
type DealUnitFlowStageCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Name NullableString `json:"name,omitempty"`
	DealUnitFlowId NullableString `json:"dealUnitFlowId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Description NullableString `json:"description,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
	ParentBusinessProcessStageId NullableString `json:"parentBusinessProcessStageId,omitempty"`
}

// NewDealUnitFlowStageCreateDto instantiates a new DealUnitFlowStageCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealUnitFlowStageCreateDto() *DealUnitFlowStageCreateDto {
	this := DealUnitFlowStageCreateDto{}
	return &this
}

// NewDealUnitFlowStageCreateDtoWithDefaults instantiates a new DealUnitFlowStageCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealUnitFlowStageCreateDtoWithDefaults() *DealUnitFlowStageCreateDto {
	this := DealUnitFlowStageCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DealUnitFlowStageCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUnitFlowStageCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DealUnitFlowStageCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DealUnitFlowStageCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUnitFlowStageCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DealUnitFlowStageCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DealUnitFlowStageCreateDto) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUnitFlowStageCreateDto) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *DealUnitFlowStageCreateDto) SetOrder(v int32) {
	o.Order = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DealUnitFlowStageCreateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetName() {
	o.Name.Unset()
}

// GetDealUnitFlowId returns the DealUnitFlowId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetDealUnitFlowId() string {
	if o == nil || IsNil(o.DealUnitFlowId.Get()) {
		var ret string
		return ret
	}
	return *o.DealUnitFlowId.Get()
}

// GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetDealUnitFlowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealUnitFlowId.Get(), o.DealUnitFlowId.IsSet()
}

// HasDealUnitFlowId returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasDealUnitFlowId() bool {
	if o != nil && o.DealUnitFlowId.IsSet() {
		return true
	}

	return false
}

// SetDealUnitFlowId gets a reference to the given NullableString and assigns it to the DealUnitFlowId field.
func (o *DealUnitFlowStageCreateDto) SetDealUnitFlowId(v string) {
	o.DealUnitFlowId.Set(&v)
}
// SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetDealUnitFlowIdNil() {
	o.DealUnitFlowId.Set(nil)
}

// UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetDealUnitFlowId() {
	o.DealUnitFlowId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DealUnitFlowStageCreateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DealUnitFlowStageCreateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *DealUnitFlowStageCreateDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

// GetParentBusinessProcessStageId returns the ParentBusinessProcessStageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageCreateDto) GetParentBusinessProcessStageId() string {
	if o == nil || IsNil(o.ParentBusinessProcessStageId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentBusinessProcessStageId.Get()
}

// GetParentBusinessProcessStageIdOk returns a tuple with the ParentBusinessProcessStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageCreateDto) GetParentBusinessProcessStageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentBusinessProcessStageId.Get(), o.ParentBusinessProcessStageId.IsSet()
}

// HasParentBusinessProcessStageId returns a boolean if a field has been set.
func (o *DealUnitFlowStageCreateDto) HasParentBusinessProcessStageId() bool {
	if o != nil && o.ParentBusinessProcessStageId.IsSet() {
		return true
	}

	return false
}

// SetParentBusinessProcessStageId gets a reference to the given NullableString and assigns it to the ParentBusinessProcessStageId field.
func (o *DealUnitFlowStageCreateDto) SetParentBusinessProcessStageId(v string) {
	o.ParentBusinessProcessStageId.Set(&v)
}
// SetParentBusinessProcessStageIdNil sets the value for ParentBusinessProcessStageId to be an explicit nil
func (o *DealUnitFlowStageCreateDto) SetParentBusinessProcessStageIdNil() {
	o.ParentBusinessProcessStageId.Set(nil)
}

// UnsetParentBusinessProcessStageId ensures that no value is present for ParentBusinessProcessStageId, not even an explicit nil
func (o *DealUnitFlowStageCreateDto) UnsetParentBusinessProcessStageId() {
	o.ParentBusinessProcessStageId.Unset()
}

func (o DealUnitFlowStageCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealUnitFlowStageCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DealUnitFlowId.IsSet() {
		toSerialize["dealUnitFlowId"] = o.DealUnitFlowId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EnrolmentId.IsSet() {
		toSerialize["enrolmentId"] = o.EnrolmentId.Get()
	}
	if o.ParentBusinessProcessStageId.IsSet() {
		toSerialize["parentBusinessProcessStageId"] = o.ParentBusinessProcessStageId.Get()
	}
	return toSerialize, nil
}

type NullableDealUnitFlowStageCreateDto struct {
	value *DealUnitFlowStageCreateDto
	isSet bool
}

func (v NullableDealUnitFlowStageCreateDto) Get() *DealUnitFlowStageCreateDto {
	return v.value
}

func (v *NullableDealUnitFlowStageCreateDto) Set(val *DealUnitFlowStageCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDealUnitFlowStageCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDealUnitFlowStageCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealUnitFlowStageCreateDto(val *DealUnitFlowStageCreateDto) *NullableDealUnitFlowStageCreateDto {
	return &NullableDealUnitFlowStageCreateDto{value: val, isSet: true}
}

func (v NullableDealUnitFlowStageCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealUnitFlowStageCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


