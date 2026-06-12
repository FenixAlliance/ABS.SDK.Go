# JobOfferCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**ExpectedHireDate** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TechnicalSkills** | Pointer to **NullableString** |  | [optional] 
**NonTechnicalSkills** | Pointer to **NullableString** |  | [optional] 
**Certifications** | Pointer to **NullableString** |  | [optional] 
**ProjectExperience** | Pointer to **NullableString** |  | [optional] 
**Technologies** | Pointer to **NullableString** |  | [optional] 
**Benefits** | Pointer to **NullableString** |  | [optional] 
**IsOfficialJobOffer** | Pointer to **bool** |  | [optional] 
**IsRemoteJobOffer** | Pointer to **bool** |  | [optional] 
**IsMidTimeJobOffer** | Pointer to **bool** |  | [optional] 
**IsUndergraduateOption** | Pointer to **bool** |  | [optional] 
**MinOverallExperienceYears** | Pointer to **int32** |  | [optional] 
**AvailiablePositionsCount** | Pointer to **int32** |  | [optional] 
**MinSalaryAmount** | Pointer to **float64** |  | [optional] 
**MaxSalaryAmount** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**JobFieldId** | Pointer to **NullableString** |  | [optional] 
**EmployerProfileId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CountryStateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**ExternalUrl** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**DataLabel** | Pointer to **NullableString** |  | [optional] 
**Data1** | Pointer to **NullableString** |  | [optional] 
**Data1Label** | Pointer to **NullableString** |  | [optional] 
**Data2** | Pointer to **NullableString** |  | [optional] 
**Data2Label** | Pointer to **NullableString** |  | [optional] 
**Data3** | Pointer to **NullableString** |  | [optional] 
**Data3Label** | Pointer to **NullableString** |  | [optional] 
**Data4** | Pointer to **NullableString** |  | [optional] 
**Data4Label** | Pointer to **NullableString** |  | [optional] 
**Data5** | Pointer to **NullableString** |  | [optional] 
**Data5Label** | Pointer to **NullableString** |  | [optional] 
**Data6** | Pointer to **NullableString** |  | [optional] 
**Data6Label** | Pointer to **NullableString** |  | [optional] 
**Data7** | Pointer to **NullableString** |  | [optional] 
**Data7Label** | Pointer to **NullableString** |  | [optional] 
**Data8** | Pointer to **NullableString** |  | [optional] 
**Data8Label** | Pointer to **NullableString** |  | [optional] 
**Data9** | Pointer to **NullableString** |  | [optional] 
**Data9Label** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferCreateDto

`func NewJobOfferCreateDto() *JobOfferCreateDto`

NewJobOfferCreateDto instantiates a new JobOfferCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferCreateDtoWithDefaults

`func NewJobOfferCreateDtoWithDefaults() *JobOfferCreateDto`

NewJobOfferCreateDtoWithDefaults instantiates a new JobOfferCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JobOfferCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRemote

