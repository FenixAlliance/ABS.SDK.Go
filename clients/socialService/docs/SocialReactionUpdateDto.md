# SocialReactionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Reaction** | Pointer to **string** |  | [optional] 
**ReactionValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialReactionUpdateDto

`func NewSocialReactionUpdateDto() *SocialReactionUpdateDto`

NewSocialReactionUpdateDto instantiates a new SocialReactionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialReactionUpdateDtoWithDefaults

`func NewSocialReactionUpdateDtoWithDefaults() *SocialReactionUpdateDto`

NewSocialReactionUpdateDtoWithDefaults instantiates a new SocialReactionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialReactionUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialReactionUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialReactionUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialReactionUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialReactionUpdateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialReactionUpdateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialReactionUpdateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialReactionUpdateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialReactionUpdateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialReactionUpdateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialReactionUpdateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialReactionUpdateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetReaction

`func (o *SocialReactionUpdateDto) GetReaction() string`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SocialReactionUpdateDto) GetReactionOk() (*string, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SocialReactionUpdateDto) SetReaction(v string)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SocialReactionUpdateDto) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetReactionValue

`func (o *SocialReactionUpdateDto) GetReactionValue() string`

GetReactionValue returns the ReactionValue field if non-nil, zero value otherwise.

### GetReactionValueOk

`func (o *SocialReactionUpdateDto) GetReactionValueOk() (*string, bool)`

GetReactionValueOk returns a tuple with the ReactionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionValue

`func (o *SocialReactionUpdateDto) SetReactionValue(v string)`

SetReactionValue sets ReactionValue field to given value.

### HasReactionValue

`func (o *SocialReactionUpdateDto) HasReactionValue() bool`

HasReactionValue returns a boolean if a field has been set.

### SetReactionValueNil

`func (o *SocialReactionUpdateDto) SetReactionValueNil(b bool)`

 SetReactionValueNil sets the value for ReactionValue to be an explicit nil

### UnsetReactionValue
`func (o *SocialReactionUpdateDto) UnsetReactionValue()`

UnsetReactionValue ensures that no value is present for ReactionValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


