# AssetRepairDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **map[string]interface{}** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetId** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
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
**AssetMaintenanceTeamName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetRepairDto

`func NewAssetRepairDto() *AssetRepairDto`

NewAssetRepairDto instantiates a new AssetRepairDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetRepairDtoWithDefaults

`func NewAssetRepairDtoWithDefaults() *AssetRepairDto`

NewAssetRepairDtoWithDefaults instantiates a new AssetRepairDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetRepairDto) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetRepairDto) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetRepairDto) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *AssetRepairDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetRepairDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetRepairDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetRepairDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetRepairDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetRepairDto) GetBusinessId() map[string]interface{}`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetRepairDto) GetBusinessIdOk() (*map[string]interface{}, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetRepairDto) SetBusinessId(v map[string]interface{})`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetRepairDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetBusinessProfileRecordId

`func (o *AssetRepairDto) GetBusinessProfileRecordId() map[string]interface{}`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetRepairDto) GetBusinessProfileRecordIdOk() (*map[string]interface{}, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetRepairDto) SetBusinessProfileRecordId(v map[string]interface{})`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetRepairDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetRepairDto) GetAssetId() map[string]interface{}`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetRepairDto) GetAssetIdOk() (*map[string]interface{}, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetRepairDto) SetAssetId(v map[string]interface{})`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetRepairDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetName

`func (o *AssetRepairDto) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetRepairDto) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetRepairDto) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AssetRepairDto) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### SetAssetNameNil

`func (o *AssetRepairDto) SetAssetNameNil(b bool)`

 SetAssetNameNil sets the value for AssetName to be an explicit nil

### UnsetAssetName
`func (o *AssetRepairDto) UnsetAssetName()`

UnsetAssetName ensures that no value is present for AssetName, not even an explicit nil
### GetRepairStatus

`func (o *AssetRepairDto) GetRepairStatus() string`

GetRepairStatus returns the RepairStatus field if non-nil, zero value otherwise.

### GetRepairStatusOk

`func (o *AssetRepairDto) GetRepairStatusOk() (*string, bool)`

GetRepairStatusOk returns a tuple with the RepairStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairStatus

`func (o *AssetRepairDto) SetRepairStatus(v string)`

SetRepairStatus sets RepairStatus field to given value.

### HasRepairStatus

`func (o *AssetRepairDto) HasRepairStatus() bool`

HasRepairStatus returns a boolean if a field has been set.

### GetScheduledDate

`func (o *AssetRepairDto) GetScheduledDate() time.Time`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *AssetRepairDto) GetScheduledDateOk() (*time.Time, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *AssetRepairDto) SetScheduledDate(v time.Time)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *AssetRepairDto) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *AssetRepairDto) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AssetRepairDto) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AssetRepairDto) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AssetRepairDto) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDateNil

`func (o *AssetRepairDto) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *AssetRepairDto) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetReportedDate

`func (o *AssetRepairDto) GetReportedDate() time.Time`

GetReportedDate returns the ReportedDate field if non-nil, zero value otherwise.

### GetReportedDateOk

`func (o *AssetRepairDto) GetReportedDateOk() (*time.Time, bool)`

GetReportedDateOk returns a tuple with the ReportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDate

`func (o *AssetRepairDto) SetReportedDate(v time.Time)`

SetReportedDate sets ReportedDate field to given value.

### HasReportedDate

`func (o *AssetRepairDto) HasReportedDate() bool`

HasReportedDate returns a boolean if a field has been set.

### GetEstimatedCost

`func (o *AssetRepairDto) GetEstimatedCost() float64`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *AssetRepairDto) GetEstimatedCostOk() (*float64, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *AssetRepairDto) SetEstimatedCost(v float64)`

SetEstimatedCost sets EstimatedCost field to given value.

### HasEstimatedCost

`func (o *AssetRepairDto) HasEstimatedCost() bool`

HasEstimatedCost returns a boolean if a field has been set.

### GetActualCost

