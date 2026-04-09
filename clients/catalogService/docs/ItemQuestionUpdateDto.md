# ItemQuestionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**NeedsRevision** | **bool** |  | 
**Question** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemQuestionUpdateDto

`func NewItemQuestionUpdateDto(needsRevision bool, ) *ItemQuestionUpdateDto`

NewItemQuestionUpdateDto instantiates a new ItemQuestionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemQuestionUpdateDtoWithDefaults

`func NewItemQuestionUpdateDtoWithDefaults() *ItemQuestionUpdateDto`

NewItemQuestionUpdateDtoWithDefaults instantiates a new ItemQuestionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ItemQuestionUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemQuestionUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemQuestionUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemQuestionUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemQuestionUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemQuestionUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetNeedsRevision

`func (o *ItemQuestionUpdateDto) GetNeedsRevision() bool`

GetNeedsRevision returns the NeedsRevision field if non-nil, zero value otherwise.

### GetNeedsRevisionOk

`func (o *ItemQuestionUpdateDto) GetNeedsRevisionOk() (*bool, bool)`

GetNeedsRevisionOk returns a tuple with the NeedsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRevision

`func (o *ItemQuestionUpdateDto) SetNeedsRevision(v bool)`

SetNeedsRevision sets NeedsRevision field to given value.


### GetQuestion

`func (o *ItemQuestionUpdateDto) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ItemQuestionUpdateDto) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ItemQuestionUpdateDto) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ItemQuestionUpdateDto) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### SetQuestionNil

`func (o *ItemQuestionUpdateDto) SetQuestionNil(b bool)`

 SetQuestionNil sets the value for Question to be an explicit nil

### UnsetQuestion
`func (o *ItemQuestionUpdateDto) UnsetQuestion()`

UnsetQuestion ensures that no value is present for Question, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


