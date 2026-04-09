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
**BusinessID** | **string** |  | 
**ItemCategoryID** | **string** |  | 
**ItemGoogleCategoryID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemTypeDto

`func NewItemTypeDto(pluralTitle string, singularTitle string, businessID string, itemCategoryID string, ) *ItemTypeDto`

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
### GetBusinessID

`func (o *ItemTypeDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemTypeDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemTypeDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetItemCategoryID

`func (o *ItemTypeDto) GetItemCategoryID() string`

GetItemCategoryID returns the ItemCategoryID field if non-nil, zero value otherwise.

### GetItemCategoryIDOk

`func (o *ItemTypeDto) GetItemCategoryIDOk() (*string, bool)`

GetItemCategoryIDOk returns a tuple with the ItemCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCategoryID

`func (o *ItemTypeDto) SetItemCategoryID(v string)`

SetItemCategoryID sets ItemCategoryID field to given value.


### GetItemGoogleCategoryID

`func (o *ItemTypeDto) GetItemGoogleCategoryID() string`

GetItemGoogleCategoryID returns the ItemGoogleCategoryID field if non-nil, zero value otherwise.

### GetItemGoogleCategoryIDOk

`func (o *ItemTypeDto) GetItemGoogleCategoryIDOk() (*string, bool)`

GetItemGoogleCategoryIDOk returns a tuple with the ItemGoogleCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGoogleCategoryID

`func (o *ItemTypeDto) SetItemGoogleCategoryID(v string)`

SetItemGoogleCategoryID sets ItemGoogleCategoryID field to given value.

### HasItemGoogleCategoryID

`func (o *ItemTypeDto) HasItemGoogleCategoryID() bool`

HasItemGoogleCategoryID returns a boolean if a field has been set.

### SetItemGoogleCategoryIDNil

`func (o *ItemTypeDto) SetItemGoogleCategoryIDNil(b bool)`

 SetItemGoogleCategoryIDNil sets the value for ItemGoogleCategoryID to be an explicit nil

### UnsetItemGoogleCategoryID
`func (o *ItemTypeDto) UnsetItemGoogleCategoryID()`

UnsetItemGoogleCategoryID ensures that no value is present for ItemGoogleCategoryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


