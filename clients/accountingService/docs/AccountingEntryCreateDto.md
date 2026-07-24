# AccountingEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**JournalEntryId** | **string** |  | 
**AccountId** | **string** |  | 
**Direction** | **string** |  | 
**TransactionAmount** | Pointer to **float64** |  | [optional] 
**TransactionCurrencyId** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewAccountingEntryCreateDto

`func NewAccountingEntryCreateDto(journalEntryId string, accountId string, direction string, transactionCurrencyId string, description string, ) *AccountingEntryCreateDto`

NewAccountingEntryCreateDto instantiates a new AccountingEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingEntryCreateDtoWithDefaults

`func NewAccountingEntryCreateDtoWithDefaults() *AccountingEntryCreateDto`

NewAccountingEntryCreateDtoWithDefaults instantiates a new AccountingEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountingEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountingEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountingEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountingEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJournalEntryId

`func (o *AccountingEntryCreateDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *AccountingEntryCreateDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *AccountingEntryCreateDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.


### GetAccountId

`func (o *AccountingEntryCreateDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountingEntryCreateDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountingEntryCreateDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDirection

`func (o *AccountingEntryCreateDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AccountingEntryCreateDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AccountingEntryCreateDto) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetTransactionAmount

`func (o *AccountingEntryCreateDto) GetTransactionAmount() float64`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *AccountingEntryCreateDto) GetTransactionAmountOk() (*float64, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *AccountingEntryCreateDto) SetTransactionAmount(v float64)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *AccountingEntryCreateDto) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionCurrencyId

`func (o *AccountingEntryCreateDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *AccountingEntryCreateDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *AccountingEntryCreateDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.


### GetDescription

`func (o *AccountingEntryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountingEntryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountingEntryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


