# WebPageCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPageCategoryDto

`func NewWebPageCategoryDto() *WebPageCategoryDto`

NewWebPageCategoryDto instantiates a new WebPageCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageCategoryDtoWithDefaults

`func NewWebPageCategoryDtoWithDefaults() *WebPageCategoryDto`

NewWebPageCategoryDtoWithDefaults instantiates a new WebPageCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPageCategoryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPageCategoryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPageCategoryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPageCategoryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebPageCategoryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebPageCategoryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WebPageCategoryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPageCategoryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPageCategoryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPageCategoryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WebPageCategoryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WebPageCategoryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSlug

`func (o *WebPageCategoryDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebPageCategoryDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebPageCategoryDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebPageCategoryDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebPageCategoryDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebPageCategoryDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetTitle

`func (o *WebPageCategoryDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPageCategoryDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPageCategoryDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPageCategoryDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPageCategoryDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPageCategoryDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WebPageCategoryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPageCategoryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPageCategoryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPageCategoryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPageCategoryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPageCategoryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSeoTitle

`func (o *WebPageCategoryDto) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *WebPageCategoryDto) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *WebPageCategoryDto) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *WebPageCategoryDto) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *WebPageCategoryDto) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *WebPageCategoryDto) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetMetaDescription

`func (o *WebPageCategoryDto) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *WebPageCategoryDto) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *WebPageCategoryDto) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *WebPageCategoryDto) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *WebPageCategoryDto) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *WebPageCategoryDto) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetCornerstoneContent

`func (o *WebPageCategoryDto) GetCornerstoneContent() bool`

GetCornerstoneContent returns the CornerstoneContent field if non-nil, zero value otherwise.

### GetCornerstoneContentOk

`func (o *WebPageCategoryDto) GetCornerstoneContentOk() (*bool, bool)`

GetCornerstoneContentOk returns a tuple with the CornerstoneContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerstoneContent

`func (o *WebPageCategoryDto) SetCornerstoneContent(v bool)`

SetCornerstoneContent sets CornerstoneContent field to given value.

### HasCornerstoneContent

`func (o *WebPageCategoryDto) HasCornerstoneContent() bool`

HasCornerstoneContent returns a boolean if a field has been set.

### GetAllowSerachEngines

`func (o *WebPageCategoryDto) GetAllowSerachEngines() bool`

GetAllowSerachEngines returns the AllowSerachEngines field if non-nil, zero value otherwise.

### GetAllowSerachEnginesOk

`func (o *WebPageCategoryDto) GetAllowSerachEnginesOk() (*bool, bool)`

GetAllowSerachEnginesOk returns a tuple with the AllowSerachEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSerachEngines

`func (o *WebPageCategoryDto) SetAllowSerachEngines(v bool)`

SetAllowSerachEngines sets AllowSerachEngines field to given value.

### HasAllowSerachEngines

`func (o *WebPageCategoryDto) HasAllowSerachEngines() bool`

HasAllowSerachEngines returns a boolean if a field has been set.

### GetSeoKeyPhrases

`func (o *WebPageCategoryDto) GetSeoKeyPhrases() string`

GetSeoKeyPhrases returns the SeoKeyPhrases field if non-nil, zero value otherwise.

### GetSeoKeyPhrasesOk

`func (o *WebPageCategoryDto) GetSeoKeyPhrasesOk() (*string, bool)`

GetSeoKeyPhrasesOk returns a tuple with the SeoKeyPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoKeyPhrases

`func (o *WebPageCategoryDto) SetSeoKeyPhrases(v string)`

SetSeoKeyPhrases sets SeoKeyPhrases field to given value.

### HasSeoKeyPhrases

`func (o *WebPageCategoryDto) HasSeoKeyPhrases() bool`

HasSeoKeyPhrases returns a boolean if a field has been set.

### SetSeoKeyPhrasesNil

`func (o *WebPageCategoryDto) SetSeoKeyPhrasesNil(b bool)`

 SetSeoKeyPhrasesNil sets the value for SeoKeyPhrases to be an explicit nil

### UnsetSeoKeyPhrases
`func (o *WebPageCategoryDto) UnsetSeoKeyPhrases()`

UnsetSeoKeyPhrases ensures that no value is present for SeoKeyPhrases, not even an explicit nil
### GetCanonicalUrl

`func (o *WebPageCategoryDto) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *WebPageCategoryDto) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *WebPageCategoryDto) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *WebPageCategoryDto) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### SetCanonicalUrlNil

`func (o *WebPageCategoryDto) SetCanonicalUrlNil(b bool)`

 SetCanonicalUrlNil sets the value for CanonicalUrl to be an explicit nil

### UnsetCanonicalUrl
`func (o *WebPageCategoryDto) UnsetCanonicalUrl()`

UnsetCanonicalUrl ensures that no value is present for CanonicalUrl, not even an explicit nil
### GetImageURL

`func (o *WebPageCategoryDto) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *WebPageCategoryDto) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *WebPageCategoryDto) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.

### HasImageURL

`func (o *WebPageCategoryDto) HasImageURL() bool`

HasImageURL returns a boolean if a field has been set.

### SetImageURLNil

`func (o *WebPageCategoryDto) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *WebPageCategoryDto) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil
### GetImage

`func (o *WebPageCategoryDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WebPageCategoryDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WebPageCategoryDto) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *WebPageCategoryDto) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *WebPageCategoryDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *WebPageCategoryDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetWebPortalID

`func (o *WebPageCategoryDto) GetWebPortalID() string`

GetWebPortalID returns the WebPortalID field if non-nil, zero value otherwise.

### GetWebPortalIDOk

`func (o *WebPageCategoryDto) GetWebPortalIDOk() (*string, bool)`

GetWebPortalIDOk returns a tuple with the WebPortalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalID

`func (o *WebPageCategoryDto) SetWebPortalID(v string)`

SetWebPortalID sets WebPortalID field to given value.

### HasWebPortalID

`func (o *WebPageCategoryDto) HasWebPortalID() bool`

HasWebPortalID returns a boolean if a field has been set.

### SetWebPortalIDNil

`func (o *WebPageCategoryDto) SetWebPortalIDNil(b bool)`

 SetWebPortalIDNil sets the value for WebPortalID to be an explicit nil

### UnsetWebPortalID
`func (o *WebPageCategoryDto) UnsetWebPortalID()`

UnsetWebPortalID ensures that no value is present for WebPortalID, not even an explicit nil
### GetBusinessID

`func (o *WebPageCategoryDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *WebPageCategoryDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *WebPageCategoryDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *WebPageCategoryDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *WebPageCategoryDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *WebPageCategoryDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *WebPageCategoryDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *WebPageCategoryDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *WebPageCategoryDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *WebPageCategoryDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *WebPageCategoryDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *WebPageCategoryDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


