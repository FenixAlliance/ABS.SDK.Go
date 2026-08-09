# AccountingEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AccountName** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TransactionAmount** | Pointer to **float64** |  | [optional] 
**TransactionCurrencyId** | Pointer to **NullableString** |  | [optional] 
**FunctionalAmount** | Pointer to **float64** |  | [optional] 
**FunctionalCurrencyId** | Pointer to **NullableString** |  | [optional] 
**AccountAmount** | Pointer to **float64** |  | [optional] 
**AccountCurrencyId** | Pointer to **NullableString** |  | [optional] 
**ReportingAmountInUsd** | Pointer to **float64** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**CostCentreId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Debit** | Pointer to **float64** |  | [optional] [readonly] 
**Credit** | Pointer to **float64** |  | [optional] [readonly] 
**Amount** | Pointer to [**Money**](Money.md) |  | [optional] 
**AmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 

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
### GetAccountName

`func (o *AccountingEntryDto) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountingEntryDto) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountingEntryDto) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AccountingEntryDto) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *AccountingEntryDto) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *AccountingEntryDto) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetDirection

`func (o *AccountingEntryDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AccountingEntryDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AccountingEntryDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AccountingEntryDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

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
### GetTransactionAmount

`func (o *AccountingEntryDto) GetTransactionAmount() float64`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *AccountingEntryDto) GetTransactionAmountOk() (*float64, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *AccountingEntryDto) SetTransactionAmount(v float64)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *AccountingEntryDto) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionCurrencyId

`func (o *AccountingEntryDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *AccountingEntryDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *AccountingEntryDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.

### HasTransactionCurrencyId

`func (o *AccountingEntryDto) HasTransactionCurrencyId() bool`

HasTransactionCurrencyId returns a boolean if a field has been set.

### SetTransactionCurrencyIdNil

`func (o *AccountingEntryDto) SetTransactionCurrencyIdNil(b bool)`

 SetTransactionCurrencyIdNil sets the value for TransactionCurrencyId to be an explicit nil

### UnsetTransactionCurrencyId
`func (o *AccountingEntryDto) UnsetTransactionCurrencyId()`

UnsetTransactionCurrencyId ensures that no value is present for TransactionCurrencyId, not even an explicit nil
### GetFunctionalAmount

`func (o *AccountingEntryDto) GetFunctionalAmount() float64`

GetFunctionalAmount returns the FunctionalAmount field if non-nil, zero value otherwise.

### GetFunctionalAmountOk

`func (o *AccountingEntryDto) GetFunctionalAmountOk() (*float64, bool)`

GetFunctionalAmountOk returns a tuple with the FunctionalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalAmount

`func (o *AccountingEntryDto) SetFunctionalAmount(v float64)`

SetFunctionalAmount sets FunctionalAmount field to given value.

### HasFunctionalAmount

`func (o *AccountingEntryDto) HasFunctionalAmount() bool`

HasFunctionalAmount returns a boolean if a field has been set.

### GetFunctionalCurrencyId

`func (o *AccountingEntryDto) GetFunctionalCurrencyId() string`

GetFunctionalCurrencyId returns the FunctionalCurrencyId field if non-nil, zero value otherwise.

### GetFunctionalCurrencyIdOk

`func (o *AccountingEntryDto) GetFunctionalCurrencyIdOk() (*string, bool)`

GetFunctionalCurrencyIdOk returns a tuple with the FunctionalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalCurrencyId

`func (o *AccountingEntryDto) SetFunctionalCurrencyId(v string)`

SetFunctionalCurrencyId sets FunctionalCurrencyId field to given value.

### HasFunctionalCurrencyId

`func (o *AccountingEntryDto) HasFunctionalCurrencyId() bool`

HasFunctionalCurrencyId returns a boolean if a field has been set.

### SetFunctionalCurrencyIdNil

`func (o *AccountingEntryDto) SetFunctionalCurrencyIdNil(b bool)`

 SetFunctionalCurrencyIdNil sets the value for FunctionalCurrencyId to be an explicit nil

### UnsetFunctionalCurrencyId
`func (o *AccountingEntryDto) UnsetFunctionalCurrencyId()`

UnsetFunctionalCurrencyId ensures that no value is present for FunctionalCurrencyId, not even an explicit nil
### GetAccountAmount

`func (o *AccountingEntryDto) GetAccountAmount() float64`

GetAccountAmount returns the AccountAmount field if non-nil, zero value otherwise.

### GetAccountAmountOk

`func (o *AccountingEntryDto) GetAccountAmountOk() (*float64, bool)`

GetAccountAmountOk returns a tuple with the AccountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAmount

`func (o *AccountingEntryDto) SetAccountAmount(v float64)`

SetAccountAmount sets AccountAmount field to given value.

### HasAccountAmount

`func (o *AccountingEntryDto) HasAccountAmount() bool`

HasAccountAmount returns a boolean if a field has been set.

### GetAccountCurrencyId

`func (o *AccountingEntryDto) GetAccountCurrencyId() string`

GetAccountCurrencyId returns the AccountCurrencyId field if non-nil, zero value otherwise.

### GetAccountCurrencyIdOk

`func (o *AccountingEntryDto) GetAccountCurrencyIdOk() (*string, bool)`

GetAccountCurrencyIdOk returns a tuple with the AccountCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCurrencyId

`func (o *AccountingEntryDto) SetAccountCurrencyId(v string)`

SetAccountCurrencyId sets AccountCurrencyId field to given value.

### HasAccountCurrencyId

`func (o *AccountingEntryDto) HasAccountCurrencyId() bool`

HasAccountCurrencyId returns a boolean if a field has been set.

### SetAccountCurrencyIdNil

`func (o *AccountingEntryDto) SetAccountCurrencyIdNil(b bool)`

 SetAccountCurrencyIdNil sets the value for AccountCurrencyId to be an explicit nil

### UnsetAccountCurrencyId
`func (o *AccountingEntryDto) UnsetAccountCurrencyId()`

UnsetAccountCurrencyId ensures that no value is present for AccountCurrencyId, not even an explicit nil
### GetReportingAmountInUsd

`func (o *AccountingEntryDto) GetReportingAmountInUsd() float64`

GetReportingAmountInUsd returns the ReportingAmountInUsd field if non-nil, zero value otherwise.

### GetReportingAmountInUsdOk

`func (o *AccountingEntryDto) GetReportingAmountInUsdOk() (*float64, bool)`

GetReportingAmountInUsdOk returns a tuple with the ReportingAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingAmountInUsd

`func (o *AccountingEntryDto) SetReportingAmountInUsd(v float64)`

SetReportingAmountInUsd sets ReportingAmountInUsd field to given value.

### HasReportingAmountInUsd

`func (o *AccountingEntryDto) HasReportingAmountInUsd() bool`

HasReportingAmountInUsd returns a boolean if a field has been set.

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

### GetForexRatesSnapshot

`func (o *AccountingEntryDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *AccountingEntryDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *AccountingEntryDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *AccountingEntryDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *AccountingEntryDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *AccountingEntryDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetCostCentreId

`func (o *AccountingEntryDto) GetCostCentreId() string`

GetCostCentreId returns the CostCentreId field if non-nil, zero value otherwise.

### GetCostCentreIdOk

`func (o *AccountingEntryDto) GetCostCentreIdOk() (*string, bool)`

GetCostCentreIdOk returns a tuple with the CostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreId

`func (o *AccountingEntryDto) SetCostCentreId(v string)`

SetCostCentreId sets CostCentreId field to given value.

### HasCostCentreId

`func (o *AccountingEntryDto) HasCostCentreId() bool`

HasCostCentreId returns a boolean if a field has been set.

### SetCostCentreIdNil

`func (o *AccountingEntryDto) SetCostCentreIdNil(b bool)`

 SetCostCentreIdNil sets the value for CostCentreId to be an explicit nil

### UnsetCostCentreId
`func (o *AccountingEntryDto) UnsetCostCentreId()`

UnsetCostCentreId ensures that no value is present for CostCentreId, not even an explicit nil
### GetProjectId

`func (o *AccountingEntryDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AccountingEntryDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AccountingEntryDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AccountingEntryDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *AccountingEntryDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AccountingEntryDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
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

### GetAmount

`func (o *AccountingEntryDto) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountingEntryDto) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountingEntryDto) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountingEntryDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountInUsd

`func (o *AccountingEntryDto) GetAmountInUsd() Money`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *AccountingEntryDto) GetAmountInUsdOk() (*Money, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *AccountingEntryDto) SetAmountInUsd(v Money)`

SetAmountInUsd sets AmountInUsd field to given value.

### HasAmountInUsd

`func (o *AccountingEntryDto) HasAmountInUsd() bool`

HasAmountInUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


