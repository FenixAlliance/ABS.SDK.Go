# CourseEnrollmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**CourseCohortID** | Pointer to **NullableString** |  | [optional] 
**StudentProfileID** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateID** | Pointer to **NullableString** |  | [optional] 

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

### GetCourseID

`func (o *CourseEnrollmentCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseEnrollmentCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseEnrollmentCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseEnrollmentCreateDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseEnrollmentCreateDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseEnrollmentCreateDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetCourseCohortID

`func (o *CourseEnrollmentCreateDto) GetCourseCohortID() string`

GetCourseCohortID returns the CourseCohortID field if non-nil, zero value otherwise.

### GetCourseCohortIDOk

`func (o *CourseEnrollmentCreateDto) GetCourseCohortIDOk() (*string, bool)`

GetCourseCohortIDOk returns a tuple with the CourseCohortID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortID

`func (o *CourseEnrollmentCreateDto) SetCourseCohortID(v string)`

SetCourseCohortID sets CourseCohortID field to given value.

### HasCourseCohortID

`func (o *CourseEnrollmentCreateDto) HasCourseCohortID() bool`

HasCourseCohortID returns a boolean if a field has been set.

### SetCourseCohortIDNil

`func (o *CourseEnrollmentCreateDto) SetCourseCohortIDNil(b bool)`

 SetCourseCohortIDNil sets the value for CourseCohortID to be an explicit nil

### UnsetCourseCohortID
`func (o *CourseEnrollmentCreateDto) UnsetCourseCohortID()`

UnsetCourseCohortID ensures that no value is present for CourseCohortID, not even an explicit nil
### GetStudentProfileID

`func (o *CourseEnrollmentCreateDto) GetStudentProfileID() string`

GetStudentProfileID returns the StudentProfileID field if non-nil, zero value otherwise.

### GetStudentProfileIDOk

`func (o *CourseEnrollmentCreateDto) GetStudentProfileIDOk() (*string, bool)`

GetStudentProfileIDOk returns a tuple with the StudentProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileID

`func (o *CourseEnrollmentCreateDto) SetStudentProfileID(v string)`

SetStudentProfileID sets StudentProfileID field to given value.

### HasStudentProfileID

`func (o *CourseEnrollmentCreateDto) HasStudentProfileID() bool`

HasStudentProfileID returns a boolean if a field has been set.

### SetStudentProfileIDNil

`func (o *CourseEnrollmentCreateDto) SetStudentProfileIDNil(b bool)`

 SetStudentProfileIDNil sets the value for StudentProfileID to be an explicit nil

### UnsetStudentProfileID
`func (o *CourseEnrollmentCreateDto) UnsetStudentProfileID()`

UnsetStudentProfileID ensures that no value is present for StudentProfileID, not even an explicit nil
### GetCourseCompletionCertificateID

`func (o *CourseEnrollmentCreateDto) GetCourseCompletionCertificateID() string`

GetCourseCompletionCertificateID returns the CourseCompletionCertificateID field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateIDOk

`func (o *CourseEnrollmentCreateDto) GetCourseCompletionCertificateIDOk() (*string, bool)`

GetCourseCompletionCertificateIDOk returns a tuple with the CourseCompletionCertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateID

`func (o *CourseEnrollmentCreateDto) SetCourseCompletionCertificateID(v string)`

SetCourseCompletionCertificateID sets CourseCompletionCertificateID field to given value.

### HasCourseCompletionCertificateID

`func (o *CourseEnrollmentCreateDto) HasCourseCompletionCertificateID() bool`

HasCourseCompletionCertificateID returns a boolean if a field has been set.

### SetCourseCompletionCertificateIDNil

`func (o *CourseEnrollmentCreateDto) SetCourseCompletionCertificateIDNil(b bool)`

 SetCourseCompletionCertificateIDNil sets the value for CourseCompletionCertificateID to be an explicit nil

### UnsetCourseCompletionCertificateID
`func (o *CourseEnrollmentCreateDto) UnsetCourseCompletionCertificateID()`

UnsetCourseCompletionCertificateID ensures that no value is present for CourseCompletionCertificateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


