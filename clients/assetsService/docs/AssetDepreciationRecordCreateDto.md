# AssetDepreciationRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**AssetDepreciationPolicyId** | Pointer to **NullableString** |  | [optional] 
**DepreciationAmount** | Pointer to **float64** |  | [optional] 
**AccumulatedDepreciation** | Pointer to **float64** |  | [optional] 
**BookValue** | Pointer to **float64** |  | [optional] 
**DepreciationDate** | Pointer to **time.Time** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 

## Methods

### NewAssetDepreciationRecordCreateDto

`func NewAssetDepreciationRecordCreateDto() *AssetDepreciationRecordCreateDto`

NewAssetDepreciationRecordCreateDto instantiates a new AssetDepreciationRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDepreciationRecordCreateDtoWithDefaults

`func NewAssetDepreciationRecordCreateDtoWithDefaults() *AssetDepreciationRecordCreateDto`

NewAssetDepreciationRecordCreateDtoWithDefaults instantiates a new AssetDepreciationRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetDepreciationRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetDepreciationRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetDepreciationRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetDepreciationRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetDepreciationRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetDepreciationRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetDepreciationRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetDepreciationRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetDepreciationRecordCreateDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetDepreciationRecordCreateDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetDepreciationRecordCreateDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetDepreciationRecordCreateDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetDepreciationRecordCreateDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetDepreciationRecordCreateDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordCreateDto) GetAssetDepreciationPolicyId() string`

GetAssetDepreciationPolicyId returns the AssetDepreciationPolicyId field if non-nil, zero value otherwise.

### GetAssetDepreciationPolicyIdOk

`func (o *AssetDepreciationRecordCreateDto) GetAssetDepreciationPolicyIdOk() (*string, bool)`

GetAssetDepreciationPolicyIdOk returns a tuple with the AssetDepreciationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordCreateDto) SetAssetDepreciationPolicyId(v string)`

SetAssetDepreciationPolicyId sets AssetDepreciationPolicyId field to given value.

### HasAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordCreateDto) HasAssetDepreciationPolicyId() bool`

HasAssetDepreciationPolicyId returns a boolean if a field has been set.

### SetAssetDepreciationPolicyIdNil

`func (o *AssetDepreciationRecordCreateDto) SetAssetDepreciationPolicyIdNil(b bool)`

 SetAssetDepreciationPolicyIdNil sets the value for AssetDepreciationPolicyId to be an explicit nil

### UnsetAssetDepreciationPolicyId
`func (o *AssetDepreciationRecordCreateDto) UnsetAssetDepreciationPolicyId()`

UnsetAssetDepreciationPolicyId ensures that no value is present for AssetDepreciationPolicyId, not even an explicit nil
### GetDepreciationAmount

`func (o *AssetDepreciationRecordCreateDto) GetDepreciationAmount() float64`

GetDepreciationAmount returns the DepreciationAmount field if non-nil, zero value otherwise.

### GetDepreciationAmountOk

`func (o *AssetDepreciationRecordCreateDto) GetDepreciationAmountOk() (*float64, bool)`

GetDepreciationAmountOk returns a tuple with the DepreciationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAmount

`func (o *AssetDepreciationRecordCreateDto) SetDepreciationAmount(v float64)`

SetDepreciationAmount sets DepreciationAmount field to given value.

### HasDepreciationAmount

`func (o *AssetDepreciationRecordCreateDto) HasDepreciationAmount() bool`

HasDepreciationAmount returns a boolean if a field has been set.

### GetAccumulatedDepreciation

`func (o *AssetDepreciationRecordCreateDto) GetAccumulatedDepreciation() float64`

GetAccumulatedDepreciation returns the AccumulatedDepreciation field if non-nil, zero value otherwise.

### GetAccumulatedDepreciationOk

`func (o *AssetDepreciationRecordCreateDto) GetAccumulatedDepreciationOk() (*float64, bool)`

GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepreciation

`func (o *AssetDepreciationRecordCreateDto) SetAccumulatedDepreciation(v float64)`

SetAccumulatedDepreciation sets AccumulatedDepreciation field to given value.

### HasAccumulatedDepreciation

`func (o *AssetDepreciationRecordCreateDto) HasAccumulatedDepreciation() bool`

HasAccumulatedDepreciation returns a boolean if a field has been set.

### GetBookValue

`func (o *AssetDepreciationRecordCreateDto) GetBookValue() float64`

GetBookValue returns the BookValue field if non-nil, zero value otherwise.

### GetBookValueOk

`func (o *AssetDepreciationRecordCreateDto) GetBookValueOk() (*float64, bool)`

GetBookValueOk returns a tuple with the BookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValue

`func (o *AssetDepreciationRecordCreateDto) SetBookValue(v float64)`

SetBookValue sets BookValue field to given value.

### HasBookValue

`func (o *AssetDepreciationRecordCreateDto) HasBookValue() bool`

HasBookValue returns a boolean if a field has been set.

### GetDepreciationDate

`func (o *AssetDepreciationRecordCreateDto) GetDepreciationDate() time.Time`

GetDepreciationDate returns the DepreciationDate field if non-nil, zero value otherwise.

### GetDepreciationDateOk

`func (o *AssetDepreciationRecordCreateDto) GetDepreciationDateOk() (*time.Time, bool)`

GetDepreciationDateOk returns a tuple with the DepreciationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationDate

`func (o *AssetDepreciationRecordCreateDto) SetDepreciationDate(v time.Time)`

SetDepreciationDate sets DepreciationDate field to given value.

### HasDepreciationDate

`func (o *AssetDepreciationRecordCreateDto) HasDepreciationDate() bool`

HasDepreciationDate returns a boolean if a field has been set.

### GetYear

`func (o *AssetDepreciationRecordCreateDto) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AssetDepreciationRecordCreateDto) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AssetDepreciationRecordCreateDto) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AssetDepreciationRecordCreateDto) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *AssetDepreciationRecordCreateDto) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *AssetDepreciationRecordCreateDto) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *AssetDepreciationRecordCreateDto) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *AssetDepreciationRecordCreateDto) HasMonth() bool`

HasMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


