# AccountingEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 
**Date** | Pointer to **NullableTime** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | **string** |  | 
**DebitAccountId** | Pointer to **NullableString** |  | [optional] 
**CreditAccountId** | Pointer to **NullableString** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**AccountingEntryType** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountingEntryCreateDto

`func NewAccountingEntryCreateDto(description string, currencyId string, ) *AccountingEntryCreateDto`

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


### GetDate

`func (o *AccountingEntryCreateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountingEntryCreateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountingEntryCreateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *AccountingEntryCreateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *AccountingEntryCreateDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AccountingEntryCreateDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetAmount

`func (o *AccountingEntryCreateDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountingEntryCreateDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountingEntryCreateDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountingEntryCreateDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AccountingEntryCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountingEntryCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountingEntryCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetDebitAccountId

`func (o *AccountingEntryCreateDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *AccountingEntryCreateDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *AccountingEntryCreateDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.

### HasDebitAccountId

`func (o *AccountingEntryCreateDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *AccountingEntryCreateDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *AccountingEntryCreateDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
### GetCreditAccountId

`func (o *AccountingEntryCreateDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *AccountingEntryCreateDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *AccountingEntryCreateDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *AccountingEntryCreateDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *AccountingEntryCreateDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *AccountingEntryCreateDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
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

### HasJournalEntryId

`func (o *AccountingEntryCreateDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *AccountingEntryCreateDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *AccountingEntryCreateDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetAccountingEntryType

`func (o *AccountingEntryCreateDto) GetAccountingEntryType() string`

GetAccountingEntryType returns the AccountingEntryType field if non-nil, zero value otherwise.

### GetAccountingEntryTypeOk

`func (o *AccountingEntryCreateDto) GetAccountingEntryTypeOk() (*string, bool)`

GetAccountingEntryTypeOk returns a tuple with the AccountingEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryType

`func (o *AccountingEntryCreateDto) SetAccountingEntryType(v string)`

SetAccountingEntryType sets AccountingEntryType field to given value.

### HasAccountingEntryType

`func (o *AccountingEntryCreateDto) HasAccountingEntryType() bool`

HasAccountingEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


