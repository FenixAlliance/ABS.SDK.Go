# TrialBalanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalPeriodId** | Pointer to **NullableString** |  | [optional] 
**FinancialBookId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Rows** | Pointer to [**[]TrialBalanceRowDto**](TrialBalanceRowDto.md) |  | [optional] 
**TotalDebit** | Pointer to **float64** |  | [optional] 
**TotalCredit** | Pointer to **float64** |  | [optional] 
**IsBalanced** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrialBalanceDto

`func NewTrialBalanceDto() *TrialBalanceDto`

NewTrialBalanceDto instantiates a new TrialBalanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBalanceDtoWithDefaults

`func NewTrialBalanceDtoWithDefaults() *TrialBalanceDto`

NewTrialBalanceDtoWithDefaults instantiates a new TrialBalanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalPeriodId

`func (o *TrialBalanceDto) GetFiscalPeriodId() string`

GetFiscalPeriodId returns the FiscalPeriodId field if non-nil, zero value otherwise.

### GetFiscalPeriodIdOk

`func (o *TrialBalanceDto) GetFiscalPeriodIdOk() (*string, bool)`

GetFiscalPeriodIdOk returns a tuple with the FiscalPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalPeriodId

`func (o *TrialBalanceDto) SetFiscalPeriodId(v string)`

SetFiscalPeriodId sets FiscalPeriodId field to given value.

### HasFiscalPeriodId

`func (o *TrialBalanceDto) HasFiscalPeriodId() bool`

HasFiscalPeriodId returns a boolean if a field has been set.

### SetFiscalPeriodIdNil

`func (o *TrialBalanceDto) SetFiscalPeriodIdNil(b bool)`

 SetFiscalPeriodIdNil sets the value for FiscalPeriodId to be an explicit nil

### UnsetFiscalPeriodId
`func (o *TrialBalanceDto) UnsetFiscalPeriodId()`

UnsetFiscalPeriodId ensures that no value is present for FiscalPeriodId, not even an explicit nil
### GetFinancialBookId

`func (o *TrialBalanceDto) GetFinancialBookId() string`

GetFinancialBookId returns the FinancialBookId field if non-nil, zero value otherwise.

### GetFinancialBookIdOk

`func (o *TrialBalanceDto) GetFinancialBookIdOk() (*string, bool)`

GetFinancialBookIdOk returns a tuple with the FinancialBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBookId

`func (o *TrialBalanceDto) SetFinancialBookId(v string)`

SetFinancialBookId sets FinancialBookId field to given value.

### HasFinancialBookId

`func (o *TrialBalanceDto) HasFinancialBookId() bool`

HasFinancialBookId returns a boolean if a field has been set.

### SetFinancialBookIdNil

`func (o *TrialBalanceDto) SetFinancialBookIdNil(b bool)`

 SetFinancialBookIdNil sets the value for FinancialBookId to be an explicit nil

### UnsetFinancialBookId
`func (o *TrialBalanceDto) UnsetFinancialBookId()`

UnsetFinancialBookId ensures that no value is present for FinancialBookId, not even an explicit nil
### GetCurrencyId

`func (o *TrialBalanceDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TrialBalanceDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TrialBalanceDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TrialBalanceDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TrialBalanceDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TrialBalanceDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetRows

`func (o *TrialBalanceDto) GetRows() []TrialBalanceRowDto`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TrialBalanceDto) GetRowsOk() (*[]TrialBalanceRowDto, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TrialBalanceDto) SetRows(v []TrialBalanceRowDto)`

SetRows sets Rows field to given value.

### HasRows

`func (o *TrialBalanceDto) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRowsNil

`func (o *TrialBalanceDto) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *TrialBalanceDto) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil
### GetTotalDebit

`func (o *TrialBalanceDto) GetTotalDebit() float64`

GetTotalDebit returns the TotalDebit field if non-nil, zero value otherwise.

### GetTotalDebitOk

`func (o *TrialBalanceDto) GetTotalDebitOk() (*float64, bool)`

GetTotalDebitOk returns a tuple with the TotalDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebit

`func (o *TrialBalanceDto) SetTotalDebit(v float64)`

SetTotalDebit sets TotalDebit field to given value.

### HasTotalDebit

`func (o *TrialBalanceDto) HasTotalDebit() bool`

HasTotalDebit returns a boolean if a field has been set.

### GetTotalCredit

`func (o *TrialBalanceDto) GetTotalCredit() float64`

GetTotalCredit returns the TotalCredit field if non-nil, zero value otherwise.

### GetTotalCreditOk

`func (o *TrialBalanceDto) GetTotalCreditOk() (*float64, bool)`

GetTotalCreditOk returns a tuple with the TotalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredit

`func (o *TrialBalanceDto) SetTotalCredit(v float64)`

SetTotalCredit sets TotalCredit field to given value.

### HasTotalCredit

`func (o *TrialBalanceDto) HasTotalCredit() bool`

HasTotalCredit returns a boolean if a field has been set.

### GetIsBalanced

`func (o *TrialBalanceDto) GetIsBalanced() bool`

GetIsBalanced returns the IsBalanced field if non-nil, zero value otherwise.

### GetIsBalancedOk

`func (o *TrialBalanceDto) GetIsBalancedOk() (*bool, bool)`

GetIsBalancedOk returns a tuple with the IsBalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBalanced

`func (o *TrialBalanceDto) SetIsBalanced(v bool)`

SetIsBalanced sets IsBalanced field to given value.

### HasIsBalanced

`func (o *TrialBalanceDto) HasIsBalanced() bool`

HasIsBalanced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


