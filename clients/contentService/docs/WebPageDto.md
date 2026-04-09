# WebPageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CodeType** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Excerpt** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HighlightImage** | Pointer to **NullableString** |  | [optional] 
**CanonicalUrl** | Pointer to **NullableString** |  | [optional] 
**EmitResult** | Pointer to **interface{}** |  | [optional] 
**SeoTitle** | Pointer to **NullableString** |  | [optional] 
**SeoKeyWords** | Pointer to **NullableString** |  | [optional] 
**SeoKeyPhrases** | Pointer to **NullableString** |  | [optional] 
**MetaDescription** | Pointer to **NullableString** |  | [optional] 
**TwitterImage** | Pointer to **NullableString** |  | [optional] 
**TwitterTitle** | Pointer to **NullableString** |  | [optional] 
**TwitterDescription** | Pointer to **NullableString** |  | [optional] 
**FacebookImage** | Pointer to **NullableString** |  | [optional] 
**FacebookTitle** | Pointer to **NullableString** |  | [optional] 
**FacebookDescription** | Pointer to **NullableString** |  | [optional] 
**FeaturedImageUrl** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**TypeName** | Pointer to **NullableString** |  | [optional] 
**GeneratedCode** | Pointer to **NullableString** |  | [optional] 
**CompilationPath** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**CSharpContent** | Pointer to **NullableString** |  | [optional] 
**RazorContent** | Pointer to **NullableString** |  | [optional] 
**CssContent** | Pointer to **NullableString** |  | [optional] 
**JsContent** | Pointer to **NullableString** |  | [optional] 
**CssFiles** | Pointer to **NullableString** |  | [optional] 
**JsFiles** | Pointer to **NullableString** |  | [optional] 
**RazorGeneratedCode** | Pointer to **NullableString** |  | [optional] 
**CSharpGeneratedCode** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**EnableComments** | Pointer to **bool** |  | [optional] 
**DisplaySocialBox** | Pointer to **bool** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**InTrashCan** | Pointer to **bool** |  | [optional] 
**SystemLocked** | Pointer to **bool** |  | [optional] 
**AllowPingBacks** | Pointer to **bool** |  | [optional] 
**AllowTrackbacks** | Pointer to **bool** |  | [optional] 
**CornerstoneContent** | Pointer to **bool** |  | [optional] 
**IsEssentialContent** | Pointer to **bool** |  | [optional] 
**AllowSearchEngineIndexing** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**WebPortalId** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentWebContentId** | Pointer to **NullableString** |  | [optional] 
**ParentWebContentVersionId** | Pointer to **NullableString** |  | [optional] 
**WebTemplateID** | Pointer to **NullableString** |  | [optional] 
**IsHomePage** | Pointer to **bool** |  | [optional] 
**IsStorePage** | Pointer to **bool** |  | [optional] 
**IsCartPage** | Pointer to **bool** |  | [optional] 
**IsBlogPage** | Pointer to **bool** |  | [optional] 
**IsAccountPage** | Pointer to **bool** |  | [optional] 
**IsCheckoutPage** | Pointer to **bool** |  | [optional] 
**IsWishListsPage** | Pointer to **bool** |  | [optional] 
**IsContactPage** | Pointer to **bool** |  | [optional] 
**IsPrivacyPolicyPage** | Pointer to **bool** |  | [optional] 
**IsTermsOfServicePage** | Pointer to **bool** |  | [optional] 

## Methods

### NewWebPageDto

`func NewWebPageDto() *WebPageDto`

NewWebPageDto instantiates a new WebPageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageDtoWithDefaults

`func NewWebPageDtoWithDefaults() *WebPageDto`

NewWebPageDtoWithDefaults instantiates a new WebPageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebPageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebPageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WebPageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WebPageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WebPageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCodeType

`func (o *WebPageDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *WebPageDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *WebPageDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *WebPageDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *WebPageDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *WebPageDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetOrder

`func (o *WebPageDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebPageDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebPageDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebPageDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSlug

`func (o *WebPageDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebPageDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebPageDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebPageDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebPageDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebPageDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *WebPageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebPageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebPageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebPageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebPageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebPageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *WebPageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExcerpt

`func (o *WebPageDto) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *WebPageDto) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *WebPageDto) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *WebPageDto) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *WebPageDto) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *WebPageDto) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetPassword

