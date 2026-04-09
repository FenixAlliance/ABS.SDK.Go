# ItemBrandDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteURL** | Pointer to **NullableString** |  | [optional] 
**LogoURL** | Pointer to **NullableString** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Trending** | Pointer to **bool** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemBrandDto

`func NewItemBrandDto() *ItemBrandDto`

NewItemBrandDto instantiates a new ItemBrandDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemBrandDtoWithDefaults

`func NewItemBrandDtoWithDefaults() *ItemBrandDto`

NewItemBrandDtoWithDefaults instantiates a new ItemBrandDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemBrandDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemBrandDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemBrandDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemBrandDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemBrandDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemBrandDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemBrandDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemBrandDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemBrandDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemBrandDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemBrandDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemBrandDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCode

`func (o *ItemBrandDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ItemBrandDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ItemBrandDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ItemBrandDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ItemBrandDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ItemBrandDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *ItemBrandDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemBrandDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemBrandDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemBrandDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ItemBrandDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ItemBrandDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ItemBrandDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemBrandDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemBrandDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemBrandDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemBrandDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemBrandDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteURL

`func (o *ItemBrandDto) GetWebsiteURL() string`

GetWebsiteURL returns the WebsiteURL field if non-nil, zero value otherwise.

### GetWebsiteURLOk

`func (o *ItemBrandDto) GetWebsiteURLOk() (*string, bool)`

GetWebsiteURLOk returns a tuple with the WebsiteURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteURL

`func (o *ItemBrandDto) SetWebsiteURL(v string)`

SetWebsiteURL sets WebsiteURL field to given value.

### HasWebsiteURL

`func (o *ItemBrandDto) HasWebsiteURL() bool`

HasWebsiteURL returns a boolean if a field has been set.

### SetWebsiteURLNil

`func (o *ItemBrandDto) SetWebsiteURLNil(b bool)`

 SetWebsiteURLNil sets the value for WebsiteURL to be an explicit nil

### UnsetWebsiteURL
`func (o *ItemBrandDto) UnsetWebsiteURL()`

UnsetWebsiteURL ensures that no value is present for WebsiteURL, not even an explicit nil
### GetLogoURL

`func (o *ItemBrandDto) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *ItemBrandDto) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *ItemBrandDto) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *ItemBrandDto) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### SetLogoURLNil

`func (o *ItemBrandDto) SetLogoURLNil(b bool)`

 SetLogoURLNil sets the value for LogoURL to be an explicit nil

### UnsetLogoURL
`func (o *ItemBrandDto) UnsetLogoURL()`

UnsetLogoURL ensures that no value is present for LogoURL, not even an explicit nil
### GetFeatured

`func (o *ItemBrandDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ItemBrandDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ItemBrandDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ItemBrandDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetTrending

`func (o *ItemBrandDto) GetTrending() bool`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *ItemBrandDto) GetTrendingOk() (*bool, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *ItemBrandDto) SetTrending(v bool)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *ItemBrandDto) HasTrending() bool`

HasTrending returns a boolean if a field has been set.

### GetBusinessID

`func (o *ItemBrandDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemBrandDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemBrandDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ItemBrandDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ItemBrandDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ItemBrandDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


