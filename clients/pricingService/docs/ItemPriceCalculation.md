# ItemPriceCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float64** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**Item** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**PriceId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**DiscountId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**EffectiveDiscountPercent** | Pointer to **float64** |  | [optional] [readonly] 
**EffectiveTaxPercent** | Pointer to **float64** |  | [optional] [readonly] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to [**CurrencyId**](CurrencyId.md) |  | [optional] 
**TotalBaseAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalProfitAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDetailAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxBaseAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalWTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingCostAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalAmount** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewItemPriceCalculation

`func NewItemPriceCalculation() *ItemPriceCalculation`

NewItemPriceCalculation instantiates a new ItemPriceCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceCalculationWithDefaults

`func NewItemPriceCalculationWithDefaults() *ItemPriceCalculation`

NewItemPriceCalculationWithDefaults instantiates a new ItemPriceCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *ItemPriceCalculation) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemPriceCalculation) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemPriceCalculation) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemPriceCalculation) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPriceCalculation) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPriceCalculation) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPriceCalculation) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemPriceCalculation) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemPriceCalculation) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemPriceCalculation) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItem

`func (o *ItemPriceCalculation) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemPriceCalculation) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemPriceCalculation) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ItemPriceCalculation) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *ItemPriceCalculation) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *ItemPriceCalculation) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetUnitId

`func (o *ItemPriceCalculation) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ItemPriceCalculation) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ItemPriceCalculation) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ItemPriceCalculation) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *ItemPriceCalculation) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *ItemPriceCalculation) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *ItemPriceCalculation) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *ItemPriceCalculation) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *ItemPriceCalculation) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *ItemPriceCalculation) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *ItemPriceCalculation) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *ItemPriceCalculation) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriceId

`func (o *ItemPriceCalculation) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *ItemPriceCalculation) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *ItemPriceCalculation) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *ItemPriceCalculation) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *ItemPriceCalculation) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *ItemPriceCalculation) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetPriceListId

`func (o *ItemPriceCalculation) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ItemPriceCalculation) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ItemPriceCalculation) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ItemPriceCalculation) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ItemPriceCalculation) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ItemPriceCalculation) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDiscountId

`func (o *ItemPriceCalculation) GetDiscountId() string`

GetDiscountId returns the DiscountId field if non-nil, zero value otherwise.

### GetDiscountIdOk

`func (o *ItemPriceCalculation) GetDiscountIdOk() (*string, bool)`

GetDiscountIdOk returns a tuple with the DiscountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountId

`func (o *ItemPriceCalculation) SetDiscountId(v string)`

SetDiscountId sets DiscountId field to given value.

### HasDiscountId

`func (o *ItemPriceCalculation) HasDiscountId() bool`

HasDiscountId returns a boolean if a field has been set.

### SetDiscountIdNil

`func (o *ItemPriceCalculation) SetDiscountIdNil(b bool)`

 SetDiscountIdNil sets the value for DiscountId to be an explicit nil

### UnsetDiscountId
`func (o *ItemPriceCalculation) UnsetDiscountId()`

UnsetDiscountId ensures that no value is present for DiscountId, not even an explicit nil
### GetDiscountListId

`func (o *ItemPriceCalculation) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *ItemPriceCalculation) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *ItemPriceCalculation) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *ItemPriceCalculation) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *ItemPriceCalculation) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *ItemPriceCalculation) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetTenantId

`func (o *ItemPriceCalculation) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPriceCalculation) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPriceCalculation) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPriceCalculation) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPriceCalculation) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPriceCalculation) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ItemPriceCalculation) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ItemPriceCalculation) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ItemPriceCalculation) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ItemPriceCalculation) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ItemPriceCalculation) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ItemPriceCalculation) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetRoundingPolicyId

`func (o *ItemPriceCalculation) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *ItemPriceCalculation) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *ItemPriceCalculation) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *ItemPriceCalculation) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *ItemPriceCalculation) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *ItemPriceCalculation) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
### GetTimestamp

`func (o *ItemPriceCalculation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPriceCalculation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPriceCalculation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPriceCalculation) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEffectiveDiscountPercent

