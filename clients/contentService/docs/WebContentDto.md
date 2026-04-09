# WebContentDto

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

## Methods

### NewWebContentDto

`func NewWebContentDto() *WebContentDto`

NewWebContentDto instantiates a new WebContentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebContentDtoWithDefaults

`func NewWebContentDtoWithDefaults() *WebContentDto`

NewWebContentDtoWithDefaults instantiates a new WebContentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebContentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebContentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebContentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebContentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebContentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebContentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WebContentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebContentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebContentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebContentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WebContentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WebContentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCodeType

`func (o *WebContentDto) GetCodeType() string`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *WebContentDto) GetCodeTypeOk() (*string, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *WebContentDto) SetCodeType(v string)`

SetCodeType sets CodeType field to given value.

### HasCodeType

`func (o *WebContentDto) HasCodeType() bool`

HasCodeType returns a boolean if a field has been set.

### SetCodeTypeNil

`func (o *WebContentDto) SetCodeTypeNil(b bool)`

 SetCodeTypeNil sets the value for CodeType to be an explicit nil

### UnsetCodeType
`func (o *WebContentDto) UnsetCodeType()`

UnsetCodeType ensures that no value is present for CodeType, not even an explicit nil
### GetOrder

`func (o *WebContentDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebContentDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebContentDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebContentDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSlug

`func (o *WebContentDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebContentDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebContentDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebContentDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebContentDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebContentDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *WebContentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebContentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebContentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebContentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebContentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebContentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *WebContentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebContentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebContentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebContentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebContentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebContentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExcerpt

`func (o *WebContentDto) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *WebContentDto) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *WebContentDto) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *WebContentDto) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *WebContentDto) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *WebContentDto) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetPassword

