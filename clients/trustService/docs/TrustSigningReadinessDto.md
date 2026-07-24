# TrustSigningReadinessDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanProceed** | Pointer to **bool** |  | [optional] 
**BlockingReasons** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**ResolvedDocumentTitle** | Pointer to **NullableString** |  | [optional] 
**ResolvedProfileDisplayName** | Pointer to **NullableString** |  | [optional] 
**ResolvedCertificateTitle** | Pointer to **NullableString** |  | [optional] 
**ExpectedSignatureFormat** | Pointer to **NullableString** |  | [optional] 
**ExpectedSignaturePurpose** | Pointer to **NullableString** |  | [optional] 
**ExpectedDigestAlgorithm** | Pointer to **NullableString** |  | [optional] 
**ExpectedSignatureAlgorithm** | Pointer to **NullableString** |  | [optional] 
**ExpectedCanonicalizationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**PolicyIdentifier** | Pointer to **NullableString** |  | [optional] 
**AuthorityProfile** | Pointer to **NullableString** |  | [optional] 
**RequiresCustodyProvider** | Pointer to **bool** |  | [optional] 
**RequiresSourceArtifact** | Pointer to **bool** |  | [optional] 
**RequiresCertificate** | Pointer to **bool** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrustSigningReadinessDto

`func NewTrustSigningReadinessDto() *TrustSigningReadinessDto`

NewTrustSigningReadinessDto instantiates a new TrustSigningReadinessDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustSigningReadinessDtoWithDefaults

`func NewTrustSigningReadinessDtoWithDefaults() *TrustSigningReadinessDto`

NewTrustSigningReadinessDtoWithDefaults instantiates a new TrustSigningReadinessDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanProceed

`func (o *TrustSigningReadinessDto) GetCanProceed() bool`

GetCanProceed returns the CanProceed field if non-nil, zero value otherwise.

### GetCanProceedOk

`func (o *TrustSigningReadinessDto) GetCanProceedOk() (*bool, bool)`

GetCanProceedOk returns a tuple with the CanProceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanProceed

`func (o *TrustSigningReadinessDto) SetCanProceed(v bool)`

SetCanProceed sets CanProceed field to given value.

### HasCanProceed

`func (o *TrustSigningReadinessDto) HasCanProceed() bool`

HasCanProceed returns a boolean if a field has been set.

### GetBlockingReasons

`func (o *TrustSigningReadinessDto) GetBlockingReasons() []string`

GetBlockingReasons returns the BlockingReasons field if non-nil, zero value otherwise.

### GetBlockingReasonsOk

`func (o *TrustSigningReadinessDto) GetBlockingReasonsOk() (*[]string, bool)`

GetBlockingReasonsOk returns a tuple with the BlockingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingReasons

`func (o *TrustSigningReadinessDto) SetBlockingReasons(v []string)`

SetBlockingReasons sets BlockingReasons field to given value.

### HasBlockingReasons

`func (o *TrustSigningReadinessDto) HasBlockingReasons() bool`

HasBlockingReasons returns a boolean if a field has been set.

### SetBlockingReasonsNil

`func (o *TrustSigningReadinessDto) SetBlockingReasonsNil(b bool)`

 SetBlockingReasonsNil sets the value for BlockingReasons to be an explicit nil

### UnsetBlockingReasons
`func (o *TrustSigningReadinessDto) UnsetBlockingReasons()`

UnsetBlockingReasons ensures that no value is present for BlockingReasons, not even an explicit nil
### GetWarnings

