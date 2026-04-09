# CourseProblemSetUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**OverallScore** | Pointer to **NullableFloat64** |  | [optional] 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**CourseGradingRubricID** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseProblemSetUpdateDto

`func NewCourseProblemSetUpdateDto() *CourseProblemSetUpdateDto`

NewCourseProblemSetUpdateDto instantiates a new CourseProblemSetUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseProblemSetUpdateDtoWithDefaults

`func NewCourseProblemSetUpdateDtoWithDefaults() *CourseProblemSetUpdateDto`

NewCourseProblemSetUpdateDtoWithDefaults instantiates a new CourseProblemSetUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CourseProblemSetUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseProblemSetUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseProblemSetUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseProblemSetUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseProblemSetUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseProblemSetUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseProblemSetUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseProblemSetUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseProblemSetUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseProblemSetUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseProblemSetUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseProblemSetUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOverallScore

`func (o *CourseProblemSetUpdateDto) GetOverallScore() float64`

GetOverallScore returns the OverallScore field if non-nil, zero value otherwise.

### GetOverallScoreOk

`func (o *CourseProblemSetUpdateDto) GetOverallScoreOk() (*float64, bool)`

GetOverallScoreOk returns a tuple with the OverallScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallScore

`func (o *CourseProblemSetUpdateDto) SetOverallScore(v float64)`

SetOverallScore sets OverallScore field to given value.

### HasOverallScore

`func (o *CourseProblemSetUpdateDto) HasOverallScore() bool`

HasOverallScore returns a boolean if a field has been set.

### SetOverallScoreNil

`func (o *CourseProblemSetUpdateDto) SetOverallScoreNil(b bool)`

 SetOverallScoreNil sets the value for OverallScore to be an explicit nil

### UnsetOverallScore
`func (o *CourseProblemSetUpdateDto) UnsetOverallScore()`

UnsetOverallScore ensures that no value is present for OverallScore, not even an explicit nil
### GetCourseUnitID

`func (o *CourseProblemSetUpdateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseProblemSetUpdateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseProblemSetUpdateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseProblemSetUpdateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseProblemSetUpdateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseProblemSetUpdateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetCourseGradingRubricID

`func (o *CourseProblemSetUpdateDto) GetCourseGradingRubricID() string`

GetCourseGradingRubricID returns the CourseGradingRubricID field if non-nil, zero value otherwise.

### GetCourseGradingRubricIDOk

`func (o *CourseProblemSetUpdateDto) GetCourseGradingRubricIDOk() (*string, bool)`

GetCourseGradingRubricIDOk returns a tuple with the CourseGradingRubricID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseGradingRubricID

`func (o *CourseProblemSetUpdateDto) SetCourseGradingRubricID(v string)`

SetCourseGradingRubricID sets CourseGradingRubricID field to given value.

### HasCourseGradingRubricID

`func (o *CourseProblemSetUpdateDto) HasCourseGradingRubricID() bool`

HasCourseGradingRubricID returns a boolean if a field has been set.

### SetCourseGradingRubricIDNil

`func (o *CourseProblemSetUpdateDto) SetCourseGradingRubricIDNil(b bool)`

 SetCourseGradingRubricIDNil sets the value for CourseGradingRubricID to be an explicit nil

### UnsetCourseGradingRubricID
`func (o *CourseProblemSetUpdateDto) UnsetCourseGradingRubricID()`

UnsetCourseGradingRubricID ensures that no value is present for CourseGradingRubricID, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseProblemSetUpdateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseProblemSetUpdateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseProblemSetUpdateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseProblemSetUpdateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseProblemSetUpdateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseProblemSetUpdateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


