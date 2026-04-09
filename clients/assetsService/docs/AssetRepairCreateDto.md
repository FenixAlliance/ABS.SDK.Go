# AssetRepairCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**RepairStatus** | Pointer to **string** |  | [optional] 
**ScheduledDate** | Pointer to **time.Time** |  | [optional] 
**CompletionDate** | Pointer to **NullableTime** |  | [optional] 
**ReportedDate** | Pointer to **time.Time** |  | [optional] 
**EstimatedCost** | Pointer to **float64** |  | [optional] 
**ActualCost** | Pointer to **NullableFloat64** |  | [optional] 
**ProblemDescription** | Pointer to **NullableString** |  | [optional] 
**RepairDescription** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**AssetMaintenanceTeamId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetRepairCreateDto

`func NewAssetRepairCreateDto() *AssetRepairCreateDto`

NewAssetRepairCreateDto instantiates a new AssetRepairCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetRepairCreateDtoWithDefaults

`func NewAssetRepairCreateDtoWithDefaults() *AssetRepairCreateDto`

NewAssetRepairCreateDtoWithDefaults instantiates a new AssetRepairCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetRepairCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetRepairCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetRepairCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetRepairCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetRepairCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetRepairCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetRepairCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetRepairCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetRepairCreateDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetRepairCreateDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetRepairCreateDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetRepairCreateDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetRepairCreateDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetRepairCreateDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetRepairStatus

`func (o *AssetRepairCreateDto) GetRepairStatus() string`

GetRepairStatus returns the RepairStatus field if non-nil, zero value otherwise.

### GetRepairStatusOk

`func (o *AssetRepairCreateDto) GetRepairStatusOk() (*string, bool)`

GetRepairStatusOk returns a tuple with the RepairStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairStatus

`func (o *AssetRepairCreateDto) SetRepairStatus(v string)`

SetRepairStatus sets RepairStatus field to given value.

### HasRepairStatus

`func (o *AssetRepairCreateDto) HasRepairStatus() bool`

HasRepairStatus returns a boolean if a field has been set.

### GetScheduledDate

`func (o *AssetRepairCreateDto) GetScheduledDate() time.Time`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *AssetRepairCreateDto) GetScheduledDateOk() (*time.Time, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *AssetRepairCreateDto) SetScheduledDate(v time.Time)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *AssetRepairCreateDto) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *AssetRepairCreateDto) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AssetRepairCreateDto) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AssetRepairCreateDto) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AssetRepairCreateDto) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *AssetRepairCreateDto) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *AssetRepairCreateDto) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetReportedDate

`func (o *AssetRepairCreateDto) GetReportedDate() time.Time`

GetReportedDate returns the ReportedDate field if non-nil, zero value otherwise.

### GetReportedDateOk

`func (o *AssetRepairCreateDto) GetReportedDateOk() (*time.Time, bool)`

GetReportedDateOk returns a tuple with the ReportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDate

`func (o *AssetRepairCreateDto) SetReportedDate(v time.Time)`

SetReportedDate sets ReportedDate field to given value.

### HasReportedDate

`func (o *AssetRepairCreateDto) HasReportedDate() bool`

HasReportedDate returns a boolean if a field has been set.

### GetEstimatedCost

`func (o *AssetRepairCreateDto) GetEstimatedCost() float64`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *AssetRepairCreateDto) GetEstimatedCostOk() (*float64, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *AssetRepairCreateDto) SetEstimatedCost(v float64)`

SetEstimatedCost sets EstimatedCost field to given value.

### HasEstimatedCost

`func (o *AssetRepairCreateDto) HasEstimatedCost() bool`

HasEstimatedCost returns a boolean if a field has been set.

### GetActualCost

`func (o *AssetRepairCreateDto) GetActualCost() float64`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *AssetRepairCreateDto) GetActualCostOk() (*float64, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *AssetRepairCreateDto) SetActualCost(v float64)`

SetActualCost sets ActualCost field to given value.

### HasActualCost

`func (o *AssetRepairCreateDto) HasActualCost() bool`

HasActualCost returns a boolean if a field has been set.

