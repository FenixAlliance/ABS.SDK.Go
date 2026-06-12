# CourseAssignmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Points** | Pointer to **float64** |  | [optional] 
**DueDateTime** | Pointer to **time.Time** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**CourseUnitId** | Pointer to **NullableString** |  | [optional] 
**CourseSectionId** | Pointer to **NullableString** |  | [optional] 
**CourseCohortId** | Pointer to **NullableString** |  | [optional] 
**CourseAssignmentTypeId** | Pointer to **NullableString** |  | [optional] 
**AsignToAllCohorts** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseAssignmentDto

`func NewCourseAssignmentDto() *CourseAssignmentDto`

NewCourseAssignmentDto instantiates a new CourseAssignmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentDtoWithDefaults

`func NewCourseAssignmentDtoWithDefaults() *CourseAssignmentDto`

NewCourseAssignmentDtoWithDefaults instantiates a new CourseAssignmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseAssignmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseAssignmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseAssignmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseAssignmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseAssignmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseAssignmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseAssignmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseAssignmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseAssignmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseAssignmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseAssignmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseAssignmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CourseAssignmentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseAssignmentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseAssignmentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseAssignmentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseAssignmentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseAssignmentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseAssignmentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseAssignmentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseAssignmentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseAssignmentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseAssignmentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseAssignmentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *CourseAssignmentDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CourseAssignmentDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CourseAssignmentDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CourseAssignmentDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CourseAssignmentDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CourseAssignmentDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetPoints

`func (o *CourseAssignmentDto) GetPoints() float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CourseAssignmentDto) GetPointsOk() (*float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CourseAssignmentDto) SetPoints(v float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CourseAssignmentDto) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetDueDateTime

`func (o *CourseAssignmentDto) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *CourseAssignmentDto) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *CourseAssignmentDto) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *CourseAssignmentDto) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### GetCourseId

