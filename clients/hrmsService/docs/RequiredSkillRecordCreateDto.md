# RequiredSkillRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**SkillId** | **string** |  | 
**JobOfferId** | Pointer to **NullableString** |  | [optional] 
**EmployerProfileId** | Pointer to **NullableString** |  | [optional] 
**ExperienceInYears** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **float64** |  | [optional] 
**RequiredSkillRecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewRequiredSkillRecordCreateDto

`func NewRequiredSkillRecordCreateDto(skillId string, ) *RequiredSkillRecordCreateDto`

NewRequiredSkillRecordCreateDto instantiates a new RequiredSkillRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredSkillRecordCreateDtoWithDefaults

`func NewRequiredSkillRecordCreateDtoWithDefaults() *RequiredSkillRecordCreateDto`

NewRequiredSkillRecordCreateDtoWithDefaults instantiates a new RequiredSkillRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequiredSkillRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequiredSkillRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequiredSkillRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequiredSkillRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *RequiredSkillRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequiredSkillRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequiredSkillRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RequiredSkillRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSkillId

`func (o *RequiredSkillRecordCreateDto) GetSkillId() string`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *RequiredSkillRecordCreateDto) GetSkillIdOk() (*string, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *RequiredSkillRecordCreateDto) SetSkillId(v string)`

SetSkillId sets SkillId field to given value.


### GetJobOfferId

`func (o *RequiredSkillRecordCreateDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *RequiredSkillRecordCreateDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *RequiredSkillRecordCreateDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *RequiredSkillRecordCreateDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *RequiredSkillRecordCreateDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *RequiredSkillRecordCreateDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil
### GetEmployerProfileId

`func (o *RequiredSkillRecordCreateDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *RequiredSkillRecordCreateDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *RequiredSkillRecordCreateDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *RequiredSkillRecordCreateDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *RequiredSkillRecordCreateDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *RequiredSkillRecordCreateDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil
### GetExperienceInYears

`func (o *RequiredSkillRecordCreateDto) GetExperienceInYears() int32`

GetExperienceInYears returns the ExperienceInYears field if non-nil, zero value otherwise.

### GetExperienceInYearsOk

`func (o *RequiredSkillRecordCreateDto) GetExperienceInYearsOk() (*int32, bool)`

GetExperienceInYearsOk returns a tuple with the ExperienceInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInYears

`func (o *RequiredSkillRecordCreateDto) SetExperienceInYears(v int32)`

SetExperienceInYears sets ExperienceInYears field to given value.

### HasExperienceInYears

`func (o *RequiredSkillRecordCreateDto) HasExperienceInYears() bool`

HasExperienceInYears returns a boolean if a field has been set.

### GetPriority

`func (o *RequiredSkillRecordCreateDto) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequiredSkillRecordCreateDto) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequiredSkillRecordCreateDto) SetPriority(v float64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequiredSkillRecordCreateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRequiredSkillRecordType

`func (o *RequiredSkillRecordCreateDto) GetRequiredSkillRecordType() string`

GetRequiredSkillRecordType returns the RequiredSkillRecordType field if non-nil, zero value otherwise.

### GetRequiredSkillRecordTypeOk

`func (o *RequiredSkillRecordCreateDto) GetRequiredSkillRecordTypeOk() (*string, bool)`

GetRequiredSkillRecordTypeOk returns a tuple with the RequiredSkillRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkillRecordType

`func (o *RequiredSkillRecordCreateDto) SetRequiredSkillRecordType(v string)`

SetRequiredSkillRecordType sets RequiredSkillRecordType field to given value.

### HasRequiredSkillRecordType

`func (o *RequiredSkillRecordCreateDto) HasRequiredSkillRecordType() bool`

HasRequiredSkillRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


