# ProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectStartDate** | Pointer to **time.Time** |  | [optional] 
**ProjectEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectDto

`func NewProjectDto() *ProjectDto`

NewProjectDto instantiates a new ProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDtoWithDefaults

`func NewProjectDtoWithDefaults() *ProjectDto`

NewProjectDtoWithDefaults instantiates a new ProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProjectDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProjectDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProjectDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ProjectDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProjectDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProjectDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProjectDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ProjectDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectStartDate

`func (o *ProjectDto) GetProjectStartDate() time.Time`

GetProjectStartDate returns the ProjectStartDate field if non-nil, zero value otherwise.

### GetProjectStartDateOk

`func (o *ProjectDto) GetProjectStartDateOk() (*time.Time, bool)`

GetProjectStartDateOk returns a tuple with the ProjectStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStartDate

`func (o *ProjectDto) SetProjectStartDate(v time.Time)`

SetProjectStartDate sets ProjectStartDate field to given value.

### HasProjectStartDate

`func (o *ProjectDto) HasProjectStartDate() bool`

HasProjectStartDate returns a boolean if a field has been set.

### GetProjectEndDate

`func (o *ProjectDto) GetProjectEndDate() time.Time`

GetProjectEndDate returns the ProjectEndDate field if non-nil, zero value otherwise.

### GetProjectEndDateOk

`func (o *ProjectDto) GetProjectEndDateOk() (*time.Time, bool)`

GetProjectEndDateOk returns a tuple with the ProjectEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEndDate

`func (o *ProjectDto) SetProjectEndDate(v time.Time)`

SetProjectEndDate sets ProjectEndDate field to given value.

### HasProjectEndDate

`func (o *ProjectDto) HasProjectEndDate() bool`

HasProjectEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


