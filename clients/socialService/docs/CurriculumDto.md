# CurriculumDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**JobApplicantProfileId** | Pointer to **NullableString** |  | [optional] 
**Undergraduate** | Pointer to **bool** |  | [optional] 
**CertifiedProfessional** | Pointer to **bool** |  | [optional] 
**Sales** | Pointer to **bool** |  | [optional] 
**Others** | Pointer to **bool** |  | [optional] 
**Finance** | Pointer to **bool** |  | [optional] 
**Research** | Pointer to **bool** |  | [optional] 
**Advocate** | Pointer to **bool** |  | [optional] 
**Marketing** | Pointer to **bool** |  | [optional] 
**Education** | Pointer to **bool** |  | [optional] 
**Executive** | Pointer to **bool** |  | [optional] 
**Commercial** | Pointer to **bool** |  | [optional] 
**GraphicDesign** | Pointer to **bool** |  | [optional] 
**Sustainability** | Pointer to **bool** |  | [optional] 
**Administrative** | Pointer to **bool** |  | [optional] 
**HumanResources** | Pointer to **bool** |  | [optional] 
**SoundEngineering** | Pointer to **bool** |  | [optional] 
**CloudEngineering** | Pointer to **bool** |  | [optional] 
**FirstLevelSupport** | Pointer to **bool** |  | [optional] 
**SecondLevelSupport** | Pointer to **bool** |  | [optional] 
**SoftwareEngineering** | Pointer to **bool** |  | [optional] 
**PartnerAccountRepresentative** | Pointer to **bool** |  | [optional] 
**StartupSuccessRepresentative** | Pointer to **bool** |  | [optional] 
**CustomerSuccessRepresentative** | Pointer to **bool** |  | [optional] 

## Methods

### NewCurriculumDto

`func NewCurriculumDto() *CurriculumDto`

NewCurriculumDto instantiates a new CurriculumDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurriculumDtoWithDefaults

`func NewCurriculumDtoWithDefaults() *CurriculumDto`

NewCurriculumDtoWithDefaults instantiates a new CurriculumDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurriculumDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurriculumDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurriculumDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurriculumDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CurriculumDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CurriculumDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CurriculumDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CurriculumDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CurriculumDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CurriculumDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CurriculumDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CurriculumDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSocialProfileId

`func (o *CurriculumDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *CurriculumDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *CurriculumDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *CurriculumDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *CurriculumDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *CurriculumDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetJobApplicantProfileId

`func (o *CurriculumDto) GetJobApplicantProfileId() string`

GetJobApplicantProfileId returns the JobApplicantProfileId field if non-nil, zero value otherwise.

### GetJobApplicantProfileIdOk

`func (o *CurriculumDto) GetJobApplicantProfileIdOk() (*string, bool)`

GetJobApplicantProfileIdOk returns a tuple with the JobApplicantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobApplicantProfileId

`func (o *CurriculumDto) SetJobApplicantProfileId(v string)`

SetJobApplicantProfileId sets JobApplicantProfileId field to given value.

### HasJobApplicantProfileId

`func (o *CurriculumDto) HasJobApplicantProfileId() bool`

HasJobApplicantProfileId returns a boolean if a field has been set.

### SetJobApplicantProfileIdNil

`func (o *CurriculumDto) SetJobApplicantProfileIdNil(b bool)`

 SetJobApplicantProfileIdNil sets the value for JobApplicantProfileId to be an explicit nil

### UnsetJobApplicantProfileId
`func (o *CurriculumDto) UnsetJobApplicantProfileId()`

UnsetJobApplicantProfileId ensures that no value is present for JobApplicantProfileId, not even an explicit nil
### GetUndergraduate

`func (o *CurriculumDto) GetUndergraduate() bool`

GetUndergraduate returns the Undergraduate field if non-nil, zero value otherwise.

### GetUndergraduateOk

`func (o *CurriculumDto) GetUndergraduateOk() (*bool, bool)`

GetUndergraduateOk returns a tuple with the Undergraduate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndergraduate

`func (o *CurriculumDto) SetUndergraduate(v bool)`

SetUndergraduate sets Undergraduate field to given value.

### HasUndergraduate

`func (o *CurriculumDto) HasUndergraduate() bool`

HasUndergraduate returns a boolean if a field has been set.

### GetCertifiedProfessional

`func (o *CurriculumDto) GetCertifiedProfessional() bool`

GetCertifiedProfessional returns the CertifiedProfessional field if non-nil, zero value otherwise.

### GetCertifiedProfessionalOk

`func (o *CurriculumDto) GetCertifiedProfessionalOk() (*bool, bool)`

GetCertifiedProfessionalOk returns a tuple with the CertifiedProfessional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedProfessional

`func (o *CurriculumDto) SetCertifiedProfessional(v bool)`

SetCertifiedProfessional sets CertifiedProfessional field to given value.

### HasCertifiedProfessional

`func (o *CurriculumDto) HasCertifiedProfessional() bool`

HasCertifiedProfessional returns a boolean if a field has been set.

### GetSales

`func (o *CurriculumDto) GetSales() bool`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CurriculumDto) GetSalesOk() (*bool, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CurriculumDto) SetSales(v bool)`

