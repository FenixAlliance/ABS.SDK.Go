# SocialFeedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SocialPostsCount** | Pointer to **int32** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialFeedDto

`func NewSocialFeedDto() *SocialFeedDto`

NewSocialFeedDto instantiates a new SocialFeedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialFeedDtoWithDefaults

`func NewSocialFeedDtoWithDefaults() *SocialFeedDto`

NewSocialFeedDtoWithDefaults instantiates a new SocialFeedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialFeedDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialFeedDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialFeedDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialFeedDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialFeedDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialFeedDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialFeedDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialFeedDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialFeedDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialFeedDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialFeedDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialFeedDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSocialPostsCount

`func (o *SocialFeedDto) GetSocialPostsCount() int32`

GetSocialPostsCount returns the SocialPostsCount field if non-nil, zero value otherwise.

### GetSocialPostsCountOk

`func (o *SocialFeedDto) GetSocialPostsCountOk() (*int32, bool)`

GetSocialPostsCountOk returns a tuple with the SocialPostsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostsCount

`func (o *SocialFeedDto) SetSocialPostsCount(v int32)`

SetSocialPostsCount sets SocialPostsCount field to given value.

### HasSocialPostsCount

`func (o *SocialFeedDto) HasSocialPostsCount() bool`

HasSocialPostsCount returns a boolean if a field has been set.

### GetSocialProfileId

`func (o *SocialFeedDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialFeedDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialFeedDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialFeedDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialFeedDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialFeedDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


