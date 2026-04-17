# TaxRateUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTaxRateUpdateDto

`func NewTaxRateUpdateDto() *TaxRateUpdateDto`

NewTaxRateUpdateDto instantiates a new TaxRateUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateUpdateDtoWithDefaults

`func NewTaxRateUpdateDtoWithDefaults() *TaxRateUpdateDto`

NewTaxRateUpdateDtoWithDefaults instantiates a new TaxRateUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxRateUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRateUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRateUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxRateUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaxRateUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaxRateUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRate

`func (o *TaxRateUpdateDto) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxRateUpdateDto) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxRateUpdateDto) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxRateUpdateDto) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetValue

`func (o *TaxRateUpdateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaxRateUpdateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaxRateUpdateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaxRateUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUm

`func (o *TaxRateUpdateDto) GetUm() string`

GetUm returns the Um field if non-nil, zero value otherwise.

### GetUmOk

`func (o *TaxRateUpdateDto) GetUmOk() (*string, bool)`

GetUmOk returns a tuple with the Um field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUm

`func (o *TaxRateUpdateDto) SetUm(v string)`

SetUm sets Um field to given value.

### HasUm

`func (o *TaxRateUpdateDto) HasUm() bool`

HasUm returns a boolean if a field has been set.

### SetUmNil

`func (o *TaxRateUpdateDto) SetUmNil(b bool)`

 SetUmNil sets the value for Um to be an explicit nil

### UnsetUm
`func (o *TaxRateUpdateDto) UnsetUm()`

UnsetUm ensures that no value is present for Um, not even an explicit nil
### GetUnitId

`func (o *TaxRateUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *TaxRateUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *TaxRateUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *TaxRateUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *TaxRateUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *TaxRateUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *TaxRateUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *TaxRateUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *TaxRateUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *TaxRateUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *TaxRateUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *TaxRateUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriority

`func (o *TaxRateUpdateDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaxRateUpdateDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaxRateUpdateDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaxRateUpdateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCompound

`func (o *TaxRateUpdateDto) GetCompound() bool`

GetCompound returns the Compound field if non-nil, zero value otherwise.

### GetCompoundOk

`func (o *TaxRateUpdateDto) GetCompoundOk() (*bool, bool)`

GetCompoundOk returns a tuple with the Compound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompound

`func (o *TaxRateUpdateDto) SetCompound(v bool)`

SetCompound sets Compound field to given value.

### HasCompound

`func (o *TaxRateUpdateDto) HasCompound() bool`

HasCompound returns a boolean if a field has been set.

### GetShipping

`func (o *TaxRateUpdateDto) GetShipping() bool`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *TaxRateUpdateDto) GetShippingOk() (*bool, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *TaxRateUpdateDto) SetShipping(v bool)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *TaxRateUpdateDto) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetWithholding

`func (o *TaxRateUpdateDto) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *TaxRateUpdateDto) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *TaxRateUpdateDto) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *TaxRateUpdateDto) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.

### GetSingleTransactionThreshold

`func (o *TaxRateUpdateDto) GetSingleTransactionThreshold() float64`

GetSingleTransactionThreshold returns the SingleTransactionThreshold field if non-nil, zero value otherwise.

### GetSingleTransactionThresholdOk

`func (o *TaxRateUpdateDto) GetSingleTransactionThresholdOk() (*float64, bool)`

GetSingleTransactionThresholdOk returns a tuple with the SingleTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTransactionThreshold

`func (o *TaxRateUpdateDto) SetSingleTransactionThreshold(v float64)`

SetSingleTransactionThreshold sets SingleTransactionThreshold field to given value.

### HasSingleTransactionThreshold

`func (o *TaxRateUpdateDto) HasSingleTransactionThreshold() bool`

HasSingleTransactionThreshold returns a boolean if a field has been set.

### GetCumulativeTransactionThreshold

`func (o *TaxRateUpdateDto) GetCumulativeTransactionThreshold() float64`

GetCumulativeTransactionThreshold returns the CumulativeTransactionThreshold field if non-nil, zero value otherwise.

### GetCumulativeTransactionThresholdOk

`func (o *TaxRateUpdateDto) GetCumulativeTransactionThresholdOk() (*float64, bool)`

GetCumulativeTransactionThresholdOk returns a tuple with the CumulativeTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeTransactionThreshold

`func (o *TaxRateUpdateDto) SetCumulativeTransactionThreshold(v float64)`

SetCumulativeTransactionThreshold sets CumulativeTransactionThreshold field to given value.

### HasCumulativeTransactionThreshold

`func (o *TaxRateUpdateDto) HasCumulativeTransactionThreshold() bool`

HasCumulativeTransactionThreshold returns a boolean if a field has been set.

### GetFiscalAuthorityId

`func (o *TaxRateUpdateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *TaxRateUpdateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *TaxRateUpdateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *TaxRateUpdateDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *TaxRateUpdateDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *TaxRateUpdateDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetFiscalYearId

`func (o *TaxRateUpdateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *TaxRateUpdateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *TaxRateUpdateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *TaxRateUpdateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *TaxRateUpdateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *TaxRateUpdateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetCountryId

`func (o *TaxRateUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TaxRateUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TaxRateUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TaxRateUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TaxRateUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TaxRateUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTaxClassId

`func (o *TaxRateUpdateDto) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *TaxRateUpdateDto) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *TaxRateUpdateDto) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *TaxRateUpdateDto) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### SetTaxClassIdNil

`func (o *TaxRateUpdateDto) SetTaxClassIdNil(b bool)`

 SetTaxClassIdNil sets the value for TaxClassId to be an explicit nil

### UnsetTaxClassId
`func (o *TaxRateUpdateDto) UnsetTaxClassId()`

UnsetTaxClassId ensures that no value is present for TaxClassId, not even an explicit nil
### GetCurrencyId

`func (o *TaxRateUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TaxRateUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TaxRateUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TaxRateUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TaxRateUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TaxRateUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTaxPolicyId

`func (o *TaxRateUpdateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *TaxRateUpdateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *TaxRateUpdateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *TaxRateUpdateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *TaxRateUpdateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *TaxRateUpdateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


