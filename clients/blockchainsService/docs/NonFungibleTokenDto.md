# NonFungibleTokenDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Minted** | Pointer to **bool** |  | [optional] 
**BlockchainId** | Pointer to **NullableString** |  | [optional] 
**BlockchainName** | Pointer to **NullableString** |  | [optional] 
**MintedTimestamp** | Pointer to **time.Time** |  | [optional] 
**MintTransactionHash** | Pointer to **NullableString** |  | [optional] 
**BlockchainBlockId** | Pointer to **NullableString** |  | [optional] 
**PrimaryImageUrl** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNonFungibleTokenDto

`func NewNonFungibleTokenDto() *NonFungibleTokenDto`

NewNonFungibleTokenDto instantiates a new NonFungibleTokenDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenDtoWithDefaults

`func NewNonFungibleTokenDtoWithDefaults() *NonFungibleTokenDto`

NewNonFungibleTokenDtoWithDefaults instantiates a new NonFungibleTokenDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NonFungibleTokenDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NonFungibleTokenDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NonFungibleTokenDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NonFungibleTokenDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NonFungibleTokenDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NonFungibleTokenDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *NonFungibleTokenDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NonFungibleTokenDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NonFungibleTokenDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NonFungibleTokenDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *NonFungibleTokenDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *NonFungibleTokenDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *NonFungibleTokenDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NonFungibleTokenDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NonFungibleTokenDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NonFungibleTokenDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *NonFungibleTokenDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NonFungibleTokenDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *NonFungibleTokenDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NonFungibleTokenDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NonFungibleTokenDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NonFungibleTokenDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NonFungibleTokenDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NonFungibleTokenDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSummary

