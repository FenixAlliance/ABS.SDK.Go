/*
SupportService

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

// checks if the SupportTicketTypeCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicketTypeCreateDto{}

// SupportTicketTypeCreateDto struct for SupportTicketTypeCreateDto
type SupportTicketTypeCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	BusinessID NullableString `json:"businessID,omitempty"`
}

// NewSupportTicketTypeCreateDto instantiates a new SupportTicketTypeCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicketTypeCreateDto() *SupportTicketTypeCreateDto {
	this := SupportTicketTypeCreateDto{}
	return &this
}

// NewSupportTicketTypeCreateDtoWithDefaults instantiates a new SupportTicketTypeCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketTypeCreateDtoWithDefaults() *SupportTicketTypeCreateDto {
	this := SupportTicketTypeCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportTicketTypeCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketTypeCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportTicketTypeCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportTicketTypeCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SupportTicketTypeCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketTypeCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SupportTicketTypeCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SupportTicketTypeCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketTypeCreateDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketTypeCreateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SupportTicketTypeCreateDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SupportTicketTypeCreateDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SupportTicketTypeCreateDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SupportTicketTypeCreateDto) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketTypeCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketTypeCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SupportTicketTypeCreateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SupportTicketTypeCreateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SupportTicketTypeCreateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SupportTicketTypeCreateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetBusinessID returns the BusinessID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketTypeCreateDto) GetBusinessID() string {
	if o == nil || IsNil(o.BusinessID.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessID.Get()
}

// GetBusinessIDOk returns a tuple with the BusinessID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketTypeCreateDto) GetBusinessIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessID.Get(), o.BusinessID.IsSet()
}

// HasBusinessID returns a boolean if a field has been set.
func (o *SupportTicketTypeCreateDto) HasBusinessID() bool {
	if o != nil && o.BusinessID.IsSet() {
		return true
	}

	return false
}

// SetBusinessID gets a reference to the given NullableString and assigns it to the BusinessID field.
func (o *SupportTicketTypeCreateDto) SetBusinessID(v string) {
	o.BusinessID.Set(&v)
}
// SetBusinessIDNil sets the value for BusinessID to be an explicit nil
func (o *SupportTicketTypeCreateDto) SetBusinessIDNil() {
	o.BusinessID.Set(nil)
}

// UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
func (o *SupportTicketTypeCreateDto) UnsetBusinessID() {
	o.BusinessID.Unset()
}

func (o SupportTicketTypeCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicketTypeCreateDto) ToMap() (map[string]interface{}, error) {
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.BusinessID.IsSet() {
		toSerialize["businessID"] = o.BusinessID.Get()
	}
	return toSerialize, nil
}

type NullableSupportTicketTypeCreateDto struct {
	value *SupportTicketTypeCreateDto
	isSet bool
}

func (v NullableSupportTicketTypeCreateDto) Get() *SupportTicketTypeCreateDto {
	return v.value
}

func (v *NullableSupportTicketTypeCreateDto) Set(val *SupportTicketTypeCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketTypeCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketTypeCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketTypeCreateDto(val *SupportTicketTypeCreateDto) *NullableSupportTicketTypeCreateDto {
	return &NullableSupportTicketTypeCreateDto{value: val, isSet: true}
}

func (v NullableSupportTicketTypeCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketTypeCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


