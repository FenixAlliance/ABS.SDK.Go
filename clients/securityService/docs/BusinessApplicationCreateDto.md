# BusinessApplicationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**AvatarURL** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**IsMultiTenant** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsSinglePageApplication** | Pointer to **bool** |  | [optional] 
**IsNativeOrDesktopApp** | Pointer to **bool** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicyURL** | Pointer to **NullableString** |  | [optional] 
**TermsAndConditionsURL** | Pointer to **NullableString** |  | [optional] 
**RequireHttps** | Pointer to **bool** |  | [optional] 
**RequireAppSecret** | Pointer to **bool** |  | [optional] 
**EnableClientOauthLogin** | Pointer to **bool** |  | [optional] 
**EnableWebOAuthLogin** | Pointer to **bool** |  | [optional] 
**EnableDeviceOAuthLogin** | Pointer to **bool** |  | [optional] 
**AllowAccessToSuiteSettings** | Pointer to **bool** |  | [optional] 
**RequireWebOAuthReauthentication** | Pointer to **bool** |  | [optional] 
**RequireTwoFactorReauthorization** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedBrowserOAuthLogin** | Pointer to **bool** |  | [optional] 
**UseStrictModeForRedirectURIs** | Pointer to **bool** |  | [optional] 
**CountryRestricted** | Pointer to **bool** |  | [optional] 
**SpaUIEngine** | Pointer to **NullableString** |  | [optional] 
**SpaStaticFilesRootPath** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeAppPath** | Pointer to **NullableString** |  | [optional] 
**SpaNpmStartScript** | Pointer to **NullableString** |  | [optional] 
**SpaNpmPublishScript** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeSourcePath** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeOutputPath** | Pointer to **NullableString** |  | [optional] 
**UseProxyToSpaDevelopmentServer** | Pointer to **bool** |  | [optional] 
**SpaDevelopmentServerUri** | Pointer to **NullableString** |  | [optional] 
**EnableGitRepoManagement** | Pointer to **bool** |  | [optional] 
**GitRepoUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBusinessApplicationCreateDto

`func NewBusinessApplicationCreateDto(name string, ) *BusinessApplicationCreateDto`

NewBusinessApplicationCreateDto instantiates a new BusinessApplicationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessApplicationCreateDtoWithDefaults

`func NewBusinessApplicationCreateDtoWithDefaults() *BusinessApplicationCreateDto`

NewBusinessApplicationCreateDtoWithDefaults instantiates a new BusinessApplicationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessApplicationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessApplicationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessApplicationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessApplicationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BusinessApplicationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessApplicationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessApplicationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessApplicationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *BusinessApplicationCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessApplicationCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessApplicationCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *BusinessApplicationCreateDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BusinessApplicationCreateDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BusinessApplicationCreateDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BusinessApplicationCreateDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *BusinessApplicationCreateDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *BusinessApplicationCreateDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetDisplayName

