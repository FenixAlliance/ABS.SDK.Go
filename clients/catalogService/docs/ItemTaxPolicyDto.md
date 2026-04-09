# ItemTaxPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Zero** | Pointer to **bool** |  | [optional] 
**Reduced** | Pointer to **bool** |  | [optional] 
**Withholding** | Pointer to **bool** |  | [optional] 
**FiscalAuthorityID** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
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
**CurrencyID** | Pointer to **NullableString** |  | [optional] 
**CountryID** | Pointer to **NullableString** |  | [optional] 
**CountryStateID** | Pointer to **NullableString** |  | [optional] 
**CustomState** | Pointer to **NullableString** |  | [optional] 
**CustomCity** | Pointer to **NullableString** |  | [optional] 
**CityID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemTaxPolicyDto

`func NewItemTaxPolicyDto() *ItemTaxPolicyDto`

NewItemTaxPolicyDto instantiates a new ItemTaxPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTaxPolicyDtoWithDefaults

`func NewItemTaxPolicyDtoWithDefaults() *ItemTaxPolicyDto`

NewItemTaxPolicyDtoWithDefaults instantiates a new ItemTaxPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemTaxPolicyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemTaxPolicyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemTaxPolicyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemTaxPolicyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemTaxPolicyDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemTaxPolicyDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemTaxPolicyDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemTaxPolicyDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemTaxPolicyDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemTaxPolicyDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemTaxPolicyDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemTaxPolicyDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetZero

`func (o *ItemTaxPolicyDto) GetZero() bool`

GetZero returns the Zero field if non-nil, zero value otherwise.

### GetZeroOk

`func (o *ItemTaxPolicyDto) GetZeroOk() (*bool, bool)`

GetZeroOk returns a tuple with the Zero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZero

`func (o *ItemTaxPolicyDto) SetZero(v bool)`

SetZero sets Zero field to given value.

### HasZero

`func (o *ItemTaxPolicyDto) HasZero() bool`

HasZero returns a boolean if a field has been set.

### GetReduced

`func (o *ItemTaxPolicyDto) GetReduced() bool`

GetReduced returns the Reduced field if non-nil, zero value otherwise.

### GetReducedOk

`func (o *ItemTaxPolicyDto) GetReducedOk() (*bool, bool)`

GetReducedOk returns a tuple with the Reduced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduced

`func (o *ItemTaxPolicyDto) SetReduced(v bool)`

SetReduced sets Reduced field to given value.

### HasReduced

`func (o *ItemTaxPolicyDto) HasReduced() bool`

HasReduced returns a boolean if a field has been set.

### GetWithholding

