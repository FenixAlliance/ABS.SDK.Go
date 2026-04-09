# WebPageCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**CodeType** | Pointer to **NullableString** |  | [optional] 
**WebTemplateID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPageCreateDto

`func NewWebPageCreateDto() *WebPageCreateDto`

NewWebPageCreateDto instantiates a new WebPageCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageCreateDtoWithDefaults

`func NewWebPageCreateDtoWithDefaults() *WebPageCreateDto`

NewWebPageCreateDtoWithDefaults instantiates a new WebPageCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPageCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPageCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPageCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPageCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebPageCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPageCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPageCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPageCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *WebPageCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WebPageCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WebPageCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WebPageCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WebPageCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WebPageCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *WebPageCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *WebPageCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *WebPageCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *WebPageCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *WebPageCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *WebPageCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTitle

`func (o *WebPageCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPageCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPageCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPageCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPageCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPageCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCode

`func (o *WebPageCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebPageCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebPageCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebPageCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WebPageCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WebPageCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPublished

`func (o *WebPageCreateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *WebPageCreateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *WebPageCreateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *WebPageCreateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetDescription

`func (o *WebPageCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPageCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPageCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPageCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPageCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPageCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHtmlContent

`func (o *WebPageCreateDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebPageCreateDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebPageCreateDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebPageCreateDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebPageCreateDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebPageCreateDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetFeaturedImageUrl

`func (o *WebPageCreateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *WebPageCreateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *WebPageCreateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *WebPageCreateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *WebPageCreateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *WebPageCreateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetCodeType

`func (o *WebPageCreateDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *WebPageCreateDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *WebPageCreateDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *WebPageCreateDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *WebPageCreateDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *WebPageCreateDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetWebTemplateID

`func (o *WebPageCreateDto) GetWebTemplateID() string`

GetWebTemplateID returns the WebTemplateID field if non-nil, zero value otherwise.

### GetWebTemplateIDOk

`func (o *WebPageCreateDto) GetWebTemplateIDOk() (*string, bool)`

GetWebTemplateIDOk returns a tuple with the WebTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTemplateID

`func (o *WebPageCreateDto) SetWebTemplateID(v string)`

SetWebTemplateID sets WebTemplateID field to given value.

### HasWebTemplateID

`func (o *WebPageCreateDto) HasWebTemplateID() bool`

HasWebTemplateID returns a boolean if a field has been set.

### SetWebTemplateIDNil

`func (o *WebPageCreateDto) SetWebTemplateIDNil(b bool)`

 SetWebTemplateIDNil sets the value for WebTemplateID to be an explicit nil

### UnsetWebTemplateID
`func (o *WebPageCreateDto) UnsetWebTemplateID()`

UnsetWebTemplateID ensures that no value is present for WebTemplateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


