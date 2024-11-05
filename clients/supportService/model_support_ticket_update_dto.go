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
)

// checks if the SupportTicketUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicketUpdateDto{}

// SupportTicketUpdateDto struct for SupportTicketUpdateDto
type SupportTicketUpdateDto struct {
	Description NullableString `json:"description,omitempty"`
	AccountHolderID NullableString `json:"accountHolderID,omitempty"`
	ContactID NullableString `json:"contactID,omitempty"`
	BusinessProfileRecordID NullableString `json:"businessProfileRecordID,omitempty"`
	SupportTicketTypeID NullableString `json:"supportTicketTypeID,omitempty"`
	SupportEntitlementID NullableString `json:"supportEntitlementID,omitempty"`
	SupportPriorityID NullableString `json:"supportPriorityID,omitempty"`
}

// NewSupportTicketUpdateDto instantiates a new SupportTicketUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicketUpdateDto() *SupportTicketUpdateDto {
	this := SupportTicketUpdateDto{}
	return &this
}

// NewSupportTicketUpdateDtoWithDefaults instantiates a new SupportTicketUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketUpdateDtoWithDefaults() *SupportTicketUpdateDto {
	this := SupportTicketUpdateDto{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SupportTicketUpdateDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SupportTicketUpdateDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetDescription() {
	o.Description.Unset()
}

// GetAccountHolderID returns the AccountHolderID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetAccountHolderID() string {
	if o == nil || IsNil(o.AccountHolderID.Get()) {
		var ret string
		return ret
	}
	return *o.AccountHolderID.Get()
}

// GetAccountHolderIDOk returns a tuple with the AccountHolderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetAccountHolderIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderID.Get(), o.AccountHolderID.IsSet()
}

// HasAccountHolderID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasAccountHolderID() bool {
	if o != nil && o.AccountHolderID.IsSet() {
		return true
	}

	return false
}

// SetAccountHolderID gets a reference to the given NullableString and assigns it to the AccountHolderID field.
func (o *SupportTicketUpdateDto) SetAccountHolderID(v string) {
	o.AccountHolderID.Set(&v)
}
// SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil
func (o *SupportTicketUpdateDto) SetAccountHolderIDNil() {
	o.AccountHolderID.Set(nil)
}

// UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetAccountHolderID() {
	o.AccountHolderID.Unset()
}

// GetContactID returns the ContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetContactID() string {
	if o == nil || IsNil(o.ContactID.Get()) {
		var ret string
		return ret
	}
	return *o.ContactID.Get()
}

// GetContactIDOk returns a tuple with the ContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetContactIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactID.Get(), o.ContactID.IsSet()
}

// HasContactID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasContactID() bool {
	if o != nil && o.ContactID.IsSet() {
		return true
	}

	return false
}

// SetContactID gets a reference to the given NullableString and assigns it to the ContactID field.
func (o *SupportTicketUpdateDto) SetContactID(v string) {
	o.ContactID.Set(&v)
}
// SetContactIDNil sets the value for ContactID to be an explicit nil
func (o *SupportTicketUpdateDto) SetContactIDNil() {
	o.ContactID.Set(nil)
}

// UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetContactID() {
	o.ContactID.Unset()
}

// GetBusinessProfileRecordID returns the BusinessProfileRecordID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetBusinessProfileRecordID() string {
	if o == nil || IsNil(o.BusinessProfileRecordID.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessProfileRecordID.Get()
}

// GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetBusinessProfileRecordIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessProfileRecordID.Get(), o.BusinessProfileRecordID.IsSet()
}

// HasBusinessProfileRecordID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasBusinessProfileRecordID() bool {
	if o != nil && o.BusinessProfileRecordID.IsSet() {
		return true
	}

	return false
}

// SetBusinessProfileRecordID gets a reference to the given NullableString and assigns it to the BusinessProfileRecordID field.
func (o *SupportTicketUpdateDto) SetBusinessProfileRecordID(v string) {
	o.BusinessProfileRecordID.Set(&v)
}
// SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil
func (o *SupportTicketUpdateDto) SetBusinessProfileRecordIDNil() {
	o.BusinessProfileRecordID.Set(nil)
}

// UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetBusinessProfileRecordID() {
	o.BusinessProfileRecordID.Unset()
}

// GetSupportTicketTypeID returns the SupportTicketTypeID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetSupportTicketTypeID() string {
	if o == nil || IsNil(o.SupportTicketTypeID.Get()) {
		var ret string
		return ret
	}
	return *o.SupportTicketTypeID.Get()
}

// GetSupportTicketTypeIDOk returns a tuple with the SupportTicketTypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetSupportTicketTypeIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportTicketTypeID.Get(), o.SupportTicketTypeID.IsSet()
}

