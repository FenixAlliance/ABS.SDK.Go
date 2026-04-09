# BlogPostCommentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**BlogPostID** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**OwnerSocialProfileID** | Pointer to **NullableString** |  | [optional] 
**SocialPostID** | Pointer to **NullableString** |  | [optional] 
**ParentCommentID** | Pointer to **NullableString** |  | [optional] 

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
### GetBlogPostID

`func (o *BlogPostCommentDto) GetBlogPostID() string`

GetBlogPostID returns the BlogPostID field if non-nil, zero value otherwise.

### GetBlogPostIDOk

`func (o *BlogPostCommentDto) GetBlogPostIDOk() (*string, bool)`

GetBlogPostIDOk returns a tuple with the BlogPostID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostID

`func (o *BlogPostCommentDto) SetBlogPostID(v string)`

SetBlogPostID sets BlogPostID field to given value.

### HasBlogPostID

`func (o *BlogPostCommentDto) HasBlogPostID() bool`

HasBlogPostID returns a boolean if a field has been set.

### SetBlogPostIDNil

`func (o *BlogPostCommentDto) SetBlogPostIDNil(b bool)`

 SetBlogPostIDNil sets the value for BlogPostID to be an explicit nil

### UnsetBlogPostID
`func (o *BlogPostCommentDto) UnsetBlogPostID()`

UnsetBlogPostID ensures that no value is present for BlogPostID, not even an explicit nil
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
### GetOwnerSocialProfileID

`func (o *BlogPostCommentDto) GetOwnerSocialProfileID() string`

GetOwnerSocialProfileID returns the OwnerSocialProfileID field if non-nil, zero value otherwise.

### GetOwnerSocialProfileIDOk

`func (o *BlogPostCommentDto) GetOwnerSocialProfileIDOk() (*string, bool)`

GetOwnerSocialProfileIDOk returns a tuple with the OwnerSocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSocialProfileID

`func (o *BlogPostCommentDto) SetOwnerSocialProfileID(v string)`

SetOwnerSocialProfileID sets OwnerSocialProfileID field to given value.

### HasOwnerSocialProfileID

`func (o *BlogPostCommentDto) HasOwnerSocialProfileID() bool`

HasOwnerSocialProfileID returns a boolean if a field has been set.

### SetOwnerSocialProfileIDNil

`func (o *BlogPostCommentDto) SetOwnerSocialProfileIDNil(b bool)`

 SetOwnerSocialProfileIDNil sets the value for OwnerSocialProfileID to be an explicit nil

### UnsetOwnerSocialProfileID
`func (o *BlogPostCommentDto) UnsetOwnerSocialProfileID()`

UnsetOwnerSocialProfileID ensures that no value is present for OwnerSocialProfileID, not even an explicit nil
### GetSocialPostID

`func (o *BlogPostCommentDto) GetSocialPostID() string`

GetSocialPostID returns the SocialPostID field if non-nil, zero value otherwise.

### GetSocialPostIDOk

`func (o *BlogPostCommentDto) GetSocialPostIDOk() (*string, bool)`

GetSocialPostIDOk returns a tuple with the SocialPostID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostID

`func (o *BlogPostCommentDto) SetSocialPostID(v string)`

SetSocialPostID sets SocialPostID field to given value.

### HasSocialPostID

`func (o *BlogPostCommentDto) HasSocialPostID() bool`

HasSocialPostID returns a boolean if a field has been set.

### SetSocialPostIDNil

`func (o *BlogPostCommentDto) SetSocialPostIDNil(b bool)`

 SetSocialPostIDNil sets the value for SocialPostID to be an explicit nil

### UnsetSocialPostID
`func (o *BlogPostCommentDto) UnsetSocialPostID()`

UnsetSocialPostID ensures that no value is present for SocialPostID, not even an explicit nil
### GetParentCommentID

`func (o *BlogPostCommentDto) GetParentCommentID() string`

GetParentCommentID returns the ParentCommentID field if non-nil, zero value otherwise.

### GetParentCommentIDOk

`func (o *BlogPostCommentDto) GetParentCommentIDOk() (*string, bool)`

GetParentCommentIDOk returns a tuple with the ParentCommentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentID

`func (o *BlogPostCommentDto) SetParentCommentID(v string)`

SetParentCommentID sets ParentCommentID field to given value.

### HasParentCommentID

`func (o *BlogPostCommentDto) HasParentCommentID() bool`

HasParentCommentID returns a boolean if a field has been set.

### SetParentCommentIDNil

`func (o *BlogPostCommentDto) SetParentCommentIDNil(b bool)`

 SetParentCommentIDNil sets the value for ParentCommentID to be an explicit nil

### UnsetParentCommentID
`func (o *BlogPostCommentDto) UnsetParentCommentID()`

UnsetParentCommentID ensures that no value is present for ParentCommentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


