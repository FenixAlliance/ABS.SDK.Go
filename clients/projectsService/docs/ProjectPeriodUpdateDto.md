# ProjectPeriodUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStartDate** | Pointer to **time.Time** |  | [optional] 
**PeriodEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectPeriodUpdateDto

`func NewProjectPeriodUpdateDto() *ProjectPeriodUpdateDto`

NewProjectPeriodUpdateDto instantiates a new ProjectPeriodUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPeriodUpdateDtoWithDefaults

`func NewProjectPeriodUpdateDtoWithDefaults() *ProjectPeriodUpdateDto`

NewProjectPeriodUpdateDtoWithDefaults instantiates a new ProjectPeriodUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStartDate

`func (o *ProjectPeriodUpdateDto) GetPeriodStartDate() time.Time`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *ProjectPeriodUpdateDto) GetPeriodStartDateOk() (*time.Time, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *ProjectPeriodUpdateDto) SetPeriodStartDate(v time.Time)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *ProjectPeriodUpdateDto) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *ProjectPeriodUpdateDto) GetPeriodEndDate() time.Time`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *ProjectPeriodUpdateDto) GetPeriodEndDateOk() (*time.Time, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *ProjectPeriodUpdateDto) SetPeriodEndDate(v time.Time)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *ProjectPeriodUpdateDto) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


