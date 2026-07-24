# SignedDocumentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Signed** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLengthInBits** | Pointer to **int64** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**DocumentStandard** | Pointer to **NullableString** |  | [optional] 
**TrustDocumentType** | Pointer to **NullableString** |  | [optional] 
**SigningStatus** | Pointer to **NullableString** |  | [optional] 
**VerificationStatus** | Pointer to **NullableString** |  | [optional] 
**SignedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 
**SourceStorageObjectId** | Pointer to **NullableString** |  | [optional] 
**SourceSha256** | Pointer to **NullableString** |  | [optional] 
**SignedStorageObjectId** | Pointer to **NullableString** |  | [optional] 
**SignedSha256** | Pointer to **NullableString** |  | [optional] 
**EvidenceStorageObjectId** | Pointer to **NullableString** |  | [optional] 
**EvidenceSha256** | Pointer to **NullableString** |  | [optional] 
**PrimaryFileUploadId** | Pointer to **NullableString** |  | [optional] 
**FrozenSourceFileUploadId** | Pointer to **NullableString** |  | [optional] 
**SignedFileUploadId** | Pointer to **NullableString** |  | [optional] 
**EvidenceFileUploadId** | Pointer to **NullableString** |  | [optional] 
**LockState** | Pointer to **NullableString** |  | [optional] 
**GraphicalRepresentationFileUploadId** | Pointer to **NullableString** |  | [optional] 
**GraphicalRepresentationStorageObjectId** | Pointer to **NullableString** |  | [optional] 
**GraphicalRepresentationSha256** | Pointer to **NullableString** |  | [optional] 
**GraphicalRepresentationContentType** | Pointer to **NullableString** |  | [optional] 
**GraphicalRepresentationGeneratedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignedDocumentDto

`func NewSignedDocumentDto() *SignedDocumentDto`

NewSignedDocumentDto instantiates a new SignedDocumentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentDtoWithDefaults

`func NewSignedDocumentDtoWithDefaults() *SignedDocumentDto`

NewSignedDocumentDtoWithDefaults instantiates a new SignedDocumentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignedDocumentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignedDocumentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignedDocumentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignedDocumentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SignedDocumentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SignedDocumentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SignedDocumentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignedDocumentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignedDocumentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignedDocumentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SignedDocumentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SignedDocumentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSigned

