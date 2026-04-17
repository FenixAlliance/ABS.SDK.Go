# ItemCategoryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**ParentItemCategoryID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemCategoryCreateDto

`func NewItemCategoryCreateDto(title string, ) *ItemCategoryCreateDto`

NewItemCategoryCreateDto instantiates a new ItemCategoryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCategoryCreateDtoWithDefaults

`func NewItemCategoryCreateDtoWithDefaults() *ItemCategoryCreateDto`

NewItemCategoryCreateDtoWithDefaults instantiates a new ItemCategoryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemCategoryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemCategoryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemCategoryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemCategoryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemCategoryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemCategoryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemCategoryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemCategoryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ItemCategoryCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemCategoryCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemCategoryCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemCategoryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemCategoryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemCategoryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemCategoryCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemCategoryCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemCategoryCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageURL

`func (o *ItemCategoryCreateDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *ItemCategoryCreateDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *ItemCategoryCreateDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *ItemCategoryCreateDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *ItemCategoryCreateDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *ItemCategoryCreateDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetParentItemCategoryID

`func (o *ItemCategoryCreateDto) GetParentItemCategoryID() string`

GetParentItemCategoryID returns the ParentItemCategoryID field if non-nil, zero value otherwise.

### GetParentItemCategoryIDOk

`func (o *ItemCategoryCreateDto) GetParentItemCategoryIDOk() (*string, bool)`

GetParentItemCategoryIDOk returns a tuple with the ParentItemCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemCategoryID

`func (o *ItemCategoryCreateDto) SetParentItemCategoryID(v string)`

SetParentItemCategoryID sets ParentItemCategoryID field to given value.

### HasParentItemCategoryID

`func (o *ItemCategoryCreateDto) HasParentItemCategoryID() bool`

HasParentItemCategoryID returns a boolean if a field has been set.

### SetParentItemCategoryIDNil

`func (o *ItemCategoryCreateDto) SetParentItemCategoryIDNil(b bool)`

 SetParentItemCategoryIDNil sets the value for ParentItemCategoryID to be an explicit nil

### UnsetParentItemCategoryID
`func (o *ItemCategoryCreateDto) UnsetParentItemCategoryID()`

UnsetParentItemCategoryID ensures that no value is present for ParentItemCategoryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


