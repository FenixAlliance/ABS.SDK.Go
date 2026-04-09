# Logo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margin** | Pointer to [**Margin**](Margin.md) |  | [optional] 
**Alignment** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**CustomLinkURL** | Pointer to **NullableString** |  | [optional] 
**BackgroundColor** | Pointer to **NullableString** |  | [optional] 
**DefaltLogoURL** | Pointer to **NullableString** |  | [optional] 
**DefaltRetinaLogoURL** | Pointer to **NullableString** |  | [optional] 
**DefaltStickyLogoURL** | Pointer to **NullableString** |  | [optional] 
**DefaltStickyRetinaLogoURL** | Pointer to **NullableString** |  | [optional] 
**DefaltMobileLogoURL** | Pointer to **NullableString** |  | [optional] 
**DefaltMobileRetinaLogoURL** | Pointer to **NullableString** |  | [optional] 
**Footer** | Pointer to **NullableString** |  | [optional] 
**Header** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogo

`func NewLogo() *Logo`

NewLogo instantiates a new Logo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoWithDefaults

`func NewLogoWithDefaults() *Logo`

NewLogoWithDefaults instantiates a new Logo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMargin

`func (o *Logo) GetMargin() Margin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *Logo) GetMarginOk() (*Margin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *Logo) SetMargin(v Margin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *Logo) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetAlignment

`func (o *Logo) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *Logo) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *Logo) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *Logo) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetWidth

`func (o *Logo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Logo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Logo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Logo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *Logo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *Logo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *Logo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Logo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Logo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Logo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *Logo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *Logo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetMaxWidth

`func (o *Logo) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *Logo) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *Logo) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *Logo) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *Logo) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *Logo) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMaxHeight

`func (o *Logo) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *Logo) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *Logo) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *Logo) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *Logo) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *Logo) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetCustomLinkURL

`func (o *Logo) GetCustomLinkURL() string`

GetCustomLinkURL returns the CustomLinkURL field if non-nil, zero value otherwise.

### GetCustomLinkURLOk

`func (o *Logo) GetCustomLinkURLOk() (*string, bool)`

GetCustomLinkURLOk returns a tuple with the CustomLinkURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinkURL

`func (o *Logo) SetCustomLinkURL(v string)`

SetCustomLinkURL sets CustomLinkURL field to given value.

### HasCustomLinkURL

`func (o *Logo) HasCustomLinkURL() bool`

HasCustomLinkURL returns a boolean if a field has been set.

### SetCustomLinkURLNil

`func (o *Logo) SetCustomLinkURLNil(b bool)`

 SetCustomLinkURLNil sets the value for CustomLinkURL to be an explicit nil

### UnsetCustomLinkURL
`func (o *Logo) UnsetCustomLinkURL()`

UnsetCustomLinkURL ensures that no value is present for CustomLinkURL, not even an explicit nil
### GetBackgroundColor

`func (o *Logo) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *Logo) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *Logo) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *Logo) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *Logo) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *Logo) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetDefaltLogoURL

`func (o *Logo) GetDefaltLogoURL() string`

GetDefaltLogoURL returns the DefaltLogoURL field if non-nil, zero value otherwise.

### GetDefaltLogoURLOk

`func (o *Logo) GetDefaltLogoURLOk() (*string, bool)`

GetDefaltLogoURLOk returns a tuple with the DefaltLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltLogoURL

`func (o *Logo) SetDefaltLogoURL(v string)`

SetDefaltLogoURL sets DefaltLogoURL field to given value.

### HasDefaltLogoURL

`func (o *Logo) HasDefaltLogoURL() bool`

HasDefaltLogoURL returns a boolean if a field has been set.

### SetDefaltLogoURLNil

`func (o *Logo) SetDefaltLogoURLNil(b bool)`

 SetDefaltLogoURLNil sets the value for DefaltLogoURL to be an explicit nil

### UnsetDefaltLogoURL
`func (o *Logo) UnsetDefaltLogoURL()`

UnsetDefaltLogoURL ensures that no value is present for DefaltLogoURL, not even an explicit nil
### GetDefaltRetinaLogoURL

`func (o *Logo) GetDefaltRetinaLogoURL() string`

GetDefaltRetinaLogoURL returns the DefaltRetinaLogoURL field if non-nil, zero value otherwise.

### GetDefaltRetinaLogoURLOk

`func (o *Logo) GetDefaltRetinaLogoURLOk() (*string, bool)`

GetDefaltRetinaLogoURLOk returns a tuple with the DefaltRetinaLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltRetinaLogoURL

`func (o *Logo) SetDefaltRetinaLogoURL(v string)`

SetDefaltRetinaLogoURL sets DefaltRetinaLogoURL field to given value.

### HasDefaltRetinaLogoURL

`func (o *Logo) HasDefaltRetinaLogoURL() bool`

HasDefaltRetinaLogoURL returns a boolean if a field has been set.

### SetDefaltRetinaLogoURLNil

`func (o *Logo) SetDefaltRetinaLogoURLNil(b bool)`

 SetDefaltRetinaLogoURLNil sets the value for DefaltRetinaLogoURL to be an explicit nil

### UnsetDefaltRetinaLogoURL
`func (o *Logo) UnsetDefaltRetinaLogoURL()`

UnsetDefaltRetinaLogoURL ensures that no value is present for DefaltRetinaLogoURL, not even an explicit nil
### GetDefaltStickyLogoURL

`func (o *Logo) GetDefaltStickyLogoURL() string`

GetDefaltStickyLogoURL returns the DefaltStickyLogoURL field if non-nil, zero value otherwise.

### GetDefaltStickyLogoURLOk

`func (o *Logo) GetDefaltStickyLogoURLOk() (*string, bool)`

