# CourseLibraryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseLibraryUpdateDto

`func NewCourseLibraryUpdateDto() *CourseLibraryUpdateDto`

NewCourseLibraryUpdateDto instantiates a new CourseLibraryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseLibraryUpdateDtoWithDefaults

`func NewCourseLibraryUpdateDtoWithDefaults() *CourseLibraryUpdateDto`

NewCourseLibraryUpdateDtoWithDefaults instantiates a new CourseLibraryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CourseLibraryUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseLibraryUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseLibraryUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseLibraryUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseLibraryUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseLibraryUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseLibraryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseLibraryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseLibraryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseLibraryUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseLibraryUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseLibraryUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCourseUnitID

`func (o *CourseLibraryUpdateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseLibraryUpdateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseLibraryUpdateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseLibraryUpdateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseLibraryUpdateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseLibraryUpdateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseLibraryUpdateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseLibraryUpdateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseLibraryUpdateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseLibraryUpdateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseLibraryUpdateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseLibraryUpdateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


