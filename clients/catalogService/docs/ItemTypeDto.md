# ItemTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**PluralTitle** | **string** |  | 
**SingularTitle** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**GoogleCategoryTaxonomy** | Pointer to **NullableString** |  | [optional] 
**TenantId** | **string** |  | 
**ItemCategoryId** | **string** |  | 
**ItemGoogleCategoryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemTypeDto

`func NewItemTypeDto(pluralTitle string, singularTitle string, tenantId string, itemCategoryId string, ) *ItemTypeDto`

NewItemTypeDto instantiates a new ItemTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTypeDtoWithDefaults

`func NewItemTypeDtoWithDefaults() *ItemTypeDto`

NewItemTypeDtoWithDefaults instantiates a new ItemTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemTypeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemTypeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemTypeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemTypeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemTypeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemTypeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemTypeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemTypeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemTypeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemTypeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemTypeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemTypeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPluralTitle

`func (o *ItemTypeDto) GetPluralTitle() string`

GetPluralTitle returns the PluralTitle field if non-nil, zero value otherwise.

### GetPluralTitleOk

`func (o *ItemTypeDto) GetPluralTitleOk() (*string, bool)`

GetPluralTitleOk returns a tuple with the PluralTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralTitle

`func (o *ItemTypeDto) SetPluralTitle(v string)`

SetPluralTitle sets PluralTitle field to given value.


### GetSingularTitle

`func (o *ItemTypeDto) GetSingularTitle() string`

GetSingularTitle returns the SingularTitle field if non-nil, zero value otherwise.

### GetSingularTitleOk

`func (o *ItemTypeDto) GetSingularTitleOk() (*string, bool)`

GetSingularTitleOk returns a tuple with the SingularTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularTitle

`func (o *ItemTypeDto) SetSingularTitle(v string)`

SetSingularTitle sets SingularTitle field to given value.


### GetDescription

`func (o *ItemTypeDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemTypeDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemTypeDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemTypeDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemTypeDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemTypeDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageURL

`func (o *ItemTypeDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *ItemTypeDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *ItemTypeDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *ItemTypeDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *ItemTypeDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *ItemTypeDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetGoogleCategoryTaxonomy

`func (o *ItemTypeDto) GetGoogleCategoryTaxonomy() string`

GetGoogleCategoryTaxonomy returns the GoogleCategoryTaxonomy field if non-nil, zero value otherwise.

### GetGoogleCategoryTaxonomyOk

`func (o *ItemTypeDto) GetGoogleCategoryTaxonomyOk() (*string, bool)`

GetGoogleCategoryTaxonomyOk returns a tuple with the GoogleCategoryTaxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCategoryTaxonomy

`func (o *ItemTypeDto) SetGoogleCategoryTaxonomy(v string)`

SetGoogleCategoryTaxonomy sets GoogleCategoryTaxonomy field to given value.

### HasGoogleCategoryTaxonomy

`func (o *ItemTypeDto) HasGoogleCategoryTaxonomy() bool`

HasGoogleCategoryTaxonomy returns a boolean if a field has been set.

### SetGoogleCategoryTaxonomyNil

`func (o *ItemTypeDto) SetGoogleCategoryTaxonomyNil(b bool)`

 SetGoogleCategoryTaxonomyNil sets the value for GoogleCategoryTaxonomy to be an explicit nil

### UnsetGoogleCategoryTaxonomy
`func (o *ItemTypeDto) UnsetGoogleCategoryTaxonomy()`

UnsetGoogleCategoryTaxonomy ensures that no value is present for GoogleCategoryTaxonomy, not even an explicit nil
### GetTenantId

`func (o *ItemTypeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemTypeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemTypeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetItemCategoryId

`func (o *ItemTypeDto) GetItemCategoryId() string`

GetItemCategoryId returns the ItemCategoryId field if non-nil, zero value otherwise.

### GetItemCategoryIdOk

`func (o *ItemTypeDto) GetItemCategoryIdOk() (*string, bool)`

GetItemCategoryIdOk returns a tuple with the ItemCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCategoryId

`func (o *ItemTypeDto) SetItemCategoryId(v string)`

SetItemCategoryId sets ItemCategoryId field to given value.


### GetItemGoogleCategoryId

`func (o *ItemTypeDto) GetItemGoogleCategoryId() string`

GetItemGoogleCategoryId returns the ItemGoogleCategoryId field if non-nil, zero value otherwise.

### GetItemGoogleCategoryIdOk

`func (o *ItemTypeDto) GetItemGoogleCategoryIdOk() (*string, bool)`

GetItemGoogleCategoryIdOk returns a tuple with the ItemGoogleCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGoogleCategoryId

`func (o *ItemTypeDto) SetItemGoogleCategoryId(v string)`

SetItemGoogleCategoryId sets ItemGoogleCategoryId field to given value.

### HasItemGoogleCategoryId

`func (o *ItemTypeDto) HasItemGoogleCategoryId() bool`

HasItemGoogleCategoryId returns a boolean if a field has been set.

### SetItemGoogleCategoryIdNil

`func (o *ItemTypeDto) SetItemGoogleCategoryIdNil(b bool)`

 SetItemGoogleCategoryIdNil sets the value for ItemGoogleCategoryId to be an explicit nil

### UnsetItemGoogleCategoryId
`func (o *ItemTypeDto) UnsetItemGoogleCategoryId()`

UnsetItemGoogleCategoryId ensures that no value is present for ItemGoogleCategoryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


