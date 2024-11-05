/*
InvoicingService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InvoiceLineUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceLineUpdateDto{}

// InvoiceLineUpdateDto struct for InvoiceLineUpdateDto
type InvoiceLineUpdateDto struct {
	Price *float64 `json:"price,omitempty"`
	UnitId NullableString `json:"unitId,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	UnitGroupId NullableString `json:"unitGroupId,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	DiscountListId NullableString `json:"discountListId,omitempty"`
	RoundingPolicyId NullableString `json:"roundingPolicyId,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	ItemId NullableString `json:"itemId,omitempty"`
	ItemPriceId NullableString `json:"itemPriceId,omitempty"`
	InvoiceLineId NullableString `json:"invoiceLineId,omitempty"`
	TaxAmountInUsd *float64 `json:"taxAmountInUsd,omitempty"`
	TaxBaseAmountInUsd *float64 `json:"taxBaseAmountInUsd,omitempty"`
}

// NewInvoiceLineUpdateDto instantiates a new InvoiceLineUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceLineUpdateDto() *InvoiceLineUpdateDto {
	this := InvoiceLineUpdateDto{}
	return &this
}

// NewInvoiceLineUpdateDtoWithDefaults instantiates a new InvoiceLineUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceLineUpdateDtoWithDefaults() *InvoiceLineUpdateDto {
	this := InvoiceLineUpdateDto{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InvoiceLineUpdateDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineUpdateDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *InvoiceLineUpdateDto) SetPrice(v float64) {
	o.Price = &v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetUnitId() string {
	if o == nil || IsNil(o.UnitId.Get()) {
		var ret string
		return ret
	}
	return *o.UnitId.Get()
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetUnitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitId.Get(), o.UnitId.IsSet()
}

// HasUnitId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasUnitId() bool {
	if o != nil && o.UnitId.IsSet() {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given NullableString and assigns it to the UnitId field.
func (o *InvoiceLineUpdateDto) SetUnitId(v string) {
	o.UnitId.Set(&v)
}
// SetUnitIdNil sets the value for UnitId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetUnitIdNil() {
	o.UnitId.Set(nil)
}

// UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetUnitId() {
	o.UnitId.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *InvoiceLineUpdateDto) GetPercent() float64 {
	if o == nil || IsNil(o.Percent) {
		var ret float64
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineUpdateDto) GetPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float64 and assigns it to the Percent field.
func (o *InvoiceLineUpdateDto) SetPercent(v float64) {
	o.Percent = &v
}

// GetUnitGroupId returns the UnitGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetUnitGroupId() string {
	if o == nil || IsNil(o.UnitGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.UnitGroupId.Get()
}

// GetUnitGroupIdOk returns a tuple with the UnitGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetUnitGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitGroupId.Get(), o.UnitGroupId.IsSet()
}

// HasUnitGroupId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasUnitGroupId() bool {
	if o != nil && o.UnitGroupId.IsSet() {
		return true
	}

	return false
}

// SetUnitGroupId gets a reference to the given NullableString and assigns it to the UnitGroupId field.
func (o *InvoiceLineUpdateDto) SetUnitGroupId(v string) {
	o.UnitGroupId.Set(&v)
}
// SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetUnitGroupIdNil() {
	o.UnitGroupId.Set(nil)
}

// UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetUnitGroupId() {
	o.UnitGroupId.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *InvoiceLineUpdateDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetDiscountListId returns the DiscountListId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetDiscountListId() string {
	if o == nil || IsNil(o.DiscountListId.Get()) {
		var ret string
		return ret
	}
	return *o.DiscountListId.Get()
}

// GetDiscountListIdOk returns a tuple with the DiscountListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetDiscountListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountListId.Get(), o.DiscountListId.IsSet()
}

// HasDiscountListId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasDiscountListId() bool {
	if o != nil && o.DiscountListId.IsSet() {
		return true
	}

	return false
}

// SetDiscountListId gets a reference to the given NullableString and assigns it to the DiscountListId field.
func (o *InvoiceLineUpdateDto) SetDiscountListId(v string) {
	o.DiscountListId.Set(&v)
}
// SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetDiscountListIdNil() {
	o.DiscountListId.Set(nil)
}

// UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetDiscountListId() {
	o.DiscountListId.Unset()
}

// GetRoundingPolicyId returns the RoundingPolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetRoundingPolicyId() string {
	if o == nil || IsNil(o.RoundingPolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.RoundingPolicyId.Get()
}

// GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetRoundingPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoundingPolicyId.Get(), o.RoundingPolicyId.IsSet()
}

// HasRoundingPolicyId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasRoundingPolicyId() bool {
	if o != nil && o.RoundingPolicyId.IsSet() {
		return true
	}

	return false
}

// SetRoundingPolicyId gets a reference to the given NullableString and assigns it to the RoundingPolicyId field.
func (o *InvoiceLineUpdateDto) SetRoundingPolicyId(v string) {
	o.RoundingPolicyId.Set(&v)
}
// SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetRoundingPolicyIdNil() {
	o.RoundingPolicyId.Set(nil)
}

// UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetRoundingPolicyId() {
	o.RoundingPolicyId.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *InvoiceLineUpdateDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineUpdateDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *InvoiceLineUpdateDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *InvoiceLineUpdateDto) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetItemId() {
	o.ItemId.Unset()
}

// GetItemPriceId returns the ItemPriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetItemPriceId() string {
	if o == nil || IsNil(o.ItemPriceId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemPriceId.Get()
}

// GetItemPriceIdOk returns a tuple with the ItemPriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetItemPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemPriceId.Get(), o.ItemPriceId.IsSet()
}

// HasItemPriceId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasItemPriceId() bool {
	if o != nil && o.ItemPriceId.IsSet() {
		return true
	}

	return false
}

// SetItemPriceId gets a reference to the given NullableString and assigns it to the ItemPriceId field.
func (o *InvoiceLineUpdateDto) SetItemPriceId(v string) {
	o.ItemPriceId.Set(&v)
}
// SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetItemPriceIdNil() {
	o.ItemPriceId.Set(nil)
}

// UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetItemPriceId() {
	o.ItemPriceId.Unset()
}

// GetInvoiceLineId returns the InvoiceLineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUpdateDto) GetInvoiceLineId() string {
	if o == nil || IsNil(o.InvoiceLineId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceLineId.Get()
}

// GetInvoiceLineIdOk returns a tuple with the InvoiceLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUpdateDto) GetInvoiceLineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceLineId.Get(), o.InvoiceLineId.IsSet()
}

// HasInvoiceLineId returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasInvoiceLineId() bool {
	if o != nil && o.InvoiceLineId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceLineId gets a reference to the given NullableString and assigns it to the InvoiceLineId field.
func (o *InvoiceLineUpdateDto) SetInvoiceLineId(v string) {
	o.InvoiceLineId.Set(&v)
}
// SetInvoiceLineIdNil sets the value for InvoiceLineId to be an explicit nil
func (o *InvoiceLineUpdateDto) SetInvoiceLineIdNil() {
	o.InvoiceLineId.Set(nil)
}

// UnsetInvoiceLineId ensures that no value is present for InvoiceLineId, not even an explicit nil
func (o *InvoiceLineUpdateDto) UnsetInvoiceLineId() {
	o.InvoiceLineId.Unset()
}

// GetTaxAmountInUsd returns the TaxAmountInUsd field value if set, zero value otherwise.
func (o *InvoiceLineUpdateDto) GetTaxAmountInUsd() float64 {
	if o == nil || IsNil(o.TaxAmountInUsd) {
		var ret float64
		return ret
	}
	return *o.TaxAmountInUsd
}

// GetTaxAmountInUsdOk returns a tuple with the TaxAmountInUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineUpdateDto) GetTaxAmountInUsdOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxAmountInUsd) {
		return nil, false
	}
	return o.TaxAmountInUsd, true
}

// HasTaxAmountInUsd returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasTaxAmountInUsd() bool {
	if o != nil && !IsNil(o.TaxAmountInUsd) {
		return true
	}

	return false
}

// SetTaxAmountInUsd gets a reference to the given float64 and assigns it to the TaxAmountInUsd field.
func (o *InvoiceLineUpdateDto) SetTaxAmountInUsd(v float64) {
	o.TaxAmountInUsd = &v
}

// GetTaxBaseAmountInUsd returns the TaxBaseAmountInUsd field value if set, zero value otherwise.
func (o *InvoiceLineUpdateDto) GetTaxBaseAmountInUsd() float64 {
	if o == nil || IsNil(o.TaxBaseAmountInUsd) {
		var ret float64
		return ret
	}
	return *o.TaxBaseAmountInUsd
}

// GetTaxBaseAmountInUsdOk returns a tuple with the TaxBaseAmountInUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineUpdateDto) GetTaxBaseAmountInUsdOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxBaseAmountInUsd) {
		return nil, false
	}
	return o.TaxBaseAmountInUsd, true
}

// HasTaxBaseAmountInUsd returns a boolean if a field has been set.
func (o *InvoiceLineUpdateDto) HasTaxBaseAmountInUsd() bool {
	if o != nil && !IsNil(o.TaxBaseAmountInUsd) {
		return true
	}

	return false
}

// SetTaxBaseAmountInUsd gets a reference to the given float64 and assigns it to the TaxBaseAmountInUsd field.
func (o *InvoiceLineUpdateDto) SetTaxBaseAmountInUsd(v float64) {
	o.TaxBaseAmountInUsd = &v
}

func (o InvoiceLineUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceLineUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.UnitId.IsSet() {
		toSerialize["unitId"] = o.UnitId.Get()
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if o.UnitGroupId.IsSet() {
		toSerialize["unitGroupId"] = o.UnitGroupId.Get()
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.DiscountListId.IsSet() {
		toSerialize["discountListId"] = o.DiscountListId.Get()
	}
	if o.RoundingPolicyId.IsSet() {
		toSerialize["roundingPolicyId"] = o.RoundingPolicyId.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ItemId.IsSet() {
		toSerialize["itemId"] = o.ItemId.Get()
	}
	if o.ItemPriceId.IsSet() {
		toSerialize["itemPriceId"] = o.ItemPriceId.Get()
	}
	if o.InvoiceLineId.IsSet() {
		toSerialize["invoiceLineId"] = o.InvoiceLineId.Get()
	}
	if !IsNil(o.TaxAmountInUsd) {
		toSerialize["taxAmountInUsd"] = o.TaxAmountInUsd
	}
	if !IsNil(o.TaxBaseAmountInUsd) {
		toSerialize["taxBaseAmountInUsd"] = o.TaxBaseAmountInUsd
	}
	return toSerialize, nil
}

type NullableInvoiceLineUpdateDto struct {
	value *InvoiceLineUpdateDto
	isSet bool
}

func (v NullableInvoiceLineUpdateDto) Get() *InvoiceLineUpdateDto {
	return v.value
}

func (v *NullableInvoiceLineUpdateDto) Set(val *InvoiceLineUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceLineUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceLineUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceLineUpdateDto(val *InvoiceLineUpdateDto) *NullableInvoiceLineUpdateDto {
	return &NullableInvoiceLineUpdateDto{value: val, isSet: true}
}

func (v NullableInvoiceLineUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceLineUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

