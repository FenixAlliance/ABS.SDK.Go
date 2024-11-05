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

// checks if the DiscountListCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountListCreateDto{}

// DiscountListCreateDto struct for DiscountListCreateDto
type DiscountListCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
}

// NewDiscountListCreateDto instantiates a new DiscountListCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountListCreateDto() *DiscountListCreateDto {
	this := DiscountListCreateDto{}
	return &this
}

// NewDiscountListCreateDtoWithDefaults instantiates a new DiscountListCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountListCreateDtoWithDefaults() *DiscountListCreateDto {
	this := DiscountListCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscountListCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountListCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiscountListCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DiscountListCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountListCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DiscountListCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountListCreateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountListCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DiscountListCreateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DiscountListCreateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DiscountListCreateDto) UnsetName() {
	o.Name.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountListCreateDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountListCreateDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *DiscountListCreateDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *DiscountListCreateDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *DiscountListCreateDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountListCreateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountListCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DiscountListCreateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DiscountListCreateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DiscountListCreateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountListCreateDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountListCreateDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *DiscountListCreateDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *DiscountListCreateDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *DiscountListCreateDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *DiscountListCreateDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

func (o DiscountListCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountListCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrolmentId.IsSet() {
		toSerialize["enrolmentId"] = o.EnrolmentId.Get()
	}
	return toSerialize, nil
}

type NullableDiscountListCreateDto struct {
	value *DiscountListCreateDto
	isSet bool
}

func (v NullableDiscountListCreateDto) Get() *DiscountListCreateDto {
	return v.value
}

func (v *NullableDiscountListCreateDto) Set(val *DiscountListCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountListCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountListCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountListCreateDto(val *DiscountListCreateDto) *NullableDiscountListCreateDto {
	return &NullableDiscountListCreateDto{value: val, isSet: true}
}

func (v NullableDiscountListCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountListCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

