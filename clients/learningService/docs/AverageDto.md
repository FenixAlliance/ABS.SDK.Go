# AverageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Average** | Pointer to **float64** |  | [optional] 

## Methods

### NewAverageDto

`func NewAverageDto() *AverageDto`

NewAverageDto instantiates a new AverageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAverageDtoWithDefaults

`func NewAverageDtoWithDefaults() *AverageDto`

NewAverageDtoWithDefaults instantiates a new AverageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AverageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AverageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AverageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AverageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AverageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AverageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AverageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AverageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AverageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AverageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AverageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AverageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetAverage

`func (o *AverageDto) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *AverageDto) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *AverageDto) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *AverageDto) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


