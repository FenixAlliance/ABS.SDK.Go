# BackgroundOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundPatternID** | Pointer to **int32** |  | [optional] 
**EnableBackgroundPattern** | Pointer to **bool** |  | [optional] 
**BackgroundImageForPage** | Pointer to **NullableString** |  | [optional] 
**BackgroundColorForPage** | Pointer to **NullableString** |  | [optional] 
**MainContentColor** | Pointer to **NullableString** |  | [optional] 
**MainContentImageURL** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackgroundOptions

`func NewBackgroundOptions() *BackgroundOptions`

NewBackgroundOptions instantiates a new BackgroundOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundOptionsWithDefaults

`func NewBackgroundOptionsWithDefaults() *BackgroundOptions`

NewBackgroundOptionsWithDefaults instantiates a new BackgroundOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundPatternID

`func (o *BackgroundOptions) GetBackgroundPatternID() int32`

GetBackgroundPatternID returns the BackgroundPatternID field if non-nil, zero value otherwise.

### GetBackgroundPatternIDOk

`func (o *BackgroundOptions) GetBackgroundPatternIDOk() (*int32, bool)`

GetBackgroundPatternIDOk returns a tuple with the BackgroundPatternID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPatternID

`func (o *BackgroundOptions) SetBackgroundPatternID(v int32)`

SetBackgroundPatternID sets BackgroundPatternID field to given value.

### HasBackgroundPatternID

`func (o *BackgroundOptions) HasBackgroundPatternID() bool`

HasBackgroundPatternID returns a boolean if a field has been set.

### GetEnableBackgroundPattern

`func (o *BackgroundOptions) GetEnableBackgroundPattern() bool`

GetEnableBackgroundPattern returns the EnableBackgroundPattern field if non-nil, zero value otherwise.

### GetEnableBackgroundPatternOk

`func (o *BackgroundOptions) GetEnableBackgroundPatternOk() (*bool, bool)`

GetEnableBackgroundPatternOk returns a tuple with the EnableBackgroundPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBackgroundPattern

`func (o *BackgroundOptions) SetEnableBackgroundPattern(v bool)`

SetEnableBackgroundPattern sets EnableBackgroundPattern field to given value.

### HasEnableBackgroundPattern

`func (o *BackgroundOptions) HasEnableBackgroundPattern() bool`

HasEnableBackgroundPattern returns a boolean if a field has been set.

### GetBackgroundImageForPage

`func (o *BackgroundOptions) GetBackgroundImageForPage() string`

GetBackgroundImageForPage returns the BackgroundImageForPage field if non-nil, zero value otherwise.

### GetBackgroundImageForPageOk

`func (o *BackgroundOptions) GetBackgroundImageForPageOk() (*string, bool)`

GetBackgroundImageForPageOk returns a tuple with the BackgroundImageForPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImageForPage

`func (o *BackgroundOptions) SetBackgroundImageForPage(v string)`

SetBackgroundImageForPage sets BackgroundImageForPage field to given value.

### HasBackgroundImageForPage

`func (o *BackgroundOptions) HasBackgroundImageForPage() bool`

HasBackgroundImageForPage returns a boolean if a field has been set.

### SetBackgroundImageForPageNil

`func (o *BackgroundOptions) SetBackgroundImageForPageNil(b bool)`

 SetBackgroundImageForPageNil sets the value for BackgroundImageForPage to be an explicit nil

### UnsetBackgroundImageForPage
`func (o *BackgroundOptions) UnsetBackgroundImageForPage()`

UnsetBackgroundImageForPage ensures that no value is present for BackgroundImageForPage, not even an explicit nil
### GetBackgroundColorForPage

`func (o *BackgroundOptions) GetBackgroundColorForPage() string`

GetBackgroundColorForPage returns the BackgroundColorForPage field if non-nil, zero value otherwise.

### GetBackgroundColorForPageOk

`func (o *BackgroundOptions) GetBackgroundColorForPageOk() (*string, bool)`

GetBackgroundColorForPageOk returns a tuple with the BackgroundColorForPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColorForPage

`func (o *BackgroundOptions) SetBackgroundColorForPage(v string)`

SetBackgroundColorForPage sets BackgroundColorForPage field to given value.

### HasBackgroundColorForPage

`func (o *BackgroundOptions) HasBackgroundColorForPage() bool`

HasBackgroundColorForPage returns a boolean if a field has been set.

### SetBackgroundColorForPageNil

`func (o *BackgroundOptions) SetBackgroundColorForPageNil(b bool)`

 SetBackgroundColorForPageNil sets the value for BackgroundColorForPage to be an explicit nil

### UnsetBackgroundColorForPage
`func (o *BackgroundOptions) UnsetBackgroundColorForPage()`

UnsetBackgroundColorForPage ensures that no value is present for BackgroundColorForPage, not even an explicit nil
### GetMainContentColor

`func (o *BackgroundOptions) GetMainContentColor() string`

GetMainContentColor returns the MainContentColor field if non-nil, zero value otherwise.

### GetMainContentColorOk

`func (o *BackgroundOptions) GetMainContentColorOk() (*string, bool)`

GetMainContentColorOk returns a tuple with the MainContentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContentColor

`func (o *BackgroundOptions) SetMainContentColor(v string)`

SetMainContentColor sets MainContentColor field to given value.

### HasMainContentColor

`func (o *BackgroundOptions) HasMainContentColor() bool`

HasMainContentColor returns a boolean if a field has been set.

### SetMainContentColorNil

`func (o *BackgroundOptions) SetMainContentColorNil(b bool)`

 SetMainContentColorNil sets the value for MainContentColor to be an explicit nil

### UnsetMainContentColor
`func (o *BackgroundOptions) UnsetMainContentColor()`

UnsetMainContentColor ensures that no value is present for MainContentColor, not even an explicit nil
### GetMainContentImageURL

`func (o *BackgroundOptions) GetMainContentImageURL() string`

GetMainContentImageURL returns the MainContentImageURL field if non-nil, zero value otherwise.

### GetMainContentImageURLOk

`func (o *BackgroundOptions) GetMainContentImageURLOk() (*string, bool)`

GetMainContentImageURLOk returns a tuple with the MainContentImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContentImageURL

`func (o *BackgroundOptions) SetMainContentImageURL(v string)`

SetMainContentImageURL sets MainContentImageURL field to given value.

### HasMainContentImageURL

`func (o *BackgroundOptions) HasMainContentImageURL() bool`

HasMainContentImageURL returns a boolean if a field has been set.

### SetMainContentImageURLNil

`func (o *BackgroundOptions) SetMainContentImageURLNil(b bool)`

 SetMainContentImageURLNil sets the value for MainContentImageURL to be an explicit nil

### UnsetMainContentImageURL
`func (o *BackgroundOptions) UnsetMainContentImageURL()`

UnsetMainContentImageURL ensures that no value is present for MainContentImageURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


