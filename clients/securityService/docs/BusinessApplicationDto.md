# BusinessApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
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
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
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
**MarkedForPublish** | Pointer to **bool** |  | [optional] 

## Methods

### NewBusinessApplicationDto

`func NewBusinessApplicationDto() *BusinessApplicationDto`

NewBusinessApplicationDto instantiates a new BusinessApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessApplicationDtoWithDefaults

`func NewBusinessApplicationDtoWithDefaults() *BusinessApplicationDto`

NewBusinessApplicationDtoWithDefaults instantiates a new BusinessApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessApplicationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessApplicationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessApplicationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessApplicationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BusinessApplicationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BusinessApplicationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BusinessApplicationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessApplicationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessApplicationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessApplicationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BusinessApplicationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BusinessApplicationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *BusinessApplicationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessApplicationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessApplicationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessApplicationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BusinessApplicationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BusinessApplicationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *BusinessApplicationDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BusinessApplicationDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BusinessApplicationDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BusinessApplicationDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *BusinessApplicationDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *BusinessApplicationDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetDisplayName

`func (o *BusinessApplicationDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BusinessApplicationDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BusinessApplicationDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BusinessApplicationDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BusinessApplicationDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BusinessApplicationDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatarURL

`func (o *BusinessApplicationDto) GetAvatarURL() string`

GetAvatarURL returns the AvatarURL field if non-nil, zero value otherwise.

### GetAvatarURLOk

`func (o *BusinessApplicationDto) GetAvatarURLOk() (*string, bool)`

GetAvatarURLOk returns a tuple with the AvatarURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarURL

`func (o *BusinessApplicationDto) SetAvatarURL(v string)`

SetAvatarURL sets AvatarURL field to given value.

### HasAvatarURL

`func (o *BusinessApplicationDto) HasAvatarURL() bool`

HasAvatarURL returns a boolean if a field has been set.

### SetAvatarURLNil

`func (o *BusinessApplicationDto) SetAvatarURLNil(b bool)`

 SetAvatarURLNil sets the value for AvatarURL to be an explicit nil

### UnsetAvatarURL
`func (o *BusinessApplicationDto) UnsetAvatarURL()`

UnsetAvatarURL ensures that no value is present for AvatarURL, not even an explicit nil
### GetWebsiteUrl

`func (o *BusinessApplicationDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *BusinessApplicationDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *BusinessApplicationDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *BusinessApplicationDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *BusinessApplicationDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *BusinessApplicationDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetIsMultiTenant

`func (o *BusinessApplicationDto) GetIsMultiTenant() bool`

GetIsMultiTenant returns the IsMultiTenant field if non-nil, zero value otherwise.

### GetIsMultiTenantOk

`func (o *BusinessApplicationDto) GetIsMultiTenantOk() (*bool, bool)`

GetIsMultiTenantOk returns a tuple with the IsMultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiTenant

`func (o *BusinessApplicationDto) SetIsMultiTenant(v bool)`

SetIsMultiTenant sets IsMultiTenant field to given value.

### HasIsMultiTenant

`func (o *BusinessApplicationDto) HasIsMultiTenant() bool`

HasIsMultiTenant returns a boolean if a field has been set.

### GetIsVerified

`func (o *BusinessApplicationDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *BusinessApplicationDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *BusinessApplicationDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *BusinessApplicationDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsDisabled

`func (o *BusinessApplicationDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *BusinessApplicationDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *BusinessApplicationDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *BusinessApplicationDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsSinglePageApplication

`func (o *BusinessApplicationDto) GetIsSinglePageApplication() bool`

GetIsSinglePageApplication returns the IsSinglePageApplication field if non-nil, zero value otherwise.

### GetIsSinglePageApplicationOk

`func (o *BusinessApplicationDto) GetIsSinglePageApplicationOk() (*bool, bool)`

GetIsSinglePageApplicationOk returns a tuple with the IsSinglePageApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSinglePageApplication

`func (o *BusinessApplicationDto) SetIsSinglePageApplication(v bool)`

SetIsSinglePageApplication sets IsSinglePageApplication field to given value.

### HasIsSinglePageApplication

`func (o *BusinessApplicationDto) HasIsSinglePageApplication() bool`

HasIsSinglePageApplication returns a boolean if a field has been set.

### GetIsNativeOrDesktopApp

`func (o *BusinessApplicationDto) GetIsNativeOrDesktopApp() bool`

GetIsNativeOrDesktopApp returns the IsNativeOrDesktopApp field if non-nil, zero value otherwise.

### GetIsNativeOrDesktopAppOk

`func (o *BusinessApplicationDto) GetIsNativeOrDesktopAppOk() (*bool, bool)`

GetIsNativeOrDesktopAppOk returns a tuple with the IsNativeOrDesktopApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeOrDesktopApp

`func (o *BusinessApplicationDto) SetIsNativeOrDesktopApp(v bool)`

SetIsNativeOrDesktopApp sets IsNativeOrDesktopApp field to given value.

### HasIsNativeOrDesktopApp

`func (o *BusinessApplicationDto) HasIsNativeOrDesktopApp() bool`

HasIsNativeOrDesktopApp returns a boolean if a field has been set.

### GetContactEmail

`func (o *BusinessApplicationDto) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *BusinessApplicationDto) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *BusinessApplicationDto) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *BusinessApplicationDto) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *BusinessApplicationDto) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *BusinessApplicationDto) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetPrivacyPolicyURL

`func (o *BusinessApplicationDto) GetPrivacyPolicyURL() string`

GetPrivacyPolicyURL returns the PrivacyPolicyURL field if non-nil, zero value otherwise.

### GetPrivacyPolicyURLOk

`func (o *BusinessApplicationDto) GetPrivacyPolicyURLOk() (*string, bool)`

GetPrivacyPolicyURLOk returns a tuple with the PrivacyPolicyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyURL

`func (o *BusinessApplicationDto) SetPrivacyPolicyURL(v string)`

SetPrivacyPolicyURL sets PrivacyPolicyURL field to given value.

### HasPrivacyPolicyURL

`func (o *BusinessApplicationDto) HasPrivacyPolicyURL() bool`

HasPrivacyPolicyURL returns a boolean if a field has been set.

### SetPrivacyPolicyURLNil

`func (o *BusinessApplicationDto) SetPrivacyPolicyURLNil(b bool)`

 SetPrivacyPolicyURLNil sets the value for PrivacyPolicyURL to be an explicit nil

### UnsetPrivacyPolicyURL
`func (o *BusinessApplicationDto) UnsetPrivacyPolicyURL()`

UnsetPrivacyPolicyURL ensures that no value is present for PrivacyPolicyURL, not even an explicit nil
### GetTermsAndConditionsURL

`func (o *BusinessApplicationDto) GetTermsAndConditionsURL() string`

GetTermsAndConditionsURL returns the TermsAndConditionsURL field if non-nil, zero value otherwise.

### GetTermsAndConditionsURLOk

`func (o *BusinessApplicationDto) GetTermsAndConditionsURLOk() (*string, bool)`

GetTermsAndConditionsURLOk returns a tuple with the TermsAndConditionsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsURL

`func (o *BusinessApplicationDto) SetTermsAndConditionsURL(v string)`

SetTermsAndConditionsURL sets TermsAndConditionsURL field to given value.

### HasTermsAndConditionsURL

`func (o *BusinessApplicationDto) HasTermsAndConditionsURL() bool`

HasTermsAndConditionsURL returns a boolean if a field has been set.

### SetTermsAndConditionsURLNil

`func (o *BusinessApplicationDto) SetTermsAndConditionsURLNil(b bool)`

 SetTermsAndConditionsURLNil sets the value for TermsAndConditionsURL to be an explicit nil

### UnsetTermsAndConditionsURL
`func (o *BusinessApplicationDto) UnsetTermsAndConditionsURL()`

UnsetTermsAndConditionsURL ensures that no value is present for TermsAndConditionsURL, not even an explicit nil
### GetBusinessID

`func (o *BusinessApplicationDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *BusinessApplicationDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *BusinessApplicationDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *BusinessApplicationDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *BusinessApplicationDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *BusinessApplicationDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *BusinessApplicationDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *BusinessApplicationDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *BusinessApplicationDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *BusinessApplicationDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *BusinessApplicationDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *BusinessApplicationDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetRequireHttps

`func (o *BusinessApplicationDto) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *BusinessApplicationDto) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *BusinessApplicationDto) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *BusinessApplicationDto) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetRequireAppSecret

`func (o *BusinessApplicationDto) GetRequireAppSecret() bool`

GetRequireAppSecret returns the RequireAppSecret field if non-nil, zero value otherwise.

### GetRequireAppSecretOk

`func (o *BusinessApplicationDto) GetRequireAppSecretOk() (*bool, bool)`

GetRequireAppSecretOk returns a tuple with the RequireAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAppSecret

`func (o *BusinessApplicationDto) SetRequireAppSecret(v bool)`

SetRequireAppSecret sets RequireAppSecret field to given value.

### HasRequireAppSecret

`func (o *BusinessApplicationDto) HasRequireAppSecret() bool`

HasRequireAppSecret returns a boolean if a field has been set.

### GetEnableClientOauthLogin

`func (o *BusinessApplicationDto) GetEnableClientOauthLogin() bool`

GetEnableClientOauthLogin returns the EnableClientOauthLogin field if non-nil, zero value otherwise.

### GetEnableClientOauthLoginOk

`func (o *BusinessApplicationDto) GetEnableClientOauthLoginOk() (*bool, bool)`

GetEnableClientOauthLoginOk returns a tuple with the EnableClientOauthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientOauthLogin

`func (o *BusinessApplicationDto) SetEnableClientOauthLogin(v bool)`

SetEnableClientOauthLogin sets EnableClientOauthLogin field to given value.

### HasEnableClientOauthLogin

`func (o *BusinessApplicationDto) HasEnableClientOauthLogin() bool`

HasEnableClientOauthLogin returns a boolean if a field has been set.

### GetEnableWebOAuthLogin

`func (o *BusinessApplicationDto) GetEnableWebOAuthLogin() bool`

GetEnableWebOAuthLogin returns the EnableWebOAuthLogin field if non-nil, zero value otherwise.

### GetEnableWebOAuthLoginOk

`func (o *BusinessApplicationDto) GetEnableWebOAuthLoginOk() (*bool, bool)`

GetEnableWebOAuthLoginOk returns a tuple with the EnableWebOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebOAuthLogin

`func (o *BusinessApplicationDto) SetEnableWebOAuthLogin(v bool)`

SetEnableWebOAuthLogin sets EnableWebOAuthLogin field to given value.

### HasEnableWebOAuthLogin

`func (o *BusinessApplicationDto) HasEnableWebOAuthLogin() bool`

HasEnableWebOAuthLogin returns a boolean if a field has been set.

### GetEnableDeviceOAuthLogin

`func (o *BusinessApplicationDto) GetEnableDeviceOAuthLogin() bool`

GetEnableDeviceOAuthLogin returns the EnableDeviceOAuthLogin field if non-nil, zero value otherwise.

### GetEnableDeviceOAuthLoginOk

`func (o *BusinessApplicationDto) GetEnableDeviceOAuthLoginOk() (*bool, bool)`

GetEnableDeviceOAuthLoginOk returns a tuple with the EnableDeviceOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceOAuthLogin

`func (o *BusinessApplicationDto) SetEnableDeviceOAuthLogin(v bool)`

SetEnableDeviceOAuthLogin sets EnableDeviceOAuthLogin field to given value.

### HasEnableDeviceOAuthLogin

`func (o *BusinessApplicationDto) HasEnableDeviceOAuthLogin() bool`

HasEnableDeviceOAuthLogin returns a boolean if a field has been set.

### GetAllowAccessToSuiteSettings

`func (o *BusinessApplicationDto) GetAllowAccessToSuiteSettings() bool`

GetAllowAccessToSuiteSettings returns the AllowAccessToSuiteSettings field if non-nil, zero value otherwise.

### GetAllowAccessToSuiteSettingsOk

`func (o *BusinessApplicationDto) GetAllowAccessToSuiteSettingsOk() (*bool, bool)`

GetAllowAccessToSuiteSettingsOk returns a tuple with the AllowAccessToSuiteSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessToSuiteSettings

`func (o *BusinessApplicationDto) SetAllowAccessToSuiteSettings(v bool)`

SetAllowAccessToSuiteSettings sets AllowAccessToSuiteSettings field to given value.

### HasAllowAccessToSuiteSettings

`func (o *BusinessApplicationDto) HasAllowAccessToSuiteSettings() bool`

HasAllowAccessToSuiteSettings returns a boolean if a field has been set.

### GetRequireWebOAuthReauthentication

`func (o *BusinessApplicationDto) GetRequireWebOAuthReauthentication() bool`

GetRequireWebOAuthReauthentication returns the RequireWebOAuthReauthentication field if non-nil, zero value otherwise.

### GetRequireWebOAuthReauthenticationOk

`func (o *BusinessApplicationDto) GetRequireWebOAuthReauthenticationOk() (*bool, bool)`

GetRequireWebOAuthReauthenticationOk returns a tuple with the RequireWebOAuthReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireWebOAuthReauthentication

`func (o *BusinessApplicationDto) SetRequireWebOAuthReauthentication(v bool)`

SetRequireWebOAuthReauthentication sets RequireWebOAuthReauthentication field to given value.

### HasRequireWebOAuthReauthentication

`func (o *BusinessApplicationDto) HasRequireWebOAuthReauthentication() bool`

HasRequireWebOAuthReauthentication returns a boolean if a field has been set.

### GetRequireTwoFactorReauthorization

`func (o *BusinessApplicationDto) GetRequireTwoFactorReauthorization() bool`

GetRequireTwoFactorReauthorization returns the RequireTwoFactorReauthorization field if non-nil, zero value otherwise.

### GetRequireTwoFactorReauthorizationOk

`func (o *BusinessApplicationDto) GetRequireTwoFactorReauthorizationOk() (*bool, bool)`

GetRequireTwoFactorReauthorizationOk returns a tuple with the RequireTwoFactorReauthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTwoFactorReauthorization

`func (o *BusinessApplicationDto) SetRequireTwoFactorReauthorization(v bool)`

SetRequireTwoFactorReauthorization sets RequireTwoFactorReauthorization field to given value.

### HasRequireTwoFactorReauthorization

`func (o *BusinessApplicationDto) HasRequireTwoFactorReauthorization() bool`

HasRequireTwoFactorReauthorization returns a boolean if a field has been set.

### GetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationDto) GetEnableEmbeddedBrowserOAuthLogin() bool`

GetEnableEmbeddedBrowserOAuthLogin returns the EnableEmbeddedBrowserOAuthLogin field if non-nil, zero value otherwise.

### GetEnableEmbeddedBrowserOAuthLoginOk

`func (o *BusinessApplicationDto) GetEnableEmbeddedBrowserOAuthLoginOk() (*bool, bool)`

GetEnableEmbeddedBrowserOAuthLoginOk returns a tuple with the EnableEmbeddedBrowserOAuthLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationDto) SetEnableEmbeddedBrowserOAuthLogin(v bool)`

SetEnableEmbeddedBrowserOAuthLogin sets EnableEmbeddedBrowserOAuthLogin field to given value.

### HasEnableEmbeddedBrowserOAuthLogin

`func (o *BusinessApplicationDto) HasEnableEmbeddedBrowserOAuthLogin() bool`

HasEnableEmbeddedBrowserOAuthLogin returns a boolean if a field has been set.

### GetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationDto) GetUseStrictModeForRedirectURIs() bool`

GetUseStrictModeForRedirectURIs returns the UseStrictModeForRedirectURIs field if non-nil, zero value otherwise.

### GetUseStrictModeForRedirectURIsOk

`func (o *BusinessApplicationDto) GetUseStrictModeForRedirectURIsOk() (*bool, bool)`

GetUseStrictModeForRedirectURIsOk returns a tuple with the UseStrictModeForRedirectURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStrictModeForRedirectURIs

`func (o *BusinessApplicationDto) SetUseStrictModeForRedirectURIs(v bool)`

SetUseStrictModeForRedirectURIs sets UseStrictModeForRedirectURIs field to given value.

### HasUseStrictModeForRedirectURIs

`func (o *BusinessApplicationDto) HasUseStrictModeForRedirectURIs() bool`

HasUseStrictModeForRedirectURIs returns a boolean if a field has been set.

### GetCountryRestricted

`func (o *BusinessApplicationDto) GetCountryRestricted() bool`

GetCountryRestricted returns the CountryRestricted field if non-nil, zero value otherwise.

### GetCountryRestrictedOk

`func (o *BusinessApplicationDto) GetCountryRestrictedOk() (*bool, bool)`

GetCountryRestrictedOk returns a tuple with the CountryRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRestricted

`func (o *BusinessApplicationDto) SetCountryRestricted(v bool)`

SetCountryRestricted sets CountryRestricted field to given value.

### HasCountryRestricted

`func (o *BusinessApplicationDto) HasCountryRestricted() bool`

HasCountryRestricted returns a boolean if a field has been set.

### GetSpaUIEngine

`func (o *BusinessApplicationDto) GetSpaUIEngine() string`

GetSpaUIEngine returns the SpaUIEngine field if non-nil, zero value otherwise.

### GetSpaUIEngineOk

`func (o *BusinessApplicationDto) GetSpaUIEngineOk() (*string, bool)`

GetSpaUIEngineOk returns a tuple with the SpaUIEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUIEngine

`func (o *BusinessApplicationDto) SetSpaUIEngine(v string)`

SetSpaUIEngine sets SpaUIEngine field to given value.

### HasSpaUIEngine

`func (o *BusinessApplicationDto) HasSpaUIEngine() bool`

HasSpaUIEngine returns a boolean if a field has been set.

### SetSpaUIEngineNil

`func (o *BusinessApplicationDto) SetSpaUIEngineNil(b bool)`

 SetSpaUIEngineNil sets the value for SpaUIEngine to be an explicit nil

### UnsetSpaUIEngine
`func (o *BusinessApplicationDto) UnsetSpaUIEngine()`

UnsetSpaUIEngine ensures that no value is present for SpaUIEngine, not even an explicit nil
### GetSpaStaticFilesRootPath

`func (o *BusinessApplicationDto) GetSpaStaticFilesRootPath() string`

GetSpaStaticFilesRootPath returns the SpaStaticFilesRootPath field if non-nil, zero value otherwise.

### GetSpaStaticFilesRootPathOk

`func (o *BusinessApplicationDto) GetSpaStaticFilesRootPathOk() (*string, bool)`

GetSpaStaticFilesRootPathOk returns a tuple with the SpaStaticFilesRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaStaticFilesRootPath

`func (o *BusinessApplicationDto) SetSpaStaticFilesRootPath(v string)`

SetSpaStaticFilesRootPath sets SpaStaticFilesRootPath field to given value.

### HasSpaStaticFilesRootPath

`func (o *BusinessApplicationDto) HasSpaStaticFilesRootPath() bool`

HasSpaStaticFilesRootPath returns a boolean if a field has been set.

### SetSpaStaticFilesRootPathNil

`func (o *BusinessApplicationDto) SetSpaStaticFilesRootPathNil(b bool)`

 SetSpaStaticFilesRootPathNil sets the value for SpaStaticFilesRootPath to be an explicit nil

### UnsetSpaStaticFilesRootPath
`func (o *BusinessApplicationDto) UnsetSpaStaticFilesRootPath()`

UnsetSpaStaticFilesRootPath ensures that no value is present for SpaStaticFilesRootPath, not even an explicit nil
### GetSpaRelativeAppPath

`func (o *BusinessApplicationDto) GetSpaRelativeAppPath() string`

GetSpaRelativeAppPath returns the SpaRelativeAppPath field if non-nil, zero value otherwise.

### GetSpaRelativeAppPathOk

`func (o *BusinessApplicationDto) GetSpaRelativeAppPathOk() (*string, bool)`

GetSpaRelativeAppPathOk returns a tuple with the SpaRelativeAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeAppPath

`func (o *BusinessApplicationDto) SetSpaRelativeAppPath(v string)`

SetSpaRelativeAppPath sets SpaRelativeAppPath field to given value.

### HasSpaRelativeAppPath

`func (o *BusinessApplicationDto) HasSpaRelativeAppPath() bool`

HasSpaRelativeAppPath returns a boolean if a field has been set.

### SetSpaRelativeAppPathNil

`func (o *BusinessApplicationDto) SetSpaRelativeAppPathNil(b bool)`

 SetSpaRelativeAppPathNil sets the value for SpaRelativeAppPath to be an explicit nil

### UnsetSpaRelativeAppPath
`func (o *BusinessApplicationDto) UnsetSpaRelativeAppPath()`

UnsetSpaRelativeAppPath ensures that no value is present for SpaRelativeAppPath, not even an explicit nil
### GetSpaNpmStartScript

`func (o *BusinessApplicationDto) GetSpaNpmStartScript() string`

GetSpaNpmStartScript returns the SpaNpmStartScript field if non-nil, zero value otherwise.

### GetSpaNpmStartScriptOk

`func (o *BusinessApplicationDto) GetSpaNpmStartScriptOk() (*string, bool)`

GetSpaNpmStartScriptOk returns a tuple with the SpaNpmStartScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmStartScript

`func (o *BusinessApplicationDto) SetSpaNpmStartScript(v string)`

SetSpaNpmStartScript sets SpaNpmStartScript field to given value.

### HasSpaNpmStartScript

`func (o *BusinessApplicationDto) HasSpaNpmStartScript() bool`

HasSpaNpmStartScript returns a boolean if a field has been set.

### SetSpaNpmStartScriptNil

`func (o *BusinessApplicationDto) SetSpaNpmStartScriptNil(b bool)`

 SetSpaNpmStartScriptNil sets the value for SpaNpmStartScript to be an explicit nil

### UnsetSpaNpmStartScript
`func (o *BusinessApplicationDto) UnsetSpaNpmStartScript()`

UnsetSpaNpmStartScript ensures that no value is present for SpaNpmStartScript, not even an explicit nil
### GetSpaNpmPublishScript

`func (o *BusinessApplicationDto) GetSpaNpmPublishScript() string`

GetSpaNpmPublishScript returns the SpaNpmPublishScript field if non-nil, zero value otherwise.

### GetSpaNpmPublishScriptOk

`func (o *BusinessApplicationDto) GetSpaNpmPublishScriptOk() (*string, bool)`

GetSpaNpmPublishScriptOk returns a tuple with the SpaNpmPublishScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaNpmPublishScript

`func (o *BusinessApplicationDto) SetSpaNpmPublishScript(v string)`

SetSpaNpmPublishScript sets SpaNpmPublishScript field to given value.

### HasSpaNpmPublishScript

`func (o *BusinessApplicationDto) HasSpaNpmPublishScript() bool`

HasSpaNpmPublishScript returns a boolean if a field has been set.

### SetSpaNpmPublishScriptNil

`func (o *BusinessApplicationDto) SetSpaNpmPublishScriptNil(b bool)`

 SetSpaNpmPublishScriptNil sets the value for SpaNpmPublishScript to be an explicit nil

### UnsetSpaNpmPublishScript
`func (o *BusinessApplicationDto) UnsetSpaNpmPublishScript()`

UnsetSpaNpmPublishScript ensures that no value is present for SpaNpmPublishScript, not even an explicit nil
### GetSpaRelativeSourcePath

`func (o *BusinessApplicationDto) GetSpaRelativeSourcePath() string`

GetSpaRelativeSourcePath returns the SpaRelativeSourcePath field if non-nil, zero value otherwise.

### GetSpaRelativeSourcePathOk

`func (o *BusinessApplicationDto) GetSpaRelativeSourcePathOk() (*string, bool)`

GetSpaRelativeSourcePathOk returns a tuple with the SpaRelativeSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeSourcePath

`func (o *BusinessApplicationDto) SetSpaRelativeSourcePath(v string)`

SetSpaRelativeSourcePath sets SpaRelativeSourcePath field to given value.

### HasSpaRelativeSourcePath

`func (o *BusinessApplicationDto) HasSpaRelativeSourcePath() bool`

HasSpaRelativeSourcePath returns a boolean if a field has been set.

### SetSpaRelativeSourcePathNil

`func (o *BusinessApplicationDto) SetSpaRelativeSourcePathNil(b bool)`

 SetSpaRelativeSourcePathNil sets the value for SpaRelativeSourcePath to be an explicit nil

### UnsetSpaRelativeSourcePath
`func (o *BusinessApplicationDto) UnsetSpaRelativeSourcePath()`

UnsetSpaRelativeSourcePath ensures that no value is present for SpaRelativeSourcePath, not even an explicit nil
### GetSpaRelativeOutputPath

`func (o *BusinessApplicationDto) GetSpaRelativeOutputPath() string`

GetSpaRelativeOutputPath returns the SpaRelativeOutputPath field if non-nil, zero value otherwise.

### GetSpaRelativeOutputPathOk

`func (o *BusinessApplicationDto) GetSpaRelativeOutputPathOk() (*string, bool)`

GetSpaRelativeOutputPathOk returns a tuple with the SpaRelativeOutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaRelativeOutputPath

`func (o *BusinessApplicationDto) SetSpaRelativeOutputPath(v string)`

SetSpaRelativeOutputPath sets SpaRelativeOutputPath field to given value.

### HasSpaRelativeOutputPath

`func (o *BusinessApplicationDto) HasSpaRelativeOutputPath() bool`

HasSpaRelativeOutputPath returns a boolean if a field has been set.

### SetSpaRelativeOutputPathNil

`func (o *BusinessApplicationDto) SetSpaRelativeOutputPathNil(b bool)`

 SetSpaRelativeOutputPathNil sets the value for SpaRelativeOutputPath to be an explicit nil

### UnsetSpaRelativeOutputPath
`func (o *BusinessApplicationDto) UnsetSpaRelativeOutputPath()`

UnsetSpaRelativeOutputPath ensures that no value is present for SpaRelativeOutputPath, not even an explicit nil
### GetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationDto) GetUseProxyToSpaDevelopmentServer() bool`

GetUseProxyToSpaDevelopmentServer returns the UseProxyToSpaDevelopmentServer field if non-nil, zero value otherwise.

### GetUseProxyToSpaDevelopmentServerOk

`func (o *BusinessApplicationDto) GetUseProxyToSpaDevelopmentServerOk() (*bool, bool)`

GetUseProxyToSpaDevelopmentServerOk returns a tuple with the UseProxyToSpaDevelopmentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationDto) SetUseProxyToSpaDevelopmentServer(v bool)`

SetUseProxyToSpaDevelopmentServer sets UseProxyToSpaDevelopmentServer field to given value.

### HasUseProxyToSpaDevelopmentServer

`func (o *BusinessApplicationDto) HasUseProxyToSpaDevelopmentServer() bool`

HasUseProxyToSpaDevelopmentServer returns a boolean if a field has been set.

### GetSpaDevelopmentServerUri

`func (o *BusinessApplicationDto) GetSpaDevelopmentServerUri() string`

GetSpaDevelopmentServerUri returns the SpaDevelopmentServerUri field if non-nil, zero value otherwise.

### GetSpaDevelopmentServerUriOk

`func (o *BusinessApplicationDto) GetSpaDevelopmentServerUriOk() (*string, bool)`

GetSpaDevelopmentServerUriOk returns a tuple with the SpaDevelopmentServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaDevelopmentServerUri

`func (o *BusinessApplicationDto) SetSpaDevelopmentServerUri(v string)`

SetSpaDevelopmentServerUri sets SpaDevelopmentServerUri field to given value.

### HasSpaDevelopmentServerUri

`func (o *BusinessApplicationDto) HasSpaDevelopmentServerUri() bool`

HasSpaDevelopmentServerUri returns a boolean if a field has been set.

### SetSpaDevelopmentServerUriNil

`func (o *BusinessApplicationDto) SetSpaDevelopmentServerUriNil(b bool)`

 SetSpaDevelopmentServerUriNil sets the value for SpaDevelopmentServerUri to be an explicit nil

### UnsetSpaDevelopmentServerUri
`func (o *BusinessApplicationDto) UnsetSpaDevelopmentServerUri()`

UnsetSpaDevelopmentServerUri ensures that no value is present for SpaDevelopmentServerUri, not even an explicit nil
### GetEnableGitRepoManagement

`func (o *BusinessApplicationDto) GetEnableGitRepoManagement() bool`

GetEnableGitRepoManagement returns the EnableGitRepoManagement field if non-nil, zero value otherwise.

### GetEnableGitRepoManagementOk

`func (o *BusinessApplicationDto) GetEnableGitRepoManagementOk() (*bool, bool)`

GetEnableGitRepoManagementOk returns a tuple with the EnableGitRepoManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGitRepoManagement

`func (o *BusinessApplicationDto) SetEnableGitRepoManagement(v bool)`

SetEnableGitRepoManagement sets EnableGitRepoManagement field to given value.

### HasEnableGitRepoManagement

`func (o *BusinessApplicationDto) HasEnableGitRepoManagement() bool`

HasEnableGitRepoManagement returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *BusinessApplicationDto) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *BusinessApplicationDto) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *BusinessApplicationDto) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *BusinessApplicationDto) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### SetGitRepoUrlNil

`func (o *BusinessApplicationDto) SetGitRepoUrlNil(b bool)`

 SetGitRepoUrlNil sets the value for GitRepoUrl to be an explicit nil

### UnsetGitRepoUrl
`func (o *BusinessApplicationDto) UnsetGitRepoUrl()`

UnsetGitRepoUrl ensures that no value is present for GitRepoUrl, not even an explicit nil
### GetMarkedForPublish

`func (o *BusinessApplicationDto) GetMarkedForPublish() bool`

GetMarkedForPublish returns the MarkedForPublish field if non-nil, zero value otherwise.

### GetMarkedForPublishOk

`func (o *BusinessApplicationDto) GetMarkedForPublishOk() (*bool, bool)`

GetMarkedForPublishOk returns a tuple with the MarkedForPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForPublish

`func (o *BusinessApplicationDto) SetMarkedForPublish(v bool)`

SetMarkedForPublish sets MarkedForPublish field to given value.

### HasMarkedForPublish

`func (o *BusinessApplicationDto) HasMarkedForPublish() bool`

HasMarkedForPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


