# TaxRateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Um** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Compound** | Pointer to **bool** |  | [optional] 
**Shipping** | Pointer to **bool** |  | [optional] 
**Withholding** | Pointer to **bool** |  | [optional] 
**SingleTransactionThreshold** | Pointer to **float64** |  | [optional] 
**CumulativeTransactionThreshold** | Pointer to **float64** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**TaxClassId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaxRateCreateDto

`func NewTaxRateCreateDto() *TaxRateCreateDto`

NewTaxRateCreateDto instantiates a new TaxRateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateCreateDtoWithDefaults

`func NewTaxRateCreateDtoWithDefaults() *TaxRateCreateDto`

NewTaxRateCreateDtoWithDefaults instantiates a new TaxRateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxRateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxRateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TaxRateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TaxRateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TaxRateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TaxRateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TaxRateCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRateCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRateCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxRateCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaxRateCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaxRateCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRate

`func (o *TaxRateCreateDto) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxRateCreateDto) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxRateCreateDto) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxRateCreateDto) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetValue

`func (o *TaxRateCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaxRateCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaxRateCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaxRateCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUm

`func (o *TaxRateCreateDto) GetUm() string`

GetUm returns the Um field if non-nil, zero value otherwise.

### GetUmOk

`func (o *TaxRateCreateDto) GetUmOk() (*string, bool)`

GetUmOk returns a tuple with the Um field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUm

`func (o *TaxRateCreateDto) SetUm(v string)`

SetUm sets Um field to given value.

### HasUm

`func (o *TaxRateCreateDto) HasUm() bool`

HasUm returns a boolean if a field has been set.

### SetUmNil

`func (o *TaxRateCreateDto) SetUmNil(b bool)`

 SetUmNil sets the value for Um to be an explicit nil

### UnsetUm
`func (o *TaxRateCreateDto) UnsetUm()`

UnsetUm ensures that no value is present for Um, not even an explicit nil
### GetUnitId

`func (o *TaxRateCreateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *TaxRateCreateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *TaxRateCreateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *TaxRateCreateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *TaxRateCreateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *TaxRateCreateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *TaxRateCreateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *TaxRateCreateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *TaxRateCreateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *TaxRateCreateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *TaxRateCreateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *TaxRateCreateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriority

`func (o *TaxRateCreateDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaxRateCreateDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaxRateCreateDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaxRateCreateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCompound

`func (o *TaxRateCreateDto) GetCompound() bool`

GetCompound returns the Compound field if non-nil, zero value otherwise.

### GetCompoundOk

`func (o *TaxRateCreateDto) GetCompoundOk() (*bool, bool)`

GetCompoundOk returns a tuple with the Compound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompound

`func (o *TaxRateCreateDto) SetCompound(v bool)`

SetCompound sets Compound field to given value.

### HasCompound

`func (o *TaxRateCreateDto) HasCompound() bool`

HasCompound returns a boolean if a field has been set.

### GetShipping

`func (o *TaxRateCreateDto) GetShipping() bool`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *TaxRateCreateDto) GetShippingOk() (*bool, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *TaxRateCreateDto) SetShipping(v bool)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *TaxRateCreateDto) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetWithholding

`func (o *TaxRateCreateDto) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *TaxRateCreateDto) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *TaxRateCreateDto) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *TaxRateCreateDto) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.

### GetSingleTransactionThreshold

`func (o *TaxRateCreateDto) GetSingleTransactionThreshold() float64`

GetSingleTransactionThreshold returns the SingleTransactionThreshold field if non-nil, zero value otherwise.

### GetSingleTransactionThresholdOk

`func (o *TaxRateCreateDto) GetSingleTransactionThresholdOk() (*float64, bool)`

GetSingleTransactionThresholdOk returns a tuple with the SingleTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTransactionThreshold

`func (o *TaxRateCreateDto) SetSingleTransactionThreshold(v float64)`

SetSingleTransactionThreshold sets SingleTransactionThreshold field to given value.

### HasSingleTransactionThreshold

`func (o *TaxRateCreateDto) HasSingleTransactionThreshold() bool`

HasSingleTransactionThreshold returns a boolean if a field has been set.

### GetCumulativeTransactionThreshold

`func (o *TaxRateCreateDto) GetCumulativeTransactionThreshold() float64`

GetCumulativeTransactionThreshold returns the CumulativeTransactionThreshold field if non-nil, zero value otherwise.

### GetCumulativeTransactionThresholdOk

`func (o *TaxRateCreateDto) GetCumulativeTransactionThresholdOk() (*float64, bool)`

GetCumulativeTransactionThresholdOk returns a tuple with the CumulativeTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeTransactionThreshold

`func (o *TaxRateCreateDto) SetCumulativeTransactionThreshold(v float64)`

SetCumulativeTransactionThreshold sets CumulativeTransactionThreshold field to given value.

### HasCumulativeTransactionThreshold

`func (o *TaxRateCreateDto) HasCumulativeTransactionThreshold() bool`

HasCumulativeTransactionThreshold returns a boolean if a field has been set.

### GetFiscalAuthorityId

`func (o *TaxRateCreateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *TaxRateCreateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *TaxRateCreateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *TaxRateCreateDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *TaxRateCreateDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *TaxRateCreateDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetFiscalYearId

`func (o *TaxRateCreateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *TaxRateCreateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *TaxRateCreateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *TaxRateCreateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *TaxRateCreateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *TaxRateCreateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetCountryId

`func (o *TaxRateCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TaxRateCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TaxRateCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TaxRateCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TaxRateCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TaxRateCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTaxClassId

`func (o *TaxRateCreateDto) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *TaxRateCreateDto) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *TaxRateCreateDto) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *TaxRateCreateDto) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### SetTaxClassIdNil

`func (o *TaxRateCreateDto) SetTaxClassIdNil(b bool)`

 SetTaxClassIdNil sets the value for TaxClassId to be an explicit nil

### UnsetTaxClassId
`func (o *TaxRateCreateDto) UnsetTaxClassId()`

UnsetTaxClassId ensures that no value is present for TaxClassId, not even an explicit nil
### GetCurrencyId

`func (o *TaxRateCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TaxRateCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TaxRateCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TaxRateCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TaxRateCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TaxRateCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTaxPolicyId

`func (o *TaxRateCreateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *TaxRateCreateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *TaxRateCreateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *TaxRateCreateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *TaxRateCreateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *TaxRateCreateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


