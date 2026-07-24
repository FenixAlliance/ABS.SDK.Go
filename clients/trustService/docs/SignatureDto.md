# SignatureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ValidationCode** | Pointer to **NullableString** |  | [optional] 
**SignatureImage** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateId** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentId** | Pointer to **NullableString** |  | [optional] 
**SignedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**SigningStatus** | Pointer to **NullableString** |  | [optional] 
**VerificationStatus** | Pointer to **NullableString** |  | [optional] 
**SignatureFormat** | Pointer to **NullableString** |  | [optional] 
**DigestAlgorithm** | Pointer to **NullableString** |  | [optional] 
**SignatureAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CanonicalizationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**PolicyIdentifier** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**DigestValue** | Pointer to **NullableString** |  | [optional] 
**SignatureValueHash** | Pointer to **NullableString** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**SigningProfileDisplayName** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateTitle** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentTitle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignatureDto

`func NewSignatureDto() *SignatureDto`

NewSignatureDto instantiates a new SignatureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureDtoWithDefaults

`func NewSignatureDtoWithDefaults() *SignatureDto`

NewSignatureDtoWithDefaults instantiates a new SignatureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignatureDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignatureDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignatureDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignatureDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SignatureDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SignatureDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SignatureDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignatureDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignatureDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignatureDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SignatureDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SignatureDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetType

`func (o *SignatureDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignatureDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignatureDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignatureDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SignatureDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SignatureDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValidationCode

`func (o *SignatureDto) GetValidationCode() string`

GetValidationCode returns the ValidationCode field if non-nil, zero value otherwise.

### GetValidationCodeOk

`func (o *SignatureDto) GetValidationCodeOk() (*string, bool)`

GetValidationCodeOk returns a tuple with the ValidationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationCode

`func (o *SignatureDto) SetValidationCode(v string)`

SetValidationCode sets ValidationCode field to given value.

### HasValidationCode

`func (o *SignatureDto) HasValidationCode() bool`

HasValidationCode returns a boolean if a field has been set.

### SetValidationCodeNil

`func (o *SignatureDto) SetValidationCodeNil(b bool)`

 SetValidationCodeNil sets the value for ValidationCode to be an explicit nil

### UnsetValidationCode
`func (o *SignatureDto) UnsetValidationCode()`

UnsetValidationCode ensures that no value is present for ValidationCode, not even an explicit nil
### GetSignatureImage

`func (o *SignatureDto) GetSignatureImage() string`

GetSignatureImage returns the SignatureImage field if non-nil, zero value otherwise.

### GetSignatureImageOk

`func (o *SignatureDto) GetSignatureImageOk() (*string, bool)`

GetSignatureImageOk returns a tuple with the SignatureImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureImage

`func (o *SignatureDto) SetSignatureImage(v string)`

SetSignatureImage sets SignatureImage field to given value.

### HasSignatureImage

`func (o *SignatureDto) HasSignatureImage() bool`

HasSignatureImage returns a boolean if a field has been set.

### SetSignatureImageNil

`func (o *SignatureDto) SetSignatureImageNil(b bool)`

 SetSignatureImageNil sets the value for SignatureImage to be an explicit nil

### UnsetSignatureImage
`func (o *SignatureDto) UnsetSignatureImage()`

UnsetSignatureImage ensures that no value is present for SignatureImage, not even an explicit nil
### GetContactId

`func (o *SignatureDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SignatureDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SignatureDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SignatureDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SignatureDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SignatureDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *SignatureDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SignatureDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SignatureDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SignatureDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SignatureDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SignatureDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *SignatureDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignatureDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SignatureDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SignatureDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SignatureDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SignatureDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEnrollmentId

`func (o *SignatureDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SignatureDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SignatureDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SignatureDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SignatureDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SignatureDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSigningProfileId

`func (o *SignatureDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *SignatureDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *SignatureDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *SignatureDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *SignatureDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *SignatureDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil
### GetSigningCertificateId

`func (o *SignatureDto) GetSigningCertificateId() string`

GetSigningCertificateId returns the SigningCertificateId field if non-nil, zero value otherwise.

### GetSigningCertificateIdOk

`func (o *SignatureDto) GetSigningCertificateIdOk() (*string, bool)`

GetSigningCertificateIdOk returns a tuple with the SigningCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateId

`func (o *SignatureDto) SetSigningCertificateId(v string)`

SetSigningCertificateId sets SigningCertificateId field to given value.

### HasSigningCertificateId

`func (o *SignatureDto) HasSigningCertificateId() bool`

HasSigningCertificateId returns a boolean if a field has been set.

### SetSigningCertificateIdNil

`func (o *SignatureDto) SetSigningCertificateIdNil(b bool)`

 SetSigningCertificateIdNil sets the value for SigningCertificateId to be an explicit nil

### UnsetSigningCertificateId
`func (o *SignatureDto) UnsetSigningCertificateId()`

UnsetSigningCertificateId ensures that no value is present for SigningCertificateId, not even an explicit nil
### GetSignedDocumentId

`func (o *SignatureDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *SignatureDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *SignatureDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.

### HasSignedDocumentId

`func (o *SignatureDto) HasSignedDocumentId() bool`

HasSignedDocumentId returns a boolean if a field has been set.

### SetSignedDocumentIdNil

`func (o *SignatureDto) SetSignedDocumentIdNil(b bool)`

 SetSignedDocumentIdNil sets the value for SignedDocumentId to be an explicit nil

### UnsetSignedDocumentId
`func (o *SignatureDto) UnsetSignedDocumentId()`

UnsetSignedDocumentId ensures that no value is present for SignedDocumentId, not even an explicit nil
### GetSignedAtUtc

`func (o *SignatureDto) GetSignedAtUtc() time.Time`

GetSignedAtUtc returns the SignedAtUtc field if non-nil, zero value otherwise.

### GetSignedAtUtcOk

`func (o *SignatureDto) GetSignedAtUtcOk() (*time.Time, bool)`

GetSignedAtUtcOk returns a tuple with the SignedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAtUtc

`func (o *SignatureDto) SetSignedAtUtc(v time.Time)`

SetSignedAtUtc sets SignedAtUtc field to given value.

### HasSignedAtUtc

`func (o *SignatureDto) HasSignedAtUtc() bool`

HasSignedAtUtc returns a boolean if a field has been set.

### SetSignedAtUtcNil

`func (o *SignatureDto) SetSignedAtUtcNil(b bool)`

 SetSignedAtUtcNil sets the value for SignedAtUtc to be an explicit nil

### UnsetSignedAtUtc
`func (o *SignatureDto) UnsetSignedAtUtc()`

UnsetSignedAtUtc ensures that no value is present for SignedAtUtc, not even an explicit nil
### GetSigningStatus

`func (o *SignatureDto) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *SignatureDto) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *SignatureDto) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *SignatureDto) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.

