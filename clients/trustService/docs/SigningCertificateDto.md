# SigningCertificateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**FileLengthInBits** | Pointer to **int64** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SecurityCertificateId** | Pointer to **NullableString** |  | [optional] 
**CertificateStatus** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**SubjectName** | Pointer to **NullableString** |  | [optional] 
**IssuerName** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**NotBeforeUtc** | Pointer to **NullableTime** |  | [optional] 
**NotAfterUtc** | Pointer to **NullableTime** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningCertificateDto

`func NewSigningCertificateDto() *SigningCertificateDto`

NewSigningCertificateDto instantiates a new SigningCertificateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningCertificateDtoWithDefaults

`func NewSigningCertificateDtoWithDefaults() *SigningCertificateDto`

NewSigningCertificateDtoWithDefaults instantiates a new SigningCertificateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningCertificateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningCertificateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningCertificateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningCertificateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningCertificateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningCertificateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SigningCertificateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningCertificateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningCertificateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningCertificateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SigningCertificateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SigningCertificateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SigningCertificateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SigningCertificateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SigningCertificateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SigningCertificateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SigningCertificateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SigningCertificateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *SigningCertificateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningCertificateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningCertificateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningCertificateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningCertificateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningCertificateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *SigningCertificateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SigningCertificateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SigningCertificateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SigningCertificateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SigningCertificateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SigningCertificateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCsr

`func (o *SigningCertificateDto) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *SigningCertificateDto) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *SigningCertificateDto) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *SigningCertificateDto) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *SigningCertificateDto) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *SigningCertificateDto) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetPublicKey

`func (o *SigningCertificateDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SigningCertificateDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SigningCertificateDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SigningCertificateDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *SigningCertificateDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *SigningCertificateDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetSignature

`func (o *SigningCertificateDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SigningCertificateDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SigningCertificateDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SigningCertificateDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SigningCertificateDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SigningCertificateDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetFileLengthInBits

`func (o *SigningCertificateDto) GetFileLengthInBits() int64`

GetFileLengthInBits returns the FileLengthInBits field if non-nil, zero value otherwise.

### GetFileLengthInBitsOk

`func (o *SigningCertificateDto) GetFileLengthInBitsOk() (*int64, bool)`

GetFileLengthInBitsOk returns a tuple with the FileLengthInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLengthInBits

`func (o *SigningCertificateDto) SetFileLengthInBits(v int64)`

SetFileLengthInBits sets FileLengthInBits field to given value.

### HasFileLengthInBits

`func (o *SigningCertificateDto) HasFileLengthInBits() bool`

HasFileLengthInBits returns a boolean if a field has been set.

### GetCertificateType

`func (o *SigningCertificateDto) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *SigningCertificateDto) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *SigningCertificateDto) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *SigningCertificateDto) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetContactId

`func (o *SigningCertificateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningCertificateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningCertificateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SigningCertificateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SigningCertificateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SigningCertificateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *SigningCertificateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningCertificateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningCertificateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningCertificateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningCertificateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningCertificateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *SigningCertificateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SigningCertificateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SigningCertificateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SigningCertificateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SigningCertificateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SigningCertificateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEnrollmentId

`func (o *SigningCertificateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SigningCertificateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SigningCertificateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SigningCertificateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SigningCertificateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SigningCertificateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSecurityCertificateId

`func (o *SigningCertificateDto) GetSecurityCertificateId() string`

GetSecurityCertificateId returns the SecurityCertificateId field if non-nil, zero value otherwise.

### GetSecurityCertificateIdOk

`func (o *SigningCertificateDto) GetSecurityCertificateIdOk() (*string, bool)`

GetSecurityCertificateIdOk returns a tuple with the SecurityCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateId

`func (o *SigningCertificateDto) SetSecurityCertificateId(v string)`

SetSecurityCertificateId sets SecurityCertificateId field to given value.

### HasSecurityCertificateId

`func (o *SigningCertificateDto) HasSecurityCertificateId() bool`

HasSecurityCertificateId returns a boolean if a field has been set.

### SetSecurityCertificateIdNil

`func (o *SigningCertificateDto) SetSecurityCertificateIdNil(b bool)`

 SetSecurityCertificateIdNil sets the value for SecurityCertificateId to be an explicit nil

### UnsetSecurityCertificateId
`func (o *SigningCertificateDto) UnsetSecurityCertificateId()`

UnsetSecurityCertificateId ensures that no value is present for SecurityCertificateId, not even an explicit nil
### GetCertificateStatus

`func (o *SigningCertificateDto) GetCertificateStatus() string`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *SigningCertificateDto) GetCertificateStatusOk() (*string, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *SigningCertificateDto) SetCertificateStatus(v string)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *SigningCertificateDto) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### SetCertificateStatusNil

