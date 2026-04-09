# BusinessApplicationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**AvatarURL** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**IsMultiTenant** | Pointer to **NullableBool** |  | [optional] 
**IsVerified** | Pointer to **NullableBool** |  | [optional] 
**IsDisabled** | Pointer to **NullableBool** |  | [optional] 
**IsSinglePageApplication** | Pointer to **NullableBool** |  | [optional] 
**IsNativeOrDesktopApp** | Pointer to **NullableBool** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicyURL** | Pointer to **NullableString** |  | [optional] 
**TermsAndConditionsURL** | Pointer to **NullableString** |  | [optional] 
**RequireHttps** | Pointer to **NullableBool** |  | [optional] 
**RequireAppSecret** | Pointer to **NullableBool** |  | [optional] 
**EnableClientOauthLogin** | Pointer to **NullableBool** |  | [optional] 
**EnableWebOAuthLogin** | Pointer to **NullableBool** |  | [optional] 
**EnableDeviceOAuthLogin** | Pointer to **NullableBool** |  | [optional] 
**AllowAccessToSuiteSettings** | Pointer to **NullableBool** |  | [optional] 
**RequireWebOAuthReauthentication** | Pointer to **NullableBool** |  | [optional] 
**RequireTwoFactorReauthorization** | Pointer to **NullableBool** |  | [optional] 
**EnableEmbeddedBrowserOAuthLogin** | Pointer to **NullableBool** |  | [optional] 
**UseStrictModeForRedirectURIs** | Pointer to **NullableBool** |  | [optional] 
**CountryRestricted** | Pointer to **NullableBool** |  | [optional] 
**SpaUIEngine** | Pointer to **NullableString** |  | [optional] 
**SpaStaticFilesRootPath** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeAppPath** | Pointer to **NullableString** |  | [optional] 
**SpaNpmStartScript** | Pointer to **NullableString** |  | [optional] 
**SpaNpmPublishScript** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeSourcePath** | Pointer to **NullableString** |  | [optional] 
**SpaRelativeOutputPath** | Pointer to **NullableString** |  | [optional] 
**UseProxyToSpaDevelopmentServer** | Pointer to **NullableBool** |  | [optional] 
**SpaDevelopmentServerUri** | Pointer to **NullableString** |  | [optional] 
**EnableGitRepoManagement** | Pointer to **NullableBool** |  | [optional] 
**GitRepoUrl** | Pointer to **NullableString** |  | [optional] 
**MarkedForPublish** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBusinessApplicationUpdateDto

`func NewBusinessApplicationUpdateDto() *BusinessApplicationUpdateDto`

NewBusinessApplicationUpdateDto instantiates a new BusinessApplicationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessApplicationUpdateDtoWithDefaults

`func NewBusinessApplicationUpdateDtoWithDefaults() *BusinessApplicationUpdateDto`

NewBusinessApplicationUpdateDtoWithDefaults instantiates a new BusinessApplicationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BusinessApplicationUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessApplicationUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessApplicationUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessApplicationUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BusinessApplicationUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BusinessApplicationUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *BusinessApplicationUpdateDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BusinessApplicationUpdateDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BusinessApplicationUpdateDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BusinessApplicationUpdateDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *BusinessApplicationUpdateDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *BusinessApplicationUpdateDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetDisplayName

