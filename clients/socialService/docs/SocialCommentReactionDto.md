# SocialCommentReactionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Reaction** | Pointer to **string** |  | [optional] 
**ReactionValue** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileName** | Pointer to **NullableString** |  | [optional] 
**SocialProfileAvatarUrl** | Pointer to **NullableString** |  | [optional] 
**SocialProfileType** | Pointer to **NullableString** |  | [optional] 
**SocialCommentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialCommentReactionDto

`func NewSocialCommentReactionDto() *SocialCommentReactionDto`

NewSocialCommentReactionDto instantiates a new SocialCommentReactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialCommentReactionDtoWithDefaults

`func NewSocialCommentReactionDtoWithDefaults() *SocialCommentReactionDto`

NewSocialCommentReactionDtoWithDefaults instantiates a new SocialCommentReactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialCommentReactionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialCommentReactionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialCommentReactionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialCommentReactionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialCommentReactionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialCommentReactionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialCommentReactionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialCommentReactionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialCommentReactionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialCommentReactionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialCommentReactionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialCommentReactionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetReaction

`func (o *SocialCommentReactionDto) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SocialCommentReactionDto) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SocialCommentReactionDto) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SocialCommentReactionDto) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetReactionValue

`func (o *SocialCommentReactionDto) GetReactionValue() string`

GetReactionValue returns the ReactionValue field if non-nil, zero value otherwise.

### GetReactionValueOk

`func (o *SocialCommentReactionDto) GetReactionValueOk() (*string, bool)`

GetReactionValueOk returns a tuple with the ReactionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionValue

`func (o *SocialCommentReactionDto) SetReactionValue(v string)`

SetReactionValue sets ReactionValue field to given value.

### HasReactionValue

`func (o *SocialCommentReactionDto) HasReactionValue() bool`

HasReactionValue returns a boolean if a field has been set.

### SetReactionValueNil

`func (o *SocialCommentReactionDto) SetReactionValueNil(b bool)`

 SetReactionValueNil sets the value for ReactionValue to be an explicit nil

### UnsetReactionValue
`func (o *SocialCommentReactionDto) UnsetReactionValue()`

UnsetReactionValue ensures that no value is present for ReactionValue, not even an explicit nil
### GetSocialProfileId

`func (o *SocialCommentReactionDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialCommentReactionDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialCommentReactionDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialCommentReactionDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialCommentReactionDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialCommentReactionDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialCommentReactionDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialCommentReactionDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialCommentReactionDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialCommentReactionDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialCommentReactionDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialCommentReactionDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialCommentReactionDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialCommentReactionDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialCommentReactionDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialCommentReactionDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialCommentReactionDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialCommentReactionDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetSocialProfileType

`func (o *SocialCommentReactionDto) GetSocialProfileType() string`

GetSocialProfileType returns the SocialProfileType field if non-nil, zero value otherwise.

### GetSocialProfileTypeOk

`func (o *SocialCommentReactionDto) GetSocialProfileTypeOk() (*string, bool)`

GetSocialProfileTypeOk returns a tuple with the SocialProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileType

`func (o *SocialCommentReactionDto) SetSocialProfileType(v string)`

SetSocialProfileType sets SocialProfileType field to given value.

### HasSocialProfileType

`func (o *SocialCommentReactionDto) HasSocialProfileType() bool`

HasSocialProfileType returns a boolean if a field has been set.

### SetSocialProfileTypeNil

`func (o *SocialCommentReactionDto) SetSocialProfileTypeNil(b bool)`

 SetSocialProfileTypeNil sets the value for SocialProfileType to be an explicit nil

### UnsetSocialProfileType
`func (o *SocialCommentReactionDto) UnsetSocialProfileType()`

UnsetSocialProfileType ensures that no value is present for SocialProfileType, not even an explicit nil
### GetSocialCommentId

`func (o *SocialCommentReactionDto) GetSocialCommentId() string`

GetSocialCommentId returns the SocialCommentId field if non-nil, zero value otherwise.

### GetSocialCommentIdOk

`func (o *SocialCommentReactionDto) GetSocialCommentIdOk() (*string, bool)`

GetSocialCommentIdOk returns a tuple with the SocialCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialCommentId

`func (o *SocialCommentReactionDto) SetSocialCommentId(v string)`

SetSocialCommentId sets SocialCommentId field to given value.

### HasSocialCommentId

`func (o *SocialCommentReactionDto) HasSocialCommentId() bool`

HasSocialCommentId returns a boolean if a field has been set.

### SetSocialCommentIdNil

`func (o *SocialCommentReactionDto) SetSocialCommentIdNil(b bool)`

 SetSocialCommentIdNil sets the value for SocialCommentId to be an explicit nil

### UnsetSocialCommentId
`func (o *SocialCommentReactionDto) UnsetSocialCommentId()`

UnsetSocialCommentId ensures that no value is present for SocialCommentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


