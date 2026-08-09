# InboxMessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**SourceSystem** | Pointer to **NullableString** |  | [optional] 
**SourceRegistrationId** | Pointer to **NullableString** |  | [optional] 
**ExternalMessageId** | Pointer to **NullableString** |  | [optional] 
**DeduplicationKey** | Pointer to **NullableString** |  | [optional] 
**DeduplicationSignature** | Pointer to **NullableString** |  | [optional] 
**PayloadDigest** | Pointer to **NullableString** |  | [optional] 
**DeliveryCount** | Pointer to **int32** |  | [optional] 
**LastDuplicateReceivedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**MessageType** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Attempts** | Pointer to **int32** |  | [optional] 
**MaxAttempts** | Pointer to **int32** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**VerificationProfile** | Pointer to **NullableString** |  | [optional] 
**VerificationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**VerifiedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 
**ReplayCount** | Pointer to **int32** |  | [optional] 
**OriginalInboxMessageId** | Pointer to **NullableString** |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**CausationId** | Pointer to **NullableString** |  | [optional] 
**LockedBy** | Pointer to **NullableString** |  | [optional] 
**LockedUntilUtc** | Pointer to **NullableTime** |  | [optional] 
**AvailableAtUtc** | Pointer to **time.Time** |  | [optional] 
**ReceivedAtUtc** | Pointer to **time.Time** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**LastAttemptAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ProcessedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**FailedAtUtc** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewInboxMessageDto

`func NewInboxMessageDto() *InboxMessageDto`

NewInboxMessageDto instantiates a new InboxMessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxMessageDtoWithDefaults

`func NewInboxMessageDtoWithDefaults() *InboxMessageDto`

NewInboxMessageDtoWithDefaults instantiates a new InboxMessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboxMessageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxMessageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxMessageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InboxMessageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InboxMessageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InboxMessageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InboxMessageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InboxMessageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InboxMessageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InboxMessageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InboxMessageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InboxMessageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *InboxMessageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InboxMessageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InboxMessageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InboxMessageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InboxMessageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InboxMessageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSourceSystem

