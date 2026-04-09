# PortalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**PortalID** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **NullableString** |  | [optional] 
**TenantID** | Pointer to **NullableString** |  | [optional] 
**HomePageID** | Pointer to **NullableString** |  | [optional] 
**BlogPageID** | Pointer to **NullableString** |  | [optional] 
**StorePageID** | Pointer to **NullableString** |  | [optional] 
**BaseEndpoint** | Pointer to **NullableString** |  | [optional] 
**StoreRoutePrefix** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**AuthToken** | Pointer to **NullableString** |  | [optional] 
**AuthTokenType** | Pointer to **NullableString** |  | [optional] 
**AuthTokenExpiration** | Pointer to **int64** |  | [optional] 
**Options** | Pointer to [**PortalOptions**](PortalOptions.md) |  | [optional] 

## Methods

### NewPortalSettings

`func NewPortalSettings() *PortalSettings`

NewPortalSettings instantiates a new PortalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSettingsWithDefaults

`func NewPortalSettingsWithDefaults() *PortalSettings`

NewPortalSettingsWithDefaults instantiates a new PortalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *PortalSettings) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PortalSettings) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PortalSettings) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PortalSettings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPortalID

`func (o *PortalSettings) GetPortalID() string`

GetPortalID returns the PortalID field if non-nil, zero value otherwise.

### GetPortalIDOk

`func (o *PortalSettings) GetPortalIDOk() (*string, bool)`

GetPortalIDOk returns a tuple with the PortalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalID

`func (o *PortalSettings) SetPortalID(v string)`

SetPortalID sets PortalID field to given value.

### HasPortalID

`func (o *PortalSettings) HasPortalID() bool`

HasPortalID returns a boolean if a field has been set.

### SetPortalIDNil

`func (o *PortalSettings) SetPortalIDNil(b bool)`

 SetPortalIDNil sets the value for PortalID to be an explicit nil

### UnsetPortalID
`func (o *PortalSettings) UnsetPortalID()`

UnsetPortalID ensures that no value is present for PortalID, not even an explicit nil
### GetScopes

`func (o *PortalSettings) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PortalSettings) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PortalSettings) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PortalSettings) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *PortalSettings) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PortalSettings) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenantID

`func (o *PortalSettings) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *PortalSettings) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *PortalSettings) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *PortalSettings) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### SetTenantIDNil

`func (o *PortalSettings) SetTenantIDNil(b bool)`

 SetTenantIDNil sets the value for TenantID to be an explicit nil

### UnsetTenantID
`func (o *PortalSettings) UnsetTenantID()`

UnsetTenantID ensures that no value is present for TenantID, not even an explicit nil
### GetHomePageID

`func (o *PortalSettings) GetHomePageID() string`

GetHomePageID returns the HomePageID field if non-nil, zero value otherwise.

### GetHomePageIDOk

`func (o *PortalSettings) GetHomePageIDOk() (*string, bool)`

GetHomePageIDOk returns a tuple with the HomePageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageID

`func (o *PortalSettings) SetHomePageID(v string)`

SetHomePageID sets HomePageID field to given value.

### HasHomePageID

`func (o *PortalSettings) HasHomePageID() bool`

HasHomePageID returns a boolean if a field has been set.

### SetHomePageIDNil

`func (o *PortalSettings) SetHomePageIDNil(b bool)`

 SetHomePageIDNil sets the value for HomePageID to be an explicit nil

### UnsetHomePageID
`func (o *PortalSettings) UnsetHomePageID()`

UnsetHomePageID ensures that no value is present for HomePageID, not even an explicit nil
### GetBlogPageID

`func (o *PortalSettings) GetBlogPageID() string`

GetBlogPageID returns the BlogPageID field if non-nil, zero value otherwise.

### GetBlogPageIDOk

`func (o *PortalSettings) GetBlogPageIDOk() (*string, bool)`

GetBlogPageIDOk returns a tuple with the BlogPageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPageID

`func (o *PortalSettings) SetBlogPageID(v string)`

SetBlogPageID sets BlogPageID field to given value.

### HasBlogPageID

`func (o *PortalSettings) HasBlogPageID() bool`

HasBlogPageID returns a boolean if a field has been set.

### SetBlogPageIDNil

`func (o *PortalSettings) SetBlogPageIDNil(b bool)`

 SetBlogPageIDNil sets the value for BlogPageID to be an explicit nil

### UnsetBlogPageID
`func (o *PortalSettings) UnsetBlogPageID()`

UnsetBlogPageID ensures that no value is present for BlogPageID, not even an explicit nil
### GetStorePageID

`func (o *PortalSettings) GetStorePageID() string`

GetStorePageID returns the StorePageID field if non-nil, zero value otherwise.

### GetStorePageIDOk

`func (o *PortalSettings) GetStorePageIDOk() (*string, bool)`

GetStorePageIDOk returns a tuple with the StorePageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePageID

`func (o *PortalSettings) SetStorePageID(v string)`

SetStorePageID sets StorePageID field to given value.

### HasStorePageID

`func (o *PortalSettings) HasStorePageID() bool`

