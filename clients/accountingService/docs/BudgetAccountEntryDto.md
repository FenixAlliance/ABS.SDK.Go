# BudgetAccountEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Debit** | Pointer to **float64** |  | [optional] 
**Credit** | Pointer to **float64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableTime** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**DebitAccountId** | Pointer to **NullableString** |  | [optional] 
**CreditAccountId** | Pointer to **NullableString** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**DebitAccountName** | Pointer to **NullableString** |  | [optional] 
**CreditAccountName** | Pointer to **NullableString** |  | [optional] 
**AccountingEntryType** | Pointer to **string** |  | [optional] 
**DebitAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**CreditAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**BudgetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBudgetAccountEntryDto

`func NewBudgetAccountEntryDto() *BudgetAccountEntryDto`

NewBudgetAccountEntryDto instantiates a new BudgetAccountEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetAccountEntryDtoWithDefaults

`func NewBudgetAccountEntryDtoWithDefaults() *BudgetAccountEntryDto`

NewBudgetAccountEntryDtoWithDefaults instantiates a new BudgetAccountEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetAccountEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetAccountEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetAccountEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetAccountEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BudgetAccountEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BudgetAccountEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BudgetAccountEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BudgetAccountEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BudgetAccountEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BudgetAccountEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BudgetAccountEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BudgetAccountEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDebit

`func (o *BudgetAccountEntryDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *BudgetAccountEntryDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *BudgetAccountEntryDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *BudgetAccountEntryDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCredit

`func (o *BudgetAccountEntryDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BudgetAccountEntryDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BudgetAccountEntryDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BudgetAccountEntryDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDescription

`func (o *BudgetAccountEntryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BudgetAccountEntryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BudgetAccountEntryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BudgetAccountEntryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BudgetAccountEntryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BudgetAccountEntryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetForexRate

`func (o *BudgetAccountEntryDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *BudgetAccountEntryDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *BudgetAccountEntryDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *BudgetAccountEntryDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetAccountId

`func (o *BudgetAccountEntryDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BudgetAccountEntryDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BudgetAccountEntryDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BudgetAccountEntryDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *BudgetAccountEntryDto) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BudgetAccountEntryDto) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetTenantId

`func (o *BudgetAccountEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetAccountEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetAccountEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BudgetAccountEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BudgetAccountEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BudgetAccountEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDate

`func (o *BudgetAccountEntryDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BudgetAccountEntryDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BudgetAccountEntryDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BudgetAccountEntryDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *BudgetAccountEntryDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *BudgetAccountEntryDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetEnrollmentId

`func (o *BudgetAccountEntryDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BudgetAccountEntryDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BudgetAccountEntryDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BudgetAccountEntryDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BudgetAccountEntryDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BudgetAccountEntryDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *BudgetAccountEntryDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BudgetAccountEntryDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BudgetAccountEntryDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BudgetAccountEntryDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BudgetAccountEntryDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BudgetAccountEntryDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDebitAccountId

`func (o *BudgetAccountEntryDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *BudgetAccountEntryDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *BudgetAccountEntryDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.

### HasDebitAccountId

`func (o *BudgetAccountEntryDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *BudgetAccountEntryDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *BudgetAccountEntryDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
### GetCreditAccountId

`func (o *BudgetAccountEntryDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *BudgetAccountEntryDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *BudgetAccountEntryDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *BudgetAccountEntryDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *BudgetAccountEntryDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *BudgetAccountEntryDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
### GetJournalEntryId

`func (o *BudgetAccountEntryDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *BudgetAccountEntryDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *BudgetAccountEntryDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.

### HasJournalEntryId

`func (o *BudgetAccountEntryDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *BudgetAccountEntryDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *BudgetAccountEntryDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetDebitAccountName

`func (o *BudgetAccountEntryDto) GetDebitAccountName() string`

GetDebitAccountName returns the DebitAccountName field if non-nil, zero value otherwise.

### GetDebitAccountNameOk

`func (o *BudgetAccountEntryDto) GetDebitAccountNameOk() (*string, bool)`

GetDebitAccountNameOk returns a tuple with the DebitAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountName

`func (o *BudgetAccountEntryDto) SetDebitAccountName(v string)`

SetDebitAccountName sets DebitAccountName field to given value.

### HasDebitAccountName

`func (o *BudgetAccountEntryDto) HasDebitAccountName() bool`

HasDebitAccountName returns a boolean if a field has been set.

### SetDebitAccountNameNil

`func (o *BudgetAccountEntryDto) SetDebitAccountNameNil(b bool)`

 SetDebitAccountNameNil sets the value for DebitAccountName to be an explicit nil

### UnsetDebitAccountName
`func (o *BudgetAccountEntryDto) UnsetDebitAccountName()`

UnsetDebitAccountName ensures that no value is present for DebitAccountName, not even an explicit nil
### GetCreditAccountName

`func (o *BudgetAccountEntryDto) GetCreditAccountName() string`

GetCreditAccountName returns the CreditAccountName field if non-nil, zero value otherwise.

### GetCreditAccountNameOk

`func (o *BudgetAccountEntryDto) GetCreditAccountNameOk() (*string, bool)`

GetCreditAccountNameOk returns a tuple with the CreditAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountName

`func (o *BudgetAccountEntryDto) SetCreditAccountName(v string)`

SetCreditAccountName sets CreditAccountName field to given value.

### HasCreditAccountName

`func (o *BudgetAccountEntryDto) HasCreditAccountName() bool`

HasCreditAccountName returns a boolean if a field has been set.

### SetCreditAccountNameNil

`func (o *BudgetAccountEntryDto) SetCreditAccountNameNil(b bool)`

 SetCreditAccountNameNil sets the value for CreditAccountName to be an explicit nil

### UnsetCreditAccountName
`func (o *BudgetAccountEntryDto) UnsetCreditAccountName()`

UnsetCreditAccountName ensures that no value is present for CreditAccountName, not even an explicit nil
### GetAccountingEntryType

`func (o *BudgetAccountEntryDto) GetAccountingEntryType() string`

GetAccountingEntryType returns the AccountingEntryType field if non-nil, zero value otherwise.

### GetAccountingEntryTypeOk

`func (o *BudgetAccountEntryDto) GetAccountingEntryTypeOk() (*string, bool)`

GetAccountingEntryTypeOk returns a tuple with the AccountingEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryType

`func (o *BudgetAccountEntryDto) SetAccountingEntryType(v string)`

SetAccountingEntryType sets AccountingEntryType field to given value.

### HasAccountingEntryType

`func (o *BudgetAccountEntryDto) HasAccountingEntryType() bool`

HasAccountingEntryType returns a boolean if a field has been set.

### GetDebitAmount

`func (o *BudgetAccountEntryDto) GetDebitAmount() Money`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *BudgetAccountEntryDto) GetDebitAmountOk() (*Money, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *BudgetAccountEntryDto) SetDebitAmount(v Money)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *BudgetAccountEntryDto) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *BudgetAccountEntryDto) GetCreditAmount() Money`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *BudgetAccountEntryDto) GetCreditAmountOk() (*Money, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *BudgetAccountEntryDto) SetCreditAmount(v Money)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *BudgetAccountEntryDto) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetBudgetId

`func (o *BudgetAccountEntryDto) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetAccountEntryDto) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetAccountEntryDto) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *BudgetAccountEntryDto) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *BudgetAccountEntryDto) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *BudgetAccountEntryDto) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