`func (o *WebPageDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebPageDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebPageDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebPageDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebPageDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebPageDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *WebPageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHighlightImage

`func (o *WebPageDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *WebPageDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *WebPageDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *WebPageDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *WebPageDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *WebPageDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetCanonicalUrl

`func (o *WebPageDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebPageDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebPageDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebPageDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebPageDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebPageDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetEmitResult

`func (o *WebPageDto) GetEmitResult() interface{}`

GetEmitResult returns the EmitResult field if non-nil, zero value otherwise.

### GetEmitResultOk

`func (o *WebPageDto) GetEmitResultOk() (*interface{}, bool)`

GetEmitResultOk returns a tuple with the EmitResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitResult

`func (o *WebPageDto) SetEmitResult(v interface{})`

SetEmitResult sets EmitResult field to given value.

### HasEmitResult

`func (o *WebPageDto) HasEmitResult() bool`

HasEmitResult returns a boolean if a field has been set.

### SetEmitResultNil

`func (o *WebPageDto) SetEmitResultNil(b bool)`

 SetEmitResultNil sets the value for EmitResult to be an explicit nil

### UnsetEmitResult
`func (o *WebPageDto) UnsetEmitResult()`

UnsetEmitResult ensures that no value is present for EmitResult, not even an explicit nil
### GetSeoTitle

`func (o *WebPageDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebPageDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebPageDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebPageDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebPageDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebPageDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoKeyWords

`func (o *WebPageDto) GetSeoKeyWords() string`

GetSeoKeyWords returns the SeoKeyWords field if non-nil, zero value otherwise.

### GetSeoKeyWordsOk

`func (o *WebPageDto) GetSeoKeyWordsOk() (*string, bool)`

GetSeoKeyWordsOk returns a tuple with the SeoKeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyWords

`func (o *WebPageDto) SetSeoKeyWords(v string)`

SetSeoKeyWords sets SeoKeyWords field to given value.

### HasSeoKeyWords

`func (o *WebPageDto) HasSeoKeyWords() bool`

HasSeoKeyWords returns a boolean if a field has been set.

### SetSeoKeyWordsNil

`func (o *WebPageDto) SetSeoKeyWordsNil(b bool)`

 SetSeoKeyWordsNil sets the value for SeoKeyWords to be an explicit nil

### UnsetSeoKeyWords
`func (o *WebPageDto) UnsetSeoKeyWords()`

UnsetSeoKeyWords ensures that no value is present for SeoKeyWords, not even an explicit nil
### GetSeoKeyPhrases

`func (o *WebPageDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebPageDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebPageDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebPageDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebPageDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebPageDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetMetaDescription

`func (o *WebPageDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebPageDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebPageDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebPageDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebPageDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebPageDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetTwitterImage

`func (o *WebPageDto) GetTwitterImage() string`

GetTwitterImage returns the TwitterImage field if non-nil, zero value otherwise.

### GetTwitterImageOk

`func (o *WebPageDto) GetTwitterImageOk() (*string, bool)`

GetTwitterImageOk returns a tuple with the TwitterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterImage

`func (o *WebPageDto) SetTwitterImage(v string)`

SetTwitterImage sets TwitterImage field to given value.

### HasTwitterImage

`func (o *WebPageDto) HasTwitterImage() bool`

HasTwitterImage returns a boolean if a field has been set.

### SetTwitterImageNil

`func (o *WebPageDto) SetTwitterImageNil(b bool)`

 SetTwitterImageNil sets the value for TwitterImage to be an explicit nil

### UnsetTwitterImage
`func (o *WebPageDto) UnsetTwitterImage()`

UnsetTwitterImage ensures that no value is present for TwitterImage, not even an explicit nil
### GetTwitterTitle

`func (o *WebPageDto) GetTwitterTitle() string`

GetTwitterTitle returns the TwitterTitle field if non-nil, zero value otherwise.

### GetTwitterTitleOk

`func (o *WebPageDto) GetTwitterTitleOk() (*string, bool)`

GetTwitterTitleOk returns a tuple with the TwitterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterTitle

`func (o *WebPageDto) SetTwitterTitle(v string)`

SetTwitterTitle sets TwitterTitle field to given value.

### HasTwitterTitle

`func (o *WebPageDto) HasTwitterTitle() bool`

HasTwitterTitle returns a boolean if a field has been set.

### SetTwitterTitleNil

`func (o *WebPageDto) SetTwitterTitleNil(b bool)`

 SetTwitterTitleNil sets the value for TwitterTitle to be an explicit nil

### UnsetTwitterTitle
`func (o *WebPageDto) UnsetTwitterTitle()`

UnsetTwitterTitle ensures that no value is present for TwitterTitle, not even an explicit nil
### GetTwitterDescription

`func (o *WebPageDto) GetTwitterDescription() string`

GetTwitterDescription returns the TwitterDescription field if non-nil, zero value otherwise.

### GetTwitterDescriptionOk

`func (o *WebPageDto) GetTwitterDescriptionOk() (*string, bool)`

GetTwitterDescriptionOk returns a tuple with the TwitterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterDescription

`func (o *WebPageDto) SetTwitterDescription(v string)`

SetTwitterDescription sets TwitterDescription field to given value.

### HasTwitterDescription

`func (o *WebPageDto) HasTwitterDescription() bool`

HasTwitterDescription returns a boolean if a field has been set.

### SetTwitterDescriptionNil

`func (o *WebPageDto) SetTwitterDescriptionNil(b bool)`

 SetTwitterDescriptionNil sets the value for TwitterDescription to be an explicit nil

### UnsetTwitterDescription
`func (o *WebPageDto) UnsetTwitterDescription()`

UnsetTwitterDescription ensures that no value is present for TwitterDescription, not even an explicit nil
### GetFacebookImage

`func (o *WebPageDto) GetFacebookImage() string`

GetFacebookImage returns the FacebookImage field if non-nil, zero value otherwise.

### GetFacebookImageOk

`func (o *WebPageDto) GetFacebookImageOk() (*string, bool)`

GetFacebookImageOk returns a tuple with the FacebookImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookImage

`func (o *WebPageDto) SetFacebookImage(v string)`

SetFacebookImage sets FacebookImage field to given value.

### HasFacebookImage

`func (o *WebPageDto) HasFacebookImage() bool`

HasFacebookImage returns a boolean if a field has been set.

### SetFacebookImageNil

`func (o *WebPageDto) SetFacebookImageNil(b bool)`

 SetFacebookImageNil sets the value for FacebookImage to be an explicit nil

### UnsetFacebookImage
`func (o *WebPageDto) UnsetFacebookImage()`

UnsetFacebookImage ensures that no value is present for FacebookImage, not even an explicit nil
### GetFacebookTitle

`func (o *WebPageDto) GetFacebookTitle() string`

GetFacebookTitle returns the FacebookTitle field if non-nil, zero value otherwise.

### GetFacebookTitleOk

`func (o *WebPageDto) GetFacebookTitleOk() (*string, bool)`

GetFacebookTitleOk returns a tuple with the FacebookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookTitle

`func (o *WebPageDto) SetFacebookTitle(v string)`

SetFacebookTitle sets FacebookTitle field to given value.

### HasFacebookTitle

`func (o *WebPageDto) HasFacebookTitle() bool`

HasFacebookTitle returns a boolean if a field has been set.

### SetFacebookTitleNil

`func (o *WebPageDto) SetFacebookTitleNil(b bool)`

 SetFacebookTitleNil sets the value for FacebookTitle to be an explicit nil

### UnsetFacebookTitle
`func (o *WebPageDto) UnsetFacebookTitle()`

UnsetFacebookTitle ensures that no value is present for FacebookTitle, not even an explicit nil
### GetFacebookDescription

`func (o *WebPageDto) GetFacebookDescription() string`

GetFacebookDescription returns the FacebookDescription field if non-nil, zero value otherwise.

### GetFacebookDescriptionOk

`func (o *WebPageDto) GetFacebookDescriptionOk() (*string, bool)`

GetFacebookDescriptionOk returns a tuple with the FacebookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookDescription

`func (o *WebPageDto) SetFacebookDescription(v string)`

SetFacebookDescription sets FacebookDescription field to given value.

### HasFacebookDescription

`func (o *WebPageDto) HasFacebookDescription() bool`

HasFacebookDescription returns a boolean if a field has been set.

### SetFacebookDescriptionNil

`func (o *WebPageDto) SetFacebookDescriptionNil(b bool)`

 SetFacebookDescriptionNil sets the value for FacebookDescription to be an explicit nil

### UnsetFacebookDescription
`func (o *WebPageDto) UnsetFacebookDescription()`

UnsetFacebookDescription ensures that no value is present for FacebookDescription, not even an explicit nil
### GetFeaturedImageUrl

`func (o *WebPageDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *WebPageDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *WebPageDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *WebPageDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *WebPageDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *WebPageDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetContent

`func (o *WebPageDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebPageDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebPageDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebPageDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebPageDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebPageDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCode

`func (o *WebPageDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebPageDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebPageDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebPageDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WebPageDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WebPageDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNamespace

`func (o *WebPageDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WebPageDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WebPageDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WebPageDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *WebPageDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *WebPageDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTypeName

`func (o *WebPageDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *WebPageDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *WebPageDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *WebPageDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *WebPageDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *WebPageDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetGeneratedCode

`func (o *WebPageDto) GetGeneratedCode() string`

GetGeneratedCode returns the GeneratedCode field if non-nil, zero value otherwise.

### GetGeneratedCodeOk

`func (o *WebPageDto) GetGeneratedCodeOk() (*string, bool)`

GetGeneratedCodeOk returns a tuple with the GeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCode

`func (o *WebPageDto) SetGeneratedCode(v string)`

SetGeneratedCode sets GeneratedCode field to given value.

### HasGeneratedCode

`func (o *WebPageDto) HasGeneratedCode() bool`

HasGeneratedCode returns a boolean if a field has been set.

### SetGeneratedCodeNil

`func (o *WebPageDto) SetGeneratedCodeNil(b bool)`

 SetGeneratedCodeNil sets the value for GeneratedCode to be an explicit nil

### UnsetGeneratedCode
`func (o *WebPageDto) UnsetGeneratedCode()`

UnsetGeneratedCode ensures that no value is present for GeneratedCode, not even an explicit nil
### GetCompilationPath

`func (o *WebPageDto) GetCompilationPath() string`

GetCompilationPath returns the CompilationPath field if non-nil, zero value otherwise.

### GetCompilationPathOk

`func (o *WebPageDto) GetCompilationPathOk() (*string, bool)`

GetCompilationPathOk returns a tuple with the CompilationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilationPath

`func (o *WebPageDto) SetCompilationPath(v string)`

SetCompilationPath sets CompilationPath field to given value.

### HasCompilationPath

`func (o *WebPageDto) HasCompilationPath() bool`

HasCompilationPath returns a boolean if a field has been set.

### SetCompilationPathNil

`func (o *WebPageDto) SetCompilationPathNil(b bool)`

 SetCompilationPathNil sets the value for CompilationPath to be an explicit nil

### UnsetCompilationPath
`func (o *WebPageDto) UnsetCompilationPath()`

UnsetCompilationPath ensures that no value is present for CompilationPath, not even an explicit nil
### GetHtmlContent

`func (o *WebPageDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebPageDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebPageDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebPageDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebPageDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebPageDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCSharpContent

`func (o *WebPageDto) GetCSharpContent() string`

GetCSharpContent returns the CSharpContent field if non-nil, zero value otherwise.

### GetCSharpContentOk

`func (o *WebPageDto) GetCSharpContentOk() (*string, bool)`

GetCSharpContentOk returns a tuple with the CSharpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpContent

`func (o *WebPageDto) SetCSharpContent(v string)`

SetCSharpContent sets CSharpContent field to given value.

### HasCSharpContent

`func (o *WebPageDto) HasCSharpContent() bool`

HasCSharpContent returns a boolean if a field has been set.

### SetCSharpContentNil

`func (o *WebPageDto) SetCSharpContentNil(b bool)`

 SetCSharpContentNil sets the value for CSharpContent to be an explicit nil

### UnsetCSharpContent
`func (o *WebPageDto) UnsetCSharpContent()`

UnsetCSharpContent ensures that no value is present for CSharpContent, not even an explicit nil
### GetRazorContent

`func (o *WebPageDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *WebPageDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *WebPageDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *WebPageDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *WebPageDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *WebPageDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetCssContent

`func (o *WebPageDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *WebPageDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *WebPageDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *WebPageDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *WebPageDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *WebPageDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *WebPageDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *WebPageDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *WebPageDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *WebPageDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *WebPageDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *WebPageDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetCssFiles

`func (o *WebPageDto) GetCssFiles() string`

GetCssFiles returns the CssFiles field if non-nil, zero value otherwise.

### GetCssFilesOk

`func (o *WebPageDto) GetCssFilesOk() (*string, bool)`

GetCssFilesOk returns a tuple with the CssFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFiles

`func (o *WebPageDto) SetCssFiles(v string)`

SetCssFiles sets CssFiles field to given value.

### HasCssFiles

`func (o *WebPageDto) HasCssFiles() bool`

HasCssFiles returns a boolean if a field has been set.

### SetCssFilesNil

`func (o *WebPageDto) SetCssFilesNil(b bool)`

 SetCssFilesNil sets the value for CssFiles to be an explicit nil

### UnsetCssFiles
`func (o *WebPageDto) UnsetCssFiles()`

UnsetCssFiles ensures that no value is present for CssFiles, not even an explicit nil
### GetJsFiles

`func (o *WebPageDto) GetJsFiles() string`

GetJsFiles returns the JsFiles field if non-nil, zero value otherwise.

### GetJsFilesOk

`func (o *WebPageDto) GetJsFilesOk() (*string, bool)`

GetJsFilesOk returns a tuple with the JsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsFiles

`func (o *WebPageDto) SetJsFiles(v string)`

SetJsFiles sets JsFiles field to given value.

### HasJsFiles

`func (o *WebPageDto) HasJsFiles() bool`

HasJsFiles returns a boolean if a field has been set.

### SetJsFilesNil

`func (o *WebPageDto) SetJsFilesNil(b bool)`

 SetJsFilesNil sets the value for JsFiles to be an explicit nil

### UnsetJsFiles
`func (o *WebPageDto) UnsetJsFiles()`

UnsetJsFiles ensures that no value is present for JsFiles, not even an explicit nil
### GetRazorGeneratedCode

`func (o *WebPageDto) GetRazorGeneratedCode() string`

GetRazorGeneratedCode returns the RazorGeneratedCode field if non-nil, zero value otherwise.

### GetRazorGeneratedCodeOk

`func (o *WebPageDto) GetRazorGeneratedCodeOk() (*string, bool)`

GetRazorGeneratedCodeOk returns a tuple with the RazorGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorGeneratedCode

`func (o *WebPageDto) SetRazorGeneratedCode(v string)`

SetRazorGeneratedCode sets RazorGeneratedCode field to given value.

### HasRazorGeneratedCode

`func (o *WebPageDto) HasRazorGeneratedCode() bool`

HasRazorGeneratedCode returns a boolean if a field has been set.

### SetRazorGeneratedCodeNil

`func (o *WebPageDto) SetRazorGeneratedCodeNil(b bool)`

 SetRazorGeneratedCodeNil sets the value for RazorGeneratedCode to be an explicit nil

### UnsetRazorGeneratedCode
`func (o *WebPageDto) UnsetRazorGeneratedCode()`

UnsetRazorGeneratedCode ensures that no value is present for RazorGeneratedCode, not even an explicit nil
### GetCSharpGeneratedCode

`func (o *WebPageDto) GetCSharpGeneratedCode() string`

GetCSharpGeneratedCode returns the CSharpGeneratedCode field if non-nil, zero value otherwise.

### GetCSharpGeneratedCodeOk

`func (o *WebPageDto) GetCSharpGeneratedCodeOk() (*string, bool)`

GetCSharpGeneratedCodeOk returns a tuple with the CSharpGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpGeneratedCode

`func (o *WebPageDto) SetCSharpGeneratedCode(v string)`

SetCSharpGeneratedCode sets CSharpGeneratedCode field to given value.

### HasCSharpGeneratedCode

`func (o *WebPageDto) HasCSharpGeneratedCode() bool`

HasCSharpGeneratedCode returns a boolean if a field has been set.

### SetCSharpGeneratedCodeNil

`func (o *WebPageDto) SetCSharpGeneratedCodeNil(b bool)`

 SetCSharpGeneratedCodeNil sets the value for CSharpGeneratedCode to be an explicit nil

### UnsetCSharpGeneratedCode
`func (o *WebPageDto) UnsetCSharpGeneratedCode()`

UnsetCSharpGeneratedCode ensures that no value is present for CSharpGeneratedCode, not even an explicit nil
### GetTemplate

`func (o *WebPageDto) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebPageDto) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebPageDto) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebPageDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefault

`func (o *WebPageDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WebPageDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WebPageDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WebPageDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnable

`func (o *WebPageDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WebPageDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WebPageDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *WebPageDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableComments

`func (o *WebPageDto) GetEnableComments() bool`

GetEnableComments returns the EnableComments field if non-nil, zero value otherwise.

### GetEnableCommentsOk

`func (o *WebPageDto) GetEnableCommentsOk() (*bool, bool)`

GetEnableCommentsOk returns a tuple with the EnableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComments

`func (o *WebPageDto) SetEnableComments(v bool)`

SetEnableComments sets EnableComments field to given value.

### HasEnableComments

`func (o *WebPageDto) HasEnableComments() bool`

HasEnableComments returns a boolean if a field has been set.

### GetDisplaySocialBox

`func (o *WebPageDto) GetDisplaySocialBox() bool`

GetDisplaySocialBox returns the DisplaySocialBox field if non-nil, zero value otherwise.

### GetDisplaySocialBoxOk

`func (o *WebPageDto) GetDisplaySocialBoxOk() (*bool, bool)`

GetDisplaySocialBoxOk returns a tuple with the DisplaySocialBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySocialBox

`func (o *WebPageDto) SetDisplaySocialBox(v bool)`

SetDisplaySocialBox sets DisplaySocialBox field to given value.

### HasDisplaySocialBox

`func (o *WebPageDto) HasDisplaySocialBox() bool`

HasDisplaySocialBox returns a boolean if a field has been set.

### GetPublished

`func (o *WebPageDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *WebPageDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *WebPageDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *WebPageDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInTrashCan

`func (o *WebPageDto) GetInTrashCan() bool`

GetInTrashCan returns the InTrashCan field if non-nil, zero value otherwise.

### GetInTrashCanOk

`func (o *WebPageDto) GetInTrashCanOk() (*bool, bool)`

GetInTrashCanOk returns a tuple with the InTrashCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrashCan

`func (o *WebPageDto) SetInTrashCan(v bool)`

SetInTrashCan sets InTrashCan field to given value.

### HasInTrashCan

`func (o *WebPageDto) HasInTrashCan() bool`

HasInTrashCan returns a boolean if a field has been set.

### GetSystemLocked

`func (o *WebPageDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *WebPageDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *WebPageDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *WebPageDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetAllowPingBacks

`func (o *WebPageDto) GetAllowPingBacks() bool`

GetAllowPingBacks returns the AllowPingBacks field if non-nil, zero value otherwise.

### GetAllowPingBacksOk

`func (o *WebPageDto) GetAllowPingBacksOk() (*bool, bool)`

GetAllowPingBacksOk returns a tuple with the AllowPingBacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPingBacks

`func (o *WebPageDto) SetAllowPingBacks(v bool)`

SetAllowPingBacks sets AllowPingBacks field to given value.

### HasAllowPingBacks

`func (o *WebPageDto) HasAllowPingBacks() bool`

HasAllowPingBacks returns a boolean if a field has been set.

### GetAllowTrackbacks

`func (o *WebPageDto) GetAllowTrackbacks() bool`

GetAllowTrackbacks returns the AllowTrackbacks field if non-nil, zero value otherwise.

### GetAllowTrackbacksOk

`func (o *WebPageDto) GetAllowTrackbacksOk() (*bool, bool)`

GetAllowTrackbacksOk returns a tuple with the AllowTrackbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrackbacks

`func (o *WebPageDto) SetAllowTrackbacks(v bool)`

SetAllowTrackbacks sets AllowTrackbacks field to given value.

### HasAllowTrackbacks

`func (o *WebPageDto) HasAllowTrackbacks() bool`

HasAllowTrackbacks returns a boolean if a field has been set.

### GetCornerstoneContent

`func (o *WebPageDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebPageDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebPageDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebPageDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetIsEssentialContent

`func (o *WebPageDto) GetIsEssentialContent() bool`

GetIsEssentialContent returns the IsEssentialContent field if non-nil, zero value otherwise.

### GetIsEssentialContentOk

`func (o *WebPageDto) GetIsEssentialContentOk() (*bool, bool)`

GetIsEssentialContentOk returns a tuple with the IsEssentialContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEssentialContent

`func (o *WebPageDto) SetIsEssentialContent(v bool)`

SetIsEssentialContent sets IsEssentialContent field to given value.

### HasIsEssentialContent

`func (o *WebPageDto) HasIsEssentialContent() bool`

HasIsEssentialContent returns a boolean if a field has been set.

### GetAllowSearchEngineIndexing

`func (o *WebPageDto) GetAllowSearchEngineIndexing() bool`

GetAllowSearchEngineIndexing returns the AllowSearchEngineIndexing field if non-nil, zero value otherwise.

### GetAllowSearchEngineIndexingOk

`func (o *WebPageDto) GetAllowSearchEngineIndexingOk() (*bool, bool)`

GetAllowSearchEngineIndexingOk returns a tuple with the AllowSearchEngineIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearchEngineIndexing

`func (o *WebPageDto) SetAllowSearchEngineIndexing(v bool)`

SetAllowSearchEngineIndexing sets AllowSearchEngineIndexing field to given value.

### HasAllowSearchEngineIndexing

`func (o *WebPageDto) HasAllowSearchEngineIndexing() bool`

HasAllowSearchEngineIndexing returns a boolean if a field has been set.

### GetTenantId

`func (o *WebPageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WebPageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WebPageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WebPageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WebPageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WebPageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWebPortalId

`func (o *WebPageDto) GetWebPortalId() string`

GetWebPortalId returns the WebPortalId field if non-nil, zero value otherwise.

### GetWebPortalIdOk

`func (o *WebPageDto) GetWebPortalIdOk() (*string, bool)`

GetWebPortalIdOk returns a tuple with the WebPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalId

`func (o *WebPageDto) SetWebPortalId(v string)`

SetWebPortalId sets WebPortalId field to given value.

### HasWebPortalId

`func (o *WebPageDto) HasWebPortalId() bool`

HasWebPortalId returns a boolean if a field has been set.

### SetWebPortalIdNil

`func (o *WebPageDto) SetWebPortalIdNil(b bool)`

 SetWebPortalIdNil sets the value for WebPortalId to be an explicit nil

### UnsetWebPortalId
`func (o *WebPageDto) UnsetWebPortalId()`

UnsetWebPortalId ensures that no value is present for WebPortalId, not even an explicit nil
### GetWebsiteThemeId

`func (o *WebPageDto) GetWebsiteThemeId() string`

GetWebsiteThemeId returns the WebsiteThemeId field if non-nil, zero value otherwise.

### GetWebsiteThemeIdOk

`func (o *WebPageDto) GetWebsiteThemeIdOk() (*string, bool)`

GetWebsiteThemeIdOk returns a tuple with the WebsiteThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeId

`func (o *WebPageDto) SetWebsiteThemeId(v string)`

SetWebsiteThemeId sets WebsiteThemeId field to given value.

### HasWebsiteThemeId

`func (o *WebPageDto) HasWebsiteThemeId() bool`

HasWebsiteThemeId returns a boolean if a field has been set.

### SetWebsiteThemeIdNil

`func (o *WebPageDto) SetWebsiteThemeIdNil(b bool)`

 SetWebsiteThemeIdNil sets the value for WebsiteThemeId to be an explicit nil

### UnsetWebsiteThemeId
`func (o *WebPageDto) UnsetWebsiteThemeId()`

UnsetWebsiteThemeId ensures that no value is present for WebsiteThemeId, not even an explicit nil
### GetEnrollmentId

`func (o *WebPageDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *WebPageDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *WebPageDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *WebPageDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *WebPageDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *WebPageDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *WebPageDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *WebPageDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *WebPageDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *WebPageDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *WebPageDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *WebPageDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetParentWebContentId

`func (o *WebPageDto) GetParentWebContentId() string`

GetParentWebContentId returns the ParentWebContentId field if non-nil, zero value otherwise.

### GetParentWebContentIdOk

`func (o *WebPageDto) GetParentWebContentIdOk() (*string, bool)`

GetParentWebContentIdOk returns a tuple with the ParentWebContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentId

`func (o *WebPageDto) SetParentWebContentId(v string)`

SetParentWebContentId sets ParentWebContentId field to given value.

### HasParentWebContentId

`func (o *WebPageDto) HasParentWebContentId() bool`

HasParentWebContentId returns a boolean if a field has been set.

### SetParentWebContentIdNil

`func (o *WebPageDto) SetParentWebContentIdNil(b bool)`

 SetParentWebContentIdNil sets the value for ParentWebContentId to be an explicit nil

### UnsetParentWebContentId
`func (o *WebPageDto) UnsetParentWebContentId()`

UnsetParentWebContentId ensures that no value is present for ParentWebContentId, not even an explicit nil
### GetParentWebContentVersionId

`func (o *WebPageDto) GetParentWebContentVersionId() string`

GetParentWebContentVersionId returns the ParentWebContentVersionId field if non-nil, zero value otherwise.

### GetParentWebContentVersionIdOk

`func (o *WebPageDto) GetParentWebContentVersionIdOk() (*string, bool)`

GetParentWebContentVersionIdOk returns a tuple with the ParentWebContentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentVersionId

`func (o *WebPageDto) SetParentWebContentVersionId(v string)`

SetParentWebContentVersionId sets ParentWebContentVersionId field to given value.

### HasParentWebContentVersionId

`func (o *WebPageDto) HasParentWebContentVersionId() bool`

HasParentWebContentVersionId returns a boolean if a field has been set.

### SetParentWebContentVersionIdNil

`func (o *WebPageDto) SetParentWebContentVersionIdNil(b bool)`

 SetParentWebContentVersionIdNil sets the value for ParentWebContentVersionId to be an explicit nil

### UnsetParentWebContentVersionId
`func (o *WebPageDto) UnsetParentWebContentVersionId()`

UnsetParentWebContentVersionId ensures that no value is present for ParentWebContentVersionId, not even an explicit nil
### GetWebTemplateID

`func (o *WebPageDto) GetWebTemplateID() string`

GetWebTemplateID returns the WebTemplateID field if non-nil, zero value otherwise.

### GetWebTemplateIDOk

`func (o *WebPageDto) GetWebTemplateIDOk() (*string, bool)`

GetWebTemplateIDOk returns a tuple with the WebTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTemplateID

`func (o *WebPageDto) SetWebTemplateID(v string)`

SetWebTemplateID sets WebTemplateID field to given value.

### HasWebTemplateID

`func (o *WebPageDto) HasWebTemplateID() bool`

HasWebTemplateID returns a boolean if a field has been set.

### SetWebTemplateIDNil

`func (o *WebPageDto) SetWebTemplateIDNil(b bool)`

 SetWebTemplateIDNil sets the value for WebTemplateID to be an explicit nil

### UnsetWebTemplateID
`func (o *WebPageDto) UnsetWebTemplateID()`

UnsetWebTemplateID ensures that no value is present for WebTemplateID, not even an explicit nil
### GetIsHomePage

`func (o *WebPageDto) GetIsHomePage() bool`

GetIsHomePage returns the IsHomePage field if non-nil, zero value otherwise.

### GetIsHomePageOk

`func (o *WebPageDto) GetIsHomePageOk() (*bool, bool)`

GetIsHomePageOk returns a tuple with the IsHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHomePage

`func (o *WebPageDto) SetIsHomePage(v bool)`

SetIsHomePage sets IsHomePage field to given value.

### HasIsHomePage

`func (o *WebPageDto) HasIsHomePage() bool`

HasIsHomePage returns a boolean if a field has been set.

### GetIsStorePage

`func (o *WebPageDto) GetIsStorePage() bool`

GetIsStorePage returns the IsStorePage field if non-nil, zero value otherwise.

### GetIsStorePageOk

`func (o *WebPageDto) GetIsStorePageOk() (*bool, bool)`

GetIsStorePageOk returns a tuple with the IsStorePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorePage

`func (o *WebPageDto) SetIsStorePage(v bool)`

SetIsStorePage sets IsStorePage field to given value.

### HasIsStorePage

`func (o *WebPageDto) HasIsStorePage() bool`

HasIsStorePage returns a boolean if a field has been set.

### GetIsCartPage

`func (o *WebPageDto) GetIsCartPage() bool`

GetIsCartPage returns the IsCartPage field if non-nil, zero value otherwise.

### GetIsCartPageOk

`func (o *WebPageDto) GetIsCartPageOk() (*bool, bool)`

GetIsCartPageOk returns a tuple with the IsCartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCartPage

`func (o *WebPageDto) SetIsCartPage(v bool)`

SetIsCartPage sets IsCartPage field to given value.

### HasIsCartPage

`func (o *WebPageDto) HasIsCartPage() bool`

HasIsCartPage returns a boolean if a field has been set.

### GetIsBlogPage

`func (o *WebPageDto) GetIsBlogPage() bool`

GetIsBlogPage returns the IsBlogPage field if non-nil, zero value otherwise.

### GetIsBlogPageOk

`func (o *WebPageDto) GetIsBlogPageOk() (*bool, bool)`

GetIsBlogPageOk returns a tuple with the IsBlogPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlogPage

`func (o *WebPageDto) SetIsBlogPage(v bool)`

SetIsBlogPage sets IsBlogPage field to given value.

### HasIsBlogPage

`func (o *WebPageDto) HasIsBlogPage() bool`

HasIsBlogPage returns a boolean if a field has been set.

### GetIsAccountPage

`func (o *WebPageDto) GetIsAccountPage() bool`

GetIsAccountPage returns the IsAccountPage field if non-nil, zero value otherwise.

### GetIsAccountPageOk

`func (o *WebPageDto) GetIsAccountPageOk() (*bool, bool)`

GetIsAccountPageOk returns a tuple with the IsAccountPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPage

`func (o *WebPageDto) SetIsAccountPage(v bool)`

SetIsAccountPage sets IsAccountPage field to given value.

### HasIsAccountPage

`func (o *WebPageDto) HasIsAccountPage() bool`

HasIsAccountPage returns a boolean if a field has been set.

### GetIsCheckoutPage

`func (o *WebPageDto) GetIsCheckoutPage() bool`

GetIsCheckoutPage returns the IsCheckoutPage field if non-nil, zero value otherwise.

### GetIsCheckoutPageOk

`func (o *WebPageDto) GetIsCheckoutPageOk() (*bool, bool)`

GetIsCheckoutPageOk returns a tuple with the IsCheckoutPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCheckoutPage

`func (o *WebPageDto) SetIsCheckoutPage(v bool)`

SetIsCheckoutPage sets IsCheckoutPage field to given value.

### HasIsCheckoutPage

`func (o *WebPageDto) HasIsCheckoutPage() bool`

HasIsCheckoutPage returns a boolean if a field has been set.

### GetIsWishListsPage

`func (o *WebPageDto) GetIsWishListsPage() bool`

GetIsWishListsPage returns the IsWishListsPage field if non-nil, zero value otherwise.

### GetIsWishListsPageOk

`func (o *WebPageDto) GetIsWishListsPageOk() (*bool, bool)`

GetIsWishListsPageOk returns a tuple with the IsWishListsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWishListsPage

`func (o *WebPageDto) SetIsWishListsPage(v bool)`

SetIsWishListsPage sets IsWishListsPage field to given value.

### HasIsWishListsPage

`func (o *WebPageDto) HasIsWishListsPage() bool`

HasIsWishListsPage returns a boolean if a field has been set.

### GetIsContactPage

`func (o *WebPageDto) GetIsContactPage() bool`

GetIsContactPage returns the IsContactPage field if non-nil, zero value otherwise.

### GetIsContactPageOk

`func (o *WebPageDto) GetIsContactPageOk() (*bool, bool)`

GetIsContactPageOk returns a tuple with the IsContactPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContactPage

`func (o *WebPageDto) SetIsContactPage(v bool)`

SetIsContactPage sets IsContactPage field to given value.

### HasIsContactPage

`func (o *WebPageDto) HasIsContactPage() bool`

HasIsContactPage returns a boolean if a field has been set.

### GetIsPrivacyPolicyPage

`func (o *WebPageDto) GetIsPrivacyPolicyPage() bool`

GetIsPrivacyPolicyPage returns the IsPrivacyPolicyPage field if non-nil, zero value otherwise.

### GetIsPrivacyPolicyPageOk

`func (o *WebPageDto) GetIsPrivacyPolicyPageOk() (*bool, bool)`

GetIsPrivacyPolicyPageOk returns a tuple with the IsPrivacyPolicyPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivacyPolicyPage

`func (o *WebPageDto) SetIsPrivacyPolicyPage(v bool)`

SetIsPrivacyPolicyPage sets IsPrivacyPolicyPage field to given value.

### HasIsPrivacyPolicyPage

`func (o *WebPageDto) HasIsPrivacyPolicyPage() bool`

HasIsPrivacyPolicyPage returns a boolean if a field has been set.

### GetIsTermsOfServicePage

`func (o *WebPageDto) GetIsTermsOfServicePage() bool`

GetIsTermsOfServicePage returns the IsTermsOfServicePage field if non-nil, zero value otherwise.

### GetIsTermsOfServicePageOk

`func (o *WebPageDto) GetIsTermsOfServicePageOk() (*bool, bool)`

GetIsTermsOfServicePageOk returns a tuple with the IsTermsOfServicePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTermsOfServicePage

`func (o *WebPageDto) SetIsTermsOfServicePage(v bool)`

SetIsTermsOfServicePage sets IsTermsOfServicePage field to given value.

### HasIsTermsOfServicePage

`func (o *WebPageDto) HasIsTermsOfServicePage() bool`

HasIsTermsOfServicePage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


