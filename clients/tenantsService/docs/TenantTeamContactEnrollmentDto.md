# TenantTeamContactEnrollmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TeamId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamContactEnrollmentDto

`func NewTenantTeamContactEnrollmentDto() *TenantTeamContactEnrollmentDto`

NewTenantTeamContactEnrollmentDto instantiates a new TenantTeamContactEnrollmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamContactEnrollmentDtoWithDefaults

`func NewTenantTeamContactEnrollmentDtoWithDefaults() *TenantTeamContactEnrollmentDto`

NewTenantTeamContactEnrollmentDtoWithDefaults instantiates a new TenantTeamContactEnrollmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamContactEnrollmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamContactEnrollmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamContactEnrollmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamContactEnrollmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantTeamContactEnrollmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantTeamContactEnrollmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantTeamContactEnrollmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamContactEnrollmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamContactEnrollmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamContactEnrollmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantTeamContactEnrollmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantTeamContactEnrollmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTeamId

`func (o *TenantTeamContactEnrollmentDto) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TenantTeamContactEnrollmentDto) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TenantTeamContactEnrollmentDto) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TenantTeamContactEnrollmentDto) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *TenantTeamContactEnrollmentDto) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *TenantTeamContactEnrollmentDto) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetTenantId

`func (o *TenantTeamContactEnrollmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantTeamContactEnrollmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantTeamContactEnrollmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantTeamContactEnrollmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantTeamContactEnrollmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantTeamContactEnrollmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantTeamContactEnrollmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantTeamContactEnrollmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantTeamContactEnrollmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantTeamContactEnrollmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantTeamContactEnrollmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantTeamContactEnrollmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetContactID

`func (o *TenantTeamContactEnrollmentDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *TenantTeamContactEnrollmentDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *TenantTeamContactEnrollmentDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *TenantTeamContactEnrollmentDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *TenantTeamContactEnrollmentDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *TenantTeamContactEnrollmentDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


