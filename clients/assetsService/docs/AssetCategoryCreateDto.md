# AssetCategoryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetCategoryCreateDto

`func NewAssetCategoryCreateDto() *AssetCategoryCreateDto`

NewAssetCategoryCreateDto instantiates a new AssetCategoryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCategoryCreateDtoWithDefaults

`func NewAssetCategoryCreateDtoWithDefaults() *AssetCategoryCreateDto`

NewAssetCategoryCreateDtoWithDefaults instantiates a new AssetCategoryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetCategoryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetCategoryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetCategoryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetCategoryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetCategoryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetCategoryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetCategoryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetCategoryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetCategoryCreateDto) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetCategoryCreateDto) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetCategoryCreateDto) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetCategoryCreateDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### SetBusinessIdNil

`func (o *AssetCategoryCreateDto) SetBusinessIdNil(b bool)`

 SetBusinessIdNil sets the value for BusinessId to be an explicit nil

### UnsetBusinessId
`func (o *AssetCategoryCreateDto) UnsetBusinessId()`

UnsetBusinessId ensures that no value is present for BusinessId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *AssetCategoryCreateDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetCategoryCreateDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetCategoryCreateDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetCategoryCreateDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *AssetCategoryCreateDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *AssetCategoryCreateDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetName

`func (o *AssetCategoryCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetCategoryCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetCategoryCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetCategoryCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AssetCategoryCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetCategoryCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AssetCategoryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetCategoryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetCategoryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetCategoryCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AssetCategoryCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetCategoryCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


