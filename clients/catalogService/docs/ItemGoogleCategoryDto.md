# ItemGoogleCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Path** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Disabled** | Pointer to **bool** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**DisplayOnMenu** | Pointer to **bool** |  | [optional] 
**HasChildren** | Pointer to **bool** |  | [optional] 
**MenuImage** | Pointer to **NullableString** |  | [optional] 
**BannerCode** | Pointer to **NullableString** |  | [optional] 
**BannerImage** | Pointer to **NullableString** |  | [optional] 
**PrimaryImage** | Pointer to **NullableString** |  | [optional] 
**ParentCategoryId** | Pointer to **NullableString** |  | [optional] 
**StartingAtAmountInUsd** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewItemGoogleCategoryDto

`func NewItemGoogleCategoryDto(path string, name string, ) *ItemGoogleCategoryDto`

NewItemGoogleCategoryDto instantiates a new ItemGoogleCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGoogleCategoryDtoWithDefaults

`func NewItemGoogleCategoryDtoWithDefaults() *ItemGoogleCategoryDto`

NewItemGoogleCategoryDtoWithDefaults instantiates a new ItemGoogleCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemGoogleCategoryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemGoogleCategoryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemGoogleCategoryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemGoogleCategoryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemGoogleCategoryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemGoogleCategoryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemGoogleCategoryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemGoogleCategoryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemGoogleCategoryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemGoogleCategoryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemGoogleCategoryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemGoogleCategoryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPath

`func (o *ItemGoogleCategoryDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ItemGoogleCategoryDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ItemGoogleCategoryDto) SetPath(v string)`

SetPath sets Path field to given value.


### GetIcon

`func (o *ItemGoogleCategoryDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ItemGoogleCategoryDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ItemGoogleCategoryDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ItemGoogleCategoryDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ItemGoogleCategoryDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ItemGoogleCategoryDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetName

`func (o *ItemGoogleCategoryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemGoogleCategoryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemGoogleCategoryDto) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *ItemGoogleCategoryDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ItemGoogleCategoryDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ItemGoogleCategoryDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ItemGoogleCategoryDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFeatured

`func (o *ItemGoogleCategoryDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ItemGoogleCategoryDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ItemGoogleCategoryDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ItemGoogleCategoryDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetDisplayOnMenu

`func (o *ItemGoogleCategoryDto) GetDisplayOnMenu() bool`

GetDisplayOnMenu returns the DisplayOnMenu field if non-nil, zero value otherwise.

### GetDisplayOnMenuOk

`func (o *ItemGoogleCategoryDto) GetDisplayOnMenuOk() (*bool, bool)`

GetDisplayOnMenuOk returns a tuple with the DisplayOnMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnMenu

`func (o *ItemGoogleCategoryDto) SetDisplayOnMenu(v bool)`

SetDisplayOnMenu sets DisplayOnMenu field to given value.

### HasDisplayOnMenu

`func (o *ItemGoogleCategoryDto) HasDisplayOnMenu() bool`

HasDisplayOnMenu returns a boolean if a field has been set.

### GetHasChildren

`func (o *ItemGoogleCategoryDto) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *ItemGoogleCategoryDto) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *ItemGoogleCategoryDto) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *ItemGoogleCategoryDto) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetMenuImage

`func (o *ItemGoogleCategoryDto) GetMenuImage() string`

GetMenuImage returns the MenuImage field if non-nil, zero value otherwise.

### GetMenuImageOk

`func (o *ItemGoogleCategoryDto) GetMenuImageOk() (*string, bool)`

GetMenuImageOk returns a tuple with the MenuImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuImage

`func (o *ItemGoogleCategoryDto) SetMenuImage(v string)`

SetMenuImage sets MenuImage field to given value.

### HasMenuImage

`func (o *ItemGoogleCategoryDto) HasMenuImage() bool`

HasMenuImage returns a boolean if a field has been set.

### SetMenuImageNil

`func (o *ItemGoogleCategoryDto) SetMenuImageNil(b bool)`

 SetMenuImageNil sets the value for MenuImage to be an explicit nil

### UnsetMenuImage
`func (o *ItemGoogleCategoryDto) UnsetMenuImage()`

UnsetMenuImage ensures that no value is present for MenuImage, not even an explicit nil
### GetBannerCode

`func (o *ItemGoogleCategoryDto) GetBannerCode() string`

GetBannerCode returns the BannerCode field if non-nil, zero value otherwise.

### GetBannerCodeOk

`func (o *ItemGoogleCategoryDto) GetBannerCodeOk() (*string, bool)`

GetBannerCodeOk returns a tuple with the BannerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerCode

`func (o *ItemGoogleCategoryDto) SetBannerCode(v string)`

SetBannerCode sets BannerCode field to given value.

### HasBannerCode

`func (o *ItemGoogleCategoryDto) HasBannerCode() bool`

HasBannerCode returns a boolean if a field has been set.

### SetBannerCodeNil

