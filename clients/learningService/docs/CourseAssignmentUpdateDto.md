# CourseAssignmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Points** | Pointer to **float64** |  | [optional] 
**CourseUnitId** | Pointer to **NullableString** |  | [optional] 
**CourseCohortId** | Pointer to **NullableString** |  | [optional] 
**CourseAssignmentTypeId** | Pointer to **NullableString** |  | [optional] 
**DueDateTime** | Pointer to **NullableTime** |  | [optional] 
**AsignToAllCohorts** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseAssignmentUpdateDto

`func NewCourseAssignmentUpdateDto() *CourseAssignmentUpdateDto`

NewCourseAssignmentUpdateDto instantiates a new CourseAssignmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentUpdateDtoWithDefaults

`func NewCourseAssignmentUpdateDtoWithDefaults() *CourseAssignmentUpdateDto`

NewCourseAssignmentUpdateDtoWithDefaults instantiates a new CourseAssignmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CourseAssignmentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseAssignmentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseAssignmentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseAssignmentUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseAssignmentUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseAssignmentUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseAssignmentUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseAssignmentUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseAssignmentUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseAssignmentUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseAssignmentUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseAssignmentUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *CourseAssignmentUpdateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CourseAssignmentUpdateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CourseAssignmentUpdateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CourseAssignmentUpdateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CourseAssignmentUpdateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CourseAssignmentUpdateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetPoints

`func (o *CourseAssignmentUpdateDto) GetPoints() float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CourseAssignmentUpdateDto) GetPointsOk() (*float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CourseAssignmentUpdateDto) SetPoints(v float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CourseAssignmentUpdateDto) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetCourseUnitId

`func (o *CourseAssignmentUpdateDto) GetCourseUnitId() string`

GetCourseUnitId returns the CourseUnitId field if non-nil, zero value otherwise.

### GetCourseUnitIdOk

`func (o *CourseAssignmentUpdateDto) GetCourseUnitIdOk() (*string, bool)`

GetCourseUnitIdOk returns a tuple with the CourseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitId

`func (o *CourseAssignmentUpdateDto) SetCourseUnitId(v string)`

SetCourseUnitId sets CourseUnitId field to given value.

### HasCourseUnitId

`func (o *CourseAssignmentUpdateDto) HasCourseUnitId() bool`

HasCourseUnitId returns a boolean if a field has been set.

### SetCourseUnitIdNil

`func (o *CourseAssignmentUpdateDto) SetCourseUnitIdNil(b bool)`

 SetCourseUnitIdNil sets the value for CourseUnitId to be an explicit nil

### UnsetCourseUnitId
`func (o *CourseAssignmentUpdateDto) UnsetCourseUnitId()`

UnsetCourseUnitId ensures that no value is present for CourseUnitId, not even an explicit nil
### GetCourseCohortId

`func (o *CourseAssignmentUpdateDto) GetCourseCohortId() string`

GetCourseCohortId returns the CourseCohortId field if non-nil, zero value otherwise.

### GetCourseCohortIdOk

`func (o *CourseAssignmentUpdateDto) GetCourseCohortIdOk() (*string, bool)`

GetCourseCohortIdOk returns a tuple with the CourseCohortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortId

`func (o *CourseAssignmentUpdateDto) SetCourseCohortId(v string)`

SetCourseCohortId sets CourseCohortId field to given value.

### HasCourseCohortId

`func (o *CourseAssignmentUpdateDto) HasCourseCohortId() bool`

HasCourseCohortId returns a boolean if a field has been set.

### SetCourseCohortIdNil

`func (o *CourseAssignmentUpdateDto) SetCourseCohortIdNil(b bool)`

 SetCourseCohortIdNil sets the value for CourseCohortId to be an explicit nil

### UnsetCourseCohortId
`func (o *CourseAssignmentUpdateDto) UnsetCourseCohortId()`

UnsetCourseCohortId ensures that no value is present for CourseCohortId, not even an explicit nil
### GetCourseAssignmentTypeId

`func (o *CourseAssignmentUpdateDto) GetCourseAssignmentTypeId() string`

GetCourseAssignmentTypeId returns the CourseAssignmentTypeId field if non-nil, zero value otherwise.

### GetCourseAssignmentTypeIdOk

`func (o *CourseAssignmentUpdateDto) GetCourseAssignmentTypeIdOk() (*string, bool)`

GetCourseAssignmentTypeIdOk returns a tuple with the CourseAssignmentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentTypeId

`func (o *CourseAssignmentUpdateDto) SetCourseAssignmentTypeId(v string)`

SetCourseAssignmentTypeId sets CourseAssignmentTypeId field to given value.

### HasCourseAssignmentTypeId

`func (o *CourseAssignmentUpdateDto) HasCourseAssignmentTypeId() bool`

HasCourseAssignmentTypeId returns a boolean if a field has been set.

### SetCourseAssignmentTypeIdNil

`func (o *CourseAssignmentUpdateDto) SetCourseAssignmentTypeIdNil(b bool)`

 SetCourseAssignmentTypeIdNil sets the value for CourseAssignmentTypeId to be an explicit nil

### UnsetCourseAssignmentTypeId
`func (o *CourseAssignmentUpdateDto) UnsetCourseAssignmentTypeId()`

UnsetCourseAssignmentTypeId ensures that no value is present for CourseAssignmentTypeId, not even an explicit nil
### GetDueDateTime

`func (o *CourseAssignmentUpdateDto) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *CourseAssignmentUpdateDto) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *CourseAssignmentUpdateDto) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *CourseAssignmentUpdateDto) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *CourseAssignmentUpdateDto) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *CourseAssignmentUpdateDto) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetAsignToAllCohorts

`func (o *CourseAssignmentUpdateDto) GetAsignToAllCohorts() bool`

GetAsignToAllCohorts returns the AsignToAllCohorts field if non-nil, zero value otherwise.

### GetAsignToAllCohortsOk

`func (o *CourseAssignmentUpdateDto) GetAsignToAllCohortsOk() (*bool, bool)`

GetAsignToAllCohortsOk returns a tuple with the AsignToAllCohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsignToAllCohorts

`func (o *CourseAssignmentUpdateDto) SetAsignToAllCohorts(v bool)`

SetAsignToAllCohorts sets AsignToAllCohorts field to given value.

### HasAsignToAllCohorts

`func (o *CourseAssignmentUpdateDto) HasAsignToAllCohorts() bool`

HasAsignToAllCohorts returns a boolean if a field has been set.

### GetResources

`func (o *CourseAssignmentUpdateDto) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CourseAssignmentUpdateDto) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CourseAssignmentUpdateDto) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CourseAssignmentUpdateDto) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *CourseAssignmentUpdateDto) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *CourseAssignmentUpdateDto) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


