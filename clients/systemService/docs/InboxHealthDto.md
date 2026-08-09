# InboxHealthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ReceivedCount** | Pointer to **int32** |  | [optional] 
**AcceptedCount** | Pointer to **int32** |  | [optional] 
**ProcessingCount** | Pointer to **int32** |  | [optional] 
**RetryScheduledCount** | Pointer to **int32** |  | [optional] 
**RejectedCount** | Pointer to **int32** |  | [optional] 
**QuarantinedCount** | Pointer to **int32** |  | [optional] 
**DeadLetterCount** | Pointer to **int32** |  | [optional] 
**CancelledCount** | Pointer to **int32** |  | [optional] 
**OldestAcceptedAgeSeconds** | Pointer to **NullableFloat64** |  | [optional] 
**LastSuccessfulProcessingUtc** | Pointer to **NullableTime** |  | [optional] 
**SuccessfulProcessingTracked** | Pointer to **bool** |  | [optional] 

## Methods

### NewInboxHealthDto

`func NewInboxHealthDto() *InboxHealthDto`

NewInboxHealthDto instantiates a new InboxHealthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxHealthDtoWithDefaults

`func NewInboxHealthDtoWithDefaults() *InboxHealthDto`

NewInboxHealthDtoWithDefaults instantiates a new InboxHealthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InboxHealthDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InboxHealthDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InboxHealthDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InboxHealthDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceivedCount

`func (o *InboxHealthDto) GetReceivedCount() int32`

GetReceivedCount returns the ReceivedCount field if non-nil, zero value otherwise.

### GetReceivedCountOk

`func (o *InboxHealthDto) GetReceivedCountOk() (*int32, bool)`

GetReceivedCountOk returns a tuple with the ReceivedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedCount

`func (o *InboxHealthDto) SetReceivedCount(v int32)`

SetReceivedCount sets ReceivedCount field to given value.

### HasReceivedCount

`func (o *InboxHealthDto) HasReceivedCount() bool`

HasReceivedCount returns a boolean if a field has been set.

### GetAcceptedCount

`func (o *InboxHealthDto) GetAcceptedCount() int32`

GetAcceptedCount returns the AcceptedCount field if non-nil, zero value otherwise.

### GetAcceptedCountOk

`func (o *InboxHealthDto) GetAcceptedCountOk() (*int32, bool)`

GetAcceptedCountOk returns a tuple with the AcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCount

`func (o *InboxHealthDto) SetAcceptedCount(v int32)`

SetAcceptedCount sets AcceptedCount field to given value.

### HasAcceptedCount

`func (o *InboxHealthDto) HasAcceptedCount() bool`

HasAcceptedCount returns a boolean if a field has been set.

### GetProcessingCount

`func (o *InboxHealthDto) GetProcessingCount() int32`

GetProcessingCount returns the ProcessingCount field if non-nil, zero value otherwise.

### GetProcessingCountOk

`func (o *InboxHealthDto) GetProcessingCountOk() (*int32, bool)`

GetProcessingCountOk returns a tuple with the ProcessingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingCount

`func (o *InboxHealthDto) SetProcessingCount(v int32)`

SetProcessingCount sets ProcessingCount field to given value.

### HasProcessingCount

`func (o *InboxHealthDto) HasProcessingCount() bool`

HasProcessingCount returns a boolean if a field has been set.

### GetRetryScheduledCount

`func (o *InboxHealthDto) GetRetryScheduledCount() int32`

GetRetryScheduledCount returns the RetryScheduledCount field if non-nil, zero value otherwise.

### GetRetryScheduledCountOk

`func (o *InboxHealthDto) GetRetryScheduledCountOk() (*int32, bool)`

GetRetryScheduledCountOk returns a tuple with the RetryScheduledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryScheduledCount

`func (o *InboxHealthDto) SetRetryScheduledCount(v int32)`

SetRetryScheduledCount sets RetryScheduledCount field to given value.

### HasRetryScheduledCount

`func (o *InboxHealthDto) HasRetryScheduledCount() bool`

HasRetryScheduledCount returns a boolean if a field has been set.

### GetRejectedCount

`func (o *InboxHealthDto) GetRejectedCount() int32`

GetRejectedCount returns the RejectedCount field if non-nil, zero value otherwise.

### GetRejectedCountOk

`func (o *InboxHealthDto) GetRejectedCountOk() (*int32, bool)`

GetRejectedCountOk returns a tuple with the RejectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCount

`func (o *InboxHealthDto) SetRejectedCount(v int32)`

SetRejectedCount sets RejectedCount field to given value.

### HasRejectedCount

`func (o *InboxHealthDto) HasRejectedCount() bool`

HasRejectedCount returns a boolean if a field has been set.

### GetQuarantinedCount

`func (o *InboxHealthDto) GetQuarantinedCount() int32`

GetQuarantinedCount returns the QuarantinedCount field if non-nil, zero value otherwise.

### GetQuarantinedCountOk

`func (o *InboxHealthDto) GetQuarantinedCountOk() (*int32, bool)`

GetQuarantinedCountOk returns a tuple with the QuarantinedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantinedCount

`func (o *InboxHealthDto) SetQuarantinedCount(v int32)`

SetQuarantinedCount sets QuarantinedCount field to given value.

### HasQuarantinedCount

`func (o *InboxHealthDto) HasQuarantinedCount() bool`

