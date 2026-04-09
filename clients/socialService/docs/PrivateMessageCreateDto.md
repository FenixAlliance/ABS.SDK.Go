# PrivateMessageCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ConversationId** | Pointer to **NullableString** |  | [optional] 
**SenderSocialProfileId** | Pointer to **NullableString** |  | [optional] 
**ReceiverSocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateMessageCreateDto

`func NewPrivateMessageCreateDto() *PrivateMessageCreateDto`

NewPrivateMessageCreateDto instantiates a new PrivateMessageCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateMessageCreateDtoWithDefaults

`func NewPrivateMessageCreateDtoWithDefaults() *PrivateMessageCreateDto`

NewPrivateMessageCreateDtoWithDefaults instantiates a new PrivateMessageCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateMessageCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateMessageCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateMessageCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateMessageCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PrivateMessageCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PrivateMessageCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PrivateMessageCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PrivateMessageCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *PrivateMessageCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrivateMessageCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrivateMessageCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PrivateMessageCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PrivateMessageCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PrivateMessageCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *PrivateMessageCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrivateMessageCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrivateMessageCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PrivateMessageCreateDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PrivateMessageCreateDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PrivateMessageCreateDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetConversationId

`func (o *PrivateMessageCreateDto) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *PrivateMessageCreateDto) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *PrivateMessageCreateDto) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *PrivateMessageCreateDto) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *PrivateMessageCreateDto) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *PrivateMessageCreateDto) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetSenderSocialProfileId

`func (o *PrivateMessageCreateDto) GetSenderSocialProfileId() string`

GetSenderSocialProfileId returns the SenderSocialProfileId field if non-nil, zero value otherwise.

### GetSenderSocialProfileIdOk

`func (o *PrivateMessageCreateDto) GetSenderSocialProfileIdOk() (*string, bool)`

GetSenderSocialProfileIdOk returns a tuple with the SenderSocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSocialProfileId

`func (o *PrivateMessageCreateDto) SetSenderSocialProfileId(v string)`

SetSenderSocialProfileId sets SenderSocialProfileId field to given value.

### HasSenderSocialProfileId

`func (o *PrivateMessageCreateDto) HasSenderSocialProfileId() bool`

HasSenderSocialProfileId returns a boolean if a field has been set.

### SetSenderSocialProfileIdNil

`func (o *PrivateMessageCreateDto) SetSenderSocialProfileIdNil(b bool)`

 SetSenderSocialProfileIdNil sets the value for SenderSocialProfileId to be an explicit nil

### UnsetSenderSocialProfileId
`func (o *PrivateMessageCreateDto) UnsetSenderSocialProfileId()`

UnsetSenderSocialProfileId ensures that no value is present for SenderSocialProfileId, not even an explicit nil
### GetReceiverSocialProfileId

`func (o *PrivateMessageCreateDto) GetReceiverSocialProfileId() string`

GetReceiverSocialProfileId returns the ReceiverSocialProfileId field if non-nil, zero value otherwise.

### GetReceiverSocialProfileIdOk

`func (o *PrivateMessageCreateDto) GetReceiverSocialProfileIdOk() (*string, bool)`

GetReceiverSocialProfileIdOk returns a tuple with the ReceiverSocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverSocialProfileId

`func (o *PrivateMessageCreateDto) SetReceiverSocialProfileId(v string)`

SetReceiverSocialProfileId sets ReceiverSocialProfileId field to given value.

### HasReceiverSocialProfileId

`func (o *PrivateMessageCreateDto) HasReceiverSocialProfileId() bool`

HasReceiverSocialProfileId returns a boolean if a field has been set.

### SetReceiverSocialProfileIdNil

`func (o *PrivateMessageCreateDto) SetReceiverSocialProfileIdNil(b bool)`

 SetReceiverSocialProfileIdNil sets the value for ReceiverSocialProfileId to be an explicit nil

### UnsetReceiverSocialProfileId
`func (o *PrivateMessageCreateDto) UnsetReceiverSocialProfileId()`

UnsetReceiverSocialProfileId ensures that no value is present for ReceiverSocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


