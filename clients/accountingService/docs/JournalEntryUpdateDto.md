# JournalEntryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalPeriodId** | **string** |  | 
**TransactionCurrencyId** | **string** |  | 
**Description** | **string** |  | 
**SourceDocumentType** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentId** | Pointer to **NullableString** |  | [optional] 
**IsOpeningBalance** | Pointer to **bool** |  | [optional] 

## Methods

### NewJournalEntryUpdateDto

`func NewJournalEntryUpdateDto(fiscalPeriodId string, transactionCurrencyId string, description string, ) *JournalEntryUpdateDto`

NewJournalEntryUpdateDto instantiates a new JournalEntryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryUpdateDtoWithDefaults

`func NewJournalEntryUpdateDtoWithDefaults() *JournalEntryUpdateDto`

NewJournalEntryUpdateDtoWithDefaults instantiates a new JournalEntryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalPeriodId

`func (o *JournalEntryUpdateDto) GetFiscalPeriodId() string`

GetFiscalPeriodId returns the FiscalPeriodId field if non-nil, zero value otherwise.

### GetFiscalPeriodIdOk

`func (o *JournalEntryUpdateDto) GetFiscalPeriodIdOk() (*string, bool)`

GetFiscalPeriodIdOk returns a tuple with the FiscalPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalPeriodId

`func (o *JournalEntryUpdateDto) SetFiscalPeriodId(v string)`

SetFiscalPeriodId sets FiscalPeriodId field to given value.


### GetTransactionCurrencyId

`func (o *JournalEntryUpdateDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *JournalEntryUpdateDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *JournalEntryUpdateDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.


### GetDescription

`func (o *JournalEntryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalEntryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalEntryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSourceDocumentType

`func (o *JournalEntryUpdateDto) GetSourceDocumentType() string`

GetSourceDocumentType returns the SourceDocumentType field if non-nil, zero value otherwise.

### GetSourceDocumentTypeOk

`func (o *JournalEntryUpdateDto) GetSourceDocumentTypeOk() (*string, bool)`

GetSourceDocumentTypeOk returns a tuple with the SourceDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentType

`func (o *JournalEntryUpdateDto) SetSourceDocumentType(v string)`

SetSourceDocumentType sets SourceDocumentType field to given value.

### HasSourceDocumentType

`func (o *JournalEntryUpdateDto) HasSourceDocumentType() bool`

HasSourceDocumentType returns a boolean if a field has been set.

### SetSourceDocumentTypeNil

`func (o *JournalEntryUpdateDto) SetSourceDocumentTypeNil(b bool)`

 SetSourceDocumentTypeNil sets the value for SourceDocumentType to be an explicit nil

### UnsetSourceDocumentType
`func (o *JournalEntryUpdateDto) UnsetSourceDocumentType()`

UnsetSourceDocumentType ensures that no value is present for SourceDocumentType, not even an explicit nil
### GetSourceDocumentId

`func (o *JournalEntryUpdateDto) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *JournalEntryUpdateDto) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *JournalEntryUpdateDto) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *JournalEntryUpdateDto) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### SetSourceDocumentIdNil

`func (o *JournalEntryUpdateDto) SetSourceDocumentIdNil(b bool)`

 SetSourceDocumentIdNil sets the value for SourceDocumentId to be an explicit nil

### UnsetSourceDocumentId
`func (o *JournalEntryUpdateDto) UnsetSourceDocumentId()`

UnsetSourceDocumentId ensures that no value is present for SourceDocumentId, not even an explicit nil
### GetIsOpeningBalance

`func (o *JournalEntryUpdateDto) GetIsOpeningBalance() bool`

GetIsOpeningBalance returns the IsOpeningBalance field if non-nil, zero value otherwise.

### GetIsOpeningBalanceOk

`func (o *JournalEntryUpdateDto) GetIsOpeningBalanceOk() (*bool, bool)`

GetIsOpeningBalanceOk returns a tuple with the IsOpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpeningBalance

`func (o *JournalEntryUpdateDto) SetIsOpeningBalance(v bool)`

SetIsOpeningBalance sets IsOpeningBalance field to given value.

### HasIsOpeningBalance

`func (o *JournalEntryUpdateDto) HasIsOpeningBalance() bool`

HasIsOpeningBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


