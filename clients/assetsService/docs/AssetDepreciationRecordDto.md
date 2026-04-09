# AssetDepreciationRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **map[string]interface{}** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
**AssetDepreciationPolicyId** | Pointer to **NullableString** |  | [optional] 
**AssetDepreciationPolicyName** | Pointer to **NullableString** |  | [optional] 
**DepreciationAmount** | Pointer to **float64** |  | [optional] 
**AccumulatedDepreciation** | Pointer to **float64** |  | [optional] 
**BookValue** | Pointer to **float64** |  | [optional] 
**DepreciationDate** | Pointer to **time.Time** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**Period** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetDepreciationRecordDto

`func NewAssetDepreciationRecordDto() *AssetDepreciationRecordDto`

NewAssetDepreciationRecordDto instantiates a new AssetDepreciationRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDepreciationRecordDtoWithDefaults

`func NewAssetDepreciationRecordDtoWithDefaults() *AssetDepreciationRecordDto`

NewAssetDepreciationRecordDtoWithDefaults instantiates a new AssetDepreciationRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetDepreciationRecordDto) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetDepreciationRecordDto) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetDepreciationRecordDto) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *AssetDepreciationRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetDepreciationRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetDepreciationRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetDepreciationRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetDepreciationRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetDepreciationRecordDto) GetBusinessId() map[string]interface{}`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetDepreciationRecordDto) GetBusinessIdOk() (*map[string]interface{}, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetDepreciationRecordDto) SetBusinessId(v map[string]interface{})`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetDepreciationRecordDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetBusinessProfileRecordId

`func (o *AssetDepreciationRecordDto) GetBusinessProfileRecordId() map[string]interface{}`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetDepreciationRecordDto) GetBusinessProfileRecordIdOk() (*map[string]interface{}, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetDepreciationRecordDto) SetBusinessProfileRecordId(v map[string]interface{})`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetDepreciationRecordDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetDepreciationRecordDto) GetAssetId() map[string]interface{}`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetDepreciationRecordDto) GetAssetIdOk() (*map[string]interface{}, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetDepreciationRecordDto) SetAssetId(v map[string]interface{})`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetDepreciationRecordDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetName

`func (o *AssetDepreciationRecordDto) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetDepreciationRecordDto) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetDepreciationRecordDto) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AssetDepreciationRecordDto) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### SetAssetNameNil

`func (o *AssetDepreciationRecordDto) SetAssetNameNil(b bool)`

 SetAssetNameNil sets the value for AssetName to be an explicit nil

### UnsetAssetName
`func (o *AssetDepreciationRecordDto) UnsetAssetName()`

UnsetAssetName ensures that no value is present for AssetName, not even an explicit nil
### GetAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordDto) GetAssetDepreciationPolicyId() string`

GetAssetDepreciationPolicyId returns the AssetDepreciationPolicyId field if non-nil, zero value otherwise.

### GetAssetDepreciationPolicyIdOk

`func (o *AssetDepreciationRecordDto) GetAssetDepreciationPolicyIdOk() (*string, bool)`