`func (o *ItemPriceCalculation) GetEffectiveDiscountPercent() float64`

GetEffectiveDiscountPercent returns the EffectiveDiscountPercent field if non-nil, zero value otherwise.

### GetEffectiveDiscountPercentOk

`func (o *ItemPriceCalculation) GetEffectiveDiscountPercentOk() (*float64, bool)`

GetEffectiveDiscountPercentOk returns a tuple with the EffectiveDiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDiscountPercent

`func (o *ItemPriceCalculation) SetEffectiveDiscountPercent(v float64)`

SetEffectiveDiscountPercent sets EffectiveDiscountPercent field to given value.

### HasEffectiveDiscountPercent

`func (o *ItemPriceCalculation) HasEffectiveDiscountPercent() bool`

HasEffectiveDiscountPercent returns a boolean if a field has been set.

### GetEffectiveTaxPercent

`func (o *ItemPriceCalculation) GetEffectiveTaxPercent() float64`

GetEffectiveTaxPercent returns the EffectiveTaxPercent field if non-nil, zero value otherwise.

### GetEffectiveTaxPercentOk

`func (o *ItemPriceCalculation) GetEffectiveTaxPercentOk() (*float64, bool)`

GetEffectiveTaxPercentOk returns a tuple with the EffectiveTaxPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTaxPercent

`func (o *ItemPriceCalculation) SetEffectiveTaxPercent(v float64)`

SetEffectiveTaxPercent sets EffectiveTaxPercent field to given value.

### HasEffectiveTaxPercent

`func (o *ItemPriceCalculation) HasEffectiveTaxPercent() bool`

HasEffectiveTaxPercent returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ItemPriceCalculation) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemPriceCalculation) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemPriceCalculation) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemPriceCalculation) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemPriceCalculation) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemPriceCalculation) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrency

`func (o *ItemPriceCalculation) GetCurrency() CurrencyId`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ItemPriceCalculation) GetCurrencyOk() (*CurrencyId, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ItemPriceCalculation) SetCurrency(v CurrencyId)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ItemPriceCalculation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalBaseAmount

`func (o *ItemPriceCalculation) GetTotalBaseAmount() Money`

GetTotalBaseAmount returns the TotalBaseAmount field if non-nil, zero value otherwise.

### GetTotalBaseAmountOk

`func (o *ItemPriceCalculation) GetTotalBaseAmountOk() (*Money, bool)`

GetTotalBaseAmountOk returns a tuple with the TotalBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseAmount

`func (o *ItemPriceCalculation) SetTotalBaseAmount(v Money)`

SetTotalBaseAmount sets TotalBaseAmount field to given value.

### HasTotalBaseAmount

`func (o *ItemPriceCalculation) HasTotalBaseAmount() bool`

HasTotalBaseAmount returns a boolean if a field has been set.

### GetTotalProfitAmount

`func (o *ItemPriceCalculation) GetTotalProfitAmount() Money`

GetTotalProfitAmount returns the TotalProfitAmount field if non-nil, zero value otherwise.

### GetTotalProfitAmountOk

`func (o *ItemPriceCalculation) GetTotalProfitAmountOk() (*Money, bool)`

GetTotalProfitAmountOk returns a tuple with the TotalProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitAmount

`func (o *ItemPriceCalculation) SetTotalProfitAmount(v Money)`

SetTotalProfitAmount sets TotalProfitAmount field to given value.

### HasTotalProfitAmount

`func (o *ItemPriceCalculation) HasTotalProfitAmount() bool`

HasTotalProfitAmount returns a boolean if a field has been set.

### GetTotalDetailAmount

`func (o *ItemPriceCalculation) GetTotalDetailAmount() Money`

GetTotalDetailAmount returns the TotalDetailAmount field if non-nil, zero value otherwise.

### GetTotalDetailAmountOk

`func (o *ItemPriceCalculation) GetTotalDetailAmountOk() (*Money, bool)`

GetTotalDetailAmountOk returns a tuple with the TotalDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmount

`func (o *ItemPriceCalculation) SetTotalDetailAmount(v Money)`

SetTotalDetailAmount sets TotalDetailAmount field to given value.

### HasTotalDetailAmount

