# ProjectTimeLogUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDate** | Pointer to **time.Time** |  | [optional] 
**TimeSpan** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskId** | Pointer to **NullableString** |  | [optional] 
**ProjectPeriodId** | Pointer to **NullableString** |  | [optional] 
**ProjectTimeLogRecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectTimeLogUpdateDto

`func NewProjectTimeLogUpdateDto() *ProjectTimeLogUpdateDto`

NewProjectTimeLogUpdateDto instantiates a new ProjectTimeLogUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTimeLogUpdateDtoWithDefaults

`func NewProjectTimeLogUpdateDtoWithDefaults() *ProjectTimeLogUpdateDto`

NewProjectTimeLogUpdateDtoWithDefaults instantiates a new ProjectTimeLogUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDate

`func (o *ProjectTimeLogUpdateDto) GetLogDate() time.Time`

GetLogDate returns the LogDate field if non-nil, zero value otherwise.

### GetLogDateOk

`func (o *ProjectTimeLogUpdateDto) GetLogDateOk() (*time.Time, bool)`

GetLogDateOk returns a tuple with the LogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDate

`func (o *ProjectTimeLogUpdateDto) SetLogDate(v time.Time)`

SetLogDate sets LogDate field to given value.

### HasLogDate

`func (o *ProjectTimeLogUpdateDto) HasLogDate() bool`

HasLogDate returns a boolean if a field has been set.

### GetTimeSpan

`func (o *ProjectTimeLogUpdateDto) GetTimeSpan() string`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *ProjectTimeLogUpdateDto) GetTimeSpanOk() (*string, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *ProjectTimeLogUpdateDto) SetTimeSpan(v string)`

SetTimeSpan sets TimeSpan field to given value.

### HasTimeSpan

`func (o *ProjectTimeLogUpdateDto) HasTimeSpan() bool`

HasTimeSpan returns a boolean if a field has been set.

### GetComments

`func (o *ProjectTimeLogUpdateDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProjectTimeLogUpdateDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProjectTimeLogUpdateDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProjectTimeLogUpdateDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ProjectTimeLogUpdateDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ProjectTimeLogUpdateDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetProjectTaskId

`func (o *ProjectTimeLogUpdateDto) GetProjectTaskId() string`

GetProjectTaskId returns the ProjectTaskId field if non-nil, zero value otherwise.

### GetProjectTaskIdOk

`func (o *ProjectTimeLogUpdateDto) GetProjectTaskIdOk() (*string, bool)`

GetProjectTaskIdOk returns a tuple with the ProjectTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskId

`func (o *ProjectTimeLogUpdateDto) SetProjectTaskId(v string)`

SetProjectTaskId sets ProjectTaskId field to given value.

### HasProjectTaskId

`func (o *ProjectTimeLogUpdateDto) HasProjectTaskId() bool`

HasProjectTaskId returns a boolean if a field has been set.

### SetProjectTaskIdNil

`func (o *ProjectTimeLogUpdateDto) SetProjectTaskIdNil(b bool)`

 SetProjectTaskIdNil sets the value for ProjectTaskId to be an explicit nil

### UnsetProjectTaskId
`func (o *ProjectTimeLogUpdateDto) UnsetProjectTaskId()`

UnsetProjectTaskId ensures that no value is present for ProjectTaskId, not even an explicit nil
### GetProjectPeriodId

`func (o *ProjectTimeLogUpdateDto) GetProjectPeriodId() string`

GetProjectPeriodId returns the ProjectPeriodId field if non-nil, zero value otherwise.

### GetProjectPeriodIdOk

`func (o *ProjectTimeLogUpdateDto) GetProjectPeriodIdOk() (*string, bool)`

GetProjectPeriodIdOk returns a tuple with the ProjectPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodId

`func (o *ProjectTimeLogUpdateDto) SetProjectPeriodId(v string)`

SetProjectPeriodId sets ProjectPeriodId field to given value.

### HasProjectPeriodId

`func (o *ProjectTimeLogUpdateDto) HasProjectPeriodId() bool`

HasProjectPeriodId returns a boolean if a field has been set.

### SetProjectPeriodIdNil

`func (o *ProjectTimeLogUpdateDto) SetProjectPeriodIdNil(b bool)`

 SetProjectPeriodIdNil sets the value for ProjectPeriodId to be an explicit nil

### UnsetProjectPeriodId
`func (o *ProjectTimeLogUpdateDto) UnsetProjectPeriodId()`

UnsetProjectPeriodId ensures that no value is present for ProjectPeriodId, not even an explicit nil
### GetProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) GetProjectTimeLogRecordType() string`

GetProjectTimeLogRecordType returns the ProjectTimeLogRecordType field if non-nil, zero value otherwise.

### GetProjectTimeLogRecordTypeOk

`func (o *ProjectTimeLogUpdateDto) GetProjectTimeLogRecordTypeOk() (*string, bool)`

GetProjectTimeLogRecordTypeOk returns a tuple with the ProjectTimeLogRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) SetProjectTimeLogRecordType(v string)`

SetProjectTimeLogRecordType sets ProjectTimeLogRecordType field to given value.

### HasProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) HasProjectTimeLogRecordType() bool`

HasProjectTimeLogRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