`func (o *BusinessApplicationCreateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BusinessApplicationCreateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BusinessApplicationCreateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BusinessApplicationCreateDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BusinessApplicationCreateDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BusinessApplicationCreateDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatarURL

`func (o *BusinessApplicationCreateDto) GetAvatarURL() string`

GetAvatarURL returns the AvatarURL field if non-nil, zero value otherwise.

### GetAvatarURLOk

`func (o *BusinessApplicationCreateDto) GetAvatarURLOk() (*string, bool)`

GetAvatarURLOk returns a tuple with the AvatarURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarURL

`func (o *BusinessApplicationCreateDto) SetAvatarURL(v string)`

SetAvatarURL sets AvatarURL field to given value.

### HasAvatarURL

`func (o *BusinessApplicationCreateDto) HasAvatarURL() bool`

HasAvatarURL returns a boolean if a field has been set.

### SetAvatarURLNil

`func (o *BusinessApplicationCreateDto) SetAvatarURLNil(b bool)`

 SetAvatarURLNil sets the value for AvatarURL to be an explicit nil

### UnsetAvatarURL
`func (o *BusinessApplicationCreateDto) UnsetAvatarURL()`

UnsetAvatarURL ensures that no value is present for AvatarURL, not even an explicit nil
### GetWebsiteUrl

`func (o *BusinessApplicationCreateDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *BusinessApplicationCreateDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *BusinessApplicationCreateDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *BusinessApplicationCreateDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *BusinessApplicationCreateDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *BusinessApplicationCreateDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetIsMultiTenant

`func (o *BusinessApplicationCreateDto) GetIsMultiTenant() bool`

GetIsMultiTenant returns the IsMultiTenant field if non-nil, zero value otherwise.

### GetIsMultiTenantOk

`func (o *BusinessApplicationCreateDto) GetIsMultiTenantOk() (*bool, bool)`

GetIsMultiTenantOk returns a tuple with the IsMultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiTenant

`func (o *BusinessApplicationCreateDto) SetIsMultiTenant(v bool)`

SetIsMultiTenant sets IsMultiTenant field to given value.

### HasIsMultiTenant

`func (o *BusinessApplicationCreateDto) HasIsMultiTenant() bool`

HasIsMultiTenant returns a boolean if a field has been set.

### GetIsVerified

`func (o *BusinessApplicationCreateDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *BusinessApplicationCreateDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *BusinessApplicationCreateDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *BusinessApplicationCreateDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsDisabled

`func (o *BusinessApplicationCreateDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *BusinessApplicationCreateDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *BusinessApplicationCreateDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *BusinessApplicationCreateDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsSinglePageApplication

`func (o *BusinessApplicationCreateDto) GetIsSinglePageApplication() bool`

GetIsSinglePageApplication returns the IsSinglePageApplication field if non-nil, zero value otherwise.

### GetIsSinglePageApplicationOk

`func (o *BusinessApplicationCreateDto) GetIsSinglePageApplicationOk() (*bool, bool)`

GetIsSinglePageApplicationOk returns a tuple with the IsSinglePageApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSinglePageApplication

`func (o *BusinessApplicationCreateDto) SetIsSinglePageApplication(v bool)`

SetIsSinglePageApplication sets IsSinglePageApplication field to given value.

### HasIsSinglePageApplication

`func (o *BusinessApplicationCreateDto) HasIsSinglePageApplication() bool`

HasIsSinglePageApplication returns a boolean if a field has been set.

### GetIsNativeOrDesktopApp

`func (o *BusinessApplicationCreateDto) GetIsNativeOrDesktopApp() bool`

GetIsNativeOrDesktopApp returns the IsNativeOrDesktopApp field if non-nil, zero value otherwise.

### GetIsNativeOrDesktopAppOk

`func (o *BusinessApplicationCreateDto) GetIsNativeOrDesktopAppOk() (*bool, bool)`

GetIsNativeOrDesktopAppOk returns a tuple with the IsNativeOrDesktopApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeOrDesktopApp

`func (o *BusinessApplicationCreateDto) SetIsNativeOrDesktopApp(v bool)`

SetIsNativeOrDesktopApp sets IsNativeOrDesktopApp field to given value.

### HasIsNativeOrDesktopApp

`func (o *BusinessApplicationCreateDto) HasIsNativeOrDesktopApp() bool`

HasIsNativeOrDesktopApp returns a boolean if a field has been set.

### GetContactEmail

`func (o *BusinessApplicationCreateDto) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *BusinessApplicationCreateDto) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *BusinessApplicationCreateDto) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *BusinessApplicationCreateDto) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *BusinessApplicationCreateDto) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *BusinessApplicationCreateDto) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetPrivacyPolicyURL

`func (o *BusinessApplicationCreateDto) GetPrivacyPolicyURL() string`

GetPrivacyPolicyURL returns the PrivacyPolicyURL field if non-nil, zero value otherwise.

### GetPrivacyPolicyURLOk

`func (o *BusinessApplicationCreateDto) GetPrivacyPolicyURLOk() (*string, bool)`

GetPrivacyPolicyURLOk returns a tuple with the PrivacyPolicyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyURL

`func (o *BusinessApplicationCreateDto) SetPrivacyPolicyURL(v string)`

SetPrivacyPolicyURL sets PrivacyPolicyURL field to given value.

### HasPrivacyPolicyURL

`func (o *BusinessApplicationCreateDto) HasPrivacyPolicyURL() bool`

HasPrivacyPolicyURL returns a boolean if a field has been set.

### SetPrivacyPolicyURLNil

`func (o *BusinessApplicationCreateDto) SetPrivacyPolicyURLNil(b bool)`

 SetPrivacyPolicyURLNil sets the value for PrivacyPolicyURL to be an explicit nil

### UnsetPrivacyPolicyURL
`func (o *BusinessApplicationCreateDto) UnsetPrivacyPolicyURL()`

UnsetPrivacyPolicyURL ensures that no value is present for PrivacyPolicyURL, not even an explicit nil
### GetTermsAndConditionsURL

`func (o *BusinessApplicationCreateDto) GetTermsAndConditionsURL() string`

GetTermsAndConditionsURL returns the TermsAndConditionsURL field if non-nil, zero value otherwise.

### GetTermsAndConditionsURLOk

`func (o *BusinessApplicationCreateDto) GetTermsAndConditionsURLOk() (*string, bool)`

GetTermsAndConditionsURLOk returns a tuple with the TermsAndConditionsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsURL

`func (o *BusinessApplicationCreateDto) SetTermsAndConditionsURL(v string)`

SetTermsAndConditionsURL sets TermsAndConditionsURL field to given value.

### HasTermsAndConditionsURL

`func (o *BusinessApplicationCreateDto) HasTermsAndConditionsURL() bool`

HasTermsAndConditionsURL returns a boolean if a field has been set.

### SetTermsAndConditionsURLNil

`func (o *BusinessApplicationCreateDto) SetTermsAndConditionsURLNil(b bool)`

 SetTermsAndConditionsURLNil sets the value for TermsAndConditionsURL to be an explicit nil

### UnsetTermsAndConditionsURL
`func (o *BusinessApplicationCreateDto) UnsetTermsAndConditionsURL()`

UnsetTermsAndConditionsURL ensures that no value is present for TermsAndConditionsURL, not even an explicit nil
### GetRequireHttps

`func (o *BusinessApplicationCreateDto) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *BusinessApplicationCreateDto) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *BusinessApplicationCreateDto) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *BusinessApplicationCreateDto) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetRequireAppSecret

