# LicenseTypeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AllowTrials** | Pointer to **bool** |  | [optional] 
**IsPerpetualLicense** | Pointer to **bool** |  | [optional] 
**MaxLicenseUsages** | Pointer to **int32** |  | [optional] 
**TrialLicenseRelativeExpirationInDays** | Pointer to **int32** |  | [optional] 
**StandardLicenseRelativeExpirationInDays** | Pointer to **int32** |  | [optional] 
**LicensingCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseTypeCreateDto

`func NewLicenseTypeCreateDto(title string, ) *LicenseTypeCreateDto`

NewLicenseTypeCreateDto instantiates a new LicenseTypeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseTypeCreateDtoWithDefaults

`func NewLicenseTypeCreateDtoWithDefaults() *LicenseTypeCreateDto`

NewLicenseTypeCreateDtoWithDefaults instantiates a new LicenseTypeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseTypeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseTypeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseTypeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseTypeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LicenseTypeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseTypeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseTypeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseTypeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *LicenseTypeCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicenseTypeCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicenseTypeCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *LicenseTypeCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseTypeCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseTypeCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseTypeCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseTypeCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseTypeCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowTrials

`func (o *LicenseTypeCreateDto) GetAllowTrials() bool`

GetAllowTrials returns the AllowTrials field if non-nil, zero value otherwise.

### GetAllowTrialsOk

`func (o *LicenseTypeCreateDto) GetAllowTrialsOk() (*bool, bool)`

GetAllowTrialsOk returns a tuple with the AllowTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrials

`func (o *LicenseTypeCreateDto) SetAllowTrials(v bool)`

SetAllowTrials sets AllowTrials field to given value.

### HasAllowTrials

`func (o *LicenseTypeCreateDto) HasAllowTrials() bool`

HasAllowTrials returns a boolean if a field has been set.

### GetIsPerpetualLicense

`func (o *LicenseTypeCreateDto) GetIsPerpetualLicense() bool`

GetIsPerpetualLicense returns the IsPerpetualLicense field if non-nil, zero value otherwise.

### GetIsPerpetualLicenseOk

`func (o *LicenseTypeCreateDto) GetIsPerpetualLicenseOk() (*bool, bool)`

GetIsPerpetualLicenseOk returns a tuple with the IsPerpetualLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerpetualLicense

`func (o *LicenseTypeCreateDto) SetIsPerpetualLicense(v bool)`

SetIsPerpetualLicense sets IsPerpetualLicense field to given value.

### HasIsPerpetualLicense

`func (o *LicenseTypeCreateDto) HasIsPerpetualLicense() bool`

HasIsPerpetualLicense returns a boolean if a field has been set.

### GetMaxLicenseUsages

`func (o *LicenseTypeCreateDto) GetMaxLicenseUsages() int32`

GetMaxLicenseUsages returns the MaxLicenseUsages field if non-nil, zero value otherwise.

### GetMaxLicenseUsagesOk

`func (o *LicenseTypeCreateDto) GetMaxLicenseUsagesOk() (*int32, bool)`

GetMaxLicenseUsagesOk returns a tuple with the MaxLicenseUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLicenseUsages

`func (o *LicenseTypeCreateDto) SetMaxLicenseUsages(v int32)`

SetMaxLicenseUsages sets MaxLicenseUsages field to given value.

### HasMaxLicenseUsages

`func (o *LicenseTypeCreateDto) HasMaxLicenseUsages() bool`

HasMaxLicenseUsages returns a boolean if a field has been set.

### GetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) GetTrialLicenseRelativeExpirationInDays() int32`

GetTrialLicenseRelativeExpirationInDays returns the TrialLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetTrialLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeCreateDto) GetTrialLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetTrialLicenseRelativeExpirationInDaysOk returns a tuple with the TrialLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) SetTrialLicenseRelativeExpirationInDays(v int32)`

SetTrialLicenseRelativeExpirationInDays sets TrialLicenseRelativeExpirationInDays field to given value.

### HasTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) HasTrialLicenseRelativeExpirationInDays() bool`

HasTrialLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### GetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) GetStandardLicenseRelativeExpirationInDays() int32`

GetStandardLicenseRelativeExpirationInDays returns the StandardLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetStandardLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeCreateDto) GetStandardLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetStandardLicenseRelativeExpirationInDaysOk returns a tuple with the StandardLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) SetStandardLicenseRelativeExpirationInDays(v int32)`

SetStandardLicenseRelativeExpirationInDays sets StandardLicenseRelativeExpirationInDays field to given value.

### HasStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeCreateDto) HasStandardLicenseRelativeExpirationInDays() bool`

HasStandardLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### GetLicensingCertificateId

`func (o *LicenseTypeCreateDto) GetLicensingCertificateId() string`

GetLicensingCertificateId returns the LicensingCertificateId field if non-nil, zero value otherwise.

### GetLicensingCertificateIdOk

`func (o *LicenseTypeCreateDto) GetLicensingCertificateIdOk() (*string, bool)`

GetLicensingCertificateIdOk returns a tuple with the LicensingCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingCertificateId

`func (o *LicenseTypeCreateDto) SetLicensingCertificateId(v string)`

SetLicensingCertificateId sets LicensingCertificateId field to given value.

### HasLicensingCertificateId

`func (o *LicenseTypeCreateDto) HasLicensingCertificateId() bool`

HasLicensingCertificateId returns a boolean if a field has been set.

### SetLicensingCertificateIdNil

`func (o *LicenseTypeCreateDto) SetLicensingCertificateIdNil(b bool)`

 SetLicensingCertificateIdNil sets the value for LicensingCertificateId to be an explicit nil

### UnsetLicensingCertificateId
`func (o *LicenseTypeCreateDto) UnsetLicensingCertificateId()`

UnsetLicensingCertificateId ensures that no value is present for LicensingCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


