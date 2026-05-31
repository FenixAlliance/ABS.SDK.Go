# ScheduleCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Sunday** | Pointer to **bool** |  | [optional] 
**Monday** | Pointer to **bool** |  | [optional] 
**Tuesday** | Pointer to **bool** |  | [optional] 
**Wednesday** | Pointer to **bool** |  | [optional] 
**Thursday** | Pointer to **bool** |  | [optional] 
**Friday** | Pointer to **bool** |  | [optional] 
**Saturday** | Pointer to **bool** |  | [optional] 
**UniqueInterval** | Pointer to **bool** |  | [optional] 
**Is24x7Interval** | Pointer to **bool** |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**HolidayScheduleId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleCreateDto

`func NewScheduleCreateDto(name string, ) *ScheduleCreateDto`

NewScheduleCreateDto instantiates a new ScheduleCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreateDtoWithDefaults

`func NewScheduleCreateDtoWithDefaults() *ScheduleCreateDto`

NewScheduleCreateDtoWithDefaults instantiates a new ScheduleCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ScheduleCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ScheduleCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ScheduleCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ScheduleCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ScheduleCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScheduleCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *ScheduleCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ScheduleCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ScheduleCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ScheduleCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSunday

`func (o *ScheduleCreateDto) GetSunday() bool`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *ScheduleCreateDto) GetSundayOk() (*bool, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *ScheduleCreateDto) SetSunday(v bool)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *ScheduleCreateDto) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### GetMonday

`func (o *ScheduleCreateDto) GetMonday() bool`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *ScheduleCreateDto) GetMondayOk() (*bool, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *ScheduleCreateDto) SetMonday(v bool)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *ScheduleCreateDto) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *ScheduleCreateDto) GetTuesday() bool`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *ScheduleCreateDto) GetTuesdayOk() (*bool, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *ScheduleCreateDto) SetTuesday(v bool)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *ScheduleCreateDto) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *ScheduleCreateDto) GetWednesday() bool`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *ScheduleCreateDto) GetWednesdayOk() (*bool, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *ScheduleCreateDto) SetWednesday(v bool)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *ScheduleCreateDto) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *ScheduleCreateDto) GetThursday() bool`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *ScheduleCreateDto) GetThursdayOk() (*bool, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *ScheduleCreateDto) SetThursday(v bool)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *ScheduleCreateDto) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *ScheduleCreateDto) GetFriday() bool`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *ScheduleCreateDto) GetFridayOk() (*bool, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *ScheduleCreateDto) SetFriday(v bool)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *ScheduleCreateDto) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *ScheduleCreateDto) GetSaturday() bool`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *ScheduleCreateDto) GetSaturdayOk() (*bool, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *ScheduleCreateDto) SetSaturday(v bool)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *ScheduleCreateDto) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetUniqueInterval

`func (o *ScheduleCreateDto) GetUniqueInterval() bool`

GetUniqueInterval returns the UniqueInterval field if non-nil, zero value otherwise.

### GetUniqueIntervalOk

`func (o *ScheduleCreateDto) GetUniqueIntervalOk() (*bool, bool)`

GetUniqueIntervalOk returns a tuple with the UniqueInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueInterval

`func (o *ScheduleCreateDto) SetUniqueInterval(v bool)`

SetUniqueInterval sets UniqueInterval field to given value.

### HasUniqueInterval

`func (o *ScheduleCreateDto) HasUniqueInterval() bool`

HasUniqueInterval returns a boolean if a field has been set.

### GetIs24x7Interval

`func (o *ScheduleCreateDto) GetIs24x7Interval() bool`

GetIs24x7Interval returns the Is24x7Interval field if non-nil, zero value otherwise.

### GetIs24x7IntervalOk

`func (o *ScheduleCreateDto) GetIs24x7IntervalOk() (*bool, bool)`

GetIs24x7IntervalOk returns a tuple with the Is24x7Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs24x7Interval

`func (o *ScheduleCreateDto) SetIs24x7Interval(v bool)`

SetIs24x7Interval sets Is24x7Interval field to given value.

### HasIs24x7Interval

`func (o *ScheduleCreateDto) HasIs24x7Interval() bool`

HasIs24x7Interval returns a boolean if a field has been set.

### GetStart

`func (o *ScheduleCreateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ScheduleCreateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ScheduleCreateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ScheduleCreateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *ScheduleCreateDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *ScheduleCreateDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *ScheduleCreateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ScheduleCreateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ScheduleCreateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ScheduleCreateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ScheduleCreateDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ScheduleCreateDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetTimezoneId

`func (o *ScheduleCreateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ScheduleCreateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ScheduleCreateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ScheduleCreateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ScheduleCreateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ScheduleCreateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetFiscalYearId

`func (o *ScheduleCreateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *ScheduleCreateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *ScheduleCreateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *ScheduleCreateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *ScheduleCreateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *ScheduleCreateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetHolidayScheduleId

`func (o *ScheduleCreateDto) GetHolidayScheduleId() string`

GetHolidayScheduleId returns the HolidayScheduleId field if non-nil, zero value otherwise.

### GetHolidayScheduleIdOk

`func (o *ScheduleCreateDto) GetHolidayScheduleIdOk() (*string, bool)`

GetHolidayScheduleIdOk returns a tuple with the HolidayScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayScheduleId

`func (o *ScheduleCreateDto) SetHolidayScheduleId(v string)`

SetHolidayScheduleId sets HolidayScheduleId field to given value.

### HasHolidayScheduleId

`func (o *ScheduleCreateDto) HasHolidayScheduleId() bool`

HasHolidayScheduleId returns a boolean if a field has been set.

### SetHolidayScheduleIdNil

`func (o *ScheduleCreateDto) SetHolidayScheduleIdNil(b bool)`

 SetHolidayScheduleIdNil sets the value for HolidayScheduleId to be an explicit nil

### UnsetHolidayScheduleId
`func (o *ScheduleCreateDto) UnsetHolidayScheduleId()`

UnsetHolidayScheduleId ensures that no value is present for HolidayScheduleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