`func (o *BusinessApplicationCreateDto) GetRequireAppSecret() bool`

GetRequireAppSecret returns the RequireAppSecret field if non-nil, zero value otherwise.

### GetRequireAppSecretOk

`func (o *BusinessApplicationCreateDto) GetRequireAppSecretOk() (*bool, bool)`

GetRequireAppSecretOk returns a tuple with the RequireAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAppSecret

`func (o *BusinessApplicationCreateDto) SetRequireAppSecret(v bool)`

SetRequireAppSecret sets RequireAppSecret field to given value.

### HasRequireAppSecret

`func (o *BusinessApplicationCreateDto) HasRequireAppSecret() bool`

HasRequireAppSecret returns a boolean if a field has been set.

### GetEnableClientOauthLogin

`func (o *BusinessApplicationCreateDto) GetEnableClientOauthLogin() bool`

GetEnableClientOauthLogin returns the EnableClientOauthLogin field if non-nil, zero value otherwise.

### GetEnableClientOauthLoginOk

`func (o *BusinessApplicationCreateDto) GetEnableClientOauthLoginOk() (*bool, bool)`

GetEnableClientOauthLoginOk returns a tuple with the EnableClientOauthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientOauthLogin

`func (o *BusinessApplicationCreateDto) SetEnableClientOauthLogin(v bool)`

SetEnableClientOauthLogin sets EnableClientOauthLogin field to given value.

### HasEnableClientOauthLogin

`func (o *BusinessApplicationCreateDto) HasEnableClientOauthLogin() bool`

HasEnableClientOauthLogin returns a boolean if a field has been set.

### GetEnableWebOAuthLogin

`func (o *BusinessApplicationCreateDto) GetEnableWebOAuthLogin() bool`

GetEnableWebOAuthLogin returns the EnableWebOAuthLogin field if non-nil, zero value otherwise.

### GetEnableWebOAuthLoginOk