`func (o *WebContentDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebContentDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebContentDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebContentDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebContentDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebContentDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *WebContentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebContentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebContentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebContentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebContentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebContentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHighlightImage

`func (o *WebContentDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *WebContentDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *WebContentDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *WebContentDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *WebContentDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *WebContentDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetCanonicalUrl

`func (o *WebContentDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebContentDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebContentDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebContentDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebContentDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebContentDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetEmitResult

`func (o *WebContentDto) GetEmitResult() interface{}`

GetEmitResult returns the EmitResult field if non-nil, zero value otherwise.

### GetEmitResultOk

`func (o *WebContentDto) GetEmitResultOk() (*interface{}, bool)`

GetEmitResultOk returns a tuple with the EmitResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitResult

`func (o *WebContentDto) SetEmitResult(v interface{})`

SetEmitResult sets EmitResult field to given value.

### HasEmitResult

`func (o *WebContentDto) HasEmitResult() bool`

HasEmitResult returns a boolean if a field has been set.

### SetEmitResultNil

`func (o *WebContentDto) SetEmitResultNil(b bool)`

 SetEmitResultNil sets the value for EmitResult to be an explicit nil

### UnsetEmitResult
`func (o *WebContentDto) UnsetEmitResult()`

UnsetEmitResult ensures that no value is present for EmitResult, not even an explicit nil
### GetSeoTitle

`func (o *WebContentDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebContentDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebContentDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebContentDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebContentDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebContentDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoKeyWords

`func (o *WebContentDto) GetSeoKeyWords() string`

GetSeoKeyWords returns the SeoKeyWords field if non-nil, zero value otherwise.

### GetSeoKeyWordsOk

`func (o *WebContentDto) GetSeoKeyWordsOk() (*string, bool)`

GetSeoKeyWordsOk returns a tuple with the SeoKeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyWords

`func (o *WebContentDto) SetSeoKeyWords(v string)`

SetSeoKeyWords sets SeoKeyWords field to given value.

### HasSeoKeyWords

`func (o *WebContentDto) HasSeoKeyWords() bool`

HasSeoKeyWords returns a boolean if a field has been set.

### SetSeoKeyWordsNil

`func (o *WebContentDto) SetSeoKeyWordsNil(b bool)`

 SetSeoKeyWordsNil sets the value for SeoKeyWords to be an explicit nil

### UnsetSeoKeyWords
`func (o *WebContentDto) UnsetSeoKeyWords()`

UnsetSeoKeyWords ensures that no value is present for SeoKeyWords, not even an explicit nil
### GetSeoKeyPhrases

`func (o *WebContentDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebContentDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebContentDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebContentDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebContentDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebContentDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetMetaDescription

`func (o *WebContentDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebContentDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebContentDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebContentDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebContentDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebContentDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetTwitterImage

`func (o *WebContentDto) GetTwitterImage() string`

GetTwitterImage returns the TwitterImage field if non-nil, zero value otherwise.

### GetTwitterImageOk

`func (o *WebContentDto) GetTwitterImageOk() (*string, bool)`

GetTwitterImageOk returns a tuple with the TwitterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterImage

`func (o *WebContentDto) SetTwitterImage(v string)`

SetTwitterImage sets TwitterImage field to given value.

### HasTwitterImage

`func (o *WebContentDto) HasTwitterImage() bool`

HasTwitterImage returns a boolean if a field has been set.

### SetTwitterImageNil

`func (o *WebContentDto) SetTwitterImageNil(b bool)`

 SetTwitterImageNil sets the value for TwitterImage to be an explicit nil

### UnsetTwitterImage
`func (o *WebContentDto) UnsetTwitterImage()`

UnsetTwitterImage ensures that no value is present for TwitterImage, not even an explicit nil
### GetTwitterTitle

`func (o *WebContentDto) GetTwitterTitle() string`

GetTwitterTitle returns the TwitterTitle field if non-nil, zero value otherwise.

### GetTwitterTitleOk

`func (o *WebContentDto) GetTwitterTitleOk() (*string, bool)`

GetTwitterTitleOk returns a tuple with the TwitterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterTitle

`func (o *WebContentDto) SetTwitterTitle(v string)`

SetTwitterTitle sets TwitterTitle field to given value.

### HasTwitterTitle

`func (o *WebContentDto) HasTwitterTitle() bool`

HasTwitterTitle returns a boolean if a field has been set.

### SetTwitterTitleNil

`func (o *WebContentDto) SetTwitterTitleNil(b bool)`

 SetTwitterTitleNil sets the value for TwitterTitle to be an explicit nil

### UnsetTwitterTitle
`func (o *WebContentDto) UnsetTwitterTitle()`

UnsetTwitterTitle ensures that no value is present for TwitterTitle, not even an explicit nil
### GetTwitterDescription

`func (o *WebContentDto) GetTwitterDescription() string`

GetTwitterDescription returns the TwitterDescription field if non-nil, zero value otherwise.

### GetTwitterDescriptionOk

`func (o *WebContentDto) GetTwitterDescriptionOk() (*string, bool)`

GetTwitterDescriptionOk returns a tuple with the TwitterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterDescription

`func (o *WebContentDto) SetTwitterDescription(v string)`

SetTwitterDescription sets TwitterDescription field to given value.

### HasTwitterDescription

`func (o *WebContentDto) HasTwitterDescription() bool`

HasTwitterDescription returns a boolean if a field has been set.

### SetTwitterDescriptionNil

`func (o *WebContentDto) SetTwitterDescriptionNil(b bool)`

 SetTwitterDescriptionNil sets the value for TwitterDescription to be an explicit nil

### UnsetTwitterDescription
`func (o *WebContentDto) UnsetTwitterDescription()`

UnsetTwitterDescription ensures that no value is present for TwitterDescription, not even an explicit nil
### GetFacebookImage

`func (o *WebContentDto) GetFacebookImage() string`

GetFacebookImage returns the FacebookImage field if non-nil, zero value otherwise.

### GetFacebookImageOk

`func (o *WebContentDto) GetFacebookImageOk() (*string, bool)`

GetFacebookImageOk returns a tuple with the FacebookImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookImage

`func (o *WebContentDto) SetFacebookImage(v string)`

SetFacebookImage sets FacebookImage field to given value.

### HasFacebookImage

`func (o *WebContentDto) HasFacebookImage() bool`

HasFacebookImage returns a boolean if a field has been set.

### SetFacebookImageNil

`func (o *WebContentDto) SetFacebookImageNil(b bool)`

 SetFacebookImageNil sets the value for FacebookImage to be an explicit nil

### UnsetFacebookImage
`func (o *WebContentDto) UnsetFacebookImage()`

UnsetFacebookImage ensures that no value is present for FacebookImage, not even an explicit nil
### GetFacebookTitle

`func (o *WebContentDto) GetFacebookTitle() string`

GetFacebookTitle returns the FacebookTitle field if non-nil, zero value otherwise.

### GetFacebookTitleOk

`func (o *WebContentDto) GetFacebookTitleOk() (*string, bool)`

GetFacebookTitleOk returns a tuple with the FacebookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookTitle

`func (o *WebContentDto) SetFacebookTitle(v string)`

SetFacebookTitle sets FacebookTitle field to given value.

### HasFacebookTitle

`func (o *WebContentDto) HasFacebookTitle() bool`

HasFacebookTitle returns a boolean if a field has been set.

### SetFacebookTitleNil

`func (o *WebContentDto) SetFacebookTitleNil(b bool)`

 SetFacebookTitleNil sets the value for FacebookTitle to be an explicit nil

### UnsetFacebookTitle
`func (o *WebContentDto) UnsetFacebookTitle()`

UnsetFacebookTitle ensures that no value is present for FacebookTitle, not even an explicit nil
### GetFacebookDescription

`func (o *WebContentDto) GetFacebookDescription() string`

GetFacebookDescription returns the FacebookDescription field if non-nil, zero value otherwise.

### GetFacebookDescriptionOk

`func (o *WebContentDto) GetFacebookDescriptionOk() (*string, bool)`

GetFacebookDescriptionOk returns a tuple with the FacebookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookDescription

`func (o *WebContentDto) SetFacebookDescription(v string)`

SetFacebookDescription sets FacebookDescription field to given value.

### HasFacebookDescription

`func (o *WebContentDto) HasFacebookDescription() bool`

HasFacebookDescription returns a boolean if a field has been set.

### SetFacebookDescriptionNil

`func (o *WebContentDto) SetFacebookDescriptionNil(b bool)`

 SetFacebookDescriptionNil sets the value for FacebookDescription to be an explicit nil

### UnsetFacebookDescription
`func (o *WebContentDto) UnsetFacebookDescription()`

UnsetFacebookDescription ensures that no value is present for FacebookDescription, not even an explicit nil
### GetFeaturedImageUrl

`func (o *WebContentDto) GetFeaturedImageUrl() string`

GetFeaturedImageUrl returns the FeaturedImageUrl field if non-nil, zero value otherwise.

### GetFeaturedImageUrlOk

`func (o *WebContentDto) GetFeaturedImageUrlOk() (*string, bool)`

GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageUrl

`func (o *WebContentDto) SetFeaturedImageUrl(v string)`

SetFeaturedImageUrl sets FeaturedImageUrl field to given value.

### HasFeaturedImageUrl

`func (o *WebContentDto) HasFeaturedImageUrl() bool`

HasFeaturedImageUrl returns a boolean if a field has been set.

### SetFeaturedImageUrlNil

`func (o *WebContentDto) SetFeaturedImageUrlNil(b bool)`

 SetFeaturedImageUrlNil sets the value for FeaturedImageUrl to be an explicit nil

### UnsetFeaturedImageUrl
`func (o *WebContentDto) UnsetFeaturedImageUrl()`

UnsetFeaturedImageUrl ensures that no value is present for FeaturedImageUrl, not even an explicit nil
### GetContent

`func (o *WebContentDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebContentDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebContentDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebContentDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebContentDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebContentDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCode

`func (o *WebContentDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebContentDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebContentDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebContentDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WebContentDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WebContentDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetNamespace

`func (o *WebContentDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WebContentDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WebContentDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WebContentDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *WebContentDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *WebContentDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTypeName

`func (o *WebContentDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *WebContentDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *WebContentDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *WebContentDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *WebContentDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *WebContentDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetGeneratedCode

`func (o *WebContentDto) GetGeneratedCode() string`

GetGeneratedCode returns the GeneratedCode field if non-nil, zero value otherwise.

### GetGeneratedCodeOk

`func (o *WebContentDto) GetGeneratedCodeOk() (*string, bool)`

GetGeneratedCodeOk returns a tuple with the GeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCode

`func (o *WebContentDto) SetGeneratedCode(v string)`

SetGeneratedCode sets GeneratedCode field to given value.

### HasGeneratedCode

`func (o *WebContentDto) HasGeneratedCode() bool`

HasGeneratedCode returns a boolean if a field has been set.

### SetGeneratedCodeNil

`func (o *WebContentDto) SetGeneratedCodeNil(b bool)`

 SetGeneratedCodeNil sets the value for GeneratedCode to be an explicit nil

### UnsetGeneratedCode
`func (o *WebContentDto) UnsetGeneratedCode()`

UnsetGeneratedCode ensures that no value is present for GeneratedCode, not even an explicit nil
### GetCompilationPath

`func (o *WebContentDto) GetCompilationPath() string`

GetCompilationPath returns the CompilationPath field if non-nil, zero value otherwise.

### GetCompilationPathOk

`func (o *WebContentDto) GetCompilationPathOk() (*string, bool)`

GetCompilationPathOk returns a tuple with the CompilationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilationPath

`func (o *WebContentDto) SetCompilationPath(v string)`

SetCompilationPath sets CompilationPath field to given value.

### HasCompilationPath

`func (o *WebContentDto) HasCompilationPath() bool`

HasCompilationPath returns a boolean if a field has been set.

### SetCompilationPathNil

`func (o *WebContentDto) SetCompilationPathNil(b bool)`

 SetCompilationPathNil sets the value for CompilationPath to be an explicit nil

### UnsetCompilationPath
`func (o *WebContentDto) UnsetCompilationPath()`

UnsetCompilationPath ensures that no value is present for CompilationPath, not even an explicit nil
### GetHtmlContent

`func (o *WebContentDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebContentDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebContentDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebContentDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebContentDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebContentDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCSharpContent

`func (o *WebContentDto) GetCSharpContent() string`

GetCSharpContent returns the CSharpContent field if non-nil, zero value otherwise.

### GetCSharpContentOk

`func (o *WebContentDto) GetCSharpContentOk() (*string, bool)`

GetCSharpContentOk returns a tuple with the CSharpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpContent

`func (o *WebContentDto) SetCSharpContent(v string)`

SetCSharpContent sets CSharpContent field to given value.

### HasCSharpContent

`func (o *WebContentDto) HasCSharpContent() bool`

HasCSharpContent returns a boolean if a field has been set.

### SetCSharpContentNil

`func (o *WebContentDto) SetCSharpContentNil(b bool)`

 SetCSharpContentNil sets the value for CSharpContent to be an explicit nil

### UnsetCSharpContent
`func (o *WebContentDto) UnsetCSharpContent()`

UnsetCSharpContent ensures that no value is present for CSharpContent, not even an explicit nil
### GetRazorContent

`func (o *WebContentDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *WebContentDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *WebContentDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *WebContentDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *WebContentDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *WebContentDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetCssContent

`func (o *WebContentDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *WebContentDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *WebContentDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *WebContentDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *WebContentDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *WebContentDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *WebContentDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *WebContentDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *WebContentDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *WebContentDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *WebContentDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *WebContentDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetCssFiles

`func (o *WebContentDto) GetCssFiles() string`

GetCssFiles returns the CssFiles field if non-nil, zero value otherwise.

### GetCssFilesOk

`func (o *WebContentDto) GetCssFilesOk() (*string, bool)`

GetCssFilesOk returns a tuple with the CssFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFiles

`func (o *WebContentDto) SetCssFiles(v string)`

SetCssFiles sets CssFiles field to given value.

### HasCssFiles

`func (o *WebContentDto) HasCssFiles() bool`

HasCssFiles returns a boolean if a field has been set.

### SetCssFilesNil

`func (o *WebContentDto) SetCssFilesNil(b bool)`

 SetCssFilesNil sets the value for CssFiles to be an explicit nil

### UnsetCssFiles
`func (o *WebContentDto) UnsetCssFiles()`

UnsetCssFiles ensures that no value is present for CssFiles, not even an explicit nil
### GetJsFiles

`func (o *WebContentDto) GetJsFiles() string`

GetJsFiles returns the JsFiles field if non-nil, zero value otherwise.

### GetJsFilesOk

`func (o *WebContentDto) GetJsFilesOk() (*string, bool)`

GetJsFilesOk returns a tuple with the JsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsFiles

`func (o *WebContentDto) SetJsFiles(v string)`

SetJsFiles sets JsFiles field to given value.

### HasJsFiles

`func (o *WebContentDto) HasJsFiles() bool`

HasJsFiles returns a boolean if a field has been set.

### SetJsFilesNil

`func (o *WebContentDto) SetJsFilesNil(b bool)`

 SetJsFilesNil sets the value for JsFiles to be an explicit nil

### UnsetJsFiles
`func (o *WebContentDto) UnsetJsFiles()`

UnsetJsFiles ensures that no value is present for JsFiles, not even an explicit nil
### GetRazorGeneratedCode

`func (o *WebContentDto) GetRazorGeneratedCode() string`

GetRazorGeneratedCode returns the RazorGeneratedCode field if non-nil, zero value otherwise.

### GetRazorGeneratedCodeOk

`func (o *WebContentDto) GetRazorGeneratedCodeOk() (*string, bool)`

GetRazorGeneratedCodeOk returns a tuple with the RazorGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorGeneratedCode

`func (o *WebContentDto) SetRazorGeneratedCode(v string)`

SetRazorGeneratedCode sets RazorGeneratedCode field to given value.

### HasRazorGeneratedCode

`func (o *WebContentDto) HasRazorGeneratedCode() bool`

HasRazorGeneratedCode returns a boolean if a field has been set.

### SetRazorGeneratedCodeNil

`func (o *WebContentDto) SetRazorGeneratedCodeNil(b bool)`

 SetRazorGeneratedCodeNil sets the value for RazorGeneratedCode to be an explicit nil

### UnsetRazorGeneratedCode
`func (o *WebContentDto) UnsetRazorGeneratedCode()`

UnsetRazorGeneratedCode ensures that no value is present for RazorGeneratedCode, not even an explicit nil
### GetCSharpGeneratedCode

`func (o *WebContentDto) GetCSharpGeneratedCode() string`

GetCSharpGeneratedCode returns the CSharpGeneratedCode field if non-nil, zero value otherwise.

### GetCSharpGeneratedCodeOk

`func (o *WebContentDto) GetCSharpGeneratedCodeOk() (*string, bool)`

GetCSharpGeneratedCodeOk returns a tuple with the CSharpGeneratedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSharpGeneratedCode

`func (o *WebContentDto) SetCSharpGeneratedCode(v string)`

SetCSharpGeneratedCode sets CSharpGeneratedCode field to given value.

### HasCSharpGeneratedCode

`func (o *WebContentDto) HasCSharpGeneratedCode() bool`

HasCSharpGeneratedCode returns a boolean if a field has been set.

### SetCSharpGeneratedCodeNil

`func (o *WebContentDto) SetCSharpGeneratedCodeNil(b bool)`

 SetCSharpGeneratedCodeNil sets the value for CSharpGeneratedCode to be an explicit nil

### UnsetCSharpGeneratedCode
`func (o *WebContentDto) UnsetCSharpGeneratedCode()`

UnsetCSharpGeneratedCode ensures that no value is present for CSharpGeneratedCode, not even an explicit nil
### GetTemplate

`func (o *WebContentDto) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebContentDto) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebContentDto) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebContentDto) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefault

`func (o *WebContentDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WebContentDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WebContentDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WebContentDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnable

`func (o *WebContentDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WebContentDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WebContentDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *WebContentDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableComments

`func (o *WebContentDto) GetEnableComments() bool`

GetEnableComments returns the EnableComments field if non-nil, zero value otherwise.

### GetEnableCommentsOk

`func (o *WebContentDto) GetEnableCommentsOk() (*bool, bool)`

GetEnableCommentsOk returns a tuple with the EnableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComments

`func (o *WebContentDto) SetEnableComments(v bool)`

SetEnableComments sets EnableComments field to given value.

### HasEnableComments

`func (o *WebContentDto) HasEnableComments() bool`

HasEnableComments returns a boolean if a field has been set.

### GetDisplaySocialBox

`func (o *WebContentDto) GetDisplaySocialBox() bool`

GetDisplaySocialBox returns the DisplaySocialBox field if non-nil, zero value otherwise.

### GetDisplaySocialBoxOk

`func (o *WebContentDto) GetDisplaySocialBoxOk() (*bool, bool)`

GetDisplaySocialBoxOk returns a tuple with the DisplaySocialBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySocialBox

`func (o *WebContentDto) SetDisplaySocialBox(v bool)`

SetDisplaySocialBox sets DisplaySocialBox field to given value.

### HasDisplaySocialBox

`func (o *WebContentDto) HasDisplaySocialBox() bool`

HasDisplaySocialBox returns a boolean if a field has been set.

### GetPublished

`func (o *WebContentDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *WebContentDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *WebContentDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *WebContentDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInTrashCan

`func (o *WebContentDto) GetInTrashCan() bool`

GetInTrashCan returns the InTrashCan field if non-nil, zero value otherwise.

### GetInTrashCanOk

`func (o *WebContentDto) GetInTrashCanOk() (*bool, bool)`

GetInTrashCanOk returns a tuple with the InTrashCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrashCan

`func (o *WebContentDto) SetInTrashCan(v bool)`

SetInTrashCan sets InTrashCan field to given value.

### HasInTrashCan

`func (o *WebContentDto) HasInTrashCan() bool`

HasInTrashCan returns a boolean if a field has been set.

### GetSystemLocked

`func (o *WebContentDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *WebContentDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *WebContentDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *WebContentDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetAllowPingBacks

`func (o *WebContentDto) GetAllowPingBacks() bool`

GetAllowPingBacks returns the AllowPingBacks field if non-nil, zero value otherwise.

### GetAllowPingBacksOk

`func (o *WebContentDto) GetAllowPingBacksOk() (*bool, bool)`

GetAllowPingBacksOk returns a tuple with the AllowPingBacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPingBacks

`func (o *WebContentDto) SetAllowPingBacks(v bool)`

SetAllowPingBacks sets AllowPingBacks field to given value.

### HasAllowPingBacks

`func (o *WebContentDto) HasAllowPingBacks() bool`

HasAllowPingBacks returns a boolean if a field has been set.

### GetAllowTrackbacks

`func (o *WebContentDto) GetAllowTrackbacks() bool`

GetAllowTrackbacks returns the AllowTrackbacks field if non-nil, zero value otherwise.

### GetAllowTrackbacksOk

`func (o *WebContentDto) GetAllowTrackbacksOk() (*bool, bool)`

GetAllowTrackbacksOk returns a tuple with the AllowTrackbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrackbacks

`func (o *WebContentDto) SetAllowTrackbacks(v bool)`

SetAllowTrackbacks sets AllowTrackbacks field to given value.

### HasAllowTrackbacks

`func (o *WebContentDto) HasAllowTrackbacks() bool`

HasAllowTrackbacks returns a boolean if a field has been set.

### GetCornerstoneContent

`func (o *WebContentDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebContentDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebContentDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebContentDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetIsEssentialContent

`func (o *WebContentDto) GetIsEssentialContent() bool`

GetIsEssentialContent returns the IsEssentialContent field if non-nil, zero value otherwise.

### GetIsEssentialContentOk

`func (o *WebContentDto) GetIsEssentialContentOk() (*bool, bool)`

GetIsEssentialContentOk returns a tuple with the IsEssentialContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEssentialContent

`func (o *WebContentDto) SetIsEssentialContent(v bool)`

SetIsEssentialContent sets IsEssentialContent field to given value.

### HasIsEssentialContent

`func (o *WebContentDto) HasIsEssentialContent() bool`

HasIsEssentialContent returns a boolean if a field has been set.

### GetAllowSearchEngineIndexing

`func (o *WebContentDto) GetAllowSearchEngineIndexing() bool`

GetAllowSearchEngineIndexing returns the AllowSearchEngineIndexing field if non-nil, zero value otherwise.

### GetAllowSearchEngineIndexingOk

`func (o *WebContentDto) GetAllowSearchEngineIndexingOk() (*bool, bool)`

GetAllowSearchEngineIndexingOk returns a tuple with the AllowSearchEngineIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearchEngineIndexing

`func (o *WebContentDto) SetAllowSearchEngineIndexing(v bool)`

SetAllowSearchEngineIndexing sets AllowSearchEngineIndexing field to given value.

### HasAllowSearchEngineIndexing

`func (o *WebContentDto) HasAllowSearchEngineIndexing() bool`

HasAllowSearchEngineIndexing returns a boolean if a field has been set.

### GetTenantId

`func (o *WebContentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WebContentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WebContentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WebContentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WebContentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WebContentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWebPortalId

`func (o *WebContentDto) GetWebPortalId() string`

GetWebPortalId returns the WebPortalId field if non-nil, zero value otherwise.

### GetWebPortalIdOk

`func (o *WebContentDto) GetWebPortalIdOk() (*string, bool)`

GetWebPortalIdOk returns a tuple with the WebPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalId

`func (o *WebContentDto) SetWebPortalId(v string)`

SetWebPortalId sets WebPortalId field to given value.

### HasWebPortalId

`func (o *WebContentDto) HasWebPortalId() bool`

HasWebPortalId returns a boolean if a field has been set.

### SetWebPortalIdNil

`func (o *WebContentDto) SetWebPortalIdNil(b bool)`

 SetWebPortalIdNil sets the value for WebPortalId to be an explicit nil

### UnsetWebPortalId
`func (o *WebContentDto) UnsetWebPortalId()`

UnsetWebPortalId ensures that no value is present for WebPortalId, not even an explicit nil
### GetWebsiteThemeId

`func (o *WebContentDto) GetWebsiteThemeId() string`

GetWebsiteThemeId returns the WebsiteThemeId field if non-nil, zero value otherwise.

### GetWebsiteThemeIdOk

`func (o *WebContentDto) GetWebsiteThemeIdOk() (*string, bool)`

GetWebsiteThemeIdOk returns a tuple with the WebsiteThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeId

`func (o *WebContentDto) SetWebsiteThemeId(v string)`

SetWebsiteThemeId sets WebsiteThemeId field to given value.

### HasWebsiteThemeId

`func (o *WebContentDto) HasWebsiteThemeId() bool`

HasWebsiteThemeId returns a boolean if a field has been set.

### SetWebsiteThemeIdNil

`func (o *WebContentDto) SetWebsiteThemeIdNil(b bool)`

 SetWebsiteThemeIdNil sets the value for WebsiteThemeId to be an explicit nil

### UnsetWebsiteThemeId
`func (o *WebContentDto) UnsetWebsiteThemeId()`

UnsetWebsiteThemeId ensures that no value is present for WebsiteThemeId, not even an explicit nil
### GetEnrollmentId

`func (o *WebContentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *WebContentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *WebContentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *WebContentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *WebContentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *WebContentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *WebContentDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *WebContentDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *WebContentDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *WebContentDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *WebContentDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *WebContentDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetParentWebContentId

`func (o *WebContentDto) GetParentWebContentId() string`

GetParentWebContentId returns the ParentWebContentId field if non-nil, zero value otherwise.

### GetParentWebContentIdOk

`func (o *WebContentDto) GetParentWebContentIdOk() (*string, bool)`

GetParentWebContentIdOk returns a tuple with the ParentWebContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentId

`func (o *WebContentDto) SetParentWebContentId(v string)`

SetParentWebContentId sets ParentWebContentId field to given value.

### HasParentWebContentId

`func (o *WebContentDto) HasParentWebContentId() bool`

HasParentWebContentId returns a boolean if a field has been set.

### SetParentWebContentIdNil

`func (o *WebContentDto) SetParentWebContentIdNil(b bool)`

 SetParentWebContentIdNil sets the value for ParentWebContentId to be an explicit nil

### UnsetParentWebContentId
`func (o *WebContentDto) UnsetParentWebContentId()`

UnsetParentWebContentId ensures that no value is present for ParentWebContentId, not even an explicit nil
### GetParentWebContentVersionId

`func (o *WebContentDto) GetParentWebContentVersionId() string`

GetParentWebContentVersionId returns the ParentWebContentVersionId field if non-nil, zero value otherwise.

### GetParentWebContentVersionIdOk

`func (o *WebContentDto) GetParentWebContentVersionIdOk() (*string, bool)`

GetParentWebContentVersionIdOk returns a tuple with the ParentWebContentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentVersionId

`func (o *WebContentDto) SetParentWebContentVersionId(v string)`

SetParentWebContentVersionId sets ParentWebContentVersionId field to given value.

### HasParentWebContentVersionId

`func (o *WebContentDto) HasParentWebContentVersionId() bool`

HasParentWebContentVersionId returns a boolean if a field has been set.

### SetParentWebContentVersionIdNil

`func (o *WebContentDto) SetParentWebContentVersionIdNil(b bool)`

 SetParentWebContentVersionIdNil sets the value for ParentWebContentVersionId to be an explicit nil

### UnsetParentWebContentVersionId
`func (o *WebContentDto) UnsetParentWebContentVersionId()`

UnsetParentWebContentVersionId ensures that no value is present for ParentWebContentVersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


