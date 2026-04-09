# TenantTeamDtoEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ActivityId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Result** | Pointer to [**TenantTeamDto**](TenantTeamDto.md) |  | [optional] 

## Methods

### NewTenantTeamDtoEnvelope

`func NewTenantTeamDtoEnvelope() *TenantTeamDtoEnvelope`

NewTenantTeamDtoEnvelope instantiates a new TenantTeamDtoEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamDtoEnvelopeWithDefaults

`func NewTenantTeamDtoEnvelopeWithDefaults() *TenantTeamDtoEnvelope`

NewTenantTeamDtoEnvelopeWithDefaults instantiates a new TenantTeamDtoEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *TenantTeamDtoEnvelope) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *TenantTeamDtoEnvelope) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *TenantTeamDtoEnvelope) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *TenantTeamDtoEnvelope) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TenantTeamDtoEnvelope) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TenantTeamDtoEnvelope) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TenantTeamDtoEnvelope) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TenantTeamDtoEnvelope) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TenantTeamDtoEnvelope) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TenantTeamDtoEnvelope) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCorrelationId

`func (o *TenantTeamDtoEnvelope) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TenantTeamDtoEnvelope) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TenantTeamDtoEnvelope) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TenantTeamDtoEnvelope) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *TenantTeamDtoEnvelope) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *TenantTeamDtoEnvelope) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetTimestamp

`func (o *TenantTeamDtoEnvelope) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamDtoEnvelope) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamDtoEnvelope) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamDtoEnvelope) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActivityId

`func (o *TenantTeamDtoEnvelope) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TenantTeamDtoEnvelope) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TenantTeamDtoEnvelope) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TenantTeamDtoEnvelope) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *TenantTeamDtoEnvelope) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *TenantTeamDtoEnvelope) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetResult

`func (o *TenantTeamDtoEnvelope) GetResult() TenantTeamDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TenantTeamDtoEnvelope) GetResultOk() (*TenantTeamDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TenantTeamDtoEnvelope) SetResult(v TenantTeamDto)`

SetResult sets Result field to given value.

### HasResult

`func (o *TenantTeamDtoEnvelope) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


