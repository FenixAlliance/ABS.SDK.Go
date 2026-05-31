# ShiftUpdateDto

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
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShiftUpdateDto

`func NewShiftUpdateDto() *ShiftUpdateDto`

NewShiftUpdateDto instantiates a new ShiftUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftUpdateDtoWithDefaults

`func NewShiftUpdateDtoWithDefaults() *ShiftUpdateDto`

NewShiftUpdateDtoWithDefaults instantiates a new ShiftUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ShiftUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShiftUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShiftUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShiftUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ShiftUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ShiftUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ShiftUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShiftUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShiftUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShiftUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShiftUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShiftUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *ShiftUpdateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ShiftUpdateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ShiftUpdateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ShiftUpdateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *ShiftUpdateDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *ShiftUpdateDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *ShiftUpdateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ShiftUpdateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ShiftUpdateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ShiftUpdateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ShiftUpdateDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ShiftUpdateDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIsBreak

`func (o *ShiftUpdateDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *ShiftUpdateDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *ShiftUpdateDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *ShiftUpdateDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### SetIsBreakNil

`func (o *ShiftUpdateDto) SetIsBreakNil(b bool)`

 SetIsBreakNil sets the value for IsBreak to be an explicit nil

### UnsetIsBreak
`func (o *ShiftUpdateDto) UnsetIsBreak()`

UnsetIsBreak ensures that no value is present for IsBreak, not even an explicit nil
### GetOccustOnMonday

`func (o *ShiftUpdateDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *ShiftUpdateDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *ShiftUpdateDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *ShiftUpdateDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### SetOccustOnMondayNil

`func (o *ShiftUpdateDto) SetOccustOnMondayNil(b bool)`

 SetOccustOnMondayNil sets the value for OccustOnMonday to be an explicit nil

### UnsetOccustOnMonday
`func (o *ShiftUpdateDto) UnsetOccustOnMonday()`

UnsetOccustOnMonday ensures that no value is present for OccustOnMonday, not even an explicit nil
### GetOccustOnTuesday

`func (o *ShiftUpdateDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *ShiftUpdateDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *ShiftUpdateDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *ShiftUpdateDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### SetOccustOnTuesdayNil

`func (o *ShiftUpdateDto) SetOccustOnTuesdayNil(b bool)`

 SetOccustOnTuesdayNil sets the value for OccustOnTuesday to be an explicit nil

### UnsetOccustOnTuesday
`func (o *ShiftUpdateDto) UnsetOccustOnTuesday()`

UnsetOccustOnTuesday ensures that no value is present for OccustOnTuesday, not even an explicit nil
### GetOccustOnWednesday

`func (o *ShiftUpdateDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *ShiftUpdateDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *ShiftUpdateDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *ShiftUpdateDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### SetOccustOnWednesdayNil

`func (o *ShiftUpdateDto) SetOccustOnWednesdayNil(b bool)`

 SetOccustOnWednesdayNil sets the value for OccustOnWednesday to be an explicit nil

### UnsetOccustOnWednesday
`func (o *ShiftUpdateDto) UnsetOccustOnWednesday()`

UnsetOccustOnWednesday ensures that no value is present for OccustOnWednesday, not even an explicit nil
### GetOccustOnThursday

`func (o *ShiftUpdateDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *ShiftUpdateDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *ShiftUpdateDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *ShiftUpdateDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### SetOccustOnThursdayNil

`func (o *ShiftUpdateDto) SetOccustOnThursdayNil(b bool)`

 SetOccustOnThursdayNil sets the value for OccustOnThursday to be an explicit nil

### UnsetOccustOnThursday
`func (o *ShiftUpdateDto) UnsetOccustOnThursday()`

UnsetOccustOnThursday ensures that no value is present for OccustOnThursday, not even an explicit nil
### GetOccustOnFriday

`func (o *ShiftUpdateDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *ShiftUpdateDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *ShiftUpdateDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *ShiftUpdateDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### SetOccustOnFridayNil

`func (o *ShiftUpdateDto) SetOccustOnFridayNil(b bool)`

 SetOccustOnFridayNil sets the value for OccustOnFriday to be an explicit nil

### UnsetOccustOnFriday
`func (o *ShiftUpdateDto) UnsetOccustOnFriday()`

UnsetOccustOnFriday ensures that no value is present for OccustOnFriday, not even an explicit nil
### GetOccustOnSaturday

`func (o *ShiftUpdateDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *ShiftUpdateDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *ShiftUpdateDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *ShiftUpdateDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### SetOccustOnSaturdayNil

`func (o *ShiftUpdateDto) SetOccustOnSaturdayNil(b bool)`

 SetOccustOnSaturdayNil sets the value for OccustOnSaturday to be an explicit nil

### UnsetOccustOnSaturday
`func (o *ShiftUpdateDto) UnsetOccustOnSaturday()`

UnsetOccustOnSaturday ensures that no value is present for OccustOnSaturday, not even an explicit nil
### GetOccustOnSunday

`func (o *ShiftUpdateDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *ShiftUpdateDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *ShiftUpdateDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *ShiftUpdateDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### SetOccustOnSundayNil

`func (o *ShiftUpdateDto) SetOccustOnSundayNil(b bool)`

 SetOccustOnSundayNil sets the value for OccustOnSunday to be an explicit nil

### UnsetOccustOnSunday
`func (o *ShiftUpdateDto) UnsetOccustOnSunday()`

UnsetOccustOnSunday ensures that no value is present for OccustOnSunday, not even an explicit nil
### GetRepeatEvery

`func (o *ShiftUpdateDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *ShiftUpdateDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *ShiftUpdateDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *ShiftUpdateDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### SetRepeatEveryNil

`func (o *ShiftUpdateDto) SetRepeatEveryNil(b bool)`

 SetRepeatEveryNil sets the value for RepeatEvery to be an explicit nil

### UnsetRepeatEvery
`func (o *ShiftUpdateDto) UnsetRepeatEvery()`

UnsetRepeatEvery ensures that no value is present for RepeatEvery, not even an explicit nil
### GetRepetitionCriteria

`func (o *ShiftUpdateDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *ShiftUpdateDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *ShiftUpdateDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *ShiftUpdateDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### SetRepetitionCriteriaNil

`func (o *ShiftUpdateDto) SetRepetitionCriteriaNil(b bool)`

 SetRepetitionCriteriaNil sets the value for RepetitionCriteria to be an explicit nil

### UnsetRepetitionCriteria
`func (o *ShiftUpdateDto) UnsetRepetitionCriteria()`

UnsetRepetitionCriteria ensures that no value is present for RepetitionCriteria, not even an explicit nil
### GetRecurrenceStart

`func (o *ShiftUpdateDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *ShiftUpdateDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *ShiftUpdateDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *ShiftUpdateDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *ShiftUpdateDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *ShiftUpdateDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *ShiftUpdateDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *ShiftUpdateDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *ShiftUpdateDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *ShiftUpdateDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *ShiftUpdateDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *ShiftUpdateDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *ShiftUpdateDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *ShiftUpdateDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *ShiftUpdateDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *ShiftUpdateDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### SetDayOfTheWeekNil

`func (o *ShiftUpdateDto) SetDayOfTheWeekNil(b bool)`

 SetDayOfTheWeekNil sets the value for DayOfTheWeek to be an explicit nil

### UnsetDayOfTheWeek
`func (o *ShiftUpdateDto) UnsetDayOfTheWeek()`

UnsetDayOfTheWeek ensures that no value is present for DayOfTheWeek, not even an explicit nil
### GetScheduleId

`func (o *ShiftUpdateDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ShiftUpdateDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ShiftUpdateDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ShiftUpdateDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ShiftUpdateDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ShiftUpdateDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *ShiftUpdateDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *ShiftUpdateDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *ShiftUpdateDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *ShiftUpdateDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *ShiftUpdateDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *ShiftUpdateDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetEmployeeProfileId

`func (o *ShiftUpdateDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *ShiftUpdateDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *ShiftUpdateDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *ShiftUpdateDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *ShiftUpdateDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *ShiftUpdateDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


