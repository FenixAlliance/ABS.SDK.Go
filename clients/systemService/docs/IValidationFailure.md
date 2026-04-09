# IValidationFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** |  | [optional] 
**HowToResolve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIValidationFailure

`func NewIValidationFailure() *IValidationFailure`

NewIValidationFailure instantiates a new IValidationFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIValidationFailureWithDefaults

`func NewIValidationFailureWithDefaults() *IValidationFailure`

NewIValidationFailureWithDefaults instantiates a new IValidationFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IValidationFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IValidationFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IValidationFailure) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IValidationFailure) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *IValidationFailure) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *IValidationFailure) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetHowToResolve

`func (o *IValidationFailure) GetHowToResolve() string`

GetHowToResolve returns the HowToResolve field if non-nil, zero value otherwise.

### GetHowToResolveOk

`func (o *IValidationFailure) GetHowToResolveOk() (*string, bool)`

GetHowToResolveOk returns a tuple with the HowToResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHowToResolve

`func (o *IValidationFailure) SetHowToResolve(v string)`

SetHowToResolve sets HowToResolve field to given value.

### HasHowToResolve

`func (o *IValidationFailure) HasHowToResolve() bool`

HasHowToResolve returns a boolean if a field has been set.

### SetHowToResolveNil

`func (o *IValidationFailure) SetHowToResolveNil(b bool)`

 SetHowToResolveNil sets the value for HowToResolve to be an explicit nil

### UnsetHowToResolve
`func (o *IValidationFailure) UnsetHowToResolve()`

UnsetHowToResolve ensures that no value is present for HowToResolve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


