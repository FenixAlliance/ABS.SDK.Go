# WebTemplateUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**CssContent** | Pointer to **NullableString** |  | [optional] 
**JsContent** | Pointer to **NullableString** |  | [optional] 
**RazorContent** | Pointer to **NullableString** |  | [optional] 
**HighlightImage** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewWebTemplateUpdateDto

`func NewWebTemplateUpdateDto() *WebTemplateUpdateDto`

NewWebTemplateUpdateDto instantiates a new WebTemplateUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebTemplateUpdateDtoWithDefaults

`func NewWebTemplateUpdateDtoWithDefaults() *WebTemplateUpdateDto`

NewWebTemplateUpdateDtoWithDefaults instantiates a new WebTemplateUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebTemplateUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebTemplateUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebTemplateUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebTemplateUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebTemplateUpdateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebTemplateUpdateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WebTemplateUpdateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebTemplateUpdateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebTemplateUpdateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebTemplateUpdateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WebTemplateUpdateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WebTemplateUpdateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSlug

`func (o *WebTemplateUpdateDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WebTemplateUpdateDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WebTemplateUpdateDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WebTemplateUpdateDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *WebTemplateUpdateDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *WebTemplateUpdateDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *WebTemplateUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebTemplateUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebTemplateUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebTemplateUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebTemplateUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebTemplateUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *WebTemplateUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebTemplateUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebTemplateUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebTemplateUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebTemplateUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebTemplateUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WebTemplateUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebTemplateUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebTemplateUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebTemplateUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebTemplateUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebTemplateUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *WebTemplateUpdateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebTemplateUpdateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebTemplateUpdateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebTemplateUpdateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebTemplateUpdateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebTemplateUpdateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetHtmlContent

`func (o *WebTemplateUpdateDto) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *WebTemplateUpdateDto) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *WebTemplateUpdateDto) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *WebTemplateUpdateDto) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *WebTemplateUpdateDto) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *WebTemplateUpdateDto) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCssContent

`func (o *WebTemplateUpdateDto) GetCssContent() string`

GetCssContent returns the CssContent field if non-nil, zero value otherwise.

### GetCssContentOk

`func (o *WebTemplateUpdateDto) GetCssContentOk() (*string, bool)`

GetCssContentOk returns a tuple with the CssContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssContent

`func (o *WebTemplateUpdateDto) SetCssContent(v string)`

SetCssContent sets CssContent field to given value.

### HasCssContent

`func (o *WebTemplateUpdateDto) HasCssContent() bool`

HasCssContent returns a boolean if a field has been set.

### SetCssContentNil

`func (o *WebTemplateUpdateDto) SetCssContentNil(b bool)`

 SetCssContentNil sets the value for CssContent to be an explicit nil

### UnsetCssContent
`func (o *WebTemplateUpdateDto) UnsetCssContent()`

UnsetCssContent ensures that no value is present for CssContent, not even an explicit nil
### GetJsContent

`func (o *WebTemplateUpdateDto) GetJsContent() string`

GetJsContent returns the JsContent field if non-nil, zero value otherwise.

### GetJsContentOk

`func (o *WebTemplateUpdateDto) GetJsContentOk() (*string, bool)`

GetJsContentOk returns a tuple with the JsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsContent

`func (o *WebTemplateUpdateDto) SetJsContent(v string)`

SetJsContent sets JsContent field to given value.

### HasJsContent

`func (o *WebTemplateUpdateDto) HasJsContent() bool`

HasJsContent returns a boolean if a field has been set.

### SetJsContentNil

`func (o *WebTemplateUpdateDto) SetJsContentNil(b bool)`

 SetJsContentNil sets the value for JsContent to be an explicit nil

### UnsetJsContent
`func (o *WebTemplateUpdateDto) UnsetJsContent()`

UnsetJsContent ensures that no value is present for JsContent, not even an explicit nil
### GetRazorContent

`func (o *WebTemplateUpdateDto) GetRazorContent() string`

GetRazorContent returns the RazorContent field if non-nil, zero value otherwise.

### GetRazorContentOk

`func (o *WebTemplateUpdateDto) GetRazorContentOk() (*string, bool)`

GetRazorContentOk returns a tuple with the RazorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorContent

`func (o *WebTemplateUpdateDto) SetRazorContent(v string)`

SetRazorContent sets RazorContent field to given value.

### HasRazorContent

`func (o *WebTemplateUpdateDto) HasRazorContent() bool`

HasRazorContent returns a boolean if a field has been set.

### SetRazorContentNil

`func (o *WebTemplateUpdateDto) SetRazorContentNil(b bool)`

 SetRazorContentNil sets the value for RazorContent to be an explicit nil

### UnsetRazorContent
`func (o *WebTemplateUpdateDto) UnsetRazorContent()`

UnsetRazorContent ensures that no value is present for RazorContent, not even an explicit nil
### GetHighlightImage

`func (o *WebTemplateUpdateDto) GetHighlightImage() string`

GetHighlightImage returns the HighlightImage field if non-nil, zero value otherwise.

### GetHighlightImageOk

`func (o *WebTemplateUpdateDto) GetHighlightImageOk() (*string, bool)`

GetHighlightImageOk returns a tuple with the HighlightImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightImage

`func (o *WebTemplateUpdateDto) SetHighlightImage(v string)`

SetHighlightImage sets HighlightImage field to given value.

### HasHighlightImage

`func (o *WebTemplateUpdateDto) HasHighlightImage() bool`

HasHighlightImage returns a boolean if a field has been set.

### SetHighlightImageNil

`func (o *WebTemplateUpdateDto) SetHighlightImageNil(b bool)`

 SetHighlightImageNil sets the value for HighlightImage to be an explicit nil

### UnsetHighlightImage
`func (o *WebTemplateUpdateDto) UnsetHighlightImage()`

UnsetHighlightImage ensures that no value is present for HighlightImage, not even an explicit nil
### GetOrder

`func (o *WebTemplateUpdateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WebTemplateUpdateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WebTemplateUpdateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WebTemplateUpdateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


