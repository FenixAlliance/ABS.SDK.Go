# JournalEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**JournalId** | Pointer to **NullableString** |  | [optional] 
**JournalName** | Pointer to **NullableString** |  | [optional] 
**JournalCode** | Pointer to **NullableString** |  | [optional] 
**FiscalPeriodId** | Pointer to **NullableString** |  | [optional] 
**FinancialBookId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EntryType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PostingDate** | Pointer to **NullableTime** |  | [optional] 
**IsOpeningBalance** | Pointer to **bool** |  | [optional] 
**TransactionCurrencyId** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentType** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentId** | Pointer to **NullableString** |  | [optional] 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**ReversalOfJournalEntryId** | Pointer to **NullableString** |  | [optional] 
**PostedBy** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**DebitInUsd** | Pointer to **float64** |  | [optional] 
**CreditInUsd** | Pointer to **float64** |  | [optional] 
**AccountingEntries** | Pointer to [**[]AccountingEntryDto**](AccountingEntryDto.md) |  | [optional] 
**TotalDebit** | Pointer to **float64** |  | [optional] [readonly] 
**TotalCredit** | Pointer to **float64** |  | [optional] [readonly] 
**TotalDebitAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalCreditAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**DebitInUsdAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**CreditInUsdAmount** | Pointer to [**Money**](Money.md) |  | [optional] 

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
### GetFiscalPeriodId

`func (o *JournalEntryDto) GetFiscalPeriodId() string`

GetFiscalPeriodId returns the FiscalPeriodId field if non-nil, zero value otherwise.

### GetFiscalPeriodIdOk

`func (o *JournalEntryDto) GetFiscalPeriodIdOk() (*string, bool)`

GetFiscalPeriodIdOk returns a tuple with the FiscalPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalPeriodId

`func (o *JournalEntryDto) SetFiscalPeriodId(v string)`

SetFiscalPeriodId sets FiscalPeriodId field to given value.

### HasFiscalPeriodId

`func (o *JournalEntryDto) HasFiscalPeriodId() bool`

HasFiscalPeriodId returns a boolean if a field has been set.

### SetFiscalPeriodIdNil

`func (o *JournalEntryDto) SetFiscalPeriodIdNil(b bool)`

 SetFiscalPeriodIdNil sets the value for FiscalPeriodId to be an explicit nil

### UnsetFiscalPeriodId
`func (o *JournalEntryDto) UnsetFiscalPeriodId()`

UnsetFiscalPeriodId ensures that no value is present for FiscalPeriodId, not even an explicit nil
### GetFinancialBookId

`func (o *JournalEntryDto) GetFinancialBookId() string`

GetFinancialBookId returns the FinancialBookId field if non-nil, zero value otherwise.

### GetFinancialBookIdOk

`func (o *JournalEntryDto) GetFinancialBookIdOk() (*string, bool)`

GetFinancialBookIdOk returns a tuple with the FinancialBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBookId

`func (o *JournalEntryDto) SetFinancialBookId(v string)`

SetFinancialBookId sets FinancialBookId field to given value.

### HasFinancialBookId

`func (o *JournalEntryDto) HasFinancialBookId() bool`

HasFinancialBookId returns a boolean if a field has been set.

### SetFinancialBookIdNil

`func (o *JournalEntryDto) SetFinancialBookIdNil(b bool)`

 SetFinancialBookIdNil sets the value for FinancialBookId to be an explicit nil

### UnsetFinancialBookId
`func (o *JournalEntryDto) UnsetFinancialBookId()`

UnsetFinancialBookId ensures that no value is present for FinancialBookId, not even an explicit nil
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
### GetEntryType

`func (o *JournalEntryDto) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalEntryDto) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalEntryDto) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *JournalEntryDto) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetStatus

`func (o *JournalEntryDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalEntryDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalEntryDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JournalEntryDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPostingDate

`func (o *JournalEntryDto) GetPostingDate() time.Time`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *JournalEntryDto) GetPostingDateOk() (*time.Time, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *JournalEntryDto) SetPostingDate(v time.Time)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *JournalEntryDto) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### SetPostingDateNil

`func (o *JournalEntryDto) SetPostingDateNil(b bool)`

 SetPostingDateNil sets the value for PostingDate to be an explicit nil