`func (o *InboxMessageDto) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *InboxMessageDto) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *InboxMessageDto) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.

### HasSourceSystem

`func (o *InboxMessageDto) HasSourceSystem() bool`

HasSourceSystem returns a boolean if a field has been set.

### SetSourceSystemNil

`func (o *InboxMessageDto) SetSourceSystemNil(b bool)`

 SetSourceSystemNil sets the value for SourceSystem to be an explicit nil

### UnsetSourceSystem
`func (o *InboxMessageDto) UnsetSourceSystem()`

UnsetSourceSystem ensures that no value is present for SourceSystem, not even an explicit nil
### GetSourceRegistrationId

`func (o *InboxMessageDto) GetSourceRegistrationId() string`

GetSourceRegistrationId returns the SourceRegistrationId field if non-nil, zero value otherwise.

### GetSourceRegistrationIdOk

`func (o *InboxMessageDto) GetSourceRegistrationIdOk() (*string, bool)`

GetSourceRegistrationIdOk returns a tuple with the SourceRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegistrationId

`func (o *InboxMessageDto) SetSourceRegistrationId(v string)`

SetSourceRegistrationId sets SourceRegistrationId field to given value.

### HasSourceRegistrationId

`func (o *InboxMessageDto) HasSourceRegistrationId() bool`

HasSourceRegistrationId returns a boolean if a field has been set.

### SetSourceRegistrationIdNil

`func (o *InboxMessageDto) SetSourceRegistrationIdNil(b bool)`

 SetSourceRegistrationIdNil sets the value for SourceRegistrationId to be an explicit nil

### UnsetSourceRegistrationId
`func (o *InboxMessageDto) UnsetSourceRegistrationId()`

UnsetSourceRegistrationId ensures that no value is present for SourceRegistrationId, not even an explicit nil
### GetExternalMessageId

`func (o *InboxMessageDto) GetExternalMessageId() string`

GetExternalMessageId returns the ExternalMessageId field if non-nil, zero value otherwise.

### GetExternalMessageIdOk

`func (o *InboxMessageDto) GetExternalMessageIdOk() (*string, bool)`

GetExternalMessageIdOk returns a tuple with the ExternalMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMessageId

`func (o *InboxMessageDto) SetExternalMessageId(v string)`

SetExternalMessageId sets ExternalMessageId field to given value.

### HasExternalMessageId

`func (o *InboxMessageDto) HasExternalMessageId() bool`

HasExternalMessageId returns a boolean if a field has been set.

### SetExternalMessageIdNil

`func (o *InboxMessageDto) SetExternalMessageIdNil(b bool)`

 SetExternalMessageIdNil sets the value for ExternalMessageId to be an explicit nil

### UnsetExternalMessageId
`func (o *InboxMessageDto) UnsetExternalMessageId()`

UnsetExternalMessageId ensures that no value is present for ExternalMessageId, not even an explicit nil
### GetDeduplicationKey

`func (o *InboxMessageDto) GetDeduplicationKey() string`

GetDeduplicationKey returns the DeduplicationKey field if non-nil, zero value otherwise.

### GetDeduplicationKeyOk

`func (o *InboxMessageDto) GetDeduplicationKeyOk() (*string, bool)`

GetDeduplicationKeyOk returns a tuple with the DeduplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationKey

`func (o *InboxMessageDto) SetDeduplicationKey(v string)`

SetDeduplicationKey sets DeduplicationKey field to given value.

### HasDeduplicationKey

`func (o *InboxMessageDto) HasDeduplicationKey() bool`

HasDeduplicationKey returns a boolean if a field has been set.

### SetDeduplicationKeyNil

`func (o *InboxMessageDto) SetDeduplicationKeyNil(b bool)`

 SetDeduplicationKeyNil sets the value for DeduplicationKey to be an explicit nil

### UnsetDeduplicationKey
`func (o *InboxMessageDto) UnsetDeduplicationKey()`

UnsetDeduplicationKey ensures that no value is present for DeduplicationKey, not even an explicit nil
### GetDeduplicationSignature

`func (o *InboxMessageDto) GetDeduplicationSignature() string`

GetDeduplicationSignature returns the DeduplicationSignature field if non-nil, zero value otherwise.

### GetDeduplicationSignatureOk

`func (o *InboxMessageDto) GetDeduplicationSignatureOk() (*string, bool)`

GetDeduplicationSignatureOk returns a tuple with the DeduplicationSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationSignature

`func (o *InboxMessageDto) SetDeduplicationSignature(v string)`

SetDeduplicationSignature sets DeduplicationSignature field to given value.

### HasDeduplicationSignature

`func (o *InboxMessageDto) HasDeduplicationSignature() bool`

HasDeduplicationSignature returns a boolean if a field has been set.

### SetDeduplicationSignatureNil

`func (o *InboxMessageDto) SetDeduplicationSignatureNil(b bool)`

 SetDeduplicationSignatureNil sets the value for DeduplicationSignature to be an explicit nil

### UnsetDeduplicationSignature
`func (o *InboxMessageDto) UnsetDeduplicationSignature()`

UnsetDeduplicationSignature ensures that no value is present for DeduplicationSignature, not even an explicit nil
### GetPayloadDigest

`func (o *InboxMessageDto) GetPayloadDigest() string`

GetPayloadDigest returns the PayloadDigest field if non-nil, zero value otherwise.

### GetPayloadDigestOk

`func (o *InboxMessageDto) GetPayloadDigestOk() (*string, bool)`

GetPayloadDigestOk returns a tuple with the PayloadDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadDigest

`func (o *InboxMessageDto) SetPayloadDigest(v string)`

SetPayloadDigest sets PayloadDigest field to given value.

### HasPayloadDigest

`func (o *InboxMessageDto) HasPayloadDigest() bool`

HasPayloadDigest returns a boolean if a field has been set.

### SetPayloadDigestNil

`func (o *InboxMessageDto) SetPayloadDigestNil(b bool)`

 SetPayloadDigestNil sets the value for PayloadDigest to be an explicit nil

### UnsetPayloadDigest
`func (o *InboxMessageDto) UnsetPayloadDigest()`

UnsetPayloadDigest ensures that no value is present for PayloadDigest, not even an explicit nil
### GetDeliveryCount

`func (o *InboxMessageDto) GetDeliveryCount() int32`

GetDeliveryCount returns the DeliveryCount field if non-nil, zero value otherwise.

### GetDeliveryCountOk

`func (o *InboxMessageDto) GetDeliveryCountOk() (*int32, bool)`

GetDeliveryCountOk returns a tuple with the DeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCount

`func (o *InboxMessageDto) SetDeliveryCount(v int32)`

SetDeliveryCount sets DeliveryCount field to given value.

### HasDeliveryCount

`func (o *InboxMessageDto) HasDeliveryCount() bool`

HasDeliveryCount returns a boolean if a field has been set.

### GetLastDuplicateReceivedAtUtc

`func (o *InboxMessageDto) GetLastDuplicateReceivedAtUtc() time.Time`

GetLastDuplicateReceivedAtUtc returns the LastDuplicateReceivedAtUtc field if non-nil, zero value otherwise.

### GetLastDuplicateReceivedAtUtcOk

`func (o *InboxMessageDto) GetLastDuplicateReceivedAtUtcOk() (*time.Time, bool)`

GetLastDuplicateReceivedAtUtcOk returns a tuple with the LastDuplicateReceivedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDuplicateReceivedAtUtc

`func (o *InboxMessageDto) SetLastDuplicateReceivedAtUtc(v time.Time)`

SetLastDuplicateReceivedAtUtc sets LastDuplicateReceivedAtUtc field to given value.

### HasLastDuplicateReceivedAtUtc

`func (o *InboxMessageDto) HasLastDuplicateReceivedAtUtc() bool`

HasLastDuplicateReceivedAtUtc returns a boolean if a field has been set.

### SetLastDuplicateReceivedAtUtcNil

`func (o *InboxMessageDto) SetLastDuplicateReceivedAtUtcNil(b bool)`

 SetLastDuplicateReceivedAtUtcNil sets the value for LastDuplicateReceivedAtUtc to be an explicit nil

### UnsetLastDuplicateReceivedAtUtc
`func (o *InboxMessageDto) UnsetLastDuplicateReceivedAtUtc()`

UnsetLastDuplicateReceivedAtUtc ensures that no value is present for LastDuplicateReceivedAtUtc, not even an explicit nil
### GetMessageType

`func (o *InboxMessageDto) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *InboxMessageDto) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *InboxMessageDto) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *InboxMessageDto) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### SetMessageTypeNil

