# ProjectPeriodCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**PeriodStartDate** | Pointer to **time.Time** |  | [optional] 
**PeriodEndDate** | Pointer to **time.Time** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectPeriodCreateDto

`func NewProjectPeriodCreateDto() *ProjectPeriodCreateDto`

NewProjectPeriodCreateDto instantiates a new ProjectPeriodCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPeriodCreateDtoWithDefaults

`func NewProjectPeriodCreateDtoWithDefaults() *ProjectPeriodCreateDto`

NewProjectPeriodCreateDtoWithDefaults instantiates a new ProjectPeriodCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectPeriodCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectPeriodCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectPeriodCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectPeriodCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProjectPeriodCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectPeriodCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectPeriodCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectPeriodCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPeriodStartDate

`func (o *ProjectPeriodCreateDto) GetPeriodStartDate() time.Time`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *ProjectPeriodCreateDto) GetPeriodStartDateOk() (*time.Time, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *ProjectPeriodCreateDto) SetPeriodStartDate(v time.Time)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *ProjectPeriodCreateDto) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *ProjectPeriodCreateDto) GetPeriodEndDate() time.Time`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *ProjectPeriodCreateDto) GetPeriodEndDateOk() (*time.Time, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *ProjectPeriodCreateDto) SetPeriodEndDate(v time.Time)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *ProjectPeriodCreateDto) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.

### GetProjectID

`func (o *ProjectPeriodCreateDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *ProjectPeriodCreateDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *ProjectPeriodCreateDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *ProjectPeriodCreateDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *ProjectPeriodCreateDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *ProjectPeriodCreateDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


