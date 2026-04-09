# CourseWikiCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseWikiCreateDto

`func NewCourseWikiCreateDto(title string, courseID string, businessID string, ) *CourseWikiCreateDto`

NewCourseWikiCreateDto instantiates a new CourseWikiCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseWikiCreateDtoWithDefaults

`func NewCourseWikiCreateDtoWithDefaults() *CourseWikiCreateDto`

NewCourseWikiCreateDtoWithDefaults instantiates a new CourseWikiCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseWikiCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseWikiCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseWikiCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseWikiCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseWikiCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseWikiCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseWikiCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseWikiCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseWikiCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseWikiCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseWikiCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseWikiCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseWikiCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseWikiCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseWikiCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseWikiCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseWikiCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCourseID

`func (o *CourseWikiCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseWikiCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseWikiCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseWikiCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseWikiCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseWikiCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetCourseUnitID

`func (o *CourseWikiCreateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseWikiCreateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseWikiCreateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseWikiCreateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseWikiCreateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseWikiCreateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseWikiCreateDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseWikiCreateDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseWikiCreateDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseWikiCreateDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTimeNil

`func (o *CourseWikiCreateDto) SetReleaseDateTimeNil(b bool)`

 SetReleaseDateTimeNil sets the value for ReleaseDateTime to be an explicit nil

### UnsetReleaseDateTime
`func (o *CourseWikiCreateDto) UnsetReleaseDateTime()`

UnsetReleaseDateTime ensures that no value is present for ReleaseDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


