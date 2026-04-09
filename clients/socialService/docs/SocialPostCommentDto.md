# SocialPostCommentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ParentCommentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialFeedPostId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileName** | Pointer to **NullableString** |  | [optional] 
**SocialProfileAvatarUrl** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostCommentDto

`func NewSocialPostCommentDto() *SocialPostCommentDto`

NewSocialPostCommentDto instantiates a new SocialPostCommentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostCommentDtoWithDefaults

`func NewSocialPostCommentDtoWithDefaults() *SocialPostCommentDto`

NewSocialPostCommentDtoWithDefaults instantiates a new SocialPostCommentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostCommentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostCommentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostCommentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostCommentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialPostCommentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialPostCommentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialPostCommentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostCommentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostCommentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostCommentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialPostCommentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialPostCommentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetMessage

`func (o *SocialPostCommentDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialPostCommentDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialPostCommentDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SocialPostCommentDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SocialPostCommentDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SocialPostCommentDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetParentCommentId

`func (o *SocialPostCommentDto) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *SocialPostCommentDto) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *SocialPostCommentDto) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *SocialPostCommentDto) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### SetParentCommentIdNil

`func (o *SocialPostCommentDto) SetParentCommentIdNil(b bool)`

 SetParentCommentIdNil sets the value for ParentCommentId to be an explicit nil

### UnsetParentCommentId
`func (o *SocialPostCommentDto) UnsetParentCommentId()`

UnsetParentCommentId ensures that no value is present for ParentCommentId, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostCommentDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostCommentDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostCommentDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostCommentDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostCommentDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostCommentDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialFeedPostId

`func (o *SocialPostCommentDto) GetSocialFeedPostId() string`

GetSocialFeedPostId returns the SocialFeedPostId field if non-nil, zero value otherwise.

### GetSocialFeedPostIdOk

`func (o *SocialPostCommentDto) GetSocialFeedPostIdOk() (*string, bool)`

GetSocialFeedPostIdOk returns a tuple with the SocialFeedPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedPostId

`func (o *SocialPostCommentDto) SetSocialFeedPostId(v string)`

SetSocialFeedPostId sets SocialFeedPostId field to given value.

### HasSocialFeedPostId

`func (o *SocialPostCommentDto) HasSocialFeedPostId() bool`

HasSocialFeedPostId returns a boolean if a field has been set.

### SetSocialFeedPostIdNil

`func (o *SocialPostCommentDto) SetSocialFeedPostIdNil(b bool)`

 SetSocialFeedPostIdNil sets the value for SocialFeedPostId to be an explicit nil

### UnsetSocialFeedPostId
`func (o *SocialPostCommentDto) UnsetSocialFeedPostId()`

UnsetSocialFeedPostId ensures that no value is present for SocialFeedPostId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialPostCommentDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialPostCommentDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialPostCommentDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialPostCommentDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialPostCommentDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialPostCommentDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialPostCommentDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialPostCommentDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialPostCommentDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialPostCommentDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialPostCommentDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialPostCommentDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetSocialPostId

`func (o *SocialPostCommentDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *SocialPostCommentDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *SocialPostCommentDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *SocialPostCommentDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *SocialPostCommentDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *SocialPostCommentDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


