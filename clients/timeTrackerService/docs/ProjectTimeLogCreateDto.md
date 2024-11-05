# ProjectTimeLogCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**TimeSpan** | Pointer to **string** |  | [optional] 
**LogDate** | Pointer to **time.Time** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskID** | **string** |  | 
**ProjectPeriodID** | **string** |  | 
**ProjectTimeLogRecordType** | Pointer to **int32** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectTimeLogCreateDto

`func NewProjectTimeLogCreateDto(projectTaskID string, projectPeriodID string, ) *ProjectTimeLogCreateDto`

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
### GetProjectTaskID

`func (o *ProjectTimeLogCreateDto) GetProjectTaskID() string`

GetProjectTaskID returns the ProjectTaskID field if non-nil, zero value otherwise.

### GetProjectTaskIDOk

`func (o *ProjectTimeLogCreateDto) GetProjectTaskIDOk() (*string, bool)`

GetProjectTaskIDOk returns a tuple with the ProjectTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskID

`func (o *ProjectTimeLogCreateDto) SetProjectTaskID(v string)`

SetProjectTaskID sets ProjectTaskID field to given value.


### GetProjectPeriodID

`func (o *ProjectTimeLogCreateDto) GetProjectPeriodID() string`

GetProjectPeriodID returns the ProjectPeriodID field if non-nil, zero value otherwise.

### GetProjectPeriodIDOk

`func (o *ProjectTimeLogCreateDto) GetProjectPeriodIDOk() (*string, bool)`

GetProjectPeriodIDOk returns a tuple with the ProjectPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodID

`func (o *ProjectTimeLogCreateDto) SetProjectPeriodID(v string)`

SetProjectPeriodID sets ProjectPeriodID field to given value.


### GetProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordType() int32`

GetProjectTimeLogRecordType returns the ProjectTimeLogRecordType field if non-nil, zero value otherwise.

### GetProjectTimeLogRecordTypeOk

`func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordTypeOk() (*int32, bool)`

GetProjectTimeLogRecordTypeOk returns a tuple with the ProjectTimeLogRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) SetProjectTimeLogRecordType(v int32)`

SetProjectTimeLogRecordType sets ProjectTimeLogRecordType field to given value.

### HasProjectTimeLogRecordType

`func (o *ProjectTimeLogCreateDto) HasProjectTimeLogRecordType() bool`

HasProjectTimeLogRecordType returns a boolean if a field has been set.

### GetProjectID

`func (o *ProjectTimeLogCreateDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *ProjectTimeLogCreateDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *ProjectTimeLogCreateDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *ProjectTimeLogCreateDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *ProjectTimeLogCreateDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *ProjectTimeLogCreateDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


