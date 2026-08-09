# CreateProviderWebhookRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderCode** | Pointer to **NullableString** |  | [optional] 
**ExternalAccountId** | Pointer to **NullableString** |  | [optional] 
**WebhookSigningSecret** | Pointer to **NullableString** |  | [optional] 
**CredentialMode** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateProviderWebhookRegistrationRequest

`func NewCreateProviderWebhookRegistrationRequest() *CreateProviderWebhookRegistrationRequest`

NewCreateProviderWebhookRegistrationRequest instantiates a new CreateProviderWebhookRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProviderWebhookRegistrationRequestWithDefaults

`func NewCreateProviderWebhookRegistrationRequestWithDefaults() *CreateProviderWebhookRegistrationRequest`

NewCreateProviderWebhookRegistrationRequestWithDefaults instantiates a new CreateProviderWebhookRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderCode

`func (o *CreateProviderWebhookRegistrationRequest) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *CreateProviderWebhookRegistrationRequest) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *CreateProviderWebhookRegistrationRequest) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *CreateProviderWebhookRegistrationRequest) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *CreateProviderWebhookRegistrationRequest) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *CreateProviderWebhookRegistrationRequest) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetExternalAccountId

`func (o *CreateProviderWebhookRegistrationRequest) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *CreateProviderWebhookRegistrationRequest) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *CreateProviderWebhookRegistrationRequest) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.

### HasExternalAccountId

`func (o *CreateProviderWebhookRegistrationRequest) HasExternalAccountId() bool`

HasExternalAccountId returns a boolean if a field has been set.

### SetExternalAccountIdNil

`func (o *CreateProviderWebhookRegistrationRequest) SetExternalAccountIdNil(b bool)`

 SetExternalAccountIdNil sets the value for ExternalAccountId to be an explicit nil

### UnsetExternalAccountId
`func (o *CreateProviderWebhookRegistrationRequest) UnsetExternalAccountId()`

UnsetExternalAccountId ensures that no value is present for ExternalAccountId, not even an explicit nil
### GetWebhookSigningSecret

`func (o *CreateProviderWebhookRegistrationRequest) GetWebhookSigningSecret() string`

GetWebhookSigningSecret returns the WebhookSigningSecret field if non-nil, zero value otherwise.

### GetWebhookSigningSecretOk

`func (o *CreateProviderWebhookRegistrationRequest) GetWebhookSigningSecretOk() (*string, bool)`

GetWebhookSigningSecretOk returns a tuple with the WebhookSigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSigningSecret

`func (o *CreateProviderWebhookRegistrationRequest) SetWebhookSigningSecret(v string)`

SetWebhookSigningSecret sets WebhookSigningSecret field to given value.

### HasWebhookSigningSecret

`func (o *CreateProviderWebhookRegistrationRequest) HasWebhookSigningSecret() bool`

HasWebhookSigningSecret returns a boolean if a field has been set.

### SetWebhookSigningSecretNil

`func (o *CreateProviderWebhookRegistrationRequest) SetWebhookSigningSecretNil(b bool)`

 SetWebhookSigningSecretNil sets the value for WebhookSigningSecret to be an explicit nil

### UnsetWebhookSigningSecret
`func (o *CreateProviderWebhookRegistrationRequest) UnsetWebhookSigningSecret()`

UnsetWebhookSigningSecret ensures that no value is present for WebhookSigningSecret, not even an explicit nil
### GetCredentialMode

`func (o *CreateProviderWebhookRegistrationRequest) GetCredentialMode() string`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *CreateProviderWebhookRegistrationRequest) GetCredentialModeOk() (*string, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *CreateProviderWebhookRegistrationRequest) SetCredentialMode(v string)`

SetCredentialMode sets CredentialMode field to given value.

### HasCredentialMode

`func (o *CreateProviderWebhookRegistrationRequest) HasCredentialMode() bool`

HasCredentialMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


