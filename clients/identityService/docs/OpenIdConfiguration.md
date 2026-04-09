# OpenIdConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **NullableString** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **NullableString** |  | [optional] 
**TokenEndpoint** | Pointer to **NullableString** |  | [optional] 
**EndSessionEndpoint** | Pointer to **NullableString** |  | [optional] 
**JwksUri** | Pointer to **NullableString** |  | [optional] 
**ResponseModesSupported** | Pointer to **[]string** |  | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** |  | [optional] 
**ScopesSupported** | Pointer to **[]string** |  | [optional] 
**SubjectTypesSupported** | Pointer to **[]string** |  | [optional] 
**IdTokenSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**ClaimsSupported** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOpenIdConfiguration

`func NewOpenIdConfiguration() *OpenIdConfiguration`

NewOpenIdConfiguration instantiates a new OpenIdConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConfigurationWithDefaults

`func NewOpenIdConfigurationWithDefaults() *OpenIdConfiguration`

NewOpenIdConfigurationWithDefaults instantiates a new OpenIdConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *OpenIdConfiguration) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenIdConfiguration) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenIdConfiguration) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OpenIdConfiguration) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *OpenIdConfiguration) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *OpenIdConfiguration) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetAuthorizationEndpoint

`func (o *OpenIdConfiguration) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *OpenIdConfiguration) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *OpenIdConfiguration) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *OpenIdConfiguration) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### SetAuthorizationEndpointNil

`func (o *OpenIdConfiguration) SetAuthorizationEndpointNil(b bool)`

 SetAuthorizationEndpointNil sets the value for AuthorizationEndpoint to be an explicit nil

### UnsetAuthorizationEndpoint
`func (o *OpenIdConfiguration) UnsetAuthorizationEndpoint()`

UnsetAuthorizationEndpoint ensures that no value is present for AuthorizationEndpoint, not even an explicit nil
### GetTokenEndpoint

`func (o *OpenIdConfiguration) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OpenIdConfiguration) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OpenIdConfiguration) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OpenIdConfiguration) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### SetTokenEndpointNil

`func (o *OpenIdConfiguration) SetTokenEndpointNil(b bool)`

 SetTokenEndpointNil sets the value for TokenEndpoint to be an explicit nil

### UnsetTokenEndpoint
`func (o *OpenIdConfiguration) UnsetTokenEndpoint()`

UnsetTokenEndpoint ensures that no value is present for TokenEndpoint, not even an explicit nil
### GetEndSessionEndpoint

`func (o *OpenIdConfiguration) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *OpenIdConfiguration) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *OpenIdConfiguration) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *OpenIdConfiguration) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### SetEndSessionEndpointNil

`func (o *OpenIdConfiguration) SetEndSessionEndpointNil(b bool)`

 SetEndSessionEndpointNil sets the value for EndSessionEndpoint to be an explicit nil

### UnsetEndSessionEndpoint
`func (o *OpenIdConfiguration) UnsetEndSessionEndpoint()`

UnsetEndSessionEndpoint ensures that no value is present for EndSessionEndpoint, not even an explicit nil
### GetJwksUri

`func (o *OpenIdConfiguration) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OpenIdConfiguration) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OpenIdConfiguration) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OpenIdConfiguration) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### SetJwksUriNil

`func (o *OpenIdConfiguration) SetJwksUriNil(b bool)`

 SetJwksUriNil sets the value for JwksUri to be an explicit nil

### UnsetJwksUri
`func (o *OpenIdConfiguration) UnsetJwksUri()`

UnsetJwksUri ensures that no value is present for JwksUri, not even an explicit nil
### GetResponseModesSupported

`func (o *OpenIdConfiguration) GetResponseModesSupported() []string`

GetResponseModesSupported returns the ResponseModesSupported field if non-nil, zero value otherwise.

### GetResponseModesSupportedOk

`func (o *OpenIdConfiguration) GetResponseModesSupportedOk() (*[]string, bool)`

GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModesSupported

`func (o *OpenIdConfiguration) SetResponseModesSupported(v []string)`

SetResponseModesSupported sets ResponseModesSupported field to given value.

### HasResponseModesSupported

`func (o *OpenIdConfiguration) HasResponseModesSupported() bool`

HasResponseModesSupported returns a boolean if a field has been set.

### SetResponseModesSupportedNil

`func (o *OpenIdConfiguration) SetResponseModesSupportedNil(b bool)`

 SetResponseModesSupportedNil sets the value for ResponseModesSupported to be an explicit nil

### UnsetResponseModesSupported
`func (o *OpenIdConfiguration) UnsetResponseModesSupported()`

UnsetResponseModesSupported ensures that no value is present for ResponseModesSupported, not even an explicit nil
### GetResponseTypesSupported

`func (o *OpenIdConfiguration) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *OpenIdConfiguration) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *OpenIdConfiguration) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *OpenIdConfiguration) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### SetResponseTypesSupportedNil

`func (o *OpenIdConfiguration) SetResponseTypesSupportedNil(b bool)`

 SetResponseTypesSupportedNil sets the value for ResponseTypesSupported to be an explicit nil

### UnsetResponseTypesSupported
`func (o *OpenIdConfiguration) UnsetResponseTypesSupported()`

UnsetResponseTypesSupported ensures that no value is present for ResponseTypesSupported, not even an explicit nil
### GetScopesSupported

`func (o *OpenIdConfiguration) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *OpenIdConfiguration) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *OpenIdConfiguration) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *OpenIdConfiguration) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### SetScopesSupportedNil

