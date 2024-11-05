# SocialMediaPostCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**SocialPostBucketId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialMediaPostCreateDto

`func NewSocialMediaPostCreateDto() *SocialMediaPostCreateDto`

NewSocialMediaPostCreateDto instantiates a new SocialMediaPostCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialMediaPostCreateDtoWithDefaults

`func NewSocialMediaPostCreateDtoWithDefaults() *SocialMediaPostCreateDto`

NewSocialMediaPostCreateDtoWithDefaults instantiates a new SocialMediaPostCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SocialMediaPostCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialMediaPostCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialMediaPostCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialMediaPostCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialMediaPostCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialMediaPostCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *SocialMediaPostCreateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialMediaPostCreateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialMediaPostCreateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialMediaPostCreateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SocialMediaPostCreateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SocialMediaPostCreateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFeaturedImageUrl

`func (o *SocialMediaPostCreateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *SocialMediaPostCreateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *SocialMediaPostCreateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *SocialMediaPostCreateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *SocialMediaPostCreateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *SocialMediaPostCreateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetTenantId

`func (o *SocialMediaPostCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SocialMediaPostCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SocialMediaPostCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SocialMediaPostCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SocialMediaPostCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SocialMediaPostCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSocialPostBucketId

`func (o *SocialMediaPostCreateDto) GetSocialPostBucketId() string`

GetSocialPostBucketId returns the SocialPostBucketId field if non-nil, zero value otherwise.

### GetSocialPostBucketIdOk

`func (o *SocialMediaPostCreateDto) GetSocialPostBucketIdOk() (*string, bool)`

GetSocialPostBucketIdOk returns a tuple with the SocialPostBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostBucketId

`func (o *SocialMediaPostCreateDto) SetSocialPostBucketId(v string)`

SetSocialPostBucketId sets SocialPostBucketId field to given value.

### HasSocialPostBucketId

`func (o *SocialMediaPostCreateDto) HasSocialPostBucketId() bool`

HasSocialPostBucketId returns a boolean if a field has been set.

### SetSocialPostBucketIdNil

`func (o *SocialMediaPostCreateDto) SetSocialPostBucketIdNil(b bool)`

 SetSocialPostBucketIdNil sets the value for SocialPostBucketId to be an explicit nil

### UnsetSocialPostBucketId
`func (o *SocialMediaPostCreateDto) UnsetSocialPostBucketId()`

UnsetSocialPostBucketId ensures that no value is present for SocialPostBucketId, not even an explicit nil
### GetEnrolmentId

`func (o *SocialMediaPostCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *SocialMediaPostCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *SocialMediaPostCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *SocialMediaPostCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *SocialMediaPostCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *SocialMediaPostCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


