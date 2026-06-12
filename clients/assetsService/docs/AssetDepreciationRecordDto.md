# AssetDepreciationRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
**AssetDepreciationPolicyId** | Pointer to **NullableString** |  | [optional] 
**AssetDepreciationPolicyName** | Pointer to **NullableString** |  | [optional] 
**FinancialBookId** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**TotalDepreciations** | Pointer to **int32** |  | [optional] 
**DepreciationFrequency** | Pointer to **int32** |  | [optional] 
**DepreciationRate** | Pointer to **float64** |  | [optional] 
**ExpectedValueAUL** | Pointer to **float64** |  | [optional] 
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

`func (o *AssetDepreciationRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetDepreciationRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetDepreciationRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetDepreciationRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssetDepreciationRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssetDepreciationRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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

### GetTenantId

`func (o *AssetDepreciationRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetDepreciationRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetDepreciationRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetDepreciationRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AssetDepreciationRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AssetDepreciationRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AssetDepreciationRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AssetDepreciationRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AssetDepreciationRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AssetDepreciationRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AssetDepreciationRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AssetDepreciationRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAssetId

`func (o *AssetDepreciationRecordDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetDepreciationRecordDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetDepreciationRecordDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetDepreciationRecordDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetDepreciationRecordDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetDepreciationRecordDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
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
### GetFinancialBookId

`func (o *AssetDepreciationRecordDto) GetFinancialBookId() string`

GetFinancialBookId returns the FinancialBookId field if non-nil, zero value otherwise.

### GetFinancialBookIdOk

`func (o *AssetDepreciationRecordDto) GetFinancialBookIdOk() (*string, bool)`

GetFinancialBookIdOk returns a tuple with the FinancialBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBookId

`func (o *AssetDepreciationRecordDto) SetFinancialBookId(v string)`

SetFinancialBookId sets FinancialBookId field to given value.

### HasFinancialBookId

`func (o *AssetDepreciationRecordDto) HasFinancialBookId() bool`

HasFinancialBookId returns a boolean if a field has been set.

### SetFinancialBookIdNil

`func (o *AssetDepreciationRecordDto) SetFinancialBookIdNil(b bool)`

 SetFinancialBookIdNil sets the value for FinancialBookId to be an explicit nil

### UnsetFinancialBookId
`func (o *AssetDepreciationRecordDto) UnsetFinancialBookId()`

UnsetFinancialBookId ensures that no value is present for FinancialBookId, not even an explicit nil
### GetStartDate

`func (o *AssetDepreciationRecordDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AssetDepreciationRecordDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AssetDepreciationRecordDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AssetDepreciationRecordDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalDepreciations

`func (o *AssetDepreciationRecordDto) GetTotalDepreciations() int32`

GetTotalDepreciations returns the TotalDepreciations field if non-nil, zero value otherwise.

### GetTotalDepreciationsOk

`func (o *AssetDepreciationRecordDto) GetTotalDepreciationsOk() (*int32, bool)`

GetTotalDepreciationsOk returns a tuple with the TotalDepreciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDepreciations

`func (o *AssetDepreciationRecordDto) SetTotalDepreciations(v int32)`

SetTotalDepreciations sets TotalDepreciations field to given value.

### HasTotalDepreciations

`func (o *AssetDepreciationRecordDto) HasTotalDepreciations() bool`

HasTotalDepreciations returns a boolean if a field has been set.

### GetDepreciationFrequency

`func (o *AssetDepreciationRecordDto) GetDepreciationFrequency() int32`

GetDepreciationFrequency returns the DepreciationFrequency field if non-nil, zero value otherwise.

### GetDepreciationFrequencyOk

`func (o *AssetDepreciationRecordDto) GetDepreciationFrequencyOk() (*int32, bool)`

GetDepreciationFrequencyOk returns a tuple with the DepreciationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationFrequency

`func (o *AssetDepreciationRecordDto) SetDepreciationFrequency(v int32)`

SetDepreciationFrequency sets DepreciationFrequency field to given value.

### HasDepreciationFrequency

`func (o *AssetDepreciationRecordDto) HasDepreciationFrequency() bool`

HasDepreciationFrequency returns a boolean if a field has been set.

### GetDepreciationRate

`func (o *AssetDepreciationRecordDto) GetDepreciationRate() float64`

GetDepreciationRate returns the DepreciationRate field if non-nil, zero value otherwise.

### GetDepreciationRateOk

`func (o *AssetDepreciationRecordDto) GetDepreciationRateOk() (*float64, bool)`

GetDepreciationRateOk returns a tuple with the DepreciationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationRate

`func (o *AssetDepreciationRecordDto) SetDepreciationRate(v float64)`

SetDepreciationRate sets DepreciationRate field to given value.

### HasDepreciationRate

`func (o *AssetDepreciationRecordDto) HasDepreciationRate() bool`

HasDepreciationRate returns a boolean if a field has been set.

### GetExpectedValueAUL

`func (o *AssetDepreciationRecordDto) GetExpectedValueAUL() float64`

GetExpectedValueAUL returns the ExpectedValueAUL field if non-nil, zero value otherwise.

### GetExpectedValueAULOk

`func (o *AssetDepreciationRecordDto) GetExpectedValueAULOk() (*float64, bool)`

GetExpectedValueAULOk returns a tuple with the ExpectedValueAUL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedValueAUL

`func (o *AssetDepreciationRecordDto) SetExpectedValueAUL(v float64)`

SetExpectedValueAUL sets ExpectedValueAUL field to given value.

### HasExpectedValueAUL

`func (o *AssetDepreciationRecordDto) HasExpectedValueAUL() bool`

HasExpectedValueAUL returns a boolean if a field has been set.

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


