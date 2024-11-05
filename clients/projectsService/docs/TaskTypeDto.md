# TaskTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TaskCategoryID** | Pointer to **NullableString** |  | [optional] 
**DisplayInTimeTracker** | Pointer to **bool** |  | [optional] 
**RequiresDescription** | Pointer to **bool** |  | [optional] 

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
### GetTaskCategoryID

`func (o *TaskTypeDto) GetTaskCategoryID() string`

GetTaskCategoryID returns the TaskCategoryID field if non-nil, zero value otherwise.

### GetTaskCategoryIDOk

`func (o *TaskTypeDto) GetTaskCategoryIDOk() (*string, bool)`

GetTaskCategoryIDOk returns a tuple with the TaskCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryID

`func (o *TaskTypeDto) SetTaskCategoryID(v string)`

SetTaskCategoryID sets TaskCategoryID field to given value.

### HasTaskCategoryID

`func (o *TaskTypeDto) HasTaskCategoryID() bool`

HasTaskCategoryID returns a boolean if a field has been set.

### SetTaskCategoryIDNil

`func (o *TaskTypeDto) SetTaskCategoryIDNil(b bool)`

 SetTaskCategoryIDNil sets the value for TaskCategoryID to be an explicit nil

### UnsetTaskCategoryID
`func (o *TaskTypeDto) UnsetTaskCategoryID()`

UnsetTaskCategoryID ensures that no value is present for TaskCategoryID, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


