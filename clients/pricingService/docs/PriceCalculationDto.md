# PriceCalculationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**PriceId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**DiscountId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 
**EffectiveDiscountPercent** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalBaseAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalAmount** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewPriceCalculationDto

`func NewPriceCalculationDto() *PriceCalculationDto`

NewPriceCalculationDto instantiates a new PriceCalculationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceCalculationDtoWithDefaults

`func NewPriceCalculationDtoWithDefaults() *PriceCalculationDto`

NewPriceCalculationDtoWithDefaults instantiates a new PriceCalculationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceCalculationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceCalculationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceCalculationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PriceCalculationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PriceCalculationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PriceCalculationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PriceCalculationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PriceCalculationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PriceCalculationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PriceCalculationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PriceCalculationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PriceCalculationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetItemId

`func (o *PriceCalculationDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PriceCalculationDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PriceCalculationDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PriceCalculationDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *PriceCalculationDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *PriceCalculationDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetUnitId

`func (o *PriceCalculationDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *PriceCalculationDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *PriceCalculationDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *PriceCalculationDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *PriceCalculationDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *PriceCalculationDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *PriceCalculationDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *PriceCalculationDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *PriceCalculationDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *PriceCalculationDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *PriceCalculationDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *PriceCalculationDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriceId

`func (o *PriceCalculationDto) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *PriceCalculationDto) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *PriceCalculationDto) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *PriceCalculationDto) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *PriceCalculationDto) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *PriceCalculationDto) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetPriceListId

`func (o *PriceCalculationDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *PriceCalculationDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *PriceCalculationDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *PriceCalculationDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *PriceCalculationDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *PriceCalculationDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDiscountId

`func (o *PriceCalculationDto) GetDiscountId() string`

GetDiscountId returns the DiscountId field if non-nil, zero value otherwise.

### GetDiscountIdOk

`func (o *PriceCalculationDto) GetDiscountIdOk() (*string, bool)`

GetDiscountIdOk returns a tuple with the DiscountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountId

`func (o *PriceCalculationDto) SetDiscountId(v string)`

SetDiscountId sets DiscountId field to given value.

### HasDiscountId

`func (o *PriceCalculationDto) HasDiscountId() bool`

HasDiscountId returns a boolean if a field has been set.

### SetDiscountIdNil

`func (o *PriceCalculationDto) SetDiscountIdNil(b bool)`

 SetDiscountIdNil sets the value for DiscountId to be an explicit nil

### UnsetDiscountId
`func (o *PriceCalculationDto) UnsetDiscountId()`

UnsetDiscountId ensures that no value is present for DiscountId, not even an explicit nil
### GetDiscountListId

`func (o *PriceCalculationDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *PriceCalculationDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *PriceCalculationDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *PriceCalculationDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *PriceCalculationDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *PriceCalculationDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetTenantId

`func (o *PriceCalculationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PriceCalculationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PriceCalculationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PriceCalculationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PriceCalculationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PriceCalculationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PriceCalculationDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PriceCalculationDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PriceCalculationDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PriceCalculationDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PriceCalculationDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PriceCalculationDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetRoundingPolicyId

`func (o *PriceCalculationDto) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *PriceCalculationDto) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *PriceCalculationDto) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *PriceCalculationDto) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *PriceCalculationDto) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *PriceCalculationDto) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
### GetEffectiveDiscountPercent

`func (o *PriceCalculationDto) GetEffectiveDiscountPercent() float64`

GetEffectiveDiscountPercent returns the EffectiveDiscountPercent field if non-nil, zero value otherwise.

### GetEffectiveDiscountPercentOk

`func (o *PriceCalculationDto) GetEffectiveDiscountPercentOk() (*float64, bool)`

GetEffectiveDiscountPercentOk returns a tuple with the EffectiveDiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDiscountPercent

`func (o *PriceCalculationDto) SetEffectiveDiscountPercent(v float64)`

SetEffectiveDiscountPercent sets EffectiveDiscountPercent field to given value.

### HasEffectiveDiscountPercent

`func (o *PriceCalculationDto) HasEffectiveDiscountPercent() bool`

HasEffectiveDiscountPercent returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PriceCalculationDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PriceCalculationDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PriceCalculationDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PriceCalculationDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *PriceCalculationDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *PriceCalculationDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalBaseAmount

`func (o *PriceCalculationDto) GetTotalBaseAmount() Money`

GetTotalBaseAmount returns the TotalBaseAmount field if non-nil, zero value otherwise.

### GetTotalBaseAmountOk

`func (o *PriceCalculationDto) GetTotalBaseAmountOk() (*Money, bool)`

GetTotalBaseAmountOk returns a tuple with the TotalBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseAmount

`func (o *PriceCalculationDto) SetTotalBaseAmount(v Money)`

SetTotalBaseAmount sets TotalBaseAmount field to given value.

### HasTotalBaseAmount

`func (o *PriceCalculationDto) HasTotalBaseAmount() bool`

HasTotalBaseAmount returns a boolean if a field has been set.

### GetTotalDiscountsAmount

`func (o *PriceCalculationDto) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *PriceCalculationDto) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *PriceCalculationDto) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *PriceCalculationDto) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurchargesAmount

`func (o *PriceCalculationDto) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *PriceCalculationDto) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *PriceCalculationDto) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *PriceCalculationDto) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalShippingAmount

`func (o *PriceCalculationDto) GetTotalShippingAmount() Money`

GetTotalShippingAmount returns the TotalShippingAmount field if non-nil, zero value otherwise.

### GetTotalShippingAmountOk

`func (o *PriceCalculationDto) GetTotalShippingAmountOk() (*Money, bool)`

GetTotalShippingAmountOk returns a tuple with the TotalShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingAmount

`func (o *PriceCalculationDto) SetTotalShippingAmount(v Money)`

SetTotalShippingAmount sets TotalShippingAmount field to given value.

### HasTotalShippingAmount

`func (o *PriceCalculationDto) HasTotalShippingAmount() bool`

HasTotalShippingAmount returns a boolean if a field has been set.

### GetTotalShippingTaxAmount

`func (o *PriceCalculationDto) GetTotalShippingTaxAmount() Money`

GetTotalShippingTaxAmount returns the TotalShippingTaxAmount field if non-nil, zero value otherwise.

### GetTotalShippingTaxAmountOk

`func (o *PriceCalculationDto) GetTotalShippingTaxAmountOk() (*Money, bool)`

GetTotalShippingTaxAmountOk returns a tuple with the TotalShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxAmount

`func (o *PriceCalculationDto) SetTotalShippingTaxAmount(v Money)`

SetTotalShippingTaxAmount sets TotalShippingTaxAmount field to given value.

### HasTotalShippingTaxAmount

`func (o *PriceCalculationDto) HasTotalShippingTaxAmount() bool`

HasTotalShippingTaxAmount returns a boolean if a field has been set.

### GetTotalTaxAmount

`func (o *PriceCalculationDto) GetTotalTaxAmount() Money`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *PriceCalculationDto) GetTotalTaxAmountOk() (*Money, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *PriceCalculationDto) SetTotalTaxAmount(v Money)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *PriceCalculationDto) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *PriceCalculationDto) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PriceCalculationDto) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PriceCalculationDto) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *PriceCalculationDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


