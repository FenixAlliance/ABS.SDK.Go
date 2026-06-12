# ItemReviewRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ReviewScore** | Pointer to **float64** |  | [optional] 
**ReviewMessage** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemReviewRecordCreateDto

`func NewItemReviewRecordCreateDto() *ItemReviewRecordCreateDto`

NewItemReviewRecordCreateDto instantiates a new ItemReviewRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReviewRecordCreateDtoWithDefaults

`func NewItemReviewRecordCreateDtoWithDefaults() *ItemReviewRecordCreateDto`

NewItemReviewRecordCreateDtoWithDefaults instantiates a new ItemReviewRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemReviewRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemReviewRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemReviewRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemReviewRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemReviewRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemReviewRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemReviewRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemReviewRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReviewScore

`func (o *ItemReviewRecordCreateDto) GetReviewScore() float64`

GetReviewScore returns the ReviewScore field if non-nil, zero value otherwise.

### GetReviewScoreOk

`func (o *ItemReviewRecordCreateDto) GetReviewScoreOk() (*float64, bool)`

GetReviewScoreOk returns a tuple with the ReviewScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewScore

`func (o *ItemReviewRecordCreateDto) SetReviewScore(v float64)`

SetReviewScore sets ReviewScore field to given value.

### HasReviewScore

`func (o *ItemReviewRecordCreateDto) HasReviewScore() bool`

HasReviewScore returns a boolean if a field has been set.

### GetReviewMessage

`func (o *ItemReviewRecordCreateDto) GetReviewMessage() string`

GetReviewMessage returns the ReviewMessage field if non-nil, zero value otherwise.

### GetReviewMessageOk

`func (o *ItemReviewRecordCreateDto) GetReviewMessageOk() (*string, bool)`

GetReviewMessageOk returns a tuple with the ReviewMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewMessage

`func (o *ItemReviewRecordCreateDto) SetReviewMessage(v string)`

SetReviewMessage sets ReviewMessage field to given value.

### HasReviewMessage

`func (o *ItemReviewRecordCreateDto) HasReviewMessage() bool`

HasReviewMessage returns a boolean if a field has been set.

### SetReviewMessageNil

`func (o *ItemReviewRecordCreateDto) SetReviewMessageNil(b bool)`

 SetReviewMessageNil sets the value for ReviewMessage to be an explicit nil

### UnsetReviewMessage
`func (o *ItemReviewRecordCreateDto) UnsetReviewMessage()`

UnsetReviewMessage ensures that no value is present for ReviewMessage, not even an explicit nil
### GetSocialProfileId

`func (o *ItemReviewRecordCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ItemReviewRecordCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ItemReviewRecordCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ItemReviewRecordCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ItemReviewRecordCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ItemReviewRecordCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