`func (o *BusinessApplicationCreateDto) GetEnableWebOAuthLoginOk() (*bool, bool)`

GetEnableWebOAuthLoginOk returns a tuple with the EnableWebOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebOAuthLogin

`func (o *BusinessApplicationCreateDto) SetEnableWebOAuthLogin(v bool)`

SetEnableWebOAuthLogin sets EnableWebOAuthLogin field to given value.

### HasEnableWebOAuthLogin

`func (o *BusinessApplicationCreateDto) HasEnableWebOAuthLogin() bool`

HasEnableWebOAuthLogin returns a boolean if a field has been set.

### GetEnableDeviceOAuthLogin

`func (o *BusinessApplicationCreateDto) GetEnableDeviceOAuthLogin() bool`

GetEnableDeviceOAuthLogin returns the EnableDeviceOAuthLogin field if non-nil, zero value otherwise.

### GetEnableDeviceOAuthLoginOk

`func (o *BusinessApplicationCreateDto) GetEnableDeviceOAuthLoginOk() (*bool, bool)`

GetEnableDeviceOAuthLoginOk returns a tuple with the EnableDeviceOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceOAuthLogin

`func (o *BusinessApplicationCreateDto) SetEnableDeviceOAuthLogin(v bool)`

SetEnableDeviceOAuthLogin sets EnableDeviceOAuthLogin field to given value.

### HasEnableDeviceOAuthLogin

`func (o *BusinessApplicationCreateDto) HasEnableDeviceOAuthLogin() bool`

HasEnableDeviceOAuthLogin returns a boolean if a field has been set.

### GetAllowAccessToSuiteSettings

`func (o *BusinessApplicationCreateDto) GetAllowAccessToSuiteSettings() bool`

GetAllowAccessToSuiteSettings returns the AllowAccessToSuiteSettings field if non-nil, zero value otherwise.

### GetAllowAccessToSuiteSettingsOk

`func (o *BusinessApplicationCreateDto) GetAllowAccessToSuiteSettingsOk() (*bool, bool)`

GetAllowAccessToSuiteSettingsOk returns a tuple with the AllowAccessToSuiteSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessToSuiteSettings

`func (o *BusinessApplicationCreateDto) SetAllowAccessToSuiteSettings(v bool)`

SetAllowAccessToSuiteSettings sets AllowAccessToSuiteSettings field to given value.

### HasAllowAccessToSuiteSettings

`func (o *BusinessApplicationCreateDto) HasAllowAccessToSuiteSettings() bool`

HasAllowAccessToSuiteSettings returns a boolean if a field has been set.

### GetRequireWebOAuthReauthentication

`func (o *BusinessApplicationCreateDto) GetRequireWebOAuthReauthentication() bool`

GetRequireWebOAuthReauthentication returns the RequireWebOAuthReauthentication field if non-nil, zero value otherwise.

### GetRequireWebOAuthReauthenticationOk

`func (o *BusinessApplicationCreateDto) GetRequireWebOAuthReauthenticationOk() (*bool, bool)`

GetRequireWebOAuthReauthenticationOk returns a tuple with the RequireWebOAuthReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireWebOAuthReauthentication

`func (o *BusinessApplicationCreateDto) SetRequireWebOAuthReauthentication(v bool)`

SetRequireWebOAuthReauthentication sets RequireWebOAuthReauthentication field to given value.

### HasRequireWebOAuthReauthentication

`func (o *BusinessApplicationCreateDto) HasRequireWebOAuthReauthentication() bool`

HasRequireWebOAuthReauthentication returns a boolean if a field has been set.

### GetRequireTwoFactorReauthorization

`func (o *BusinessApplicationCreateDto) GetRequireTwoFactorReauthorization() bool`

GetRequireTwoFactorReauthorization returns the RequireTwoFactorReauthorization field if non-nil, zero value otherwise.

### GetRequireTwoFactorReauthorizationOk

`func (o *BusinessApplicationCreateDto) GetRequireTwoFactorReauthorizationOk() (*bool, bool)`

GetRequireTwoFactorReauthorizationOk returns a tuple with the RequireTwoFactorReauthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTwoFactorReauthorization

