# PortalOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PortalUiEngine** | Pointer to **string** |  | [optional] 
**Seo** | Pointer to [**SeoOptions**](SeoOptions.md) |  | [optional] 
**Store** | Pointer to [**StoreOptions**](StoreOptions.md) |  | [optional] 
**Theming** | Pointer to [**ThemingOptions**](ThemingOptions.md) |  | [optional] 
**Branding** | Pointer to [**BrandingOptions**](BrandingOptions.md) |  | [optional] 
**Social** | Pointer to [**SocialMediaOptions**](SocialMediaOptions.md) |  | [optional] 
**Privacy** | Pointer to [**PrivacyOptions**](PrivacyOptions.md) |  | [optional] 
**Blog** | Pointer to [**BlogOptions**](BlogOptions.md) |  | [optional] 
**Forums** | Pointer to [**ForumOptions**](ForumOptions.md) |  | [optional] 
**Footer** | Pointer to [**FooterOptions**](FooterOptions.md) |  | [optional] 
**Background** | Pointer to [**BackgroundOptions**](BackgroundOptions.md) |  | [optional] 
**Breadcrumbs** | Pointer to [**BreadcrumbsOptions**](BreadcrumbsOptions.md) |  | [optional] 
**Contact** | Pointer to [**ContactOptions**](ContactOptions.md) |  | [optional] 
**Color** | Pointer to [**ColorOptions**](ColorOptions.md) |  | [optional] 
**Dashboard** | Pointer to [**DashboardOptions**](DashboardOptions.md) |  | [optional] 
**Header** | Pointer to [**HeaderOptions**](HeaderOptions.md) |  | [optional] 
**TitleBar** | Pointer to [**TitleBarOptions**](TitleBarOptions.md) |  | [optional] 
**Typography** | Pointer to [**TypographyOptions**](TypographyOptions.md) |  | [optional] 
**SocialMedia** | Pointer to [**SocialMediaOptions**](SocialMediaOptions.md) |  | [optional] 
**SlidingBar** | Pointer to [**SlidingBarOptions**](SlidingBarOptions.md) |  | [optional] 
**Slideshow** | Pointer to **map[string]interface{}** |  | [optional] 
**Slider** | Pointer to **map[string]interface{}** |  | [optional] 
**Sidebar** | Pointer to **map[string]interface{}** |  | [optional] 
**Search** | Pointer to **map[string]interface{}** |  | [optional] 
**Responsive** | Pointer to [**ResponsiveOptions**](ResponsiveOptions.md) |  | [optional] 
**Portfolio** | Pointer to **map[string]interface{}** |  | [optional] 
**Performance** | Pointer to **map[string]interface{}** |  | [optional] 
**Pagination** | Pointer to **map[string]interface{}** |  | [optional] 
**Miscellaneous** | Pointer to **map[string]interface{}** |  | [optional] 
**Menu** | Pointer to [**MenuOptions**](MenuOptions.md) |  | [optional] 
**Grid** | Pointer to **map[string]interface{}** |  | [optional] 
**Mansory** | Pointer to **map[string]interface{}** |  | [optional] 
**Lightbox** | Pointer to **map[string]interface{}** |  | [optional] 
**Layout** | Pointer to [**LayoutOptions**](LayoutOptions.md) |  | [optional] 
**CodeFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**Forms** | Pointer to **map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to [**IntegrationsOptions**](IntegrationsOptions.md) |  | [optional] 
**Emails** | Pointer to [**EmailsOptions**](EmailsOptions.md) |  | [optional] 

## Methods

### NewPortalOptions

`func NewPortalOptions() *PortalOptions`

NewPortalOptions instantiates a new PortalOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalOptionsWithDefaults

`func NewPortalOptionsWithDefaults() *PortalOptions`

NewPortalOptionsWithDefaults instantiates a new PortalOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PortalOptions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PortalOptions) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PortalOptions) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PortalOptions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PortalOptions) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PortalOptions) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PortalOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortalOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortalOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortalOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PortalOptions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PortalOptions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPortalUiEngine

`func (o *PortalOptions) GetPortalUiEngine() string`

