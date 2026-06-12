# ItemReviewCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ReviewScore** | Pointer to **float64** |  | [optional] 
**ReviewMessage** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemReviewCreateDto

`func NewItemReviewCreateDto() *ItemReviewCreateDto`

NewItemReviewCreateDto instantiates a new ItemReviewCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReviewCreateDtoWithDefaults

`func NewItemReviewCreateDtoWithDefaults() *ItemReviewCreateDto`

NewItemReviewCreateDtoWithDefaults instantiates a new ItemReviewCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemReviewCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemReviewCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemReviewCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemReviewCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemReviewCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemReviewCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemReviewCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemReviewCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemId

`func (o *ItemReviewCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemReviewCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemReviewCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemReviewCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemReviewCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemReviewCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetReviewScore

`func (o *ItemReviewCreateDto) GetReviewScore() float64`

GetReviewScore returns the ReviewScore field if non-nil, zero value otherwise.

### GetReviewScoreOk

`func (o *ItemReviewCreateDto) GetReviewScoreOk() (*float64, bool)`

GetReviewScoreOk returns a tuple with the ReviewScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewScore

`func (o *ItemReviewCreateDto) SetReviewScore(v float64)`

SetReviewScore sets ReviewScore field to given value.

### HasReviewScore

`func (o *ItemReviewCreateDto) HasReviewScore() bool`

HasReviewScore returns a boolean if a field has been set.

### GetReviewMessage

`func (o *ItemReviewCreateDto) GetReviewMessage() string`

GetReviewMessage returns the ReviewMessage field if non-nil, zero value otherwise.

### GetReviewMessageOk

`func (o *ItemReviewCreateDto) GetReviewMessageOk() (*string, bool)`

GetReviewMessageOk returns a tuple with the ReviewMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewMessage

`func (o *ItemReviewCreateDto) SetReviewMessage(v string)`

SetReviewMessage sets ReviewMessage field to given value.

### HasReviewMessage

`func (o *ItemReviewCreateDto) HasReviewMessage() bool`

HasReviewMessage returns a boolean if a field has been set.

### SetReviewMessageNil

`func (o *ItemReviewCreateDto) SetReviewMessageNil(b bool)`

 SetReviewMessageNil sets the value for ReviewMessage to be an explicit nil

### UnsetReviewMessage
`func (o *ItemReviewCreateDto) UnsetReviewMessage()`

UnsetReviewMessage ensures that no value is present for ReviewMessage, not even an explicit nil
### GetSocialProfileId

`func (o *ItemReviewCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ItemReviewCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ItemReviewCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ItemReviewCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ItemReviewCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ItemReviewCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


