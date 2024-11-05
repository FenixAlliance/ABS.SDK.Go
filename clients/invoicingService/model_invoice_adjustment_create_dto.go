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
	"time"
)

// checks if the InvoiceAdjustmentCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceAdjustmentCreateDto{}

// InvoiceAdjustmentCreateDto struct for InvoiceAdjustmentCreateDto
type InvoiceAdjustmentCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	InvoiceId NullableString `json:"invoiceId,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	EnrollmentId NullableString `json:"enrollmentId,omitempty"`
	Description NullableString `json:"description,omitempty"`
	SurchargePercent *float64 `json:"surchargePercent,omitempty"`
	SurchargeAmount *float64 `json:"surchargeAmount,omitempty"`
	DiscountPercent *float64 `json:"discountPercent,omitempty"`
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	TotalSurcharge *float64 `json:"totalSurcharge,omitempty"`
	TotalDiscount *float64 `json:"totalDiscount,omitempty"`
	Type *int32 `json:"type,omitempty"`
}

// NewInvoiceAdjustmentCreateDto instantiates a new InvoiceAdjustmentCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceAdjustmentCreateDto() *InvoiceAdjustmentCreateDto {
	this := InvoiceAdjustmentCreateDto{}
	return &this
}

// NewInvoiceAdjustmentCreateDtoWithDefaults instantiates a new InvoiceAdjustmentCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceAdjustmentCreateDtoWithDefaults() *InvoiceAdjustmentCreateDto {
	this := InvoiceAdjustmentCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvoiceAdjustmentCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *InvoiceAdjustmentCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceAdjustmentCreateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceAdjustmentCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *InvoiceAdjustmentCreateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *InvoiceAdjustmentCreateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *InvoiceAdjustmentCreateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceAdjustmentCreateDto) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceAdjustmentCreateDto) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *InvoiceAdjustmentCreateDto) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *InvoiceAdjustmentCreateDto) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *InvoiceAdjustmentCreateDto) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceAdjustmentCreateDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceAdjustmentCreateDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *InvoiceAdjustmentCreateDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *InvoiceAdjustmentCreateDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *InvoiceAdjustmentCreateDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetEnrollmentId returns the EnrollmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceAdjustmentCreateDto) GetEnrollmentId() string {
	if o == nil || IsNil(o.EnrollmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentId.Get()
}

// GetEnrollmentIdOk returns a tuple with the EnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceAdjustmentCreateDto) GetEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentId.Get(), o.EnrollmentId.IsSet()
}

// HasEnrollmentId returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasEnrollmentId() bool {
	if o != nil && o.EnrollmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentId gets a reference to the given NullableString and assigns it to the EnrollmentId field.
func (o *InvoiceAdjustmentCreateDto) SetEnrollmentId(v string) {
	o.EnrollmentId.Set(&v)
}
// SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil
func (o *InvoiceAdjustmentCreateDto) SetEnrollmentIdNil() {
	o.EnrollmentId.Set(nil)
}

// UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
func (o *InvoiceAdjustmentCreateDto) UnsetEnrollmentId() {
	o.EnrollmentId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceAdjustmentCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceAdjustmentCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InvoiceAdjustmentCreateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InvoiceAdjustmentCreateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InvoiceAdjustmentCreateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetSurchargePercent returns the SurchargePercent field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetSurchargePercent() float64 {
	if o == nil || IsNil(o.SurchargePercent) {
		var ret float64
		return ret
	}
	return *o.SurchargePercent
}

// GetSurchargePercentOk returns a tuple with the SurchargePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetSurchargePercentOk() (*float64, bool) {
	if o == nil || IsNil(o.SurchargePercent) {
		return nil, false
	}
	return o.SurchargePercent, true
}

// HasSurchargePercent returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasSurchargePercent() bool {
	if o != nil && !IsNil(o.SurchargePercent) {
		return true
	}

	return false
}

// SetSurchargePercent gets a reference to the given float64 and assigns it to the SurchargePercent field.
func (o *InvoiceAdjustmentCreateDto) SetSurchargePercent(v float64) {
	o.SurchargePercent = &v
}

// GetSurchargeAmount returns the SurchargeAmount field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetSurchargeAmount() float64 {
	if o == nil || IsNil(o.SurchargeAmount) {
		var ret float64
		return ret
	}
	return *o.SurchargeAmount
}

// GetSurchargeAmountOk returns a tuple with the SurchargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetSurchargeAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.SurchargeAmount) {
		return nil, false
	}
	return o.SurchargeAmount, true
}

// HasSurchargeAmount returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasSurchargeAmount() bool {
	if o != nil && !IsNil(o.SurchargeAmount) {
		return true
	}

	return false
}

// SetSurchargeAmount gets a reference to the given float64 and assigns it to the SurchargeAmount field.
func (o *InvoiceAdjustmentCreateDto) SetSurchargeAmount(v float64) {
	o.SurchargeAmount = &v
}

// GetDiscountPercent returns the DiscountPercent field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetDiscountPercent() float64 {
	if o == nil || IsNil(o.DiscountPercent) {
		var ret float64
		return ret
	}
	return *o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetDiscountPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountPercent) {
		return nil, false
	}
	return o.DiscountPercent, true
}

// HasDiscountPercent returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasDiscountPercent() bool {
	if o != nil && !IsNil(o.DiscountPercent) {
		return true
	}

	return false
}

// SetDiscountPercent gets a reference to the given float64 and assigns it to the DiscountPercent field.
func (o *InvoiceAdjustmentCreateDto) SetDiscountPercent(v float64) {
	o.DiscountPercent = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetDiscountAmount() float64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret float64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetDiscountAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given float64 and assigns it to the DiscountAmount field.
func (o *InvoiceAdjustmentCreateDto) SetDiscountAmount(v float64) {
	o.DiscountAmount = &v
}

// GetTotalSurcharge returns the TotalSurcharge field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetTotalSurcharge() float64 {
	if o == nil || IsNil(o.TotalSurcharge) {
		var ret float64
		return ret
	}
	return *o.TotalSurcharge
}

// GetTotalSurchargeOk returns a tuple with the TotalSurcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetTotalSurchargeOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalSurcharge) {
		return nil, false
	}
	return o.TotalSurcharge, true
}

// HasTotalSurcharge returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasTotalSurcharge() bool {
	if o != nil && !IsNil(o.TotalSurcharge) {
		return true
	}

	return false
}

// SetTotalSurcharge gets a reference to the given float64 and assigns it to the TotalSurcharge field.
func (o *InvoiceAdjustmentCreateDto) SetTotalSurcharge(v float64) {
	o.TotalSurcharge = &v
}

// GetTotalDiscount returns the TotalDiscount field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetTotalDiscount() float64 {
	if o == nil || IsNil(o.TotalDiscount) {
		var ret float64
		return ret
	}
	return *o.TotalDiscount
}

// GetTotalDiscountOk returns a tuple with the TotalDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetTotalDiscountOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDiscount) {
		return nil, false
	}
	return o.TotalDiscount, true
}

// HasTotalDiscount returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasTotalDiscount() bool {
	if o != nil && !IsNil(o.TotalDiscount) {
		return true
	}

	return false
}

// SetTotalDiscount gets a reference to the given float64 and assigns it to the TotalDiscount field.
func (o *InvoiceAdjustmentCreateDto) SetTotalDiscount(v float64) {
	o.TotalDiscount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InvoiceAdjustmentCreateDto) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustmentCreateDto) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InvoiceAdjustmentCreateDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *InvoiceAdjustmentCreateDto) SetType(v int32) {
	o.Type = &v
}

func (o InvoiceAdjustmentCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceAdjustmentCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoiceId"] = o.InvoiceId.Get()
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.EnrollmentId.IsSet() {
		toSerialize["enrollmentId"] = o.EnrollmentId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.SurchargePercent) {
		toSerialize["surchargePercent"] = o.SurchargePercent
	}
	if !IsNil(o.SurchargeAmount) {
		toSerialize["surchargeAmount"] = o.SurchargeAmount
	}
	if !IsNil(o.DiscountPercent) {
		toSerialize["discountPercent"] = o.DiscountPercent
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.TotalSurcharge) {
		toSerialize["totalSurcharge"] = o.TotalSurcharge
	}
	if !IsNil(o.TotalDiscount) {
		toSerialize["totalDiscount"] = o.TotalDiscount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInvoiceAdjustmentCreateDto struct {
	value *InvoiceAdjustmentCreateDto
	isSet bool
}

func (v NullableInvoiceAdjustmentCreateDto) Get() *InvoiceAdjustmentCreateDto {
	return v.value
}

func (v *NullableInvoiceAdjustmentCreateDto) Set(val *InvoiceAdjustmentCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceAdjustmentCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceAdjustmentCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceAdjustmentCreateDto(val *InvoiceAdjustmentCreateDto) *NullableInvoiceAdjustmentCreateDto {
	return &NullableInvoiceAdjustmentCreateDto{value: val, isSet: true}
}

func (v NullableInvoiceAdjustmentCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceAdjustmentCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

