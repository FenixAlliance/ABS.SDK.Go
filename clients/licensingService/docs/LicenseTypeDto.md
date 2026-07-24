# LicenseTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AllowTrials** | Pointer to **bool** |  | [optional] 
**IsPerpetualLicense** | Pointer to **bool** |  | [optional] 
**MaxLicenseUsages** | Pointer to **int32** |  | [optional] 
**TrialLicenseRelativeExpirationInDays** | Pointer to **int32** |  | [optional] 
**StandardLicenseRelativeExpirationInDays** | Pointer to **int32** |  | [optional] 
**LicensingCertificateId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseTypeDto

`func NewLicenseTypeDto() *LicenseTypeDto`

NewLicenseTypeDto instantiates a new LicenseTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseTypeDtoWithDefaults

`func NewLicenseTypeDtoWithDefaults() *LicenseTypeDto`

NewLicenseTypeDtoWithDefaults instantiates a new LicenseTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseTypeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseTypeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseTypeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseTypeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LicenseTypeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LicenseTypeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LicenseTypeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseTypeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseTypeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseTypeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LicenseTypeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LicenseTypeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *LicenseTypeDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicenseTypeDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicenseTypeDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LicenseTypeDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LicenseTypeDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LicenseTypeDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetName

`func (o *LicenseTypeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseTypeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseTypeDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseTypeDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LicenseTypeDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LicenseTypeDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LicenseTypeDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseTypeDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseTypeDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseTypeDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseTypeDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseTypeDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowTrials

`func (o *LicenseTypeDto) GetAllowTrials() bool`

GetAllowTrials returns the AllowTrials field if non-nil, zero value otherwise.

### GetAllowTrialsOk

`func (o *LicenseTypeDto) GetAllowTrialsOk() (*bool, bool)`

GetAllowTrialsOk returns a tuple with the AllowTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrials

`func (o *LicenseTypeDto) SetAllowTrials(v bool)`

SetAllowTrials sets AllowTrials field to given value.

### HasAllowTrials

`func (o *LicenseTypeDto) HasAllowTrials() bool`

HasAllowTrials returns a boolean if a field has been set.

### GetIsPerpetualLicense

`func (o *LicenseTypeDto) GetIsPerpetualLicense() bool`

GetIsPerpetualLicense returns the IsPerpetualLicense field if non-nil, zero value otherwise.

### GetIsPerpetualLicenseOk

`func (o *LicenseTypeDto) GetIsPerpetualLicenseOk() (*bool, bool)`

GetIsPerpetualLicenseOk returns a tuple with the IsPerpetualLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerpetualLicense

`func (o *LicenseTypeDto) SetIsPerpetualLicense(v bool)`

SetIsPerpetualLicense sets IsPerpetualLicense field to given value.

### HasIsPerpetualLicense

`func (o *LicenseTypeDto) HasIsPerpetualLicense() bool`

HasIsPerpetualLicense returns a boolean if a field has been set.

### GetMaxLicenseUsages

`func (o *LicenseTypeDto) GetMaxLicenseUsages() int32`

GetMaxLicenseUsages returns the MaxLicenseUsages field if non-nil, zero value otherwise.

### GetMaxLicenseUsagesOk

`func (o *LicenseTypeDto) GetMaxLicenseUsagesOk() (*int32, bool)`

GetMaxLicenseUsagesOk returns a tuple with the MaxLicenseUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLicenseUsages

`func (o *LicenseTypeDto) SetMaxLicenseUsages(v int32)`

SetMaxLicenseUsages sets MaxLicenseUsages field to given value.

### HasMaxLicenseUsages

`func (o *LicenseTypeDto) HasMaxLicenseUsages() bool`

HasMaxLicenseUsages returns a boolean if a field has been set.

### GetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) GetTrialLicenseRelativeExpirationInDays() int32`

GetTrialLicenseRelativeExpirationInDays returns the TrialLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetTrialLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeDto) GetTrialLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetTrialLicenseRelativeExpirationInDaysOk returns a tuple with the TrialLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) SetTrialLicenseRelativeExpirationInDays(v int32)`

SetTrialLicenseRelativeExpirationInDays sets TrialLicenseRelativeExpirationInDays field to given value.

### HasTrialLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) HasTrialLicenseRelativeExpirationInDays() bool`

HasTrialLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### GetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) GetStandardLicenseRelativeExpirationInDays() int32`

GetStandardLicenseRelativeExpirationInDays returns the StandardLicenseRelativeExpirationInDays field if non-nil, zero value otherwise.

### GetStandardLicenseRelativeExpirationInDaysOk

`func (o *LicenseTypeDto) GetStandardLicenseRelativeExpirationInDaysOk() (*int32, bool)`

GetStandardLicenseRelativeExpirationInDaysOk returns a tuple with the StandardLicenseRelativeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) SetStandardLicenseRelativeExpirationInDays(v int32)`

SetStandardLicenseRelativeExpirationInDays sets StandardLicenseRelativeExpirationInDays field to given value.

### HasStandardLicenseRelativeExpirationInDays

`func (o *LicenseTypeDto) HasStandardLicenseRelativeExpirationInDays() bool`

HasStandardLicenseRelativeExpirationInDays returns a boolean if a field has been set.

### GetLicensingCertificateId

`func (o *LicenseTypeDto) GetLicensingCertificateId() string`

GetLicensingCertificateId returns the LicensingCertificateId field if non-nil, zero value otherwise.

### GetLicensingCertificateIdOk

`func (o *LicenseTypeDto) GetLicensingCertificateIdOk() (*string, bool)`

GetLicensingCertificateIdOk returns a tuple with the LicensingCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingCertificateId

`func (o *LicenseTypeDto) SetLicensingCertificateId(v string)`

SetLicensingCertificateId sets LicensingCertificateId field to given value.

### HasLicensingCertificateId

`func (o *LicenseTypeDto) HasLicensingCertificateId() bool`

HasLicensingCertificateId returns a boolean if a field has been set.

### SetLicensingCertificateIdNil

`func (o *LicenseTypeDto) SetLicensingCertificateIdNil(b bool)`

 SetLicensingCertificateIdNil sets the value for LicensingCertificateId to be an explicit nil

### UnsetLicensingCertificateId
`func (o *LicenseTypeDto) UnsetLicensingCertificateId()`

UnsetLicensingCertificateId ensures that no value is present for LicensingCertificateId, not even an explicit nil
### GetTenantId

`func (o *LicenseTypeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LicenseTypeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LicenseTypeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LicenseTypeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LicenseTypeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LicenseTypeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LicenseTypeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LicenseTypeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LicenseTypeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LicenseTypeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LicenseTypeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LicenseTypeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


