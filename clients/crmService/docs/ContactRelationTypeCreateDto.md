# ContactRelationTypeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**BackName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactRelationTypeCreateDto

`func NewContactRelationTypeCreateDto(name string, ) *ContactRelationTypeCreateDto`

NewContactRelationTypeCreateDto instantiates a new ContactRelationTypeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRelationTypeCreateDtoWithDefaults

`func NewContactRelationTypeCreateDtoWithDefaults() *ContactRelationTypeCreateDto`

NewContactRelationTypeCreateDtoWithDefaults instantiates a new ContactRelationTypeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactRelationTypeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactRelationTypeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactRelationTypeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactRelationTypeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContactRelationTypeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactRelationTypeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactRelationTypeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactRelationTypeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ContactRelationTypeCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactRelationTypeCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactRelationTypeCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetBackName

`func (o *ContactRelationTypeCreateDto) GetBackName() string`

GetBackName returns the BackName field if non-nil, zero value otherwise.

### GetBackNameOk

`func (o *ContactRelationTypeCreateDto) GetBackNameOk() (*string, bool)`

GetBackNameOk returns a tuple with the BackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackName

`func (o *ContactRelationTypeCreateDto) SetBackName(v string)`

SetBackName sets BackName field to given value.

### HasBackName

`func (o *ContactRelationTypeCreateDto) HasBackName() bool`

HasBackName returns a boolean if a field has been set.

### SetBackNameNil

`func (o *ContactRelationTypeCreateDto) SetBackNameNil(b bool)`

 SetBackNameNil sets the value for BackName to be an explicit nil

### UnsetBackName
`func (o *ContactRelationTypeCreateDto) UnsetBackName()`

UnsetBackName ensures that no value is present for BackName, not even an explicit nil
### GetDescription

`func (o *ContactRelationTypeCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactRelationTypeCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactRelationTypeCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactRelationTypeCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContactRelationTypeCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContactRelationTypeCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


