# SocialPostDto

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

## Methods

### NewSocialPostDto

`func NewSocialPostDto() *SocialPostDto`

NewSocialPostDto instantiates a new SocialPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostDtoWithDefaults

`func NewSocialPostDtoWithDefaults() *SocialPostDto`

NewSocialPostDtoWithDefaults instantiates a new SocialPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialPostDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialPostDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialPostDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialPostDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialPostDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SocialPostDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialPostDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialPostDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialPostDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialPostDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialPostDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *SocialPostDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialPostDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialPostDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SocialPostDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SocialPostDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SocialPostDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialPostDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialPostDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialPostDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialPostDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialPostDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialPostDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialPostDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialPostDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialPostDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialPostDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialPostDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialPostDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetCommentsCount

`func (o *SocialPostDto) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *SocialPostDto) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *SocialPostDto) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *SocialPostDto) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetReactionsCount

`func (o *SocialPostDto) GetReactionsCount() int32`

GetReactionsCount returns the ReactionsCount field if non-nil, zero value otherwise.

### GetReactionsCountOk

`func (o *SocialPostDto) GetReactionsCountOk() (*int32, bool)`

GetReactionsCountOk returns a tuple with the ReactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionsCount

`func (o *SocialPostDto) SetReactionsCount(v int32)`

SetReactionsCount sets ReactionsCount field to given value.

### HasReactionsCount

`func (o *SocialPostDto) HasReactionsCount() bool`

HasReactionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


