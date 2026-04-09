# TaxPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsFree** | Pointer to **bool** |  | [optional] 
**Reduce** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**AllowInternational** | Pointer to **bool** |  | [optional] 
**Hours** | Pointer to **int32** |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**Weeks** | Pointer to **int32** |  | [optional] 
**Months** | Pointer to **int32** |  | [optional] 
**Years** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Percentage** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CountryStateId** | Pointer to **NullableString** |  | [optional] 
**CustomState** | Pointer to **NullableString** |  | [optional] 
**CustomCity** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Zero** | Pointer to **bool** |  | [optional] 
**Reduced** | Pointer to **bool** |  | [optional] 
**Withholding** | Pointer to **bool** |  | [optional] 
**FiscalAuthorityID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaxPolicyDto

`func NewTaxPolicyDto() *TaxPolicyDto`

NewTaxPolicyDto instantiates a new TaxPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxPolicyDtoWithDefaults

`func NewTaxPolicyDtoWithDefaults() *TaxPolicyDto`

NewTaxPolicyDtoWithDefaults instantiates a new TaxPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxPolicyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxPolicyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxPolicyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxPolicyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaxPolicyDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxPolicyDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TaxPolicyDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TaxPolicyDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TaxPolicyDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TaxPolicyDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TaxPolicyDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TaxPolicyDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCode

`func (o *TaxPolicyDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxPolicyDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxPolicyDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxPolicyDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TaxPolicyDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TaxPolicyDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTitle

`func (o *TaxPolicyDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaxPolicyDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaxPolicyDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaxPolicyDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TaxPolicyDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TaxPolicyDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TaxPolicyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxPolicyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxPolicyDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxPolicyDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaxPolicyDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaxPolicyDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *TaxPolicyDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *TaxPolicyDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *TaxPolicyDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *TaxPolicyDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *TaxPolicyDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *TaxPolicyDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *TaxPolicyDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *TaxPolicyDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *TaxPolicyDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *TaxPolicyDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *TaxPolicyDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *TaxPolicyDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *TaxPolicyDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TaxPolicyDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TaxPolicyDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TaxPolicyDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *TaxPolicyDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *TaxPolicyDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *TaxPolicyDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *TaxPolicyDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *TaxPolicyDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TaxPolicyDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TaxPolicyDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TaxPolicyDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *TaxPolicyDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TaxPolicyDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TaxPolicyDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *TaxPolicyDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *TaxPolicyDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *TaxPolicyDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *TaxPolicyDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *TaxPolicyDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *TaxPolicyDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *TaxPolicyDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *TaxPolicyDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *TaxPolicyDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *TaxPolicyDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *TaxPolicyDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *TaxPolicyDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *TaxPolicyDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *TaxPolicyDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaxPolicyDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaxPolicyDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaxPolicyDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *TaxPolicyDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *TaxPolicyDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *TaxPolicyDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *TaxPolicyDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyId

`func (o *TaxPolicyDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TaxPolicyDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TaxPolicyDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TaxPolicyDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TaxPolicyDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TaxPolicyDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCountryId

`func (o *TaxPolicyDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TaxPolicyDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TaxPolicyDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TaxPolicyDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TaxPolicyDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TaxPolicyDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *TaxPolicyDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *TaxPolicyDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *TaxPolicyDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *TaxPolicyDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *TaxPolicyDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *TaxPolicyDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCustomState

`func (o *TaxPolicyDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *TaxPolicyDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *TaxPolicyDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *TaxPolicyDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *TaxPolicyDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *TaxPolicyDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *TaxPolicyDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *TaxPolicyDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *TaxPolicyDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *TaxPolicyDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *TaxPolicyDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *TaxPolicyDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityId

`func (o *TaxPolicyDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *TaxPolicyDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *TaxPolicyDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *TaxPolicyDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *TaxPolicyDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *TaxPolicyDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetEnrollmentId

`func (o *TaxPolicyDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TaxPolicyDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TaxPolicyDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TaxPolicyDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TaxPolicyDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TaxPolicyDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTenantId

`func (o *TaxPolicyDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TaxPolicyDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TaxPolicyDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TaxPolicyDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TaxPolicyDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TaxPolicyDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetZero

`func (o *TaxPolicyDto) GetZero() bool`

GetZero returns the Zero field if non-nil, zero value otherwise.

### GetZeroOk

`func (o *TaxPolicyDto) GetZeroOk() (*bool, bool)`

GetZeroOk returns a tuple with the Zero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZero

`func (o *TaxPolicyDto) SetZero(v bool)`

SetZero sets Zero field to given value.

### HasZero

`func (o *TaxPolicyDto) HasZero() bool`

HasZero returns a boolean if a field has been set.

### GetReduced

`func (o *TaxPolicyDto) GetReduced() bool`

GetReduced returns the Reduced field if non-nil, zero value otherwise.

### GetReducedOk

`func (o *TaxPolicyDto) GetReducedOk() (*bool, bool)`

GetReducedOk returns a tuple with the Reduced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduced

`func (o *TaxPolicyDto) SetReduced(v bool)`

SetReduced sets Reduced field to given value.

### HasReduced

`func (o *TaxPolicyDto) HasReduced() bool`

HasReduced returns a boolean if a field has been set.

### GetWithholding

`func (o *TaxPolicyDto) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *TaxPolicyDto) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *TaxPolicyDto) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *TaxPolicyDto) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.

### GetFiscalAuthorityID

`func (o *TaxPolicyDto) GetFiscalAuthorityID() string`

GetFiscalAuthorityID returns the FiscalAuthorityID field if non-nil, zero value otherwise.

### GetFiscalAuthorityIDOk

`func (o *TaxPolicyDto) GetFiscalAuthorityIDOk() (*string, bool)`

GetFiscalAuthorityIDOk returns a tuple with the FiscalAuthorityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityID

`func (o *TaxPolicyDto) SetFiscalAuthorityID(v string)`

SetFiscalAuthorityID sets FiscalAuthorityID field to given value.

### HasFiscalAuthorityID

`func (o *TaxPolicyDto) HasFiscalAuthorityID() bool`

HasFiscalAuthorityID returns a boolean if a field has been set.

### SetFiscalAuthorityIDNil

`func (o *TaxPolicyDto) SetFiscalAuthorityIDNil(b bool)`

 SetFiscalAuthorityIDNil sets the value for FiscalAuthorityID to be an explicit nil

### UnsetFiscalAuthorityID
`func (o *TaxPolicyDto) UnsetFiscalAuthorityID()`

UnsetFiscalAuthorityID ensures that no value is present for FiscalAuthorityID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


