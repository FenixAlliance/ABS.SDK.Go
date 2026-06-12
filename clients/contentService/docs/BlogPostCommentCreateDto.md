# BlogPostCommentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Message** | **string** |  | 
**OwnerSocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 
**ParentCommentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlogPostCommentCreateDto

`func NewBlogPostCommentCreateDto(message string, ) *BlogPostCommentCreateDto`

NewBlogPostCommentCreateDto instantiates a new BlogPostCommentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostCommentCreateDtoWithDefaults

`func NewBlogPostCommentCreateDtoWithDefaults() *BlogPostCommentCreateDto`

NewBlogPostCommentCreateDtoWithDefaults instantiates a new BlogPostCommentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostCommentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostCommentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostCommentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostCommentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlogPostCommentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlogPostCommentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlogPostCommentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlogPostCommentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *BlogPostCommentCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BlogPostCommentCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BlogPostCommentCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOwnerSocialProfileId

`func (o *BlogPostCommentCreateDto) GetOwnerSocialProfileId() string`

GetOwnerSocialProfileId returns the OwnerSocialProfileId field if non-nil, zero value otherwise.

### GetOwnerSocialProfileIdOk

`func (o *BlogPostCommentCreateDto) GetOwnerSocialProfileIdOk() (*string, bool)`

GetOwnerSocialProfileIdOk returns a tuple with the OwnerSocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSocialProfileId

`func (o *BlogPostCommentCreateDto) SetOwnerSocialProfileId(v string)`

SetOwnerSocialProfileId sets OwnerSocialProfileId field to given value.

### HasOwnerSocialProfileId

`func (o *BlogPostCommentCreateDto) HasOwnerSocialProfileId() bool`

HasOwnerSocialProfileId returns a boolean if a field has been set.

### SetOwnerSocialProfileIdNil

`func (o *BlogPostCommentCreateDto) SetOwnerSocialProfileIdNil(b bool)`

 SetOwnerSocialProfileIdNil sets the value for OwnerSocialProfileId to be an explicit nil

### UnsetOwnerSocialProfileId
`func (o *BlogPostCommentCreateDto) UnsetOwnerSocialProfileId()`

UnsetOwnerSocialProfileId ensures that no value is present for OwnerSocialProfileId, not even an explicit nil
### GetSocialPostId

`func (o *BlogPostCommentCreateDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *BlogPostCommentCreateDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *BlogPostCommentCreateDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *BlogPostCommentCreateDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *BlogPostCommentCreateDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *BlogPostCommentCreateDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil
### GetParentCommentId

`func (o *BlogPostCommentCreateDto) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *BlogPostCommentCreateDto) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *BlogPostCommentCreateDto) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *BlogPostCommentCreateDto) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### SetParentCommentIdNil

`func (o *BlogPostCommentCreateDto) SetParentCommentIdNil(b bool)`

 SetParentCommentIdNil sets the value for ParentCommentId to be an explicit nil

### UnsetParentCommentId
`func (o *BlogPostCommentCreateDto) UnsetParentCommentId()`

UnsetParentCommentId ensures that no value is present for ParentCommentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


