# CourseUnitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateTime** | Pointer to **time.Time** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**CourseSectionID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CourseHandouts** | Pointer to [**[]CourseHandoutDto**](CourseHandoutDto.md) |  | [optional] 
**CourseAssignments** | Pointer to [**[]CourseAssignmentDto**](CourseAssignmentDto.md) |  | [optional] 
**CourseComponents** | Pointer to [**[]CourseUnitComponentDto**](CourseUnitComponentDto.md) |  | [optional] 

## Methods

### NewCourseUnitDto

`func NewCourseUnitDto() *CourseUnitDto`

NewCourseUnitDto instantiates a new CourseUnitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseUnitDtoWithDefaults

`func NewCourseUnitDtoWithDefaults() *CourseUnitDto`

NewCourseUnitDtoWithDefaults instantiates a new CourseUnitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseUnitDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseUnitDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseUnitDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseUnitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseUnitDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseUnitDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseUnitDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseUnitDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseUnitDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseUnitDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseUnitDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseUnitDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CourseUnitDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseUnitDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseUnitDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseUnitDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseUnitDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseUnitDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseUnitDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseUnitDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseUnitDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseUnitDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseUnitDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseUnitDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReleaseDateTime

`func (o *CourseUnitDto) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *CourseUnitDto) GetReleaseDateTimeOk() (*time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateTime

`func (o *CourseUnitDto) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime sets ReleaseDateTime field to given value.

### HasReleaseDateTime

`func (o *CourseUnitDto) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### GetContent

`func (o *CourseUnitDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseUnitDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseUnitDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseUnitDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseUnitDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseUnitDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCourseID

`func (o *CourseUnitDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseUnitDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseUnitDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseUnitDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseUnitDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseUnitDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetCourseSectionID

`func (o *CourseUnitDto) GetCourseSectionID() string`

GetCourseSectionID returns the CourseSectionID field if non-nil, zero value otherwise.

### GetCourseSectionIDOk

`func (o *CourseUnitDto) GetCourseSectionIDOk() (*string, bool)`

GetCourseSectionIDOk returns a tuple with the CourseSectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseSectionID

`func (o *CourseUnitDto) SetCourseSectionID(v string)`

SetCourseSectionID sets CourseSectionID field to given value.

### HasCourseSectionID

`func (o *CourseUnitDto) HasCourseSectionID() bool`

HasCourseSectionID returns a boolean if a field has been set.

### SetCourseSectionIDNil

`func (o *CourseUnitDto) SetCourseSectionIDNil(b bool)`

 SetCourseSectionIDNil sets the value for CourseSectionID to be an explicit nil

### UnsetCourseSectionID
`func (o *CourseUnitDto) UnsetCourseSectionID()`

UnsetCourseSectionID ensures that no value is present for CourseSectionID, not even an explicit nil
### GetTenantId

`func (o *CourseUnitDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseUnitDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseUnitDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseUnitDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseUnitDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseUnitDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCourseHandouts

`func (o *CourseUnitDto) GetCourseHandouts() []CourseHandoutDto`

GetCourseHandouts returns the CourseHandouts field if non-nil, zero value otherwise.

### GetCourseHandoutsOk

`func (o *CourseUnitDto) GetCourseHandoutsOk() (*[]CourseHandoutDto, bool)`

GetCourseHandoutsOk returns a tuple with the CourseHandouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseHandouts

`func (o *CourseUnitDto) SetCourseHandouts(v []CourseHandoutDto)`

SetCourseHandouts sets CourseHandouts field to given value.

### HasCourseHandouts

`func (o *CourseUnitDto) HasCourseHandouts() bool`

HasCourseHandouts returns a boolean if a field has been set.

### SetCourseHandoutsNil

`func (o *CourseUnitDto) SetCourseHandoutsNil(b bool)`

 SetCourseHandoutsNil sets the value for CourseHandouts to be an explicit nil

### UnsetCourseHandouts
`func (o *CourseUnitDto) UnsetCourseHandouts()`

UnsetCourseHandouts ensures that no value is present for CourseHandouts, not even an explicit nil
### GetCourseAssignments

`func (o *CourseUnitDto) GetCourseAssignments() []CourseAssignmentDto`

GetCourseAssignments returns the CourseAssignments field if non-nil, zero value otherwise.

### GetCourseAssignmentsOk

`func (o *CourseUnitDto) GetCourseAssignmentsOk() (*[]CourseAssignmentDto, bool)`

GetCourseAssignmentsOk returns a tuple with the CourseAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignments

`func (o *CourseUnitDto) SetCourseAssignments(v []CourseAssignmentDto)`

SetCourseAssignments sets CourseAssignments field to given value.

### HasCourseAssignments

`func (o *CourseUnitDto) HasCourseAssignments() bool`

HasCourseAssignments returns a boolean if a field has been set.

### SetCourseAssignmentsNil

`func (o *CourseUnitDto) SetCourseAssignmentsNil(b bool)`

 SetCourseAssignmentsNil sets the value for CourseAssignments to be an explicit nil

### UnsetCourseAssignments
`func (o *CourseUnitDto) UnsetCourseAssignments()`

UnsetCourseAssignments ensures that no value is present for CourseAssignments, not even an explicit nil
### GetCourseComponents

`func (o *CourseUnitDto) GetCourseComponents() []CourseUnitComponentDto`

GetCourseComponents returns the CourseComponents field if non-nil, zero value otherwise.

### GetCourseComponentsOk

`func (o *CourseUnitDto) GetCourseComponentsOk() (*[]CourseUnitComponentDto, bool)`

GetCourseComponentsOk returns a tuple with the CourseComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseComponents

`func (o *CourseUnitDto) SetCourseComponents(v []CourseUnitComponentDto)`

SetCourseComponents sets CourseComponents field to given value.

### HasCourseComponents

`func (o *CourseUnitDto) HasCourseComponents() bool`

HasCourseComponents returns a boolean if a field has been set.

### SetCourseComponentsNil

`func (o *CourseUnitDto) SetCourseComponentsNil(b bool)`

 SetCourseComponentsNil sets the value for CourseComponents to be an explicit nil

### UnsetCourseComponents
`func (o *CourseUnitDto) UnsetCourseComponents()`

UnsetCourseComponents ensures that no value is present for CourseComponents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


