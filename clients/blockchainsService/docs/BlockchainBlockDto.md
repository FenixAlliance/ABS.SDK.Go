# BlockchainBlockDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Nonce** | Pointer to **int64** |  | [optional] 
**PreviousHash** | Pointer to **NullableString** |  | [optional] 
**BlockchainId** | Pointer to **NullableString** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlockchainBlockDto

`func NewBlockchainBlockDto() *BlockchainBlockDto`

NewBlockchainBlockDto instantiates a new BlockchainBlockDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainBlockDtoWithDefaults

`func NewBlockchainBlockDtoWithDefaults() *BlockchainBlockDto`

NewBlockchainBlockDtoWithDefaults instantiates a new BlockchainBlockDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockchainBlockDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockchainBlockDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockchainBlockDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlockchainBlockDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BlockchainBlockDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BlockchainBlockDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BlockchainBlockDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockchainBlockDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockchainBlockDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockchainBlockDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BlockchainBlockDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BlockchainBlockDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetIndex

`func (o *BlockchainBlockDto) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BlockchainBlockDto) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BlockchainBlockDto) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BlockchainBlockDto) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetHash

`func (o *BlockchainBlockDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockchainBlockDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockchainBlockDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *BlockchainBlockDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *BlockchainBlockDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *BlockchainBlockDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetData

`func (o *BlockchainBlockDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BlockchainBlockDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BlockchainBlockDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BlockchainBlockDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BlockchainBlockDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BlockchainBlockDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNonce

`func (o *BlockchainBlockDto) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockchainBlockDto) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockchainBlockDto) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BlockchainBlockDto) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPreviousHash

`func (o *BlockchainBlockDto) GetPreviousHash() string`

GetPreviousHash returns the PreviousHash field if non-nil, zero value otherwise.

### GetPreviousHashOk

`func (o *BlockchainBlockDto) GetPreviousHashOk() (*string, bool)`

GetPreviousHashOk returns a tuple with the PreviousHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousHash

`func (o *BlockchainBlockDto) SetPreviousHash(v string)`

SetPreviousHash sets PreviousHash field to given value.

### HasPreviousHash

`func (o *BlockchainBlockDto) HasPreviousHash() bool`

HasPreviousHash returns a boolean if a field has been set.

### SetPreviousHashNil

`func (o *BlockchainBlockDto) SetPreviousHashNil(b bool)`

 SetPreviousHashNil sets the value for PreviousHash to be an explicit nil

### UnsetPreviousHash
`func (o *BlockchainBlockDto) UnsetPreviousHash()`

UnsetPreviousHash ensures that no value is present for PreviousHash, not even an explicit nil
### GetBlockchainId

`func (o *BlockchainBlockDto) GetBlockchainId() string`

GetBlockchainId returns the BlockchainId field if non-nil, zero value otherwise.

### GetBlockchainIdOk

`func (o *BlockchainBlockDto) GetBlockchainIdOk() (*string, bool)`

GetBlockchainIdOk returns a tuple with the BlockchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainId

`func (o *BlockchainBlockDto) SetBlockchainId(v string)`

SetBlockchainId sets BlockchainId field to given value.

### HasBlockchainId

`func (o *BlockchainBlockDto) HasBlockchainId() bool`

HasBlockchainId returns a boolean if a field has been set.

### SetBlockchainIdNil

`func (o *BlockchainBlockDto) SetBlockchainIdNil(b bool)`

 SetBlockchainIdNil sets the value for BlockchainId to be an explicit nil

### UnsetBlockchainId
`func (o *BlockchainBlockDto) UnsetBlockchainId()`

UnsetBlockchainId ensures that no value is present for BlockchainId, not even an explicit nil
### GetWalletAccountId

`func (o *BlockchainBlockDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *BlockchainBlockDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *BlockchainBlockDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *BlockchainBlockDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *BlockchainBlockDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *BlockchainBlockDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetTenantId

`func (o *BlockchainBlockDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BlockchainBlockDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BlockchainBlockDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BlockchainBlockDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BlockchainBlockDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BlockchainBlockDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


