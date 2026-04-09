# ItemWarrantyPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**IsExtendedWarranty** | Pointer to **bool** |  | [optional] 
**Code** | **string** |  | 
**Title** | **string** |  | 
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
**BusinessID** | **string** |  | 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemWarrantyPolicyDto

`func NewItemWarrantyPolicyDto(code string, title string, businessID string, ) *ItemWarrantyPolicyDto`

NewItemWarrantyPolicyDto instantiates a new ItemWarrantyPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWarrantyPolicyDtoWithDefaults

`func NewItemWarrantyPolicyDtoWithDefaults() *ItemWarrantyPolicyDto`

NewItemWarrantyPolicyDtoWithDefaults instantiates a new ItemWarrantyPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemWarrantyPolicyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemWarrantyPolicyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemWarrantyPolicyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemWarrantyPolicyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemWarrantyPolicyDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemWarrantyPolicyDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemWarrantyPolicyDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemWarrantyPolicyDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemWarrantyPolicyDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemWarrantyPolicyDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemWarrantyPolicyDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemWarrantyPolicyDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetIsExtendedWarranty

`func (o *ItemWarrantyPolicyDto) GetIsExtendedWarranty() bool`

GetIsExtendedWarranty returns the IsExtendedWarranty field if non-nil, zero value otherwise.

### GetIsExtendedWarrantyOk

`func (o *ItemWarrantyPolicyDto) GetIsExtendedWarrantyOk() (*bool, bool)`

GetIsExtendedWarrantyOk returns a tuple with the IsExtendedWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedWarranty

`func (o *ItemWarrantyPolicyDto) SetIsExtendedWarranty(v bool)`

SetIsExtendedWarranty sets IsExtendedWarranty field to given value.

### HasIsExtendedWarranty

`func (o *ItemWarrantyPolicyDto) HasIsExtendedWarranty() bool`

HasIsExtendedWarranty returns a boolean if a field has been set.

### GetCode

`func (o *ItemWarrantyPolicyDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ItemWarrantyPolicyDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ItemWarrantyPolicyDto) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ItemWarrantyPolicyDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemWarrantyPolicyDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemWarrantyPolicyDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItemWarrantyPolicyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemWarrantyPolicyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemWarrantyPolicyDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemWarrantyPolicyDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemWarrantyPolicyDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemWarrantyPolicyDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *ItemWarrantyPolicyDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ItemWarrantyPolicyDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ItemWarrantyPolicyDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ItemWarrantyPolicyDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetReduce

`func (o *ItemWarrantyPolicyDto) GetReduce() bool`

GetReduce returns the Reduce field if non-nil, zero value otherwise.

### GetReduceOk

`func (o *ItemWarrantyPolicyDto) GetReduceOk() (*bool, bool)`

GetReduceOk returns a tuple with the Reduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduce

`func (o *ItemWarrantyPolicyDto) SetReduce(v bool)`

SetReduce sets Reduce field to given value.

### HasReduce

`func (o *ItemWarrantyPolicyDto) HasReduce() bool`

HasReduce returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ItemWarrantyPolicyDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ItemWarrantyPolicyDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ItemWarrantyPolicyDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ItemWarrantyPolicyDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *ItemWarrantyPolicyDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ItemWarrantyPolicyDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ItemWarrantyPolicyDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ItemWarrantyPolicyDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAllowInternational

`func (o *ItemWarrantyPolicyDto) GetAllowInternational() bool`

GetAllowInternational returns the AllowInternational field if non-nil, zero value otherwise.

### GetAllowInternationalOk

`func (o *ItemWarrantyPolicyDto) GetAllowInternationalOk() (*bool, bool)`

GetAllowInternationalOk returns a tuple with the AllowInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternational

`func (o *ItemWarrantyPolicyDto) SetAllowInternational(v bool)`

SetAllowInternational sets AllowInternational field to given value.

### HasAllowInternational

`func (o *ItemWarrantyPolicyDto) HasAllowInternational() bool`

HasAllowInternational returns a boolean if a field has been set.

### GetHours

`func (o *ItemWarrantyPolicyDto) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ItemWarrantyPolicyDto) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ItemWarrantyPolicyDto) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ItemWarrantyPolicyDto) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *ItemWarrantyPolicyDto) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemWarrantyPolicyDto) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemWarrantyPolicyDto) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ItemWarrantyPolicyDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetWeeks

`func (o *ItemWarrantyPolicyDto) GetWeeks() int32`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ItemWarrantyPolicyDto) GetWeeksOk() (*int32, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ItemWarrantyPolicyDto) SetWeeks(v int32)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ItemWarrantyPolicyDto) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### GetMonths

