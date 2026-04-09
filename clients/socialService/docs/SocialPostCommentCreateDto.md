# SocialPostCommentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Message** | **string** |  | 
**ParentCommentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialFeedPostId** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostCommentCreateDto

`func NewSocialPostCommentCreateDto(message string, ) *SocialPostCommentCreateDto`

NewSocialPostCommentCreateDto instantiates a new SocialPostCommentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostCommentCreateDtoWithDefaults

`func NewSocialPostCommentCreateDtoWithDefaults() *SocialPostCommentCreateDto`

NewSocialPostCommentCreateDtoWithDefaults instantiates a new SocialPostCommentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostCommentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostCommentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostCommentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostCommentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SocialPostCommentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostCommentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostCommentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostCommentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *SocialPostCommentCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialPostCommentCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialPostCommentCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParentCommentId

`func (o *SocialPostCommentCreateDto) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *SocialPostCommentCreateDto) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *SocialPostCommentCreateDto) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *SocialPostCommentCreateDto) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### SetParentCommentIdNil

`func (o *SocialPostCommentCreateDto) SetParentCommentIdNil(b bool)`

 SetParentCommentIdNil sets the value for ParentCommentId to be an explicit nil

### UnsetParentCommentId
`func (o *SocialPostCommentCreateDto) UnsetParentCommentId()`

UnsetParentCommentId ensures that no value is present for ParentCommentId, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostCommentCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostCommentCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostCommentCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostCommentCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostCommentCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostCommentCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialFeedPostId

`func (o *SocialPostCommentCreateDto) GetSocialFeedPostId() string`

GetSocialFeedPostId returns the SocialFeedPostId field if non-nil, zero value otherwise.

### GetSocialFeedPostIdOk

`func (o *SocialPostCommentCreateDto) GetSocialFeedPostIdOk() (*string, bool)`

GetSocialFeedPostIdOk returns a tuple with the SocialFeedPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedPostId

`func (o *SocialPostCommentCreateDto) SetSocialFeedPostId(v string)`

SetSocialFeedPostId sets SocialFeedPostId field to given value.

### HasSocialFeedPostId

`func (o *SocialPostCommentCreateDto) HasSocialFeedPostId() bool`

HasSocialFeedPostId returns a boolean if a field has been set.

### SetSocialFeedPostIdNil

`func (o *SocialPostCommentCreateDto) SetSocialFeedPostIdNil(b bool)`

 SetSocialFeedPostIdNil sets the value for SocialFeedPostId to be an explicit nil

### UnsetSocialFeedPostId
`func (o *SocialPostCommentCreateDto) UnsetSocialFeedPostId()`

UnsetSocialFeedPostId ensures that no value is present for SocialFeedPostId, not even an explicit nil
### GetSocialPostId

`func (o *SocialPostCommentCreateDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *SocialPostCommentCreateDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *SocialPostCommentCreateDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *SocialPostCommentCreateDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *SocialPostCommentCreateDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *SocialPostCommentCreateDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


