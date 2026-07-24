# BudgetAccountEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 
**PlannedAmount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | **string** |  | 
**DebitAccountId** | **string** |  | 
**CreditAccountId** | **string** |  | 
**BudgetId** | **string** |  | 

## Methods

### NewBudgetAccountEntryCreateDto

`func NewBudgetAccountEntryCreateDto(description string, currencyId string, debitAccountId string, creditAccountId string, budgetId string, ) *BudgetAccountEntryCreateDto`

NewBudgetAccountEntryCreateDto instantiates a new BudgetAccountEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetAccountEntryCreateDtoWithDefaults

`func NewBudgetAccountEntryCreateDtoWithDefaults() *BudgetAccountEntryCreateDto`

NewBudgetAccountEntryCreateDtoWithDefaults instantiates a new BudgetAccountEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetAccountEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetAccountEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetAccountEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetAccountEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BudgetAccountEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BudgetAccountEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BudgetAccountEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BudgetAccountEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *BudgetAccountEntryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BudgetAccountEntryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BudgetAccountEntryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlannedAmount

`func (o *BudgetAccountEntryCreateDto) GetPlannedAmount() float64`

GetPlannedAmount returns the PlannedAmount field if non-nil, zero value otherwise.

### GetPlannedAmountOk

`func (o *BudgetAccountEntryCreateDto) GetPlannedAmountOk() (*float64, bool)`

GetPlannedAmountOk returns a tuple with the PlannedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedAmount

`func (o *BudgetAccountEntryCreateDto) SetPlannedAmount(v float64)`

SetPlannedAmount sets PlannedAmount field to given value.

### HasPlannedAmount

`func (o *BudgetAccountEntryCreateDto) HasPlannedAmount() bool`

HasPlannedAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *BudgetAccountEntryCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BudgetAccountEntryCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BudgetAccountEntryCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetDebitAccountId

`func (o *BudgetAccountEntryCreateDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *BudgetAccountEntryCreateDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *BudgetAccountEntryCreateDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.


### GetCreditAccountId

`func (o *BudgetAccountEntryCreateDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *BudgetAccountEntryCreateDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *BudgetAccountEntryCreateDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.


### GetBudgetId

`func (o *BudgetAccountEntryCreateDto) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetAccountEntryCreateDto) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetAccountEntryCreateDto) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