`func (o *BusinessApplicationUpdateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BusinessApplicationUpdateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BusinessApplicationUpdateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BusinessApplicationUpdateDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BusinessApplicationUpdateDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BusinessApplicationUpdateDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatarURL

`func (o *BusinessApplicationUpdateDto) GetAvatarURL() string`

GetAvatarURL returns the AvatarURL field if non-nil, zero value otherwise.

### GetAvatarURLOk

`func (o *BusinessApplicationUpdateDto) GetAvatarURLOk() (*string, bool)`

GetAvatarURLOk returns a tuple with the AvatarURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarURL

`func (o *BusinessApplicationUpdateDto) SetAvatarURL(v string)`

SetAvatarURL sets AvatarURL field to given value.

### HasAvatarURL

`func (o *BusinessApplicationUpdateDto) HasAvatarURL() bool`

HasAvatarURL returns a boolean if a field has been set.

### SetAvatarURLNil

`func (o *BusinessApplicationUpdateDto) SetAvatarURLNil(b bool)`

 SetAvatarURLNil sets the value for AvatarURL to be an explicit nil

### UnsetAvatarURL
`func (o *BusinessApplicationUpdateDto) UnsetAvatarURL()`

UnsetAvatarURL ensures that no value is present for AvatarURL, not even an explicit nil
### GetWebsiteUrl

`func (o *BusinessApplicationUpdateDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *BusinessApplicationUpdateDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *BusinessApplicationUpdateDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *BusinessApplicationUpdateDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *BusinessApplicationUpdateDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *BusinessApplicationUpdateDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetIsMultiTenant

`func (o *BusinessApplicationUpdateDto) GetIsMultiTenant() bool`

GetIsMultiTenant returns the IsMultiTenant field if non-nil, zero value otherwise.

### GetIsMultiTenantOk

`func (o *BusinessApplicationUpdateDto) GetIsMultiTenantOk() (*bool, bool)`

GetIsMultiTenantOk returns a tuple with the IsMultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiTenant

`func (o *BusinessApplicationUpdateDto) SetIsMultiTenant(v bool)`

SetIsMultiTenant sets IsMultiTenant field to given value.

### HasIsMultiTenant

`func (o *BusinessApplicationUpdateDto) HasIsMultiTenant() bool`

HasIsMultiTenant returns a boolean if a field has been set.

### SetIsMultiTenantNil

`func (o *BusinessApplicationUpdateDto) SetIsMultiTenantNil(b bool)`

 SetIsMultiTenantNil sets the value for IsMultiTenant to be an explicit nil

### UnsetIsMultiTenant
`func (o *BusinessApplicationUpdateDto) UnsetIsMultiTenant()`

UnsetIsMultiTenant ensures that no value is present for IsMultiTenant, not even an explicit nil
### GetIsVerified

`func (o *BusinessApplicationUpdateDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *BusinessApplicationUpdateDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *BusinessApplicationUpdateDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *BusinessApplicationUpdateDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### SetIsVerifiedNil

`func (o *BusinessApplicationUpdateDto) SetIsVerifiedNil(b bool)`

 SetIsVerifiedNil sets the value for IsVerified to be an explicit nil

### UnsetIsVerified
`func (o *BusinessApplicationUpdateDto) UnsetIsVerified()`

UnsetIsVerified ensures that no value is present for IsVerified, not even an explicit nil
### GetIsDisabled

`func (o *BusinessApplicationUpdateDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *BusinessApplicationUpdateDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *BusinessApplicationUpdateDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *BusinessApplicationUpdateDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### SetIsDisabledNil

`func (o *BusinessApplicationUpdateDto) SetIsDisabledNil(b bool)`

 SetIsDisabledNil sets the value for IsDisabled to be an explicit nil

### UnsetIsDisabled
`func (o *BusinessApplicationUpdateDto) UnsetIsDisabled()`

UnsetIsDisabled ensures that no value is present for IsDisabled, not even an explicit nil
### GetIsSinglePageApplication

`func (o *BusinessApplicationUpdateDto) GetIsSinglePageApplication() bool`

GetIsSinglePageApplication returns the IsSinglePageApplication field if non-nil, zero value otherwise.

### GetIsSinglePageApplicationOk

`func (o *BusinessApplicationUpdateDto) GetIsSinglePageApplicationOk() (*bool, bool)`

GetIsSinglePageApplicationOk returns a tuple with the IsSinglePageApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSinglePageApplication

`func (o *BusinessApplicationUpdateDto) SetIsSinglePageApplication(v bool)`

SetIsSinglePageApplication sets IsSinglePageApplication field to given value.

### HasIsSinglePageApplication

`func (o *BusinessApplicationUpdateDto) HasIsSinglePageApplication() bool`

HasIsSinglePageApplication returns a boolean if a field has been set.

### SetIsSinglePageApplicationNil

`func (o *BusinessApplicationUpdateDto) SetIsSinglePageApplicationNil(b bool)`

 SetIsSinglePageApplicationNil sets the value for IsSinglePageApplication to be an explicit nil

### UnsetIsSinglePageApplication
`func (o *BusinessApplicationUpdateDto) UnsetIsSinglePageApplication()`

UnsetIsSinglePageApplication ensures that no value is present for IsSinglePageApplication, not even an explicit nil
### GetIsNativeOrDesktopApp

`func (o *BusinessApplicationUpdateDto) GetIsNativeOrDesktopApp() bool`

GetIsNativeOrDesktopApp returns the IsNativeOrDesktopApp field if non-nil, zero value otherwise.

### GetIsNativeOrDesktopAppOk

`func (o *BusinessApplicationUpdateDto) GetIsNativeOrDesktopAppOk() (*bool, bool)`

GetIsNativeOrDesktopAppOk returns a tuple with the IsNativeOrDesktopApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeOrDesktopApp

`func (o *BusinessApplicationUpdateDto) SetIsNativeOrDesktopApp(v bool)`

SetIsNativeOrDesktopApp sets IsNativeOrDesktopApp field to given value.

### HasIsNativeOrDesktopApp

`func (o *BusinessApplicationUpdateDto) HasIsNativeOrDesktopApp() bool`

HasIsNativeOrDesktopApp returns a boolean if a field has been set.

### SetIsNativeOrDesktopAppNil

`func (o *BusinessApplicationUpdateDto) SetIsNativeOrDesktopAppNil(b bool)`

 SetIsNativeOrDesktopAppNil sets the value for IsNativeOrDesktopApp to be an explicit nil

### UnsetIsNativeOrDesktopApp
`func (o *BusinessApplicationUpdateDto) UnsetIsNativeOrDesktopApp()`

UnsetIsNativeOrDesktopApp ensures that no value is present for IsNativeOrDesktopApp, not even an explicit nil
### GetContactEmail

`func (o *BusinessApplicationUpdateDto) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *BusinessApplicationUpdateDto) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *BusinessApplicationUpdateDto) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *BusinessApplicationUpdateDto) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *BusinessApplicationUpdateDto) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *BusinessApplicationUpdateDto) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetPrivacyPolicyURL

`func (o *BusinessApplicationUpdateDto) GetPrivacyPolicyURL() string`

GetPrivacyPolicyURL returns the PrivacyPolicyURL field if non-nil, zero value otherwise.

### GetPrivacyPolicyURLOk

`func (o *BusinessApplicationUpdateDto) GetPrivacyPolicyURLOk() (*string, bool)`

GetPrivacyPolicyURLOk returns a tuple with the PrivacyPolicyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyURL

`func (o *BusinessApplicationUpdateDto) SetPrivacyPolicyURL(v string)`

SetPrivacyPolicyURL sets PrivacyPolicyURL field to given value.

### HasPrivacyPolicyURL

`func (o *BusinessApplicationUpdateDto) HasPrivacyPolicyURL() bool`

HasPrivacyPolicyURL returns a boolean if a field has been set.

### SetPrivacyPolicyURLNil

`func (o *BusinessApplicationUpdateDto) SetPrivacyPolicyURLNil(b bool)`

 SetPrivacyPolicyURLNil sets the value for PrivacyPolicyURL to be an explicit nil

### UnsetPrivacyPolicyURL
`func (o *BusinessApplicationUpdateDto) UnsetPrivacyPolicyURL()`

UnsetPrivacyPolicyURL ensures that no value is present for PrivacyPolicyURL, not even an explicit nil
### GetTermsAndConditionsURL

`func (o *BusinessApplicationUpdateDto) GetTermsAndConditionsURL() string`

GetTermsAndConditionsURL returns the TermsAndConditionsURL field if non-nil, zero value otherwise.

### GetTermsAndConditionsURLOk

`func (o *BusinessApplicationUpdateDto) GetTermsAndConditionsURLOk() (*string, bool)`

GetTermsAndConditionsURLOk returns a tuple with the TermsAndConditionsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsURL

`func (o *BusinessApplicationUpdateDto) SetTermsAndConditionsURL(v string)`

SetTermsAndConditionsURL sets TermsAndConditionsURL field to given value.

### HasTermsAndConditionsURL

`func (o *BusinessApplicationUpdateDto) HasTermsAndConditionsURL() bool`

HasTermsAndConditionsURL returns a boolean if a field has been set.

### SetTermsAndConditionsURLNil

`func (o *BusinessApplicationUpdateDto) SetTermsAndConditionsURLNil(b bool)`

 SetTermsAndConditionsURLNil sets the value for TermsAndConditionsURL to be an explicit nil

### UnsetTermsAndConditionsURL
`func (o *BusinessApplicationUpdateDto) UnsetTermsAndConditionsURL()`

UnsetTermsAndConditionsURL ensures that no value is present for TermsAndConditionsURL, not even an explicit nil
### GetRequireHttps

`func (o *BusinessApplicationUpdateDto) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *BusinessApplicationUpdateDto) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *BusinessApplicationUpdateDto) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *BusinessApplicationUpdateDto) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### SetRequireHttpsNil

