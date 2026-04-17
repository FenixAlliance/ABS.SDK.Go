# SocialMediaPostCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**SocialPostBucketId** | Pointer to **NullableString** |  | [optional] 

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

### GetId

`func (o *SocialMediaPostCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialMediaPostCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialMediaPostCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialMediaPostCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SocialMediaPostCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialMediaPostCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialMediaPostCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialMediaPostCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


