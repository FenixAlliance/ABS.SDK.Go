# TaskTypeUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**DisplayInTimeTracker** | Pointer to **bool** |  | [optional] 
**RequiresDescription** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskTypeUpdateDto

`func NewTaskTypeUpdateDto() *TaskTypeUpdateDto`

NewTaskTypeUpdateDto instantiates a new TaskTypeUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTypeUpdateDtoWithDefaults

`func NewTaskTypeUpdateDtoWithDefaults() *TaskTypeUpdateDto`

NewTaskTypeUpdateDtoWithDefaults instantiates a new TaskTypeUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TaskTypeUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskTypeUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskTypeUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskTypeUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TaskTypeUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TaskTypeUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDisplayInTimeTracker

`func (o *TaskTypeUpdateDto) GetDisplayInTimeTracker() bool`

GetDisplayInTimeTracker returns the DisplayInTimeTracker field if non-nil, zero value otherwise.

### GetDisplayInTimeTrackerOk

`func (o *TaskTypeUpdateDto) GetDisplayInTimeTrackerOk() (*bool, bool)`

GetDisplayInTimeTrackerOk returns a tuple with the DisplayInTimeTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInTimeTracker

`func (o *TaskTypeUpdateDto) SetDisplayInTimeTracker(v bool)`

SetDisplayInTimeTracker sets DisplayInTimeTracker field to given value.

### HasDisplayInTimeTracker

`func (o *TaskTypeUpdateDto) HasDisplayInTimeTracker() bool`

HasDisplayInTimeTracker returns a boolean if a field has been set.

### GetRequiresDescription

`func (o *TaskTypeUpdateDto) GetRequiresDescription() bool`

GetRequiresDescription returns the RequiresDescription field if non-nil, zero value otherwise.

### GetRequiresDescriptionOk

`func (o *TaskTypeUpdateDto) GetRequiresDescriptionOk() (*bool, bool)`

GetRequiresDescriptionOk returns a tuple with the RequiresDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDescription

`func (o *TaskTypeUpdateDto) SetRequiresDescription(v bool)`

SetRequiresDescription sets RequiresDescription field to given value.

### HasRequiresDescription

`func (o *TaskTypeUpdateDto) HasRequiresDescription() bool`

HasRequiresDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


