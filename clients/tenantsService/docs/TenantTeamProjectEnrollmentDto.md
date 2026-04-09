# TenantTeamProjectEnrollmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TeamId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamProjectEnrollmentDto

`func NewTenantTeamProjectEnrollmentDto() *TenantTeamProjectEnrollmentDto`

NewTenantTeamProjectEnrollmentDto instantiates a new TenantTeamProjectEnrollmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamProjectEnrollmentDtoWithDefaults

`func NewTenantTeamProjectEnrollmentDtoWithDefaults() *TenantTeamProjectEnrollmentDto`

NewTenantTeamProjectEnrollmentDtoWithDefaults instantiates a new TenantTeamProjectEnrollmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamProjectEnrollmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamProjectEnrollmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamProjectEnrollmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamProjectEnrollmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantTeamProjectEnrollmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantTeamProjectEnrollmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantTeamProjectEnrollmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamProjectEnrollmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamProjectEnrollmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamProjectEnrollmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantTeamProjectEnrollmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantTeamProjectEnrollmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTeamId

`func (o *TenantTeamProjectEnrollmentDto) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TenantTeamProjectEnrollmentDto) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TenantTeamProjectEnrollmentDto) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TenantTeamProjectEnrollmentDto) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *TenantTeamProjectEnrollmentDto) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *TenantTeamProjectEnrollmentDto) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetTenantId

`func (o *TenantTeamProjectEnrollmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantTeamProjectEnrollmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantTeamProjectEnrollmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantTeamProjectEnrollmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantTeamProjectEnrollmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantTeamProjectEnrollmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantTeamProjectEnrollmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantTeamProjectEnrollmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantTeamProjectEnrollmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantTeamProjectEnrollmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantTeamProjectEnrollmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantTeamProjectEnrollmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetProjectID

`func (o *TenantTeamProjectEnrollmentDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *TenantTeamProjectEnrollmentDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *TenantTeamProjectEnrollmentDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *TenantTeamProjectEnrollmentDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *TenantTeamProjectEnrollmentDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *TenantTeamProjectEnrollmentDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