HasQuarantinedCount returns a boolean if a field has been set.

### GetDeadLetterCount

`func (o *InboxHealthDto) GetDeadLetterCount() int32`

GetDeadLetterCount returns the DeadLetterCount field if non-nil, zero value otherwise.

### GetDeadLetterCountOk

`func (o *InboxHealthDto) GetDeadLetterCountOk() (*int32, bool)`

GetDeadLetterCountOk returns a tuple with the DeadLetterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterCount

`func (o *InboxHealthDto) SetDeadLetterCount(v int32)`

SetDeadLetterCount sets DeadLetterCount field to given value.

### HasDeadLetterCount

`func (o *InboxHealthDto) HasDeadLetterCount() bool`

HasDeadLetterCount returns a boolean if a field has been set.

### GetCancelledCount

`func (o *InboxHealthDto) GetCancelledCount() int32`

GetCancelledCount returns the CancelledCount field if non-nil, zero value otherwise.

### GetCancelledCountOk

`func (o *InboxHealthDto) GetCancelledCountOk() (*int32, bool)`

GetCancelledCountOk returns a tuple with the CancelledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledCount

`func (o *InboxHealthDto) SetCancelledCount(v int32)`

SetCancelledCount sets CancelledCount field to given value.

### HasCancelledCount

`func (o *InboxHealthDto) HasCancelledCount() bool`

HasCancelledCount returns a boolean if a field has been set.

### GetOldestAcceptedAgeSeconds

`func (o *InboxHealthDto) GetOldestAcceptedAgeSeconds() float64`

GetOldestAcceptedAgeSeconds returns the OldestAcceptedAgeSeconds field if non-nil, zero value otherwise.

### GetOldestAcceptedAgeSecondsOk

`func (o *InboxHealthDto) GetOldestAcceptedAgeSecondsOk() (*float64, bool)`

GetOldestAcceptedAgeSecondsOk returns a tuple with the OldestAcceptedAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestAcceptedAgeSeconds

`func (o *InboxHealthDto) SetOldestAcceptedAgeSeconds(v float64)`

SetOldestAcceptedAgeSeconds sets OldestAcceptedAgeSeconds field to given value.

### HasOldestAcceptedAgeSeconds

`func (o *InboxHealthDto) HasOldestAcceptedAgeSeconds() bool`

HasOldestAcceptedAgeSeconds returns a boolean if a field has been set.

### SetOldestAcceptedAgeSecondsNil

`func (o *InboxHealthDto) SetOldestAcceptedAgeSecondsNil(b bool)`

 SetOldestAcceptedAgeSecondsNil sets the value for OldestAcceptedAgeSeconds to be an explicit nil

### UnsetOldestAcceptedAgeSeconds
`func (o *InboxHealthDto) UnsetOldestAcceptedAgeSeconds()`

UnsetOldestAcceptedAgeSeconds ensures that no value is present for OldestAcceptedAgeSeconds, not even an explicit nil
### GetLastSuccessfulProcessingUtc

`func (o *InboxHealthDto) GetLastSuccessfulProcessingUtc() time.Time`

GetLastSuccessfulProcessingUtc returns the LastSuccessfulProcessingUtc field if non-nil, zero value otherwise.

### GetLastSuccessfulProcessingUtcOk

`func (o *InboxHealthDto) GetLastSuccessfulProcessingUtcOk() (*time.Time, bool)`

GetLastSuccessfulProcessingUtcOk returns a tuple with the LastSuccessfulProcessingUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulProcessingUtc

`func (o *InboxHealthDto) SetLastSuccessfulProcessingUtc(v time.Time)`

SetLastSuccessfulProcessingUtc sets LastSuccessfulProcessingUtc field to given value.

### HasLastSuccessfulProcessingUtc

`func (o *InboxHealthDto) HasLastSuccessfulProcessingUtc() bool`

HasLastSuccessfulProcessingUtc returns a boolean if a field has been set.

### SetLastSuccessfulProcessingUtcNil

`func (o *InboxHealthDto) SetLastSuccessfulProcessingUtcNil(b bool)`

 SetLastSuccessfulProcessingUtcNil sets the value for LastSuccessfulProcessingUtc to be an explicit nil

### UnsetLastSuccessfulProcessingUtc
`func (o *InboxHealthDto) UnsetLastSuccessfulProcessingUtc()`

UnsetLastSuccessfulProcessingUtc ensures that no value is present for LastSuccessfulProcessingUtc, not even an explicit nil
### GetSuccessfulProcessingTracked

`func (o *InboxHealthDto) GetSuccessfulProcessingTracked() bool`

GetSuccessfulProcessingTracked returns the SuccessfulProcessingTracked field if non-nil, zero value otherwise.

### GetSuccessfulProcessingTrackedOk

`func (o *InboxHealthDto) GetSuccessfulProcessingTrackedOk() (*bool, bool)`

GetSuccessfulProcessingTrackedOk returns a tuple with the SuccessfulProcessingTracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulProcessingTracked

`func (o *InboxHealthDto) SetSuccessfulProcessingTracked(v bool)`

SetSuccessfulProcessingTracked sets SuccessfulProcessingTracked field to given value.

### HasSuccessfulProcessingTracked

`func (o *InboxHealthDto) HasSuccessfulProcessingTracked() bool`

HasSuccessfulProcessingTracked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