`func (o *ItemGoogleCategoryDto) SetBannerCodeNil(b bool)`

 SetBannerCodeNil sets the value for BannerCode to be an explicit nil

### UnsetBannerCode
`func (o *ItemGoogleCategoryDto) UnsetBannerCode()`

UnsetBannerCode ensures that no value is present for BannerCode, not even an explicit nil
### GetBannerImage

`func (o *ItemGoogleCategoryDto) GetBannerImage() string`

GetBannerImage returns the BannerImage field if non-nil, zero value otherwise.

### GetBannerImageOk

`func (o *ItemGoogleCategoryDto) GetBannerImageOk() (*string, bool)`

GetBannerImageOk returns a tuple with the BannerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImage

`func (o *ItemGoogleCategoryDto) SetBannerImage(v string)`

SetBannerImage sets BannerImage field to given value.

### HasBannerImage

`func (o *ItemGoogleCategoryDto) HasBannerImage() bool`

HasBannerImage returns a boolean if a field has been set.

### SetBannerImageNil

`func (o *ItemGoogleCategoryDto) SetBannerImageNil(b bool)`

 SetBannerImageNil sets the value for BannerImage to be an explicit nil

### UnsetBannerImage
`func (o *ItemGoogleCategoryDto) UnsetBannerImage()`

UnsetBannerImage ensures that no value is present for BannerImage, not even an explicit nil
### GetPrimaryImage

`func (o *ItemGoogleCategoryDto) GetPrimaryImage() string`

GetPrimaryImage returns the PrimaryImage field if non-nil, zero value otherwise.

### GetPrimaryImageOk

`func (o *ItemGoogleCategoryDto) GetPrimaryImageOk() (*string, bool)`

GetPrimaryImageOk returns a tuple with the PrimaryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImage

`func (o *ItemGoogleCategoryDto) SetPrimaryImage(v string)`

SetPrimaryImage sets PrimaryImage field to given value.

### HasPrimaryImage

`func (o *ItemGoogleCategoryDto) HasPrimaryImage() bool`

HasPrimaryImage returns a boolean if a field has been set.

### SetPrimaryImageNil

`func (o *ItemGoogleCategoryDto) SetPrimaryImageNil(b bool)`

 SetPrimaryImageNil sets the value for PrimaryImage to be an explicit nil

### UnsetPrimaryImage
`func (o *ItemGoogleCategoryDto) UnsetPrimaryImage()`

UnsetPrimaryImage ensures that no value is present for PrimaryImage, not even an explicit nil
### GetParentCategoryId

`func (o *ItemGoogleCategoryDto) GetParentCategoryId() string`

GetParentCategoryId returns the ParentCategoryId field if non-nil, zero value otherwise.

### GetParentCategoryIdOk

`func (o *ItemGoogleCategoryDto) GetParentCategoryIdOk() (*string, bool)`

GetParentCategoryIdOk returns a tuple with the ParentCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategoryId

`func (o *ItemGoogleCategoryDto) SetParentCategoryId(v string)`

SetParentCategoryId sets ParentCategoryId field to given value.

### HasParentCategoryId

`func (o *ItemGoogleCategoryDto) HasParentCategoryId() bool`

HasParentCategoryId returns a boolean if a field has been set.

### SetParentCategoryIdNil

`func (o *ItemGoogleCategoryDto) SetParentCategoryIdNil(b bool)`

 SetParentCategoryIdNil sets the value for ParentCategoryId to be an explicit nil

### UnsetParentCategoryId
`func (o *ItemGoogleCategoryDto) UnsetParentCategoryId()`

UnsetParentCategoryId ensures that no value is present for ParentCategoryId, not even an explicit nil
### GetStartingAtAmountInUsd

`func (o *ItemGoogleCategoryDto) GetStartingAtAmountInUsd() float64`

GetStartingAtAmountInUsd returns the StartingAtAmountInUsd field if non-nil, zero value otherwise.

### GetStartingAtAmountInUsdOk

`func (o *ItemGoogleCategoryDto) GetStartingAtAmountInUsdOk() (*float64, bool)`

GetStartingAtAmountInUsdOk returns a tuple with the StartingAtAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAtAmountInUsd

`func (o *ItemGoogleCategoryDto) SetStartingAtAmountInUsd(v float64)`

SetStartingAtAmountInUsd sets StartingAtAmountInUsd field to given value.

### HasStartingAtAmountInUsd

`func (o *ItemGoogleCategoryDto) HasStartingAtAmountInUsd() bool`

HasStartingAtAmountInUsd returns a boolean if a field has been set.

### SetStartingAtAmountInUsdNil

`func (o *ItemGoogleCategoryDto) SetStartingAtAmountInUsdNil(b bool)`

 SetStartingAtAmountInUsdNil sets the value for StartingAtAmountInUsd to be an explicit nil

### UnsetStartingAtAmountInUsd
`func (o *ItemGoogleCategoryDto) UnsetStartingAtAmountInUsd()`

UnsetStartingAtAmountInUsd ensures that no value is present for StartingAtAmountInUsd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


