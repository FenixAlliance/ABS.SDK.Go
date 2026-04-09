# SocialFeedPostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileName** | Pointer to **NullableString** |  | [optional] 
**SocialProfileAvatarUrl** | Pointer to **NullableString** |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**ReactionsCount** | Pointer to **int32** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialFeedPostDto

`func NewSocialFeedPostDto() *SocialFeedPostDto`

NewSocialFeedPostDto instantiates a new SocialFeedPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialFeedPostDtoWithDefaults

`func NewSocialFeedPostDtoWithDefaults() *SocialFeedPostDto`

NewSocialFeedPostDtoWithDefaults instantiates a new SocialFeedPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialFeedPostDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialFeedPostDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialFeedPostDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialFeedPostDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialFeedPostDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialFeedPostDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialFeedPostDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialFeedPostDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialFeedPostDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialFeedPostDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialFeedPostDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialFeedPostDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SocialFeedPostDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialFeedPostDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialFeedPostDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialFeedPostDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialFeedPostDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialFeedPostDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *SocialFeedPostDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialFeedPostDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialFeedPostDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SocialFeedPostDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SocialFeedPostDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SocialFeedPostDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSocialProfileId

`func (o *SocialFeedPostDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialFeedPostDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialFeedPostDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialFeedPostDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialFeedPostDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialFeedPostDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialFeedPostDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialFeedPostDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialFeedPostDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialFeedPostDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialFeedPostDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialFeedPostDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialFeedPostDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialFeedPostDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialFeedPostDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetCommentsCount

`func (o *SocialFeedPostDto) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *SocialFeedPostDto) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *SocialFeedPostDto) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *SocialFeedPostDto) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetReactionsCount

`func (o *SocialFeedPostDto) GetReactionsCount() int32`

GetReactionsCount returns the ReactionsCount field if non-nil, zero value otherwise.

### GetReactionsCountOk

`func (o *SocialFeedPostDto) GetReactionsCountOk() (*int32, bool)`

GetReactionsCountOk returns a tuple with the ReactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionsCount

`func (o *SocialFeedPostDto) SetReactionsCount(v int32)`

SetReactionsCount sets ReactionsCount field to given value.

### HasReactionsCount

`func (o *SocialFeedPostDto) HasReactionsCount() bool`

HasReactionsCount returns a boolean if a field has been set.

### GetSocialFeedId

`func (o *SocialFeedPostDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *SocialFeedPostDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *SocialFeedPostDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *SocialFeedPostDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *SocialFeedPostDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *SocialFeedPostDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


