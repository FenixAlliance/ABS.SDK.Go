# BankTransactionUpdateDto

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
**BankProfileId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankTransactionUpdateDto

`func NewBankTransactionUpdateDto() *BankTransactionUpdateDto`

NewBankTransactionUpdateDto instantiates a new BankTransactionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionUpdateDtoWithDefaults

`func NewBankTransactionUpdateDtoWithDefaults() *BankTransactionUpdateDto`

NewBankTransactionUpdateDtoWithDefaults instantiates a new BankTransactionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BankTransactionUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransactionUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransactionUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTransactionUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BankTransactionUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BankTransactionUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *BankTransactionUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BankTransactionUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BankTransactionUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BankTransactionUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *BankTransactionUpdateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BankTransactionUpdateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BankTransactionUpdateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BankTransactionUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExternalDescription

`func (o *BankTransactionUpdateDto) GetExternalDescription() string`

GetExternalDescription returns the ExternalDescription field if non-nil, zero value otherwise.

### GetExternalDescriptionOk

`func (o *BankTransactionUpdateDto) GetExternalDescriptionOk() (*string, bool)`

GetExternalDescriptionOk returns a tuple with the ExternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDescription

`func (o *BankTransactionUpdateDto) SetExternalDescription(v string)`

SetExternalDescription sets ExternalDescription field to given value.

### HasExternalDescription

`func (o *BankTransactionUpdateDto) HasExternalDescription() bool`

HasExternalDescription returns a boolean if a field has been set.

### SetExternalDescriptionNil

`func (o *BankTransactionUpdateDto) SetExternalDescriptionNil(b bool)`

 SetExternalDescriptionNil sets the value for ExternalDescription to be an explicit nil

### UnsetExternalDescription
`func (o *BankTransactionUpdateDto) UnsetExternalDescription()`

UnsetExternalDescription ensures that no value is present for ExternalDescription, not even an explicit nil
### GetBasisQuantity

`func (o *BankTransactionUpdateDto) GetBasisQuantity() float64`

GetBasisQuantity returns the BasisQuantity field if non-nil, zero value otherwise.

### GetBasisQuantityOk

`func (o *BankTransactionUpdateDto) GetBasisQuantityOk() (*float64, bool)`

GetBasisQuantityOk returns a tuple with the BasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisQuantity

`func (o *BankTransactionUpdateDto) SetBasisQuantity(v float64)`

SetBasisQuantity sets BasisQuantity field to given value.

### HasBasisQuantity

`func (o *BankTransactionUpdateDto) HasBasisQuantity() bool`

HasBasisQuantity returns a boolean if a field has been set.

### GetBasisAmount

`func (o *BankTransactionUpdateDto) GetBasisAmount() float64`

GetBasisAmount returns the BasisAmount field if non-nil, zero value otherwise.

### GetBasisAmountOk

`func (o *BankTransactionUpdateDto) GetBasisAmountOk() (*float64, bool)`

GetBasisAmountOk returns a tuple with the BasisAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisAmount

`func (o *BankTransactionUpdateDto) SetBasisAmount(v float64)`

SetBasisAmount sets BasisAmount field to given value.

### HasBasisAmount

`func (o *BankTransactionUpdateDto) HasBasisAmount() bool`

HasBasisAmount returns a boolean if a field has been set.

### GetPercent

`func (o *BankTransactionUpdateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *BankTransactionUpdateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *BankTransactionUpdateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *BankTransactionUpdateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *BankTransactionUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *BankTransactionUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *BankTransactionUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *BankTransactionUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *BankTransactionUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *BankTransactionUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetUnitId

`func (o *BankTransactionUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *BankTransactionUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *BankTransactionUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *BankTransactionUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *BankTransactionUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *BankTransactionUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTransactionCategoryId

`func (o *BankTransactionUpdateDto) GetTransactionCategoryId() string`

GetTransactionCategoryId returns the TransactionCategoryId field if non-nil, zero value otherwise.

### GetTransactionCategoryIdOk

`func (o *BankTransactionUpdateDto) GetTransactionCategoryIdOk() (*string, bool)`

GetTransactionCategoryIdOk returns a tuple with the TransactionCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryId

`func (o *BankTransactionUpdateDto) SetTransactionCategoryId(v string)`

SetTransactionCategoryId sets TransactionCategoryId field to given value.

### HasTransactionCategoryId

`func (o *BankTransactionUpdateDto) HasTransactionCategoryId() bool`

HasTransactionCategoryId returns a boolean if a field has been set.

### SetTransactionCategoryIdNil

`func (o *BankTransactionUpdateDto) SetTransactionCategoryIdNil(b bool)`

 SetTransactionCategoryIdNil sets the value for TransactionCategoryId to be an explicit nil

### UnsetTransactionCategoryId
`func (o *BankTransactionUpdateDto) UnsetTransactionCategoryId()`

UnsetTransactionCategoryId ensures that no value is present for TransactionCategoryId, not even an explicit nil
### GetCurrencyId

`func (o *BankTransactionUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankTransactionUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankTransactionUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankTransactionUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankTransactionUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankTransactionUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetBankProfileId

`func (o *BankTransactionUpdateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankTransactionUpdateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankTransactionUpdateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankTransactionUpdateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankTransactionUpdateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankTransactionUpdateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankTransactionUpdateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankTransactionUpdateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankTransactionUpdateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankTransactionUpdateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankTransactionUpdateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankTransactionUpdateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


