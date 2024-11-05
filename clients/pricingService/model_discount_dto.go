/*
PricingService

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

// checks if the DiscountDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountDto{}

// DiscountDto struct for DiscountDto
type DiscountDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	ItemId NullableString `json:"itemId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
	DiscountListId NullableString `json:"discountListId,omitempty"`
	EndQuantity *float64 `json:"endQuantity,omitempty"`
	BeginQuantity *float64 `json:"beginQuantity,omitempty"`
}

// NewDiscountDto instantiates a new DiscountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountDto() *DiscountDto {
	this := DiscountDto{}
	return &this
}

// NewDiscountDtoWithDefaults instantiates a new DiscountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountDtoWithDefaults() *DiscountDto {
	this := DiscountDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DiscountDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DiscountDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DiscountDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DiscountDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiscountDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *DiscountDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DiscountDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DiscountDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscountDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DiscountDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DiscountDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DiscountDto) UnsetDescription() {
	o.Description.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DiscountDto) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountDto) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DiscountDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DiscountDto) SetValue(v float64) {
	o.Value = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *DiscountDto) GetPercent() float64 {
	if o == nil || IsNil(o.Percent) {
		var ret float64
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountDto) GetPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *DiscountDto) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float64 and assigns it to the Percent field.
func (o *DiscountDto) SetPercent(v float64) {
	o.Percent = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *DiscountDto) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *DiscountDto) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *DiscountDto) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *DiscountDto) UnsetItemId() {
	o.ItemId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DiscountDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DiscountDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DiscountDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DiscountDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *DiscountDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *DiscountDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *DiscountDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *DiscountDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

// GetDiscountListId returns the DiscountListId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountDto) GetDiscountListId() string {
	if o == nil || IsNil(o.DiscountListId.Get()) {
		var ret string
		return ret
	}
	return *o.DiscountListId.Get()
}

// GetDiscountListIdOk returns a tuple with the DiscountListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountDto) GetDiscountListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountListId.Get(), o.DiscountListId.IsSet()
}

// HasDiscountListId returns a boolean if a field has been set.
func (o *DiscountDto) HasDiscountListId() bool {
	if o != nil && o.DiscountListId.IsSet() {
		return true
	}

	return false
}

// SetDiscountListId gets a reference to the given NullableString and assigns it to the DiscountListId field.
func (o *DiscountDto) SetDiscountListId(v string) {
	o.DiscountListId.Set(&v)
}
// SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil
func (o *DiscountDto) SetDiscountListIdNil() {
	o.DiscountListId.Set(nil)
}

// UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
func (o *DiscountDto) UnsetDiscountListId() {
	o.DiscountListId.Unset()
}

// GetEndQuantity returns the EndQuantity field value if set, zero value otherwise.
func (o *DiscountDto) GetEndQuantity() float64 {
	if o == nil || IsNil(o.EndQuantity) {
		var ret float64
		return ret
	}
	return *o.EndQuantity
}

// GetEndQuantityOk returns a tuple with the EndQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountDto) GetEndQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.EndQuantity) {
		return nil, false
	}
	return o.EndQuantity, true
}

// HasEndQuantity returns a boolean if a field has been set.
func (o *DiscountDto) HasEndQuantity() bool {
	if o != nil && !IsNil(o.EndQuantity) {
		return true
	}

	return false
}

// SetEndQuantity gets a reference to the given float64 and assigns it to the EndQuantity field.
func (o *DiscountDto) SetEndQuantity(v float64) {
	o.EndQuantity = &v
}

// GetBeginQuantity returns the BeginQuantity field value if set, zero value otherwise.
func (o *DiscountDto) GetBeginQuantity() float64 {
	if o == nil || IsNil(o.BeginQuantity) {
		var ret float64
		return ret
	}
	return *o.BeginQuantity
}

// GetBeginQuantityOk returns a tuple with the BeginQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountDto) GetBeginQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.BeginQuantity) {
		return nil, false
	}
	return o.BeginQuantity, true
}

// HasBeginQuantity returns a boolean if a field has been set.
func (o *DiscountDto) HasBeginQuantity() bool {
	if o != nil && !IsNil(o.BeginQuantity) {
		return true
	}

	return false
}

// SetBeginQuantity gets a reference to the given float64 and assigns it to the BeginQuantity field.
func (o *DiscountDto) SetBeginQuantity(v float64) {
	o.BeginQuantity = &v
}

func (o DiscountDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if o.ItemId.IsSet() {
		toSerialize["itemId"] = o.ItemId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrolmentId.IsSet() {
		toSerialize["enrolmentId"] = o.EnrolmentId.Get()
	}
	if o.DiscountListId.IsSet() {
		toSerialize["discountListId"] = o.DiscountListId.Get()
	}
	if !IsNil(o.EndQuantity) {
		toSerialize["endQuantity"] = o.EndQuantity
	}
	if !IsNil(o.BeginQuantity) {
		toSerialize["beginQuantity"] = o.BeginQuantity
	}
	return toSerialize, nil
}

type NullableDiscountDto struct {
	value *DiscountDto
	isSet bool
}

func (v NullableDiscountDto) Get() *DiscountDto {
	return v.value
}

func (v *NullableDiscountDto) Set(val *DiscountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountDto(val *DiscountDto) *NullableDiscountDto {
	return &NullableDiscountDto{value: val, isSet: true}
}

func (v NullableDiscountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


