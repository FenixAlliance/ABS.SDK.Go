# WebPageCategoryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SeoTitle** | Pointer to **NullableString** |  | [optional] 
**MetaDescription** | Pointer to **NullableString** |  | [optional] 
**CornerstoneContent** | Pointer to **bool** |  | [optional] 
**AllowSerachEngines** | Pointer to **bool** |  | [optional] 
**SeoKeyPhrases** | Pointer to **NullableString** |  | [optional] 
**CanonicalUrl** | Pointer to **NullableString** |  | [optional] 
**ImageURL** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**WebPortalID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPageCategoryCreateDto

`func NewWebPageCategoryCreateDto() *WebPageCategoryCreateDto`

NewWebPageCategoryCreateDto instantiates a new WebPageCategoryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageCategoryCreateDtoWithDefaults

`func NewWebPageCategoryCreateDtoWithDefaults() *WebPageCategoryCreateDto`

NewWebPageCategoryCreateDtoWithDefaults instantiates a new WebPageCategoryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPageCategoryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPageCategoryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPageCategoryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPageCategoryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebPageCategoryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPageCategoryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPageCategoryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPageCategoryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSlug

`func (o *WebPageCategoryCreateDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebPageCategoryCreateDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebPageCategoryCreateDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebPageCategoryCreateDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebPageCategoryCreateDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebPageCategoryCreateDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetTitle

`func (o *WebPageCategoryCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPageCategoryCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPageCategoryCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPageCategoryCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPageCategoryCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPageCategoryCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WebPageCategoryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPageCategoryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPageCategoryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPageCategoryCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPageCategoryCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPageCategoryCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSeoTitle

`func (o *WebPageCategoryCreateDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebPageCategoryCreateDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebPageCategoryCreateDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebPageCategoryCreateDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebPageCategoryCreateDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebPageCategoryCreateDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetMetaDescription

`func (o *WebPageCategoryCreateDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebPageCategoryCreateDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebPageCategoryCreateDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebPageCategoryCreateDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebPageCategoryCreateDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebPageCategoryCreateDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetCornerstoneContent

`func (o *WebPageCategoryCreateDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebPageCategoryCreateDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebPageCategoryCreateDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebPageCategoryCreateDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetAllowSerachEngines

`func (o *WebPageCategoryCreateDto) GetAllowSerachEngines() bool`

GetAllowSerachEngines returns the AllowSerachEngines field if non-nil, zero value otherwise.

### GetAllowSerachEnginesOk

`func (o *WebPageCategoryCreateDto) GetAllowSerachEnginesOk() (*bool, bool)`

GetAllowSerachEnginesOk returns a tuple with the AllowSerachEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSerachEngines

`func (o *WebPageCategoryCreateDto) SetAllowSerachEngines(v bool)`

SetAllowSerachEngines sets AllowSerachEngines field to given value.

### HasAllowSerachEngines

`func (o *WebPageCategoryCreateDto) HasAllowSerachEngines() bool`

HasAllowSerachEngines returns a boolean if a field has been set.

### GetSeoKeyPhrases

`func (o *WebPageCategoryCreateDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebPageCategoryCreateDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebPageCategoryCreateDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebPageCategoryCreateDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebPageCategoryCreateDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebPageCategoryCreateDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetCanonicalUrl

`func (o *WebPageCategoryCreateDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebPageCategoryCreateDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebPageCategoryCreateDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebPageCategoryCreateDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebPageCategoryCreateDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebPageCategoryCreateDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetImageURL

`func (o *WebPageCategoryCreateDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *WebPageCategoryCreateDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *WebPageCategoryCreateDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *WebPageCategoryCreateDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *WebPageCategoryCreateDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *WebPageCategoryCreateDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetImage

`func (o *WebPageCategoryCreateDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WebPageCategoryCreateDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WebPageCategoryCreateDto) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *WebPageCategoryCreateDto) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *WebPageCategoryCreateDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *WebPageCategoryCreateDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetWebPortalID

`func (o *WebPageCategoryCreateDto) GetWebPortalID() string`

GetWebPortalID returns the WebPortalID field if non-nil, zero value otherwise.

### GetWebPortalIDOk

`func (o *WebPageCategoryCreateDto) GetWebPortalIDOk() (*string, bool)`

GetWebPortalIDOk returns a tuple with the WebPortalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalID

`func (o *WebPageCategoryCreateDto) SetWebPortalID(v string)`

SetWebPortalID sets WebPortalID field to given value.

### HasWebPortalID

`func (o *WebPageCategoryCreateDto) HasWebPortalID() bool`

HasWebPortalID returns a boolean if a field has been set.

### SetWebPortalIDNil

`func (o *WebPageCategoryCreateDto) SetWebPortalIDNil(b bool)`

 SetWebPortalIDNil sets the value for WebPortalID to be an explicit nil

### UnsetWebPortalID
`func (o *WebPageCategoryCreateDto) UnsetWebPortalID()`

UnsetWebPortalID ensures that no value is present for WebPortalID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *WebPageCategoryCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *WebPageCategoryCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *WebPageCategoryCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *WebPageCategoryCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *WebPageCategoryCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *WebPageCategoryCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


