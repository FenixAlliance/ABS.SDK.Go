# EmployeeAppraisalSessionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**EmployeeProfileId** | **string** |  | 
**AppraisalWorkflowId** | **string** |  | 
**AppraisalStageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmployeeAppraisalSessionCreateDto

`func NewEmployeeAppraisalSessionCreateDto(employeeProfileId string, appraisalWorkflowId string, ) *EmployeeAppraisalSessionCreateDto`

NewEmployeeAppraisalSessionCreateDto instantiates a new EmployeeAppraisalSessionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeAppraisalSessionCreateDtoWithDefaults

`func NewEmployeeAppraisalSessionCreateDtoWithDefaults() *EmployeeAppraisalSessionCreateDto`

NewEmployeeAppraisalSessionCreateDtoWithDefaults instantiates a new EmployeeAppraisalSessionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeAppraisalSessionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeAppraisalSessionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeAppraisalSessionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeAppraisalSessionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EmployeeAppraisalSessionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmployeeAppraisalSessionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmployeeAppraisalSessionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmployeeAppraisalSessionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEmployeeProfileId

`func (o *EmployeeAppraisalSessionCreateDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *EmployeeAppraisalSessionCreateDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *EmployeeAppraisalSessionCreateDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.


### GetAppraisalWorkflowId

`func (o *EmployeeAppraisalSessionCreateDto) GetAppraisalWorkflowId() string`

GetAppraisalWorkflowId returns the AppraisalWorkflowId field if non-nil, zero value otherwise.

### GetAppraisalWorkflowIdOk

`func (o *EmployeeAppraisalSessionCreateDto) GetAppraisalWorkflowIdOk() (*string, bool)`

GetAppraisalWorkflowIdOk returns a tuple with the AppraisalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalWorkflowId

`func (o *EmployeeAppraisalSessionCreateDto) SetAppraisalWorkflowId(v string)`

SetAppraisalWorkflowId sets AppraisalWorkflowId field to given value.


### GetAppraisalStageId

`func (o *EmployeeAppraisalSessionCreateDto) GetAppraisalStageId() string`

GetAppraisalStageId returns the AppraisalStageId field if non-nil, zero value otherwise.

### GetAppraisalStageIdOk

`func (o *EmployeeAppraisalSessionCreateDto) GetAppraisalStageIdOk() (*string, bool)`

GetAppraisalStageIdOk returns a tuple with the AppraisalStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalStageId

`func (o *EmployeeAppraisalSessionCreateDto) SetAppraisalStageId(v string)`

SetAppraisalStageId sets AppraisalStageId field to given value.

### HasAppraisalStageId

`func (o *EmployeeAppraisalSessionCreateDto) HasAppraisalStageId() bool`

HasAppraisalStageId returns a boolean if a field has been set.

### SetAppraisalStageIdNil

`func (o *EmployeeAppraisalSessionCreateDto) SetAppraisalStageIdNil(b bool)`

 SetAppraisalStageIdNil sets the value for AppraisalStageId to be an explicit nil

### UnsetAppraisalStageId
`func (o *EmployeeAppraisalSessionCreateDto) UnsetAppraisalStageId()`

UnsetAppraisalStageId ensures that no value is present for AppraisalStageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