`func (o *BusinessApplicationCreateDto) SetRequireTwoFactorReauthorization(v bool)`

SetRequireTwoFactorReauthorization sets RequireTwoFactorReauthorization field to given value.

### HasRequireTwoFactorReauthorization

`func (o *BusinessApplicationCreateDto) HasRequireTwoFactorReauthorization() bool`

HasRequireTwoFactorReauthorization returns a boolean if a field has been set.

### GetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationCreateDto) GetEnableEmbeddedBrowserOAuthLogin() bool`

GetEnableEmbeddedBrowserOAuthLogin returns the EnableEmbeddedBrowserOAuthLogin field if non-nil, zero value otherwise.

### GetEnableEmbeddedBrowserOAuthLoginOk

`func (o *BusinessApplicationCreateDto) GetEnableEmbeddedBrowserOAuthLoginOk() (*bool, bool)`

GetEnableEmbeddedBrowserOAuthLoginOk returns a tuple with the EnableEmbeddedBrowserOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationCreateDto) SetEnableEmbeddedBrowserOAuthLogin(v bool)`

SetEnableEmbeddedBrowserOAuthLogin sets EnableEmbeddedBrowserOAuthLogin field to given value.

### HasEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationCreateDto) HasEnableEmbeddedBrowserOAuthLogin() bool`

HasEnableEmbeddedBrowserOAuthLogin returns a boolean if a field has been set.

### GetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationCreateDto) GetUseStrictModeForRedirectURIs() bool`

GetUseStrictModeForRedirectURIs returns the UseStrictModeForRedirectURIs field if non-nil, zero value otherwise.

### GetUseStrictModeForRedirectURIsOk

`func (o *BusinessApplicationCreateDto) GetUseStrictModeForRedirectURIsOk() (*bool, bool)`

