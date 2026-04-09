# AssetRepairUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepairStatus** | Pointer to **NullableString** |  | [optional] 
**ScheduledDate** | Pointer to **NullableTime** |  | [optional] 
**CompletionDate** | Pointer to **NullableTime** |  | [optional] 
**EstimatedCost** | Pointer to **NullableFloat64** |  | [optional] 
**ActualCost** | Pointer to **NullableFloat64** |  | [optional] 
**ProblemDescription** | Pointer to **NullableString** |  | [optional] 
**RepairDescription** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**AssetMaintenanceTeamId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetRepairUpdateDto

`func NewAssetRepairUpdateDto() *AssetRepairUpdateDto`

NewAssetRepairUpdateDto instantiates a new AssetRepairUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetRepairUpdateDtoWithDefaults

`func NewAssetRepairUpdateDtoWithDefaults() *AssetRepairUpdateDto`

NewAssetRepairUpdateDtoWithDefaults instantiates a new AssetRepairUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepairStatus

`func (o *AssetRepairUpdateDto) GetRepairStatus() string`

GetRepairStatus returns the RepairStatus field if non-nil, zero value otherwise.

### GetRepairStatusOk

`func (o *AssetRepairUpdateDto) GetRepairStatusOk() (*string, bool)`

GetRepairStatusOk returns a tuple with the RepairStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairStatus

`func (o *AssetRepairUpdateDto) SetRepairStatus(v string)`

SetRepairStatus sets RepairStatus field to given value.

### HasRepairStatus

`func (o *AssetRepairUpdateDto) HasRepairStatus() bool`

HasRepairStatus returns a boolean if a field has been set.

### SetRepairStatusNil

`func (o *AssetRepairUpdateDto) SetRepairStatusNil(b bool)`

 SetRepairStatusNil sets the value for RepairStatus to be an explicit nil

### UnsetRepairStatus
`func (o *AssetRepairUpdateDto) UnsetRepairStatus()`

UnsetRepairStatus ensures that no value is present for RepairStatus, not even an explicit nil
### GetScheduledDate

`func (o *AssetRepairUpdateDto) GetScheduledDate() time.Time`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *AssetRepairUpdateDto) GetScheduledDateOk() (*time.Time, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *AssetRepairUpdateDto) SetScheduledDate(v time.Time)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *AssetRepairUpdateDto) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### SetScheduledDateNil

`func (o *AssetRepairUpdateDto) SetScheduledDateNil(b bool)`

 SetScheduledDateNil sets the value for ScheduledDate to be an explicit nil

### UnsetScheduledDate
`func (o *AssetRepairUpdateDto) UnsetScheduledDate()`

UnsetScheduledDate ensures that no value is present for ScheduledDate, not even an explicit nil
### GetCompletionDate

`func (o *AssetRepairUpdateDto) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AssetRepairUpdateDto) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AssetRepairUpdateDto) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AssetRepairUpdateDto) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *AssetRepairUpdateDto) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *AssetRepairUpdateDto) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetEstimatedCost

`func (o *AssetRepairUpdateDto) GetEstimatedCost() float64`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *AssetRepairUpdateDto) GetEstimatedCostOk() (*float64, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *AssetRepairUpdateDto) SetEstimatedCost(v float64)`

SetEstimatedCost sets EstimatedCost field to given value.

### HasEstimatedCost

`func (o *AssetRepairUpdateDto) HasEstimatedCost() bool`

HasEstimatedCost returns a boolean if a field has been set.

### SetEstimatedCostNil

`func (o *AssetRepairUpdateDto) SetEstimatedCostNil(b bool)`

 SetEstimatedCostNil sets the value for EstimatedCost to be an explicit nil

### UnsetEstimatedCost
`func (o *AssetRepairUpdateDto) UnsetEstimatedCost()`

UnsetEstimatedCost ensures that no value is present for EstimatedCost, not even an explicit nil
### GetActualCost

