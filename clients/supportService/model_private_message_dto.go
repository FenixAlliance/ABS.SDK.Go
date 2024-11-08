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

// checks if the PrivateMessageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateMessageDto{}

// PrivateMessageDto struct for PrivateMessageDto
type PrivateMessageDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Read *bool `json:"read,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Message NullableString `json:"message,omitempty"`
	ConversationId NullableString `json:"conversationId,omitempty"`
	SenderSocialProfileId NullableString `json:"senderSocialProfileId,omitempty"`
	ReceiverSocialProfileID NullableString `json:"receiverSocialProfileID,omitempty"`
	SentTimestamp *time.Time `json:"sentTimestamp,omitempty"`
	ReadTimestamp *time.Time `json:"readTimestamp,omitempty"`
	ReceivedTimestamp *time.Time `json:"receivedTimestamp,omitempty"`
}

// NewPrivateMessageDto instantiates a new PrivateMessageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateMessageDto() *PrivateMessageDto {
	this := PrivateMessageDto{}
	return &this
}

// NewPrivateMessageDtoWithDefaults instantiates a new PrivateMessageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateMessageDtoWithDefaults() *PrivateMessageDto {
	this := PrivateMessageDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PrivateMessageDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PrivateMessageDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PrivateMessageDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *PrivateMessageDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *PrivateMessageDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *PrivateMessageDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *PrivateMessageDto) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateMessageDto) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *PrivateMessageDto) SetRead(v bool) {
	o.Read = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *PrivateMessageDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *PrivateMessageDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *PrivateMessageDto) UnsetTitle() {
	o.Title.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *PrivateMessageDto) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *PrivateMessageDto) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *PrivateMessageDto) UnsetMessage() {
	o.Message.Unset()
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConversationId.Get()
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConversationId.Get(), o.ConversationId.IsSet()
}

// HasConversationId returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasConversationId() bool {
	if o != nil && o.ConversationId.IsSet() {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given NullableString and assigns it to the ConversationId field.
func (o *PrivateMessageDto) SetConversationId(v string) {
	o.ConversationId.Set(&v)
}
// SetConversationIdNil sets the value for ConversationId to be an explicit nil
func (o *PrivateMessageDto) SetConversationIdNil() {
	o.ConversationId.Set(nil)
}

// UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
func (o *PrivateMessageDto) UnsetConversationId() {
	o.ConversationId.Unset()
}

// GetSenderSocialProfileId returns the SenderSocialProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetSenderSocialProfileId() string {
	if o == nil || IsNil(o.SenderSocialProfileId.Get()) {
		var ret string
		return ret
	}
	return *o.SenderSocialProfileId.Get()
}

// GetSenderSocialProfileIdOk returns a tuple with the SenderSocialProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetSenderSocialProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderSocialProfileId.Get(), o.SenderSocialProfileId.IsSet()
}

// HasSenderSocialProfileId returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasSenderSocialProfileId() bool {
	if o != nil && o.SenderSocialProfileId.IsSet() {
		return true
	}

	return false
}

// SetSenderSocialProfileId gets a reference to the given NullableString and assigns it to the SenderSocialProfileId field.
func (o *PrivateMessageDto) SetSenderSocialProfileId(v string) {
	o.SenderSocialProfileId.Set(&v)
}
// SetSenderSocialProfileIdNil sets the value for SenderSocialProfileId to be an explicit nil
func (o *PrivateMessageDto) SetSenderSocialProfileIdNil() {
	o.SenderSocialProfileId.Set(nil)
}

// UnsetSenderSocialProfileId ensures that no value is present for SenderSocialProfileId, not even an explicit nil
func (o *PrivateMessageDto) UnsetSenderSocialProfileId() {
	o.SenderSocialProfileId.Unset()
}

// GetReceiverSocialProfileID returns the ReceiverSocialProfileID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateMessageDto) GetReceiverSocialProfileID() string {
	if o == nil || IsNil(o.ReceiverSocialProfileID.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverSocialProfileID.Get()
}

// GetReceiverSocialProfileIDOk returns a tuple with the ReceiverSocialProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateMessageDto) GetReceiverSocialProfileIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverSocialProfileID.Get(), o.ReceiverSocialProfileID.IsSet()
}

// HasReceiverSocialProfileID returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasReceiverSocialProfileID() bool {
	if o != nil && o.ReceiverSocialProfileID.IsSet() {
		return true
	}

	return false
}

// SetReceiverSocialProfileID gets a reference to the given NullableString and assigns it to the ReceiverSocialProfileID field.
func (o *PrivateMessageDto) SetReceiverSocialProfileID(v string) {
	o.ReceiverSocialProfileID.Set(&v)
}
// SetReceiverSocialProfileIDNil sets the value for ReceiverSocialProfileID to be an explicit nil
func (o *PrivateMessageDto) SetReceiverSocialProfileIDNil() {
	o.ReceiverSocialProfileID.Set(nil)
}

// UnsetReceiverSocialProfileID ensures that no value is present for ReceiverSocialProfileID, not even an explicit nil
func (o *PrivateMessageDto) UnsetReceiverSocialProfileID() {
	o.ReceiverSocialProfileID.Unset()
}

// GetSentTimestamp returns the SentTimestamp field value if set, zero value otherwise.
func (o *PrivateMessageDto) GetSentTimestamp() time.Time {
	if o == nil || IsNil(o.SentTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SentTimestamp
}

// GetSentTimestampOk returns a tuple with the SentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateMessageDto) GetSentTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentTimestamp) {
		return nil, false
	}
	return o.SentTimestamp, true
}

// HasSentTimestamp returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasSentTimestamp() bool {
	if o != nil && !IsNil(o.SentTimestamp) {
		return true
	}

	return false
}

// SetSentTimestamp gets a reference to the given time.Time and assigns it to the SentTimestamp field.
func (o *PrivateMessageDto) SetSentTimestamp(v time.Time) {
	o.SentTimestamp = &v
}

// GetReadTimestamp returns the ReadTimestamp field value if set, zero value otherwise.
func (o *PrivateMessageDto) GetReadTimestamp() time.Time {
	if o == nil || IsNil(o.ReadTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ReadTimestamp
}

// GetReadTimestampOk returns a tuple with the ReadTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateMessageDto) GetReadTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReadTimestamp) {
		return nil, false
	}
	return o.ReadTimestamp, true
}

// HasReadTimestamp returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasReadTimestamp() bool {
	if o != nil && !IsNil(o.ReadTimestamp) {
		return true
	}

	return false
}

// SetReadTimestamp gets a reference to the given time.Time and assigns it to the ReadTimestamp field.
func (o *PrivateMessageDto) SetReadTimestamp(v time.Time) {
	o.ReadTimestamp = &v
}

// GetReceivedTimestamp returns the ReceivedTimestamp field value if set, zero value otherwise.
func (o *PrivateMessageDto) GetReceivedTimestamp() time.Time {
	if o == nil || IsNil(o.ReceivedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ReceivedTimestamp
}

// GetReceivedTimestampOk returns a tuple with the ReceivedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateMessageDto) GetReceivedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReceivedTimestamp) {
		return nil, false
	}
	return o.ReceivedTimestamp, true
}

// HasReceivedTimestamp returns a boolean if a field has been set.
func (o *PrivateMessageDto) HasReceivedTimestamp() bool {
	if o != nil && !IsNil(o.ReceivedTimestamp) {
		return true
	}

	return false
}

// SetReceivedTimestamp gets a reference to the given time.Time and assigns it to the ReceivedTimestamp field.
func (o *PrivateMessageDto) SetReceivedTimestamp(v time.Time) {
	o.ReceivedTimestamp = &v
}

func (o PrivateMessageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateMessageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.ConversationId.IsSet() {
		toSerialize["conversationId"] = o.ConversationId.Get()
	}
	if o.SenderSocialProfileId.IsSet() {
		toSerialize["senderSocialProfileId"] = o.SenderSocialProfileId.Get()
	}
	if o.ReceiverSocialProfileID.IsSet() {
		toSerialize["receiverSocialProfileID"] = o.ReceiverSocialProfileID.Get()
	}
	if !IsNil(o.SentTimestamp) {
		toSerialize["sentTimestamp"] = o.SentTimestamp
	}
	if !IsNil(o.ReadTimestamp) {
		toSerialize["readTimestamp"] = o.ReadTimestamp
	}
	if !IsNil(o.ReceivedTimestamp) {
		toSerialize["receivedTimestamp"] = o.ReceivedTimestamp
	}
	return toSerialize, nil
}

type NullablePrivateMessageDto struct {
	value *PrivateMessageDto
	isSet bool
}

func (v NullablePrivateMessageDto) Get() *PrivateMessageDto {
	return v.value
}

func (v *NullablePrivateMessageDto) Set(val *PrivateMessageDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateMessageDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateMessageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateMessageDto(val *PrivateMessageDto) *NullablePrivateMessageDto {
	return &NullablePrivateMessageDto{value: val, isSet: true}
}

func (v NullablePrivateMessageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateMessageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


