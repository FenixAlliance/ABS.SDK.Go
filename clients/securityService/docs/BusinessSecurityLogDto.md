# BusinessSecurityLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**SecurityEvent** | Pointer to **NullableString** |  | [optional] 
**RequiresAttention** | Pointer to **bool** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBusinessSecurityLogDto

`func NewBusinessSecurityLogDto() *BusinessSecurityLogDto`

NewBusinessSecurityLogDto instantiates a new BusinessSecurityLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessSecurityLogDtoWithDefaults

`func NewBusinessSecurityLogDtoWithDefaults() *BusinessSecurityLogDto`

NewBusinessSecurityLogDtoWithDefaults instantiates a new BusinessSecurityLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessSecurityLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessSecurityLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessSecurityLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessSecurityLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BusinessSecurityLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BusinessSecurityLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BusinessSecurityLogDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessSecurityLogDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessSecurityLogDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessSecurityLogDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *BusinessSecurityLogDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BusinessSecurityLogDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BusinessSecurityLogDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BusinessSecurityLogDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BusinessSecurityLogDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BusinessSecurityLogDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *BusinessSecurityLogDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BusinessSecurityLogDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BusinessSecurityLogDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BusinessSecurityLogDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BusinessSecurityLogDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BusinessSecurityLogDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSecurityEvent

`func (o *BusinessSecurityLogDto) GetSecurityEvent() string`

GetSecurityEvent returns the SecurityEvent field if non-nil, zero value otherwise.

### GetSecurityEventOk

`func (o *BusinessSecurityLogDto) GetSecurityEventOk() (*string, bool)`

GetSecurityEventOk returns a tuple with the SecurityEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEvent

`func (o *BusinessSecurityLogDto) SetSecurityEvent(v string)`

SetSecurityEvent sets SecurityEvent field to given value.

### HasSecurityEvent

`func (o *BusinessSecurityLogDto) HasSecurityEvent() bool`

HasSecurityEvent returns a boolean if a field has been set.

### SetSecurityEventNil

`func (o *BusinessSecurityLogDto) SetSecurityEventNil(b bool)`

 SetSecurityEventNil sets the value for SecurityEvent to be an explicit nil

### UnsetSecurityEvent
`func (o *BusinessSecurityLogDto) UnsetSecurityEvent()`

UnsetSecurityEvent ensures that no value is present for SecurityEvent, not even an explicit nil
### GetRequiresAttention

`func (o *BusinessSecurityLogDto) GetRequiresAttention() bool`

GetRequiresAttention returns the RequiresAttention field if non-nil, zero value otherwise.

### GetRequiresAttentionOk

`func (o *BusinessSecurityLogDto) GetRequiresAttentionOk() (*bool, bool)`

GetRequiresAttentionOk returns a tuple with the RequiresAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAttention

`func (o *BusinessSecurityLogDto) SetRequiresAttention(v bool)`

SetRequiresAttention sets RequiresAttention field to given value.

### HasRequiresAttention

`func (o *BusinessSecurityLogDto) HasRequiresAttention() bool`

HasRequiresAttention returns a boolean if a field has been set.

### GetBusinessID

`func (o *BusinessSecurityLogDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *BusinessSecurityLogDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *BusinessSecurityLogDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *BusinessSecurityLogDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *BusinessSecurityLogDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *BusinessSecurityLogDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


