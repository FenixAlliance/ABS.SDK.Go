# TenantDtoListEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ActivityId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Result** | Pointer to [**[]TenantDto**](TenantDto.md) |  | [optional] 

## Methods

### NewTenantDtoListEnvelope

`func NewTenantDtoListEnvelope() *TenantDtoListEnvelope`

NewTenantDtoListEnvelope instantiates a new TenantDtoListEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDtoListEnvelopeWithDefaults

`func NewTenantDtoListEnvelopeWithDefaults() *TenantDtoListEnvelope`

NewTenantDtoListEnvelopeWithDefaults instantiates a new TenantDtoListEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *TenantDtoListEnvelope) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *TenantDtoListEnvelope) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *TenantDtoListEnvelope) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *TenantDtoListEnvelope) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TenantDtoListEnvelope) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TenantDtoListEnvelope) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TenantDtoListEnvelope) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TenantDtoListEnvelope) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TenantDtoListEnvelope) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TenantDtoListEnvelope) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCorrelationId

`func (o *TenantDtoListEnvelope) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TenantDtoListEnvelope) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TenantDtoListEnvelope) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TenantDtoListEnvelope) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *TenantDtoListEnvelope) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *TenantDtoListEnvelope) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetTimestamp

`func (o *TenantDtoListEnvelope) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantDtoListEnvelope) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantDtoListEnvelope) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantDtoListEnvelope) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActivityId

`func (o *TenantDtoListEnvelope) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TenantDtoListEnvelope) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TenantDtoListEnvelope) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TenantDtoListEnvelope) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *TenantDtoListEnvelope) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *TenantDtoListEnvelope) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetResult

`func (o *TenantDtoListEnvelope) GetResult() []TenantDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TenantDtoListEnvelope) GetResultOk() (*[]TenantDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TenantDtoListEnvelope) SetResult(v []TenantDto)`

SetResult sets Result field to given value.

### HasResult

`func (o *TenantDtoListEnvelope) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TenantDtoListEnvelope) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TenantDtoListEnvelope) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

