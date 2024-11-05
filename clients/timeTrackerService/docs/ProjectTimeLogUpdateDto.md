# ProjectTimeLogUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDate** | Pointer to **time.Time** |  | [optional] 
**TimeSpan** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskID** | Pointer to **NullableString** |  | [optional] 
**ProjectPeriodID** | Pointer to **NullableString** |  | [optional] 
**ProjectTimeLogRecordType** | Pointer to **int32** |  | [optional] 

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
### GetProjectTaskID

`func (o *ProjectTimeLogUpdateDto) GetProjectTaskID() string`

GetProjectTaskID returns the ProjectTaskID field if non-nil, zero value otherwise.

### GetProjectTaskIDOk

`func (o *ProjectTimeLogUpdateDto) GetProjectTaskIDOk() (*string, bool)`

GetProjectTaskIDOk returns a tuple with the ProjectTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskID

`func (o *ProjectTimeLogUpdateDto) SetProjectTaskID(v string)`

SetProjectTaskID sets ProjectTaskID field to given value.

### HasProjectTaskID

`func (o *ProjectTimeLogUpdateDto) HasProjectTaskID() bool`

HasProjectTaskID returns a boolean if a field has been set.

### SetProjectTaskIDNil

`func (o *ProjectTimeLogUpdateDto) SetProjectTaskIDNil(b bool)`

 SetProjectTaskIDNil sets the value for ProjectTaskID to be an explicit nil

### UnsetProjectTaskID
`func (o *ProjectTimeLogUpdateDto) UnsetProjectTaskID()`

UnsetProjectTaskID ensures that no value is present for ProjectTaskID, not even an explicit nil
### GetProjectPeriodID

`func (o *ProjectTimeLogUpdateDto) GetProjectPeriodID() string`

GetProjectPeriodID returns the ProjectPeriodID field if non-nil, zero value otherwise.

### GetProjectPeriodIDOk

`func (o *ProjectTimeLogUpdateDto) GetProjectPeriodIDOk() (*string, bool)`

GetProjectPeriodIDOk returns a tuple with the ProjectPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodID

`func (o *ProjectTimeLogUpdateDto) SetProjectPeriodID(v string)`

SetProjectPeriodID sets ProjectPeriodID field to given value.

### HasProjectPeriodID

`func (o *ProjectTimeLogUpdateDto) HasProjectPeriodID() bool`

HasProjectPeriodID returns a boolean if a field has been set.

### SetProjectPeriodIDNil

`func (o *ProjectTimeLogUpdateDto) SetProjectPeriodIDNil(b bool)`

 SetProjectPeriodIDNil sets the value for ProjectPeriodID to be an explicit nil

### UnsetProjectPeriodID
`func (o *ProjectTimeLogUpdateDto) UnsetProjectPeriodID()`

UnsetProjectPeriodID ensures that no value is present for ProjectPeriodID, not even an explicit nil
### GetProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) GetProjectTimeLogRecordType() int32`

GetProjectTimeLogRecordType returns the ProjectTimeLogRecordType field if non-nil, zero value otherwise.

### GetProjectTimeLogRecordTypeOk

`func (o *ProjectTimeLogUpdateDto) GetProjectTimeLogRecordTypeOk() (*int32, bool)`

GetProjectTimeLogRecordTypeOk returns a tuple with the ProjectTimeLogRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) SetProjectTimeLogRecordType(v int32)`

SetProjectTimeLogRecordType sets ProjectTimeLogRecordType field to given value.

### HasProjectTimeLogRecordType

`func (o *ProjectTimeLogUpdateDto) HasProjectTimeLogRecordType() bool`

HasProjectTimeLogRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