GetUseStrictModeForRedirectURIsOk returns a tuple with the UseStrictModeForRedirectURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationCreateDto) SetUseStrictModeForRedirectURIs(v bool)`

SetUseStrictModeForRedirectURIs sets UseStrictModeForRedirectURIs field to given value.

### HasUseStrictModeForRedirectURIs

`func (o *BusinessApplicationCreateDto) HasUseStrictModeForRedirectURIs() bool`

HasUseStrictModeForRedirectURIs returns a boolean if a field has been set.

### GetCountryRestricted

`func (o *BusinessApplicationCreateDto) GetCountryRestricted() bool`

GetCountryRestricted returns the CountryRestricted field if non-nil, zero value otherwise.

### GetCountryRestrictedOk

`func (o *BusinessApplicationCreateDto) GetCountryRestrictedOk() (*bool, bool)`

GetCountryRestrictedOk returns a tuple with the CountryRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRestricted

`func (o *BusinessApplicationCreateDto) SetCountryRestricted(v bool)`

SetCountryRestricted sets CountryRestricted field to given value.

### HasCountryRestricted

`func (o *BusinessApplicationCreateDto) HasCountryRestricted() bool`

HasCountryRestricted returns a boolean if a field has been set.

### GetSpaUIEngine

`func (o *BusinessApplicationCreateDto) GetSpaUIEngine() string`

GetSpaUIEngine returns the SpaUIEngine field if non-nil, zero value otherwise.

### GetSpaUIEngineOk

`func (o *BusinessApplicationCreateDto) GetSpaUIEngineOk() (*string, bool)`

GetSpaUIEngineOk returns a tuple with the SpaUIEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUIEngine

`func (o *BusinessApplicationCreateDto) SetSpaUIEngine(v string)`

SetSpaUIEngine sets SpaUIEngine field to given value.

### HasSpaUIEngine

`func (o *BusinessApplicationCreateDto) HasSpaUIEngine() bool`

HasSpaUIEngine returns a boolean if a field has been set.

### SetSpaUIEngineNil

`func (o *BusinessApplicationCreateDto) SetSpaUIEngineNil(b bool)`

 SetSpaUIEngineNil sets the value for SpaUIEngine to be an explicit nil

### UnsetSpaUIEngine
`func (o *BusinessApplicationCreateDto) UnsetSpaUIEngine()`

UnsetSpaUIEngine ensures that no value is present for SpaUIEngine, not even an explicit nil
### GetSpaStaticFilesRootPath

`func (o *BusinessApplicationCreateDto) GetSpaStaticFilesRootPath() string`

GetSpaStaticFilesRootPath returns the SpaStaticFilesRootPath field if non-nil, zero value otherwise.

### GetSpaStaticFilesRootPathOk

`func (o *BusinessApplicationCreateDto) GetSpaStaticFilesRootPathOk() (*string, bool)`

GetSpaStaticFilesRootPathOk returns a tuple with the SpaStaticFilesRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaStaticFilesRootPath

`func (o *BusinessApplicationCreateDto) SetSpaStaticFilesRootPath(v string)`

SetSpaStaticFilesRootPath sets SpaStaticFilesRootPath field to given value.

### HasSpaStaticFilesRootPath

`func (o *BusinessApplicationCreateDto) HasSpaStaticFilesRootPath() bool`

HasSpaStaticFilesRootPath returns a boolean if a field has been set.

### SetSpaStaticFilesRootPathNil

`func (o *BusinessApplicationCreateDto) SetSpaStaticFilesRootPathNil(b bool)`

 SetSpaStaticFilesRootPathNil sets the value for SpaStaticFilesRootPath to be an explicit nil

### UnsetSpaStaticFilesRootPath
`func (o *BusinessApplicationCreateDto) UnsetSpaStaticFilesRootPath()`

UnsetSpaStaticFilesRootPath ensures that no value is present for SpaStaticFilesRootPath, not even an explicit nil
### GetSpaRelativeAppPath

`func (o *BusinessApplicationCreateDto) GetSpaRelativeAppPath() string`

GetSpaRelativeAppPath returns the SpaRelativeAppPath field if non-nil, zero value otherwise.

### GetSpaRelativeAppPathOk

`func (o *BusinessApplicationCreateDto) GetSpaRelativeAppPathOk() (*string, bool)`

GetSpaRelativeAppPathOk returns a tuple with the SpaRelativeAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeAppPath

`func (o *BusinessApplicationCreateDto) SetSpaRelativeAppPath(v string)`

SetSpaRelativeAppPath sets SpaRelativeAppPath field to given value.

### HasSpaRelativeAppPath

`func (o *BusinessApplicationCreateDto) HasSpaRelativeAppPath() bool`

HasSpaRelativeAppPath returns a boolean if a field has been set.

### SetSpaRelativeAppPathNil

`func (o *BusinessApplicationCreateDto) SetSpaRelativeAppPathNil(b bool)`

 SetSpaRelativeAppPathNil sets the value for SpaRelativeAppPath to be an explicit nil

### UnsetSpaRelativeAppPath
`func (o *BusinessApplicationCreateDto) UnsetSpaRelativeAppPath()`

UnsetSpaRelativeAppPath ensures that no value is present for SpaRelativeAppPath, not even an explicit nil
### GetSpaNpmStartScript

`func (o *BusinessApplicationCreateDto) GetSpaNpmStartScript() string`

GetSpaNpmStartScript returns the SpaNpmStartScript field if non-nil, zero value otherwise.

### GetSpaNpmStartScriptOk

`func (o *BusinessApplicationCreateDto) GetSpaNpmStartScriptOk() (*string, bool)`

GetSpaNpmStartScriptOk returns a tuple with the SpaNpmStartScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmStartScript

`func (o *BusinessApplicationCreateDto) SetSpaNpmStartScript(v string)`

SetSpaNpmStartScript sets SpaNpmStartScript field to given value.

### HasSpaNpmStartScript

`func (o *BusinessApplicationCreateDto) HasSpaNpmStartScript() bool`

HasSpaNpmStartScript returns a boolean if a field has been set.

### SetSpaNpmStartScriptNil

`func (o *BusinessApplicationCreateDto) SetSpaNpmStartScriptNil(b bool)`

 SetSpaNpmStartScriptNil sets the value for SpaNpmStartScript to be an explicit nil

### UnsetSpaNpmStartScript
`func (o *BusinessApplicationCreateDto) UnsetSpaNpmStartScript()`

UnsetSpaNpmStartScript ensures that no value is present for SpaNpmStartScript, not even an explicit nil
### GetSpaNpmPublishScript

`func (o *BusinessApplicationCreateDto) GetSpaNpmPublishScript() string`

GetSpaNpmPublishScript returns the SpaNpmPublishScript field if non-nil, zero value otherwise.

### GetSpaNpmPublishScriptOk

`func (o *BusinessApplicationCreateDto) GetSpaNpmPublishScriptOk() (*string, bool)`

GetSpaNpmPublishScriptOk returns a tuple with the SpaNpmPublishScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmPublishScript

`func (o *BusinessApplicationCreateDto) SetSpaNpmPublishScript(v string)`

SetSpaNpmPublishScript sets SpaNpmPublishScript field to given value.

### HasSpaNpmPublishScript

`func (o *BusinessApplicationCreateDto) HasSpaNpmPublishScript() bool`

HasSpaNpmPublishScript returns a boolean if a field has been set.

### SetSpaNpmPublishScriptNil

`func (o *BusinessApplicationCreateDto) SetSpaNpmPublishScriptNil(b bool)`

 SetSpaNpmPublishScriptNil sets the value for SpaNpmPublishScript to be an explicit nil

### UnsetSpaNpmPublishScript
`func (o *BusinessApplicationCreateDto) UnsetSpaNpmPublishScript()`

UnsetSpaNpmPublishScript ensures that no value is present for SpaNpmPublishScript, not even an explicit nil
### GetSpaRelativeSourcePath

`func (o *BusinessApplicationCreateDto) GetSpaRelativeSourcePath() string`

GetSpaRelativeSourcePath returns the SpaRelativeSourcePath field if non-nil, zero value otherwise.

### GetSpaRelativeSourcePathOk

`func (o *BusinessApplicationCreateDto) GetSpaRelativeSourcePathOk() (*string, bool)`

GetSpaRelativeSourcePathOk returns a tuple with the SpaRelativeSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeSourcePath

`func (o *BusinessApplicationCreateDto) SetSpaRelativeSourcePath(v string)`

SetSpaRelativeSourcePath sets SpaRelativeSourcePath field to given value.

### HasSpaRelativeSourcePath

`func (o *BusinessApplicationCreateDto) HasSpaRelativeSourcePath() bool`

HasSpaRelativeSourcePath returns a boolean if a field has been set.

### SetSpaRelativeSourcePathNil

`func (o *BusinessApplicationCreateDto) SetSpaRelativeSourcePathNil(b bool)`

 SetSpaRelativeSourcePathNil sets the value for SpaRelativeSourcePath to be an explicit nil

### UnsetSpaRelativeSourcePath
`func (o *BusinessApplicationCreateDto) UnsetSpaRelativeSourcePath()`

UnsetSpaRelativeSourcePath ensures that no value is present for SpaRelativeSourcePath, not even an explicit nil
### GetSpaRelativeOutputPath

`func (o *BusinessApplicationCreateDto) GetSpaRelativeOutputPath() string`

GetSpaRelativeOutputPath returns the SpaRelativeOutputPath field if non-nil, zero value otherwise.

### GetSpaRelativeOutputPathOk

`func (o *BusinessApplicationCreateDto) GetSpaRelativeOutputPathOk() (*string, bool)`

GetSpaRelativeOutputPathOk returns a tuple with the SpaRelativeOutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeOutputPath

`func (o *BusinessApplicationCreateDto) SetSpaRelativeOutputPath(v string)`

SetSpaRelativeOutputPath sets SpaRelativeOutputPath field to given value.

### HasSpaRelativeOutputPath

`func (o *BusinessApplicationCreateDto) HasSpaRelativeOutputPath() bool`

HasSpaRelativeOutputPath returns a boolean if a field has been set.

### SetSpaRelativeOutputPathNil

`func (o *BusinessApplicationCreateDto) SetSpaRelativeOutputPathNil(b bool)`

 SetSpaRelativeOutputPathNil sets the value for SpaRelativeOutputPath to be an explicit nil

### UnsetSpaRelativeOutputPath
`func (o *BusinessApplicationCreateDto) UnsetSpaRelativeOutputPath()`

UnsetSpaRelativeOutputPath ensures that no value is present for SpaRelativeOutputPath, not even an explicit nil
### GetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationCreateDto) GetUseProxyToSpaDevelopmentServer() bool`

