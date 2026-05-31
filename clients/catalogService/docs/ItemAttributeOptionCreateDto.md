# ItemAttributeOptionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ItemAttributeId** | **string** |  | 

## Methods

### NewItemAttributeOptionCreateDto

`func NewItemAttributeOptionCreateDto(name string, itemAttributeId string, ) *ItemAttributeOptionCreateDto`

NewItemAttributeOptionCreateDto instantiates a new ItemAttributeOptionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAttributeOptionCreateDtoWithDefaults

`func NewItemAttributeOptionCreateDtoWithDefaults() *ItemAttributeOptionCreateDto`

NewItemAttributeOptionCreateDtoWithDefaults instantiates a new ItemAttributeOptionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemAttributeOptionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemAttributeOptionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemAttributeOptionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemAttributeOptionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemAttributeOptionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemAttributeOptionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemAttributeOptionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemAttributeOptionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ItemAttributeOptionCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemAttributeOptionCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemAttributeOptionCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ItemAttributeOptionCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemAttributeOptionCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemAttributeOptionCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemAttributeOptionCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemAttributeOptionCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemAttributeOptionCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetItemAttributeId

`func (o *ItemAttributeOptionCreateDto) GetItemAttributeId() string`

GetItemAttributeId returns the ItemAttributeId field if non-nil, zero value otherwise.

### GetItemAttributeIdOk

`func (o *ItemAttributeOptionCreateDto) GetItemAttributeIdOk() (*string, bool)`

GetItemAttributeIdOk returns a tuple with the ItemAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAttributeId

`func (o *ItemAttributeOptionCreateDto) SetItemAttributeId(v string)`

SetItemAttributeId sets ItemAttributeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


