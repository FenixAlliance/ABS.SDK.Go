# ItemQuestionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**NeedsRevision** | **bool** |  | 
**Question** | **string** |  | 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**ItemID** | **string** |  | 

## Methods

### NewItemQuestionCreateDto

`func NewItemQuestionCreateDto(title string, needsRevision bool, question string, itemID string, ) *ItemQuestionCreateDto`

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


### GetSocialProfileID

`func (o *ItemQuestionCreateDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemQuestionCreateDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemQuestionCreateDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemQuestionCreateDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemQuestionCreateDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemQuestionCreateDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetItemID

`func (o *ItemQuestionCreateDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemQuestionCreateDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemQuestionCreateDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


