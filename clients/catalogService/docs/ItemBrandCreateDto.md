# ItemBrandCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteURL** | Pointer to **NullableString** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Trending** | Pointer to **bool** |  | [optional] 
**BusinessID** | **string** |  | 

## Methods

### NewItemBrandCreateDto

`func NewItemBrandCreateDto(name string, businessID string, ) *ItemBrandCreateDto`

NewItemBrandCreateDto instantiates a new ItemBrandCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemBrandCreateDtoWithDefaults

`func NewItemBrandCreateDtoWithDefaults() *ItemBrandCreateDto`

NewItemBrandCreateDtoWithDefaults instantiates a new ItemBrandCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemBrandCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemBrandCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemBrandCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemBrandCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemBrandCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemBrandCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemBrandCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemBrandCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCode

`func (o *ItemBrandCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ItemBrandCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ItemBrandCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ItemBrandCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ItemBrandCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ItemBrandCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *ItemBrandCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemBrandCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemBrandCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ItemBrandCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemBrandCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemBrandCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemBrandCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemBrandCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemBrandCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteURL

`func (o *ItemBrandCreateDto) GetWebsiteURL() string`

GetWebsiteURL returns the WebsiteURL field if non-nil, zero value otherwise.

### GetWebsiteURLOk

`func (o *ItemBrandCreateDto) GetWebsiteURLOk() (*string, bool)`

GetWebsiteURLOk returns a tuple with the WebsiteURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteURL

`func (o *ItemBrandCreateDto) SetWebsiteURL(v string)`

SetWebsiteURL sets WebsiteURL field to given value.

### HasWebsiteURL

`func (o *ItemBrandCreateDto) HasWebsiteURL() bool`

HasWebsiteURL returns a boolean if a field has been set.

### SetWebsiteURLNil

`func (o *ItemBrandCreateDto) SetWebsiteURLNil(b bool)`

 SetWebsiteURLNil sets the value for WebsiteURL to be an explicit nil

### UnsetWebsiteURL
`func (o *ItemBrandCreateDto) UnsetWebsiteURL()`

UnsetWebsiteURL ensures that no value is present for WebsiteURL, not even an explicit nil
### GetFeatured

`func (o *ItemBrandCreateDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ItemBrandCreateDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ItemBrandCreateDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ItemBrandCreateDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetTrending

`func (o *ItemBrandCreateDto) GetTrending() bool`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *ItemBrandCreateDto) GetTrendingOk() (*bool, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *ItemBrandCreateDto) SetTrending(v bool)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *ItemBrandCreateDto) HasTrending() bool`

HasTrending returns a boolean if a field has been set.

### GetBusinessID

`func (o *ItemBrandCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemBrandCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemBrandCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


