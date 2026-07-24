# SigningRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentId** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentTitle** | Pointer to **NullableString** |  | [optional] 
**FrozenSourceFileUploadId** | Pointer to **NullableString** |  | [optional] 
**SourceSha256** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RoutingMode** | Pointer to **string** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**SentAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CompletedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**VoidedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**VoidedReason** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningRequestDto

`func NewSigningRequestDto() *SigningRequestDto`

NewSigningRequestDto instantiates a new SigningRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningRequestDtoWithDefaults

`func NewSigningRequestDtoWithDefaults() *SigningRequestDto`

NewSigningRequestDtoWithDefaults instantiates a new SigningRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningRequestDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningRequestDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningRequestDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *SigningRequestDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningRequestDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningRequestDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningRequestDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningRequestDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningRequestDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSignedDocumentId

`func (o *SigningRequestDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *SigningRequestDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *SigningRequestDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.

### HasSignedDocumentId

`func (o *SigningRequestDto) HasSignedDocumentId() bool`

HasSignedDocumentId returns a boolean if a field has been set.

### SetSignedDocumentIdNil

`func (o *SigningRequestDto) SetSignedDocumentIdNil(b bool)`

 SetSignedDocumentIdNil sets the value for SignedDocumentId to be an explicit nil

### UnsetSignedDocumentId
`func (o *SigningRequestDto) UnsetSignedDocumentId()`

UnsetSignedDocumentId ensures that no value is present for SignedDocumentId, not even an explicit nil
### GetSignedDocumentTitle

`func (o *SigningRequestDto) GetSignedDocumentTitle() string`

GetSignedDocumentTitle returns the SignedDocumentTitle field if non-nil, zero value otherwise.

### GetSignedDocumentTitleOk

`func (o *SigningRequestDto) GetSignedDocumentTitleOk() (*string, bool)`

GetSignedDocumentTitleOk returns a tuple with the SignedDocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentTitle

`func (o *SigningRequestDto) SetSignedDocumentTitle(v string)`

SetSignedDocumentTitle sets SignedDocumentTitle field to given value.

### HasSignedDocumentTitle

`func (o *SigningRequestDto) HasSignedDocumentTitle() bool`

HasSignedDocumentTitle returns a boolean if a field has been set.

### SetSignedDocumentTitleNil

`func (o *SigningRequestDto) SetSignedDocumentTitleNil(b bool)`

 SetSignedDocumentTitleNil sets the value for SignedDocumentTitle to be an explicit nil

### UnsetSignedDocumentTitle
`func (o *SigningRequestDto) UnsetSignedDocumentTitle()`

UnsetSignedDocumentTitle ensures that no value is present for SignedDocumentTitle, not even an explicit nil
### GetFrozenSourceFileUploadId

`func (o *SigningRequestDto) GetFrozenSourceFileUploadId() string`

GetFrozenSourceFileUploadId returns the FrozenSourceFileUploadId field if non-nil, zero value otherwise.

### GetFrozenSourceFileUploadIdOk

`func (o *SigningRequestDto) GetFrozenSourceFileUploadIdOk() (*string, bool)`

GetFrozenSourceFileUploadIdOk returns a tuple with the FrozenSourceFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenSourceFileUploadId

`func (o *SigningRequestDto) SetFrozenSourceFileUploadId(v string)`

SetFrozenSourceFileUploadId sets FrozenSourceFileUploadId field to given value.

### HasFrozenSourceFileUploadId

`func (o *SigningRequestDto) HasFrozenSourceFileUploadId() bool`

HasFrozenSourceFileUploadId returns a boolean if a field has been set.

### SetFrozenSourceFileUploadIdNil

`func (o *SigningRequestDto) SetFrozenSourceFileUploadIdNil(b bool)`

 SetFrozenSourceFileUploadIdNil sets the value for FrozenSourceFileUploadId to be an explicit nil

### UnsetFrozenSourceFileUploadId
`func (o *SigningRequestDto) UnsetFrozenSourceFileUploadId()`

UnsetFrozenSourceFileUploadId ensures that no value is present for FrozenSourceFileUploadId, not even an explicit nil
### GetSourceSha256

`func (o *SigningRequestDto) GetSourceSha256() string`

GetSourceSha256 returns the SourceSha256 field if non-nil, zero value otherwise.

### GetSourceSha256Ok

`func (o *SigningRequestDto) GetSourceSha256Ok() (*string, bool)`

GetSourceSha256Ok returns a tuple with the SourceSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSha256

`func (o *SigningRequestDto) SetSourceSha256(v string)`

SetSourceSha256 sets SourceSha256 field to given value.

### HasSourceSha256

`func (o *SigningRequestDto) HasSourceSha256() bool`

HasSourceSha256 returns a boolean if a field has been set.

### SetSourceSha256Nil

`func (o *SigningRequestDto) SetSourceSha256Nil(b bool)`

 SetSourceSha256Nil sets the value for SourceSha256 to be an explicit nil

### UnsetSourceSha256
`func (o *SigningRequestDto) UnsetSourceSha256()`

UnsetSourceSha256 ensures that no value is present for SourceSha256, not even an explicit nil
### GetStatus

`func (o *SigningRequestDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SigningRequestDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SigningRequestDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SigningRequestDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRoutingMode

`func (o *SigningRequestDto) GetRoutingMode() string`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *SigningRequestDto) GetRoutingModeOk() (*string, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *SigningRequestDto) SetRoutingMode(v string)`

SetRoutingMode sets RoutingMode field to given value.

### HasRoutingMode

`func (o *SigningRequestDto) HasRoutingMode() bool`

HasRoutingMode returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *SigningRequestDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *SigningRequestDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *SigningRequestDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *SigningRequestDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetSentAtUtc

`func (o *SigningRequestDto) GetSentAtUtc() time.Time`

GetSentAtUtc returns the SentAtUtc field if non-nil, zero value otherwise.

### GetSentAtUtcOk

`func (o *SigningRequestDto) GetSentAtUtcOk() (*time.Time, bool)`

GetSentAtUtcOk returns a tuple with the SentAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAtUtc

`func (o *SigningRequestDto) SetSentAtUtc(v time.Time)`

SetSentAtUtc sets SentAtUtc field to given value.

### HasSentAtUtc

`func (o *SigningRequestDto) HasSentAtUtc() bool`

HasSentAtUtc returns a boolean if a field has been set.

### SetSentAtUtcNil

`func (o *SigningRequestDto) SetSentAtUtcNil(b bool)`

 SetSentAtUtcNil sets the value for SentAtUtc to be an explicit nil

### UnsetSentAtUtc
`func (o *SigningRequestDto) UnsetSentAtUtc()`

UnsetSentAtUtc ensures that no value is present for SentAtUtc, not even an explicit nil
### GetCompletedAtUtc

`func (o *SigningRequestDto) GetCompletedAtUtc() time.Time`

GetCompletedAtUtc returns the CompletedAtUtc field if non-nil, zero value otherwise.

### GetCompletedAtUtcOk

`func (o *SigningRequestDto) GetCompletedAtUtcOk() (*time.Time, bool)`

GetCompletedAtUtcOk returns a tuple with the CompletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtUtc

`func (o *SigningRequestDto) SetCompletedAtUtc(v time.Time)`

SetCompletedAtUtc sets CompletedAtUtc field to given value.

### HasCompletedAtUtc

`func (o *SigningRequestDto) HasCompletedAtUtc() bool`

HasCompletedAtUtc returns a boolean if a field has been set.

### SetCompletedAtUtcNil

`func (o *SigningRequestDto) SetCompletedAtUtcNil(b bool)`

 SetCompletedAtUtcNil sets the value for CompletedAtUtc to be an explicit nil

### UnsetCompletedAtUtc
`func (o *SigningRequestDto) UnsetCompletedAtUtc()`

UnsetCompletedAtUtc ensures that no value is present for CompletedAtUtc, not even an explicit nil
### GetExpiresAtUtc

`func (o *SigningRequestDto) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *SigningRequestDto) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *SigningRequestDto) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *SigningRequestDto) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *SigningRequestDto) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *SigningRequestDto) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetVoidedAtUtc

