# AuthResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **map[string]interface{}** |  | [optional] 
**TenantId** | Pointer to **map[string]interface{}** |  | [optional] 
**PortalId** | Pointer to **map[string]interface{}** |  | [optional] 
**ApplicationId** | Pointer to **map[string]interface{}** |  | [optional] 
**EnrollmentId** | Pointer to **map[string]interface{}** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthorizationResult

`func NewAuthorizationResult() *AuthResult`

NewAuthorizationResult instantiates a new AuthResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResultWithDefaults

`func NewAuthorizationResultWithDefaults() *AuthResult`

NewAuthorizationResultWithDefaults instantiates a new AuthResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AuthResult) GetUserId() map[string]interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthResult) GetUserIdOk() (*map[string]interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthResult) SetUserId(v map[string]interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *AuthResult) GetTenantId() map[string]interface{}`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AuthResult) GetTenantIdOk() (*map[string]interface{}, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AuthResult) SetTenantId(v map[string]interface{})`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AuthResult) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPortalId

`func (o *AuthResult) GetPortalId() map[string]interface{}`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *AuthResult) GetPortalIdOk() (*map[string]interface{}, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *AuthResult) SetPortalId(v map[string]interface{})`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *AuthResult) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetApplicationId

`func (o *AuthResult) GetApplicationId() map[string]interface{}`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AuthResult) GetApplicationIdOk() (*map[string]interface{}, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AuthResult) SetApplicationId(v map[string]interface{})`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AuthResult) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *AuthResult) GetEnrollmentId() map[string]interface{}`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AuthResult) GetEnrollmentIdOk() (*map[string]interface{}, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AuthResult) SetEnrollmentId(v map[string]interface{})`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AuthResult) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetCorrelationId

`func (o *AuthResult) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *AuthResult) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *AuthResult) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *AuthResult) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *AuthResult) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *AuthResult) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetScopes

`func (o *AuthResult) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthResult) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthResult) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthResult) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *AuthResult) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *AuthResult) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetError

`func (o *AuthResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AuthResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AuthResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


