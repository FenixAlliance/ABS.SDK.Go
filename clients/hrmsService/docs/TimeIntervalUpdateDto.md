# TimeIntervalUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsBreak** | Pointer to **NullableBool** |  | [optional] 
**OccustOnMonday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnTuesday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnWednesday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnThursday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnFriday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnSaturday** | Pointer to **NullableBool** |  | [optional] 
**OccustOnSunday** | Pointer to **NullableBool** |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**RepeatEvery** | Pointer to **NullableInt32** |  | [optional] 
**ParentTimeIntervalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTimeIntervalUpdateDto

`func NewTimeIntervalUpdateDto() *TimeIntervalUpdateDto`

NewTimeIntervalUpdateDto instantiates a new TimeIntervalUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeIntervalUpdateDtoWithDefaults

`func NewTimeIntervalUpdateDtoWithDefaults() *TimeIntervalUpdateDto`

NewTimeIntervalUpdateDtoWithDefaults instantiates a new TimeIntervalUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TimeIntervalUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimeIntervalUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimeIntervalUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TimeIntervalUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TimeIntervalUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TimeIntervalUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TimeIntervalUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimeIntervalUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimeIntervalUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimeIntervalUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TimeIntervalUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TimeIntervalUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsBreak

`func (o *TimeIntervalUpdateDto) GetIsBreak() bool`

GetIsBreak returns the IsBreak field if non-nil, zero value otherwise.

### GetIsBreakOk

`func (o *TimeIntervalUpdateDto) GetIsBreakOk() (*bool, bool)`

GetIsBreakOk returns a tuple with the IsBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreak

`func (o *TimeIntervalUpdateDto) SetIsBreak(v bool)`

SetIsBreak sets IsBreak field to given value.

### HasIsBreak

`func (o *TimeIntervalUpdateDto) HasIsBreak() bool`

HasIsBreak returns a boolean if a field has been set.

### SetIsBreakNil

`func (o *TimeIntervalUpdateDto) SetIsBreakNil(b bool)`

 SetIsBreakNil sets the value for IsBreak to be an explicit nil

### UnsetIsBreak
`func (o *TimeIntervalUpdateDto) UnsetIsBreak()`

UnsetIsBreak ensures that no value is present for IsBreak, not even an explicit nil
### GetOccustOnMonday

`func (o *TimeIntervalUpdateDto) GetOccustOnMonday() bool`

GetOccustOnMonday returns the OccustOnMonday field if non-nil, zero value otherwise.

### GetOccustOnMondayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnMondayOk() (*bool, bool)`

GetOccustOnMondayOk returns a tuple with the OccustOnMonday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnMonday

`func (o *TimeIntervalUpdateDto) SetOccustOnMonday(v bool)`

SetOccustOnMonday sets OccustOnMonday field to given value.

### HasOccustOnMonday

`func (o *TimeIntervalUpdateDto) HasOccustOnMonday() bool`

HasOccustOnMonday returns a boolean if a field has been set.

### SetOccustOnMondayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnMondayNil(b bool)`

 SetOccustOnMondayNil sets the value for OccustOnMonday to be an explicit nil

### UnsetOccustOnMonday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnMonday()`

UnsetOccustOnMonday ensures that no value is present for OccustOnMonday, not even an explicit nil
### GetOccustOnTuesday

`func (o *TimeIntervalUpdateDto) GetOccustOnTuesday() bool`

GetOccustOnTuesday returns the OccustOnTuesday field if non-nil, zero value otherwise.

### GetOccustOnTuesdayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnTuesdayOk() (*bool, bool)`

GetOccustOnTuesdayOk returns a tuple with the OccustOnTuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnTuesday

`func (o *TimeIntervalUpdateDto) SetOccustOnTuesday(v bool)`

SetOccustOnTuesday sets OccustOnTuesday field to given value.

### HasOccustOnTuesday

`func (o *TimeIntervalUpdateDto) HasOccustOnTuesday() bool`

HasOccustOnTuesday returns a boolean if a field has been set.

### SetOccustOnTuesdayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnTuesdayNil(b bool)`

 SetOccustOnTuesdayNil sets the value for OccustOnTuesday to be an explicit nil

### UnsetOccustOnTuesday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnTuesday()`

UnsetOccustOnTuesday ensures that no value is present for OccustOnTuesday, not even an explicit nil
### GetOccustOnWednesday

`func (o *TimeIntervalUpdateDto) GetOccustOnWednesday() bool`

GetOccustOnWednesday returns the OccustOnWednesday field if non-nil, zero value otherwise.

### GetOccustOnWednesdayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnWednesdayOk() (*bool, bool)`

GetOccustOnWednesdayOk returns a tuple with the OccustOnWednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnWednesday

`func (o *TimeIntervalUpdateDto) SetOccustOnWednesday(v bool)`

SetOccustOnWednesday sets OccustOnWednesday field to given value.

### HasOccustOnWednesday

`func (o *TimeIntervalUpdateDto) HasOccustOnWednesday() bool`

HasOccustOnWednesday returns a boolean if a field has been set.

### SetOccustOnWednesdayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnWednesdayNil(b bool)`

 SetOccustOnWednesdayNil sets the value for OccustOnWednesday to be an explicit nil

### UnsetOccustOnWednesday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnWednesday()`

UnsetOccustOnWednesday ensures that no value is present for OccustOnWednesday, not even an explicit nil
### GetOccustOnThursday

`func (o *TimeIntervalUpdateDto) GetOccustOnThursday() bool`

GetOccustOnThursday returns the OccustOnThursday field if non-nil, zero value otherwise.

### GetOccustOnThursdayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnThursdayOk() (*bool, bool)`

GetOccustOnThursdayOk returns a tuple with the OccustOnThursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnThursday

`func (o *TimeIntervalUpdateDto) SetOccustOnThursday(v bool)`

SetOccustOnThursday sets OccustOnThursday field to given value.

### HasOccustOnThursday

`func (o *TimeIntervalUpdateDto) HasOccustOnThursday() bool`

HasOccustOnThursday returns a boolean if a field has been set.

### SetOccustOnThursdayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnThursdayNil(b bool)`

 SetOccustOnThursdayNil sets the value for OccustOnThursday to be an explicit nil

### UnsetOccustOnThursday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnThursday()`

UnsetOccustOnThursday ensures that no value is present for OccustOnThursday, not even an explicit nil
### GetOccustOnFriday

`func (o *TimeIntervalUpdateDto) GetOccustOnFriday() bool`

GetOccustOnFriday returns the OccustOnFriday field if non-nil, zero value otherwise.

### GetOccustOnFridayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnFridayOk() (*bool, bool)`

GetOccustOnFridayOk returns a tuple with the OccustOnFriday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnFriday

`func (o *TimeIntervalUpdateDto) SetOccustOnFriday(v bool)`

SetOccustOnFriday sets OccustOnFriday field to given value.

### HasOccustOnFriday

`func (o *TimeIntervalUpdateDto) HasOccustOnFriday() bool`

HasOccustOnFriday returns a boolean if a field has been set.

### SetOccustOnFridayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnFridayNil(b bool)`

 SetOccustOnFridayNil sets the value for OccustOnFriday to be an explicit nil

### UnsetOccustOnFriday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnFriday()`

UnsetOccustOnFriday ensures that no value is present for OccustOnFriday, not even an explicit nil
### GetOccustOnSaturday

`func (o *TimeIntervalUpdateDto) GetOccustOnSaturday() bool`

GetOccustOnSaturday returns the OccustOnSaturday field if non-nil, zero value otherwise.

### GetOccustOnSaturdayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnSaturdayOk() (*bool, bool)`

GetOccustOnSaturdayOk returns a tuple with the OccustOnSaturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSaturday

`func (o *TimeIntervalUpdateDto) SetOccustOnSaturday(v bool)`

SetOccustOnSaturday sets OccustOnSaturday field to given value.

### HasOccustOnSaturday

`func (o *TimeIntervalUpdateDto) HasOccustOnSaturday() bool`

