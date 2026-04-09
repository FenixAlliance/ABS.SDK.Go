# BrandingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLang** | Pointer to **NullableString** |  | [optional] 
**PrimaryColor** | Pointer to **NullableString** |  | [optional] 
**SecondaryColor** | Pointer to **NullableString** |  | [optional] 
**HeaderLogo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**FooterLogo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**Favicons** | Pointer to [**Favicons**](Favicons.md) |  | [optional] 
**AppleIcon** | Pointer to [**AppleIcons**](AppleIcons.md) |  | [optional] 
**MsAppTile** | Pointer to [**MSAppTile**](MSAppTile.md) |  | [optional] 
**Dashboard** | Pointer to [**DashboardOptions**](DashboardOptions.md) |  | [optional] 
**Studio** | Pointer to [**StudioOptions**](StudioOptions.md) |  | [optional] 

## Methods

### NewBrandingOptions

`func NewBrandingOptions() *BrandingOptions`

NewBrandingOptions instantiates a new BrandingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingOptionsWithDefaults

`func NewBrandingOptionsWithDefaults() *BrandingOptions`

NewBrandingOptionsWithDefaults instantiates a new BrandingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLang

`func (o *BrandingOptions) GetDefaultLang() string`

GetDefaultLang returns the DefaultLang field if non-nil, zero value otherwise.

### GetDefaultLangOk

`func (o *BrandingOptions) GetDefaultLangOk() (*string, bool)`

GetDefaultLangOk returns a tuple with the DefaultLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLang

`func (o *BrandingOptions) SetDefaultLang(v string)`

SetDefaultLang sets DefaultLang field to given value.

### HasDefaultLang

`func (o *BrandingOptions) HasDefaultLang() bool`

HasDefaultLang returns a boolean if a field has been set.

### SetDefaultLangNil

`func (o *BrandingOptions) SetDefaultLangNil(b bool)`

 SetDefaultLangNil sets the value for DefaultLang to be an explicit nil

### UnsetDefaultLang
`func (o *BrandingOptions) UnsetDefaultLang()`

UnsetDefaultLang ensures that no value is present for DefaultLang, not even an explicit nil
### GetPrimaryColor

`func (o *BrandingOptions) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *BrandingOptions) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *BrandingOptions) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *BrandingOptions) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### SetPrimaryColorNil

`func (o *BrandingOptions) SetPrimaryColorNil(b bool)`

 SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil

### UnsetPrimaryColor
`func (o *BrandingOptions) UnsetPrimaryColor()`

UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
### GetSecondaryColor

`func (o *BrandingOptions) GetSecondaryColor() string`

GetSecondaryColor returns the SecondaryColor field if non-nil, zero value otherwise.

### GetSecondaryColorOk

`func (o *BrandingOptions) GetSecondaryColorOk() (*string, bool)`

GetSecondaryColorOk returns a tuple with the SecondaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColor

`func (o *BrandingOptions) SetSecondaryColor(v string)`

SetSecondaryColor sets SecondaryColor field to given value.

### HasSecondaryColor

`func (o *BrandingOptions) HasSecondaryColor() bool`

HasSecondaryColor returns a boolean if a field has been set.

### SetSecondaryColorNil

`func (o *BrandingOptions) SetSecondaryColorNil(b bool)`

 SetSecondaryColorNil sets the value for SecondaryColor to be an explicit nil

### UnsetSecondaryColor
`func (o *BrandingOptions) UnsetSecondaryColor()`

UnsetSecondaryColor ensures that no value is present for SecondaryColor, not even an explicit nil
### GetHeaderLogo

`func (o *BrandingOptions) GetHeaderLogo() Logo`

GetHeaderLogo returns the HeaderLogo field if non-nil, zero value otherwise.

### GetHeaderLogoOk

`func (o *BrandingOptions) GetHeaderLogoOk() (*Logo, bool)`

GetHeaderLogoOk returns a tuple with the HeaderLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderLogo

`func (o *BrandingOptions) SetHeaderLogo(v Logo)`

SetHeaderLogo sets HeaderLogo field to given value.

### HasHeaderLogo

`func (o *BrandingOptions) HasHeaderLogo() bool`

HasHeaderLogo returns a boolean if a field has been set.

### GetFooterLogo

`func (o *BrandingOptions) GetFooterLogo() Logo`

GetFooterLogo returns the FooterLogo field if non-nil, zero value otherwise.

### GetFooterLogoOk

`func (o *BrandingOptions) GetFooterLogoOk() (*Logo, bool)`

GetFooterLogoOk returns a tuple with the FooterLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLogo

`func (o *BrandingOptions) SetFooterLogo(v Logo)`

SetFooterLogo sets FooterLogo field to given value.

### HasFooterLogo

`func (o *BrandingOptions) HasFooterLogo() bool`

HasFooterLogo returns a boolean if a field has been set.

### GetFavicons

`func (o *BrandingOptions) GetFavicons() Favicons`

GetFavicons returns the Favicons field if non-nil, zero value otherwise.

### GetFaviconsOk

`func (o *BrandingOptions) GetFaviconsOk() (*Favicons, bool)`

GetFaviconsOk returns a tuple with the Favicons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicons

`func (o *BrandingOptions) SetFavicons(v Favicons)`

SetFavicons sets Favicons field to given value.

### HasFavicons

`func (o *BrandingOptions) HasFavicons() bool`

HasFavicons returns a boolean if a field has been set.

### GetAppleIcon

`func (o *BrandingOptions) GetAppleIcon() AppleIcons`

GetAppleIcon returns the AppleIcon field if non-nil, zero value otherwise.

### GetAppleIconOk

`func (o *BrandingOptions) GetAppleIconOk() (*AppleIcons, bool)`

GetAppleIconOk returns a tuple with the AppleIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleIcon

`func (o *BrandingOptions) SetAppleIcon(v AppleIcons)`

SetAppleIcon sets AppleIcon field to given value.

### HasAppleIcon

`func (o *BrandingOptions) HasAppleIcon() bool`

HasAppleIcon returns a boolean if a field has been set.

### GetMsAppTile

`func (o *BrandingOptions) GetMsAppTile() MSAppTile`

GetMsAppTile returns the MsAppTile field if non-nil, zero value otherwise.

### GetMsAppTileOk

`func (o *BrandingOptions) GetMsAppTileOk() (*MSAppTile, bool)`

GetMsAppTileOk returns a tuple with the MsAppTile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAppTile

`func (o *BrandingOptions) SetMsAppTile(v MSAppTile)`

SetMsAppTile sets MsAppTile field to given value.

### HasMsAppTile

`func (o *BrandingOptions) HasMsAppTile() bool`

HasMsAppTile returns a boolean if a field has been set.

### GetDashboard

`func (o *BrandingOptions) GetDashboard() DashboardOptions`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *BrandingOptions) GetDashboardOk() (*DashboardOptions, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *BrandingOptions) SetDashboard(v DashboardOptions)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *BrandingOptions) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetStudio

`func (o *BrandingOptions) GetStudio() StudioOptions`

GetStudio returns the Studio field if non-nil, zero value otherwise.

### GetStudioOk

`func (o *BrandingOptions) GetStudioOk() (*StudioOptions, bool)`

GetStudioOk returns a tuple with the Studio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudio

`func (o *BrandingOptions) SetStudio(v StudioOptions)`

SetStudio sets Studio field to given value.

### HasStudio

`func (o *BrandingOptions) HasStudio() bool`

HasStudio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


