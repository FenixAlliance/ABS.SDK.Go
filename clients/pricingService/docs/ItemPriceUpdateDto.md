# ItemPriceUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPriceUpdateDto

`func NewItemPriceUpdateDto() *ItemPriceUpdateDto`

NewItemPriceUpdateDto instantiates a new ItemPriceUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceUpdateDtoWithDefaults

`func NewItemPriceUpdateDtoWithDefaults() *ItemPriceUpdateDto`

NewItemPriceUpdateDtoWithDefaults instantiates a new ItemPriceUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ItemPriceUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemPriceUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemPriceUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ItemPriceUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPriceUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPriceUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPriceUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemPriceUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemPriceUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemPriceUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetUnitId

`func (o *ItemPriceUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ItemPriceUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ItemPriceUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ItemPriceUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *ItemPriceUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *ItemPriceUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetPercent

`func (o *ItemPriceUpdateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ItemPriceUpdateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ItemPriceUpdateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ItemPriceUpdateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *ItemPriceUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *ItemPriceUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *ItemPriceUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *ItemPriceUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *ItemPriceUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *ItemPriceUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetCurrencyId

`func (o *ItemPriceUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemPriceUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemPriceUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemPriceUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemPriceUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemPriceUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDiscountListId

`func (o *ItemPriceUpdateDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *ItemPriceUpdateDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *ItemPriceUpdateDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *ItemPriceUpdateDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *ItemPriceUpdateDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *ItemPriceUpdateDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetRoundingPolicyId

`func (o *ItemPriceUpdateDto) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *ItemPriceUpdateDto) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *ItemPriceUpdateDto) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *ItemPriceUpdateDto) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *ItemPriceUpdateDto) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *ItemPriceUpdateDto) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


