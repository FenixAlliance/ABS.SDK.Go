# OAuthApplicationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**ConsentType** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**Requirements** | Pointer to **NullableString** |  | [optional] 
**RedirectUris** | Pointer to **NullableString** |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuthApplicationUpdateDto

`func NewOAuthApplicationUpdateDto() *OAuthApplicationUpdateDto`

NewOAuthApplicationUpdateDto instantiates a new OAuthApplicationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthApplicationUpdateDtoWithDefaults

`func NewOAuthApplicationUpdateDtoWithDefaults() *OAuthApplicationUpdateDto`

NewOAuthApplicationUpdateDtoWithDefaults instantiates a new OAuthApplicationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *OAuthApplicationUpdateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuthApplicationUpdateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuthApplicationUpdateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuthApplicationUpdateDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OAuthApplicationUpdateDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OAuthApplicationUpdateDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetClientSecret

`func (o *OAuthApplicationUpdateDto) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthApplicationUpdateDto) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthApplicationUpdateDto) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthApplicationUpdateDto) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuthApplicationUpdateDto) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuthApplicationUpdateDto) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetConsentType

`func (o *OAuthApplicationUpdateDto) GetConsentType() string`

GetConsentType returns the ConsentType field if non-nil, zero value otherwise.

### GetConsentTypeOk

`func (o *OAuthApplicationUpdateDto) GetConsentTypeOk() (*string, bool)`

GetConsentTypeOk returns a tuple with the ConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentType

`func (o *OAuthApplicationUpdateDto) SetConsentType(v string)`

SetConsentType sets ConsentType field to given value.

### HasConsentType

`func (o *OAuthApplicationUpdateDto) HasConsentType() bool`

HasConsentType returns a boolean if a field has been set.

### SetConsentTypeNil

`func (o *OAuthApplicationUpdateDto) SetConsentTypeNil(b bool)`

 SetConsentTypeNil sets the value for ConsentType to be an explicit nil

### UnsetConsentType
`func (o *OAuthApplicationUpdateDto) UnsetConsentType()`

UnsetConsentType ensures that no value is present for ConsentType, not even an explicit nil
### GetPermissions

`func (o *OAuthApplicationUpdateDto) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OAuthApplicationUpdateDto) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OAuthApplicationUpdateDto) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OAuthApplicationUpdateDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *OAuthApplicationUpdateDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *OAuthApplicationUpdateDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRequirements

`func (o *OAuthApplicationUpdateDto) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *OAuthApplicationUpdateDto) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *OAuthApplicationUpdateDto) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *OAuthApplicationUpdateDto) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *OAuthApplicationUpdateDto) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *OAuthApplicationUpdateDto) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetRedirectUris

`func (o *OAuthApplicationUpdateDto) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuthApplicationUpdateDto) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuthApplicationUpdateDto) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuthApplicationUpdateDto) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### SetRedirectUrisNil

`func (o *OAuthApplicationUpdateDto) SetRedirectUrisNil(b bool)`

 SetRedirectUrisNil sets the value for RedirectUris to be an explicit nil

### UnsetRedirectUris
`func (o *OAuthApplicationUpdateDto) UnsetRedirectUris()`

UnsetRedirectUris ensures that no value is present for RedirectUris, not even an explicit nil
### GetPostLogoutRedirectUris

`func (o *OAuthApplicationUpdateDto) GetPostLogoutRedirectUris() string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *OAuthApplicationUpdateDto) GetPostLogoutRedirectUrisOk() (*string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *OAuthApplicationUpdateDto) SetPostLogoutRedirectUris(v string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *OAuthApplicationUpdateDto) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### SetPostLogoutRedirectUrisNil

`func (o *OAuthApplicationUpdateDto) SetPostLogoutRedirectUrisNil(b bool)`

 SetPostLogoutRedirectUrisNil sets the value for PostLogoutRedirectUris to be an explicit nil

### UnsetPostLogoutRedirectUris
`func (o *OAuthApplicationUpdateDto) UnsetPostLogoutRedirectUris()`

UnsetPostLogoutRedirectUris ensures that no value is present for PostLogoutRedirectUris, not even an explicit nil
### GetLogo

`func (o *OAuthApplicationUpdateDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *OAuthApplicationUpdateDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *OAuthApplicationUpdateDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *OAuthApplicationUpdateDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *OAuthApplicationUpdateDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *OAuthApplicationUpdateDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