### UnsetPostingDate
`func (o *JournalEntryDto) UnsetPostingDate()`

UnsetPostingDate ensures that no value is present for PostingDate, not even an explicit nil
### GetIsOpeningBalance

`func (o *JournalEntryDto) GetIsOpeningBalance() bool`

GetIsOpeningBalance returns the IsOpeningBalance field if non-nil, zero value otherwise.

### GetIsOpeningBalanceOk

`func (o *JournalEntryDto) GetIsOpeningBalanceOk() (*bool, bool)`

GetIsOpeningBalanceOk returns a tuple with the IsOpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpeningBalance

`func (o *JournalEntryDto) SetIsOpeningBalance(v bool)`

SetIsOpeningBalance sets IsOpeningBalance field to given value.

### HasIsOpeningBalance

`func (o *JournalEntryDto) HasIsOpeningBalance() bool`

HasIsOpeningBalance returns a boolean if a field has been set.

### GetTransactionCurrencyId

`func (o *JournalEntryDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *JournalEntryDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *JournalEntryDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.

### HasTransactionCurrencyId

`func (o *JournalEntryDto) HasTransactionCurrencyId() bool`

HasTransactionCurrencyId returns a boolean if a field has been set.

### SetTransactionCurrencyIdNil

`func (o *JournalEntryDto) SetTransactionCurrencyIdNil(b bool)`

 SetTransactionCurrencyIdNil sets the value for TransactionCurrencyId to be an explicit nil

### UnsetTransactionCurrencyId
`func (o *JournalEntryDto) UnsetTransactionCurrencyId()`

UnsetTransactionCurrencyId ensures that no value is present for TransactionCurrencyId, not even an explicit nil
### GetSourceDocumentType

`func (o *JournalEntryDto) GetSourceDocumentType() string`

GetSourceDocumentType returns the SourceDocumentType field if non-nil, zero value otherwise.

### GetSourceDocumentTypeOk

`func (o *JournalEntryDto) GetSourceDocumentTypeOk() (*string, bool)`

GetSourceDocumentTypeOk returns a tuple with the SourceDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentType

`func (o *JournalEntryDto) SetSourceDocumentType(v string)`

SetSourceDocumentType sets SourceDocumentType field to given value.

### HasSourceDocumentType

`func (o *JournalEntryDto) HasSourceDocumentType() bool`

HasSourceDocumentType returns a boolean if a field has been set.

### SetSourceDocumentTypeNil

`func (o *JournalEntryDto) SetSourceDocumentTypeNil(b bool)`

 SetSourceDocumentTypeNil sets the value for SourceDocumentType to be an explicit nil

### UnsetSourceDocumentType
`func (o *JournalEntryDto) UnsetSourceDocumentType()`

UnsetSourceDocumentType ensures that no value is present for SourceDocumentType, not even an explicit nil
### GetSourceDocumentId

`func (o *JournalEntryDto) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *JournalEntryDto) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *JournalEntryDto) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *JournalEntryDto) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### SetSourceDocumentIdNil

`func (o *JournalEntryDto) SetSourceDocumentIdNil(b bool)`

 SetSourceDocumentIdNil sets the value for SourceDocumentId to be an explicit nil

### UnsetSourceDocumentId
`func (o *JournalEntryDto) UnsetSourceDocumentId()`

UnsetSourceDocumentId ensures that no value is present for SourceDocumentId, not even an explicit nil
### GetIdempotencyKey

