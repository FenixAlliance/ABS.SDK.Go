# AccountingEntryDto

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

## Methods

### NewAccountingEntryDto

`func NewAccountingEntryDto() *AccountingEntryDto`

NewAccountingEntryDto instantiates a new AccountingEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingEntryDtoWithDefaults

`func NewAccountingEntryDtoWithDefaults() *AccountingEntryDto`

NewAccountingEntryDtoWithDefaults instantiates a new AccountingEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AccountingEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AccountingEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AccountingEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountingEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountingEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountingEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AccountingEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AccountingEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDebit

`func (o *AccountingEntryDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *AccountingEntryDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *AccountingEntryDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *AccountingEntryDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCredit

`func (o *AccountingEntryDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AccountingEntryDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AccountingEntryDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AccountingEntryDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDescription

`func (o *AccountingEntryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountingEntryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountingEntryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountingEntryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountingEntryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountingEntryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetForexRate

`func (o *AccountingEntryDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *AccountingEntryDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *AccountingEntryDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *AccountingEntryDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetAccountId

`func (o *AccountingEntryDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountingEntryDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountingEntryDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountingEntryDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AccountingEntryDto) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AccountingEntryDto) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetTenantId

`func (o *AccountingEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountingEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountingEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccountingEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AccountingEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AccountingEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDate

`func (o *AccountingEntryDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountingEntryDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountingEntryDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *AccountingEntryDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *AccountingEntryDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AccountingEntryDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetEnrollmentId

`func (o *AccountingEntryDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AccountingEntryDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AccountingEntryDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AccountingEntryDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AccountingEntryDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AccountingEntryDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *AccountingEntryDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountingEntryDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountingEntryDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AccountingEntryDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AccountingEntryDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AccountingEntryDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDebitAccountId

`func (o *AccountingEntryDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *AccountingEntryDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *AccountingEntryDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.

### HasDebitAccountId

`func (o *AccountingEntryDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *AccountingEntryDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *AccountingEntryDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
### GetCreditAccountId

`func (o *AccountingEntryDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *AccountingEntryDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *AccountingEntryDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *AccountingEntryDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *AccountingEntryDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *AccountingEntryDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
### GetJournalEntryId

`func (o *AccountingEntryDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *AccountingEntryDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *AccountingEntryDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.

### HasJournalEntryId

`func (o *AccountingEntryDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *AccountingEntryDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *AccountingEntryDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetDebitAccountName

`func (o *AccountingEntryDto) GetDebitAccountName() string`

GetDebitAccountName returns the DebitAccountName field if non-nil, zero value otherwise.

### GetDebitAccountNameOk

`func (o *AccountingEntryDto) GetDebitAccountNameOk() (*string, bool)`

GetDebitAccountNameOk returns a tuple with the DebitAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountName

`func (o *AccountingEntryDto) SetDebitAccountName(v string)`

SetDebitAccountName sets DebitAccountName field to given value.

### HasDebitAccountName

`func (o *AccountingEntryDto) HasDebitAccountName() bool`

HasDebitAccountName returns a boolean if a field has been set.

### SetDebitAccountNameNil

`func (o *AccountingEntryDto) SetDebitAccountNameNil(b bool)`

 SetDebitAccountNameNil sets the value for DebitAccountName to be an explicit nil

### UnsetDebitAccountName
`func (o *AccountingEntryDto) UnsetDebitAccountName()`

UnsetDebitAccountName ensures that no value is present for DebitAccountName, not even an explicit nil
### GetCreditAccountName

`func (o *AccountingEntryDto) GetCreditAccountName() string`

GetCreditAccountName returns the CreditAccountName field if non-nil, zero value otherwise.

### GetCreditAccountNameOk

`func (o *AccountingEntryDto) GetCreditAccountNameOk() (*string, bool)`

GetCreditAccountNameOk returns a tuple with the CreditAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountName

`func (o *AccountingEntryDto) SetCreditAccountName(v string)`

SetCreditAccountName sets CreditAccountName field to given value.

### HasCreditAccountName

`func (o *AccountingEntryDto) HasCreditAccountName() bool`

HasCreditAccountName returns a boolean if a field has been set.

### SetCreditAccountNameNil

`func (o *AccountingEntryDto) SetCreditAccountNameNil(b bool)`

 SetCreditAccountNameNil sets the value for CreditAccountName to be an explicit nil

### UnsetCreditAccountName
`func (o *AccountingEntryDto) UnsetCreditAccountName()`

UnsetCreditAccountName ensures that no value is present for CreditAccountName, not even an explicit nil
### GetAccountingEntryType

`func (o *AccountingEntryDto) GetAccountingEntryType() string`

GetAccountingEntryType returns the AccountingEntryType field if non-nil, zero value otherwise.

### GetAccountingEntryTypeOk

`func (o *AccountingEntryDto) GetAccountingEntryTypeOk() (*string, bool)`

GetAccountingEntryTypeOk returns a tuple with the AccountingEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryType

`func (o *AccountingEntryDto) SetAccountingEntryType(v string)`

SetAccountingEntryType sets AccountingEntryType field to given value.

### HasAccountingEntryType

`func (o *AccountingEntryDto) HasAccountingEntryType() bool`

HasAccountingEntryType returns a boolean if a field has been set.

### GetDebitAmount

`func (o *AccountingEntryDto) GetDebitAmount() Money`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *AccountingEntryDto) GetDebitAmountOk() (*Money, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *AccountingEntryDto) SetDebitAmount(v Money)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *AccountingEntryDto) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *AccountingEntryDto) GetCreditAmount() Money`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *AccountingEntryDto) GetCreditAmountOk() (*Money, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *AccountingEntryDto) SetCreditAmount(v Money)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *AccountingEntryDto) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


