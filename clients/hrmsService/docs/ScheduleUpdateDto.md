# ScheduleUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Sunday** | Pointer to **NullableBool** |  | [optional] 
**Monday** | Pointer to **NullableBool** |  | [optional] 
**Tuesday** | Pointer to **NullableBool** |  | [optional] 
**Wednesday** | Pointer to **NullableBool** |  | [optional] 
**Thursday** | Pointer to **NullableBool** |  | [optional] 
**Friday** | Pointer to **NullableBool** |  | [optional] 
**Saturday** | Pointer to **NullableBool** |  | [optional] 
**UniqueInterval** | Pointer to **NullableBool** |  | [optional] 
**Is24x7Interval** | Pointer to **NullableBool** |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**HolidayScheduleId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleUpdateDto

`func NewScheduleUpdateDto() *ScheduleUpdateDto`

NewScheduleUpdateDto instantiates a new ScheduleUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleUpdateDtoWithDefaults

`func NewScheduleUpdateDtoWithDefaults() *ScheduleUpdateDto`

NewScheduleUpdateDtoWithDefaults instantiates a new ScheduleUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScheduleUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ScheduleUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScheduleUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ScheduleUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *ScheduleUpdateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ScheduleUpdateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ScheduleUpdateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ScheduleUpdateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *ScheduleUpdateDto) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *ScheduleUpdateDto) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetSunday

`func (o *ScheduleUpdateDto) GetSunday() bool`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *ScheduleUpdateDto) GetSundayOk() (*bool, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *ScheduleUpdateDto) SetSunday(v bool)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *ScheduleUpdateDto) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### SetSundayNil

`func (o *ScheduleUpdateDto) SetSundayNil(b bool)`

 SetSundayNil sets the value for Sunday to be an explicit nil

### UnsetSunday
`func (o *ScheduleUpdateDto) UnsetSunday()`

UnsetSunday ensures that no value is present for Sunday, not even an explicit nil
### GetMonday

`func (o *ScheduleUpdateDto) GetMonday() bool`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *ScheduleUpdateDto) GetMondayOk() (*bool, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *ScheduleUpdateDto) SetMonday(v bool)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *ScheduleUpdateDto) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### SetMondayNil

`func (o *ScheduleUpdateDto) SetMondayNil(b bool)`

 SetMondayNil sets the value for Monday to be an explicit nil

### UnsetMonday
`func (o *ScheduleUpdateDto) UnsetMonday()`

UnsetMonday ensures that no value is present for Monday, not even an explicit nil
### GetTuesday

`func (o *ScheduleUpdateDto) GetTuesday() bool`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *ScheduleUpdateDto) GetTuesdayOk() (*bool, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *ScheduleUpdateDto) SetTuesday(v bool)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *ScheduleUpdateDto) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### SetTuesdayNil

`func (o *ScheduleUpdateDto) SetTuesdayNil(b bool)`

 SetTuesdayNil sets the value for Tuesday to be an explicit nil

### UnsetTuesday
`func (o *ScheduleUpdateDto) UnsetTuesday()`

UnsetTuesday ensures that no value is present for Tuesday, not even an explicit nil
### GetWednesday

`func (o *ScheduleUpdateDto) GetWednesday() bool`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *ScheduleUpdateDto) GetWednesdayOk() (*bool, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *ScheduleUpdateDto) SetWednesday(v bool)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *ScheduleUpdateDto) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### SetWednesdayNil

`func (o *ScheduleUpdateDto) SetWednesdayNil(b bool)`

 SetWednesdayNil sets the value for Wednesday to be an explicit nil

### UnsetWednesday
`func (o *ScheduleUpdateDto) UnsetWednesday()`

UnsetWednesday ensures that no value is present for Wednesday, not even an explicit nil
### GetThursday

`func (o *ScheduleUpdateDto) GetThursday() bool`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *ScheduleUpdateDto) GetThursdayOk() (*bool, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *ScheduleUpdateDto) SetThursday(v bool)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *ScheduleUpdateDto) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### SetThursdayNil

`func (o *ScheduleUpdateDto) SetThursdayNil(b bool)`

 SetThursdayNil sets the value for Thursday to be an explicit nil

### UnsetThursday
`func (o *ScheduleUpdateDto) UnsetThursday()`

UnsetThursday ensures that no value is present for Thursday, not even an explicit nil
### GetFriday

`func (o *ScheduleUpdateDto) GetFriday() bool`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *ScheduleUpdateDto) GetFridayOk() (*bool, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *ScheduleUpdateDto) SetFriday(v bool)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *ScheduleUpdateDto) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### SetFridayNil

`func (o *ScheduleUpdateDto) SetFridayNil(b bool)`

 SetFridayNil sets the value for Friday to be an explicit nil

### UnsetFriday
`func (o *ScheduleUpdateDto) UnsetFriday()`

UnsetFriday ensures that no value is present for Friday, not even an explicit nil
### GetSaturday

