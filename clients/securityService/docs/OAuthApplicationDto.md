# OAuthApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ApplicationType** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ConcurrencyToken** | Pointer to **NullableString** |  | [optional] 
**ConsentType** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**DisplayNames** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**RedirectUris** | Pointer to **NullableString** |  | [optional] 
**Requirements** | Pointer to **NullableString** |  | [optional] 
**Settings** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**BusinessApplicationID** | Pointer to **NullableString** |  | [optional] 
**AuthorizationsCount** | Pointer to **int32** |  | [optional] 
**TokensCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOAuthApplicationDto

`func NewOAuthApplicationDto() *OAuthApplicationDto`

NewOAuthApplicationDto instantiates a new OAuthApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthApplicationDtoWithDefaults

`func NewOAuthApplicationDtoWithDefaults() *OAuthApplicationDto`

NewOAuthApplicationDtoWithDefaults instantiates a new OAuthApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthApplicationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthApplicationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthApplicationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuthApplicationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OAuthApplicationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OAuthApplicationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetApplicationType

`func (o *OAuthApplicationDto) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *OAuthApplicationDto) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *OAuthApplicationDto) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *OAuthApplicationDto) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *OAuthApplicationDto) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *OAuthApplicationDto) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetClientId

`func (o *OAuthApplicationDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthApplicationDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthApplicationDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthApplicationDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *OAuthApplicationDto) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *OAuthApplicationDto) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetConcurrencyToken

`func (o *OAuthApplicationDto) GetConcurrencyToken() string`

GetConcurrencyToken returns the ConcurrencyToken field if non-nil, zero value otherwise.

### GetConcurrencyTokenOk

`func (o *OAuthApplicationDto) GetConcurrencyTokenOk() (*string, bool)`

GetConcurrencyTokenOk returns a tuple with the ConcurrencyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyToken

`func (o *OAuthApplicationDto) SetConcurrencyToken(v string)`

SetConcurrencyToken sets ConcurrencyToken field to given value.

### HasConcurrencyToken

`func (o *OAuthApplicationDto) HasConcurrencyToken() bool`

HasConcurrencyToken returns a boolean if a field has been set.

### SetConcurrencyTokenNil

`func (o *OAuthApplicationDto) SetConcurrencyTokenNil(b bool)`

 SetConcurrencyTokenNil sets the value for ConcurrencyToken to be an explicit nil

### UnsetConcurrencyToken
`func (o *OAuthApplicationDto) UnsetConcurrencyToken()`

UnsetConcurrencyToken ensures that no value is present for ConcurrencyToken, not even an explicit nil
### GetConsentType

`func (o *OAuthApplicationDto) GetConsentType() string`

GetConsentType returns the ConsentType field if non-nil, zero value otherwise.

### GetConsentTypeOk

`func (o *OAuthApplicationDto) GetConsentTypeOk() (*string, bool)`

GetConsentTypeOk returns a tuple with the ConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentType

`func (o *OAuthApplicationDto) SetConsentType(v string)`

SetConsentType sets ConsentType field to given value.

### HasConsentType

`func (o *OAuthApplicationDto) HasConsentType() bool`

HasConsentType returns a boolean if a field has been set.

### SetConsentTypeNil

`func (o *OAuthApplicationDto) SetConsentTypeNil(b bool)`

 SetConsentTypeNil sets the value for ConsentType to be an explicit nil

### UnsetConsentType
`func (o *OAuthApplicationDto) UnsetConsentType()`

UnsetConsentType ensures that no value is present for ConsentType, not even an explicit nil
### GetDisplayName

`func (o *OAuthApplicationDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuthApplicationDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuthApplicationDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuthApplicationDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OAuthApplicationDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OAuthApplicationDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDisplayNames

`func (o *OAuthApplicationDto) GetDisplayNames() string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *OAuthApplicationDto) GetDisplayNamesOk() (*string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *OAuthApplicationDto) SetDisplayNames(v string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *OAuthApplicationDto) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *OAuthApplicationDto) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *OAuthApplicationDto) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetPermissions

`func (o *OAuthApplicationDto) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OAuthApplicationDto) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OAuthApplicationDto) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OAuthApplicationDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *OAuthApplicationDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *OAuthApplicationDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPostLogoutRedirectUris

`func (o *OAuthApplicationDto) GetPostLogoutRedirectUris() string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *OAuthApplicationDto) GetPostLogoutRedirectUrisOk() (*string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *OAuthApplicationDto) SetPostLogoutRedirectUris(v string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *OAuthApplicationDto) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### SetPostLogoutRedirectUrisNil

`func (o *OAuthApplicationDto) SetPostLogoutRedirectUrisNil(b bool)`

 SetPostLogoutRedirectUrisNil sets the value for PostLogoutRedirectUris to be an explicit nil

