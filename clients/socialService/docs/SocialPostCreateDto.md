# SocialPostCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostCreateDto

`func NewSocialPostCreateDto() *SocialPostCreateDto`

NewSocialPostCreateDto instantiates a new SocialPostCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostCreateDtoWithDefaults

`func NewSocialPostCreateDtoWithDefaults() *SocialPostCreateDto`

NewSocialPostCreateDtoWithDefaults instantiates a new SocialPostCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SocialPostCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SocialPostCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialPostCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialPostCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialPostCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialPostCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialPostCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *SocialPostCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialPostCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialPostCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SocialPostCreateDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SocialPostCreateDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SocialPostCreateDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSocialFeedId

`func (o *SocialPostCreateDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *SocialPostCreateDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *SocialPostCreateDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *SocialPostCreateDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *SocialPostCreateDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *SocialPostCreateDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