### SetSigningStatusNil

`func (o *SignatureDto) SetSigningStatusNil(b bool)`

 SetSigningStatusNil sets the value for SigningStatus to be an explicit nil

### UnsetSigningStatus
`func (o *SignatureDto) UnsetSigningStatus()`

UnsetSigningStatus ensures that no value is present for SigningStatus, not even an explicit nil
### GetVerificationStatus

`func (o *SignatureDto) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *SignatureDto) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *SignatureDto) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *SignatureDto) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### SetVerificationStatusNil

`func (o *SignatureDto) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *SignatureDto) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
### GetSignatureFormat

`func (o *SignatureDto) GetSignatureFormat() string`

GetSignatureFormat returns the SignatureFormat field if non-nil, zero value otherwise.

### GetSignatureFormatOk

`func (o *SignatureDto) GetSignatureFormatOk() (*string, bool)`

GetSignatureFormatOk returns a tuple with the SignatureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFormat

`func (o *SignatureDto) SetSignatureFormat(v string)`

SetSignatureFormat sets SignatureFormat field to given value.

### HasSignatureFormat

`func (o *SignatureDto) HasSignatureFormat() bool`

HasSignatureFormat returns a boolean if a field has been set.

### SetSignatureFormatNil

`func (o *SignatureDto) SetSignatureFormatNil(b bool)`

 SetSignatureFormatNil sets the value for SignatureFormat to be an explicit nil

### UnsetSignatureFormat
`func (o *SignatureDto) UnsetSignatureFormat()`

UnsetSignatureFormat ensures that no value is present for SignatureFormat, not even an explicit nil
### GetDigestAlgorithm

`func (o *SignatureDto) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SignatureDto) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SignatureDto) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SignatureDto) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### SetDigestAlgorithmNil

`func (o *SignatureDto) SetDigestAlgorithmNil(b bool)`

 SetDigestAlgorithmNil sets the value for DigestAlgorithm to be an explicit nil

### UnsetDigestAlgorithm
`func (o *SignatureDto) UnsetDigestAlgorithm()`

UnsetDigestAlgorithm ensures that no value is present for DigestAlgorithm, not even an explicit nil
### GetSignatureAlgorithm

`func (o *SignatureDto) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SignatureDto) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SignatureDto) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SignatureDto) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### SetSignatureAlgorithmNil

