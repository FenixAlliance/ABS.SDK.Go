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
**CourseID** | Pointer to **NullableString** |  | [optional] 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 
**CourseSectionID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

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

### GetCourseID

`func (o *CourseAssignmentDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseAssignmentDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseAssignmentDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseAssignmentDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseAssignmentDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseAssignmentDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetCourseUnitID

`func (o *CourseAssignmentDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseAssignmentDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseAssignmentDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseAssignmentDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseAssignmentDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseAssignmentDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil
### GetCourseSectionID

`func (o *CourseAssignmentDto) GetCourseSectionID() string`

GetCourseSectionID returns the CourseSectionID field if non-nil, zero value otherwise.

### GetCourseSectionIDOk

`func (o *CourseAssignmentDto) GetCourseSectionIDOk() (*string, bool)`

GetCourseSectionIDOk returns a tuple with the CourseSectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseSectionID

`func (o *CourseAssignmentDto) SetCourseSectionID(v string)`

SetCourseSectionID sets CourseSectionID field to given value.

### HasCourseSectionID

`func (o *CourseAssignmentDto) HasCourseSectionID() bool`

HasCourseSectionID returns a boolean if a field has been set.

### SetCourseSectionIDNil

`func (o *CourseAssignmentDto) SetCourseSectionIDNil(b bool)`

 SetCourseSectionIDNil sets the value for CourseSectionID to be an explicit nil

### UnsetCourseSectionID
`func (o *CourseAssignmentDto) UnsetCourseSectionID()`

UnsetCourseSectionID ensures that no value is present for CourseSectionID, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


