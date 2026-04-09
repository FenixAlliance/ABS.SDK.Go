# AssetValueAmendCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**PreviousValue** | Pointer to **float64** |  | [optional] 
**NewValue** | Pointer to **float64** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**AmendmentDate** | Pointer to **time.Time** |  | [optional] 
**ApprovedBy** | Pointer to **NullableString** |  | [optional] 
**ApprovalDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAssetValueAmendCreateDto

`func NewAssetValueAmendCreateDto() *AssetValueAmendCreateDto`

NewAssetValueAmendCreateDto instantiates a new AssetValueAmendCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetValueAmendCreateDtoWithDefaults

`func NewAssetValueAmendCreateDtoWithDefaults() *AssetValueAmendCreateDto`

NewAssetValueAmendCreateDtoWithDefaults instantiates a new AssetValueAmendCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetValueAmendCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetValueAmendCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetValueAmendCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetValueAmendCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetValueAmendCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetValueAmendCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetValueAmendCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetValueAmendCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetValueAmendCreateDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetValueAmendCreateDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetValueAmendCreateDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetValueAmendCreateDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetValueAmendCreateDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetValueAmendCreateDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetPreviousValue

`func (o *AssetValueAmendCreateDto) GetPreviousValue() float64`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *AssetValueAmendCreateDto) GetPreviousValueOk() (*float64, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *AssetValueAmendCreateDto) SetPreviousValue(v float64)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *AssetValueAmendCreateDto) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### GetNewValue

`func (o *AssetValueAmendCreateDto) GetNewValue() float64`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AssetValueAmendCreateDto) GetNewValueOk() (*float64, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AssetValueAmendCreateDto) SetNewValue(v float64)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AssetValueAmendCreateDto) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetReason

`func (o *AssetValueAmendCreateDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssetValueAmendCreateDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssetValueAmendCreateDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AssetValueAmendCreateDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AssetValueAmendCreateDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AssetValueAmendCreateDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetAmendmentDate

`func (o *AssetValueAmendCreateDto) GetAmendmentDate() time.Time`

GetAmendmentDate returns the AmendmentDate field if non-nil, zero value otherwise.

### GetAmendmentDateOk

`func (o *AssetValueAmendCreateDto) GetAmendmentDateOk() (*time.Time, bool)`

GetAmendmentDateOk returns a tuple with the AmendmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentDate

`func (o *AssetValueAmendCreateDto) SetAmendmentDate(v time.Time)`

SetAmendmentDate sets AmendmentDate field to given value.

### HasAmendmentDate

`func (o *AssetValueAmendCreateDto) HasAmendmentDate() bool`

HasAmendmentDate returns a boolean if a field has been set.

### GetApprovedBy

`func (o *AssetValueAmendCreateDto) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *AssetValueAmendCreateDto) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *AssetValueAmendCreateDto) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *AssetValueAmendCreateDto) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### SetApprovedByNil

`func (o *AssetValueAmendCreateDto) SetApprovedByNil(b bool)`

 SetApprovedByNil sets the value for ApprovedBy to be an explicit nil

### UnsetApprovedBy
`func (o *AssetValueAmendCreateDto) UnsetApprovedBy()`

UnsetApprovedBy ensures that no value is present for ApprovedBy, not even an explicit nil
### GetApprovalDate

`func (o *AssetValueAmendCreateDto) GetApprovalDate() time.Time`

GetApprovalDate returns the ApprovalDate field if non-nil, zero value otherwise.

### GetApprovalDateOk

`func (o *AssetValueAmendCreateDto) GetApprovalDateOk() (*time.Time, bool)`

GetApprovalDateOk returns a tuple with the ApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalDate

`func (o *AssetValueAmendCreateDto) SetApprovalDate(v time.Time)`

SetApprovalDate sets ApprovalDate field to given value.

### HasApprovalDate

`func (o *AssetValueAmendCreateDto) HasApprovalDate() bool`

HasApprovalDate returns a boolean if a field has been set.

### SetApprovalDateNil

`func (o *AssetValueAmendCreateDto) SetApprovalDateNil(b bool)`

 SetApprovalDateNil sets the value for ApprovalDate to be an explicit nil

### UnsetApprovalDate
`func (o *AssetValueAmendCreateDto) UnsetApprovalDate()`

UnsetApprovalDate ensures that no value is present for ApprovalDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


