# ItemPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**DiscountId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 

## Methods

### NewItemPriceDto

`func NewItemPriceDto() *ItemPriceDto`

NewItemPriceDto instantiates a new ItemPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceDtoWithDefaults

`func NewItemPriceDtoWithDefaults() *ItemPriceDto`

NewItemPriceDtoWithDefaults instantiates a new ItemPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPriceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPriceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPriceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPriceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemPriceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemPriceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemPriceDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPriceDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPriceDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPriceDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemPriceDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemPriceDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetItemId

`func (o *ItemPriceDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPriceDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPriceDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemPriceDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemPriceDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemPriceDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetUnitId

`func (o *ItemPriceDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ItemPriceDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ItemPriceDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ItemPriceDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *ItemPriceDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *ItemPriceDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetCurrencyId

`func (o *ItemPriceDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemPriceDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemPriceDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemPriceDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemPriceDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemPriceDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDiscountId

`func (o *ItemPriceDto) GetDiscountId() string`

GetDiscountId returns the DiscountId field if non-nil, zero value otherwise.

### GetDiscountIdOk

`func (o *ItemPriceDto) GetDiscountIdOk() (*string, bool)`

GetDiscountIdOk returns a tuple with the DiscountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountId

`func (o *ItemPriceDto) SetDiscountId(v string)`

SetDiscountId sets DiscountId field to given value.

### HasDiscountId

`func (o *ItemPriceDto) HasDiscountId() bool`

HasDiscountId returns a boolean if a field has been set.

### SetDiscountIdNil

`func (o *ItemPriceDto) SetDiscountIdNil(b bool)`

 SetDiscountIdNil sets the value for DiscountId to be an explicit nil

### UnsetDiscountId
`func (o *ItemPriceDto) UnsetDiscountId()`

UnsetDiscountId ensures that no value is present for DiscountId, not even an explicit nil
### GetUnitGroupId

`func (o *ItemPriceDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *ItemPriceDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *ItemPriceDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *ItemPriceDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *ItemPriceDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *ItemPriceDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriceListId

`func (o *ItemPriceDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ItemPriceDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ItemPriceDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ItemPriceDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ItemPriceDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ItemPriceDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDiscountListId

`func (o *ItemPriceDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *ItemPriceDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *ItemPriceDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *ItemPriceDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *ItemPriceDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *ItemPriceDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetRoundingPolicyId

`func (o *ItemPriceDto) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *ItemPriceDto) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *ItemPriceDto) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *ItemPriceDto) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *ItemPriceDto) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *ItemPriceDto) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
### GetEnrollmentId

`func (o *ItemPriceDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ItemPriceDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ItemPriceDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ItemPriceDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ItemPriceDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ItemPriceDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTenantId

`func (o *ItemPriceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPriceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPriceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPriceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPriceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPriceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetPrice

`func (o *ItemPriceDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemPriceDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemPriceDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ItemPriceDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPercent

`func (o *ItemPriceDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ItemPriceDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ItemPriceDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ItemPriceDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