`func (o *BusinessApplicationUpdateDto) SetRequireHttpsNil(b bool)`

 SetRequireHttpsNil sets the value for RequireHttps to be an explicit nil

### UnsetRequireHttps
`func (o *BusinessApplicationUpdateDto) UnsetRequireHttps()`

UnsetRequireHttps ensures that no value is present for RequireHttps, not even an explicit nil
### GetRequireAppSecret

`func (o *BusinessApplicationUpdateDto) GetRequireAppSecret() bool`

GetRequireAppSecret returns the RequireAppSecret field if non-nil, zero value otherwise.

### GetRequireAppSecretOk

`func (o *BusinessApplicationUpdateDto) GetRequireAppSecretOk() (*bool, bool)`

GetRequireAppSecretOk returns a tuple with the RequireAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAppSecret

`func (o *BusinessApplicationUpdateDto) SetRequireAppSecret(v bool)`

SetRequireAppSecret sets RequireAppSecret field to given value.

### HasRequireAppSecret

`func (o *BusinessApplicationUpdateDto) HasRequireAppSecret() bool`

HasRequireAppSecret returns a boolean if a field has been set.

### SetRequireAppSecretNil

`func (o *BusinessApplicationUpdateDto) SetRequireAppSecretNil(b bool)`

 SetRequireAppSecretNil sets the value for RequireAppSecret to be an explicit nil

### UnsetRequireAppSecret
`func (o *BusinessApplicationUpdateDto) UnsetRequireAppSecret()`

UnsetRequireAppSecret ensures that no value is present for RequireAppSecret, not even an explicit nil
### GetEnableClientOauthLogin

`func (o *BusinessApplicationUpdateDto) GetEnableClientOauthLogin() bool`

GetEnableClientOauthLogin returns the EnableClientOauthLogin field if non-nil, zero value otherwise.

### GetEnableClientOauthLoginOk

`func (o *BusinessApplicationUpdateDto) GetEnableClientOauthLoginOk() (*bool, bool)`

GetEnableClientOauthLoginOk returns a tuple with the EnableClientOauthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientOauthLogin

`func (o *BusinessApplicationUpdateDto) SetEnableClientOauthLogin(v bool)`

SetEnableClientOauthLogin sets EnableClientOauthLogin field to given value.

### HasEnableClientOauthLogin

`func (o *BusinessApplicationUpdateDto) HasEnableClientOauthLogin() bool`

HasEnableClientOauthLogin returns a boolean if a field has been set.

### SetEnableClientOauthLoginNil

`func (o *BusinessApplicationUpdateDto) SetEnableClientOauthLoginNil(b bool)`

 SetEnableClientOauthLoginNil sets the value for EnableClientOauthLogin to be an explicit nil

### UnsetEnableClientOauthLogin
`func (o *BusinessApplicationUpdateDto) UnsetEnableClientOauthLogin()`

UnsetEnableClientOauthLogin ensures that no value is present for EnableClientOauthLogin, not even an explicit nil
### GetEnableWebOAuthLogin

