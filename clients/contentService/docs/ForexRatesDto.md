# ForexRatesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **NullableString** |  | [optional] 
**Base** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 
**Rates** | Pointer to **map[string]float64** |  | [optional] 

## Methods

### NewForexRatesDto

`func NewForexRatesDto() *ForexRatesDto`

NewForexRatesDto instantiates a new ForexRatesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexRatesDtoWithDefaults

`func NewForexRatesDtoWithDefaults() *ForexRatesDto`

NewForexRatesDtoWithDefaults instantiates a new ForexRatesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ForexRatesDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ForexRatesDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ForexRatesDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ForexRatesDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetDate

`func (o *ForexRatesDto) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ForexRatesDto) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ForexRatesDto) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ForexRatesDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ForexRatesDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ForexRatesDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetBase

`func (o *ForexRatesDto) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ForexRatesDto) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ForexRatesDto) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *ForexRatesDto) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *ForexRatesDto) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *ForexRatesDto) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetTimestamp

`func (o *ForexRatesDto) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ForexRatesDto) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ForexRatesDto) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ForexRatesDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *ForexRatesDto) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *ForexRatesDto) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *ForexRatesDto) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *ForexRatesDto) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetRates

`func (o *ForexRatesDto) GetRates() map[string]float64`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ForexRatesDto) GetRatesOk() (*map[string]float64, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ForexRatesDto) SetRates(v map[string]float64)`

SetRates sets Rates field to given value.

### HasRates

`func (o *ForexRatesDto) HasRates() bool`

HasRates returns a boolean if a field has been set.

### SetRatesNil

`func (o *ForexRatesDto) SetRatesNil(b bool)`

 SetRatesNil sets the value for Rates to be an explicit nil

### UnsetRates
`func (o *ForexRatesDto) UnsetRates()`

UnsetRates ensures that no value is present for Rates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