HasStorePageID returns a boolean if a field has been set.

### SetStorePageIDNil

`func (o *PortalSettings) SetStorePageIDNil(b bool)`

 SetStorePageIDNil sets the value for StorePageID to be an explicit nil

### UnsetStorePageID
`func (o *PortalSettings) UnsetStorePageID()`

UnsetStorePageID ensures that no value is present for StorePageID, not even an explicit nil
### GetBaseEndpoint

`func (o *PortalSettings) GetBaseEndpoint() string`

GetBaseEndpoint returns the BaseEndpoint field if non-nil, zero value otherwise.

### GetBaseEndpointOk

`func (o *PortalSettings) GetBaseEndpointOk() (*string, bool)`

GetBaseEndpointOk returns a tuple with the BaseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEndpoint

`func (o *PortalSettings) SetBaseEndpoint(v string)`

SetBaseEndpoint sets BaseEndpoint field to given value.

### HasBaseEndpoint

`func (o *PortalSettings) HasBaseEndpoint() bool`

HasBaseEndpoint returns a boolean if a field has been set.

### SetBaseEndpointNil

`func (o *PortalSettings) SetBaseEndpointNil(b bool)`

 SetBaseEndpointNil sets the value for BaseEndpoint to be an explicit nil

### UnsetBaseEndpoint
`func (o *PortalSettings) UnsetBaseEndpoint()`

UnsetBaseEndpoint ensures that no value is present for BaseEndpoint, not even an explicit nil
### GetStoreRoutePrefix

`func (o *PortalSettings) GetStoreRoutePrefix() string`

GetStoreRoutePrefix returns the StoreRoutePrefix field if non-nil, zero value otherwise.

### GetStoreRoutePrefixOk

`func (o *PortalSettings) GetStoreRoutePrefixOk() (*string, bool)`

GetStoreRoutePrefixOk returns a tuple with the StoreRoutePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreRoutePrefix

`func (o *PortalSettings) SetStoreRoutePrefix(v string)`

SetStoreRoutePrefix sets StoreRoutePrefix field to given value.

### HasStoreRoutePrefix

`func (o *PortalSettings) HasStoreRoutePrefix() bool`

HasStoreRoutePrefix returns a boolean if a field has been set.

### SetStoreRoutePrefixNil

`func (o *PortalSettings) SetStoreRoutePrefixNil(b bool)`

 SetStoreRoutePrefixNil sets the value for StoreRoutePrefix to be an explicit nil

### UnsetStoreRoutePrefix
`func (o *PortalSettings) UnsetStoreRoutePrefix()`

UnsetStoreRoutePrefix ensures that no value is present for StoreRoutePrefix, not even an explicit nil
### GetPublicKey

`func (o *PortalSettings) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PortalSettings) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PortalSettings) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PortalSettings) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *PortalSettings) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *PortalSettings) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *PortalSettings) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PortalSettings) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PortalSettings) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PortalSettings) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *PortalSettings) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *PortalSettings) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetAuthToken

`func (o *PortalSettings) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *PortalSettings) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *PortalSettings) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *PortalSettings) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *PortalSettings) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *PortalSettings) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetAuthTokenType

`func (o *PortalSettings) GetAuthTokenType() string`

GetAuthTokenType returns the AuthTokenType field if non-nil, zero value otherwise.

### GetAuthTokenTypeOk

`func (o *PortalSettings) GetAuthTokenTypeOk() (*string, bool)`

GetAuthTokenTypeOk returns a tuple with the AuthTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenType

`func (o *PortalSettings) SetAuthTokenType(v string)`

SetAuthTokenType sets AuthTokenType field to given value.

### HasAuthTokenType

`func (o *PortalSettings) HasAuthTokenType() bool`

HasAuthTokenType returns a boolean if a field has been set.

### SetAuthTokenTypeNil

`func (o *PortalSettings) SetAuthTokenTypeNil(b bool)`

 SetAuthTokenTypeNil sets the value for AuthTokenType to be an explicit nil

### UnsetAuthTokenType
`func (o *PortalSettings) UnsetAuthTokenType()`

UnsetAuthTokenType ensures that no value is present for AuthTokenType, not even an explicit nil
### GetAuthTokenExpiration

`func (o *PortalSettings) GetAuthTokenExpiration() int64`

GetAuthTokenExpiration returns the AuthTokenExpiration field if non-nil, zero value otherwise.

### GetAuthTokenExpirationOk

`func (o *PortalSettings) GetAuthTokenExpirationOk() (*int64, bool)`

GetAuthTokenExpirationOk returns a tuple with the AuthTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenExpiration

`func (o *PortalSettings) SetAuthTokenExpiration(v int64)`

SetAuthTokenExpiration sets AuthTokenExpiration field to given value.

### HasAuthTokenExpiration

`func (o *PortalSettings) HasAuthTokenExpiration() bool`

HasAuthTokenExpiration returns a boolean if a field has been set.

### GetOptions

`func (o *PortalSettings) GetOptions() PortalOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PortalSettings) GetOptionsOk() (*PortalOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PortalSettings) SetOptions(v PortalOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PortalSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