`func (o *BusinessApplicationUpdateDto) GetEnableWebOAuthLogin() bool`

GetEnableWebOAuthLogin returns the EnableWebOAuthLogin field if non-nil, zero value otherwise.

### GetEnableWebOAuthLoginOk

`func (o *BusinessApplicationUpdateDto) GetEnableWebOAuthLoginOk() (*bool, bool)`

GetEnableWebOAuthLoginOk returns a tuple with the EnableWebOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebOAuthLogin

`func (o *BusinessApplicationUpdateDto) SetEnableWebOAuthLogin(v bool)`

SetEnableWebOAuthLogin sets EnableWebOAuthLogin field to given value.

### HasEnableWebOAuthLogin

`func (o *BusinessApplicationUpdateDto) HasEnableWebOAuthLogin() bool`

HasEnableWebOAuthLogin returns a boolean if a field has been set.

### SetEnableWebOAuthLoginNil

`func (o *BusinessApplicationUpdateDto) SetEnableWebOAuthLoginNil(b bool)`

 SetEnableWebOAuthLoginNil sets the value for EnableWebOAuthLogin to be an explicit nil

### UnsetEnableWebOAuthLogin
`func (o *BusinessApplicationUpdateDto) UnsetEnableWebOAuthLogin()`

UnsetEnableWebOAuthLogin ensures that no value is present for EnableWebOAuthLogin, not even an explicit nil
### GetEnableDeviceOAuthLogin

`func (o *BusinessApplicationUpdateDto) GetEnableDeviceOAuthLogin() bool`

GetEnableDeviceOAuthLogin returns the EnableDeviceOAuthLogin field if non-nil, zero value otherwise.

### GetEnableDeviceOAuthLoginOk

`func (o *BusinessApplicationUpdateDto) GetEnableDeviceOAuthLoginOk() (*bool, bool)`

GetEnableDeviceOAuthLoginOk returns a tuple with the EnableDeviceOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceOAuthLogin

`func (o *BusinessApplicationUpdateDto) SetEnableDeviceOAuthLogin(v bool)`

SetEnableDeviceOAuthLogin sets EnableDeviceOAuthLogin field to given value.

### HasEnableDeviceOAuthLogin

`func (o *BusinessApplicationUpdateDto) HasEnableDeviceOAuthLogin() bool`

HasEnableDeviceOAuthLogin returns a boolean if a field has been set.

### SetEnableDeviceOAuthLoginNil

`func (o *BusinessApplicationUpdateDto) SetEnableDeviceOAuthLoginNil(b bool)`

 SetEnableDeviceOAuthLoginNil sets the value for EnableDeviceOAuthLogin to be an explicit nil

### UnsetEnableDeviceOAuthLogin
`func (o *BusinessApplicationUpdateDto) UnsetEnableDeviceOAuthLogin()`

UnsetEnableDeviceOAuthLogin ensures that no value is present for EnableDeviceOAuthLogin, not even an explicit nil
### GetAllowAccessToSuiteSettings

`func (o *BusinessApplicationUpdateDto) GetAllowAccessToSuiteSettings() bool`

GetAllowAccessToSuiteSettings returns the AllowAccessToSuiteSettings field if non-nil, zero value otherwise.

### GetAllowAccessToSuiteSettingsOk

`func (o *BusinessApplicationUpdateDto) GetAllowAccessToSuiteSettingsOk() (*bool, bool)`

GetAllowAccessToSuiteSettingsOk returns a tuple with the AllowAccessToSuiteSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessToSuiteSettings

`func (o *BusinessApplicationUpdateDto) SetAllowAccessToSuiteSettings(v bool)`

SetAllowAccessToSuiteSettings sets AllowAccessToSuiteSettings field to given value.

### HasAllowAccessToSuiteSettings

`func (o *BusinessApplicationUpdateDto) HasAllowAccessToSuiteSettings() bool`

HasAllowAccessToSuiteSettings returns a boolean if a field has been set.

### SetAllowAccessToSuiteSettingsNil

`func (o *BusinessApplicationUpdateDto) SetAllowAccessToSuiteSettingsNil(b bool)`

 SetAllowAccessToSuiteSettingsNil sets the value for AllowAccessToSuiteSettings to be an explicit nil

### UnsetAllowAccessToSuiteSettings
`func (o *BusinessApplicationUpdateDto) UnsetAllowAccessToSuiteSettings()`

UnsetAllowAccessToSuiteSettings ensures that no value is present for AllowAccessToSuiteSettings, not even an explicit nil
### GetRequireWebOAuthReauthentication

`func (o *BusinessApplicationUpdateDto) GetRequireWebOAuthReauthentication() bool`

GetRequireWebOAuthReauthentication returns the RequireWebOAuthReauthentication field if non-nil, zero value otherwise.

### GetRequireWebOAuthReauthenticationOk

`func (o *BusinessApplicationUpdateDto) GetRequireWebOAuthReauthenticationOk() (*bool, bool)`

GetRequireWebOAuthReauthenticationOk returns a tuple with the RequireWebOAuthReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireWebOAuthReauthentication

`func (o *BusinessApplicationUpdateDto) SetRequireWebOAuthReauthentication(v bool)`

SetRequireWebOAuthReauthentication sets RequireWebOAuthReauthentication field to given value.

### HasRequireWebOAuthReauthentication

`func (o *BusinessApplicationUpdateDto) HasRequireWebOAuthReauthentication() bool`

