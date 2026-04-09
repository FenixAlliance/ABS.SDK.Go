# CourseCompletionCertificateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**StudentProfileID** | **string** |  | 
**CourseEnrollmentID** | **string** |  | 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**CourseCompletionCertificateTemplateID** | Pointer to **NullableString** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseCompletionCertificateCreateDto

`func NewCourseCompletionCertificateCreateDto(studentProfileID string, courseEnrollmentID string, ) *CourseCompletionCertificateCreateDto`

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

### GetStudentProfileID

`func (o *CourseCompletionCertificateCreateDto) GetStudentProfileID() string`

GetStudentProfileID returns the StudentProfileID field if non-nil, zero value otherwise.

### GetStudentProfileIDOk

`func (o *CourseCompletionCertificateCreateDto) GetStudentProfileIDOk() (*string, bool)`

GetStudentProfileIDOk returns a tuple with the StudentProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentProfileID

`func (o *CourseCompletionCertificateCreateDto) SetStudentProfileID(v string)`

SetStudentProfileID sets StudentProfileID field to given value.


### GetCourseEnrollmentID

`func (o *CourseCompletionCertificateCreateDto) GetCourseEnrollmentID() string`

GetCourseEnrollmentID returns the CourseEnrollmentID field if non-nil, zero value otherwise.

### GetCourseEnrollmentIDOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseEnrollmentIDOk() (*string, bool)`

GetCourseEnrollmentIDOk returns a tuple with the CourseEnrollmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseEnrollmentID

`func (o *CourseCompletionCertificateCreateDto) SetCourseEnrollmentID(v string)`

SetCourseEnrollmentID sets CourseEnrollmentID field to given value.


### GetBusinessID

`func (o *CourseCompletionCertificateCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseCompletionCertificateCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseCompletionCertificateCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *CourseCompletionCertificateCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *CourseCompletionCertificateCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *CourseCompletionCertificateCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *CourseCompletionCertificateCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *CourseCompletionCertificateCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *CourseCompletionCertificateCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *CourseCompletionCertificateCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *CourseCompletionCertificateCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *CourseCompletionCertificateCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateCreateDto) GetCourseCompletionCertificateTemplateID() string`

GetCourseCompletionCertificateTemplateID returns the CourseCompletionCertificateTemplateID field if non-nil, zero value otherwise.

### GetCourseCompletionCertificateTemplateIDOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseCompletionCertificateTemplateIDOk() (*string, bool)`

GetCourseCompletionCertificateTemplateIDOk returns a tuple with the CourseCompletionCertificateTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateCreateDto) SetCourseCompletionCertificateTemplateID(v string)`

SetCourseCompletionCertificateTemplateID sets CourseCompletionCertificateTemplateID field to given value.

### HasCourseCompletionCertificateTemplateID

`func (o *CourseCompletionCertificateCreateDto) HasCourseCompletionCertificateTemplateID() bool`

HasCourseCompletionCertificateTemplateID returns a boolean if a field has been set.

### SetCourseCompletionCertificateTemplateIDNil

`func (o *CourseCompletionCertificateCreateDto) SetCourseCompletionCertificateTemplateIDNil(b bool)`

 SetCourseCompletionCertificateTemplateIDNil sets the value for CourseCompletionCertificateTemplateID to be an explicit nil

### UnsetCourseCompletionCertificateTemplateID
`func (o *CourseCompletionCertificateCreateDto) UnsetCourseCompletionCertificateTemplateID()`

UnsetCourseCompletionCertificateTemplateID ensures that no value is present for CourseCompletionCertificateTemplateID, not even an explicit nil
### GetCourseID

`func (o *CourseCompletionCertificateCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseCompletionCertificateCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseCompletionCertificateCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseCompletionCertificateCreateDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseCompletionCertificateCreateDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseCompletionCertificateCreateDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


