# TenantTeamEmployeeEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | **string** |  | 
**BusinessProfileRecordID** | **string** |  | 
**BusinessTeamID** | **string** |  | 
**EmployeeProfileID** | **string** |  | 

## Methods

### NewTenantTeamEmployeeEnrollmentCreateDto

`func NewTenantTeamEmployeeEnrollmentCreateDto(businessID string, businessProfileRecordID string, businessTeamID string, employeeProfileID string, ) *TenantTeamEmployeeEnrollmentCreateDto`

NewTenantTeamEmployeeEnrollmentCreateDto instantiates a new TenantTeamEmployeeEnrollmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamEmployeeEnrollmentCreateDtoWithDefaults

`func NewTenantTeamEmployeeEnrollmentCreateDtoWithDefaults() *TenantTeamEmployeeEnrollmentCreateDto`

NewTenantTeamEmployeeEnrollmentCreateDtoWithDefaults instantiates a new TenantTeamEmployeeEnrollmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamEmployeeEnrollmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamEmployeeEnrollmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetBusinessProfileRecordID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.


### GetBusinessTeamID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessTeamID() string`

GetBusinessTeamID returns the BusinessTeamID field if non-nil, zero value otherwise.

### GetBusinessTeamIDOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetBusinessTeamIDOk() (*string, bool)`

GetBusinessTeamIDOk returns a tuple with the BusinessTeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetBusinessTeamID(v string)`

SetBusinessTeamID sets BusinessTeamID field to given value.


### GetEmployeeProfileID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetEmployeeProfileID() string`

GetEmployeeProfileID returns the EmployeeProfileID field if non-nil, zero value otherwise.

### GetEmployeeProfileIDOk

`func (o *TenantTeamEmployeeEnrollmentCreateDto) GetEmployeeProfileIDOk() (*string, bool)`

GetEmployeeProfileIDOk returns a tuple with the EmployeeProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileID

`func (o *TenantTeamEmployeeEnrollmentCreateDto) SetEmployeeProfileID(v string)`

SetEmployeeProfileID sets EmployeeProfileID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