HasRequireWebOAuthReauthentication returns a boolean if a field has been set.

### SetRequireWebOAuthReauthenticationNil

`func (o *BusinessApplicationUpdateDto) SetRequireWebOAuthReauthenticationNil(b bool)`

 SetRequireWebOAuthReauthenticationNil sets the value for RequireWebOAuthReauthentication to be an explicit nil

### UnsetRequireWebOAuthReauthentication
`func (o *BusinessApplicationUpdateDto) UnsetRequireWebOAuthReauthentication()`

UnsetRequireWebOAuthReauthentication ensures that no value is present for RequireWebOAuthReauthentication, not even an explicit nil
### GetRequireTwoFactorReauthorization

`func (o *BusinessApplicationUpdateDto) GetRequireTwoFactorReauthorization() bool`

GetRequireTwoFactorReauthorization returns the RequireTwoFactorReauthorization field if non-nil, zero value otherwise.

### GetRequireTwoFactorReauthorizationOk

`func (o *BusinessApplicationUpdateDto) GetRequireTwoFactorReauthorizationOk() (*bool, bool)`

GetRequireTwoFactorReauthorizationOk returns a tuple with the RequireTwoFactorReauthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTwoFactorReauthorization

`func (o *BusinessApplicationUpdateDto) SetRequireTwoFactorReauthorization(v bool)`

SetRequireTwoFactorReauthorization sets RequireTwoFactorReauthorization field to given value.

### HasRequireTwoFactorReauthorization

`func (o *BusinessApplicationUpdateDto) HasRequireTwoFactorReauthorization() bool`

HasRequireTwoFactorReauthorization returns a boolean if a field has been set.

### SetRequireTwoFactorReauthorizationNil

`func (o *BusinessApplicationUpdateDto) SetRequireTwoFactorReauthorizationNil(b bool)`

 SetRequireTwoFactorReauthorizationNil sets the value for RequireTwoFactorReauthorization to be an explicit nil

### UnsetRequireTwoFactorReauthorization
`func (o *BusinessApplicationUpdateDto) UnsetRequireTwoFactorReauthorization()`

UnsetRequireTwoFactorReauthorization ensures that no value is present for RequireTwoFactorReauthorization, not even an explicit nil
### GetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationUpdateDto) GetEnableEmbeddedBrowserOAuthLogin() bool`

GetEnableEmbeddedBrowserOAuthLogin returns the EnableEmbeddedBrowserOAuthLogin field if non-nil, zero value otherwise.

### GetEnableEmbeddedBrowserOAuthLoginOk

`func (o *BusinessApplicationUpdateDto) GetEnableEmbeddedBrowserOAuthLoginOk() (*bool, bool)`

GetEnableEmbeddedBrowserOAuthLoginOk returns a tuple with the EnableEmbeddedBrowserOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationUpdateDto) SetEnableEmbeddedBrowserOAuthLogin(v bool)`

SetEnableEmbeddedBrowserOAuthLogin sets EnableEmbeddedBrowserOAuthLogin field to given value.

### HasEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationUpdateDto) HasEnableEmbeddedBrowserOAuthLogin() bool`

HasEnableEmbeddedBrowserOAuthLogin returns a boolean if a field has been set.

### SetEnableEmbeddedBrowserOAuthLoginNil

`func (o *BusinessApplicationUpdateDto) SetEnableEmbeddedBrowserOAuthLoginNil(b bool)`

 SetEnableEmbeddedBrowserOAuthLoginNil sets the value for EnableEmbeddedBrowserOAuthLogin to be an explicit nil

### UnsetEnableEmbeddedBrowserOAuthLogin
`func (o *BusinessApplicationUpdateDto) UnsetEnableEmbeddedBrowserOAuthLogin()`

UnsetEnableEmbeddedBrowserOAuthLogin ensures that no value is present for EnableEmbeddedBrowserOAuthLogin, not even an explicit nil
### GetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationUpdateDto) GetUseStrictModeForRedirectURIs() bool`

GetUseStrictModeForRedirectURIs returns the UseStrictModeForRedirectURIs field if non-nil, zero value otherwise.

### GetUseStrictModeForRedirectURIsOk

`func (o *BusinessApplicationUpdateDto) GetUseStrictModeForRedirectURIsOk() (*bool, bool)`

