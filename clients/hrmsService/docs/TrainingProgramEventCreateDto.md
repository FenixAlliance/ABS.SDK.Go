# TrainingProgramEventCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Start** | **time.Time** |  | 
**End** | **time.Time** |  | 
**IsBreak** | Pointer to **bool** |  | [optional] 
**OccustOnMonday** | Pointer to **bool** |  | [optional] 
**OccustOnTuesday** | Pointer to **bool** |  | [optional] 
**OccustOnWednesday** | Pointer to **bool** |  | [optional] 
**OccustOnThursday** | Pointer to **bool** |  | [optional] 
**OccustOnFriday** | Pointer to **bool** |  | [optional] 
**OccustOnSaturday** | Pointer to **bool** |  | [optional] 
**OccustOnSunday** | Pointer to **bool** |  | [optional] 
**RepeatEvery** | Pointer to **int32** |  | [optional] 
**RepetitionCriteria** | Pointer to **string** |  | [optional] 
**RecurrenceStart** | Pointer to **NullableTime** |  | [optional] 
**RecurrenceEnd** | Pointer to **NullableTime** |  | [optional] 
**DayOfTheWeek** | Pointer to **string** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**ParentTimeIntervalId** | Pointer to **NullableString** |  | [optional] 
**TrainingProgramId** | **string** |  | 

## Methods

### NewTrainingProgramEventCreateDto

`func NewTrainingProgramEventCreateDto(title string, start time.Time, end time.Time, trainingProgramId string, ) *TrainingProgramEventCreateDto`

NewTrainingProgramEventCreateDto instantiates a new TrainingProgramEventCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingProgramEventCreateDtoWithDefaults

`func NewTrainingProgramEventCreateDtoWithDefaults() *TrainingProgramEventCreateDto`

NewTrainingProgramEventCreateDtoWithDefaults instantiates a new TrainingProgramEventCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainingProgramEventCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainingProgramEventCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainingProgramEventCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainingProgramEventCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TrainingProgramEventCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrainingProgramEventCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrainingProgramEventCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TrainingProgramEventCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *TrainingProgramEventCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrainingProgramEventCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrainingProgramEventCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TrainingProgramEventCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrainingProgramEventCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrainingProgramEventCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrainingProgramEventCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TrainingProgramEventCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TrainingProgramEventCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *TrainingProgramEventCreateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TrainingProgramEventCreateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TrainingProgramEventCreateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TrainingProgramEventCreateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TrainingProgramEventCreateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TrainingProgramEventCreateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetIsBreak

`func (o *TrainingProgramEventCreateDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *TrainingProgramEventCreateDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *TrainingProgramEventCreateDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *TrainingProgramEventCreateDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### GetOccustOnMonday

`func (o *TrainingProgramEventCreateDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *TrainingProgramEventCreateDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *TrainingProgramEventCreateDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### GetOccustOnTuesday

`func (o *TrainingProgramEventCreateDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *TrainingProgramEventCreateDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *TrainingProgramEventCreateDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### GetOccustOnWednesday

`func (o *TrainingProgramEventCreateDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *TrainingProgramEventCreateDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *TrainingProgramEventCreateDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### GetOccustOnThursday

`func (o *TrainingProgramEventCreateDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *TrainingProgramEventCreateDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *TrainingProgramEventCreateDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### GetOccustOnFriday

`func (o *TrainingProgramEventCreateDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *TrainingProgramEventCreateDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *TrainingProgramEventCreateDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### GetOccustOnSaturday

`func (o *TrainingProgramEventCreateDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *TrainingProgramEventCreateDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *TrainingProgramEventCreateDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### GetOccustOnSunday

`func (o *TrainingProgramEventCreateDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *TrainingProgramEventCreateDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *TrainingProgramEventCreateDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *TrainingProgramEventCreateDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *TrainingProgramEventCreateDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *TrainingProgramEventCreateDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *TrainingProgramEventCreateDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *TrainingProgramEventCreateDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetRepetitionCriteria

`func (o *TrainingProgramEventCreateDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *TrainingProgramEventCreateDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *TrainingProgramEventCreateDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *TrainingProgramEventCreateDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetRecurrenceStart

`func (o *TrainingProgramEventCreateDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *TrainingProgramEventCreateDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *TrainingProgramEventCreateDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *TrainingProgramEventCreateDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *TrainingProgramEventCreateDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *TrainingProgramEventCreateDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *TrainingProgramEventCreateDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *TrainingProgramEventCreateDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *TrainingProgramEventCreateDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *TrainingProgramEventCreateDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *TrainingProgramEventCreateDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *TrainingProgramEventCreateDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *TrainingProgramEventCreateDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *TrainingProgramEventCreateDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *TrainingProgramEventCreateDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *TrainingProgramEventCreateDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### GetScheduleId

`func (o *TrainingProgramEventCreateDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *TrainingProgramEventCreateDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *TrainingProgramEventCreateDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *TrainingProgramEventCreateDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *TrainingProgramEventCreateDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *TrainingProgramEventCreateDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *TrainingProgramEventCreateDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *TrainingProgramEventCreateDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *TrainingProgramEventCreateDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *TrainingProgramEventCreateDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *TrainingProgramEventCreateDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *TrainingProgramEventCreateDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetTrainingProgramId

`func (o *TrainingProgramEventCreateDto) GetTrainingProgramId() string`

GetTrainingProgramId returns the TrainingProgramId field if non-nil, zero value otherwise.

### GetTrainingProgramIdOk

`func (o *TrainingProgramEventCreateDto) GetTrainingProgramIdOk() (*string, bool)`

GetTrainingProgramIdOk returns a tuple with the TrainingProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingProgramId

`func (o *TrainingProgramEventCreateDto) SetTrainingProgramId(v string)`

SetTrainingProgramId sets TrainingProgramId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


