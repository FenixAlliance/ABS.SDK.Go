# CreateSigningRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingMode** | Pointer to **string** |  | [optional] 
**ExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSigningRequestDto

`func NewCreateSigningRequestDto() *CreateSigningRequestDto`

NewCreateSigningRequestDto instantiates a new CreateSigningRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSigningRequestDtoWithDefaults

`func NewCreateSigningRequestDtoWithDefaults() *CreateSigningRequestDto`

NewCreateSigningRequestDtoWithDefaults instantiates a new CreateSigningRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingMode

`func (o *CreateSigningRequestDto) GetRoutingMode() string`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *CreateSigningRequestDto) GetRoutingModeOk() (*string, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *CreateSigningRequestDto) SetRoutingMode(v string)`

SetRoutingMode sets RoutingMode field to given value.

### HasRoutingMode

`func (o *CreateSigningRequestDto) HasRoutingMode() bool`

HasRoutingMode returns a boolean if a field has been set.

### GetExpiresAtUtc

`func (o *CreateSigningRequestDto) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *CreateSigningRequestDto) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *CreateSigningRequestDto) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *CreateSigningRequestDto) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *CreateSigningRequestDto) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *CreateSigningRequestDto) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetMessage

`func (o *CreateSigningRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateSigningRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateSigningRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateSigningRequestDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateSigningRequestDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateSigningRequestDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCorrelationId

`func (o *CreateSigningRequestDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *CreateSigningRequestDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *CreateSigningRequestDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *CreateSigningRequestDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *CreateSigningRequestDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *CreateSigningRequestDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *CreateSigningRequestDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *CreateSigningRequestDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *CreateSigningRequestDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *CreateSigningRequestDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *CreateSigningRequestDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *CreateSigningRequestDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