`func (o *OpenIdConfiguration) SetScopesSupportedNil(b bool)`

 SetScopesSupportedNil sets the value for ScopesSupported to be an explicit nil

### UnsetScopesSupported
`func (o *OpenIdConfiguration) UnsetScopesSupported()`

UnsetScopesSupported ensures that no value is present for ScopesSupported, not even an explicit nil
### GetSubjectTypesSupported

`func (o *OpenIdConfiguration) GetSubjectTypesSupported() []string`

GetSubjectTypesSupported returns the SubjectTypesSupported field if non-nil, zero value otherwise.

### GetSubjectTypesSupportedOk

`func (o *OpenIdConfiguration) GetSubjectTypesSupportedOk() (*[]string, bool)`

GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTypesSupported

`func (o *OpenIdConfiguration) SetSubjectTypesSupported(v []string)`

SetSubjectTypesSupported sets SubjectTypesSupported field to given value.

### HasSubjectTypesSupported

`func (o *OpenIdConfiguration) HasSubjectTypesSupported() bool`

HasSubjectTypesSupported returns a boolean if a field has been set.

### SetSubjectTypesSupportedNil

`func (o *OpenIdConfiguration) SetSubjectTypesSupportedNil(b bool)`

 SetSubjectTypesSupportedNil sets the value for SubjectTypesSupported to be an explicit nil

### UnsetSubjectTypesSupported
`func (o *OpenIdConfiguration) UnsetSubjectTypesSupported()`

UnsetSubjectTypesSupported ensures that no value is present for SubjectTypesSupported, not even an explicit nil
### GetIdTokenSigningAlgValuesSupported

`func (o *OpenIdConfiguration) GetIdTokenSigningAlgValuesSupported() []string`

GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgValuesSupportedOk

`func (o *OpenIdConfiguration) GetIdTokenSigningAlgValuesSupportedOk() (*[]string, bool)`

GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgValuesSupported

`func (o *OpenIdConfiguration) SetIdTokenSigningAlgValuesSupported(v []string)`

SetIdTokenSigningAlgValuesSupported sets IdTokenSigningAlgValuesSupported field to given value.

### HasIdTokenSigningAlgValuesSupported

`func (o *OpenIdConfiguration) HasIdTokenSigningAlgValuesSupported() bool`

HasIdTokenSigningAlgValuesSupported returns a boolean if a field has been set.

### SetIdTokenSigningAlgValuesSupportedNil

`func (o *OpenIdConfiguration) SetIdTokenSigningAlgValuesSupportedNil(b bool)`

 SetIdTokenSigningAlgValuesSupportedNil sets the value for IdTokenSigningAlgValuesSupported to be an explicit nil

### UnsetIdTokenSigningAlgValuesSupported
`func (o *OpenIdConfiguration) UnsetIdTokenSigningAlgValuesSupported()`

UnsetIdTokenSigningAlgValuesSupported ensures that no value is present for IdTokenSigningAlgValuesSupported, not even an explicit nil
### GetTokenEndpointAuthMethodsSupported

`func (o *OpenIdConfiguration) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *OpenIdConfiguration) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *OpenIdConfiguration) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *OpenIdConfiguration) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.

### SetTokenEndpointAuthMethodsSupportedNil

`func (o *OpenIdConfiguration) SetTokenEndpointAuthMethodsSupportedNil(b bool)`

 SetTokenEndpointAuthMethodsSupportedNil sets the value for TokenEndpointAuthMethodsSupported to be an explicit nil

### UnsetTokenEndpointAuthMethodsSupported
`func (o *OpenIdConfiguration) UnsetTokenEndpointAuthMethodsSupported()`

UnsetTokenEndpointAuthMethodsSupported ensures that no value is present for TokenEndpointAuthMethodsSupported, not even an explicit nil
### GetClaimsSupported

`func (o *OpenIdConfiguration) GetClaimsSupported() []string`

GetClaimsSupported returns the ClaimsSupported field if non-nil, zero value otherwise.

### GetClaimsSupportedOk

`func (o *OpenIdConfiguration) GetClaimsSupportedOk() (*[]string, bool)`

GetClaimsSupportedOk returns a tuple with the ClaimsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsSupported

`func (o *OpenIdConfiguration) SetClaimsSupported(v []string)`

SetClaimsSupported sets ClaimsSupported field to given value.

### HasClaimsSupported

`func (o *OpenIdConfiguration) HasClaimsSupported() bool`

HasClaimsSupported returns a boolean if a field has been set.

### SetClaimsSupportedNil

`func (o *OpenIdConfiguration) SetClaimsSupportedNil(b bool)`

 SetClaimsSupportedNil sets the value for ClaimsSupported to be an explicit nil

### UnsetClaimsSupported
`func (o *OpenIdConfiguration) UnsetClaimsSupported()`

UnsetClaimsSupported ensures that no value is present for ClaimsSupported, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


