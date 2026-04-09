# ItemCategoryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**IsFeatured** | Pointer to **bool** |  | [optional] 
**EnableForCourses** | Pointer to **bool** |  | [optional] 
**EnableForProducts** | Pointer to **bool** |  | [optional] 
**EnableForLicenses** | Pointer to **bool** |  | [optional] 
**EnableForServices** | Pointer to **bool** |  | [optional] 
**EnableForSubscriptions** | Pointer to **bool** |  | [optional] 

## Methods

### NewItemCategoryUpdateDto

`func NewItemCategoryUpdateDto(title string, ) *ItemCategoryUpdateDto`

NewItemCategoryUpdateDto instantiates a new ItemCategoryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCategoryUpdateDtoWithDefaults

`func NewItemCategoryUpdateDtoWithDefaults() *ItemCategoryUpdateDto`

NewItemCategoryUpdateDtoWithDefaults instantiates a new ItemCategoryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ItemCategoryUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemCategoryUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemCategoryUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemCategoryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemCategoryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemCategoryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemCategoryUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemCategoryUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemCategoryUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageURL

`func (o *ItemCategoryUpdateDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *ItemCategoryUpdateDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *ItemCategoryUpdateDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *ItemCategoryUpdateDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *ItemCategoryUpdateDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *ItemCategoryUpdateDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetIsFeatured

`func (o *ItemCategoryUpdateDto) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *ItemCategoryUpdateDto) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *ItemCategoryUpdateDto) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.

### HasIsFeatured

`func (o *ItemCategoryUpdateDto) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### GetEnableForCourses

`func (o *ItemCategoryUpdateDto) GetEnableForCourses() bool`

GetEnableForCourses returns the EnableForCourses field if non-nil, zero value otherwise.

### GetEnableForCoursesOk

`func (o *ItemCategoryUpdateDto) GetEnableForCoursesOk() (*bool, bool)`

GetEnableForCoursesOk returns a tuple with the EnableForCourses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForCourses

`func (o *ItemCategoryUpdateDto) SetEnableForCourses(v bool)`

SetEnableForCourses sets EnableForCourses field to given value.

### HasEnableForCourses

`func (o *ItemCategoryUpdateDto) HasEnableForCourses() bool`

HasEnableForCourses returns a boolean if a field has been set.

### GetEnableForProducts

`func (o *ItemCategoryUpdateDto) GetEnableForProducts() bool`

GetEnableForProducts returns the EnableForProducts field if non-nil, zero value otherwise.

### GetEnableForProductsOk

`func (o *ItemCategoryUpdateDto) GetEnableForProductsOk() (*bool, bool)`

GetEnableForProductsOk returns a tuple with the EnableForProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForProducts

`func (o *ItemCategoryUpdateDto) SetEnableForProducts(v bool)`

SetEnableForProducts sets EnableForProducts field to given value.

### HasEnableForProducts

`func (o *ItemCategoryUpdateDto) HasEnableForProducts() bool`

HasEnableForProducts returns a boolean if a field has been set.

### GetEnableForLicenses

`func (o *ItemCategoryUpdateDto) GetEnableForLicenses() bool`

GetEnableForLicenses returns the EnableForLicenses field if non-nil, zero value otherwise.

### GetEnableForLicensesOk

`func (o *ItemCategoryUpdateDto) GetEnableForLicensesOk() (*bool, bool)`

GetEnableForLicensesOk returns a tuple with the EnableForLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForLicenses

`func (o *ItemCategoryUpdateDto) SetEnableForLicenses(v bool)`

SetEnableForLicenses sets EnableForLicenses field to given value.

### HasEnableForLicenses

`func (o *ItemCategoryUpdateDto) HasEnableForLicenses() bool`

HasEnableForLicenses returns a boolean if a field has been set.

### GetEnableForServices

`func (o *ItemCategoryUpdateDto) GetEnableForServices() bool`

GetEnableForServices returns the EnableForServices field if non-nil, zero value otherwise.

### GetEnableForServicesOk

`func (o *ItemCategoryUpdateDto) GetEnableForServicesOk() (*bool, bool)`

GetEnableForServicesOk returns a tuple with the EnableForServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForServices

`func (o *ItemCategoryUpdateDto) SetEnableForServices(v bool)`

SetEnableForServices sets EnableForServices field to given value.

### HasEnableForServices

`func (o *ItemCategoryUpdateDto) HasEnableForServices() bool`

HasEnableForServices returns a boolean if a field has been set.

### GetEnableForSubscriptions

`func (o *ItemCategoryUpdateDto) GetEnableForSubscriptions() bool`

GetEnableForSubscriptions returns the EnableForSubscriptions field if non-nil, zero value otherwise.

### GetEnableForSubscriptionsOk

`func (o *ItemCategoryUpdateDto) GetEnableForSubscriptionsOk() (*bool, bool)`

GetEnableForSubscriptionsOk returns a tuple with the EnableForSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForSubscriptions

`func (o *ItemCategoryUpdateDto) SetEnableForSubscriptions(v bool)`

SetEnableForSubscriptions sets EnableForSubscriptions field to given value.

### HasEnableForSubscriptions

`func (o *ItemCategoryUpdateDto) HasEnableForSubscriptions() bool`

HasEnableForSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


