# CourseTeamMembershipDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**InstructorProfileId** | Pointer to **NullableString** |  | [optional] 
**CourseTeamMembershipType** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseTeamMembershipDto

`func NewCourseTeamMembershipDto() *CourseTeamMembershipDto`

NewCourseTeamMembershipDto instantiates a new CourseTeamMembershipDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseTeamMembershipDtoWithDefaults

`func NewCourseTeamMembershipDtoWithDefaults() *CourseTeamMembershipDto`

NewCourseTeamMembershipDtoWithDefaults instantiates a new CourseTeamMembershipDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseTeamMembershipDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseTeamMembershipDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseTeamMembershipDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseTeamMembershipDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseTeamMembershipDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseTeamMembershipDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseTeamMembershipDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseTeamMembershipDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseTeamMembershipDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseTeamMembershipDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseTeamMembershipDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseTeamMembershipDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCourseId

`func (o *CourseTeamMembershipDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseTeamMembershipDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseTeamMembershipDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseTeamMembershipDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseTeamMembershipDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseTeamMembershipDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetInstructorProfileId

`func (o *CourseTeamMembershipDto) GetInstructorProfileId() string`

GetInstructorProfileId returns the InstructorProfileId field if non-nil, zero value otherwise.

### GetInstructorProfileIdOk

`func (o *CourseTeamMembershipDto) GetInstructorProfileIdOk() (*string, bool)`

GetInstructorProfileIdOk returns a tuple with the InstructorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorProfileId

`func (o *CourseTeamMembershipDto) SetInstructorProfileId(v string)`

SetInstructorProfileId sets InstructorProfileId field to given value.

### HasInstructorProfileId

`func (o *CourseTeamMembershipDto) HasInstructorProfileId() bool`

HasInstructorProfileId returns a boolean if a field has been set.

### SetInstructorProfileIdNil

`func (o *CourseTeamMembershipDto) SetInstructorProfileIdNil(b bool)`

 SetInstructorProfileIdNil sets the value for InstructorProfileId to be an explicit nil

### UnsetInstructorProfileId
`func (o *CourseTeamMembershipDto) UnsetInstructorProfileId()`

UnsetInstructorProfileId ensures that no value is present for InstructorProfileId, not even an explicit nil
### GetCourseTeamMembershipType

`func (o *CourseTeamMembershipDto) GetCourseTeamMembershipType() string`

GetCourseTeamMembershipType returns the CourseTeamMembershipType field if non-nil, zero value otherwise.

### GetCourseTeamMembershipTypeOk

`func (o *CourseTeamMembershipDto) GetCourseTeamMembershipTypeOk() (*string, bool)`

GetCourseTeamMembershipTypeOk returns a tuple with the CourseTeamMembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTeamMembershipType

`func (o *CourseTeamMembershipDto) SetCourseTeamMembershipType(v string)`

SetCourseTeamMembershipType sets CourseTeamMembershipType field to given value.

### HasCourseTeamMembershipType

`func (o *CourseTeamMembershipDto) HasCourseTeamMembershipType() bool`

HasCourseTeamMembershipType returns a boolean if a field has been set.

### GetTenantId

`func (o *CourseTeamMembershipDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseTeamMembershipDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseTeamMembershipDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseTeamMembershipDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseTeamMembershipDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseTeamMembershipDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseTeamMembershipDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseTeamMembershipDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseTeamMembershipDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseTeamMembershipDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseTeamMembershipDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseTeamMembershipDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


