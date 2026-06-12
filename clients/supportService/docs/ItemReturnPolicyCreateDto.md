# ItemReturnPolicyCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 
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

### NewItemReturnPolicyCreateDto

`func NewItemReturnPolicyCreateDto(title string, ) *ItemReturnPolicyCreateDto`

NewItemReturnPolicyCreateDto instantiates a new ItemReturnPolicyCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReturnPolicyCreateDtoWithDefaults

`func NewItemReturnPolicyCreateDtoWithDefaults() *ItemReturnPolicyCreateDto`

NewItemReturnPolicyCreateDtoWithDefaults instantiates a new ItemReturnPolicyCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemReturnPolicyCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemReturnPolicyCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemReturnPolicyCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemReturnPolicyCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemReturnPolicyCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemReturnPolicyCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemReturnPolicyCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemReturnPolicyCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ItemReturnPolicyCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemReturnPolicyCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemReturnPolicyCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemReturnPolicyCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemReturnPolicyCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemReturnPolicyCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemReturnPolicyCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemReturnPolicyCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemReturnPolicyCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShippingCourierId

`func (o *ItemReturnPolicyCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *ItemReturnPolicyCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *ItemReturnPolicyCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *ItemReturnPolicyCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *ItemReturnPolicyCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *ItemReturnPolicyCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetIsFree

`func (o *ItemReturnPolicyCreateDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ItemReturnPolicyCreateDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ItemReturnPolicyCreateDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ItemReturnPolicyCreateDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *ItemReturnPolicyCreateDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *ItemReturnPolicyCreateDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *ItemReturnPolicyCreateDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *ItemReturnPolicyCreateDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ItemReturnPolicyCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ItemReturnPolicyCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ItemReturnPolicyCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ItemReturnPolicyCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *ItemReturnPolicyCreateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ItemReturnPolicyCreateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ItemReturnPolicyCreateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ItemReturnPolicyCreateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *ItemReturnPolicyCreateDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *ItemReturnPolicyCreateDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *ItemReturnPolicyCreateDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *ItemReturnPolicyCreateDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *ItemReturnPolicyCreateDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ItemReturnPolicyCreateDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ItemReturnPolicyCreateDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ItemReturnPolicyCreateDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *ItemReturnPolicyCreateDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemReturnPolicyCreateDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemReturnPolicyCreateDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ItemReturnPolicyCreateDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *ItemReturnPolicyCreateDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ItemReturnPolicyCreateDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ItemReturnPolicyCreateDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ItemReturnPolicyCreateDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *ItemReturnPolicyCreateDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ItemReturnPolicyCreateDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ItemReturnPolicyCreateDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ItemReturnPolicyCreateDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *ItemReturnPolicyCreateDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ItemReturnPolicyCreateDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ItemReturnPolicyCreateDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ItemReturnPolicyCreateDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *ItemReturnPolicyCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemReturnPolicyCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemReturnPolicyCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemReturnPolicyCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ItemReturnPolicyCreateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ItemReturnPolicyCreateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ItemReturnPolicyCreateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ItemReturnPolicyCreateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ItemReturnPolicyCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ItemReturnPolicyCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ItemReturnPolicyCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ItemReturnPolicyCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ItemReturnPolicyCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ItemReturnPolicyCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCountryId

`func (o *ItemReturnPolicyCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ItemReturnPolicyCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ItemReturnPolicyCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ItemReturnPolicyCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ItemReturnPolicyCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ItemReturnPolicyCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *ItemReturnPolicyCreateDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *ItemReturnPolicyCreateDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *ItemReturnPolicyCreateDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *ItemReturnPolicyCreateDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *ItemReturnPolicyCreateDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *ItemReturnPolicyCreateDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCustomState

`func (o *ItemReturnPolicyCreateDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ItemReturnPolicyCreateDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ItemReturnPolicyCreateDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ItemReturnPolicyCreateDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ItemReturnPolicyCreateDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ItemReturnPolicyCreateDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *ItemReturnPolicyCreateDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *ItemReturnPolicyCreateDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *ItemReturnPolicyCreateDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *ItemReturnPolicyCreateDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *ItemReturnPolicyCreateDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *ItemReturnPolicyCreateDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityId

`func (o *ItemReturnPolicyCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ItemReturnPolicyCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ItemReturnPolicyCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ItemReturnPolicyCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ItemReturnPolicyCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ItemReturnPolicyCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


