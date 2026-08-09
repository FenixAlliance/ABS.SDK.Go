# OutboxMessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Attempts** | Pointer to **int32** |  | [optional] 
**MaxAttempts** | Pointer to **int32** |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**LockedBy** | Pointer to **NullableString** |  | [optional] 
**LockedUntilUtc** | Pointer to **NullableTime** |  | [optional] 
**AvailableAtUtc** | Pointer to **time.Time** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**LastAttemptAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ProcessedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**FailedAtUtc** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOutboxMessageDto

`func NewOutboxMessageDto() *OutboxMessageDto`

NewOutboxMessageDto instantiates a new OutboxMessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxMessageDtoWithDefaults

`func NewOutboxMessageDtoWithDefaults() *OutboxMessageDto`

NewOutboxMessageDtoWithDefaults instantiates a new OutboxMessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutboxMessageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboxMessageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboxMessageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutboxMessageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OutboxMessageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OutboxMessageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *OutboxMessageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OutboxMessageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OutboxMessageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OutboxMessageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *OutboxMessageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *OutboxMessageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *OutboxMessageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OutboxMessageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OutboxMessageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OutboxMessageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *OutboxMessageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *OutboxMessageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetKind

`func (o *OutboxMessageDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutboxMessageDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutboxMessageDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OutboxMessageDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessageType

`func (o *OutboxMessageDto) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *OutboxMessageDto) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *OutboxMessageDto) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *OutboxMessageDto) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### SetMessageTypeNil

`func (o *OutboxMessageDto) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *OutboxMessageDto) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetStatus

`func (o *OutboxMessageDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutboxMessageDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutboxMessageDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OutboxMessageDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttempts

`func (o *OutboxMessageDto) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *OutboxMessageDto) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *OutboxMessageDto) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *OutboxMessageDto) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *OutboxMessageDto) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *OutboxMessageDto) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *OutboxMessageDto) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *OutboxMessageDto) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetFailureCode

`func (o *OutboxMessageDto) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *OutboxMessageDto) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *OutboxMessageDto) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *OutboxMessageDto) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *OutboxMessageDto) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *OutboxMessageDto) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureReason

`func (o *OutboxMessageDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OutboxMessageDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OutboxMessageDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OutboxMessageDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *OutboxMessageDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *OutboxMessageDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetIdempotencyKey

`func (o *OutboxMessageDto) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OutboxMessageDto) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OutboxMessageDto) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OutboxMessageDto) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *OutboxMessageDto) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *OutboxMessageDto) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetCorrelationId

`func (o *OutboxMessageDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *OutboxMessageDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *OutboxMessageDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *OutboxMessageDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *OutboxMessageDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *OutboxMessageDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetLockedBy

`func (o *OutboxMessageDto) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *OutboxMessageDto) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *OutboxMessageDto) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *OutboxMessageDto) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### SetLockedByNil

`func (o *OutboxMessageDto) SetLockedByNil(b bool)`

 SetLockedByNil sets the value for LockedBy to be an explicit nil

### UnsetLockedBy
`func (o *OutboxMessageDto) UnsetLockedBy()`

UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
### GetLockedUntilUtc

`func (o *OutboxMessageDto) GetLockedUntilUtc() time.Time`

GetLockedUntilUtc returns the LockedUntilUtc field if non-nil, zero value otherwise.

### GetLockedUntilUtcOk

`func (o *OutboxMessageDto) GetLockedUntilUtcOk() (*time.Time, bool)`

GetLockedUntilUtcOk returns a tuple with the LockedUntilUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntilUtc

`func (o *OutboxMessageDto) SetLockedUntilUtc(v time.Time)`

SetLockedUntilUtc sets LockedUntilUtc field to given value.

### HasLockedUntilUtc

`func (o *OutboxMessageDto) HasLockedUntilUtc() bool`

HasLockedUntilUtc returns a boolean if a field has been set.

### SetLockedUntilUtcNil

`func (o *OutboxMessageDto) SetLockedUntilUtcNil(b bool)`

 SetLockedUntilUtcNil sets the value for LockedUntilUtc to be an explicit nil

### UnsetLockedUntilUtc
`func (o *OutboxMessageDto) UnsetLockedUntilUtc()`

UnsetLockedUntilUtc ensures that no value is present for LockedUntilUtc, not even an explicit nil
### GetAvailableAtUtc

`func (o *OutboxMessageDto) GetAvailableAtUtc() time.Time`

