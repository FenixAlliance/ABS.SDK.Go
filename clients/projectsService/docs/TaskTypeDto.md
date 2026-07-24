# TaskTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TaskCategoryId** | Pointer to **NullableString** |  | [optional] 
**DisplayInTimeTracker** | Pointer to **bool** |  | [optional] 
**RequiresDescription** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskTypeDto

`func NewTaskTypeDto() *TaskTypeDto`

NewTaskTypeDto instantiates a new TaskTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTypeDtoWithDefaults

`func NewTaskTypeDtoWithDefaults() *TaskTypeDto`

NewTaskTypeDtoWithDefaults instantiates a new TaskTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskTypeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskTypeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskTypeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskTypeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskTypeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskTypeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TaskTypeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TaskTypeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TaskTypeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TaskTypeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TaskTypeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TaskTypeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *TaskTypeDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskTypeDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskTypeDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskTypeDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TaskTypeDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TaskTypeDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTaskCategoryId

`func (o *TaskTypeDto) GetTaskCategoryId() string`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *TaskTypeDto) GetTaskCategoryIdOk() (*string, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *TaskTypeDto) SetTaskCategoryId(v string)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *TaskTypeDto) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### SetTaskCategoryIdNil

`func (o *TaskTypeDto) SetTaskCategoryIdNil(b bool)`

 SetTaskCategoryIdNil sets the value for TaskCategoryId to be an explicit nil

### UnsetTaskCategoryId
`func (o *TaskTypeDto) UnsetTaskCategoryId()`

UnsetTaskCategoryId ensures that no value is present for TaskCategoryId, not even an explicit nil
### GetDisplayInTimeTracker

`func (o *TaskTypeDto) GetDisplayInTimeTracker() bool`

GetDisplayInTimeTracker returns the DisplayInTimeTracker field if non-nil, zero value otherwise.

### GetDisplayInTimeTrackerOk

`func (o *TaskTypeDto) GetDisplayInTimeTrackerOk() (*bool, bool)`

GetDisplayInTimeTrackerOk returns a tuple with the DisplayInTimeTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInTimeTracker

`func (o *TaskTypeDto) SetDisplayInTimeTracker(v bool)`

SetDisplayInTimeTracker sets DisplayInTimeTracker field to given value.

### HasDisplayInTimeTracker

`func (o *TaskTypeDto) HasDisplayInTimeTracker() bool`

HasDisplayInTimeTracker returns a boolean if a field has been set.

### GetRequiresDescription

`func (o *TaskTypeDto) GetRequiresDescription() bool`

GetRequiresDescription returns the RequiresDescription field if non-nil, zero value otherwise.

### GetRequiresDescriptionOk

`func (o *TaskTypeDto) GetRequiresDescriptionOk() (*bool, bool)`

GetRequiresDescriptionOk returns a tuple with the RequiresDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDescription

`func (o *TaskTypeDto) SetRequiresDescription(v bool)`

SetRequiresDescription sets RequiresDescription field to given value.

### HasRequiresDescription

`func (o *TaskTypeDto) HasRequiresDescription() bool`

HasRequiresDescription returns a boolean if a field has been set.

### GetTenantId

`func (o *TaskTypeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TaskTypeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TaskTypeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TaskTypeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TaskTypeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TaskTypeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TaskTypeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TaskTypeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TaskTypeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TaskTypeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TaskTypeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TaskTypeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