`func (o *InboxMessageDto) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *InboxMessageDto) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetVersion

`func (o *InboxMessageDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InboxMessageDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InboxMessageDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InboxMessageDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *InboxMessageDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *InboxMessageDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetContentType

`func (o *InboxMessageDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InboxMessageDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InboxMessageDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InboxMessageDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *InboxMessageDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *InboxMessageDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetStatus

`func (o *InboxMessageDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboxMessageDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboxMessageDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InboxMessageDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttempts

`func (o *InboxMessageDto) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *InboxMessageDto) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *InboxMessageDto) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *InboxMessageDto) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *InboxMessageDto) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *InboxMessageDto) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *InboxMessageDto) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *InboxMessageDto) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *InboxMessageDto) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *InboxMessageDto) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *InboxMessageDto) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *InboxMessageDto) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetVerificationProfile

`func (o *InboxMessageDto) GetVerificationProfile() string`

GetVerificationProfile returns the VerificationProfile field if non-nil, zero value otherwise.

### GetVerificationProfileOk

`func (o *InboxMessageDto) GetVerificationProfileOk() (*string, bool)`

GetVerificationProfileOk returns a tuple with the VerificationProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProfile

`func (o *InboxMessageDto) SetVerificationProfile(v string)`

SetVerificationProfile sets VerificationProfile field to given value.

### HasVerificationProfile

`func (o *InboxMessageDto) HasVerificationProfile() bool`

HasVerificationProfile returns a boolean if a field has been set.

### SetVerificationProfileNil

`func (o *InboxMessageDto) SetVerificationProfileNil(b bool)`

 SetVerificationProfileNil sets the value for VerificationProfile to be an explicit nil

### UnsetVerificationProfile
`func (o *InboxMessageDto) UnsetVerificationProfile()`

UnsetVerificationProfile ensures that no value is present for VerificationProfile, not even an explicit nil
### GetVerificationAlgorithm

`func (o *InboxMessageDto) GetVerificationAlgorithm() string`

GetVerificationAlgorithm returns the VerificationAlgorithm field if non-nil, zero value otherwise.

### GetVerificationAlgorithmOk

`func (o *InboxMessageDto) GetVerificationAlgorithmOk() (*string, bool)`

GetVerificationAlgorithmOk returns a tuple with the VerificationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAlgorithm

`func (o *InboxMessageDto) SetVerificationAlgorithm(v string)`

SetVerificationAlgorithm sets VerificationAlgorithm field to given value.

### HasVerificationAlgorithm

`func (o *InboxMessageDto) HasVerificationAlgorithm() bool`

HasVerificationAlgorithm returns a boolean if a field has been set.

### SetVerificationAlgorithmNil

`func (o *InboxMessageDto) SetVerificationAlgorithmNil(b bool)`

 SetVerificationAlgorithmNil sets the value for VerificationAlgorithm to be an explicit nil

### UnsetVerificationAlgorithm
`func (o *InboxMessageDto) UnsetVerificationAlgorithm()`

UnsetVerificationAlgorithm ensures that no value is present for VerificationAlgorithm, not even an explicit nil
### GetVerifiedAtUtc

`func (o *InboxMessageDto) GetVerifiedAtUtc() time.Time`

GetVerifiedAtUtc returns the VerifiedAtUtc field if non-nil, zero value otherwise.

### GetVerifiedAtUtcOk

`func (o *InboxMessageDto) GetVerifiedAtUtcOk() (*time.Time, bool)`

GetVerifiedAtUtcOk returns a tuple with the VerifiedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAtUtc

`func (o *InboxMessageDto) SetVerifiedAtUtc(v time.Time)`

SetVerifiedAtUtc sets VerifiedAtUtc field to given value.

### HasVerifiedAtUtc

`func (o *InboxMessageDto) HasVerifiedAtUtc() bool`

HasVerifiedAtUtc returns a boolean if a field has been set.

### SetVerifiedAtUtcNil

`func (o *InboxMessageDto) SetVerifiedAtUtcNil(b bool)`

 SetVerifiedAtUtcNil sets the value for VerifiedAtUtc to be an explicit nil

### UnsetVerifiedAtUtc
`func (o *InboxMessageDto) UnsetVerifiedAtUtc()`

UnsetVerifiedAtUtc ensures that no value is present for VerifiedAtUtc, not even an explicit nil
### GetGeneration

`func (o *InboxMessageDto) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *InboxMessageDto) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *InboxMessageDto) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *InboxMessageDto) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetReplayCount

