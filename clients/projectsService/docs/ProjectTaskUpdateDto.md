# ProjectTaskUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**DueLine** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectTaskUpdateDto

`func NewProjectTaskUpdateDto() *ProjectTaskUpdateDto`

NewProjectTaskUpdateDto instantiates a new ProjectTaskUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTaskUpdateDtoWithDefaults

`func NewProjectTaskUpdateDtoWithDefaults() *ProjectTaskUpdateDto`

NewProjectTaskUpdateDtoWithDefaults instantiates a new ProjectTaskUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ProjectTaskUpdateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectTaskUpdateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectTaskUpdateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectTaskUpdateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueLine

`func (o *ProjectTaskUpdateDto) GetDueLine() time.Time`

GetDueLine returns the DueLine field if non-nil, zero value otherwise.

### GetDueLineOk

`func (o *ProjectTaskUpdateDto) GetDueLineOk() (*time.Time, bool)`

GetDueLineOk returns a tuple with the DueLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueLine

`func (o *ProjectTaskUpdateDto) SetDueLine(v time.Time)`

SetDueLine sets DueLine field to given value.

### HasDueLine

`func (o *ProjectTaskUpdateDto) HasDueLine() bool`

HasDueLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