// HasSupportTicketTypeID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasSupportTicketTypeID() bool {
	if o != nil && o.SupportTicketTypeID.IsSet() {
		return true
	}

	return false
}

// SetSupportTicketTypeID gets a reference to the given NullableString and assigns it to the SupportTicketTypeID field.
func (o *SupportTicketUpdateDto) SetSupportTicketTypeID(v string) {
	o.SupportTicketTypeID.Set(&v)
}
// SetSupportTicketTypeIDNil sets the value for SupportTicketTypeID to be an explicit nil
func (o *SupportTicketUpdateDto) SetSupportTicketTypeIDNil() {
	o.SupportTicketTypeID.Set(nil)
}

// UnsetSupportTicketTypeID ensures that no value is present for SupportTicketTypeID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetSupportTicketTypeID() {
	o.SupportTicketTypeID.Unset()
}

// GetSupportEntitlementID returns the SupportEntitlementID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetSupportEntitlementID() string {
	if o == nil || IsNil(o.SupportEntitlementID.Get()) {
		var ret string
		return ret
	}
	return *o.SupportEntitlementID.Get()
}

// GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetSupportEntitlementIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportEntitlementID.Get(), o.SupportEntitlementID.IsSet()
}

// HasSupportEntitlementID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasSupportEntitlementID() bool {
	if o != nil && o.SupportEntitlementID.IsSet() {
		return true
	}

	return false
}

// SetSupportEntitlementID gets a reference to the given NullableString and assigns it to the SupportEntitlementID field.
func (o *SupportTicketUpdateDto) SetSupportEntitlementID(v string) {
	o.SupportEntitlementID.Set(&v)
}
// SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil
func (o *SupportTicketUpdateDto) SetSupportEntitlementIDNil() {
	o.SupportEntitlementID.Set(nil)
}

// UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetSupportEntitlementID() {
	o.SupportEntitlementID.Unset()
}

// GetSupportPriorityID returns the SupportPriorityID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketUpdateDto) GetSupportPriorityID() string {
	if o == nil || IsNil(o.SupportPriorityID.Get()) {
		var ret string
		return ret
	}
	return *o.SupportPriorityID.Get()
}

// GetSupportPriorityIDOk returns a tuple with the SupportPriorityID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketUpdateDto) GetSupportPriorityIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportPriorityID.Get(), o.SupportPriorityID.IsSet()
}

// HasSupportPriorityID returns a boolean if a field has been set.
func (o *SupportTicketUpdateDto) HasSupportPriorityID() bool {
	if o != nil && o.SupportPriorityID.IsSet() {
		return true
	}

	return false
}

// SetSupportPriorityID gets a reference to the given NullableString and assigns it to the SupportPriorityID field.
func (o *SupportTicketUpdateDto) SetSupportPriorityID(v string) {
	o.SupportPriorityID.Set(&v)
}
// SetSupportPriorityIDNil sets the value for SupportPriorityID to be an explicit nil
func (o *SupportTicketUpdateDto) SetSupportPriorityIDNil() {
	o.SupportPriorityID.Set(nil)
}

// UnsetSupportPriorityID ensures that no value is present for SupportPriorityID, not even an explicit nil
func (o *SupportTicketUpdateDto) UnsetSupportPriorityID() {
	o.SupportPriorityID.Unset()
}

func (o SupportTicketUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicketUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AccountHolderID.IsSet() {
		toSerialize["accountHolderID"] = o.AccountHolderID.Get()
	}
	if o.ContactID.IsSet() {
		toSerialize["contactID"] = o.ContactID.Get()
	}
	if o.BusinessProfileRecordID.IsSet() {
		toSerialize["businessProfileRecordID"] = o.BusinessProfileRecordID.Get()
	}
	if o.SupportTicketTypeID.IsSet() {
		toSerialize["supportTicketTypeID"] = o.SupportTicketTypeID.Get()
	}
	if o.SupportEntitlementID.IsSet() {
		toSerialize["supportEntitlementID"] = o.SupportEntitlementID.Get()
	}
	if o.SupportPriorityID.IsSet() {
		toSerialize["supportPriorityID"] = o.SupportPriorityID.Get()
	}
	return toSerialize, nil
}

type NullableSupportTicketUpdateDto struct {
	value *SupportTicketUpdateDto
	isSet bool
}

func (v NullableSupportTicketUpdateDto) Get() *SupportTicketUpdateDto {
	return v.value
}

func (v *NullableSupportTicketUpdateDto) Set(val *SupportTicketUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketUpdateDto(val *SupportTicketUpdateDto) *NullableSupportTicketUpdateDto {
	return &NullableSupportTicketUpdateDto{value: val, isSet: true}
}

func (v NullableSupportTicketUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