`func (o *AssetRepairUpdateDto) GetActualCost() float64`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *AssetRepairUpdateDto) GetActualCostOk() (*float64, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *AssetRepairUpdateDto) SetActualCost(v float64)`

SetActualCost sets ActualCost field to given value.

### HasActualCost

`func (o *AssetRepairUpdateDto) HasActualCost() bool`

HasActualCost returns a boolean if a field has been set.

### SetActualCostNil

`func (o *AssetRepairUpdateDto) SetActualCostNil(b bool)`

 SetActualCostNil sets the value for ActualCost to be an explicit nil

### UnsetActualCost
`func (o *AssetRepairUpdateDto) UnsetActualCost()`

UnsetActualCost ensures that no value is present for ActualCost, not even an explicit nil
### GetProblemDescription

`func (o *AssetRepairUpdateDto) GetProblemDescription() string`

GetProblemDescription returns the ProblemDescription field if non-nil, zero value otherwise.

### GetProblemDescriptionOk

`func (o *AssetRepairUpdateDto) GetProblemDescriptionOk() (*string, bool)`

GetProblemDescriptionOk returns a tuple with the ProblemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDescription

`func (o *AssetRepairUpdateDto) SetProblemDescription(v string)`

SetProblemDescription sets ProblemDescription field to given value.

### HasProblemDescription

`func (o *AssetRepairUpdateDto) HasProblemDescription() bool`

HasProblemDescription returns a boolean if a field has been set.

### SetProblemDescriptionNil

`func (o *AssetRepairUpdateDto) SetProblemDescriptionNil(b bool)`

 SetProblemDescriptionNil sets the value for ProblemDescription to be an explicit nil

### UnsetProblemDescription
`func (o *AssetRepairUpdateDto) UnsetProblemDescription()`

UnsetProblemDescription ensures that no value is present for ProblemDescription, not even an explicit nil
### GetRepairDescription

`func (o *AssetRepairUpdateDto) GetRepairDescription() string`

GetRepairDescription returns the RepairDescription field if non-nil, zero value otherwise.

### GetRepairDescriptionOk

`func (o *AssetRepairUpdateDto) GetRepairDescriptionOk() (*string, bool)`

GetRepairDescriptionOk returns a tuple with the RepairDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairDescription

`func (o *AssetRepairUpdateDto) SetRepairDescription(v string)`

SetRepairDescription sets RepairDescription field to given value.

### HasRepairDescription

`func (o *AssetRepairUpdateDto) HasRepairDescription() bool`

HasRepairDescription returns a boolean if a field has been set.

### SetRepairDescriptionNil

`func (o *AssetRepairUpdateDto) SetRepairDescriptionNil(b bool)`

 SetRepairDescriptionNil sets the value for RepairDescription to be an explicit nil

### UnsetRepairDescription
`func (o *AssetRepairUpdateDto) UnsetRepairDescription()`

UnsetRepairDescription ensures that no value is present for RepairDescription, not even an explicit nil
### GetNotes

`func (o *AssetRepairUpdateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssetRepairUpdateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssetRepairUpdateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssetRepairUpdateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AssetRepairUpdateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AssetRepairUpdateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetAssetMaintenanceTeamId

`func (o *AssetRepairUpdateDto) GetAssetMaintenanceTeamId() string`

GetAssetMaintenanceTeamId returns the AssetMaintenanceTeamId field if non-nil, zero value otherwise.

### GetAssetMaintenanceTeamIdOk

`func (o *AssetRepairUpdateDto) GetAssetMaintenanceTeamIdOk() (*string, bool)`

GetAssetMaintenanceTeamIdOk returns a tuple with the AssetMaintenanceTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMaintenanceTeamId

`func (o *AssetRepairUpdateDto) SetAssetMaintenanceTeamId(v string)`

SetAssetMaintenanceTeamId sets AssetMaintenanceTeamId field to given value.

### HasAssetMaintenanceTeamId

`func (o *AssetRepairUpdateDto) HasAssetMaintenanceTeamId() bool`

HasAssetMaintenanceTeamId returns a boolean if a field has been set.

### SetAssetMaintenanceTeamIdNil

`func (o *AssetRepairUpdateDto) SetAssetMaintenanceTeamIdNil(b bool)`

 SetAssetMaintenanceTeamIdNil sets the value for AssetMaintenanceTeamId to be an explicit nil

### UnsetAssetMaintenanceTeamId
`func (o *AssetRepairUpdateDto) UnsetAssetMaintenanceTeamId()`

UnsetAssetMaintenanceTeamId ensures that no value is present for AssetMaintenanceTeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


