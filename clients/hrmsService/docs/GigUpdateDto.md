# GigUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGigUpdateDto

`func NewGigUpdateDto() *GigUpdateDto`

NewGigUpdateDto instantiates a new GigUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGigUpdateDtoWithDefaults

`func NewGigUpdateDtoWithDefaults() *GigUpdateDto`

NewGigUpdateDtoWithDefaults instantiates a new GigUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GigUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GigUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GigUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GigUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GigUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GigUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *GigUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GigUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GigUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GigUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GigUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GigUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *GigUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GigUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GigUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GigUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *GigUpdateDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *GigUpdateDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCurrency

`func (o *GigUpdateDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GigUpdateDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GigUpdateDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GigUpdateDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *GigUpdateDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *GigUpdateDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetLocation

`func (o *GigUpdateDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GigUpdateDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GigUpdateDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GigUpdateDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GigUpdateDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GigUpdateDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetStartDate

`func (o *GigUpdateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GigUpdateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GigUpdateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GigUpdateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GigUpdateDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GigUpdateDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *GigUpdateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GigUpdateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GigUpdateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GigUpdateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GigUpdateDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GigUpdateDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetCategory

`func (o *GigUpdateDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GigUpdateDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GigUpdateDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GigUpdateDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GigUpdateDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GigUpdateDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetTags

`func (o *GigUpdateDto) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GigUpdateDto) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GigUpdateDto) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GigUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GigUpdateDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GigUpdateDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