`func (o *ItemWarrantyPolicyDto) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ItemWarrantyPolicyDto) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ItemWarrantyPolicyDto) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ItemWarrantyPolicyDto) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetYears

`func (o *ItemWarrantyPolicyDto) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ItemWarrantyPolicyDto) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ItemWarrantyPolicyDto) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ItemWarrantyPolicyDto) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetValue

`func (o *ItemWarrantyPolicyDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemWarrantyPolicyDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemWarrantyPolicyDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemWarrantyPolicyDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercentage

`func (o *ItemWarrantyPolicyDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ItemWarrantyPolicyDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ItemWarrantyPolicyDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ItemWarrantyPolicyDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrencyID

`func (o *ItemWarrantyPolicyDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *ItemWarrantyPolicyDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *ItemWarrantyPolicyDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.

### HasCurrencyID

`func (o *ItemWarrantyPolicyDto) HasCurrencyID() bool`

HasCurrencyID returns a boolean if a field has been set.

### SetCurrencyIDNil

`func (o *ItemWarrantyPolicyDto) SetCurrencyIDNil(b bool)`

 SetCurrencyIDNil sets the value for CurrencyID to be an explicit nil

### UnsetCurrencyID
`func (o *ItemWarrantyPolicyDto) UnsetCurrencyID()`

UnsetCurrencyID ensures that no value is present for CurrencyID, not even an explicit nil
### GetCountryID

`func (o *ItemWarrantyPolicyDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *ItemWarrantyPolicyDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *ItemWarrantyPolicyDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *ItemWarrantyPolicyDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *ItemWarrantyPolicyDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *ItemWarrantyPolicyDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetCountryStateID

`func (o *ItemWarrantyPolicyDto) GetCountryStateID() string`

GetCountryStateID returns the CountryStateID field if non-nil, zero value otherwise.

### GetCountryStateIDOk

`func (o *ItemWarrantyPolicyDto) GetCountryStateIDOk() (*string, bool)`

GetCountryStateIDOk returns a tuple with the CountryStateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateID

`func (o *ItemWarrantyPolicyDto) SetCountryStateID(v string)`

SetCountryStateID sets CountryStateID field to given value.

### HasCountryStateID

`func (o *ItemWarrantyPolicyDto) HasCountryStateID() bool`

HasCountryStateID returns a boolean if a field has been set.

### SetCountryStateIDNil

`func (o *ItemWarrantyPolicyDto) SetCountryStateIDNil(b bool)`

 SetCountryStateIDNil sets the value for CountryStateID to be an explicit nil

### UnsetCountryStateID
`func (o *ItemWarrantyPolicyDto) UnsetCountryStateID()`

UnsetCountryStateID ensures that no value is present for CountryStateID, not even an explicit nil
### GetCustomState

`func (o *ItemWarrantyPolicyDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ItemWarrantyPolicyDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ItemWarrantyPolicyDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ItemWarrantyPolicyDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ItemWarrantyPolicyDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ItemWarrantyPolicyDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetCustomCity

`func (o *ItemWarrantyPolicyDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *ItemWarrantyPolicyDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *ItemWarrantyPolicyDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *ItemWarrantyPolicyDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *ItemWarrantyPolicyDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *ItemWarrantyPolicyDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCityID

`func (o *ItemWarrantyPolicyDto) GetCityID() string`

GetCityID returns the CityID field if non-nil, zero value otherwise.

### GetCityIDOk

`func (o *ItemWarrantyPolicyDto) GetCityIDOk() (*string, bool)`

GetCityIDOk returns a tuple with the CityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityID

`func (o *ItemWarrantyPolicyDto) SetCityID(v string)`

SetCityID sets CityID field to given value.

### HasCityID

`func (o *ItemWarrantyPolicyDto) HasCityID() bool`

HasCityID returns a boolean if a field has been set.

### SetCityIDNil

`func (o *ItemWarrantyPolicyDto) SetCityIDNil(b bool)`

 SetCityIDNil sets the value for CityID to be an explicit nil

### UnsetCityID
`func (o *ItemWarrantyPolicyDto) UnsetCityID()`

UnsetCityID ensures that no value is present for CityID, not even an explicit nil
### GetBusinessID

`func (o *ItemWarrantyPolicyDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemWarrantyPolicyDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemWarrantyPolicyDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetBusinessProfileRecordID

`func (o *ItemWarrantyPolicyDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *ItemWarrantyPolicyDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *ItemWarrantyPolicyDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *ItemWarrantyPolicyDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *ItemWarrantyPolicyDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *ItemWarrantyPolicyDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


