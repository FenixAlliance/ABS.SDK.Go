# LicenseTypeUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AllowTrials** | Pointer to **NullableBool** |  | [optional] 
**IsPerpetualLicense** | Pointer to **NullableBool** |  | [optional] 
**MaxLicenseUsages** | Pointer to **NullableInt32** |  | [optional] 
**TrialLicenseRelativeExpirationInDays** | Pointer to **NullableInt32** |  | [optional] 
**StandardLicenseRelativeExpirationInDays** | Pointer to **NullableInt32** |  | [optional] 
**LicensingCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseTypeUpdateDto

`func NewLicenseTypeUpdateDto() *LicenseTypeUpdateDto`

NewLicenseTypeUpdateDto instantiates a new LicenseTypeUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseTypeUpdateDtoWithDefaults

`func NewLicenseTypeUpdateDtoWithDefaults() *LicenseTypeUpdateDto`

NewLicenseTypeUpdateDtoWithDefaults instantiates a new LicenseTypeUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LicenseTypeUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicenseTypeUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicenseTypeUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LicenseTypeUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LicenseTypeUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LicenseTypeUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *LicenseTypeUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseTypeUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseTypeUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseTypeUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseTypeUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseTypeUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowTrials

`func (o *LicenseTypeUpdateDto) GetAllowTrials() bool`

GetAllowTrials returns the AllowTrials field if non-nil, zero value otherwise.

### GetAllowTrialsOk

`func (o *LicenseTypeUpdateDto) GetAllowTrialsOk() (*bool, bool)`

GetAllowTrialsOk returns a tuple with the AllowTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrials

`func (o *LicenseTypeUpdateDto) SetAllowTrials(v bool)`

SetAllowTrials sets AllowTrials field to given value.

### HasAllowTrials

`func (o *LicenseTypeUpdateDto) HasAllowTrials() bool`

HasAllowTrials returns a boolean if a field has been set.

### SetAllowTrialsNil

`func (o *LicenseTypeUpdateDto) SetAllowTrialsNil(b bool)`

 SetAllowTrialsNil sets the value for AllowTrials to be an explicit nil

### UnsetAllowTrials
`func (o *LicenseTypeUpdateDto) UnsetAllowTrials()`

UnsetAllowTrials ensures that no value is present for AllowTrials, not even an explicit nil
### GetIsPerpetualLicense

`func (o *LicenseTypeUpdateDto) GetIsPerpetualLicense() bool`

GetIsPerpetualLicense returns the IsPerpetualLicense field if non-nil, zero value otherwise.

### GetIsPerpetualLicenseOk

`func (o *LicenseTypeUpdateDto) GetIsPerpetualLicenseOk() (*bool, bool)`

GetIsPerpetualLicenseOk returns a tuple with the IsPerpetualLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerpetualLicense

`func (o *LicenseTypeUpdateDto) SetIsPerpetualLicense(v bool)`

SetIsPerpetualLicense sets IsPerpetualLicense field to given value.

### HasIsPerpetualLicense

`func (o *LicenseTypeUpdateDto) HasIsPerpetualLicense() bool`

HasIsPerpetualLicense returns a boolean if a field has been set.

### SetIsPerpetualLicenseNil

`func (o *LicenseTypeUpdateDto) SetIsPerpetualLicenseNil(b bool)`

 SetIsPerpetualLicenseNil sets the value for IsPerpetualLicense to be an explicit nil

### UnsetIsPerpetualLicense
`func (o *LicenseTypeUpdateDto) UnsetIsPerpetualLicense()`

UnsetIsPerpetualLicense ensures that no value is present for IsPerpetualLicense, not even an explicit nil
### GetMaxLicenseUsages

`func (o *LicenseTypeUpdateDto) GetMaxLicenseUsages() int32`

GetMaxLicenseUsages returns the MaxLicenseUsages field if non-nil, zero value otherwise.

### GetMaxLicenseUsagesOk

`func (o *LicenseTypeUpdateDto) GetMaxLicenseUsagesOk() (*int32, bool)`

GetMaxLicenseUsagesOk returns a tuple with the MaxLicenseUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLicenseUsages

`func (o *LicenseTypeUpdateDto) SetMaxLicenseUsages(v int32)`

SetMaxLicenseUsages sets MaxLicenseUsages field to given value.

### HasMaxLicenseUsages

