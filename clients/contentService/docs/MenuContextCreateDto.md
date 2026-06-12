# MenuContextCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Category** | Pointer to **NullableString** |  | [optional] 
**Component** | Pointer to **NullableString** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**StudioMenu** | Pointer to **bool** |  | [optional] 
**CustomCss** | Pointer to **NullableString** |  | [optional] 
**CustomJs** | Pointer to **NullableString** |  | [optional] 
**CustomHtml** | Pointer to **NullableString** |  | [optional] 
**LoggedInOnly** | Pointer to **NullableString** |  | [optional] 
**BackgroundImage** | Pointer to **NullableString** |  | [optional] 
**WebPortalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMenuContextCreateDto

`func NewMenuContextCreateDto(name string, ) *MenuContextCreateDto`

NewMenuContextCreateDto instantiates a new MenuContextCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuContextCreateDtoWithDefaults

`func NewMenuContextCreateDtoWithDefaults() *MenuContextCreateDto`

NewMenuContextCreateDtoWithDefaults instantiates a new MenuContextCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuContextCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuContextCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuContextCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuContextCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *MenuContextCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MenuContextCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MenuContextCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MenuContextCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *MenuContextCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuContextCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuContextCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *MenuContextCreateDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MenuContextCreateDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MenuContextCreateDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MenuContextCreateDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MenuContextCreateDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MenuContextCreateDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetComponent

`func (o *MenuContextCreateDto) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MenuContextCreateDto) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MenuContextCreateDto) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *MenuContextCreateDto) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *MenuContextCreateDto) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *MenuContextCreateDto) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetEnable

`func (o *MenuContextCreateDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MenuContextCreateDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MenuContextCreateDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MenuContextCreateDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetStudioMenu

`func (o *MenuContextCreateDto) GetStudioMenu() bool`

GetStudioMenu returns the StudioMenu field if non-nil, zero value otherwise.

### GetStudioMenuOk

`func (o *MenuContextCreateDto) GetStudioMenuOk() (*bool, bool)`

GetStudioMenuOk returns a tuple with the StudioMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudioMenu

`func (o *MenuContextCreateDto) SetStudioMenu(v bool)`

SetStudioMenu sets StudioMenu field to given value.

### HasStudioMenu

`func (o *MenuContextCreateDto) HasStudioMenu() bool`

HasStudioMenu returns a boolean if a field has been set.

### GetCustomCss

`func (o *MenuContextCreateDto) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *MenuContextCreateDto) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *MenuContextCreateDto) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *MenuContextCreateDto) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *MenuContextCreateDto) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *MenuContextCreateDto) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomJs

`func (o *MenuContextCreateDto) GetCustomJs() string`

GetCustomJs returns the CustomJs field if non-nil, zero value otherwise.

### GetCustomJsOk

`func (o *MenuContextCreateDto) GetCustomJsOk() (*string, bool)`

GetCustomJsOk returns a tuple with the CustomJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomJs

`func (o *MenuContextCreateDto) SetCustomJs(v string)`

SetCustomJs sets CustomJs field to given value.

### HasCustomJs

`func (o *MenuContextCreateDto) HasCustomJs() bool`

HasCustomJs returns a boolean if a field has been set.

### SetCustomJsNil

`func (o *MenuContextCreateDto) SetCustomJsNil(b bool)`

 SetCustomJsNil sets the value for CustomJs to be an explicit nil

### UnsetCustomJs
`func (o *MenuContextCreateDto) UnsetCustomJs()`

UnsetCustomJs ensures that no value is present for CustomJs, not even an explicit nil
### GetCustomHtml

`func (o *MenuContextCreateDto) GetCustomHtml() string`

GetCustomHtml returns the CustomHtml field if non-nil, zero value otherwise.

### GetCustomHtmlOk

`func (o *MenuContextCreateDto) GetCustomHtmlOk() (*string, bool)`

GetCustomHtmlOk returns a tuple with the CustomHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHtml

`func (o *MenuContextCreateDto) SetCustomHtml(v string)`

SetCustomHtml sets CustomHtml field to given value.

### HasCustomHtml

`func (o *MenuContextCreateDto) HasCustomHtml() bool`

HasCustomHtml returns a boolean if a field has been set.

### SetCustomHtmlNil

`func (o *MenuContextCreateDto) SetCustomHtmlNil(b bool)`

 SetCustomHtmlNil sets the value for CustomHtml to be an explicit nil

### UnsetCustomHtml
`func (o *MenuContextCreateDto) UnsetCustomHtml()`

UnsetCustomHtml ensures that no value is present for CustomHtml, not even an explicit nil
### GetLoggedInOnly

`func (o *MenuContextCreateDto) GetLoggedInOnly() string`

GetLoggedInOnly returns the LoggedInOnly field if non-nil, zero value otherwise.

### GetLoggedInOnlyOk

`func (o *MenuContextCreateDto) GetLoggedInOnlyOk() (*string, bool)`

GetLoggedInOnlyOk returns a tuple with the LoggedInOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInOnly

`func (o *MenuContextCreateDto) SetLoggedInOnly(v string)`

SetLoggedInOnly sets LoggedInOnly field to given value.

### HasLoggedInOnly

`func (o *MenuContextCreateDto) HasLoggedInOnly() bool`

HasLoggedInOnly returns a boolean if a field has been set.

### SetLoggedInOnlyNil

`func (o *MenuContextCreateDto) SetLoggedInOnlyNil(b bool)`

 SetLoggedInOnlyNil sets the value for LoggedInOnly to be an explicit nil

### UnsetLoggedInOnly
`func (o *MenuContextCreateDto) UnsetLoggedInOnly()`

UnsetLoggedInOnly ensures that no value is present for LoggedInOnly, not even an explicit nil
### GetBackgroundImage

`func (o *MenuContextCreateDto) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *MenuContextCreateDto) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *MenuContextCreateDto) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *MenuContextCreateDto) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### SetBackgroundImageNil

`func (o *MenuContextCreateDto) SetBackgroundImageNil(b bool)`

 SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil

### UnsetBackgroundImage
`func (o *MenuContextCreateDto) UnsetBackgroundImage()`

UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
### GetWebPortalId

`func (o *MenuContextCreateDto) GetWebPortalId() string`

GetWebPortalId returns the WebPortalId field if non-nil, zero value otherwise.

### GetWebPortalIdOk

`func (o *MenuContextCreateDto) GetWebPortalIdOk() (*string, bool)`

GetWebPortalIdOk returns a tuple with the WebPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalId

`func (o *MenuContextCreateDto) SetWebPortalId(v string)`

SetWebPortalId sets WebPortalId field to given value.

### HasWebPortalId

`func (o *MenuContextCreateDto) HasWebPortalId() bool`

HasWebPortalId returns a boolean if a field has been set.

### SetWebPortalIdNil

`func (o *MenuContextCreateDto) SetWebPortalIdNil(b bool)`

 SetWebPortalIdNil sets the value for WebPortalId to be an explicit nil

### UnsetWebPortalId
`func (o *MenuContextCreateDto) UnsetWebPortalId()`

UnsetWebPortalId ensures that no value is present for WebPortalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


