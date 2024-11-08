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

// checks if the SupportTicketConversationCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicketConversationCreateDto{}

// SupportTicketConversationCreateDto struct for SupportTicketConversationCreateDto
type SupportTicketConversationCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	Closed *bool `json:"closed,omitempty"`
	ClosedTimestamp *time.Time `json:"closedTimestamp,omitempty"`
	SocialProfileID NullableString `json:"socialProfileID,omitempty"`
}

// NewSupportTicketConversationCreateDto instantiates a new SupportTicketConversationCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicketConversationCreateDto() *SupportTicketConversationCreateDto {
	this := SupportTicketConversationCreateDto{}
	return &this
}

// NewSupportTicketConversationCreateDtoWithDefaults instantiates a new SupportTicketConversationCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketConversationCreateDtoWithDefaults() *SupportTicketConversationCreateDto {
	this := SupportTicketConversationCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportTicketConversationCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketConversationCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportTicketConversationCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SupportTicketConversationCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketConversationCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SupportTicketConversationCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketConversationCreateDto) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketConversationCreateDto) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *SupportTicketConversationCreateDto) SetTopic(v string) {
	o.Topic.Set(&v)
}
// SetTopicNil sets the value for Topic to be an explicit nil
func (o *SupportTicketConversationCreateDto) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *SupportTicketConversationCreateDto) UnsetTopic() {
	o.Topic.Unset()
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *SupportTicketConversationCreateDto) GetClosed() bool {
	if o == nil || IsNil(o.Closed) {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketConversationCreateDto) GetClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.Closed) {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasClosed() bool {
	if o != nil && !IsNil(o.Closed) {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *SupportTicketConversationCreateDto) SetClosed(v bool) {
	o.Closed = &v
}

// GetClosedTimestamp returns the ClosedTimestamp field value if set, zero value otherwise.
func (o *SupportTicketConversationCreateDto) GetClosedTimestamp() time.Time {
	if o == nil || IsNil(o.ClosedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ClosedTimestamp
}

// GetClosedTimestampOk returns a tuple with the ClosedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketConversationCreateDto) GetClosedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedTimestamp) {
		return nil, false
	}
	return o.ClosedTimestamp, true
}

// HasClosedTimestamp returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasClosedTimestamp() bool {
	if o != nil && !IsNil(o.ClosedTimestamp) {
		return true
	}

	return false
}

// SetClosedTimestamp gets a reference to the given time.Time and assigns it to the ClosedTimestamp field.
func (o *SupportTicketConversationCreateDto) SetClosedTimestamp(v time.Time) {
	o.ClosedTimestamp = &v
}

// GetSocialProfileID returns the SocialProfileID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicketConversationCreateDto) GetSocialProfileID() string {
	if o == nil || IsNil(o.SocialProfileID.Get()) {
		var ret string
		return ret
	}
	return *o.SocialProfileID.Get()
}

// GetSocialProfileIDOk returns a tuple with the SocialProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicketConversationCreateDto) GetSocialProfileIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialProfileID.Get(), o.SocialProfileID.IsSet()
}

// HasSocialProfileID returns a boolean if a field has been set.
func (o *SupportTicketConversationCreateDto) HasSocialProfileID() bool {
	if o != nil && o.SocialProfileID.IsSet() {
		return true
	}

	return false
}

// SetSocialProfileID gets a reference to the given NullableString and assigns it to the SocialProfileID field.
func (o *SupportTicketConversationCreateDto) SetSocialProfileID(v string) {
	o.SocialProfileID.Set(&v)
}
// SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil
func (o *SupportTicketConversationCreateDto) SetSocialProfileIDNil() {
	o.SocialProfileID.Set(nil)
}

// UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
func (o *SupportTicketConversationCreateDto) UnsetSocialProfileID() {
	o.SocialProfileID.Unset()
}

func (o SupportTicketConversationCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicketConversationCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if !IsNil(o.Closed) {
		toSerialize["closed"] = o.Closed
	}
	if !IsNil(o.ClosedTimestamp) {
		toSerialize["closedTimestamp"] = o.ClosedTimestamp
	}
	if o.SocialProfileID.IsSet() {
		toSerialize["socialProfileID"] = o.SocialProfileID.Get()
	}
	return toSerialize, nil
}

type NullableSupportTicketConversationCreateDto struct {
	value *SupportTicketConversationCreateDto
	isSet bool
}

func (v NullableSupportTicketConversationCreateDto) Get() *SupportTicketConversationCreateDto {
	return v.value
}

func (v *NullableSupportTicketConversationCreateDto) Set(val *SupportTicketConversationCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketConversationCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketConversationCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketConversationCreateDto(val *SupportTicketConversationCreateDto) *NullableSupportTicketConversationCreateDto {
	return &NullableSupportTicketConversationCreateDto{value: val, isSet: true}
}

func (v NullableSupportTicketConversationCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketConversationCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


