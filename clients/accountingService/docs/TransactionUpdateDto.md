# TransactionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**ExternalDescription** | Pointer to **NullableString** |  | [optional] 
**BasisQuantity** | Pointer to **float64** |  | [optional] 
**BasisAmount** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**TransactionCategoryId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransactionUpdateDto

`func NewTransactionUpdateDto() *TransactionUpdateDto`

NewTransactionUpdateDto instantiates a new TransactionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUpdateDtoWithDefaults

`func NewTransactionUpdateDtoWithDefaults() *TransactionUpdateDto`

NewTransactionUpdateDtoWithDefaults instantiates a new TransactionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TransactionUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *TransactionUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TransactionUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TransactionUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TransactionUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *TransactionUpdateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TransactionUpdateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TransactionUpdateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TransactionUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExternalDescription

`func (o *TransactionUpdateDto) GetExternalDescription() string`

GetExternalDescription returns the ExternalDescription field if non-nil, zero value otherwise.

### GetExternalDescriptionOk

`func (o *TransactionUpdateDto) GetExternalDescriptionOk() (*string, bool)`

GetExternalDescriptionOk returns a tuple with the ExternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDescription

`func (o *TransactionUpdateDto) SetExternalDescription(v string)`

SetExternalDescription sets ExternalDescription field to given value.

### HasExternalDescription

`func (o *TransactionUpdateDto) HasExternalDescription() bool`

HasExternalDescription returns a boolean if a field has been set.

### SetExternalDescriptionNil

`func (o *TransactionUpdateDto) SetExternalDescriptionNil(b bool)`

 SetExternalDescriptionNil sets the value for ExternalDescription to be an explicit nil

### UnsetExternalDescription
`func (o *TransactionUpdateDto) UnsetExternalDescription()`

UnsetExternalDescription ensures that no value is present for ExternalDescription, not even an explicit nil
### GetBasisQuantity

`func (o *TransactionUpdateDto) GetBasisQuantity() float64`

GetBasisQuantity returns the BasisQuantity field if non-nil, zero value otherwise.

### GetBasisQuantityOk

`func (o *TransactionUpdateDto) GetBasisQuantityOk() (*float64, bool)`

GetBasisQuantityOk returns a tuple with the BasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisQuantity

`func (o *TransactionUpdateDto) SetBasisQuantity(v float64)`

SetBasisQuantity sets BasisQuantity field to given value.

### HasBasisQuantity

`func (o *TransactionUpdateDto) HasBasisQuantity() bool`

HasBasisQuantity returns a boolean if a field has been set.

### GetBasisAmount

`func (o *TransactionUpdateDto) GetBasisAmount() float64`

GetBasisAmount returns the BasisAmount field if non-nil, zero value otherwise.

### GetBasisAmountOk

`func (o *TransactionUpdateDto) GetBasisAmountOk() (*float64, bool)`

GetBasisAmountOk returns a tuple with the BasisAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisAmount

`func (o *TransactionUpdateDto) SetBasisAmount(v float64)`

SetBasisAmount sets BasisAmount field to given value.

### HasBasisAmount

`func (o *TransactionUpdateDto) HasBasisAmount() bool`

HasBasisAmount returns a boolean if a field has been set.

### GetPercent

`func (o *TransactionUpdateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *TransactionUpdateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *TransactionUpdateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *TransactionUpdateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *TransactionUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *TransactionUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *TransactionUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *TransactionUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *TransactionUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *TransactionUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetUnitId

`func (o *TransactionUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *TransactionUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *TransactionUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *TransactionUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *TransactionUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *TransactionUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTransactionCategoryId

`func (o *TransactionUpdateDto) GetTransactionCategoryId() string`

GetTransactionCategoryId returns the TransactionCategoryId field if non-nil, zero value otherwise.

### GetTransactionCategoryIdOk

`func (o *TransactionUpdateDto) GetTransactionCategoryIdOk() (*string, bool)`

GetTransactionCategoryIdOk returns a tuple with the TransactionCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryId

`func (o *TransactionUpdateDto) SetTransactionCategoryId(v string)`

SetTransactionCategoryId sets TransactionCategoryId field to given value.

### HasTransactionCategoryId

`func (o *TransactionUpdateDto) HasTransactionCategoryId() bool`

HasTransactionCategoryId returns a boolean if a field has been set.

### SetTransactionCategoryIdNil

`func (o *TransactionUpdateDto) SetTransactionCategoryIdNil(b bool)`

 SetTransactionCategoryIdNil sets the value for TransactionCategoryId to be an explicit nil

### UnsetTransactionCategoryId
`func (o *TransactionUpdateDto) UnsetTransactionCategoryId()`

UnsetTransactionCategoryId ensures that no value is present for TransactionCategoryId, not even an explicit nil
### GetCurrencyId

`func (o *TransactionUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TransactionUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TransactionUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TransactionUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TransactionUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TransactionUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *TransactionUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TransactionUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TransactionUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TransactionUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TransactionUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TransactionUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TransactionUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TransactionUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TransactionUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TransactionUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TransactionUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TransactionUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