`func (o *SigningCertificateDto) SetCertificateStatusNil(b bool)`

 SetCertificateStatusNil sets the value for CertificateStatus to be an explicit nil

### UnsetCertificateStatus
`func (o *SigningCertificateDto) UnsetCertificateStatus()`

UnsetCertificateStatus ensures that no value is present for CertificateStatus, not even an explicit nil
### GetThumbprint

`func (o *SigningCertificateDto) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *SigningCertificateDto) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *SigningCertificateDto) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *SigningCertificateDto) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *SigningCertificateDto) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *SigningCertificateDto) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetSubjectName

`func (o *SigningCertificateDto) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *SigningCertificateDto) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *SigningCertificateDto) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *SigningCertificateDto) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *SigningCertificateDto) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *SigningCertificateDto) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetIssuerName

`func (o *SigningCertificateDto) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *SigningCertificateDto) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *SigningCertificateDto) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *SigningCertificateDto) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *SigningCertificateDto) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *SigningCertificateDto) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetSerialNumber

`func (o *SigningCertificateDto) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SigningCertificateDto) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SigningCertificateDto) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SigningCertificateDto) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *SigningCertificateDto) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *SigningCertificateDto) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetNotBeforeUtc

`func (o *SigningCertificateDto) GetNotBeforeUtc() time.Time`

GetNotBeforeUtc returns the NotBeforeUtc field if non-nil, zero value otherwise.

### GetNotBeforeUtcOk

`func (o *SigningCertificateDto) GetNotBeforeUtcOk() (*time.Time, bool)`

GetNotBeforeUtcOk returns a tuple with the NotBeforeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeUtc

`func (o *SigningCertificateDto) SetNotBeforeUtc(v time.Time)`

SetNotBeforeUtc sets NotBeforeUtc field to given value.

### HasNotBeforeUtc

`func (o *SigningCertificateDto) HasNotBeforeUtc() bool`

HasNotBeforeUtc returns a boolean if a field has been set.

### SetNotBeforeUtcNil

`func (o *SigningCertificateDto) SetNotBeforeUtcNil(b bool)`

 SetNotBeforeUtcNil sets the value for NotBeforeUtc to be an explicit nil

### UnsetNotBeforeUtc
`func (o *SigningCertificateDto) UnsetNotBeforeUtc()`

UnsetNotBeforeUtc ensures that no value is present for NotBeforeUtc, not even an explicit nil
### GetNotAfterUtc

`func (o *SigningCertificateDto) GetNotAfterUtc() time.Time`

GetNotAfterUtc returns the NotAfterUtc field if non-nil, zero value otherwise.

### GetNotAfterUtcOk

`func (o *SigningCertificateDto) GetNotAfterUtcOk() (*time.Time, bool)`

GetNotAfterUtcOk returns a tuple with the NotAfterUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfterUtc

`func (o *SigningCertificateDto) SetNotAfterUtc(v time.Time)`

SetNotAfterUtc sets NotAfterUtc field to given value.

### HasNotAfterUtc

`func (o *SigningCertificateDto) HasNotAfterUtc() bool`

HasNotAfterUtc returns a boolean if a field has been set.

### SetNotAfterUtcNil

`func (o *SigningCertificateDto) SetNotAfterUtcNil(b bool)`

 SetNotAfterUtcNil sets the value for NotAfterUtc to be an explicit nil

### UnsetNotAfterUtc
`func (o *SigningCertificateDto) UnsetNotAfterUtc()`

UnsetNotAfterUtc ensures that no value is present for NotAfterUtc, not even an explicit nil
### GetContactName

`func (o *SigningCertificateDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SigningCertificateDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SigningCertificateDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SigningCertificateDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *SigningCertificateDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *SigningCertificateDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


