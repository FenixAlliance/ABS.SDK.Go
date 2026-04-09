# ForexRates

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

### NewForexRates

`func NewForexRates() *ForexRates`

NewForexRates instantiates a new ForexRates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexRatesWithDefaults

`func NewForexRatesWithDefaults() *ForexRates`

NewForexRatesWithDefaults instantiates a new ForexRates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ForexRates) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ForexRates) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ForexRates) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ForexRates) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetDate

`func (o *ForexRates) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ForexRates) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ForexRates) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ForexRates) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ForexRates) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ForexRates) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetBase

`func (o *ForexRates) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ForexRates) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ForexRates) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *ForexRates) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *ForexRates) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *ForexRates) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetTimestamp

`func (o *ForexRates) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ForexRates) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ForexRates) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ForexRates) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *ForexRates) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *ForexRates) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *ForexRates) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *ForexRates) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetRates

`func (o *ForexRates) GetRates() map[string]float64`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ForexRates) GetRatesOk() (*map[string]float64, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ForexRates) SetRates(v map[string]float64)`

SetRates sets Rates field to given value.

### HasRates

`func (o *ForexRates) HasRates() bool`

HasRates returns a boolean if a field has been set.

### SetRatesNil

`func (o *ForexRates) SetRatesNil(b bool)`

 SetRatesNil sets the value for Rates to be an explicit nil

### UnsetRates
`func (o *ForexRates) UnsetRates()`

UnsetRates ensures that no value is present for Rates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


