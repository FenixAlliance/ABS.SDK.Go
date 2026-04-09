# BlogPostCommentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Message** | **string** |  | 
**OwnerSocialProfileID** | Pointer to **NullableString** |  | [optional] 
**SocialPostID** | Pointer to **NullableString** |  | [optional] 
**ParentCommentID** | Pointer to **NullableString** |  | [optional] 

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


### GetOwnerSocialProfileID

`func (o *BlogPostCommentCreateDto) GetOwnerSocialProfileID() string`

GetOwnerSocialProfileID returns the OwnerSocialProfileID field if non-nil, zero value otherwise.

### GetOwnerSocialProfileIDOk

`func (o *BlogPostCommentCreateDto) GetOwnerSocialProfileIDOk() (*string, bool)`

GetOwnerSocialProfileIDOk returns a tuple with the OwnerSocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSocialProfileID

`func (o *BlogPostCommentCreateDto) SetOwnerSocialProfileID(v string)`

SetOwnerSocialProfileID sets OwnerSocialProfileID field to given value.

### HasOwnerSocialProfileID

`func (o *BlogPostCommentCreateDto) HasOwnerSocialProfileID() bool`

HasOwnerSocialProfileID returns a boolean if a field has been set.

### SetOwnerSocialProfileIDNil

`func (o *BlogPostCommentCreateDto) SetOwnerSocialProfileIDNil(b bool)`

 SetOwnerSocialProfileIDNil sets the value for OwnerSocialProfileID to be an explicit nil

### UnsetOwnerSocialProfileID
`func (o *BlogPostCommentCreateDto) UnsetOwnerSocialProfileID()`

UnsetOwnerSocialProfileID ensures that no value is present for OwnerSocialProfileID, not even an explicit nil
### GetSocialPostID

`func (o *BlogPostCommentCreateDto) GetSocialPostID() string`

GetSocialPostID returns the SocialPostID field if non-nil, zero value otherwise.

### GetSocialPostIDOk

`func (o *BlogPostCommentCreateDto) GetSocialPostIDOk() (*string, bool)`

GetSocialPostIDOk returns a tuple with the SocialPostID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostID

`func (o *BlogPostCommentCreateDto) SetSocialPostID(v string)`

SetSocialPostID sets SocialPostID field to given value.

### HasSocialPostID

`func (o *BlogPostCommentCreateDto) HasSocialPostID() bool`

HasSocialPostID returns a boolean if a field has been set.

### SetSocialPostIDNil

`func (o *BlogPostCommentCreateDto) SetSocialPostIDNil(b bool)`

 SetSocialPostIDNil sets the value for SocialPostID to be an explicit nil

### UnsetSocialPostID
`func (o *BlogPostCommentCreateDto) UnsetSocialPostID()`

UnsetSocialPostID ensures that no value is present for SocialPostID, not even an explicit nil
### GetParentCommentID

`func (o *BlogPostCommentCreateDto) GetParentCommentID() string`

GetParentCommentID returns the ParentCommentID field if non-nil, zero value otherwise.

### GetParentCommentIDOk

`func (o *BlogPostCommentCreateDto) GetParentCommentIDOk() (*string, bool)`

GetParentCommentIDOk returns a tuple with the ParentCommentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentID

`func (o *BlogPostCommentCreateDto) SetParentCommentID(v string)`

SetParentCommentID sets ParentCommentID field to given value.

### HasParentCommentID

`func (o *BlogPostCommentCreateDto) HasParentCommentID() bool`

HasParentCommentID returns a boolean if a field has been set.

### SetParentCommentIDNil

`func (o *BlogPostCommentCreateDto) SetParentCommentIDNil(b bool)`

 SetParentCommentIDNil sets the value for ParentCommentID to be an explicit nil

### UnsetParentCommentID
`func (o *BlogPostCommentCreateDto) UnsetParentCommentID()`

UnsetParentCommentID ensures that no value is present for ParentCommentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


