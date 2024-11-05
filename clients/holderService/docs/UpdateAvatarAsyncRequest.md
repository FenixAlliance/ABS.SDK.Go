# UpdateAvatarAsyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUpdateAvatarAsyncRequest

`func NewUpdateAvatarAsyncRequest() *UpdateAvatarAsyncRequest`

NewUpdateAvatarAsyncRequest instantiates a new UpdateAvatarAsyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAvatarAsyncRequestWithDefaults

`func NewUpdateAvatarAsyncRequestWithDefaults() *UpdateAvatarAsyncRequest`

NewUpdateAvatarAsyncRequestWithDefaults instantiates a new UpdateAvatarAsyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UpdateAvatarAsyncRequest) GetAvatar() *os.File`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UpdateAvatarAsyncRequest) GetAvatarOk() (**os.File, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UpdateAvatarAsyncRequest) SetAvatar(v *os.File)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UpdateAvatarAsyncRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


