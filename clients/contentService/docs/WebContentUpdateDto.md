# WebContentUpdateDto

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

## Methods

### NewWebContentUpdateDto

`func NewWebContentUpdateDto() *WebContentUpdateDto`

NewWebContentUpdateDto instantiates a new WebContentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebContentUpdateDtoWithDefaults

`func NewWebContentUpdateDtoWithDefaults() *WebContentUpdateDto`

NewWebContentUpdateDtoWithDefaults instantiates a new WebContentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *WebContentUpdateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebContentUpdateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebContentUpdateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebContentUpdateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSlug

`func (o *WebContentUpdateDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebContentUpdateDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebContentUpdateDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebContentUpdateDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebContentUpdateDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebContentUpdateDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *WebContentUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebContentUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebContentUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebContentUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebContentUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebContentUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *WebContentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebContentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebContentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebContentUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebContentUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebContentUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExcerpt

`func (o *WebContentUpdateDto) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *WebContentUpdateDto) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *WebContentUpdateDto) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *WebContentUpdateDto) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *WebContentUpdateDto) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *WebContentUpdateDto) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetPassword

`func (o *WebContentUpdateDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebContentUpdateDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebContentUpdateDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebContentUpdateDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebContentUpdateDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebContentUpdateDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *WebContentUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebContentUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebContentUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebContentUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebContentUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebContentUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHighlightImage

`func (o *WebContentUpdateDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *WebContentUpdateDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *WebContentUpdateDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *WebContentUpdateDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *WebContentUpdateDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *WebContentUpdateDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetCanonicalUrl

`func (o *WebContentUpdateDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebContentUpdateDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebContentUpdateDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebContentUpdateDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebContentUpdateDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebContentUpdateDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetSeoTitle

`func (o *WebContentUpdateDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebContentUpdateDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebContentUpdateDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebContentUpdateDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebContentUpdateDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebContentUpdateDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoKeyWords

`func (o *WebContentUpdateDto) GetSeoKeyWords() string`

GetSeoKeyWords returns the SeoKeyWords field if non-nil, zero value otherwise.

### GetSeoKeyWordsOk

`func (o *WebContentUpdateDto) GetSeoKeyWordsOk() (*string, bool)`

GetSeoKeyWordsOk returns a tuple with the SeoKeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyWords

`func (o *WebContentUpdateDto) SetSeoKeyWords(v string)`

SetSeoKeyWords sets SeoKeyWords field to given value.

### HasSeoKeyWords

`func (o *WebContentUpdateDto) HasSeoKeyWords() bool`

HasSeoKeyWords returns a boolean if a field has been set.

### SetSeoKeyWordsNil

`func (o *WebContentUpdateDto) SetSeoKeyWordsNil(b bool)`

 SetSeoKeyWordsNil sets the value for SeoKeyWords to be an explicit nil

### UnsetSeoKeyWords
`func (o *WebContentUpdateDto) UnsetSeoKeyWords()`

UnsetSeoKeyWords ensures that no value is present for SeoKeyWords, not even an explicit nil
### GetSeoKeyPhrases

`func (o *WebContentUpdateDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebContentUpdateDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebContentUpdateDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebContentUpdateDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebContentUpdateDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebContentUpdateDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetMetaDescription

`func (o *WebContentUpdateDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebContentUpdateDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebContentUpdateDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebContentUpdateDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebContentUpdateDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebContentUpdateDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetTwitterImage

`func (o *WebContentUpdateDto) GetTwitterImage() string`

GetTwitterImage returns the TwitterImage field if non-nil, zero value otherwise.

### GetTwitterImageOk

`func (o *WebContentUpdateDto) GetTwitterImageOk() (*string, bool)`

GetTwitterImageOk returns a tuple with the TwitterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterImage

`func (o *WebContentUpdateDto) SetTwitterImage(v string)`

SetTwitterImage sets TwitterImage field to given value.

### HasTwitterImage

`func (o *WebContentUpdateDto) HasTwitterImage() bool`

HasTwitterImage returns a boolean if a field has been set.

### SetTwitterImageNil

`func (o *WebContentUpdateDto) SetTwitterImageNil(b bool)`

 SetTwitterImageNil sets the value for TwitterImage to be an explicit nil

### UnsetTwitterImage
`func (o *WebContentUpdateDto) UnsetTwitterImage()`

UnsetTwitterImage ensures that no value is present for TwitterImage, not even an explicit nil
### GetTwitterTitle

`func (o *WebContentUpdateDto) GetTwitterTitle() string`

GetTwitterTitle returns the TwitterTitle field if non-nil, zero value otherwise.

### GetTwitterTitleOk

`func (o *WebContentUpdateDto) GetTwitterTitleOk() (*string, bool)`

GetTwitterTitleOk returns a tuple with the TwitterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterTitle

`func (o *WebContentUpdateDto) SetTwitterTitle(v string)`

SetTwitterTitle sets TwitterTitle field to given value.

### HasTwitterTitle

`func (o *WebContentUpdateDto) HasTwitterTitle() bool`

HasTwitterTitle returns a boolean if a field has been set.

### SetTwitterTitleNil

`func (o *WebContentUpdateDto) SetTwitterTitleNil(b bool)`

 SetTwitterTitleNil sets the value for TwitterTitle to be an explicit nil

### UnsetTwitterTitle
`func (o *WebContentUpdateDto) UnsetTwitterTitle()`

UnsetTwitterTitle ensures that no value is present for TwitterTitle, not even an explicit nil
### GetTwitterDescription

`func (o *WebContentUpdateDto) GetTwitterDescription() string`

GetTwitterDescription returns the TwitterDescription field if non-nil, zero value otherwise.

### GetTwitterDescriptionOk

`func (o *WebContentUpdateDto) GetTwitterDescriptionOk() (*string, bool)`

GetTwitterDescriptionOk returns a tuple with the TwitterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterDescription

`func (o *WebContentUpdateDto) SetTwitterDescription(v string)`

SetTwitterDescription sets TwitterDescription field to given value.

### HasTwitterDescription

`func (o *WebContentUpdateDto) HasTwitterDescription() bool`

HasTwitterDescription returns a boolean if a field has been set.

### SetTwitterDescriptionNil

`func (o *WebContentUpdateDto) SetTwitterDescriptionNil(b bool)`

 SetTwitterDescriptionNil sets the value for TwitterDescription to be an explicit nil

### UnsetTwitterDescription
`func (o *WebContentUpdateDto) UnsetTwitterDescription()`

UnsetTwitterDescription ensures that no value is present for TwitterDescription, not even an explicit nil
### GetFacebookImage

`func (o *WebContentUpdateDto) GetFacebookImage() string`

GetFacebookImage returns the FacebookImage field if non-nil, zero value otherwise.

### GetFacebookImageOk

`func (o *WebContentUpdateDto) GetFacebookImageOk() (*string, bool)`

GetFacebookImageOk returns a tuple with the FacebookImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookImage

`func (o *WebContentUpdateDto) SetFacebookImage(v string)`

SetFacebookImage sets FacebookImage field to given value.

### HasFacebookImage

`func (o *WebContentUpdateDto) HasFacebookImage() bool`

HasFacebookImage returns a boolean if a field has been set.

### SetFacebookImageNil

`func (o *WebContentUpdateDto) SetFacebookImageNil(b bool)`

 SetFacebookImageNil sets the value for FacebookImage to be an explicit nil

### UnsetFacebookImage
`func (o *WebContentUpdateDto) UnsetFacebookImage()`

UnsetFacebookImage ensures that no value is present for FacebookImage, not even an explicit nil
### GetFacebookTitle

`func (o *WebContentUpdateDto) GetFacebookTitle() string`

GetFacebookTitle returns the FacebookTitle field if non-nil, zero value otherwise.

### GetFacebookTitleOk

`func (o *WebContentUpdateDto) GetFacebookTitleOk() (*string, bool)`

GetFacebookTitleOk returns a tuple with the FacebookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookTitle

`func (o *WebContentUpdateDto) SetFacebookTitle(v string)`

SetFacebookTitle sets FacebookTitle field to given value.

### HasFacebookTitle

`func (o *WebContentUpdateDto) HasFacebookTitle() bool`

HasFacebookTitle returns a boolean if a field has been set.

### SetFacebookTitleNil

`func (o *WebContentUpdateDto) SetFacebookTitleNil(b bool)`

 SetFacebookTitleNil sets the value for FacebookTitle to be an explicit nil

### UnsetFacebookTitle
`func (o *WebContentUpdateDto) UnsetFacebookTitle()`

UnsetFacebookTitle ensures that no value is present for FacebookTitle, not even an explicit nil
### GetFacebookDescription

`func (o *WebContentUpdateDto) GetFacebookDescription() string`

GetFacebookDescription returns the FacebookDescription field if non-nil, zero value otherwise.

### GetFacebookDescriptionOk

`func (o *WebContentUpdateDto) GetFacebookDescriptionOk() (*string, bool)`

GetFacebookDescriptionOk returns a tuple with the FacebookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookDescription

`func (o *WebContentUpdateDto) SetFacebookDescription(v string)`

SetFacebookDescription sets FacebookDescription field to given value.

### HasFacebookDescription

`func (o *WebContentUpdateDto) HasFacebookDescription() bool`

HasFacebookDescription returns a boolean if a field has been set.

### SetFacebookDescriptionNil

`func (o *WebContentUpdateDto) SetFacebookDescriptionNil(b bool)`

 SetFacebookDescriptionNil sets the value for FacebookDescription to be an explicit nil

### UnsetFacebookDescription
`func (o *WebContentUpdateDto) UnsetFacebookDescription()`

UnsetFacebookDescription ensures that no value is present for FacebookDescription, not even an explicit nil
### GetFeaturedImageUrl

`func (o *WebContentUpdateDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *WebContentUpdateDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *WebContentUpdateDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *WebContentUpdateDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *WebContentUpdateDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *WebContentUpdateDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetContent

`func (o *WebContentUpdateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebContentUpdateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebContentUpdateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebContentUpdateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebContentUpdateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebContentUpdateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCode

`func (o *WebContentUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebContentUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebContentUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebContentUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WebContentUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WebContentUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNamespace

`func (o *WebContentUpdateDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WebContentUpdateDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WebContentUpdateDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WebContentUpdateDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *WebContentUpdateDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *WebContentUpdateDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTypeName

`func (o *WebContentUpdateDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *WebContentUpdateDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *WebContentUpdateDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *WebContentUpdateDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *WebContentUpdateDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *WebContentUpdateDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetGeneratedCode

`func (o *WebContentUpdateDto) GetGeneratedCode() string`

GetGeneratedCode returns the GeneratedCode field if non-nil, zero value otherwise.

### GetGeneratedCodeOk

`func (o *WebContentUpdateDto) GetGeneratedCodeOk() (*string, bool)`

GetGeneratedCodeOk returns a tuple with the GeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCode

`func (o *WebContentUpdateDto) SetGeneratedCode(v string)`

SetGeneratedCode sets GeneratedCode field to given value.

### HasGeneratedCode

`func (o *WebContentUpdateDto) HasGeneratedCode() bool`

HasGeneratedCode returns a boolean if a field has been set.

### SetGeneratedCodeNil

`func (o *WebContentUpdateDto) SetGeneratedCodeNil(b bool)`

 SetGeneratedCodeNil sets the value for GeneratedCode to be an explicit nil

### UnsetGeneratedCode
`func (o *WebContentUpdateDto) UnsetGeneratedCode()`

UnsetGeneratedCode ensures that no value is present for GeneratedCode, not even an explicit nil
### GetCompilationPath

`func (o *WebContentUpdateDto) GetCompilationPath() string`

GetCompilationPath returns the CompilationPath field if non-nil, zero value otherwise.

### GetCompilationPathOk

`func (o *WebContentUpdateDto) GetCompilationPathOk() (*string, bool)`

GetCompilationPathOk returns a tuple with the CompilationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilationPath

`func (o *WebContentUpdateDto) SetCompilationPath(v string)`

SetCompilationPath sets CompilationPath field to given value.

### HasCompilationPath

`func (o *WebContentUpdateDto) HasCompilationPath() bool`

HasCompilationPath returns a boolean if a field has been set.

### SetCompilationPathNil

`func (o *WebContentUpdateDto) SetCompilationPathNil(b bool)`

 SetCompilationPathNil sets the value for CompilationPath to be an explicit nil

### UnsetCompilationPath
`func (o *WebContentUpdateDto) UnsetCompilationPath()`

UnsetCompilationPath ensures that no value is present for CompilationPath, not even an explicit nil
### GetHtmlContent

`func (o *WebContentUpdateDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebContentUpdateDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebContentUpdateDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebContentUpdateDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebContentUpdateDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebContentUpdateDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCodeType

`func (o *WebContentUpdateDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *WebContentUpdateDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *WebContentUpdateDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *WebContentUpdateDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *WebContentUpdateDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *WebContentUpdateDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetCSharpContent

`func (o *WebContentUpdateDto) GetCSharpContent() string`

GetCSharpContent returns the CSharpContent field if non-nil, zero value otherwise.

### GetCSharpContentOk

`func (o *WebContentUpdateDto) GetCSharpContentOk() (*string, bool)`

GetCSharpContentOk returns a tuple with the CSharpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpContent

`func (o *WebContentUpdateDto) SetCSharpContent(v string)`

SetCSharpContent sets CSharpContent field to given value.

### HasCSharpContent

`func (o *WebContentUpdateDto) HasCSharpContent() bool`

HasCSharpContent returns a boolean if a field has been set.

### SetCSharpContentNil

`func (o *WebContentUpdateDto) SetCSharpContentNil(b bool)`

 SetCSharpContentNil sets the value for CSharpContent to be an explicit nil

### UnsetCSharpContent
`func (o *WebContentUpdateDto) UnsetCSharpContent()`

UnsetCSharpContent ensures that no value is present for CSharpContent, not even an explicit nil
### GetRazorContent

`func (o *WebContentUpdateDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *WebContentUpdateDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *WebContentUpdateDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *WebContentUpdateDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *WebContentUpdateDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *WebContentUpdateDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetCssContent

`func (o *WebContentUpdateDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *WebContentUpdateDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *WebContentUpdateDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *WebContentUpdateDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *WebContentUpdateDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *WebContentUpdateDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *WebContentUpdateDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *WebContentUpdateDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *WebContentUpdateDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *WebContentUpdateDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *WebContentUpdateDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *WebContentUpdateDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetCssFiles

`func (o *WebContentUpdateDto) GetCssFiles() string`

GetCssFiles returns the CssFiles field if non-nil, zero value otherwise.

### GetCssFilesOk

`func (o *WebContentUpdateDto) GetCssFilesOk() (*string, bool)`

GetCssFilesOk returns a tuple with the CssFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFiles

`func (o *WebContentUpdateDto) SetCssFiles(v string)`

SetCssFiles sets CssFiles field to given value.

### HasCssFiles

`func (o *WebContentUpdateDto) HasCssFiles() bool`

HasCssFiles returns a boolean if a field has been set.

### SetCssFilesNil

`func (o *WebContentUpdateDto) SetCssFilesNil(b bool)`

 SetCssFilesNil sets the value for CssFiles to be an explicit nil

### UnsetCssFiles
`func (o *WebContentUpdateDto) UnsetCssFiles()`

UnsetCssFiles ensures that no value is present for CssFiles, not even an explicit nil
### GetJsFiles

`func (o *WebContentUpdateDto) GetJsFiles() string`

GetJsFiles returns the JsFiles field if non-nil, zero value otherwise.

### GetJsFilesOk

`func (o *WebContentUpdateDto) GetJsFilesOk() (*string, bool)`

GetJsFilesOk returns a tuple with the JsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsFiles

`func (o *WebContentUpdateDto) SetJsFiles(v string)`

SetJsFiles sets JsFiles field to given value.

### HasJsFiles

`func (o *WebContentUpdateDto) HasJsFiles() bool`

HasJsFiles returns a boolean if a field has been set.

### SetJsFilesNil

`func (o *WebContentUpdateDto) SetJsFilesNil(b bool)`

 SetJsFilesNil sets the value for JsFiles to be an explicit nil

### UnsetJsFiles
`func (o *WebContentUpdateDto) UnsetJsFiles()`

UnsetJsFiles ensures that no value is present for JsFiles, not even an explicit nil
### GetRazorGeneratedCode

`func (o *WebContentUpdateDto) GetRazorGeneratedCode() string`

GetRazorGeneratedCode returns the RazorGeneratedCode field if non-nil, zero value otherwise.

### GetRazorGeneratedCodeOk

`func (o *WebContentUpdateDto) GetRazorGeneratedCodeOk() (*string, bool)`

GetRazorGeneratedCodeOk returns a tuple with the RazorGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorGeneratedCode

`func (o *WebContentUpdateDto) SetRazorGeneratedCode(v string)`

SetRazorGeneratedCode sets RazorGeneratedCode field to given value.

### HasRazorGeneratedCode

`func (o *WebContentUpdateDto) HasRazorGeneratedCode() bool`

HasRazorGeneratedCode returns a boolean if a field has been set.

### SetRazorGeneratedCodeNil

`func (o *WebContentUpdateDto) SetRazorGeneratedCodeNil(b bool)`

 SetRazorGeneratedCodeNil sets the value for RazorGeneratedCode to be an explicit nil

### UnsetRazorGeneratedCode
`func (o *WebContentUpdateDto) UnsetRazorGeneratedCode()`

UnsetRazorGeneratedCode ensures that no value is present for RazorGeneratedCode, not even an explicit nil
### GetCSharpGeneratedCode

`func (o *WebContentUpdateDto) GetCSharpGeneratedCode() string`

GetCSharpGeneratedCode returns the CSharpGeneratedCode field if non-nil, zero value otherwise.

### GetCSharpGeneratedCodeOk

`func (o *WebContentUpdateDto) GetCSharpGeneratedCodeOk() (*string, bool)`

GetCSharpGeneratedCodeOk returns a tuple with the CSharpGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpGeneratedCode

`func (o *WebContentUpdateDto) SetCSharpGeneratedCode(v string)`

SetCSharpGeneratedCode sets CSharpGeneratedCode field to given value.

### HasCSharpGeneratedCode

`func (o *WebContentUpdateDto) HasCSharpGeneratedCode() bool`

HasCSharpGeneratedCode returns a boolean if a field has been set.

### SetCSharpGeneratedCodeNil

`func (o *WebContentUpdateDto) SetCSharpGeneratedCodeNil(b bool)`

 SetCSharpGeneratedCodeNil sets the value for CSharpGeneratedCode to be an explicit nil

### UnsetCSharpGeneratedCode
`func (o *WebContentUpdateDto) UnsetCSharpGeneratedCode()`

UnsetCSharpGeneratedCode ensures that no value is present for CSharpGeneratedCode, not even an explicit nil
### GetPrecompiledLogicSize

`func (o *WebContentUpdateDto) GetPrecompiledLogicSize() int32`

GetPrecompiledLogicSize returns the PrecompiledLogicSize field if non-nil, zero value otherwise.

### GetPrecompiledLogicSizeOk

`func (o *WebContentUpdateDto) GetPrecompiledLogicSizeOk() (*int32, bool)`

GetPrecompiledLogicSizeOk returns a tuple with the PrecompiledLogicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicSize

`func (o *WebContentUpdateDto) SetPrecompiledLogicSize(v int32)`

SetPrecompiledLogicSize sets PrecompiledLogicSize field to given value.

### HasPrecompiledLogicSize

`func (o *WebContentUpdateDto) HasPrecompiledLogicSize() bool`

HasPrecompiledLogicSize returns a boolean if a field has been set.

### GetPrecompiledLogicSizeLong

`func (o *WebContentUpdateDto) GetPrecompiledLogicSizeLong() int64`

GetPrecompiledLogicSizeLong returns the PrecompiledLogicSizeLong field if non-nil, zero value otherwise.

### GetPrecompiledLogicSizeLongOk

`func (o *WebContentUpdateDto) GetPrecompiledLogicSizeLongOk() (*int64, bool)`

GetPrecompiledLogicSizeLongOk returns a tuple with the PrecompiledLogicSizeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicSizeLong

`func (o *WebContentUpdateDto) SetPrecompiledLogicSizeLong(v int64)`

SetPrecompiledLogicSizeLong sets PrecompiledLogicSizeLong field to given value.

### HasPrecompiledLogicSizeLong

`func (o *WebContentUpdateDto) HasPrecompiledLogicSizeLong() bool`

HasPrecompiledLogicSizeLong returns a boolean if a field has been set.

### GetPrecompiledViewSize

`func (o *WebContentUpdateDto) GetPrecompiledViewSize() int32`

GetPrecompiledViewSize returns the PrecompiledViewSize field if non-nil, zero value otherwise.

### GetPrecompiledViewSizeOk

`func (o *WebContentUpdateDto) GetPrecompiledViewSizeOk() (*int32, bool)`

GetPrecompiledViewSizeOk returns a tuple with the PrecompiledViewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledViewSize

`func (o *WebContentUpdateDto) SetPrecompiledViewSize(v int32)`

SetPrecompiledViewSize sets PrecompiledViewSize field to given value.

### HasPrecompiledViewSize

`func (o *WebContentUpdateDto) HasPrecompiledViewSize() bool`

HasPrecompiledViewSize returns a boolean if a field has been set.

### GetPrecompiledViewSizeLong

`func (o *WebContentUpdateDto) GetPrecompiledViewSizeLong() int64`

GetPrecompiledViewSizeLong returns the PrecompiledViewSizeLong field if non-nil, zero value otherwise.

### GetPrecompiledViewSizeLongOk

`func (o *WebContentUpdateDto) GetPrecompiledViewSizeLongOk() (*int64, bool)`

GetPrecompiledViewSizeLongOk returns a tuple with the PrecompiledViewSizeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledViewSizeLong

`func (o *WebContentUpdateDto) SetPrecompiledViewSizeLong(v int64)`

SetPrecompiledViewSizeLong sets PrecompiledViewSizeLong field to given value.

### HasPrecompiledViewSizeLong

`func (o *WebContentUpdateDto) HasPrecompiledViewSizeLong() bool`

HasPrecompiledViewSizeLong returns a boolean if a field has been set.

### GetPrecompiledLogicViewSize

`func (o *WebContentUpdateDto) GetPrecompiledLogicViewSize() int32`

GetPrecompiledLogicViewSize returns the PrecompiledLogicViewSize field if non-nil, zero value otherwise.

### GetPrecompiledLogicViewSizeOk

`func (o *WebContentUpdateDto) GetPrecompiledLogicViewSizeOk() (*int32, bool)`

GetPrecompiledLogicViewSizeOk returns a tuple with the PrecompiledLogicViewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecompiledLogicViewSize

`func (o *WebContentUpdateDto) SetPrecompiledLogicViewSize(v int32)`

SetPrecompiledLogicViewSize sets PrecompiledLogicViewSize field to given value.

### HasPrecompiledLogicViewSize

`func (o *WebContentUpdateDto) HasPrecompiledLogicViewSize() bool`

HasPrecompiledLogicViewSize returns a boolean if a field has been set.

### GetTemplate

`func (o *WebContentUpdateDto) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebContentUpdateDto) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebContentUpdateDto) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebContentUpdateDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefault

`func (o *WebContentUpdateDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WebContentUpdateDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WebContentUpdateDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WebContentUpdateDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnable

`func (o *WebContentUpdateDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WebContentUpdateDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WebContentUpdateDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *WebContentUpdateDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableComments

`func (o *WebContentUpdateDto) GetEnableComments() bool`

GetEnableComments returns the EnableComments field if non-nil, zero value otherwise.

### GetEnableCommentsOk

`func (o *WebContentUpdateDto) GetEnableCommentsOk() (*bool, bool)`

GetEnableCommentsOk returns a tuple with the EnableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComments

`func (o *WebContentUpdateDto) SetEnableComments(v bool)`

SetEnableComments sets EnableComments field to given value.

### HasEnableComments

`func (o *WebContentUpdateDto) HasEnableComments() bool`

HasEnableComments returns a boolean if a field has been set.

### GetDisplaySocialBox

`func (o *WebContentUpdateDto) GetDisplaySocialBox() bool`

GetDisplaySocialBox returns the DisplaySocialBox field if non-nil, zero value otherwise.

### GetDisplaySocialBoxOk

`func (o *WebContentUpdateDto) GetDisplaySocialBoxOk() (*bool, bool)`

GetDisplaySocialBoxOk returns a tuple with the DisplaySocialBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySocialBox

`func (o *WebContentUpdateDto) SetDisplaySocialBox(v bool)`

SetDisplaySocialBox sets DisplaySocialBox field to given value.

### HasDisplaySocialBox

`func (o *WebContentUpdateDto) HasDisplaySocialBox() bool`

HasDisplaySocialBox returns a boolean if a field has been set.

### GetPublished

`func (o *WebContentUpdateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *WebContentUpdateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *WebContentUpdateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *WebContentUpdateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInTrashCan

`func (o *WebContentUpdateDto) GetInTrashCan() bool`

GetInTrashCan returns the InTrashCan field if non-nil, zero value otherwise.

### GetInTrashCanOk

`func (o *WebContentUpdateDto) GetInTrashCanOk() (*bool, bool)`

GetInTrashCanOk returns a tuple with the InTrashCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrashCan

`func (o *WebContentUpdateDto) SetInTrashCan(v bool)`

SetInTrashCan sets InTrashCan field to given value.

### HasInTrashCan

`func (o *WebContentUpdateDto) HasInTrashCan() bool`

HasInTrashCan returns a boolean if a field has been set.

### GetSystemLocked

`func (o *WebContentUpdateDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *WebContentUpdateDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *WebContentUpdateDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *WebContentUpdateDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetAllowPingbacks

`func (o *WebContentUpdateDto) GetAllowPingbacks() bool`

GetAllowPingbacks returns the AllowPingbacks field if non-nil, zero value otherwise.

### GetAllowPingbacksOk

`func (o *WebContentUpdateDto) GetAllowPingbacksOk() (*bool, bool)`

GetAllowPingbacksOk returns a tuple with the AllowPingbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPingbacks

`func (o *WebContentUpdateDto) SetAllowPingbacks(v bool)`

SetAllowPingbacks sets AllowPingbacks field to given value.

### HasAllowPingbacks

`func (o *WebContentUpdateDto) HasAllowPingbacks() bool`

HasAllowPingbacks returns a boolean if a field has been set.

### GetAllowTrackbacks

`func (o *WebContentUpdateDto) GetAllowTrackbacks() bool`

GetAllowTrackbacks returns the AllowTrackbacks field if non-nil, zero value otherwise.

### GetAllowTrackbacksOk

`func (o *WebContentUpdateDto) GetAllowTrackbacksOk() (*bool, bool)`

GetAllowTrackbacksOk returns a tuple with the AllowTrackbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrackbacks

`func (o *WebContentUpdateDto) SetAllowTrackbacks(v bool)`

SetAllowTrackbacks sets AllowTrackbacks field to given value.

### HasAllowTrackbacks

`func (o *WebContentUpdateDto) HasAllowTrackbacks() bool`

HasAllowTrackbacks returns a boolean if a field has been set.

### GetCornerstoneContent

`func (o *WebContentUpdateDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebContentUpdateDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebContentUpdateDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebContentUpdateDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetIsEssentialContent

`func (o *WebContentUpdateDto) GetIsEssentialContent() bool`

GetIsEssentialContent returns the IsEssentialContent field if non-nil, zero value otherwise.

### GetIsEssentialContentOk

`func (o *WebContentUpdateDto) GetIsEssentialContentOk() (*bool, bool)`

GetIsEssentialContentOk returns a tuple with the IsEssentialContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEssentialContent

`func (o *WebContentUpdateDto) SetIsEssentialContent(v bool)`

SetIsEssentialContent sets IsEssentialContent field to given value.

### HasIsEssentialContent

`func (o *WebContentUpdateDto) HasIsEssentialContent() bool`

HasIsEssentialContent returns a boolean if a field has been set.

### GetAllowSearchEngineIndexing

`func (o *WebContentUpdateDto) GetAllowSearchEngineIndexing() bool`

GetAllowSearchEngineIndexing returns the AllowSearchEngineIndexing field if non-nil, zero value otherwise.

### GetAllowSearchEngineIndexingOk

`func (o *WebContentUpdateDto) GetAllowSearchEngineIndexingOk() (*bool, bool)`

GetAllowSearchEngineIndexingOk returns a tuple with the AllowSearchEngineIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearchEngineIndexing

`func (o *WebContentUpdateDto) SetAllowSearchEngineIndexing(v bool)`

SetAllowSearchEngineIndexing sets AllowSearchEngineIndexing field to given value.

### HasAllowSearchEngineIndexing

`func (o *WebContentUpdateDto) HasAllowSearchEngineIndexing() bool`

HasAllowSearchEngineIndexing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


