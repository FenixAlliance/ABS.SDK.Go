# TrainingProgramEventUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**IsBreak** | Pointer to **NullableBool** |  | [optional] 
**OccustOnMonday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnTuesday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnWednesday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnThursday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnFriday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnSaturday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnSunday** | Pointer to **NullableBool** |  | [optional] 
**RepeatEvery** | Pointer to **NullableInt32** |  | [optional] 
**RepetitionCriteria** | Pointer to **NullableString** |  | [optional] 
**RecurrenceStart** | Pointer to **NullableTime** |  | [optional] 
**RecurrenceEnd** | Pointer to **NullableTime** |  | [optional] 
**DayOfTheWeek** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**ParentTimeIntervalId** | Pointer to **NullableString** |  | [optional] 
**TrainingProgramId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrainingProgramEventUpdateDto

`func NewTrainingProgramEventUpdateDto() *TrainingProgramEventUpdateDto`

NewTrainingProgramEventUpdateDto instantiates a new TrainingProgramEventUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingProgramEventUpdateDtoWithDefaults

`func NewTrainingProgramEventUpdateDtoWithDefaults() *TrainingProgramEventUpdateDto`

NewTrainingProgramEventUpdateDtoWithDefaults instantiates a new TrainingProgramEventUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TrainingProgramEventUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrainingProgramEventUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrainingProgramEventUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrainingProgramEventUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrainingProgramEventUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrainingProgramEventUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TrainingProgramEventUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrainingProgramEventUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrainingProgramEventUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrainingProgramEventUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TrainingProgramEventUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TrainingProgramEventUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *TrainingProgramEventUpdateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TrainingProgramEventUpdateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TrainingProgramEventUpdateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *TrainingProgramEventUpdateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *TrainingProgramEventUpdateDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *TrainingProgramEventUpdateDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *TrainingProgramEventUpdateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TrainingProgramEventUpdateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TrainingProgramEventUpdateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TrainingProgramEventUpdateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *TrainingProgramEventUpdateDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *TrainingProgramEventUpdateDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIsBreak

`func (o *TrainingProgramEventUpdateDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *TrainingProgramEventUpdateDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *TrainingProgramEventUpdateDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *TrainingProgramEventUpdateDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### SetIsBreakNil

`func (o *TrainingProgramEventUpdateDto) SetIsBreakNil(b bool)`

 SetIsBreakNil sets the value for IsBreak to be an explicit nil

### UnsetIsBreak
`func (o *TrainingProgramEventUpdateDto) UnsetIsBreak()`

UnsetIsBreak ensures that no value is present for IsBreak, not even an explicit nil
### GetOccustOnMonday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### SetOccustOnMondayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnMondayNil(b bool)`

 SetOccustOnMondayNil sets the value for OccustOnMonday to be an explicit nil

### UnsetOccustOnMonday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnMonday()`

UnsetOccustOnMonday ensures that no value is present for OccustOnMonday, not even an explicit nil
### GetOccustOnTuesday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### SetOccustOnTuesdayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnTuesdayNil(b bool)`

 SetOccustOnTuesdayNil sets the value for OccustOnTuesday to be an explicit nil

### UnsetOccustOnTuesday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnTuesday()`

UnsetOccustOnTuesday ensures that no value is present for OccustOnTuesday, not even an explicit nil
### GetOccustOnWednesday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### SetOccustOnWednesdayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnWednesdayNil(b bool)`

 SetOccustOnWednesdayNil sets the value for OccustOnWednesday to be an explicit nil

### UnsetOccustOnWednesday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnWednesday()`

UnsetOccustOnWednesday ensures that no value is present for OccustOnWednesday, not even an explicit nil
### GetOccustOnThursday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### SetOccustOnThursdayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnThursdayNil(b bool)`

 SetOccustOnThursdayNil sets the value for OccustOnThursday to be an explicit nil

### UnsetOccustOnThursday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnThursday()`

UnsetOccustOnThursday ensures that no value is present for OccustOnThursday, not even an explicit nil
### GetOccustOnFriday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### SetOccustOnFridayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnFridayNil(b bool)`

 SetOccustOnFridayNil sets the value for OccustOnFriday to be an explicit nil

### UnsetOccustOnFriday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnFriday()`

UnsetOccustOnFriday ensures that no value is present for OccustOnFriday, not even an explicit nil
### GetOccustOnSaturday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### SetOccustOnSaturdayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnSaturdayNil(b bool)`

 SetOccustOnSaturdayNil sets the value for OccustOnSaturday to be an explicit nil

### UnsetOccustOnSaturday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnSaturday()`

UnsetOccustOnSaturday ensures that no value is present for OccustOnSaturday, not even an explicit nil
### GetOccustOnSunday

`func (o *TrainingProgramEventUpdateDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *TrainingProgramEventUpdateDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *TrainingProgramEventUpdateDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *TrainingProgramEventUpdateDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### SetOccustOnSundayNil

`func (o *TrainingProgramEventUpdateDto) SetOccustOnSundayNil(b bool)`

 SetOccustOnSundayNil sets the value for OccustOnSunday to be an explicit nil

### UnsetOccustOnSunday
`func (o *TrainingProgramEventUpdateDto) UnsetOccustOnSunday()`

UnsetOccustOnSunday ensures that no value is present for OccustOnSunday, not even an explicit nil
### GetRepeatEvery

`func (o *TrainingProgramEventUpdateDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *TrainingProgramEventUpdateDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *TrainingProgramEventUpdateDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *TrainingProgramEventUpdateDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### SetRepeatEveryNil

`func (o *TrainingProgramEventUpdateDto) SetRepeatEveryNil(b bool)`

 SetRepeatEveryNil sets the value for RepeatEvery to be an explicit nil

### UnsetRepeatEvery
`func (o *TrainingProgramEventUpdateDto) UnsetRepeatEvery()`

UnsetRepeatEvery ensures that no value is present for RepeatEvery, not even an explicit nil
### GetRepetitionCriteria

`func (o *TrainingProgramEventUpdateDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *TrainingProgramEventUpdateDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *TrainingProgramEventUpdateDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *TrainingProgramEventUpdateDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### SetRepetitionCriteriaNil

`func (o *TrainingProgramEventUpdateDto) SetRepetitionCriteriaNil(b bool)`

 SetRepetitionCriteriaNil sets the value for RepetitionCriteria to be an explicit nil

### UnsetRepetitionCriteria
`func (o *TrainingProgramEventUpdateDto) UnsetRepetitionCriteria()`

UnsetRepetitionCriteria ensures that no value is present for RepetitionCriteria, not even an explicit nil
### GetRecurrenceStart

`func (o *TrainingProgramEventUpdateDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *TrainingProgramEventUpdateDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *TrainingProgramEventUpdateDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *TrainingProgramEventUpdateDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *TrainingProgramEventUpdateDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *TrainingProgramEventUpdateDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *TrainingProgramEventUpdateDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *TrainingProgramEventUpdateDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *TrainingProgramEventUpdateDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *TrainingProgramEventUpdateDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *TrainingProgramEventUpdateDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *TrainingProgramEventUpdateDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *TrainingProgramEventUpdateDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *TrainingProgramEventUpdateDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *TrainingProgramEventUpdateDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *TrainingProgramEventUpdateDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### SetDayOfTheWeekNil

`func (o *TrainingProgramEventUpdateDto) SetDayOfTheWeekNil(b bool)`

 SetDayOfTheWeekNil sets the value for DayOfTheWeek to be an explicit nil

### UnsetDayOfTheWeek
`func (o *TrainingProgramEventUpdateDto) UnsetDayOfTheWeek()`

UnsetDayOfTheWeek ensures that no value is present for DayOfTheWeek, not even an explicit nil
### GetScheduleId

`func (o *TrainingProgramEventUpdateDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *TrainingProgramEventUpdateDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *TrainingProgramEventUpdateDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *TrainingProgramEventUpdateDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *TrainingProgramEventUpdateDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *TrainingProgramEventUpdateDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *TrainingProgramEventUpdateDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *TrainingProgramEventUpdateDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *TrainingProgramEventUpdateDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *TrainingProgramEventUpdateDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *TrainingProgramEventUpdateDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *TrainingProgramEventUpdateDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetTrainingProgramId

`func (o *TrainingProgramEventUpdateDto) GetTrainingProgramId() string`

GetTrainingProgramId returns the TrainingProgramId field if non-nil, zero value otherwise.

### GetTrainingProgramIdOk

`func (o *TrainingProgramEventUpdateDto) GetTrainingProgramIdOk() (*string, bool)`

GetTrainingProgramIdOk returns a tuple with the TrainingProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingProgramId

`func (o *TrainingProgramEventUpdateDto) SetTrainingProgramId(v string)`

SetTrainingProgramId sets TrainingProgramId field to given value.

### HasTrainingProgramId

`func (o *TrainingProgramEventUpdateDto) HasTrainingProgramId() bool`

HasTrainingProgramId returns a boolean if a field has been set.

### SetTrainingProgramIdNil

`func (o *TrainingProgramEventUpdateDto) SetTrainingProgramIdNil(b bool)`

 SetTrainingProgramIdNil sets the value for TrainingProgramId to be an explicit nil

### UnsetTrainingProgramId
`func (o *TrainingProgramEventUpdateDto) UnsetTrainingProgramId()`

UnsetTrainingProgramId ensures that no value is present for TrainingProgramId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


