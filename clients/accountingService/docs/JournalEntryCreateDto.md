# JournalEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Group** | Pointer to **bool** |  | [optional] 
**Opening** | Pointer to **bool** |  | [optional] 
**Description** | **string** |  | 
**Date** | **time.Time** |  | 
**Debit** | Pointer to **float64** |  | [optional] 
**Credit** | Pointer to **float64** |  | [optional] 
**JournalId** | **string** |  | 
**CurrencyId** | **string** |  | 
**DebitAccountId** | **string** |  | 
**CreditAccountId** | **string** |  | 
**ParentJournalEntryId** | Pointer to **NullableString** |  | [optional] 
**InvoiceCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJournalEntryCreateDto

`func NewJournalEntryCreateDto(description string, date time.Time, journalId string, currencyId string, debitAccountId string, creditAccountId string, ) *JournalEntryCreateDto`

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

### GetGroup

`func (o *JournalEntryCreateDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JournalEntryCreateDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JournalEntryCreateDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *JournalEntryCreateDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOpening

`func (o *JournalEntryCreateDto) GetOpening() bool`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *JournalEntryCreateDto) GetOpeningOk() (*bool, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *JournalEntryCreateDto) SetOpening(v bool)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *JournalEntryCreateDto) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

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


### GetDate

`func (o *JournalEntryCreateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *JournalEntryCreateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *JournalEntryCreateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDebit

`func (o *JournalEntryCreateDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *JournalEntryCreateDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *JournalEntryCreateDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *JournalEntryCreateDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCredit

`func (o *JournalEntryCreateDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *JournalEntryCreateDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *JournalEntryCreateDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *JournalEntryCreateDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

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


### GetCurrencyId

`func (o *JournalEntryCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JournalEntryCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JournalEntryCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetDebitAccountId

`func (o *JournalEntryCreateDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *JournalEntryCreateDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *JournalEntryCreateDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.


### GetCreditAccountId

`func (o *JournalEntryCreateDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *JournalEntryCreateDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *JournalEntryCreateDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.


### GetParentJournalEntryId

`func (o *JournalEntryCreateDto) GetParentJournalEntryId() string`

GetParentJournalEntryId returns the ParentJournalEntryId field if non-nil, zero value otherwise.

### GetParentJournalEntryIdOk

`func (o *JournalEntryCreateDto) GetParentJournalEntryIdOk() (*string, bool)`

GetParentJournalEntryIdOk returns a tuple with the ParentJournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalEntryId

`func (o *JournalEntryCreateDto) SetParentJournalEntryId(v string)`

SetParentJournalEntryId sets ParentJournalEntryId field to given value.

### HasParentJournalEntryId

`func (o *JournalEntryCreateDto) HasParentJournalEntryId() bool`

HasParentJournalEntryId returns a boolean if a field has been set.

### SetParentJournalEntryIdNil

`func (o *JournalEntryCreateDto) SetParentJournalEntryIdNil(b bool)`

 SetParentJournalEntryIdNil sets the value for ParentJournalEntryId to be an explicit nil

### UnsetParentJournalEntryId
`func (o *JournalEntryCreateDto) UnsetParentJournalEntryId()`

UnsetParentJournalEntryId ensures that no value is present for ParentJournalEntryId, not even an explicit nil
### GetInvoiceCode

`func (o *JournalEntryCreateDto) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *JournalEntryCreateDto) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *JournalEntryCreateDto) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.

### HasInvoiceCode

`func (o *JournalEntryCreateDto) HasInvoiceCode() bool`

HasInvoiceCode returns a boolean if a field has been set.

### SetInvoiceCodeNil

`func (o *JournalEntryCreateDto) SetInvoiceCodeNil(b bool)`

 SetInvoiceCodeNil sets the value for InvoiceCode to be an explicit nil

### UnsetInvoiceCode
`func (o *JournalEntryCreateDto) UnsetInvoiceCode()`

UnsetInvoiceCode ensures that no value is present for InvoiceCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


