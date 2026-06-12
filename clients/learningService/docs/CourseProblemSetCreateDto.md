# CourseProblemSetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**OverallScore** | Pointer to **float64** |  | [optional] 
**CourseId** | **string** |  | 
**CourseUnitId** | Pointer to **NullableString** |  | [optional] 
**CourseGradingRubricId** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseProblemSetCreateDto

`func NewCourseProblemSetCreateDto(title string, courseId string, ) *CourseProblemSetCreateDto`

NewCourseProblemSetCreateDto instantiates a new CourseProblemSetCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseProblemSetCreateDtoWithDefaults

`func NewCourseProblemSetCreateDtoWithDefaults() *CourseProblemSetCreateDto`

NewCourseProblemSetCreateDtoWithDefaults instantiates a new CourseProblemSetCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseProblemSetCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseProblemSetCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseProblemSetCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseProblemSetCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseProblemSetCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseProblemSetCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseProblemSetCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseProblemSetCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseProblemSetCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseProblemSetCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseProblemSetCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseProblemSetCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseProblemSetCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseProblemSetCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseProblemSetCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseProblemSetCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseProblemSetCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOverallScore

`func (o *CourseProblemSetCreateDto) GetOverallScore() float64`

GetOverallScore returns the OverallScore field if non-nil, zero value otherwise.

### GetOverallScoreOk

`func (o *CourseProblemSetCreateDto) GetOverallScoreOk() (*float64, bool)`

GetOverallScoreOk returns a tuple with the OverallScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallScore

`func (o *CourseProblemSetCreateDto) SetOverallScore(v float64)`

SetOverallScore sets OverallScore field to given value.

### HasOverallScore

`func (o *CourseProblemSetCreateDto) HasOverallScore() bool`

HasOverallScore returns a boolean if a field has been set.

### GetCourseId

`func (o *CourseProblemSetCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseProblemSetCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseProblemSetCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.


### GetCourseUnitId

`func (o *CourseProblemSetCreateDto) GetCourseUnitId() string`

GetCourseUnitId returns the CourseUnitId field if non-nil, zero value otherwise.

### GetCourseUnitIdOk

`func (o *CourseProblemSetCreateDto) GetCourseUnitIdOk() (*string, bool)`

GetCourseUnitIdOk returns a tuple with the CourseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitId

`func (o *CourseProblemSetCreateDto) SetCourseUnitId(v string)`

SetCourseUnitId sets CourseUnitId field to given value.

### HasCourseUnitId

`func (o *CourseProblemSetCreateDto) HasCourseUnitId() bool`

HasCourseUnitId returns a boolean if a field has been set.

### SetCourseUnitIdNil

`func (o *CourseProblemSetCreateDto) SetCourseUnitIdNil(b bool)`

 SetCourseUnitIdNil sets the value for CourseUnitId to be an explicit nil

### UnsetCourseUnitId
`func (o *CourseProblemSetCreateDto) UnsetCourseUnitId()`

UnsetCourseUnitId ensures that no value is present for CourseUnitId, not even an explicit nil
### GetCourseGradingRubricId

`func (o *CourseProblemSetCreateDto) GetCourseGradingRubricId() string`

GetCourseGradingRubricId returns the CourseGradingRubricId field if non-nil, zero value otherwise.

### GetCourseGradingRubricIdOk

`func (o *CourseProblemSetCreateDto) GetCourseGradingRubricIdOk() (*string, bool)`

GetCourseGradingRubricIdOk returns a tuple with the CourseGradingRubricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseGradingRubricId

`func (o *CourseProblemSetCreateDto) SetCourseGradingRubricId(v string)`

SetCourseGradingRubricId sets CourseGradingRubricId field to given value.

### HasCourseGradingRubricId

`func (o *CourseProblemSetCreateDto) HasCourseGradingRubricId() bool`

HasCourseGradingRubricId returns a boolean if a field has been set.

### SetCourseGradingRubricIdNil

`func (o *CourseProblemSetCreateDto) SetCourseGradingRubricIdNil(b bool)`

 SetCourseGradingRubricIdNil sets the value for CourseGradingRubricId to be an explicit nil

### UnsetCourseGradingRubricId
`func (o *CourseProblemSetCreateDto) UnsetCourseGradingRubricId()`

UnsetCourseGradingRubricId ensures that no value is present for CourseGradingRubricId, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseProblemSetCreateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseProblemSetCreateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseProblemSetCreateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseProblemSetCreateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseProblemSetCreateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseProblemSetCreateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