`func (o *JournalEntryDto) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *JournalEntryDto) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *JournalEntryDto) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *JournalEntryDto) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *JournalEntryDto) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *JournalEntryDto) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetReversalOfJournalEntryId

`func (o *JournalEntryDto) GetReversalOfJournalEntryId() string`

GetReversalOfJournalEntryId returns the ReversalOfJournalEntryId field if non-nil, zero value otherwise.

### GetReversalOfJournalEntryIdOk

`func (o *JournalEntryDto) GetReversalOfJournalEntryIdOk() (*string, bool)`

GetReversalOfJournalEntryIdOk returns a tuple with the ReversalOfJournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversalOfJournalEntryId

`func (o *JournalEntryDto) SetReversalOfJournalEntryId(v string)`

SetReversalOfJournalEntryId sets ReversalOfJournalEntryId field to given value.

### HasReversalOfJournalEntryId

`func (o *JournalEntryDto) HasReversalOfJournalEntryId() bool`

HasReversalOfJournalEntryId returns a boolean if a field has been set.

### SetReversalOfJournalEntryIdNil

`func (o *JournalEntryDto) SetReversalOfJournalEntryIdNil(b bool)`

 SetReversalOfJournalEntryIdNil sets the value for ReversalOfJournalEntryId to be an explicit nil

### UnsetReversalOfJournalEntryId
`func (o *JournalEntryDto) UnsetReversalOfJournalEntryId()`

UnsetReversalOfJournalEntryId ensures that no value is present for ReversalOfJournalEntryId, not even an explicit nil
### GetPostedBy

`func (o *JournalEntryDto) GetPostedBy() string`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *JournalEntryDto) GetPostedByOk() (*string, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *JournalEntryDto) SetPostedBy(v string)`

SetPostedBy sets PostedBy field to given value.

### HasPostedBy

`func (o *JournalEntryDto) HasPostedBy() bool`

HasPostedBy returns a boolean if a field has been set.

### SetPostedByNil

`func (o *JournalEntryDto) SetPostedByNil(b bool)`

 SetPostedByNil sets the value for PostedBy to be an explicit nil

### UnsetPostedBy
`func (o *JournalEntryDto) UnsetPostedBy()`

UnsetPostedBy ensures that no value is present for PostedBy, not even an explicit nil
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
### GetDebitInUsd

`func (o *JournalEntryDto) GetDebitInUsd() float64`

GetDebitInUsd returns the DebitInUsd field if non-nil, zero value otherwise.

### GetDebitInUsdOk

`func (o *JournalEntryDto) GetDebitInUsdOk() (*float64, bool)`

GetDebitInUsdOk returns a tuple with the DebitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitInUsd

`func (o *JournalEntryDto) SetDebitInUsd(v float64)`

SetDebitInUsd sets DebitInUsd field to given value.

### HasDebitInUsd

`func (o *JournalEntryDto) HasDebitInUsd() bool`

HasDebitInUsd returns a boolean if a field has been set.

### GetCreditInUsd

`func (o *JournalEntryDto) GetCreditInUsd() float64`

GetCreditInUsd returns the CreditInUsd field if non-nil, zero value otherwise.

### GetCreditInUsdOk

`func (o *JournalEntryDto) GetCreditInUsdOk() (*float64, bool)`

GetCreditInUsdOk returns a tuple with the CreditInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditInUsd

`func (o *JournalEntryDto) SetCreditInUsd(v float64)`

SetCreditInUsd sets CreditInUsd field to given value.

### HasCreditInUsd

`func (o *JournalEntryDto) HasCreditInUsd() bool`

HasCreditInUsd returns a boolean if a field has been set.

### GetAccountingEntries

`func (o *JournalEntryDto) GetAccountingEntries() []AccountingEntryDto`

GetAccountingEntries returns the AccountingEntries field if non-nil, zero value otherwise.

### GetAccountingEntriesOk

`func (o *JournalEntryDto) GetAccountingEntriesOk() (*[]AccountingEntryDto, bool)`

GetAccountingEntriesOk returns a tuple with the AccountingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntries

`func (o *JournalEntryDto) SetAccountingEntries(v []AccountingEntryDto)`

SetAccountingEntries sets AccountingEntries field to given value.

### HasAccountingEntries

`func (o *JournalEntryDto) HasAccountingEntries() bool`

HasAccountingEntries returns a boolean if a field has been set.

### SetAccountingEntriesNil

`func (o *JournalEntryDto) SetAccountingEntriesNil(b bool)`

 SetAccountingEntriesNil sets the value for AccountingEntries to be an explicit nil

### UnsetAccountingEntries
`func (o *JournalEntryDto) UnsetAccountingEntries()`

UnsetAccountingEntries ensures that no value is present for AccountingEntries, not even an explicit nil
### GetTotalDebit

`func (o *JournalEntryDto) GetTotalDebit() float64`

