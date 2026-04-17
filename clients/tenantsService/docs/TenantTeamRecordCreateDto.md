# TenantTeamRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessTeamID** | **string** |  | 

## Methods

### NewTenantTeamRecordCreateDto

`func NewTenantTeamRecordCreateDto(businessTeamID string, ) *TenantTeamRecordCreateDto`

NewTenantTeamRecordCreateDto instantiates a new TenantTeamRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamRecordCreateDtoWithDefaults

`func NewTenantTeamRecordCreateDtoWithDefaults() *TenantTeamRecordCreateDto`

NewTenantTeamRecordCreateDtoWithDefaults instantiates a new TenantTeamRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTeamRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessTeamID

`func (o *TenantTeamRecordCreateDto) GetBusinessTeamID() string`

GetBusinessTeamID returns the BusinessTeamID field if non-nil, zero value otherwise.

### GetBusinessTeamIDOk

`func (o *TenantTeamRecordCreateDto) GetBusinessTeamIDOk() (*string, bool)`

GetBusinessTeamIDOk returns a tuple with the BusinessTeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTeamID

`func (o *TenantTeamRecordCreateDto) SetBusinessTeamID(v string)`

SetBusinessTeamID sets BusinessTeamID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


