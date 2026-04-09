# TenantSegmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Revenue** | Pointer to **NullableString** |  | [optional] 
**MinEmployees** | Pointer to **float64** |  | [optional] 
**MaxEmployees** | Pointer to **float64** |  | [optional] 

## Methods

### NewTenantSegmentDto

`func NewTenantSegmentDto() *TenantSegmentDto`

NewTenantSegmentDto instantiates a new TenantSegmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSegmentDtoWithDefaults

`func NewTenantSegmentDtoWithDefaults() *TenantSegmentDto`

NewTenantSegmentDtoWithDefaults instantiates a new TenantSegmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantSegmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantSegmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantSegmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantSegmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantSegmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantSegmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantSegmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantSegmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantSegmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantSegmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantSegmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantSegmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRevenue

`func (o *TenantSegmentDto) GetRevenue() string`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *TenantSegmentDto) GetRevenueOk() (*string, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *TenantSegmentDto) SetRevenue(v string)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *TenantSegmentDto) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *TenantSegmentDto) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *TenantSegmentDto) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetMinEmployees

`func (o *TenantSegmentDto) GetMinEmployees() float64`

GetMinEmployees returns the MinEmployees field if non-nil, zero value otherwise.

### GetMinEmployeesOk

`func (o *TenantSegmentDto) GetMinEmployeesOk() (*float64, bool)`

GetMinEmployeesOk returns a tuple with the MinEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEmployees

`func (o *TenantSegmentDto) SetMinEmployees(v float64)`

SetMinEmployees sets MinEmployees field to given value.

### HasMinEmployees

`func (o *TenantSegmentDto) HasMinEmployees() bool`

HasMinEmployees returns a boolean if a field has been set.

### GetMaxEmployees

`func (o *TenantSegmentDto) GetMaxEmployees() float64`

GetMaxEmployees returns the MaxEmployees field if non-nil, zero value otherwise.

### GetMaxEmployeesOk

`func (o *TenantSegmentDto) GetMaxEmployeesOk() (*float64, bool)`

GetMaxEmployeesOk returns a tuple with the MaxEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEmployees

`func (o *TenantSegmentDto) SetMaxEmployees(v float64)`

SetMaxEmployees sets MaxEmployees field to given value.

### HasMaxEmployees

`func (o *TenantSegmentDto) HasMaxEmployees() bool`

HasMaxEmployees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


