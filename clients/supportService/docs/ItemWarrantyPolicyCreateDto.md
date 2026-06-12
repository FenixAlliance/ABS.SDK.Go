# ItemWarrantyPolicyCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsExtendedWarranty** | Pointer to **bool** |  | [optional] 
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

## Methods

### NewItemWarrantyPolicyCreateDto

`func NewItemWarrantyPolicyCreateDto(title string, ) *ItemWarrantyPolicyCreateDto`

NewItemWarrantyPolicyCreateDto instantiates a new ItemWarrantyPolicyCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWarrantyPolicyCreateDtoWithDefaults

`func NewItemWarrantyPolicyCreateDtoWithDefaults() *ItemWarrantyPolicyCreateDto`

NewItemWarrantyPolicyCreateDtoWithDefaults instantiates a new ItemWarrantyPolicyCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemWarrantyPolicyCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemWarrantyPolicyCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemWarrantyPolicyCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemWarrantyPolicyCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemWarrantyPolicyCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemWarrantyPolicyCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemWarrantyPolicyCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemWarrantyPolicyCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ItemWarrantyPolicyCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemWarrantyPolicyCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemWarrantyPolicyCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemWarrantyPolicyCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemWarrantyPolicyCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemWarrantyPolicyCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemWarrantyPolicyCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemWarrantyPolicyCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemWarrantyPolicyCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsExtendedWarranty

`func (o *ItemWarrantyPolicyCreateDto) GetIsExtendedWarranty() bool`

GetIsExtendedWarranty returns the IsExtendedWarranty field if non-nil, zero value otherwise.

### GetIsExtendedWarrantyOk

`func (o *ItemWarrantyPolicyCreateDto) GetIsExtendedWarrantyOk() (*bool, bool)`

GetIsExtendedWarrantyOk returns a tuple with the IsExtendedWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedWarranty

`func (o *ItemWarrantyPolicyCreateDto) SetIsExtendedWarranty(v bool)`

SetIsExtendedWarranty sets IsExtendedWarranty field to given value.

### HasIsExtendedWarranty

`func (o *ItemWarrantyPolicyCreateDto) HasIsExtendedWarranty() bool`

HasIsExtendedWarranty returns a boolean if a field has been set.

### GetIsFree

`func (o *ItemWarrantyPolicyCreateDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ItemWarrantyPolicyCreateDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ItemWarrantyPolicyCreateDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ItemWarrantyPolicyCreateDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *ItemWarrantyPolicyCreateDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *ItemWarrantyPolicyCreateDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *ItemWarrantyPolicyCreateDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *ItemWarrantyPolicyCreateDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ItemWarrantyPolicyCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ItemWarrantyPolicyCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ItemWarrantyPolicyCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ItemWarrantyPolicyCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *ItemWarrantyPolicyCreateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ItemWarrantyPolicyCreateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ItemWarrantyPolicyCreateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ItemWarrantyPolicyCreateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *ItemWarrantyPolicyCreateDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *ItemWarrantyPolicyCreateDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *ItemWarrantyPolicyCreateDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *ItemWarrantyPolicyCreateDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *ItemWarrantyPolicyCreateDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ItemWarrantyPolicyCreateDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ItemWarrantyPolicyCreateDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ItemWarrantyPolicyCreateDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *ItemWarrantyPolicyCreateDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemWarrantyPolicyCreateDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemWarrantyPolicyCreateDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ItemWarrantyPolicyCreateDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *ItemWarrantyPolicyCreateDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ItemWarrantyPolicyCreateDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ItemWarrantyPolicyCreateDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ItemWarrantyPolicyCreateDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *ItemWarrantyPolicyCreateDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ItemWarrantyPolicyCreateDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ItemWarrantyPolicyCreateDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ItemWarrantyPolicyCreateDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *ItemWarrantyPolicyCreateDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ItemWarrantyPolicyCreateDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ItemWarrantyPolicyCreateDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ItemWarrantyPolicyCreateDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *ItemWarrantyPolicyCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemWarrantyPolicyCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemWarrantyPolicyCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemWarrantyPolicyCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ItemWarrantyPolicyCreateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ItemWarrantyPolicyCreateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ItemWarrantyPolicyCreateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ItemWarrantyPolicyCreateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ItemWarrantyPolicyCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemWarrantyPolicyCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemWarrantyPolicyCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemWarrantyPolicyCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemWarrantyPolicyCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemWarrantyPolicyCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCountryId

`func (o *ItemWarrantyPolicyCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ItemWarrantyPolicyCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ItemWarrantyPolicyCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ItemWarrantyPolicyCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ItemWarrantyPolicyCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ItemWarrantyPolicyCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *ItemWarrantyPolicyCreateDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *ItemWarrantyPolicyCreateDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *ItemWarrantyPolicyCreateDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *ItemWarrantyPolicyCreateDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *ItemWarrantyPolicyCreateDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *ItemWarrantyPolicyCreateDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCustomState

`func (o *ItemWarrantyPolicyCreateDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ItemWarrantyPolicyCreateDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ItemWarrantyPolicyCreateDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ItemWarrantyPolicyCreateDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ItemWarrantyPolicyCreateDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ItemWarrantyPolicyCreateDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *ItemWarrantyPolicyCreateDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *ItemWarrantyPolicyCreateDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *ItemWarrantyPolicyCreateDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *ItemWarrantyPolicyCreateDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *ItemWarrantyPolicyCreateDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *ItemWarrantyPolicyCreateDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityId

`func (o *ItemWarrantyPolicyCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ItemWarrantyPolicyCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ItemWarrantyPolicyCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ItemWarrantyPolicyCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ItemWarrantyPolicyCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ItemWarrantyPolicyCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


