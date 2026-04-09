# CourseCompletionCertificateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**StudentProfileID** | Pointer to **NullableString** |  | [optional] 
**CourseEnrollmentID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateTemplateID** | Pointer to **NullableString** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
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
### GetStudentProfileID

`func (o *CourseCompletionCertificateDto) GetStudentProfileID() string`

GetStudentProfileID returns the StudentProfileID field if non-nil, zero value otherwise.

### GetStudentProfileIDOk

`func (o *CourseCompletionCertificateDto) GetStudentProfileIDOk() (*string, bool)`

GetStudentProfileIDOk returns a tuple with the StudentProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileID

`func (o *CourseCompletionCertificateDto) SetStudentProfileID(v string)`

SetStudentProfileID sets StudentProfileID field to given value.

### HasStudentProfileID

`func (o *CourseCompletionCertificateDto) HasStudentProfileID() bool`

HasStudentProfileID returns a boolean if a field has been set.

### SetStudentProfileIDNil

`func (o *CourseCompletionCertificateDto) SetStudentProfileIDNil(b bool)`

 SetStudentProfileIDNil sets the value for StudentProfileID to be an explicit nil

### UnsetStudentProfileID
`func (o *CourseCompletionCertificateDto) UnsetStudentProfileID()`

UnsetStudentProfileID ensures that no value is present for StudentProfileID, not even an explicit nil
### GetCourseEnrollmentID

`func (o *CourseCompletionCertificateDto) GetCourseEnrollmentID() string`

GetCourseEnrollmentID returns the CourseEnrollmentID field if non-nil, zero value otherwise.

### GetCourseEnrollmentIDOk

`func (o *CourseCompletionCertificateDto) GetCourseEnrollmentIDOk() (*string, bool)`

GetCourseEnrollmentIDOk returns a tuple with the CourseEnrollmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseEnrollmentID

`func (o *CourseCompletionCertificateDto) SetCourseEnrollmentID(v string)`

SetCourseEnrollmentID sets CourseEnrollmentID field to given value.

### HasCourseEnrollmentID

`func (o *CourseCompletionCertificateDto) HasCourseEnrollmentID() bool`

HasCourseEnrollmentID returns a boolean if a field has been set.

### SetCourseEnrollmentIDNil

`func (o *CourseCompletionCertificateDto) SetCourseEnrollmentIDNil(b bool)`

 SetCourseEnrollmentIDNil sets the value for CourseEnrollmentID to be an explicit nil

### UnsetCourseEnrollmentID
`func (o *CourseCompletionCertificateDto) UnsetCourseEnrollmentID()`

UnsetCourseEnrollmentID ensures that no value is present for CourseEnrollmentID, not even an explicit nil
### GetBusinessID

`func (o *CourseCompletionCertificateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseCompletionCertificateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseCompletionCertificateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *CourseCompletionCertificateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *CourseCompletionCertificateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *CourseCompletionCertificateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *CourseCompletionCertificateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *CourseCompletionCertificateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *CourseCompletionCertificateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *CourseCompletionCertificateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *CourseCompletionCertificateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *CourseCompletionCertificateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateDto) GetCourseCompletionCertificateTemplateID() string`

GetCourseCompletionCertificateTemplateID returns the CourseCompletionCertificateTemplateID field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateTemplateIDOk

`func (o *CourseCompletionCertificateDto) GetCourseCompletionCertificateTemplateIDOk() (*string, bool)`

GetCourseCompletionCertificateTemplateIDOk returns a tuple with the CourseCompletionCertificateTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateDto) SetCourseCompletionCertificateTemplateID(v string)`

SetCourseCompletionCertificateTemplateID sets CourseCompletionCertificateTemplateID field to given value.

### HasCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateDto) HasCourseCompletionCertificateTemplateID() bool`

HasCourseCompletionCertificateTemplateID returns a boolean if a field has been set.

### SetCourseCompletionCertificateTemplateIDNil

`func (o *CourseCompletionCertificateDto) SetCourseCompletionCertificateTemplateIDNil(b bool)`

 SetCourseCompletionCertificateTemplateIDNil sets the value for CourseCompletionCertificateTemplateID to be an explicit nil

### UnsetCourseCompletionCertificateTemplateID
`func (o *CourseCompletionCertificateDto) UnsetCourseCompletionCertificateTemplateID()`

UnsetCourseCompletionCertificateTemplateID ensures that no value is present for CourseCompletionCertificateTemplateID, not even an explicit nil
### GetCourseID

`func (o *CourseCompletionCertificateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseCompletionCertificateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseCompletionCertificateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseCompletionCertificateDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseCompletionCertificateDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseCompletionCertificateDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
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


