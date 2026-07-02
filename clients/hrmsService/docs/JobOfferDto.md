# JobOfferDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
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
**ViewsCount** | Pointer to **int32** |  | [optional] 
**MinSalaryAmount** | Pointer to **float64** |  | [optional] 
**MaxSalaryAmount** | Pointer to **float64** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
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
**ExternalUrl** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**JobFieldId** | Pointer to **NullableString** |  | [optional] 
**EmployerProfileId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CountryStateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferDto

`func NewJobOfferDto() *JobOfferDto`

NewJobOfferDto instantiates a new JobOfferDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferDtoWithDefaults

`func NewJobOfferDtoWithDefaults() *JobOfferDto`

NewJobOfferDtoWithDefaults instantiates a new JobOfferDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobOfferDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobOfferDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JobOfferDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JobOfferDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JobOfferDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetStatus

`func (o *JobOfferDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobOfferDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobOfferDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobOfferDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRemote

`func (o *JobOfferDto) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *JobOfferDto) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *JobOfferDto) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *JobOfferDto) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetExpectedHireDate

`func (o *JobOfferDto) GetExpectedHireDate() time.Time`

GetExpectedHireDate returns the ExpectedHireDate field if non-nil, zero value otherwise.

### GetExpectedHireDateOk

`func (o *JobOfferDto) GetExpectedHireDateOk() (*time.Time, bool)`

GetExpectedHireDateOk returns a tuple with the ExpectedHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHireDate

`func (o *JobOfferDto) SetExpectedHireDate(v time.Time)`

SetExpectedHireDate sets ExpectedHireDate field to given value.

### HasExpectedHireDate

`func (o *JobOfferDto) HasExpectedHireDate() bool`

HasExpectedHireDate returns a boolean if a field has been set.

### GetTitle

