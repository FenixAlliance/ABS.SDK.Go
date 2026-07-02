# CurriculumExperienceUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewCurriculumExperienceUpdateDto

`func NewCurriculumExperienceUpdateDto() *CurriculumExperienceUpdateDto`

NewCurriculumExperienceUpdateDto instantiates a new CurriculumExperienceUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurriculumExperienceUpdateDtoWithDefaults

`func NewCurriculumExperienceUpdateDtoWithDefaults() *CurriculumExperienceUpdateDto`

NewCurriculumExperienceUpdateDtoWithDefaults instantiates a new CurriculumExperienceUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CurriculumExperienceUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CurriculumExperienceUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CurriculumExperienceUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CurriculumExperienceUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CurriculumExperienceUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CurriculumExperienceUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CurriculumExperienceUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurriculumExperienceUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurriculumExperienceUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurriculumExperienceUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurriculumExperienceUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurriculumExperienceUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *CurriculumExperienceUpdateDto) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CurriculumExperienceUpdateDto) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CurriculumExperienceUpdateDto) SetPriority(v float64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CurriculumExperienceUpdateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrganization

`func (o *CurriculumExperienceUpdateDto) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CurriculumExperienceUpdateDto) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CurriculumExperienceUpdateDto) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CurriculumExperienceUpdateDto) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CurriculumExperienceUpdateDto) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CurriculumExperienceUpdateDto) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetFeatured

`func (o *CurriculumExperienceUpdateDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CurriculumExperienceUpdateDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CurriculumExperienceUpdateDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CurriculumExperienceUpdateDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetCurrent

`func (o *CurriculumExperienceUpdateDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CurriculumExperienceUpdateDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CurriculumExperienceUpdateDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CurriculumExperienceUpdateDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetValidated

`func (o *CurriculumExperienceUpdateDto) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *CurriculumExperienceUpdateDto) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *CurriculumExperienceUpdateDto) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *CurriculumExperienceUpdateDto) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetImageUrl

`func (o *CurriculumExperienceUpdateDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CurriculumExperienceUpdateDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CurriculumExperienceUpdateDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CurriculumExperienceUpdateDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CurriculumExperienceUpdateDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CurriculumExperienceUpdateDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDateFrom

`func (o *CurriculumExperienceUpdateDto) GetDateFrom() time.Time`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *CurriculumExperienceUpdateDto) GetDateFromOk() (*time.Time, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *CurriculumExperienceUpdateDto) SetDateFrom(v time.Time)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *CurriculumExperienceUpdateDto) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *CurriculumExperienceUpdateDto) GetDateTo() time.Time`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *CurriculumExperienceUpdateDto) GetDateToOk() (*time.Time, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *CurriculumExperienceUpdateDto) SetDateTo(v time.Time)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *CurriculumExperienceUpdateDto) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetDate

`func (o *CurriculumExperienceUpdateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CurriculumExperienceUpdateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CurriculumExperienceUpdateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *CurriculumExperienceUpdateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCountryLanguageId

`func (o *CurriculumExperienceUpdateDto) GetCountryLanguageId() string`

GetCountryLanguageId returns the CountryLanguageId field if non-nil, zero value otherwise.

### GetCountryLanguageIdOk

`func (o *CurriculumExperienceUpdateDto) GetCountryLanguageIdOk() (*string, bool)`

GetCountryLanguageIdOk returns a tuple with the CountryLanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLanguageId

`func (o *CurriculumExperienceUpdateDto) SetCountryLanguageId(v string)`

SetCountryLanguageId sets CountryLanguageId field to given value.

### HasCountryLanguageId

`func (o *CurriculumExperienceUpdateDto) HasCountryLanguageId() bool`

HasCountryLanguageId returns a boolean if a field has been set.

### SetCountryLanguageIdNil

`func (o *CurriculumExperienceUpdateDto) SetCountryLanguageIdNil(b bool)`

 SetCountryLanguageIdNil sets the value for CountryLanguageId to be an explicit nil

### UnsetCountryLanguageId
`func (o *CurriculumExperienceUpdateDto) UnsetCountryLanguageId()`

UnsetCountryLanguageId ensures that no value is present for CountryLanguageId, not even an explicit nil
### GetProficiencyRatingValueId

`func (o *CurriculumExperienceUpdateDto) GetProficiencyRatingValueId() string`

GetProficiencyRatingValueId returns the ProficiencyRatingValueId field if non-nil, zero value otherwise.

### GetProficiencyRatingValueIdOk

`func (o *CurriculumExperienceUpdateDto) GetProficiencyRatingValueIdOk() (*string, bool)`

GetProficiencyRatingValueIdOk returns a tuple with the ProficiencyRatingValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingValueId

`func (o *CurriculumExperienceUpdateDto) SetProficiencyRatingValueId(v string)`

SetProficiencyRatingValueId sets ProficiencyRatingValueId field to given value.

### HasProficiencyRatingValueId

`func (o *CurriculumExperienceUpdateDto) HasProficiencyRatingValueId() bool`

HasProficiencyRatingValueId returns a boolean if a field has been set.

### SetProficiencyRatingValueIdNil

`func (o *CurriculumExperienceUpdateDto) SetProficiencyRatingValueIdNil(b bool)`

 SetProficiencyRatingValueIdNil sets the value for ProficiencyRatingValueId to be an explicit nil

### UnsetProficiencyRatingValueId
`func (o *CurriculumExperienceUpdateDto) UnsetProficiencyRatingValueId()`

UnsetProficiencyRatingValueId ensures that no value is present for ProficiencyRatingValueId, not even an explicit nil
### GetProficiencyRatingModelId

`func (o *CurriculumExperienceUpdateDto) GetProficiencyRatingModelId() string`

GetProficiencyRatingModelId returns the ProficiencyRatingModelId field if non-nil, zero value otherwise.

### GetProficiencyRatingModelIdOk

`func (o *CurriculumExperienceUpdateDto) GetProficiencyRatingModelIdOk() (*string, bool)`

GetProficiencyRatingModelIdOk returns a tuple with the ProficiencyRatingModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingModelId

`func (o *CurriculumExperienceUpdateDto) SetProficiencyRatingModelId(v string)`

SetProficiencyRatingModelId sets ProficiencyRatingModelId field to given value.

### HasProficiencyRatingModelId

`func (o *CurriculumExperienceUpdateDto) HasProficiencyRatingModelId() bool`

HasProficiencyRatingModelId returns a boolean if a field has been set.

### SetProficiencyRatingModelIdNil

`func (o *CurriculumExperienceUpdateDto) SetProficiencyRatingModelIdNil(b bool)`

 SetProficiencyRatingModelIdNil sets the value for ProficiencyRatingModelId to be an explicit nil

### UnsetProficiencyRatingModelId
`func (o *CurriculumExperienceUpdateDto) UnsetProficiencyRatingModelId()`

UnsetProficiencyRatingModelId ensures that no value is present for ProficiencyRatingModelId, not even an explicit nil
### GetAchievements

`func (o *CurriculumExperienceUpdateDto) GetAchievements() string`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *CurriculumExperienceUpdateDto) GetAchievementsOk() (*string, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *CurriculumExperienceUpdateDto) SetAchievements(v string)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *CurriculumExperienceUpdateDto) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### SetAchievementsNil

`func (o *CurriculumExperienceUpdateDto) SetAchievementsNil(b bool)`

 SetAchievementsNil sets the value for Achievements to be an explicit nil

### UnsetAchievements
`func (o *CurriculumExperienceUpdateDto) UnsetAchievements()`

UnsetAchievements ensures that no value is present for Achievements, not even an explicit nil
### GetResponsibilities

`func (o *CurriculumExperienceUpdateDto) GetResponsibilities() string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *CurriculumExperienceUpdateDto) GetResponsibilitiesOk() (*string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibilities

`func (o *CurriculumExperienceUpdateDto) SetResponsibilities(v string)`

SetResponsibilities sets Responsibilities field to given value.

### HasResponsibilities

`func (o *CurriculumExperienceUpdateDto) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### SetResponsibilitiesNil

`func (o *CurriculumExperienceUpdateDto) SetResponsibilitiesNil(b bool)`

 SetResponsibilitiesNil sets the value for Responsibilities to be an explicit nil

### UnsetResponsibilities
`func (o *CurriculumExperienceUpdateDto) UnsetResponsibilities()`

UnsetResponsibilities ensures that no value is present for Responsibilities, not even an explicit nil
### GetEmployerProfileId

`func (o *CurriculumExperienceUpdateDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *CurriculumExperienceUpdateDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *CurriculumExperienceUpdateDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *CurriculumExperienceUpdateDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *CurriculumExperienceUpdateDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *CurriculumExperienceUpdateDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


