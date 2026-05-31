# TrainingProgramEventDto

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
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TrainingProgramId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrainingProgramEventDto

`func NewTrainingProgramEventDto() *TrainingProgramEventDto`

NewTrainingProgramEventDto instantiates a new TrainingProgramEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingProgramEventDtoWithDefaults

`func NewTrainingProgramEventDtoWithDefaults() *TrainingProgramEventDto`

NewTrainingProgramEventDtoWithDefaults instantiates a new TrainingProgramEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainingProgramEventDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainingProgramEventDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainingProgramEventDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainingProgramEventDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TrainingProgramEventDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TrainingProgramEventDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TrainingProgramEventDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrainingProgramEventDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrainingProgramEventDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TrainingProgramEventDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TrainingProgramEventDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TrainingProgramEventDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *TrainingProgramEventDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrainingProgramEventDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrainingProgramEventDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrainingProgramEventDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrainingProgramEventDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrainingProgramEventDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TrainingProgramEventDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrainingProgramEventDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrainingProgramEventDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrainingProgramEventDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TrainingProgramEventDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TrainingProgramEventDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *TrainingProgramEventDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TrainingProgramEventDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TrainingProgramEventDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *TrainingProgramEventDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *TrainingProgramEventDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TrainingProgramEventDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TrainingProgramEventDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TrainingProgramEventDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIsBreak

`func (o *TrainingProgramEventDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *TrainingProgramEventDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *TrainingProgramEventDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *TrainingProgramEventDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### GetOccustOnMonday

`func (o *TrainingProgramEventDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *TrainingProgramEventDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *TrainingProgramEventDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *TrainingProgramEventDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### GetOccustOnTuesday

`func (o *TrainingProgramEventDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *TrainingProgramEventDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *TrainingProgramEventDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *TrainingProgramEventDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### GetOccustOnWednesday

`func (o *TrainingProgramEventDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *TrainingProgramEventDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *TrainingProgramEventDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *TrainingProgramEventDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### GetOccustOnThursday

`func (o *TrainingProgramEventDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *TrainingProgramEventDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *TrainingProgramEventDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *TrainingProgramEventDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### GetOccustOnFriday

`func (o *TrainingProgramEventDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *TrainingProgramEventDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *TrainingProgramEventDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *TrainingProgramEventDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### GetOccustOnSaturday

`func (o *TrainingProgramEventDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *TrainingProgramEventDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *TrainingProgramEventDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *TrainingProgramEventDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### GetOccustOnSunday

`func (o *TrainingProgramEventDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *TrainingProgramEventDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *TrainingProgramEventDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *TrainingProgramEventDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *TrainingProgramEventDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *TrainingProgramEventDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *TrainingProgramEventDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *TrainingProgramEventDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetRepetitionCriteria

`func (o *TrainingProgramEventDto) GetRepetitionCriteria() string`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *TrainingProgramEventDto) GetRepetitionCriteriaOk() (*string, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *TrainingProgramEventDto) SetRepetitionCriteria(v string)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *TrainingProgramEventDto) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetRecurrenceStart

`func (o *TrainingProgramEventDto) GetRecurrenceStart() time.Time`

GetRecurrenceStart returns the RecurrenceStart field if non-nil, zero value otherwise.

### GetRecurrenceStartOk

`func (o *TrainingProgramEventDto) GetRecurrenceStartOk() (*time.Time, bool)`

GetRecurrenceStartOk returns a tuple with the RecurrenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStart

`func (o *TrainingProgramEventDto) SetRecurrenceStart(v time.Time)`

SetRecurrenceStart sets RecurrenceStart field to given value.

### HasRecurrenceStart

`func (o *TrainingProgramEventDto) HasRecurrenceStart() bool`

HasRecurrenceStart returns a boolean if a field has been set.

### SetRecurrenceStartNil

`func (o *TrainingProgramEventDto) SetRecurrenceStartNil(b bool)`

 SetRecurrenceStartNil sets the value for RecurrenceStart to be an explicit nil

### UnsetRecurrenceStart
`func (o *TrainingProgramEventDto) UnsetRecurrenceStart()`

UnsetRecurrenceStart ensures that no value is present for RecurrenceStart, not even an explicit nil
### GetRecurrenceEnd

`func (o *TrainingProgramEventDto) GetRecurrenceEnd() time.Time`

GetRecurrenceEnd returns the RecurrenceEnd field if non-nil, zero value otherwise.

### GetRecurrenceEndOk

`func (o *TrainingProgramEventDto) GetRecurrenceEndOk() (*time.Time, bool)`

GetRecurrenceEndOk returns a tuple with the RecurrenceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceEnd

`func (o *TrainingProgramEventDto) SetRecurrenceEnd(v time.Time)`

SetRecurrenceEnd sets RecurrenceEnd field to given value.

### HasRecurrenceEnd

`func (o *TrainingProgramEventDto) HasRecurrenceEnd() bool`

HasRecurrenceEnd returns a boolean if a field has been set.

### SetRecurrenceEndNil

`func (o *TrainingProgramEventDto) SetRecurrenceEndNil(b bool)`

 SetRecurrenceEndNil sets the value for RecurrenceEnd to be an explicit nil

### UnsetRecurrenceEnd
`func (o *TrainingProgramEventDto) UnsetRecurrenceEnd()`

UnsetRecurrenceEnd ensures that no value is present for RecurrenceEnd, not even an explicit nil
### GetDayOfTheWeek

`func (o *TrainingProgramEventDto) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *TrainingProgramEventDto) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *TrainingProgramEventDto) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *TrainingProgramEventDto) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### GetScheduleId

`func (o *TrainingProgramEventDto) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *TrainingProgramEventDto) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *TrainingProgramEventDto) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *TrainingProgramEventDto) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *TrainingProgramEventDto) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *TrainingProgramEventDto) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetParentTimeIntervalId

`func (o *TrainingProgramEventDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *TrainingProgramEventDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *TrainingProgramEventDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *TrainingProgramEventDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *TrainingProgramEventDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *TrainingProgramEventDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil
### GetTenantId

`func (o *TrainingProgramEventDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TrainingProgramEventDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TrainingProgramEventDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TrainingProgramEventDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TrainingProgramEventDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TrainingProgramEventDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TrainingProgramEventDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TrainingProgramEventDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TrainingProgramEventDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TrainingProgramEventDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TrainingProgramEventDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TrainingProgramEventDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTrainingProgramId

`func (o *TrainingProgramEventDto) GetTrainingProgramId() string`

GetTrainingProgramId returns the TrainingProgramId field if non-nil, zero value otherwise.

### GetTrainingProgramIdOk

`func (o *TrainingProgramEventDto) GetTrainingProgramIdOk() (*string, bool)`

GetTrainingProgramIdOk returns a tuple with the TrainingProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingProgramId

`func (o *TrainingProgramEventDto) SetTrainingProgramId(v string)`

SetTrainingProgramId sets TrainingProgramId field to given value.

### HasTrainingProgramId

`func (o *TrainingProgramEventDto) HasTrainingProgramId() bool`

HasTrainingProgramId returns a boolean if a field has been set.

### SetTrainingProgramIdNil

`func (o *TrainingProgramEventDto) SetTrainingProgramIdNil(b bool)`

 SetTrainingProgramIdNil sets the value for TrainingProgramId to be an explicit nil

### UnsetTrainingProgramId
`func (o *TrainingProgramEventDto) UnsetTrainingProgramId()`

UnsetTrainingProgramId ensures that no value is present for TrainingProgramId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


