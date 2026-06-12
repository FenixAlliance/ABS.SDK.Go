# NonFungibleTokenUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Published** | Pointer to **NullableBool** |  | [optional] 
**Minted** | Pointer to **NullableBool** |  | [optional] 
**MintTransactionHash** | Pointer to **NullableString** |  | [optional] 
**BlockchainBlockId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNonFungibleTokenUpdateDto

`func NewNonFungibleTokenUpdateDto() *NonFungibleTokenUpdateDto`

NewNonFungibleTokenUpdateDto instantiates a new NonFungibleTokenUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenUpdateDtoWithDefaults

`func NewNonFungibleTokenUpdateDtoWithDefaults() *NonFungibleTokenUpdateDto`

NewNonFungibleTokenUpdateDtoWithDefaults instantiates a new NonFungibleTokenUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NonFungibleTokenUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NonFungibleTokenUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NonFungibleTokenUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NonFungibleTokenUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *NonFungibleTokenUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NonFungibleTokenUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *NonFungibleTokenUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NonFungibleTokenUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NonFungibleTokenUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NonFungibleTokenUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NonFungibleTokenUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NonFungibleTokenUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSummary

`func (o *NonFungibleTokenUpdateDto) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NonFungibleTokenUpdateDto) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NonFungibleTokenUpdateDto) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NonFungibleTokenUpdateDto) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *NonFungibleTokenUpdateDto) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *NonFungibleTokenUpdateDto) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetPrice

`func (o *NonFungibleTokenUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NonFungibleTokenUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NonFungibleTokenUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NonFungibleTokenUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *NonFungibleTokenUpdateDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *NonFungibleTokenUpdateDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCurrencyId

`func (o *NonFungibleTokenUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *NonFungibleTokenUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *NonFungibleTokenUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *NonFungibleTokenUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *NonFungibleTokenUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *NonFungibleTokenUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPublished

`func (o *NonFungibleTokenUpdateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *NonFungibleTokenUpdateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *NonFungibleTokenUpdateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *NonFungibleTokenUpdateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *NonFungibleTokenUpdateDto) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *NonFungibleTokenUpdateDto) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetMinted

`func (o *NonFungibleTokenUpdateDto) GetMinted() bool`

GetMinted returns the Minted field if non-nil, zero value otherwise.

### GetMintedOk

`func (o *NonFungibleTokenUpdateDto) GetMintedOk() (*bool, bool)`

GetMintedOk returns a tuple with the Minted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinted

`func (o *NonFungibleTokenUpdateDto) SetMinted(v bool)`

SetMinted sets Minted field to given value.

### HasMinted

`func (o *NonFungibleTokenUpdateDto) HasMinted() bool`

HasMinted returns a boolean if a field has been set.

### SetMintedNil

`func (o *NonFungibleTokenUpdateDto) SetMintedNil(b bool)`

 SetMintedNil sets the value for Minted to be an explicit nil

### UnsetMinted
`func (o *NonFungibleTokenUpdateDto) UnsetMinted()`

UnsetMinted ensures that no value is present for Minted, not even an explicit nil
### GetMintTransactionHash

`func (o *NonFungibleTokenUpdateDto) GetMintTransactionHash() string`

GetMintTransactionHash returns the MintTransactionHash field if non-nil, zero value otherwise.

### GetMintTransactionHashOk

`func (o *NonFungibleTokenUpdateDto) GetMintTransactionHashOk() (*string, bool)`

GetMintTransactionHashOk returns a tuple with the MintTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTransactionHash

`func (o *NonFungibleTokenUpdateDto) SetMintTransactionHash(v string)`

SetMintTransactionHash sets MintTransactionHash field to given value.

### HasMintTransactionHash

`func (o *NonFungibleTokenUpdateDto) HasMintTransactionHash() bool`

HasMintTransactionHash returns a boolean if a field has been set.

### SetMintTransactionHashNil

`func (o *NonFungibleTokenUpdateDto) SetMintTransactionHashNil(b bool)`

 SetMintTransactionHashNil sets the value for MintTransactionHash to be an explicit nil

### UnsetMintTransactionHash
`func (o *NonFungibleTokenUpdateDto) UnsetMintTransactionHash()`

UnsetMintTransactionHash ensures that no value is present for MintTransactionHash, not even an explicit nil
### GetBlockchainBlockId

`func (o *NonFungibleTokenUpdateDto) GetBlockchainBlockId() string`

GetBlockchainBlockId returns the BlockchainBlockId field if non-nil, zero value otherwise.

### GetBlockchainBlockIdOk

`func (o *NonFungibleTokenUpdateDto) GetBlockchainBlockIdOk() (*string, bool)`

GetBlockchainBlockIdOk returns a tuple with the BlockchainBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainBlockId

`func (o *NonFungibleTokenUpdateDto) SetBlockchainBlockId(v string)`

SetBlockchainBlockId sets BlockchainBlockId field to given value.

### HasBlockchainBlockId

`func (o *NonFungibleTokenUpdateDto) HasBlockchainBlockId() bool`

HasBlockchainBlockId returns a boolean if a field has been set.

### SetBlockchainBlockIdNil

`func (o *NonFungibleTokenUpdateDto) SetBlockchainBlockIdNil(b bool)`

 SetBlockchainBlockIdNil sets the value for BlockchainBlockId to be an explicit nil

### UnsetBlockchainBlockId
`func (o *NonFungibleTokenUpdateDto) UnsetBlockchainBlockId()`

UnsetBlockchainBlockId ensures that no value is present for BlockchainBlockId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


