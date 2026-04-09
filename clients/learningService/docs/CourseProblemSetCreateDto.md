# CourseProblemSetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**OverallScore** | Pointer to **float64** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**CourseGradingRubricID** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseProblemSetCreateDto

`func NewCourseProblemSetCreateDto(title string, courseID string, businessID string, ) *CourseProblemSetCreateDto`

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

### GetCourseID

`func (o *CourseProblemSetCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseProblemSetCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseProblemSetCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseProblemSetCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseProblemSetCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseProblemSetCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetCourseUnitID

`func (o *CourseProblemSetCreateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseProblemSetCreateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseProblemSetCreateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseProblemSetCreateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseProblemSetCreateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseProblemSetCreateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetCourseGradingRubricID

`func (o *CourseProblemSetCreateDto) GetCourseGradingRubricID() string`

GetCourseGradingRubricID returns the CourseGradingRubricID field if non-nil, zero value otherwise.

### GetCourseGradingRubricIDOk

`func (o *CourseProblemSetCreateDto) GetCourseGradingRubricIDOk() (*string, bool)`

GetCourseGradingRubricIDOk returns a tuple with the CourseGradingRubricID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseGradingRubricID

`func (o *CourseProblemSetCreateDto) SetCourseGradingRubricID(v string)`

SetCourseGradingRubricID sets CourseGradingRubricID field to given value.

### HasCourseGradingRubricID

`func (o *CourseProblemSetCreateDto) HasCourseGradingRubricID() bool`

HasCourseGradingRubricID returns a boolean if a field has been set.

### SetCourseGradingRubricIDNil

`func (o *CourseProblemSetCreateDto) SetCourseGradingRubricIDNil(b bool)`

 SetCourseGradingRubricIDNil sets the value for CourseGradingRubricID to be an explicit nil

### UnsetCourseGradingRubricID
`func (o *CourseProblemSetCreateDto) UnsetCourseGradingRubricID()`

UnsetCourseGradingRubricID ensures that no value is present for CourseGradingRubricID, not even an explicit nil
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


