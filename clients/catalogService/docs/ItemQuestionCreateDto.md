# ItemQuestionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**NeedsRevision** | **bool** |  | 
**Question** | **string** |  | 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | **string** |  | 

## Methods

### NewItemQuestionCreateDto

`func NewItemQuestionCreateDto(title string, needsRevision bool, question string, itemId string, ) *ItemQuestionCreateDto`

NewItemQuestionCreateDto instantiates a new ItemQuestionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemQuestionCreateDtoWithDefaults

`func NewItemQuestionCreateDtoWithDefaults() *ItemQuestionCreateDto`

NewItemQuestionCreateDtoWithDefaults instantiates a new ItemQuestionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemQuestionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemQuestionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemQuestionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemQuestionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemQuestionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemQuestionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemQuestionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemQuestionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ItemQuestionCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemQuestionCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemQuestionCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetNeedsRevision

`func (o *ItemQuestionCreateDto) GetNeedsRevision() bool`

GetNeedsRevision returns the NeedsRevision field if non-nil, zero value otherwise.

### GetNeedsRevisionOk

`func (o *ItemQuestionCreateDto) GetNeedsRevisionOk() (*bool, bool)`

GetNeedsRevisionOk returns a tuple with the NeedsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRevision

`func (o *ItemQuestionCreateDto) SetNeedsRevision(v bool)`

SetNeedsRevision sets NeedsRevision field to given value.


### GetQuestion

`func (o *ItemQuestionCreateDto) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ItemQuestionCreateDto) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ItemQuestionCreateDto) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetSocialProfileId

`func (o *ItemQuestionCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ItemQuestionCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ItemQuestionCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ItemQuestionCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ItemQuestionCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ItemQuestionCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetItemId

`func (o *ItemQuestionCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemQuestionCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemQuestionCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


