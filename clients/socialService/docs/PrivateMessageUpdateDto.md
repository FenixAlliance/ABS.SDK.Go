# PrivateMessageUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateMessageUpdateDto

`func NewPrivateMessageUpdateDto() *PrivateMessageUpdateDto`

NewPrivateMessageUpdateDto instantiates a new PrivateMessageUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateMessageUpdateDtoWithDefaults

`func NewPrivateMessageUpdateDtoWithDefaults() *PrivateMessageUpdateDto`

NewPrivateMessageUpdateDtoWithDefaults instantiates a new PrivateMessageUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PrivateMessageUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrivateMessageUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrivateMessageUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PrivateMessageUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PrivateMessageUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PrivateMessageUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *PrivateMessageUpdateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrivateMessageUpdateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrivateMessageUpdateDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PrivateMessageUpdateDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PrivateMessageUpdateDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PrivateMessageUpdateDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