### SetActualCostNil

`func (o *AssetRepairCreateDto) SetActualCostNil(b bool)`

 SetActualCostNil sets the value for ActualCost to be an explicit nil

### UnsetActualCost
`func (o *AssetRepairCreateDto) UnsetActualCost()`

UnsetActualCost ensures that no value is present for ActualCost, not even an explicit nil
### GetProblemDescription

`func (o *AssetRepairCreateDto) GetProblemDescription() string`

GetProblemDescription returns the ProblemDescription field if non-nil, zero value otherwise.

### GetProblemDescriptionOk

`func (o *AssetRepairCreateDto) GetProblemDescriptionOk() (*string, bool)`

GetProblemDescriptionOk returns a tuple with the ProblemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDescription

`func (o *AssetRepairCreateDto) SetProblemDescription(v string)`

SetProblemDescription sets ProblemDescription field to given value.

### HasProblemDescription

`func (o *AssetRepairCreateDto) HasProblemDescription() bool`

HasProblemDescription returns a boolean if a field has been set.

### SetProblemDescriptionNil

`func (o *AssetRepairCreateDto) SetProblemDescriptionNil(b bool)`

 SetProblemDescriptionNil sets the value for ProblemDescription to be an explicit nil

### UnsetProblemDescription
`func (o *AssetRepairCreateDto) UnsetProblemDescription()`

UnsetProblemDescription ensures that no value is present for ProblemDescription, not even an explicit nil
### GetRepairDescription

`func (o *AssetRepairCreateDto) GetRepairDescription() string`

GetRepairDescription returns the RepairDescription field if non-nil, zero value otherwise.

### GetRepairDescriptionOk

`func (o *AssetRepairCreateDto) GetRepairDescriptionOk() (*string, bool)`

GetRepairDescriptionOk returns a tuple with the RepairDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairDescription

`func (o *AssetRepairCreateDto) SetRepairDescription(v string)`

SetRepairDescription sets RepairDescription field to given value.

### HasRepairDescription

`func (o *AssetRepairCreateDto) HasRepairDescription() bool`

HasRepairDescription returns a boolean if a field has been set.

### SetRepairDescriptionNil

`func (o *AssetRepairCreateDto) SetRepairDescriptionNil(b bool)`

 SetRepairDescriptionNil sets the value for RepairDescription to be an explicit nil

### UnsetRepairDescription
`func (o *AssetRepairCreateDto) UnsetRepairDescription()`

UnsetRepairDescription ensures that no value is present for RepairDescription, not even an explicit nil
### GetNotes

`func (o *AssetRepairCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssetRepairCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssetRepairCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssetRepairCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AssetRepairCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AssetRepairCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetAssetMaintenanceTeamId

`func (o *AssetRepairCreateDto) GetAssetMaintenanceTeamId() string`

GetAssetMaintenanceTeamId returns the AssetMaintenanceTeamId field if non-nil, zero value otherwise.

### GetAssetMaintenanceTeamIdOk

`func (o *AssetRepairCreateDto) GetAssetMaintenanceTeamIdOk() (*string, bool)`

GetAssetMaintenanceTeamIdOk returns a tuple with the AssetMaintenanceTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMaintenanceTeamId

`func (o *AssetRepairCreateDto) SetAssetMaintenanceTeamId(v string)`

SetAssetMaintenanceTeamId sets AssetMaintenanceTeamId field to given value.

### HasAssetMaintenanceTeamId

`func (o *AssetRepairCreateDto) HasAssetMaintenanceTeamId() bool`

HasAssetMaintenanceTeamId returns a boolean if a field has been set.

### SetAssetMaintenanceTeamIdNil

`func (o *AssetRepairCreateDto) SetAssetMaintenanceTeamIdNil(b bool)`

 SetAssetMaintenanceTeamIdNil sets the value for AssetMaintenanceTeamId to be an explicit nil

### UnsetAssetMaintenanceTeamId
`func (o *AssetRepairCreateDto) UnsetAssetMaintenanceTeamId()`

UnsetAssetMaintenanceTeamId ensures that no value is present for AssetMaintenanceTeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