`func (o *ItemTaxPolicyDto) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *ItemTaxPolicyDto) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *ItemTaxPolicyDto) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *ItemTaxPolicyDto) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.

### GetFiscalAuthorityID

`func (o *ItemTaxPolicyDto) GetFiscalAuthorityID() string`

GetFiscalAuthorityID returns the FiscalAuthorityID field if non-nil, zero value otherwise.

### GetFiscalAuthorityIDOk

`func (o *ItemTaxPolicyDto) GetFiscalAuthorityIDOk() (*string, bool)`

GetFiscalAuthorityIDOk returns a tuple with the FiscalAuthorityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityID

`func (o *ItemTaxPolicyDto) SetFiscalAuthorityID(v string)`

SetFiscalAuthorityID sets FiscalAuthorityID field to given value.

### HasFiscalAuthorityID

`func (o *ItemTaxPolicyDto) HasFiscalAuthorityID() bool`

HasFiscalAuthorityID returns a boolean if a field has been set.

### SetFiscalAuthorityIDNil

`func (o *ItemTaxPolicyDto) SetFiscalAuthorityIDNil(b bool)`

 SetFiscalAuthorityIDNil sets the value for FiscalAuthorityID to be an explicit nil

### UnsetFiscalAuthorityID
`func (o *ItemTaxPolicyDto) UnsetFiscalAuthorityID()`

UnsetFiscalAuthorityID ensures that no value is present for FiscalAuthorityID, not even an explicit nil
### GetType

`func (o *ItemTaxPolicyDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemTaxPolicyDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemTaxPolicyDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ItemTaxPolicyDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ItemTaxPolicyDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ItemTaxPolicyDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *ItemTaxPolicyDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ItemTaxPolicyDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ItemTaxPolicyDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ItemTaxPolicyDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ItemTaxPolicyDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ItemTaxPolicyDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTitle

`func (o *ItemTaxPolicyDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemTaxPolicyDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemTaxPolicyDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemTaxPolicyDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemTaxPolicyDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemTaxPolicyDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ItemTaxPolicyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemTaxPolicyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemTaxPolicyDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemTaxPolicyDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemTaxPolicyDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemTaxPolicyDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *ItemTaxPolicyDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ItemTaxPolicyDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ItemTaxPolicyDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ItemTaxPolicyDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *ItemTaxPolicyDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *ItemTaxPolicyDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *ItemTaxPolicyDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *ItemTaxPolicyDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ItemTaxPolicyDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ItemTaxPolicyDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ItemTaxPolicyDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ItemTaxPolicyDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *ItemTaxPolicyDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ItemTaxPolicyDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ItemTaxPolicyDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ItemTaxPolicyDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *ItemTaxPolicyDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *ItemTaxPolicyDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *ItemTaxPolicyDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *ItemTaxPolicyDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *ItemTaxPolicyDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ItemTaxPolicyDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ItemTaxPolicyDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ItemTaxPolicyDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *ItemTaxPolicyDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemTaxPolicyDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemTaxPolicyDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ItemTaxPolicyDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *ItemTaxPolicyDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ItemTaxPolicyDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ItemTaxPolicyDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ItemTaxPolicyDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *ItemTaxPolicyDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ItemTaxPolicyDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ItemTaxPolicyDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ItemTaxPolicyDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *ItemTaxPolicyDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ItemTaxPolicyDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ItemTaxPolicyDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ItemTaxPolicyDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *ItemTaxPolicyDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemTaxPolicyDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemTaxPolicyDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemTaxPolicyDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ItemTaxPolicyDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ItemTaxPolicyDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ItemTaxPolicyDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ItemTaxPolicyDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyID

`func (o *ItemTaxPolicyDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *ItemTaxPolicyDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *ItemTaxPolicyDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.

### HasCurrencyID

`func (o *ItemTaxPolicyDto) HasCurrencyID() bool`

HasCurrencyID returns a boolean if a field has been set.

### SetCurrencyIDNil

`func (o *ItemTaxPolicyDto) SetCurrencyIDNil(b bool)`

 SetCurrencyIDNil sets the value for CurrencyID to be an explicit nil

### UnsetCurrencyID
`func (o *ItemTaxPolicyDto) UnsetCurrencyID()`

UnsetCurrencyID ensures that no value is present for CurrencyID, not even an explicit nil
### GetCountryID

`func (o *ItemTaxPolicyDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *ItemTaxPolicyDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *ItemTaxPolicyDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *ItemTaxPolicyDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *ItemTaxPolicyDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *ItemTaxPolicyDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetCountryStateID

`func (o *ItemTaxPolicyDto) GetCountryStateID() string`

GetCountryStateID returns the CountryStateID field if non-nil, zero value otherwise.

### GetCountryStateIDOk

`func (o *ItemTaxPolicyDto) GetCountryStateIDOk() (*string, bool)`

GetCountryStateIDOk returns a tuple with the CountryStateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateID

`func (o *ItemTaxPolicyDto) SetCountryStateID(v string)`

SetCountryStateID sets CountryStateID field to given value.

### HasCountryStateID

`func (o *ItemTaxPolicyDto) HasCountryStateID() bool`

HasCountryStateID returns a boolean if a field has been set.

### SetCountryStateIDNil

`func (o *ItemTaxPolicyDto) SetCountryStateIDNil(b bool)`

 SetCountryStateIDNil sets the value for CountryStateID to be an explicit nil

### UnsetCountryStateID
`func (o *ItemTaxPolicyDto) UnsetCountryStateID()`

UnsetCountryStateID ensures that no value is present for CountryStateID, not even an explicit nil
### GetCustomState

`func (o *ItemTaxPolicyDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ItemTaxPolicyDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ItemTaxPolicyDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ItemTaxPolicyDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ItemTaxPolicyDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ItemTaxPolicyDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *ItemTaxPolicyDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *ItemTaxPolicyDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *ItemTaxPolicyDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *ItemTaxPolicyDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *ItemTaxPolicyDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *ItemTaxPolicyDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityID

`func (o *ItemTaxPolicyDto) GetCityID() string`

GetCityID returns the CityID field if non-nil, zero value otherwise.

### GetCityIDOk

`func (o *ItemTaxPolicyDto) GetCityIDOk() (*string, bool)`

GetCityIDOk returns a tuple with the CityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityID

`func (o *ItemTaxPolicyDto) SetCityID(v string)`

SetCityID sets CityID field to given value.

### HasCityID

`func (o *ItemTaxPolicyDto) HasCityID() bool`

HasCityID returns a boolean if a field has been set.

### SetCityIDNil

`func (o *ItemTaxPolicyDto) SetCityIDNil(b bool)`

 SetCityIDNil sets the value for CityID to be an explicit nil

### UnsetCityID
`func (o *ItemTaxPolicyDto) UnsetCityID()`

UnsetCityID ensures that no value is present for CityID, not even an explicit nil
### GetBusinessID

`func (o *ItemTaxPolicyDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemTaxPolicyDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemTaxPolicyDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ItemTaxPolicyDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ItemTaxPolicyDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ItemTaxPolicyDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *ItemTaxPolicyDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *ItemTaxPolicyDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *ItemTaxPolicyDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *ItemTaxPolicyDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *ItemTaxPolicyDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *ItemTaxPolicyDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


