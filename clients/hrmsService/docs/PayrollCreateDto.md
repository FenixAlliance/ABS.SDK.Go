# PayrollCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**PayrollPeriodId** | **string** |  | 

## Methods

### NewPayrollCreateDto

`func NewPayrollCreateDto(payrollPeriodId string, ) *PayrollCreateDto`

NewPayrollCreateDto instantiates a new PayrollCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayrollCreateDtoWithDefaults

`func NewPayrollCreateDtoWithDefaults() *PayrollCreateDto`

NewPayrollCreateDtoWithDefaults instantiates a new PayrollCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayrollCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayrollCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayrollCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayrollCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PayrollCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PayrollCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PayrollCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PayrollCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPayrollPeriodId

`func (o *PayrollCreateDto) GetPayrollPeriodId() string`

GetPayrollPeriodId returns the PayrollPeriodId field if non-nil, zero value otherwise.

### GetPayrollPeriodIdOk

`func (o *PayrollCreateDto) GetPayrollPeriodIdOk() (*string, bool)`

GetPayrollPeriodIdOk returns a tuple with the PayrollPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollPeriodId

`func (o *PayrollCreateDto) SetPayrollPeriodId(v string)`

SetPayrollPeriodId sets PayrollPeriodId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