`func (o *JobOfferCreateDto) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *JobOfferCreateDto) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *JobOfferCreateDto) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *JobOfferCreateDto) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetExpectedHireDate

`func (o *JobOfferCreateDto) GetExpectedHireDate() time.Time`

GetExpectedHireDate returns the ExpectedHireDate field if non-nil, zero value otherwise.

### GetExpectedHireDateOk

`func (o *JobOfferCreateDto) GetExpectedHireDateOk() (*time.Time, bool)`

GetExpectedHireDateOk returns a tuple with the ExpectedHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHireDate

`func (o *JobOfferCreateDto) SetExpectedHireDate(v time.Time)`

SetExpectedHireDate sets ExpectedHireDate field to given value.

### HasExpectedHireDate

`func (o *JobOfferCreateDto) HasExpectedHireDate() bool`

HasExpectedHireDate returns a boolean if a field has been set.

### GetTitle

`func (o *JobOfferCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JobOfferCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JobOfferCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JobOfferCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *JobOfferCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *JobOfferCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *JobOfferCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobOfferCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobOfferCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobOfferCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JobOfferCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JobOfferCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTechnicalSkills

`func (o *JobOfferCreateDto) GetTechnicalSkills() string`

GetTechnicalSkills returns the TechnicalSkills field if non-nil, zero value otherwise.

### GetTechnicalSkillsOk

`func (o *JobOfferCreateDto) GetTechnicalSkillsOk() (*string, bool)`

GetTechnicalSkillsOk returns a tuple with the TechnicalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalSkills

`func (o *JobOfferCreateDto) SetTechnicalSkills(v string)`

SetTechnicalSkills sets TechnicalSkills field to given value.

### HasTechnicalSkills

`func (o *JobOfferCreateDto) HasTechnicalSkills() bool`

HasTechnicalSkills returns a boolean if a field has been set.

### SetTechnicalSkillsNil

`func (o *JobOfferCreateDto) SetTechnicalSkillsNil(b bool)`

 SetTechnicalSkillsNil sets the value for TechnicalSkills to be an explicit nil

### UnsetTechnicalSkills
`func (o *JobOfferCreateDto) UnsetTechnicalSkills()`

UnsetTechnicalSkills ensures that no value is present for TechnicalSkills, not even an explicit nil
### GetNonTechnicalSkills

`func (o *JobOfferCreateDto) GetNonTechnicalSkills() string`

GetNonTechnicalSkills returns the NonTechnicalSkills field if non-nil, zero value otherwise.

### GetNonTechnicalSkillsOk

`func (o *JobOfferCreateDto) GetNonTechnicalSkillsOk() (*string, bool)`

GetNonTechnicalSkillsOk returns a tuple with the NonTechnicalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTechnicalSkills

`func (o *JobOfferCreateDto) SetNonTechnicalSkills(v string)`

SetNonTechnicalSkills sets NonTechnicalSkills field to given value.

### HasNonTechnicalSkills

`func (o *JobOfferCreateDto) HasNonTechnicalSkills() bool`

HasNonTechnicalSkills returns a boolean if a field has been set.

### SetNonTechnicalSkillsNil

`func (o *JobOfferCreateDto) SetNonTechnicalSkillsNil(b bool)`

 SetNonTechnicalSkillsNil sets the value for NonTechnicalSkills to be an explicit nil

### UnsetNonTechnicalSkills
`func (o *JobOfferCreateDto) UnsetNonTechnicalSkills()`

UnsetNonTechnicalSkills ensures that no value is present for NonTechnicalSkills, not even an explicit nil
### GetCertifications

`func (o *JobOfferCreateDto) GetCertifications() string`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *JobOfferCreateDto) GetCertificationsOk() (*string, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *JobOfferCreateDto) SetCertifications(v string)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *JobOfferCreateDto) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### SetCertificationsNil

`func (o *JobOfferCreateDto) SetCertificationsNil(b bool)`

 SetCertificationsNil sets the value for Certifications to be an explicit nil

### UnsetCertifications
`func (o *JobOfferCreateDto) UnsetCertifications()`

UnsetCertifications ensures that no value is present for Certifications, not even an explicit nil
### GetProjectExperience

`func (o *JobOfferCreateDto) GetProjectExperience() string`

GetProjectExperience returns the ProjectExperience field if non-nil, zero value otherwise.

### GetProjectExperienceOk

`func (o *JobOfferCreateDto) GetProjectExperienceOk() (*string, bool)`

GetProjectExperienceOk returns a tuple with the ProjectExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectExperience

`func (o *JobOfferCreateDto) SetProjectExperience(v string)`

SetProjectExperience sets ProjectExperience field to given value.

### HasProjectExperience

`func (o *JobOfferCreateDto) HasProjectExperience() bool`

HasProjectExperience returns a boolean if a field has been set.

### SetProjectExperienceNil

`func (o *JobOfferCreateDto) SetProjectExperienceNil(b bool)`

 SetProjectExperienceNil sets the value for ProjectExperience to be an explicit nil

### UnsetProjectExperience
`func (o *JobOfferCreateDto) UnsetProjectExperience()`

UnsetProjectExperience ensures that no value is present for ProjectExperience, not even an explicit nil
### GetTechnologies

`func (o *JobOfferCreateDto) GetTechnologies() string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *JobOfferCreateDto) GetTechnologiesOk() (*string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *JobOfferCreateDto) SetTechnologies(v string)`

