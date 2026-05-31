# ShiftDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EmployeeProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShiftDto

`func NewShiftDto() *ShiftDto`

NewShiftDto instantiates a new ShiftDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftDtoWithDefaults

`func NewShiftDtoWithDefaults() *ShiftDto`

NewShiftDtoWithDefaults instantiates a new ShiftDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShiftDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShiftDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShiftDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShiftDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShiftDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShiftDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ShiftDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShiftDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShiftDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShiftDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ShiftDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ShiftDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ShiftDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShiftDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShiftDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShiftDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ShiftDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ShiftDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ShiftDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShiftDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShiftDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShiftDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShiftDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShiftDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *ShiftDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ShiftDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ShiftDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ShiftDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ShiftDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ShiftDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ShiftDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ShiftDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIsBreak

`func (o *ShiftDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *ShiftDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *ShiftDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *ShiftDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### GetOccustOnMonday

`func (o *ShiftDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *ShiftDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *ShiftDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *ShiftDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### GetOccustOnTuesday

`func (o *ShiftDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *ShiftDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *ShiftDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *ShiftDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### GetOccustOnWednesday

`func (o *ShiftDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *ShiftDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *ShiftDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *ShiftDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### GetOccustOnThursday

`func (o *ShiftDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *ShiftDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *ShiftDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *ShiftDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### GetOccustOnFriday

`func (o *ShiftDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *ShiftDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *ShiftDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *ShiftDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### GetOccustOnSaturday

`func (o *ShiftDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *ShiftDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *ShiftDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *ShiftDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### GetOccustOnSunday

`func (o *ShiftDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *ShiftDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *ShiftDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *ShiftDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *ShiftDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *ShiftDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *ShiftDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *ShiftDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetRepetitionCriteria

`func (o *ShiftDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *ShiftDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *ShiftDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *ShiftDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetRecurrenceStart

`func (o *ShiftDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *ShiftDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *ShiftDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *ShiftDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *ShiftDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *ShiftDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *ShiftDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *ShiftDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *ShiftDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *ShiftDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *ShiftDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *ShiftDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *ShiftDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *ShiftDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *ShiftDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *ShiftDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### GetScheduleId

`func (o *ShiftDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ShiftDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ShiftDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ShiftDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ShiftDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ShiftDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *ShiftDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *ShiftDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *ShiftDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *ShiftDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *ShiftDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *ShiftDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetTenantId

`func (o *ShiftDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShiftDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShiftDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShiftDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShiftDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShiftDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEmployeeProfileId

`func (o *ShiftDto) GetEmployeeProfileId() string`

GetEmployeeProfileId returns the EmployeeProfileId field if non-nil, zero value otherwise.

### GetEmployeeProfileIdOk

`func (o *ShiftDto) GetEmployeeProfileIdOk() (*string, bool)`

GetEmployeeProfileIdOk returns a tuple with the EmployeeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeProfileId

`func (o *ShiftDto) SetEmployeeProfileId(v string)`

SetEmployeeProfileId sets EmployeeProfileId field to given value.

### HasEmployeeProfileId

`func (o *ShiftDto) HasEmployeeProfileId() bool`

HasEmployeeProfileId returns a boolean if a field has been set.

### SetEmployeeProfileIdNil

`func (o *ShiftDto) SetEmployeeProfileIdNil(b bool)`

 SetEmployeeProfileIdNil sets the value for EmployeeProfileId to be an explicit nil

### UnsetEmployeeProfileId
`func (o *ShiftDto) UnsetEmployeeProfileId()`

UnsetEmployeeProfileId ensures that no value is present for EmployeeProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


