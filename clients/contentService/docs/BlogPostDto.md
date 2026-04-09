# BlogPostDto

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
**BlogPostCategoryID** | Pointer to **NullableString** |  | [optional] 
**WebTemplateID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlogPostDto

`func NewBlogPostDto() *BlogPostDto`

NewBlogPostDto instantiates a new BlogPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostDtoWithDefaults

`func NewBlogPostDtoWithDefaults() *BlogPostDto`

NewBlogPostDtoWithDefaults instantiates a new BlogPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BlogPostDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BlogPostDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BlogPostDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlogPostDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlogPostDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlogPostDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BlogPostDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BlogPostDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCodeType

`func (o *BlogPostDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *BlogPostDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *BlogPostDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *BlogPostDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *BlogPostDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *BlogPostDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetOrder

`func (o *BlogPostDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BlogPostDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BlogPostDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BlogPostDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSlug

`func (o *BlogPostDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BlogPostDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BlogPostDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BlogPostDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *BlogPostDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *BlogPostDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *BlogPostDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlogPostDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlogPostDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlogPostDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BlogPostDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BlogPostDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *BlogPostDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPostDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BlogPostDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BlogPostDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExcerpt

`func (o *BlogPostDto) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *BlogPostDto) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *BlogPostDto) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *BlogPostDto) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *BlogPostDto) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *BlogPostDto) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetPassword

`func (o *BlogPostDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BlogPostDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BlogPostDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BlogPostDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *BlogPostDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *BlogPostDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *BlogPostDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlogPostDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlogPostDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlogPostDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BlogPostDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BlogPostDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHighlightImage

`func (o *BlogPostDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *BlogPostDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *BlogPostDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *BlogPostDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *BlogPostDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *BlogPostDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetCanonicalUrl

`func (o *BlogPostDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *BlogPostDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *BlogPostDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *BlogPostDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *BlogPostDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *BlogPostDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetEmitResult

`func (o *BlogPostDto) GetEmitResult() interface{}`

GetEmitResult returns the EmitResult field if non-nil, zero value otherwise.

### GetEmitResultOk

`func (o *BlogPostDto) GetEmitResultOk() (*interface{}, bool)`

GetEmitResultOk returns a tuple with the EmitResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitResult

`func (o *BlogPostDto) SetEmitResult(v interface{})`

SetEmitResult sets EmitResult field to given value.

### HasEmitResult

`func (o *BlogPostDto) HasEmitResult() bool`

HasEmitResult returns a boolean if a field has been set.

### SetEmitResultNil

`func (o *BlogPostDto) SetEmitResultNil(b bool)`

 SetEmitResultNil sets the value for EmitResult to be an explicit nil

### UnsetEmitResult
`func (o *BlogPostDto) UnsetEmitResult()`

UnsetEmitResult ensures that no value is present for EmitResult, not even an explicit nil
### GetSeoTitle

`func (o *BlogPostDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *BlogPostDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *BlogPostDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *BlogPostDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *BlogPostDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *BlogPostDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoKeyWords

`func (o *BlogPostDto) GetSeoKeyWords() string`

GetSeoKeyWords returns the SeoKeyWords field if non-nil, zero value otherwise.

### GetSeoKeyWordsOk

`func (o *BlogPostDto) GetSeoKeyWordsOk() (*string, bool)`

GetSeoKeyWordsOk returns a tuple with the SeoKeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyWords

`func (o *BlogPostDto) SetSeoKeyWords(v string)`

SetSeoKeyWords sets SeoKeyWords field to given value.

### HasSeoKeyWords

`func (o *BlogPostDto) HasSeoKeyWords() bool`

HasSeoKeyWords returns a boolean if a field has been set.

### SetSeoKeyWordsNil

`func (o *BlogPostDto) SetSeoKeyWordsNil(b bool)`

 SetSeoKeyWordsNil sets the value for SeoKeyWords to be an explicit nil

### UnsetSeoKeyWords
`func (o *BlogPostDto) UnsetSeoKeyWords()`

UnsetSeoKeyWords ensures that no value is present for SeoKeyWords, not even an explicit nil
### GetSeoKeyPhrases

`func (o *BlogPostDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *BlogPostDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *BlogPostDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *BlogPostDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *BlogPostDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *BlogPostDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetMetaDescription

`func (o *BlogPostDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *BlogPostDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *BlogPostDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *BlogPostDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *BlogPostDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *BlogPostDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetTwitterImage

`func (o *BlogPostDto) GetTwitterImage() string`

GetTwitterImage returns the TwitterImage field if non-nil, zero value otherwise.

### GetTwitterImageOk

`func (o *BlogPostDto) GetTwitterImageOk() (*string, bool)`

GetTwitterImageOk returns a tuple with the TwitterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterImage

`func (o *BlogPostDto) SetTwitterImage(v string)`

SetTwitterImage sets TwitterImage field to given value.

### HasTwitterImage

`func (o *BlogPostDto) HasTwitterImage() bool`

HasTwitterImage returns a boolean if a field has been set.

### SetTwitterImageNil

`func (o *BlogPostDto) SetTwitterImageNil(b bool)`

 SetTwitterImageNil sets the value for TwitterImage to be an explicit nil

### UnsetTwitterImage
`func (o *BlogPostDto) UnsetTwitterImage()`

UnsetTwitterImage ensures that no value is present for TwitterImage, not even an explicit nil
### GetTwitterTitle

`func (o *BlogPostDto) GetTwitterTitle() string`

GetTwitterTitle returns the TwitterTitle field if non-nil, zero value otherwise.

### GetTwitterTitleOk

`func (o *BlogPostDto) GetTwitterTitleOk() (*string, bool)`

GetTwitterTitleOk returns a tuple with the TwitterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterTitle

`func (o *BlogPostDto) SetTwitterTitle(v string)`

SetTwitterTitle sets TwitterTitle field to given value.

### HasTwitterTitle

`func (o *BlogPostDto) HasTwitterTitle() bool`

HasTwitterTitle returns a boolean if a field has been set.

### SetTwitterTitleNil

`func (o *BlogPostDto) SetTwitterTitleNil(b bool)`

 SetTwitterTitleNil sets the value for TwitterTitle to be an explicit nil

### UnsetTwitterTitle
`func (o *BlogPostDto) UnsetTwitterTitle()`

UnsetTwitterTitle ensures that no value is present for TwitterTitle, not even an explicit nil
### GetTwitterDescription

`func (o *BlogPostDto) GetTwitterDescription() string`

GetTwitterDescription returns the TwitterDescription field if non-nil, zero value otherwise.

### GetTwitterDescriptionOk

`func (o *BlogPostDto) GetTwitterDescriptionOk() (*string, bool)`

GetTwitterDescriptionOk returns a tuple with the TwitterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterDescription

`func (o *BlogPostDto) SetTwitterDescription(v string)`

SetTwitterDescription sets TwitterDescription field to given value.

### HasTwitterDescription

`func (o *BlogPostDto) HasTwitterDescription() bool`

HasTwitterDescription returns a boolean if a field has been set.

### SetTwitterDescriptionNil

`func (o *BlogPostDto) SetTwitterDescriptionNil(b bool)`

 SetTwitterDescriptionNil sets the value for TwitterDescription to be an explicit nil

### UnsetTwitterDescription
`func (o *BlogPostDto) UnsetTwitterDescription()`

UnsetTwitterDescription ensures that no value is present for TwitterDescription, not even an explicit nil
### GetFacebookImage

`func (o *BlogPostDto) GetFacebookImage() string`

GetFacebookImage returns the FacebookImage field if non-nil, zero value otherwise.

### GetFacebookImageOk

`func (o *BlogPostDto) GetFacebookImageOk() (*string, bool)`

GetFacebookImageOk returns a tuple with the FacebookImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookImage

`func (o *BlogPostDto) SetFacebookImage(v string)`

SetFacebookImage sets FacebookImage field to given value.

### HasFacebookImage

`func (o *BlogPostDto) HasFacebookImage() bool`

HasFacebookImage returns a boolean if a field has been set.

### SetFacebookImageNil

`func (o *BlogPostDto) SetFacebookImageNil(b bool)`

 SetFacebookImageNil sets the value for FacebookImage to be an explicit nil

### UnsetFacebookImage
`func (o *BlogPostDto) UnsetFacebookImage()`

UnsetFacebookImage ensures that no value is present for FacebookImage, not even an explicit nil
### GetFacebookTitle

`func (o *BlogPostDto) GetFacebookTitle() string`

GetFacebookTitle returns the FacebookTitle field if non-nil, zero value otherwise.

### GetFacebookTitleOk

`func (o *BlogPostDto) GetFacebookTitleOk() (*string, bool)`

GetFacebookTitleOk returns a tuple with the FacebookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookTitle

`func (o *BlogPostDto) SetFacebookTitle(v string)`

SetFacebookTitle sets FacebookTitle field to given value.

### HasFacebookTitle

`func (o *BlogPostDto) HasFacebookTitle() bool`

HasFacebookTitle returns a boolean if a field has been set.

### SetFacebookTitleNil

`func (o *BlogPostDto) SetFacebookTitleNil(b bool)`

 SetFacebookTitleNil sets the value for FacebookTitle to be an explicit nil

### UnsetFacebookTitle
`func (o *BlogPostDto) UnsetFacebookTitle()`

UnsetFacebookTitle ensures that no value is present for FacebookTitle, not even an explicit nil
### GetFacebookDescription

`func (o *BlogPostDto) GetFacebookDescription() string`

GetFacebookDescription returns the FacebookDescription field if non-nil, zero value otherwise.

### GetFacebookDescriptionOk

`func (o *BlogPostDto) GetFacebookDescriptionOk() (*string, bool)`

GetFacebookDescriptionOk returns a tuple with the FacebookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookDescription

`func (o *BlogPostDto) SetFacebookDescription(v string)`

SetFacebookDescription sets FacebookDescription field to given value.

### HasFacebookDescription

`func (o *BlogPostDto) HasFacebookDescription() bool`

HasFacebookDescription returns a boolean if a field has been set.

### SetFacebookDescriptionNil

`func (o *BlogPostDto) SetFacebookDescriptionNil(b bool)`

 SetFacebookDescriptionNil sets the value for FacebookDescription to be an explicit nil

### UnsetFacebookDescription
`func (o *BlogPostDto) UnsetFacebookDescription()`

UnsetFacebookDescription ensures that no value is present for FacebookDescription, not even an explicit nil
### GetFeaturedImageUrl

`func (o *BlogPostDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *BlogPostDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *BlogPostDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *BlogPostDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *BlogPostDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *BlogPostDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetContent

`func (o *BlogPostDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BlogPostDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BlogPostDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BlogPostDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *BlogPostDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *BlogPostDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCode

`func (o *BlogPostDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BlogPostDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BlogPostDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BlogPostDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BlogPostDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BlogPostDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNamespace

`func (o *BlogPostDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BlogPostDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BlogPostDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BlogPostDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *BlogPostDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *BlogPostDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTypeName

`func (o *BlogPostDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BlogPostDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BlogPostDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BlogPostDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *BlogPostDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *BlogPostDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetGeneratedCode

`func (o *BlogPostDto) GetGeneratedCode() string`

GetGeneratedCode returns the GeneratedCode field if non-nil, zero value otherwise.

### GetGeneratedCodeOk

`func (o *BlogPostDto) GetGeneratedCodeOk() (*string, bool)`

GetGeneratedCodeOk returns a tuple with the GeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCode

`func (o *BlogPostDto) SetGeneratedCode(v string)`

SetGeneratedCode sets GeneratedCode field to given value.

### HasGeneratedCode

`func (o *BlogPostDto) HasGeneratedCode() bool`

HasGeneratedCode returns a boolean if a field has been set.

### SetGeneratedCodeNil

`func (o *BlogPostDto) SetGeneratedCodeNil(b bool)`

 SetGeneratedCodeNil sets the value for GeneratedCode to be an explicit nil

### UnsetGeneratedCode
`func (o *BlogPostDto) UnsetGeneratedCode()`

UnsetGeneratedCode ensures that no value is present for GeneratedCode, not even an explicit nil
### GetCompilationPath

`func (o *BlogPostDto) GetCompilationPath() string`

GetCompilationPath returns the CompilationPath field if non-nil, zero value otherwise.

### GetCompilationPathOk

`func (o *BlogPostDto) GetCompilationPathOk() (*string, bool)`

GetCompilationPathOk returns a tuple with the CompilationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilationPath

`func (o *BlogPostDto) SetCompilationPath(v string)`

SetCompilationPath sets CompilationPath field to given value.

### HasCompilationPath

`func (o *BlogPostDto) HasCompilationPath() bool`

HasCompilationPath returns a boolean if a field has been set.

### SetCompilationPathNil

`func (o *BlogPostDto) SetCompilationPathNil(b bool)`

 SetCompilationPathNil sets the value for CompilationPath to be an explicit nil

### UnsetCompilationPath
`func (o *BlogPostDto) UnsetCompilationPath()`

UnsetCompilationPath ensures that no value is present for CompilationPath, not even an explicit nil
### GetHtmlContent

`func (o *BlogPostDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *BlogPostDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *BlogPostDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *BlogPostDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *BlogPostDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *BlogPostDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCSharpContent

`func (o *BlogPostDto) GetCSharpContent() string`

GetCSharpContent returns the CSharpContent field if non-nil, zero value otherwise.

### GetCSharpContentOk

`func (o *BlogPostDto) GetCSharpContentOk() (*string, bool)`

GetCSharpContentOk returns a tuple with the CSharpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpContent

`func (o *BlogPostDto) SetCSharpContent(v string)`

SetCSharpContent sets CSharpContent field to given value.

### HasCSharpContent

`func (o *BlogPostDto) HasCSharpContent() bool`

HasCSharpContent returns a boolean if a field has been set.

### SetCSharpContentNil

`func (o *BlogPostDto) SetCSharpContentNil(b bool)`

 SetCSharpContentNil sets the value for CSharpContent to be an explicit nil

### UnsetCSharpContent
`func (o *BlogPostDto) UnsetCSharpContent()`

UnsetCSharpContent ensures that no value is present for CSharpContent, not even an explicit nil
### GetRazorContent

`func (o *BlogPostDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *BlogPostDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *BlogPostDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *BlogPostDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *BlogPostDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *BlogPostDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetCssContent

`func (o *BlogPostDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *BlogPostDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *BlogPostDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *BlogPostDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *BlogPostDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *BlogPostDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *BlogPostDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *BlogPostDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *BlogPostDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *BlogPostDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *BlogPostDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *BlogPostDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetCssFiles

`func (o *BlogPostDto) GetCssFiles() string`

GetCssFiles returns the CssFiles field if non-nil, zero value otherwise.

### GetCssFilesOk

`func (o *BlogPostDto) GetCssFilesOk() (*string, bool)`

GetCssFilesOk returns a tuple with the CssFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFiles

`func (o *BlogPostDto) SetCssFiles(v string)`

SetCssFiles sets CssFiles field to given value.

### HasCssFiles

`func (o *BlogPostDto) HasCssFiles() bool`

HasCssFiles returns a boolean if a field has been set.

### SetCssFilesNil

`func (o *BlogPostDto) SetCssFilesNil(b bool)`

 SetCssFilesNil sets the value for CssFiles to be an explicit nil

### UnsetCssFiles
`func (o *BlogPostDto) UnsetCssFiles()`

UnsetCssFiles ensures that no value is present for CssFiles, not even an explicit nil
### GetJsFiles

`func (o *BlogPostDto) GetJsFiles() string`

GetJsFiles returns the JsFiles field if non-nil, zero value otherwise.

### GetJsFilesOk

`func (o *BlogPostDto) GetJsFilesOk() (*string, bool)`

GetJsFilesOk returns a tuple with the JsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsFiles

`func (o *BlogPostDto) SetJsFiles(v string)`

SetJsFiles sets JsFiles field to given value.

### HasJsFiles

`func (o *BlogPostDto) HasJsFiles() bool`

HasJsFiles returns a boolean if a field has been set.

### SetJsFilesNil

`func (o *BlogPostDto) SetJsFilesNil(b bool)`

 SetJsFilesNil sets the value for JsFiles to be an explicit nil

### UnsetJsFiles
`func (o *BlogPostDto) UnsetJsFiles()`

UnsetJsFiles ensures that no value is present for JsFiles, not even an explicit nil
### GetRazorGeneratedCode

`func (o *BlogPostDto) GetRazorGeneratedCode() string`

GetRazorGeneratedCode returns the RazorGeneratedCode field if non-nil, zero value otherwise.

### GetRazorGeneratedCodeOk

`func (o *BlogPostDto) GetRazorGeneratedCodeOk() (*string, bool)`

GetRazorGeneratedCodeOk returns a tuple with the RazorGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorGeneratedCode

`func (o *BlogPostDto) SetRazorGeneratedCode(v string)`

SetRazorGeneratedCode sets RazorGeneratedCode field to given value.

### HasRazorGeneratedCode

`func (o *BlogPostDto) HasRazorGeneratedCode() bool`

HasRazorGeneratedCode returns a boolean if a field has been set.

### SetRazorGeneratedCodeNil

`func (o *BlogPostDto) SetRazorGeneratedCodeNil(b bool)`

 SetRazorGeneratedCodeNil sets the value for RazorGeneratedCode to be an explicit nil

### UnsetRazorGeneratedCode
`func (o *BlogPostDto) UnsetRazorGeneratedCode()`

UnsetRazorGeneratedCode ensures that no value is present for RazorGeneratedCode, not even an explicit nil
### GetCSharpGeneratedCode

`func (o *BlogPostDto) GetCSharpGeneratedCode() string`

GetCSharpGeneratedCode returns the CSharpGeneratedCode field if non-nil, zero value otherwise.

### GetCSharpGeneratedCodeOk

`func (o *BlogPostDto) GetCSharpGeneratedCodeOk() (*string, bool)`

GetCSharpGeneratedCodeOk returns a tuple with the CSharpGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpGeneratedCode

`func (o *BlogPostDto) SetCSharpGeneratedCode(v string)`

SetCSharpGeneratedCode sets CSharpGeneratedCode field to given value.

### HasCSharpGeneratedCode

`func (o *BlogPostDto) HasCSharpGeneratedCode() bool`

HasCSharpGeneratedCode returns a boolean if a field has been set.

### SetCSharpGeneratedCodeNil

`func (o *BlogPostDto) SetCSharpGeneratedCodeNil(b bool)`

 SetCSharpGeneratedCodeNil sets the value for CSharpGeneratedCode to be an explicit nil

### UnsetCSharpGeneratedCode
`func (o *BlogPostDto) UnsetCSharpGeneratedCode()`

UnsetCSharpGeneratedCode ensures that no value is present for CSharpGeneratedCode, not even an explicit nil
### GetTemplate

`func (o *BlogPostDto) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *BlogPostDto) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *BlogPostDto) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *BlogPostDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefault

`func (o *BlogPostDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BlogPostDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BlogPostDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *BlogPostDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnable

`func (o *BlogPostDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BlogPostDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BlogPostDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *BlogPostDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableComments

`func (o *BlogPostDto) GetEnableComments() bool`

GetEnableComments returns the EnableComments field if non-nil, zero value otherwise.

### GetEnableCommentsOk

`func (o *BlogPostDto) GetEnableCommentsOk() (*bool, bool)`

GetEnableCommentsOk returns a tuple with the EnableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComments

`func (o *BlogPostDto) SetEnableComments(v bool)`

SetEnableComments sets EnableComments field to given value.

### HasEnableComments

`func (o *BlogPostDto) HasEnableComments() bool`

HasEnableComments returns a boolean if a field has been set.

### GetDisplaySocialBox

`func (o *BlogPostDto) GetDisplaySocialBox() bool`

GetDisplaySocialBox returns the DisplaySocialBox field if non-nil, zero value otherwise.

### GetDisplaySocialBoxOk

`func (o *BlogPostDto) GetDisplaySocialBoxOk() (*bool, bool)`

GetDisplaySocialBoxOk returns a tuple with the DisplaySocialBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySocialBox

`func (o *BlogPostDto) SetDisplaySocialBox(v bool)`

SetDisplaySocialBox sets DisplaySocialBox field to given value.

### HasDisplaySocialBox

`func (o *BlogPostDto) HasDisplaySocialBox() bool`

HasDisplaySocialBox returns a boolean if a field has been set.

### GetPublished

`func (o *BlogPostDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BlogPostDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BlogPostDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BlogPostDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInTrashCan

`func (o *BlogPostDto) GetInTrashCan() bool`

GetInTrashCan returns the InTrashCan field if non-nil, zero value otherwise.

### GetInTrashCanOk

`func (o *BlogPostDto) GetInTrashCanOk() (*bool, bool)`

GetInTrashCanOk returns a tuple with the InTrashCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrashCan

`func (o *BlogPostDto) SetInTrashCan(v bool)`

SetInTrashCan sets InTrashCan field to given value.

### HasInTrashCan

`func (o *BlogPostDto) HasInTrashCan() bool`

HasInTrashCan returns a boolean if a field has been set.

### GetSystemLocked

`func (o *BlogPostDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *BlogPostDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *BlogPostDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *BlogPostDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetAllowPingBacks

`func (o *BlogPostDto) GetAllowPingBacks() bool`

GetAllowPingBacks returns the AllowPingBacks field if non-nil, zero value otherwise.

### GetAllowPingBacksOk

`func (o *BlogPostDto) GetAllowPingBacksOk() (*bool, bool)`

GetAllowPingBacksOk returns a tuple with the AllowPingBacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPingBacks

`func (o *BlogPostDto) SetAllowPingBacks(v bool)`

SetAllowPingBacks sets AllowPingBacks field to given value.

### HasAllowPingBacks

`func (o *BlogPostDto) HasAllowPingBacks() bool`

HasAllowPingBacks returns a boolean if a field has been set.

### GetAllowTrackbacks

`func (o *BlogPostDto) GetAllowTrackbacks() bool`

GetAllowTrackbacks returns the AllowTrackbacks field if non-nil, zero value otherwise.

### GetAllowTrackbacksOk

`func (o *BlogPostDto) GetAllowTrackbacksOk() (*bool, bool)`

GetAllowTrackbacksOk returns a tuple with the AllowTrackbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrackbacks

`func (o *BlogPostDto) SetAllowTrackbacks(v bool)`

SetAllowTrackbacks sets AllowTrackbacks field to given value.

### HasAllowTrackbacks

`func (o *BlogPostDto) HasAllowTrackbacks() bool`

HasAllowTrackbacks returns a boolean if a field has been set.

### GetCornerstoneContent

`func (o *BlogPostDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *BlogPostDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *BlogPostDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *BlogPostDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetIsEssentialContent

`func (o *BlogPostDto) GetIsEssentialContent() bool`

GetIsEssentialContent returns the IsEssentialContent field if non-nil, zero value otherwise.

### GetIsEssentialContentOk

`func (o *BlogPostDto) GetIsEssentialContentOk() (*bool, bool)`

GetIsEssentialContentOk returns a tuple with the IsEssentialContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEssentialContent

`func (o *BlogPostDto) SetIsEssentialContent(v bool)`

SetIsEssentialContent sets IsEssentialContent field to given value.

### HasIsEssentialContent

`func (o *BlogPostDto) HasIsEssentialContent() bool`

HasIsEssentialContent returns a boolean if a field has been set.

### GetAllowSearchEngineIndexing

`func (o *BlogPostDto) GetAllowSearchEngineIndexing() bool`

GetAllowSearchEngineIndexing returns the AllowSearchEngineIndexing field if non-nil, zero value otherwise.

### GetAllowSearchEngineIndexingOk

`func (o *BlogPostDto) GetAllowSearchEngineIndexingOk() (*bool, bool)`

GetAllowSearchEngineIndexingOk returns a tuple with the AllowSearchEngineIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearchEngineIndexing

`func (o *BlogPostDto) SetAllowSearchEngineIndexing(v bool)`

SetAllowSearchEngineIndexing sets AllowSearchEngineIndexing field to given value.

### HasAllowSearchEngineIndexing

`func (o *BlogPostDto) HasAllowSearchEngineIndexing() bool`

HasAllowSearchEngineIndexing returns a boolean if a field has been set.

### GetTenantId

`func (o *BlogPostDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BlogPostDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BlogPostDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BlogPostDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BlogPostDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BlogPostDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWebPortalId

`func (o *BlogPostDto) GetWebPortalId() string`

GetWebPortalId returns the WebPortalId field if non-nil, zero value otherwise.

### GetWebPortalIdOk

`func (o *BlogPostDto) GetWebPortalIdOk() (*string, bool)`

GetWebPortalIdOk returns a tuple with the WebPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalId

`func (o *BlogPostDto) SetWebPortalId(v string)`

SetWebPortalId sets WebPortalId field to given value.

### HasWebPortalId

`func (o *BlogPostDto) HasWebPortalId() bool`

HasWebPortalId returns a boolean if a field has been set.

### SetWebPortalIdNil

`func (o *BlogPostDto) SetWebPortalIdNil(b bool)`

 SetWebPortalIdNil sets the value for WebPortalId to be an explicit nil

### UnsetWebPortalId
`func (o *BlogPostDto) UnsetWebPortalId()`

UnsetWebPortalId ensures that no value is present for WebPortalId, not even an explicit nil
### GetWebsiteThemeId

`func (o *BlogPostDto) GetWebsiteThemeId() string`

GetWebsiteThemeId returns the WebsiteThemeId field if non-nil, zero value otherwise.

### GetWebsiteThemeIdOk

`func (o *BlogPostDto) GetWebsiteThemeIdOk() (*string, bool)`

GetWebsiteThemeIdOk returns a tuple with the WebsiteThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeId

`func (o *BlogPostDto) SetWebsiteThemeId(v string)`

SetWebsiteThemeId sets WebsiteThemeId field to given value.

### HasWebsiteThemeId

`func (o *BlogPostDto) HasWebsiteThemeId() bool`

HasWebsiteThemeId returns a boolean if a field has been set.

### SetWebsiteThemeIdNil

`func (o *BlogPostDto) SetWebsiteThemeIdNil(b bool)`

 SetWebsiteThemeIdNil sets the value for WebsiteThemeId to be an explicit nil

### UnsetWebsiteThemeId
`func (o *BlogPostDto) UnsetWebsiteThemeId()`

UnsetWebsiteThemeId ensures that no value is present for WebsiteThemeId, not even an explicit nil
### GetEnrollmentId

`func (o *BlogPostDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BlogPostDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BlogPostDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BlogPostDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BlogPostDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BlogPostDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *BlogPostDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *BlogPostDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *BlogPostDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *BlogPostDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *BlogPostDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *BlogPostDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetParentWebContentId

`func (o *BlogPostDto) GetParentWebContentId() string`

GetParentWebContentId returns the ParentWebContentId field if non-nil, zero value otherwise.

### GetParentWebContentIdOk

`func (o *BlogPostDto) GetParentWebContentIdOk() (*string, bool)`

GetParentWebContentIdOk returns a tuple with the ParentWebContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentId

`func (o *BlogPostDto) SetParentWebContentId(v string)`

SetParentWebContentId sets ParentWebContentId field to given value.

### HasParentWebContentId

`func (o *BlogPostDto) HasParentWebContentId() bool`

HasParentWebContentId returns a boolean if a field has been set.

### SetParentWebContentIdNil

`func (o *BlogPostDto) SetParentWebContentIdNil(b bool)`

 SetParentWebContentIdNil sets the value for ParentWebContentId to be an explicit nil

### UnsetParentWebContentId
`func (o *BlogPostDto) UnsetParentWebContentId()`

UnsetParentWebContentId ensures that no value is present for ParentWebContentId, not even an explicit nil
### GetParentWebContentVersionId

`func (o *BlogPostDto) GetParentWebContentVersionId() string`

GetParentWebContentVersionId returns the ParentWebContentVersionId field if non-nil, zero value otherwise.

### GetParentWebContentVersionIdOk

`func (o *BlogPostDto) GetParentWebContentVersionIdOk() (*string, bool)`

GetParentWebContentVersionIdOk returns a tuple with the ParentWebContentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentVersionId

`func (o *BlogPostDto) SetParentWebContentVersionId(v string)`

SetParentWebContentVersionId sets ParentWebContentVersionId field to given value.

### HasParentWebContentVersionId

`func (o *BlogPostDto) HasParentWebContentVersionId() bool`

HasParentWebContentVersionId returns a boolean if a field has been set.

### SetParentWebContentVersionIdNil

`func (o *BlogPostDto) SetParentWebContentVersionIdNil(b bool)`

 SetParentWebContentVersionIdNil sets the value for ParentWebContentVersionId to be an explicit nil

### UnsetParentWebContentVersionId
`func (o *BlogPostDto) UnsetParentWebContentVersionId()`

UnsetParentWebContentVersionId ensures that no value is present for ParentWebContentVersionId, not even an explicit nil
### GetBlogPostCategoryID

`func (o *BlogPostDto) GetBlogPostCategoryID() string`

GetBlogPostCategoryID returns the BlogPostCategoryID field if non-nil, zero value otherwise.

### GetBlogPostCategoryIDOk

`func (o *BlogPostDto) GetBlogPostCategoryIDOk() (*string, bool)`

GetBlogPostCategoryIDOk returns a tuple with the BlogPostCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostCategoryID

`func (o *BlogPostDto) SetBlogPostCategoryID(v string)`

SetBlogPostCategoryID sets BlogPostCategoryID field to given value.

### HasBlogPostCategoryID

`func (o *BlogPostDto) HasBlogPostCategoryID() bool`

HasBlogPostCategoryID returns a boolean if a field has been set.

### SetBlogPostCategoryIDNil

`func (o *BlogPostDto) SetBlogPostCategoryIDNil(b bool)`

 SetBlogPostCategoryIDNil sets the value for BlogPostCategoryID to be an explicit nil

### UnsetBlogPostCategoryID
`func (o *BlogPostDto) UnsetBlogPostCategoryID()`

UnsetBlogPostCategoryID ensures that no value is present for BlogPostCategoryID, not even an explicit nil
### GetWebTemplateID

`func (o *BlogPostDto) GetWebTemplateID() string`

GetWebTemplateID returns the WebTemplateID field if non-nil, zero value otherwise.

### GetWebTemplateIDOk

`func (o *BlogPostDto) GetWebTemplateIDOk() (*string, bool)`

GetWebTemplateIDOk returns a tuple with the WebTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTemplateID

`func (o *BlogPostDto) SetWebTemplateID(v string)`

SetWebTemplateID sets WebTemplateID field to given value.

### HasWebTemplateID

`func (o *BlogPostDto) HasWebTemplateID() bool`

HasWebTemplateID returns a boolean if a field has been set.

### SetWebTemplateIDNil

`func (o *BlogPostDto) SetWebTemplateIDNil(b bool)`

 SetWebTemplateIDNil sets the value for WebTemplateID to be an explicit nil

### UnsetWebTemplateID
`func (o *BlogPostDto) UnsetWebTemplateID()`

UnsetWebTemplateID ensures that no value is present for WebTemplateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


