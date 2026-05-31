# ScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
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

### NewScheduleDto

`func NewScheduleDto() *ScheduleDto`

NewScheduleDto instantiates a new ScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDtoWithDefaults

`func NewScheduleDtoWithDefaults() *ScheduleDto`

NewScheduleDtoWithDefaults instantiates a new ScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ScheduleDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ScheduleDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ScheduleDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ScheduleDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ScheduleDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ScheduleDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ScheduleDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ScheduleDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *ScheduleDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ScheduleDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ScheduleDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ScheduleDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ScheduleDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ScheduleDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetName

`func (o *ScheduleDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ScheduleDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScheduleDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ScheduleDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *ScheduleDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ScheduleDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ScheduleDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ScheduleDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSunday

`func (o *ScheduleDto) GetSunday() bool`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *ScheduleDto) GetSundayOk() (*bool, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *ScheduleDto) SetSunday(v bool)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *ScheduleDto) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### GetMonday

`func (o *ScheduleDto) GetMonday() bool`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *ScheduleDto) GetMondayOk() (*bool, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *ScheduleDto) SetMonday(v bool)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *ScheduleDto) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *ScheduleDto) GetTuesday() bool`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *ScheduleDto) GetTuesdayOk() (*bool, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *ScheduleDto) SetTuesday(v bool)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *ScheduleDto) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *ScheduleDto) GetWednesday() bool`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *ScheduleDto) GetWednesdayOk() (*bool, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *ScheduleDto) SetWednesday(v bool)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *ScheduleDto) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *ScheduleDto) GetThursday() bool`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *ScheduleDto) GetThursdayOk() (*bool, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *ScheduleDto) SetThursday(v bool)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *ScheduleDto) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *ScheduleDto) GetFriday() bool`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *ScheduleDto) GetFridayOk() (*bool, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *ScheduleDto) SetFriday(v bool)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *ScheduleDto) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *ScheduleDto) GetSaturday() bool`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *ScheduleDto) GetSaturdayOk() (*bool, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *ScheduleDto) SetSaturday(v bool)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *ScheduleDto) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetUniqueInterval

`func (o *ScheduleDto) GetUniqueInterval() bool`

GetUniqueInterval returns the UniqueInterval field if non-nil, zero value otherwise.

### GetUniqueIntervalOk

`func (o *ScheduleDto) GetUniqueIntervalOk() (*bool, bool)`

GetUniqueIntervalOk returns a tuple with the UniqueInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueInterval

`func (o *ScheduleDto) SetUniqueInterval(v bool)`

SetUniqueInterval sets UniqueInterval field to given value.

### HasUniqueInterval

`func (o *ScheduleDto) HasUniqueInterval() bool`

HasUniqueInterval returns a boolean if a field has been set.

### GetIs24x7Interval

`func (o *ScheduleDto) GetIs24x7Interval() bool`

GetIs24x7Interval returns the Is24x7Interval field if non-nil, zero value otherwise.

### GetIs24x7IntervalOk

`func (o *ScheduleDto) GetIs24x7IntervalOk() (*bool, bool)`

GetIs24x7IntervalOk returns a tuple with the Is24x7Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs24x7Interval

`func (o *ScheduleDto) SetIs24x7Interval(v bool)`

SetIs24x7Interval sets Is24x7Interval field to given value.

### HasIs24x7Interval

`func (o *ScheduleDto) HasIs24x7Interval() bool`

HasIs24x7Interval returns a boolean if a field has been set.

### GetStart

`func (o *ScheduleDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ScheduleDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ScheduleDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ScheduleDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *ScheduleDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *ScheduleDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *ScheduleDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ScheduleDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ScheduleDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ScheduleDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *ScheduleDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *ScheduleDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetTimezoneId

`func (o *ScheduleDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ScheduleDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ScheduleDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ScheduleDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ScheduleDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ScheduleDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetFiscalYearId

`func (o *ScheduleDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *ScheduleDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *ScheduleDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *ScheduleDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *ScheduleDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *ScheduleDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetHolidayScheduleId

`func (o *ScheduleDto) GetHolidayScheduleId() string`

GetHolidayScheduleId returns the HolidayScheduleId field if non-nil, zero value otherwise.

### GetHolidayScheduleIdOk

`func (o *ScheduleDto) GetHolidayScheduleIdOk() (*string, bool)`

GetHolidayScheduleIdOk returns a tuple with the HolidayScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayScheduleId

`func (o *ScheduleDto) SetHolidayScheduleId(v string)`

SetHolidayScheduleId sets HolidayScheduleId field to given value.

### HasHolidayScheduleId

`func (o *ScheduleDto) HasHolidayScheduleId() bool`

HasHolidayScheduleId returns a boolean if a field has been set.

### SetHolidayScheduleIdNil

`func (o *ScheduleDto) SetHolidayScheduleIdNil(b bool)`

 SetHolidayScheduleIdNil sets the value for HolidayScheduleId to be an explicit nil

### UnsetHolidayScheduleId
`func (o *ScheduleDto) UnsetHolidayScheduleId()`

UnsetHolidayScheduleId ensures that no value is present for HolidayScheduleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


