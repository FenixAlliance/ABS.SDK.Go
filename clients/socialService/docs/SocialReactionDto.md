# SocialReactionDto

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

## Methods

### NewSocialReactionDto

`func NewSocialReactionDto() *SocialReactionDto`

NewSocialReactionDto instantiates a new SocialReactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialReactionDtoWithDefaults

`func NewSocialReactionDtoWithDefaults() *SocialReactionDto`

NewSocialReactionDtoWithDefaults instantiates a new SocialReactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialReactionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialReactionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialReactionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialReactionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialReactionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialReactionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialReactionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialReactionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialReactionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialReactionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialReactionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialReactionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetReaction

`func (o *SocialReactionDto) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SocialReactionDto) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SocialReactionDto) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SocialReactionDto) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetReactionValue

`func (o *SocialReactionDto) GetReactionValue() string`

GetReactionValue returns the ReactionValue field if non-nil, zero value otherwise.

### GetReactionValueOk

`func (o *SocialReactionDto) GetReactionValueOk() (*string, bool)`

GetReactionValueOk returns a tuple with the ReactionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionValue

`func (o *SocialReactionDto) SetReactionValue(v string)`

SetReactionValue sets ReactionValue field to given value.

### HasReactionValue

`func (o *SocialReactionDto) HasReactionValue() bool`

HasReactionValue returns a boolean if a field has been set.

### SetReactionValueNil

`func (o *SocialReactionDto) SetReactionValueNil(b bool)`

 SetReactionValueNil sets the value for ReactionValue to be an explicit nil

### UnsetReactionValue
`func (o *SocialReactionDto) UnsetReactionValue()`

UnsetReactionValue ensures that no value is present for ReactionValue, not even an explicit nil
### GetSocialProfileId

`func (o *SocialReactionDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialReactionDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialReactionDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialReactionDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialReactionDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialReactionDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialReactionDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialReactionDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialReactionDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialReactionDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialReactionDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialReactionDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialReactionDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialReactionDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialReactionDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialReactionDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialReactionDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialReactionDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


