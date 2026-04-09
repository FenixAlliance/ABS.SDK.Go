# JournalEntryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **bool** |  | [optional] 
**Opening** | Pointer to **bool** |  | [optional] 
**Description** | **string** |  | 
**Date** | **time.Time** |  | 
**Debit** | Pointer to **float64** |  | [optional] 
**Credit** | Pointer to **float64** |  | [optional] 
**JournalId** | **string** |  | 
**CurrencyId** | **string** |  | 
**InvoiceCode** | Pointer to **NullableString** |  | [optional] 
**DebitAccountId** | **string** |  | 
**CreditAccountId** | **string** |  | 
**ParentJournalEntryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJournalEntryUpdateDto

`func NewJournalEntryUpdateDto(description string, date time.Time, journalId string, currencyId string, debitAccountId string, creditAccountId string, ) *JournalEntryUpdateDto`

NewJournalEntryUpdateDto instantiates a new JournalEntryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryUpdateDtoWithDefaults

`func NewJournalEntryUpdateDtoWithDefaults() *JournalEntryUpdateDto`

NewJournalEntryUpdateDtoWithDefaults instantiates a new JournalEntryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *JournalEntryUpdateDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JournalEntryUpdateDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JournalEntryUpdateDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *JournalEntryUpdateDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOpening

`func (o *JournalEntryUpdateDto) GetOpening() bool`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *JournalEntryUpdateDto) GetOpeningOk() (*bool, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *JournalEntryUpdateDto) SetOpening(v bool)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *JournalEntryUpdateDto) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

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


### GetDate

`func (o *JournalEntryUpdateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *JournalEntryUpdateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *JournalEntryUpdateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDebit

`func (o *JournalEntryUpdateDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *JournalEntryUpdateDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *JournalEntryUpdateDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *JournalEntryUpdateDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCredit

`func (o *JournalEntryUpdateDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *JournalEntryUpdateDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *JournalEntryUpdateDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *JournalEntryUpdateDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetJournalId

`func (o *JournalEntryUpdateDto) GetJournalId() string`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *JournalEntryUpdateDto) GetJournalIdOk() (*string, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *JournalEntryUpdateDto) SetJournalId(v string)`

SetJournalId sets JournalId field to given value.


### GetCurrencyId

`func (o *JournalEntryUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JournalEntryUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JournalEntryUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetInvoiceCode

`func (o *JournalEntryUpdateDto) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *JournalEntryUpdateDto) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *JournalEntryUpdateDto) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.

### HasInvoiceCode

`func (o *JournalEntryUpdateDto) HasInvoiceCode() bool`

HasInvoiceCode returns a boolean if a field has been set.

### SetInvoiceCodeNil

`func (o *JournalEntryUpdateDto) SetInvoiceCodeNil(b bool)`

 SetInvoiceCodeNil sets the value for InvoiceCode to be an explicit nil

### UnsetInvoiceCode
`func (o *JournalEntryUpdateDto) UnsetInvoiceCode()`

UnsetInvoiceCode ensures that no value is present for InvoiceCode, not even an explicit nil
### GetDebitAccountId

`func (o *JournalEntryUpdateDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *JournalEntryUpdateDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *JournalEntryUpdateDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.


### GetCreditAccountId

`func (o *JournalEntryUpdateDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *JournalEntryUpdateDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *JournalEntryUpdateDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.


### GetParentJournalEntryId

`func (o *JournalEntryUpdateDto) GetParentJournalEntryId() string`

GetParentJournalEntryId returns the ParentJournalEntryId field if non-nil, zero value otherwise.

### GetParentJournalEntryIdOk

`func (o *JournalEntryUpdateDto) GetParentJournalEntryIdOk() (*string, bool)`

GetParentJournalEntryIdOk returns a tuple with the ParentJournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalEntryId

`func (o *JournalEntryUpdateDto) SetParentJournalEntryId(v string)`

SetParentJournalEntryId sets ParentJournalEntryId field to given value.

### HasParentJournalEntryId

`func (o *JournalEntryUpdateDto) HasParentJournalEntryId() bool`

HasParentJournalEntryId returns a boolean if a field has been set.

### SetParentJournalEntryIdNil

`func (o *JournalEntryUpdateDto) SetParentJournalEntryIdNil(b bool)`

 SetParentJournalEntryIdNil sets the value for ParentJournalEntryId to be an explicit nil

### UnsetParentJournalEntryId
`func (o *JournalEntryUpdateDto) UnsetParentJournalEntryId()`

UnsetParentJournalEntryId ensures that no value is present for ParentJournalEntryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