`func (o *SigningRequestDto) GetVoidedAtUtc() time.Time`

GetVoidedAtUtc returns the VoidedAtUtc field if non-nil, zero value otherwise.

### GetVoidedAtUtcOk

`func (o *SigningRequestDto) GetVoidedAtUtcOk() (*time.Time, bool)`

GetVoidedAtUtcOk returns a tuple with the VoidedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAtUtc

`func (o *SigningRequestDto) SetVoidedAtUtc(v time.Time)`

SetVoidedAtUtc sets VoidedAtUtc field to given value.

### HasVoidedAtUtc

`func (o *SigningRequestDto) HasVoidedAtUtc() bool`

HasVoidedAtUtc returns a boolean if a field has been set.

### SetVoidedAtUtcNil

`func (o *SigningRequestDto) SetVoidedAtUtcNil(b bool)`

 SetVoidedAtUtcNil sets the value for VoidedAtUtc to be an explicit nil

### UnsetVoidedAtUtc
`func (o *SigningRequestDto) UnsetVoidedAtUtc()`

UnsetVoidedAtUtc ensures that no value is present for VoidedAtUtc, not even an explicit nil
### GetVoidedReason

`func (o *SigningRequestDto) GetVoidedReason() string`

GetVoidedReason returns the VoidedReason field if non-nil, zero value otherwise.

### GetVoidedReasonOk

`func (o *SigningRequestDto) GetVoidedReasonOk() (*string, bool)`

GetVoidedReasonOk returns a tuple with the VoidedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedReason

`func (o *SigningRequestDto) SetVoidedReason(v string)`

SetVoidedReason sets VoidedReason field to given value.

### HasVoidedReason

`func (o *SigningRequestDto) HasVoidedReason() bool`

HasVoidedReason returns a boolean if a field has been set.

### SetVoidedReasonNil

`func (o *SigningRequestDto) SetVoidedReasonNil(b bool)`

 SetVoidedReasonNil sets the value for VoidedReason to be an explicit nil

### UnsetVoidedReason
`func (o *SigningRequestDto) UnsetVoidedReason()`

UnsetVoidedReason ensures that no value is present for VoidedReason, not even an explicit nil
### GetMessage

`func (o *SigningRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SigningRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SigningRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SigningRequestDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SigningRequestDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SigningRequestDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCorrelationId

`func (o *SigningRequestDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SigningRequestDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SigningRequestDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SigningRequestDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SigningRequestDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SigningRequestDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *SigningRequestDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *SigningRequestDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *SigningRequestDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *SigningRequestDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *SigningRequestDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *SigningRequestDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


