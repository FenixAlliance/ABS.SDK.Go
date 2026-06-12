# CourseCompletionCertificateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**StudentProfileId** | Pointer to **NullableString** |  | [optional] 
**CourseEnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateTemplateId** | Pointer to **NullableString** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**StudentName** | Pointer to **NullableString** |  | [optional] 
**StudentLastName** | Pointer to **NullableString** |  | [optional] 
**CourseTitle** | Pointer to **NullableString** |  | [optional] 
**TotalEffortInHours** | Pointer to **NullableFloat64** |  | [optional] 
**InstructorName** | Pointer to **NullableString** |  | [optional] 
**InstructorLastName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseCompletionCertificateDto

`func NewCourseCompletionCertificateDto() *CourseCompletionCertificateDto`

NewCourseCompletionCertificateDto instantiates a new CourseCompletionCertificateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCompletionCertificateDtoWithDefaults

`func NewCourseCompletionCertificateDtoWithDefaults() *CourseCompletionCertificateDto`

NewCourseCompletionCertificateDtoWithDefaults instantiates a new CourseCompletionCertificateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCompletionCertificateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCompletionCertificateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCompletionCertificateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCompletionCertificateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseCompletionCertificateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseCompletionCertificateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseCompletionCertificateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCompletionCertificateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCompletionCertificateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCompletionCertificateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseCompletionCertificateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseCompletionCertificateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetStudentProfileId

`func (o *CourseCompletionCertificateDto) GetStudentProfileId() string`

GetStudentProfileId returns the StudentProfileId field if non-nil, zero value otherwise.

### GetStudentProfileIdOk

`func (o *CourseCompletionCertificateDto) GetStudentProfileIdOk() (*string, bool)`

GetStudentProfileIdOk returns a tuple with the StudentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileId

`func (o *CourseCompletionCertificateDto) SetStudentProfileId(v string)`

SetStudentProfileId sets StudentProfileId field to given value.

### HasStudentProfileId

`func (o *CourseCompletionCertificateDto) HasStudentProfileId() bool`

HasStudentProfileId returns a boolean if a field has been set.

### SetStudentProfileIdNil

`func (o *CourseCompletionCertificateDto) SetStudentProfileIdNil(b bool)`

 SetStudentProfileIdNil sets the value for StudentProfileId to be an explicit nil

### UnsetStudentProfileId
`func (o *CourseCompletionCertificateDto) UnsetStudentProfileId()`

UnsetStudentProfileId ensures that no value is present for StudentProfileId, not even an explicit nil
### GetCourseEnrollmentId

`func (o *CourseCompletionCertificateDto) GetCourseEnrollmentId() string`

GetCourseEnrollmentId returns the CourseEnrollmentId field if non-nil, zero value otherwise.

### GetCourseEnrollmentIdOk

`func (o *CourseCompletionCertificateDto) GetCourseEnrollmentIdOk() (*string, bool)`

GetCourseEnrollmentIdOk returns a tuple with the CourseEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseEnrollmentId

`func (o *CourseCompletionCertificateDto) SetCourseEnrollmentId(v string)`

SetCourseEnrollmentId sets CourseEnrollmentId field to given value.

### HasCourseEnrollmentId

`func (o *CourseCompletionCertificateDto) HasCourseEnrollmentId() bool`

HasCourseEnrollmentId returns a boolean if a field has been set.

### SetCourseEnrollmentIdNil

`func (o *CourseCompletionCertificateDto) SetCourseEnrollmentIdNil(b bool)`

 SetCourseEnrollmentIdNil sets the value for CourseEnrollmentId to be an explicit nil

### UnsetCourseEnrollmentId
`func (o *CourseCompletionCertificateDto) UnsetCourseEnrollmentId()`

UnsetCourseEnrollmentId ensures that no value is present for CourseEnrollmentId, not even an explicit nil
### GetTenantId

`func (o *CourseCompletionCertificateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseCompletionCertificateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseCompletionCertificateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseCompletionCertificateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseCompletionCertificateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseCompletionCertificateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseCompletionCertificateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseCompletionCertificateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseCompletionCertificateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseCompletionCertificateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseCompletionCertificateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseCompletionCertificateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateDto) GetCourseCompletionCertificateTemplateId() string`

