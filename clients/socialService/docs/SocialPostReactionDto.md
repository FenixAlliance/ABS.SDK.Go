# SocialPostReactionDto

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
**SocialPostId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostReactionDto

`func NewSocialPostReactionDto() *SocialPostReactionDto`

NewSocialPostReactionDto instantiates a new SocialPostReactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostReactionDtoWithDefaults

`func NewSocialPostReactionDtoWithDefaults() *SocialPostReactionDto`

NewSocialPostReactionDtoWithDefaults instantiates a new SocialPostReactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostReactionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostReactionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostReactionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostReactionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialPostReactionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialPostReactionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialPostReactionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostReactionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostReactionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostReactionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialPostReactionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialPostReactionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetReaction

`func (o *SocialPostReactionDto) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SocialPostReactionDto) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SocialPostReactionDto) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SocialPostReactionDto) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetReactionValue

`func (o *SocialPostReactionDto) GetReactionValue() string`

GetReactionValue returns the ReactionValue field if non-nil, zero value otherwise.

### GetReactionValueOk

`func (o *SocialPostReactionDto) GetReactionValueOk() (*string, bool)`

GetReactionValueOk returns a tuple with the ReactionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionValue

`func (o *SocialPostReactionDto) SetReactionValue(v string)`

SetReactionValue sets ReactionValue field to given value.

### HasReactionValue

`func (o *SocialPostReactionDto) HasReactionValue() bool`

HasReactionValue returns a boolean if a field has been set.

### SetReactionValueNil

`func (o *SocialPostReactionDto) SetReactionValueNil(b bool)`

 SetReactionValueNil sets the value for ReactionValue to be an explicit nil

### UnsetReactionValue
`func (o *SocialPostReactionDto) UnsetReactionValue()`

UnsetReactionValue ensures that no value is present for ReactionValue, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostReactionDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostReactionDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostReactionDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostReactionDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostReactionDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostReactionDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialPostReactionDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialPostReactionDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialPostReactionDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialPostReactionDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialPostReactionDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialPostReactionDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialPostReactionDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialPostReactionDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialPostReactionDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialPostReactionDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialPostReactionDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialPostReactionDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetSocialProfileType

`func (o *SocialPostReactionDto) GetSocialProfileType() string`

GetSocialProfileType returns the SocialProfileType field if non-nil, zero value otherwise.

### GetSocialProfileTypeOk

`func (o *SocialPostReactionDto) GetSocialProfileTypeOk() (*string, bool)`

GetSocialProfileTypeOk returns a tuple with the SocialProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileType

`func (o *SocialPostReactionDto) SetSocialProfileType(v string)`

SetSocialProfileType sets SocialProfileType field to given value.

### HasSocialProfileType

`func (o *SocialPostReactionDto) HasSocialProfileType() bool`

HasSocialProfileType returns a boolean if a field has been set.

### SetSocialProfileTypeNil

`func (o *SocialPostReactionDto) SetSocialProfileTypeNil(b bool)`

 SetSocialProfileTypeNil sets the value for SocialProfileType to be an explicit nil

### UnsetSocialProfileType
`func (o *SocialPostReactionDto) UnsetSocialProfileType()`

UnsetSocialProfileType ensures that no value is present for SocialProfileType, not even an explicit nil
### GetSocialPostId

`func (o *SocialPostReactionDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *SocialPostReactionDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *SocialPostReactionDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *SocialPostReactionDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *SocialPostReactionDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *SocialPostReactionDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


