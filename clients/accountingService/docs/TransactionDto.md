# TransactionDto

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

## Methods

### NewTransactionDto

`func NewTransactionDto() *TransactionDto`

NewTransactionDto instantiates a new TransactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDtoWithDefaults

`func NewTransactionDtoWithDefaults() *TransactionDto`

NewTransactionDtoWithDefaults instantiates a new TransactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransactionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransactionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TransactionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransactionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransactionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TransactionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TransactionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TransactionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDescription

`func (o *TransactionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *TransactionDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TransactionDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TransactionDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TransactionDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *TransactionDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TransactionDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TransactionDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TransactionDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExternalDescription

`func (o *TransactionDto) GetExternalDescription() string`

GetExternalDescription returns the ExternalDescription field if non-nil, zero value otherwise.

### GetExternalDescriptionOk

`func (o *TransactionDto) GetExternalDescriptionOk() (*string, bool)`

GetExternalDescriptionOk returns a tuple with the ExternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDescription

`func (o *TransactionDto) SetExternalDescription(v string)`

SetExternalDescription sets ExternalDescription field to given value.

### HasExternalDescription

`func (o *TransactionDto) HasExternalDescription() bool`

HasExternalDescription returns a boolean if a field has been set.

### SetExternalDescriptionNil

`func (o *TransactionDto) SetExternalDescriptionNil(b bool)`

 SetExternalDescriptionNil sets the value for ExternalDescription to be an explicit nil

### UnsetExternalDescription
`func (o *TransactionDto) UnsetExternalDescription()`

UnsetExternalDescription ensures that no value is present for ExternalDescription, not even an explicit nil
### GetBasisQuantity

`func (o *TransactionDto) GetBasisQuantity() float64`

GetBasisQuantity returns the BasisQuantity field if non-nil, zero value otherwise.

### GetBasisQuantityOk

`func (o *TransactionDto) GetBasisQuantityOk() (*float64, bool)`

GetBasisQuantityOk returns a tuple with the BasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisQuantity

`func (o *TransactionDto) SetBasisQuantity(v float64)`

SetBasisQuantity sets BasisQuantity field to given value.

### HasBasisQuantity

`func (o *TransactionDto) HasBasisQuantity() bool`

HasBasisQuantity returns a boolean if a field has been set.

### GetBasisAmount

`func (o *TransactionDto) GetBasisAmount() float64`

GetBasisAmount returns the BasisAmount field if non-nil, zero value otherwise.

### GetBasisAmountOk

`func (o *TransactionDto) GetBasisAmountOk() (*float64, bool)`

GetBasisAmountOk returns a tuple with the BasisAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisAmount

`func (o *TransactionDto) SetBasisAmount(v float64)`

SetBasisAmount sets BasisAmount field to given value.

### HasBasisAmount

`func (o *TransactionDto) HasBasisAmount() bool`

HasBasisAmount returns a boolean if a field has been set.

### GetPercent

`func (o *TransactionDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *TransactionDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *TransactionDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *TransactionDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUnitGroupId

`func (o *TransactionDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *TransactionDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *TransactionDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *TransactionDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *TransactionDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *TransactionDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetUnitId

`func (o *TransactionDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *TransactionDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *TransactionDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *TransactionDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *TransactionDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *TransactionDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTransactionCategoryId

`func (o *TransactionDto) GetTransactionCategoryId() string`

GetTransactionCategoryId returns the TransactionCategoryId field if non-nil, zero value otherwise.

### GetTransactionCategoryIdOk

`func (o *TransactionDto) GetTransactionCategoryIdOk() (*string, bool)`

GetTransactionCategoryIdOk returns a tuple with the TransactionCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryId

`func (o *TransactionDto) SetTransactionCategoryId(v string)`

SetTransactionCategoryId sets TransactionCategoryId field to given value.

### HasTransactionCategoryId

`func (o *TransactionDto) HasTransactionCategoryId() bool`

HasTransactionCategoryId returns a boolean if a field has been set.

### SetTransactionCategoryIdNil

`func (o *TransactionDto) SetTransactionCategoryIdNil(b bool)`

 SetTransactionCategoryIdNil sets the value for TransactionCategoryId to be an explicit nil

### UnsetTransactionCategoryId
`func (o *TransactionDto) UnsetTransactionCategoryId()`

UnsetTransactionCategoryId ensures that no value is present for TransactionCategoryId, not even an explicit nil
### GetCurrencyId

`func (o *TransactionDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TransactionDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TransactionDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TransactionDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TransactionDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TransactionDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *TransactionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TransactionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TransactionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TransactionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TransactionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TransactionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TransactionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TransactionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TransactionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TransactionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TransactionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TransactionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