GetPortalUiEngine returns the PortalUiEngine field if non-nil, zero value otherwise.

### GetPortalUiEngineOk

`func (o *PortalOptions) GetPortalUiEngineOk() (*string, bool)`

GetPortalUiEngineOk returns a tuple with the PortalUiEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUiEngine

`func (o *PortalOptions) SetPortalUiEngine(v string)`

SetPortalUiEngine sets PortalUiEngine field to given value.

### HasPortalUiEngine

`func (o *PortalOptions) HasPortalUiEngine() bool`

HasPortalUiEngine returns a boolean if a field has been set.

### GetSeo

`func (o *PortalOptions) GetSeo() SeoOptions`

GetSeo returns the Seo field if non-nil, zero value otherwise.

### GetSeoOk

`func (o *PortalOptions) GetSeoOk() (*SeoOptions, bool)`

GetSeoOk returns a tuple with the Seo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeo

`func (o *PortalOptions) SetSeo(v SeoOptions)`

SetSeo sets Seo field to given value.

### HasSeo

`func (o *PortalOptions) HasSeo() bool`

HasSeo returns a boolean if a field has been set.

### GetStore

`func (o *PortalOptions) GetStore() StoreOptions`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PortalOptions) GetStoreOk() (*StoreOptions, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PortalOptions) SetStore(v StoreOptions)`

SetStore sets Store field to given value.

### HasStore

`func (o *PortalOptions) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTheming

`func (o *PortalOptions) GetTheming() ThemingOptions`

GetTheming returns the Theming field if non-nil, zero value otherwise.

### GetThemingOk

`func (o *PortalOptions) GetThemingOk() (*ThemingOptions, bool)`

GetThemingOk returns a tuple with the Theming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheming

`func (o *PortalOptions) SetTheming(v ThemingOptions)`

SetTheming sets Theming field to given value.

### HasTheming

`func (o *PortalOptions) HasTheming() bool`

HasTheming returns a boolean if a field has been set.

### GetBranding

`func (o *PortalOptions) GetBranding() BrandingOptions`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *PortalOptions) GetBrandingOk() (*BrandingOptions, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *PortalOptions) SetBranding(v BrandingOptions)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *PortalOptions) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetSocial

`func (o *PortalOptions) GetSocial() SocialMediaOptions`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *PortalOptions) GetSocialOk() (*SocialMediaOptions, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *PortalOptions) SetSocial(v SocialMediaOptions)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *PortalOptions) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetPrivacy

`func (o *PortalOptions) GetPrivacy() PrivacyOptions`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *PortalOptions) GetPrivacyOk() (*PrivacyOptions, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *PortalOptions) SetPrivacy(v PrivacyOptions)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *PortalOptions) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetBlog

`func (o *PortalOptions) GetBlog() BlogOptions`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *PortalOptions) GetBlogOk() (*BlogOptions, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *PortalOptions) SetBlog(v BlogOptions)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *PortalOptions) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetForums

`func (o *PortalOptions) GetForums() ForumOptions`

GetForums returns the Forums field if non-nil, zero value otherwise.

### GetForumsOk

`func (o *PortalOptions) GetForumsOk() (*ForumOptions, bool)`

GetForumsOk returns a tuple with the Forums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForums

`func (o *PortalOptions) SetForums(v ForumOptions)`

SetForums sets Forums field to given value.

### HasForums

`func (o *PortalOptions) HasForums() bool`

HasForums returns a boolean if a field has been set.

### GetFooter

`func (o *PortalOptions) GetFooter() FooterOptions`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *PortalOptions) GetFooterOk() (*FooterOptions, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *PortalOptions) SetFooter(v FooterOptions)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *PortalOptions) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetBackground

`func (o *PortalOptions) GetBackground() BackgroundOptions`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PortalOptions) GetBackgroundOk() (*BackgroundOptions, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PortalOptions) SetBackground(v BackgroundOptions)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *PortalOptions) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBreadcrumbs

`func (o *PortalOptions) GetBreadcrumbs() BreadcrumbsOptions`

GetBreadcrumbs returns the Breadcrumbs field if non-nil, zero value otherwise.

