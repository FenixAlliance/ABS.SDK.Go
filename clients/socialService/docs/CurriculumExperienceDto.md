# CurriculumExperienceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CurriculumId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **float64** |  | [optional] 
**Organization** | Pointer to **NullableString** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**Validated** | Pointer to **bool** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**DateFrom** | Pointer to **time.Time** |  | [optional] 
**DateTo** | Pointer to **time.Time** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**CountryLanguageId** | Pointer to **NullableString** |  | [optional] 
**ProficiencyRatingValueId** | Pointer to **NullableString** |  | [optional] 
**ProficiencyRatingModelId** | Pointer to **NullableString** |  | [optional] 
**Achievements** | Pointer to **NullableString** |  | [optional] 
**Responsibilities** | Pointer to **NullableString** |  | [optional] 
**EmployerProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCurriculumExperienceDto

`func NewCurriculumExperienceDto() *CurriculumExperienceDto`

NewCurriculumExperienceDto instantiates a new CurriculumExperienceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurriculumExperienceDtoWithDefaults

`func NewCurriculumExperienceDtoWithDefaults() *CurriculumExperienceDto`

NewCurriculumExperienceDtoWithDefaults instantiates a new CurriculumExperienceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurriculumExperienceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurriculumExperienceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurriculumExperienceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurriculumExperienceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CurriculumExperienceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CurriculumExperienceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CurriculumExperienceDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CurriculumExperienceDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CurriculumExperienceDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CurriculumExperienceDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CurriculumExperienceDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CurriculumExperienceDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCurriculumId

`func (o *CurriculumExperienceDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *CurriculumExperienceDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *CurriculumExperienceDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *CurriculumExperienceDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *CurriculumExperienceDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *CurriculumExperienceDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetTitle

`func (o *CurriculumExperienceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CurriculumExperienceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CurriculumExperienceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CurriculumExperienceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CurriculumExperienceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CurriculumExperienceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CurriculumExperienceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurriculumExperienceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurriculumExperienceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurriculumExperienceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurriculumExperienceDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurriculumExperienceDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *CurriculumExperienceDto) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CurriculumExperienceDto) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CurriculumExperienceDto) SetPriority(v float64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CurriculumExperienceDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrganization

`func (o *CurriculumExperienceDto) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CurriculumExperienceDto) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CurriculumExperienceDto) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CurriculumExperienceDto) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CurriculumExperienceDto) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CurriculumExperienceDto) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetFeatured

`func (o *CurriculumExperienceDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CurriculumExperienceDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CurriculumExperienceDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CurriculumExperienceDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetCurrent

`func (o *CurriculumExperienceDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CurriculumExperienceDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CurriculumExperienceDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CurriculumExperienceDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetValidated

`func (o *CurriculumExperienceDto) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *CurriculumExperienceDto) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *CurriculumExperienceDto) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *CurriculumExperienceDto) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetImageUrl

`func (o *CurriculumExperienceDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CurriculumExperienceDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CurriculumExperienceDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CurriculumExperienceDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CurriculumExperienceDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CurriculumExperienceDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDateFrom

`func (o *CurriculumExperienceDto) GetDateFrom() time.Time`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *CurriculumExperienceDto) GetDateFromOk() (*time.Time, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *CurriculumExperienceDto) SetDateFrom(v time.Time)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *CurriculumExperienceDto) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *CurriculumExperienceDto) GetDateTo() time.Time`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *CurriculumExperienceDto) GetDateToOk() (*time.Time, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *CurriculumExperienceDto) SetDateTo(v time.Time)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *CurriculumExperienceDto) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetDate

`func (o *CurriculumExperienceDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CurriculumExperienceDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CurriculumExperienceDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *CurriculumExperienceDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCountryLanguageId

`func (o *CurriculumExperienceDto) GetCountryLanguageId() string`

GetCountryLanguageId returns the CountryLanguageId field if non-nil, zero value otherwise.

### GetCountryLanguageIdOk

`func (o *CurriculumExperienceDto) GetCountryLanguageIdOk() (*string, bool)`

GetCountryLanguageIdOk returns a tuple with the CountryLanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLanguageId

`func (o *CurriculumExperienceDto) SetCountryLanguageId(v string)`

SetCountryLanguageId sets CountryLanguageId field to given value.

### HasCountryLanguageId

`func (o *CurriculumExperienceDto) HasCountryLanguageId() bool`

HasCountryLanguageId returns a boolean if a field has been set.

### SetCountryLanguageIdNil

`func (o *CurriculumExperienceDto) SetCountryLanguageIdNil(b bool)`

 SetCountryLanguageIdNil sets the value for CountryLanguageId to be an explicit nil

### UnsetCountryLanguageId
`func (o *CurriculumExperienceDto) UnsetCountryLanguageId()`

UnsetCountryLanguageId ensures that no value is present for CountryLanguageId, not even an explicit nil
### GetProficiencyRatingValueId

`func (o *CurriculumExperienceDto) GetProficiencyRatingValueId() string`

GetProficiencyRatingValueId returns the ProficiencyRatingValueId field if non-nil, zero value otherwise.

### GetProficiencyRatingValueIdOk

`func (o *CurriculumExperienceDto) GetProficiencyRatingValueIdOk() (*string, bool)`

GetProficiencyRatingValueIdOk returns a tuple with the ProficiencyRatingValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingValueId

`func (o *CurriculumExperienceDto) SetProficiencyRatingValueId(v string)`

SetProficiencyRatingValueId sets ProficiencyRatingValueId field to given value.

### HasProficiencyRatingValueId

`func (o *CurriculumExperienceDto) HasProficiencyRatingValueId() bool`

HasProficiencyRatingValueId returns a boolean if a field has been set.

### SetProficiencyRatingValueIdNil

`func (o *CurriculumExperienceDto) SetProficiencyRatingValueIdNil(b bool)`

 SetProficiencyRatingValueIdNil sets the value for ProficiencyRatingValueId to be an explicit nil

### UnsetProficiencyRatingValueId
`func (o *CurriculumExperienceDto) UnsetProficiencyRatingValueId()`

UnsetProficiencyRatingValueId ensures that no value is present for ProficiencyRatingValueId, not even an explicit nil
### GetProficiencyRatingModelId

`func (o *CurriculumExperienceDto) GetProficiencyRatingModelId() string`

GetProficiencyRatingModelId returns the ProficiencyRatingModelId field if non-nil, zero value otherwise.

### GetProficiencyRatingModelIdOk

`func (o *CurriculumExperienceDto) GetProficiencyRatingModelIdOk() (*string, bool)`

GetProficiencyRatingModelIdOk returns a tuple with the ProficiencyRatingModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingModelId

`func (o *CurriculumExperienceDto) SetProficiencyRatingModelId(v string)`

SetProficiencyRatingModelId sets ProficiencyRatingModelId field to given value.

### HasProficiencyRatingModelId

`func (o *CurriculumExperienceDto) HasProficiencyRatingModelId() bool`

HasProficiencyRatingModelId returns a boolean if a field has been set.

### SetProficiencyRatingModelIdNil

`func (o *CurriculumExperienceDto) SetProficiencyRatingModelIdNil(b bool)`

 SetProficiencyRatingModelIdNil sets the value for ProficiencyRatingModelId to be an explicit nil

### UnsetProficiencyRatingModelId
`func (o *CurriculumExperienceDto) UnsetProficiencyRatingModelId()`

UnsetProficiencyRatingModelId ensures that no value is present for ProficiencyRatingModelId, not even an explicit nil
### GetAchievements

`func (o *CurriculumExperienceDto) GetAchievements() string`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *CurriculumExperienceDto) GetAchievementsOk() (*string, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *CurriculumExperienceDto) SetAchievements(v string)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *CurriculumExperienceDto) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### SetAchievementsNil

`func (o *CurriculumExperienceDto) SetAchievementsNil(b bool)`

 SetAchievementsNil sets the value for Achievements to be an explicit nil

### UnsetAchievements
`func (o *CurriculumExperienceDto) UnsetAchievements()`

UnsetAchievements ensures that no value is present for Achievements, not even an explicit nil
### GetResponsibilities

`func (o *CurriculumExperienceDto) GetResponsibilities() string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *CurriculumExperienceDto) GetResponsibilitiesOk() (*string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibilities

`func (o *CurriculumExperienceDto) SetResponsibilities(v string)`

SetResponsibilities sets Responsibilities field to given value.

### HasResponsibilities

`func (o *CurriculumExperienceDto) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### SetResponsibilitiesNil

`func (o *CurriculumExperienceDto) SetResponsibilitiesNil(b bool)`

 SetResponsibilitiesNil sets the value for Responsibilities to be an explicit nil

### UnsetResponsibilities
`func (o *CurriculumExperienceDto) UnsetResponsibilities()`

UnsetResponsibilities ensures that no value is present for Responsibilities, not even an explicit nil
### GetEmployerProfileId

`func (o *CurriculumExperienceDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *CurriculumExperienceDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *CurriculumExperienceDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *CurriculumExperienceDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *CurriculumExperienceDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *CurriculumExperienceDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