GetUseProxyToSpaDevelopmentServer returns the UseProxyToSpaDevelopmentServer field if non-nil, zero value otherwise.

### GetUseProxyToSpaDevelopmentServerOk

`func (o *BusinessApplicationCreateDto) GetUseProxyToSpaDevelopmentServerOk() (*bool, bool)`

GetUseProxyToSpaDevelopmentServerOk returns a tuple with the UseProxyToSpaDevelopmentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationCreateDto) SetUseProxyToSpaDevelopmentServer(v bool)`

SetUseProxyToSpaDevelopmentServer sets UseProxyToSpaDevelopmentServer field to given value.

### HasUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationCreateDto) HasUseProxyToSpaDevelopmentServer() bool`

HasUseProxyToSpaDevelopmentServer returns a boolean if a field has been set.

### GetSpaDevelopmentServerUri

`func (o *BusinessApplicationCreateDto) GetSpaDevelopmentServerUri() string`

GetSpaDevelopmentServerUri returns the SpaDevelopmentServerUri field if non-nil, zero value otherwise.

### GetSpaDevelopmentServerUriOk

`func (o *BusinessApplicationCreateDto) GetSpaDevelopmentServerUriOk() (*string, bool)`

GetSpaDevelopmentServerUriOk returns a tuple with the SpaDevelopmentServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaDevelopmentServerUri

