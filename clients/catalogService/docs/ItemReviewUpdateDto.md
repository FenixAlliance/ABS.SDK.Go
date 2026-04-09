# ItemReviewUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewScore** | Pointer to **float64** |  | [optional] 
**ReviewMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemReviewUpdateDto

`func NewItemReviewUpdateDto() *ItemReviewUpdateDto`

NewItemReviewUpdateDto instantiates a new ItemReviewUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReviewUpdateDtoWithDefaults

`func NewItemReviewUpdateDtoWithDefaults() *ItemReviewUpdateDto`

NewItemReviewUpdateDtoWithDefaults instantiates a new ItemReviewUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewScore

`func (o *ItemReviewUpdateDto) GetReviewScore() float64`

GetReviewScore returns the ReviewScore field if non-nil, zero value otherwise.

### GetReviewScoreOk

`func (o *ItemReviewUpdateDto) GetReviewScoreOk() (*float64, bool)`

GetReviewScoreOk returns a tuple with the ReviewScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewScore

`func (o *ItemReviewUpdateDto) SetReviewScore(v float64)`

SetReviewScore sets ReviewScore field to given value.

### HasReviewScore

`func (o *ItemReviewUpdateDto) HasReviewScore() bool`

HasReviewScore returns a boolean if a field has been set.

### GetReviewMessage

`func (o *ItemReviewUpdateDto) GetReviewMessage() string`

GetReviewMessage returns the ReviewMessage field if non-nil, zero value otherwise.

### GetReviewMessageOk

`func (o *ItemReviewUpdateDto) GetReviewMessageOk() (*string, bool)`

GetReviewMessageOk returns a tuple with the ReviewMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewMessage

`func (o *ItemReviewUpdateDto) SetReviewMessage(v string)`

SetReviewMessage sets ReviewMessage field to given value.

### HasReviewMessage

`func (o *ItemReviewUpdateDto) HasReviewMessage() bool`

HasReviewMessage returns a boolean if a field has been set.

### SetReviewMessageNil

`func (o *ItemReviewUpdateDto) SetReviewMessageNil(b bool)`

 SetReviewMessageNil sets the value for ReviewMessage to be an explicit nil

### UnsetReviewMessage
`func (o *ItemReviewUpdateDto) UnsetReviewMessage()`

UnsetReviewMessage ensures that no value is present for ReviewMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


