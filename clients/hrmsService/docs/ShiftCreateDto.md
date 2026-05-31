# ShiftCreateDto

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
**EmployeeProfileId** | **string** |  | 

## Methods

### NewShiftCreateDto

`func NewShiftCreateDto(title string, start time.Time, end time.Time, employeeProfileId string, ) *ShiftCreateDto`

NewShiftCreateDto instantiates a new ShiftCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftCreateDtoWithDefaults

`func NewShiftCreateDtoWithDefaults() *ShiftCreateDto`

NewShiftCreateDtoWithDefaults instantiates a new ShiftCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShiftCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShiftCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShiftCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShiftCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShiftCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShiftCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShiftCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShiftCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ShiftCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShiftCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShiftCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ShiftCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShiftCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShiftCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShiftCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShiftCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShiftCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *ShiftCreateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ShiftCreateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ShiftCreateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *ShiftCreateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ShiftCreateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ShiftCreateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetIsBreak

`func (o *ShiftCreateDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *ShiftCreateDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *ShiftCreateDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *ShiftCreateDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### GetOccustOnMonday

`func (o *ShiftCreateDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *ShiftCreateDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *ShiftCreateDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *ShiftCreateDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### GetOccustOnTuesday

`func (o *ShiftCreateDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *ShiftCreateDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *ShiftCreateDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *ShiftCreateDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### GetOccustOnWednesday

`func (o *ShiftCreateDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *ShiftCreateDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *ShiftCreateDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *ShiftCreateDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### GetOccustOnThursday

`func (o *ShiftCreateDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *ShiftCreateDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *ShiftCreateDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *ShiftCreateDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### GetOccustOnFriday

`func (o *ShiftCreateDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *ShiftCreateDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *ShiftCreateDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *ShiftCreateDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### GetOccustOnSaturday

`func (o *ShiftCreateDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *ShiftCreateDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *ShiftCreateDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *ShiftCreateDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### GetOccustOnSunday

`func (o *ShiftCreateDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *ShiftCreateDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *ShiftCreateDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *ShiftCreateDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *ShiftCreateDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *ShiftCreateDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *ShiftCreateDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *ShiftCreateDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetRepetitionCriteria

`func (o *ShiftCreateDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *ShiftCreateDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *ShiftCreateDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *ShiftCreateDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetRecurrenceStart

`func (o *ShiftCreateDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *ShiftCreateDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *ShiftCreateDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *ShiftCreateDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *ShiftCreateDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *ShiftCreateDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *ShiftCreateDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *ShiftCreateDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *ShiftCreateDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *ShiftCreateDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *ShiftCreateDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *ShiftCreateDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *ShiftCreateDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *ShiftCreateDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *ShiftCreateDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *ShiftCreateDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### GetScheduleId

`func (o *ShiftCreateDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ShiftCreateDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ShiftCreateDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ShiftCreateDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ShiftCreateDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ShiftCreateDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *ShiftCreateDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *ShiftCreateDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *ShiftCreateDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *ShiftCreateDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *ShiftCreateDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *ShiftCreateDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetEmployeeProfileId

`func (o *ShiftCreateDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *ShiftCreateDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *ShiftCreateDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


