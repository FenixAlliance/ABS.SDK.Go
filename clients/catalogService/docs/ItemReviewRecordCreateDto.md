# ItemReviewRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ReviewScore** | Pointer to **float64** |  | [optional] 
**ReviewMessage** | Pointer to **NullableString** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 

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
### GetSocialProfileID

`func (o *ItemReviewRecordCreateDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemReviewRecordCreateDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemReviewRecordCreateDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemReviewRecordCreateDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemReviewRecordCreateDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemReviewRecordCreateDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