`func (o *InboxMessageDto) GetReplayCount() int32`

GetReplayCount returns the ReplayCount field if non-nil, zero value otherwise.

### GetReplayCountOk

`func (o *InboxMessageDto) GetReplayCountOk() (*int32, bool)`

GetReplayCountOk returns a tuple with the ReplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayCount

`func (o *InboxMessageDto) SetReplayCount(v int32)`

SetReplayCount sets ReplayCount field to given value.

### HasReplayCount

`func (o *InboxMessageDto) HasReplayCount() bool`

HasReplayCount returns a boolean if a field has been set.

### GetOriginalInboxMessageId

`func (o *InboxMessageDto) GetOriginalInboxMessageId() string`

GetOriginalInboxMessageId returns the OriginalInboxMessageId field if non-nil, zero value otherwise.

### GetOriginalInboxMessageIdOk

`func (o *InboxMessageDto) GetOriginalInboxMessageIdOk() (*string, bool)`

GetOriginalInboxMessageIdOk returns a tuple with the OriginalInboxMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInboxMessageId

`func (o *InboxMessageDto) SetOriginalInboxMessageId(v string)`

SetOriginalInboxMessageId sets OriginalInboxMessageId field to given value.

### HasOriginalInboxMessageId

`func (o *InboxMessageDto) HasOriginalInboxMessageId() bool`

HasOriginalInboxMessageId returns a boolean if a field has been set.

### SetOriginalInboxMessageIdNil

`func (o *InboxMessageDto) SetOriginalInboxMessageIdNil(b bool)`

 SetOriginalInboxMessageIdNil sets the value for OriginalInboxMessageId to be an explicit nil

