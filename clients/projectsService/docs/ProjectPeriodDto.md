# ProjectPeriodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**PeriodStartDate** | Pointer to **time.Time** |  | [optional] 
**PeriodEndDate** | Pointer to **time.Time** |  | [optional] 
**ProjectID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectPeriodDto

`func NewProjectPeriodDto() *ProjectPeriodDto`

NewProjectPeriodDto instantiates a new ProjectPeriodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPeriodDtoWithDefaults

`func NewProjectPeriodDtoWithDefaults() *ProjectPeriodDto`

NewProjectPeriodDtoWithDefaults instantiates a new ProjectPeriodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectPeriodDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectPeriodDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectPeriodDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectPeriodDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProjectPeriodDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProjectPeriodDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProjectPeriodDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectPeriodDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectPeriodDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectPeriodDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProjectPeriodDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProjectPeriodDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPeriodStartDate

`func (o *ProjectPeriodDto) GetPeriodStartDate() time.Time`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *ProjectPeriodDto) GetPeriodStartDateOk() (*time.Time, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *ProjectPeriodDto) SetPeriodStartDate(v time.Time)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *ProjectPeriodDto) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *ProjectPeriodDto) GetPeriodEndDate() time.Time`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *ProjectPeriodDto) GetPeriodEndDateOk() (*time.Time, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *ProjectPeriodDto) SetPeriodEndDate(v time.Time)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *ProjectPeriodDto) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.

### GetProjectID

`func (o *ProjectPeriodDto) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *ProjectPeriodDto) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *ProjectPeriodDto) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *ProjectPeriodDto) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### SetProjectIDNil

`func (o *ProjectPeriodDto) SetProjectIDNil(b bool)`

 SetProjectIDNil sets the value for ProjectID to be an explicit nil

### UnsetProjectID
`func (o *ProjectPeriodDto) UnsetProjectID()`

UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


