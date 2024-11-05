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

// checks if the DealUnitFlowStageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealUnitFlowStageDto{}

// DealUnitFlowStageDto struct for DealUnitFlowStageDto
type DealUnitFlowStageDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Name NullableString `json:"name,omitempty"`
	DealUnitFlowId NullableString `json:"dealUnitFlowId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Description NullableString `json:"description,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
	ParentBusinessProcessStageId NullableString `json:"parentBusinessProcessStageId,omitempty"`
}

// NewDealUnitFlowStageDto instantiates a new DealUnitFlowStageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealUnitFlowStageDto() *DealUnitFlowStageDto {
	this := DealUnitFlowStageDto{}
	return &this
}

// NewDealUnitFlowStageDtoWithDefaults instantiates a new DealUnitFlowStageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealUnitFlowStageDtoWithDefaults() *DealUnitFlowStageDto {
	this := DealUnitFlowStageDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DealUnitFlowStageDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DealUnitFlowStageDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *DealUnitFlowStageDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DealUnitFlowStageDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DealUnitFlowStageDto) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealUnitFlowStageDto) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *DealUnitFlowStageDto) SetOrder(v int32) {
	o.Order = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DealUnitFlowStageDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DealUnitFlowStageDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetName() {
	o.Name.Unset()
}

// GetDealUnitFlowId returns the DealUnitFlowId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetDealUnitFlowId() string {
	if o == nil || IsNil(o.DealUnitFlowId.Get()) {
		var ret string
		return ret
	}
	return *o.DealUnitFlowId.Get()
}

// GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetDealUnitFlowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealUnitFlowId.Get(), o.DealUnitFlowId.IsSet()
}

// HasDealUnitFlowId returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasDealUnitFlowId() bool {
	if o != nil && o.DealUnitFlowId.IsSet() {
		return true
	}

	return false
}

// SetDealUnitFlowId gets a reference to the given NullableString and assigns it to the DealUnitFlowId field.
func (o *DealUnitFlowStageDto) SetDealUnitFlowId(v string) {
	o.DealUnitFlowId.Set(&v)
}
// SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil
func (o *DealUnitFlowStageDto) SetDealUnitFlowIdNil() {
	o.DealUnitFlowId.Set(nil)
}

// UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetDealUnitFlowId() {
	o.DealUnitFlowId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DealUnitFlowStageDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DealUnitFlowStageDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DealUnitFlowStageDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DealUnitFlowStageDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetDescription() {
	o.Description.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *DealUnitFlowStageDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *DealUnitFlowStageDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

// GetParentBusinessProcessStageId returns the ParentBusinessProcessStageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealUnitFlowStageDto) GetParentBusinessProcessStageId() string {
	if o == nil || IsNil(o.ParentBusinessProcessStageId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentBusinessProcessStageId.Get()
}

// GetParentBusinessProcessStageIdOk returns a tuple with the ParentBusinessProcessStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealUnitFlowStageDto) GetParentBusinessProcessStageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentBusinessProcessStageId.Get(), o.ParentBusinessProcessStageId.IsSet()
}

// HasParentBusinessProcessStageId returns a boolean if a field has been set.
func (o *DealUnitFlowStageDto) HasParentBusinessProcessStageId() bool {
	if o != nil && o.ParentBusinessProcessStageId.IsSet() {
		return true
	}

	return false
}

// SetParentBusinessProcessStageId gets a reference to the given NullableString and assigns it to the ParentBusinessProcessStageId field.
func (o *DealUnitFlowStageDto) SetParentBusinessProcessStageId(v string) {
	o.ParentBusinessProcessStageId.Set(&v)
}
// SetParentBusinessProcessStageIdNil sets the value for ParentBusinessProcessStageId to be an explicit nil
func (o *DealUnitFlowStageDto) SetParentBusinessProcessStageIdNil() {
	o.ParentBusinessProcessStageId.Set(nil)
}

// UnsetParentBusinessProcessStageId ensures that no value is present for ParentBusinessProcessStageId, not even an explicit nil
func (o *DealUnitFlowStageDto) UnsetParentBusinessProcessStageId() {
	o.ParentBusinessProcessStageId.Unset()
}

func (o DealUnitFlowStageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealUnitFlowStageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
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

type NullableDealUnitFlowStageDto struct {
	value *DealUnitFlowStageDto
	isSet bool
}

func (v NullableDealUnitFlowStageDto) Get() *DealUnitFlowStageDto {
	return v.value
}

func (v *NullableDealUnitFlowStageDto) Set(val *DealUnitFlowStageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDealUnitFlowStageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDealUnitFlowStageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealUnitFlowStageDto(val *DealUnitFlowStageDto) *NullableDealUnitFlowStageDto {
	return &NullableDealUnitFlowStageDto{value: val, isSet: true}
}

func (v NullableDealUnitFlowStageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealUnitFlowStageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


