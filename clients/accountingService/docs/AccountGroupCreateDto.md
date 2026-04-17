# AccountGroupCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentAccountGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccountGroupCreateDto

`func NewAccountGroupCreateDto() *AccountGroupCreateDto`

NewAccountGroupCreateDto instantiates a new AccountGroupCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupCreateDtoWithDefaults

`func NewAccountGroupCreateDtoWithDefaults() *AccountGroupCreateDto`

NewAccountGroupCreateDtoWithDefaults instantiates a new AccountGroupCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountGroupCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountGroupCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountGroupCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountGroupCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountGroupCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *AccountGroupCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountGroupCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountGroupCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AccountGroupCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AccountGroupCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AccountGroupCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AccountGroupCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountGroupCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountGroupCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountGroupCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountGroupCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountGroupCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentAccountGroupId

`func (o *AccountGroupCreateDto) GetParentAccountGroupId() string`

GetParentAccountGroupId returns the ParentAccountGroupId field if non-nil, zero value otherwise.

### GetParentAccountGroupIdOk

`func (o *AccountGroupCreateDto) GetParentAccountGroupIdOk() (*string, bool)`

GetParentAccountGroupIdOk returns a tuple with the ParentAccountGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountGroupId

`func (o *AccountGroupCreateDto) SetParentAccountGroupId(v string)`

SetParentAccountGroupId sets ParentAccountGroupId field to given value.

### HasParentAccountGroupId

`func (o *AccountGroupCreateDto) HasParentAccountGroupId() bool`

HasParentAccountGroupId returns a boolean if a field has been set.

### SetParentAccountGroupIdNil

`func (o *AccountGroupCreateDto) SetParentAccountGroupIdNil(b bool)`

 SetParentAccountGroupIdNil sets the value for ParentAccountGroupId to be an explicit nil

### UnsetParentAccountGroupId
`func (o *AccountGroupCreateDto) UnsetParentAccountGroupId()`

UnsetParentAccountGroupId ensures that no value is present for ParentAccountGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