GetUseStrictModeForRedirectURIsOk returns a tuple with the UseStrictModeForRedirectURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationUpdateDto) SetUseStrictModeForRedirectURIs(v bool)`

SetUseStrictModeForRedirectURIs sets UseStrictModeForRedirectURIs field to given value.

### HasUseStrictModeForRedirectURIs

`func (o *BusinessApplicationUpdateDto) HasUseStrictModeForRedirectURIs() bool`

HasUseStrictModeForRedirectURIs returns a boolean if a field has been set.

### SetUseStrictModeForRedirectURIsNil

`func (o *BusinessApplicationUpdateDto) SetUseStrictModeForRedirectURIsNil(b bool)`

 SetUseStrictModeForRedirectURIsNil sets the value for UseStrictModeForRedirectURIs to be an explicit nil

### UnsetUseStrictModeForRedirectURIs
`func (o *BusinessApplicationUpdateDto) UnsetUseStrictModeForRedirectURIs()`

UnsetUseStrictModeForRedirectURIs ensures that no value is present for UseStrictModeForRedirectURIs, not even an explicit nil
### GetCountryRestricted

`func (o *BusinessApplicationUpdateDto) GetCountryRestricted() bool`

GetCountryRestricted returns the CountryRestricted field if non-nil, zero value otherwise.

### GetCountryRestrictedOk

`func (o *BusinessApplicationUpdateDto) GetCountryRestrictedOk() (*bool, bool)`

GetCountryRestrictedOk returns a tuple with the CountryRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRestricted

`func (o *BusinessApplicationUpdateDto) SetCountryRestricted(v bool)`

SetCountryRestricted sets CountryRestricted field to given value.

### HasCountryRestricted

`func (o *BusinessApplicationUpdateDto) HasCountryRestricted() bool`

HasCountryRestricted returns a boolean if a field has been set.

### SetCountryRestrictedNil

`func (o *BusinessApplicationUpdateDto) SetCountryRestrictedNil(b bool)`

 SetCountryRestrictedNil sets the value for CountryRestricted to be an explicit nil

### UnsetCountryRestricted
`func (o *BusinessApplicationUpdateDto) UnsetCountryRestricted()`

UnsetCountryRestricted ensures that no value is present for CountryRestricted, not even an explicit nil
### GetSpaUIEngine

`func (o *BusinessApplicationUpdateDto) GetSpaUIEngine() string`

GetSpaUIEngine returns the SpaUIEngine field if non-nil, zero value otherwise.

### GetSpaUIEngineOk

`func (o *BusinessApplicationUpdateDto) GetSpaUIEngineOk() (*string, bool)`

GetSpaUIEngineOk returns a tuple with the SpaUIEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUIEngine

`func (o *BusinessApplicationUpdateDto) SetSpaUIEngine(v string)`

SetSpaUIEngine sets SpaUIEngine field to given value.

### HasSpaUIEngine

`func (o *BusinessApplicationUpdateDto) HasSpaUIEngine() bool`

HasSpaUIEngine returns a boolean if a field has been set.

### SetSpaUIEngineNil

`func (o *BusinessApplicationUpdateDto) SetSpaUIEngineNil(b bool)`

 SetSpaUIEngineNil sets the value for SpaUIEngine to be an explicit nil

### UnsetSpaUIEngine
`func (o *BusinessApplicationUpdateDto) UnsetSpaUIEngine()`

UnsetSpaUIEngine ensures that no value is present for SpaUIEngine, not even an explicit nil
### GetSpaStaticFilesRootPath

`func (o *BusinessApplicationUpdateDto) GetSpaStaticFilesRootPath() string`

GetSpaStaticFilesRootPath returns the SpaStaticFilesRootPath field if non-nil, zero value otherwise.

### GetSpaStaticFilesRootPathOk

`func (o *BusinessApplicationUpdateDto) GetSpaStaticFilesRootPathOk() (*string, bool)`

GetSpaStaticFilesRootPathOk returns a tuple with the SpaStaticFilesRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaStaticFilesRootPath

`func (o *BusinessApplicationUpdateDto) SetSpaStaticFilesRootPath(v string)`

SetSpaStaticFilesRootPath sets SpaStaticFilesRootPath field to given value.

### HasSpaStaticFilesRootPath

`func (o *BusinessApplicationUpdateDto) HasSpaStaticFilesRootPath() bool`

HasSpaStaticFilesRootPath returns a boolean if a field has been set.

### SetSpaStaticFilesRootPathNil

`func (o *BusinessApplicationUpdateDto) SetSpaStaticFilesRootPathNil(b bool)`

 SetSpaStaticFilesRootPathNil sets the value for SpaStaticFilesRootPath to be an explicit nil

### UnsetSpaStaticFilesRootPath
`func (o *BusinessApplicationUpdateDto) UnsetSpaStaticFilesRootPath()`

UnsetSpaStaticFilesRootPath ensures that no value is present for SpaStaticFilesRootPath, not even an explicit nil
### GetSpaRelativeAppPath

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeAppPath() string`

GetSpaRelativeAppPath returns the SpaRelativeAppPath field if non-nil, zero value otherwise.

### GetSpaRelativeAppPathOk

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeAppPathOk() (*string, bool)`

GetSpaRelativeAppPathOk returns a tuple with the SpaRelativeAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeAppPath

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeAppPath(v string)`

SetSpaRelativeAppPath sets SpaRelativeAppPath field to given value.

### HasSpaRelativeAppPath

`func (o *BusinessApplicationUpdateDto) HasSpaRelativeAppPath() bool`

HasSpaRelativeAppPath returns a boolean if a field has been set.

