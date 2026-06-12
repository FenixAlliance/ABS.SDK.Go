# SalaryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSalaryUpdateDto

`func NewSalaryUpdateDto() *SalaryUpdateDto`

NewSalaryUpdateDto instantiates a new SalaryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryUpdateDtoWithDefaults

`func NewSalaryUpdateDtoWithDefaults() *SalaryUpdateDto`

NewSalaryUpdateDtoWithDefaults instantiates a new SalaryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SalaryUpdateDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SalaryUpdateDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SalaryUpdateDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SalaryUpdateDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *SalaryUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *SalaryUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *SalaryUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *SalaryUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *SalaryUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *SalaryUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetEmployeeProfileId

`func (o *SalaryUpdateDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *SalaryUpdateDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *SalaryUpdateDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *SalaryUpdateDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *SalaryUpdateDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *SalaryUpdateDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