### UnsetOriginalInboxMessageId
`func (o *InboxMessageDto) UnsetOriginalInboxMessageId()`

UnsetOriginalInboxMessageId ensures that no value is present for OriginalInboxMessageId, not even an explicit nil
### GetFailureCode

`func (o *InboxMessageDto) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *InboxMessageDto) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *InboxMessageDto) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *InboxMessageDto) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *InboxMessageDto) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *InboxMessageDto) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureReason

`func (o *InboxMessageDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *InboxMessageDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *InboxMessageDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *InboxMessageDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *InboxMessageDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *InboxMessageDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetCorrelationId

`func (o *InboxMessageDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *InboxMessageDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *InboxMessageDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *InboxMessageDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *InboxMessageDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *InboxMessageDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetCausationId

`func (o *InboxMessageDto) GetCausationId() string`

GetCausationId returns the CausationId field if non-nil, zero value otherwise.

### GetCausationIdOk

`func (o *InboxMessageDto) GetCausationIdOk() (*string, bool)`

GetCausationIdOk returns a tuple with the CausationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausationId

`func (o *InboxMessageDto) SetCausationId(v string)`

SetCausationId sets CausationId field to given value.

### HasCausationId

`func (o *InboxMessageDto) HasCausationId() bool`

HasCausationId returns a boolean if a field has been set.

### SetCausationIdNil

`func (o *InboxMessageDto) SetCausationIdNil(b bool)`

 SetCausationIdNil sets the value for CausationId to be an explicit nil

### UnsetCausationId
`func (o *InboxMessageDto) UnsetCausationId()`

UnsetCausationId ensures that no value is present for CausationId, not even an explicit nil
### GetLockedBy

`func (o *InboxMessageDto) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *InboxMessageDto) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *InboxMessageDto) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *InboxMessageDto) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### SetLockedByNil

`func (o *InboxMessageDto) SetLockedByNil(b bool)`

 SetLockedByNil sets the value for LockedBy to be an explicit nil

### UnsetLockedBy
`func (o *InboxMessageDto) UnsetLockedBy()`

UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
### GetLockedUntilUtc

`func (o *InboxMessageDto) GetLockedUntilUtc() time.Time`

GetLockedUntilUtc returns the LockedUntilUtc field if non-nil, zero value otherwise.

### GetLockedUntilUtcOk

`func (o *InboxMessageDto) GetLockedUntilUtcOk() (*time.Time, bool)`

GetLockedUntilUtcOk returns a tuple with the LockedUntilUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntilUtc

`func (o *InboxMessageDto) SetLockedUntilUtc(v time.Time)`

SetLockedUntilUtc sets LockedUntilUtc field to given value.

### HasLockedUntilUtc

`func (o *InboxMessageDto) HasLockedUntilUtc() bool`

HasLockedUntilUtc returns a boolean if a field has been set.

### SetLockedUntilUtcNil

`func (o *InboxMessageDto) SetLockedUntilUtcNil(b bool)`

 SetLockedUntilUtcNil sets the value for LockedUntilUtc to be an explicit nil

### UnsetLockedUntilUtc
`func (o *InboxMessageDto) UnsetLockedUntilUtc()`

UnsetLockedUntilUtc ensures that no value is present for LockedUntilUtc, not even an explicit nil
### GetAvailableAtUtc

`func (o *InboxMessageDto) GetAvailableAtUtc() time.Time`

GetAvailableAtUtc returns the AvailableAtUtc field if non-nil, zero value otherwise.

### GetAvailableAtUtcOk

`func (o *InboxMessageDto) GetAvailableAtUtcOk() (*time.Time, bool)`

GetAvailableAtUtcOk returns a tuple with the AvailableAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAtUtc

`func (o *InboxMessageDto) SetAvailableAtUtc(v time.Time)`

SetAvailableAtUtc sets AvailableAtUtc field to given value.

### HasAvailableAtUtc

`func (o *InboxMessageDto) HasAvailableAtUtc() bool`

HasAvailableAtUtc returns a boolean if a field has been set.

### GetReceivedAtUtc

`func (o *InboxMessageDto) GetReceivedAtUtc() time.Time`

GetReceivedAtUtc returns the ReceivedAtUtc field if non-nil, zero value otherwise.