`func (o *NonFungibleTokenDto) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NonFungibleTokenDto) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NonFungibleTokenDto) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NonFungibleTokenDto) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *NonFungibleTokenDto) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *NonFungibleTokenDto) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetPrice

`func (o *NonFungibleTokenDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NonFungibleTokenDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NonFungibleTokenDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NonFungibleTokenDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *NonFungibleTokenDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *NonFungibleTokenDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCurrencyId

`func (o *NonFungibleTokenDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *NonFungibleTokenDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *NonFungibleTokenDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *NonFungibleTokenDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *NonFungibleTokenDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *NonFungibleTokenDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPublished

`func (o *NonFungibleTokenDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *NonFungibleTokenDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *NonFungibleTokenDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *NonFungibleTokenDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetMinted

`func (o *NonFungibleTokenDto) GetMinted() bool`

GetMinted returns the Minted field if non-nil, zero value otherwise.

### GetMintedOk

`func (o *NonFungibleTokenDto) GetMintedOk() (*bool, bool)`

GetMintedOk returns a tuple with the Minted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinted

`func (o *NonFungibleTokenDto) SetMinted(v bool)`

SetMinted sets Minted field to given value.

### HasMinted

`func (o *NonFungibleTokenDto) HasMinted() bool`

HasMinted returns a boolean if a field has been set.

### GetBlockchainId

`func (o *NonFungibleTokenDto) GetBlockchainId() string`

GetBlockchainId returns the BlockchainId field if non-nil, zero value otherwise.

### GetBlockchainIdOk

`func (o *NonFungibleTokenDto) GetBlockchainIdOk() (*string, bool)`

GetBlockchainIdOk returns a tuple with the BlockchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainId

`func (o *NonFungibleTokenDto) SetBlockchainId(v string)`

SetBlockchainId sets BlockchainId field to given value.

### HasBlockchainId

`func (o *NonFungibleTokenDto) HasBlockchainId() bool`

HasBlockchainId returns a boolean if a field has been set.

### SetBlockchainIdNil

`func (o *NonFungibleTokenDto) SetBlockchainIdNil(b bool)`

 SetBlockchainIdNil sets the value for BlockchainId to be an explicit nil

### UnsetBlockchainId
`func (o *NonFungibleTokenDto) UnsetBlockchainId()`

UnsetBlockchainId ensures that no value is present for BlockchainId, not even an explicit nil
### GetBlockchainName

`func (o *NonFungibleTokenDto) GetBlockchainName() string`

GetBlockchainName returns the BlockchainName field if non-nil, zero value otherwise.

### GetBlockchainNameOk

`func (o *NonFungibleTokenDto) GetBlockchainNameOk() (*string, bool)`

GetBlockchainNameOk returns a tuple with the BlockchainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainName

`func (o *NonFungibleTokenDto) SetBlockchainName(v string)`

SetBlockchainName sets BlockchainName field to given value.

### HasBlockchainName

`func (o *NonFungibleTokenDto) HasBlockchainName() bool`

HasBlockchainName returns a boolean if a field has been set.

### SetBlockchainNameNil

`func (o *NonFungibleTokenDto) SetBlockchainNameNil(b bool)`

 SetBlockchainNameNil sets the value for BlockchainName to be an explicit nil

### UnsetBlockchainName
`func (o *NonFungibleTokenDto) UnsetBlockchainName()`

UnsetBlockchainName ensures that no value is present for BlockchainName, not even an explicit nil
### GetMintedTimestamp

`func (o *NonFungibleTokenDto) GetMintedTimestamp() time.Time`

GetMintedTimestamp returns the MintedTimestamp field if non-nil, zero value otherwise.

### GetMintedTimestampOk

`func (o *NonFungibleTokenDto) GetMintedTimestampOk() (*time.Time, bool)`

GetMintedTimestampOk returns a tuple with the MintedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintedTimestamp

`func (o *NonFungibleTokenDto) SetMintedTimestamp(v time.Time)`

SetMintedTimestamp sets MintedTimestamp field to given value.

### HasMintedTimestamp

`func (o *NonFungibleTokenDto) HasMintedTimestamp() bool`

HasMintedTimestamp returns a boolean if a field has been set.

### GetMintTransactionHash

`func (o *NonFungibleTokenDto) GetMintTransactionHash() string`

GetMintTransactionHash returns the MintTransactionHash field if non-nil, zero value otherwise.

### GetMintTransactionHashOk

`func (o *NonFungibleTokenDto) GetMintTransactionHashOk() (*string, bool)`

GetMintTransactionHashOk returns a tuple with the MintTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTransactionHash

`func (o *NonFungibleTokenDto) SetMintTransactionHash(v string)`

SetMintTransactionHash sets MintTransactionHash field to given value.

### HasMintTransactionHash

`func (o *NonFungibleTokenDto) HasMintTransactionHash() bool`

HasMintTransactionHash returns a boolean if a field has been set.

### SetMintTransactionHashNil

`func (o *NonFungibleTokenDto) SetMintTransactionHashNil(b bool)`

 SetMintTransactionHashNil sets the value for MintTransactionHash to be an explicit nil

### UnsetMintTransactionHash
`func (o *NonFungibleTokenDto) UnsetMintTransactionHash()`

UnsetMintTransactionHash ensures that no value is present for MintTransactionHash, not even an explicit nil
### GetBlockchainBlockId

`func (o *NonFungibleTokenDto) GetBlockchainBlockId() string`

GetBlockchainBlockId returns the BlockchainBlockId field if non-nil, zero value otherwise.

### GetBlockchainBlockIdOk

`func (o *NonFungibleTokenDto) GetBlockchainBlockIdOk() (*string, bool)`

GetBlockchainBlockIdOk returns a tuple with the BlockchainBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainBlockId

`func (o *NonFungibleTokenDto) SetBlockchainBlockId(v string)`

SetBlockchainBlockId sets BlockchainBlockId field to given value.

### HasBlockchainBlockId

`func (o *NonFungibleTokenDto) HasBlockchainBlockId() bool`

HasBlockchainBlockId returns a boolean if a field has been set.

### SetBlockchainBlockIdNil

`func (o *NonFungibleTokenDto) SetBlockchainBlockIdNil(b bool)`

 SetBlockchainBlockIdNil sets the value for BlockchainBlockId to be an explicit nil

### UnsetBlockchainBlockId
`func (o *NonFungibleTokenDto) UnsetBlockchainBlockId()`

UnsetBlockchainBlockId ensures that no value is present for BlockchainBlockId, not even an explicit nil
### GetPrimaryImageUrl

`func (o *NonFungibleTokenDto) GetPrimaryImageUrl() string`

GetPrimaryImageUrl returns the PrimaryImageUrl field if non-nil, zero value otherwise.

### GetPrimaryImageUrlOk

`func (o *NonFungibleTokenDto) GetPrimaryImageUrlOk() (*string, bool)`

GetPrimaryImageUrlOk returns a tuple with the PrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageUrl

`func (o *NonFungibleTokenDto) SetPrimaryImageUrl(v string)`

SetPrimaryImageUrl sets PrimaryImageUrl field to given value.

### HasPrimaryImageUrl

`func (o *NonFungibleTokenDto) HasPrimaryImageUrl() bool`

HasPrimaryImageUrl returns a boolean if a field has been set.

### SetPrimaryImageUrlNil

`func (o *NonFungibleTokenDto) SetPrimaryImageUrlNil(b bool)`

 SetPrimaryImageUrlNil sets the value for PrimaryImageUrl to be an explicit nil

### UnsetPrimaryImageUrl
`func (o *NonFungibleTokenDto) UnsetPrimaryImageUrl()`

UnsetPrimaryImageUrl ensures that no value is present for PrimaryImageUrl, not even an explicit nil
### GetCoverUrl

`func (o *NonFungibleTokenDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *NonFungibleTokenDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *NonFungibleTokenDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *NonFungibleTokenDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *NonFungibleTokenDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *NonFungibleTokenDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetTenantId

`func (o *NonFungibleTokenDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NonFungibleTokenDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NonFungibleTokenDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NonFungibleTokenDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *NonFungibleTokenDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *NonFungibleTokenDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *NonFungibleTokenDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *NonFungibleTokenDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *NonFungibleTokenDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *NonFungibleTokenDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *NonFungibleTokenDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *NonFungibleTokenDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


