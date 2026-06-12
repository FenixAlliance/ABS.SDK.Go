# BlockchainBlockCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Hash** | **string** |  | 
**Data** | Pointer to **NullableString** |  | [optional] 
**Nonce** | Pointer to **int64** |  | [optional] 
**PreviousHash** | Pointer to **NullableString** |  | [optional] 
**BlockchainId** | **string** |  | 
**WalletAccountId** | **string** |  | 

## Methods

### NewBlockchainBlockCreateDto

`func NewBlockchainBlockCreateDto(hash string, blockchainId string, walletAccountId string, ) *BlockchainBlockCreateDto`

NewBlockchainBlockCreateDto instantiates a new BlockchainBlockCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainBlockCreateDtoWithDefaults

`func NewBlockchainBlockCreateDtoWithDefaults() *BlockchainBlockCreateDto`

NewBlockchainBlockCreateDtoWithDefaults instantiates a new BlockchainBlockCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockchainBlockCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockchainBlockCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockchainBlockCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlockchainBlockCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlockchainBlockCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockchainBlockCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockchainBlockCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockchainBlockCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHash

`func (o *BlockchainBlockCreateDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockchainBlockCreateDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockchainBlockCreateDto) SetHash(v string)`

SetHash sets Hash field to given value.


### GetData

`func (o *BlockchainBlockCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BlockchainBlockCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BlockchainBlockCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BlockchainBlockCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BlockchainBlockCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BlockchainBlockCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNonce

`func (o *BlockchainBlockCreateDto) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockchainBlockCreateDto) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockchainBlockCreateDto) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BlockchainBlockCreateDto) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPreviousHash

`func (o *BlockchainBlockCreateDto) GetPreviousHash() string`

GetPreviousHash returns the PreviousHash field if non-nil, zero value otherwise.

### GetPreviousHashOk

`func (o *BlockchainBlockCreateDto) GetPreviousHashOk() (*string, bool)`

GetPreviousHashOk returns a tuple with the PreviousHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousHash

`func (o *BlockchainBlockCreateDto) SetPreviousHash(v string)`

SetPreviousHash sets PreviousHash field to given value.

### HasPreviousHash

`func (o *BlockchainBlockCreateDto) HasPreviousHash() bool`

HasPreviousHash returns a boolean if a field has been set.

### SetPreviousHashNil

`func (o *BlockchainBlockCreateDto) SetPreviousHashNil(b bool)`

 SetPreviousHashNil sets the value for PreviousHash to be an explicit nil

### UnsetPreviousHash
`func (o *BlockchainBlockCreateDto) UnsetPreviousHash()`

UnsetPreviousHash ensures that no value is present for PreviousHash, not even an explicit nil
### GetBlockchainId

`func (o *BlockchainBlockCreateDto) GetBlockchainId() string`

GetBlockchainId returns the BlockchainId field if non-nil, zero value otherwise.

### GetBlockchainIdOk

`func (o *BlockchainBlockCreateDto) GetBlockchainIdOk() (*string, bool)`

GetBlockchainIdOk returns a tuple with the BlockchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainId

`func (o *BlockchainBlockCreateDto) SetBlockchainId(v string)`

SetBlockchainId sets BlockchainId field to given value.


### GetWalletAccountId

`func (o *BlockchainBlockCreateDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *BlockchainBlockCreateDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *BlockchainBlockCreateDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


