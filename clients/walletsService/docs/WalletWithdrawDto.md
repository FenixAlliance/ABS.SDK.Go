# WalletWithdrawDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**WithdrawStatus** | Pointer to **string** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**WalletWithdrawRequestId** | Pointer to **NullableString** |  | [optional] 
**BalanceBeforeWithdraw** | Pointer to **float64** |  | [optional] 
**BalanceAfterWithdraw** | Pointer to **float64** |  | [optional] 
**WithdrawedAmount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWalletWithdrawDto

`func NewWalletWithdrawDto() *WalletWithdrawDto`

NewWalletWithdrawDto instantiates a new WalletWithdrawDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithdrawDtoWithDefaults

`func NewWalletWithdrawDtoWithDefaults() *WalletWithdrawDto`

NewWalletWithdrawDtoWithDefaults instantiates a new WalletWithdrawDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletWithdrawDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletWithdrawDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletWithdrawDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletWithdrawDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WalletWithdrawDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WalletWithdrawDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WalletWithdrawDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WalletWithdrawDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WalletWithdrawDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WalletWithdrawDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WalletWithdrawDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WalletWithdrawDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetWithdrawStatus

`func (o *WalletWithdrawDto) GetWithdrawStatus() string`

GetWithdrawStatus returns the WithdrawStatus field if non-nil, zero value otherwise.

### GetWithdrawStatusOk

`func (o *WalletWithdrawDto) GetWithdrawStatusOk() (*string, bool)`

GetWithdrawStatusOk returns a tuple with the WithdrawStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawStatus

`func (o *WalletWithdrawDto) SetWithdrawStatus(v string)`

SetWithdrawStatus sets WithdrawStatus field to given value.

### HasWithdrawStatus

`func (o *WalletWithdrawDto) HasWithdrawStatus() bool`

HasWithdrawStatus returns a boolean if a field has been set.

### GetWalletAccountId

`func (o *WalletWithdrawDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *WalletWithdrawDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *WalletWithdrawDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *WalletWithdrawDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *WalletWithdrawDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *WalletWithdrawDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetWalletWithdrawRequestId

`func (o *WalletWithdrawDto) GetWalletWithdrawRequestId() string`

GetWalletWithdrawRequestId returns the WalletWithdrawRequestId field if non-nil, zero value otherwise.

### GetWalletWithdrawRequestIdOk

`func (o *WalletWithdrawDto) GetWalletWithdrawRequestIdOk() (*string, bool)`

GetWalletWithdrawRequestIdOk returns a tuple with the WalletWithdrawRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletWithdrawRequestId

`func (o *WalletWithdrawDto) SetWalletWithdrawRequestId(v string)`

SetWalletWithdrawRequestId sets WalletWithdrawRequestId field to given value.

### HasWalletWithdrawRequestId

`func (o *WalletWithdrawDto) HasWalletWithdrawRequestId() bool`

HasWalletWithdrawRequestId returns a boolean if a field has been set.

### SetWalletWithdrawRequestIdNil

`func (o *WalletWithdrawDto) SetWalletWithdrawRequestIdNil(b bool)`

 SetWalletWithdrawRequestIdNil sets the value for WalletWithdrawRequestId to be an explicit nil

### UnsetWalletWithdrawRequestId
`func (o *WalletWithdrawDto) UnsetWalletWithdrawRequestId()`

UnsetWalletWithdrawRequestId ensures that no value is present for WalletWithdrawRequestId, not even an explicit nil
### GetBalanceBeforeWithdraw

`func (o *WalletWithdrawDto) GetBalanceBeforeWithdraw() float64`

GetBalanceBeforeWithdraw returns the BalanceBeforeWithdraw field if non-nil, zero value otherwise.

### GetBalanceBeforeWithdrawOk

`func (o *WalletWithdrawDto) GetBalanceBeforeWithdrawOk() (*float64, bool)`

GetBalanceBeforeWithdrawOk returns a tuple with the BalanceBeforeWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBeforeWithdraw

`func (o *WalletWithdrawDto) SetBalanceBeforeWithdraw(v float64)`

SetBalanceBeforeWithdraw sets BalanceBeforeWithdraw field to given value.

### HasBalanceBeforeWithdraw

`func (o *WalletWithdrawDto) HasBalanceBeforeWithdraw() bool`

HasBalanceBeforeWithdraw returns a boolean if a field has been set.

### GetBalanceAfterWithdraw

`func (o *WalletWithdrawDto) GetBalanceAfterWithdraw() float64`

GetBalanceAfterWithdraw returns the BalanceAfterWithdraw field if non-nil, zero value otherwise.

### GetBalanceAfterWithdrawOk

`func (o *WalletWithdrawDto) GetBalanceAfterWithdrawOk() (*float64, bool)`

GetBalanceAfterWithdrawOk returns a tuple with the BalanceAfterWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfterWithdraw

`func (o *WalletWithdrawDto) SetBalanceAfterWithdraw(v float64)`

SetBalanceAfterWithdraw sets BalanceAfterWithdraw field to given value.

### HasBalanceAfterWithdraw

`func (o *WalletWithdrawDto) HasBalanceAfterWithdraw() bool`

HasBalanceAfterWithdraw returns a boolean if a field has been set.

### GetWithdrawedAmount

`func (o *WalletWithdrawDto) GetWithdrawedAmount() float64`

GetWithdrawedAmount returns the WithdrawedAmount field if non-nil, zero value otherwise.

### GetWithdrawedAmountOk

`func (o *WalletWithdrawDto) GetWithdrawedAmountOk() (*float64, bool)`

GetWithdrawedAmountOk returns a tuple with the WithdrawedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawedAmount

`func (o *WalletWithdrawDto) SetWithdrawedAmount(v float64)`

SetWithdrawedAmount sets WithdrawedAmount field to given value.

### HasWithdrawedAmount

`func (o *WalletWithdrawDto) HasWithdrawedAmount() bool`

HasWithdrawedAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *WalletWithdrawDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WalletWithdrawDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WalletWithdrawDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WalletWithdrawDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WalletWithdrawDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WalletWithdrawDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