SetSales sets Sales field to given value.

### HasSales

`func (o *CurriculumDto) HasSales() bool`

HasSales returns a boolean if a field has been set.

### GetOthers

`func (o *CurriculumDto) GetOthers() bool`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *CurriculumDto) GetOthersOk() (*bool, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *CurriculumDto) SetOthers(v bool)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *CurriculumDto) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetFinance

`func (o *CurriculumDto) GetFinance() bool`

GetFinance returns the Finance field if non-nil, zero value otherwise.

### GetFinanceOk

`func (o *CurriculumDto) GetFinanceOk() (*bool, bool)`

GetFinanceOk returns a tuple with the Finance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinance

`func (o *CurriculumDto) SetFinance(v bool)`

SetFinance sets Finance field to given value.

### HasFinance

`func (o *CurriculumDto) HasFinance() bool`

HasFinance returns a boolean if a field has been set.

### GetResearch

`func (o *CurriculumDto) GetResearch() bool`

GetResearch returns the Research field if non-nil, zero value otherwise.

### GetResearchOk

`func (o *CurriculumDto) GetResearchOk() (*bool, bool)`

GetResearchOk returns a tuple with the Research field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearch

`func (o *CurriculumDto) SetResearch(v bool)`

SetResearch sets Research field to given value.

### HasResearch

`func (o *CurriculumDto) HasResearch() bool`

HasResearch returns a boolean if a field has been set.

### GetAdvocate

`func (o *CurriculumDto) GetAdvocate() bool`

GetAdvocate returns the Advocate field if non-nil, zero value otherwise.

### GetAdvocateOk

`func (o *CurriculumDto) GetAdvocateOk() (*bool, bool)`

GetAdvocateOk returns a tuple with the Advocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocate

`func (o *CurriculumDto) SetAdvocate(v bool)`

SetAdvocate sets Advocate field to given value.

### HasAdvocate

`func (o *CurriculumDto) HasAdvocate() bool`

HasAdvocate returns a boolean if a field has been set.

### GetMarketing

`func (o *CurriculumDto) GetMarketing() bool`

GetMarketing returns the Marketing field if non-nil, zero value otherwise.

### GetMarketingOk

`func (o *CurriculumDto) GetMarketingOk() (*bool, bool)`

GetMarketingOk returns a tuple with the Marketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketing

`func (o *CurriculumDto) SetMarketing(v bool)`

SetMarketing sets Marketing field to given value.

### HasMarketing

`func (o *CurriculumDto) HasMarketing() bool`

HasMarketing returns a boolean if a field has been set.

### GetEducation

`func (o *CurriculumDto) GetEducation() bool`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *CurriculumDto) GetEducationOk() (*bool, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *CurriculumDto) SetEducation(v bool)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *CurriculumDto) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetExecutive

`func (o *CurriculumDto) GetExecutive() bool`

GetExecutive returns the Executive field if non-nil, zero value otherwise.

### GetExecutiveOk

`func (o *CurriculumDto) GetExecutiveOk() (*bool, bool)`

GetExecutiveOk returns a tuple with the Executive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutive

`func (o *CurriculumDto) SetExecutive(v bool)`

SetExecutive sets Executive field to given value.

### HasExecutive

`func (o *CurriculumDto) HasExecutive() bool`

HasExecutive returns a boolean if a field has been set.

### GetCommercial

`func (o *CurriculumDto) GetCommercial() bool`

GetCommercial returns the Commercial field if non-nil, zero value otherwise.

### GetCommercialOk

`func (o *CurriculumDto) GetCommercialOk() (*bool, bool)`

GetCommercialOk returns a tuple with the Commercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercial

