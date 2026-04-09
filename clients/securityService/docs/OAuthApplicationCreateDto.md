# OAuthApplicationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**ConsentType** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**Requirements** | Pointer to **NullableString** |  | [optional] 
**RedirectUris** | Pointer to **NullableString** |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuthApplicationCreateDto

`func NewOAuthApplicationCreateDto(displayName string, ) *OAuthApplicationCreateDto`

NewOAuthApplicationCreateDto instantiates a new OAuthApplicationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthApplicationCreateDtoWithDefaults

`func NewOAuthApplicationCreateDtoWithDefaults() *OAuthApplicationCreateDto`

NewOAuthApplicationCreateDtoWithDefaults instantiates a new OAuthApplicationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *OAuthApplicationCreateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuthApplicationCreateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuthApplicationCreateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetClientId

`func (o *OAuthApplicationCreateDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthApplicationCreateDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthApplicationCreateDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthApplicationCreateDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *OAuthApplicationCreateDto) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *OAuthApplicationCreateDto) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *OAuthApplicationCreateDto) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthApplicationCreateDto) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthApplicationCreateDto) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthApplicationCreateDto) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuthApplicationCreateDto) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuthApplicationCreateDto) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetConsentType

`func (o *OAuthApplicationCreateDto) GetConsentType() string`

GetConsentType returns the ConsentType field if non-nil, zero value otherwise.

### GetConsentTypeOk

`func (o *OAuthApplicationCreateDto) GetConsentTypeOk() (*string, bool)`

GetConsentTypeOk returns a tuple with the ConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentType

`func (o *OAuthApplicationCreateDto) SetConsentType(v string)`

SetConsentType sets ConsentType field to given value.

### HasConsentType

`func (o *OAuthApplicationCreateDto) HasConsentType() bool`

HasConsentType returns a boolean if a field has been set.

### SetConsentTypeNil

`func (o *OAuthApplicationCreateDto) SetConsentTypeNil(b bool)`

 SetConsentTypeNil sets the value for ConsentType to be an explicit nil

### UnsetConsentType
`func (o *OAuthApplicationCreateDto) UnsetConsentType()`

UnsetConsentType ensures that no value is present for ConsentType, not even an explicit nil
### GetPermissions

`func (o *OAuthApplicationCreateDto) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OAuthApplicationCreateDto) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OAuthApplicationCreateDto) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OAuthApplicationCreateDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *OAuthApplicationCreateDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *OAuthApplicationCreateDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRequirements

`func (o *OAuthApplicationCreateDto) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *OAuthApplicationCreateDto) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *OAuthApplicationCreateDto) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *OAuthApplicationCreateDto) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *OAuthApplicationCreateDto) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *OAuthApplicationCreateDto) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetRedirectUris

`func (o *OAuthApplicationCreateDto) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuthApplicationCreateDto) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuthApplicationCreateDto) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuthApplicationCreateDto) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### SetRedirectUrisNil

`func (o *OAuthApplicationCreateDto) SetRedirectUrisNil(b bool)`

 SetRedirectUrisNil sets the value for RedirectUris to be an explicit nil

### UnsetRedirectUris
`func (o *OAuthApplicationCreateDto) UnsetRedirectUris()`

UnsetRedirectUris ensures that no value is present for RedirectUris, not even an explicit nil
### GetPostLogoutRedirectUris

`func (o *OAuthApplicationCreateDto) GetPostLogoutRedirectUris() string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *OAuthApplicationCreateDto) GetPostLogoutRedirectUrisOk() (*string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *OAuthApplicationCreateDto) SetPostLogoutRedirectUris(v string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *OAuthApplicationCreateDto) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### SetPostLogoutRedirectUrisNil

`func (o *OAuthApplicationCreateDto) SetPostLogoutRedirectUrisNil(b bool)`

 SetPostLogoutRedirectUrisNil sets the value for PostLogoutRedirectUris to be an explicit nil

### UnsetPostLogoutRedirectUris
`func (o *OAuthApplicationCreateDto) UnsetPostLogoutRedirectUris()`

UnsetPostLogoutRedirectUris ensures that no value is present for PostLogoutRedirectUris, not even an explicit nil
### GetLogo

`func (o *OAuthApplicationCreateDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *OAuthApplicationCreateDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *OAuthApplicationCreateDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *OAuthApplicationCreateDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *OAuthApplicationCreateDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *OAuthApplicationCreateDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetBusinessID

`func (o *OAuthApplicationCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *OAuthApplicationCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *OAuthApplicationCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *OAuthApplicationCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *OAuthApplicationCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *OAuthApplicationCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *OAuthApplicationCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *OAuthApplicationCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *OAuthApplicationCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *OAuthApplicationCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *OAuthApplicationCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *OAuthApplicationCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


