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

// checks if the ItemPriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPriceDto{}

// ItemPriceDto struct for ItemPriceDto
type ItemPriceDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	ItemId NullableString `json:"itemId,omitempty"`
	UnitId NullableString `json:"unitId,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	DiscountId NullableString `json:"discountId,omitempty"`
	UnitGroupId NullableString `json:"unitGroupId,omitempty"`
	PriceListId NullableString `json:"priceListId,omitempty"`
	DiscountListId NullableString `json:"discountListId,omitempty"`
	RoundingPolicyId NullableString `json:"roundingPolicyId,omitempty"`
	EnrollmentId NullableString `json:"enrollmentId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
}

// NewItemPriceDto instantiates a new ItemPriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPriceDto() *ItemPriceDto {
	this := ItemPriceDto{}
	return &this
}

// NewItemPriceDtoWithDefaults instantiates a new ItemPriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPriceDtoWithDefaults() *ItemPriceDto {
	this := ItemPriceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ItemPriceDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ItemPriceDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ItemPriceDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ItemPriceDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *ItemPriceDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ItemPriceDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ItemPriceDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *ItemPriceDto) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *ItemPriceDto) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *ItemPriceDto) UnsetItemId() {
	o.ItemId.Unset()
}

// GetUnitId returns the UnitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetUnitId() string {
	if o == nil || IsNil(o.UnitId.Get()) {
		var ret string
		return ret
	}
	return *o.UnitId.Get()
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetUnitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitId.Get(), o.UnitId.IsSet()
}

// HasUnitId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasUnitId() bool {
	if o != nil && o.UnitId.IsSet() {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given NullableString and assigns it to the UnitId field.
func (o *ItemPriceDto) SetUnitId(v string) {
	o.UnitId.Set(&v)
}
// SetUnitIdNil sets the value for UnitId to be an explicit nil
func (o *ItemPriceDto) SetUnitIdNil() {
	o.UnitId.Set(nil)
}

// UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
func (o *ItemPriceDto) UnsetUnitId() {
	o.UnitId.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *ItemPriceDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *ItemPriceDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *ItemPriceDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetDiscountId returns the DiscountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetDiscountId() string {
	if o == nil || IsNil(o.DiscountId.Get()) {
		var ret string
		return ret
	}
	return *o.DiscountId.Get()
}

// GetDiscountIdOk returns a tuple with the DiscountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetDiscountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountId.Get(), o.DiscountId.IsSet()
}

// HasDiscountId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasDiscountId() bool {
	if o != nil && o.DiscountId.IsSet() {
		return true
	}

	return false
}

// SetDiscountId gets a reference to the given NullableString and assigns it to the DiscountId field.
func (o *ItemPriceDto) SetDiscountId(v string) {
	o.DiscountId.Set(&v)
}
// SetDiscountIdNil sets the value for DiscountId to be an explicit nil
func (o *ItemPriceDto) SetDiscountIdNil() {
	o.DiscountId.Set(nil)
}

// UnsetDiscountId ensures that no value is present for DiscountId, not even an explicit nil
func (o *ItemPriceDto) UnsetDiscountId() {
	o.DiscountId.Unset()
}

// GetUnitGroupId returns the UnitGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetUnitGroupId() string {
	if o == nil || IsNil(o.UnitGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.UnitGroupId.Get()
}

// GetUnitGroupIdOk returns a tuple with the UnitGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetUnitGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitGroupId.Get(), o.UnitGroupId.IsSet()
}

// HasUnitGroupId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasUnitGroupId() bool {
	if o != nil && o.UnitGroupId.IsSet() {
		return true
	}

	return false
}

// SetUnitGroupId gets a reference to the given NullableString and assigns it to the UnitGroupId field.
func (o *ItemPriceDto) SetUnitGroupId(v string) {
	o.UnitGroupId.Set(&v)
}
// SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil
func (o *ItemPriceDto) SetUnitGroupIdNil() {
	o.UnitGroupId.Set(nil)
}

// UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
func (o *ItemPriceDto) UnsetUnitGroupId() {
	o.UnitGroupId.Unset()
}

// GetPriceListId returns the PriceListId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetPriceListId() string {
	if o == nil || IsNil(o.PriceListId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceListId.Get()
}

// GetPriceListIdOk returns a tuple with the PriceListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetPriceListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceListId.Get(), o.PriceListId.IsSet()
}

// HasPriceListId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasPriceListId() bool {
	if o != nil && o.PriceListId.IsSet() {
		return true
	}

	return false
}

// SetPriceListId gets a reference to the given NullableString and assigns it to the PriceListId field.
func (o *ItemPriceDto) SetPriceListId(v string) {
	o.PriceListId.Set(&v)
}
// SetPriceListIdNil sets the value for PriceListId to be an explicit nil
func (o *ItemPriceDto) SetPriceListIdNil() {
	o.PriceListId.Set(nil)
}

// UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
func (o *ItemPriceDto) UnsetPriceListId() {
	o.PriceListId.Unset()
}

// GetDiscountListId returns the DiscountListId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetDiscountListId() string {
	if o == nil || IsNil(o.DiscountListId.Get()) {
		var ret string
		return ret
	}
	return *o.DiscountListId.Get()
}

// GetDiscountListIdOk returns a tuple with the DiscountListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetDiscountListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountListId.Get(), o.DiscountListId.IsSet()
}

// HasDiscountListId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasDiscountListId() bool {
	if o != nil && o.DiscountListId.IsSet() {
		return true
	}

	return false
}

// SetDiscountListId gets a reference to the given NullableString and assigns it to the DiscountListId field.
func (o *ItemPriceDto) SetDiscountListId(v string) {
	o.DiscountListId.Set(&v)
}
// SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil
func (o *ItemPriceDto) SetDiscountListIdNil() {
	o.DiscountListId.Set(nil)
}

// UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
func (o *ItemPriceDto) UnsetDiscountListId() {
	o.DiscountListId.Unset()
}

// GetRoundingPolicyId returns the RoundingPolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetRoundingPolicyId() string {
	if o == nil || IsNil(o.RoundingPolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.RoundingPolicyId.Get()
}

// GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetRoundingPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoundingPolicyId.Get(), o.RoundingPolicyId.IsSet()
}

// HasRoundingPolicyId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasRoundingPolicyId() bool {
	if o != nil && o.RoundingPolicyId.IsSet() {
		return true
	}

	return false
}

// SetRoundingPolicyId gets a reference to the given NullableString and assigns it to the RoundingPolicyId field.
func (o *ItemPriceDto) SetRoundingPolicyId(v string) {
	o.RoundingPolicyId.Set(&v)
}
// SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil
func (o *ItemPriceDto) SetRoundingPolicyIdNil() {
	o.RoundingPolicyId.Set(nil)
}

// UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
func (o *ItemPriceDto) UnsetRoundingPolicyId() {
	o.RoundingPolicyId.Unset()
}

// GetEnrollmentId returns the EnrollmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetEnrollmentId() string {
	if o == nil || IsNil(o.EnrollmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentId.Get()
}

// GetEnrollmentIdOk returns a tuple with the EnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentId.Get(), o.EnrollmentId.IsSet()
}

// HasEnrollmentId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasEnrollmentId() bool {
	if o != nil && o.EnrollmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentId gets a reference to the given NullableString and assigns it to the EnrollmentId field.
func (o *ItemPriceDto) SetEnrollmentId(v string) {
	o.EnrollmentId.Set(&v)
}
// SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil
func (o *ItemPriceDto) SetEnrollmentIdNil() {
	o.EnrollmentId.Set(nil)
}

// UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
func (o *ItemPriceDto) UnsetEnrollmentId() {
	o.EnrollmentId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemPriceDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemPriceDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ItemPriceDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ItemPriceDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ItemPriceDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ItemPriceDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ItemPriceDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPriceDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ItemPriceDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *ItemPriceDto) SetPrice(v float64) {
	o.Price = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *ItemPriceDto) GetPercent() float64 {
	if o == nil || IsNil(o.Percent) {
		var ret float64
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPriceDto) GetPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *ItemPriceDto) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float64 and assigns it to the Percent field.
func (o *ItemPriceDto) SetPercent(v float64) {
	o.Percent = &v
}

func (o ItemPriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemPriceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["itemId"] = o.ItemId.Get()
	}
	if o.UnitId.IsSet() {
		toSerialize["unitId"] = o.UnitId.Get()
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.DiscountId.IsSet() {
		toSerialize["discountId"] = o.DiscountId.Get()
	}
	if o.UnitGroupId.IsSet() {
		toSerialize["unitGroupId"] = o.UnitGroupId.Get()
	}
	if o.PriceListId.IsSet() {
		toSerialize["priceListId"] = o.PriceListId.Get()
	}
	if o.DiscountListId.IsSet() {
		toSerialize["discountListId"] = o.DiscountListId.Get()
	}
	if o.RoundingPolicyId.IsSet() {
		toSerialize["roundingPolicyId"] = o.RoundingPolicyId.Get()
	}
	if o.EnrollmentId.IsSet() {
		toSerialize["enrollmentId"] = o.EnrollmentId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	return toSerialize, nil
}

type NullableItemPriceDto struct {
	value *ItemPriceDto
	isSet bool
}

func (v NullableItemPriceDto) Get() *ItemPriceDto {
	return v.value
}

func (v *NullableItemPriceDto) Set(val *ItemPriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPriceDto(val *ItemPriceDto) *NullableItemPriceDto {
	return &NullableItemPriceDto{value: val, isSet: true}
}

func (v NullableItemPriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


