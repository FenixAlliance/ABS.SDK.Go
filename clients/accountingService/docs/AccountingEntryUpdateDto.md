# AccountingEntryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Date** | Pointer to **NullableTime** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**DebitAccountId** | Pointer to **NullableString** |  | [optional] 
**CreditAccountId** | Pointer to **NullableString** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**AccountingEntryType** | Pointer to **string** |  | [optional] 

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

### GetTenantId

`func (o *AccountingEntryUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountingEntryUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountingEntryUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccountingEntryUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AccountingEntryUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AccountingEntryUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AccountingEntryUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AccountingEntryUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AccountingEntryUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AccountingEntryUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AccountingEntryUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AccountingEntryUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
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
### GetAmount

`func (o *AccountingEntryUpdateDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountingEntryUpdateDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountingEntryUpdateDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountingEntryUpdateDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *AccountingEntryUpdateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountingEntryUpdateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountingEntryUpdateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *AccountingEntryUpdateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *AccountingEntryUpdateDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AccountingEntryUpdateDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetCurrencyId

`func (o *AccountingEntryUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountingEntryUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountingEntryUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AccountingEntryUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AccountingEntryUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AccountingEntryUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDebitAccountId

`func (o *AccountingEntryUpdateDto) GetDebitAccountId() string`

GetDebitAccountId returns the DebitAccountId field if non-nil, zero value otherwise.

### GetDebitAccountIdOk

`func (o *AccountingEntryUpdateDto) GetDebitAccountIdOk() (*string, bool)`

GetDebitAccountIdOk returns a tuple with the DebitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountId

`func (o *AccountingEntryUpdateDto) SetDebitAccountId(v string)`

SetDebitAccountId sets DebitAccountId field to given value.

### HasDebitAccountId

`func (o *AccountingEntryUpdateDto) HasDebitAccountId() bool`

HasDebitAccountId returns a boolean if a field has been set.

### SetDebitAccountIdNil

`func (o *AccountingEntryUpdateDto) SetDebitAccountIdNil(b bool)`

 SetDebitAccountIdNil sets the value for DebitAccountId to be an explicit nil

### UnsetDebitAccountId
`func (o *AccountingEntryUpdateDto) UnsetDebitAccountId()`

UnsetDebitAccountId ensures that no value is present for DebitAccountId, not even an explicit nil
### GetCreditAccountId

`func (o *AccountingEntryUpdateDto) GetCreditAccountId() string`

GetCreditAccountId returns the CreditAccountId field if non-nil, zero value otherwise.

### GetCreditAccountIdOk

`func (o *AccountingEntryUpdateDto) GetCreditAccountIdOk() (*string, bool)`

GetCreditAccountIdOk returns a tuple with the CreditAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountId

`func (o *AccountingEntryUpdateDto) SetCreditAccountId(v string)`

SetCreditAccountId sets CreditAccountId field to given value.

### HasCreditAccountId

`func (o *AccountingEntryUpdateDto) HasCreditAccountId() bool`

HasCreditAccountId returns a boolean if a field has been set.

### SetCreditAccountIdNil

`func (o *AccountingEntryUpdateDto) SetCreditAccountIdNil(b bool)`

 SetCreditAccountIdNil sets the value for CreditAccountId to be an explicit nil

### UnsetCreditAccountId
`func (o *AccountingEntryUpdateDto) UnsetCreditAccountId()`

UnsetCreditAccountId ensures that no value is present for CreditAccountId, not even an explicit nil
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
### GetAccountingEntryType

`func (o *AccountingEntryUpdateDto) GetAccountingEntryType() string`

GetAccountingEntryType returns the AccountingEntryType field if non-nil, zero value otherwise.

### GetAccountingEntryTypeOk

`func (o *AccountingEntryUpdateDto) GetAccountingEntryTypeOk() (*string, bool)`

GetAccountingEntryTypeOk returns a tuple with the AccountingEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryType

`func (o *AccountingEntryUpdateDto) SetAccountingEntryType(v string)`

SetAccountingEntryType sets AccountingEntryType field to given value.

### HasAccountingEntryType

`func (o *AccountingEntryUpdateDto) HasAccountingEntryType() bool`

HasAccountingEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


