# ConversationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**StartedTimestamp** | Pointer to **time.Time** |  | [optional] 
**LastMessageTimestamp** | Pointer to **time.Time** |  | [optional] 
**LastMessage** | Pointer to **NullableString** |  | [optional] 
**SocialProfileName** | Pointer to **NullableString** |  | [optional] 
**SocialProfileAvatarUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConversationDto

`func NewConversationDto() *ConversationDto`

NewConversationDto instantiates a new ConversationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationDtoWithDefaults

`func NewConversationDtoWithDefaults() *ConversationDto`

NewConversationDtoWithDefaults instantiates a new ConversationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConversationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConversationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConversationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConversationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ConversationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConversationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConversationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConversationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ConversationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ConversationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSubject

`func (o *ConversationDto) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ConversationDto) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ConversationDto) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ConversationDto) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ConversationDto) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ConversationDto) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSocialProfileId

`func (o *ConversationDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ConversationDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ConversationDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ConversationDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ConversationDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ConversationDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetStartedTimestamp

`func (o *ConversationDto) GetStartedTimestamp() time.Time`

GetStartedTimestamp returns the StartedTimestamp field if non-nil, zero value otherwise.

### GetStartedTimestampOk

`func (o *ConversationDto) GetStartedTimestampOk() (*time.Time, bool)`

GetStartedTimestampOk returns a tuple with the StartedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTimestamp

`func (o *ConversationDto) SetStartedTimestamp(v time.Time)`

SetStartedTimestamp sets StartedTimestamp field to given value.

### HasStartedTimestamp

`func (o *ConversationDto) HasStartedTimestamp() bool`

HasStartedTimestamp returns a boolean if a field has been set.

### GetLastMessageTimestamp

`func (o *ConversationDto) GetLastMessageTimestamp() time.Time`

GetLastMessageTimestamp returns the LastMessageTimestamp field if non-nil, zero value otherwise.

### GetLastMessageTimestampOk

`func (o *ConversationDto) GetLastMessageTimestampOk() (*time.Time, bool)`

GetLastMessageTimestampOk returns a tuple with the LastMessageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageTimestamp

`func (o *ConversationDto) SetLastMessageTimestamp(v time.Time)`

SetLastMessageTimestamp sets LastMessageTimestamp field to given value.

### HasLastMessageTimestamp

`func (o *ConversationDto) HasLastMessageTimestamp() bool`

HasLastMessageTimestamp returns a boolean if a field has been set.

### GetLastMessage

`func (o *ConversationDto) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ConversationDto) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ConversationDto) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ConversationDto) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *ConversationDto) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *ConversationDto) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetSocialProfileName

`func (o *ConversationDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *ConversationDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *ConversationDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *ConversationDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *ConversationDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *ConversationDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *ConversationDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *ConversationDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *ConversationDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *ConversationDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *ConversationDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *ConversationDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


