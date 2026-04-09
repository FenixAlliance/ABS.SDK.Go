# GigCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Budget** | Pointer to **float64** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**SkillsRequired** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGigCreateDto

`func NewGigCreateDto() *GigCreateDto`

NewGigCreateDto instantiates a new GigCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGigCreateDtoWithDefaults

`func NewGigCreateDtoWithDefaults() *GigCreateDto`

NewGigCreateDtoWithDefaults instantiates a new GigCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GigCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GigCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GigCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GigCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GigCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GigCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GigCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GigCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *GigCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GigCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GigCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GigCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GigCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GigCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *GigCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GigCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GigCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GigCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GigCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GigCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDate

`func (o *GigCreateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GigCreateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GigCreateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GigCreateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GigCreateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GigCreateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GigCreateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GigCreateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetBudget

`func (o *GigCreateDto) GetBudget() float64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *GigCreateDto) GetBudgetOk() (*float64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *GigCreateDto) SetBudget(v float64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *GigCreateDto) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetLocation

`func (o *GigCreateDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GigCreateDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GigCreateDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GigCreateDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GigCreateDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GigCreateDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSkillsRequired

`func (o *GigCreateDto) GetSkillsRequired() string`

GetSkillsRequired returns the SkillsRequired field if non-nil, zero value otherwise.

### GetSkillsRequiredOk

`func (o *GigCreateDto) GetSkillsRequiredOk() (*string, bool)`

GetSkillsRequiredOk returns a tuple with the SkillsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillsRequired

`func (o *GigCreateDto) SetSkillsRequired(v string)`

SetSkillsRequired sets SkillsRequired field to given value.

### HasSkillsRequired

`func (o *GigCreateDto) HasSkillsRequired() bool`

HasSkillsRequired returns a boolean if a field has been set.

### SetSkillsRequiredNil

`func (o *GigCreateDto) SetSkillsRequiredNil(b bool)`

 SetSkillsRequiredNil sets the value for SkillsRequired to be an explicit nil

### UnsetSkillsRequired
`func (o *GigCreateDto) UnsetSkillsRequired()`

UnsetSkillsRequired ensures that no value is present for SkillsRequired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


