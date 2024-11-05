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

// checks if the InvoiceReferenceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceReferenceDto{}

// InvoiceReferenceDto struct for InvoiceReferenceDto
type InvoiceReferenceDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrollmentId NullableString `json:"enrollmentId,omitempty"`
	ReferralInvoiceId NullableString `json:"referralInvoiceId,omitempty"`
	ReferencedInvoiceId NullableString `json:"referencedInvoiceId,omitempty"`
}

// NewInvoiceReferenceDto instantiates a new InvoiceReferenceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceReferenceDto() *InvoiceReferenceDto {
	this := InvoiceReferenceDto{}
	return &this
}

// NewInvoiceReferenceDtoWithDefaults instantiates a new InvoiceReferenceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceReferenceDtoWithDefaults() *InvoiceReferenceDto {
	this := InvoiceReferenceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *InvoiceReferenceDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InvoiceReferenceDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *InvoiceReferenceDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *InvoiceReferenceDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *InvoiceReferenceDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *InvoiceReferenceDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrollmentId returns the EnrollmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetEnrollmentId() string {
	if o == nil || IsNil(o.EnrollmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentId.Get()
}

// GetEnrollmentIdOk returns a tuple with the EnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentId.Get(), o.EnrollmentId.IsSet()
}

// HasEnrollmentId returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasEnrollmentId() bool {
	if o != nil && o.EnrollmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentId gets a reference to the given NullableString and assigns it to the EnrollmentId field.
func (o *InvoiceReferenceDto) SetEnrollmentId(v string) {
	o.EnrollmentId.Set(&v)
}
// SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil
func (o *InvoiceReferenceDto) SetEnrollmentIdNil() {
	o.EnrollmentId.Set(nil)
}

// UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetEnrollmentId() {
	o.EnrollmentId.Unset()
}

// GetReferralInvoiceId returns the ReferralInvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetReferralInvoiceId() string {
	if o == nil || IsNil(o.ReferralInvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferralInvoiceId.Get()
}

// GetReferralInvoiceIdOk returns a tuple with the ReferralInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetReferralInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferralInvoiceId.Get(), o.ReferralInvoiceId.IsSet()
}

// HasReferralInvoiceId returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasReferralInvoiceId() bool {
	if o != nil && o.ReferralInvoiceId.IsSet() {
		return true
	}

	return false
}

// SetReferralInvoiceId gets a reference to the given NullableString and assigns it to the ReferralInvoiceId field.
func (o *InvoiceReferenceDto) SetReferralInvoiceId(v string) {
	o.ReferralInvoiceId.Set(&v)
}
// SetReferralInvoiceIdNil sets the value for ReferralInvoiceId to be an explicit nil
func (o *InvoiceReferenceDto) SetReferralInvoiceIdNil() {
	o.ReferralInvoiceId.Set(nil)
}

// UnsetReferralInvoiceId ensures that no value is present for ReferralInvoiceId, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetReferralInvoiceId() {
	o.ReferralInvoiceId.Unset()
}

// GetReferencedInvoiceId returns the ReferencedInvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceReferenceDto) GetReferencedInvoiceId() string {
	if o == nil || IsNil(o.ReferencedInvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferencedInvoiceId.Get()
}

// GetReferencedInvoiceIdOk returns a tuple with the ReferencedInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceReferenceDto) GetReferencedInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferencedInvoiceId.Get(), o.ReferencedInvoiceId.IsSet()
}

// HasReferencedInvoiceId returns a boolean if a field has been set.
func (o *InvoiceReferenceDto) HasReferencedInvoiceId() bool {
	if o != nil && o.ReferencedInvoiceId.IsSet() {
		return true
	}

	return false
}

// SetReferencedInvoiceId gets a reference to the given NullableString and assigns it to the ReferencedInvoiceId field.
func (o *InvoiceReferenceDto) SetReferencedInvoiceId(v string) {
	o.ReferencedInvoiceId.Set(&v)
}
// SetReferencedInvoiceIdNil sets the value for ReferencedInvoiceId to be an explicit nil
func (o *InvoiceReferenceDto) SetReferencedInvoiceIdNil() {
	o.ReferencedInvoiceId.Set(nil)
}

// UnsetReferencedInvoiceId ensures that no value is present for ReferencedInvoiceId, not even an explicit nil
func (o *InvoiceReferenceDto) UnsetReferencedInvoiceId() {
	o.ReferencedInvoiceId.Unset()
}

func (o InvoiceReferenceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceReferenceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrollmentId.IsSet() {
		toSerialize["enrollmentId"] = o.EnrollmentId.Get()
	}
	if o.ReferralInvoiceId.IsSet() {
		toSerialize["referralInvoiceId"] = o.ReferralInvoiceId.Get()
	}
	if o.ReferencedInvoiceId.IsSet() {
		toSerialize["referencedInvoiceId"] = o.ReferencedInvoiceId.Get()
	}
	return toSerialize, nil
}

type NullableInvoiceReferenceDto struct {
	value *InvoiceReferenceDto
	isSet bool
}

func (v NullableInvoiceReferenceDto) Get() *InvoiceReferenceDto {
	return v.value
}

func (v *NullableInvoiceReferenceDto) Set(val *InvoiceReferenceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceReferenceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceReferenceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceReferenceDto(val *InvoiceReferenceDto) *NullableInvoiceReferenceDto {
	return &NullableInvoiceReferenceDto{value: val, isSet: true}
}

func (v NullableInvoiceReferenceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceReferenceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

