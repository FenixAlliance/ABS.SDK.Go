# SocialReactionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Reaction** | Pointer to **string** |  | [optional] 
**ReactionValue** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialReactionCreateDto

`func NewSocialReactionCreateDto() *SocialReactionCreateDto`

NewSocialReactionCreateDto instantiates a new SocialReactionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialReactionCreateDtoWithDefaults

`func NewSocialReactionCreateDtoWithDefaults() *SocialReactionCreateDto`

NewSocialReactionCreateDtoWithDefaults instantiates a new SocialReactionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialReactionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialReactionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialReactionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialReactionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SocialReactionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialReactionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialReactionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialReactionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReaction

`func (o *SocialReactionCreateDto) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SocialReactionCreateDto) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SocialReactionCreateDto) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SocialReactionCreateDto) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetReactionValue

`func (o *SocialReactionCreateDto) GetReactionValue() string`

GetReactionValue returns the ReactionValue field if non-nil, zero value otherwise.

### GetReactionValueOk

`func (o *SocialReactionCreateDto) GetReactionValueOk() (*string, bool)`

GetReactionValueOk returns a tuple with the ReactionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionValue

`func (o *SocialReactionCreateDto) SetReactionValue(v string)`

SetReactionValue sets ReactionValue field to given value.

### HasReactionValue

`func (o *SocialReactionCreateDto) HasReactionValue() bool`

HasReactionValue returns a boolean if a field has been set.

### SetReactionValueNil

`func (o *SocialReactionCreateDto) SetReactionValueNil(b bool)`

 SetReactionValueNil sets the value for ReactionValue to be an explicit nil

### UnsetReactionValue
`func (o *SocialReactionCreateDto) UnsetReactionValue()`

UnsetReactionValue ensures that no value is present for ReactionValue, not even an explicit nil
### GetSocialProfileId

`func (o *SocialReactionCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialReactionCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialReactionCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialReactionCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialReactionCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialReactionCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


