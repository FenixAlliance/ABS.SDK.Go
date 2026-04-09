# CourseSectionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 
**HideFromStudents** | Pointer to **bool** |  | [optional] 

## Methods

### NewCourseSectionUpdateDto

`func NewCourseSectionUpdateDto() *CourseSectionUpdateDto`

NewCourseSectionUpdateDto instantiates a new CourseSectionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseSectionUpdateDtoWithDefaults

`func NewCourseSectionUpdateDtoWithDefaults() *CourseSectionUpdateDto`

NewCourseSectionUpdateDtoWithDefaults instantiates a new CourseSectionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CourseSectionUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseSectionUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseSectionUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CourseSectionUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CourseSectionUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CourseSectionUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *CourseSectionUpdateDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CourseSectionUpdateDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CourseSectionUpdateDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CourseSectionUpdateDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CourseSectionUpdateDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CourseSectionUpdateDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *CourseSectionUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseSectionUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseSectionUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseSectionUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseSectionUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseSectionUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseSectionUpdateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseSectionUpdateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseSectionUpdateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseSectionUpdateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseSectionUpdateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseSectionUpdateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil
### GetHideFromStudents

`func (o *CourseSectionUpdateDto) GetHideFromStudents() bool`

GetHideFromStudents returns the HideFromStudents field if non-nil, zero value otherwise.

### GetHideFromStudentsOk

`func (o *CourseSectionUpdateDto) GetHideFromStudentsOk() (*bool, bool)`

GetHideFromStudentsOk returns a tuple with the HideFromStudents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromStudents

`func (o *CourseSectionUpdateDto) SetHideFromStudents(v bool)`

SetHideFromStudents sets HideFromStudents field to given value.

### HasHideFromStudents

`func (o *CourseSectionUpdateDto) HasHideFromStudents() bool`

HasHideFromStudents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


