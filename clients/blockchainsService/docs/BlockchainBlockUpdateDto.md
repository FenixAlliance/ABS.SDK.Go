# BlockchainBlockUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** |  | [optional] 
**Nonce** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewBlockchainBlockUpdateDto

`func NewBlockchainBlockUpdateDto() *BlockchainBlockUpdateDto`

NewBlockchainBlockUpdateDto instantiates a new BlockchainBlockUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainBlockUpdateDtoWithDefaults

`func NewBlockchainBlockUpdateDtoWithDefaults() *BlockchainBlockUpdateDto`

NewBlockchainBlockUpdateDtoWithDefaults instantiates a new BlockchainBlockUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BlockchainBlockUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BlockchainBlockUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BlockchainBlockUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BlockchainBlockUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BlockchainBlockUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BlockchainBlockUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNonce

`func (o *BlockchainBlockUpdateDto) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockchainBlockUpdateDto) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockchainBlockUpdateDto) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BlockchainBlockUpdateDto) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *BlockchainBlockUpdateDto) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *BlockchainBlockUpdateDto) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


