# CourseAssignmentComponentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**CourseAssignmentId** | **string** |  | 
**CourseId** | **string** |  | 

## Methods

### NewCourseAssignmentComponentCreateDto

`func NewCourseAssignmentComponentCreateDto(title string, courseAssignmentId string, courseId string, ) *CourseAssignmentComponentCreateDto`

NewCourseAssignmentComponentCreateDto instantiates a new CourseAssignmentComponentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentComponentCreateDtoWithDefaults

`func NewCourseAssignmentComponentCreateDtoWithDefaults() *CourseAssignmentComponentCreateDto`

NewCourseAssignmentComponentCreateDtoWithDefaults instantiates a new CourseAssignmentComponentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseAssignmentComponentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseAssignmentComponentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseAssignmentComponentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseAssignmentComponentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseAssignmentComponentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseAssignmentComponentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseAssignmentComponentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseAssignmentComponentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseAssignmentComponentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseAssignmentComponentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseAssignmentComponentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseAssignmentComponentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseAssignmentComponentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseAssignmentComponentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseAssignmentComponentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseAssignmentComponentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseAssignmentComponentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseAssignmentComponentCreateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseAssignmentComponentCreateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseAssignmentComponentCreateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseAssignmentComponentCreateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseAssignmentComponentCreateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseAssignmentComponentCreateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetOrder

`func (o *CourseAssignmentComponentCreateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CourseAssignmentComponentCreateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CourseAssignmentComponentCreateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CourseAssignmentComponentCreateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCourseAssignmentId

`func (o *CourseAssignmentComponentCreateDto) GetCourseAssignmentId() string`

GetCourseAssignmentId returns the CourseAssignmentId field if non-nil, zero value otherwise.

### GetCourseAssignmentIdOk

`func (o *CourseAssignmentComponentCreateDto) GetCourseAssignmentIdOk() (*string, bool)`

GetCourseAssignmentIdOk returns a tuple with the CourseAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseAssignmentId

`func (o *CourseAssignmentComponentCreateDto) SetCourseAssignmentId(v string)`

SetCourseAssignmentId sets CourseAssignmentId field to given value.


### GetCourseId

`func (o *CourseAssignmentComponentCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseAssignmentComponentCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseAssignmentComponentCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


