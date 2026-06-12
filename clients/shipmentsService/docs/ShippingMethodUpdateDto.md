# ShippingMethodUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**TaxIncluded** | Pointer to **bool** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ShippingClassCalculationType** | Pointer to **string** |  | [optional] 

## Methods

### NewShippingMethodUpdateDto

`func NewShippingMethodUpdateDto() *ShippingMethodUpdateDto`

NewShippingMethodUpdateDto instantiates a new ShippingMethodUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodUpdateDtoWithDefaults

`func NewShippingMethodUpdateDtoWithDefaults() *ShippingMethodUpdateDto`

NewShippingMethodUpdateDtoWithDefaults instantiates a new ShippingMethodUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingMethodUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethodUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethodUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingMethodUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShippingMethodUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShippingMethodUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ShippingMethodUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShippingMethodUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShippingMethodUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShippingMethodUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShippingMethodUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShippingMethodUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCost

`func (o *ShippingMethodUpdateDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ShippingMethodUpdateDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ShippingMethodUpdateDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ShippingMethodUpdateDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTaxable

`func (o *ShippingMethodUpdateDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ShippingMethodUpdateDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ShippingMethodUpdateDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ShippingMethodUpdateDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *ShippingMethodUpdateDto) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *ShippingMethodUpdateDto) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *ShippingMethodUpdateDto) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *ShippingMethodUpdateDto) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ShippingMethodUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ShippingMethodUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ShippingMethodUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ShippingMethodUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ShippingMethodUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ShippingMethodUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetShippingClassCalculationType

`func (o *ShippingMethodUpdateDto) GetShippingClassCalculationType() string`

GetShippingClassCalculationType returns the ShippingClassCalculationType field if non-nil, zero value otherwise.

### GetShippingClassCalculationTypeOk

`func (o *ShippingMethodUpdateDto) GetShippingClassCalculationTypeOk() (*string, bool)`

GetShippingClassCalculationTypeOk returns a tuple with the ShippingClassCalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingClassCalculationType

`func (o *ShippingMethodUpdateDto) SetShippingClassCalculationType(v string)`

SetShippingClassCalculationType sets ShippingClassCalculationType field to given value.

### HasShippingClassCalculationType

`func (o *ShippingMethodUpdateDto) HasShippingClassCalculationType() bool`

HasShippingClassCalculationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


