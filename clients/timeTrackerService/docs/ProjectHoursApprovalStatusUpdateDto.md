# ProjectHoursApprovalStatusUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalStatus** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectHoursApprovalStatusUpdateDto

`func NewProjectHoursApprovalStatusUpdateDto() *ProjectHoursApprovalStatusUpdateDto`

NewProjectHoursApprovalStatusUpdateDto instantiates a new ProjectHoursApprovalStatusUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectHoursApprovalStatusUpdateDtoWithDefaults

`func NewProjectHoursApprovalStatusUpdateDtoWithDefaults() *ProjectHoursApprovalStatusUpdateDto`

NewProjectHoursApprovalStatusUpdateDtoWithDefaults instantiates a new ProjectHoursApprovalStatusUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalStatus

`func (o *ProjectHoursApprovalStatusUpdateDto) GetApprovalStatus() int32`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *ProjectHoursApprovalStatusUpdateDto) GetApprovalStatusOk() (*int32, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *ProjectHoursApprovalStatusUpdateDto) SetApprovalStatus(v int32)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *ProjectHoursApprovalStatusUpdateDto) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetComments

`func (o *ProjectHoursApprovalStatusUpdateDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProjectHoursApprovalStatusUpdateDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProjectHoursApprovalStatusUpdateDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProjectHoursApprovalStatusUpdateDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ProjectHoursApprovalStatusUpdateDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ProjectHoursApprovalStatusUpdateDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


