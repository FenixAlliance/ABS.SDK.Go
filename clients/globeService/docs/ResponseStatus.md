# ResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**CorrelationID** | Pointer to **NullableString** |  | [optional] 
**UtcTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResponseStatus

`func NewResponseStatus() *ResponseStatus`

NewResponseStatus instantiates a new ResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStatusWithDefaults

`func NewResponseStatusWithDefaults() *ResponseStatus`

NewResponseStatusWithDefaults instantiates a new ResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseStatus) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseStatus) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseStatus) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ResponseStatus) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseStatus) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseStatus) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCorrelationID

`func (o *ResponseStatus) GetCorrelationID() string`

GetCorrelationID returns the CorrelationID field if non-nil, zero value otherwise.

### GetCorrelationIDOk

`func (o *ResponseStatus) GetCorrelationIDOk() (*string, bool)`

GetCorrelationIDOk returns a tuple with the CorrelationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationID

`func (o *ResponseStatus) SetCorrelationID(v string)`

SetCorrelationID sets CorrelationID field to given value.

### HasCorrelationID

`func (o *ResponseStatus) HasCorrelationID() bool`

HasCorrelationID returns a boolean if a field has been set.

### SetCorrelationIDNil

`func (o *ResponseStatus) SetCorrelationIDNil(b bool)`

 SetCorrelationIDNil sets the value for CorrelationID to be an explicit nil

### UnsetCorrelationID
`func (o *ResponseStatus) UnsetCorrelationID()`

UnsetCorrelationID ensures that no value is present for CorrelationID, not even an explicit nil
### GetUtcTimestamp

`func (o *ResponseStatus) GetUtcTimestamp() time.Time`

GetUtcTimestamp returns the UtcTimestamp field if non-nil, zero value otherwise.

### GetUtcTimestampOk

`func (o *ResponseStatus) GetUtcTimestampOk() (*time.Time, bool)`

GetUtcTimestampOk returns a tuple with the UtcTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcTimestamp

`func (o *ResponseStatus) SetUtcTimestamp(v time.Time)`

SetUtcTimestamp sets UtcTimestamp field to given value.

### HasUtcTimestamp

`func (o *ResponseStatus) HasUtcTimestamp() bool`

HasUtcTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


