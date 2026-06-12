# CourseAssignmentComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**CourseAssignmentId** | Pointer to **NullableString** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseAssignmentComponentDto

`func NewCourseAssignmentComponentDto() *CourseAssignmentComponentDto`

NewCourseAssignmentComponentDto instantiates a new CourseAssignmentComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentComponentDtoWithDefaults

`func NewCourseAssignmentComponentDtoWithDefaults() *CourseAssignmentComponentDto`

NewCourseAssignmentComponentDtoWithDefaults instantiates a new CourseAssignmentComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseAssignmentComponentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseAssignmentComponentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseAssignmentComponentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseAssignmentComponentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseAssignmentComponentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseAssignmentComponentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseAssignmentComponentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseAssignmentComponentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseAssignmentComponentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseAssignmentComponentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseAssignmentComponentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseAssignmentComponentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CourseAssignmentComponentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseAssignmentComponentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseAssignmentComponentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseAssignmentComponentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseAssignmentComponentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseAssignmentComponentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseAssignmentComponentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseAssignmentComponentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseAssignmentComponentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseAssignmentComponentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseAssignmentComponentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseAssignmentComponentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseAssignmentComponentDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseAssignmentComponentDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseAssignmentComponentDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseAssignmentComponentDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseAssignmentComponentDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseAssignmentComponentDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetOrder

`func (o *CourseAssignmentComponentDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CourseAssignmentComponentDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CourseAssignmentComponentDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CourseAssignmentComponentDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCourseAssignmentId

`func (o *CourseAssignmentComponentDto) GetCourseAssignmentId() string`

GetCourseAssignmentId returns the CourseAssignmentId field if non-nil, zero value otherwise.

### GetCourseAssignmentIdOk

`func (o *CourseAssignmentComponentDto) GetCourseAssignmentIdOk() (*string, bool)`

GetCourseAssignmentIdOk returns a tuple with the CourseAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentId

`func (o *CourseAssignmentComponentDto) SetCourseAssignmentId(v string)`

SetCourseAssignmentId sets CourseAssignmentId field to given value.

### HasCourseAssignmentId

`func (o *CourseAssignmentComponentDto) HasCourseAssignmentId() bool`

HasCourseAssignmentId returns a boolean if a field has been set.

### SetCourseAssignmentIdNil

`func (o *CourseAssignmentComponentDto) SetCourseAssignmentIdNil(b bool)`

 SetCourseAssignmentIdNil sets the value for CourseAssignmentId to be an explicit nil

### UnsetCourseAssignmentId
`func (o *CourseAssignmentComponentDto) UnsetCourseAssignmentId()`

UnsetCourseAssignmentId ensures that no value is present for CourseAssignmentId, not even an explicit nil
### GetCourseId

`func (o *CourseAssignmentComponentDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseAssignmentComponentDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseAssignmentComponentDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseAssignmentComponentDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseAssignmentComponentDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseAssignmentComponentDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetTenantId

`func (o *CourseAssignmentComponentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseAssignmentComponentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseAssignmentComponentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseAssignmentComponentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseAssignmentComponentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseAssignmentComponentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseAssignmentComponentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseAssignmentComponentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseAssignmentComponentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseAssignmentComponentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseAssignmentComponentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseAssignmentComponentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


