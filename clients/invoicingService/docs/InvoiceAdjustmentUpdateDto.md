# InvoiceAdjustmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SurchargePercent** | Pointer to **float64** |  | [optional] 
**SurchargeAmount** | Pointer to **float64** |  | [optional] 
**DiscountPercent** | Pointer to **float64** |  | [optional] 
**DiscountAmount** | Pointer to **float64** |  | [optional] 
**TotalSurcharge** | Pointer to **float64** |  | [optional] 
**TotalDiscount** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewInvoiceAdjustmentUpdateDto

`func NewInvoiceAdjustmentUpdateDto() *InvoiceAdjustmentUpdateDto`

NewInvoiceAdjustmentUpdateDto instantiates a new InvoiceAdjustmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAdjustmentUpdateDtoWithDefaults

`func NewInvoiceAdjustmentUpdateDtoWithDefaults() *InvoiceAdjustmentUpdateDto`

NewInvoiceAdjustmentUpdateDtoWithDefaults instantiates a new InvoiceAdjustmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *InvoiceAdjustmentUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceAdjustmentUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceAdjustmentUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceAdjustmentUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceAdjustmentUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceAdjustmentUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *InvoiceAdjustmentUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceAdjustmentUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceAdjustmentUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceAdjustmentUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceAdjustmentUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceAdjustmentUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSurchargePercent

`func (o *InvoiceAdjustmentUpdateDto) GetSurchargePercent() float64`

GetSurchargePercent returns the SurchargePercent field if non-nil, zero value otherwise.

### GetSurchargePercentOk

`func (o *InvoiceAdjustmentUpdateDto) GetSurchargePercentOk() (*float64, bool)`

GetSurchargePercentOk returns a tuple with the SurchargePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargePercent

`func (o *InvoiceAdjustmentUpdateDto) SetSurchargePercent(v float64)`

SetSurchargePercent sets SurchargePercent field to given value.

### HasSurchargePercent

`func (o *InvoiceAdjustmentUpdateDto) HasSurchargePercent() bool`

HasSurchargePercent returns a boolean if a field has been set.

### GetSurchargeAmount

`func (o *InvoiceAdjustmentUpdateDto) GetSurchargeAmount() float64`

GetSurchargeAmount returns the SurchargeAmount field if non-nil, zero value otherwise.

### GetSurchargeAmountOk

`func (o *InvoiceAdjustmentUpdateDto) GetSurchargeAmountOk() (*float64, bool)`

GetSurchargeAmountOk returns a tuple with the SurchargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargeAmount

`func (o *InvoiceAdjustmentUpdateDto) SetSurchargeAmount(v float64)`

SetSurchargeAmount sets SurchargeAmount field to given value.

### HasSurchargeAmount

`func (o *InvoiceAdjustmentUpdateDto) HasSurchargeAmount() bool`

HasSurchargeAmount returns a boolean if a field has been set.

### GetDiscountPercent

`func (o *InvoiceAdjustmentUpdateDto) GetDiscountPercent() float64`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *InvoiceAdjustmentUpdateDto) GetDiscountPercentOk() (*float64, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *InvoiceAdjustmentUpdateDto) SetDiscountPercent(v float64)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *InvoiceAdjustmentUpdateDto) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *InvoiceAdjustmentUpdateDto) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceAdjustmentUpdateDto) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceAdjustmentUpdateDto) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *InvoiceAdjustmentUpdateDto) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTotalSurcharge

`func (o *InvoiceAdjustmentUpdateDto) GetTotalSurcharge() float64`

GetTotalSurcharge returns the TotalSurcharge field if non-nil, zero value otherwise.

### GetTotalSurchargeOk

`func (o *InvoiceAdjustmentUpdateDto) GetTotalSurchargeOk() (*float64, bool)`

GetTotalSurchargeOk returns a tuple with the TotalSurcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharge

`func (o *InvoiceAdjustmentUpdateDto) SetTotalSurcharge(v float64)`

SetTotalSurcharge sets TotalSurcharge field to given value.

### HasTotalSurcharge

`func (o *InvoiceAdjustmentUpdateDto) HasTotalSurcharge() bool`

HasTotalSurcharge returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *InvoiceAdjustmentUpdateDto) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *InvoiceAdjustmentUpdateDto) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *InvoiceAdjustmentUpdateDto) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *InvoiceAdjustmentUpdateDto) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetType

`func (o *InvoiceAdjustmentUpdateDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceAdjustmentUpdateDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceAdjustmentUpdateDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceAdjustmentUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