### GetBreadcrumbsOk

`func (o *PortalOptions) GetBreadcrumbsOk() (*BreadcrumbsOptions, bool)`

GetBreadcrumbsOk returns a tuple with the Breadcrumbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreadcrumbs

`func (o *PortalOptions) SetBreadcrumbs(v BreadcrumbsOptions)`

SetBreadcrumbs sets Breadcrumbs field to given value.

### HasBreadcrumbs

`func (o *PortalOptions) HasBreadcrumbs() bool`

HasBreadcrumbs returns a boolean if a field has been set.

### GetContact

`func (o *PortalOptions) GetContact() ContactOptions`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PortalOptions) GetContactOk() (*ContactOptions, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PortalOptions) SetContact(v ContactOptions)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PortalOptions) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetColor

`func (o *PortalOptions) GetColor() ColorOptions`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PortalOptions) GetColorOk() (*ColorOptions, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PortalOptions) SetColor(v ColorOptions)`

SetColor sets Color field to given value.

### HasColor

`func (o *PortalOptions) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDashboard

`func (o *PortalOptions) GetDashboard() DashboardOptions`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *PortalOptions) GetDashboardOk() (*DashboardOptions, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *PortalOptions) SetDashboard(v DashboardOptions)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *PortalOptions) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetHeader

`func (o *PortalOptions) GetHeader() HeaderOptions`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *PortalOptions) GetHeaderOk() (*HeaderOptions, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *PortalOptions) SetHeader(v HeaderOptions)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *PortalOptions) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTitleBar

`func (o *PortalOptions) GetTitleBar() TitleBarOptions`

GetTitleBar returns the TitleBar field if non-nil, zero value otherwise.

### GetTitleBarOk

`func (o *PortalOptions) GetTitleBarOk() (*TitleBarOptions, bool)`

GetTitleBarOk returns a tuple with the TitleBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleBar

`func (o *PortalOptions) SetTitleBar(v TitleBarOptions)`

SetTitleBar sets TitleBar field to given value.

### HasTitleBar

`func (o *PortalOptions) HasTitleBar() bool`

HasTitleBar returns a boolean if a field has been set.

### GetTypography

`func (o *PortalOptions) GetTypography() TypographyOptions`

GetTypography returns the Typography field if non-nil, zero value otherwise.

### GetTypographyOk

`func (o *PortalOptions) GetTypographyOk() (*TypographyOptions, bool)`

GetTypographyOk returns a tuple with the Typography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypography

`func (o *PortalOptions) SetTypography(v TypographyOptions)`

SetTypography sets Typography field to given value.

### HasTypography

`func (o *PortalOptions) HasTypography() bool`

HasTypography returns a boolean if a field has been set.

### GetSocialMedia

`func (o *PortalOptions) GetSocialMedia() SocialMediaOptions`

GetSocialMedia returns the SocialMedia field if non-nil, zero value otherwise.

### GetSocialMediaOk

`func (o *PortalOptions) GetSocialMediaOk() (*SocialMediaOptions, bool)`

GetSocialMediaOk returns a tuple with the SocialMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialMedia

`func (o *PortalOptions) SetSocialMedia(v SocialMediaOptions)`

SetSocialMedia sets SocialMedia field to given value.

### HasSocialMedia

`func (o *PortalOptions) HasSocialMedia() bool`

HasSocialMedia returns a boolean if a field has been set.

### GetSlidingBar

`func (o *PortalOptions) GetSlidingBar() SlidingBarOptions`

GetSlidingBar returns the SlidingBar field if non-nil, zero value otherwise.

### GetSlidingBarOk

`func (o *PortalOptions) GetSlidingBarOk() (*SlidingBarOptions, bool)`

GetSlidingBarOk returns a tuple with the SlidingBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlidingBar

`func (o *PortalOptions) SetSlidingBar(v SlidingBarOptions)`

SetSlidingBar sets SlidingBar field to given value.

### HasSlidingBar

`func (o *PortalOptions) HasSlidingBar() bool`

HasSlidingBar returns a boolean if a field has been set.

