# PostingExecutionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**PostingIntentId** | Pointer to **NullableString** |  | [optional] 
**PostingIdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**IntentType** | Pointer to **NullableString** |  | [optional] 
**PostingOperation** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentType** | Pointer to **NullableString** |  | [optional] 
**SourceDocumentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FailureKind** | Pointer to **NullableString** |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**CausationId** | Pointer to **NullableString** |  | [optional] 
**ReceivedAtUtc** | Pointer to **time.Time** |  | [optional] 
**ProcessingStartedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CompletedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**BookResults** | Pointer to [**[]PostingBookResultDto**](PostingBookResultDto.md) |  | [optional] 
**FailureClass** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewPostingExecutionDto

`func NewPostingExecutionDto() *PostingExecutionDto`

NewPostingExecutionDto instantiates a new PostingExecutionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostingExecutionDtoWithDefaults

`func NewPostingExecutionDtoWithDefaults() *PostingExecutionDto`

NewPostingExecutionDtoWithDefaults instantiates a new PostingExecutionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostingExecutionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostingExecutionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostingExecutionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostingExecutionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PostingExecutionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PostingExecutionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PostingExecutionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PostingExecutionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PostingExecutionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PostingExecutionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PostingExecutionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PostingExecutionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *PostingExecutionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PostingExecutionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PostingExecutionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PostingExecutionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PostingExecutionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PostingExecutionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PostingExecutionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PostingExecutionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PostingExecutionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PostingExecutionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PostingExecutionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PostingExecutionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetPostingIntentId

`func (o *PostingExecutionDto) GetPostingIntentId() string`

GetPostingIntentId returns the PostingIntentId field if non-nil, zero value otherwise.

### GetPostingIntentIdOk

`func (o *PostingExecutionDto) GetPostingIntentIdOk() (*string, bool)`

GetPostingIntentIdOk returns a tuple with the PostingIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingIntentId

`func (o *PostingExecutionDto) SetPostingIntentId(v string)`

SetPostingIntentId sets PostingIntentId field to given value.

### HasPostingIntentId

`func (o *PostingExecutionDto) HasPostingIntentId() bool`

HasPostingIntentId returns a boolean if a field has been set.

### SetPostingIntentIdNil

`func (o *PostingExecutionDto) SetPostingIntentIdNil(b bool)`

 SetPostingIntentIdNil sets the value for PostingIntentId to be an explicit nil

### UnsetPostingIntentId
`func (o *PostingExecutionDto) UnsetPostingIntentId()`

UnsetPostingIntentId ensures that no value is present for PostingIntentId, not even an explicit nil
### GetPostingIdempotencyKey

`func (o *PostingExecutionDto) GetPostingIdempotencyKey() string`

GetPostingIdempotencyKey returns the PostingIdempotencyKey field if non-nil, zero value otherwise.

### GetPostingIdempotencyKeyOk

`func (o *PostingExecutionDto) GetPostingIdempotencyKeyOk() (*string, bool)`

GetPostingIdempotencyKeyOk returns a tuple with the PostingIdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingIdempotencyKey

`func (o *PostingExecutionDto) SetPostingIdempotencyKey(v string)`

SetPostingIdempotencyKey sets PostingIdempotencyKey field to given value.

### HasPostingIdempotencyKey

`func (o *PostingExecutionDto) HasPostingIdempotencyKey() bool`

HasPostingIdempotencyKey returns a boolean if a field has been set.

### SetPostingIdempotencyKeyNil

`func (o *PostingExecutionDto) SetPostingIdempotencyKeyNil(b bool)`

 SetPostingIdempotencyKeyNil sets the value for PostingIdempotencyKey to be an explicit nil

### UnsetPostingIdempotencyKey
`func (o *PostingExecutionDto) UnsetPostingIdempotencyKey()`

UnsetPostingIdempotencyKey ensures that no value is present for PostingIdempotencyKey, not even an explicit nil
### GetIntentType

`func (o *PostingExecutionDto) GetIntentType() string`

GetIntentType returns the IntentType field if non-nil, zero value otherwise.

### GetIntentTypeOk

`func (o *PostingExecutionDto) GetIntentTypeOk() (*string, bool)`

GetIntentTypeOk returns a tuple with the IntentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentType

`func (o *PostingExecutionDto) SetIntentType(v string)`

SetIntentType sets IntentType field to given value.

### HasIntentType

`func (o *PostingExecutionDto) HasIntentType() bool`

HasIntentType returns a boolean if a field has been set.

### SetIntentTypeNil

`func (o *PostingExecutionDto) SetIntentTypeNil(b bool)`

 SetIntentTypeNil sets the value for IntentType to be an explicit nil

### UnsetIntentType
`func (o *PostingExecutionDto) UnsetIntentType()`

UnsetIntentType ensures that no value is present for IntentType, not even an explicit nil
### GetPostingOperation

`func (o *PostingExecutionDto) GetPostingOperation() string`

GetPostingOperation returns the PostingOperation field if non-nil, zero value otherwise.

### GetPostingOperationOk

`func (o *PostingExecutionDto) GetPostingOperationOk() (*string, bool)`

GetPostingOperationOk returns a tuple with the PostingOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingOperation

`func (o *PostingExecutionDto) SetPostingOperation(v string)`

SetPostingOperation sets PostingOperation field to given value.

### HasPostingOperation

`func (o *PostingExecutionDto) HasPostingOperation() bool`

HasPostingOperation returns a boolean if a field has been set.

### SetPostingOperationNil

`func (o *PostingExecutionDto) SetPostingOperationNil(b bool)`

 SetPostingOperationNil sets the value for PostingOperation to be an explicit nil

### UnsetPostingOperation
`func (o *PostingExecutionDto) UnsetPostingOperation()`

UnsetPostingOperation ensures that no value is present for PostingOperation, not even an explicit nil
### GetSourceDocumentType

`func (o *PostingExecutionDto) GetSourceDocumentType() string`

GetSourceDocumentType returns the SourceDocumentType field if non-nil, zero value otherwise.

### GetSourceDocumentTypeOk

`func (o *PostingExecutionDto) GetSourceDocumentTypeOk() (*string, bool)`

GetSourceDocumentTypeOk returns a tuple with the SourceDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentType

`func (o *PostingExecutionDto) SetSourceDocumentType(v string)`

SetSourceDocumentType sets SourceDocumentType field to given value.

### HasSourceDocumentType

`func (o *PostingExecutionDto) HasSourceDocumentType() bool`

HasSourceDocumentType returns a boolean if a field has been set.

### SetSourceDocumentTypeNil

`func (o *PostingExecutionDto) SetSourceDocumentTypeNil(b bool)`

 SetSourceDocumentTypeNil sets the value for SourceDocumentType to be an explicit nil

### UnsetSourceDocumentType
`func (o *PostingExecutionDto) UnsetSourceDocumentType()`

UnsetSourceDocumentType ensures that no value is present for SourceDocumentType, not even an explicit nil
### GetSourceDocumentId

`func (o *PostingExecutionDto) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *PostingExecutionDto) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *PostingExecutionDto) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *PostingExecutionDto) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### SetSourceDocumentIdNil

`func (o *PostingExecutionDto) SetSourceDocumentIdNil(b bool)`

 SetSourceDocumentIdNil sets the value for SourceDocumentId to be an explicit nil

### UnsetSourceDocumentId
`func (o *PostingExecutionDto) UnsetSourceDocumentId()`

UnsetSourceDocumentId ensures that no value is present for SourceDocumentId, not even an explicit nil
### GetStatus

`func (o *PostingExecutionDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostingExecutionDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostingExecutionDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostingExecutionDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureKind

`func (o *PostingExecutionDto) GetFailureKind() string`

GetFailureKind returns the FailureKind field if non-nil, zero value otherwise.

### GetFailureKindOk

`func (o *PostingExecutionDto) GetFailureKindOk() (*string, bool)`

GetFailureKindOk returns a tuple with the FailureKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureKind

`func (o *PostingExecutionDto) SetFailureKind(v string)`

SetFailureKind sets FailureKind field to given value.

### HasFailureKind

`func (o *PostingExecutionDto) HasFailureKind() bool`

HasFailureKind returns a boolean if a field has been set.

### SetFailureKindNil

`func (o *PostingExecutionDto) SetFailureKindNil(b bool)`

 SetFailureKindNil sets the value for FailureKind to be an explicit nil

### UnsetFailureKind
`func (o *PostingExecutionDto) UnsetFailureKind()`

UnsetFailureKind ensures that no value is present for FailureKind, not even an explicit nil
### GetFailureCode

`func (o *PostingExecutionDto) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PostingExecutionDto) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PostingExecutionDto) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PostingExecutionDto) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PostingExecutionDto) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PostingExecutionDto) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetRetryable

`func (o *PostingExecutionDto) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *PostingExecutionDto) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *PostingExecutionDto) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *PostingExecutionDto) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetCorrelationId

`func (o *PostingExecutionDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *PostingExecutionDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *PostingExecutionDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *PostingExecutionDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *PostingExecutionDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *PostingExecutionDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetCausationId

`func (o *PostingExecutionDto) GetCausationId() string`

GetCausationId returns the CausationId field if non-nil, zero value otherwise.

### GetCausationIdOk

`func (o *PostingExecutionDto) GetCausationIdOk() (*string, bool)`

GetCausationIdOk returns a tuple with the CausationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausationId

`func (o *PostingExecutionDto) SetCausationId(v string)`

SetCausationId sets CausationId field to given value.

### HasCausationId

`func (o *PostingExecutionDto) HasCausationId() bool`

HasCausationId returns a boolean if a field has been set.

### SetCausationIdNil

`func (o *PostingExecutionDto) SetCausationIdNil(b bool)`

 SetCausationIdNil sets the value for CausationId to be an explicit nil

### UnsetCausationId
`func (o *PostingExecutionDto) UnsetCausationId()`

UnsetCausationId ensures that no value is present for CausationId, not even an explicit nil
### GetReceivedAtUtc

`func (o *PostingExecutionDto) GetReceivedAtUtc() time.Time`

GetReceivedAtUtc returns the ReceivedAtUtc field if non-nil, zero value otherwise.

### GetReceivedAtUtcOk

`func (o *PostingExecutionDto) GetReceivedAtUtcOk() (*time.Time, bool)`

GetReceivedAtUtcOk returns a tuple with the ReceivedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAtUtc

`func (o *PostingExecutionDto) SetReceivedAtUtc(v time.Time)`

SetReceivedAtUtc sets ReceivedAtUtc field to given value.

### HasReceivedAtUtc

`func (o *PostingExecutionDto) HasReceivedAtUtc() bool`

HasReceivedAtUtc returns a boolean if a field has been set.

### GetProcessingStartedAtUtc

`func (o *PostingExecutionDto) GetProcessingStartedAtUtc() time.Time`

GetProcessingStartedAtUtc returns the ProcessingStartedAtUtc field if non-nil, zero value otherwise.

### GetProcessingStartedAtUtcOk

`func (o *PostingExecutionDto) GetProcessingStartedAtUtcOk() (*time.Time, bool)`

GetProcessingStartedAtUtcOk returns a tuple with the ProcessingStartedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStartedAtUtc

`func (o *PostingExecutionDto) SetProcessingStartedAtUtc(v time.Time)`

SetProcessingStartedAtUtc sets ProcessingStartedAtUtc field to given value.

### HasProcessingStartedAtUtc

`func (o *PostingExecutionDto) HasProcessingStartedAtUtc() bool`

HasProcessingStartedAtUtc returns a boolean if a field has been set.

### SetProcessingStartedAtUtcNil

`func (o *PostingExecutionDto) SetProcessingStartedAtUtcNil(b bool)`

 SetProcessingStartedAtUtcNil sets the value for ProcessingStartedAtUtc to be an explicit nil

### UnsetProcessingStartedAtUtc
`func (o *PostingExecutionDto) UnsetProcessingStartedAtUtc()`

UnsetProcessingStartedAtUtc ensures that no value is present for ProcessingStartedAtUtc, not even an explicit nil
### GetCompletedAtUtc

`func (o *PostingExecutionDto) GetCompletedAtUtc() time.Time`

GetCompletedAtUtc returns the CompletedAtUtc field if non-nil, zero value otherwise.

### GetCompletedAtUtcOk

`func (o *PostingExecutionDto) GetCompletedAtUtcOk() (*time.Time, bool)`

GetCompletedAtUtcOk returns a tuple with the CompletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtUtc

`func (o *PostingExecutionDto) SetCompletedAtUtc(v time.Time)`

SetCompletedAtUtc sets CompletedAtUtc field to given value.

### HasCompletedAtUtc

`func (o *PostingExecutionDto) HasCompletedAtUtc() bool`

HasCompletedAtUtc returns a boolean if a field has been set.

### SetCompletedAtUtcNil

`func (o *PostingExecutionDto) SetCompletedAtUtcNil(b bool)`

 SetCompletedAtUtcNil sets the value for CompletedAtUtc to be an explicit nil

### UnsetCompletedAtUtc
`func (o *PostingExecutionDto) UnsetCompletedAtUtc()`

UnsetCompletedAtUtc ensures that no value is present for CompletedAtUtc, not even an explicit nil
### GetBookResults

`func (o *PostingExecutionDto) GetBookResults() []PostingBookResultDto`

GetBookResults returns the BookResults field if non-nil, zero value otherwise.

### GetBookResultsOk

`func (o *PostingExecutionDto) GetBookResultsOk() (*[]PostingBookResultDto, bool)`

GetBookResultsOk returns a tuple with the BookResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookResults

`func (o *PostingExecutionDto) SetBookResults(v []PostingBookResultDto)`

SetBookResults sets BookResults field to given value.

### HasBookResults

`func (o *PostingExecutionDto) HasBookResults() bool`

HasBookResults returns a boolean if a field has been set.

### SetBookResultsNil

`func (o *PostingExecutionDto) SetBookResultsNil(b bool)`

 SetBookResultsNil sets the value for BookResults to be an explicit nil

### UnsetBookResults
`func (o *PostingExecutionDto) UnsetBookResults()`

UnsetBookResults ensures that no value is present for BookResults, not even an explicit nil
### GetFailureClass

`func (o *PostingExecutionDto) GetFailureClass() string`

GetFailureClass returns the FailureClass field if non-nil, zero value otherwise.

### GetFailureClassOk

`func (o *PostingExecutionDto) GetFailureClassOk() (*string, bool)`

GetFailureClassOk returns a tuple with the FailureClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClass

`func (o *PostingExecutionDto) SetFailureClass(v string)`

SetFailureClass sets FailureClass field to given value.

### HasFailureClass

`func (o *PostingExecutionDto) HasFailureClass() bool`

HasFailureClass returns a boolean if a field has been set.

### SetFailureClassNil

`func (o *PostingExecutionDto) SetFailureClassNil(b bool)`

 SetFailureClassNil sets the value for FailureClass to be an explicit nil

### UnsetFailureClass
`func (o *PostingExecutionDto) UnsetFailureClass()`

UnsetFailureClass ensures that no value is present for FailureClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


