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
**CourseID** | **string** |  | 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**CourseCohortID** | Pointer to **NullableString** |  | [optional] 
**CourseAssignmentTypeID** | Pointer to **NullableString** |  | [optional] 
**DueDateTime** | Pointer to **NullableTime** |  | [optional] 
**AsignToAllCohorts** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseAssignmentCreateDto

`func NewCourseAssignmentCreateDto(title string, courseID string, ) *CourseAssignmentCreateDto`

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

### GetCourseID

`func (o *CourseAssignmentCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseAssignmentCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseAssignmentCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetCourseUnitID

`func (o *CourseAssignmentCreateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseAssignmentCreateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseAssignmentCreateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseAssignmentCreateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseAssignmentCreateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseAssignmentCreateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetCourseCohortID

`func (o *CourseAssignmentCreateDto) GetCourseCohortID() string`

GetCourseCohortID returns the CourseCohortID field if non-nil, zero value otherwise.

### GetCourseCohortIDOk

`func (o *CourseAssignmentCreateDto) GetCourseCohortIDOk() (*string, bool)`

GetCourseCohortIDOk returns a tuple with the CourseCohortID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCohortID

`func (o *CourseAssignmentCreateDto) SetCourseCohortID(v string)`

SetCourseCohortID sets CourseCohortID field to given value.

### HasCourseCohortID

`func (o *CourseAssignmentCreateDto) HasCourseCohortID() bool`

HasCourseCohortID returns a boolean if a field has been set.

### SetCourseCohortIDNil

`func (o *CourseAssignmentCreateDto) SetCourseCohortIDNil(b bool)`

 SetCourseCohortIDNil sets the value for CourseCohortID to be an explicit nil

### UnsetCourseCohortID
`func (o *CourseAssignmentCreateDto) UnsetCourseCohortID()`

UnsetCourseCohortID ensures that no value is present for CourseCohortID, not even an explicit nil
### GetCourseAssignmentTypeID

`func (o *CourseAssignmentCreateDto) GetCourseAssignmentTypeID() string`

GetCourseAssignmentTypeID returns the CourseAssignmentTypeID field if non-nil, zero value otherwise.

### GetCourseAssignmentTypeIDOk

`func (o *CourseAssignmentCreateDto) GetCourseAssignmentTypeIDOk() (*string, bool)`

GetCourseAssignmentTypeIDOk returns a tuple with the CourseAssignmentTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentTypeID

`func (o *CourseAssignmentCreateDto) SetCourseAssignmentTypeID(v string)`

SetCourseAssignmentTypeID sets CourseAssignmentTypeID field to given value.

### HasCourseAssignmentTypeID

`func (o *CourseAssignmentCreateDto) HasCourseAssignmentTypeID() bool`

HasCourseAssignmentTypeID returns a boolean if a field has been set.

### SetCourseAssignmentTypeIDNil

`func (o *CourseAssignmentCreateDto) SetCourseAssignmentTypeIDNil(b bool)`

 SetCourseAssignmentTypeIDNil sets the value for CourseAssignmentTypeID to be an explicit nil

### UnsetCourseAssignmentTypeID
`func (o *CourseAssignmentCreateDto) UnsetCourseAssignmentTypeID()`

UnsetCourseAssignmentTypeID ensures that no value is present for CourseAssignmentTypeID, not even an explicit nil
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