GetAssetDepreciationPolicyIdOk returns a tuple with the AssetDepreciationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordDto) SetAssetDepreciationPolicyId(v string)`

SetAssetDepreciationPolicyId sets AssetDepreciationPolicyId field to given value.

### HasAssetDepreciationPolicyId

`func (o *AssetDepreciationRecordDto) HasAssetDepreciationPolicyId() bool`

HasAssetDepreciationPolicyId returns a boolean if a field has been set.

### SetAssetDepreciationPolicyIdNil

`func (o *AssetDepreciationRecordDto) SetAssetDepreciationPolicyIdNil(b bool)`

 SetAssetDepreciationPolicyIdNil sets the value for AssetDepreciationPolicyId to be an explicit nil

### UnsetAssetDepreciationPolicyId
`func (o *AssetDepreciationRecordDto) UnsetAssetDepreciationPolicyId()`

UnsetAssetDepreciationPolicyId ensures that no value is present for AssetDepreciationPolicyId, not even an explicit nil
### GetAssetDepreciationPolicyName

`func (o *AssetDepreciationRecordDto) GetAssetDepreciationPolicyName() string`

GetAssetDepreciationPolicyName returns the AssetDepreciationPolicyName field if non-nil, zero value otherwise.

### GetAssetDepreciationPolicyNameOk

`func (o *AssetDepreciationRecordDto) GetAssetDepreciationPolicyNameOk() (*string, bool)`

GetAssetDepreciationPolicyNameOk returns a tuple with the AssetDepreciationPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDepreciationPolicyName

`func (o *AssetDepreciationRecordDto) SetAssetDepreciationPolicyName(v string)`

SetAssetDepreciationPolicyName sets AssetDepreciationPolicyName field to given value.

### HasAssetDepreciationPolicyName

`func (o *AssetDepreciationRecordDto) HasAssetDepreciationPolicyName() bool`

HasAssetDepreciationPolicyName returns a boolean if a field has been set.

### SetAssetDepreciationPolicyNameNil

`func (o *AssetDepreciationRecordDto) SetAssetDepreciationPolicyNameNil(b bool)`

 SetAssetDepreciationPolicyNameNil sets the value for AssetDepreciationPolicyName to be an explicit nil

### UnsetAssetDepreciationPolicyName
`func (o *AssetDepreciationRecordDto) UnsetAssetDepreciationPolicyName()`

UnsetAssetDepreciationPolicyName ensures that no value is present for AssetDepreciationPolicyName, not even an explicit nil
### GetDepreciationAmount

`func (o *AssetDepreciationRecordDto) GetDepreciationAmount() float64`

GetDepreciationAmount returns the DepreciationAmount field if non-nil, zero value otherwise.

### GetDepreciationAmountOk

`func (o *AssetDepreciationRecordDto) GetDepreciationAmountOk() (*float64, bool)`

GetDepreciationAmountOk returns a tuple with the DepreciationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAmount

`func (o *AssetDepreciationRecordDto) SetDepreciationAmount(v float64)`

SetDepreciationAmount sets DepreciationAmount field to given value.

### HasDepreciationAmount

`func (o *AssetDepreciationRecordDto) HasDepreciationAmount() bool`

HasDepreciationAmount returns a boolean if a field has been set.

### GetAccumulatedDepreciation

`func (o *AssetDepreciationRecordDto) GetAccumulatedDepreciation() float64`

GetAccumulatedDepreciation returns the AccumulatedDepreciation field if non-nil, zero value otherwise.

### GetAccumulatedDepreciationOk

`func (o *AssetDepreciationRecordDto) GetAccumulatedDepreciationOk() (*float64, bool)`

GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepreciation

`func (o *AssetDepreciationRecordDto) SetAccumulatedDepreciation(v float64)`

SetAccumulatedDepreciation sets AccumulatedDepreciation field to given value.

### HasAccumulatedDepreciation

`func (o *AssetDepreciationRecordDto) HasAccumulatedDepreciation() bool`

HasAccumulatedDepreciation returns a boolean if a field has been set.

### GetBookValue

`func (o *AssetDepreciationRecordDto) GetBookValue() float64`

GetBookValue returns the BookValue field if non-nil, zero value otherwise.

### GetBookValueOk

`func (o *AssetDepreciationRecordDto) GetBookValueOk() (*float64, bool)`

GetBookValueOk returns a tuple with the BookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValue

`func (o *AssetDepreciationRecordDto) SetBookValue(v float64)`

SetBookValue sets BookValue field to given value.

### HasBookValue

`func (o *AssetDepreciationRecordDto) HasBookValue() bool`

HasBookValue returns a boolean if a field has been set.

### GetDepreciationDate

`func (o *AssetDepreciationRecordDto) GetDepreciationDate() time.Time`

GetDepreciationDate returns the DepreciationDate field if non-nil, zero value otherwise.

### GetDepreciationDateOk

`func (o *AssetDepreciationRecordDto) GetDepreciationDateOk() (*time.Time, bool)`

GetDepreciationDateOk returns a tuple with the DepreciationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationDate

`func (o *AssetDepreciationRecordDto) SetDepreciationDate(v time.Time)`

SetDepreciationDate sets DepreciationDate field to given value.

### HasDepreciationDate

`func (o *AssetDepreciationRecordDto) HasDepreciationDate() bool`

HasDepreciationDate returns a boolean if a field has been set.

### GetYear

`func (o *AssetDepreciationRecordDto) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AssetDepreciationRecordDto) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AssetDepreciationRecordDto) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AssetDepreciationRecordDto) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *AssetDepreciationRecordDto) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *AssetDepreciationRecordDto) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *AssetDepreciationRecordDto) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *AssetDepreciationRecordDto) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetPeriod

`func (o *AssetDepreciationRecordDto) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AssetDepreciationRecordDto) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AssetDepreciationRecordDto) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AssetDepreciationRecordDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *AssetDepreciationRecordDto) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *AssetDepreciationRecordDto) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


