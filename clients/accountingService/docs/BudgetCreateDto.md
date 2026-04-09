# BudgetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBudgetCreateDto

`func NewBudgetCreateDto() *BudgetCreateDto`

NewBudgetCreateDto instantiates a new BudgetCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetCreateDtoWithDefaults

`func NewBudgetCreateDtoWithDefaults() *BudgetCreateDto`

NewBudgetCreateDtoWithDefaults instantiates a new BudgetCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BudgetCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BudgetCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BudgetCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BudgetCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *BudgetCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BudgetCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BudgetCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BudgetCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BudgetCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BudgetCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *BudgetCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BudgetCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BudgetCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BudgetCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetFiscalYearId

`func (o *BudgetCreateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *BudgetCreateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *BudgetCreateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *BudgetCreateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *BudgetCreateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *BudgetCreateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


