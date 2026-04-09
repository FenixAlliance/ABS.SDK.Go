# TaxRateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**TaxClassId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaxRateDto

`func NewTaxRateDto() *TaxRateDto`

NewTaxRateDto instantiates a new TaxRateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateDtoWithDefaults

`func NewTaxRateDtoWithDefaults() *TaxRateDto`

NewTaxRateDtoWithDefaults instantiates a new TaxRateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxRateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxRateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaxRateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxRateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TaxRateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TaxRateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TaxRateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TaxRateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TaxRateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TaxRateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *TaxRateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxRateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaxRateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaxRateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRate

`func (o *TaxRateDto) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxRateDto) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxRateDto) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxRateDto) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetValue

`func (o *TaxRateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaxRateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaxRateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaxRateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUm

`func (o *TaxRateDto) GetUm() string`

GetUm returns the Um field if non-nil, zero value otherwise.

### GetUmOk

`func (o *TaxRateDto) GetUmOk() (*string, bool)`

GetUmOk returns a tuple with the Um field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUm

`func (o *TaxRateDto) SetUm(v string)`

SetUm sets Um field to given value.

### HasUm

`func (o *TaxRateDto) HasUm() bool`

HasUm returns a boolean if a field has been set.

### SetUmNil

`func (o *TaxRateDto) SetUmNil(b bool)`

 SetUmNil sets the value for Um to be an explicit nil

### UnsetUm
`func (o *TaxRateDto) UnsetUm()`

UnsetUm ensures that no value is present for Um, not even an explicit nil
### GetUnitId

`func (o *TaxRateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *TaxRateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *TaxRateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *TaxRateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *TaxRateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *TaxRateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *TaxRateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *TaxRateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *TaxRateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *TaxRateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *TaxRateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *TaxRateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetPriority

`func (o *TaxRateDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaxRateDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaxRateDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaxRateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCompound

`func (o *TaxRateDto) GetCompound() bool`

GetCompound returns the Compound field if non-nil, zero value otherwise.

### GetCompoundOk

`func (o *TaxRateDto) GetCompoundOk() (*bool, bool)`

GetCompoundOk returns a tuple with the Compound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompound

`func (o *TaxRateDto) SetCompound(v bool)`

SetCompound sets Compound field to given value.

### HasCompound

`func (o *TaxRateDto) HasCompound() bool`

HasCompound returns a boolean if a field has been set.

### GetShipping

`func (o *TaxRateDto) GetShipping() bool`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *TaxRateDto) GetShippingOk() (*bool, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *TaxRateDto) SetShipping(v bool)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *TaxRateDto) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetWithholding

`func (o *TaxRateDto) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *TaxRateDto) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *TaxRateDto) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *TaxRateDto) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.

### GetSingleTransactionThreshold

`func (o *TaxRateDto) GetSingleTransactionThreshold() float64`

GetSingleTransactionThreshold returns the SingleTransactionThreshold field if non-nil, zero value otherwise.

### GetSingleTransactionThresholdOk

`func (o *TaxRateDto) GetSingleTransactionThresholdOk() (*float64, bool)`

GetSingleTransactionThresholdOk returns a tuple with the SingleTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTransactionThreshold

`func (o *TaxRateDto) SetSingleTransactionThreshold(v float64)`

SetSingleTransactionThreshold sets SingleTransactionThreshold field to given value.

### HasSingleTransactionThreshold

`func (o *TaxRateDto) HasSingleTransactionThreshold() bool`

HasSingleTransactionThreshold returns a boolean if a field has been set.

### GetCumulativeTransactionThreshold

`func (o *TaxRateDto) GetCumulativeTransactionThreshold() float64`

GetCumulativeTransactionThreshold returns the CumulativeTransactionThreshold field if non-nil, zero value otherwise.

### GetCumulativeTransactionThresholdOk

`func (o *TaxRateDto) GetCumulativeTransactionThresholdOk() (*float64, bool)`

GetCumulativeTransactionThresholdOk returns a tuple with the CumulativeTransactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeTransactionThreshold

`func (o *TaxRateDto) SetCumulativeTransactionThreshold(v float64)`

SetCumulativeTransactionThreshold sets CumulativeTransactionThreshold field to given value.

### HasCumulativeTransactionThreshold

`func (o *TaxRateDto) HasCumulativeTransactionThreshold() bool`

HasCumulativeTransactionThreshold returns a boolean if a field has been set.

### GetFiscalAuthorityId

`func (o *TaxRateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *TaxRateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *TaxRateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *TaxRateDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *TaxRateDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *TaxRateDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetFiscalYearId

`func (o *TaxRateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *TaxRateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *TaxRateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *TaxRateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *TaxRateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *TaxRateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetTenantId

`func (o *TaxRateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TaxRateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TaxRateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TaxRateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TaxRateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TaxRateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCountryId

`func (o *TaxRateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TaxRateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TaxRateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TaxRateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TaxRateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TaxRateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTaxClassId

`func (o *TaxRateDto) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *TaxRateDto) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *TaxRateDto) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *TaxRateDto) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### SetTaxClassIdNil

`func (o *TaxRateDto) SetTaxClassIdNil(b bool)`

 SetTaxClassIdNil sets the value for TaxClassId to be an explicit nil

### UnsetTaxClassId
`func (o *TaxRateDto) UnsetTaxClassId()`

UnsetTaxClassId ensures that no value is present for TaxClassId, not even an explicit nil
### GetCurrencyId

`func (o *TaxRateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TaxRateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TaxRateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TaxRateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TaxRateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TaxRateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTaxPolicyId

`func (o *TaxRateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *TaxRateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *TaxRateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *TaxRateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *TaxRateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *TaxRateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetEnrollmentId

`func (o *TaxRateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TaxRateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TaxRateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TaxRateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TaxRateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TaxRateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


