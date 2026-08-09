# RequiredSkillRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ExperienceInYears** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **float64** |  | [optional] 
**RequiredSkillRecordType** | Pointer to **string** |  | [optional] 
**SkillId** | Pointer to **NullableString** |  | [optional] 
**JobOfferId** | Pointer to **NullableString** |  | [optional] 
**EmployerProfileId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequiredSkillRecordDto

`func NewRequiredSkillRecordDto() *RequiredSkillRecordDto`

NewRequiredSkillRecordDto instantiates a new RequiredSkillRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredSkillRecordDtoWithDefaults

`func NewRequiredSkillRecordDtoWithDefaults() *RequiredSkillRecordDto`

NewRequiredSkillRecordDtoWithDefaults instantiates a new RequiredSkillRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequiredSkillRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequiredSkillRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequiredSkillRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequiredSkillRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RequiredSkillRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RequiredSkillRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *RequiredSkillRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequiredSkillRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequiredSkillRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RequiredSkillRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *RequiredSkillRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *RequiredSkillRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetExperienceInYears

`func (o *RequiredSkillRecordDto) GetExperienceInYears() int32`

GetExperienceInYears returns the ExperienceInYears field if non-nil, zero value otherwise.

### GetExperienceInYearsOk

`func (o *RequiredSkillRecordDto) GetExperienceInYearsOk() (*int32, bool)`

GetExperienceInYearsOk returns a tuple with the ExperienceInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInYears

`func (o *RequiredSkillRecordDto) SetExperienceInYears(v int32)`

SetExperienceInYears sets ExperienceInYears field to given value.

### HasExperienceInYears

`func (o *RequiredSkillRecordDto) HasExperienceInYears() bool`

HasExperienceInYears returns a boolean if a field has been set.

### GetPriority

`func (o *RequiredSkillRecordDto) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequiredSkillRecordDto) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequiredSkillRecordDto) SetPriority(v float64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequiredSkillRecordDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRequiredSkillRecordType

`func (o *RequiredSkillRecordDto) GetRequiredSkillRecordType() string`

GetRequiredSkillRecordType returns the RequiredSkillRecordType field if non-nil, zero value otherwise.

### GetRequiredSkillRecordTypeOk

`func (o *RequiredSkillRecordDto) GetRequiredSkillRecordTypeOk() (*string, bool)`

GetRequiredSkillRecordTypeOk returns a tuple with the RequiredSkillRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkillRecordType

`func (o *RequiredSkillRecordDto) SetRequiredSkillRecordType(v string)`

SetRequiredSkillRecordType sets RequiredSkillRecordType field to given value.

### HasRequiredSkillRecordType

`func (o *RequiredSkillRecordDto) HasRequiredSkillRecordType() bool`

HasRequiredSkillRecordType returns a boolean if a field has been set.

### GetSkillId

`func (o *RequiredSkillRecordDto) GetSkillId() string`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *RequiredSkillRecordDto) GetSkillIdOk() (*string, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *RequiredSkillRecordDto) SetSkillId(v string)`

SetSkillId sets SkillId field to given value.

### HasSkillId

`func (o *RequiredSkillRecordDto) HasSkillId() bool`

HasSkillId returns a boolean if a field has been set.

### SetSkillIdNil

`func (o *RequiredSkillRecordDto) SetSkillIdNil(b bool)`

 SetSkillIdNil sets the value for SkillId to be an explicit nil

### UnsetSkillId
`func (o *RequiredSkillRecordDto) UnsetSkillId()`

UnsetSkillId ensures that no value is present for SkillId, not even an explicit nil
### GetJobOfferId

`func (o *RequiredSkillRecordDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *RequiredSkillRecordDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *RequiredSkillRecordDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *RequiredSkillRecordDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *RequiredSkillRecordDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *RequiredSkillRecordDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil
### GetEmployerProfileId

`func (o *RequiredSkillRecordDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *RequiredSkillRecordDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *RequiredSkillRecordDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *RequiredSkillRecordDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *RequiredSkillRecordDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *RequiredSkillRecordDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil
### GetTenantId

`func (o *RequiredSkillRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RequiredSkillRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RequiredSkillRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RequiredSkillRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RequiredSkillRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RequiredSkillRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *RequiredSkillRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *RequiredSkillRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *RequiredSkillRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *RequiredSkillRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *RequiredSkillRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *RequiredSkillRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


