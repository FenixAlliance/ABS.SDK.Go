/*
MarketingService

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

// checks if the NewsletterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewsletterDto{}

// NewsletterDto struct for NewsletterDto
type NewsletterDto struct {
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Title NullableString `json:"title,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
}

// NewNewsletterDto instantiates a new NewsletterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsletterDto() *NewsletterDto {
	this := NewsletterDto{}
	return &this
}

// NewNewsletterDtoWithDefaults instantiates a new NewsletterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsletterDtoWithDefaults() *NewsletterDto {
	this := NewsletterDto{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *NewsletterDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *NewsletterDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *NewsletterDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *NewsletterDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *NewsletterDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *NewsletterDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *NewsletterDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *NewsletterDto) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NewsletterDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NewsletterDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NewsletterDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NewsletterDto) UnsetName() {
	o.Name.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *NewsletterDto) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *NewsletterDto) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *NewsletterDto) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *NewsletterDto) UnsetCode() {
	o.Code.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *NewsletterDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *NewsletterDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *NewsletterDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *NewsletterDto) UnsetTitle() {
	o.Title.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *NewsletterDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *NewsletterDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *NewsletterDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *NewsletterDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewsletterDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewsletterDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *NewsletterDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *NewsletterDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *NewsletterDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *NewsletterDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

func (o NewsletterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewsletterDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrolmentId.IsSet() {
		toSerialize["enrolmentId"] = o.EnrolmentId.Get()
	}
	return toSerialize, nil
}

type NullableNewsletterDto struct {
	value *NewsletterDto
	isSet bool
}

func (v NullableNewsletterDto) Get() *NewsletterDto {
	return v.value
}

func (v *NullableNewsletterDto) Set(val *NewsletterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsletterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsletterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsletterDto(val *NewsletterDto) *NullableNewsletterDto {
	return &NullableNewsletterDto{value: val, isSet: true}
}

func (v NullableNewsletterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsletterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


