# CourseTeamMembershipDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**InstructorProfileID** | Pointer to **NullableString** |  | [optional] 
**CourseTeamMembershipType** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

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
### GetCourseID

`func (o *CourseTeamMembershipDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseTeamMembershipDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseTeamMembershipDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseTeamMembershipDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseTeamMembershipDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseTeamMembershipDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetInstructorProfileID

`func (o *CourseTeamMembershipDto) GetInstructorProfileID() string`

GetInstructorProfileID returns the InstructorProfileID field if non-nil, zero value otherwise.

### GetInstructorProfileIDOk

`func (o *CourseTeamMembershipDto) GetInstructorProfileIDOk() (*string, bool)`

GetInstructorProfileIDOk returns a tuple with the InstructorProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorProfileID

`func (o *CourseTeamMembershipDto) SetInstructorProfileID(v string)`

SetInstructorProfileID sets InstructorProfileID field to given value.

### HasInstructorProfileID

`func (o *CourseTeamMembershipDto) HasInstructorProfileID() bool`

HasInstructorProfileID returns a boolean if a field has been set.

### SetInstructorProfileIDNil

`func (o *CourseTeamMembershipDto) SetInstructorProfileIDNil(b bool)`

 SetInstructorProfileIDNil sets the value for InstructorProfileID to be an explicit nil

### UnsetInstructorProfileID
`func (o *CourseTeamMembershipDto) UnsetInstructorProfileID()`

UnsetInstructorProfileID ensures that no value is present for InstructorProfileID, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


