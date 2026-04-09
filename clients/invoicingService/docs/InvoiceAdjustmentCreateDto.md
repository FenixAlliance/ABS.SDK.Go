# InvoiceAdjustmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SurchargePercent** | Pointer to **float64** |  | [optional] 
**SurchargeAmount** | Pointer to **float64** |  | [optional] 
**DiscountPercent** | Pointer to **float64** |  | [optional] 
**DiscountAmount** | Pointer to **float64** |  | [optional] 
**TotalSurcharge** | Pointer to **float64** |  | [optional] 
**TotalDiscount** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceAdjustmentCreateDto

`func NewInvoiceAdjustmentCreateDto() *InvoiceAdjustmentCreateDto`

NewInvoiceAdjustmentCreateDto instantiates a new InvoiceAdjustmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAdjustmentCreateDtoWithDefaults

`func NewInvoiceAdjustmentCreateDtoWithDefaults() *InvoiceAdjustmentCreateDto`

NewInvoiceAdjustmentCreateDtoWithDefaults instantiates a new InvoiceAdjustmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceAdjustmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceAdjustmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceAdjustmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceAdjustmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceAdjustmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceAdjustmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceAdjustmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceAdjustmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCurrencyId

`func (o *InvoiceAdjustmentCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceAdjustmentCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceAdjustmentCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceAdjustmentCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceAdjustmentCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceAdjustmentCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *InvoiceAdjustmentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceAdjustmentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceAdjustmentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceAdjustmentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceAdjustmentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceAdjustmentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSurchargePercent

`func (o *InvoiceAdjustmentCreateDto) GetSurchargePercent() float64`

GetSurchargePercent returns the SurchargePercent field if non-nil, zero value otherwise.

### GetSurchargePercentOk

`func (o *InvoiceAdjustmentCreateDto) GetSurchargePercentOk() (*float64, bool)`

GetSurchargePercentOk returns a tuple with the SurchargePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargePercent

`func (o *InvoiceAdjustmentCreateDto) SetSurchargePercent(v float64)`

SetSurchargePercent sets SurchargePercent field to given value.

### HasSurchargePercent

`func (o *InvoiceAdjustmentCreateDto) HasSurchargePercent() bool`

HasSurchargePercent returns a boolean if a field has been set.

### GetSurchargeAmount

`func (o *InvoiceAdjustmentCreateDto) GetSurchargeAmount() float64`

GetSurchargeAmount returns the SurchargeAmount field if non-nil, zero value otherwise.

### GetSurchargeAmountOk

`func (o *InvoiceAdjustmentCreateDto) GetSurchargeAmountOk() (*float64, bool)`

GetSurchargeAmountOk returns a tuple with the SurchargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargeAmount

`func (o *InvoiceAdjustmentCreateDto) SetSurchargeAmount(v float64)`

SetSurchargeAmount sets SurchargeAmount field to given value.

### HasSurchargeAmount

`func (o *InvoiceAdjustmentCreateDto) HasSurchargeAmount() bool`

HasSurchargeAmount returns a boolean if a field has been set.

### GetDiscountPercent

`func (o *InvoiceAdjustmentCreateDto) GetDiscountPercent() float64`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *InvoiceAdjustmentCreateDto) GetDiscountPercentOk() (*float64, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *InvoiceAdjustmentCreateDto) SetDiscountPercent(v float64)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *InvoiceAdjustmentCreateDto) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *InvoiceAdjustmentCreateDto) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceAdjustmentCreateDto) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceAdjustmentCreateDto) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *InvoiceAdjustmentCreateDto) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTotalSurcharge

`func (o *InvoiceAdjustmentCreateDto) GetTotalSurcharge() float64`

GetTotalSurcharge returns the TotalSurcharge field if non-nil, zero value otherwise.

### GetTotalSurchargeOk

`func (o *InvoiceAdjustmentCreateDto) GetTotalSurchargeOk() (*float64, bool)`

GetTotalSurchargeOk returns a tuple with the TotalSurcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharge

`func (o *InvoiceAdjustmentCreateDto) SetTotalSurcharge(v float64)`

SetTotalSurcharge sets TotalSurcharge field to given value.

### HasTotalSurcharge

`func (o *InvoiceAdjustmentCreateDto) HasTotalSurcharge() bool`

HasTotalSurcharge returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *InvoiceAdjustmentCreateDto) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *InvoiceAdjustmentCreateDto) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *InvoiceAdjustmentCreateDto) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *InvoiceAdjustmentCreateDto) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetType

`func (o *InvoiceAdjustmentCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceAdjustmentCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceAdjustmentCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceAdjustmentCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


