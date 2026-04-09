# TenantSegmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Revenue** | Pointer to **NullableString** |  | [optional] 
**MinEmployees** | Pointer to **float64** |  | [optional] 
**MaxEmployees** | Pointer to **float64** |  | [optional] 

## Methods

### NewTenantSegmentCreateDto

`func NewTenantSegmentCreateDto() *TenantSegmentCreateDto`

NewTenantSegmentCreateDto instantiates a new TenantSegmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSegmentCreateDtoWithDefaults

`func NewTenantSegmentCreateDtoWithDefaults() *TenantSegmentCreateDto`

NewTenantSegmentCreateDtoWithDefaults instantiates a new TenantSegmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantSegmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantSegmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantSegmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantSegmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantSegmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantSegmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantSegmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantSegmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRevenue

`func (o *TenantSegmentCreateDto) GetRevenue() string`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *TenantSegmentCreateDto) GetRevenueOk() (*string, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *TenantSegmentCreateDto) SetRevenue(v string)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *TenantSegmentCreateDto) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *TenantSegmentCreateDto) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *TenantSegmentCreateDto) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetMinEmployees

`func (o *TenantSegmentCreateDto) GetMinEmployees() float64`

GetMinEmployees returns the MinEmployees field if non-nil, zero value otherwise.

### GetMinEmployeesOk

`func (o *TenantSegmentCreateDto) GetMinEmployeesOk() (*float64, bool)`

GetMinEmployeesOk returns a tuple with the MinEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEmployees

`func (o *TenantSegmentCreateDto) SetMinEmployees(v float64)`

SetMinEmployees sets MinEmployees field to given value.

### HasMinEmployees

`func (o *TenantSegmentCreateDto) HasMinEmployees() bool`

HasMinEmployees returns a boolean if a field has been set.

### GetMaxEmployees

`func (o *TenantSegmentCreateDto) GetMaxEmployees() float64`

GetMaxEmployees returns the MaxEmployees field if non-nil, zero value otherwise.

### GetMaxEmployeesOk

`func (o *TenantSegmentCreateDto) GetMaxEmployeesOk() (*float64, bool)`

GetMaxEmployeesOk returns a tuple with the MaxEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEmployees

`func (o *TenantSegmentCreateDto) SetMaxEmployees(v float64)`

SetMaxEmployees sets MaxEmployees field to given value.

### HasMaxEmployees

`func (o *TenantSegmentCreateDto) HasMaxEmployees() bool`

HasMaxEmployees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


