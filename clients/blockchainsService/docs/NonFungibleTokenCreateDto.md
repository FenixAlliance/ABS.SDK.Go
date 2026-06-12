# NonFungibleTokenCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**BlockchainId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNonFungibleTokenCreateDto

`func NewNonFungibleTokenCreateDto(title string, description string, ) *NonFungibleTokenCreateDto`

NewNonFungibleTokenCreateDto instantiates a new NonFungibleTokenCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenCreateDtoWithDefaults

`func NewNonFungibleTokenCreateDtoWithDefaults() *NonFungibleTokenCreateDto`

NewNonFungibleTokenCreateDtoWithDefaults instantiates a new NonFungibleTokenCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NonFungibleTokenCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NonFungibleTokenCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NonFungibleTokenCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NonFungibleTokenCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *NonFungibleTokenCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NonFungibleTokenCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NonFungibleTokenCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NonFungibleTokenCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *NonFungibleTokenCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NonFungibleTokenCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NonFungibleTokenCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NonFungibleTokenCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NonFungibleTokenCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NonFungibleTokenCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSummary

`func (o *NonFungibleTokenCreateDto) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NonFungibleTokenCreateDto) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NonFungibleTokenCreateDto) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NonFungibleTokenCreateDto) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *NonFungibleTokenCreateDto) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *NonFungibleTokenCreateDto) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetPrice

`func (o *NonFungibleTokenCreateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NonFungibleTokenCreateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NonFungibleTokenCreateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NonFungibleTokenCreateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *NonFungibleTokenCreateDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *NonFungibleTokenCreateDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCurrencyId

`func (o *NonFungibleTokenCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *NonFungibleTokenCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *NonFungibleTokenCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *NonFungibleTokenCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *NonFungibleTokenCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *NonFungibleTokenCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetBlockchainId

`func (o *NonFungibleTokenCreateDto) GetBlockchainId() string`

GetBlockchainId returns the BlockchainId field if non-nil, zero value otherwise.

### GetBlockchainIdOk

`func (o *NonFungibleTokenCreateDto) GetBlockchainIdOk() (*string, bool)`

GetBlockchainIdOk returns a tuple with the BlockchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainId

`func (o *NonFungibleTokenCreateDto) SetBlockchainId(v string)`

SetBlockchainId sets BlockchainId field to given value.

### HasBlockchainId

`func (o *NonFungibleTokenCreateDto) HasBlockchainId() bool`

HasBlockchainId returns a boolean if a field has been set.

### SetBlockchainIdNil

`func (o *NonFungibleTokenCreateDto) SetBlockchainIdNil(b bool)`

 SetBlockchainIdNil sets the value for BlockchainId to be an explicit nil

### UnsetBlockchainId
`func (o *NonFungibleTokenCreateDto) UnsetBlockchainId()`

UnsetBlockchainId ensures that no value is present for BlockchainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


