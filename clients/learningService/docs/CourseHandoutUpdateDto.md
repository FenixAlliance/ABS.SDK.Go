# CourseHandoutUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 
**CourseUnitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseHandoutUpdateDto

`func NewCourseHandoutUpdateDto() *CourseHandoutUpdateDto`

NewCourseHandoutUpdateDto instantiates a new CourseHandoutUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseHandoutUpdateDtoWithDefaults

`func NewCourseHandoutUpdateDtoWithDefaults() *CourseHandoutUpdateDto`

NewCourseHandoutUpdateDtoWithDefaults instantiates a new CourseHandoutUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CourseHandoutUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseHandoutUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseHandoutUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CourseHandoutUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CourseHandoutUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CourseHandoutUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CourseHandoutUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseHandoutUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseHandoutUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseHandoutUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseHandoutUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseHandoutUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseHandoutUpdateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseHandoutUpdateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseHandoutUpdateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseHandoutUpdateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseHandoutUpdateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseHandoutUpdateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetUrl

`func (o *CourseHandoutUpdateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CourseHandoutUpdateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CourseHandoutUpdateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CourseHandoutUpdateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CourseHandoutUpdateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CourseHandoutUpdateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseHandoutUpdateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseHandoutUpdateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseHandoutUpdateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseHandoutUpdateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseHandoutUpdateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseHandoutUpdateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil
### GetCourseUnitId

`func (o *CourseHandoutUpdateDto) GetCourseUnitId() string`

GetCourseUnitId returns the CourseUnitId field if non-nil, zero value otherwise.

### GetCourseUnitIdOk

`func (o *CourseHandoutUpdateDto) GetCourseUnitIdOk() (*string, bool)`

GetCourseUnitIdOk returns a tuple with the CourseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitId

`func (o *CourseHandoutUpdateDto) SetCourseUnitId(v string)`

SetCourseUnitId sets CourseUnitId field to given value.

### HasCourseUnitId

`func (o *CourseHandoutUpdateDto) HasCourseUnitId() bool`

HasCourseUnitId returns a boolean if a field has been set.

### SetCourseUnitIdNil

`func (o *CourseHandoutUpdateDto) SetCourseUnitIdNil(b bool)`

 SetCourseUnitIdNil sets the value for CourseUnitId to be an explicit nil

### UnsetCourseUnitId
`func (o *CourseHandoutUpdateDto) UnsetCourseUnitId()`

UnsetCourseUnitId ensures that no value is present for CourseUnitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


