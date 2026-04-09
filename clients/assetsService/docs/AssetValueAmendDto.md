# AssetValueAmendDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **map[string]interface{}** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
**PreviousValue** | Pointer to **float64** |  | [optional] 
**NewValue** | Pointer to **float64** |  | [optional] 
**AmendmentAmount** | Pointer to **float64** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**AmendmentDate** | Pointer to **time.Time** |  | [optional] 
**ApprovedBy** | Pointer to **NullableString** |  | [optional] 
**ApprovalDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAssetValueAmendDto

`func NewAssetValueAmendDto() *AssetValueAmendDto`

NewAssetValueAmendDto instantiates a new AssetValueAmendDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetValueAmendDtoWithDefaults

`func NewAssetValueAmendDtoWithDefaults() *AssetValueAmendDto`

NewAssetValueAmendDtoWithDefaults instantiates a new AssetValueAmendDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetValueAmendDto) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetValueAmendDto) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetValueAmendDto) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *AssetValueAmendDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetValueAmendDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetValueAmendDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetValueAmendDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetValueAmendDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetValueAmendDto) GetBusinessId() map[string]interface{}`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetValueAmendDto) GetBusinessIdOk() (*map[string]interface{}, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetValueAmendDto) SetBusinessId(v map[string]interface{})`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetValueAmendDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetBusinessProfileRecordId

`func (o *AssetValueAmendDto) GetBusinessProfileRecordId() map[string]interface{}`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetValueAmendDto) GetBusinessProfileRecordIdOk() (*map[string]interface{}, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetValueAmendDto) SetBusinessProfileRecordId(v map[string]interface{})`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetValueAmendDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetValueAmendDto) GetAssetId() map[string]interface{}`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetValueAmendDto) GetAssetIdOk() (*map[string]interface{}, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetValueAmendDto) SetAssetId(v map[string]interface{})`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetValueAmendDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetName

`func (o *AssetValueAmendDto) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetValueAmendDto) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetValueAmendDto) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AssetValueAmendDto) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### SetAssetNameNil

`func (o *AssetValueAmendDto) SetAssetNameNil(b bool)`

 SetAssetNameNil sets the value for AssetName to be an explicit nil

### UnsetAssetName
`func (o *AssetValueAmendDto) UnsetAssetName()`

UnsetAssetName ensures that no value is present for AssetName, not even an explicit nil
### GetPreviousValue

`func (o *AssetValueAmendDto) GetPreviousValue() float64`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *AssetValueAmendDto) GetPreviousValueOk() (*float64, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *AssetValueAmendDto) SetPreviousValue(v float64)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *AssetValueAmendDto) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### GetNewValue

`func (o *AssetValueAmendDto) GetNewValue() float64`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AssetValueAmendDto) GetNewValueOk() (*float64, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AssetValueAmendDto) SetNewValue(v float64)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AssetValueAmendDto) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetAmendmentAmount

`func (o *AssetValueAmendDto) GetAmendmentAmount() float64`

GetAmendmentAmount returns the AmendmentAmount field if non-nil, zero value otherwise.

### GetAmendmentAmountOk

`func (o *AssetValueAmendDto) GetAmendmentAmountOk() (*float64, bool)`

GetAmendmentAmountOk returns a tuple with the AmendmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentAmount

`func (o *AssetValueAmendDto) SetAmendmentAmount(v float64)`

SetAmendmentAmount sets AmendmentAmount field to given value.

### HasAmendmentAmount

`func (o *AssetValueAmendDto) HasAmendmentAmount() bool`

HasAmendmentAmount returns a boolean if a field has been set.

### GetReason

`func (o *AssetValueAmendDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssetValueAmendDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssetValueAmendDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AssetValueAmendDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AssetValueAmendDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AssetValueAmendDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetAmendmentDate

`func (o *AssetValueAmendDto) GetAmendmentDate() time.Time`

GetAmendmentDate returns the AmendmentDate field if non-nil, zero value otherwise.

### GetAmendmentDateOk

`func (o *AssetValueAmendDto) GetAmendmentDateOk() (*time.Time, bool)`

GetAmendmentDateOk returns a tuple with the AmendmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentDate

`func (o *AssetValueAmendDto) SetAmendmentDate(v time.Time)`

SetAmendmentDate sets AmendmentDate field to given value.

### HasAmendmentDate

`func (o *AssetValueAmendDto) HasAmendmentDate() bool`

HasAmendmentDate returns a boolean if a field has been set.

### GetApprovedBy

`func (o *AssetValueAmendDto) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *AssetValueAmendDto) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *AssetValueAmendDto) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *AssetValueAmendDto) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### SetApprovedByNil

`func (o *AssetValueAmendDto) SetApprovedByNil(b bool)`

 SetApprovedByNil sets the value for ApprovedBy to be an explicit nil

### UnsetApprovedBy
`func (o *AssetValueAmendDto) UnsetApprovedBy()`

UnsetApprovedBy ensures that no value is present for ApprovedBy, not even an explicit nil
### GetApprovalDate

`func (o *AssetValueAmendDto) GetApprovalDate() time.Time`

GetApprovalDate returns the ApprovalDate field if non-nil, zero value otherwise.

### GetApprovalDateOk

`func (o *AssetValueAmendDto) GetApprovalDateOk() (*time.Time, bool)`

GetApprovalDateOk returns a tuple with the ApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalDate

`func (o *AssetValueAmendDto) SetApprovalDate(v time.Time)`

SetApprovalDate sets ApprovalDate field to given value.

### HasApprovalDate

`func (o *AssetValueAmendDto) HasApprovalDate() bool`

HasApprovalDate returns a boolean if a field has been set.

### SetApprovalDateNil

`func (o *AssetValueAmendDto) SetApprovalDateNil(b bool)`

 SetApprovalDateNil sets the value for ApprovalDate to be an explicit nil

### UnsetApprovalDate
`func (o *AssetValueAmendDto) UnsetApprovalDate()`

UnsetApprovalDate ensures that no value is present for ApprovalDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


