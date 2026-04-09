# ItemCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**IsFeatured** | Pointer to **bool** |  | [optional] 
**EnableForCourses** | Pointer to **bool** |  | [optional] 
**EnableForProducts** | Pointer to **bool** |  | [optional] 
**EnableForLicenses** | Pointer to **bool** |  | [optional] 
**EnableForServices** | Pointer to **bool** |  | [optional] 
**EnableForSubscriptions** | Pointer to **bool** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**ParentItemCategoryID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemCategoryDto

`func NewItemCategoryDto() *ItemCategoryDto`

NewItemCategoryDto instantiates a new ItemCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCategoryDtoWithDefaults

`func NewItemCategoryDtoWithDefaults() *ItemCategoryDto`

NewItemCategoryDtoWithDefaults instantiates a new ItemCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemCategoryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemCategoryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemCategoryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemCategoryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemCategoryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemCategoryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemCategoryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemCategoryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemCategoryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemCategoryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemCategoryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemCategoryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ItemCategoryDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemCategoryDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemCategoryDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemCategoryDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemCategoryDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemCategoryDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ItemCategoryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemCategoryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemCategoryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemCategoryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemCategoryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemCategoryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageURL

`func (o *ItemCategoryDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *ItemCategoryDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *ItemCategoryDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *ItemCategoryDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *ItemCategoryDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *ItemCategoryDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetIsFeatured

`func (o *ItemCategoryDto) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *ItemCategoryDto) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *ItemCategoryDto) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.

### HasIsFeatured

`func (o *ItemCategoryDto) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### GetEnableForCourses

`func (o *ItemCategoryDto) GetEnableForCourses() bool`

GetEnableForCourses returns the EnableForCourses field if non-nil, zero value otherwise.

### GetEnableForCoursesOk

`func (o *ItemCategoryDto) GetEnableForCoursesOk() (*bool, bool)`

GetEnableForCoursesOk returns a tuple with the EnableForCourses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForCourses

`func (o *ItemCategoryDto) SetEnableForCourses(v bool)`

SetEnableForCourses sets EnableForCourses field to given value.

### HasEnableForCourses

`func (o *ItemCategoryDto) HasEnableForCourses() bool`

HasEnableForCourses returns a boolean if a field has been set.

### GetEnableForProducts

`func (o *ItemCategoryDto) GetEnableForProducts() bool`

GetEnableForProducts returns the EnableForProducts field if non-nil, zero value otherwise.

### GetEnableForProductsOk

`func (o *ItemCategoryDto) GetEnableForProductsOk() (*bool, bool)`

GetEnableForProductsOk returns a tuple with the EnableForProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForProducts

`func (o *ItemCategoryDto) SetEnableForProducts(v bool)`

SetEnableForProducts sets EnableForProducts field to given value.

### HasEnableForProducts

`func (o *ItemCategoryDto) HasEnableForProducts() bool`

HasEnableForProducts returns a boolean if a field has been set.

### GetEnableForLicenses

`func (o *ItemCategoryDto) GetEnableForLicenses() bool`

GetEnableForLicenses returns the EnableForLicenses field if non-nil, zero value otherwise.

### GetEnableForLicensesOk

`func (o *ItemCategoryDto) GetEnableForLicensesOk() (*bool, bool)`

GetEnableForLicensesOk returns a tuple with the EnableForLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForLicenses

`func (o *ItemCategoryDto) SetEnableForLicenses(v bool)`

SetEnableForLicenses sets EnableForLicenses field to given value.

### HasEnableForLicenses

`func (o *ItemCategoryDto) HasEnableForLicenses() bool`

HasEnableForLicenses returns a boolean if a field has been set.

### GetEnableForServices

`func (o *ItemCategoryDto) GetEnableForServices() bool`

GetEnableForServices returns the EnableForServices field if non-nil, zero value otherwise.

### GetEnableForServicesOk

`func (o *ItemCategoryDto) GetEnableForServicesOk() (*bool, bool)`

GetEnableForServicesOk returns a tuple with the EnableForServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForServices

`func (o *ItemCategoryDto) SetEnableForServices(v bool)`

SetEnableForServices sets EnableForServices field to given value.

### HasEnableForServices

`func (o *ItemCategoryDto) HasEnableForServices() bool`

HasEnableForServices returns a boolean if a field has been set.

### GetEnableForSubscriptions

`func (o *ItemCategoryDto) GetEnableForSubscriptions() bool`

GetEnableForSubscriptions returns the EnableForSubscriptions field if non-nil, zero value otherwise.

### GetEnableForSubscriptionsOk

`func (o *ItemCategoryDto) GetEnableForSubscriptionsOk() (*bool, bool)`

GetEnableForSubscriptionsOk returns a tuple with the EnableForSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForSubscriptions

`func (o *ItemCategoryDto) SetEnableForSubscriptions(v bool)`

SetEnableForSubscriptions sets EnableForSubscriptions field to given value.

### HasEnableForSubscriptions

`func (o *ItemCategoryDto) HasEnableForSubscriptions() bool`

HasEnableForSubscriptions returns a boolean if a field has been set.

### GetBusinessID

`func (o *ItemCategoryDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemCategoryDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemCategoryDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ItemCategoryDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ItemCategoryDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ItemCategoryDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *ItemCategoryDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *ItemCategoryDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *ItemCategoryDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *ItemCategoryDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *ItemCategoryDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *ItemCategoryDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetParentItemCategoryID

`func (o *ItemCategoryDto) GetParentItemCategoryID() string`

GetParentItemCategoryID returns the ParentItemCategoryID field if non-nil, zero value otherwise.

### GetParentItemCategoryIDOk

`func (o *ItemCategoryDto) GetParentItemCategoryIDOk() (*string, bool)`

GetParentItemCategoryIDOk returns a tuple with the ParentItemCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemCategoryID

`func (o *ItemCategoryDto) SetParentItemCategoryID(v string)`

SetParentItemCategoryID sets ParentItemCategoryID field to given value.

### HasParentItemCategoryID

`func (o *ItemCategoryDto) HasParentItemCategoryID() bool`

HasParentItemCategoryID returns a boolean if a field has been set.

### SetParentItemCategoryIDNil

`func (o *ItemCategoryDto) SetParentItemCategoryIDNil(b bool)`

 SetParentItemCategoryIDNil sets the value for ParentItemCategoryID to be an explicit nil

### UnsetParentItemCategoryID
`func (o *ItemCategoryDto) UnsetParentItemCategoryID()`

UnsetParentItemCategoryID ensures that no value is present for ParentItemCategoryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


