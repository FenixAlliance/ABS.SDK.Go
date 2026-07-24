# SigningLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**LogType** | Pointer to **string** |  | [optional] 
**SecurityEvent** | Pointer to **NullableString** |  | [optional] 
**RequiresAttention** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateId** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentId** | Pointer to **NullableString** |  | [optional] 
**OperationType** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**InputHash** | Pointer to **NullableString** |  | [optional] 
**OutputHash** | Pointer to **NullableString** |  | [optional] 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**ResultCode** | Pointer to **NullableString** |  | [optional] 
**SigningProfileDisplayName** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateTitle** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentTitle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningLogDto

`func NewSigningLogDto() *SigningLogDto`

NewSigningLogDto instantiates a new SigningLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningLogDtoWithDefaults

`func NewSigningLogDtoWithDefaults() *SigningLogDto`

NewSigningLogDtoWithDefaults instantiates a new SigningLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SigningLogDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningLogDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningLogDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningLogDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SigningLogDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SigningLogDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetType

`func (o *SigningLogDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningLogDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningLogDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningLogDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningLogDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningLogDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *SigningLogDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SigningLogDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SigningLogDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SigningLogDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SigningLogDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SigningLogDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLogType

`func (o *SigningLogDto) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *SigningLogDto) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *SigningLogDto) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *SigningLogDto) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetSecurityEvent

`func (o *SigningLogDto) GetSecurityEvent() string`

GetSecurityEvent returns the SecurityEvent field if non-nil, zero value otherwise.

### GetSecurityEventOk

`func (o *SigningLogDto) GetSecurityEventOk() (*string, bool)`

GetSecurityEventOk returns a tuple with the SecurityEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEvent

`func (o *SigningLogDto) SetSecurityEvent(v string)`

SetSecurityEvent sets SecurityEvent field to given value.

### HasSecurityEvent

`func (o *SigningLogDto) HasSecurityEvent() bool`

HasSecurityEvent returns a boolean if a field has been set.

### SetSecurityEventNil

`func (o *SigningLogDto) SetSecurityEventNil(b bool)`

 SetSecurityEventNil sets the value for SecurityEvent to be an explicit nil

### UnsetSecurityEvent
`func (o *SigningLogDto) UnsetSecurityEvent()`

UnsetSecurityEvent ensures that no value is present for SecurityEvent, not even an explicit nil
### GetRequiresAttention

`func (o *SigningLogDto) GetRequiresAttention() bool`

GetRequiresAttention returns the RequiresAttention field if non-nil, zero value otherwise.

### GetRequiresAttentionOk

`func (o *SigningLogDto) GetRequiresAttentionOk() (*bool, bool)`

GetRequiresAttentionOk returns a tuple with the RequiresAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAttention

`func (o *SigningLogDto) SetRequiresAttention(v bool)`

SetRequiresAttention sets RequiresAttention field to given value.

### HasRequiresAttention

`func (o *SigningLogDto) HasRequiresAttention() bool`

HasRequiresAttention returns a boolean if a field has been set.

### GetTenantId

`func (o *SigningLogDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningLogDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningLogDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningLogDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningLogDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningLogDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *SigningLogDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SigningLogDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SigningLogDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SigningLogDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SigningLogDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SigningLogDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEnrollmentId

`func (o *SigningLogDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SigningLogDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SigningLogDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SigningLogDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SigningLogDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SigningLogDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetContactId

`func (o *SigningLogDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningLogDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningLogDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SigningLogDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SigningLogDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SigningLogDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetSigningProfileId

`func (o *SigningLogDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *SigningLogDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *SigningLogDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *SigningLogDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *SigningLogDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *SigningLogDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil
### GetSigningCertificateId

`func (o *SigningLogDto) GetSigningCertificateId() string`

GetSigningCertificateId returns the SigningCertificateId field if non-nil, zero value otherwise.

### GetSigningCertificateIdOk

`func (o *SigningLogDto) GetSigningCertificateIdOk() (*string, bool)`

GetSigningCertificateIdOk returns a tuple with the SigningCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateId

`func (o *SigningLogDto) SetSigningCertificateId(v string)`

SetSigningCertificateId sets SigningCertificateId field to given value.

### HasSigningCertificateId

`func (o *SigningLogDto) HasSigningCertificateId() bool`

HasSigningCertificateId returns a boolean if a field has been set.

### SetSigningCertificateIdNil

`func (o *SigningLogDto) SetSigningCertificateIdNil(b bool)`

 SetSigningCertificateIdNil sets the value for SigningCertificateId to be an explicit nil

### UnsetSigningCertificateId
`func (o *SigningLogDto) UnsetSigningCertificateId()`

UnsetSigningCertificateId ensures that no value is present for SigningCertificateId, not even an explicit nil
### GetSignedDocumentId

`func (o *SigningLogDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *SigningLogDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *SigningLogDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.

### HasSignedDocumentId

`func (o *SigningLogDto) HasSignedDocumentId() bool`

HasSignedDocumentId returns a boolean if a field has been set.

### SetSignedDocumentIdNil

`func (o *SigningLogDto) SetSignedDocumentIdNil(b bool)`

 SetSignedDocumentIdNil sets the value for SignedDocumentId to be an explicit nil

### UnsetSignedDocumentId
`func (o *SigningLogDto) UnsetSignedDocumentId()`

UnsetSignedDocumentId ensures that no value is present for SignedDocumentId, not even an explicit nil
### GetOperationType

`func (o *SigningLogDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SigningLogDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SigningLogDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *SigningLogDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *SigningLogDto) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *SigningLogDto) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetCorrelationId

`func (o *SigningLogDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SigningLogDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SigningLogDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SigningLogDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SigningLogDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SigningLogDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetInputHash

`func (o *SigningLogDto) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *SigningLogDto) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *SigningLogDto) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *SigningLogDto) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### SetInputHashNil

