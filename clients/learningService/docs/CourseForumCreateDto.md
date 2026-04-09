# CourseForumCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 

## Methods

### NewCourseForumCreateDto

`func NewCourseForumCreateDto(title string, courseID string, businessID string, ) *CourseForumCreateDto`

NewCourseForumCreateDto instantiates a new CourseForumCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseForumCreateDtoWithDefaults

`func NewCourseForumCreateDtoWithDefaults() *CourseForumCreateDto`

NewCourseForumCreateDtoWithDefaults instantiates a new CourseForumCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseForumCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseForumCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseForumCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseForumCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseForumCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseForumCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseForumCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseForumCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseForumCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseForumCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseForumCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseForumCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseForumCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseForumCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseForumCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseForumCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseForumCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCourseID

`func (o *CourseForumCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseForumCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseForumCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseForumCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseForumCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseForumCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


