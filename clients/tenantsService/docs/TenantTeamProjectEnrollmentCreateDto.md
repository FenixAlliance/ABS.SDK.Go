# TenantTeamProjectEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | **string** |  | 
**BusinessProfileRecordID** | **string** |  | 
**BusinessTeamID** | **string** |  | 
**ProjectID** | **string** |  | 

## Methods

### NewTenantTeamProjectEnrollmentCreateDto

`func NewTenantTeamProjectEnrollmentCreateDto(businessID string, businessProfileRecordID string, businessTeamID string, projectID string, ) *TenantTeamProjectEnrollmentCreateDto`

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

### GetBusinessID

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTeamProjectEnrollmentCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetBusinessProfileRecordID

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTeamProjectEnrollmentCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.


### GetBusinessTeamID

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessTeamID() string`

GetBusinessTeamID returns the BusinessTeamID field if non-nil, zero value otherwise.

### GetBusinessTeamIDOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetBusinessTeamIDOk() (*string, bool)`

GetBusinessTeamIDOk returns a tuple with the BusinessTeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamID

`func (o *TenantTeamProjectEnrollmentCreateDto) SetBusinessTeamID(v string)`

SetBusinessTeamID sets BusinessTeamID field to given value.


### GetProjectID

`func (o *TenantTeamProjectEnrollmentCreateDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *TenantTeamProjectEnrollmentCreateDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *TenantTeamProjectEnrollmentCreateDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


