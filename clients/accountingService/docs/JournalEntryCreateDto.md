# JournalEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**JournalId** | **string** |  | 
**FiscalPeriodId** | **string** |  | 
**TransactionCurrencyId** | **string** |  | 
**Description** | **string** |  | 
**SourceDocumentType** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentId** | Pointer to **NullableString** |  | [optional] 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**IsOpeningBalance** | Pointer to **bool** |  | [optional] 
**AccountingEntries** | Pointer to [**[]AccountingEntryCreateDto**](AccountingEntryCreateDto.md) |  | [optional] 

## Methods

### NewJournalEntryCreateDto

`func NewJournalEntryCreateDto(journalId string, fiscalPeriodId string, transactionCurrencyId string, description string, ) *JournalEntryCreateDto`

NewJournalEntryCreateDto instantiates a new JournalEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryCreateDtoWithDefaults

`func NewJournalEntryCreateDtoWithDefaults() *JournalEntryCreateDto`

NewJournalEntryCreateDtoWithDefaults instantiates a new JournalEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JournalEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JournalEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JournalEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JournalEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JournalEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJournalId

`func (o *JournalEntryCreateDto) GetJournalId() string`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *JournalEntryCreateDto) GetJournalIdOk() (*string, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *JournalEntryCreateDto) SetJournalId(v string)`

SetJournalId sets JournalId field to given value.


### GetFiscalPeriodId

`func (o *JournalEntryCreateDto) GetFiscalPeriodId() string`

GetFiscalPeriodId returns the FiscalPeriodId field if non-nil, zero value otherwise.

### GetFiscalPeriodIdOk

`func (o *JournalEntryCreateDto) GetFiscalPeriodIdOk() (*string, bool)`

GetFiscalPeriodIdOk returns a tuple with the FiscalPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalPeriodId

`func (o *JournalEntryCreateDto) SetFiscalPeriodId(v string)`

SetFiscalPeriodId sets FiscalPeriodId field to given value.


### GetTransactionCurrencyId

`func (o *JournalEntryCreateDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *JournalEntryCreateDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *JournalEntryCreateDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.


### GetDescription

`func (o *JournalEntryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalEntryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalEntryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSourceDocumentType

`func (o *JournalEntryCreateDto) GetSourceDocumentType() string`

GetSourceDocumentType returns the SourceDocumentType field if non-nil, zero value otherwise.

### GetSourceDocumentTypeOk

`func (o *JournalEntryCreateDto) GetSourceDocumentTypeOk() (*string, bool)`

GetSourceDocumentTypeOk returns a tuple with the SourceDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentType

`func (o *JournalEntryCreateDto) SetSourceDocumentType(v string)`

SetSourceDocumentType sets SourceDocumentType field to given value.

### HasSourceDocumentType

`func (o *JournalEntryCreateDto) HasSourceDocumentType() bool`

HasSourceDocumentType returns a boolean if a field has been set.

### SetSourceDocumentTypeNil

`func (o *JournalEntryCreateDto) SetSourceDocumentTypeNil(b bool)`

 SetSourceDocumentTypeNil sets the value for SourceDocumentType to be an explicit nil

### UnsetSourceDocumentType
`func (o *JournalEntryCreateDto) UnsetSourceDocumentType()`

UnsetSourceDocumentType ensures that no value is present for SourceDocumentType, not even an explicit nil
### GetSourceDocumentId

`func (o *JournalEntryCreateDto) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *JournalEntryCreateDto) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *JournalEntryCreateDto) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *JournalEntryCreateDto) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### SetSourceDocumentIdNil

`func (o *JournalEntryCreateDto) SetSourceDocumentIdNil(b bool)`

 SetSourceDocumentIdNil sets the value for SourceDocumentId to be an explicit nil

### UnsetSourceDocumentId
`func (o *JournalEntryCreateDto) UnsetSourceDocumentId()`

UnsetSourceDocumentId ensures that no value is present for SourceDocumentId, not even an explicit nil
### GetIdempotencyKey

`func (o *JournalEntryCreateDto) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *JournalEntryCreateDto) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *JournalEntryCreateDto) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *JournalEntryCreateDto) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *JournalEntryCreateDto) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *JournalEntryCreateDto) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetIsOpeningBalance

`func (o *JournalEntryCreateDto) GetIsOpeningBalance() bool`

GetIsOpeningBalance returns the IsOpeningBalance field if non-nil, zero value otherwise.

### GetIsOpeningBalanceOk

`func (o *JournalEntryCreateDto) GetIsOpeningBalanceOk() (*bool, bool)`

GetIsOpeningBalanceOk returns a tuple with the IsOpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpeningBalance

`func (o *JournalEntryCreateDto) SetIsOpeningBalance(v bool)`

SetIsOpeningBalance sets IsOpeningBalance field to given value.

### HasIsOpeningBalance

`func (o *JournalEntryCreateDto) HasIsOpeningBalance() bool`

HasIsOpeningBalance returns a boolean if a field has been set.

### GetAccountingEntries

`func (o *JournalEntryCreateDto) GetAccountingEntries() []AccountingEntryCreateDto`

GetAccountingEntries returns the AccountingEntries field if non-nil, zero value otherwise.

### GetAccountingEntriesOk

`func (o *JournalEntryCreateDto) GetAccountingEntriesOk() (*[]AccountingEntryCreateDto, bool)`

GetAccountingEntriesOk returns a tuple with the AccountingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntries

`func (o *JournalEntryCreateDto) SetAccountingEntries(v []AccountingEntryCreateDto)`

SetAccountingEntries sets AccountingEntries field to given value.

### HasAccountingEntries

`func (o *JournalEntryCreateDto) HasAccountingEntries() bool`

HasAccountingEntries returns a boolean if a field has been set.

### SetAccountingEntriesNil

`func (o *JournalEntryCreateDto) SetAccountingEntriesNil(b bool)`

 SetAccountingEntriesNil sets the value for AccountingEntries to be an explicit nil

### UnsetAccountingEntries
`func (o *JournalEntryCreateDto) UnsetAccountingEntries()`

UnsetAccountingEntries ensures that no value is present for AccountingEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


