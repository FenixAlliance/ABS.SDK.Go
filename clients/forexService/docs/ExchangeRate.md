# ExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**Money**](Money.md) |  | [optional] 
**Target** | Pointer to [**Money**](Money.md) |  | [optional] 
**Rate** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewExchangeRate

`func NewExchangeRate() *ExchangeRate`

NewExchangeRate instantiates a new ExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateWithDefaults

`func NewExchangeRateWithDefaults() *ExchangeRate`

NewExchangeRateWithDefaults instantiates a new ExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ExchangeRate) GetSource() Money`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExchangeRate) GetSourceOk() (*Money, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExchangeRate) SetSource(v Money)`

SetSource sets Source field to given value.

### HasSource

`func (o *ExchangeRate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ExchangeRate) GetTarget() Money`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ExchangeRate) GetTargetOk() (*Money, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ExchangeRate) SetTarget(v Money)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ExchangeRate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetRate

`func (o *ExchangeRate) GetRate() Money`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeRate) GetRateOk() (*Money, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeRate) SetRate(v Money)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ExchangeRate) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


