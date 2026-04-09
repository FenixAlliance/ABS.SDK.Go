# WebPageUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Excerpt** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HighlightImage** | Pointer to **NullableString** |  | [optional] 
**CanonicalUrl** | Pointer to **NullableString** |  | [optional] 
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
**CodeType** | Pointer to **NullableString** |  | [optional] 
**CSharpContent** | Pointer to **NullableString** |  | [optional] 
**RazorContent** | Pointer to **NullableString** |  | [optional] 
**CssContent** | Pointer to **NullableString** |  | [optional] 
**JsContent** | Pointer to **NullableString** |  | [optional] 
**CssFiles** | Pointer to **NullableString** |  | [optional] 
**JsFiles** | Pointer to **NullableString** |  | [optional] 
**RazorGeneratedCode** | Pointer to **NullableString** |  | [optional] 
**CSharpGeneratedCode** | Pointer to **NullableString** |  | [optional] 
**PrecompiledLogicSize** | Pointer to **int32** |  | [optional] 
**PrecompiledLogicSizeLong** | Pointer to **int64** |  | [optional] 
**PrecompiledViewSize** | Pointer to **int32** |  | [optional] 
**PrecompiledViewSizeLong** | Pointer to **int64** |  | [optional] 
**PrecompiledLogicViewSize** | Pointer to **int32** |  | [optional] 
**Template** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**EnableComments** | Pointer to **bool** |  | [optional] 
**DisplaySocialBox** | Pointer to **bool** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**InTrashCan** | Pointer to **bool** |  | [optional] 
**SystemLocked** | Pointer to **bool** |  | [optional] 
**AllowPingbacks** | Pointer to **bool** |  | [optional] 
**AllowTrackbacks** | Pointer to **bool** |  | [optional] 
**CornerstoneContent** | Pointer to **bool** |  | [optional] 
**IsEssentialContent** | Pointer to **bool** |  | [optional] 
**AllowSearchEngineIndexing** | Pointer to **bool** |  | [optional] 
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

### NewWebPageUpdateDto

`func NewWebPageUpdateDto() *WebPageUpdateDto`

NewWebPageUpdateDto instantiates a new WebPageUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageUpdateDtoWithDefaults

`func NewWebPageUpdateDtoWithDefaults() *WebPageUpdateDto`

NewWebPageUpdateDtoWithDefaults instantiates a new WebPageUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *WebPageUpdateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebPageUpdateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebPageUpdateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebPageUpdateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSlug

`func (o *WebPageUpdateDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebPageUpdateDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebPageUpdateDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebPageUpdateDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebPageUpdateDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebPageUpdateDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *WebPageUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebPageUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebPageUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebPageUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebPageUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebPageUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *WebPageUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPageUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPageUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPageUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPageUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPageUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExcerpt

`func (o *WebPageUpdateDto) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *WebPageUpdateDto) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *WebPageUpdateDto) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *WebPageUpdateDto) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *WebPageUpdateDto) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *WebPageUpdateDto) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetPassword

`func (o *WebPageUpdateDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebPageUpdateDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebPageUpdateDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebPageUpdateDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebPageUpdateDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebPageUpdateDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *WebPageUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPageUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPageUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPageUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPageUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPageUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHighlightImage

`func (o *WebPageUpdateDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *WebPageUpdateDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *WebPageUpdateDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *WebPageUpdateDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *WebPageUpdateDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *WebPageUpdateDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetCanonicalUrl

`func (o *WebPageUpdateDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebPageUpdateDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebPageUpdateDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebPageUpdateDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebPageUpdateDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebPageUpdateDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetSeoTitle

`func (o *WebPageUpdateDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebPageUpdateDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebPageUpdateDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebPageUpdateDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebPageUpdateDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebPageUpdateDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoKeyWords

`func (o *WebPageUpdateDto) GetSeoKeyWords() string`

GetSeoKeyWords returns the SeoKeyWords field if non-nil, zero value otherwise.

### GetSeoKeyWordsOk

`func (o *WebPageUpdateDto) GetSeoKeyWordsOk() (*string, bool)`

GetSeoKeyWordsOk returns a tuple with the SeoKeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyWords

`func (o *WebPageUpdateDto) SetSeoKeyWords(v string)`

SetSeoKeyWords sets SeoKeyWords field to given value.

### HasSeoKeyWords

`func (o *WebPageUpdateDto) HasSeoKeyWords() bool`

HasSeoKeyWords returns a boolean if a field has been set.

### SetSeoKeyWordsNil

`func (o *WebPageUpdateDto) SetSeoKeyWordsNil(b bool)`

 SetSeoKeyWordsNil sets the value for SeoKeyWords to be an explicit nil

### UnsetSeoKeyWords
`func (o *WebPageUpdateDto) UnsetSeoKeyWords()`

UnsetSeoKeyWords ensures that no value is present for SeoKeyWords, not even an explicit nil
### GetSeoKeyPhrases

`func (o *WebPageUpdateDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebPageUpdateDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebPageUpdateDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebPageUpdateDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebPageUpdateDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebPageUpdateDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetMetaDescription

`func (o *WebPageUpdateDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebPageUpdateDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebPageUpdateDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebPageUpdateDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebPageUpdateDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebPageUpdateDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetTwitterImage

`func (o *WebPageUpdateDto) GetTwitterImage() string`

GetTwitterImage returns the TwitterImage field if non-nil, zero value otherwise.

### GetTwitterImageOk

`func (o *WebPageUpdateDto) GetTwitterImageOk() (*string, bool)`

GetTwitterImageOk returns a tuple with the TwitterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterImage

`func (o *WebPageUpdateDto) SetTwitterImage(v string)`

SetTwitterImage sets TwitterImage field to given value.

### HasTwitterImage

`func (o *WebPageUpdateDto) HasTwitterImage() bool`

HasTwitterImage returns a boolean if a field has been set.

### SetTwitterImageNil

`func (o *WebPageUpdateDto) SetTwitterImageNil(b bool)`

 SetTwitterImageNil sets the value for TwitterImage to be an explicit nil

### UnsetTwitterImage
`func (o *WebPageUpdateDto) UnsetTwitterImage()`

UnsetTwitterImage ensures that no value is present for TwitterImage, not even an explicit nil
### GetTwitterTitle

`func (o *WebPageUpdateDto) GetTwitterTitle() string`

GetTwitterTitle returns the TwitterTitle field if non-nil, zero value otherwise.

### GetTwitterTitleOk

`func (o *WebPageUpdateDto) GetTwitterTitleOk() (*string, bool)`

GetTwitterTitleOk returns a tuple with the TwitterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterTitle

`func (o *WebPageUpdateDto) SetTwitterTitle(v string)`

SetTwitterTitle sets TwitterTitle field to given value.

### HasTwitterTitle

`func (o *WebPageUpdateDto) HasTwitterTitle() bool`

HasTwitterTitle returns a boolean if a field has been set.

### SetTwitterTitleNil

`func (o *WebPageUpdateDto) SetTwitterTitleNil(b bool)`

 SetTwitterTitleNil sets the value for TwitterTitle to be an explicit nil

### UnsetTwitterTitle
`func (o *WebPageUpdateDto) UnsetTwitterTitle()`

UnsetTwitterTitle ensures that no value is present for TwitterTitle, not even an explicit nil
### GetTwitterDescription

`func (o *WebPageUpdateDto) GetTwitterDescription() string`

GetTwitterDescription returns the TwitterDescription field if non-nil, zero value otherwise.

### GetTwitterDescriptionOk

`func (o *WebPageUpdateDto) GetTwitterDescriptionOk() (*string, bool)`

GetTwitterDescriptionOk returns a tuple with the TwitterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterDescription

`func (o *WebPageUpdateDto) SetTwitterDescription(v string)`

SetTwitterDescription sets TwitterDescription field to given value.

### HasTwitterDescription

`func (o *WebPageUpdateDto) HasTwitterDescription() bool`

HasTwitterDescription returns a boolean if a field has been set.

### SetTwitterDescriptionNil

`func (o *WebPageUpdateDto) SetTwitterDescriptionNil(b bool)`

 SetTwitterDescriptionNil sets the value for TwitterDescription to be an explicit nil

### UnsetTwitterDescription
`func (o *WebPageUpdateDto) UnsetTwitterDescription()`

UnsetTwitterDescription ensures that no value is present for TwitterDescription, not even an explicit nil
### GetFacebookImage

`func (o *WebPageUpdateDto) GetFacebookImage() string`

GetFacebookImage returns the FacebookImage field if non-nil, zero value otherwise.

### GetFacebookImageOk

`func (o *WebPageUpdateDto) GetFacebookImageOk() (*string, bool)`

GetFacebookImageOk returns a tuple with the FacebookImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookImage

`func (o *WebPageUpdateDto) SetFacebookImage(v string)`

SetFacebookImage sets FacebookImage field to given value.

### HasFacebookImage

`func (o *WebPageUpdateDto) HasFacebookImage() bool`

HasFacebookImage returns a boolean if a field has been set.

### SetFacebookImageNil

`func (o *WebPageUpdateDto) SetFacebookImageNil(b bool)`

 SetFacebookImageNil sets the value for FacebookImage to be an explicit nil

### UnsetFacebookImage
`func (o *WebPageUpdateDto) UnsetFacebookImage()`

UnsetFacebookImage ensures that no value is present for FacebookImage, not even an explicit nil
### GetFacebookTitle

`func (o *WebPageUpdateDto) GetFacebookTitle() string`

GetFacebookTitle returns the FacebookTitle field if non-nil, zero value otherwise.

### GetFacebookTitleOk

`func (o *WebPageUpdateDto) GetFacebookTitleOk() (*string, bool)`

GetFacebookTitleOk returns a tuple with the FacebookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookTitle

`func (o *WebPageUpdateDto) SetFacebookTitle(v string)`

SetFacebookTitle sets FacebookTitle field to given value.

### HasFacebookTitle

`func (o *WebPageUpdateDto) HasFacebookTitle() bool`

HasFacebookTitle returns a boolean if a field has been set.

### SetFacebookTitleNil

`func (o *WebPageUpdateDto) SetFacebookTitleNil(b bool)`

 SetFacebookTitleNil sets the value for FacebookTitle to be an explicit nil

### UnsetFacebookTitle
`func (o *WebPageUpdateDto) UnsetFacebookTitle()`

UnsetFacebookTitle ensures that no value is present for FacebookTitle, not even an explicit nil
### GetFacebookDescription

`func (o *WebPageUpdateDto) GetFacebookDescription() string`

GetFacebookDescription returns the FacebookDescription field if non-nil, zero value otherwise.

### GetFacebookDescriptionOk

`func (o *WebPageUpdateDto) GetFacebookDescriptionOk() (*string, bool)`

GetFacebookDescriptionOk returns a tuple with the FacebookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookDescription

`func (o *WebPageUpdateDto) SetFacebookDescription(v string)`

SetFacebookDescription sets FacebookDescription field to given value.

### HasFacebookDescription

`func (o *WebPageUpdateDto) HasFacebookDescription() bool`

HasFacebookDescription returns a boolean if a field has been set.

### SetFacebookDescriptionNil

`func (o *WebPageUpdateDto) SetFacebookDescriptionNil(b bool)`

 SetFacebookDescriptionNil sets the value for FacebookDescription to be an explicit nil

### UnsetFacebookDescription
`func (o *WebPageUpdateDto) UnsetFacebookDescription()`

UnsetFacebookDescription ensures that no value is present for FacebookDescription, not even an explicit nil
### GetFeaturedImageUrl

`func (o *WebPageUpdateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *WebPageUpdateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *WebPageUpdateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *WebPageUpdateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *WebPageUpdateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *WebPageUpdateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetContent

`func (o *WebPageUpdateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebPageUpdateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebPageUpdateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebPageUpdateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebPageUpdateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebPageUpdateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCode

`func (o *WebPageUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebPageUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebPageUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebPageUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WebPageUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WebPageUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNamespace

`func (o *WebPageUpdateDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WebPageUpdateDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WebPageUpdateDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WebPageUpdateDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *WebPageUpdateDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *WebPageUpdateDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTypeName

`func (o *WebPageUpdateDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *WebPageUpdateDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *WebPageUpdateDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *WebPageUpdateDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *WebPageUpdateDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *WebPageUpdateDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetGeneratedCode

`func (o *WebPageUpdateDto) GetGeneratedCode() string`

GetGeneratedCode returns the GeneratedCode field if non-nil, zero value otherwise.

### GetGeneratedCodeOk

`func (o *WebPageUpdateDto) GetGeneratedCodeOk() (*string, bool)`

GetGeneratedCodeOk returns a tuple with the GeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCode

`func (o *WebPageUpdateDto) SetGeneratedCode(v string)`

SetGeneratedCode sets GeneratedCode field to given value.

### HasGeneratedCode

`func (o *WebPageUpdateDto) HasGeneratedCode() bool`

HasGeneratedCode returns a boolean if a field has been set.

### SetGeneratedCodeNil

`func (o *WebPageUpdateDto) SetGeneratedCodeNil(b bool)`

 SetGeneratedCodeNil sets the value for GeneratedCode to be an explicit nil

### UnsetGeneratedCode
`func (o *WebPageUpdateDto) UnsetGeneratedCode()`

UnsetGeneratedCode ensures that no value is present for GeneratedCode, not even an explicit nil
### GetCompilationPath

`func (o *WebPageUpdateDto) GetCompilationPath() string`

GetCompilationPath returns the CompilationPath field if non-nil, zero value otherwise.

### GetCompilationPathOk

`func (o *WebPageUpdateDto) GetCompilationPathOk() (*string, bool)`

GetCompilationPathOk returns a tuple with the CompilationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilationPath

`func (o *WebPageUpdateDto) SetCompilationPath(v string)`

SetCompilationPath sets CompilationPath field to given value.

### HasCompilationPath

`func (o *WebPageUpdateDto) HasCompilationPath() bool`

HasCompilationPath returns a boolean if a field has been set.

### SetCompilationPathNil

`func (o *WebPageUpdateDto) SetCompilationPathNil(b bool)`

 SetCompilationPathNil sets the value for CompilationPath to be an explicit nil

### UnsetCompilationPath
`func (o *WebPageUpdateDto) UnsetCompilationPath()`

UnsetCompilationPath ensures that no value is present for CompilationPath, not even an explicit nil
### GetHtmlContent

`func (o *WebPageUpdateDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebPageUpdateDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebPageUpdateDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebPageUpdateDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebPageUpdateDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebPageUpdateDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCodeType

`func (o *WebPageUpdateDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *WebPageUpdateDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *WebPageUpdateDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *WebPageUpdateDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *WebPageUpdateDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *WebPageUpdateDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetCSharpContent

`func (o *WebPageUpdateDto) GetCSharpContent() string`

GetCSharpContent returns the CSharpContent field if non-nil, zero value otherwise.

### GetCSharpContentOk

`func (o *WebPageUpdateDto) GetCSharpContentOk() (*string, bool)`

GetCSharpContentOk returns a tuple with the CSharpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpContent

`func (o *WebPageUpdateDto) SetCSharpContent(v string)`

SetCSharpContent sets CSharpContent field to given value.

### HasCSharpContent

`func (o *WebPageUpdateDto) HasCSharpContent() bool`

HasCSharpContent returns a boolean if a field has been set.

### SetCSharpContentNil

`func (o *WebPageUpdateDto) SetCSharpContentNil(b bool)`

 SetCSharpContentNil sets the value for CSharpContent to be an explicit nil

### UnsetCSharpContent
`func (o *WebPageUpdateDto) UnsetCSharpContent()`

UnsetCSharpContent ensures that no value is present for CSharpContent, not even an explicit nil
### GetRazorContent

`func (o *WebPageUpdateDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *WebPageUpdateDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *WebPageUpdateDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *WebPageUpdateDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *WebPageUpdateDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *WebPageUpdateDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetCssContent

`func (o *WebPageUpdateDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *WebPageUpdateDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *WebPageUpdateDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *WebPageUpdateDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *WebPageUpdateDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *WebPageUpdateDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *WebPageUpdateDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *WebPageUpdateDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *WebPageUpdateDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *WebPageUpdateDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *WebPageUpdateDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *WebPageUpdateDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetCssFiles

`func (o *WebPageUpdateDto) GetCssFiles() string`

GetCssFiles returns the CssFiles field if non-nil, zero value otherwise.

### GetCssFilesOk

`func (o *WebPageUpdateDto) GetCssFilesOk() (*string, bool)`

GetCssFilesOk returns a tuple with the CssFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFiles

`func (o *WebPageUpdateDto) SetCssFiles(v string)`

SetCssFiles sets CssFiles field to given value.

### HasCssFiles

`func (o *WebPageUpdateDto) HasCssFiles() bool`

HasCssFiles returns a boolean if a field has been set.

### SetCssFilesNil

`func (o *WebPageUpdateDto) SetCssFilesNil(b bool)`

 SetCssFilesNil sets the value for CssFiles to be an explicit nil

### UnsetCssFiles
`func (o *WebPageUpdateDto) UnsetCssFiles()`

UnsetCssFiles ensures that no value is present for CssFiles, not even an explicit nil
### GetJsFiles

`func (o *WebPageUpdateDto) GetJsFiles() string`

GetJsFiles returns the JsFiles field if non-nil, zero value otherwise.

### GetJsFilesOk

`func (o *WebPageUpdateDto) GetJsFilesOk() (*string, bool)`

GetJsFilesOk returns a tuple with the JsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsFiles

`func (o *WebPageUpdateDto) SetJsFiles(v string)`

SetJsFiles sets JsFiles field to given value.

### HasJsFiles

`func (o *WebPageUpdateDto) HasJsFiles() bool`

HasJsFiles returns a boolean if a field has been set.

### SetJsFilesNil

`func (o *WebPageUpdateDto) SetJsFilesNil(b bool)`

 SetJsFilesNil sets the value for JsFiles to be an explicit nil

### UnsetJsFiles
`func (o *WebPageUpdateDto) UnsetJsFiles()`

UnsetJsFiles ensures that no value is present for JsFiles, not even an explicit nil
### GetRazorGeneratedCode

`func (o *WebPageUpdateDto) GetRazorGeneratedCode() string`

GetRazorGeneratedCode returns the RazorGeneratedCode field if non-nil, zero value otherwise.

### GetRazorGeneratedCodeOk

`func (o *WebPageUpdateDto) GetRazorGeneratedCodeOk() (*string, bool)`

GetRazorGeneratedCodeOk returns a tuple with the RazorGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorGeneratedCode

`func (o *WebPageUpdateDto) SetRazorGeneratedCode(v string)`

SetRazorGeneratedCode sets RazorGeneratedCode field to given value.

### HasRazorGeneratedCode

`func (o *WebPageUpdateDto) HasRazorGeneratedCode() bool`

HasRazorGeneratedCode returns a boolean if a field has been set.

### SetRazorGeneratedCodeNil

`func (o *WebPageUpdateDto) SetRazorGeneratedCodeNil(b bool)`

 SetRazorGeneratedCodeNil sets the value for RazorGeneratedCode to be an explicit nil

### UnsetRazorGeneratedCode
`func (o *WebPageUpdateDto) UnsetRazorGeneratedCode()`

UnsetRazorGeneratedCode ensures that no value is present for RazorGeneratedCode, not even an explicit nil
### GetCSharpGeneratedCode

`func (o *WebPageUpdateDto) GetCSharpGeneratedCode() string`

GetCSharpGeneratedCode returns the CSharpGeneratedCode field if non-nil, zero value otherwise.

### GetCSharpGeneratedCodeOk

`func (o *WebPageUpdateDto) GetCSharpGeneratedCodeOk() (*string, bool)`

GetCSharpGeneratedCodeOk returns a tuple with the CSharpGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpGeneratedCode

`func (o *WebPageUpdateDto) SetCSharpGeneratedCode(v string)`

SetCSharpGeneratedCode sets CSharpGeneratedCode field to given value.

### HasCSharpGeneratedCode

`func (o *WebPageUpdateDto) HasCSharpGeneratedCode() bool`

HasCSharpGeneratedCode returns a boolean if a field has been set.

### SetCSharpGeneratedCodeNil

`func (o *WebPageUpdateDto) SetCSharpGeneratedCodeNil(b bool)`

 SetCSharpGeneratedCodeNil sets the value for CSharpGeneratedCode to be an explicit nil

### UnsetCSharpGeneratedCode
`func (o *WebPageUpdateDto) UnsetCSharpGeneratedCode()`

UnsetCSharpGeneratedCode ensures that no value is present for CSharpGeneratedCode, not even an explicit nil
### GetPrecompiledLogicSize

`func (o *WebPageUpdateDto) GetPrecompiledLogicSize() int32`

GetPrecompiledLogicSize returns the PrecompiledLogicSize field if non-nil, zero value otherwise.

### GetPrecompiledLogicSizeOk

`func (o *WebPageUpdateDto) GetPrecompiledLogicSizeOk() (*int32, bool)`

GetPrecompiledLogicSizeOk returns a tuple with the PrecompiledLogicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicSize

`func (o *WebPageUpdateDto) SetPrecompiledLogicSize(v int32)`

SetPrecompiledLogicSize sets PrecompiledLogicSize field to given value.

### HasPrecompiledLogicSize

`func (o *WebPageUpdateDto) HasPrecompiledLogicSize() bool`

HasPrecompiledLogicSize returns a boolean if a field has been set.

### GetPrecompiledLogicSizeLong

`func (o *WebPageUpdateDto) GetPrecompiledLogicSizeLong() int64`

GetPrecompiledLogicSizeLong returns the PrecompiledLogicSizeLong field if non-nil, zero value otherwise.

### GetPrecompiledLogicSizeLongOk

`func (o *WebPageUpdateDto) GetPrecompiledLogicSizeLongOk() (*int64, bool)`

GetPrecompiledLogicSizeLongOk returns a tuple with the PrecompiledLogicSizeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicSizeLong

`func (o *WebPageUpdateDto) SetPrecompiledLogicSizeLong(v int64)`

SetPrecompiledLogicSizeLong sets PrecompiledLogicSizeLong field to given value.

### HasPrecompiledLogicSizeLong

`func (o *WebPageUpdateDto) HasPrecompiledLogicSizeLong() bool`

HasPrecompiledLogicSizeLong returns a boolean if a field has been set.

### GetPrecompiledViewSize

`func (o *WebPageUpdateDto) GetPrecompiledViewSize() int32`

GetPrecompiledViewSize returns the PrecompiledViewSize field if non-nil, zero value otherwise.

### GetPrecompiledViewSizeOk

`func (o *WebPageUpdateDto) GetPrecompiledViewSizeOk() (*int32, bool)`

GetPrecompiledViewSizeOk returns a tuple with the PrecompiledViewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledViewSize

`func (o *WebPageUpdateDto) SetPrecompiledViewSize(v int32)`

SetPrecompiledViewSize sets PrecompiledViewSize field to given value.

### HasPrecompiledViewSize

`func (o *WebPageUpdateDto) HasPrecompiledViewSize() bool`

HasPrecompiledViewSize returns a boolean if a field has been set.

### GetPrecompiledViewSizeLong

`func (o *WebPageUpdateDto) GetPrecompiledViewSizeLong() int64`

GetPrecompiledViewSizeLong returns the PrecompiledViewSizeLong field if non-nil, zero value otherwise.

### GetPrecompiledViewSizeLongOk

`func (o *WebPageUpdateDto) GetPrecompiledViewSizeLongOk() (*int64, bool)`

GetPrecompiledViewSizeLongOk returns a tuple with the PrecompiledViewSizeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledViewSizeLong

`func (o *WebPageUpdateDto) SetPrecompiledViewSizeLong(v int64)`

SetPrecompiledViewSizeLong sets PrecompiledViewSizeLong field to given value.

### HasPrecompiledViewSizeLong

`func (o *WebPageUpdateDto) HasPrecompiledViewSizeLong() bool`

HasPrecompiledViewSizeLong returns a boolean if a field has been set.

### GetPrecompiledLogicViewSize

`func (o *WebPageUpdateDto) GetPrecompiledLogicViewSize() int32`

GetPrecompiledLogicViewSize returns the PrecompiledLogicViewSize field if non-nil, zero value otherwise.

### GetPrecompiledLogicViewSizeOk

`func (o *WebPageUpdateDto) GetPrecompiledLogicViewSizeOk() (*int32, bool)`

GetPrecompiledLogicViewSizeOk returns a tuple with the PrecompiledLogicViewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicViewSize

`func (o *WebPageUpdateDto) SetPrecompiledLogicViewSize(v int32)`

SetPrecompiledLogicViewSize sets PrecompiledLogicViewSize field to given value.

### HasPrecompiledLogicViewSize

`func (o *WebPageUpdateDto) HasPrecompiledLogicViewSize() bool`

HasPrecompiledLogicViewSize returns a boolean if a field has been set.

### GetTemplate

`func (o *WebPageUpdateDto) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebPageUpdateDto) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebPageUpdateDto) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebPageUpdateDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefault

`func (o *WebPageUpdateDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WebPageUpdateDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WebPageUpdateDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WebPageUpdateDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnable

`func (o *WebPageUpdateDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WebPageUpdateDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WebPageUpdateDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *WebPageUpdateDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableComments

`func (o *WebPageUpdateDto) GetEnableComments() bool`

GetEnableComments returns the EnableComments field if non-nil, zero value otherwise.

### GetEnableCommentsOk

`func (o *WebPageUpdateDto) GetEnableCommentsOk() (*bool, bool)`

GetEnableCommentsOk returns a tuple with the EnableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComments

`func (o *WebPageUpdateDto) SetEnableComments(v bool)`

SetEnableComments sets EnableComments field to given value.

### HasEnableComments

`func (o *WebPageUpdateDto) HasEnableComments() bool`

HasEnableComments returns a boolean if a field has been set.

### GetDisplaySocialBox

`func (o *WebPageUpdateDto) GetDisplaySocialBox() bool`

GetDisplaySocialBox returns the DisplaySocialBox field if non-nil, zero value otherwise.

### GetDisplaySocialBoxOk

`func (o *WebPageUpdateDto) GetDisplaySocialBoxOk() (*bool, bool)`

GetDisplaySocialBoxOk returns a tuple with the DisplaySocialBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySocialBox

`func (o *WebPageUpdateDto) SetDisplaySocialBox(v bool)`

SetDisplaySocialBox sets DisplaySocialBox field to given value.

### HasDisplaySocialBox

`func (o *WebPageUpdateDto) HasDisplaySocialBox() bool`

HasDisplaySocialBox returns a boolean if a field has been set.

### GetPublished

`func (o *WebPageUpdateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *WebPageUpdateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *WebPageUpdateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *WebPageUpdateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInTrashCan

`func (o *WebPageUpdateDto) GetInTrashCan() bool`

GetInTrashCan returns the InTrashCan field if non-nil, zero value otherwise.

### GetInTrashCanOk

`func (o *WebPageUpdateDto) GetInTrashCanOk() (*bool, bool)`

GetInTrashCanOk returns a tuple with the InTrashCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrashCan

`func (o *WebPageUpdateDto) SetInTrashCan(v bool)`

SetInTrashCan sets InTrashCan field to given value.

### HasInTrashCan

`func (o *WebPageUpdateDto) HasInTrashCan() bool`

HasInTrashCan returns a boolean if a field has been set.

### GetSystemLocked

`func (o *WebPageUpdateDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *WebPageUpdateDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *WebPageUpdateDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *WebPageUpdateDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetAllowPingbacks

`func (o *WebPageUpdateDto) GetAllowPingbacks() bool`

GetAllowPingbacks returns the AllowPingbacks field if non-nil, zero value otherwise.

### GetAllowPingbacksOk

`func (o *WebPageUpdateDto) GetAllowPingbacksOk() (*bool, bool)`

GetAllowPingbacksOk returns a tuple with the AllowPingbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPingbacks

`func (o *WebPageUpdateDto) SetAllowPingbacks(v bool)`

SetAllowPingbacks sets AllowPingbacks field to given value.

### HasAllowPingbacks

`func (o *WebPageUpdateDto) HasAllowPingbacks() bool`

HasAllowPingbacks returns a boolean if a field has been set.

### GetAllowTrackbacks

`func (o *WebPageUpdateDto) GetAllowTrackbacks() bool`

GetAllowTrackbacks returns the AllowTrackbacks field if non-nil, zero value otherwise.

### GetAllowTrackbacksOk

`func (o *WebPageUpdateDto) GetAllowTrackbacksOk() (*bool, bool)`

GetAllowTrackbacksOk returns a tuple with the AllowTrackbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrackbacks

`func (o *WebPageUpdateDto) SetAllowTrackbacks(v bool)`

SetAllowTrackbacks sets AllowTrackbacks field to given value.

### HasAllowTrackbacks

`func (o *WebPageUpdateDto) HasAllowTrackbacks() bool`

HasAllowTrackbacks returns a boolean if a field has been set.

### GetCornerstoneContent

`func (o *WebPageUpdateDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebPageUpdateDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebPageUpdateDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebPageUpdateDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetIsEssentialContent

`func (o *WebPageUpdateDto) GetIsEssentialContent() bool`

GetIsEssentialContent returns the IsEssentialContent field if non-nil, zero value otherwise.

### GetIsEssentialContentOk

`func (o *WebPageUpdateDto) GetIsEssentialContentOk() (*bool, bool)`

GetIsEssentialContentOk returns a tuple with the IsEssentialContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEssentialContent

`func (o *WebPageUpdateDto) SetIsEssentialContent(v bool)`

SetIsEssentialContent sets IsEssentialContent field to given value.

### HasIsEssentialContent

`func (o *WebPageUpdateDto) HasIsEssentialContent() bool`

HasIsEssentialContent returns a boolean if a field has been set.

### GetAllowSearchEngineIndexing

`func (o *WebPageUpdateDto) GetAllowSearchEngineIndexing() bool`

GetAllowSearchEngineIndexing returns the AllowSearchEngineIndexing field if non-nil, zero value otherwise.

### GetAllowSearchEngineIndexingOk

`func (o *WebPageUpdateDto) GetAllowSearchEngineIndexingOk() (*bool, bool)`

GetAllowSearchEngineIndexingOk returns a tuple with the AllowSearchEngineIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearchEngineIndexing

`func (o *WebPageUpdateDto) SetAllowSearchEngineIndexing(v bool)`

SetAllowSearchEngineIndexing sets AllowSearchEngineIndexing field to given value.

### HasAllowSearchEngineIndexing

`func (o *WebPageUpdateDto) HasAllowSearchEngineIndexing() bool`

HasAllowSearchEngineIndexing returns a boolean if a field has been set.

### GetIsHomePage

`func (o *WebPageUpdateDto) GetIsHomePage() bool`

GetIsHomePage returns the IsHomePage field if non-nil, zero value otherwise.

### GetIsHomePageOk

`func (o *WebPageUpdateDto) GetIsHomePageOk() (*bool, bool)`

GetIsHomePageOk returns a tuple with the IsHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHomePage

`func (o *WebPageUpdateDto) SetIsHomePage(v bool)`

SetIsHomePage sets IsHomePage field to given value.

### HasIsHomePage

`func (o *WebPageUpdateDto) HasIsHomePage() bool`

HasIsHomePage returns a boolean if a field has been set.

### GetIsStorePage

`func (o *WebPageUpdateDto) GetIsStorePage() bool`

GetIsStorePage returns the IsStorePage field if non-nil, zero value otherwise.

### GetIsStorePageOk

`func (o *WebPageUpdateDto) GetIsStorePageOk() (*bool, bool)`

GetIsStorePageOk returns a tuple with the IsStorePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorePage

`func (o *WebPageUpdateDto) SetIsStorePage(v bool)`

SetIsStorePage sets IsStorePage field to given value.

### HasIsStorePage

`func (o *WebPageUpdateDto) HasIsStorePage() bool`

HasIsStorePage returns a boolean if a field has been set.

### GetIsCartPage

`func (o *WebPageUpdateDto) GetIsCartPage() bool`

GetIsCartPage returns the IsCartPage field if non-nil, zero value otherwise.

### GetIsCartPageOk

`func (o *WebPageUpdateDto) GetIsCartPageOk() (*bool, bool)`

GetIsCartPageOk returns a tuple with the IsCartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCartPage

`func (o *WebPageUpdateDto) SetIsCartPage(v bool)`

SetIsCartPage sets IsCartPage field to given value.

### HasIsCartPage

`func (o *WebPageUpdateDto) HasIsCartPage() bool`

HasIsCartPage returns a boolean if a field has been set.

### GetIsBlogPage

`func (o *WebPageUpdateDto) GetIsBlogPage() bool`

GetIsBlogPage returns the IsBlogPage field if non-nil, zero value otherwise.

### GetIsBlogPageOk

`func (o *WebPageUpdateDto) GetIsBlogPageOk() (*bool, bool)`

GetIsBlogPageOk returns a tuple with the IsBlogPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlogPage

`func (o *WebPageUpdateDto) SetIsBlogPage(v bool)`

SetIsBlogPage sets IsBlogPage field to given value.

### HasIsBlogPage

`func (o *WebPageUpdateDto) HasIsBlogPage() bool`

HasIsBlogPage returns a boolean if a field has been set.

### GetIsAccountPage

`func (o *WebPageUpdateDto) GetIsAccountPage() bool`

GetIsAccountPage returns the IsAccountPage field if non-nil, zero value otherwise.

### GetIsAccountPageOk

`func (o *WebPageUpdateDto) GetIsAccountPageOk() (*bool, bool)`

GetIsAccountPageOk returns a tuple with the IsAccountPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPage

`func (o *WebPageUpdateDto) SetIsAccountPage(v bool)`

SetIsAccountPage sets IsAccountPage field to given value.

### HasIsAccountPage

`func (o *WebPageUpdateDto) HasIsAccountPage() bool`

HasIsAccountPage returns a boolean if a field has been set.

### GetIsCheckoutPage

`func (o *WebPageUpdateDto) GetIsCheckoutPage() bool`

GetIsCheckoutPage returns the IsCheckoutPage field if non-nil, zero value otherwise.

### GetIsCheckoutPageOk

`func (o *WebPageUpdateDto) GetIsCheckoutPageOk() (*bool, bool)`

GetIsCheckoutPageOk returns a tuple with the IsCheckoutPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCheckoutPage

`func (o *WebPageUpdateDto) SetIsCheckoutPage(v bool)`

SetIsCheckoutPage sets IsCheckoutPage field to given value.

### HasIsCheckoutPage

`func (o *WebPageUpdateDto) HasIsCheckoutPage() bool`

HasIsCheckoutPage returns a boolean if a field has been set.

### GetIsWishListsPage

`func (o *WebPageUpdateDto) GetIsWishListsPage() bool`

GetIsWishListsPage returns the IsWishListsPage field if non-nil, zero value otherwise.

### GetIsWishListsPageOk

`func (o *WebPageUpdateDto) GetIsWishListsPageOk() (*bool, bool)`

GetIsWishListsPageOk returns a tuple with the IsWishListsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWishListsPage

`func (o *WebPageUpdateDto) SetIsWishListsPage(v bool)`

SetIsWishListsPage sets IsWishListsPage field to given value.

### HasIsWishListsPage

`func (o *WebPageUpdateDto) HasIsWishListsPage() bool`

HasIsWishListsPage returns a boolean if a field has been set.

### GetIsContactPage

`func (o *WebPageUpdateDto) GetIsContactPage() bool`

GetIsContactPage returns the IsContactPage field if non-nil, zero value otherwise.

### GetIsContactPageOk

`func (o *WebPageUpdateDto) GetIsContactPageOk() (*bool, bool)`

GetIsContactPageOk returns a tuple with the IsContactPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContactPage

`func (o *WebPageUpdateDto) SetIsContactPage(v bool)`

SetIsContactPage sets IsContactPage field to given value.

### HasIsContactPage

`func (o *WebPageUpdateDto) HasIsContactPage() bool`

HasIsContactPage returns a boolean if a field has been set.

### GetIsPrivacyPolicyPage

`func (o *WebPageUpdateDto) GetIsPrivacyPolicyPage() bool`

GetIsPrivacyPolicyPage returns the IsPrivacyPolicyPage field if non-nil, zero value otherwise.

### GetIsPrivacyPolicyPageOk

`func (o *WebPageUpdateDto) GetIsPrivacyPolicyPageOk() (*bool, bool)`

GetIsPrivacyPolicyPageOk returns a tuple with the IsPrivacyPolicyPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivacyPolicyPage

`func (o *WebPageUpdateDto) SetIsPrivacyPolicyPage(v bool)`

SetIsPrivacyPolicyPage sets IsPrivacyPolicyPage field to given value.

### HasIsPrivacyPolicyPage

`func (o *WebPageUpdateDto) HasIsPrivacyPolicyPage() bool`

HasIsPrivacyPolicyPage returns a boolean if a field has been set.

### GetIsTermsOfServicePage

`func (o *WebPageUpdateDto) GetIsTermsOfServicePage() bool`

GetIsTermsOfServicePage returns the IsTermsOfServicePage field if non-nil, zero value otherwise.

### GetIsTermsOfServicePageOk

`func (o *WebPageUpdateDto) GetIsTermsOfServicePageOk() (*bool, bool)`

GetIsTermsOfServicePageOk returns a tuple with the IsTermsOfServicePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTermsOfServicePage

`func (o *WebPageUpdateDto) SetIsTermsOfServicePage(v bool)`

SetIsTermsOfServicePage sets IsTermsOfServicePage field to given value.

### HasIsTermsOfServicePage

`func (o *WebPageUpdateDto) HasIsTermsOfServicePage() bool`

HasIsTermsOfServicePage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


