# TenantTeamContactEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessTeamId** | **string** |  | 
**ContactId** | **string** |  | 

## Methods

### NewTenantTeamContactEnrollmentCreateDto

`func NewTenantTeamContactEnrollmentCreateDto(businessTeamId string, contactId string, ) *TenantTeamContactEnrollmentCreateDto`

NewTenantTeamContactEnrollmentCreateDto instantiates a new TenantTeamContactEnrollmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamContactEnrollmentCreateDtoWithDefaults

`func NewTenantTeamContactEnrollmentCreateDtoWithDefaults() *TenantTeamContactEnrollmentCreateDto`

NewTenantTeamContactEnrollmentCreateDtoWithDefaults instantiates a new TenantTeamContactEnrollmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamContactEnrollmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamContactEnrollmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamContactEnrollmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTeamContactEnrollmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamContactEnrollmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamContactEnrollmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessTeamId

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessTeamId() string`

GetBusinessTeamId returns the BusinessTeamId field if non-nil, zero value otherwise.

### GetBusinessTeamIdOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessTeamIdOk() (*string, bool)`

GetBusinessTeamIdOk returns a tuple with the BusinessTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamId

`func (o *TenantTeamContactEnrollmentCreateDto) SetBusinessTeamId(v string)`

SetBusinessTeamId sets BusinessTeamId field to given value.


### GetContactId

`func (o *TenantTeamContactEnrollmentCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TenantTeamContactEnrollmentCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