`func (o *CourseAssignmentDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseAssignmentDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseAssignmentDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseAssignmentDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseAssignmentDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseAssignmentDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetCourseUnitId

`func (o *CourseAssignmentDto) GetCourseUnitId() string`

GetCourseUnitId returns the CourseUnitId field if non-nil, zero value otherwise.

### GetCourseUnitIdOk

`func (o *CourseAssignmentDto) GetCourseUnitIdOk() (*string, bool)`

GetCourseUnitIdOk returns a tuple with the CourseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitId

`func (o *CourseAssignmentDto) SetCourseUnitId(v string)`

SetCourseUnitId sets CourseUnitId field to given value.

### HasCourseUnitId

`func (o *CourseAssignmentDto) HasCourseUnitId() bool`

HasCourseUnitId returns a boolean if a field has been set.

### SetCourseUnitIdNil

`func (o *CourseAssignmentDto) SetCourseUnitIdNil(b bool)`

 SetCourseUnitIdNil sets the value for CourseUnitId to be an explicit nil

### UnsetCourseUnitId
`func (o *CourseAssignmentDto) UnsetCourseUnitId()`

UnsetCourseUnitId ensures that no value is present for CourseUnitId, not even an explicit nil
### GetCourseSectionId

`func (o *CourseAssignmentDto) GetCourseSectionId() string`

GetCourseSectionId returns the CourseSectionId field if non-nil, zero value otherwise.

### GetCourseSectionIdOk

`func (o *CourseAssignmentDto) GetCourseSectionIdOk() (*string, bool)`

GetCourseSectionIdOk returns a tuple with the CourseSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseSectionId

`func (o *CourseAssignmentDto) SetCourseSectionId(v string)`

SetCourseSectionId sets CourseSectionId field to given value.

### HasCourseSectionId

`func (o *CourseAssignmentDto) HasCourseSectionId() bool`

HasCourseSectionId returns a boolean if a field has been set.

### SetCourseSectionIdNil

`func (o *CourseAssignmentDto) SetCourseSectionIdNil(b bool)`

 SetCourseSectionIdNil sets the value for CourseSectionId to be an explicit nil

### UnsetCourseSectionId
`func (o *CourseAssignmentDto) UnsetCourseSectionId()`

UnsetCourseSectionId ensures that no value is present for CourseSectionId, not even an explicit nil
### GetCourseCohortId

`func (o *CourseAssignmentDto) GetCourseCohortId() string`

GetCourseCohortId returns the CourseCohortId field if non-nil, zero value otherwise.

### GetCourseCohortIdOk

`func (o *CourseAssignmentDto) GetCourseCohortIdOk() (*string, bool)`

GetCourseCohortIdOk returns a tuple with the CourseCohortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortId

`func (o *CourseAssignmentDto) SetCourseCohortId(v string)`

SetCourseCohortId sets CourseCohortId field to given value.

### HasCourseCohortId

`func (o *CourseAssignmentDto) HasCourseCohortId() bool`

HasCourseCohortId returns a boolean if a field has been set.

### SetCourseCohortIdNil

`func (o *CourseAssignmentDto) SetCourseCohortIdNil(b bool)`

 SetCourseCohortIdNil sets the value for CourseCohortId to be an explicit nil

### UnsetCourseCohortId
`func (o *CourseAssignmentDto) UnsetCourseCohortId()`

UnsetCourseCohortId ensures that no value is present for CourseCohortId, not even an explicit nil
### GetCourseAssignmentTypeId

`func (o *CourseAssignmentDto) GetCourseAssignmentTypeId() string`

GetCourseAssignmentTypeId returns the CourseAssignmentTypeId field if non-nil, zero value otherwise.

### GetCourseAssignmentTypeIdOk

`func (o *CourseAssignmentDto) GetCourseAssignmentTypeIdOk() (*string, bool)`

GetCourseAssignmentTypeIdOk returns a tuple with the CourseAssignmentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentTypeId

`func (o *CourseAssignmentDto) SetCourseAssignmentTypeId(v string)`

SetCourseAssignmentTypeId sets CourseAssignmentTypeId field to given value.

### HasCourseAssignmentTypeId

`func (o *CourseAssignmentDto) HasCourseAssignmentTypeId() bool`

HasCourseAssignmentTypeId returns a boolean if a field has been set.

### SetCourseAssignmentTypeIdNil

`func (o *CourseAssignmentDto) SetCourseAssignmentTypeIdNil(b bool)`

 SetCourseAssignmentTypeIdNil sets the value for CourseAssignmentTypeId to be an explicit nil

### UnsetCourseAssignmentTypeId
`func (o *CourseAssignmentDto) UnsetCourseAssignmentTypeId()`

UnsetCourseAssignmentTypeId ensures that no value is present for CourseAssignmentTypeId, not even an explicit nil
### GetAsignToAllCohorts

`func (o *CourseAssignmentDto) GetAsignToAllCohorts() bool`

GetAsignToAllCohorts returns the AsignToAllCohorts field if non-nil, zero value otherwise.

### GetAsignToAllCohortsOk

`func (o *CourseAssignmentDto) GetAsignToAllCohortsOk() (*bool, bool)`

GetAsignToAllCohortsOk returns a tuple with the AsignToAllCohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsignToAllCohorts

`func (o *CourseAssignmentDto) SetAsignToAllCohorts(v bool)`

SetAsignToAllCohorts sets AsignToAllCohorts field to given value.

### HasAsignToAllCohorts

`func (o *CourseAssignmentDto) HasAsignToAllCohorts() bool`

HasAsignToAllCohorts returns a boolean if a field has been set.

### GetResources

`func (o *CourseAssignmentDto) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CourseAssignmentDto) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CourseAssignmentDto) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CourseAssignmentDto) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *CourseAssignmentDto) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *CourseAssignmentDto) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetTenantId

`func (o *CourseAssignmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseAssignmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseAssignmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseAssignmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseAssignmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseAssignmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseAssignmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseAssignmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseAssignmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseAssignmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseAssignmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseAssignmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