GetCourseCompletionCertificateTemplateId returns the CourseCompletionCertificateTemplateId field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateTemplateIdOk

`func (o *CourseCompletionCertificateDto) GetCourseCompletionCertificateTemplateIdOk() (*string, bool)`

GetCourseCompletionCertificateTemplateIdOk returns a tuple with the CourseCompletionCertificateTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateDto) SetCourseCompletionCertificateTemplateId(v string)`

SetCourseCompletionCertificateTemplateId sets CourseCompletionCertificateTemplateId field to given value.

### HasCourseCompletionCertificateTemplateId

`func (o *CourseCompletionCertificateDto) HasCourseCompletionCertificateTemplateId() bool`

HasCourseCompletionCertificateTemplateId returns a boolean if a field has been set.

### SetCourseCompletionCertificateTemplateIdNil

`func (o *CourseCompletionCertificateDto) SetCourseCompletionCertificateTemplateIdNil(b bool)`

 SetCourseCompletionCertificateTemplateIdNil sets the value for CourseCompletionCertificateTemplateId to be an explicit nil

### UnsetCourseCompletionCertificateTemplateId
`func (o *CourseCompletionCertificateDto) UnsetCourseCompletionCertificateTemplateId()`

UnsetCourseCompletionCertificateTemplateId ensures that no value is present for CourseCompletionCertificateTemplateId, not even an explicit nil
### GetCourseId

`func (o *CourseCompletionCertificateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseCompletionCertificateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseCompletionCertificateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseCompletionCertificateDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseCompletionCertificateDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseCompletionCertificateDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetStudentName

`func (o *CourseCompletionCertificateDto) GetStudentName() string`

GetStudentName returns the StudentName field if non-nil, zero value otherwise.

### GetStudentNameOk

`func (o *CourseCompletionCertificateDto) GetStudentNameOk() (*string, bool)`

GetStudentNameOk returns a tuple with the StudentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentName

`func (o *CourseCompletionCertificateDto) SetStudentName(v string)`

SetStudentName sets StudentName field to given value.

### HasStudentName

`func (o *CourseCompletionCertificateDto) HasStudentName() bool`

HasStudentName returns a boolean if a field has been set.

### SetStudentNameNil

`func (o *CourseCompletionCertificateDto) SetStudentNameNil(b bool)`

 SetStudentNameNil sets the value for StudentName to be an explicit nil

### UnsetStudentName
`func (o *CourseCompletionCertificateDto) UnsetStudentName()`

UnsetStudentName ensures that no value is present for StudentName, not even an explicit nil
### GetStudentLastName

`func (o *CourseCompletionCertificateDto) GetStudentLastName() string`

GetStudentLastName returns the StudentLastName field if non-nil, zero value otherwise.

### GetStudentLastNameOk

`func (o *CourseCompletionCertificateDto) GetStudentLastNameOk() (*string, bool)`

GetStudentLastNameOk returns a tuple with the StudentLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentLastName

`func (o *CourseCompletionCertificateDto) SetStudentLastName(v string)`

SetStudentLastName sets StudentLastName field to given value.

### HasStudentLastName

`func (o *CourseCompletionCertificateDto) HasStudentLastName() bool`

HasStudentLastName returns a boolean if a field has been set.

### SetStudentLastNameNil

`func (o *CourseCompletionCertificateDto) SetStudentLastNameNil(b bool)`

 SetStudentLastNameNil sets the value for StudentLastName to be an explicit nil

### UnsetStudentLastName
`func (o *CourseCompletionCertificateDto) UnsetStudentLastName()`

UnsetStudentLastName ensures that no value is present for StudentLastName, not even an explicit nil
### GetCourseTitle

`func (o *CourseCompletionCertificateDto) GetCourseTitle() string`

GetCourseTitle returns the CourseTitle field if non-nil, zero value otherwise.

### GetCourseTitleOk

`func (o *CourseCompletionCertificateDto) GetCourseTitleOk() (*string, bool)`

GetCourseTitleOk returns a tuple with the CourseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTitle

`func (o *CourseCompletionCertificateDto) SetCourseTitle(v string)`

SetCourseTitle sets CourseTitle field to given value.

### HasCourseTitle

`func (o *CourseCompletionCertificateDto) HasCourseTitle() bool`

HasCourseTitle returns a boolean if a field has been set.

### SetCourseTitleNil

`func (o *CourseCompletionCertificateDto) SetCourseTitleNil(b bool)`

 SetCourseTitleNil sets the value for CourseTitle to be an explicit nil

### UnsetCourseTitle
`func (o *CourseCompletionCertificateDto) UnsetCourseTitle()`

UnsetCourseTitle ensures that no value is present for CourseTitle, not even an explicit nil
### GetTotalEffortInHours

`func (o *CourseCompletionCertificateDto) GetTotalEffortInHours() float64`

GetTotalEffortInHours returns the TotalEffortInHours field if non-nil, zero value otherwise.

### GetTotalEffortInHoursOk

`func (o *CourseCompletionCertificateDto) GetTotalEffortInHoursOk() (*float64, bool)`

GetTotalEffortInHoursOk returns a tuple with the TotalEffortInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEffortInHours

`func (o *CourseCompletionCertificateDto) SetTotalEffortInHours(v float64)`

SetTotalEffortInHours sets TotalEffortInHours field to given value.

### HasTotalEffortInHours

`func (o *CourseCompletionCertificateDto) HasTotalEffortInHours() bool`

HasTotalEffortInHours returns a boolean if a field has been set.

### SetTotalEffortInHoursNil

`func (o *CourseCompletionCertificateDto) SetTotalEffortInHoursNil(b bool)`

 SetTotalEffortInHoursNil sets the value for TotalEffortInHours to be an explicit nil

### UnsetTotalEffortInHours
`func (o *CourseCompletionCertificateDto) UnsetTotalEffortInHours()`

UnsetTotalEffortInHours ensures that no value is present for TotalEffortInHours, not even an explicit nil
### GetInstructorName

`func (o *CourseCompletionCertificateDto) GetInstructorName() string`

GetInstructorName returns the InstructorName field if non-nil, zero value otherwise.

### GetInstructorNameOk

`func (o *CourseCompletionCertificateDto) GetInstructorNameOk() (*string, bool)`

GetInstructorNameOk returns a tuple with the InstructorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorName

`func (o *CourseCompletionCertificateDto) SetInstructorName(v string)`

SetInstructorName sets InstructorName field to given value.

### HasInstructorName

`func (o *CourseCompletionCertificateDto) HasInstructorName() bool`

HasInstructorName returns a boolean if a field has been set.

### SetInstructorNameNil

`func (o *CourseCompletionCertificateDto) SetInstructorNameNil(b bool)`

 SetInstructorNameNil sets the value for InstructorName to be an explicit nil

### UnsetInstructorName
`func (o *CourseCompletionCertificateDto) UnsetInstructorName()`

UnsetInstructorName ensures that no value is present for InstructorName, not even an explicit nil
### GetInstructorLastName

`func (o *CourseCompletionCertificateDto) GetInstructorLastName() string`

GetInstructorLastName returns the InstructorLastName field if non-nil, zero value otherwise.

### GetInstructorLastNameOk

`func (o *CourseCompletionCertificateDto) GetInstructorLastNameOk() (*string, bool)`

GetInstructorLastNameOk returns a tuple with the InstructorLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorLastName

`func (o *CourseCompletionCertificateDto) SetInstructorLastName(v string)`

SetInstructorLastName sets InstructorLastName field to given value.

### HasInstructorLastName

`func (o *CourseCompletionCertificateDto) HasInstructorLastName() bool`

HasInstructorLastName returns a boolean if a field has been set.

### SetInstructorLastNameNil

`func (o *CourseCompletionCertificateDto) SetInstructorLastNameNil(b bool)`

 SetInstructorLastNameNil sets the value for InstructorLastName to be an explicit nil

### UnsetInstructorLastName
`func (o *CourseCompletionCertificateDto) UnsetInstructorLastName()`

UnsetInstructorLastName ensures that no value is present for InstructorLastName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


