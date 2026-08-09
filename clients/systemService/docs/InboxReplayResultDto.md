# InboxReplayResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewInboxMessageId** | Pointer to **NullableString** |  | [optional] 
**RootInboxMessageId** | Pointer to **NullableString** |  | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 

## Methods

### NewInboxReplayResultDto

`func NewInboxReplayResultDto() *InboxReplayResultDto`

NewInboxReplayResultDto instantiates a new InboxReplayResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxReplayResultDtoWithDefaults

`func NewInboxReplayResultDtoWithDefaults() *InboxReplayResultDto`

NewInboxReplayResultDtoWithDefaults instantiates a new InboxReplayResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewInboxMessageId

`func (o *InboxReplayResultDto) GetNewInboxMessageId() string`

GetNewInboxMessageId returns the NewInboxMessageId field if non-nil, zero value otherwise.

### GetNewInboxMessageIdOk

`func (o *InboxReplayResultDto) GetNewInboxMessageIdOk() (*string, bool)`

GetNewInboxMessageIdOk returns a tuple with the NewInboxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewInboxMessageId

`func (o *InboxReplayResultDto) SetNewInboxMessageId(v string)`

SetNewInboxMessageId sets NewInboxMessageId field to given value.

### HasNewInboxMessageId

`func (o *InboxReplayResultDto) HasNewInboxMessageId() bool`

HasNewInboxMessageId returns a boolean if a field has been set.

### SetNewInboxMessageIdNil

`func (o *InboxReplayResultDto) SetNewInboxMessageIdNil(b bool)`

 SetNewInboxMessageIdNil sets the value for NewInboxMessageId to be an explicit nil

### UnsetNewInboxMessageId
`func (o *InboxReplayResultDto) UnsetNewInboxMessageId()`

UnsetNewInboxMessageId ensures that no value is present for NewInboxMessageId, not even an explicit nil
### GetRootInboxMessageId

`func (o *InboxReplayResultDto) GetRootInboxMessageId() string`

GetRootInboxMessageId returns the RootInboxMessageId field if non-nil, zero value otherwise.

### GetRootInboxMessageIdOk

`func (o *InboxReplayResultDto) GetRootInboxMessageIdOk() (*string, bool)`

GetRootInboxMessageIdOk returns a tuple with the RootInboxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootInboxMessageId

`func (o *InboxReplayResultDto) SetRootInboxMessageId(v string)`

SetRootInboxMessageId sets RootInboxMessageId field to given value.

### HasRootInboxMessageId

`func (o *InboxReplayResultDto) HasRootInboxMessageId() bool`

HasRootInboxMessageId returns a boolean if a field has been set.

### SetRootInboxMessageIdNil

`func (o *InboxReplayResultDto) SetRootInboxMessageIdNil(b bool)`

 SetRootInboxMessageIdNil sets the value for RootInboxMessageId to be an explicit nil

### UnsetRootInboxMessageId
`func (o *InboxReplayResultDto) UnsetRootInboxMessageId()`

UnsetRootInboxMessageId ensures that no value is present for RootInboxMessageId, not even an explicit nil
### GetGeneration

`func (o *InboxReplayResultDto) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *InboxReplayResultDto) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *InboxReplayResultDto) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *InboxReplayResultDto) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


