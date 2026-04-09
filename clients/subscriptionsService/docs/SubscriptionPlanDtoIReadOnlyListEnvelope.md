# SubscriptionPlanDtoIReadOnlyListEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ActivityId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Result** | Pointer to [**[]SubscriptionPlanDto**](SubscriptionPlanDto.md) |  | [optional] 

## Methods

### NewSubscriptionPlanDtoIReadOnlyListEnvelope

`func NewSubscriptionPlanDtoIReadOnlyListEnvelope() *SubscriptionPlanDtoIReadOnlyListEnvelope`

NewSubscriptionPlanDtoIReadOnlyListEnvelope instantiates a new SubscriptionPlanDtoIReadOnlyListEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanDtoIReadOnlyListEnvelopeWithDefaults

`func NewSubscriptionPlanDtoIReadOnlyListEnvelopeWithDefaults() *SubscriptionPlanDtoIReadOnlyListEnvelope`

NewSubscriptionPlanDtoIReadOnlyListEnvelopeWithDefaults instantiates a new SubscriptionPlanDtoIReadOnlyListEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCorrelationId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetTimestamp

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActivityId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetResult

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetResult() []SubscriptionPlanDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) GetResultOk() (*[]SubscriptionPlanDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetResult(v []SubscriptionPlanDto)`

SetResult sets Result field to given value.

### HasResult

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *SubscriptionPlanDtoIReadOnlyListEnvelope) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


