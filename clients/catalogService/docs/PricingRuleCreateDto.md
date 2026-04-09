# PricingRuleCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewPricingRuleCreateDto

`func NewPricingRuleCreateDto() *PricingRuleCreateDto`

NewPricingRuleCreateDto instantiates a new PricingRuleCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingRuleCreateDtoWithDefaults

`func NewPricingRuleCreateDtoWithDefaults() *PricingRuleCreateDto`

NewPricingRuleCreateDtoWithDefaults instantiates a new PricingRuleCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingRuleCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingRuleCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingRuleCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingRuleCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PricingRuleCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PricingRuleCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PricingRuleCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PricingRuleCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCode

`func (o *PricingRuleCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PricingRuleCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PricingRuleCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PricingRuleCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PricingRuleCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PricingRuleCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTitle

`func (o *PricingRuleCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PricingRuleCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PricingRuleCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PricingRuleCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PricingRuleCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PricingRuleCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PricingRuleCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingRuleCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingRuleCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingRuleCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PricingRuleCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PricingRuleCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *PricingRuleCreateDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *PricingRuleCreateDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *PricingRuleCreateDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *PricingRuleCreateDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *PricingRuleCreateDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *PricingRuleCreateDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *PricingRuleCreateDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *PricingRuleCreateDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PricingRuleCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PricingRuleCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PricingRuleCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PricingRuleCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *PricingRuleCreateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PricingRuleCreateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PricingRuleCreateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PricingRuleCreateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *PricingRuleCreateDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *PricingRuleCreateDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *PricingRuleCreateDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *PricingRuleCreateDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *PricingRuleCreateDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *PricingRuleCreateDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *PricingRuleCreateDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *PricingRuleCreateDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *PricingRuleCreateDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PricingRuleCreateDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PricingRuleCreateDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PricingRuleCreateDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *PricingRuleCreateDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *PricingRuleCreateDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *PricingRuleCreateDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *PricingRuleCreateDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *PricingRuleCreateDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *PricingRuleCreateDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *PricingRuleCreateDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *PricingRuleCreateDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *PricingRuleCreateDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *PricingRuleCreateDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *PricingRuleCreateDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *PricingRuleCreateDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *PricingRuleCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PricingRuleCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PricingRuleCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *PricingRuleCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *PricingRuleCreateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PricingRuleCreateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PricingRuleCreateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PricingRuleCreateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyID

`func (o *PricingRuleCreateDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *PricingRuleCreateDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *PricingRuleCreateDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.

### HasCurrencyID

`func (o *PricingRuleCreateDto) HasCurrencyID() bool`

HasCurrencyID returns a boolean if a field has been set.

### SetCurrencyIDNil

`func (o *PricingRuleCreateDto) SetCurrencyIDNil(b bool)`

 SetCurrencyIDNil sets the value for CurrencyID to be an explicit nil

### UnsetCurrencyID
`func (o *PricingRuleCreateDto) UnsetCurrencyID()`

UnsetCurrencyID ensures that no value is present for CurrencyID, not even an explicit nil
### GetCountryID

`func (o *PricingRuleCreateDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *PricingRuleCreateDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *PricingRuleCreateDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *PricingRuleCreateDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *PricingRuleCreateDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *PricingRuleCreateDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetCountryStateID

`func (o *PricingRuleCreateDto) GetCountryStateID() string`

GetCountryStateID returns the CountryStateID field if non-nil, zero value otherwise.

### GetCountryStateIDOk

`func (o *PricingRuleCreateDto) GetCountryStateIDOk() (*string, bool)`

GetCountryStateIDOk returns a tuple with the CountryStateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateID

`func (o *PricingRuleCreateDto) SetCountryStateID(v string)`

SetCountryStateID sets CountryStateID field to given value.

### HasCountryStateID

`func (o *PricingRuleCreateDto) HasCountryStateID() bool`

HasCountryStateID returns a boolean if a field has been set.

### SetCountryStateIDNil

`func (o *PricingRuleCreateDto) SetCountryStateIDNil(b bool)`

 SetCountryStateIDNil sets the value for CountryStateID to be an explicit nil

### UnsetCountryStateID
`func (o *PricingRuleCreateDto) UnsetCountryStateID()`

UnsetCountryStateID ensures that no value is present for CountryStateID, not even an explicit nil
### GetCustomState

`func (o *PricingRuleCreateDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *PricingRuleCreateDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *PricingRuleCreateDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *PricingRuleCreateDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *PricingRuleCreateDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *PricingRuleCreateDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *PricingRuleCreateDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *PricingRuleCreateDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *PricingRuleCreateDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *PricingRuleCreateDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *PricingRuleCreateDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *PricingRuleCreateDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityID

`func (o *PricingRuleCreateDto) GetCityID() string`

GetCityID returns the CityID field if non-nil, zero value otherwise.

### GetCityIDOk

`func (o *PricingRuleCreateDto) GetCityIDOk() (*string, bool)`

GetCityIDOk returns a tuple with the CityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityID

`func (o *PricingRuleCreateDto) SetCityID(v string)`

SetCityID sets CityID field to given value.

### HasCityID

`func (o *PricingRuleCreateDto) HasCityID() bool`

HasCityID returns a boolean if a field has been set.

### SetCityIDNil

`func (o *PricingRuleCreateDto) SetCityIDNil(b bool)`

 SetCityIDNil sets the value for CityID to be an explicit nil

### UnsetCityID
`func (o *PricingRuleCreateDto) UnsetCityID()`

UnsetCityID ensures that no value is present for CityID, not even an explicit nil
### GetBusinessID

`func (o *PricingRuleCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *PricingRuleCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *PricingRuleCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *PricingRuleCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *PricingRuleCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *PricingRuleCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *PricingRuleCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *PricingRuleCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *PricingRuleCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *PricingRuleCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *PricingRuleCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *PricingRuleCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