`func (o *JobOfferDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JobOfferDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JobOfferDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JobOfferDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *JobOfferDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *JobOfferDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *JobOfferDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobOfferDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobOfferDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobOfferDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JobOfferDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JobOfferDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTechnicalSkills

`func (o *JobOfferDto) GetTechnicalSkills() string`

GetTechnicalSkills returns the TechnicalSkills field if non-nil, zero value otherwise.

### GetTechnicalSkillsOk

`func (o *JobOfferDto) GetTechnicalSkillsOk() (*string, bool)`

GetTechnicalSkillsOk returns a tuple with the TechnicalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalSkills

`func (o *JobOfferDto) SetTechnicalSkills(v string)`

SetTechnicalSkills sets TechnicalSkills field to given value.

### HasTechnicalSkills

`func (o *JobOfferDto) HasTechnicalSkills() bool`

HasTechnicalSkills returns a boolean if a field has been set.

### SetTechnicalSkillsNil

`func (o *JobOfferDto) SetTechnicalSkillsNil(b bool)`

 SetTechnicalSkillsNil sets the value for TechnicalSkills to be an explicit nil

### UnsetTechnicalSkills
`func (o *JobOfferDto) UnsetTechnicalSkills()`

UnsetTechnicalSkills ensures that no value is present for TechnicalSkills, not even an explicit nil
### GetNonTechnicalSkills

`func (o *JobOfferDto) GetNonTechnicalSkills() string`

GetNonTechnicalSkills returns the NonTechnicalSkills field if non-nil, zero value otherwise.

### GetNonTechnicalSkillsOk

`func (o *JobOfferDto) GetNonTechnicalSkillsOk() (*string, bool)`

GetNonTechnicalSkillsOk returns a tuple with the NonTechnicalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTechnicalSkills

`func (o *JobOfferDto) SetNonTechnicalSkills(v string)`

SetNonTechnicalSkills sets NonTechnicalSkills field to given value.

### HasNonTechnicalSkills

`func (o *JobOfferDto) HasNonTechnicalSkills() bool`

HasNonTechnicalSkills returns a boolean if a field has been set.

### SetNonTechnicalSkillsNil

`func (o *JobOfferDto) SetNonTechnicalSkillsNil(b bool)`

 SetNonTechnicalSkillsNil sets the value for NonTechnicalSkills to be an explicit nil

### UnsetNonTechnicalSkills
`func (o *JobOfferDto) UnsetNonTechnicalSkills()`

UnsetNonTechnicalSkills ensures that no value is present for NonTechnicalSkills, not even an explicit nil
### GetCertifications

`func (o *JobOfferDto) GetCertifications() string`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *JobOfferDto) GetCertificationsOk() (*string, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *JobOfferDto) SetCertifications(v string)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *JobOfferDto) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### SetCertificationsNil

`func (o *JobOfferDto) SetCertificationsNil(b bool)`

 SetCertificationsNil sets the value for Certifications to be an explicit nil

### UnsetCertifications
`func (o *JobOfferDto) UnsetCertifications()`

UnsetCertifications ensures that no value is present for Certifications, not even an explicit nil
### GetProjectExperience

`func (o *JobOfferDto) GetProjectExperience() string`

GetProjectExperience returns the ProjectExperience field if non-nil, zero value otherwise.

### GetProjectExperienceOk

`func (o *JobOfferDto) GetProjectExperienceOk() (*string, bool)`

GetProjectExperienceOk returns a tuple with the ProjectExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectExperience

`func (o *JobOfferDto) SetProjectExperience(v string)`

SetProjectExperience sets ProjectExperience field to given value.

### HasProjectExperience

`func (o *JobOfferDto) HasProjectExperience() bool`

HasProjectExperience returns a boolean if a field has been set.

### SetProjectExperienceNil

`func (o *JobOfferDto) SetProjectExperienceNil(b bool)`

 SetProjectExperienceNil sets the value for ProjectExperience to be an explicit nil

### UnsetProjectExperience
`func (o *JobOfferDto) UnsetProjectExperience()`

UnsetProjectExperience ensures that no value is present for ProjectExperience, not even an explicit nil
### GetTechnologies

`func (o *JobOfferDto) GetTechnologies() string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *JobOfferDto) GetTechnologiesOk() (*string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *JobOfferDto) SetTechnologies(v string)`

SetTechnologies sets Technologies field to given value.

### HasTechnologies

`func (o *JobOfferDto) HasTechnologies() bool`

HasTechnologies returns a boolean if a field has been set.

### SetTechnologiesNil

`func (o *JobOfferDto) SetTechnologiesNil(b bool)`

 SetTechnologiesNil sets the value for Technologies to be an explicit nil

### UnsetTechnologies
`func (o *JobOfferDto) UnsetTechnologies()`

UnsetTechnologies ensures that no value is present for Technologies, not even an explicit nil
### GetBenefits

`func (o *JobOfferDto) GetBenefits() string`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *JobOfferDto) GetBenefitsOk() (*string, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *JobOfferDto) SetBenefits(v string)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *JobOfferDto) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *JobOfferDto) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *JobOfferDto) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetIsOfficialJobOffer

`func (o *JobOfferDto) GetIsOfficialJobOffer() bool`

GetIsOfficialJobOffer returns the IsOfficialJobOffer field if non-nil, zero value otherwise.

### GetIsOfficialJobOfferOk

`func (o *JobOfferDto) GetIsOfficialJobOfferOk() (*bool, bool)`

GetIsOfficialJobOfferOk returns a tuple with the IsOfficialJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficialJobOffer

`func (o *JobOfferDto) SetIsOfficialJobOffer(v bool)`

SetIsOfficialJobOffer sets IsOfficialJobOffer field to given value.

### HasIsOfficialJobOffer

`func (o *JobOfferDto) HasIsOfficialJobOffer() bool`

HasIsOfficialJobOffer returns a boolean if a field has been set.

### GetIsRemoteJobOffer

`func (o *JobOfferDto) GetIsRemoteJobOffer() bool`

GetIsRemoteJobOffer returns the IsRemoteJobOffer field if non-nil, zero value otherwise.

### GetIsRemoteJobOfferOk

`func (o *JobOfferDto) GetIsRemoteJobOfferOk() (*bool, bool)`

GetIsRemoteJobOfferOk returns a tuple with the IsRemoteJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteJobOffer

`func (o *JobOfferDto) SetIsRemoteJobOffer(v bool)`

SetIsRemoteJobOffer sets IsRemoteJobOffer field to given value.

### HasIsRemoteJobOffer

`func (o *JobOfferDto) HasIsRemoteJobOffer() bool`

HasIsRemoteJobOffer returns a boolean if a field has been set.

### GetIsMidTimeJobOffer

`func (o *JobOfferDto) GetIsMidTimeJobOffer() bool`

GetIsMidTimeJobOffer returns the IsMidTimeJobOffer field if non-nil, zero value otherwise.

### GetIsMidTimeJobOfferOk

`func (o *JobOfferDto) GetIsMidTimeJobOfferOk() (*bool, bool)`

GetIsMidTimeJobOfferOk returns a tuple with the IsMidTimeJobOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMidTimeJobOffer

`func (o *JobOfferDto) SetIsMidTimeJobOffer(v bool)`

SetIsMidTimeJobOffer sets IsMidTimeJobOffer field to given value.

### HasIsMidTimeJobOffer

`func (o *JobOfferDto) HasIsMidTimeJobOffer() bool`

HasIsMidTimeJobOffer returns a boolean if a field has been set.

### GetIsUndergraduateOption

`func (o *JobOfferDto) GetIsUndergraduateOption() bool`

GetIsUndergraduateOption returns the IsUndergraduateOption field if non-nil, zero value otherwise.

### GetIsUndergraduateOptionOk

`func (o *JobOfferDto) GetIsUndergraduateOptionOk() (*bool, bool)`

GetIsUndergraduateOptionOk returns a tuple with the IsUndergraduateOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUndergraduateOption

`func (o *JobOfferDto) SetIsUndergraduateOption(v bool)`

SetIsUndergraduateOption sets IsUndergraduateOption field to given value.

### HasIsUndergraduateOption

`func (o *JobOfferDto) HasIsUndergraduateOption() bool`

HasIsUndergraduateOption returns a boolean if a field has been set.

### GetMinOverallExperienceYears

`func (o *JobOfferDto) GetMinOverallExperienceYears() int32`

GetMinOverallExperienceYears returns the MinOverallExperienceYears field if non-nil, zero value otherwise.

### GetMinOverallExperienceYearsOk

`func (o *JobOfferDto) GetMinOverallExperienceYearsOk() (*int32, bool)`

GetMinOverallExperienceYearsOk returns a tuple with the MinOverallExperienceYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOverallExperienceYears

`func (o *JobOfferDto) SetMinOverallExperienceYears(v int32)`

SetMinOverallExperienceYears sets MinOverallExperienceYears field to given value.

### HasMinOverallExperienceYears

`func (o *JobOfferDto) HasMinOverallExperienceYears() bool`

HasMinOverallExperienceYears returns a boolean if a field has been set.

### GetAvailiablePositionsCount

`func (o *JobOfferDto) GetAvailiablePositionsCount() int32`

GetAvailiablePositionsCount returns the AvailiablePositionsCount field if non-nil, zero value otherwise.

### GetAvailiablePositionsCountOk

`func (o *JobOfferDto) GetAvailiablePositionsCountOk() (*int32, bool)`

GetAvailiablePositionsCountOk returns a tuple with the AvailiablePositionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailiablePositionsCount

`func (o *JobOfferDto) SetAvailiablePositionsCount(v int32)`

SetAvailiablePositionsCount sets AvailiablePositionsCount field to given value.

### HasAvailiablePositionsCount

`func (o *JobOfferDto) HasAvailiablePositionsCount() bool`

HasAvailiablePositionsCount returns a boolean if a field has been set.

### GetViewsCount

`func (o *JobOfferDto) GetViewsCount() int32`

GetViewsCount returns the ViewsCount field if non-nil, zero value otherwise.

### GetViewsCountOk

`func (o *JobOfferDto) GetViewsCountOk() (*int32, bool)`

GetViewsCountOk returns a tuple with the ViewsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsCount

`func (o *JobOfferDto) SetViewsCount(v int32)`

SetViewsCount sets ViewsCount field to given value.

### HasViewsCount

`func (o *JobOfferDto) HasViewsCount() bool`

HasViewsCount returns a boolean if a field has been set.

### GetMinSalaryAmount

`func (o *JobOfferDto) GetMinSalaryAmount() float64`

GetMinSalaryAmount returns the MinSalaryAmount field if non-nil, zero value otherwise.

### GetMinSalaryAmountOk

`func (o *JobOfferDto) GetMinSalaryAmountOk() (*float64, bool)`

GetMinSalaryAmountOk returns a tuple with the MinSalaryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSalaryAmount

`func (o *JobOfferDto) SetMinSalaryAmount(v float64)`

SetMinSalaryAmount sets MinSalaryAmount field to given value.

### HasMinSalaryAmount

`func (o *JobOfferDto) HasMinSalaryAmount() bool`

HasMinSalaryAmount returns a boolean if a field has been set.

### GetMaxSalaryAmount

`func (o *JobOfferDto) GetMaxSalaryAmount() float64`

GetMaxSalaryAmount returns the MaxSalaryAmount field if non-nil, zero value otherwise.

### GetMaxSalaryAmountOk

`func (o *JobOfferDto) GetMaxSalaryAmountOk() (*float64, bool)`

GetMaxSalaryAmountOk returns a tuple with the MaxSalaryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSalaryAmount

`func (o *JobOfferDto) SetMaxSalaryAmount(v float64)`

SetMaxSalaryAmount sets MaxSalaryAmount field to given value.

### HasMaxSalaryAmount

`func (o *JobOfferDto) HasMaxSalaryAmount() bool`

HasMaxSalaryAmount returns a boolean if a field has been set.

### GetImageUrl

`func (o *JobOfferDto) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *JobOfferDto) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *JobOfferDto) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *JobOfferDto) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *JobOfferDto) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *JobOfferDto) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetLocation

