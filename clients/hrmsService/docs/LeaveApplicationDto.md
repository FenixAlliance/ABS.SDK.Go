# LeaveApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Justification** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**OnReview** | Pointer to **bool** |  | [optional] 
**LeaveTypeId** | Pointer to **NullableString** |  | [optional] 
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLeaveApplicationDto

`func NewLeaveApplicationDto() *LeaveApplicationDto`

NewLeaveApplicationDto instantiates a new LeaveApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaveApplicationDtoWithDefaults

`func NewLeaveApplicationDtoWithDefaults() *LeaveApplicationDto`

NewLeaveApplicationDtoWithDefaults instantiates a new LeaveApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeaveApplicationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaveApplicationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaveApplicationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LeaveApplicationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LeaveApplicationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LeaveApplicationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LeaveApplicationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LeaveApplicationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LeaveApplicationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LeaveApplicationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LeaveApplicationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LeaveApplicationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetJustification

`func (o *LeaveApplicationDto) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *LeaveApplicationDto) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *LeaveApplicationDto) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *LeaveApplicationDto) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *LeaveApplicationDto) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *LeaveApplicationDto) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetApproved

`func (o *LeaveApplicationDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *LeaveApplicationDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *LeaveApplicationDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *LeaveApplicationDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetOnReview

`func (o *LeaveApplicationDto) GetOnReview() bool`

GetOnReview returns the OnReview field if non-nil, zero value otherwise.

### GetOnReviewOk

`func (o *LeaveApplicationDto) GetOnReviewOk() (*bool, bool)`

GetOnReviewOk returns a tuple with the OnReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReview

`func (o *LeaveApplicationDto) SetOnReview(v bool)`

SetOnReview sets OnReview field to given value.

### HasOnReview

`func (o *LeaveApplicationDto) HasOnReview() bool`

HasOnReview returns a boolean if a field has been set.

### GetLeaveTypeId

`func (o *LeaveApplicationDto) GetLeaveTypeId() string`

GetLeaveTypeId returns the LeaveTypeId field if non-nil, zero value otherwise.

### GetLeaveTypeIdOk

`func (o *LeaveApplicationDto) GetLeaveTypeIdOk() (*string, bool)`

GetLeaveTypeIdOk returns a tuple with the LeaveTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveTypeId

`func (o *LeaveApplicationDto) SetLeaveTypeId(v string)`

SetLeaveTypeId sets LeaveTypeId field to given value.

### HasLeaveTypeId

`func (o *LeaveApplicationDto) HasLeaveTypeId() bool`

HasLeaveTypeId returns a boolean if a field has been set.

### SetLeaveTypeIdNil

`func (o *LeaveApplicationDto) SetLeaveTypeIdNil(b bool)`

 SetLeaveTypeIdNil sets the value for LeaveTypeId to be an explicit nil

### UnsetLeaveTypeId
`func (o *LeaveApplicationDto) UnsetLeaveTypeId()`

UnsetLeaveTypeId ensures that no value is present for LeaveTypeId, not even an explicit nil
### GetEmployeeProfileId

`func (o *LeaveApplicationDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *LeaveApplicationDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *LeaveApplicationDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *LeaveApplicationDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *LeaveApplicationDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *LeaveApplicationDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil
### GetTenantId

`func (o *LeaveApplicationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LeaveApplicationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LeaveApplicationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LeaveApplicationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LeaveApplicationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LeaveApplicationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LeaveApplicationDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LeaveApplicationDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LeaveApplicationDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LeaveApplicationDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LeaveApplicationDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LeaveApplicationDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