HasOccustOnSaturday returns a boolean if a field has been set.

### SetOccustOnSaturdayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnSaturdayNil(b bool)`

 SetOccustOnSaturdayNil sets the value for OccustOnSaturday to be an explicit nil

### UnsetOccustOnSaturday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnSaturday()`

UnsetOccustOnSaturday ensures that no value is present for OccustOnSaturday, not even an explicit nil
### GetOccustOnSunday

`func (o *TimeIntervalUpdateDto) GetOccustOnSunday() bool`

GetOccustOnSunday returns the OccustOnSunday field if non-nil, zero value otherwise.

### GetOccustOnSundayOk

`func (o *TimeIntervalUpdateDto) GetOccustOnSundayOk() (*bool, bool)`

GetOccustOnSundayOk returns a tuple with the OccustOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccustOnSunday

`func (o *TimeIntervalUpdateDto) SetOccustOnSunday(v bool)`

SetOccustOnSunday sets OccustOnSunday field to given value.

### HasOccustOnSunday

`func (o *TimeIntervalUpdateDto) HasOccustOnSunday() bool`

HasOccustOnSunday returns a boolean if a field has been set.

### SetOccustOnSundayNil

`func (o *TimeIntervalUpdateDto) SetOccustOnSundayNil(b bool)`

 SetOccustOnSundayNil sets the value for OccustOnSunday to be an explicit nil

### UnsetOccustOnSunday
`func (o *TimeIntervalUpdateDto) UnsetOccustOnSunday()`

UnsetOccustOnSunday ensures that no value is present for OccustOnSunday, not even an explicit nil
### GetStart

`func (o *TimeIntervalUpdateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TimeIntervalUpdateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TimeIntervalUpdateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *TimeIntervalUpdateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *TimeIntervalUpdateDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *TimeIntervalUpdateDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *TimeIntervalUpdateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimeIntervalUpdateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimeIntervalUpdateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TimeIntervalUpdateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *TimeIntervalUpdateDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *TimeIntervalUpdateDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetRepeatEvery

`func (o *TimeIntervalUpdateDto) GetRepeatEvery() int32`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *TimeIntervalUpdateDto) GetRepeatEveryOk() (*int32, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *TimeIntervalUpdateDto) SetRepeatEvery(v int32)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *TimeIntervalUpdateDto) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### SetRepeatEveryNil

`func (o *TimeIntervalUpdateDto) SetRepeatEveryNil(b bool)`

 SetRepeatEveryNil sets the value for RepeatEvery to be an explicit nil

### UnsetRepeatEvery
`func (o *TimeIntervalUpdateDto) UnsetRepeatEvery()`

UnsetRepeatEvery ensures that no value is present for RepeatEvery, not even an explicit nil
### GetParentTimeIntervalId

`func (o *TimeIntervalUpdateDto) GetParentTimeIntervalId() string`

GetParentTimeIntervalId returns the ParentTimeIntervalId field if non-nil, zero value otherwise.

### GetParentTimeIntervalIdOk

`func (o *TimeIntervalUpdateDto) GetParentTimeIntervalIdOk() (*string, bool)`

GetParentTimeIntervalIdOk returns a tuple with the ParentTimeIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeIntervalId

`func (o *TimeIntervalUpdateDto) SetParentTimeIntervalId(v string)`

SetParentTimeIntervalId sets ParentTimeIntervalId field to given value.

### HasParentTimeIntervalId

`func (o *TimeIntervalUpdateDto) HasParentTimeIntervalId() bool`

HasParentTimeIntervalId returns a boolean if a field has been set.

### SetParentTimeIntervalIdNil

`func (o *TimeIntervalUpdateDto) SetParentTimeIntervalIdNil(b bool)`

 SetParentTimeIntervalIdNil sets the value for ParentTimeIntervalId to be an explicit nil

### UnsetParentTimeIntervalId
`func (o *TimeIntervalUpdateDto) UnsetParentTimeIntervalId()`

UnsetParentTimeIntervalId ensures that no value is present for ParentTimeIntervalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


