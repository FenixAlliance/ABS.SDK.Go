# EmailTemplateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**AuthorId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**MarketingCampaignId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailTemplateCreateDto

`func NewEmailTemplateCreateDto() *EmailTemplateCreateDto`

NewEmailTemplateCreateDto instantiates a new EmailTemplateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateCreateDtoWithDefaults

`func NewEmailTemplateCreateDtoWithDefaults() *EmailTemplateCreateDto`

NewEmailTemplateCreateDtoWithDefaults instantiates a new EmailTemplateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailTemplateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailTemplateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailTemplateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailTemplateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EmailTemplateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmailTemplateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmailTemplateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmailTemplateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *EmailTemplateCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmailTemplateCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmailTemplateCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmailTemplateCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmailTemplateCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmailTemplateCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *EmailTemplateCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *EmailTemplateCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *EmailTemplateCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *EmailTemplateCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *EmailTemplateCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *EmailTemplateCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetTitle

`func (o *EmailTemplateCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmailTemplateCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmailTemplateCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EmailTemplateCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EmailTemplateCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EmailTemplateCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthorId

`func (o *EmailTemplateCreateDto) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *EmailTemplateCreateDto) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *EmailTemplateCreateDto) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *EmailTemplateCreateDto) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *EmailTemplateCreateDto) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *EmailTemplateCreateDto) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetDescription

`func (o *EmailTemplateCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailTemplateCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailTemplateCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailTemplateCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EmailTemplateCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EmailTemplateCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHtmlContent

`func (o *EmailTemplateCreateDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *EmailTemplateCreateDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *EmailTemplateCreateDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *EmailTemplateCreateDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *EmailTemplateCreateDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *EmailTemplateCreateDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetFeaturedImageUrl

`func (o *EmailTemplateCreateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *EmailTemplateCreateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *EmailTemplateCreateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *EmailTemplateCreateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *EmailTemplateCreateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *EmailTemplateCreateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetMarketingCampaignId

`func (o *EmailTemplateCreateDto) GetMarketingCampaignId() string`

GetMarketingCampaignId returns the MarketingCampaignId field if non-nil, zero value otherwise.

### GetMarketingCampaignIdOk

`func (o *EmailTemplateCreateDto) GetMarketingCampaignIdOk() (*string, bool)`

GetMarketingCampaignIdOk returns a tuple with the MarketingCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingCampaignId

`func (o *EmailTemplateCreateDto) SetMarketingCampaignId(v string)`

SetMarketingCampaignId sets MarketingCampaignId field to given value.

### HasMarketingCampaignId

`func (o *EmailTemplateCreateDto) HasMarketingCampaignId() bool`

HasMarketingCampaignId returns a boolean if a field has been set.

### SetMarketingCampaignIdNil

`func (o *EmailTemplateCreateDto) SetMarketingCampaignIdNil(b bool)`

 SetMarketingCampaignIdNil sets the value for MarketingCampaignId to be an explicit nil

### UnsetMarketingCampaignId
`func (o *EmailTemplateCreateDto) UnsetMarketingCampaignId()`

UnsetMarketingCampaignId ensures that no value is present for MarketingCampaignId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


