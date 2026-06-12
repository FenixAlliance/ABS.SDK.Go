# BlogPostCommentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**BlogPostId** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**OwnerSocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 
**ParentCommentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlogPostCommentDto

`func NewBlogPostCommentDto() *BlogPostCommentDto`

NewBlogPostCommentDto instantiates a new BlogPostCommentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostCommentDtoWithDefaults

`func NewBlogPostCommentDtoWithDefaults() *BlogPostCommentDto`

NewBlogPostCommentDtoWithDefaults instantiates a new BlogPostCommentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostCommentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostCommentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostCommentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostCommentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BlogPostCommentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BlogPostCommentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BlogPostCommentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlogPostCommentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlogPostCommentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlogPostCommentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BlogPostCommentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BlogPostCommentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBlogPostId

`func (o *BlogPostCommentDto) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *BlogPostCommentDto) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *BlogPostCommentDto) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *BlogPostCommentDto) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### SetBlogPostIdNil

`func (o *BlogPostCommentDto) SetBlogPostIdNil(b bool)`

 SetBlogPostIdNil sets the value for BlogPostId to be an explicit nil

### UnsetBlogPostId
`func (o *BlogPostCommentDto) UnsetBlogPostId()`

UnsetBlogPostId ensures that no value is present for BlogPostId, not even an explicit nil
### GetMessage

`func (o *BlogPostCommentDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BlogPostCommentDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BlogPostCommentDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BlogPostCommentDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BlogPostCommentDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BlogPostCommentDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOwnerSocialProfileId

`func (o *BlogPostCommentDto) GetOwnerSocialProfileId() string`

GetOwnerSocialProfileId returns the OwnerSocialProfileId field if non-nil, zero value otherwise.

### GetOwnerSocialProfileIdOk

`func (o *BlogPostCommentDto) GetOwnerSocialProfileIdOk() (*string, bool)`

GetOwnerSocialProfileIdOk returns a tuple with the OwnerSocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSocialProfileId

`func (o *BlogPostCommentDto) SetOwnerSocialProfileId(v string)`

SetOwnerSocialProfileId sets OwnerSocialProfileId field to given value.

### HasOwnerSocialProfileId

`func (o *BlogPostCommentDto) HasOwnerSocialProfileId() bool`

HasOwnerSocialProfileId returns a boolean if a field has been set.

### SetOwnerSocialProfileIdNil

`func (o *BlogPostCommentDto) SetOwnerSocialProfileIdNil(b bool)`

 SetOwnerSocialProfileIdNil sets the value for OwnerSocialProfileId to be an explicit nil

### UnsetOwnerSocialProfileId
`func (o *BlogPostCommentDto) UnsetOwnerSocialProfileId()`

UnsetOwnerSocialProfileId ensures that no value is present for OwnerSocialProfileId, not even an explicit nil
### GetSocialPostId

`func (o *BlogPostCommentDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *BlogPostCommentDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *BlogPostCommentDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *BlogPostCommentDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *BlogPostCommentDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *BlogPostCommentDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil
### GetParentCommentId

`func (o *BlogPostCommentDto) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *BlogPostCommentDto) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *BlogPostCommentDto) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *BlogPostCommentDto) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### SetParentCommentIdNil

`func (o *BlogPostCommentDto) SetParentCommentIdNil(b bool)`

 SetParentCommentIdNil sets the value for ParentCommentId to be an explicit nil

### UnsetParentCommentId
`func (o *BlogPostCommentDto) UnsetParentCommentId()`

UnsetParentCommentId ensures that no value is present for ParentCommentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