`func (o *SignedDocumentDto) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *SignedDocumentDto) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *SignedDocumentDto) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *SignedDocumentDto) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetUrl

`func (o *SignedDocumentDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SignedDocumentDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SignedDocumentDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SignedDocumentDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SignedDocumentDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SignedDocumentDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetType

`func (o *SignedDocumentDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignedDocumentDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignedDocumentDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignedDocumentDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SignedDocumentDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SignedDocumentDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SignedDocumentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SignedDocumentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SignedDocumentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SignedDocumentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContentType

`func (o *SignedDocumentDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SignedDocumentDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SignedDocumentDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SignedDocumentDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SignedDocumentDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SignedDocumentDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLengthInBits

`func (o *SignedDocumentDto) GetFileLengthInBits() int64`

GetFileLengthInBits returns the FileLengthInBits field if non-nil, zero value otherwise.

### GetFileLengthInBitsOk

`func (o *SignedDocumentDto) GetFileLengthInBitsOk() (*int64, bool)`

GetFileLengthInBitsOk returns a tuple with the FileLengthInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLengthInBits

`func (o *SignedDocumentDto) SetFileLengthInBits(v int64)`

SetFileLengthInBits sets FileLengthInBits field to given value.

### HasFileLengthInBits

`func (o *SignedDocumentDto) HasFileLengthInBits() bool`

HasFileLengthInBits returns a boolean if a field has been set.

### GetTenantId

`func (o *SignedDocumentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SignedDocumentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SignedDocumentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SignedDocumentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SignedDocumentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SignedDocumentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetContactId

`func (o *SignedDocumentDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SignedDocumentDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SignedDocumentDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SignedDocumentDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SignedDocumentDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SignedDocumentDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetUserId

`func (o *SignedDocumentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignedDocumentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SignedDocumentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SignedDocumentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SignedDocumentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SignedDocumentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEnrollmentId

`func (o *SignedDocumentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SignedDocumentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SignedDocumentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SignedDocumentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SignedDocumentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SignedDocumentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDocumentStandard

`func (o *SignedDocumentDto) GetDocumentStandard() string`

GetDocumentStandard returns the DocumentStandard field if non-nil, zero value otherwise.

### GetDocumentStandardOk

`func (o *SignedDocumentDto) GetDocumentStandardOk() (*string, bool)`

GetDocumentStandardOk returns a tuple with the DocumentStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentStandard

`func (o *SignedDocumentDto) SetDocumentStandard(v string)`

SetDocumentStandard sets DocumentStandard field to given value.

### HasDocumentStandard

`func (o *SignedDocumentDto) HasDocumentStandard() bool`

HasDocumentStandard returns a boolean if a field has been set.

### SetDocumentStandardNil

`func (o *SignedDocumentDto) SetDocumentStandardNil(b bool)`

 SetDocumentStandardNil sets the value for DocumentStandard to be an explicit nil

### UnsetDocumentStandard
`func (o *SignedDocumentDto) UnsetDocumentStandard()`

UnsetDocumentStandard ensures that no value is present for DocumentStandard, not even an explicit nil
### GetTrustDocumentType

`func (o *SignedDocumentDto) GetTrustDocumentType() string`

GetTrustDocumentType returns the TrustDocumentType field if non-nil, zero value otherwise.

### GetTrustDocumentTypeOk

`func (o *SignedDocumentDto) GetTrustDocumentTypeOk() (*string, bool)`

GetTrustDocumentTypeOk returns a tuple with the TrustDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustDocumentType

`func (o *SignedDocumentDto) SetTrustDocumentType(v string)`

SetTrustDocumentType sets TrustDocumentType field to given value.

### HasTrustDocumentType

`func (o *SignedDocumentDto) HasTrustDocumentType() bool`

HasTrustDocumentType returns a boolean if a field has been set.

### SetTrustDocumentTypeNil

`func (o *SignedDocumentDto) SetTrustDocumentTypeNil(b bool)`

 SetTrustDocumentTypeNil sets the value for TrustDocumentType to be an explicit nil

### UnsetTrustDocumentType
`func (o *SignedDocumentDto) UnsetTrustDocumentType()`

UnsetTrustDocumentType ensures that no value is present for TrustDocumentType, not even an explicit nil
### GetSigningStatus

`func (o *SignedDocumentDto) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *SignedDocumentDto) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *SignedDocumentDto) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *SignedDocumentDto) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.

### SetSigningStatusNil

`func (o *SignedDocumentDto) SetSigningStatusNil(b bool)`

 SetSigningStatusNil sets the value for SigningStatus to be an explicit nil

### UnsetSigningStatus
`func (o *SignedDocumentDto) UnsetSigningStatus()`

UnsetSigningStatus ensures that no value is present for SigningStatus, not even an explicit nil
### GetVerificationStatus

`func (o *SignedDocumentDto) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *SignedDocumentDto) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *SignedDocumentDto) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *SignedDocumentDto) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### SetVerificationStatusNil

`func (o *SignedDocumentDto) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *SignedDocumentDto) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
### GetSignedAtUtc

`func (o *SignedDocumentDto) GetSignedAtUtc() time.Time`

GetSignedAtUtc returns the SignedAtUtc field if non-nil, zero value otherwise.

### GetSignedAtUtcOk

`func (o *SignedDocumentDto) GetSignedAtUtcOk() (*time.Time, bool)`

GetSignedAtUtcOk returns a tuple with the SignedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAtUtc

`func (o *SignedDocumentDto) SetSignedAtUtc(v time.Time)`

SetSignedAtUtc sets SignedAtUtc field to given value.

### HasSignedAtUtc

`func (o *SignedDocumentDto) HasSignedAtUtc() bool`

HasSignedAtUtc returns a boolean if a field has been set.

### SetSignedAtUtcNil

`func (o *SignedDocumentDto) SetSignedAtUtcNil(b bool)`

 SetSignedAtUtcNil sets the value for SignedAtUtc to be an explicit nil

### UnsetSignedAtUtc
`func (o *SignedDocumentDto) UnsetSignedAtUtc()`

UnsetSignedAtUtc ensures that no value is present for SignedAtUtc, not even an explicit nil
### GetCorrelationId

`func (o *SignedDocumentDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignedDocumentDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SignedDocumentDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SignedDocumentDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SignedDocumentDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SignedDocumentDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *SignedDocumentDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *SignedDocumentDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *SignedDocumentDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *SignedDocumentDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *SignedDocumentDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *SignedDocumentDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil
### GetSourceStorageObjectId

`func (o *SignedDocumentDto) GetSourceStorageObjectId() string`

GetSourceStorageObjectId returns the SourceStorageObjectId field if non-nil, zero value otherwise.

### GetSourceStorageObjectIdOk

`func (o *SignedDocumentDto) GetSourceStorageObjectIdOk() (*string, bool)`

GetSourceStorageObjectIdOk returns a tuple with the SourceStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStorageObjectId

`func (o *SignedDocumentDto) SetSourceStorageObjectId(v string)`

SetSourceStorageObjectId sets SourceStorageObjectId field to given value.

### HasSourceStorageObjectId

`func (o *SignedDocumentDto) HasSourceStorageObjectId() bool`

HasSourceStorageObjectId returns a boolean if a field has been set.

### SetSourceStorageObjectIdNil

`func (o *SignedDocumentDto) SetSourceStorageObjectIdNil(b bool)`

 SetSourceStorageObjectIdNil sets the value for SourceStorageObjectId to be an explicit nil

### UnsetSourceStorageObjectId
`func (o *SignedDocumentDto) UnsetSourceStorageObjectId()`

UnsetSourceStorageObjectId ensures that no value is present for SourceStorageObjectId, not even an explicit nil
### GetSourceSha256

`func (o *SignedDocumentDto) GetSourceSha256() string`

GetSourceSha256 returns the SourceSha256 field if non-nil, zero value otherwise.

### GetSourceSha256Ok

`func (o *SignedDocumentDto) GetSourceSha256Ok() (*string, bool)`

GetSourceSha256Ok returns a tuple with the SourceSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSha256

`func (o *SignedDocumentDto) SetSourceSha256(v string)`

SetSourceSha256 sets SourceSha256 field to given value.

### HasSourceSha256

`func (o *SignedDocumentDto) HasSourceSha256() bool`

HasSourceSha256 returns a boolean if a field has been set.

### SetSourceSha256Nil

`func (o *SignedDocumentDto) SetSourceSha256Nil(b bool)`

 SetSourceSha256Nil sets the value for SourceSha256 to be an explicit nil

### UnsetSourceSha256
`func (o *SignedDocumentDto) UnsetSourceSha256()`

UnsetSourceSha256 ensures that no value is present for SourceSha256, not even an explicit nil
### GetSignedStorageObjectId

`func (o *SignedDocumentDto) GetSignedStorageObjectId() string`

GetSignedStorageObjectId returns the SignedStorageObjectId field if non-nil, zero value otherwise.

### GetSignedStorageObjectIdOk

`func (o *SignedDocumentDto) GetSignedStorageObjectIdOk() (*string, bool)`

GetSignedStorageObjectIdOk returns a tuple with the SignedStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedStorageObjectId

`func (o *SignedDocumentDto) SetSignedStorageObjectId(v string)`

SetSignedStorageObjectId sets SignedStorageObjectId field to given value.

### HasSignedStorageObjectId

`func (o *SignedDocumentDto) HasSignedStorageObjectId() bool`

HasSignedStorageObjectId returns a boolean if a field has been set.

### SetSignedStorageObjectIdNil

`func (o *SignedDocumentDto) SetSignedStorageObjectIdNil(b bool)`

 SetSignedStorageObjectIdNil sets the value for SignedStorageObjectId to be an explicit nil

### UnsetSignedStorageObjectId
`func (o *SignedDocumentDto) UnsetSignedStorageObjectId()`

UnsetSignedStorageObjectId ensures that no value is present for SignedStorageObjectId, not even an explicit nil
### GetSignedSha256

`func (o *SignedDocumentDto) GetSignedSha256() string`

GetSignedSha256 returns the SignedSha256 field if non-nil, zero value otherwise.

### GetSignedSha256Ok

`func (o *SignedDocumentDto) GetSignedSha256Ok() (*string, bool)`

GetSignedSha256Ok returns a tuple with the SignedSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedSha256

`func (o *SignedDocumentDto) SetSignedSha256(v string)`

SetSignedSha256 sets SignedSha256 field to given value.

### HasSignedSha256

`func (o *SignedDocumentDto) HasSignedSha256() bool`

HasSignedSha256 returns a boolean if a field has been set.

### SetSignedSha256Nil

`func (o *SignedDocumentDto) SetSignedSha256Nil(b bool)`

 SetSignedSha256Nil sets the value for SignedSha256 to be an explicit nil

### UnsetSignedSha256
`func (o *SignedDocumentDto) UnsetSignedSha256()`

UnsetSignedSha256 ensures that no value is present for SignedSha256, not even an explicit nil
### GetEvidenceStorageObjectId

`func (o *SignedDocumentDto) GetEvidenceStorageObjectId() string`

GetEvidenceStorageObjectId returns the EvidenceStorageObjectId field if non-nil, zero value otherwise.

### GetEvidenceStorageObjectIdOk

`func (o *SignedDocumentDto) GetEvidenceStorageObjectIdOk() (*string, bool)`

GetEvidenceStorageObjectIdOk returns a tuple with the EvidenceStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceStorageObjectId

`func (o *SignedDocumentDto) SetEvidenceStorageObjectId(v string)`

SetEvidenceStorageObjectId sets EvidenceStorageObjectId field to given value.

### HasEvidenceStorageObjectId

`func (o *SignedDocumentDto) HasEvidenceStorageObjectId() bool`

HasEvidenceStorageObjectId returns a boolean if a field has been set.

### SetEvidenceStorageObjectIdNil

`func (o *SignedDocumentDto) SetEvidenceStorageObjectIdNil(b bool)`

 SetEvidenceStorageObjectIdNil sets the value for EvidenceStorageObjectId to be an explicit nil

### UnsetEvidenceStorageObjectId
`func (o *SignedDocumentDto) UnsetEvidenceStorageObjectId()`

UnsetEvidenceStorageObjectId ensures that no value is present for EvidenceStorageObjectId, not even an explicit nil
### GetEvidenceSha256

`func (o *SignedDocumentDto) GetEvidenceSha256() string`

GetEvidenceSha256 returns the EvidenceSha256 field if non-nil, zero value otherwise.

### GetEvidenceSha256Ok

`func (o *SignedDocumentDto) GetEvidenceSha256Ok() (*string, bool)`

GetEvidenceSha256Ok returns a tuple with the EvidenceSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSha256

`func (o *SignedDocumentDto) SetEvidenceSha256(v string)`

SetEvidenceSha256 sets EvidenceSha256 field to given value.

### HasEvidenceSha256

`func (o *SignedDocumentDto) HasEvidenceSha256() bool`

HasEvidenceSha256 returns a boolean if a field has been set.

### SetEvidenceSha256Nil

`func (o *SignedDocumentDto) SetEvidenceSha256Nil(b bool)`

 SetEvidenceSha256Nil sets the value for EvidenceSha256 to be an explicit nil

### UnsetEvidenceSha256
`func (o *SignedDocumentDto) UnsetEvidenceSha256()`

UnsetEvidenceSha256 ensures that no value is present for EvidenceSha256, not even an explicit nil
### GetPrimaryFileUploadId

`func (o *SignedDocumentDto) GetPrimaryFileUploadId() string`

GetPrimaryFileUploadId returns the PrimaryFileUploadId field if non-nil, zero value otherwise.

### GetPrimaryFileUploadIdOk

`func (o *SignedDocumentDto) GetPrimaryFileUploadIdOk() (*string, bool)`

GetPrimaryFileUploadIdOk returns a tuple with the PrimaryFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFileUploadId

`func (o *SignedDocumentDto) SetPrimaryFileUploadId(v string)`

SetPrimaryFileUploadId sets PrimaryFileUploadId field to given value.

### HasPrimaryFileUploadId

`func (o *SignedDocumentDto) HasPrimaryFileUploadId() bool`

HasPrimaryFileUploadId returns a boolean if a field has been set.

### SetPrimaryFileUploadIdNil

`func (o *SignedDocumentDto) SetPrimaryFileUploadIdNil(b bool)`

 SetPrimaryFileUploadIdNil sets the value for PrimaryFileUploadId to be an explicit nil

### UnsetPrimaryFileUploadId
`func (o *SignedDocumentDto) UnsetPrimaryFileUploadId()`

UnsetPrimaryFileUploadId ensures that no value is present for PrimaryFileUploadId, not even an explicit nil
### GetFrozenSourceFileUploadId

`func (o *SignedDocumentDto) GetFrozenSourceFileUploadId() string`

GetFrozenSourceFileUploadId returns the FrozenSourceFileUploadId field if non-nil, zero value otherwise.

### GetFrozenSourceFileUploadIdOk

`func (o *SignedDocumentDto) GetFrozenSourceFileUploadIdOk() (*string, bool)`

GetFrozenSourceFileUploadIdOk returns a tuple with the FrozenSourceFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenSourceFileUploadId

`func (o *SignedDocumentDto) SetFrozenSourceFileUploadId(v string)`

SetFrozenSourceFileUploadId sets FrozenSourceFileUploadId field to given value.

### HasFrozenSourceFileUploadId

`func (o *SignedDocumentDto) HasFrozenSourceFileUploadId() bool`

HasFrozenSourceFileUploadId returns a boolean if a field has been set.

### SetFrozenSourceFileUploadIdNil

`func (o *SignedDocumentDto) SetFrozenSourceFileUploadIdNil(b bool)`

 SetFrozenSourceFileUploadIdNil sets the value for FrozenSourceFileUploadId to be an explicit nil

### UnsetFrozenSourceFileUploadId
`func (o *SignedDocumentDto) UnsetFrozenSourceFileUploadId()`

UnsetFrozenSourceFileUploadId ensures that no value is present for FrozenSourceFileUploadId, not even an explicit nil
### GetSignedFileUploadId

`func (o *SignedDocumentDto) GetSignedFileUploadId() string`

GetSignedFileUploadId returns the SignedFileUploadId field if non-nil, zero value otherwise.

### GetSignedFileUploadIdOk

`func (o *SignedDocumentDto) GetSignedFileUploadIdOk() (*string, bool)`

GetSignedFileUploadIdOk returns a tuple with the SignedFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedFileUploadId

`func (o *SignedDocumentDto) SetSignedFileUploadId(v string)`

SetSignedFileUploadId sets SignedFileUploadId field to given value.

### HasSignedFileUploadId

`func (o *SignedDocumentDto) HasSignedFileUploadId() bool`

HasSignedFileUploadId returns a boolean if a field has been set.

### SetSignedFileUploadIdNil

`func (o *SignedDocumentDto) SetSignedFileUploadIdNil(b bool)`

 SetSignedFileUploadIdNil sets the value for SignedFileUploadId to be an explicit nil

### UnsetSignedFileUploadId
`func (o *SignedDocumentDto) UnsetSignedFileUploadId()`

UnsetSignedFileUploadId ensures that no value is present for SignedFileUploadId, not even an explicit nil
### GetEvidenceFileUploadId

`func (o *SignedDocumentDto) GetEvidenceFileUploadId() string`

GetEvidenceFileUploadId returns the EvidenceFileUploadId field if non-nil, zero value otherwise.

### GetEvidenceFileUploadIdOk

`func (o *SignedDocumentDto) GetEvidenceFileUploadIdOk() (*string, bool)`

GetEvidenceFileUploadIdOk returns a tuple with the EvidenceFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceFileUploadId

`func (o *SignedDocumentDto) SetEvidenceFileUploadId(v string)`

SetEvidenceFileUploadId sets EvidenceFileUploadId field to given value.

### HasEvidenceFileUploadId

`func (o *SignedDocumentDto) HasEvidenceFileUploadId() bool`

HasEvidenceFileUploadId returns a boolean if a field has been set.

### SetEvidenceFileUploadIdNil

`func (o *SignedDocumentDto) SetEvidenceFileUploadIdNil(b bool)`

 SetEvidenceFileUploadIdNil sets the value for EvidenceFileUploadId to be an explicit nil

### UnsetEvidenceFileUploadId
`func (o *SignedDocumentDto) UnsetEvidenceFileUploadId()`

UnsetEvidenceFileUploadId ensures that no value is present for EvidenceFileUploadId, not even an explicit nil
### GetLockState

`func (o *SignedDocumentDto) GetLockState() string`

GetLockState returns the LockState field if non-nil, zero value otherwise.

### GetLockStateOk

`func (o *SignedDocumentDto) GetLockStateOk() (*string, bool)`

GetLockStateOk returns a tuple with the LockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockState

`func (o *SignedDocumentDto) SetLockState(v string)`

SetLockState sets LockState field to given value.

### HasLockState

`func (o *SignedDocumentDto) HasLockState() bool`

HasLockState returns a boolean if a field has been set.

### SetLockStateNil

`func (o *SignedDocumentDto) SetLockStateNil(b bool)`

 SetLockStateNil sets the value for LockState to be an explicit nil

### UnsetLockState
`func (o *SignedDocumentDto) UnsetLockState()`

UnsetLockState ensures that no value is present for LockState, not even an explicit nil
### GetGraphicalRepresentationFileUploadId

`func (o *SignedDocumentDto) GetGraphicalRepresentationFileUploadId() string`

GetGraphicalRepresentationFileUploadId returns the GraphicalRepresentationFileUploadId field if non-nil, zero value otherwise.

### GetGraphicalRepresentationFileUploadIdOk

`func (o *SignedDocumentDto) GetGraphicalRepresentationFileUploadIdOk() (*string, bool)`

GetGraphicalRepresentationFileUploadIdOk returns a tuple with the GraphicalRepresentationFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicalRepresentationFileUploadId

`func (o *SignedDocumentDto) SetGraphicalRepresentationFileUploadId(v string)`

SetGraphicalRepresentationFileUploadId sets GraphicalRepresentationFileUploadId field to given value.

### HasGraphicalRepresentationFileUploadId

`func (o *SignedDocumentDto) HasGraphicalRepresentationFileUploadId() bool`

HasGraphicalRepresentationFileUploadId returns a boolean if a field has been set.

### SetGraphicalRepresentationFileUploadIdNil

`func (o *SignedDocumentDto) SetGraphicalRepresentationFileUploadIdNil(b bool)`

 SetGraphicalRepresentationFileUploadIdNil sets the value for GraphicalRepresentationFileUploadId to be an explicit nil

### UnsetGraphicalRepresentationFileUploadId
`func (o *SignedDocumentDto) UnsetGraphicalRepresentationFileUploadId()`

UnsetGraphicalRepresentationFileUploadId ensures that no value is present for GraphicalRepresentationFileUploadId, not even an explicit nil
### GetGraphicalRepresentationStorageObjectId

`func (o *SignedDocumentDto) GetGraphicalRepresentationStorageObjectId() string`

GetGraphicalRepresentationStorageObjectId returns the GraphicalRepresentationStorageObjectId field if non-nil, zero value otherwise.

### GetGraphicalRepresentationStorageObjectIdOk

`func (o *SignedDocumentDto) GetGraphicalRepresentationStorageObjectIdOk() (*string, bool)`

GetGraphicalRepresentationStorageObjectIdOk returns a tuple with the GraphicalRepresentationStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicalRepresentationStorageObjectId

`func (o *SignedDocumentDto) SetGraphicalRepresentationStorageObjectId(v string)`

SetGraphicalRepresentationStorageObjectId sets GraphicalRepresentationStorageObjectId field to given value.

### HasGraphicalRepresentationStorageObjectId

`func (o *SignedDocumentDto) HasGraphicalRepresentationStorageObjectId() bool`

HasGraphicalRepresentationStorageObjectId returns a boolean if a field has been set.

### SetGraphicalRepresentationStorageObjectIdNil

`func (o *SignedDocumentDto) SetGraphicalRepresentationStorageObjectIdNil(b bool)`

 SetGraphicalRepresentationStorageObjectIdNil sets the value for GraphicalRepresentationStorageObjectId to be an explicit nil

### UnsetGraphicalRepresentationStorageObjectId
`func (o *SignedDocumentDto) UnsetGraphicalRepresentationStorageObjectId()`

UnsetGraphicalRepresentationStorageObjectId ensures that no value is present for GraphicalRepresentationStorageObjectId, not even an explicit nil
### GetGraphicalRepresentationSha256

`func (o *SignedDocumentDto) GetGraphicalRepresentationSha256() string`

GetGraphicalRepresentationSha256 returns the GraphicalRepresentationSha256 field if non-nil, zero value otherwise.

### GetGraphicalRepresentationSha256Ok

`func (o *SignedDocumentDto) GetGraphicalRepresentationSha256Ok() (*string, bool)`

GetGraphicalRepresentationSha256Ok returns a tuple with the GraphicalRepresentationSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicalRepresentationSha256

`func (o *SignedDocumentDto) SetGraphicalRepresentationSha256(v string)`

SetGraphicalRepresentationSha256 sets GraphicalRepresentationSha256 field to given value.

### HasGraphicalRepresentationSha256

`func (o *SignedDocumentDto) HasGraphicalRepresentationSha256() bool`

HasGraphicalRepresentationSha256 returns a boolean if a field has been set.

### SetGraphicalRepresentationSha256Nil

`func (o *SignedDocumentDto) SetGraphicalRepresentationSha256Nil(b bool)`

 SetGraphicalRepresentationSha256Nil sets the value for GraphicalRepresentationSha256 to be an explicit nil

### UnsetGraphicalRepresentationSha256
`func (o *SignedDocumentDto) UnsetGraphicalRepresentationSha256()`

UnsetGraphicalRepresentationSha256 ensures that no value is present for GraphicalRepresentationSha256, not even an explicit nil
### GetGraphicalRepresentationContentType

`func (o *SignedDocumentDto) GetGraphicalRepresentationContentType() string`

GetGraphicalRepresentationContentType returns the GraphicalRepresentationContentType field if non-nil, zero value otherwise.

### GetGraphicalRepresentationContentTypeOk

`func (o *SignedDocumentDto) GetGraphicalRepresentationContentTypeOk() (*string, bool)`

GetGraphicalRepresentationContentTypeOk returns a tuple with the GraphicalRepresentationContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicalRepresentationContentType

`func (o *SignedDocumentDto) SetGraphicalRepresentationContentType(v string)`

SetGraphicalRepresentationContentType sets GraphicalRepresentationContentType field to given value.

### HasGraphicalRepresentationContentType

`func (o *SignedDocumentDto) HasGraphicalRepresentationContentType() bool`

HasGraphicalRepresentationContentType returns a boolean if a field has been set.

### SetGraphicalRepresentationContentTypeNil

`func (o *SignedDocumentDto) SetGraphicalRepresentationContentTypeNil(b bool)`

 SetGraphicalRepresentationContentTypeNil sets the value for GraphicalRepresentationContentType to be an explicit nil

### UnsetGraphicalRepresentationContentType
`func (o *SignedDocumentDto) UnsetGraphicalRepresentationContentType()`

UnsetGraphicalRepresentationContentType ensures that no value is present for GraphicalRepresentationContentType, not even an explicit nil
### GetGraphicalRepresentationGeneratedAtUtc

`func (o *SignedDocumentDto) GetGraphicalRepresentationGeneratedAtUtc() time.Time`

GetGraphicalRepresentationGeneratedAtUtc returns the GraphicalRepresentationGeneratedAtUtc field if non-nil, zero value otherwise.

### GetGraphicalRepresentationGeneratedAtUtcOk

`func (o *SignedDocumentDto) GetGraphicalRepresentationGeneratedAtUtcOk() (*time.Time, bool)`

GetGraphicalRepresentationGeneratedAtUtcOk returns a tuple with the GraphicalRepresentationGeneratedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicalRepresentationGeneratedAtUtc

`func (o *SignedDocumentDto) SetGraphicalRepresentationGeneratedAtUtc(v time.Time)`

SetGraphicalRepresentationGeneratedAtUtc sets GraphicalRepresentationGeneratedAtUtc field to given value.

### HasGraphicalRepresentationGeneratedAtUtc

`func (o *SignedDocumentDto) HasGraphicalRepresentationGeneratedAtUtc() bool`

HasGraphicalRepresentationGeneratedAtUtc returns a boolean if a field has been set.

### SetGraphicalRepresentationGeneratedAtUtcNil

`func (o *SignedDocumentDto) SetGraphicalRepresentationGeneratedAtUtcNil(b bool)`

 SetGraphicalRepresentationGeneratedAtUtcNil sets the value for GraphicalRepresentationGeneratedAtUtc to be an explicit nil

### UnsetGraphicalRepresentationGeneratedAtUtc
`func (o *SignedDocumentDto) UnsetGraphicalRepresentationGeneratedAtUtc()`

UnsetGraphicalRepresentationGeneratedAtUtc ensures that no value is present for GraphicalRepresentationGeneratedAtUtc, not even an explicit nil
### GetContactName

`func (o *SignedDocumentDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SignedDocumentDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SignedDocumentDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SignedDocumentDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *SignedDocumentDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *SignedDocumentDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


