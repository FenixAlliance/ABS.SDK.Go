# TrustSigningRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedDocumentId** | Pointer to **NullableString** |  | [optional] 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**RequestedFormat** | Pointer to **NullableString** |  | [optional] 
**RequestedPurpose** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**SourceStorageObjectId** | Pointer to **NullableString** |  | [optional] 
**SourceSha256** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrustSigningRequestDto

`func NewTrustSigningRequestDto() *TrustSigningRequestDto`

NewTrustSigningRequestDto instantiates a new TrustSigningRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustSigningRequestDtoWithDefaults

`func NewTrustSigningRequestDtoWithDefaults() *TrustSigningRequestDto`

NewTrustSigningRequestDtoWithDefaults instantiates a new TrustSigningRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedDocumentId

`func (o *TrustSigningRequestDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *TrustSigningRequestDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *TrustSigningRequestDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.

### HasSignedDocumentId

`func (o *TrustSigningRequestDto) HasSignedDocumentId() bool`

HasSignedDocumentId returns a boolean if a field has been set.

### SetSignedDocumentIdNil

`func (o *TrustSigningRequestDto) SetSignedDocumentIdNil(b bool)`

 SetSignedDocumentIdNil sets the value for SignedDocumentId to be an explicit nil

### UnsetSignedDocumentId
`func (o *TrustSigningRequestDto) UnsetSignedDocumentId()`

UnsetSignedDocumentId ensures that no value is present for SignedDocumentId, not even an explicit nil
### GetSigningProfileId

`func (o *TrustSigningRequestDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *TrustSigningRequestDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *TrustSigningRequestDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *TrustSigningRequestDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *TrustSigningRequestDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *TrustSigningRequestDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil
### GetSigningCertificateId

`func (o *TrustSigningRequestDto) GetSigningCertificateId() string`

GetSigningCertificateId returns the SigningCertificateId field if non-nil, zero value otherwise.

### GetSigningCertificateIdOk

`func (o *TrustSigningRequestDto) GetSigningCertificateIdOk() (*string, bool)`

GetSigningCertificateIdOk returns a tuple with the SigningCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateId

`func (o *TrustSigningRequestDto) SetSigningCertificateId(v string)`

SetSigningCertificateId sets SigningCertificateId field to given value.

### HasSigningCertificateId

`func (o *TrustSigningRequestDto) HasSigningCertificateId() bool`

HasSigningCertificateId returns a boolean if a field has been set.

### SetSigningCertificateIdNil

`func (o *TrustSigningRequestDto) SetSigningCertificateIdNil(b bool)`

 SetSigningCertificateIdNil sets the value for SigningCertificateId to be an explicit nil

### UnsetSigningCertificateId
`func (o *TrustSigningRequestDto) UnsetSigningCertificateId()`

UnsetSigningCertificateId ensures that no value is present for SigningCertificateId, not even an explicit nil
### GetContactId

`func (o *TrustSigningRequestDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TrustSigningRequestDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TrustSigningRequestDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TrustSigningRequestDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *TrustSigningRequestDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *TrustSigningRequestDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetRequestedFormat

`func (o *TrustSigningRequestDto) GetRequestedFormat() string`

GetRequestedFormat returns the RequestedFormat field if non-nil, zero value otherwise.

### GetRequestedFormatOk

`func (o *TrustSigningRequestDto) GetRequestedFormatOk() (*string, bool)`

GetRequestedFormatOk returns a tuple with the RequestedFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFormat

`func (o *TrustSigningRequestDto) SetRequestedFormat(v string)`

SetRequestedFormat sets RequestedFormat field to given value.

### HasRequestedFormat

`func (o *TrustSigningRequestDto) HasRequestedFormat() bool`

HasRequestedFormat returns a boolean if a field has been set.

### SetRequestedFormatNil

`func (o *TrustSigningRequestDto) SetRequestedFormatNil(b bool)`

 SetRequestedFormatNil sets the value for RequestedFormat to be an explicit nil

### UnsetRequestedFormat
`func (o *TrustSigningRequestDto) UnsetRequestedFormat()`

UnsetRequestedFormat ensures that no value is present for RequestedFormat, not even an explicit nil
### GetRequestedPurpose

`func (o *TrustSigningRequestDto) GetRequestedPurpose() string`

GetRequestedPurpose returns the RequestedPurpose field if non-nil, zero value otherwise.

### GetRequestedPurposeOk

`func (o *TrustSigningRequestDto) GetRequestedPurposeOk() (*string, bool)`

GetRequestedPurposeOk returns a tuple with the RequestedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPurpose

`func (o *TrustSigningRequestDto) SetRequestedPurpose(v string)`

SetRequestedPurpose sets RequestedPurpose field to given value.

### HasRequestedPurpose

`func (o *TrustSigningRequestDto) HasRequestedPurpose() bool`

HasRequestedPurpose returns a boolean if a field has been set.

### SetRequestedPurposeNil

`func (o *TrustSigningRequestDto) SetRequestedPurposeNil(b bool)`

 SetRequestedPurposeNil sets the value for RequestedPurpose to be an explicit nil

### UnsetRequestedPurpose
`func (o *TrustSigningRequestDto) UnsetRequestedPurpose()`

UnsetRequestedPurpose ensures that no value is present for RequestedPurpose, not even an explicit nil
### GetCorrelationId

`func (o *TrustSigningRequestDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TrustSigningRequestDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TrustSigningRequestDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TrustSigningRequestDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *TrustSigningRequestDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *TrustSigningRequestDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetSourceStorageObjectId

`func (o *TrustSigningRequestDto) GetSourceStorageObjectId() string`

GetSourceStorageObjectId returns the SourceStorageObjectId field if non-nil, zero value otherwise.

### GetSourceStorageObjectIdOk

`func (o *TrustSigningRequestDto) GetSourceStorageObjectIdOk() (*string, bool)`

GetSourceStorageObjectIdOk returns a tuple with the SourceStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStorageObjectId

`func (o *TrustSigningRequestDto) SetSourceStorageObjectId(v string)`

SetSourceStorageObjectId sets SourceStorageObjectId field to given value.

### HasSourceStorageObjectId

`func (o *TrustSigningRequestDto) HasSourceStorageObjectId() bool`

HasSourceStorageObjectId returns a boolean if a field has been set.

### SetSourceStorageObjectIdNil

`func (o *TrustSigningRequestDto) SetSourceStorageObjectIdNil(b bool)`

 SetSourceStorageObjectIdNil sets the value for SourceStorageObjectId to be an explicit nil

### UnsetSourceStorageObjectId
`func (o *TrustSigningRequestDto) UnsetSourceStorageObjectId()`

UnsetSourceStorageObjectId ensures that no value is present for SourceStorageObjectId, not even an explicit nil
### GetSourceSha256

`func (o *TrustSigningRequestDto) GetSourceSha256() string`

GetSourceSha256 returns the SourceSha256 field if non-nil, zero value otherwise.

### GetSourceSha256Ok

`func (o *TrustSigningRequestDto) GetSourceSha256Ok() (*string, bool)`

GetSourceSha256Ok returns a tuple with the SourceSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSha256

`func (o *TrustSigningRequestDto) SetSourceSha256(v string)`

SetSourceSha256 sets SourceSha256 field to given value.

### HasSourceSha256

`func (o *TrustSigningRequestDto) HasSourceSha256() bool`

HasSourceSha256 returns a boolean if a field has been set.

### SetSourceSha256Nil

`func (o *TrustSigningRequestDto) SetSourceSha256Nil(b bool)`

 SetSourceSha256Nil sets the value for SourceSha256 to be an explicit nil

### UnsetSourceSha256
`func (o *TrustSigningRequestDto) UnsetSourceSha256()`

UnsetSourceSha256 ensures that no value is present for SourceSha256, not even an explicit nil
### GetExternalReference

`func (o *TrustSigningRequestDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *TrustSigningRequestDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *TrustSigningRequestDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *TrustSigningRequestDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *TrustSigningRequestDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *TrustSigningRequestDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil
### GetDryRun

`func (o *TrustSigningRequestDto) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *TrustSigningRequestDto) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *TrustSigningRequestDto) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *TrustSigningRequestDto) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