### GetSlideshow

`func (o *PortalOptions) GetSlideshow() map[string]interface{}`

GetSlideshow returns the Slideshow field if non-nil, zero value otherwise.

### GetSlideshowOk

`func (o *PortalOptions) GetSlideshowOk() (*map[string]interface{}, bool)`

GetSlideshowOk returns a tuple with the Slideshow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlideshow

`func (o *PortalOptions) SetSlideshow(v map[string]interface{})`

SetSlideshow sets Slideshow field to given value.

### HasSlideshow

`func (o *PortalOptions) HasSlideshow() bool`

HasSlideshow returns a boolean if a field has been set.

### GetSlider

`func (o *PortalOptions) GetSlider() map[string]interface{}`

GetSlider returns the Slider field if non-nil, zero value otherwise.

### GetSliderOk

`func (o *PortalOptions) GetSliderOk() (*map[string]interface{}, bool)`

GetSliderOk returns a tuple with the Slider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlider

`func (o *PortalOptions) SetSlider(v map[string]interface{})`

SetSlider sets Slider field to given value.

### HasSlider

`func (o *PortalOptions) HasSlider() bool`

HasSlider returns a boolean if a field has been set.

### GetSidebar

`func (o *PortalOptions) GetSidebar() map[string]interface{}`

GetSidebar returns the Sidebar field if non-nil, zero value otherwise.

### GetSidebarOk

`func (o *PortalOptions) GetSidebarOk() (*map[string]interface{}, bool)`

GetSidebarOk returns a tuple with the Sidebar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidebar

`func (o *PortalOptions) SetSidebar(v map[string]interface{})`

SetSidebar sets Sidebar field to given value.

### HasSidebar

`func (o *PortalOptions) HasSidebar() bool`

HasSidebar returns a boolean if a field has been set.

### GetSearch

`func (o *PortalOptions) GetSearch() map[string]interface{}`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *PortalOptions) GetSearchOk() (*map[string]interface{}, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *PortalOptions) SetSearch(v map[string]interface{})`

SetSearch sets Search field to given value.

### HasSearch

`func (o *PortalOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetResponsive

`func (o *PortalOptions) GetResponsive() ResponsiveOptions`

GetResponsive returns the Responsive field if non-nil, zero value otherwise.

### GetResponsiveOk

`func (o *PortalOptions) GetResponsiveOk() (*ResponsiveOptions, bool)`

GetResponsiveOk returns a tuple with the Responsive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsive

`func (o *PortalOptions) SetResponsive(v ResponsiveOptions)`

SetResponsive sets Responsive field to given value.

### HasResponsive

`func (o *PortalOptions) HasResponsive() bool`

HasResponsive returns a boolean if a field has been set.

### GetPortfolio

`func (o *PortalOptions) GetPortfolio() map[string]interface{}`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *PortalOptions) GetPortfolioOk() (*map[string]interface{}, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *PortalOptions) SetPortfolio(v map[string]interface{})`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *PortalOptions) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### GetPerformance

`func (o *PortalOptions) GetPerformance() map[string]interface{}`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *PortalOptions) GetPerformanceOk() (*map[string]interface{}, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *PortalOptions) SetPerformance(v map[string]interface{})`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *PortalOptions) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetPagination

`func (o *PortalOptions) GetPagination() map[string]interface{}`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PortalOptions) GetPaginationOk() (*map[string]interface{}, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PortalOptions) SetPagination(v map[string]interface{})`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PortalOptions) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetMiscellaneous

`func (o *PortalOptions) GetMiscellaneous() map[string]interface{}`

GetMiscellaneous returns the Miscellaneous field if non-nil, zero value otherwise.

### GetMiscellaneousOk

`func (o *PortalOptions) GetMiscellaneousOk() (*map[string]interface{}, bool)`

GetMiscellaneousOk returns a tuple with the Miscellaneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscellaneous

`func (o *PortalOptions) SetMiscellaneous(v map[string]interface{})`

SetMiscellaneous sets Miscellaneous field to given value.

### HasMiscellaneous

`func (o *PortalOptions) HasMiscellaneous() bool`

HasMiscellaneous returns a boolean if a field has been set.

### GetMenu

`func (o *PortalOptions) GetMenu() MenuOptions`

GetMenu returns the Menu field if non-nil, zero value otherwise.

### GetMenuOk

`func (o *PortalOptions) GetMenuOk() (*MenuOptions, bool)`

GetMenuOk returns a tuple with the Menu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenu

`func (o *PortalOptions) SetMenu(v MenuOptions)`

SetMenu sets Menu field to given value.

### HasMenu

`func (o *PortalOptions) HasMenu() bool`

HasMenu returns a boolean if a field has been set.

### GetGrid

`func (o *PortalOptions) GetGrid() map[string]interface{}`

GetGrid returns the Grid field if non-nil, zero value otherwise.

### GetGridOk

`func (o *PortalOptions) GetGridOk() (*map[string]interface{}, bool)`

GetGridOk returns a tuple with the Grid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrid

`func (o *PortalOptions) SetGrid(v map[string]interface{})`

SetGrid sets Grid field to given value.

### HasGrid

`func (o *PortalOptions) HasGrid() bool`

HasGrid returns a boolean if a field has been set.

### GetMansory

`func (o *PortalOptions) GetMansory() map[string]interface{}`

GetMansory returns the Mansory field if non-nil, zero value otherwise.

### GetMansoryOk

`func (o *PortalOptions) GetMansoryOk() (*map[string]interface{}, bool)`

GetMansoryOk returns a tuple with the Mansory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMansory

`func (o *PortalOptions) SetMansory(v map[string]interface{})`

SetMansory sets Mansory field to given value.

### HasMansory

`func (o *PortalOptions) HasMansory() bool`

HasMansory returns a boolean if a field has been set.

### GetLightbox

`func (o *PortalOptions) GetLightbox() map[string]interface{}`

GetLightbox returns the Lightbox field if non-nil, zero value otherwise.

### GetLightboxOk

`func (o *PortalOptions) GetLightboxOk() (*map[string]interface{}, bool)`

GetLightboxOk returns a tuple with the Lightbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightbox

`func (o *PortalOptions) SetLightbox(v map[string]interface{})`

SetLightbox sets Lightbox field to given value.

### HasLightbox

`func (o *PortalOptions) HasLightbox() bool`

HasLightbox returns a boolean if a field has been set.

### GetLayout

`func (o *PortalOptions) GetLayout() LayoutOptions`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *PortalOptions) GetLayoutOk() (*LayoutOptions, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *PortalOptions) SetLayout(v LayoutOptions)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *PortalOptions) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetCodeFields

`func (o *PortalOptions) GetCodeFields() map[string]interface{}`

GetCodeFields returns the CodeFields field if non-nil, zero value otherwise.

### GetCodeFieldsOk

`func (o *PortalOptions) GetCodeFieldsOk() (*map[string]interface{}, bool)`

GetCodeFieldsOk returns a tuple with the CodeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeFields

`func (o *PortalOptions) SetCodeFields(v map[string]interface{})`

SetCodeFields sets CodeFields field to given value.

### HasCodeFields

`func (o *PortalOptions) HasCodeFields() bool`

HasCodeFields returns a boolean if a field has been set.

### GetFeatures

`func (o *PortalOptions) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PortalOptions) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PortalOptions) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PortalOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetForms

`func (o *PortalOptions) GetForms() map[string]interface{}`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *PortalOptions) GetFormsOk() (*map[string]interface{}, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *PortalOptions) SetForms(v map[string]interface{})`

SetForms sets Forms field to given value.

### HasForms

`func (o *PortalOptions) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetIntegrations

`func (o *PortalOptions) GetIntegrations() IntegrationsOptions`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *PortalOptions) GetIntegrationsOk() (*IntegrationsOptions, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *PortalOptions) SetIntegrations(v IntegrationsOptions)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *PortalOptions) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetEmails

`func (o *PortalOptions) GetEmails() EmailsOptions`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *PortalOptions) GetEmailsOk() (*EmailsOptions, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *PortalOptions) SetEmails(v EmailsOptions)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *PortalOptions) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