GetTotalDebit returns the TotalDebit field if non-nil, zero value otherwise.

### GetTotalDebitOk

`func (o *JournalEntryDto) GetTotalDebitOk() (*float64, bool)`

GetTotalDebitOk returns a tuple with the TotalDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebit

`func (o *JournalEntryDto) SetTotalDebit(v float64)`

SetTotalDebit sets TotalDebit field to given value.

### HasTotalDebit

`func (o *JournalEntryDto) HasTotalDebit() bool`

HasTotalDebit returns a boolean if a field has been set.

### GetTotalCredit

`func (o *JournalEntryDto) GetTotalCredit() float64`

GetTotalCredit returns the TotalCredit field if non-nil, zero value otherwise.

### GetTotalCreditOk

`func (o *JournalEntryDto) GetTotalCreditOk() (*float64, bool)`

GetTotalCreditOk returns a tuple with the TotalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredit

`func (o *JournalEntryDto) SetTotalCredit(v float64)`

SetTotalCredit sets TotalCredit field to given value.

### HasTotalCredit

`func (o *JournalEntryDto) HasTotalCredit() bool`

HasTotalCredit returns a boolean if a field has been set.

### GetTotalDebitAmount

`func (o *JournalEntryDto) GetTotalDebitAmount() Money`

GetTotalDebitAmount returns the TotalDebitAmount field if non-nil, zero value otherwise.

### GetTotalDebitAmountOk

`func (o *JournalEntryDto) GetTotalDebitAmountOk() (*Money, bool)`

GetTotalDebitAmountOk returns a tuple with the TotalDebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebitAmount

`func (o *JournalEntryDto) SetTotalDebitAmount(v Money)`

SetTotalDebitAmount sets TotalDebitAmount field to given value.

### HasTotalDebitAmount

`func (o *JournalEntryDto) HasTotalDebitAmount() bool`

HasTotalDebitAmount returns a boolean if a field has been set.

### GetTotalCreditAmount

`func (o *JournalEntryDto) GetTotalCreditAmount() Money`

GetTotalCreditAmount returns the TotalCreditAmount field if non-nil, zero value otherwise.

### GetTotalCreditAmountOk

`func (o *JournalEntryDto) GetTotalCreditAmountOk() (*Money, bool)`

GetTotalCreditAmountOk returns a tuple with the TotalCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditAmount

`func (o *JournalEntryDto) SetTotalCreditAmount(v Money)`

SetTotalCreditAmount sets TotalCreditAmount field to given value.

### HasTotalCreditAmount

`func (o *JournalEntryDto) HasTotalCreditAmount() bool`

HasTotalCreditAmount returns a boolean if a field has been set.

### GetDebitInUsdAmount

`func (o *JournalEntryDto) GetDebitInUsdAmount() Money`

GetDebitInUsdAmount returns the DebitInUsdAmount field if non-nil, zero value otherwise.

### GetDebitInUsdAmountOk

`func (o *JournalEntryDto) GetDebitInUsdAmountOk() (*Money, bool)`

GetDebitInUsdAmountOk returns a tuple with the DebitInUsdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitInUsdAmount

`func (o *JournalEntryDto) SetDebitInUsdAmount(v Money)`

SetDebitInUsdAmount sets DebitInUsdAmount field to given value.

### HasDebitInUsdAmount

`func (o *JournalEntryDto) HasDebitInUsdAmount() bool`

HasDebitInUsdAmount returns a boolean if a field has been set.

### GetCreditInUsdAmount

`func (o *JournalEntryDto) GetCreditInUsdAmount() Money`

GetCreditInUsdAmount returns the CreditInUsdAmount field if non-nil, zero value otherwise.

### GetCreditInUsdAmountOk

`func (o *JournalEntryDto) GetCreditInUsdAmountOk() (*Money, bool)`

GetCreditInUsdAmountOk returns a tuple with the CreditInUsdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditInUsdAmount

`func (o *JournalEntryDto) SetCreditInUsdAmount(v Money)`

SetCreditInUsdAmount sets CreditInUsdAmount field to given value.

### HasCreditInUsdAmount

`func (o *JournalEntryDto) HasCreditInUsdAmount() bool`

HasCreditInUsdAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