GetDefaltStickyLogoURLOk returns a tuple with the DefaltStickyLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltStickyLogoURL

`func (o *Logo) SetDefaltStickyLogoURL(v string)`

SetDefaltStickyLogoURL sets DefaltStickyLogoURL field to given value.

### HasDefaltStickyLogoURL

`func (o *Logo) HasDefaltStickyLogoURL() bool`

HasDefaltStickyLogoURL returns a boolean if a field has been set.

### SetDefaltStickyLogoURLNil

`func (o *Logo) SetDefaltStickyLogoURLNil(b bool)`

 SetDefaltStickyLogoURLNil sets the value for DefaltStickyLogoURL to be an explicit nil

### UnsetDefaltStickyLogoURL
`func (o *Logo) UnsetDefaltStickyLogoURL()`

UnsetDefaltStickyLogoURL ensures that no value is present for DefaltStickyLogoURL, not even an explicit nil
### GetDefaltStickyRetinaLogoURL

`func (o *Logo) GetDefaltStickyRetinaLogoURL() string`

GetDefaltStickyRetinaLogoURL returns the DefaltStickyRetinaLogoURL field if non-nil, zero value otherwise.

### GetDefaltStickyRetinaLogoURLOk

`func (o *Logo) GetDefaltStickyRetinaLogoURLOk() (*string, bool)`

GetDefaltStickyRetinaLogoURLOk returns a tuple with the DefaltStickyRetinaLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltStickyRetinaLogoURL

`func (o *Logo) SetDefaltStickyRetinaLogoURL(v string)`

SetDefaltStickyRetinaLogoURL sets DefaltStickyRetinaLogoURL field to given value.

### HasDefaltStickyRetinaLogoURL

`func (o *Logo) HasDefaltStickyRetinaLogoURL() bool`

HasDefaltStickyRetinaLogoURL returns a boolean if a field has been set.

### SetDefaltStickyRetinaLogoURLNil

`func (o *Logo) SetDefaltStickyRetinaLogoURLNil(b bool)`

 SetDefaltStickyRetinaLogoURLNil sets the value for DefaltStickyRetinaLogoURL to be an explicit nil

### UnsetDefaltStickyRetinaLogoURL
`func (o *Logo) UnsetDefaltStickyRetinaLogoURL()`

UnsetDefaltStickyRetinaLogoURL ensures that no value is present for DefaltStickyRetinaLogoURL, not even an explicit nil
### GetDefaltMobileLogoURL

`func (o *Logo) GetDefaltMobileLogoURL() string`

GetDefaltMobileLogoURL returns the DefaltMobileLogoURL field if non-nil, zero value otherwise.

### GetDefaltMobileLogoURLOk

`func (o *Logo) GetDefaltMobileLogoURLOk() (*string, bool)`

GetDefaltMobileLogoURLOk returns a tuple with the DefaltMobileLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltMobileLogoURL

`func (o *Logo) SetDefaltMobileLogoURL(v string)`

SetDefaltMobileLogoURL sets DefaltMobileLogoURL field to given value.

### HasDefaltMobileLogoURL

`func (o *Logo) HasDefaltMobileLogoURL() bool`

HasDefaltMobileLogoURL returns a boolean if a field has been set.

### SetDefaltMobileLogoURLNil

`func (o *Logo) SetDefaltMobileLogoURLNil(b bool)`

 SetDefaltMobileLogoURLNil sets the value for DefaltMobileLogoURL to be an explicit nil

### UnsetDefaltMobileLogoURL
`func (o *Logo) UnsetDefaltMobileLogoURL()`

UnsetDefaltMobileLogoURL ensures that no value is present for DefaltMobileLogoURL, not even an explicit nil
### GetDefaltMobileRetinaLogoURL

`func (o *Logo) GetDefaltMobileRetinaLogoURL() string`

GetDefaltMobileRetinaLogoURL returns the DefaltMobileRetinaLogoURL field if non-nil, zero value otherwise.

### GetDefaltMobileRetinaLogoURLOk

`func (o *Logo) GetDefaltMobileRetinaLogoURLOk() (*string, bool)`

GetDefaltMobileRetinaLogoURLOk returns a tuple with the DefaltMobileRetinaLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaltMobileRetinaLogoURL

`func (o *Logo) SetDefaltMobileRetinaLogoURL(v string)`

SetDefaltMobileRetinaLogoURL sets DefaltMobileRetinaLogoURL field to given value.

### HasDefaltMobileRetinaLogoURL

`func (o *Logo) HasDefaltMobileRetinaLogoURL() bool`

HasDefaltMobileRetinaLogoURL returns a boolean if a field has been set.

### SetDefaltMobileRetinaLogoURLNil

`func (o *Logo) SetDefaltMobileRetinaLogoURLNil(b bool)`

 SetDefaltMobileRetinaLogoURLNil sets the value for DefaltMobileRetinaLogoURL to be an explicit nil

### UnsetDefaltMobileRetinaLogoURL
`func (o *Logo) UnsetDefaltMobileRetinaLogoURL()`

UnsetDefaltMobileRetinaLogoURL ensures that no value is present for DefaltMobileRetinaLogoURL, not even an explicit nil
### GetFooter

`func (o *Logo) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Logo) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Logo) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Logo) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### SetFooterNil

`func (o *Logo) SetFooterNil(b bool)`

 SetFooterNil sets the value for Footer to be an explicit nil

### UnsetFooter
`func (o *Logo) UnsetFooter()`

UnsetFooter ensures that no value is present for Footer, not even an explicit nil
### GetHeader

`func (o *Logo) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *Logo) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *Logo) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *Logo) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *Logo) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *Logo) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