`func (o *CurriculumDto) SetCommercial(v bool)`

SetCommercial sets Commercial field to given value.

### HasCommercial

`func (o *CurriculumDto) HasCommercial() bool`

HasCommercial returns a boolean if a field has been set.

### GetGraphicDesign

`func (o *CurriculumDto) GetGraphicDesign() bool`

GetGraphicDesign returns the GraphicDesign field if non-nil, zero value otherwise.

### GetGraphicDesignOk

`func (o *CurriculumDto) GetGraphicDesignOk() (*bool, bool)`

GetGraphicDesignOk returns a tuple with the GraphicDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicDesign

`func (o *CurriculumDto) SetGraphicDesign(v bool)`

SetGraphicDesign sets GraphicDesign field to given value.

### HasGraphicDesign

`func (o *CurriculumDto) HasGraphicDesign() bool`

HasGraphicDesign returns a boolean if a field has been set.

### GetSustainability

`func (o *CurriculumDto) GetSustainability() bool`

GetSustainability returns the Sustainability field if non-nil, zero value otherwise.

### GetSustainabilityOk

`func (o *CurriculumDto) GetSustainabilityOk() (*bool, bool)`

GetSustainabilityOk returns a tuple with the Sustainability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainability

`func (o *CurriculumDto) SetSustainability(v bool)`

SetSustainability sets Sustainability field to given value.

### HasSustainability

`func (o *CurriculumDto) HasSustainability() bool`

HasSustainability returns a boolean if a field has been set.

### GetAdministrative

`func (o *CurriculumDto) GetAdministrative() bool`

GetAdministrative returns the Administrative field if non-nil, zero value otherwise.

### GetAdministrativeOk

`func (o *CurriculumDto) GetAdministrativeOk() (*bool, bool)`

GetAdministrativeOk returns a tuple with the Administrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrative

`func (o *CurriculumDto) SetAdministrative(v bool)`

SetAdministrative sets Administrative field to given value.

### HasAdministrative

`func (o *CurriculumDto) HasAdministrative() bool`

HasAdministrative returns a boolean if a field has been set.

### GetHumanResources

`func (o *CurriculumDto) GetHumanResources() bool`

GetHumanResources returns the HumanResources field if non-nil, zero value otherwise.

### GetHumanResourcesOk

`func (o *CurriculumDto) GetHumanResourcesOk() (*bool, bool)`

GetHumanResourcesOk returns a tuple with the HumanResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanResources

`func (o *CurriculumDto) SetHumanResources(v bool)`

SetHumanResources sets HumanResources field to given value.

### HasHumanResources

`func (o *CurriculumDto) HasHumanResources() bool`

HasHumanResources returns a boolean if a field has been set.

### GetSoundEngineering

`func (o *CurriculumDto) GetSoundEngineering() bool`

GetSoundEngineering returns the SoundEngineering field if non-nil, zero value otherwise.

### GetSoundEngineeringOk

`func (o *CurriculumDto) GetSoundEngineeringOk() (*bool, bool)`

GetSoundEngineeringOk returns a tuple with the SoundEngineering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundEngineering

`func (o *CurriculumDto) SetSoundEngineering(v bool)`

SetSoundEngineering sets SoundEngineering field to given value.

### HasSoundEngineering

`func (o *CurriculumDto) HasSoundEngineering() bool`

HasSoundEngineering returns a boolean if a field has been set.

### GetCloudEngineering

`func (o *CurriculumDto) GetCloudEngineering() bool`

GetCloudEngineering returns the CloudEngineering field if non-nil, zero value otherwise.

### GetCloudEngineeringOk

`func (o *CurriculumDto) GetCloudEngineeringOk() (*bool, bool)`

GetCloudEngineeringOk returns a tuple with the CloudEngineering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEngineering

`func (o *CurriculumDto) SetCloudEngineering(v bool)`

SetCloudEngineering sets CloudEngineering field to given value.

### HasCloudEngineering

`func (o *CurriculumDto) HasCloudEngineering() bool`

HasCloudEngineering returns a boolean if a field has been set.

### GetFirstLevelSupport

`func (o *CurriculumDto) GetFirstLevelSupport() bool`

GetFirstLevelSupport returns the FirstLevelSupport field if non-nil, zero value otherwise.

### GetFirstLevelSupportOk

`func (o *CurriculumDto) GetFirstLevelSupportOk() (*bool, bool)`

