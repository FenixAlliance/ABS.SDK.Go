# PrepareAndCreateAsyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to ***os.File** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **string** |  | [optional] 
**RoutingMode** | Pointer to **string** |  | [optional] 
**ExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**ExternalReference** | Pointer to **string** |  | [optional] 
**Signers** | Pointer to **string** |  | [optional] 

## Methods

### NewPrepareAndCreateAsyncRequest

`func NewPrepareAndCreateAsyncRequest() *PrepareAndCreateAsyncRequest`

NewPrepareAndCreateAsyncRequest instantiates a new PrepareAndCreateAsyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAndCreateAsyncRequestWithDefaults

`func NewPrepareAndCreateAsyncRequestWithDefaults() *PrepareAndCreateAsyncRequest`

NewPrepareAndCreateAsyncRequestWithDefaults instantiates a new PrepareAndCreateAsyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *PrepareAndCreateAsyncRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PrepareAndCreateAsyncRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PrepareAndCreateAsyncRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PrepareAndCreateAsyncRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetTitle

`func (o *PrepareAndCreateAsyncRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrepareAndCreateAsyncRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrepareAndCreateAsyncRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PrepareAndCreateAsyncRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContactId

`func (o *PrepareAndCreateAsyncRequest) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *PrepareAndCreateAsyncRequest) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *PrepareAndCreateAsyncRequest) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *PrepareAndCreateAsyncRequest) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetRoutingMode

`func (o *PrepareAndCreateAsyncRequest) GetRoutingMode() string`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *PrepareAndCreateAsyncRequest) GetRoutingModeOk() (*string, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *PrepareAndCreateAsyncRequest) SetRoutingMode(v string)`

SetRoutingMode sets RoutingMode field to given value.

### HasRoutingMode

`func (o *PrepareAndCreateAsyncRequest) HasRoutingMode() bool`

HasRoutingMode returns a boolean if a field has been set.

### GetExpiresAtUtc

`func (o *PrepareAndCreateAsyncRequest) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *PrepareAndCreateAsyncRequest) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *PrepareAndCreateAsyncRequest) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *PrepareAndCreateAsyncRequest) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### GetMessage

`func (o *PrepareAndCreateAsyncRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrepareAndCreateAsyncRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrepareAndCreateAsyncRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PrepareAndCreateAsyncRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCorrelationId

`func (o *PrepareAndCreateAsyncRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *PrepareAndCreateAsyncRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *PrepareAndCreateAsyncRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *PrepareAndCreateAsyncRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetExternalReference

`func (o *PrepareAndCreateAsyncRequest) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PrepareAndCreateAsyncRequest) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PrepareAndCreateAsyncRequest) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PrepareAndCreateAsyncRequest) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetSigners

`func (o *PrepareAndCreateAsyncRequest) GetSigners() string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *PrepareAndCreateAsyncRequest) GetSignersOk() (*string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *PrepareAndCreateAsyncRequest) SetSigners(v string)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *PrepareAndCreateAsyncRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


