# CourseAssignmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Points** | Pointer to **float64** |  | [optional] 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**CourseCohortID** | Pointer to **NullableString** |  | [optional] 
**CourseAssignmentTypeID** | Pointer to **NullableString** |  | [optional] 
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

### GetCourseUnitID

`func (o *CourseAssignmentUpdateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseAssignmentUpdateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseAssignmentUpdateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseAssignmentUpdateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseAssignmentUpdateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseAssignmentUpdateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetCourseCohortID

`func (o *CourseAssignmentUpdateDto) GetCourseCohortID() string`

GetCourseCohortID returns the CourseCohortID field if non-nil, zero value otherwise.

### GetCourseCohortIDOk

`func (o *CourseAssignmentUpdateDto) GetCourseCohortIDOk() (*string, bool)`

GetCourseCohortIDOk returns a tuple with the CourseCohortID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortID

`func (o *CourseAssignmentUpdateDto) SetCourseCohortID(v string)`

SetCourseCohortID sets CourseCohortID field to given value.

### HasCourseCohortID

`func (o *CourseAssignmentUpdateDto) HasCourseCohortID() bool`

HasCourseCohortID returns a boolean if a field has been set.

### SetCourseCohortIDNil

`func (o *CourseAssignmentUpdateDto) SetCourseCohortIDNil(b bool)`

 SetCourseCohortIDNil sets the value for CourseCohortID to be an explicit nil

### UnsetCourseCohortID
`func (o *CourseAssignmentUpdateDto) UnsetCourseCohortID()`

UnsetCourseCohortID ensures that no value is present for CourseCohortID, not even an explicit nil
### GetCourseAssignmentTypeID

`func (o *CourseAssignmentUpdateDto) GetCourseAssignmentTypeID() string`

GetCourseAssignmentTypeID returns the CourseAssignmentTypeID field if non-nil, zero value otherwise.

### GetCourseAssignmentTypeIDOk

`func (o *CourseAssignmentUpdateDto) GetCourseAssignmentTypeIDOk() (*string, bool)`

GetCourseAssignmentTypeIDOk returns a tuple with the CourseAssignmentTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentTypeID

`func (o *CourseAssignmentUpdateDto) SetCourseAssignmentTypeID(v string)`

SetCourseAssignmentTypeID sets CourseAssignmentTypeID field to given value.

### HasCourseAssignmentTypeID

`func (o *CourseAssignmentUpdateDto) HasCourseAssignmentTypeID() bool`

HasCourseAssignmentTypeID returns a boolean if a field has been set.

### SetCourseAssignmentTypeIDNil

`func (o *CourseAssignmentUpdateDto) SetCourseAssignmentTypeIDNil(b bool)`

 SetCourseAssignmentTypeIDNil sets the value for CourseAssignmentTypeID to be an explicit nil

### UnsetCourseAssignmentTypeID
`func (o *CourseAssignmentUpdateDto) UnsetCourseAssignmentTypeID()`

UnsetCourseAssignmentTypeID ensures that no value is present for CourseAssignmentTypeID, not even an explicit nil
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


