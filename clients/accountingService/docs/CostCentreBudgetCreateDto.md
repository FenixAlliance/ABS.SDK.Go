# CostCentreBudgetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**CostCentreId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreBudgetCreateDto

`func NewCostCentreBudgetCreateDto() *CostCentreBudgetCreateDto`

NewCostCentreBudgetCreateDto instantiates a new CostCentreBudgetCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreBudgetCreateDtoWithDefaults

`func NewCostCentreBudgetCreateDtoWithDefaults() *CostCentreBudgetCreateDto`

NewCostCentreBudgetCreateDtoWithDefaults instantiates a new CostCentreBudgetCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreBudgetCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreBudgetCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreBudgetCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreBudgetCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CostCentreBudgetCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreBudgetCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreBudgetCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreBudgetCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CostCentreBudgetCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreBudgetCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreBudgetCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreBudgetCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreBudgetCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreBudgetCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *CostCentreBudgetCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreBudgetCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreBudgetCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreBudgetCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreBudgetCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreBudgetCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetFiscalYearId

`func (o *CostCentreBudgetCreateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *CostCentreBudgetCreateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *CostCentreBudgetCreateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *CostCentreBudgetCreateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *CostCentreBudgetCreateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *CostCentreBudgetCreateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetCostCentreId

`func (o *CostCentreBudgetCreateDto) GetCostCentreId() string`

GetCostCentreId returns the CostCentreId field if non-nil, zero value otherwise.

### GetCostCentreIdOk

`func (o *CostCentreBudgetCreateDto) GetCostCentreIdOk() (*string, bool)`

GetCostCentreIdOk returns a tuple with the CostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreId

`func (o *CostCentreBudgetCreateDto) SetCostCentreId(v string)`

SetCostCentreId sets CostCentreId field to given value.

### HasCostCentreId

`func (o *CostCentreBudgetCreateDto) HasCostCentreId() bool`

HasCostCentreId returns a boolean if a field has been set.

### SetCostCentreIdNil

`func (o *CostCentreBudgetCreateDto) SetCostCentreIdNil(b bool)`

 SetCostCentreIdNil sets the value for CostCentreId to be an explicit nil

### UnsetCostCentreId
`func (o *CostCentreBudgetCreateDto) UnsetCostCentreId()`

UnsetCostCentreId ensures that no value is present for CostCentreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