GetAvailableAtUtc returns the AvailableAtUtc field if non-nil, zero value otherwise.

### GetAvailableAtUtcOk

`func (o *OutboxMessageDto) GetAvailableAtUtcOk() (*time.Time, bool)`

GetAvailableAtUtcOk returns a tuple with the AvailableAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAtUtc

`func (o *OutboxMessageDto) SetAvailableAtUtc(v time.Time)`

SetAvailableAtUtc sets AvailableAtUtc field to given value.

### HasAvailableAtUtc

`func (o *OutboxMessageDto) HasAvailableAtUtc() bool`

HasAvailableAtUtc returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *OutboxMessageDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OutboxMessageDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OutboxMessageDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *OutboxMessageDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetLastAttemptAtUtc

`func (o *OutboxMessageDto) GetLastAttemptAtUtc() time.Time`

GetLastAttemptAtUtc returns the LastAttemptAtUtc field if non-nil, zero value otherwise.

### GetLastAttemptAtUtcOk

`func (o *OutboxMessageDto) GetLastAttemptAtUtcOk() (*time.Time, bool)`

GetLastAttemptAtUtcOk returns a tuple with the LastAttemptAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptAtUtc

`func (o *OutboxMessageDto) SetLastAttemptAtUtc(v time.Time)`

SetLastAttemptAtUtc sets LastAttemptAtUtc field to given value.

### HasLastAttemptAtUtc

`func (o *OutboxMessageDto) HasLastAttemptAtUtc() bool`

HasLastAttemptAtUtc returns a boolean if a field has been set.

### SetLastAttemptAtUtcNil

`func (o *OutboxMessageDto) SetLastAttemptAtUtcNil(b bool)`

 SetLastAttemptAtUtcNil sets the value for LastAttemptAtUtc to be an explicit nil

### UnsetLastAttemptAtUtc
`func (o *OutboxMessageDto) UnsetLastAttemptAtUtc()`

UnsetLastAttemptAtUtc ensures that no value is present for LastAttemptAtUtc, not even an explicit nil
### GetProcessedAtUtc

`func (o *OutboxMessageDto) GetProcessedAtUtc() time.Time`

GetProcessedAtUtc returns the ProcessedAtUtc field if non-nil, zero value otherwise.

### GetProcessedAtUtcOk

`func (o *OutboxMessageDto) GetProcessedAtUtcOk() (*time.Time, bool)`

GetProcessedAtUtcOk returns a tuple with the ProcessedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAtUtc

`func (o *OutboxMessageDto) SetProcessedAtUtc(v time.Time)`

SetProcessedAtUtc sets ProcessedAtUtc field to given value.

### HasProcessedAtUtc

`func (o *OutboxMessageDto) HasProcessedAtUtc() bool`

HasProcessedAtUtc returns a boolean if a field has been set.

### SetProcessedAtUtcNil

`func (o *OutboxMessageDto) SetProcessedAtUtcNil(b bool)`

 SetProcessedAtUtcNil sets the value for ProcessedAtUtc to be an explicit nil

### UnsetProcessedAtUtc
`func (o *OutboxMessageDto) UnsetProcessedAtUtc()`

UnsetProcessedAtUtc ensures that no value is present for ProcessedAtUtc, not even an explicit nil
### GetFailedAtUtc

`func (o *OutboxMessageDto) GetFailedAtUtc() time.Time`

GetFailedAtUtc returns the FailedAtUtc field if non-nil, zero value otherwise.

### GetFailedAtUtcOk

`func (o *OutboxMessageDto) GetFailedAtUtcOk() (*time.Time, bool)`

GetFailedAtUtcOk returns a tuple with the FailedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAtUtc

`func (o *OutboxMessageDto) SetFailedAtUtc(v time.Time)`

SetFailedAtUtc sets FailedAtUtc field to given value.

### HasFailedAtUtc

`func (o *OutboxMessageDto) HasFailedAtUtc() bool`

HasFailedAtUtc returns a boolean if a field has been set.

### SetFailedAtUtcNil

`func (o *OutboxMessageDto) SetFailedAtUtcNil(b bool)`

 SetFailedAtUtcNil sets the value for FailedAtUtc to be an explicit nil

### UnsetFailedAtUtc
`func (o *OutboxMessageDto) UnsetFailedAtUtc()`

UnsetFailedAtUtc ensures that no value is present for FailedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


