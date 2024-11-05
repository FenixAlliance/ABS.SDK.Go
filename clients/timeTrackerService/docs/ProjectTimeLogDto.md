# ProjectTimeLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskId** | Pointer to **NullableString** |  | [optional] 
**TaskCategoryId** | Pointer to **NullableString** |  | [optional] 
**ProjectPeriodId** | Pointer to **NullableString** |  | [optional] 
**ResponsibleContactId** | Pointer to **NullableString** |  | [optional] 
**CreatorContactId** | Pointer to **NullableString** |  | [optional] 
**RecordType** | Pointer to **int32** |  | [optional] 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**TimeSpan** | Pointer to **string** |  | [optional] 
**LogDate** | Pointer to **time.Time** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectTimeLogDto

`func NewProjectTimeLogDto() *ProjectTimeLogDto`

NewProjectTimeLogDto instantiates a new ProjectTimeLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTimeLogDtoWithDefaults

`func NewProjectTimeLogDtoWithDefaults() *ProjectTimeLogDto`

NewProjectTimeLogDtoWithDefaults instantiates a new ProjectTimeLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTimeLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTimeLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTimeLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTimeLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectTimeLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectTimeLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProjectTimeLogDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectTimeLogDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectTimeLogDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectTimeLogDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProjectTimeLogDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProjectTimeLogDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetProjectId

`func (o *ProjectTimeLogDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectTimeLogDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectTimeLogDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectTimeLogDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectTimeLogDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectTimeLogDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectTaskId

`func (o *ProjectTimeLogDto) GetProjectTaskId() string`

GetProjectTaskId returns the ProjectTaskId field if non-nil, zero value otherwise.

### GetProjectTaskIdOk

`func (o *ProjectTimeLogDto) GetProjectTaskIdOk() (*string, bool)`

GetProjectTaskIdOk returns a tuple with the ProjectTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskId

`func (o *ProjectTimeLogDto) SetProjectTaskId(v string)`

SetProjectTaskId sets ProjectTaskId field to given value.

### HasProjectTaskId

`func (o *ProjectTimeLogDto) HasProjectTaskId() bool`

HasProjectTaskId returns a boolean if a field has been set.

### SetProjectTaskIdNil

`func (o *ProjectTimeLogDto) SetProjectTaskIdNil(b bool)`

 SetProjectTaskIdNil sets the value for ProjectTaskId to be an explicit nil

### UnsetProjectTaskId
`func (o *ProjectTimeLogDto) UnsetProjectTaskId()`

UnsetProjectTaskId ensures that no value is present for ProjectTaskId, not even an explicit nil
### GetTaskCategoryId

`func (o *ProjectTimeLogDto) GetTaskCategoryId() string`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *ProjectTimeLogDto) GetTaskCategoryIdOk() (*string, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *ProjectTimeLogDto) SetTaskCategoryId(v string)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *ProjectTimeLogDto) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### SetTaskCategoryIdNil

`func (o *ProjectTimeLogDto) SetTaskCategoryIdNil(b bool)`

 SetTaskCategoryIdNil sets the value for TaskCategoryId to be an explicit nil

### UnsetTaskCategoryId
`func (o *ProjectTimeLogDto) UnsetTaskCategoryId()`

UnsetTaskCategoryId ensures that no value is present for TaskCategoryId, not even an explicit nil
### GetProjectPeriodId

`func (o *ProjectTimeLogDto) GetProjectPeriodId() string`

GetProjectPeriodId returns the ProjectPeriodId field if non-nil, zero value otherwise.

### GetProjectPeriodIdOk

`func (o *ProjectTimeLogDto) GetProjectPeriodIdOk() (*string, bool)`

GetProjectPeriodIdOk returns a tuple with the ProjectPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodId

`func (o *ProjectTimeLogDto) SetProjectPeriodId(v string)`

SetProjectPeriodId sets ProjectPeriodId field to given value.

### HasProjectPeriodId

`func (o *ProjectTimeLogDto) HasProjectPeriodId() bool`

HasProjectPeriodId returns a boolean if a field has been set.

### SetProjectPeriodIdNil

`func (o *ProjectTimeLogDto) SetProjectPeriodIdNil(b bool)`

 SetProjectPeriodIdNil sets the value for ProjectPeriodId to be an explicit nil

### UnsetProjectPeriodId
`func (o *ProjectTimeLogDto) UnsetProjectPeriodId()`

UnsetProjectPeriodId ensures that no value is present for ProjectPeriodId, not even an explicit nil
### GetResponsibleContactId

`func (o *ProjectTimeLogDto) GetResponsibleContactId() string`

GetResponsibleContactId returns the ResponsibleContactId field if non-nil, zero value otherwise.

### GetResponsibleContactIdOk

`func (o *ProjectTimeLogDto) GetResponsibleContactIdOk() (*string, bool)`

GetResponsibleContactIdOk returns a tuple with the ResponsibleContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleContactId

`func (o *ProjectTimeLogDto) SetResponsibleContactId(v string)`

SetResponsibleContactId sets ResponsibleContactId field to given value.

### HasResponsibleContactId

`func (o *ProjectTimeLogDto) HasResponsibleContactId() bool`

HasResponsibleContactId returns a boolean if a field has been set.

### SetResponsibleContactIdNil

`func (o *ProjectTimeLogDto) SetResponsibleContactIdNil(b bool)`

 SetResponsibleContactIdNil sets the value for ResponsibleContactId to be an explicit nil

### UnsetResponsibleContactId
`func (o *ProjectTimeLogDto) UnsetResponsibleContactId()`

UnsetResponsibleContactId ensures that no value is present for ResponsibleContactId, not even an explicit nil
### GetCreatorContactId

`func (o *ProjectTimeLogDto) GetCreatorContactId() string`

GetCreatorContactId returns the CreatorContactId field if non-nil, zero value otherwise.

### GetCreatorContactIdOk

`func (o *ProjectTimeLogDto) GetCreatorContactIdOk() (*string, bool)`

GetCreatorContactIdOk returns a tuple with the CreatorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorContactId

`func (o *ProjectTimeLogDto) SetCreatorContactId(v string)`

SetCreatorContactId sets CreatorContactId field to given value.

### HasCreatorContactId

`func (o *ProjectTimeLogDto) HasCreatorContactId() bool`

HasCreatorContactId returns a boolean if a field has been set.

### SetCreatorContactIdNil

`func (o *ProjectTimeLogDto) SetCreatorContactIdNil(b bool)`

 SetCreatorContactIdNil sets the value for CreatorContactId to be an explicit nil

### UnsetCreatorContactId
`func (o *ProjectTimeLogDto) UnsetCreatorContactId()`

UnsetCreatorContactId ensures that no value is present for CreatorContactId, not even an explicit nil
### GetRecordType

`func (o *ProjectTimeLogDto) GetRecordType() int32`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ProjectTimeLogDto) GetRecordTypeOk() (*int32, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ProjectTimeLogDto) SetRecordType(v int32)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ProjectTimeLogDto) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ProjectTimeLogDto) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ProjectTimeLogDto) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ProjectTimeLogDto) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ProjectTimeLogDto) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetTimeSpan

`func (o *ProjectTimeLogDto) GetTimeSpan() string`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *ProjectTimeLogDto) GetTimeSpanOk() (*string, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *ProjectTimeLogDto) SetTimeSpan(v string)`

SetTimeSpan sets TimeSpan field to given value.

### HasTimeSpan

`func (o *ProjectTimeLogDto) HasTimeSpan() bool`

HasTimeSpan returns a boolean if a field has been set.

### GetLogDate

`func (o *ProjectTimeLogDto) GetLogDate() time.Time`

GetLogDate returns the LogDate field if non-nil, zero value otherwise.

### GetLogDateOk

`func (o *ProjectTimeLogDto) GetLogDateOk() (*time.Time, bool)`

GetLogDateOk returns a tuple with the LogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDate

`func (o *ProjectTimeLogDto) SetLogDate(v time.Time)`

SetLogDate sets LogDate field to given value.

### HasLogDate

`func (o *ProjectTimeLogDto) HasLogDate() bool`

HasLogDate returns a boolean if a field has been set.

### GetComments

`func (o *ProjectTimeLogDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProjectTimeLogDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProjectTimeLogDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProjectTimeLogDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ProjectTimeLogDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ProjectTimeLogDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetType

`func (o *ProjectTimeLogDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectTimeLogDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectTimeLogDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectTimeLogDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ProjectTimeLogDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProjectTimeLogDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