GetFirstLevelSupportOk returns a tuple with the FirstLevelSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLevelSupport

`func (o *CurriculumDto) SetFirstLevelSupport(v bool)`

SetFirstLevelSupport sets FirstLevelSupport field to given value.

### HasFirstLevelSupport

`func (o *CurriculumDto) HasFirstLevelSupport() bool`

HasFirstLevelSupport returns a boolean if a field has been set.

### GetSecondLevelSupport

`func (o *CurriculumDto) GetSecondLevelSupport() bool`

GetSecondLevelSupport returns the SecondLevelSupport field if non-nil, zero value otherwise.

### GetSecondLevelSupportOk

`func (o *CurriculumDto) GetSecondLevelSupportOk() (*bool, bool)`

GetSecondLevelSupportOk returns a tuple with the SecondLevelSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLevelSupport

`func (o *CurriculumDto) SetSecondLevelSupport(v bool)`

SetSecondLevelSupport sets SecondLevelSupport field to given value.

### HasSecondLevelSupport

`func (o *CurriculumDto) HasSecondLevelSupport() bool`

HasSecondLevelSupport returns a boolean if a field has been set.

### GetSoftwareEngineering

`func (o *CurriculumDto) GetSoftwareEngineering() bool`

GetSoftwareEngineering returns the SoftwareEngineering field if non-nil, zero value otherwise.

### GetSoftwareEngineeringOk

`func (o *CurriculumDto) GetSoftwareEngineeringOk() (*bool, bool)`

GetSoftwareEngineeringOk returns a tuple with the SoftwareEngineering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEngineering

`func (o *CurriculumDto) SetSoftwareEngineering(v bool)`

SetSoftwareEngineering sets SoftwareEngineering field to given value.

### HasSoftwareEngineering

`func (o *CurriculumDto) HasSoftwareEngineering() bool`

HasSoftwareEngineering returns a boolean if a field has been set.

### GetPartnerAccountRepresentative

`func (o *CurriculumDto) GetPartnerAccountRepresentative() bool`

GetPartnerAccountRepresentative returns the PartnerAccountRepresentative field if non-nil, zero value otherwise.

### GetPartnerAccountRepresentativeOk

`func (o *CurriculumDto) GetPartnerAccountRepresentativeOk() (*bool, bool)`

GetPartnerAccountRepresentativeOk returns a tuple with the PartnerAccountRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAccountRepresentative

`func (o *CurriculumDto) SetPartnerAccountRepresentative(v bool)`

SetPartnerAccountRepresentative sets PartnerAccountRepresentative field to given value.

### HasPartnerAccountRepresentative

`func (o *CurriculumDto) HasPartnerAccountRepresentative() bool`

HasPartnerAccountRepresentative returns a boolean if a field has been set.

### GetStartupSuccessRepresentative

`func (o *CurriculumDto) GetStartupSuccessRepresentative() bool`

GetStartupSuccessRepresentative returns the StartupSuccessRepresentative field if non-nil, zero value otherwise.

### GetStartupSuccessRepresentativeOk

`func (o *CurriculumDto) GetStartupSuccessRepresentativeOk() (*bool, bool)`

GetStartupSuccessRepresentativeOk returns a tuple with the StartupSuccessRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupSuccessRepresentative

`func (o *CurriculumDto) SetStartupSuccessRepresentative(v bool)`

SetStartupSuccessRepresentative sets StartupSuccessRepresentative field to given value.

### HasStartupSuccessRepresentative

`func (o *CurriculumDto) HasStartupSuccessRepresentative() bool`

HasStartupSuccessRepresentative returns a boolean if a field has been set.

### GetCustomerSuccessRepresentative

`func (o *CurriculumDto) GetCustomerSuccessRepresentative() bool`

GetCustomerSuccessRepresentative returns the CustomerSuccessRepresentative field if non-nil, zero value otherwise.

### GetCustomerSuccessRepresentativeOk

`func (o *CurriculumDto) GetCustomerSuccessRepresentativeOk() (*bool, bool)`

GetCustomerSuccessRepresentativeOk returns a tuple with the CustomerSuccessRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSuccessRepresentative

`func (o *CurriculumDto) SetCustomerSuccessRepresentative(v bool)`

SetCustomerSuccessRepresentative sets CustomerSuccessRepresentative field to given value.

### HasCustomerSuccessRepresentative

`func (o *CurriculumDto) HasCustomerSuccessRepresentative() bool`

HasCustomerSuccessRepresentative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