`func (o *ItemPriceCalculation) HasTotalDetailAmount() bool`

HasTotalDetailAmount returns a boolean if a field has been set.

### GetTotalDiscountsAmount

`func (o *ItemPriceCalculation) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *ItemPriceCalculation) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *ItemPriceCalculation) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *ItemPriceCalculation) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurchargesAmount

`func (o *ItemPriceCalculation) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *ItemPriceCalculation) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *ItemPriceCalculation) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *ItemPriceCalculation) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalTaxBaseAmount

`func (o *ItemPriceCalculation) GetTotalTaxBaseAmount() Money`

GetTotalTaxBaseAmount returns the TotalTaxBaseAmount field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountOk

`func (o *ItemPriceCalculation) GetTotalTaxBaseAmountOk() (*Money, bool)`

GetTotalTaxBaseAmountOk returns a tuple with the TotalTaxBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmount

`func (o *ItemPriceCalculation) SetTotalTaxBaseAmount(v Money)`

SetTotalTaxBaseAmount sets TotalTaxBaseAmount field to given value.

### HasTotalTaxBaseAmount

`func (o *ItemPriceCalculation) HasTotalTaxBaseAmount() bool`

HasTotalTaxBaseAmount returns a boolean if a field has been set.

### GetTotalTaxAmount

`func (o *ItemPriceCalculation) GetTotalTaxAmount() Money`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *ItemPriceCalculation) GetTotalTaxAmountOk() (*Money, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *ItemPriceCalculation) SetTotalTaxAmount(v Money)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *ItemPriceCalculation) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### GetTotalWTaxAmount

`func (o *ItemPriceCalculation) GetTotalWTaxAmount() Money`

GetTotalWTaxAmount returns the TotalWTaxAmount field if non-nil, zero value otherwise.

### GetTotalWTaxAmountOk

`func (o *ItemPriceCalculation) GetTotalWTaxAmountOk() (*Money, bool)`

GetTotalWTaxAmountOk returns a tuple with the TotalWTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWTaxAmount

`func (o *ItemPriceCalculation) SetTotalWTaxAmount(v Money)`

SetTotalWTaxAmount sets TotalWTaxAmount field to given value.

### HasTotalWTaxAmount

`func (o *ItemPriceCalculation) HasTotalWTaxAmount() bool`

HasTotalWTaxAmount returns a boolean if a field has been set.

### GetTotalShippingCostAmount

`func (o *ItemPriceCalculation) GetTotalShippingCostAmount() Money`

GetTotalShippingCostAmount returns the TotalShippingCostAmount field if non-nil, zero value otherwise.

### GetTotalShippingCostAmountOk

`func (o *ItemPriceCalculation) GetTotalShippingCostAmountOk() (*Money, bool)`

GetTotalShippingCostAmountOk returns a tuple with the TotalShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostAmount

`func (o *ItemPriceCalculation) SetTotalShippingCostAmount(v Money)`

SetTotalShippingCostAmount sets TotalShippingCostAmount field to given value.

### HasTotalShippingCostAmount

`func (o *ItemPriceCalculation) HasTotalShippingCostAmount() bool`

HasTotalShippingCostAmount returns a boolean if a field has been set.

### GetTotalShippingTaxAmount

`func (o *ItemPriceCalculation) GetTotalShippingTaxAmount() Money`

GetTotalShippingTaxAmount returns the TotalShippingTaxAmount field if non-nil, zero value otherwise.

### GetTotalShippingTaxAmountOk

`func (o *ItemPriceCalculation) GetTotalShippingTaxAmountOk() (*Money, bool)`

GetTotalShippingTaxAmountOk returns a tuple with the TotalShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxAmount

`func (o *ItemPriceCalculation) SetTotalShippingTaxAmount(v Money)`

SetTotalShippingTaxAmount sets TotalShippingTaxAmount field to given value.

### HasTotalShippingTaxAmount

`func (o *ItemPriceCalculation) HasTotalShippingTaxAmount() bool`

HasTotalShippingTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *ItemPriceCalculation) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ItemPriceCalculation) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ItemPriceCalculation) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ItemPriceCalculation) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


