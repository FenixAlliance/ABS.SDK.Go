# CourseGradingRubricCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnablePoints** | Pointer to **bool** |  | [optional] 
**CourseId** | **string** |  | 

## Methods

### NewCourseGradingRubricCreateDto

`func NewCourseGradingRubricCreateDto(title string, courseId string, ) *CourseGradingRubricCreateDto`

NewCourseGradingRubricCreateDto instantiates a new CourseGradingRubricCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseGradingRubricCreateDtoWithDefaults

`func NewCourseGradingRubricCreateDtoWithDefaults() *CourseGradingRubricCreateDto`

NewCourseGradingRubricCreateDtoWithDefaults instantiates a new CourseGradingRubricCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseGradingRubricCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseGradingRubricCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseGradingRubricCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseGradingRubricCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseGradingRubricCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseGradingRubricCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseGradingRubricCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseGradingRubricCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseGradingRubricCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseGradingRubricCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseGradingRubricCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseGradingRubricCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseGradingRubricCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseGradingRubricCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseGradingRubricCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseGradingRubricCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseGradingRubricCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnablePoints

`func (o *CourseGradingRubricCreateDto) GetEnablePoints() bool`

GetEnablePoints returns the EnablePoints field if non-nil, zero value otherwise.

### GetEnablePointsOk

`func (o *CourseGradingRubricCreateDto) GetEnablePointsOk() (*bool, bool)`

GetEnablePointsOk returns a tuple with the EnablePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePoints

`func (o *CourseGradingRubricCreateDto) SetEnablePoints(v bool)`

SetEnablePoints sets EnablePoints field to given value.

### HasEnablePoints

`func (o *CourseGradingRubricCreateDto) HasEnablePoints() bool`

HasEnablePoints returns a boolean if a field has been set.

### GetCourseId

`func (o *CourseGradingRubricCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseGradingRubricCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseGradingRubricCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


