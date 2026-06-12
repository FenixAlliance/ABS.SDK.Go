# BankTransactionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewBankTransactionCreateDto

`func NewBankTransactionCreateDto() *BankTransactionCreateDto`

NewBankTransactionCreateDto instantiates a new BankTransactionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionCreateDtoWithDefaults

`func NewBankTransactionCreateDtoWithDefaults() *BankTransactionCreateDto`

NewBankTransactionCreateDtoWithDefaults instantiates a new BankTransactionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransactionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransactionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransactionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankTransactionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BankTransactionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankTransactionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankTransactionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankTransactionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *BankTransactionCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransactionCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransactionCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTransactionCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BankTransactionCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BankTransactionCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *BankTransactionCreateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BankTransactionCreateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BankTransactionCreateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BankTransactionCreateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *BankTransactionCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BankTransactionCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BankTransactionCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BankTransactionCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExternalDescription

`func (o *BankTransactionCreateDto) GetExternalDescription() string`

GetExternalDescription returns the ExternalDescription field if non-nil, zero value otherwise.

### GetExternalDescriptionOk

`func (o *BankTransactionCreateDto) GetExternalDescriptionOk() (*string, bool)`

GetExternalDescriptionOk returns a tuple with the ExternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDescription

`func (o *BankTransactionCreateDto) SetExternalDescription(v string)`

SetExternalDescription sets ExternalDescription field to given value.

### HasExternalDescription

`func (o *BankTransactionCreateDto) HasExternalDescription() bool`

HasExternalDescription returns a boolean if a field has been set.

### SetExternalDescriptionNil

`func (o *BankTransactionCreateDto) SetExternalDescriptionNil(b bool)`

 SetExternalDescriptionNil sets the value for ExternalDescription to be an explicit nil

### UnsetExternalDescription
`func (o *BankTransactionCreateDto) UnsetExternalDescription()`

UnsetExternalDescription ensures that no value is present for ExternalDescription, not even an explicit nil
### GetBasisQuantity

`func (o *BankTransactionCreateDto) GetBasisQuantity() float64`

GetBasisQuantity returns the BasisQuantity field if non-nil, zero value otherwise.

### GetBasisQuantityOk

`func (o *BankTransactionCreateDto) GetBasisQuantityOk() (*float64, bool)`

GetBasisQuantityOk returns a tuple with the BasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisQuantity

`func (o *BankTransactionCreateDto) SetBasisQuantity(v float64)`

SetBasisQuantity sets BasisQuantity field to given value.

### HasBasisQuantity

`func (o *BankTransactionCreateDto) HasBasisQuantity() bool`

HasBasisQuantity returns a boolean if a field has been set.

### GetBasisAmount

`func (o *BankTransactionCreateDto) GetBasisAmount() float64`

GetBasisAmount returns the BasisAmount field if non-nil, zero value otherwise.

### GetBasisAmountOk

`func (o *BankTransactionCreateDto) GetBasisAmountOk() (*float64, bool)`

GetBasisAmountOk returns a tuple with the BasisAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisAmount

`func (o *BankTransactionCreateDto) SetBasisAmount(v float64)`

SetBasisAmount sets BasisAmount field to given value.

### HasBasisAmount

`func (o *BankTransactionCreateDto) HasBasisAmount() bool`

HasBasisAmount returns a boolean if a field has been set.

### GetPercent

`func (o *BankTransactionCreateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *BankTransactionCreateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *BankTransactionCreateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *BankTransactionCreateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *BankTransactionCreateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *BankTransactionCreateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *BankTransactionCreateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *BankTransactionCreateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *BankTransactionCreateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *BankTransactionCreateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetUnitId

`func (o *BankTransactionCreateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *BankTransactionCreateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *BankTransactionCreateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *BankTransactionCreateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *BankTransactionCreateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *BankTransactionCreateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTransactionCategoryId

`func (o *BankTransactionCreateDto) GetTransactionCategoryId() string`

GetTransactionCategoryId returns the TransactionCategoryId field if non-nil, zero value otherwise.

### GetTransactionCategoryIdOk

`func (o *BankTransactionCreateDto) GetTransactionCategoryIdOk() (*string, bool)`

GetTransactionCategoryIdOk returns a tuple with the TransactionCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryId

`func (o *BankTransactionCreateDto) SetTransactionCategoryId(v string)`

SetTransactionCategoryId sets TransactionCategoryId field to given value.

### HasTransactionCategoryId

`func (o *BankTransactionCreateDto) HasTransactionCategoryId() bool`

HasTransactionCategoryId returns a boolean if a field has been set.

### SetTransactionCategoryIdNil

`func (o *BankTransactionCreateDto) SetTransactionCategoryIdNil(b bool)`

 SetTransactionCategoryIdNil sets the value for TransactionCategoryId to be an explicit nil

### UnsetTransactionCategoryId
`func (o *BankTransactionCreateDto) UnsetTransactionCategoryId()`

UnsetTransactionCategoryId ensures that no value is present for TransactionCategoryId, not even an explicit nil
### GetCurrencyId

`func (o *BankTransactionCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankTransactionCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankTransactionCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankTransactionCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankTransactionCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankTransactionCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetBankProfileId

`func (o *BankTransactionCreateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankTransactionCreateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankTransactionCreateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankTransactionCreateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankTransactionCreateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankTransactionCreateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankTransactionCreateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankTransactionCreateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankTransactionCreateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankTransactionCreateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankTransactionCreateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankTransactionCreateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