`func (o *JobOfferDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *JobOfferDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *JobOfferDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *JobOfferDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *JobOfferDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *JobOfferDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetData

`func (o *JobOfferDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobOfferDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobOfferDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JobOfferDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobOfferDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobOfferDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *JobOfferDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *JobOfferDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *JobOfferDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *JobOfferDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *JobOfferDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *JobOfferDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *JobOfferDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *JobOfferDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *JobOfferDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *JobOfferDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *JobOfferDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *JobOfferDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *JobOfferDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *JobOfferDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *JobOfferDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *JobOfferDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *JobOfferDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *JobOfferDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *JobOfferDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *JobOfferDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *JobOfferDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *JobOfferDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *JobOfferDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *JobOfferDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *JobOfferDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *JobOfferDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *JobOfferDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *JobOfferDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *JobOfferDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *JobOfferDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *JobOfferDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *JobOfferDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *JobOfferDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *JobOfferDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *JobOfferDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *JobOfferDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *JobOfferDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *JobOfferDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *JobOfferDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *JobOfferDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *JobOfferDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *JobOfferDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *JobOfferDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *JobOfferDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *JobOfferDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *JobOfferDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *JobOfferDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *JobOfferDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *JobOfferDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *JobOfferDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *JobOfferDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *JobOfferDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *JobOfferDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *JobOfferDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *JobOfferDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *JobOfferDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *JobOfferDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *JobOfferDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *JobOfferDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *JobOfferDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *JobOfferDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *JobOfferDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *JobOfferDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *JobOfferDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *JobOfferDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *JobOfferDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *JobOfferDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *JobOfferDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *JobOfferDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *JobOfferDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *JobOfferDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *JobOfferDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *JobOfferDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *JobOfferDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *JobOfferDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *JobOfferDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *JobOfferDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *JobOfferDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *JobOfferDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *JobOfferDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *JobOfferDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *JobOfferDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *JobOfferDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *JobOfferDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *JobOfferDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *JobOfferDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *JobOfferDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *JobOfferDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *JobOfferDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *JobOfferDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *JobOfferDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *JobOfferDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *JobOfferDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *JobOfferDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *JobOfferDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *JobOfferDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *JobOfferDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *JobOfferDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *JobOfferDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *JobOfferDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *JobOfferDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *JobOfferDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *JobOfferDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *JobOfferDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *JobOfferDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *JobOfferDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *JobOfferDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *JobOfferDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *JobOfferDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *JobOfferDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *JobOfferDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *JobOfferDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *JobOfferDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *JobOfferDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetExternalUrl

`func (o *JobOfferDto) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *JobOfferDto) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *JobOfferDto) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *JobOfferDto) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### SetExternalUrlNil

`func (o *JobOfferDto) SetExternalUrlNil(b bool)`

 SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil

### UnsetExternalUrl
`func (o *JobOfferDto) UnsetExternalUrl()`

UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil
### GetTenantId

`func (o *JobOfferDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JobOfferDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JobOfferDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JobOfferDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JobOfferDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JobOfferDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *JobOfferDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *JobOfferDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *JobOfferDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *JobOfferDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *JobOfferDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *JobOfferDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *JobOfferDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobOfferDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobOfferDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobOfferDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobOfferDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobOfferDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetJobFieldId

`func (o *JobOfferDto) GetJobFieldId() string`

GetJobFieldId returns the JobFieldId field if non-nil, zero value otherwise.

### GetJobFieldIdOk

`func (o *JobOfferDto) GetJobFieldIdOk() (*string, bool)`

GetJobFieldIdOk returns a tuple with the JobFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFieldId

`func (o *JobOfferDto) SetJobFieldId(v string)`

SetJobFieldId sets JobFieldId field to given value.

### HasJobFieldId

`func (o *JobOfferDto) HasJobFieldId() bool`

HasJobFieldId returns a boolean if a field has been set.

### SetJobFieldIdNil

`func (o *JobOfferDto) SetJobFieldIdNil(b bool)`

 SetJobFieldIdNil sets the value for JobFieldId to be an explicit nil

### UnsetJobFieldId
`func (o *JobOfferDto) UnsetJobFieldId()`

UnsetJobFieldId ensures that no value is present for JobFieldId, not even an explicit nil
### GetEmployerProfileId

`func (o *JobOfferDto) GetEmployerProfileId() string`

GetEmployerProfileId returns the EmployerProfileId field if non-nil, zero value otherwise.

### GetEmployerProfileIdOk

`func (o *JobOfferDto) GetEmployerProfileIdOk() (*string, bool)`

GetEmployerProfileIdOk returns a tuple with the EmployerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerProfileId

`func (o *JobOfferDto) SetEmployerProfileId(v string)`

SetEmployerProfileId sets EmployerProfileId field to given value.

### HasEmployerProfileId

`func (o *JobOfferDto) HasEmployerProfileId() bool`

HasEmployerProfileId returns a boolean if a field has been set.

### SetEmployerProfileIdNil

`func (o *JobOfferDto) SetEmployerProfileIdNil(b bool)`

 SetEmployerProfileIdNil sets the value for EmployerProfileId to be an explicit nil

### UnsetEmployerProfileId
`func (o *JobOfferDto) UnsetEmployerProfileId()`

UnsetEmployerProfileId ensures that no value is present for EmployerProfileId, not even an explicit nil
### GetCountryId

`func (o *JobOfferDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *JobOfferDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *JobOfferDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *JobOfferDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *JobOfferDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *JobOfferDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *JobOfferDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *JobOfferDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *JobOfferDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *JobOfferDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *JobOfferDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *JobOfferDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCityId

`func (o *JobOfferDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *JobOfferDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *JobOfferDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *JobOfferDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *JobOfferDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *JobOfferDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


