# CourseEnrollmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CourseCohortID** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseEnrollmentUpdateDto

`func NewCourseEnrollmentUpdateDto() *CourseEnrollmentUpdateDto`

NewCourseEnrollmentUpdateDto instantiates a new CourseEnrollmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseEnrollmentUpdateDtoWithDefaults

`func NewCourseEnrollmentUpdateDtoWithDefaults() *CourseEnrollmentUpdateDto`

NewCourseEnrollmentUpdateDtoWithDefaults instantiates a new CourseEnrollmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseEnrollmentUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseEnrollmentUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseEnrollmentUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseEnrollmentUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseEnrollmentUpdateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseEnrollmentUpdateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseEnrollmentUpdateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseEnrollmentUpdateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCourseCohortID

`func (o *CourseEnrollmentUpdateDto) GetCourseCohortID() string`

GetCourseCohortID returns the CourseCohortID field if non-nil, zero value otherwise.

### GetCourseCohortIDOk

`func (o *CourseEnrollmentUpdateDto) GetCourseCohortIDOk() (*string, bool)`

GetCourseCohortIDOk returns a tuple with the CourseCohortID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortID

`func (o *CourseEnrollmentUpdateDto) SetCourseCohortID(v string)`

SetCourseCohortID sets CourseCohortID field to given value.

### HasCourseCohortID

`func (o *CourseEnrollmentUpdateDto) HasCourseCohortID() bool`

HasCourseCohortID returns a boolean if a field has been set.

### SetCourseCohortIDNil

`func (o *CourseEnrollmentUpdateDto) SetCourseCohortIDNil(b bool)`

 SetCourseCohortIDNil sets the value for CourseCohortID to be an explicit nil

### UnsetCourseCohortID
`func (o *CourseEnrollmentUpdateDto) UnsetCourseCohortID()`

UnsetCourseCohortID ensures that no value is present for CourseCohortID, not even an explicit nil
### GetCourseCompletionCertificateID

`func (o *CourseEnrollmentUpdateDto) GetCourseCompletionCertificateID() string`

GetCourseCompletionCertificateID returns the CourseCompletionCertificateID field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateIDOk

`func (o *CourseEnrollmentUpdateDto) GetCourseCompletionCertificateIDOk() (*string, bool)`

GetCourseCompletionCertificateIDOk returns a tuple with the CourseCompletionCertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateID

`func (o *CourseEnrollmentUpdateDto) SetCourseCompletionCertificateID(v string)`

SetCourseCompletionCertificateID sets CourseCompletionCertificateID field to given value.

### HasCourseCompletionCertificateID

`func (o *CourseEnrollmentUpdateDto) HasCourseCompletionCertificateID() bool`

HasCourseCompletionCertificateID returns a boolean if a field has been set.

### SetCourseCompletionCertificateIDNil

`func (o *CourseEnrollmentUpdateDto) SetCourseCompletionCertificateIDNil(b bool)`

 SetCourseCompletionCertificateIDNil sets the value for CourseCompletionCertificateID to be an explicit nil

### UnsetCourseCompletionCertificateID
`func (o *CourseEnrollmentUpdateDto) UnsetCourseCompletionCertificateID()`

UnsetCourseCompletionCertificateID ensures that no value is present for CourseCompletionCertificateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


