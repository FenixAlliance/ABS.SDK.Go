# TenantSizeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EmployeeLowRangeValue** | Pointer to **int32** |  | [optional] 
**EmployeeHighRangeValue** | Pointer to **int32** |  | [optional] 

## Methods

### NewTenantSizeCreateDto

`func NewTenantSizeCreateDto() *TenantSizeCreateDto`

NewTenantSizeCreateDto instantiates a new TenantSizeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSizeCreateDtoWithDefaults

`func NewTenantSizeCreateDtoWithDefaults() *TenantSizeCreateDto`

NewTenantSizeCreateDtoWithDefaults instantiates a new TenantSizeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantSizeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantSizeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantSizeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantSizeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantSizeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantSizeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantSizeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantSizeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TenantSizeCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantSizeCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantSizeCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantSizeCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantSizeCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantSizeCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmployeeLowRangeValue

`func (o *TenantSizeCreateDto) GetEmployeeLowRangeValue() int32`

GetEmployeeLowRangeValue returns the EmployeeLowRangeValue field if non-nil, zero value otherwise.

### GetEmployeeLowRangeValueOk

`func (o *TenantSizeCreateDto) GetEmployeeLowRangeValueOk() (*int32, bool)`

GetEmployeeLowRangeValueOk returns a tuple with the EmployeeLowRangeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeLowRangeValue

`func (o *TenantSizeCreateDto) SetEmployeeLowRangeValue(v int32)`

SetEmployeeLowRangeValue sets EmployeeLowRangeValue field to given value.

### HasEmployeeLowRangeValue

`func (o *TenantSizeCreateDto) HasEmployeeLowRangeValue() bool`

HasEmployeeLowRangeValue returns a boolean if a field has been set.

### GetEmployeeHighRangeValue

`func (o *TenantSizeCreateDto) GetEmployeeHighRangeValue() int32`

GetEmployeeHighRangeValue returns the EmployeeHighRangeValue field if non-nil, zero value otherwise.

### GetEmployeeHighRangeValueOk

`func (o *TenantSizeCreateDto) GetEmployeeHighRangeValueOk() (*int32, bool)`

GetEmployeeHighRangeValueOk returns a tuple with the EmployeeHighRangeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeHighRangeValue

`func (o *TenantSizeCreateDto) SetEmployeeHighRangeValue(v int32)`

SetEmployeeHighRangeValue sets EmployeeHighRangeValue field to given value.

### HasEmployeeHighRangeValue

`func (o *TenantSizeCreateDto) HasEmployeeHighRangeValue() bool`

HasEmployeeHighRangeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


