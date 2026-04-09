# AuthorizationResult

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

`func NewAuthorizationResult() *AuthorizationResult`

NewAuthorizationResult instantiates a new AuthorizationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResultWithDefaults

`func NewAuthorizationResultWithDefaults() *AuthorizationResult`

NewAuthorizationResultWithDefaults instantiates a new AuthorizationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AuthorizationResult) GetUserId() map[string]interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthorizationResult) GetUserIdOk() (*map[string]interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthorizationResult) SetUserId(v map[string]interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthorizationResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *AuthorizationResult) GetTenantId() map[string]interface{}`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AuthorizationResult) GetTenantIdOk() (*map[string]interface{}, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AuthorizationResult) SetTenantId(v map[string]interface{})`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AuthorizationResult) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPortalId

`func (o *AuthorizationResult) GetPortalId() map[string]interface{}`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *AuthorizationResult) GetPortalIdOk() (*map[string]interface{}, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *AuthorizationResult) SetPortalId(v map[string]interface{})`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *AuthorizationResult) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetApplicationId

`func (o *AuthorizationResult) GetApplicationId() map[string]interface{}`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AuthorizationResult) GetApplicationIdOk() (*map[string]interface{}, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AuthorizationResult) SetApplicationId(v map[string]interface{})`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AuthorizationResult) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *AuthorizationResult) GetEnrollmentId() map[string]interface{}`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AuthorizationResult) GetEnrollmentIdOk() (*map[string]interface{}, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AuthorizationResult) SetEnrollmentId(v map[string]interface{})`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AuthorizationResult) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetCorrelationId

`func (o *AuthorizationResult) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *AuthorizationResult) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *AuthorizationResult) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *AuthorizationResult) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *AuthorizationResult) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *AuthorizationResult) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetScopes

`func (o *AuthorizationResult) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationResult) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationResult) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationResult) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *AuthorizationResult) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *AuthorizationResult) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetError

`func (o *AuthorizationResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthorizationResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthorizationResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthorizationResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AuthorizationResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AuthorizationResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