`func (o *SigningLogDto) SetInputHashNil(b bool)`

 SetInputHashNil sets the value for InputHash to be an explicit nil

### UnsetInputHash
`func (o *SigningLogDto) UnsetInputHash()`

UnsetInputHash ensures that no value is present for InputHash, not even an explicit nil
### GetOutputHash

`func (o *SigningLogDto) GetOutputHash() string`

GetOutputHash returns the OutputHash field if non-nil, zero value otherwise.

### GetOutputHashOk

`func (o *SigningLogDto) GetOutputHashOk() (*string, bool)`

GetOutputHashOk returns a tuple with the OutputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputHash

`func (o *SigningLogDto) SetOutputHash(v string)`

SetOutputHash sets OutputHash field to given value.

### HasOutputHash

`func (o *SigningLogDto) HasOutputHash() bool`

HasOutputHash returns a boolean if a field has been set.

### SetOutputHashNil

`func (o *SigningLogDto) SetOutputHashNil(b bool)`

 SetOutputHashNil sets the value for OutputHash to be an explicit nil

### UnsetOutputHash
`func (o *SigningLogDto) UnsetOutputHash()`

UnsetOutputHash ensures that no value is present for OutputHash, not even an explicit nil
### GetProviderName

`func (o *SigningLogDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *SigningLogDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *SigningLogDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *SigningLogDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *SigningLogDto) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *SigningLogDto) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetResultCode

`func (o *SigningLogDto) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *SigningLogDto) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *SigningLogDto) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *SigningLogDto) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### SetResultCodeNil

`func (o *SigningLogDto) SetResultCodeNil(b bool)`

 SetResultCodeNil sets the value for ResultCode to be an explicit nil

### UnsetResultCode
`func (o *SigningLogDto) UnsetResultCode()`

UnsetResultCode ensures that no value is present for ResultCode, not even an explicit nil
### GetSigningProfileDisplayName

`func (o *SigningLogDto) GetSigningProfileDisplayName() string`

GetSigningProfileDisplayName returns the SigningProfileDisplayName field if non-nil, zero value otherwise.

### GetSigningProfileDisplayNameOk

`func (o *SigningLogDto) GetSigningProfileDisplayNameOk() (*string, bool)`

GetSigningProfileDisplayNameOk returns a tuple with the SigningProfileDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileDisplayName

`func (o *SigningLogDto) SetSigningProfileDisplayName(v string)`

SetSigningProfileDisplayName sets SigningProfileDisplayName field to given value.

### HasSigningProfileDisplayName

`func (o *SigningLogDto) HasSigningProfileDisplayName() bool`

HasSigningProfileDisplayName returns a boolean if a field has been set.

### SetSigningProfileDisplayNameNil

`func (o *SigningLogDto) SetSigningProfileDisplayNameNil(b bool)`

 SetSigningProfileDisplayNameNil sets the value for SigningProfileDisplayName to be an explicit nil

### UnsetSigningProfileDisplayName
`func (o *SigningLogDto) UnsetSigningProfileDisplayName()`

UnsetSigningProfileDisplayName ensures that no value is present for SigningProfileDisplayName, not even an explicit nil
### GetSigningCertificateTitle

`func (o *SigningLogDto) GetSigningCertificateTitle() string`

GetSigningCertificateTitle returns the SigningCertificateTitle field if non-nil, zero value otherwise.

### GetSigningCertificateTitleOk

`func (o *SigningLogDto) GetSigningCertificateTitleOk() (*string, bool)`

GetSigningCertificateTitleOk returns a tuple with the SigningCertificateTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateTitle

`func (o *SigningLogDto) SetSigningCertificateTitle(v string)`

SetSigningCertificateTitle sets SigningCertificateTitle field to given value.

### HasSigningCertificateTitle

`func (o *SigningLogDto) HasSigningCertificateTitle() bool`

HasSigningCertificateTitle returns a boolean if a field has been set.

### SetSigningCertificateTitleNil

`func (o *SigningLogDto) SetSigningCertificateTitleNil(b bool)`

 SetSigningCertificateTitleNil sets the value for SigningCertificateTitle to be an explicit nil

### UnsetSigningCertificateTitle
`func (o *SigningLogDto) UnsetSigningCertificateTitle()`

UnsetSigningCertificateTitle ensures that no value is present for SigningCertificateTitle, not even an explicit nil
### GetSignedDocumentTitle

`func (o *SigningLogDto) GetSignedDocumentTitle() string`

GetSignedDocumentTitle returns the SignedDocumentTitle field if non-nil, zero value otherwise.

### GetSignedDocumentTitleOk

`func (o *SigningLogDto) GetSignedDocumentTitleOk() (*string, bool)`

GetSignedDocumentTitleOk returns a tuple with the SignedDocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentTitle

`func (o *SigningLogDto) SetSignedDocumentTitle(v string)`

SetSignedDocumentTitle sets SignedDocumentTitle field to given value.

### HasSignedDocumentTitle

`func (o *SigningLogDto) HasSignedDocumentTitle() bool`

HasSignedDocumentTitle returns a boolean if a field has been set.

### SetSignedDocumentTitleNil

`func (o *SigningLogDto) SetSignedDocumentTitleNil(b bool)`

 SetSignedDocumentTitleNil sets the value for SignedDocumentTitle to be an explicit nil

### UnsetSignedDocumentTitle
`func (o *SigningLogDto) UnsetSignedDocumentTitle()`

UnsetSignedDocumentTitle ensures that no value is present for SignedDocumentTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


