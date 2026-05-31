# TrainingProgramCourseCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TrainingProgramId** | **string** |  | 
**CourseId** | **string** |  | 

## Methods

### NewTrainingProgramCourseCreateDto

`func NewTrainingProgramCourseCreateDto(trainingProgramId string, courseId string, ) *TrainingProgramCourseCreateDto`

NewTrainingProgramCourseCreateDto instantiates a new TrainingProgramCourseCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingProgramCourseCreateDtoWithDefaults

`func NewTrainingProgramCourseCreateDtoWithDefaults() *TrainingProgramCourseCreateDto`

NewTrainingProgramCourseCreateDtoWithDefaults instantiates a new TrainingProgramCourseCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainingProgramCourseCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainingProgramCourseCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainingProgramCourseCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainingProgramCourseCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TrainingProgramCourseCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrainingProgramCourseCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrainingProgramCourseCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TrainingProgramCourseCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrainingProgramId

`func (o *TrainingProgramCourseCreateDto) GetTrainingProgramId() string`

GetTrainingProgramId returns the TrainingProgramId field if non-nil, zero value otherwise.

### GetTrainingProgramIdOk

`func (o *TrainingProgramCourseCreateDto) GetTrainingProgramIdOk() (*string, bool)`

GetTrainingProgramIdOk returns a tuple with the TrainingProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingProgramId

`func (o *TrainingProgramCourseCreateDto) SetTrainingProgramId(v string)`

SetTrainingProgramId sets TrainingProgramId field to given value.


### GetCourseId

`func (o *TrainingProgramCourseCreateDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *TrainingProgramCourseCreateDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *TrainingProgramCourseCreateDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