`func (o *SignatureDto) SetSignatureAlgorithmNil(b bool)`

 SetSignatureAlgorithmNil sets the value for SignatureAlgorithm to be an explicit nil

### UnsetSignatureAlgorithm
`func (o *SignatureDto) UnsetSignatureAlgorithm()`

UnsetSignatureAlgorithm ensures that no value is present for SignatureAlgorithm, not even an explicit nil
### GetCanonicalizationAlgorithm

`func (o *SignatureDto) GetCanonicalizationAlgorithm() string`

GetCanonicalizationAlgorithm returns the CanonicalizationAlgorithm field if non-nil, zero value otherwise.

### GetCanonicalizationAlgorithmOk

`func (o *SignatureDto) GetCanonicalizationAlgorithmOk() (*string, bool)`

GetCanonicalizationAlgorithmOk returns a tuple with the CanonicalizationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizationAlgorithm

`func (o *SignatureDto) SetCanonicalizationAlgorithm(v string)`

SetCanonicalizationAlgorithm sets CanonicalizationAlgorithm field to given value.

### HasCanonicalizationAlgorithm

`func (o *SignatureDto) HasCanonicalizationAlgorithm() bool`

HasCanonicalizationAlgorithm returns a boolean if a field has been set.

### SetCanonicalizationAlgorithmNil

`func (o *SignatureDto) SetCanonicalizationAlgorithmNil(b bool)`

 SetCanonicalizationAlgorithmNil sets the value for CanonicalizationAlgorithm to be an explicit nil

### UnsetCanonicalizationAlgorithm
`func (o *SignatureDto) UnsetCanonicalizationAlgorithm()`

UnsetCanonicalizationAlgorithm ensures that no value is present for CanonicalizationAlgorithm, not even an explicit nil
### GetPolicyIdentifier

`func (o *SignatureDto) GetPolicyIdentifier() string`

GetPolicyIdentifier returns the PolicyIdentifier field if non-nil, zero value otherwise.

### GetPolicyIdentifierOk

`func (o *SignatureDto) GetPolicyIdentifierOk() (*string, bool)`

GetPolicyIdentifierOk returns a tuple with the PolicyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifier

`func (o *SignatureDto) SetPolicyIdentifier(v string)`

SetPolicyIdentifier sets PolicyIdentifier field to given value.

### HasPolicyIdentifier

`func (o *SignatureDto) HasPolicyIdentifier() bool`

HasPolicyIdentifier returns a boolean if a field has been set.

### SetPolicyIdentifierNil

`func (o *SignatureDto) SetPolicyIdentifierNil(b bool)`

 SetPolicyIdentifierNil sets the value for PolicyIdentifier to be an explicit nil

### UnsetPolicyIdentifier
`func (o *SignatureDto) UnsetPolicyIdentifier()`

UnsetPolicyIdentifier ensures that no value is present for PolicyIdentifier, not even an explicit nil
### GetCorrelationId