### UnsetPostLogoutRedirectUris
`func (o *OAuthApplicationDto) UnsetPostLogoutRedirectUris()`

UnsetPostLogoutRedirectUris ensures that no value is present for PostLogoutRedirectUris, not even an explicit nil
### GetProperties

`func (o *OAuthApplicationDto) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *OAuthApplicationDto) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *OAuthApplicationDto) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *OAuthApplicationDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *OAuthApplicationDto) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *OAuthApplicationDto) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRedirectUris

`func (o *OAuthApplicationDto) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuthApplicationDto) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuthApplicationDto) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuthApplicationDto) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### SetRedirectUrisNil

`func (o *OAuthApplicationDto) SetRedirectUrisNil(b bool)`

 SetRedirectUrisNil sets the value for RedirectUris to be an explicit nil

### UnsetRedirectUris
`func (o *OAuthApplicationDto) UnsetRedirectUris()`

UnsetRedirectUris ensures that no value is present for RedirectUris, not even an explicit nil
### GetRequirements

`func (o *OAuthApplicationDto) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *OAuthApplicationDto) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *OAuthApplicationDto) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *OAuthApplicationDto) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *OAuthApplicationDto) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *OAuthApplicationDto) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetSettings

`func (o *OAuthApplicationDto) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OAuthApplicationDto) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OAuthApplicationDto) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *OAuthApplicationDto) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *OAuthApplicationDto) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *OAuthApplicationDto) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetType

`func (o *OAuthApplicationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthApplicationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthApplicationDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuthApplicationDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OAuthApplicationDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OAuthApplicationDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLogo

`func (o *OAuthApplicationDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *OAuthApplicationDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *OAuthApplicationDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *OAuthApplicationDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *OAuthApplicationDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *OAuthApplicationDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetBusinessID

`func (o *OAuthApplicationDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *OAuthApplicationDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *OAuthApplicationDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *OAuthApplicationDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *OAuthApplicationDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *OAuthApplicationDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *OAuthApplicationDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *OAuthApplicationDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *OAuthApplicationDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *OAuthApplicationDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *OAuthApplicationDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *OAuthApplicationDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetBusinessApplicationID

`func (o *OAuthApplicationDto) GetBusinessApplicationID() string`

GetBusinessApplicationID returns the BusinessApplicationID field if non-nil, zero value otherwise.

### GetBusinessApplicationIDOk

`func (o *OAuthApplicationDto) GetBusinessApplicationIDOk() (*string, bool)`

GetBusinessApplicationIDOk returns a tuple with the BusinessApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationID

`func (o *OAuthApplicationDto) SetBusinessApplicationID(v string)`

SetBusinessApplicationID sets BusinessApplicationID field to given value.

### HasBusinessApplicationID

`func (o *OAuthApplicationDto) HasBusinessApplicationID() bool`

HasBusinessApplicationID returns a boolean if a field has been set.

### SetBusinessApplicationIDNil

`func (o *OAuthApplicationDto) SetBusinessApplicationIDNil(b bool)`

 SetBusinessApplicationIDNil sets the value for BusinessApplicationID to be an explicit nil

### UnsetBusinessApplicationID
`func (o *OAuthApplicationDto) UnsetBusinessApplicationID()`

UnsetBusinessApplicationID ensures that no value is present for BusinessApplicationID, not even an explicit nil
### GetAuthorizationsCount

`func (o *OAuthApplicationDto) GetAuthorizationsCount() int32`

GetAuthorizationsCount returns the AuthorizationsCount field if non-nil, zero value otherwise.

### GetAuthorizationsCountOk

`func (o *OAuthApplicationDto) GetAuthorizationsCountOk() (*int32, bool)`

GetAuthorizationsCountOk returns a tuple with the AuthorizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationsCount

`func (o *OAuthApplicationDto) SetAuthorizationsCount(v int32)`

SetAuthorizationsCount sets AuthorizationsCount field to given value.

### HasAuthorizationsCount

`func (o *OAuthApplicationDto) HasAuthorizationsCount() bool`

HasAuthorizationsCount returns a boolean if a field has been set.

### GetTokensCount

`func (o *OAuthApplicationDto) GetTokensCount() int32`

GetTokensCount returns the TokensCount field if non-nil, zero value otherwise.

### GetTokensCountOk

`func (o *OAuthApplicationDto) GetTokensCountOk() (*int32, bool)`

GetTokensCountOk returns a tuple with the TokensCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensCount

`func (o *OAuthApplicationDto) SetTokensCount(v int32)`

SetTokensCount sets TokensCount field to given value.

### HasTokensCount

`func (o *OAuthApplicationDto) HasTokensCount() bool`

HasTokensCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


