# ItemShippingPolicyCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Code** | **string** |  | 
**IsExpressShipmentPolicy** | Pointer to **bool** |  | [optional] 
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
**CurrencyID** | **string** |  | 
**CountryID** | Pointer to **NullableString** |  | [optional] 
**CountryStateID** | Pointer to **NullableString** |  | [optional] 
**CustomState** | Pointer to **NullableString** |  | [optional] 
**CustomCity** | Pointer to **NullableString** |  | [optional] 
**CityID** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierID** | **string** |  | 

## Methods

### NewItemShippingPolicyCreateDto

`func NewItemShippingPolicyCreateDto(title string, type_ string, code string, currencyID string, shippingCourierID string, ) *ItemShippingPolicyCreateDto`

NewItemShippingPolicyCreateDto instantiates a new ItemShippingPolicyCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemShippingPolicyCreateDtoWithDefaults

`func NewItemShippingPolicyCreateDtoWithDefaults() *ItemShippingPolicyCreateDto`

NewItemShippingPolicyCreateDtoWithDefaults instantiates a new ItemShippingPolicyCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemShippingPolicyCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemShippingPolicyCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemShippingPolicyCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemShippingPolicyCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemShippingPolicyCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemShippingPolicyCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemShippingPolicyCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemShippingPolicyCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ItemShippingPolicyCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemShippingPolicyCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemShippingPolicyCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemShippingPolicyCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemShippingPolicyCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemShippingPolicyCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemShippingPolicyCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemShippingPolicyCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemShippingPolicyCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *ItemShippingPolicyCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemShippingPolicyCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemShippingPolicyCreateDto) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ItemShippingPolicyCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ItemShippingPolicyCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ItemShippingPolicyCreateDto) SetCode(v string)`

SetCode sets Code field to given value.


### GetIsExpressShipmentPolicy

`func (o *ItemShippingPolicyCreateDto) GetIsExpressShipmentPolicy() bool`

GetIsExpressShipmentPolicy returns the IsExpressShipmentPolicy field if non-nil, zero value otherwise.

### GetIsExpressShipmentPolicyOk

`func (o *ItemShippingPolicyCreateDto) GetIsExpressShipmentPolicyOk() (*bool, bool)`

GetIsExpressShipmentPolicyOk returns a tuple with the IsExpressShipmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpressShipmentPolicy

`func (o *ItemShippingPolicyCreateDto) SetIsExpressShipmentPolicy(v bool)`

SetIsExpressShipmentPolicy sets IsExpressShipmentPolicy field to given value.

### HasIsExpressShipmentPolicy

`func (o *ItemShippingPolicyCreateDto) HasIsExpressShipmentPolicy() bool`

HasIsExpressShipmentPolicy returns a boolean if a field has been set.

### GetIsFree

`func (o *ItemShippingPolicyCreateDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ItemShippingPolicyCreateDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ItemShippingPolicyCreateDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ItemShippingPolicyCreateDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *ItemShippingPolicyCreateDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *ItemShippingPolicyCreateDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *ItemShippingPolicyCreateDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *ItemShippingPolicyCreateDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ItemShippingPolicyCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ItemShippingPolicyCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ItemShippingPolicyCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ItemShippingPolicyCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *ItemShippingPolicyCreateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ItemShippingPolicyCreateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ItemShippingPolicyCreateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ItemShippingPolicyCreateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *ItemShippingPolicyCreateDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *ItemShippingPolicyCreateDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *ItemShippingPolicyCreateDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *ItemShippingPolicyCreateDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *ItemShippingPolicyCreateDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ItemShippingPolicyCreateDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ItemShippingPolicyCreateDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ItemShippingPolicyCreateDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *ItemShippingPolicyCreateDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemShippingPolicyCreateDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemShippingPolicyCreateDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ItemShippingPolicyCreateDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *ItemShippingPolicyCreateDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ItemShippingPolicyCreateDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ItemShippingPolicyCreateDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ItemShippingPolicyCreateDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *ItemShippingPolicyCreateDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ItemShippingPolicyCreateDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ItemShippingPolicyCreateDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ItemShippingPolicyCreateDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *ItemShippingPolicyCreateDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ItemShippingPolicyCreateDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ItemShippingPolicyCreateDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ItemShippingPolicyCreateDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *ItemShippingPolicyCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemShippingPolicyCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemShippingPolicyCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemShippingPolicyCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ItemShippingPolicyCreateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ItemShippingPolicyCreateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ItemShippingPolicyCreateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ItemShippingPolicyCreateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyID

`func (o *ItemShippingPolicyCreateDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *ItemShippingPolicyCreateDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *ItemShippingPolicyCreateDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.


### GetCountryID

`func (o *ItemShippingPolicyCreateDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *ItemShippingPolicyCreateDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *ItemShippingPolicyCreateDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *ItemShippingPolicyCreateDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *ItemShippingPolicyCreateDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *ItemShippingPolicyCreateDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetCountryStateID

`func (o *ItemShippingPolicyCreateDto) GetCountryStateID() string`

GetCountryStateID returns the CountryStateID field if non-nil, zero value otherwise.

### GetCountryStateIDOk

`func (o *ItemShippingPolicyCreateDto) GetCountryStateIDOk() (*string, bool)`

GetCountryStateIDOk returns a tuple with the CountryStateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateID

`func (o *ItemShippingPolicyCreateDto) SetCountryStateID(v string)`

SetCountryStateID sets CountryStateID field to given value.

### HasCountryStateID

`func (o *ItemShippingPolicyCreateDto) HasCountryStateID() bool`

HasCountryStateID returns a boolean if a field has been set.

### SetCountryStateIDNil

`func (o *ItemShippingPolicyCreateDto) SetCountryStateIDNil(b bool)`

 SetCountryStateIDNil sets the value for CountryStateID to be an explicit nil

### UnsetCountryStateID
`func (o *ItemShippingPolicyCreateDto) UnsetCountryStateID()`

UnsetCountryStateID ensures that no value is present for CountryStateID, not even an explicit nil
### GetCustomState

`func (o *ItemShippingPolicyCreateDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ItemShippingPolicyCreateDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ItemShippingPolicyCreateDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ItemShippingPolicyCreateDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ItemShippingPolicyCreateDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ItemShippingPolicyCreateDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *ItemShippingPolicyCreateDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *ItemShippingPolicyCreateDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *ItemShippingPolicyCreateDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *ItemShippingPolicyCreateDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *ItemShippingPolicyCreateDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *ItemShippingPolicyCreateDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityID

`func (o *ItemShippingPolicyCreateDto) GetCityID() string`

GetCityID returns the CityID field if non-nil, zero value otherwise.

### GetCityIDOk

`func (o *ItemShippingPolicyCreateDto) GetCityIDOk() (*string, bool)`

GetCityIDOk returns a tuple with the CityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityID

`func (o *ItemShippingPolicyCreateDto) SetCityID(v string)`

SetCityID sets CityID field to given value.

### HasCityID

`func (o *ItemShippingPolicyCreateDto) HasCityID() bool`

HasCityID returns a boolean if a field has been set.

### SetCityIDNil

`func (o *ItemShippingPolicyCreateDto) SetCityIDNil(b bool)`

 SetCityIDNil sets the value for CityID to be an explicit nil

### UnsetCityID
`func (o *ItemShippingPolicyCreateDto) UnsetCityID()`

UnsetCityID ensures that no value is present for CityID, not even an explicit nil
### GetShippingCourierID

`func (o *ItemShippingPolicyCreateDto) GetShippingCourierID() string`

GetShippingCourierID returns the ShippingCourierID field if non-nil, zero value otherwise.

### GetShippingCourierIDOk

`func (o *ItemShippingPolicyCreateDto) GetShippingCourierIDOk() (*string, bool)`

GetShippingCourierIDOk returns a tuple with the ShippingCourierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierID

`func (o *ItemShippingPolicyCreateDto) SetShippingCourierID(v string)`

SetShippingCourierID sets ShippingCourierID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


