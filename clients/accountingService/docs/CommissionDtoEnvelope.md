# CommissionDtoEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ActivityId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Result** | Pointer to [**CommissionDto**](CommissionDto.md) |  | [optional] 

## Methods

### NewCommissionDtoEnvelope

`func NewCommissionDtoEnvelope() *CommissionDtoEnvelope`

NewCommissionDtoEnvelope instantiates a new CommissionDtoEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionDtoEnvelopeWithDefaults

`func NewCommissionDtoEnvelopeWithDefaults() *CommissionDtoEnvelope`

NewCommissionDtoEnvelopeWithDefaults instantiates a new CommissionDtoEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *CommissionDtoEnvelope) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *CommissionDtoEnvelope) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *CommissionDtoEnvelope) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *CommissionDtoEnvelope) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CommissionDtoEnvelope) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CommissionDtoEnvelope) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CommissionDtoEnvelope) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CommissionDtoEnvelope) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CommissionDtoEnvelope) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CommissionDtoEnvelope) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCorrelationId

`func (o *CommissionDtoEnvelope) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *CommissionDtoEnvelope) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *CommissionDtoEnvelope) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *CommissionDtoEnvelope) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *CommissionDtoEnvelope) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *CommissionDtoEnvelope) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetTimestamp

`func (o *CommissionDtoEnvelope) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CommissionDtoEnvelope) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CommissionDtoEnvelope) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CommissionDtoEnvelope) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActivityId

`func (o *CommissionDtoEnvelope) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *CommissionDtoEnvelope) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *CommissionDtoEnvelope) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *CommissionDtoEnvelope) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *CommissionDtoEnvelope) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *CommissionDtoEnvelope) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetResult

`func (o *CommissionDtoEnvelope) GetResult() CommissionDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CommissionDtoEnvelope) GetResultOk() (*CommissionDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CommissionDtoEnvelope) SetResult(v CommissionDto)`

SetResult sets Result field to given value.

### HasResult

`func (o *CommissionDtoEnvelope) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


