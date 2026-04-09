# PrivateMessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Read** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ConversationId** | Pointer to **NullableString** |  | [optional] 
**SenderSocialProfileId** | Pointer to **NullableString** |  | [optional] 
**ReceiverSocialProfileID** | Pointer to **NullableString** |  | [optional] 
**SentTimestamp** | Pointer to **time.Time** |  | [optional] 
**ReadTimestamp** | Pointer to **time.Time** |  | [optional] 
**ReceivedTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPrivateMessageDto

`func NewPrivateMessageDto() *PrivateMessageDto`

NewPrivateMessageDto instantiates a new PrivateMessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateMessageDtoWithDefaults

`func NewPrivateMessageDtoWithDefaults() *PrivateMessageDto`

NewPrivateMessageDtoWithDefaults instantiates a new PrivateMessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateMessageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateMessageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateMessageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateMessageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PrivateMessageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PrivateMessageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PrivateMessageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PrivateMessageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PrivateMessageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PrivateMessageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PrivateMessageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PrivateMessageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRead

`func (o *PrivateMessageDto) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PrivateMessageDto) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PrivateMessageDto) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *PrivateMessageDto) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetTitle

`func (o *PrivateMessageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrivateMessageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrivateMessageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PrivateMessageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PrivateMessageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PrivateMessageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *PrivateMessageDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrivateMessageDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrivateMessageDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PrivateMessageDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PrivateMessageDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PrivateMessageDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetConversationId

`func (o *PrivateMessageDto) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *PrivateMessageDto) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *PrivateMessageDto) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *PrivateMessageDto) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *PrivateMessageDto) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *PrivateMessageDto) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetSenderSocialProfileId

`func (o *PrivateMessageDto) GetSenderSocialProfileId() string`

GetSenderSocialProfileId returns the SenderSocialProfileId field if non-nil, zero value otherwise.

### GetSenderSocialProfileIdOk

`func (o *PrivateMessageDto) GetSenderSocialProfileIdOk() (*string, bool)`

GetSenderSocialProfileIdOk returns a tuple with the SenderSocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSocialProfileId

`func (o *PrivateMessageDto) SetSenderSocialProfileId(v string)`

SetSenderSocialProfileId sets SenderSocialProfileId field to given value.

### HasSenderSocialProfileId

`func (o *PrivateMessageDto) HasSenderSocialProfileId() bool`

HasSenderSocialProfileId returns a boolean if a field has been set.

### SetSenderSocialProfileIdNil

`func (o *PrivateMessageDto) SetSenderSocialProfileIdNil(b bool)`

 SetSenderSocialProfileIdNil sets the value for SenderSocialProfileId to be an explicit nil

### UnsetSenderSocialProfileId
`func (o *PrivateMessageDto) UnsetSenderSocialProfileId()`

UnsetSenderSocialProfileId ensures that no value is present for SenderSocialProfileId, not even an explicit nil
### GetReceiverSocialProfileID

`func (o *PrivateMessageDto) GetReceiverSocialProfileID() string`

GetReceiverSocialProfileID returns the ReceiverSocialProfileID field if non-nil, zero value otherwise.

### GetReceiverSocialProfileIDOk

`func (o *PrivateMessageDto) GetReceiverSocialProfileIDOk() (*string, bool)`

GetReceiverSocialProfileIDOk returns a tuple with the ReceiverSocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverSocialProfileID

`func (o *PrivateMessageDto) SetReceiverSocialProfileID(v string)`

SetReceiverSocialProfileID sets ReceiverSocialProfileID field to given value.

### HasReceiverSocialProfileID

`func (o *PrivateMessageDto) HasReceiverSocialProfileID() bool`

HasReceiverSocialProfileID returns a boolean if a field has been set.

### SetReceiverSocialProfileIDNil

`func (o *PrivateMessageDto) SetReceiverSocialProfileIDNil(b bool)`

 SetReceiverSocialProfileIDNil sets the value for ReceiverSocialProfileID to be an explicit nil

### UnsetReceiverSocialProfileID
`func (o *PrivateMessageDto) UnsetReceiverSocialProfileID()`

UnsetReceiverSocialProfileID ensures that no value is present for ReceiverSocialProfileID, not even an explicit nil
### GetSentTimestamp

`func (o *PrivateMessageDto) GetSentTimestamp() time.Time`

GetSentTimestamp returns the SentTimestamp field if non-nil, zero value otherwise.

### GetSentTimestampOk

`func (o *PrivateMessageDto) GetSentTimestampOk() (*time.Time, bool)`

GetSentTimestampOk returns a tuple with the SentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimestamp

`func (o *PrivateMessageDto) SetSentTimestamp(v time.Time)`

SetSentTimestamp sets SentTimestamp field to given value.

### HasSentTimestamp

`func (o *PrivateMessageDto) HasSentTimestamp() bool`

HasSentTimestamp returns a boolean if a field has been set.

### GetReadTimestamp

`func (o *PrivateMessageDto) GetReadTimestamp() time.Time`

GetReadTimestamp returns the ReadTimestamp field if non-nil, zero value otherwise.

### GetReadTimestampOk

`func (o *PrivateMessageDto) GetReadTimestampOk() (*time.Time, bool)`

GetReadTimestampOk returns a tuple with the ReadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimestamp

`func (o *PrivateMessageDto) SetReadTimestamp(v time.Time)`

SetReadTimestamp sets ReadTimestamp field to given value.

### HasReadTimestamp

`func (o *PrivateMessageDto) HasReadTimestamp() bool`

HasReadTimestamp returns a boolean if a field has been set.

### GetReceivedTimestamp

`func (o *PrivateMessageDto) GetReceivedTimestamp() time.Time`

GetReceivedTimestamp returns the ReceivedTimestamp field if non-nil, zero value otherwise.

### GetReceivedTimestampOk

`func (o *PrivateMessageDto) GetReceivedTimestampOk() (*time.Time, bool)`

GetReceivedTimestampOk returns a tuple with the ReceivedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimestamp

`func (o *PrivateMessageDto) SetReceivedTimestamp(v time.Time)`

SetReceivedTimestamp sets ReceivedTimestamp field to given value.

### HasReceivedTimestamp

`func (o *PrivateMessageDto) HasReceivedTimestamp() bool`

HasReceivedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


