# BankTransactionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**BankProfileId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankTransactionDto

`func NewBankTransactionDto() *BankTransactionDto`

NewBankTransactionDto instantiates a new BankTransactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionDtoWithDefaults

`func NewBankTransactionDtoWithDefaults() *BankTransactionDto`

NewBankTransactionDtoWithDefaults instantiates a new BankTransactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransactionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransactionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransactionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankTransactionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BankTransactionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BankTransactionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BankTransactionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankTransactionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankTransactionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankTransactionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BankTransactionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BankTransactionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDescription

`func (o *BankTransactionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransactionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransactionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTransactionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BankTransactionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BankTransactionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *BankTransactionDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BankTransactionDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BankTransactionDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BankTransactionDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *BankTransactionDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BankTransactionDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BankTransactionDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BankTransactionDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExternalDescription

`func (o *BankTransactionDto) GetExternalDescription() string`

GetExternalDescription returns the ExternalDescription field if non-nil, zero value otherwise.

### GetExternalDescriptionOk

`func (o *BankTransactionDto) GetExternalDescriptionOk() (*string, bool)`

GetExternalDescriptionOk returns a tuple with the ExternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDescription

`func (o *BankTransactionDto) SetExternalDescription(v string)`

SetExternalDescription sets ExternalDescription field to given value.

### HasExternalDescription

`func (o *BankTransactionDto) HasExternalDescription() bool`

HasExternalDescription returns a boolean if a field has been set.

### SetExternalDescriptionNil

`func (o *BankTransactionDto) SetExternalDescriptionNil(b bool)`

 SetExternalDescriptionNil sets the value for ExternalDescription to be an explicit nil

### UnsetExternalDescription
`func (o *BankTransactionDto) UnsetExternalDescription()`

UnsetExternalDescription ensures that no value is present for ExternalDescription, not even an explicit nil
### GetBasisQuantity

`func (o *BankTransactionDto) GetBasisQuantity() float64`

GetBasisQuantity returns the BasisQuantity field if non-nil, zero value otherwise.

### GetBasisQuantityOk

`func (o *BankTransactionDto) GetBasisQuantityOk() (*float64, bool)`

GetBasisQuantityOk returns a tuple with the BasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisQuantity

`func (o *BankTransactionDto) SetBasisQuantity(v float64)`

SetBasisQuantity sets BasisQuantity field to given value.

### HasBasisQuantity

`func (o *BankTransactionDto) HasBasisQuantity() bool`

HasBasisQuantity returns a boolean if a field has been set.

### GetBasisAmount

`func (o *BankTransactionDto) GetBasisAmount() float64`

GetBasisAmount returns the BasisAmount field if non-nil, zero value otherwise.

### GetBasisAmountOk

`func (o *BankTransactionDto) GetBasisAmountOk() (*float64, bool)`

GetBasisAmountOk returns a tuple with the BasisAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisAmount

`func (o *BankTransactionDto) SetBasisAmount(v float64)`

SetBasisAmount sets BasisAmount field to given value.

### HasBasisAmount

`func (o *BankTransactionDto) HasBasisAmount() bool`

HasBasisAmount returns a boolean if a field has been set.

### GetPercent

`func (o *BankTransactionDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *BankTransactionDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *BankTransactionDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *BankTransactionDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *BankTransactionDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *BankTransactionDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *BankTransactionDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *BankTransactionDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *BankTransactionDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *BankTransactionDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetUnitId

`func (o *BankTransactionDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *BankTransactionDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *BankTransactionDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *BankTransactionDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *BankTransactionDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *BankTransactionDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTransactionCategoryId

`func (o *BankTransactionDto) GetTransactionCategoryId() string`

GetTransactionCategoryId returns the TransactionCategoryId field if non-nil, zero value otherwise.

### GetTransactionCategoryIdOk

`func (o *BankTransactionDto) GetTransactionCategoryIdOk() (*string, bool)`

GetTransactionCategoryIdOk returns a tuple with the TransactionCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryId

`func (o *BankTransactionDto) SetTransactionCategoryId(v string)`

SetTransactionCategoryId sets TransactionCategoryId field to given value.

### HasTransactionCategoryId

`func (o *BankTransactionDto) HasTransactionCategoryId() bool`

HasTransactionCategoryId returns a boolean if a field has been set.

### SetTransactionCategoryIdNil

`func (o *BankTransactionDto) SetTransactionCategoryIdNil(b bool)`

 SetTransactionCategoryIdNil sets the value for TransactionCategoryId to be an explicit nil

### UnsetTransactionCategoryId
`func (o *BankTransactionDto) UnsetTransactionCategoryId()`

UnsetTransactionCategoryId ensures that no value is present for TransactionCategoryId, not even an explicit nil
### GetCurrencyId

`func (o *BankTransactionDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankTransactionDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankTransactionDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankTransactionDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankTransactionDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankTransactionDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *BankTransactionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankTransactionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankTransactionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankTransactionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankTransactionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankTransactionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *BankTransactionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankTransactionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankTransactionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankTransactionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankTransactionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankTransactionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetBankProfileId

`func (o *BankTransactionDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankTransactionDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankTransactionDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankTransactionDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankTransactionDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankTransactionDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankTransactionDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankTransactionDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankTransactionDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankTransactionDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankTransactionDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankTransactionDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


