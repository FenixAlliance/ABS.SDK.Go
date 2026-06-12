# CourseAssignmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Points** | Pointer to **float64** |  | [optional] 
**CourseId** | **string** |  | 
**CourseUnitId** | Pointer to **NullableString** |  | [optional] 
**CourseCohortId** | Pointer to **NullableString** |  | [optional] 
**CourseAssignmentTypeId** | Pointer to **NullableString** |  | [optional] 
**DueDateTime** | Pointer to **NullableTime** |  | [optional] 
**AsignToAllCohorts** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseAssignmentCreateDto

`func NewCourseAssignmentCreateDto(title string, courseId string, ) *CourseAssignmentCreateDto`

NewCourseAssignmentCreateDto instantiates a new CourseAssignmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentCreateDtoWithDefaults

`func NewCourseAssignmentCreateDtoWithDefaults() *CourseAssignmentCreateDto`

NewCourseAssignmentCreateDtoWithDefaults instantiates a new CourseAssignmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseAssignmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseAssignmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseAssignmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseAssignmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseAssignmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseAssignmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseAssignmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseAssignmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseAssignmentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseAssignmentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseAssignmentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseAssignmentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseAssignmentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseAssignmentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseAssignmentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseAssignmentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseAssignmentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *CourseAssignmentCreateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CourseAssignmentCreateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CourseAssignmentCreateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CourseAssignmentCreateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CourseAssignmentCreateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CourseAssignmentCreateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetPoints

`func (o *CourseAssignmentCreateDto) GetPoints() float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CourseAssignmentCreateDto) GetPointsOk() (*float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CourseAssignmentCreateDto) SetPoints(v float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CourseAssignmentCreateDto) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetCourseId

`func (o *CourseAssignmentCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseAssignmentCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseAssignmentCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.


### GetCourseUnitId

`func (o *CourseAssignmentCreateDto) GetCourseUnitId() string`

GetCourseUnitId returns the CourseUnitId field if non-nil, zero value otherwise.

### GetCourseUnitIdOk

`func (o *CourseAssignmentCreateDto) GetCourseUnitIdOk() (*string, bool)`

GetCourseUnitIdOk returns a tuple with the CourseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitId

`func (o *CourseAssignmentCreateDto) SetCourseUnitId(v string)`

SetCourseUnitId sets CourseUnitId field to given value.

### HasCourseUnitId

`func (o *CourseAssignmentCreateDto) HasCourseUnitId() bool`

HasCourseUnitId returns a boolean if a field has been set.

### SetCourseUnitIdNil

`func (o *CourseAssignmentCreateDto) SetCourseUnitIdNil(b bool)`

 SetCourseUnitIdNil sets the value for CourseUnitId to be an explicit nil

### UnsetCourseUnitId
`func (o *CourseAssignmentCreateDto) UnsetCourseUnitId()`

UnsetCourseUnitId ensures that no value is present for CourseUnitId, not even an explicit nil
### GetCourseCohortId

`func (o *CourseAssignmentCreateDto) GetCourseCohortId() string`

GetCourseCohortId returns the CourseCohortId field if non-nil, zero value otherwise.

### GetCourseCohortIdOk

`func (o *CourseAssignmentCreateDto) GetCourseCohortIdOk() (*string, bool)`

GetCourseCohortIdOk returns a tuple with the CourseCohortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortId

`func (o *CourseAssignmentCreateDto) SetCourseCohortId(v string)`

SetCourseCohortId sets CourseCohortId field to given value.

### HasCourseCohortId

`func (o *CourseAssignmentCreateDto) HasCourseCohortId() bool`

HasCourseCohortId returns a boolean if a field has been set.

### SetCourseCohortIdNil

`func (o *CourseAssignmentCreateDto) SetCourseCohortIdNil(b bool)`

 SetCourseCohortIdNil sets the value for CourseCohortId to be an explicit nil

### UnsetCourseCohortId
`func (o *CourseAssignmentCreateDto) UnsetCourseCohortId()`

UnsetCourseCohortId ensures that no value is present for CourseCohortId, not even an explicit nil
### GetCourseAssignmentTypeId

`func (o *CourseAssignmentCreateDto) GetCourseAssignmentTypeId() string`

GetCourseAssignmentTypeId returns the CourseAssignmentTypeId field if non-nil, zero value otherwise.

### GetCourseAssignmentTypeIdOk

`func (o *CourseAssignmentCreateDto) GetCourseAssignmentTypeIdOk() (*string, bool)`

GetCourseAssignmentTypeIdOk returns a tuple with the CourseAssignmentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentTypeId

`func (o *CourseAssignmentCreateDto) SetCourseAssignmentTypeId(v string)`

SetCourseAssignmentTypeId sets CourseAssignmentTypeId field to given value.

### HasCourseAssignmentTypeId

`func (o *CourseAssignmentCreateDto) HasCourseAssignmentTypeId() bool`

HasCourseAssignmentTypeId returns a boolean if a field has been set.

### SetCourseAssignmentTypeIdNil

`func (o *CourseAssignmentCreateDto) SetCourseAssignmentTypeIdNil(b bool)`

 SetCourseAssignmentTypeIdNil sets the value for CourseAssignmentTypeId to be an explicit nil

### UnsetCourseAssignmentTypeId
`func (o *CourseAssignmentCreateDto) UnsetCourseAssignmentTypeId()`

UnsetCourseAssignmentTypeId ensures that no value is present for CourseAssignmentTypeId, not even an explicit nil
### GetDueDateTime

`func (o *CourseAssignmentCreateDto) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *CourseAssignmentCreateDto) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *CourseAssignmentCreateDto) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *CourseAssignmentCreateDto) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *CourseAssignmentCreateDto) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *CourseAssignmentCreateDto) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetAsignToAllCohorts

`func (o *CourseAssignmentCreateDto) GetAsignToAllCohorts() bool`

GetAsignToAllCohorts returns the AsignToAllCohorts field if non-nil, zero value otherwise.

### GetAsignToAllCohortsOk

`func (o *CourseAssignmentCreateDto) GetAsignToAllCohortsOk() (*bool, bool)`

GetAsignToAllCohortsOk returns a tuple with the AsignToAllCohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsignToAllCohorts

`func (o *CourseAssignmentCreateDto) SetAsignToAllCohorts(v bool)`

SetAsignToAllCohorts sets AsignToAllCohorts field to given value.

### HasAsignToAllCohorts

`func (o *CourseAssignmentCreateDto) HasAsignToAllCohorts() bool`

HasAsignToAllCohorts returns a boolean if a field has been set.

### GetResources

`func (o *CourseAssignmentCreateDto) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CourseAssignmentCreateDto) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CourseAssignmentCreateDto) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CourseAssignmentCreateDto) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *CourseAssignmentCreateDto) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *CourseAssignmentCreateDto) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


