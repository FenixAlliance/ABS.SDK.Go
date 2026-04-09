# JournalEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Group** | Pointer to **bool** |  | [optional] 
**Opening** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableTime** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**Credit** | Pointer to **float64** |  | [optional] 
**Debit** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**JournalId** | Pointer to **NullableString** |  | [optional] 
**JournalName** | Pointer to **NullableString** |  | [optional] 
**JournalCode** | Pointer to **NullableString** |  | [optional] 
**CreditAccountId** | Pointer to **NullableString** |  | [optional] 
**CreditAccountName** | Pointer to **NullableString** |  | [optional] 
**DebitAccountId** | Pointer to **NullableString** |  | [optional] 
**DebitAccountName** | Pointer to **NullableString** |  | [optional] 
**InvoiceCode** | Pointer to **NullableString** |  | [optional] 
**ParentJournalEntryId** | Pointer to **NullableString** |  | [optional] 
**CreditAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**DebitAmount** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewJournalEntryDto

`func NewJournalEntryDto() *JournalEntryDto`

NewJournalEntryDto instantiates a new JournalEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryDtoWithDefaults

`func NewJournalEntryDtoWithDefaults() *JournalEntryDto`

NewJournalEntryDtoWithDefaults instantiates a new JournalEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JournalEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JournalEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JournalEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JournalEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JournalEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JournalEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JournalEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JournalEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JournalEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetGroup

