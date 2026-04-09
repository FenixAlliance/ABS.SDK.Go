# SeoOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**SocialImage** | Pointer to **NullableString** |  | [optional] 
**TitleSuffix** | Pointer to **NullableString** |  | [optional] 
**BingVerificationCode** | Pointer to **NullableString** |  | [optional] 
**GoogleVerificationCode** | Pointer to **NullableString** |  | [optional] 
**PinterestVerificationCode** | Pointer to **NullableString** |  | [optional] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**SameAs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSeoOptions

`func NewSeoOptions() *SeoOptions`

NewSeoOptions instantiates a new SeoOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoOptionsWithDefaults

`func NewSeoOptionsWithDefaults() *SeoOptions`

NewSeoOptionsWithDefaults instantiates a new SeoOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *SeoOptions) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoOptions) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoOptions) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoOptions) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *SeoOptions) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *SeoOptions) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetDescription

`func (o *SeoOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeoOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeoOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeoOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SeoOptions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SeoOptions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLogo

`func (o *SeoOptions) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SeoOptions) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SeoOptions) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SeoOptions) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *SeoOptions) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *SeoOptions) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetSocialImage

`func (o *SeoOptions) GetSocialImage() string`

GetSocialImage returns the SocialImage field if non-nil, zero value otherwise.

### GetSocialImageOk

`func (o *SeoOptions) GetSocialImageOk() (*string, bool)`

GetSocialImageOk returns a tuple with the SocialImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialImage

`func (o *SeoOptions) SetSocialImage(v string)`

SetSocialImage sets SocialImage field to given value.

### HasSocialImage

`func (o *SeoOptions) HasSocialImage() bool`

HasSocialImage returns a boolean if a field has been set.

### SetSocialImageNil

`func (o *SeoOptions) SetSocialImageNil(b bool)`

 SetSocialImageNil sets the value for SocialImage to be an explicit nil

### UnsetSocialImage
`func (o *SeoOptions) UnsetSocialImage()`

UnsetSocialImage ensures that no value is present for SocialImage, not even an explicit nil
### GetTitleSuffix

`func (o *SeoOptions) GetTitleSuffix() string`

GetTitleSuffix returns the TitleSuffix field if non-nil, zero value otherwise.

### GetTitleSuffixOk

`func (o *SeoOptions) GetTitleSuffixOk() (*string, bool)`

GetTitleSuffixOk returns a tuple with the TitleSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSuffix

`func (o *SeoOptions) SetTitleSuffix(v string)`

SetTitleSuffix sets TitleSuffix field to given value.

### HasTitleSuffix

`func (o *SeoOptions) HasTitleSuffix() bool`

HasTitleSuffix returns a boolean if a field has been set.

### SetTitleSuffixNil

`func (o *SeoOptions) SetTitleSuffixNil(b bool)`

 SetTitleSuffixNil sets the value for TitleSuffix to be an explicit nil

### UnsetTitleSuffix
`func (o *SeoOptions) UnsetTitleSuffix()`

UnsetTitleSuffix ensures that no value is present for TitleSuffix, not even an explicit nil
### GetBingVerificationCode

`func (o *SeoOptions) GetBingVerificationCode() string`

GetBingVerificationCode returns the BingVerificationCode field if non-nil, zero value otherwise.

### GetBingVerificationCodeOk

`func (o *SeoOptions) GetBingVerificationCodeOk() (*string, bool)`

GetBingVerificationCodeOk returns a tuple with the BingVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBingVerificationCode

`func (o *SeoOptions) SetBingVerificationCode(v string)`

SetBingVerificationCode sets BingVerificationCode field to given value.

### HasBingVerificationCode

`func (o *SeoOptions) HasBingVerificationCode() bool`

HasBingVerificationCode returns a boolean if a field has been set.

### SetBingVerificationCodeNil

`func (o *SeoOptions) SetBingVerificationCodeNil(b bool)`

 SetBingVerificationCodeNil sets the value for BingVerificationCode to be an explicit nil

### UnsetBingVerificationCode
`func (o *SeoOptions) UnsetBingVerificationCode()`

UnsetBingVerificationCode ensures that no value is present for BingVerificationCode, not even an explicit nil
### GetGoogleVerificationCode

`func (o *SeoOptions) GetGoogleVerificationCode() string`

GetGoogleVerificationCode returns the GoogleVerificationCode field if non-nil, zero value otherwise.

### GetGoogleVerificationCodeOk

`func (o *SeoOptions) GetGoogleVerificationCodeOk() (*string, bool)`

GetGoogleVerificationCodeOk returns a tuple with the GoogleVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleVerificationCode

`func (o *SeoOptions) SetGoogleVerificationCode(v string)`

SetGoogleVerificationCode sets GoogleVerificationCode field to given value.

### HasGoogleVerificationCode

`func (o *SeoOptions) HasGoogleVerificationCode() bool`

HasGoogleVerificationCode returns a boolean if a field has been set.

### SetGoogleVerificationCodeNil

`func (o *SeoOptions) SetGoogleVerificationCodeNil(b bool)`

 SetGoogleVerificationCodeNil sets the value for GoogleVerificationCode to be an explicit nil

### UnsetGoogleVerificationCode
`func (o *SeoOptions) UnsetGoogleVerificationCode()`

UnsetGoogleVerificationCode ensures that no value is present for GoogleVerificationCode, not even an explicit nil
### GetPinterestVerificationCode

`func (o *SeoOptions) GetPinterestVerificationCode() string`

GetPinterestVerificationCode returns the PinterestVerificationCode field if non-nil, zero value otherwise.

### GetPinterestVerificationCodeOk

`func (o *SeoOptions) GetPinterestVerificationCodeOk() (*string, bool)`

GetPinterestVerificationCodeOk returns a tuple with the PinterestVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinterestVerificationCode

`func (o *SeoOptions) SetPinterestVerificationCode(v string)`

SetPinterestVerificationCode sets PinterestVerificationCode field to given value.

### HasPinterestVerificationCode

`func (o *SeoOptions) HasPinterestVerificationCode() bool`

HasPinterestVerificationCode returns a boolean if a field has been set.

### SetPinterestVerificationCodeNil

`func (o *SeoOptions) SetPinterestVerificationCodeNil(b bool)`

 SetPinterestVerificationCodeNil sets the value for PinterestVerificationCode to be an explicit nil

### UnsetPinterestVerificationCode
`func (o *SeoOptions) UnsetPinterestVerificationCode()`

UnsetPinterestVerificationCode ensures that no value is present for PinterestVerificationCode, not even an explicit nil
### GetCreator

`func (o *SeoOptions) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SeoOptions) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SeoOptions) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SeoOptions) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOrganization

`func (o *SeoOptions) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SeoOptions) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SeoOptions) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SeoOptions) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSameAs

`func (o *SeoOptions) GetSameAs() []string`

GetSameAs returns the SameAs field if non-nil, zero value otherwise.

### GetSameAsOk

`func (o *SeoOptions) GetSameAsOk() (*[]string, bool)`

GetSameAsOk returns a tuple with the SameAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameAs

`func (o *SeoOptions) SetSameAs(v []string)`

SetSameAs sets SameAs field to given value.

### HasSameAs

`func (o *SeoOptions) HasSameAs() bool`

HasSameAs returns a boolean if a field has been set.

### SetSameAsNil

`func (o *SeoOptions) SetSameAsNil(b bool)`

 SetSameAsNil sets the value for SameAs to be an explicit nil

### UnsetSameAs
`func (o *SeoOptions) UnsetSameAs()`

UnsetSameAs ensures that no value is present for SameAs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


