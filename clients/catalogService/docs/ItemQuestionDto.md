# ItemQuestionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**NeedsRevision** | Pointer to **bool** |  | [optional] 
**Question** | Pointer to **NullableString** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**ItemID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemQuestionDto

`func NewItemQuestionDto() *ItemQuestionDto`

NewItemQuestionDto instantiates a new ItemQuestionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemQuestionDtoWithDefaults

`func NewItemQuestionDtoWithDefaults() *ItemQuestionDto`

NewItemQuestionDtoWithDefaults instantiates a new ItemQuestionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemQuestionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemQuestionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemQuestionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemQuestionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemQuestionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemQuestionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemQuestionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemQuestionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemQuestionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemQuestionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemQuestionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemQuestionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ItemQuestionDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemQuestionDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemQuestionDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemQuestionDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemQuestionDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemQuestionDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetNeedsRevision

`func (o *ItemQuestionDto) GetNeedsRevision() bool`

GetNeedsRevision returns the NeedsRevision field if non-nil, zero value otherwise.

### GetNeedsRevisionOk

`func (o *ItemQuestionDto) GetNeedsRevisionOk() (*bool, bool)`

GetNeedsRevisionOk returns a tuple with the NeedsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRevision

`func (o *ItemQuestionDto) SetNeedsRevision(v bool)`

SetNeedsRevision sets NeedsRevision field to given value.

### HasNeedsRevision

`func (o *ItemQuestionDto) HasNeedsRevision() bool`

HasNeedsRevision returns a boolean if a field has been set.

### GetQuestion

`func (o *ItemQuestionDto) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ItemQuestionDto) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ItemQuestionDto) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ItemQuestionDto) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### SetQuestionNil

`func (o *ItemQuestionDto) SetQuestionNil(b bool)`

 SetQuestionNil sets the value for Question to be an explicit nil

### UnsetQuestion
`func (o *ItemQuestionDto) UnsetQuestion()`

UnsetQuestion ensures that no value is present for Question, not even an explicit nil
### GetSocialProfileID

`func (o *ItemQuestionDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemQuestionDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemQuestionDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemQuestionDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemQuestionDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemQuestionDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetBusinessID

`func (o *ItemQuestionDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemQuestionDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemQuestionDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ItemQuestionDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ItemQuestionDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ItemQuestionDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetItemID

`func (o *ItemQuestionDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemQuestionDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemQuestionDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemQuestionDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemQuestionDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemQuestionDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


