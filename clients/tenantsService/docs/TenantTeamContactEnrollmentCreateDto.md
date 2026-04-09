# TenantTeamContactEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | **string** |  | 
**BusinessProfileRecordID** | **string** |  | 
**BusinessTeamID** | **string** |  | 
**ContactID** | **string** |  | 

## Methods

### NewTenantTeamContactEnrollmentCreateDto

`func NewTenantTeamContactEnrollmentCreateDto(businessID string, businessProfileRecordID string, businessTeamID string, contactID string, ) *TenantTeamContactEnrollmentCreateDto`

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

### GetBusinessID

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTeamContactEnrollmentCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetBusinessProfileRecordID

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTeamContactEnrollmentCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.


### GetBusinessTeamID

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessTeamID() string`

GetBusinessTeamID returns the BusinessTeamID field if non-nil, zero value otherwise.

### GetBusinessTeamIDOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetBusinessTeamIDOk() (*string, bool)`

GetBusinessTeamIDOk returns a tuple with the BusinessTeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamID

`func (o *TenantTeamContactEnrollmentCreateDto) SetBusinessTeamID(v string)`

SetBusinessTeamID sets BusinessTeamID field to given value.


### GetContactID

`func (o *TenantTeamContactEnrollmentCreateDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *TenantTeamContactEnrollmentCreateDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *TenantTeamContactEnrollmentCreateDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


