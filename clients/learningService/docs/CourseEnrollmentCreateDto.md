# CourseEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**CourseCohortId** | Pointer to **NullableString** |  | [optional] 
**StudentProfileId** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseEnrollmentCreateDto

`func NewCourseEnrollmentCreateDto() *CourseEnrollmentCreateDto`

NewCourseEnrollmentCreateDto instantiates a new CourseEnrollmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseEnrollmentCreateDtoWithDefaults

`func NewCourseEnrollmentCreateDtoWithDefaults() *CourseEnrollmentCreateDto`

NewCourseEnrollmentCreateDtoWithDefaults instantiates a new CourseEnrollmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseEnrollmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseEnrollmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseEnrollmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseEnrollmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseEnrollmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseEnrollmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseEnrollmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseEnrollmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCourseId

`func (o *CourseEnrollmentCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseEnrollmentCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseEnrollmentCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseEnrollmentCreateDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseEnrollmentCreateDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseEnrollmentCreateDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetCourseCohortId

`func (o *CourseEnrollmentCreateDto) GetCourseCohortId() string`

GetCourseCohortId returns the CourseCohortId field if non-nil, zero value otherwise.

### GetCourseCohortIdOk

`func (o *CourseEnrollmentCreateDto) GetCourseCohortIdOk() (*string, bool)`

GetCourseCohortIdOk returns a tuple with the CourseCohortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortId

`func (o *CourseEnrollmentCreateDto) SetCourseCohortId(v string)`

SetCourseCohortId sets CourseCohortId field to given value.

### HasCourseCohortId

`func (o *CourseEnrollmentCreateDto) HasCourseCohortId() bool`

HasCourseCohortId returns a boolean if a field has been set.

### SetCourseCohortIdNil

`func (o *CourseEnrollmentCreateDto) SetCourseCohortIdNil(b bool)`

 SetCourseCohortIdNil sets the value for CourseCohortId to be an explicit nil

### UnsetCourseCohortId
`func (o *CourseEnrollmentCreateDto) UnsetCourseCohortId()`

UnsetCourseCohortId ensures that no value is present for CourseCohortId, not even an explicit nil
### GetStudentProfileId

`func (o *CourseEnrollmentCreateDto) GetStudentProfileId() string`

GetStudentProfileId returns the StudentProfileId field if non-nil, zero value otherwise.

### GetStudentProfileIdOk

`func (o *CourseEnrollmentCreateDto) GetStudentProfileIdOk() (*string, bool)`

GetStudentProfileIdOk returns a tuple with the StudentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileId

`func (o *CourseEnrollmentCreateDto) SetStudentProfileId(v string)`

SetStudentProfileId sets StudentProfileId field to given value.

### HasStudentProfileId

`func (o *CourseEnrollmentCreateDto) HasStudentProfileId() bool`

HasStudentProfileId returns a boolean if a field has been set.

### SetStudentProfileIdNil

`func (o *CourseEnrollmentCreateDto) SetStudentProfileIdNil(b bool)`

 SetStudentProfileIdNil sets the value for StudentProfileId to be an explicit nil

### UnsetStudentProfileId
`func (o *CourseEnrollmentCreateDto) UnsetStudentProfileId()`

UnsetStudentProfileId ensures that no value is present for StudentProfileId, not even an explicit nil
### GetCourseCompletionCertificateId

`func (o *CourseEnrollmentCreateDto) GetCourseCompletionCertificateId() string`

GetCourseCompletionCertificateId returns the CourseCompletionCertificateId field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateIdOk

`func (o *CourseEnrollmentCreateDto) GetCourseCompletionCertificateIdOk() (*string, bool)`

GetCourseCompletionCertificateIdOk returns a tuple with the CourseCompletionCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateId

`func (o *CourseEnrollmentCreateDto) SetCourseCompletionCertificateId(v string)`

SetCourseCompletionCertificateId sets CourseCompletionCertificateId field to given value.

### HasCourseCompletionCertificateId

`func (o *CourseEnrollmentCreateDto) HasCourseCompletionCertificateId() bool`

HasCourseCompletionCertificateId returns a boolean if a field has been set.

### SetCourseCompletionCertificateIdNil

`func (o *CourseEnrollmentCreateDto) SetCourseCompletionCertificateIdNil(b bool)`

 SetCourseCompletionCertificateIdNil sets the value for CourseCompletionCertificateId to be an explicit nil

### UnsetCourseCompletionCertificateId
`func (o *CourseEnrollmentCreateDto) UnsetCourseCompletionCertificateId()`

UnsetCourseCompletionCertificateId ensures that no value is present for CourseCompletionCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


