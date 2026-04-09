# CourseCohortCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**ExpectedStartDateTime** | Pointer to **NullableTime** |  | [optional] 
**ExpectedEndDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseCohortCreateDto

`func NewCourseCohortCreateDto(name string, courseID string, businessID string, ) *CourseCohortCreateDto`

NewCourseCohortCreateDto instantiates a new CourseCohortCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCohortCreateDtoWithDefaults

`func NewCourseCohortCreateDtoWithDefaults() *CourseCohortCreateDto`

NewCourseCohortCreateDtoWithDefaults instantiates a new CourseCohortCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCohortCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCohortCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCohortCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCohortCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseCohortCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCohortCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCohortCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCohortCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CourseCohortCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseCohortCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseCohortCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCourseID

`func (o *CourseCohortCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseCohortCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseCohortCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseCohortCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseCohortCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseCohortCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetStartDateTime

`func (o *CourseCohortCreateDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CourseCohortCreateDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CourseCohortCreateDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CourseCohortCreateDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *CourseCohortCreateDto) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *CourseCohortCreateDto) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetEndDateTime

`func (o *CourseCohortCreateDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CourseCohortCreateDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CourseCohortCreateDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CourseCohortCreateDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *CourseCohortCreateDto) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *CourseCohortCreateDto) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetExpectedStartDateTime

`func (o *CourseCohortCreateDto) GetExpectedStartDateTime() time.Time`

GetExpectedStartDateTime returns the ExpectedStartDateTime field if non-nil, zero value otherwise.

### GetExpectedStartDateTimeOk

`func (o *CourseCohortCreateDto) GetExpectedStartDateTimeOk() (*time.Time, bool)`

GetExpectedStartDateTimeOk returns a tuple with the ExpectedStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartDateTime

`func (o *CourseCohortCreateDto) SetExpectedStartDateTime(v time.Time)`

SetExpectedStartDateTime sets ExpectedStartDateTime field to given value.

### HasExpectedStartDateTime

`func (o *CourseCohortCreateDto) HasExpectedStartDateTime() bool`

HasExpectedStartDateTime returns a boolean if a field has been set.

### SetExpectedStartDateTimeNil

`func (o *CourseCohortCreateDto) SetExpectedStartDateTimeNil(b bool)`

 SetExpectedStartDateTimeNil sets the value for ExpectedStartDateTime to be an explicit nil

### UnsetExpectedStartDateTime
`func (o *CourseCohortCreateDto) UnsetExpectedStartDateTime()`

UnsetExpectedStartDateTime ensures that no value is present for ExpectedStartDateTime, not even an explicit nil
### GetExpectedEndDateTime

`func (o *CourseCohortCreateDto) GetExpectedEndDateTime() time.Time`

GetExpectedEndDateTime returns the ExpectedEndDateTime field if non-nil, zero value otherwise.

### GetExpectedEndDateTimeOk

`func (o *CourseCohortCreateDto) GetExpectedEndDateTimeOk() (*time.Time, bool)`

GetExpectedEndDateTimeOk returns a tuple with the ExpectedEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedEndDateTime

`func (o *CourseCohortCreateDto) SetExpectedEndDateTime(v time.Time)`

SetExpectedEndDateTime sets ExpectedEndDateTime field to given value.

### HasExpectedEndDateTime

`func (o *CourseCohortCreateDto) HasExpectedEndDateTime() bool`

HasExpectedEndDateTime returns a boolean if a field has been set.

### SetExpectedEndDateTimeNil

`func (o *CourseCohortCreateDto) SetExpectedEndDateTimeNil(b bool)`

 SetExpectedEndDateTimeNil sets the value for ExpectedEndDateTime to be an explicit nil

### UnsetExpectedEndDateTime
`func (o *CourseCohortCreateDto) UnsetExpectedEndDateTime()`

UnsetExpectedEndDateTime ensures that no value is present for ExpectedEndDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