### SetSpaRelativeAppPathNil

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeAppPathNil(b bool)`

 SetSpaRelativeAppPathNil sets the value for SpaRelativeAppPath to be an explicit nil

### UnsetSpaRelativeAppPath
`func (o *BusinessApplicationUpdateDto) UnsetSpaRelativeAppPath()`

UnsetSpaRelativeAppPath ensures that no value is present for SpaRelativeAppPath, not even an explicit nil
### GetSpaNpmStartScript

`func (o *BusinessApplicationUpdateDto) GetSpaNpmStartScript() string`

GetSpaNpmStartScript returns the SpaNpmStartScript field if non-nil, zero value otherwise.

### GetSpaNpmStartScriptOk

`func (o *BusinessApplicationUpdateDto) GetSpaNpmStartScriptOk() (*string, bool)`

GetSpaNpmStartScriptOk returns a tuple with the SpaNpmStartScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmStartScript

`func (o *BusinessApplicationUpdateDto) SetSpaNpmStartScript(v string)`

SetSpaNpmStartScript sets SpaNpmStartScript field to given value.

### HasSpaNpmStartScript

`func (o *BusinessApplicationUpdateDto) HasSpaNpmStartScript() bool`

HasSpaNpmStartScript returns a boolean if a field has been set.

### SetSpaNpmStartScriptNil

`func (o *BusinessApplicationUpdateDto) SetSpaNpmStartScriptNil(b bool)`

 SetSpaNpmStartScriptNil sets the value for SpaNpmStartScript to be an explicit nil

### UnsetSpaNpmStartScript
`func (o *BusinessApplicationUpdateDto) UnsetSpaNpmStartScript()`

UnsetSpaNpmStartScript ensures that no value is present for SpaNpmStartScript, not even an explicit nil
### GetSpaNpmPublishScript

`func (o *BusinessApplicationUpdateDto) GetSpaNpmPublishScript() string`

GetSpaNpmPublishScript returns the SpaNpmPublishScript field if non-nil, zero value otherwise.

### GetSpaNpmPublishScriptOk

`func (o *BusinessApplicationUpdateDto) GetSpaNpmPublishScriptOk() (*string, bool)`

GetSpaNpmPublishScriptOk returns a tuple with the SpaNpmPublishScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmPublishScript

`func (o *BusinessApplicationUpdateDto) SetSpaNpmPublishScript(v string)`

SetSpaNpmPublishScript sets SpaNpmPublishScript field to given value.

### HasSpaNpmPublishScript

`func (o *BusinessApplicationUpdateDto) HasSpaNpmPublishScript() bool`

HasSpaNpmPublishScript returns a boolean if a field has been set.

### SetSpaNpmPublishScriptNil

`func (o *BusinessApplicationUpdateDto) SetSpaNpmPublishScriptNil(b bool)`

 SetSpaNpmPublishScriptNil sets the value for SpaNpmPublishScript to be an explicit nil

### UnsetSpaNpmPublishScript
`func (o *BusinessApplicationUpdateDto) UnsetSpaNpmPublishScript()`

UnsetSpaNpmPublishScript ensures that no value is present for SpaNpmPublishScript, not even an explicit nil
### GetSpaRelativeSourcePath

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeSourcePath() string`

GetSpaRelativeSourcePath returns the SpaRelativeSourcePath field if non-nil, zero value otherwise.

### GetSpaRelativeSourcePathOk

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeSourcePathOk() (*string, bool)`

GetSpaRelativeSourcePathOk returns a tuple with the SpaRelativeSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeSourcePath

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeSourcePath(v string)`

SetSpaRelativeSourcePath sets SpaRelativeSourcePath field to given value.

### HasSpaRelativeSourcePath

`func (o *BusinessApplicationUpdateDto) HasSpaRelativeSourcePath() bool`

HasSpaRelativeSourcePath returns a boolean if a field has been set.

### SetSpaRelativeSourcePathNil

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeSourcePathNil(b bool)`

 SetSpaRelativeSourcePathNil sets the value for SpaRelativeSourcePath to be an explicit nil

### UnsetSpaRelativeSourcePath
`func (o *BusinessApplicationUpdateDto) UnsetSpaRelativeSourcePath()`

UnsetSpaRelativeSourcePath ensures that no value is present for SpaRelativeSourcePath, not even an explicit nil
### GetSpaRelativeOutputPath

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeOutputPath() string`

GetSpaRelativeOutputPath returns the SpaRelativeOutputPath field if non-nil, zero value otherwise.

### GetSpaRelativeOutputPathOk

`func (o *BusinessApplicationUpdateDto) GetSpaRelativeOutputPathOk() (*string, bool)`

GetSpaRelativeOutputPathOk returns a tuple with the SpaRelativeOutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeOutputPath

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeOutputPath(v string)`

SetSpaRelativeOutputPath sets SpaRelativeOutputPath field to given value.

### HasSpaRelativeOutputPath

`func (o *BusinessApplicationUpdateDto) HasSpaRelativeOutputPath() bool`

HasSpaRelativeOutputPath returns a boolean if a field has been set.

### SetSpaRelativeOutputPathNil

`func (o *BusinessApplicationUpdateDto) SetSpaRelativeOutputPathNil(b bool)`

 SetSpaRelativeOutputPathNil sets the value for SpaRelativeOutputPath to be an explicit nil

### UnsetSpaRelativeOutputPath
`func (o *BusinessApplicationUpdateDto) UnsetSpaRelativeOutputPath()`

UnsetSpaRelativeOutputPath ensures that no value is present for SpaRelativeOutputPath, not even an explicit nil
### GetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationUpdateDto) GetUseProxyToSpaDevelopmentServer() bool`

GetUseProxyToSpaDevelopmentServer returns the UseProxyToSpaDevelopmentServer field if non-nil, zero value otherwise.

### GetUseProxyToSpaDevelopmentServerOk

`func (o *BusinessApplicationUpdateDto) GetUseProxyToSpaDevelopmentServerOk() (*bool, bool)`

GetUseProxyToSpaDevelopmentServerOk returns a tuple with the UseProxyToSpaDevelopmentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationUpdateDto) SetUseProxyToSpaDevelopmentServer(v bool)`

SetUseProxyToSpaDevelopmentServer sets UseProxyToSpaDevelopmentServer field to given value.

### HasUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationUpdateDto) HasUseProxyToSpaDevelopmentServer() bool`

