# InvoiceAdjustmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SurchargePercent** | Pointer to **float64** |  | [optional] 
**SurchargeAmount** | Pointer to **float64** |  | [optional] 
**DiscountPercent** | Pointer to **float64** |  | [optional] 
**DiscountAmount** | Pointer to **float64** |  | [optional] 
**TotalSurcharge** | Pointer to **float64** |  | [optional] 
**TotalDiscount** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewInvoiceAdjustmentDto

`func NewInvoiceAdjustmentDto() *InvoiceAdjustmentDto`

NewInvoiceAdjustmentDto instantiates a new InvoiceAdjustmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAdjustmentDtoWithDefaults

`func NewInvoiceAdjustmentDtoWithDefaults() *InvoiceAdjustmentDto`

NewInvoiceAdjustmentDtoWithDefaults instantiates a new InvoiceAdjustmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceAdjustmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceAdjustmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceAdjustmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceAdjustmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceAdjustmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceAdjustmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InvoiceAdjustmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceAdjustmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceAdjustmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceAdjustmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InvoiceAdjustmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InvoiceAdjustmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *InvoiceAdjustmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceAdjustmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceAdjustmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceAdjustmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceAdjustmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceAdjustmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetInvoiceId

`func (o *InvoiceAdjustmentDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceAdjustmentDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceAdjustmentDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceAdjustmentDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *InvoiceAdjustmentDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceAdjustmentDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetCurrencyId

`func (o *InvoiceAdjustmentDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceAdjustmentDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceAdjustmentDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceAdjustmentDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceAdjustmentDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceAdjustmentDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceAdjustmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceAdjustmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceAdjustmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceAdjustmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceAdjustmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceAdjustmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDescription

`func (o *InvoiceAdjustmentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceAdjustmentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceAdjustmentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceAdjustmentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceAdjustmentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceAdjustmentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSurchargePercent

`func (o *InvoiceAdjustmentDto) GetSurchargePercent() float64`

GetSurchargePercent returns the SurchargePercent field if non-nil, zero value otherwise.

### GetSurchargePercentOk

`func (o *InvoiceAdjustmentDto) GetSurchargePercentOk() (*float64, bool)`

GetSurchargePercentOk returns a tuple with the SurchargePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargePercent

`func (o *InvoiceAdjustmentDto) SetSurchargePercent(v float64)`

SetSurchargePercent sets SurchargePercent field to given value.

### HasSurchargePercent

`func (o *InvoiceAdjustmentDto) HasSurchargePercent() bool`

HasSurchargePercent returns a boolean if a field has been set.

### GetSurchargeAmount

`func (o *InvoiceAdjustmentDto) GetSurchargeAmount() float64`

GetSurchargeAmount returns the SurchargeAmount field if non-nil, zero value otherwise.

### GetSurchargeAmountOk

`func (o *InvoiceAdjustmentDto) GetSurchargeAmountOk() (*float64, bool)`

GetSurchargeAmountOk returns a tuple with the SurchargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargeAmount

`func (o *InvoiceAdjustmentDto) SetSurchargeAmount(v float64)`

SetSurchargeAmount sets SurchargeAmount field to given value.

### HasSurchargeAmount

`func (o *InvoiceAdjustmentDto) HasSurchargeAmount() bool`

HasSurchargeAmount returns a boolean if a field has been set.

### GetDiscountPercent

`func (o *InvoiceAdjustmentDto) GetDiscountPercent() float64`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *InvoiceAdjustmentDto) GetDiscountPercentOk() (*float64, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *InvoiceAdjustmentDto) SetDiscountPercent(v float64)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *InvoiceAdjustmentDto) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *InvoiceAdjustmentDto) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceAdjustmentDto) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceAdjustmentDto) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *InvoiceAdjustmentDto) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTotalSurcharge

`func (o *InvoiceAdjustmentDto) GetTotalSurcharge() float64`

GetTotalSurcharge returns the TotalSurcharge field if non-nil, zero value otherwise.

### GetTotalSurchargeOk

`func (o *InvoiceAdjustmentDto) GetTotalSurchargeOk() (*float64, bool)`

GetTotalSurchargeOk returns a tuple with the TotalSurcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharge

`func (o *InvoiceAdjustmentDto) SetTotalSurcharge(v float64)`

SetTotalSurcharge sets TotalSurcharge field to given value.

### HasTotalSurcharge

`func (o *InvoiceAdjustmentDto) HasTotalSurcharge() bool`

HasTotalSurcharge returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *InvoiceAdjustmentDto) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *InvoiceAdjustmentDto) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *InvoiceAdjustmentDto) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *InvoiceAdjustmentDto) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetType

`func (o *InvoiceAdjustmentDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceAdjustmentDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceAdjustmentDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceAdjustmentDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


