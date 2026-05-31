# WalletWithdrawRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**RequestedWithdrawAmount** | Pointer to **float64** |  | [optional] 
**RequestedWithdrawAmountInUSD** | Pointer to **float64** |  | [optional] 
**WalletWithdrawRequestStatus** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**BusinessId** | Pointer to **NullableString** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWalletWithdrawRequestDto

`func NewWalletWithdrawRequestDto() *WalletWithdrawRequestDto`

NewWalletWithdrawRequestDto instantiates a new WalletWithdrawRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithdrawRequestDtoWithDefaults

`func NewWalletWithdrawRequestDtoWithDefaults() *WalletWithdrawRequestDto`

NewWalletWithdrawRequestDtoWithDefaults instantiates a new WalletWithdrawRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletWithdrawRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletWithdrawRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletWithdrawRequestDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletWithdrawRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WalletWithdrawRequestDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WalletWithdrawRequestDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WalletWithdrawRequestDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WalletWithdrawRequestDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WalletWithdrawRequestDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WalletWithdrawRequestDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WalletWithdrawRequestDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WalletWithdrawRequestDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequestedWithdrawAmount

`func (o *WalletWithdrawRequestDto) GetRequestedWithdrawAmount() float64`

GetRequestedWithdrawAmount returns the RequestedWithdrawAmount field if non-nil, zero value otherwise.

### GetRequestedWithdrawAmountOk

`func (o *WalletWithdrawRequestDto) GetRequestedWithdrawAmountOk() (*float64, bool)`

GetRequestedWithdrawAmountOk returns a tuple with the RequestedWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWithdrawAmount

`func (o *WalletWithdrawRequestDto) SetRequestedWithdrawAmount(v float64)`

SetRequestedWithdrawAmount sets RequestedWithdrawAmount field to given value.

### HasRequestedWithdrawAmount

`func (o *WalletWithdrawRequestDto) HasRequestedWithdrawAmount() bool`

HasRequestedWithdrawAmount returns a boolean if a field has been set.

### GetRequestedWithdrawAmountInUSD

`func (o *WalletWithdrawRequestDto) GetRequestedWithdrawAmountInUSD() float64`

GetRequestedWithdrawAmountInUSD returns the RequestedWithdrawAmountInUSD field if non-nil, zero value otherwise.

### GetRequestedWithdrawAmountInUSDOk

`func (o *WalletWithdrawRequestDto) GetRequestedWithdrawAmountInUSDOk() (*float64, bool)`

GetRequestedWithdrawAmountInUSDOk returns a tuple with the RequestedWithdrawAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWithdrawAmountInUSD

`func (o *WalletWithdrawRequestDto) SetRequestedWithdrawAmountInUSD(v float64)`

SetRequestedWithdrawAmountInUSD sets RequestedWithdrawAmountInUSD field to given value.

### HasRequestedWithdrawAmountInUSD

`func (o *WalletWithdrawRequestDto) HasRequestedWithdrawAmountInUSD() bool`

HasRequestedWithdrawAmountInUSD returns a boolean if a field has been set.

### GetWalletWithdrawRequestStatus

`func (o *WalletWithdrawRequestDto) GetWalletWithdrawRequestStatus() string`

GetWalletWithdrawRequestStatus returns the WalletWithdrawRequestStatus field if non-nil, zero value otherwise.

### GetWalletWithdrawRequestStatusOk

`func (o *WalletWithdrawRequestDto) GetWalletWithdrawRequestStatusOk() (*string, bool)`

GetWalletWithdrawRequestStatusOk returns a tuple with the WalletWithdrawRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletWithdrawRequestStatus

`func (o *WalletWithdrawRequestDto) SetWalletWithdrawRequestStatus(v string)`

SetWalletWithdrawRequestStatus sets WalletWithdrawRequestStatus field to given value.

### HasWalletWithdrawRequestStatus

`func (o *WalletWithdrawRequestDto) HasWalletWithdrawRequestStatus() bool`

HasWalletWithdrawRequestStatus returns a boolean if a field has been set.

### GetCurrencyId

`func (o *WalletWithdrawRequestDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WalletWithdrawRequestDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WalletWithdrawRequestDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WalletWithdrawRequestDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WalletWithdrawRequestDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WalletWithdrawRequestDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetBusinessId

`func (o *WalletWithdrawRequestDto) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *WalletWithdrawRequestDto) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *WalletWithdrawRequestDto) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *WalletWithdrawRequestDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### SetBusinessIdNil

`func (o *WalletWithdrawRequestDto) SetBusinessIdNil(b bool)`

 SetBusinessIdNil sets the value for BusinessId to be an explicit nil

### UnsetBusinessId
`func (o *WalletWithdrawRequestDto) UnsetBusinessId()`

UnsetBusinessId ensures that no value is present for BusinessId, not even an explicit nil
### GetWalletAccountId

`func (o *WalletWithdrawRequestDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *WalletWithdrawRequestDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *WalletWithdrawRequestDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *WalletWithdrawRequestDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *WalletWithdrawRequestDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *WalletWithdrawRequestDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetBankAccountId

`func (o *WalletWithdrawRequestDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *WalletWithdrawRequestDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *WalletWithdrawRequestDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *WalletWithdrawRequestDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *WalletWithdrawRequestDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *WalletWithdrawRequestDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


