# PriceListUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**PartnerVisible** | Pointer to **bool** |  | [optional] 
**UnitOfMeasureDependant** | Pointer to **bool** |  | [optional] 

## Methods

### NewPriceListUpdateDto

`func NewPriceListUpdateDto(name string, ) *PriceListUpdateDto`

NewPriceListUpdateDto instantiates a new PriceListUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListUpdateDtoWithDefaults

`func NewPriceListUpdateDtoWithDefaults() *PriceListUpdateDto`

NewPriceListUpdateDtoWithDefaults instantiates a new PriceListUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PriceListUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceListUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceListUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PriceListUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PriceListUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PriceListUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PriceListUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PriceListUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PriceListUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContext

`func (o *PriceListUpdateDto) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PriceListUpdateDto) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PriceListUpdateDto) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PriceListUpdateDto) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetStartDate

`func (o *PriceListUpdateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PriceListUpdateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PriceListUpdateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PriceListUpdateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PriceListUpdateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PriceListUpdateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PriceListUpdateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PriceListUpdateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PriceListUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PriceListUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PriceListUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PriceListUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *PriceListUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *PriceListUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetUnitId

`func (o *PriceListUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *PriceListUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *PriceListUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *PriceListUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *PriceListUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *PriceListUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *PriceListUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *PriceListUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *PriceListUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *PriceListUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *PriceListUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *PriceListUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPartnerVisible

`func (o *PriceListUpdateDto) GetPartnerVisible() bool`

GetPartnerVisible returns the PartnerVisible field if non-nil, zero value otherwise.

### GetPartnerVisibleOk

`func (o *PriceListUpdateDto) GetPartnerVisibleOk() (*bool, bool)`

GetPartnerVisibleOk returns a tuple with the PartnerVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerVisible

`func (o *PriceListUpdateDto) SetPartnerVisible(v bool)`

SetPartnerVisible sets PartnerVisible field to given value.

### HasPartnerVisible

`func (o *PriceListUpdateDto) HasPartnerVisible() bool`

HasPartnerVisible returns a boolean if a field has been set.

### GetUnitOfMeasureDependant

`func (o *PriceListUpdateDto) GetUnitOfMeasureDependant() bool`

GetUnitOfMeasureDependant returns the UnitOfMeasureDependant field if non-nil, zero value otherwise.

### GetUnitOfMeasureDependantOk

`func (o *PriceListUpdateDto) GetUnitOfMeasureDependantOk() (*bool, bool)`

GetUnitOfMeasureDependantOk returns a tuple with the UnitOfMeasureDependant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureDependant

`func (o *PriceListUpdateDto) SetUnitOfMeasureDependant(v bool)`

SetUnitOfMeasureDependant sets UnitOfMeasureDependant field to given value.

### HasUnitOfMeasureDependant

`func (o *PriceListUpdateDto) HasUnitOfMeasureDependant() bool`

HasUnitOfMeasureDependant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


