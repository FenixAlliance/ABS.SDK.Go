# WalletDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Balance** | Pointer to **float64** |  | [optional] 
**CryptoBalance** | Pointer to **float64** |  | [optional] 
**TestMode** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**BalanceInUsd** | Pointer to **float64** |  | [optional] 
**MainNetEtherBalance** | Pointer to **float64** |  | [optional] 
**EthereumAddress** | Pointer to **NullableString** |  | [optional] 
**EthereumPublicKey** | Pointer to **NullableString** |  | [optional] 
**EthereumPrivateKey** | Pointer to **NullableString** |  | [optional] 
**RollingReservePercent** | Pointer to **float64** |  | [optional] 

## Methods

### NewWalletDto

`func NewWalletDto() *WalletDto`

NewWalletDto instantiates a new WalletDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletDtoWithDefaults

`func NewWalletDtoWithDefaults() *WalletDto`

NewWalletDtoWithDefaults instantiates a new WalletDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WalletDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WalletDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WalletDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WalletDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WalletDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WalletDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WalletDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBalance

`func (o *WalletDto) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletDto) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletDto) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *WalletDto) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCryptoBalance

`func (o *WalletDto) GetCryptoBalance() float64`

GetCryptoBalance returns the CryptoBalance field if non-nil, zero value otherwise.

### GetCryptoBalanceOk

`func (o *WalletDto) GetCryptoBalanceOk() (*float64, bool)`

GetCryptoBalanceOk returns a tuple with the CryptoBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoBalance

`func (o *WalletDto) SetCryptoBalance(v float64)`

SetCryptoBalance sets CryptoBalance field to given value.

### HasCryptoBalance

`func (o *WalletDto) HasCryptoBalance() bool`

HasCryptoBalance returns a boolean if a field has been set.

### GetTestMode

`func (o *WalletDto) GetTestMode() bool`

GetTestMode returns the TestMode field if non-nil, zero value otherwise.

### GetTestModeOk

`func (o *WalletDto) GetTestModeOk() (*bool, bool)`

GetTestModeOk returns a tuple with the TestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMode

`func (o *WalletDto) SetTestMode(v bool)`

SetTestMode sets TestMode field to given value.

### HasTestMode

`func (o *WalletDto) HasTestMode() bool`

HasTestMode returns a boolean if a field has been set.

### GetVerified

`func (o *WalletDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *WalletDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *WalletDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *WalletDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetType

`func (o *WalletDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WalletDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WalletDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WalletDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WalletDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WalletDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCurrencyId

`func (o *WalletDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WalletDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WalletDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WalletDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WalletDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WalletDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetForexRate

`func (o *WalletDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *WalletDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *WalletDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *WalletDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetBalanceInUsd

`func (o *WalletDto) GetBalanceInUsd() float64`

GetBalanceInUsd returns the BalanceInUsd field if non-nil, zero value otherwise.

### GetBalanceInUsdOk

`func (o *WalletDto) GetBalanceInUsdOk() (*float64, bool)`

GetBalanceInUsdOk returns a tuple with the BalanceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceInUsd

`func (o *WalletDto) SetBalanceInUsd(v float64)`

SetBalanceInUsd sets BalanceInUsd field to given value.

### HasBalanceInUsd

`func (o *WalletDto) HasBalanceInUsd() bool`

HasBalanceInUsd returns a boolean if a field has been set.

### GetMainNetEtherBalance

`func (o *WalletDto) GetMainNetEtherBalance() float64`

GetMainNetEtherBalance returns the MainNetEtherBalance field if non-nil, zero value otherwise.

### GetMainNetEtherBalanceOk

`func (o *WalletDto) GetMainNetEtherBalanceOk() (*float64, bool)`

GetMainNetEtherBalanceOk returns a tuple with the MainNetEtherBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainNetEtherBalance

`func (o *WalletDto) SetMainNetEtherBalance(v float64)`

SetMainNetEtherBalance sets MainNetEtherBalance field to given value.

### HasMainNetEtherBalance

`func (o *WalletDto) HasMainNetEtherBalance() bool`

HasMainNetEtherBalance returns a boolean if a field has been set.

### GetEthereumAddress

`func (o *WalletDto) GetEthereumAddress() string`

GetEthereumAddress returns the EthereumAddress field if non-nil, zero value otherwise.

### GetEthereumAddressOk

`func (o *WalletDto) GetEthereumAddressOk() (*string, bool)`

GetEthereumAddressOk returns a tuple with the EthereumAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthereumAddress

`func (o *WalletDto) SetEthereumAddress(v string)`

SetEthereumAddress sets EthereumAddress field to given value.

### HasEthereumAddress

`func (o *WalletDto) HasEthereumAddress() bool`

HasEthereumAddress returns a boolean if a field has been set.

### SetEthereumAddressNil

`func (o *WalletDto) SetEthereumAddressNil(b bool)`

 SetEthereumAddressNil sets the value for EthereumAddress to be an explicit nil

### UnsetEthereumAddress
`func (o *WalletDto) UnsetEthereumAddress()`

UnsetEthereumAddress ensures that no value is present for EthereumAddress, not even an explicit nil
### GetEthereumPublicKey

`func (o *WalletDto) GetEthereumPublicKey() string`

GetEthereumPublicKey returns the EthereumPublicKey field if non-nil, zero value otherwise.

### GetEthereumPublicKeyOk

`func (o *WalletDto) GetEthereumPublicKeyOk() (*string, bool)`

GetEthereumPublicKeyOk returns a tuple with the EthereumPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthereumPublicKey

`func (o *WalletDto) SetEthereumPublicKey(v string)`

SetEthereumPublicKey sets EthereumPublicKey field to given value.

### HasEthereumPublicKey

`func (o *WalletDto) HasEthereumPublicKey() bool`

HasEthereumPublicKey returns a boolean if a field has been set.

### SetEthereumPublicKeyNil

`func (o *WalletDto) SetEthereumPublicKeyNil(b bool)`

 SetEthereumPublicKeyNil sets the value for EthereumPublicKey to be an explicit nil

### UnsetEthereumPublicKey
`func (o *WalletDto) UnsetEthereumPublicKey()`

UnsetEthereumPublicKey ensures that no value is present for EthereumPublicKey, not even an explicit nil
### GetEthereumPrivateKey

`func (o *WalletDto) GetEthereumPrivateKey() string`

GetEthereumPrivateKey returns the EthereumPrivateKey field if non-nil, zero value otherwise.

### GetEthereumPrivateKeyOk

`func (o *WalletDto) GetEthereumPrivateKeyOk() (*string, bool)`

GetEthereumPrivateKeyOk returns a tuple with the EthereumPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthereumPrivateKey

`func (o *WalletDto) SetEthereumPrivateKey(v string)`

SetEthereumPrivateKey sets EthereumPrivateKey field to given value.

### HasEthereumPrivateKey

`func (o *WalletDto) HasEthereumPrivateKey() bool`

HasEthereumPrivateKey returns a boolean if a field has been set.

### SetEthereumPrivateKeyNil

`func (o *WalletDto) SetEthereumPrivateKeyNil(b bool)`

 SetEthereumPrivateKeyNil sets the value for EthereumPrivateKey to be an explicit nil

### UnsetEthereumPrivateKey
`func (o *WalletDto) UnsetEthereumPrivateKey()`

UnsetEthereumPrivateKey ensures that no value is present for EthereumPrivateKey, not even an explicit nil
### GetRollingReservePercent

`func (o *WalletDto) GetRollingReservePercent() float64`

GetRollingReservePercent returns the RollingReservePercent field if non-nil, zero value otherwise.

### GetRollingReservePercentOk

`func (o *WalletDto) GetRollingReservePercentOk() (*float64, bool)`

GetRollingReservePercentOk returns a tuple with the RollingReservePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingReservePercent

`func (o *WalletDto) SetRollingReservePercent(v float64)`

SetRollingReservePercent sets RollingReservePercent field to given value.

### HasRollingReservePercent

`func (o *WalletDto) HasRollingReservePercent() bool`

HasRollingReservePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


