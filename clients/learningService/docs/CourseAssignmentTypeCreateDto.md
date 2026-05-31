# CourseAssignmentTypeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Abbreviation** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **float64** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Excluded** | Pointer to **int32** |  | [optional] 
**CourseID** | **string** |  | 

## Methods

### NewCourseAssignmentTypeCreateDto

`func NewCourseAssignmentTypeCreateDto(name string, courseID string, ) *CourseAssignmentTypeCreateDto`

NewCourseAssignmentTypeCreateDto instantiates a new CourseAssignmentTypeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseAssignmentTypeCreateDtoWithDefaults

`func NewCourseAssignmentTypeCreateDtoWithDefaults() *CourseAssignmentTypeCreateDto`

NewCourseAssignmentTypeCreateDtoWithDefaults instantiates a new CourseAssignmentTypeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseAssignmentTypeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseAssignmentTypeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseAssignmentTypeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseAssignmentTypeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseAssignmentTypeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseAssignmentTypeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseAssignmentTypeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseAssignmentTypeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CourseAssignmentTypeCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseAssignmentTypeCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseAssignmentTypeCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetAbbreviation

`func (o *CourseAssignmentTypeCreateDto) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *CourseAssignmentTypeCreateDto) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *CourseAssignmentTypeCreateDto) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *CourseAssignmentTypeCreateDto) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *CourseAssignmentTypeCreateDto) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *CourseAssignmentTypeCreateDto) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetWeight

`func (o *CourseAssignmentTypeCreateDto) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CourseAssignmentTypeCreateDto) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CourseAssignmentTypeCreateDto) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CourseAssignmentTypeCreateDto) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetQuantity

`func (o *CourseAssignmentTypeCreateDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CourseAssignmentTypeCreateDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CourseAssignmentTypeCreateDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CourseAssignmentTypeCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExcluded

`func (o *CourseAssignmentTypeCreateDto) GetExcluded() int32`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *CourseAssignmentTypeCreateDto) GetExcludedOk() (*int32, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *CourseAssignmentTypeCreateDto) SetExcluded(v int32)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *CourseAssignmentTypeCreateDto) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseAssignmentTypeCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseAssignmentTypeCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseAssignmentTypeCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