SetTechnologies sets Technologies field to given value.

### HasTechnologies

`func (o *JobOfferCreateDto) HasTechnologies() bool`

HasTechnologies returns a boolean if a field has been set.

### SetTechnologiesNil

`func (o *JobOfferCreateDto) SetTechnologiesNil(b bool)`

 SetTechnologiesNil sets the value for Technologies to be an explicit nil

### UnsetTechnologies
`func (o *JobOfferCreateDto) UnsetTechnologies()`

UnsetTechnologies ensures that no value is present for Technologies, not even an explicit nil
### GetBenefits

`func (o *JobOfferCreateDto) GetBenefits() string`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *JobOfferCreateDto) GetBenefitsOk() (*string, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *JobOfferCreateDto) SetBenefits(v string)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *JobOfferCreateDto) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *JobOfferCreateDto) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *JobOfferCreateDto) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetIsOfficialJobOffer

`func (o *JobOfferCreateDto) GetIsOfficialJobOffer() bool`

GetIsOfficialJobOffer returns the IsOfficialJobOffer field if non-nil, zero value otherwise.

### GetIsOfficialJobOfferOk

`func (o *JobOfferCreateDto) GetIsOfficialJobOfferOk() (*bool, bool)`

GetIsOfficialJobOfferOk returns a tuple with the IsOfficialJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficialJobOffer

`func (o *JobOfferCreateDto) SetIsOfficialJobOffer(v bool)`

SetIsOfficialJobOffer sets IsOfficialJobOffer field to given value.

### HasIsOfficialJobOffer

`func (o *JobOfferCreateDto) HasIsOfficialJobOffer() bool`

HasIsOfficialJobOffer returns a boolean if a field has been set.

### GetIsRemoteJobOffer

`func (o *JobOfferCreateDto) GetIsRemoteJobOffer() bool`

GetIsRemoteJobOffer returns the IsRemoteJobOffer field if non-nil, zero value otherwise.

### GetIsRemoteJobOfferOk

`func (o *JobOfferCreateDto) GetIsRemoteJobOfferOk() (*bool, bool)`

GetIsRemoteJobOfferOk returns a tuple with the IsRemoteJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteJobOffer

`func (o *JobOfferCreateDto) SetIsRemoteJobOffer(v bool)`

SetIsRemoteJobOffer sets IsRemoteJobOffer field to given value.

### HasIsRemoteJobOffer

`func (o *JobOfferCreateDto) HasIsRemoteJobOffer() bool`

HasIsRemoteJobOffer returns a boolean if a field has been set.

### GetIsMidTimeJobOffer

`func (o *JobOfferCreateDto) GetIsMidTimeJobOffer() bool`

GetIsMidTimeJobOffer returns the IsMidTimeJobOffer field if non-nil, zero value otherwise.

### GetIsMidTimeJobOfferOk

`func (o *JobOfferCreateDto) GetIsMidTimeJobOfferOk() (*bool, bool)`

GetIsMidTimeJobOfferOk returns a tuple with the IsMidTimeJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMidTimeJobOffer

`func (o *JobOfferCreateDto) SetIsMidTimeJobOffer(v bool)`

SetIsMidTimeJobOffer sets IsMidTimeJobOffer field to given value.

### HasIsMidTimeJobOffer

`func (o *JobOfferCreateDto) HasIsMidTimeJobOffer() bool`

HasIsMidTimeJobOffer returns a boolean if a field has been set.

### GetIsUndergraduateOption

`func (o *JobOfferCreateDto) GetIsUndergraduateOption() bool`

GetIsUndergraduateOption returns the IsUndergraduateOption field if non-nil, zero value otherwise.

### GetIsUndergraduateOptionOk

`func (o *JobOfferCreateDto) GetIsUndergraduateOptionOk() (*bool, bool)`

GetIsUndergraduateOptionOk returns a tuple with the IsUndergraduateOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUndergraduateOption

`func (o *JobOfferCreateDto) SetIsUndergraduateOption(v bool)`

SetIsUndergraduateOption sets IsUndergraduateOption field to given value.

### HasIsUndergraduateOption

`func (o *JobOfferCreateDto) HasIsUndergraduateOption() bool`

