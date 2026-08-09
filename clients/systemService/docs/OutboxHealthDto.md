# OutboxHealthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**PendingCount** | Pointer to **int32** |  | [optional] 
**ProcessingCount** | Pointer to **int32** |  | [optional] 
**FailedCount** | Pointer to **int32** |  | [optional] 
**DeadLetterCount** | Pointer to **int32** |  | [optional] 
**OldestPendingAgeSeconds** | Pointer to **NullableFloat64** |  | [optional] 
**LastSuccessfulDispatchUtc** | Pointer to **NullableTime** |  | [optional] 
**SuccessfulDispatchTracked** | Pointer to **bool** |  | [optional] 

## Methods

### NewOutboxHealthDto

`func NewOutboxHealthDto() *OutboxHealthDto`

NewOutboxHealthDto instantiates a new OutboxHealthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxHealthDtoWithDefaults

`func NewOutboxHealthDtoWithDefaults() *OutboxHealthDto`

NewOutboxHealthDtoWithDefaults instantiates a new OutboxHealthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OutboxHealthDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutboxHealthDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutboxHealthDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OutboxHealthDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPendingCount

`func (o *OutboxHealthDto) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *OutboxHealthDto) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *OutboxHealthDto) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *OutboxHealthDto) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetProcessingCount

`func (o *OutboxHealthDto) GetProcessingCount() int32`

GetProcessingCount returns the ProcessingCount field if non-nil, zero value otherwise.

### GetProcessingCountOk

`func (o *OutboxHealthDto) GetProcessingCountOk() (*int32, bool)`

GetProcessingCountOk returns a tuple with the ProcessingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingCount

`func (o *OutboxHealthDto) SetProcessingCount(v int32)`

SetProcessingCount sets ProcessingCount field to given value.

### HasProcessingCount

`func (o *OutboxHealthDto) HasProcessingCount() bool`

HasProcessingCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *OutboxHealthDto) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *OutboxHealthDto) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *OutboxHealthDto) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *OutboxHealthDto) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetDeadLetterCount

`func (o *OutboxHealthDto) GetDeadLetterCount() int32`

GetDeadLetterCount returns the DeadLetterCount field if non-nil, zero value otherwise.

### GetDeadLetterCountOk

`func (o *OutboxHealthDto) GetDeadLetterCountOk() (*int32, bool)`

GetDeadLetterCountOk returns a tuple with the DeadLetterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterCount

`func (o *OutboxHealthDto) SetDeadLetterCount(v int32)`

SetDeadLetterCount sets DeadLetterCount field to given value.

### HasDeadLetterCount

`func (o *OutboxHealthDto) HasDeadLetterCount() bool`

HasDeadLetterCount returns a boolean if a field has been set.

### GetOldestPendingAgeSeconds

`func (o *OutboxHealthDto) GetOldestPendingAgeSeconds() float64`

GetOldestPendingAgeSeconds returns the OldestPendingAgeSeconds field if non-nil, zero value otherwise.

### GetOldestPendingAgeSecondsOk

`func (o *OutboxHealthDto) GetOldestPendingAgeSecondsOk() (*float64, bool)`

GetOldestPendingAgeSecondsOk returns a tuple with the OldestPendingAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestPendingAgeSeconds

`func (o *OutboxHealthDto) SetOldestPendingAgeSeconds(v float64)`

SetOldestPendingAgeSeconds sets OldestPendingAgeSeconds field to given value.

### HasOldestPendingAgeSeconds

`func (o *OutboxHealthDto) HasOldestPendingAgeSeconds() bool`

HasOldestPendingAgeSeconds returns a boolean if a field has been set.

### SetOldestPendingAgeSecondsNil

`func (o *OutboxHealthDto) SetOldestPendingAgeSecondsNil(b bool)`

 SetOldestPendingAgeSecondsNil sets the value for OldestPendingAgeSeconds to be an explicit nil

### UnsetOldestPendingAgeSeconds
`func (o *OutboxHealthDto) UnsetOldestPendingAgeSeconds()`

UnsetOldestPendingAgeSeconds ensures that no value is present for OldestPendingAgeSeconds, not even an explicit nil
### GetLastSuccessfulDispatchUtc

`func (o *OutboxHealthDto) GetLastSuccessfulDispatchUtc() time.Time`

GetLastSuccessfulDispatchUtc returns the LastSuccessfulDispatchUtc field if non-nil, zero value otherwise.

### GetLastSuccessfulDispatchUtcOk

`func (o *OutboxHealthDto) GetLastSuccessfulDispatchUtcOk() (*time.Time, bool)`

GetLastSuccessfulDispatchUtcOk returns a tuple with the LastSuccessfulDispatchUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulDispatchUtc

`func (o *OutboxHealthDto) SetLastSuccessfulDispatchUtc(v time.Time)`

SetLastSuccessfulDispatchUtc sets LastSuccessfulDispatchUtc field to given value.

### HasLastSuccessfulDispatchUtc

`func (o *OutboxHealthDto) HasLastSuccessfulDispatchUtc() bool`

HasLastSuccessfulDispatchUtc returns a boolean if a field has been set.

### SetLastSuccessfulDispatchUtcNil

`func (o *OutboxHealthDto) SetLastSuccessfulDispatchUtcNil(b bool)`

 SetLastSuccessfulDispatchUtcNil sets the value for LastSuccessfulDispatchUtc to be an explicit nil

### UnsetLastSuccessfulDispatchUtc
`func (o *OutboxHealthDto) UnsetLastSuccessfulDispatchUtc()`

UnsetLastSuccessfulDispatchUtc ensures that no value is present for LastSuccessfulDispatchUtc, not even an explicit nil
### GetSuccessfulDispatchTracked

`func (o *OutboxHealthDto) GetSuccessfulDispatchTracked() bool`

GetSuccessfulDispatchTracked returns the SuccessfulDispatchTracked field if non-nil, zero value otherwise.

### GetSuccessfulDispatchTrackedOk

`func (o *OutboxHealthDto) GetSuccessfulDispatchTrackedOk() (*bool, bool)`

GetSuccessfulDispatchTrackedOk returns a tuple with the SuccessfulDispatchTracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulDispatchTracked

`func (o *OutboxHealthDto) SetSuccessfulDispatchTracked(v bool)`

SetSuccessfulDispatchTracked sets SuccessfulDispatchTracked field to given value.

### HasSuccessfulDispatchTracked

`func (o *OutboxHealthDto) HasSuccessfulDispatchTracked() bool`

HasSuccessfulDispatchTracked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