`func (o *TrustSigningReadinessDto) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TrustSigningReadinessDto) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TrustSigningReadinessDto) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *TrustSigningReadinessDto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *TrustSigningReadinessDto) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *TrustSigningReadinessDto) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetResolvedDocumentTitle

`func (o *TrustSigningReadinessDto) GetResolvedDocumentTitle() string`

GetResolvedDocumentTitle returns the ResolvedDocumentTitle field if non-nil, zero value otherwise.

### GetResolvedDocumentTitleOk

`func (o *TrustSigningReadinessDto) GetResolvedDocumentTitleOk() (*string, bool)`

GetResolvedDocumentTitleOk returns a tuple with the ResolvedDocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDocumentTitle

`func (o *TrustSigningReadinessDto) SetResolvedDocumentTitle(v string)`

SetResolvedDocumentTitle sets ResolvedDocumentTitle field to given value.

### HasResolvedDocumentTitle

`func (o *TrustSigningReadinessDto) HasResolvedDocumentTitle() bool`

HasResolvedDocumentTitle returns a boolean if a field has been set.

### SetResolvedDocumentTitleNil

`func (o *TrustSigningReadinessDto) SetResolvedDocumentTitleNil(b bool)`

 SetResolvedDocumentTitleNil sets the value for ResolvedDocumentTitle to be an explicit nil

### UnsetResolvedDocumentTitle
`func (o *TrustSigningReadinessDto) UnsetResolvedDocumentTitle()`

UnsetResolvedDocumentTitle ensures that no value is present for ResolvedDocumentTitle, not even an explicit nil
### GetResolvedProfileDisplayName

`func (o *TrustSigningReadinessDto) GetResolvedProfileDisplayName() string`

GetResolvedProfileDisplayName returns the ResolvedProfileDisplayName field if non-nil, zero value otherwise.

### GetResolvedProfileDisplayNameOk

`func (o *TrustSigningReadinessDto) GetResolvedProfileDisplayNameOk() (*string, bool)`

GetResolvedProfileDisplayNameOk returns a tuple with the ResolvedProfileDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedProfileDisplayName

`func (o *TrustSigningReadinessDto) SetResolvedProfileDisplayName(v string)`

SetResolvedProfileDisplayName sets ResolvedProfileDisplayName field to given value.

### HasResolvedProfileDisplayName

`func (o *TrustSigningReadinessDto) HasResolvedProfileDisplayName() bool`

HasResolvedProfileDisplayName returns a boolean if a field has been set.

### SetResolvedProfileDisplayNameNil

`func (o *TrustSigningReadinessDto) SetResolvedProfileDisplayNameNil(b bool)`

 SetResolvedProfileDisplayNameNil sets the value for ResolvedProfileDisplayName to be an explicit nil

### UnsetResolvedProfileDisplayName
`func (o *TrustSigningReadinessDto) UnsetResolvedProfileDisplayName()`

UnsetResolvedProfileDisplayName ensures that no value is present for ResolvedProfileDisplayName, not even an explicit nil
### GetResolvedCertificateTitle

`func (o *TrustSigningReadinessDto) GetResolvedCertificateTitle() string`

GetResolvedCertificateTitle returns the ResolvedCertificateTitle field if non-nil, zero value otherwise.

### GetResolvedCertificateTitleOk

`func (o *TrustSigningReadinessDto) GetResolvedCertificateTitleOk() (*string, bool)`

GetResolvedCertificateTitleOk returns a tuple with the ResolvedCertificateTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedCertificateTitle

`func (o *TrustSigningReadinessDto) SetResolvedCertificateTitle(v string)`

SetResolvedCertificateTitle sets ResolvedCertificateTitle field to given value.

### HasResolvedCertificateTitle

`func (o *TrustSigningReadinessDto) HasResolvedCertificateTitle() bool`

HasResolvedCertificateTitle returns a boolean if a field has been set.

### SetResolvedCertificateTitleNil

`func (o *TrustSigningReadinessDto) SetResolvedCertificateTitleNil(b bool)`

 SetResolvedCertificateTitleNil sets the value for ResolvedCertificateTitle to be an explicit nil

### UnsetResolvedCertificateTitle
`func (o *TrustSigningReadinessDto) UnsetResolvedCertificateTitle()`

UnsetResolvedCertificateTitle ensures that no value is present for ResolvedCertificateTitle, not even an explicit nil
### GetExpectedSignatureFormat

`func (o *TrustSigningReadinessDto) GetExpectedSignatureFormat() string`

GetExpectedSignatureFormat returns the ExpectedSignatureFormat field if non-nil, zero value otherwise.

### GetExpectedSignatureFormatOk

`func (o *TrustSigningReadinessDto) GetExpectedSignatureFormatOk() (*string, bool)`

GetExpectedSignatureFormatOk returns a tuple with the ExpectedSignatureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSignatureFormat

`func (o *TrustSigningReadinessDto) SetExpectedSignatureFormat(v string)`

SetExpectedSignatureFormat sets ExpectedSignatureFormat field to given value.

### HasExpectedSignatureFormat

`func (o *TrustSigningReadinessDto) HasExpectedSignatureFormat() bool`

HasExpectedSignatureFormat returns a boolean if a field has been set.

### SetExpectedSignatureFormatNil

`func (o *TrustSigningReadinessDto) SetExpectedSignatureFormatNil(b bool)`

 SetExpectedSignatureFormatNil sets the value for ExpectedSignatureFormat to be an explicit nil

### UnsetExpectedSignatureFormat
`func (o *TrustSigningReadinessDto) UnsetExpectedSignatureFormat()`

UnsetExpectedSignatureFormat ensures that no value is present for ExpectedSignatureFormat, not even an explicit nil
### GetExpectedSignaturePurpose

`func (o *TrustSigningReadinessDto) GetExpectedSignaturePurpose() string`

GetExpectedSignaturePurpose returns the ExpectedSignaturePurpose field if non-nil, zero value otherwise.

### GetExpectedSignaturePurposeOk

`func (o *TrustSigningReadinessDto) GetExpectedSignaturePurposeOk() (*string, bool)`

GetExpectedSignaturePurposeOk returns a tuple with the ExpectedSignaturePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSignaturePurpose

`func (o *TrustSigningReadinessDto) SetExpectedSignaturePurpose(v string)`

SetExpectedSignaturePurpose sets ExpectedSignaturePurpose field to given value.

### HasExpectedSignaturePurpose

`func (o *TrustSigningReadinessDto) HasExpectedSignaturePurpose() bool`

HasExpectedSignaturePurpose returns a boolean if a field has been set.

### SetExpectedSignaturePurposeNil

`func (o *TrustSigningReadinessDto) SetExpectedSignaturePurposeNil(b bool)`

 SetExpectedSignaturePurposeNil sets the value for ExpectedSignaturePurpose to be an explicit nil

### UnsetExpectedSignaturePurpose
`func (o *TrustSigningReadinessDto) UnsetExpectedSignaturePurpose()`

UnsetExpectedSignaturePurpose ensures that no value is present for ExpectedSignaturePurpose, not even an explicit nil
### GetExpectedDigestAlgorithm

`func (o *TrustSigningReadinessDto) GetExpectedDigestAlgorithm() string`

GetExpectedDigestAlgorithm returns the ExpectedDigestAlgorithm field if non-nil, zero value otherwise.

### GetExpectedDigestAlgorithmOk

`func (o *TrustSigningReadinessDto) GetExpectedDigestAlgorithmOk() (*string, bool)`

GetExpectedDigestAlgorithmOk returns a tuple with the ExpectedDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDigestAlgorithm

`func (o *TrustSigningReadinessDto) SetExpectedDigestAlgorithm(v string)`

SetExpectedDigestAlgorithm sets ExpectedDigestAlgorithm field to given value.

### HasExpectedDigestAlgorithm

`func (o *TrustSigningReadinessDto) HasExpectedDigestAlgorithm() bool`

HasExpectedDigestAlgorithm returns a boolean if a field has been set.

### SetExpectedDigestAlgorithmNil

`func (o *TrustSigningReadinessDto) SetExpectedDigestAlgorithmNil(b bool)`

 SetExpectedDigestAlgorithmNil sets the value for ExpectedDigestAlgorithm to be an explicit nil

### UnsetExpectedDigestAlgorithm
`func (o *TrustSigningReadinessDto) UnsetExpectedDigestAlgorithm()`

UnsetExpectedDigestAlgorithm ensures that no value is present for ExpectedDigestAlgorithm, not even an explicit nil
### GetExpectedSignatureAlgorithm

`func (o *TrustSigningReadinessDto) GetExpectedSignatureAlgorithm() string`

GetExpectedSignatureAlgorithm returns the ExpectedSignatureAlgorithm field if non-nil, zero value otherwise.

### GetExpectedSignatureAlgorithmOk

`func (o *TrustSigningReadinessDto) GetExpectedSignatureAlgorithmOk() (*string, bool)`

GetExpectedSignatureAlgorithmOk returns a tuple with the ExpectedSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSignatureAlgorithm

`func (o *TrustSigningReadinessDto) SetExpectedSignatureAlgorithm(v string)`

SetExpectedSignatureAlgorithm sets ExpectedSignatureAlgorithm field to given value.

### HasExpectedSignatureAlgorithm

`func (o *TrustSigningReadinessDto) HasExpectedSignatureAlgorithm() bool`

HasExpectedSignatureAlgorithm returns a boolean if a field has been set.

### SetExpectedSignatureAlgorithmNil

`func (o *TrustSigningReadinessDto) SetExpectedSignatureAlgorithmNil(b bool)`

 SetExpectedSignatureAlgorithmNil sets the value for ExpectedSignatureAlgorithm to be an explicit nil

### UnsetExpectedSignatureAlgorithm
`func (o *TrustSigningReadinessDto) UnsetExpectedSignatureAlgorithm()`

UnsetExpectedSignatureAlgorithm ensures that no value is present for ExpectedSignatureAlgorithm, not even an explicit nil
### GetExpectedCanonicalizationAlgorithm

`func (o *TrustSigningReadinessDto) GetExpectedCanonicalizationAlgorithm() string`

GetExpectedCanonicalizationAlgorithm returns the ExpectedCanonicalizationAlgorithm field if non-nil, zero value otherwise.

### GetExpectedCanonicalizationAlgorithmOk

`func (o *TrustSigningReadinessDto) GetExpectedCanonicalizationAlgorithmOk() (*string, bool)`

GetExpectedCanonicalizationAlgorithmOk returns a tuple with the ExpectedCanonicalizationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCanonicalizationAlgorithm

`func (o *TrustSigningReadinessDto) SetExpectedCanonicalizationAlgorithm(v string)`

SetExpectedCanonicalizationAlgorithm sets ExpectedCanonicalizationAlgorithm field to given value.

### HasExpectedCanonicalizationAlgorithm

`func (o *TrustSigningReadinessDto) HasExpectedCanonicalizationAlgorithm() bool`

HasExpectedCanonicalizationAlgorithm returns a boolean if a field has been set.

### SetExpectedCanonicalizationAlgorithmNil

`func (o *TrustSigningReadinessDto) SetExpectedCanonicalizationAlgorithmNil(b bool)`

 SetExpectedCanonicalizationAlgorithmNil sets the value for ExpectedCanonicalizationAlgorithm to be an explicit nil

### UnsetExpectedCanonicalizationAlgorithm
`func (o *TrustSigningReadinessDto) UnsetExpectedCanonicalizationAlgorithm()`

UnsetExpectedCanonicalizationAlgorithm ensures that no value is present for ExpectedCanonicalizationAlgorithm, not even an explicit nil
### GetPolicyIdentifier

`func (o *TrustSigningReadinessDto) GetPolicyIdentifier() string`

GetPolicyIdentifier returns the PolicyIdentifier field if non-nil, zero value otherwise.

### GetPolicyIdentifierOk

`func (o *TrustSigningReadinessDto) GetPolicyIdentifierOk() (*string, bool)`

GetPolicyIdentifierOk returns a tuple with the PolicyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifier

`func (o *TrustSigningReadinessDto) SetPolicyIdentifier(v string)`

SetPolicyIdentifier sets PolicyIdentifier field to given value.

### HasPolicyIdentifier

`func (o *TrustSigningReadinessDto) HasPolicyIdentifier() bool`

HasPolicyIdentifier returns a boolean if a field has been set.

### SetPolicyIdentifierNil

`func (o *TrustSigningReadinessDto) SetPolicyIdentifierNil(b bool)`

 SetPolicyIdentifierNil sets the value for PolicyIdentifier to be an explicit nil

### UnsetPolicyIdentifier
`func (o *TrustSigningReadinessDto) UnsetPolicyIdentifier()`

UnsetPolicyIdentifier ensures that no value is present for PolicyIdentifier, not even an explicit nil
### GetAuthorityProfile

`func (o *TrustSigningReadinessDto) GetAuthorityProfile() string`

GetAuthorityProfile returns the AuthorityProfile field if non-nil, zero value otherwise.

### GetAuthorityProfileOk

`func (o *TrustSigningReadinessDto) GetAuthorityProfileOk() (*string, bool)`

GetAuthorityProfileOk returns a tuple with the AuthorityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityProfile

`func (o *TrustSigningReadinessDto) SetAuthorityProfile(v string)`

SetAuthorityProfile sets AuthorityProfile field to given value.

### HasAuthorityProfile

`func (o *TrustSigningReadinessDto) HasAuthorityProfile() bool`

HasAuthorityProfile returns a boolean if a field has been set.

### SetAuthorityProfileNil

`func (o *TrustSigningReadinessDto) SetAuthorityProfileNil(b bool)`

 SetAuthorityProfileNil sets the value for AuthorityProfile to be an explicit nil

### UnsetAuthorityProfile
`func (o *TrustSigningReadinessDto) UnsetAuthorityProfile()`

UnsetAuthorityProfile ensures that no value is present for AuthorityProfile, not even an explicit nil
### GetRequiresCustodyProvider

`func (o *TrustSigningReadinessDto) GetRequiresCustodyProvider() bool`

GetRequiresCustodyProvider returns the RequiresCustodyProvider field if non-nil, zero value otherwise.

### GetRequiresCustodyProviderOk

`func (o *TrustSigningReadinessDto) GetRequiresCustodyProviderOk() (*bool, bool)`

GetRequiresCustodyProviderOk returns a tuple with the RequiresCustodyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCustodyProvider

`func (o *TrustSigningReadinessDto) SetRequiresCustodyProvider(v bool)`

SetRequiresCustodyProvider sets RequiresCustodyProvider field to given value.

### HasRequiresCustodyProvider

`func (o *TrustSigningReadinessDto) HasRequiresCustodyProvider() bool`

HasRequiresCustodyProvider returns a boolean if a field has been set.

### GetRequiresSourceArtifact

`func (o *TrustSigningReadinessDto) GetRequiresSourceArtifact() bool`

GetRequiresSourceArtifact returns the RequiresSourceArtifact field if non-nil, zero value otherwise.

### GetRequiresSourceArtifactOk

`func (o *TrustSigningReadinessDto) GetRequiresSourceArtifactOk() (*bool, bool)`

GetRequiresSourceArtifactOk returns a tuple with the RequiresSourceArtifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSourceArtifact

`func (o *TrustSigningReadinessDto) SetRequiresSourceArtifact(v bool)`

SetRequiresSourceArtifact sets RequiresSourceArtifact field to given value.

### HasRequiresSourceArtifact

`func (o *TrustSigningReadinessDto) HasRequiresSourceArtifact() bool`

HasRequiresSourceArtifact returns a boolean if a field has been set.

### GetRequiresCertificate

`func (o *TrustSigningReadinessDto) GetRequiresCertificate() bool`

GetRequiresCertificate returns the RequiresCertificate field if non-nil, zero value otherwise.

### GetRequiresCertificateOk

`func (o *TrustSigningReadinessDto) GetRequiresCertificateOk() (*bool, bool)`

GetRequiresCertificateOk returns a tuple with the RequiresCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCertificate

`func (o *TrustSigningReadinessDto) SetRequiresCertificate(v bool)`

SetRequiresCertificate sets RequiresCertificate field to given value.

### HasRequiresCertificate

`func (o *TrustSigningReadinessDto) HasRequiresCertificate() bool`

HasRequiresCertificate returns a boolean if a field has been set.

### GetCorrelationId

`func (o *TrustSigningReadinessDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TrustSigningReadinessDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TrustSigningReadinessDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TrustSigningReadinessDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *TrustSigningReadinessDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *TrustSigningReadinessDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


