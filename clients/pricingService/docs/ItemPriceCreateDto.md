# ItemPriceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ItemId** | **string** |  | 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**RoundingPolicyId** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 

## Methods

### NewItemPriceCreateDto

`func NewItemPriceCreateDto(itemId string, ) *ItemPriceCreateDto`

NewItemPriceCreateDto instantiates a new ItemPriceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceCreateDtoWithDefaults

`func NewItemPriceCreateDtoWithDefaults() *ItemPriceCreateDto`

NewItemPriceCreateDtoWithDefaults instantiates a new ItemPriceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPriceCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPriceCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPriceCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPriceCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemPriceCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPriceCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPriceCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPriceCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPriceCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPriceCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPriceCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetUnitId

`func (o *ItemPriceCreateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ItemPriceCreateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ItemPriceCreateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ItemPriceCreateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *ItemPriceCreateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *ItemPriceCreateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTenantId

`func (o *ItemPriceCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPriceCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPriceCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPriceCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPriceCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPriceCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCurrencyId

`func (o *ItemPriceCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemPriceCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemPriceCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemPriceCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemPriceCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemPriceCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPriceListId

`func (o *ItemPriceCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ItemPriceCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ItemPriceCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ItemPriceCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ItemPriceCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ItemPriceCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetUnitGroupId

`func (o *ItemPriceCreateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *ItemPriceCreateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *ItemPriceCreateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *ItemPriceCreateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *ItemPriceCreateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *ItemPriceCreateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetEnrollmentId

`func (o *ItemPriceCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ItemPriceCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ItemPriceCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ItemPriceCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ItemPriceCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ItemPriceCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDiscountListId

`func (o *ItemPriceCreateDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *ItemPriceCreateDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *ItemPriceCreateDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *ItemPriceCreateDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *ItemPriceCreateDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *ItemPriceCreateDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetRoundingPolicyId

`func (o *ItemPriceCreateDto) GetRoundingPolicyId() string`

GetRoundingPolicyId returns the RoundingPolicyId field if non-nil, zero value otherwise.

### GetRoundingPolicyIdOk

`func (o *ItemPriceCreateDto) GetRoundingPolicyIdOk() (*string, bool)`

GetRoundingPolicyIdOk returns a tuple with the RoundingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPolicyId

`func (o *ItemPriceCreateDto) SetRoundingPolicyId(v string)`

SetRoundingPolicyId sets RoundingPolicyId field to given value.

### HasRoundingPolicyId

`func (o *ItemPriceCreateDto) HasRoundingPolicyId() bool`

HasRoundingPolicyId returns a boolean if a field has been set.

### SetRoundingPolicyIdNil

`func (o *ItemPriceCreateDto) SetRoundingPolicyIdNil(b bool)`

 SetRoundingPolicyIdNil sets the value for RoundingPolicyId to be an explicit nil

### UnsetRoundingPolicyId
`func (o *ItemPriceCreateDto) UnsetRoundingPolicyId()`

UnsetRoundingPolicyId ensures that no value is present for RoundingPolicyId, not even an explicit nil
### GetPrice

`func (o *ItemPriceCreateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemPriceCreateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemPriceCreateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ItemPriceCreateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPercent

`func (o *ItemPriceCreateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ItemPriceCreateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ItemPriceCreateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ItemPriceCreateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


