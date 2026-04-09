# ItemTypeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**PluralTitle** | Pointer to **NullableString** |  | [optional] 
**SingularTitle** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**GoogleCategoryTaxonomy** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | **string** |  | 
**ItemCategoryID** | **string** |  | 
**ItemGoogleCategoryID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemTypeCreateDto

`func NewItemTypeCreateDto(businessID string, itemCategoryID string, ) *ItemTypeCreateDto`

NewItemTypeCreateDto instantiates a new ItemTypeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTypeCreateDtoWithDefaults

`func NewItemTypeCreateDtoWithDefaults() *ItemTypeCreateDto`

NewItemTypeCreateDtoWithDefaults instantiates a new ItemTypeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemTypeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemTypeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemTypeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemTypeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemTypeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemTypeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemTypeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemTypeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPluralTitle

`func (o *ItemTypeCreateDto) GetPluralTitle() string`

GetPluralTitle returns the PluralTitle field if non-nil, zero value otherwise.

### GetPluralTitleOk

`func (o *ItemTypeCreateDto) GetPluralTitleOk() (*string, bool)`

GetPluralTitleOk returns a tuple with the PluralTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralTitle

`func (o *ItemTypeCreateDto) SetPluralTitle(v string)`

SetPluralTitle sets PluralTitle field to given value.

### HasPluralTitle

`func (o *ItemTypeCreateDto) HasPluralTitle() bool`

HasPluralTitle returns a boolean if a field has been set.

### SetPluralTitleNil

`func (o *ItemTypeCreateDto) SetPluralTitleNil(b bool)`

 SetPluralTitleNil sets the value for PluralTitle to be an explicit nil

### UnsetPluralTitle
`func (o *ItemTypeCreateDto) UnsetPluralTitle()`

UnsetPluralTitle ensures that no value is present for PluralTitle, not even an explicit nil
### GetSingularTitle

`func (o *ItemTypeCreateDto) GetSingularTitle() string`

GetSingularTitle returns the SingularTitle field if non-nil, zero value otherwise.

### GetSingularTitleOk

`func (o *ItemTypeCreateDto) GetSingularTitleOk() (*string, bool)`

GetSingularTitleOk returns a tuple with the SingularTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularTitle

`func (o *ItemTypeCreateDto) SetSingularTitle(v string)`

SetSingularTitle sets SingularTitle field to given value.

### HasSingularTitle

`func (o *ItemTypeCreateDto) HasSingularTitle() bool`

HasSingularTitle returns a boolean if a field has been set.

### SetSingularTitleNil

`func (o *ItemTypeCreateDto) SetSingularTitleNil(b bool)`

 SetSingularTitleNil sets the value for SingularTitle to be an explicit nil

### UnsetSingularTitle
`func (o *ItemTypeCreateDto) UnsetSingularTitle()`

UnsetSingularTitle ensures that no value is present for SingularTitle, not even an explicit nil
### GetDescription

`func (o *ItemTypeCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemTypeCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemTypeCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemTypeCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemTypeCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemTypeCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageURL

`func (o *ItemTypeCreateDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *ItemTypeCreateDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *ItemTypeCreateDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *ItemTypeCreateDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *ItemTypeCreateDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *ItemTypeCreateDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetGoogleCategoryTaxonomy

`func (o *ItemTypeCreateDto) GetGoogleCategoryTaxonomy() string`

GetGoogleCategoryTaxonomy returns the GoogleCategoryTaxonomy field if non-nil, zero value otherwise.

### GetGoogleCategoryTaxonomyOk

`func (o *ItemTypeCreateDto) GetGoogleCategoryTaxonomyOk() (*string, bool)`

GetGoogleCategoryTaxonomyOk returns a tuple with the GoogleCategoryTaxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCategoryTaxonomy

`func (o *ItemTypeCreateDto) SetGoogleCategoryTaxonomy(v string)`

SetGoogleCategoryTaxonomy sets GoogleCategoryTaxonomy field to given value.

### HasGoogleCategoryTaxonomy

`func (o *ItemTypeCreateDto) HasGoogleCategoryTaxonomy() bool`

HasGoogleCategoryTaxonomy returns a boolean if a field has been set.

### SetGoogleCategoryTaxonomyNil

`func (o *ItemTypeCreateDto) SetGoogleCategoryTaxonomyNil(b bool)`

 SetGoogleCategoryTaxonomyNil sets the value for GoogleCategoryTaxonomy to be an explicit nil

### UnsetGoogleCategoryTaxonomy
`func (o *ItemTypeCreateDto) UnsetGoogleCategoryTaxonomy()`

UnsetGoogleCategoryTaxonomy ensures that no value is present for GoogleCategoryTaxonomy, not even an explicit nil
### GetBusinessID

`func (o *ItemTypeCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemTypeCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemTypeCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetItemCategoryID

`func (o *ItemTypeCreateDto) GetItemCategoryID() string`

GetItemCategoryID returns the ItemCategoryID field if non-nil, zero value otherwise.

### GetItemCategoryIDOk

`func (o *ItemTypeCreateDto) GetItemCategoryIDOk() (*string, bool)`

GetItemCategoryIDOk returns a tuple with the ItemCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCategoryID

`func (o *ItemTypeCreateDto) SetItemCategoryID(v string)`

SetItemCategoryID sets ItemCategoryID field to given value.


### GetItemGoogleCategoryID

`func (o *ItemTypeCreateDto) GetItemGoogleCategoryID() string`

GetItemGoogleCategoryID returns the ItemGoogleCategoryID field if non-nil, zero value otherwise.

### GetItemGoogleCategoryIDOk

`func (o *ItemTypeCreateDto) GetItemGoogleCategoryIDOk() (*string, bool)`

GetItemGoogleCategoryIDOk returns a tuple with the ItemGoogleCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGoogleCategoryID

`func (o *ItemTypeCreateDto) SetItemGoogleCategoryID(v string)`

SetItemGoogleCategoryID sets ItemGoogleCategoryID field to given value.

### HasItemGoogleCategoryID

`func (o *ItemTypeCreateDto) HasItemGoogleCategoryID() bool`

HasItemGoogleCategoryID returns a boolean if a field has been set.

### SetItemGoogleCategoryIDNil

`func (o *ItemTypeCreateDto) SetItemGoogleCategoryIDNil(b bool)`

 SetItemGoogleCategoryIDNil sets the value for ItemGoogleCategoryID to be an explicit nil

### UnsetItemGoogleCategoryID
`func (o *ItemTypeCreateDto) UnsetItemGoogleCategoryID()`

UnsetItemGoogleCategoryID ensures that no value is present for ItemGoogleCategoryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


