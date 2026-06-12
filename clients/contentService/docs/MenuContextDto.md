# MenuContextDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]MenuContextItemDto**](MenuContextItemDto.md) |  | [optional] 

## Methods

### NewMenuContextDto

`func NewMenuContextDto() *MenuContextDto`

NewMenuContextDto instantiates a new MenuContextDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuContextDtoWithDefaults

`func NewMenuContextDtoWithDefaults() *MenuContextDto`

NewMenuContextDtoWithDefaults instantiates a new MenuContextDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuContextDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuContextDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuContextDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuContextDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MenuContextDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MenuContextDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *MenuContextDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MenuContextDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MenuContextDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MenuContextDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MenuContextDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MenuContextDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *MenuContextDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuContextDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuContextDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MenuContextDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MenuContextDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MenuContextDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCategory

`func (o *MenuContextDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MenuContextDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MenuContextDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MenuContextDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MenuContextDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MenuContextDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetComponent

`func (o *MenuContextDto) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MenuContextDto) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MenuContextDto) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *MenuContextDto) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *MenuContextDto) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *MenuContextDto) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetEnable

`func (o *MenuContextDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MenuContextDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MenuContextDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MenuContextDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetStudioMenu

`func (o *MenuContextDto) GetStudioMenu() bool`

GetStudioMenu returns the StudioMenu field if non-nil, zero value otherwise.

### GetStudioMenuOk

`func (o *MenuContextDto) GetStudioMenuOk() (*bool, bool)`

GetStudioMenuOk returns a tuple with the StudioMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudioMenu

`func (o *MenuContextDto) SetStudioMenu(v bool)`

SetStudioMenu sets StudioMenu field to given value.

### HasStudioMenu

`func (o *MenuContextDto) HasStudioMenu() bool`

HasStudioMenu returns a boolean if a field has been set.

### GetCustomCss

`func (o *MenuContextDto) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *MenuContextDto) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *MenuContextDto) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *MenuContextDto) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *MenuContextDto) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *MenuContextDto) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomJs

`func (o *MenuContextDto) GetCustomJs() string`

GetCustomJs returns the CustomJs field if non-nil, zero value otherwise.

### GetCustomJsOk

`func (o *MenuContextDto) GetCustomJsOk() (*string, bool)`

GetCustomJsOk returns a tuple with the CustomJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomJs

`func (o *MenuContextDto) SetCustomJs(v string)`

SetCustomJs sets CustomJs field to given value.

### HasCustomJs

`func (o *MenuContextDto) HasCustomJs() bool`

HasCustomJs returns a boolean if a field has been set.

### SetCustomJsNil

`func (o *MenuContextDto) SetCustomJsNil(b bool)`

 SetCustomJsNil sets the value for CustomJs to be an explicit nil

### UnsetCustomJs
`func (o *MenuContextDto) UnsetCustomJs()`

UnsetCustomJs ensures that no value is present for CustomJs, not even an explicit nil
### GetCustomHtml

`func (o *MenuContextDto) GetCustomHtml() string`

GetCustomHtml returns the CustomHtml field if non-nil, zero value otherwise.

### GetCustomHtmlOk

`func (o *MenuContextDto) GetCustomHtmlOk() (*string, bool)`

GetCustomHtmlOk returns a tuple with the CustomHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHtml

`func (o *MenuContextDto) SetCustomHtml(v string)`

SetCustomHtml sets CustomHtml field to given value.

### HasCustomHtml

`func (o *MenuContextDto) HasCustomHtml() bool`

HasCustomHtml returns a boolean if a field has been set.

### SetCustomHtmlNil

`func (o *MenuContextDto) SetCustomHtmlNil(b bool)`

 SetCustomHtmlNil sets the value for CustomHtml to be an explicit nil

### UnsetCustomHtml
`func (o *MenuContextDto) UnsetCustomHtml()`

UnsetCustomHtml ensures that no value is present for CustomHtml, not even an explicit nil
### GetLoggedInOnly

`func (o *MenuContextDto) GetLoggedInOnly() string`

GetLoggedInOnly returns the LoggedInOnly field if non-nil, zero value otherwise.

### GetLoggedInOnlyOk

`func (o *MenuContextDto) GetLoggedInOnlyOk() (*string, bool)`

GetLoggedInOnlyOk returns a tuple with the LoggedInOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedInOnly

`func (o *MenuContextDto) SetLoggedInOnly(v string)`

SetLoggedInOnly sets LoggedInOnly field to given value.

### HasLoggedInOnly

`func (o *MenuContextDto) HasLoggedInOnly() bool`

HasLoggedInOnly returns a boolean if a field has been set.

### SetLoggedInOnlyNil

`func (o *MenuContextDto) SetLoggedInOnlyNil(b bool)`

 SetLoggedInOnlyNil sets the value for LoggedInOnly to be an explicit nil

### UnsetLoggedInOnly
`func (o *MenuContextDto) UnsetLoggedInOnly()`

UnsetLoggedInOnly ensures that no value is present for LoggedInOnly, not even an explicit nil
### GetBackgroundImage

`func (o *MenuContextDto) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *MenuContextDto) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *MenuContextDto) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *MenuContextDto) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### SetBackgroundImageNil

`func (o *MenuContextDto) SetBackgroundImageNil(b bool)`

 SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil

### UnsetBackgroundImage
`func (o *MenuContextDto) UnsetBackgroundImage()`

UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
### GetWebPortalId

`func (o *MenuContextDto) GetWebPortalId() string`

GetWebPortalId returns the WebPortalId field if non-nil, zero value otherwise.

### GetWebPortalIdOk

`func (o *MenuContextDto) GetWebPortalIdOk() (*string, bool)`

GetWebPortalIdOk returns a tuple with the WebPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalId

`func (o *MenuContextDto) SetWebPortalId(v string)`

SetWebPortalId sets WebPortalId field to given value.

### HasWebPortalId

`func (o *MenuContextDto) HasWebPortalId() bool`

HasWebPortalId returns a boolean if a field has been set.

### SetWebPortalIdNil

`func (o *MenuContextDto) SetWebPortalIdNil(b bool)`

 SetWebPortalIdNil sets the value for WebPortalId to be an explicit nil

### UnsetWebPortalId
`func (o *MenuContextDto) UnsetWebPortalId()`

UnsetWebPortalId ensures that no value is present for WebPortalId, not even an explicit nil
### GetTenantId

`func (o *MenuContextDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MenuContextDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MenuContextDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MenuContextDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MenuContextDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MenuContextDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetItems

`func (o *MenuContextDto) GetItems() []MenuContextItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MenuContextDto) GetItemsOk() (*[]MenuContextItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MenuContextDto) SetItems(v []MenuContextItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *MenuContextDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *MenuContextDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *MenuContextDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


