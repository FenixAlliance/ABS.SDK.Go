# ProjectCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectStartDate** | Pointer to **time.Time** |  | [optional] 
**ProjectEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectCreateDto

`func NewProjectCreateDto() *ProjectCreateDto`

NewProjectCreateDto instantiates a new ProjectCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateDtoWithDefaults

`func NewProjectCreateDtoWithDefaults() *ProjectCreateDto`

NewProjectCreateDtoWithDefaults instantiates a new ProjectCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProjectCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ProjectCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProjectCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProjectCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProjectCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ProjectCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectStartDate

`func (o *ProjectCreateDto) GetProjectStartDate() time.Time`

GetProjectStartDate returns the ProjectStartDate field if non-nil, zero value otherwise.

### GetProjectStartDateOk

`func (o *ProjectCreateDto) GetProjectStartDateOk() (*time.Time, bool)`

GetProjectStartDateOk returns a tuple with the ProjectStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStartDate

`func (o *ProjectCreateDto) SetProjectStartDate(v time.Time)`

SetProjectStartDate sets ProjectStartDate field to given value.

### HasProjectStartDate

`func (o *ProjectCreateDto) HasProjectStartDate() bool`

HasProjectStartDate returns a boolean if a field has been set.

### GetProjectEndDate

`func (o *ProjectCreateDto) GetProjectEndDate() time.Time`

GetProjectEndDate returns the ProjectEndDate field if non-nil, zero value otherwise.

### GetProjectEndDateOk

`func (o *ProjectCreateDto) GetProjectEndDateOk() (*time.Time, bool)`

GetProjectEndDateOk returns a tuple with the ProjectEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEndDate

`func (o *ProjectCreateDto) SetProjectEndDate(v time.Time)`

SetProjectEndDate sets ProjectEndDate field to given value.

### HasProjectEndDate

`func (o *ProjectCreateDto) HasProjectEndDate() bool`

HasProjectEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


