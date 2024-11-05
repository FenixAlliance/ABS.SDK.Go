# InvoiceLineUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemPriceId** | Pointer to **NullableString** |  | [optional] 
**InvoiceLineId** | Pointer to **NullableString** |  | [optional] 
**TaxAmountInUsd** | Pointer to **float64** |  | [optional] 
**TaxBaseAmountInUsd** | Pointer to **float64** |  | [optional] 

## Methods

### NewInvoiceLineUpdateDto

`func NewInvoiceLineUpdateDto() *InvoiceLineUpdateDto`

NewInvoiceLineUpdateDto instantiates a new InvoiceLineUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUpdateDtoWithDefaults

`func NewInvoiceLineUpdateDtoWithDefaults() *InvoiceLineUpdateDto`

NewInvoiceLineUpdateDtoWithDefaults instantiates a new InvoiceLineUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *InvoiceLineUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceLineUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceLineUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceLineUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetUnitId

`func (o *InvoiceLineUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *InvoiceLineUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *InvoiceLineUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *InvoiceLineUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *InvoiceLineUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *InvoiceLineUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetPercent

`func (o *InvoiceLineUpdateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *InvoiceLineUpdateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *InvoiceLineUpdateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *InvoiceLineUpdateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *InvoiceLineUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *InvoiceLineUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *InvoiceLineUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *InvoiceLineUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *InvoiceLineUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *InvoiceLineUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetCurrencyId

`func (o *InvoiceLineUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceLineUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceLineUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceLineUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceLineUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceLineUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDiscountListId

`func (o *InvoiceLineUpdateDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *InvoiceLineUpdateDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *InvoiceLineUpdateDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *InvoiceLineUpdateDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *InvoiceLineUpdateDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *InvoiceLineUpdateDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetRoundingPolicyId

`func (o *InvoiceLineUpdateDto) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *InvoiceLineUpdateDto) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *InvoiceLineUpdateDto) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *InvoiceLineUpdateDto) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *InvoiceLineUpdateDto) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *InvoiceLineUpdateDto) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
### GetQuantity

`func (o *InvoiceLineUpdateDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineUpdateDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineUpdateDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLineUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetItemId

`func (o *InvoiceLineUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceLineUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceLineUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceLineUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *InvoiceLineUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *InvoiceLineUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemPriceId

`func (o *InvoiceLineUpdateDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *InvoiceLineUpdateDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *InvoiceLineUpdateDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *InvoiceLineUpdateDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *InvoiceLineUpdateDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *InvoiceLineUpdateDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetInvoiceLineId

`func (o *InvoiceLineUpdateDto) GetInvoiceLineId() string`

GetInvoiceLineId returns the InvoiceLineId field if non-nil, zero value otherwise.

### GetInvoiceLineIdOk

`func (o *InvoiceLineUpdateDto) GetInvoiceLineIdOk() (*string, bool)`

GetInvoiceLineIdOk returns a tuple with the InvoiceLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineId

`func (o *InvoiceLineUpdateDto) SetInvoiceLineId(v string)`

SetInvoiceLineId sets InvoiceLineId field to given value.

### HasInvoiceLineId

`func (o *InvoiceLineUpdateDto) HasInvoiceLineId() bool`

HasInvoiceLineId returns a boolean if a field has been set.

### SetInvoiceLineIdNil

`func (o *InvoiceLineUpdateDto) SetInvoiceLineIdNil(b bool)`

 SetInvoiceLineIdNil sets the value for InvoiceLineId to be an explicit nil

### UnsetInvoiceLineId
`func (o *InvoiceLineUpdateDto) UnsetInvoiceLineId()`

UnsetInvoiceLineId ensures that no value is present for InvoiceLineId, not even an explicit nil
### GetTaxAmountInUsd

`func (o *InvoiceLineUpdateDto) GetTaxAmountInUsd() float64`

GetTaxAmountInUsd returns the TaxAmountInUsd field if non-nil, zero value otherwise.

### GetTaxAmountInUsdOk

`func (o *InvoiceLineUpdateDto) GetTaxAmountInUsdOk() (*float64, bool)`

GetTaxAmountInUsdOk returns a tuple with the TaxAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUsd

`func (o *InvoiceLineUpdateDto) SetTaxAmountInUsd(v float64)`

SetTaxAmountInUsd sets TaxAmountInUsd field to given value.

### HasTaxAmountInUsd

`func (o *InvoiceLineUpdateDto) HasTaxAmountInUsd() bool`

HasTaxAmountInUsd returns a boolean if a field has been set.

### GetTaxBaseAmountInUsd

`func (o *InvoiceLineUpdateDto) GetTaxBaseAmountInUsd() float64`

GetTaxBaseAmountInUsd returns the TaxBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUsdOk

`func (o *InvoiceLineUpdateDto) GetTaxBaseAmountInUsdOk() (*float64, bool)`

GetTaxBaseAmountInUsdOk returns a tuple with the TaxBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUsd

`func (o *InvoiceLineUpdateDto) SetTaxBaseAmountInUsd(v float64)`

SetTaxBaseAmountInUsd sets TaxBaseAmountInUsd field to given value.

### HasTaxBaseAmountInUsd

`func (o *InvoiceLineUpdateDto) HasTaxBaseAmountInUsd() bool`

HasTaxBaseAmountInUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


