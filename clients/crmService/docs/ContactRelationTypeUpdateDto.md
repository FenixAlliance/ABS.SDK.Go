# ContactRelationTypeUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**BackName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactRelationTypeUpdateDto

`func NewContactRelationTypeUpdateDto(name string, ) *ContactRelationTypeUpdateDto`

NewContactRelationTypeUpdateDto instantiates a new ContactRelationTypeUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRelationTypeUpdateDtoWithDefaults

`func NewContactRelationTypeUpdateDtoWithDefaults() *ContactRelationTypeUpdateDto`

NewContactRelationTypeUpdateDtoWithDefaults instantiates a new ContactRelationTypeUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactRelationTypeUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactRelationTypeUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactRelationTypeUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetBackName

`func (o *ContactRelationTypeUpdateDto) GetBackName() string`

GetBackName returns the BackName field if non-nil, zero value otherwise.

### GetBackNameOk

`func (o *ContactRelationTypeUpdateDto) GetBackNameOk() (*string, bool)`

GetBackNameOk returns a tuple with the BackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackName

`func (o *ContactRelationTypeUpdateDto) SetBackName(v string)`

SetBackName sets BackName field to given value.

### HasBackName

`func (o *ContactRelationTypeUpdateDto) HasBackName() bool`

HasBackName returns a boolean if a field has been set.

### SetBackNameNil

`func (o *ContactRelationTypeUpdateDto) SetBackNameNil(b bool)`

 SetBackNameNil sets the value for BackName to be an explicit nil

### UnsetBackName
`func (o *ContactRelationTypeUpdateDto) UnsetBackName()`

UnsetBackName ensures that no value is present for BackName, not even an explicit nil
### GetDescription

`func (o *ContactRelationTypeUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactRelationTypeUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactRelationTypeUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactRelationTypeUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContactRelationTypeUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContactRelationTypeUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


