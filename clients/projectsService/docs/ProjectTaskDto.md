# ProjectTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**DueLine** | Pointer to **time.Time** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 
**ProjectTaskBucketID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectTaskDto

`func NewProjectTaskDto() *ProjectTaskDto`

NewProjectTaskDto instantiates a new ProjectTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTaskDtoWithDefaults

`func NewProjectTaskDtoWithDefaults() *ProjectTaskDto`

NewProjectTaskDtoWithDefaults instantiates a new ProjectTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTaskDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTaskDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectTaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectTaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProjectTaskDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectTaskDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectTaskDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectTaskDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProjectTaskDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProjectTaskDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetStartDate

`func (o *ProjectTaskDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectTaskDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectTaskDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectTaskDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueLine

`func (o *ProjectTaskDto) GetDueLine() time.Time`

GetDueLine returns the DueLine field if non-nil, zero value otherwise.

### GetDueLineOk

`func (o *ProjectTaskDto) GetDueLineOk() (*time.Time, bool)`

GetDueLineOk returns a tuple with the DueLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueLine

`func (o *ProjectTaskDto) SetDueLine(v time.Time)`

SetDueLine sets DueLine field to given value.

### HasDueLine

`func (o *ProjectTaskDto) HasDueLine() bool`

HasDueLine returns a boolean if a field has been set.

### GetProjectID

`func (o *ProjectTaskDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *ProjectTaskDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *ProjectTaskDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *ProjectTaskDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *ProjectTaskDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *ProjectTaskDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil
### GetProjectTaskBucketID

`func (o *ProjectTaskDto) GetProjectTaskBucketID() string`

GetProjectTaskBucketID returns the ProjectTaskBucketID field if non-nil, zero value otherwise.

### GetProjectTaskBucketIDOk

`func (o *ProjectTaskDto) GetProjectTaskBucketIDOk() (*string, bool)`

GetProjectTaskBucketIDOk returns a tuple with the ProjectTaskBucketID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskBucketID

`func (o *ProjectTaskDto) SetProjectTaskBucketID(v string)`

SetProjectTaskBucketID sets ProjectTaskBucketID field to given value.

### HasProjectTaskBucketID

`func (o *ProjectTaskDto) HasProjectTaskBucketID() bool`

HasProjectTaskBucketID returns a boolean if a field has been set.

### SetProjectTaskBucketIDNil

`func (o *ProjectTaskDto) SetProjectTaskBucketIDNil(b bool)`

 SetProjectTaskBucketIDNil sets the value for ProjectTaskBucketID to be an explicit nil

### UnsetProjectTaskBucketID
`func (o *ProjectTaskDto) UnsetProjectTaskBucketID()`

UnsetProjectTaskBucketID ensures that no value is present for ProjectTaskBucketID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


