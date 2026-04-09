# ItemReviewCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemID** | Pointer to **NullableString** |  | [optional] 
**ReviewScore** | Pointer to **float64** |  | [optional] 
**ReviewMessage** | Pointer to **NullableString** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 

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

### GetItemID

`func (o *ItemReviewCreateDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemReviewCreateDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemReviewCreateDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemReviewCreateDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemReviewCreateDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemReviewCreateDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil
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
### GetSocialProfileID

`func (o *ItemReviewCreateDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemReviewCreateDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemReviewCreateDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemReviewCreateDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemReviewCreateDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemReviewCreateDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


