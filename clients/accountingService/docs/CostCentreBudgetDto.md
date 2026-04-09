# CostCentreBudgetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**CostCentreId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreBudgetDto

`func NewCostCentreBudgetDto() *CostCentreBudgetDto`

NewCostCentreBudgetDto instantiates a new CostCentreBudgetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreBudgetDtoWithDefaults

`func NewCostCentreBudgetDtoWithDefaults() *CostCentreBudgetDto`

NewCostCentreBudgetDtoWithDefaults instantiates a new CostCentreBudgetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreBudgetDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreBudgetDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreBudgetDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreBudgetDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CostCentreBudgetDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CostCentreBudgetDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CostCentreBudgetDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreBudgetDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreBudgetDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreBudgetDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CostCentreBudgetDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CostCentreBudgetDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CostCentreBudgetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreBudgetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreBudgetDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreBudgetDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreBudgetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreBudgetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *CostCentreBudgetDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreBudgetDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreBudgetDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreBudgetDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreBudgetDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreBudgetDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetFiscalYearId

`func (o *CostCentreBudgetDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *CostCentreBudgetDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *CostCentreBudgetDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *CostCentreBudgetDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *CostCentreBudgetDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *CostCentreBudgetDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetCostCentreId

`func (o *CostCentreBudgetDto) GetCostCentreId() string`

GetCostCentreId returns the CostCentreId field if non-nil, zero value otherwise.

### GetCostCentreIdOk

`func (o *CostCentreBudgetDto) GetCostCentreIdOk() (*string, bool)`

GetCostCentreIdOk returns a tuple with the CostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreId

`func (o *CostCentreBudgetDto) SetCostCentreId(v string)`

SetCostCentreId sets CostCentreId field to given value.

### HasCostCentreId

`func (o *CostCentreBudgetDto) HasCostCentreId() bool`

HasCostCentreId returns a boolean if a field has been set.

### SetCostCentreIdNil

`func (o *CostCentreBudgetDto) SetCostCentreIdNil(b bool)`

 SetCostCentreIdNil sets the value for CostCentreId to be an explicit nil

### UnsetCostCentreId
`func (o *CostCentreBudgetDto) UnsetCostCentreId()`

UnsetCostCentreId ensures that no value is present for CostCentreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


