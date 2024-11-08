/*
HolderService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CartDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartDto{}

// CartDto struct for CartDto
type CartDto struct {
	Id NullableString `json:"id,omitempty"`
	Ip NullableString `json:"ip,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Total *float64 `json:"total,omitempty"`
	Taxes *float64 `json:"taxes,omitempty"`
	Freight *float64 `json:"freight,omitempty"`
	SubTotal *float64 `json:"subTotal,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	CountryId NullableString `json:"countryId,omitempty"`
	ItemCartRecordsCount NullableInt32 `json:"itemCartRecordsCount,omitempty"`
	ItemToCompareRecordsCount NullableInt32 `json:"itemToCompareRecordsCount,omitempty"`
}

// NewCartDto instantiates a new CartDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartDto() *CartDto {
	this := CartDto{}
	return &this
}

// NewCartDtoWithDefaults instantiates a new CartDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartDtoWithDefaults() *CartDto {
	this := CartDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CartDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CartDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CartDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CartDto) UnsetId() {
	o.Id.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *CartDto) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *CartDto) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *CartDto) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *CartDto) UnsetIp() {
	o.Ip.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CartDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CartDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CartDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CartDto) UnsetType() {
	o.Type.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CartDto) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDto) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CartDto) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *CartDto) SetTotal(v float64) {
	o.Total = &v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *CartDto) GetTaxes() float64 {
	if o == nil || IsNil(o.Taxes) {
		var ret float64
		return ret
	}
	return *o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDto) GetTaxesOk() (*float64, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *CartDto) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given float64 and assigns it to the Taxes field.
func (o *CartDto) SetTaxes(v float64) {
	o.Taxes = &v
}

// GetFreight returns the Freight field value if set, zero value otherwise.
func (o *CartDto) GetFreight() float64 {
	if o == nil || IsNil(o.Freight) {
		var ret float64
		return ret
	}
	return *o.Freight
}

// GetFreightOk returns a tuple with the Freight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDto) GetFreightOk() (*float64, bool) {
	if o == nil || IsNil(o.Freight) {
		return nil, false
	}
	return o.Freight, true
}

// HasFreight returns a boolean if a field has been set.
func (o *CartDto) HasFreight() bool {
	if o != nil && !IsNil(o.Freight) {
		return true
	}

	return false
}

// SetFreight gets a reference to the given float64 and assigns it to the Freight field.
func (o *CartDto) SetFreight(v float64) {
	o.Freight = &v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *CartDto) GetSubTotal() float64 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float64
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartDto) GetSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotal returns a boolean if a field has been set.
func (o *CartDto) HasSubTotal() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float64 and assigns it to the SubTotal field.
func (o *CartDto) SetSubTotal(v float64) {
	o.SubTotal = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *CartDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *CartDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *CartDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *CartDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetCountryId returns the CountryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetCountryId() string {
	if o == nil || IsNil(o.CountryId.Get()) {
		var ret string
		return ret
	}
	return *o.CountryId.Get()
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryId.Get(), o.CountryId.IsSet()
}

// HasCountryId returns a boolean if a field has been set.
func (o *CartDto) HasCountryId() bool {
	if o != nil && o.CountryId.IsSet() {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given NullableString and assigns it to the CountryId field.
func (o *CartDto) SetCountryId(v string) {
	o.CountryId.Set(&v)
}
// SetCountryIdNil sets the value for CountryId to be an explicit nil
func (o *CartDto) SetCountryIdNil() {
	o.CountryId.Set(nil)
}

// UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
func (o *CartDto) UnsetCountryId() {
	o.CountryId.Unset()
}

// GetItemCartRecordsCount returns the ItemCartRecordsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetItemCartRecordsCount() int32 {
	if o == nil || IsNil(o.ItemCartRecordsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ItemCartRecordsCount.Get()
}

// GetItemCartRecordsCountOk returns a tuple with the ItemCartRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetItemCartRecordsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemCartRecordsCount.Get(), o.ItemCartRecordsCount.IsSet()
}

// HasItemCartRecordsCount returns a boolean if a field has been set.
func (o *CartDto) HasItemCartRecordsCount() bool {
	if o != nil && o.ItemCartRecordsCount.IsSet() {
		return true
	}

	return false
}

// SetItemCartRecordsCount gets a reference to the given NullableInt32 and assigns it to the ItemCartRecordsCount field.
func (o *CartDto) SetItemCartRecordsCount(v int32) {
	o.ItemCartRecordsCount.Set(&v)
}
// SetItemCartRecordsCountNil sets the value for ItemCartRecordsCount to be an explicit nil
func (o *CartDto) SetItemCartRecordsCountNil() {
	o.ItemCartRecordsCount.Set(nil)
}

// UnsetItemCartRecordsCount ensures that no value is present for ItemCartRecordsCount, not even an explicit nil
func (o *CartDto) UnsetItemCartRecordsCount() {
	o.ItemCartRecordsCount.Unset()
}

// GetItemToCompareRecordsCount returns the ItemToCompareRecordsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartDto) GetItemToCompareRecordsCount() int32 {
	if o == nil || IsNil(o.ItemToCompareRecordsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ItemToCompareRecordsCount.Get()
}

// GetItemToCompareRecordsCountOk returns a tuple with the ItemToCompareRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartDto) GetItemToCompareRecordsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemToCompareRecordsCount.Get(), o.ItemToCompareRecordsCount.IsSet()
}

// HasItemToCompareRecordsCount returns a boolean if a field has been set.
func (o *CartDto) HasItemToCompareRecordsCount() bool {
	if o != nil && o.ItemToCompareRecordsCount.IsSet() {
		return true
	}

	return false
}

// SetItemToCompareRecordsCount gets a reference to the given NullableInt32 and assigns it to the ItemToCompareRecordsCount field.
func (o *CartDto) SetItemToCompareRecordsCount(v int32) {
	o.ItemToCompareRecordsCount.Set(&v)
}
// SetItemToCompareRecordsCountNil sets the value for ItemToCompareRecordsCount to be an explicit nil
func (o *CartDto) SetItemToCompareRecordsCountNil() {
	o.ItemToCompareRecordsCount.Set(nil)
}

// UnsetItemToCompareRecordsCount ensures that no value is present for ItemToCompareRecordsCount, not even an explicit nil
func (o *CartDto) UnsetItemToCompareRecordsCount() {
	o.ItemToCompareRecordsCount.Unset()
}

func (o CartDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !IsNil(o.Freight) {
		toSerialize["freight"] = o.Freight
	}
	if !IsNil(o.SubTotal) {
		toSerialize["subTotal"] = o.SubTotal
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.CountryId.IsSet() {
		toSerialize["countryId"] = o.CountryId.Get()
	}
	if o.ItemCartRecordsCount.IsSet() {
		toSerialize["itemCartRecordsCount"] = o.ItemCartRecordsCount.Get()
	}
	if o.ItemToCompareRecordsCount.IsSet() {
		toSerialize["itemToCompareRecordsCount"] = o.ItemToCompareRecordsCount.Get()
	}
	return toSerialize, nil
}

type NullableCartDto struct {
	value *CartDto
	isSet bool
}

func (v NullableCartDto) Get() *CartDto {
	return v.value
}

func (v *NullableCartDto) Set(val *CartDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCartDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCartDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartDto(val *CartDto) *NullableCartDto {
	return &NullableCartDto{value: val, isSet: true}
}

func (v NullableCartDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


