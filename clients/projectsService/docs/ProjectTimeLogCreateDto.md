# ProjectTimeLogCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TimeSpan** | Pointer to **string** |  | [optional] 
**LogDate** | Pointer to **time.Time** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskId** | **string** |  | 
**ProjectPeriodId** | **string** |  | 
**ProjectTimeLogRecordType** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectTimeLogCreateDto

`func NewProjectTimeLogCreateDto(projectTaskId string, projectPeriodId string, ) *ProjectTimeLogCreateDto`

NewProjectTimeLogCreateDto instantiates a new ProjectTimeLogCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTimeLogCreateDtoWithDefaults

`func NewProjectTimeLogCreateDtoWithDefaults() *ProjectTimeLogCreateDto`

NewProjectTimeLogCreateDtoWithDefaults instantiates a new ProjectTimeLogCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTimeLogCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTimeLogCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTimeLogCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTimeLogCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProjectTimeLogCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectTimeLogCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectTimeLogCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectTimeLogCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimeSpan

`func (o *ProjectTimeLogCreateDto) GetTimeSpan() string`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *ProjectTimeLogCreateDto) GetTimeSpanOk() (*string, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *ProjectTimeLogCreateDto) SetTimeSpan(v string)`

SetTimeSpan sets TimeSpan field to given value.

### HasTimeSpan

`func (o *ProjectTimeLogCreateDto) HasTimeSpan() bool`

HasTimeSpan returns a boolean if a field has been set.

### GetLogDate

`func (o *ProjectTimeLogCreateDto) GetLogDate() time.Time`

GetLogDate returns the LogDate field if non-nil, zero value otherwise.

### GetLogDateOk

`func (o *ProjectTimeLogCreateDto) GetLogDateOk() (*time.Time, bool)`

GetLogDateOk returns a tuple with the LogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDate

`func (o *ProjectTimeLogCreateDto) SetLogDate(v time.Time)`

SetLogDate sets LogDate field to given value.

### HasLogDate

`func (o *ProjectTimeLogCreateDto) HasLogDate() bool`

HasLogDate returns a boolean if a field has been set.

### GetComments

`func (o *ProjectTimeLogCreateDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProjectTimeLogCreateDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProjectTimeLogCreateDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProjectTimeLogCreateDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ProjectTimeLogCreateDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ProjectTimeLogCreateDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetProjectTaskId

`func (o *ProjectTimeLogCreateDto) GetProjectTaskId() string`

GetProjectTaskId returns the ProjectTaskId field if non-nil, zero value otherwise.

### GetProjectTaskIdOk

`func (o *ProjectTimeLogCreateDto) GetProjectTaskIdOk() (*string, bool)`

GetProjectTaskIdOk returns a tuple with the ProjectTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskId

`func (o *ProjectTimeLogCreateDto) SetProjectTaskId(v string)`

SetProjectTaskId sets ProjectTaskId field to given value.


### GetProjectPeriodId

`func (o *ProjectTimeLogCreateDto) GetProjectPeriodId() string`

GetProjectPeriodId returns the ProjectPeriodId field if non-nil, zero value otherwise.

### GetProjectPeriodIdOk

`func (o *ProjectTimeLogCreateDto) GetProjectPeriodIdOk() (*string, bool)`

GetProjectPeriodIdOk returns a tuple with the ProjectPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodId

`func (o *ProjectTimeLogCreateDto) SetProjectPeriodId(v string)`

SetProjectPeriodId sets ProjectPeriodId field to given value.


### GetProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordType() string`

GetProjectTimeLogRecordType returns the ProjectTimeLogRecordType field if non-nil, zero value otherwise.

### GetProjectTimeLogRecordTypeOk

`func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordTypeOk() (*string, bool)`

GetProjectTimeLogRecordTypeOk returns a tuple with the ProjectTimeLogRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) SetProjectTimeLogRecordType(v string)`

SetProjectTimeLogRecordType sets ProjectTimeLogRecordType field to given value.

### HasProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) HasProjectTimeLogRecordType() bool`

HasProjectTimeLogRecordType returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectTimeLogCreateDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectTimeLogCreateDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectTimeLogCreateDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectTimeLogCreateDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectTimeLogCreateDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectTimeLogCreateDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


