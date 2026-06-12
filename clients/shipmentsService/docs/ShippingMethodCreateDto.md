# ShippingMethodCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**TaxIncluded** | Pointer to **bool** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ShippingClassCalculationType** | Pointer to **string** |  | [optional] 

## Methods

### NewShippingMethodCreateDto

`func NewShippingMethodCreateDto(name string, ) *ShippingMethodCreateDto`

NewShippingMethodCreateDto instantiates a new ShippingMethodCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodCreateDtoWithDefaults

`func NewShippingMethodCreateDtoWithDefaults() *ShippingMethodCreateDto`

NewShippingMethodCreateDtoWithDefaults instantiates a new ShippingMethodCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingMethodCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethodCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethodCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingMethodCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShippingMethodCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShippingMethodCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShippingMethodCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShippingMethodCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ShippingMethodCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingMethodCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingMethodCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShippingMethodCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShippingMethodCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShippingMethodCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShippingMethodCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShippingMethodCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShippingMethodCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCost

`func (o *ShippingMethodCreateDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ShippingMethodCreateDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ShippingMethodCreateDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ShippingMethodCreateDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTaxable

`func (o *ShippingMethodCreateDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ShippingMethodCreateDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ShippingMethodCreateDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ShippingMethodCreateDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *ShippingMethodCreateDto) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *ShippingMethodCreateDto) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *ShippingMethodCreateDto) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *ShippingMethodCreateDto) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ShippingMethodCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ShippingMethodCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ShippingMethodCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ShippingMethodCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ShippingMethodCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ShippingMethodCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetShippingClassCalculationType

`func (o *ShippingMethodCreateDto) GetShippingClassCalculationType() string`

GetShippingClassCalculationType returns the ShippingClassCalculationType field if non-nil, zero value otherwise.

### GetShippingClassCalculationTypeOk

`func (o *ShippingMethodCreateDto) GetShippingClassCalculationTypeOk() (*string, bool)`

GetShippingClassCalculationTypeOk returns a tuple with the ShippingClassCalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingClassCalculationType

`func (o *ShippingMethodCreateDto) SetShippingClassCalculationType(v string)`

SetShippingClassCalculationType sets ShippingClassCalculationType field to given value.

### HasShippingClassCalculationType

`func (o *ShippingMethodCreateDto) HasShippingClassCalculationType() bool`

HasShippingClassCalculationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


