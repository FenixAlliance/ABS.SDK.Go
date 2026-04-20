# BlogPostCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Published** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Markup** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**CodeType** | Pointer to **NullableString** |  | [optional] 
**BlogPostCategoryId** | Pointer to **NullableString** |  | [optional] 
**WebTemplateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlogPostCreateDto

`func NewBlogPostCreateDto(title string, ) *BlogPostCreateDto`

NewBlogPostCreateDto instantiates a new BlogPostCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostCreateDtoWithDefaults

`func NewBlogPostCreateDtoWithDefaults() *BlogPostCreateDto`

NewBlogPostCreateDtoWithDefaults instantiates a new BlogPostCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlogPostCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlogPostCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlogPostCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlogPostCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPostCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPublished

`func (o *BlogPostCreateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BlogPostCreateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BlogPostCreateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BlogPostCreateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetDescription

`func (o *BlogPostCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlogPostCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlogPostCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlogPostCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BlogPostCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BlogPostCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *BlogPostCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BlogPostCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BlogPostCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BlogPostCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BlogPostCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BlogPostCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMarkup

`func (o *BlogPostCreateDto) GetMarkup() string`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *BlogPostCreateDto) GetMarkupOk() (*string, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *BlogPostCreateDto) SetMarkup(v string)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *BlogPostCreateDto) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *BlogPostCreateDto) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *BlogPostCreateDto) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetFeaturedImageUrl

`func (o *BlogPostCreateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *BlogPostCreateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *BlogPostCreateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *BlogPostCreateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *BlogPostCreateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *BlogPostCreateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetCodeType

`func (o *BlogPostCreateDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *BlogPostCreateDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *BlogPostCreateDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *BlogPostCreateDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *BlogPostCreateDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *BlogPostCreateDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetBlogPostCategoryId

`func (o *BlogPostCreateDto) GetBlogPostCategoryId() string`

GetBlogPostCategoryId returns the BlogPostCategoryId field if non-nil, zero value otherwise.

### GetBlogPostCategoryIdOk

`func (o *BlogPostCreateDto) GetBlogPostCategoryIdOk() (*string, bool)`

GetBlogPostCategoryIdOk returns a tuple with the BlogPostCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostCategoryId

`func (o *BlogPostCreateDto) SetBlogPostCategoryId(v string)`

SetBlogPostCategoryId sets BlogPostCategoryId field to given value.

### HasBlogPostCategoryId

`func (o *BlogPostCreateDto) HasBlogPostCategoryId() bool`

HasBlogPostCategoryId returns a boolean if a field has been set.

### SetBlogPostCategoryIdNil

`func (o *BlogPostCreateDto) SetBlogPostCategoryIdNil(b bool)`

 SetBlogPostCategoryIdNil sets the value for BlogPostCategoryId to be an explicit nil

### UnsetBlogPostCategoryId
`func (o *BlogPostCreateDto) UnsetBlogPostCategoryId()`

UnsetBlogPostCategoryId ensures that no value is present for BlogPostCategoryId, not even an explicit nil
### GetWebTemplateId

`func (o *BlogPostCreateDto) GetWebTemplateId() string`

GetWebTemplateId returns the WebTemplateId field if non-nil, zero value otherwise.

### GetWebTemplateIdOk

`func (o *BlogPostCreateDto) GetWebTemplateIdOk() (*string, bool)`

GetWebTemplateIdOk returns a tuple with the WebTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTemplateId

`func (o *BlogPostCreateDto) SetWebTemplateId(v string)`

SetWebTemplateId sets WebTemplateId field to given value.

### HasWebTemplateId

`func (o *BlogPostCreateDto) HasWebTemplateId() bool`

HasWebTemplateId returns a boolean if a field has been set.

### SetWebTemplateIdNil

`func (o *BlogPostCreateDto) SetWebTemplateIdNil(b bool)`

 SetWebTemplateIdNil sets the value for WebTemplateId to be an explicit nil

### UnsetWebTemplateId
`func (o *BlogPostCreateDto) UnsetWebTemplateId()`

UnsetWebTemplateId ensures that no value is present for WebTemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