`func (o *ScheduleUpdateDto) GetSaturday() bool`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *ScheduleUpdateDto) GetSaturdayOk() (*bool, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *ScheduleUpdateDto) SetSaturday(v bool)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *ScheduleUpdateDto) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### SetSaturdayNil

`func (o *ScheduleUpdateDto) SetSaturdayNil(b bool)`

 SetSaturdayNil sets the value for Saturday to be an explicit nil

### UnsetSaturday
`func (o *ScheduleUpdateDto) UnsetSaturday()`

UnsetSaturday ensures that no value is present for Saturday, not even an explicit nil
### GetUniqueInterval

`func (o *ScheduleUpdateDto) GetUniqueInterval() bool`

GetUniqueInterval returns the UniqueInterval field if non-nil, zero value otherwise.

### GetUniqueIntervalOk

`func (o *ScheduleUpdateDto) GetUniqueIntervalOk() (*bool, bool)`

GetUniqueIntervalOk returns a tuple with the UniqueInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueInterval

`func (o *ScheduleUpdateDto) SetUniqueInterval(v bool)`

SetUniqueInterval sets UniqueInterval field to given value.

### HasUniqueInterval

`func (o *ScheduleUpdateDto) HasUniqueInterval() bool`

HasUniqueInterval returns a boolean if a field has been set.

### SetUniqueIntervalNil

`func (o *ScheduleUpdateDto) SetUniqueIntervalNil(b bool)`

 SetUniqueIntervalNil sets the value for UniqueInterval to be an explicit nil

### UnsetUniqueInterval
`func (o *ScheduleUpdateDto) UnsetUniqueInterval()`

UnsetUniqueInterval ensures that no value is present for UniqueInterval, not even an explicit nil
### GetIs24x7Interval

`func (o *ScheduleUpdateDto) GetIs24x7Interval() bool`

GetIs24x7Interval returns the Is24x7Interval field if non-nil, zero value otherwise.

### GetIs24x7IntervalOk

`func (o *ScheduleUpdateDto) GetIs24x7IntervalOk() (*bool, bool)`

GetIs24x7IntervalOk returns a tuple with the Is24x7Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs24x7Interval

`func (o *ScheduleUpdateDto) SetIs24x7Interval(v bool)`

SetIs24x7Interval sets Is24x7Interval field to given value.

### HasIs24x7Interval

`func (o *ScheduleUpdateDto) HasIs24x7Interval() bool`

HasIs24x7Interval returns a boolean if a field has been set.

### SetIs24x7IntervalNil

`func (o *ScheduleUpdateDto) SetIs24x7IntervalNil(b bool)`

 SetIs24x7IntervalNil sets the value for Is24x7Interval to be an explicit nil

### UnsetIs24x7Interval
`func (o *ScheduleUpdateDto) UnsetIs24x7Interval()`

UnsetIs24x7Interval ensures that no value is present for Is24x7Interval, not even an explicit nil
### GetStart

`func (o *ScheduleUpdateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ScheduleUpdateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ScheduleUpdateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ScheduleUpdateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *ScheduleUpdateDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *ScheduleUpdateDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *ScheduleUpdateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ScheduleUpdateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ScheduleUpdateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ScheduleUpdateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ScheduleUpdateDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ScheduleUpdateDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetTimezoneId

`func (o *ScheduleUpdateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ScheduleUpdateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ScheduleUpdateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ScheduleUpdateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ScheduleUpdateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ScheduleUpdateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetFiscalYearId

`func (o *ScheduleUpdateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *ScheduleUpdateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *ScheduleUpdateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *ScheduleUpdateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *ScheduleUpdateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *ScheduleUpdateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetHolidayScheduleId

`func (o *ScheduleUpdateDto) GetHolidayScheduleId() string`

GetHolidayScheduleId returns the HolidayScheduleId field if non-nil, zero value otherwise.

### GetHolidayScheduleIdOk

`func (o *ScheduleUpdateDto) GetHolidayScheduleIdOk() (*string, bool)`

GetHolidayScheduleIdOk returns a tuple with the HolidayScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayScheduleId

`func (o *ScheduleUpdateDto) SetHolidayScheduleId(v string)`

SetHolidayScheduleId sets HolidayScheduleId field to given value.

### HasHolidayScheduleId

`func (o *ScheduleUpdateDto) HasHolidayScheduleId() bool`

HasHolidayScheduleId returns a boolean if a field has been set.

### SetHolidayScheduleIdNil

`func (o *ScheduleUpdateDto) SetHolidayScheduleIdNil(b bool)`

 SetHolidayScheduleIdNil sets the value for HolidayScheduleId to be an explicit nil

### UnsetHolidayScheduleId
`func (o *ScheduleUpdateDto) UnsetHolidayScheduleId()`

UnsetHolidayScheduleId ensures that no value is present for HolidayScheduleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


