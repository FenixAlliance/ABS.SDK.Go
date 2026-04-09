# CourseSectionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 
**HideFromStudents** | Pointer to **bool** |  | [optional] 

## Methods

### NewCourseSectionCreateDto

`func NewCourseSectionCreateDto(name string, courseID string, businessID string, ) *CourseSectionCreateDto`

NewCourseSectionCreateDto instantiates a new CourseSectionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseSectionCreateDtoWithDefaults

`func NewCourseSectionCreateDtoWithDefaults() *CourseSectionCreateDto`

NewCourseSectionCreateDtoWithDefaults instantiates a new CourseSectionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseSectionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseSectionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseSectionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseSectionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseSectionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseSectionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseSectionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseSectionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CourseSectionCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseSectionCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseSectionCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *CourseSectionCreateDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CourseSectionCreateDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CourseSectionCreateDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CourseSectionCreateDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CourseSectionCreateDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CourseSectionCreateDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *CourseSectionCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseSectionCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseSectionCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseSectionCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseSectionCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseSectionCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCourseID

`func (o *CourseSectionCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseSectionCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseSectionCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseSectionCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseSectionCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseSectionCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetReleaseDateTime

`func (o *CourseSectionCreateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseSectionCreateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseSectionCreateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseSectionCreateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseSectionCreateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseSectionCreateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil
### GetHideFromStudents

`func (o *CourseSectionCreateDto) GetHideFromStudents() bool`

GetHideFromStudents returns the HideFromStudents field if non-nil, zero value otherwise.

### GetHideFromStudentsOk

`func (o *CourseSectionCreateDto) GetHideFromStudentsOk() (*bool, bool)`

GetHideFromStudentsOk returns a tuple with the HideFromStudents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromStudents

`func (o *CourseSectionCreateDto) SetHideFromStudents(v bool)`

SetHideFromStudents sets HideFromStudents field to given value.

### HasHideFromStudents

`func (o *CourseSectionCreateDto) HasHideFromStudents() bool`

HasHideFromStudents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