HasUseProxyToSpaDevelopmentServer returns a boolean if a field has been set.

### SetUseProxyToSpaDevelopmentServerNil

`func (o *BusinessApplicationUpdateDto) SetUseProxyToSpaDevelopmentServerNil(b bool)`

 SetUseProxyToSpaDevelopmentServerNil sets the value for UseProxyToSpaDevelopmentServer to be an explicit nil

### UnsetUseProxyToSpaDevelopmentServer
`func (o *BusinessApplicationUpdateDto) UnsetUseProxyToSpaDevelopmentServer()`

UnsetUseProxyToSpaDevelopmentServer ensures that no value is present for UseProxyToSpaDevelopmentServer, not even an explicit nil
### GetSpaDevelopmentServerUri

`func (o *BusinessApplicationUpdateDto) GetSpaDevelopmentServerUri() string`

GetSpaDevelopmentServerUri returns the SpaDevelopmentServerUri field if non-nil, zero value otherwise.

### GetSpaDevelopmentServerUriOk

`func (o *BusinessApplicationUpdateDto) GetSpaDevelopmentServerUriOk() (*string, bool)`

GetSpaDevelopmentServerUriOk returns a tuple with the SpaDevelopmentServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaDevelopmentServerUri

`func (o *BusinessApplicationUpdateDto) SetSpaDevelopmentServerUri(v string)`

SetSpaDevelopmentServerUri sets SpaDevelopmentServerUri field to given value.

### HasSpaDevelopmentServerUri

`func (o *BusinessApplicationUpdateDto) HasSpaDevelopmentServerUri() bool`

HasSpaDevelopmentServerUri returns a boolean if a field has been set.

### SetSpaDevelopmentServerUriNil

`func (o *BusinessApplicationUpdateDto) SetSpaDevelopmentServerUriNil(b bool)`

 SetSpaDevelopmentServerUriNil sets the value for SpaDevelopmentServerUri to be an explicit nil

### UnsetSpaDevelopmentServerUri
`func (o *BusinessApplicationUpdateDto) UnsetSpaDevelopmentServerUri()`

UnsetSpaDevelopmentServerUri ensures that no value is present for SpaDevelopmentServerUri, not even an explicit nil
### GetEnableGitRepoManagement

`func (o *BusinessApplicationUpdateDto) GetEnableGitRepoManagement() bool`

GetEnableGitRepoManagement returns the EnableGitRepoManagement field if non-nil, zero value otherwise.

### GetEnableGitRepoManagementOk

`func (o *BusinessApplicationUpdateDto) GetEnableGitRepoManagementOk() (*bool, bool)`

GetEnableGitRepoManagementOk returns a tuple with the EnableGitRepoManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGitRepoManagement

`func (o *BusinessApplicationUpdateDto) SetEnableGitRepoManagement(v bool)`

SetEnableGitRepoManagement sets EnableGitRepoManagement field to given value.

### HasEnableGitRepoManagement

`func (o *BusinessApplicationUpdateDto) HasEnableGitRepoManagement() bool`

HasEnableGitRepoManagement returns a boolean if a field has been set.

### SetEnableGitRepoManagementNil

`func (o *BusinessApplicationUpdateDto) SetEnableGitRepoManagementNil(b bool)`

 SetEnableGitRepoManagementNil sets the value for EnableGitRepoManagement to be an explicit nil

### UnsetEnableGitRepoManagement
`func (o *BusinessApplicationUpdateDto) UnsetEnableGitRepoManagement()`

UnsetEnableGitRepoManagement ensures that no value is present for EnableGitRepoManagement, not even an explicit nil
### GetGitRepoUrl

`func (o *BusinessApplicationUpdateDto) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *BusinessApplicationUpdateDto) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *BusinessApplicationUpdateDto) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *BusinessApplicationUpdateDto) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### SetGitRepoUrlNil

`func (o *BusinessApplicationUpdateDto) SetGitRepoUrlNil(b bool)`

 SetGitRepoUrlNil sets the value for GitRepoUrl to be an explicit nil

### UnsetGitRepoUrl
`func (o *BusinessApplicationUpdateDto) UnsetGitRepoUrl()`

UnsetGitRepoUrl ensures that no value is present for GitRepoUrl, not even an explicit nil
### GetMarkedForPublish

`func (o *BusinessApplicationUpdateDto) GetMarkedForPublish() bool`

GetMarkedForPublish returns the MarkedForPublish field if non-nil, zero value otherwise.

### GetMarkedForPublishOk

`func (o *BusinessApplicationUpdateDto) GetMarkedForPublishOk() (*bool, bool)`

GetMarkedForPublishOk returns a tuple with the MarkedForPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForPublish

`func (o *BusinessApplicationUpdateDto) SetMarkedForPublish(v bool)`

SetMarkedForPublish sets MarkedForPublish field to given value.

### HasMarkedForPublish

`func (o *BusinessApplicationUpdateDto) HasMarkedForPublish() bool`

HasMarkedForPublish returns a boolean if a field has been set.

### SetMarkedForPublishNil

`func (o *BusinessApplicationUpdateDto) SetMarkedForPublishNil(b bool)`

 SetMarkedForPublishNil sets the value for MarkedForPublish to be an explicit nil

### UnsetMarkedForPublish
`func (o *BusinessApplicationUpdateDto) UnsetMarkedForPublish()`

UnsetMarkedForPublish ensures that no value is present for MarkedForPublish, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


