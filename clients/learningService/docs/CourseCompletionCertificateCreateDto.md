# CourseCompletionCertificateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**StudentProfileId** | **string** |  | 
**CourseEnrollmentId** | **string** |  | 
**CourseCompletionCertificateTemplateId** | Pointer to **NullableString** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseCompletionCertificateCreateDto

`func NewCourseCompletionCertificateCreateDto(studentProfileId string, courseEnrollmentId string, ) *CourseCompletionCertificateCreateDto`

NewCourseCompletionCertificateCreateDto instantiates a new CourseCompletionCertificateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCompletionCertificateCreateDtoWithDefaults

`func NewCourseCompletionCertificateCreateDtoWithDefaults() *CourseCompletionCertificateCreateDto`

NewCourseCompletionCertificateCreateDtoWithDefaults instantiates a new CourseCompletionCertificateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCompletionCertificateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCompletionCertificateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCompletionCertificateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCompletionCertificateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseCompletionCertificateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCompletionCertificateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCompletionCertificateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCompletionCertificateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStudentProfileId

`func (o *CourseCompletionCertificateCreateDto) GetStudentProfileId() string`

GetStudentProfileId returns the StudentProfileId field if non-nil, zero value otherwise.

### GetStudentProfileIdOk

`func (o *CourseCompletionCertificateCreateDto) GetStudentProfileIdOk() (*string, bool)`

GetStudentProfileIdOk returns a tuple with the StudentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileId

`func (o *CourseCompletionCertificateCreateDto) SetStudentProfileId(v string)`

SetStudentProfileId sets StudentProfileId field to given value.


### GetCourseEnrollmentId

`func (o *CourseCompletionCertificateCreateDto) GetCourseEnrollmentId() string`

GetCourseEnrollmentId returns the CourseEnrollmentId field if non-nil, zero value otherwise.

### GetCourseEnrollmentIdOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseEnrollmentIdOk() (*string, bool)`

GetCourseEnrollmentIdOk returns a tuple with the CourseEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseEnrollmentId

`func (o *CourseCompletionCertificateCreateDto) SetCourseEnrollmentId(v string)`

SetCourseEnrollmentId sets CourseEnrollmentId field to given value.


### GetCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateCreateDto) GetCourseCompletionCertificateTemplateId() string`

GetCourseCompletionCertificateTemplateId returns the CourseCompletionCertificateTemplateId field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateTemplateIdOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseCompletionCertificateTemplateIdOk() (*string, bool)`

GetCourseCompletionCertificateTemplateIdOk returns a tuple with the CourseCompletionCertificateTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateCreateDto) SetCourseCompletionCertificateTemplateId(v string)`

SetCourseCompletionCertificateTemplateId sets CourseCompletionCertificateTemplateId field to given value.

### HasCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateCreateDto) HasCourseCompletionCertificateTemplateId() bool`

HasCourseCompletionCertificateTemplateId returns a boolean if a field has been set.

### SetCourseCompletionCertificateTemplateIdNil

`func (o *CourseCompletionCertificateCreateDto) SetCourseCompletionCertificateTemplateIdNil(b bool)`

 SetCourseCompletionCertificateTemplateIdNil sets the value for CourseCompletionCertificateTemplateId to be an explicit nil

### UnsetCourseCompletionCertificateTemplateId
`func (o *CourseCompletionCertificateCreateDto) UnsetCourseCompletionCertificateTemplateId()`

UnsetCourseCompletionCertificateTemplateId ensures that no value is present for CourseCompletionCertificateTemplateId, not even an explicit nil
### GetCourseId

`func (o *CourseCompletionCertificateCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseCompletionCertificateCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseCompletionCertificateCreateDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseCompletionCertificateCreateDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseCompletionCertificateCreateDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


