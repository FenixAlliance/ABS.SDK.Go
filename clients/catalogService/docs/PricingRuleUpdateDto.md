# PricingRuleUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewPricingRuleUpdateDto

`func NewPricingRuleUpdateDto() *PricingRuleUpdateDto`

NewPricingRuleUpdateDto instantiates a new PricingRuleUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingRuleUpdateDtoWithDefaults

`func NewPricingRuleUpdateDtoWithDefaults() *PricingRuleUpdateDto`

NewPricingRuleUpdateDtoWithDefaults instantiates a new PricingRuleUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PricingRuleUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PricingRuleUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PricingRuleUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PricingRuleUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PricingRuleUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PricingRuleUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PricingRuleUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingRuleUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingRuleUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingRuleUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PricingRuleUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PricingRuleUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *PricingRuleUpdateDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *PricingRuleUpdateDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *PricingRuleUpdateDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *PricingRuleUpdateDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *PricingRuleUpdateDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *PricingRuleUpdateDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *PricingRuleUpdateDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *PricingRuleUpdateDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PricingRuleUpdateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PricingRuleUpdateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PricingRuleUpdateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PricingRuleUpdateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *PricingRuleUpdateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PricingRuleUpdateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PricingRuleUpdateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PricingRuleUpdateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *PricingRuleUpdateDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *PricingRuleUpdateDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *PricingRuleUpdateDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *PricingRuleUpdateDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *PricingRuleUpdateDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *PricingRuleUpdateDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *PricingRuleUpdateDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *PricingRuleUpdateDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *PricingRuleUpdateDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PricingRuleUpdateDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PricingRuleUpdateDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PricingRuleUpdateDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *PricingRuleUpdateDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *PricingRuleUpdateDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *PricingRuleUpdateDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *PricingRuleUpdateDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *PricingRuleUpdateDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *PricingRuleUpdateDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *PricingRuleUpdateDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *PricingRuleUpdateDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *PricingRuleUpdateDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *PricingRuleUpdateDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *PricingRuleUpdateDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *PricingRuleUpdateDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *PricingRuleUpdateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PricingRuleUpdateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PricingRuleUpdateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *PricingRuleUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *PricingRuleUpdateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PricingRuleUpdateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PricingRuleUpdateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PricingRuleUpdateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyID

`func (o *PricingRuleUpdateDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *PricingRuleUpdateDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *PricingRuleUpdateDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.

### HasCurrencyID

`func (o *PricingRuleUpdateDto) HasCurrencyID() bool`

HasCurrencyID returns a boolean if a field has been set.

### SetCurrencyIDNil

`func (o *PricingRuleUpdateDto) SetCurrencyIDNil(b bool)`

 SetCurrencyIDNil sets the value for CurrencyID to be an explicit nil

### UnsetCurrencyID
`func (o *PricingRuleUpdateDto) UnsetCurrencyID()`

UnsetCurrencyID ensures that no value is present for CurrencyID, not even an explicit nil
### GetCountryID

`func (o *PricingRuleUpdateDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *PricingRuleUpdateDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *PricingRuleUpdateDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *PricingRuleUpdateDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *PricingRuleUpdateDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *PricingRuleUpdateDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetCountryStateID

`func (o *PricingRuleUpdateDto) GetCountryStateID() string`

GetCountryStateID returns the CountryStateID field if non-nil, zero value otherwise.

### GetCountryStateIDOk

`func (o *PricingRuleUpdateDto) GetCountryStateIDOk() (*string, bool)`

GetCountryStateIDOk returns a tuple with the CountryStateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateID

`func (o *PricingRuleUpdateDto) SetCountryStateID(v string)`

SetCountryStateID sets CountryStateID field to given value.

### HasCountryStateID

`func (o *PricingRuleUpdateDto) HasCountryStateID() bool`

HasCountryStateID returns a boolean if a field has been set.

### SetCountryStateIDNil

`func (o *PricingRuleUpdateDto) SetCountryStateIDNil(b bool)`

 SetCountryStateIDNil sets the value for CountryStateID to be an explicit nil

### UnsetCountryStateID
`func (o *PricingRuleUpdateDto) UnsetCountryStateID()`

UnsetCountryStateID ensures that no value is present for CountryStateID, not even an explicit nil
### GetCustomState

`func (o *PricingRuleUpdateDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *PricingRuleUpdateDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *PricingRuleUpdateDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *PricingRuleUpdateDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *PricingRuleUpdateDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *PricingRuleUpdateDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *PricingRuleUpdateDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *PricingRuleUpdateDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *PricingRuleUpdateDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *PricingRuleUpdateDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *PricingRuleUpdateDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *PricingRuleUpdateDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityID

`func (o *PricingRuleUpdateDto) GetCityID() string`

GetCityID returns the CityID field if non-nil, zero value otherwise.

### GetCityIDOk

`func (o *PricingRuleUpdateDto) GetCityIDOk() (*string, bool)`

GetCityIDOk returns a tuple with the CityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityID

`func (o *PricingRuleUpdateDto) SetCityID(v string)`

SetCityID sets CityID field to given value.

### HasCityID

`func (o *PricingRuleUpdateDto) HasCityID() bool`

HasCityID returns a boolean if a field has been set.

### SetCityIDNil

`func (o *PricingRuleUpdateDto) SetCityIDNil(b bool)`

 SetCityIDNil sets the value for CityID to be an explicit nil

### UnsetCityID
`func (o *PricingRuleUpdateDto) UnsetCityID()`

UnsetCityID ensures that no value is present for CityID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


