# CourseCohortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**EndDateTime** | Pointer to **time.Time** |  | [optional] 
**ExpectedStartDateTime** | Pointer to **time.Time** |  | [optional] 
**ExpectedEndDateTime** | Pointer to **time.Time** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseCohortDto

`func NewCourseCohortDto() *CourseCohortDto`

NewCourseCohortDto instantiates a new CourseCohortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCohortDtoWithDefaults

`func NewCourseCohortDtoWithDefaults() *CourseCohortDto`

NewCourseCohortDtoWithDefaults instantiates a new CourseCohortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCohortDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCohortDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCohortDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCohortDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseCohortDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseCohortDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseCohortDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCohortDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCohortDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCohortDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseCohortDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseCohortDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CourseCohortDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseCohortDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseCohortDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CourseCohortDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CourseCohortDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CourseCohortDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDateTime

`func (o *CourseCohortDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CourseCohortDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CourseCohortDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CourseCohortDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *CourseCohortDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CourseCohortDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CourseCohortDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CourseCohortDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetExpectedStartDateTime

`func (o *CourseCohortDto) GetExpectedStartDateTime() time.Time`

GetExpectedStartDateTime returns the ExpectedStartDateTime field if non-nil, zero value otherwise.

### GetExpectedStartDateTimeOk

`func (o *CourseCohortDto) GetExpectedStartDateTimeOk() (*time.Time, bool)`

GetExpectedStartDateTimeOk returns a tuple with the ExpectedStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartDateTime

`func (o *CourseCohortDto) SetExpectedStartDateTime(v time.Time)`

SetExpectedStartDateTime sets ExpectedStartDateTime field to given value.

### HasExpectedStartDateTime

`func (o *CourseCohortDto) HasExpectedStartDateTime() bool`

HasExpectedStartDateTime returns a boolean if a field has been set.

### GetExpectedEndDateTime

`func (o *CourseCohortDto) GetExpectedEndDateTime() time.Time`

GetExpectedEndDateTime returns the ExpectedEndDateTime field if non-nil, zero value otherwise.

### GetExpectedEndDateTimeOk

`func (o *CourseCohortDto) GetExpectedEndDateTimeOk() (*time.Time, bool)`

GetExpectedEndDateTimeOk returns a tuple with the ExpectedEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedEndDateTime

`func (o *CourseCohortDto) SetExpectedEndDateTime(v time.Time)`

SetExpectedEndDateTime sets ExpectedEndDateTime field to given value.

### HasExpectedEndDateTime

`func (o *CourseCohortDto) HasExpectedEndDateTime() bool`

HasExpectedEndDateTime returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseCohortDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseCohortDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseCohortDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseCohortDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseCohortDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseCohortDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetTenantId

`func (o *CourseCohortDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseCohortDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseCohortDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseCohortDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseCohortDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseCohortDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


