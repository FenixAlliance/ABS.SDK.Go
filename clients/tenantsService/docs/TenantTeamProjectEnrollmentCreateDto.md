# TenantTeamProjectEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessTeamId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewTenantTeamProjectEnrollmentCreateDto

`func NewTenantTeamProjectEnrollmentCreateDto(businessTeamId string, projectId string, ) *TenantTeamProjectEnrollmentCreateDto`

NewTenantTeamProjectEnrollmentCreateDto instantiates a new TenantTeamProjectEnrollmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamProjectEnrollmentCreateDtoWithDefaults

`func NewTenantTeamProjectEnrollmentCreateDtoWithDefaults() *TenantTeamProjectEnrollmentCreateDto`

NewTenantTeamProjectEnrollmentCreateDtoWithDefaults instantiates a new TenantTeamProjectEnrollmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamProjectEnrollmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamProjectEnrollmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamProjectEnrollmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTeamProjectEnrollmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamProjectEnrollmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamProjectEnrollmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessTeamId

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessTeamId() string`

GetBusinessTeamId returns the BusinessTeamId field if non-nil, zero value otherwise.

### GetBusinessTeamIdOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessTeamIdOk() (*string, bool)`

GetBusinessTeamIdOk returns a tuple with the BusinessTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamId

`func (o *TenantTeamProjectEnrollmentCreateDto) SetBusinessTeamId(v string)`

SetBusinessTeamId sets BusinessTeamId field to given value.


### GetProjectId

`func (o *TenantTeamProjectEnrollmentCreateDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TenantTeamProjectEnrollmentCreateDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


