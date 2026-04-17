# CourseUnitCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**CourseID** | **string** |  | 
**CourseSectionID** | **string** |  | 
**CourseContentGroupID** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseUnitCreateDto

`func NewCourseUnitCreateDto(title string, courseID string, courseSectionID string, ) *CourseUnitCreateDto`

NewCourseUnitCreateDto instantiates a new CourseUnitCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseUnitCreateDtoWithDefaults

`func NewCourseUnitCreateDtoWithDefaults() *CourseUnitCreateDto`

NewCourseUnitCreateDtoWithDefaults instantiates a new CourseUnitCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseUnitCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseUnitCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseUnitCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseUnitCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseUnitCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseUnitCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseUnitCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseUnitCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseUnitCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseUnitCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseUnitCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseUnitCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseUnitCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseUnitCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseUnitCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseUnitCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseUnitCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseUnitCreateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseUnitCreateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseUnitCreateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseUnitCreateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseUnitCreateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseUnitCreateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCourseID

`func (o *CourseUnitCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseUnitCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseUnitCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetCourseSectionID

`func (o *CourseUnitCreateDto) GetCourseSectionID() string`

GetCourseSectionID returns the CourseSectionID field if non-nil, zero value otherwise.

### GetCourseSectionIDOk

`func (o *CourseUnitCreateDto) GetCourseSectionIDOk() (*string, bool)`

GetCourseSectionIDOk returns a tuple with the CourseSectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseSectionID

`func (o *CourseUnitCreateDto) SetCourseSectionID(v string)`

SetCourseSectionID sets CourseSectionID field to given value.


### GetCourseContentGroupID

`func (o *CourseUnitCreateDto) GetCourseContentGroupID() string`

GetCourseContentGroupID returns the CourseContentGroupID field if non-nil, zero value otherwise.

### GetCourseContentGroupIDOk

`func (o *CourseUnitCreateDto) GetCourseContentGroupIDOk() (*string, bool)`

GetCourseContentGroupIDOk returns a tuple with the CourseContentGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseContentGroupID

`func (o *CourseUnitCreateDto) SetCourseContentGroupID(v string)`

SetCourseContentGroupID sets CourseContentGroupID field to given value.

### HasCourseContentGroupID

`func (o *CourseUnitCreateDto) HasCourseContentGroupID() bool`

HasCourseContentGroupID returns a boolean if a field has been set.

### SetCourseContentGroupIDNil

`func (o *CourseUnitCreateDto) SetCourseContentGroupIDNil(b bool)`

 SetCourseContentGroupIDNil sets the value for CourseContentGroupID to be an explicit nil

### UnsetCourseContentGroupID
`func (o *CourseUnitCreateDto) UnsetCourseContentGroupID()`

UnsetCourseContentGroupID ensures that no value is present for CourseContentGroupID, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseUnitCreateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseUnitCreateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseUnitCreateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseUnitCreateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseUnitCreateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseUnitCreateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


