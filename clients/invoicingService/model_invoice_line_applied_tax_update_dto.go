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

// checks if the InvoiceLineAppliedTaxUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceLineAppliedTaxUpdateDto{}

// InvoiceLineAppliedTaxUpdateDto struct for InvoiceLineAppliedTaxUpdateDto
type InvoiceLineAppliedTaxUpdateDto struct {
	TaxPolicyId NullableString `json:"taxPolicyId,omitempty"`
}

// NewInvoiceLineAppliedTaxUpdateDto instantiates a new InvoiceLineAppliedTaxUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceLineAppliedTaxUpdateDto() *InvoiceLineAppliedTaxUpdateDto {
	this := InvoiceLineAppliedTaxUpdateDto{}
	return &this
}

// NewInvoiceLineAppliedTaxUpdateDtoWithDefaults instantiates a new InvoiceLineAppliedTaxUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceLineAppliedTaxUpdateDtoWithDefaults() *InvoiceLineAppliedTaxUpdateDto {
	this := InvoiceLineAppliedTaxUpdateDto{}
	return &this
}

// GetTaxPolicyId returns the TaxPolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineAppliedTaxUpdateDto) GetTaxPolicyId() string {
	if o == nil || IsNil(o.TaxPolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.TaxPolicyId.Get()
}

// GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineAppliedTaxUpdateDto) GetTaxPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxPolicyId.Get(), o.TaxPolicyId.IsSet()
}

// HasTaxPolicyId returns a boolean if a field has been set.
func (o *InvoiceLineAppliedTaxUpdateDto) HasTaxPolicyId() bool {
	if o != nil && o.TaxPolicyId.IsSet() {
		return true
	}

	return false
}

// SetTaxPolicyId gets a reference to the given NullableString and assigns it to the TaxPolicyId field.
func (o *InvoiceLineAppliedTaxUpdateDto) SetTaxPolicyId(v string) {
	o.TaxPolicyId.Set(&v)
}
// SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil
func (o *InvoiceLineAppliedTaxUpdateDto) SetTaxPolicyIdNil() {
	o.TaxPolicyId.Set(nil)
}

// UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
func (o *InvoiceLineAppliedTaxUpdateDto) UnsetTaxPolicyId() {
	o.TaxPolicyId.Unset()
}

func (o InvoiceLineAppliedTaxUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceLineAppliedTaxUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxPolicyId.IsSet() {
		toSerialize["taxPolicyId"] = o.TaxPolicyId.Get()
	}
	return toSerialize, nil
}

type NullableInvoiceLineAppliedTaxUpdateDto struct {
	value *InvoiceLineAppliedTaxUpdateDto
	isSet bool
}

func (v NullableInvoiceLineAppliedTaxUpdateDto) Get() *InvoiceLineAppliedTaxUpdateDto {
	return v.value
}

func (v *NullableInvoiceLineAppliedTaxUpdateDto) Set(val *InvoiceLineAppliedTaxUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceLineAppliedTaxUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceLineAppliedTaxUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceLineAppliedTaxUpdateDto(val *InvoiceLineAppliedTaxUpdateDto) *NullableInvoiceLineAppliedTaxUpdateDto {
	return &NullableInvoiceLineAppliedTaxUpdateDto{value: val, isSet: true}
}

func (v NullableInvoiceLineAppliedTaxUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceLineAppliedTaxUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


