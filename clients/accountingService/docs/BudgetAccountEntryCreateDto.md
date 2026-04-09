# BudgetAccountEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Date** | Pointer to **NullableTime** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | **string** |  | 
**DebitAccountId** | Pointer to **NullableString** |  | [optional] 
**CreditAccountId** | Pointer to **NullableString** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**AccountingEntryType** | Pointer to **string** |  | [optional] 
**BudgetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBudgetAccountEntryCreateDto

`func NewBudgetAccountEntryCreateDto(description string, currencyId string, ) *BudgetAccountEntryCreateDto`

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

### GetTenantId

`func (o *BudgetAccountEntryCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetAccountEntryCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetAccountEntryCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BudgetAccountEntryCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BudgetAccountEntryCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BudgetAccountEntryCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *BudgetAccountEntryCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BudgetAccountEntryCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BudgetAccountEntryCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BudgetAccountEntryCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BudgetAccountEntryCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BudgetAccountEntryCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
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


### GetDate

`func (o *BudgetAccountEntryCreateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BudgetAccountEntryCreateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BudgetAccountEntryCreateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BudgetAccountEntryCreateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *BudgetAccountEntryCreateDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *BudgetAccountEntryCreateDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetAmount

`func (o *BudgetAccountEntryCreateDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetAccountEntryCreateDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetAccountEntryCreateDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BudgetAccountEntryCreateDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

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

### HasDebitAccountId

`func (o *BudgetAccountEntryCreateDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *BudgetAccountEntryCreateDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *BudgetAccountEntryCreateDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
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

### HasCreditAccountId

`func (o *BudgetAccountEntryCreateDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *BudgetAccountEntryCreateDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *BudgetAccountEntryCreateDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
### GetJournalEntryId

`func (o *BudgetAccountEntryCreateDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *BudgetAccountEntryCreateDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *BudgetAccountEntryCreateDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.

### HasJournalEntryId

`func (o *BudgetAccountEntryCreateDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *BudgetAccountEntryCreateDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *BudgetAccountEntryCreateDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetAccountingEntryType

`func (o *BudgetAccountEntryCreateDto) GetAccountingEntryType() string`

GetAccountingEntryType returns the AccountingEntryType field if non-nil, zero value otherwise.

### GetAccountingEntryTypeOk

`func (o *BudgetAccountEntryCreateDto) GetAccountingEntryTypeOk() (*string, bool)`

GetAccountingEntryTypeOk returns a tuple with the AccountingEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryType

`func (o *BudgetAccountEntryCreateDto) SetAccountingEntryType(v string)`

SetAccountingEntryType sets AccountingEntryType field to given value.

### HasAccountingEntryType

`func (o *BudgetAccountEntryCreateDto) HasAccountingEntryType() bool`

HasAccountingEntryType returns a boolean if a field has been set.

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

### HasBudgetId

`func (o *BudgetAccountEntryCreateDto) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *BudgetAccountEntryCreateDto) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *BudgetAccountEntryCreateDto) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


