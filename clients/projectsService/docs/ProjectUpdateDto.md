# ProjectUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectStartDate** | Pointer to **time.Time** |  | [optional] 
**ProjectEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectUpdateDto

`func NewProjectUpdateDto() *ProjectUpdateDto`

NewProjectUpdateDto instantiates a new ProjectUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateDtoWithDefaults

`func NewProjectUpdateDtoWithDefaults() *ProjectUpdateDto`

NewProjectUpdateDtoWithDefaults instantiates a new ProjectUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ProjectUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProjectUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProjectUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProjectUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ProjectUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectStartDate

`func (o *ProjectUpdateDto) GetProjectStartDate() time.Time`

GetProjectStartDate returns the ProjectStartDate field if non-nil, zero value otherwise.

### GetProjectStartDateOk

`func (o *ProjectUpdateDto) GetProjectStartDateOk() (*time.Time, bool)`

GetProjectStartDateOk returns a tuple with the ProjectStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStartDate

`func (o *ProjectUpdateDto) SetProjectStartDate(v time.Time)`

SetProjectStartDate sets ProjectStartDate field to given value.

### HasProjectStartDate

`func (o *ProjectUpdateDto) HasProjectStartDate() bool`

HasProjectStartDate returns a boolean if a field has been set.

### GetProjectEndDate

`func (o *ProjectUpdateDto) GetProjectEndDate() time.Time`

GetProjectEndDate returns the ProjectEndDate field if non-nil, zero value otherwise.

### GetProjectEndDateOk

`func (o *ProjectUpdateDto) GetProjectEndDateOk() (*time.Time, bool)`

GetProjectEndDateOk returns a tuple with the ProjectEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEndDate

`func (o *ProjectUpdateDto) SetProjectEndDate(v time.Time)`

SetProjectEndDate sets ProjectEndDate field to given value.

### HasProjectEndDate

`func (o *ProjectUpdateDto) HasProjectEndDate() bool`

HasProjectEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


