# WalletWithdrawRequestCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**RequestedWithdrawAmount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWalletWithdrawRequestCreateDto

`func NewWalletWithdrawRequestCreateDto() *WalletWithdrawRequestCreateDto`

NewWalletWithdrawRequestCreateDto instantiates a new WalletWithdrawRequestCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithdrawRequestCreateDtoWithDefaults

`func NewWalletWithdrawRequestCreateDtoWithDefaults() *WalletWithdrawRequestCreateDto`

NewWalletWithdrawRequestCreateDtoWithDefaults instantiates a new WalletWithdrawRequestCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletWithdrawRequestCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletWithdrawRequestCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletWithdrawRequestCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletWithdrawRequestCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WalletWithdrawRequestCreateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WalletWithdrawRequestCreateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WalletWithdrawRequestCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WalletWithdrawRequestCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WalletWithdrawRequestCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WalletWithdrawRequestCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WalletWithdrawRequestCreateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WalletWithdrawRequestCreateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequestedWithdrawAmount

`func (o *WalletWithdrawRequestCreateDto) GetRequestedWithdrawAmount() float64`

GetRequestedWithdrawAmount returns the RequestedWithdrawAmount field if non-nil, zero value otherwise.

### GetRequestedWithdrawAmountOk

`func (o *WalletWithdrawRequestCreateDto) GetRequestedWithdrawAmountOk() (*float64, bool)`

GetRequestedWithdrawAmountOk returns a tuple with the RequestedWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWithdrawAmount

`func (o *WalletWithdrawRequestCreateDto) SetRequestedWithdrawAmount(v float64)`

SetRequestedWithdrawAmount sets RequestedWithdrawAmount field to given value.

### HasRequestedWithdrawAmount

`func (o *WalletWithdrawRequestCreateDto) HasRequestedWithdrawAmount() bool`

HasRequestedWithdrawAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *WalletWithdrawRequestCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WalletWithdrawRequestCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WalletWithdrawRequestCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WalletWithdrawRequestCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WalletWithdrawRequestCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WalletWithdrawRequestCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetBankAccountId

`func (o *WalletWithdrawRequestCreateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *WalletWithdrawRequestCreateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *WalletWithdrawRequestCreateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *WalletWithdrawRequestCreateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *WalletWithdrawRequestCreateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *WalletWithdrawRequestCreateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