HasIsUndergraduateOption returns a boolean if a field has been set.

### GetMinOverallExperienceYears

`func (o *JobOfferCreateDto) GetMinOverallExperienceYears() int32`

GetMinOverallExperienceYears returns the MinOverallExperienceYears field if non-nil, zero value otherwise.

### GetMinOverallExperienceYearsOk

`func (o *JobOfferCreateDto) GetMinOverallExperienceYearsOk() (*int32, bool)`

GetMinOverallExperienceYearsOk returns a tuple with the MinOverallExperienceYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOverallExperienceYears

`func (o *JobOfferCreateDto) SetMinOverallExperienceYears(v int32)`

SetMinOverallExperienceYears sets MinOverallExperienceYears field to given value.

### HasMinOverallExperienceYears

`func (o *JobOfferCreateDto) HasMinOverallExperienceYears() bool`

HasMinOverallExperienceYears returns a boolean if a field has been set.

### GetAvailiablePositionsCount

`func (o *JobOfferCreateDto) GetAvailiablePositionsCount() int32`

GetAvailiablePositionsCount returns the AvailiablePositionsCount field if non-nil, zero value otherwise.

### GetAvailiablePositionsCountOk

`func (o *JobOfferCreateDto) GetAvailiablePositionsCountOk() (*int32, bool)`

GetAvailiablePositionsCountOk returns a tuple with the AvailiablePositionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailiablePositionsCount

`func (o *JobOfferCreateDto) SetAvailiablePositionsCount(v int32)`

SetAvailiablePositionsCount sets AvailiablePositionsCount field to given value.

### HasAvailiablePositionsCount

`func (o *JobOfferCreateDto) HasAvailiablePositionsCount() bool`

HasAvailiablePositionsCount returns a boolean if a field has been set.

### GetMinSalaryAmount

`func (o *JobOfferCreateDto) GetMinSalaryAmount() float64`

GetMinSalaryAmount returns the MinSalaryAmount field if non-nil, zero value otherwise.

### GetMinSalaryAmountOk

`func (o *JobOfferCreateDto) GetMinSalaryAmountOk() (*float64, bool)`

GetMinSalaryAmountOk returns a tuple with the MinSalaryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSalaryAmount

`func (o *JobOfferCreateDto) SetMinSalaryAmount(v float64)`

SetMinSalaryAmount sets MinSalaryAmount field to given value.

### HasMinSalaryAmount

`func (o *JobOfferCreateDto) HasMinSalaryAmount() bool`

HasMinSalaryAmount returns a boolean if a field has been set.

### GetMaxSalaryAmount

`func (o *JobOfferCreateDto) GetMaxSalaryAmount() float64`

GetMaxSalaryAmount returns the MaxSalaryAmount field if non-nil, zero value otherwise.

### GetMaxSalaryAmountOk

`func (o *JobOfferCreateDto) GetMaxSalaryAmountOk() (*float64, bool)`

GetMaxSalaryAmountOk returns a tuple with the MaxSalaryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSalaryAmount

`func (o *JobOfferCreateDto) SetMaxSalaryAmount(v float64)`

SetMaxSalaryAmount sets MaxSalaryAmount field to given value.

### HasMaxSalaryAmount

`func (o *JobOfferCreateDto) HasMaxSalaryAmount() bool`

HasMaxSalaryAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobOfferCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobOfferCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobOfferCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobOfferCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobOfferCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobOfferCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetJobFieldId

`func (o *JobOfferCreateDto) GetJobFieldId() string`

GetJobFieldId returns the JobFieldId field if non-nil, zero value otherwise.

### GetJobFieldIdOk

`func (o *JobOfferCreateDto) GetJobFieldIdOk() (*string, bool)`

GetJobFieldIdOk returns a tuple with the JobFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFieldId

`func (o *JobOfferCreateDto) SetJobFieldId(v string)`

SetJobFieldId sets JobFieldId field to given value.

### HasJobFieldId

`func (o *JobOfferCreateDto) HasJobFieldId() bool`

HasJobFieldId returns a boolean if a field has been set.

### SetJobFieldIdNil

