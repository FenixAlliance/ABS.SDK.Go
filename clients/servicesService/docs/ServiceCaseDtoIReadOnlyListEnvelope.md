# ServiceCaseDtoIReadOnlyListEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**ActivityId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Result** | Pointer to [**[]ServiceCaseDto**](ServiceCaseDto.md) |  | [optional] 

## Methods

### NewServiceCaseDtoIReadOnlyListEnvelope

`func NewServiceCaseDtoIReadOnlyListEnvelope() *ServiceCaseDtoIReadOnlyListEnvelope`

NewServiceCaseDtoIReadOnlyListEnvelope instantiates a new ServiceCaseDtoIReadOnlyListEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCaseDtoIReadOnlyListEnvelopeWithDefaults

`func NewServiceCaseDtoIReadOnlyListEnvelopeWithDefaults() *ServiceCaseDtoIReadOnlyListEnvelope`

NewServiceCaseDtoIReadOnlyListEnvelopeWithDefaults instantiates a new ServiceCaseDtoIReadOnlyListEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ServiceCaseDtoIReadOnlyListEnvelope) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCorrelationId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *ServiceCaseDtoIReadOnlyListEnvelope) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetTimestamp

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActivityId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ServiceCaseDtoIReadOnlyListEnvelope) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetResult

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetResult() []ServiceCaseDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) GetResultOk() (*[]ServiceCaseDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetResult(v []ServiceCaseDto)`

SetResult sets Result field to given value.

### HasResult

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ServiceCaseDtoIReadOnlyListEnvelope) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ServiceCaseDtoIReadOnlyListEnvelope) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