`func (o *BusinessApplicationCreateDto) SetSpaDevelopmentServerUri(v string)`

SetSpaDevelopmentServerUri sets SpaDevelopmentServerUri field to given value.

### HasSpaDevelopmentServerUri

`func (o *BusinessApplicationCreateDto) HasSpaDevelopmentServerUri() bool`

HasSpaDevelopmentServerUri returns a boolean if a field has been set.

### SetSpaDevelopmentServerUriNil

`func (o *BusinessApplicationCreateDto) SetSpaDevelopmentServerUriNil(b bool)`

 SetSpaDevelopmentServerUriNil sets the value for SpaDevelopmentServerUri to be an explicit nil

### UnsetSpaDevelopmentServerUri
`func (o *BusinessApplicationCreateDto) UnsetSpaDevelopmentServerUri()`

UnsetSpaDevelopmentServerUri ensures that no value is present for SpaDevelopmentServerUri, not even an explicit nil
### GetEnableGitRepoManagement

`func (o *BusinessApplicationCreateDto) GetEnableGitRepoManagement() bool`

GetEnableGitRepoManagement returns the EnableGitRepoManagement field if non-nil, zero value otherwise.

### GetEnableGitRepoManagementOk

`func (o *BusinessApplicationCreateDto) GetEnableGitRepoManagementOk() (*bool, bool)`

GetEnableGitRepoManagementOk returns a tuple with the EnableGitRepoManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGitRepoManagement

`func (o *BusinessApplicationCreateDto) SetEnableGitRepoManagement(v bool)`

SetEnableGitRepoManagement sets EnableGitRepoManagement field to given value.

### HasEnableGitRepoManagement

`func (o *BusinessApplicationCreateDto) HasEnableGitRepoManagement() bool`

HasEnableGitRepoManagement returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *BusinessApplicationCreateDto) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *BusinessApplicationCreateDto) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *BusinessApplicationCreateDto) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *BusinessApplicationCreateDto) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### SetGitRepoUrlNil

`func (o *BusinessApplicationCreateDto) SetGitRepoUrlNil(b bool)`

 SetGitRepoUrlNil sets the value for GitRepoUrl to be an explicit nil

### UnsetGitRepoUrl
`func (o *BusinessApplicationCreateDto) UnsetGitRepoUrl()`

UnsetGitRepoUrl ensures that no value is present for GitRepoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


