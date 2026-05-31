# CourseTeamMembershipCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CourseID** | **string** |  | 
**InstructorProfileID** | **string** |  | 
**CourseTeamMembershipType** | Pointer to **string** |  | [optional] 

## Methods

### NewCourseTeamMembershipCreateDto

`func NewCourseTeamMembershipCreateDto(courseID string, instructorProfileID string, ) *CourseTeamMembershipCreateDto`

NewCourseTeamMembershipCreateDto instantiates a new CourseTeamMembershipCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseTeamMembershipCreateDtoWithDefaults

`func NewCourseTeamMembershipCreateDtoWithDefaults() *CourseTeamMembershipCreateDto`

NewCourseTeamMembershipCreateDtoWithDefaults instantiates a new CourseTeamMembershipCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseTeamMembershipCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseTeamMembershipCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseTeamMembershipCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseTeamMembershipCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseTeamMembershipCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseTeamMembershipCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseTeamMembershipCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseTeamMembershipCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseTeamMembershipCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseTeamMembershipCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseTeamMembershipCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetInstructorProfileID

`func (o *CourseTeamMembershipCreateDto) GetInstructorProfileID() string`

GetInstructorProfileID returns the InstructorProfileID field if non-nil, zero value otherwise.

### GetInstructorProfileIDOk

`func (o *CourseTeamMembershipCreateDto) GetInstructorProfileIDOk() (*string, bool)`

GetInstructorProfileIDOk returns a tuple with the InstructorProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorProfileID

`func (o *CourseTeamMembershipCreateDto) SetInstructorProfileID(v string)`

SetInstructorProfileID sets InstructorProfileID field to given value.


### GetCourseTeamMembershipType

`func (o *CourseTeamMembershipCreateDto) GetCourseTeamMembershipType() string`

GetCourseTeamMembershipType returns the CourseTeamMembershipType field if non-nil, zero value otherwise.

### GetCourseTeamMembershipTypeOk

`func (o *CourseTeamMembershipCreateDto) GetCourseTeamMembershipTypeOk() (*string, bool)`

GetCourseTeamMembershipTypeOk returns a tuple with the CourseTeamMembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTeamMembershipType

`func (o *CourseTeamMembershipCreateDto) SetCourseTeamMembershipType(v string)`

SetCourseTeamMembershipType sets CourseTeamMembershipType field to given value.

### HasCourseTeamMembershipType

`func (o *CourseTeamMembershipCreateDto) HasCourseTeamMembershipType() bool`

HasCourseTeamMembershipType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


