# ProjectTaskCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**DueLine** | Pointer to **time.Time** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectTaskCreateDto

`func NewProjectTaskCreateDto() *ProjectTaskCreateDto`

NewProjectTaskCreateDto instantiates a new ProjectTaskCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTaskCreateDtoWithDefaults

`func NewProjectTaskCreateDtoWithDefaults() *ProjectTaskCreateDto`

NewProjectTaskCreateDtoWithDefaults instantiates a new ProjectTaskCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTaskCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTaskCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTaskCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTaskCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProjectTaskCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectTaskCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectTaskCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectTaskCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ProjectTaskCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectTaskCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectTaskCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProjectTaskCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProjectTaskCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProjectTaskCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ProjectTaskCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTaskCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTaskCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTaskCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectTaskCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectTaskCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDate

`func (o *ProjectTaskCreateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectTaskCreateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectTaskCreateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectTaskCreateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueLine

`func (o *ProjectTaskCreateDto) GetDueLine() time.Time`

GetDueLine returns the DueLine field if non-nil, zero value otherwise.

### GetDueLineOk

`func (o *ProjectTaskCreateDto) GetDueLineOk() (*time.Time, bool)`

GetDueLineOk returns a tuple with the DueLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueLine

`func (o *ProjectTaskCreateDto) SetDueLine(v time.Time)`

SetDueLine sets DueLine field to given value.

### HasDueLine

`func (o *ProjectTaskCreateDto) HasDueLine() bool`

HasDueLine returns a boolean if a field has been set.

### GetProjectID

`func (o *ProjectTaskCreateDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *ProjectTaskCreateDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *ProjectTaskCreateDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *ProjectTaskCreateDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *ProjectTaskCreateDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *ProjectTaskCreateDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