`func (o *JobOfferCreateDto) SetJobFieldIdNil(b bool)`

 SetJobFieldIdNil sets the value for JobFieldId to be an explicit nil

### UnsetJobFieldId
`func (o *JobOfferCreateDto) UnsetJobFieldId()`

UnsetJobFieldId ensures that no value is present for JobFieldId, not even an explicit nil
### GetEmployerProfileId

`func (o *JobOfferCreateDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *JobOfferCreateDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *JobOfferCreateDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *JobOfferCreateDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *JobOfferCreateDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *JobOfferCreateDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil
### GetCountryId

`func (o *JobOfferCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *JobOfferCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *JobOfferCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *JobOfferCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *JobOfferCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *JobOfferCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *JobOfferCreateDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *JobOfferCreateDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *JobOfferCreateDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *JobOfferCreateDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *JobOfferCreateDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *JobOfferCreateDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCityId

`func (o *JobOfferCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *JobOfferCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *JobOfferCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *JobOfferCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *JobOfferCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *JobOfferCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetImageUrl

`func (o *JobOfferCreateDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *JobOfferCreateDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *JobOfferCreateDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *JobOfferCreateDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *JobOfferCreateDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *JobOfferCreateDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetLocation

`func (o *JobOfferCreateDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *JobOfferCreateDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *JobOfferCreateDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *JobOfferCreateDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *JobOfferCreateDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *JobOfferCreateDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetExternalUrl

`func (o *JobOfferCreateDto) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *JobOfferCreateDto) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *JobOfferCreateDto) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *JobOfferCreateDto) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### SetExternalUrlNil

`func (o *JobOfferCreateDto) SetExternalUrlNil(b bool)`

 SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil

### UnsetExternalUrl
`func (o *JobOfferCreateDto) UnsetExternalUrl()`

UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil
### GetData

`func (o *JobOfferCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobOfferCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobOfferCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JobOfferCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobOfferCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobOfferCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *JobOfferCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *JobOfferCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *JobOfferCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *JobOfferCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *JobOfferCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *JobOfferCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *JobOfferCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *JobOfferCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *JobOfferCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *JobOfferCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *JobOfferCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *JobOfferCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *JobOfferCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *JobOfferCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *JobOfferCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *JobOfferCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *JobOfferCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *JobOfferCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *JobOfferCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *JobOfferCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *JobOfferCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *JobOfferCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *JobOfferCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *JobOfferCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *JobOfferCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *JobOfferCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *JobOfferCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *JobOfferCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *JobOfferCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *JobOfferCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *JobOfferCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *JobOfferCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *JobOfferCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *JobOfferCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *JobOfferCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *JobOfferCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *JobOfferCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *JobOfferCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *JobOfferCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *JobOfferCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *JobOfferCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *JobOfferCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *JobOfferCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *JobOfferCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *JobOfferCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *JobOfferCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *JobOfferCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *JobOfferCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *JobOfferCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *JobOfferCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *JobOfferCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *JobOfferCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *JobOfferCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *JobOfferCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *JobOfferCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *JobOfferCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *JobOfferCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *JobOfferCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *JobOfferCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *JobOfferCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *JobOfferCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *JobOfferCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *JobOfferCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *JobOfferCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *JobOfferCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *JobOfferCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *JobOfferCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *JobOfferCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *JobOfferCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *JobOfferCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *JobOfferCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *JobOfferCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *JobOfferCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *JobOfferCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *JobOfferCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *JobOfferCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *JobOfferCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *JobOfferCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *JobOfferCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *JobOfferCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *JobOfferCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *JobOfferCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *JobOfferCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *JobOfferCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *JobOfferCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *JobOfferCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *JobOfferCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *JobOfferCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *JobOfferCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *JobOfferCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *JobOfferCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *JobOfferCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *JobOfferCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *JobOfferCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *JobOfferCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *JobOfferCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *JobOfferCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *JobOfferCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *JobOfferCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *JobOfferCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *JobOfferCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *JobOfferCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *JobOfferCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *JobOfferCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *JobOfferCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *JobOfferCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *JobOfferCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *JobOfferCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *JobOfferCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *JobOfferCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *JobOfferCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *JobOfferCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *JobOfferCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *JobOfferCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