`func (o *AssetRepairDto) GetActualCost() float64`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *AssetRepairDto) GetActualCostOk() (*float64, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *AssetRepairDto) SetActualCost(v float64)`

SetActualCost sets ActualCost field to given value.

### HasActualCost

`func (o *AssetRepairDto) HasActualCost() bool`

HasActualCost returns a boolean if a field has been set.

### SetActualCostNil

`func (o *AssetRepairDto) SetActualCostNil(b bool)`

 SetActualCostNil sets the value for ActualCost to be an explicit nil

### UnsetActualCost
`func (o *AssetRepairDto) UnsetActualCost()`

UnsetActualCost ensures that no value is present for ActualCost, not even an explicit nil
### GetProblemDescription

`func (o *AssetRepairDto) GetProblemDescription() string`

GetProblemDescription returns the ProblemDescription field if non-nil, zero value otherwise.

### GetProblemDescriptionOk

`func (o *AssetRepairDto) GetProblemDescriptionOk() (*string, bool)`

GetProblemDescriptionOk returns a tuple with the ProblemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDescription

`func (o *AssetRepairDto) SetProblemDescription(v string)`

SetProblemDescription sets ProblemDescription field to given value.

### HasProblemDescription

`func (o *AssetRepairDto) HasProblemDescription() bool`

HasProblemDescription returns a boolean if a field has been set.

### SetProblemDescriptionNil

`func (o *AssetRepairDto) SetProblemDescriptionNil(b bool)`

 SetProblemDescriptionNil sets the value for ProblemDescription to be an explicit nil

### UnsetProblemDescription
`func (o *AssetRepairDto) UnsetProblemDescription()`

UnsetProblemDescription ensures that no value is present for ProblemDescription, not even an explicit nil
### GetRepairDescription

`func (o *AssetRepairDto) GetRepairDescription() string`

GetRepairDescription returns the RepairDescription field if non-nil, zero value otherwise.

### GetRepairDescriptionOk

`func (o *AssetRepairDto) GetRepairDescriptionOk() (*string, bool)`

GetRepairDescriptionOk returns a tuple with the RepairDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairDescription

`func (o *AssetRepairDto) SetRepairDescription(v string)`

SetRepairDescription sets RepairDescription field to given value.

### HasRepairDescription

`func (o *AssetRepairDto) HasRepairDescription() bool`

HasRepairDescription returns a boolean if a field has been set.

### SetRepairDescriptionNil

`func (o *AssetRepairDto) SetRepairDescriptionNil(b bool)`

 SetRepairDescriptionNil sets the value for RepairDescription to be an explicit nil

### UnsetRepairDescription
`func (o *AssetRepairDto) UnsetRepairDescription()`

UnsetRepairDescription ensures that no value is present for RepairDescription, not even an explicit nil
### GetNotes

`func (o *AssetRepairDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssetRepairDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssetRepairDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssetRepairDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AssetRepairDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AssetRepairDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetAssetMaintenanceTeamId

`func (o *AssetRepairDto) GetAssetMaintenanceTeamId() string`

GetAssetMaintenanceTeamId returns the AssetMaintenanceTeamId field if non-nil, zero value otherwise.

### GetAssetMaintenanceTeamIdOk

`func (o *AssetRepairDto) GetAssetMaintenanceTeamIdOk() (*string, bool)`

GetAssetMaintenanceTeamIdOk returns a tuple with the AssetMaintenanceTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMaintenanceTeamId

`func (o *AssetRepairDto) SetAssetMaintenanceTeamId(v string)`

SetAssetMaintenanceTeamId sets AssetMaintenanceTeamId field to given value.

### HasAssetMaintenanceTeamId

`func (o *AssetRepairDto) HasAssetMaintenanceTeamId() bool`

HasAssetMaintenanceTeamId returns a boolean if a field has been set.

### SetAssetMaintenanceTeamIdNil

`func (o *AssetRepairDto) SetAssetMaintenanceTeamIdNil(b bool)`

 SetAssetMaintenanceTeamIdNil sets the value for AssetMaintenanceTeamId to be an explicit nil

### UnsetAssetMaintenanceTeamId
`func (o *AssetRepairDto) UnsetAssetMaintenanceTeamId()`

UnsetAssetMaintenanceTeamId ensures that no value is present for AssetMaintenanceTeamId, not even an explicit nil
### GetAssetMaintenanceTeamName

`func (o *AssetRepairDto) GetAssetMaintenanceTeamName() string`

GetAssetMaintenanceTeamName returns the AssetMaintenanceTeamName field if non-nil, zero value otherwise.

### GetAssetMaintenanceTeamNameOk

`func (o *AssetRepairDto) GetAssetMaintenanceTeamNameOk() (*string, bool)`

GetAssetMaintenanceTeamNameOk returns a tuple with the AssetMaintenanceTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMaintenanceTeamName

`func (o *AssetRepairDto) SetAssetMaintenanceTeamName(v string)`

SetAssetMaintenanceTeamName sets AssetMaintenanceTeamName field to given value.

### HasAssetMaintenanceTeamName

`func (o *AssetRepairDto) HasAssetMaintenanceTeamName() bool`

HasAssetMaintenanceTeamName returns a boolean if a field has been set.

### SetAssetMaintenanceTeamNameNil

`func (o *AssetRepairDto) SetAssetMaintenanceTeamNameNil(b bool)`

 SetAssetMaintenanceTeamNameNil sets the value for AssetMaintenanceTeamName to be an explicit nil

### UnsetAssetMaintenanceTeamName
`func (o *AssetRepairDto) UnsetAssetMaintenanceTeamName()`

UnsetAssetMaintenanceTeamName ensures that no value is present for AssetMaintenanceTeamName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