`func (o *JournalEntryDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JournalEntryDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JournalEntryDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *JournalEntryDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOpening

`func (o *JournalEntryDto) GetOpening() bool`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *JournalEntryDto) GetOpeningOk() (*bool, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *JournalEntryDto) SetOpening(v bool)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *JournalEntryDto) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

### GetDescription

`func (o *JournalEntryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalEntryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalEntryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalEntryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JournalEntryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JournalEntryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDate

`func (o *JournalEntryDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *JournalEntryDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *JournalEntryDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *JournalEntryDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *JournalEntryDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *JournalEntryDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetForexRatesSnapshot

`func (o *JournalEntryDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *JournalEntryDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *JournalEntryDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *JournalEntryDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *JournalEntryDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *JournalEntryDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetForexRate

`func (o *JournalEntryDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *JournalEntryDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *JournalEntryDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *JournalEntryDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCredit

`func (o *JournalEntryDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *JournalEntryDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *JournalEntryDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *JournalEntryDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *JournalEntryDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *JournalEntryDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *JournalEntryDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *JournalEntryDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JournalEntryDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JournalEntryDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JournalEntryDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JournalEntryDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JournalEntryDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JournalEntryDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *JournalEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JournalEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JournalEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JournalEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JournalEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JournalEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *JournalEntryDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *JournalEntryDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *JournalEntryDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *JournalEntryDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *JournalEntryDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *JournalEntryDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetJournalId

`func (o *JournalEntryDto) GetJournalId() string`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *JournalEntryDto) GetJournalIdOk() (*string, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *JournalEntryDto) SetJournalId(v string)`

SetJournalId sets JournalId field to given value.

### HasJournalId

`func (o *JournalEntryDto) HasJournalId() bool`

HasJournalId returns a boolean if a field has been set.

### SetJournalIdNil

`func (o *JournalEntryDto) SetJournalIdNil(b bool)`

 SetJournalIdNil sets the value for JournalId to be an explicit nil

### UnsetJournalId
`func (o *JournalEntryDto) UnsetJournalId()`

UnsetJournalId ensures that no value is present for JournalId, not even an explicit nil
### GetJournalName

`func (o *JournalEntryDto) GetJournalName() string`

GetJournalName returns the JournalName field if non-nil, zero value otherwise.

### GetJournalNameOk

`func (o *JournalEntryDto) GetJournalNameOk() (*string, bool)`

GetJournalNameOk returns a tuple with the JournalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalName

`func (o *JournalEntryDto) SetJournalName(v string)`

SetJournalName sets JournalName field to given value.

### HasJournalName

`func (o *JournalEntryDto) HasJournalName() bool`

HasJournalName returns a boolean if a field has been set.

### SetJournalNameNil

`func (o *JournalEntryDto) SetJournalNameNil(b bool)`

 SetJournalNameNil sets the value for JournalName to be an explicit nil

### UnsetJournalName
`func (o *JournalEntryDto) UnsetJournalName()`

UnsetJournalName ensures that no value is present for JournalName, not even an explicit nil
### GetJournalCode

`func (o *JournalEntryDto) GetJournalCode() string`

GetJournalCode returns the JournalCode field if non-nil, zero value otherwise.

### GetJournalCodeOk

`func (o *JournalEntryDto) GetJournalCodeOk() (*string, bool)`

GetJournalCodeOk returns a tuple with the JournalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalCode

`func (o *JournalEntryDto) SetJournalCode(v string)`

SetJournalCode sets JournalCode field to given value.

### HasJournalCode

`func (o *JournalEntryDto) HasJournalCode() bool`

HasJournalCode returns a boolean if a field has been set.

### SetJournalCodeNil

`func (o *JournalEntryDto) SetJournalCodeNil(b bool)`

 SetJournalCodeNil sets the value for JournalCode to be an explicit nil

### UnsetJournalCode
`func (o *JournalEntryDto) UnsetJournalCode()`

UnsetJournalCode ensures that no value is present for JournalCode, not even an explicit nil
### GetCreditAccountId

`func (o *JournalEntryDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *JournalEntryDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *JournalEntryDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *JournalEntryDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *JournalEntryDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *JournalEntryDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
### GetCreditAccountName

`func (o *JournalEntryDto) GetCreditAccountName() string`

GetCreditAccountName returns the CreditAccountName field if non-nil, zero value otherwise.

### GetCreditAccountNameOk

`func (o *JournalEntryDto) GetCreditAccountNameOk() (*string, bool)`

GetCreditAccountNameOk returns a tuple with the CreditAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountName

`func (o *JournalEntryDto) SetCreditAccountName(v string)`

SetCreditAccountName sets CreditAccountName field to given value.

### HasCreditAccountName

`func (o *JournalEntryDto) HasCreditAccountName() bool`

HasCreditAccountName returns a boolean if a field has been set.

### SetCreditAccountNameNil

`func (o *JournalEntryDto) SetCreditAccountNameNil(b bool)`

 SetCreditAccountNameNil sets the value for CreditAccountName to be an explicit nil

### UnsetCreditAccountName
`func (o *JournalEntryDto) UnsetCreditAccountName()`

UnsetCreditAccountName ensures that no value is present for CreditAccountName, not even an explicit nil
### GetDebitAccountId

`func (o *JournalEntryDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *JournalEntryDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *JournalEntryDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.

### HasDebitAccountId

`func (o *JournalEntryDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *JournalEntryDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *JournalEntryDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
### GetDebitAccountName

`func (o *JournalEntryDto) GetDebitAccountName() string`

GetDebitAccountName returns the DebitAccountName field if non-nil, zero value otherwise.

### GetDebitAccountNameOk

`func (o *JournalEntryDto) GetDebitAccountNameOk() (*string, bool)`

GetDebitAccountNameOk returns a tuple with the DebitAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountName

`func (o *JournalEntryDto) SetDebitAccountName(v string)`

SetDebitAccountName sets DebitAccountName field to given value.

### HasDebitAccountName

`func (o *JournalEntryDto) HasDebitAccountName() bool`

HasDebitAccountName returns a boolean if a field has been set.

### SetDebitAccountNameNil

`func (o *JournalEntryDto) SetDebitAccountNameNil(b bool)`

 SetDebitAccountNameNil sets the value for DebitAccountName to be an explicit nil

### UnsetDebitAccountName
`func (o *JournalEntryDto) UnsetDebitAccountName()`

UnsetDebitAccountName ensures that no value is present for DebitAccountName, not even an explicit nil
### GetInvoiceCode

`func (o *JournalEntryDto) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *JournalEntryDto) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *JournalEntryDto) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.

### HasInvoiceCode

`func (o *JournalEntryDto) HasInvoiceCode() bool`

HasInvoiceCode returns a boolean if a field has been set.

### SetInvoiceCodeNil

`func (o *JournalEntryDto) SetInvoiceCodeNil(b bool)`

 SetInvoiceCodeNil sets the value for InvoiceCode to be an explicit nil

### UnsetInvoiceCode
`func (o *JournalEntryDto) UnsetInvoiceCode()`

UnsetInvoiceCode ensures that no value is present for InvoiceCode, not even an explicit nil
### GetParentJournalEntryId

`func (o *JournalEntryDto) GetParentJournalEntryId() string`

GetParentJournalEntryId returns the ParentJournalEntryId field if non-nil, zero value otherwise.

### GetParentJournalEntryIdOk

`func (o *JournalEntryDto) GetParentJournalEntryIdOk() (*string, bool)`

GetParentJournalEntryIdOk returns a tuple with the ParentJournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalEntryId

`func (o *JournalEntryDto) SetParentJournalEntryId(v string)`

SetParentJournalEntryId sets ParentJournalEntryId field to given value.

### HasParentJournalEntryId

`func (o *JournalEntryDto) HasParentJournalEntryId() bool`

HasParentJournalEntryId returns a boolean if a field has been set.

### SetParentJournalEntryIdNil

`func (o *JournalEntryDto) SetParentJournalEntryIdNil(b bool)`

 SetParentJournalEntryIdNil sets the value for ParentJournalEntryId to be an explicit nil

### UnsetParentJournalEntryId
`func (o *JournalEntryDto) UnsetParentJournalEntryId()`

UnsetParentJournalEntryId ensures that no value is present for ParentJournalEntryId, not even an explicit nil
### GetCreditAmount

`func (o *JournalEntryDto) GetCreditAmount() Money`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *JournalEntryDto) GetCreditAmountOk() (*Money, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *JournalEntryDto) SetCreditAmount(v Money)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *JournalEntryDto) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetDebitAmount

`func (o *JournalEntryDto) GetDebitAmount() Money`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *JournalEntryDto) GetDebitAmountOk() (*Money, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *JournalEntryDto) SetDebitAmount(v Money)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *JournalEntryDto) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


