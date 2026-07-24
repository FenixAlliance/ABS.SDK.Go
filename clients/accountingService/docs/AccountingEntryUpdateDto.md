# AccountingEntryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**TransactionAmount** | Pointer to **float64** |  | [optional] 
**TransactionCurrencyId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccountingEntryUpdateDto

`func NewAccountingEntryUpdateDto() *AccountingEntryUpdateDto`

NewAccountingEntryUpdateDto instantiates a new AccountingEntryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingEntryUpdateDtoWithDefaults

`func NewAccountingEntryUpdateDtoWithDefaults() *AccountingEntryUpdateDto`

NewAccountingEntryUpdateDtoWithDefaults instantiates a new AccountingEntryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJournalEntryId

`func (o *AccountingEntryUpdateDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *AccountingEntryUpdateDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *AccountingEntryUpdateDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.

### HasJournalEntryId

`func (o *AccountingEntryUpdateDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *AccountingEntryUpdateDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *AccountingEntryUpdateDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetAccountId

`func (o *AccountingEntryUpdateDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountingEntryUpdateDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountingEntryUpdateDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountingEntryUpdateDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AccountingEntryUpdateDto) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AccountingEntryUpdateDto) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetDirection

`func (o *AccountingEntryUpdateDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AccountingEntryUpdateDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AccountingEntryUpdateDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AccountingEntryUpdateDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *AccountingEntryUpdateDto) GetTransactionAmount() float64`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *AccountingEntryUpdateDto) GetTransactionAmountOk() (*float64, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *AccountingEntryUpdateDto) SetTransactionAmount(v float64)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *AccountingEntryUpdateDto) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionCurrencyId

`func (o *AccountingEntryUpdateDto) GetTransactionCurrencyId() string`

GetTransactionCurrencyId returns the TransactionCurrencyId field if non-nil, zero value otherwise.

### GetTransactionCurrencyIdOk

`func (o *AccountingEntryUpdateDto) GetTransactionCurrencyIdOk() (*string, bool)`

GetTransactionCurrencyIdOk returns a tuple with the TransactionCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCurrencyId

`func (o *AccountingEntryUpdateDto) SetTransactionCurrencyId(v string)`

SetTransactionCurrencyId sets TransactionCurrencyId field to given value.

### HasTransactionCurrencyId

`func (o *AccountingEntryUpdateDto) HasTransactionCurrencyId() bool`

HasTransactionCurrencyId returns a boolean if a field has been set.

### SetTransactionCurrencyIdNil

`func (o *AccountingEntryUpdateDto) SetTransactionCurrencyIdNil(b bool)`

 SetTransactionCurrencyIdNil sets the value for TransactionCurrencyId to be an explicit nil

### UnsetTransactionCurrencyId
`func (o *AccountingEntryUpdateDto) UnsetTransactionCurrencyId()`

UnsetTransactionCurrencyId ensures that no value is present for TransactionCurrencyId, not even an explicit nil
### GetDescription

`func (o *AccountingEntryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountingEntryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountingEntryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountingEntryUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountingEntryUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountingEntryUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


