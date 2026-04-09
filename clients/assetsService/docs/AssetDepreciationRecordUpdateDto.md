# AssetDepreciationRecordUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepreciationAmount** | Pointer to **NullableFloat64** |  | [optional] 
**AccumulatedDepreciation** | Pointer to **NullableFloat64** |  | [optional] 
**BookValue** | Pointer to **NullableFloat64** |  | [optional] 
**DepreciationDate** | Pointer to **NullableTime** |  | [optional] 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Month** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAssetDepreciationRecordUpdateDto

`func NewAssetDepreciationRecordUpdateDto() *AssetDepreciationRecordUpdateDto`

NewAssetDepreciationRecordUpdateDto instantiates a new AssetDepreciationRecordUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDepreciationRecordUpdateDtoWithDefaults

`func NewAssetDepreciationRecordUpdateDtoWithDefaults() *AssetDepreciationRecordUpdateDto`

NewAssetDepreciationRecordUpdateDtoWithDefaults instantiates a new AssetDepreciationRecordUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepreciationAmount

`func (o *AssetDepreciationRecordUpdateDto) GetDepreciationAmount() float64`

GetDepreciationAmount returns the DepreciationAmount field if non-nil, zero value otherwise.

### GetDepreciationAmountOk

`func (o *AssetDepreciationRecordUpdateDto) GetDepreciationAmountOk() (*float64, bool)`

GetDepreciationAmountOk returns a tuple with the DepreciationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAmount

`func (o *AssetDepreciationRecordUpdateDto) SetDepreciationAmount(v float64)`

SetDepreciationAmount sets DepreciationAmount field to given value.

### HasDepreciationAmount

`func (o *AssetDepreciationRecordUpdateDto) HasDepreciationAmount() bool`

HasDepreciationAmount returns a boolean if a field has been set.

### SetDepreciationAmountNil

`func (o *AssetDepreciationRecordUpdateDto) SetDepreciationAmountNil(b bool)`

 SetDepreciationAmountNil sets the value for DepreciationAmount to be an explicit nil

### UnsetDepreciationAmount
`func (o *AssetDepreciationRecordUpdateDto) UnsetDepreciationAmount()`

UnsetDepreciationAmount ensures that no value is present for DepreciationAmount, not even an explicit nil
### GetAccumulatedDepreciation

`func (o *AssetDepreciationRecordUpdateDto) GetAccumulatedDepreciation() float64`

GetAccumulatedDepreciation returns the AccumulatedDepreciation field if non-nil, zero value otherwise.

### GetAccumulatedDepreciationOk

`func (o *AssetDepreciationRecordUpdateDto) GetAccumulatedDepreciationOk() (*float64, bool)`

GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepreciation

`func (o *AssetDepreciationRecordUpdateDto) SetAccumulatedDepreciation(v float64)`

SetAccumulatedDepreciation sets AccumulatedDepreciation field to given value.

### HasAccumulatedDepreciation

`func (o *AssetDepreciationRecordUpdateDto) HasAccumulatedDepreciation() bool`

HasAccumulatedDepreciation returns a boolean if a field has been set.

### SetAccumulatedDepreciationNil

`func (o *AssetDepreciationRecordUpdateDto) SetAccumulatedDepreciationNil(b bool)`

 SetAccumulatedDepreciationNil sets the value for AccumulatedDepreciation to be an explicit nil

### UnsetAccumulatedDepreciation
`func (o *AssetDepreciationRecordUpdateDto) UnsetAccumulatedDepreciation()`

UnsetAccumulatedDepreciation ensures that no value is present for AccumulatedDepreciation, not even an explicit nil
### GetBookValue

`func (o *AssetDepreciationRecordUpdateDto) GetBookValue() float64`

GetBookValue returns the BookValue field if non-nil, zero value otherwise.

### GetBookValueOk

`func (o *AssetDepreciationRecordUpdateDto) GetBookValueOk() (*float64, bool)`

GetBookValueOk returns a tuple with the BookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValue

`func (o *AssetDepreciationRecordUpdateDto) SetBookValue(v float64)`

SetBookValue sets BookValue field to given value.

### HasBookValue

`func (o *AssetDepreciationRecordUpdateDto) HasBookValue() bool`

HasBookValue returns a boolean if a field has been set.

### SetBookValueNil

`func (o *AssetDepreciationRecordUpdateDto) SetBookValueNil(b bool)`

 SetBookValueNil sets the value for BookValue to be an explicit nil

### UnsetBookValue
`func (o *AssetDepreciationRecordUpdateDto) UnsetBookValue()`

UnsetBookValue ensures that no value is present for BookValue, not even an explicit nil
### GetDepreciationDate

`func (o *AssetDepreciationRecordUpdateDto) GetDepreciationDate() time.Time`

GetDepreciationDate returns the DepreciationDate field if non-nil, zero value otherwise.

### GetDepreciationDateOk

`func (o *AssetDepreciationRecordUpdateDto) GetDepreciationDateOk() (*time.Time, bool)`

GetDepreciationDateOk returns a tuple with the DepreciationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationDate

`func (o *AssetDepreciationRecordUpdateDto) SetDepreciationDate(v time.Time)`

SetDepreciationDate sets DepreciationDate field to given value.

### HasDepreciationDate

`func (o *AssetDepreciationRecordUpdateDto) HasDepreciationDate() bool`

HasDepreciationDate returns a boolean if a field has been set.

### SetDepreciationDateNil

`func (o *AssetDepreciationRecordUpdateDto) SetDepreciationDateNil(b bool)`

 SetDepreciationDateNil sets the value for DepreciationDate to be an explicit nil

### UnsetDepreciationDate
`func (o *AssetDepreciationRecordUpdateDto) UnsetDepreciationDate()`

UnsetDepreciationDate ensures that no value is present for DepreciationDate, not even an explicit nil
### GetYear

`func (o *AssetDepreciationRecordUpdateDto) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AssetDepreciationRecordUpdateDto) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AssetDepreciationRecordUpdateDto) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AssetDepreciationRecordUpdateDto) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *AssetDepreciationRecordUpdateDto) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *AssetDepreciationRecordUpdateDto) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetMonth

`func (o *AssetDepreciationRecordUpdateDto) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *AssetDepreciationRecordUpdateDto) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *AssetDepreciationRecordUpdateDto) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *AssetDepreciationRecordUpdateDto) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonthNil

`func (o *AssetDepreciationRecordUpdateDto) SetMonthNil(b bool)`

 SetMonthNil sets the value for Month to be an explicit nil

### UnsetMonth
`func (o *AssetDepreciationRecordUpdateDto) UnsetMonth()`

UnsetMonth ensures that no value is present for Month, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


