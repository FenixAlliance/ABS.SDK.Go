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

// checks if the SalesLiteratureCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesLiteratureCreateDto{}

// SalesLiteratureCreateDto struct for SalesLiteratureCreateDto
type SalesLiteratureCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Content NullableString `json:"content,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrolmentId NullableString `json:"enrolmentId,omitempty"`
	SalesLiteratureTypeId NullableString `json:"salesLiteratureTypeId,omitempty"`
}

// NewSalesLiteratureCreateDto instantiates a new SalesLiteratureCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesLiteratureCreateDto() *SalesLiteratureCreateDto {
	this := SalesLiteratureCreateDto{}
	return &this
}

// NewSalesLiteratureCreateDtoWithDefaults instantiates a new SalesLiteratureCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesLiteratureCreateDtoWithDefaults() *SalesLiteratureCreateDto {
	this := SalesLiteratureCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SalesLiteratureCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesLiteratureCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SalesLiteratureCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SalesLiteratureCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesLiteratureCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SalesLiteratureCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SalesLiteratureCreateDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SalesLiteratureCreateDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetTitle() {
	o.Title.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *SalesLiteratureCreateDto) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *SalesLiteratureCreateDto) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetContent() {
	o.Content.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SalesLiteratureCreateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SalesLiteratureCreateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *SalesLiteratureCreateDto) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesLiteratureCreateDto) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasModifiedDate() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *SalesLiteratureCreateDto) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SalesLiteratureCreateDto) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesLiteratureCreateDto) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *SalesLiteratureCreateDto) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *SalesLiteratureCreateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *SalesLiteratureCreateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrolmentId returns the EnrolmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetEnrolmentId() string {
	if o == nil || IsNil(o.EnrolmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrolmentId.Get()
}

// GetEnrolmentIdOk returns a tuple with the EnrolmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetEnrolmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrolmentId.Get(), o.EnrolmentId.IsSet()
}

// HasEnrolmentId returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasEnrolmentId() bool {
	if o != nil && o.EnrolmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrolmentId gets a reference to the given NullableString and assigns it to the EnrolmentId field.
func (o *SalesLiteratureCreateDto) SetEnrolmentId(v string) {
	o.EnrolmentId.Set(&v)
}
// SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil
func (o *SalesLiteratureCreateDto) SetEnrolmentIdNil() {
	o.EnrolmentId.Set(nil)
}

// UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetEnrolmentId() {
	o.EnrolmentId.Unset()
}

// GetSalesLiteratureTypeId returns the SalesLiteratureTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesLiteratureCreateDto) GetSalesLiteratureTypeId() string {
	if o == nil || IsNil(o.SalesLiteratureTypeId.Get()) {
		var ret string
		return ret
	}
	return *o.SalesLiteratureTypeId.Get()
}

// GetSalesLiteratureTypeIdOk returns a tuple with the SalesLiteratureTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesLiteratureCreateDto) GetSalesLiteratureTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SalesLiteratureTypeId.Get(), o.SalesLiteratureTypeId.IsSet()
}

// HasSalesLiteratureTypeId returns a boolean if a field has been set.
func (o *SalesLiteratureCreateDto) HasSalesLiteratureTypeId() bool {
	if o != nil && o.SalesLiteratureTypeId.IsSet() {
		return true
	}

	return false
}

// SetSalesLiteratureTypeId gets a reference to the given NullableString and assigns it to the SalesLiteratureTypeId field.
func (o *SalesLiteratureCreateDto) SetSalesLiteratureTypeId(v string) {
	o.SalesLiteratureTypeId.Set(&v)
}
// SetSalesLiteratureTypeIdNil sets the value for SalesLiteratureTypeId to be an explicit nil
func (o *SalesLiteratureCreateDto) SetSalesLiteratureTypeIdNil() {
	o.SalesLiteratureTypeId.Set(nil)
}

// UnsetSalesLiteratureTypeId ensures that no value is present for SalesLiteratureTypeId, not even an explicit nil
func (o *SalesLiteratureCreateDto) UnsetSalesLiteratureTypeId() {
	o.SalesLiteratureTypeId.Unset()
}

func (o SalesLiteratureCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesLiteratureCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrolmentId.IsSet() {
		toSerialize["enrolmentId"] = o.EnrolmentId.Get()
	}
	if o.SalesLiteratureTypeId.IsSet() {
		toSerialize["salesLiteratureTypeId"] = o.SalesLiteratureTypeId.Get()
	}
	return toSerialize, nil
}

type NullableSalesLiteratureCreateDto struct {
	value *SalesLiteratureCreateDto
	isSet bool
}

func (v NullableSalesLiteratureCreateDto) Get() *SalesLiteratureCreateDto {
	return v.value
}

func (v *NullableSalesLiteratureCreateDto) Set(val *SalesLiteratureCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesLiteratureCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesLiteratureCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesLiteratureCreateDto(val *SalesLiteratureCreateDto) *NullableSalesLiteratureCreateDto {
	return &NullableSalesLiteratureCreateDto{value: val, isSet: true}
}

func (v NullableSalesLiteratureCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesLiteratureCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