### GetReceivedAtUtcOk

`func (o *InboxMessageDto) GetReceivedAtUtcOk() (*time.Time, bool)`

GetReceivedAtUtcOk returns a tuple with the ReceivedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAtUtc

`func (o *InboxMessageDto) SetReceivedAtUtc(v time.Time)`

SetReceivedAtUtc sets ReceivedAtUtc field to given value.

### HasReceivedAtUtc

`func (o *InboxMessageDto) HasReceivedAtUtc() bool`

HasReceivedAtUtc returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *InboxMessageDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *InboxMessageDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *InboxMessageDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *InboxMessageDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetLastAttemptAtUtc

`func (o *InboxMessageDto) GetLastAttemptAtUtc() time.Time`

GetLastAttemptAtUtc returns the LastAttemptAtUtc field if non-nil, zero value otherwise.

### GetLastAttemptAtUtcOk

`func (o *InboxMessageDto) GetLastAttemptAtUtcOk() (*time.Time, bool)`

GetLastAttemptAtUtcOk returns a tuple with the LastAttemptAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptAtUtc

`func (o *InboxMessageDto) SetLastAttemptAtUtc(v time.Time)`

SetLastAttemptAtUtc sets LastAttemptAtUtc field to given value.

### HasLastAttemptAtUtc

`func (o *InboxMessageDto) HasLastAttemptAtUtc() bool`

HasLastAttemptAtUtc returns a boolean if a field has been set.

### SetLastAttemptAtUtcNil

`func (o *InboxMessageDto) SetLastAttemptAtUtcNil(b bool)`

 SetLastAttemptAtUtcNil sets the value for LastAttemptAtUtc to be an explicit nil

### UnsetLastAttemptAtUtc
`func (o *InboxMessageDto) UnsetLastAttemptAtUtc()`

UnsetLastAttemptAtUtc ensures that no value is present for LastAttemptAtUtc, not even an explicit nil
### GetProcessedAtUtc

`func (o *InboxMessageDto) GetProcessedAtUtc() time.Time`

GetProcessedAtUtc returns the ProcessedAtUtc field if non-nil, zero value otherwise.

### GetProcessedAtUtcOk

`func (o *InboxMessageDto) GetProcessedAtUtcOk() (*time.Time, bool)`

GetProcessedAtUtcOk returns a tuple with the ProcessedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAtUtc

`func (o *InboxMessageDto) SetProcessedAtUtc(v time.Time)`

SetProcessedAtUtc sets ProcessedAtUtc field to given value.

### HasProcessedAtUtc

`func (o *InboxMessageDto) HasProcessedAtUtc() bool`

HasProcessedAtUtc returns a boolean if a field has been set.

### SetProcessedAtUtcNil

`func (o *InboxMessageDto) SetProcessedAtUtcNil(b bool)`

 SetProcessedAtUtcNil sets the value for ProcessedAtUtc to be an explicit nil

### UnsetProcessedAtUtc
`func (o *InboxMessageDto) UnsetProcessedAtUtc()`

UnsetProcessedAtUtc ensures that no value is present for ProcessedAtUtc, not even an explicit nil
### GetFailedAtUtc

`func (o *InboxMessageDto) GetFailedAtUtc() time.Time`

GetFailedAtUtc returns the FailedAtUtc field if non-nil, zero value otherwise.

### GetFailedAtUtcOk

`func (o *InboxMessageDto) GetFailedAtUtcOk() (*time.Time, bool)`

GetFailedAtUtcOk returns a tuple with the FailedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAtUtc

`func (o *InboxMessageDto) SetFailedAtUtc(v time.Time)`

SetFailedAtUtc sets FailedAtUtc field to given value.

### HasFailedAtUtc

`func (o *InboxMessageDto) HasFailedAtUtc() bool`

HasFailedAtUtc returns a boolean if a field has been set.

### SetFailedAtUtcNil

`func (o *InboxMessageDto) SetFailedAtUtcNil(b bool)`

 SetFailedAtUtcNil sets the value for FailedAtUtc to be an explicit nil

### UnsetFailedAtUtc
`func (o *InboxMessageDto) UnsetFailedAtUtc()`

UnsetFailedAtUtc ensures that no value is present for FailedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


