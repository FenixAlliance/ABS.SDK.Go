# CurriculumExperienceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewCurriculumExperienceCreateDto

`func NewCurriculumExperienceCreateDto() *CurriculumExperienceCreateDto`

NewCurriculumExperienceCreateDto instantiates a new CurriculumExperienceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurriculumExperienceCreateDtoWithDefaults

`func NewCurriculumExperienceCreateDtoWithDefaults() *CurriculumExperienceCreateDto`

NewCurriculumExperienceCreateDtoWithDefaults instantiates a new CurriculumExperienceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurriculumExperienceCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurriculumExperienceCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurriculumExperienceCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurriculumExperienceCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CurriculumExperienceCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CurriculumExperienceCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CurriculumExperienceCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CurriculumExperienceCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCurriculumId

`func (o *CurriculumExperienceCreateDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *CurriculumExperienceCreateDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *CurriculumExperienceCreateDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *CurriculumExperienceCreateDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *CurriculumExperienceCreateDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *CurriculumExperienceCreateDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetTitle

`func (o *CurriculumExperienceCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CurriculumExperienceCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CurriculumExperienceCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CurriculumExperienceCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CurriculumExperienceCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CurriculumExperienceCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CurriculumExperienceCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurriculumExperienceCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurriculumExperienceCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurriculumExperienceCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurriculumExperienceCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurriculumExperienceCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *CurriculumExperienceCreateDto) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CurriculumExperienceCreateDto) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CurriculumExperienceCreateDto) SetPriority(v float64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CurriculumExperienceCreateDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrganization

`func (o *CurriculumExperienceCreateDto) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CurriculumExperienceCreateDto) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CurriculumExperienceCreateDto) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CurriculumExperienceCreateDto) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CurriculumExperienceCreateDto) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CurriculumExperienceCreateDto) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetFeatured

`func (o *CurriculumExperienceCreateDto) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CurriculumExperienceCreateDto) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CurriculumExperienceCreateDto) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CurriculumExperienceCreateDto) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetCurrent

`func (o *CurriculumExperienceCreateDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CurriculumExperienceCreateDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CurriculumExperienceCreateDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CurriculumExperienceCreateDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetValidated

`func (o *CurriculumExperienceCreateDto) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *CurriculumExperienceCreateDto) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *CurriculumExperienceCreateDto) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *CurriculumExperienceCreateDto) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetImageUrl

`func (o *CurriculumExperienceCreateDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CurriculumExperienceCreateDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CurriculumExperienceCreateDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CurriculumExperienceCreateDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CurriculumExperienceCreateDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CurriculumExperienceCreateDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDateFrom

`func (o *CurriculumExperienceCreateDto) GetDateFrom() time.Time`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *CurriculumExperienceCreateDto) GetDateFromOk() (*time.Time, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *CurriculumExperienceCreateDto) SetDateFrom(v time.Time)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *CurriculumExperienceCreateDto) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *CurriculumExperienceCreateDto) GetDateTo() time.Time`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *CurriculumExperienceCreateDto) GetDateToOk() (*time.Time, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *CurriculumExperienceCreateDto) SetDateTo(v time.Time)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *CurriculumExperienceCreateDto) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetDate

`func (o *CurriculumExperienceCreateDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CurriculumExperienceCreateDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CurriculumExperienceCreateDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *CurriculumExperienceCreateDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCountryLanguageId

`func (o *CurriculumExperienceCreateDto) GetCountryLanguageId() string`

GetCountryLanguageId returns the CountryLanguageId field if non-nil, zero value otherwise.

### GetCountryLanguageIdOk

`func (o *CurriculumExperienceCreateDto) GetCountryLanguageIdOk() (*string, bool)`

GetCountryLanguageIdOk returns a tuple with the CountryLanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLanguageId

`func (o *CurriculumExperienceCreateDto) SetCountryLanguageId(v string)`

SetCountryLanguageId sets CountryLanguageId field to given value.

### HasCountryLanguageId

`func (o *CurriculumExperienceCreateDto) HasCountryLanguageId() bool`

HasCountryLanguageId returns a boolean if a field has been set.

### SetCountryLanguageIdNil

`func (o *CurriculumExperienceCreateDto) SetCountryLanguageIdNil(b bool)`

 SetCountryLanguageIdNil sets the value for CountryLanguageId to be an explicit nil

### UnsetCountryLanguageId
`func (o *CurriculumExperienceCreateDto) UnsetCountryLanguageId()`

UnsetCountryLanguageId ensures that no value is present for CountryLanguageId, not even an explicit nil
### GetProficiencyRatingValueId

`func (o *CurriculumExperienceCreateDto) GetProficiencyRatingValueId() string`

GetProficiencyRatingValueId returns the ProficiencyRatingValueId field if non-nil, zero value otherwise.

### GetProficiencyRatingValueIdOk

`func (o *CurriculumExperienceCreateDto) GetProficiencyRatingValueIdOk() (*string, bool)`

GetProficiencyRatingValueIdOk returns a tuple with the ProficiencyRatingValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingValueId

`func (o *CurriculumExperienceCreateDto) SetProficiencyRatingValueId(v string)`

SetProficiencyRatingValueId sets ProficiencyRatingValueId field to given value.

### HasProficiencyRatingValueId

`func (o *CurriculumExperienceCreateDto) HasProficiencyRatingValueId() bool`

HasProficiencyRatingValueId returns a boolean if a field has been set.

### SetProficiencyRatingValueIdNil

`func (o *CurriculumExperienceCreateDto) SetProficiencyRatingValueIdNil(b bool)`

 SetProficiencyRatingValueIdNil sets the value for ProficiencyRatingValueId to be an explicit nil

### UnsetProficiencyRatingValueId
`func (o *CurriculumExperienceCreateDto) UnsetProficiencyRatingValueId()`

UnsetProficiencyRatingValueId ensures that no value is present for ProficiencyRatingValueId, not even an explicit nil
### GetProficiencyRatingModelId

`func (o *CurriculumExperienceCreateDto) GetProficiencyRatingModelId() string`

GetProficiencyRatingModelId returns the ProficiencyRatingModelId field if non-nil, zero value otherwise.

### GetProficiencyRatingModelIdOk

`func (o *CurriculumExperienceCreateDto) GetProficiencyRatingModelIdOk() (*string, bool)`

GetProficiencyRatingModelIdOk returns a tuple with the ProficiencyRatingModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProficiencyRatingModelId

`func (o *CurriculumExperienceCreateDto) SetProficiencyRatingModelId(v string)`

SetProficiencyRatingModelId sets ProficiencyRatingModelId field to given value.

### HasProficiencyRatingModelId

`func (o *CurriculumExperienceCreateDto) HasProficiencyRatingModelId() bool`

HasProficiencyRatingModelId returns a boolean if a field has been set.

### SetProficiencyRatingModelIdNil

`func (o *CurriculumExperienceCreateDto) SetProficiencyRatingModelIdNil(b bool)`

 SetProficiencyRatingModelIdNil sets the value for ProficiencyRatingModelId to be an explicit nil

### UnsetProficiencyRatingModelId
`func (o *CurriculumExperienceCreateDto) UnsetProficiencyRatingModelId()`

UnsetProficiencyRatingModelId ensures that no value is present for ProficiencyRatingModelId, not even an explicit nil
### GetAchievements

`func (o *CurriculumExperienceCreateDto) GetAchievements() string`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *CurriculumExperienceCreateDto) GetAchievementsOk() (*string, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *CurriculumExperienceCreateDto) SetAchievements(v string)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *CurriculumExperienceCreateDto) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### SetAchievementsNil

`func (o *CurriculumExperienceCreateDto) SetAchievementsNil(b bool)`

 SetAchievementsNil sets the value for Achievements to be an explicit nil

### UnsetAchievements
`func (o *CurriculumExperienceCreateDto) UnsetAchievements()`

UnsetAchievements ensures that no value is present for Achievements, not even an explicit nil
### GetResponsibilities

`func (o *CurriculumExperienceCreateDto) GetResponsibilities() string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *CurriculumExperienceCreateDto) GetResponsibilitiesOk() (*string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibilities

`func (o *CurriculumExperienceCreateDto) SetResponsibilities(v string)`

SetResponsibilities sets Responsibilities field to given value.

### HasResponsibilities

`func (o *CurriculumExperienceCreateDto) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### SetResponsibilitiesNil

`func (o *CurriculumExperienceCreateDto) SetResponsibilitiesNil(b bool)`

 SetResponsibilitiesNil sets the value for Responsibilities to be an explicit nil

### UnsetResponsibilities
`func (o *CurriculumExperienceCreateDto) UnsetResponsibilities()`

UnsetResponsibilities ensures that no value is present for Responsibilities, not even an explicit nil
### GetEmployerProfileId

`func (o *CurriculumExperienceCreateDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *CurriculumExperienceCreateDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *CurriculumExperienceCreateDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *CurriculumExperienceCreateDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *CurriculumExperienceCreateDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *CurriculumExperienceCreateDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


