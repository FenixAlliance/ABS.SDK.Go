# CourseEnrollmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**CourseCohortID** | Pointer to **NullableString** |  | [optional] 
**StudentProfileID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateID** | Pointer to **NullableString** |  | [optional] 

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
### GetCourseID

`func (o *CourseEnrollmentDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseEnrollmentDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseEnrollmentDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseEnrollmentDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseEnrollmentDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseEnrollmentDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetCourseCohortID

`func (o *CourseEnrollmentDto) GetCourseCohortID() string`

GetCourseCohortID returns the CourseCohortID field if non-nil, zero value otherwise.

### GetCourseCohortIDOk

`func (o *CourseEnrollmentDto) GetCourseCohortIDOk() (*string, bool)`

GetCourseCohortIDOk returns a tuple with the CourseCohortID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortID

`func (o *CourseEnrollmentDto) SetCourseCohortID(v string)`

SetCourseCohortID sets CourseCohortID field to given value.

### HasCourseCohortID

`func (o *CourseEnrollmentDto) HasCourseCohortID() bool`

HasCourseCohortID returns a boolean if a field has been set.

### SetCourseCohortIDNil

`func (o *CourseEnrollmentDto) SetCourseCohortIDNil(b bool)`

 SetCourseCohortIDNil sets the value for CourseCohortID to be an explicit nil

### UnsetCourseCohortID
`func (o *CourseEnrollmentDto) UnsetCourseCohortID()`

UnsetCourseCohortID ensures that no value is present for CourseCohortID, not even an explicit nil
### GetStudentProfileID

`func (o *CourseEnrollmentDto) GetStudentProfileID() string`

GetStudentProfileID returns the StudentProfileID field if non-nil, zero value otherwise.

### GetStudentProfileIDOk

`func (o *CourseEnrollmentDto) GetStudentProfileIDOk() (*string, bool)`

GetStudentProfileIDOk returns a tuple with the StudentProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileID

`func (o *CourseEnrollmentDto) SetStudentProfileID(v string)`

SetStudentProfileID sets StudentProfileID field to given value.

### HasStudentProfileID

`func (o *CourseEnrollmentDto) HasStudentProfileID() bool`

HasStudentProfileID returns a boolean if a field has been set.

### SetStudentProfileIDNil

`func (o *CourseEnrollmentDto) SetStudentProfileIDNil(b bool)`

 SetStudentProfileIDNil sets the value for StudentProfileID to be an explicit nil

### UnsetStudentProfileID
`func (o *CourseEnrollmentDto) UnsetStudentProfileID()`

UnsetStudentProfileID ensures that no value is present for StudentProfileID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *CourseEnrollmentDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *CourseEnrollmentDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *CourseEnrollmentDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *CourseEnrollmentDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *CourseEnrollmentDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *CourseEnrollmentDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetCourseCompletionCertificateID

`func (o *CourseEnrollmentDto) GetCourseCompletionCertificateID() string`

GetCourseCompletionCertificateID returns the CourseCompletionCertificateID field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateIDOk

`func (o *CourseEnrollmentDto) GetCourseCompletionCertificateIDOk() (*string, bool)`

GetCourseCompletionCertificateIDOk returns a tuple with the CourseCompletionCertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateID

`func (o *CourseEnrollmentDto) SetCourseCompletionCertificateID(v string)`

SetCourseCompletionCertificateID sets CourseCompletionCertificateID field to given value.

### HasCourseCompletionCertificateID

`func (o *CourseEnrollmentDto) HasCourseCompletionCertificateID() bool`

HasCourseCompletionCertificateID returns a boolean if a field has been set.

### SetCourseCompletionCertificateIDNil

`func (o *CourseEnrollmentDto) SetCourseCompletionCertificateIDNil(b bool)`

 SetCourseCompletionCertificateIDNil sets the value for CourseCompletionCertificateID to be an explicit nil

### UnsetCourseCompletionCertificateID
`func (o *CourseEnrollmentDto) UnsetCourseCompletionCertificateID()`

UnsetCourseCompletionCertificateID ensures that no value is present for CourseCompletionCertificateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