`func (o *LicenseTypeUpdateDto) HasMaxLicenseUsages() bool`

HasMaxLicenseUsages returns a boolean if a field has been set.

### SetMaxLicenseUsagesNil

`func (o *LicenseTypeUpdateDto) SetMaxLicenseUsagesNil(b bool)`

 SetMaxLicenseUsagesNil sets the value for MaxLicenseUsages to be an explicit nil

### UnsetMaxLicenseUsages
`func (o *LicenseTypeUpdateDto) UnsetMaxLicenseUsages()`

UnsetMaxLicenseUsages ensures that no value is present for MaxLicenseUsages, not even an explicit nil
### GetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) GetTrialLicenseRelativeExpirationInDays() int32`

GetTrialLicenseRelativeExpirationInDays returns the TrialLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetTrialLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeUpdateDto) GetTrialLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetTrialLicenseRelativeExpirationInDaysOk returns a tuple with the TrialLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) SetTrialLicenseRelativeExpirationInDays(v int32)`

SetTrialLicenseRelativeExpirationInDays sets TrialLicenseRelativeExpirationInDays field to given value.

### HasTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) HasTrialLicenseRelativeExpirationInDays() bool`

HasTrialLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### SetTrialLicenseRelativeExpirationInDaysNil

`func (o *LicenseTypeUpdateDto) SetTrialLicenseRelativeExpirationInDaysNil(b bool)`

 SetTrialLicenseRelativeExpirationInDaysNil sets the value for TrialLicenseRelativeExpirationInDays to be an explicit nil

### UnsetTrialLicenseRelativeExpirationInDays
`func (o *LicenseTypeUpdateDto) UnsetTrialLicenseRelativeExpirationInDays()`

UnsetTrialLicenseRelativeExpirationInDays ensures that no value is present for TrialLicenseRelativeExpirationInDays, not even an explicit nil
### GetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) GetStandardLicenseRelativeExpirationInDays() int32`

GetStandardLicenseRelativeExpirationInDays returns the StandardLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetStandardLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeUpdateDto) GetStandardLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetStandardLicenseRelativeExpirationInDaysOk returns a tuple with the StandardLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) SetStandardLicenseRelativeExpirationInDays(v int32)`

SetStandardLicenseRelativeExpirationInDays sets StandardLicenseRelativeExpirationInDays field to given value.

### HasStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeUpdateDto) HasStandardLicenseRelativeExpirationInDays() bool`

HasStandardLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### SetStandardLicenseRelativeExpirationInDaysNil

`func (o *LicenseTypeUpdateDto) SetStandardLicenseRelativeExpirationInDaysNil(b bool)`

 SetStandardLicenseRelativeExpirationInDaysNil sets the value for StandardLicenseRelativeExpirationInDays to be an explicit nil

### UnsetStandardLicenseRelativeExpirationInDays
`func (o *LicenseTypeUpdateDto) UnsetStandardLicenseRelativeExpirationInDays()`

UnsetStandardLicenseRelativeExpirationInDays ensures that no value is present for StandardLicenseRelativeExpirationInDays, not even an explicit nil
### GetLicensingCertificateId

`func (o *LicenseTypeUpdateDto) GetLicensingCertificateId() string`

GetLicensingCertificateId returns the LicensingCertificateId field if non-nil, zero value otherwise.

### GetLicensingCertificateIdOk

`func (o *LicenseTypeUpdateDto) GetLicensingCertificateIdOk() (*string, bool)`

GetLicensingCertificateIdOk returns a tuple with the LicensingCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingCertificateId

`func (o *LicenseTypeUpdateDto) SetLicensingCertificateId(v string)`

SetLicensingCertificateId sets LicensingCertificateId field to given value.

### HasLicensingCertificateId

`func (o *LicenseTypeUpdateDto) HasLicensingCertificateId() bool`

HasLicensingCertificateId returns a boolean if a field has been set.

### SetLicensingCertificateIdNil

`func (o *LicenseTypeUpdateDto) SetLicensingCertificateIdNil(b bool)`

 SetLicensingCertificateIdNil sets the value for LicensingCertificateId to be an explicit nil

### UnsetLicensingCertificateId
`func (o *LicenseTypeUpdateDto) UnsetLicensingCertificateId()`

UnsetLicensingCertificateId ensures that no value is present for LicensingCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


