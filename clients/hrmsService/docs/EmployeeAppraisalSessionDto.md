# EmployeeAppraisalSessionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 
**AppraisalWorkflowId** | Pointer to **NullableString** |  | [optional] 
**AppraisalStageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmployeeAppraisalSessionDto

`func NewEmployeeAppraisalSessionDto() *EmployeeAppraisalSessionDto`

NewEmployeeAppraisalSessionDto instantiates a new EmployeeAppraisalSessionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeAppraisalSessionDtoWithDefaults

`func NewEmployeeAppraisalSessionDtoWithDefaults() *EmployeeAppraisalSessionDto`

NewEmployeeAppraisalSessionDtoWithDefaults instantiates a new EmployeeAppraisalSessionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeAppraisalSessionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeAppraisalSessionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeAppraisalSessionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeAppraisalSessionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EmployeeAppraisalSessionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EmployeeAppraisalSessionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *EmployeeAppraisalSessionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmployeeAppraisalSessionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmployeeAppraisalSessionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmployeeAppraisalSessionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *EmployeeAppraisalSessionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *EmployeeAppraisalSessionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *EmployeeAppraisalSessionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmployeeAppraisalSessionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmployeeAppraisalSessionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmployeeAppraisalSessionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmployeeAppraisalSessionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmployeeAppraisalSessionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *EmployeeAppraisalSessionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *EmployeeAppraisalSessionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *EmployeeAppraisalSessionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *EmployeeAppraisalSessionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *EmployeeAppraisalSessionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *EmployeeAppraisalSessionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetEmployeeProfileId

`func (o *EmployeeAppraisalSessionDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *EmployeeAppraisalSessionDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *EmployeeAppraisalSessionDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *EmployeeAppraisalSessionDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *EmployeeAppraisalSessionDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *EmployeeAppraisalSessionDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil
### GetAppraisalWorkflowId

`func (o *EmployeeAppraisalSessionDto) GetAppraisalWorkflowId() string`

GetAppraisalWorkflowId returns the AppraisalWorkflowId field if non-nil, zero value otherwise.

### GetAppraisalWorkflowIdOk

`func (o *EmployeeAppraisalSessionDto) GetAppraisalWorkflowIdOk() (*string, bool)`

GetAppraisalWorkflowIdOk returns a tuple with the AppraisalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalWorkflowId

`func (o *EmployeeAppraisalSessionDto) SetAppraisalWorkflowId(v string)`

SetAppraisalWorkflowId sets AppraisalWorkflowId field to given value.

### HasAppraisalWorkflowId

`func (o *EmployeeAppraisalSessionDto) HasAppraisalWorkflowId() bool`

HasAppraisalWorkflowId returns a boolean if a field has been set.

### SetAppraisalWorkflowIdNil

`func (o *EmployeeAppraisalSessionDto) SetAppraisalWorkflowIdNil(b bool)`

 SetAppraisalWorkflowIdNil sets the value for AppraisalWorkflowId to be an explicit nil

### UnsetAppraisalWorkflowId
`func (o *EmployeeAppraisalSessionDto) UnsetAppraisalWorkflowId()`

UnsetAppraisalWorkflowId ensures that no value is present for AppraisalWorkflowId, not even an explicit nil
### GetAppraisalStageId

`func (o *EmployeeAppraisalSessionDto) GetAppraisalStageId() string`

GetAppraisalStageId returns the AppraisalStageId field if non-nil, zero value otherwise.

### GetAppraisalStageIdOk

`func (o *EmployeeAppraisalSessionDto) GetAppraisalStageIdOk() (*string, bool)`

GetAppraisalStageIdOk returns a tuple with the AppraisalStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalStageId

`func (o *EmployeeAppraisalSessionDto) SetAppraisalStageId(v string)`

SetAppraisalStageId sets AppraisalStageId field to given value.

### HasAppraisalStageId

`func (o *EmployeeAppraisalSessionDto) HasAppraisalStageId() bool`

HasAppraisalStageId returns a boolean if a field has been set.

### SetAppraisalStageIdNil

`func (o *EmployeeAppraisalSessionDto) SetAppraisalStageIdNil(b bool)`

 SetAppraisalStageIdNil sets the value for AppraisalStageId to be an explicit nil

### UnsetAppraisalStageId
`func (o *EmployeeAppraisalSessionDto) UnsetAppraisalStageId()`

UnsetAppraisalStageId ensures that no value is present for AppraisalStageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