`func (o *SignatureDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignatureDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SignatureDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SignatureDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SignatureDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SignatureDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetDigestValue

`func (o *SignatureDto) GetDigestValue() string`

GetDigestValue returns the DigestValue field if non-nil, zero value otherwise.

### GetDigestValueOk

`func (o *SignatureDto) GetDigestValueOk() (*string, bool)`

GetDigestValueOk returns a tuple with the DigestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestValue

`func (o *SignatureDto) SetDigestValue(v string)`

SetDigestValue sets DigestValue field to given value.

### HasDigestValue

`func (o *SignatureDto) HasDigestValue() bool`

HasDigestValue returns a boolean if a field has been set.

### SetDigestValueNil

`func (o *SignatureDto) SetDigestValueNil(b bool)`

 SetDigestValueNil sets the value for DigestValue to be an explicit nil

### UnsetDigestValue
`func (o *SignatureDto) UnsetDigestValue()`

UnsetDigestValue ensures that no value is present for DigestValue, not even an explicit nil
### GetSignatureValueHash

`func (o *SignatureDto) GetSignatureValueHash() string`

GetSignatureValueHash returns the SignatureValueHash field if non-nil, zero value otherwise.

### GetSignatureValueHashOk

`func (o *SignatureDto) GetSignatureValueHashOk() (*string, bool)`

GetSignatureValueHashOk returns a tuple with the SignatureValueHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureValueHash

`func (o *SignatureDto) SetSignatureValueHash(v string)`

SetSignatureValueHash sets SignatureValueHash field to given value.

### HasSignatureValueHash

`func (o *SignatureDto) HasSignatureValueHash() bool`

HasSignatureValueHash returns a boolean if a field has been set.

### SetSignatureValueHashNil

`func (o *SignatureDto) SetSignatureValueHashNil(b bool)`

 SetSignatureValueHashNil sets the value for SignatureValueHash to be an explicit nil

### UnsetSignatureValueHash
`func (o *SignatureDto) UnsetSignatureValueHash()`

UnsetSignatureValueHash ensures that no value is present for SignatureValueHash, not even an explicit nil
### GetContactName

`func (o *SignatureDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SignatureDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SignatureDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SignatureDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *SignatureDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *SignatureDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetSigningProfileDisplayName

`func (o *SignatureDto) GetSigningProfileDisplayName() string`

GetSigningProfileDisplayName returns the SigningProfileDisplayName field if non-nil, zero value otherwise.

### GetSigningProfileDisplayNameOk

`func (o *SignatureDto) GetSigningProfileDisplayNameOk() (*string, bool)`

GetSigningProfileDisplayNameOk returns a tuple with the SigningProfileDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileDisplayName

`func (o *SignatureDto) SetSigningProfileDisplayName(v string)`

SetSigningProfileDisplayName sets SigningProfileDisplayName field to given value.

### HasSigningProfileDisplayName

`func (o *SignatureDto) HasSigningProfileDisplayName() bool`

HasSigningProfileDisplayName returns a boolean if a field has been set.

### SetSigningProfileDisplayNameNil

`func (o *SignatureDto) SetSigningProfileDisplayNameNil(b bool)`

 SetSigningProfileDisplayNameNil sets the value for SigningProfileDisplayName to be an explicit nil

### UnsetSigningProfileDisplayName
`func (o *SignatureDto) UnsetSigningProfileDisplayName()`

UnsetSigningProfileDisplayName ensures that no value is present for SigningProfileDisplayName, not even an explicit nil
### GetSigningCertificateTitle

`func (o *SignatureDto) GetSigningCertificateTitle() string`

GetSigningCertificateTitle returns the SigningCertificateTitle field if non-nil, zero value otherwise.

### GetSigningCertificateTitleOk

`func (o *SignatureDto) GetSigningCertificateTitleOk() (*string, bool)`

GetSigningCertificateTitleOk returns a tuple with the SigningCertificateTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateTitle

`func (o *SignatureDto) SetSigningCertificateTitle(v string)`

SetSigningCertificateTitle sets SigningCertificateTitle field to given value.

### HasSigningCertificateTitle

`func (o *SignatureDto) HasSigningCertificateTitle() bool`

HasSigningCertificateTitle returns a boolean if a field has been set.

### SetSigningCertificateTitleNil

`func (o *SignatureDto) SetSigningCertificateTitleNil(b bool)`

 SetSigningCertificateTitleNil sets the value for SigningCertificateTitle to be an explicit nil

### UnsetSigningCertificateTitle
`func (o *SignatureDto) UnsetSigningCertificateTitle()`

UnsetSigningCertificateTitle ensures that no value is present for SigningCertificateTitle, not even an explicit nil
### GetSignedDocumentTitle

`func (o *SignatureDto) GetSignedDocumentTitle() string`

GetSignedDocumentTitle returns the SignedDocumentTitle field if non-nil, zero value otherwise.

### GetSignedDocumentTitleOk

`func (o *SignatureDto) GetSignedDocumentTitleOk() (*string, bool)`

GetSignedDocumentTitleOk returns a tuple with the SignedDocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentTitle

`func (o *SignatureDto) SetSignedDocumentTitle(v string)`

SetSignedDocumentTitle sets SignedDocumentTitle field to given value.

### HasSignedDocumentTitle

`func (o *SignatureDto) HasSignedDocumentTitle() bool`

HasSignedDocumentTitle returns a boolean if a field has been set.

### SetSignedDocumentTitleNil

`func (o *SignatureDto) SetSignedDocumentTitleNil(b bool)`

 SetSignedDocumentTitleNil sets the value for SignedDocumentTitle to be an explicit nil

### UnsetSignedDocumentTitle
`func (o *SignatureDto) UnsetSignedDocumentTitle()`

UnsetSignedDocumentTitle ensures that no value is present for SignedDocumentTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


