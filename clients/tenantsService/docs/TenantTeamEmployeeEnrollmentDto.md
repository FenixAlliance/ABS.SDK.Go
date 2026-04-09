# TenantTeamEmployeeEnrollmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TeamId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamEmployeeEnrollmentDto

`func NewTenantTeamEmployeeEnrollmentDto() *TenantTeamEmployeeEnrollmentDto`

NewTenantTeamEmployeeEnrollmentDto instantiates a new TenantTeamEmployeeEnrollmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamEmployeeEnrollmentDtoWithDefaults

`func NewTenantTeamEmployeeEnrollmentDtoWithDefaults() *TenantTeamEmployeeEnrollmentDto`

NewTenantTeamEmployeeEnrollmentDtoWithDefaults instantiates a new TenantTeamEmployeeEnrollmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamEmployeeEnrollmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamEmployeeEnrollmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamEmployeeEnrollmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantTeamEmployeeEnrollmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamEmployeeEnrollmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamEmployeeEnrollmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTeamId

`func (o *TenantTeamEmployeeEnrollmentDto) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TenantTeamEmployeeEnrollmentDto) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TenantTeamEmployeeEnrollmentDto) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetTenantId

`func (o *TenantTeamEmployeeEnrollmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantTeamEmployeeEnrollmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantTeamEmployeeEnrollmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantTeamEmployeeEnrollmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantTeamEmployeeEnrollmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantTeamEmployeeEnrollmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetEmployeeProfileId

`func (o *TenantTeamEmployeeEnrollmentDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *TenantTeamEmployeeEnrollmentDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *TenantTeamEmployeeEnrollmentDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *TenantTeamEmployeeEnrollmentDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *TenantTeamEmployeeEnrollmentDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *TenantTeamEmployeeEnrollmentDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


