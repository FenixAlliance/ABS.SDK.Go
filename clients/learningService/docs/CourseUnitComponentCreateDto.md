# CourseUnitComponentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**CourseUnitID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseUnitComponentCreateDto

`func NewCourseUnitComponentCreateDto(title string, courseID string, businessID string, ) *CourseUnitComponentCreateDto`

NewCourseUnitComponentCreateDto instantiates a new CourseUnitComponentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseUnitComponentCreateDtoWithDefaults

`func NewCourseUnitComponentCreateDtoWithDefaults() *CourseUnitComponentCreateDto`

NewCourseUnitComponentCreateDtoWithDefaults instantiates a new CourseUnitComponentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseUnitComponentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseUnitComponentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseUnitComponentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseUnitComponentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseUnitComponentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseUnitComponentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseUnitComponentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseUnitComponentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseUnitComponentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseUnitComponentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseUnitComponentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseUnitComponentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseUnitComponentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseUnitComponentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseUnitComponentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseUnitComponentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseUnitComponentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseUnitComponentCreateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseUnitComponentCreateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseUnitComponentCreateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseUnitComponentCreateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseUnitComponentCreateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseUnitComponentCreateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetOrder

`func (o *CourseUnitComponentCreateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CourseUnitComponentCreateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CourseUnitComponentCreateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CourseUnitComponentCreateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseUnitComponentCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseUnitComponentCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseUnitComponentCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseUnitComponentCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseUnitComponentCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseUnitComponentCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetCourseUnitID

`func (o *CourseUnitComponentCreateDto) GetCourseUnitID() string`

GetCourseUnitID returns the CourseUnitID field if non-nil, zero value otherwise.

### GetCourseUnitIDOk

`func (o *CourseUnitComponentCreateDto) GetCourseUnitIDOk() (*string, bool)`

GetCourseUnitIDOk returns a tuple with the CourseUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseUnitID

`func (o *CourseUnitComponentCreateDto) SetCourseUnitID(v string)`

SetCourseUnitID sets CourseUnitID field to given value.

### HasCourseUnitID

`func (o *CourseUnitComponentCreateDto) HasCourseUnitID() bool`

HasCourseUnitID returns a boolean if a field has been set.

### SetCourseUnitIDNil

`func (o *CourseUnitComponentCreateDto) SetCourseUnitIDNil(b bool)`

 SetCourseUnitIDNil sets the value for CourseUnitID to be an explicit nil

### UnsetCourseUnitID
`func (o *CourseUnitComponentCreateDto) UnsetCourseUnitID()`

UnsetCourseUnitID ensures that no value is present for CourseUnitID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


