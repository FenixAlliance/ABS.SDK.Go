# TaskTypeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TaskCategoryId** | Pointer to **NullableString** |  | [optional] 
**DisplayInTimeTracker** | Pointer to **bool** |  | [optional] 
**RequiresDescription** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskTypeCreateDto

`func NewTaskTypeCreateDto() *TaskTypeCreateDto`

NewTaskTypeCreateDto instantiates a new TaskTypeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTypeCreateDtoWithDefaults

`func NewTaskTypeCreateDtoWithDefaults() *TaskTypeCreateDto`

NewTaskTypeCreateDtoWithDefaults instantiates a new TaskTypeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskTypeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskTypeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskTypeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskTypeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TaskTypeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TaskTypeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TaskTypeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TaskTypeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *TaskTypeCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskTypeCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskTypeCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskTypeCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TaskTypeCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TaskTypeCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTaskCategoryId

`func (o *TaskTypeCreateDto) GetTaskCategoryId() string`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *TaskTypeCreateDto) GetTaskCategoryIdOk() (*string, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *TaskTypeCreateDto) SetTaskCategoryId(v string)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *TaskTypeCreateDto) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### SetTaskCategoryIdNil

`func (o *TaskTypeCreateDto) SetTaskCategoryIdNil(b bool)`

 SetTaskCategoryIdNil sets the value for TaskCategoryId to be an explicit nil

### UnsetTaskCategoryId
`func (o *TaskTypeCreateDto) UnsetTaskCategoryId()`

UnsetTaskCategoryId ensures that no value is present for TaskCategoryId, not even an explicit nil
### GetDisplayInTimeTracker

`func (o *TaskTypeCreateDto) GetDisplayInTimeTracker() bool`

GetDisplayInTimeTracker returns the DisplayInTimeTracker field if non-nil, zero value otherwise.

### GetDisplayInTimeTrackerOk

`func (o *TaskTypeCreateDto) GetDisplayInTimeTrackerOk() (*bool, bool)`

GetDisplayInTimeTrackerOk returns a tuple with the DisplayInTimeTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInTimeTracker

`func (o *TaskTypeCreateDto) SetDisplayInTimeTracker(v bool)`

SetDisplayInTimeTracker sets DisplayInTimeTracker field to given value.

### HasDisplayInTimeTracker

`func (o *TaskTypeCreateDto) HasDisplayInTimeTracker() bool`

HasDisplayInTimeTracker returns a boolean if a field has been set.

### GetRequiresDescription

`func (o *TaskTypeCreateDto) GetRequiresDescription() bool`

GetRequiresDescription returns the RequiresDescription field if non-nil, zero value otherwise.

### GetRequiresDescriptionOk

`func (o *TaskTypeCreateDto) GetRequiresDescriptionOk() (*bool, bool)`

GetRequiresDescriptionOk returns a tuple with the RequiresDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDescription

`func (o *TaskTypeCreateDto) SetRequiresDescription(v bool)`

SetRequiresDescription sets RequiresDescription field to given value.

### HasRequiresDescription

`func (o *TaskTypeCreateDto) HasRequiresDescription() bool`

HasRequiresDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


