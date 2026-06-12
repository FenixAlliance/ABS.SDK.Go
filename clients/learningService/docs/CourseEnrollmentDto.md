# CourseEnrollmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**CourseCohortId** | Pointer to **NullableString** |  | [optional] 
**StudentProfileId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseEnrollmentDto

`func NewCourseEnrollmentDto() *CourseEnrollmentDto`

NewCourseEnrollmentDto instantiates a new CourseEnrollmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseEnrollmentDtoWithDefaults

`func NewCourseEnrollmentDtoWithDefaults() *CourseEnrollmentDto`

NewCourseEnrollmentDtoWithDefaults instantiates a new CourseEnrollmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseEnrollmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseEnrollmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseEnrollmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseEnrollmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseEnrollmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseEnrollmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseEnrollmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseEnrollmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseEnrollmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseEnrollmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseEnrollmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseEnrollmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCourseId

`func (o *CourseEnrollmentDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseEnrollmentDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseEnrollmentDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseEnrollmentDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseEnrollmentDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseEnrollmentDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetCourseCohortId

`func (o *CourseEnrollmentDto) GetCourseCohortId() string`

GetCourseCohortId returns the CourseCohortId field if non-nil, zero value otherwise.

### GetCourseCohortIdOk

`func (o *CourseEnrollmentDto) GetCourseCohortIdOk() (*string, bool)`

GetCourseCohortIdOk returns a tuple with the CourseCohortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortId

`func (o *CourseEnrollmentDto) SetCourseCohortId(v string)`

SetCourseCohortId sets CourseCohortId field to given value.

### HasCourseCohortId

`func (o *CourseEnrollmentDto) HasCourseCohortId() bool`

HasCourseCohortId returns a boolean if a field has been set.

### SetCourseCohortIdNil

`func (o *CourseEnrollmentDto) SetCourseCohortIdNil(b bool)`

 SetCourseCohortIdNil sets the value for CourseCohortId to be an explicit nil

### UnsetCourseCohortId
`func (o *CourseEnrollmentDto) UnsetCourseCohortId()`

UnsetCourseCohortId ensures that no value is present for CourseCohortId, not even an explicit nil
### GetStudentProfileId

`func (o *CourseEnrollmentDto) GetStudentProfileId() string`

GetStudentProfileId returns the StudentProfileId field if non-nil, zero value otherwise.

### GetStudentProfileIdOk

`func (o *CourseEnrollmentDto) GetStudentProfileIdOk() (*string, bool)`

GetStudentProfileIdOk returns a tuple with the StudentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileId

`func (o *CourseEnrollmentDto) SetStudentProfileId(v string)`

SetStudentProfileId sets StudentProfileId field to given value.

### HasStudentProfileId

`func (o *CourseEnrollmentDto) HasStudentProfileId() bool`

HasStudentProfileId returns a boolean if a field has been set.

### SetStudentProfileIdNil

`func (o *CourseEnrollmentDto) SetStudentProfileIdNil(b bool)`

 SetStudentProfileIdNil sets the value for StudentProfileId to be an explicit nil

### UnsetStudentProfileId
`func (o *CourseEnrollmentDto) UnsetStudentProfileId()`

UnsetStudentProfileId ensures that no value is present for StudentProfileId, not even an explicit nil
### GetTenantId

`func (o *CourseEnrollmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseEnrollmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseEnrollmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseEnrollmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseEnrollmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseEnrollmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseEnrollmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseEnrollmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseEnrollmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseEnrollmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseEnrollmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseEnrollmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCourseCompletionCertificateId

`func (o *CourseEnrollmentDto) GetCourseCompletionCertificateId() string`

GetCourseCompletionCertificateId returns the CourseCompletionCertificateId field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateIdOk

`func (o *CourseEnrollmentDto) GetCourseCompletionCertificateIdOk() (*string, bool)`

GetCourseCompletionCertificateIdOk returns a tuple with the CourseCompletionCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateId

`func (o *CourseEnrollmentDto) SetCourseCompletionCertificateId(v string)`

SetCourseCompletionCertificateId sets CourseCompletionCertificateId field to given value.

### HasCourseCompletionCertificateId

`func (o *CourseEnrollmentDto) HasCourseCompletionCertificateId() bool`

HasCourseCompletionCertificateId returns a boolean if a field has been set.

### SetCourseCompletionCertificateIdNil

`func (o *CourseEnrollmentDto) SetCourseCompletionCertificateIdNil(b bool)`

 SetCourseCompletionCertificateIdNil sets the value for CourseCompletionCertificateId to be an explicit nil

### UnsetCourseCompletionCertificateId
`func (o *CourseEnrollmentDto) UnsetCourseCompletionCertificateId()`

UnsetCourseCompletionCertificateId ensures that no value is present for CourseCompletionCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


