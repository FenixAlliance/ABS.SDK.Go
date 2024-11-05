# SocialMediaPostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**SocialPostBucketId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialMediaPostDto

`func NewSocialMediaPostDto() *SocialMediaPostDto`

NewSocialMediaPostDto instantiates a new SocialMediaPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialMediaPostDtoWithDefaults

`func NewSocialMediaPostDtoWithDefaults() *SocialMediaPostDto`

NewSocialMediaPostDtoWithDefaults instantiates a new SocialMediaPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialMediaPostDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialMediaPostDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialMediaPostDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialMediaPostDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialMediaPostDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialMediaPostDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialMediaPostDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialMediaPostDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialMediaPostDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialMediaPostDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialMediaPostDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialMediaPostDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SocialMediaPostDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialMediaPostDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialMediaPostDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialMediaPostDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialMediaPostDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialMediaPostDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *SocialMediaPostDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialMediaPostDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialMediaPostDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialMediaPostDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SocialMediaPostDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SocialMediaPostDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFeaturedImageUrl

`func (o *SocialMediaPostDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *SocialMediaPostDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *SocialMediaPostDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *SocialMediaPostDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *SocialMediaPostDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *SocialMediaPostDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetTenantId

`func (o *SocialMediaPostDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SocialMediaPostDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SocialMediaPostDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SocialMediaPostDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SocialMediaPostDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SocialMediaPostDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSocialPostBucketId

`func (o *SocialMediaPostDto) GetSocialPostBucketId() string`

GetSocialPostBucketId returns the SocialPostBucketId field if non-nil, zero value otherwise.

### GetSocialPostBucketIdOk

`func (o *SocialMediaPostDto) GetSocialPostBucketIdOk() (*string, bool)`

GetSocialPostBucketIdOk returns a tuple with the SocialPostBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostBucketId

`func (o *SocialMediaPostDto) SetSocialPostBucketId(v string)`

SetSocialPostBucketId sets SocialPostBucketId field to given value.

### HasSocialPostBucketId

`func (o *SocialMediaPostDto) HasSocialPostBucketId() bool`

HasSocialPostBucketId returns a boolean if a field has been set.

### SetSocialPostBucketIdNil

`func (o *SocialMediaPostDto) SetSocialPostBucketIdNil(b bool)`

 SetSocialPostBucketIdNil sets the value for SocialPostBucketId to be an explicit nil

### UnsetSocialPostBucketId
`func (o *SocialMediaPostDto) UnsetSocialPostBucketId()`

UnsetSocialPostBucketId ensures that no value is present for SocialPostBucketId, not even an explicit nil
### GetEnrolmentId

`func (o *SocialMediaPostDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *SocialMediaPostDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *SocialMediaPostDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *SocialMediaPostDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *SocialMediaPostDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *SocialMediaPostDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


