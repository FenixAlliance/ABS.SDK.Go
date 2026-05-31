# LeaveApplicationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Justification** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**OnReview** | Pointer to **bool** |  | [optional] 
**LeaveTypeId** | **string** |  | 
**EmployeeProfileId** | **string** |  | 

## Methods

### NewLeaveApplicationCreateDto

`func NewLeaveApplicationCreateDto(leaveTypeId string, employeeProfileId string, ) *LeaveApplicationCreateDto`

NewLeaveApplicationCreateDto instantiates a new LeaveApplicationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaveApplicationCreateDtoWithDefaults

`func NewLeaveApplicationCreateDtoWithDefaults() *LeaveApplicationCreateDto`

NewLeaveApplicationCreateDtoWithDefaults instantiates a new LeaveApplicationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeaveApplicationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaveApplicationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaveApplicationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LeaveApplicationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LeaveApplicationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LeaveApplicationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LeaveApplicationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LeaveApplicationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJustification

`func (o *LeaveApplicationCreateDto) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *LeaveApplicationCreateDto) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *LeaveApplicationCreateDto) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *LeaveApplicationCreateDto) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *LeaveApplicationCreateDto) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *LeaveApplicationCreateDto) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetApproved

`func (o *LeaveApplicationCreateDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *LeaveApplicationCreateDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *LeaveApplicationCreateDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *LeaveApplicationCreateDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetOnReview

`func (o *LeaveApplicationCreateDto) GetOnReview() bool`

GetOnReview returns the OnReview field if non-nil, zero value otherwise.

### GetOnReviewOk

`func (o *LeaveApplicationCreateDto) GetOnReviewOk() (*bool, bool)`

GetOnReviewOk returns a tuple with the OnReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReview

`func (o *LeaveApplicationCreateDto) SetOnReview(v bool)`

SetOnReview sets OnReview field to given value.

### HasOnReview

`func (o *LeaveApplicationCreateDto) HasOnReview() bool`

HasOnReview returns a boolean if a field has been set.

### GetLeaveTypeId

`func (o *LeaveApplicationCreateDto) GetLeaveTypeId() string`

GetLeaveTypeId returns the LeaveTypeId field if non-nil, zero value otherwise.

### GetLeaveTypeIdOk

`func (o *LeaveApplicationCreateDto) GetLeaveTypeIdOk() (*string, bool)`

GetLeaveTypeIdOk returns a tuple with the LeaveTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveTypeId

`func (o *LeaveApplicationCreateDto) SetLeaveTypeId(v string)`

SetLeaveTypeId sets LeaveTypeId field to given value.


### GetEmployeeProfileId

`func (o *LeaveApplicationCreateDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *LeaveApplicationCreateDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *LeaveApplicationCreateDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


