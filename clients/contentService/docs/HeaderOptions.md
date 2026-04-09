# HeaderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderPadding** | Pointer to [**Padding**](Padding.md) |  | [optional] 
**HeaderBackgroundImageURL** | Pointer to **NullableString** |  | [optional] 
**HeaderBackgroundColor** | Pointer to **NullableString** |  | [optional] 
**HeaderBorderColor** | Pointer to **NullableString** |  | [optional] 
**EnableStickyHeader** | Pointer to **bool** |  | [optional] 
**EnableHeaderShadow** | Pointer to **bool** |  | [optional] 
**EnableFullWidthHeader** | Pointer to **bool** |  | [optional] 
**HeaderLayout** | Pointer to **string** |  | [optional] 
**HeaderPosition** | Pointer to **string** |  | [optional] 
**TopHeaderContentType2** | Pointer to **string** |  | [optional] 
**TopHeaderContentType1** | Pointer to **string** |  | [optional] 
**TopHeaderBackgroundColor** | Pointer to **NullableString** |  | [optional] 
**TopHeaderContent1** | Pointer to **NullableString** |  | [optional] 
**TopHeaderContent2** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHeaderOptions

`func NewHeaderOptions() *HeaderOptions`

NewHeaderOptions instantiates a new HeaderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderOptionsWithDefaults

`func NewHeaderOptionsWithDefaults() *HeaderOptions`

NewHeaderOptionsWithDefaults instantiates a new HeaderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderPadding

`func (o *HeaderOptions) GetHeaderPadding() Padding`

GetHeaderPadding returns the HeaderPadding field if non-nil, zero value otherwise.

### GetHeaderPaddingOk

`func (o *HeaderOptions) GetHeaderPaddingOk() (*Padding, bool)`

GetHeaderPaddingOk returns a tuple with the HeaderPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPadding

`func (o *HeaderOptions) SetHeaderPadding(v Padding)`

SetHeaderPadding sets HeaderPadding field to given value.

### HasHeaderPadding

`func (o *HeaderOptions) HasHeaderPadding() bool`

HasHeaderPadding returns a boolean if a field has been set.

### GetHeaderBackgroundImageURL

`func (o *HeaderOptions) GetHeaderBackgroundImageURL() string`

GetHeaderBackgroundImageURL returns the HeaderBackgroundImageURL field if non-nil, zero value otherwise.

### GetHeaderBackgroundImageURLOk

`func (o *HeaderOptions) GetHeaderBackgroundImageURLOk() (*string, bool)`

GetHeaderBackgroundImageURLOk returns a tuple with the HeaderBackgroundImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBackgroundImageURL

`func (o *HeaderOptions) SetHeaderBackgroundImageURL(v string)`

SetHeaderBackgroundImageURL sets HeaderBackgroundImageURL field to given value.

### HasHeaderBackgroundImageURL

`func (o *HeaderOptions) HasHeaderBackgroundImageURL() bool`

HasHeaderBackgroundImageURL returns a boolean if a field has been set.

### SetHeaderBackgroundImageURLNil

`func (o *HeaderOptions) SetHeaderBackgroundImageURLNil(b bool)`

 SetHeaderBackgroundImageURLNil sets the value for HeaderBackgroundImageURL to be an explicit nil

### UnsetHeaderBackgroundImageURL
`func (o *HeaderOptions) UnsetHeaderBackgroundImageURL()`

UnsetHeaderBackgroundImageURL ensures that no value is present for HeaderBackgroundImageURL, not even an explicit nil
### GetHeaderBackgroundColor

`func (o *HeaderOptions) GetHeaderBackgroundColor() string`

GetHeaderBackgroundColor returns the HeaderBackgroundColor field if non-nil, zero value otherwise.

### GetHeaderBackgroundColorOk

`func (o *HeaderOptions) GetHeaderBackgroundColorOk() (*string, bool)`

GetHeaderBackgroundColorOk returns a tuple with the HeaderBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBackgroundColor

`func (o *HeaderOptions) SetHeaderBackgroundColor(v string)`

SetHeaderBackgroundColor sets HeaderBackgroundColor field to given value.

### HasHeaderBackgroundColor

`func (o *HeaderOptions) HasHeaderBackgroundColor() bool`

HasHeaderBackgroundColor returns a boolean if a field has been set.

### SetHeaderBackgroundColorNil

`func (o *HeaderOptions) SetHeaderBackgroundColorNil(b bool)`

 SetHeaderBackgroundColorNil sets the value for HeaderBackgroundColor to be an explicit nil

### UnsetHeaderBackgroundColor
`func (o *HeaderOptions) UnsetHeaderBackgroundColor()`

UnsetHeaderBackgroundColor ensures that no value is present for HeaderBackgroundColor, not even an explicit nil
### GetHeaderBorderColor

`func (o *HeaderOptions) GetHeaderBorderColor() string`

GetHeaderBorderColor returns the HeaderBorderColor field if non-nil, zero value otherwise.

### GetHeaderBorderColorOk

`func (o *HeaderOptions) GetHeaderBorderColorOk() (*string, bool)`

GetHeaderBorderColorOk returns a tuple with the HeaderBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBorderColor

`func (o *HeaderOptions) SetHeaderBorderColor(v string)`

SetHeaderBorderColor sets HeaderBorderColor field to given value.

### HasHeaderBorderColor

`func (o *HeaderOptions) HasHeaderBorderColor() bool`

HasHeaderBorderColor returns a boolean if a field has been set.

### SetHeaderBorderColorNil

`func (o *HeaderOptions) SetHeaderBorderColorNil(b bool)`

 SetHeaderBorderColorNil sets the value for HeaderBorderColor to be an explicit nil

### UnsetHeaderBorderColor
`func (o *HeaderOptions) UnsetHeaderBorderColor()`

UnsetHeaderBorderColor ensures that no value is present for HeaderBorderColor, not even an explicit nil
### GetEnableStickyHeader

`func (o *HeaderOptions) GetEnableStickyHeader() bool`

GetEnableStickyHeader returns the EnableStickyHeader field if non-nil, zero value otherwise.

### GetEnableStickyHeaderOk

`func (o *HeaderOptions) GetEnableStickyHeaderOk() (*bool, bool)`

GetEnableStickyHeaderOk returns a tuple with the EnableStickyHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStickyHeader

`func (o *HeaderOptions) SetEnableStickyHeader(v bool)`

SetEnableStickyHeader sets EnableStickyHeader field to given value.

### HasEnableStickyHeader

`func (o *HeaderOptions) HasEnableStickyHeader() bool`

HasEnableStickyHeader returns a boolean if a field has been set.

### GetEnableHeaderShadow

`func (o *HeaderOptions) GetEnableHeaderShadow() bool`

GetEnableHeaderShadow returns the EnableHeaderShadow field if non-nil, zero value otherwise.

### GetEnableHeaderShadowOk

`func (o *HeaderOptions) GetEnableHeaderShadowOk() (*bool, bool)`

GetEnableHeaderShadowOk returns a tuple with the EnableHeaderShadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHeaderShadow

`func (o *HeaderOptions) SetEnableHeaderShadow(v bool)`

SetEnableHeaderShadow sets EnableHeaderShadow field to given value.

### HasEnableHeaderShadow

`func (o *HeaderOptions) HasEnableHeaderShadow() bool`

HasEnableHeaderShadow returns a boolean if a field has been set.

### GetEnableFullWidthHeader

`func (o *HeaderOptions) GetEnableFullWidthHeader() bool`

GetEnableFullWidthHeader returns the EnableFullWidthHeader field if non-nil, zero value otherwise.

### GetEnableFullWidthHeaderOk

`func (o *HeaderOptions) GetEnableFullWidthHeaderOk() (*bool, bool)`

GetEnableFullWidthHeaderOk returns a tuple with the EnableFullWidthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullWidthHeader

`func (o *HeaderOptions) SetEnableFullWidthHeader(v bool)`

SetEnableFullWidthHeader sets EnableFullWidthHeader field to given value.

### HasEnableFullWidthHeader

`func (o *HeaderOptions) HasEnableFullWidthHeader() bool`

HasEnableFullWidthHeader returns a boolean if a field has been set.

### GetHeaderLayout

`func (o *HeaderOptions) GetHeaderLayout() string`

GetHeaderLayout returns the HeaderLayout field if non-nil, zero value otherwise.

### GetHeaderLayoutOk

`func (o *HeaderOptions) GetHeaderLayoutOk() (*string, bool)`

GetHeaderLayoutOk returns a tuple with the HeaderLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderLayout

`func (o *HeaderOptions) SetHeaderLayout(v string)`

SetHeaderLayout sets HeaderLayout field to given value.

### HasHeaderLayout

`func (o *HeaderOptions) HasHeaderLayout() bool`

HasHeaderLayout returns a boolean if a field has been set.

### GetHeaderPosition

`func (o *HeaderOptions) GetHeaderPosition() string`

GetHeaderPosition returns the HeaderPosition field if non-nil, zero value otherwise.

### GetHeaderPositionOk

`func (o *HeaderOptions) GetHeaderPositionOk() (*string, bool)`

GetHeaderPositionOk returns a tuple with the HeaderPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPosition

`func (o *HeaderOptions) SetHeaderPosition(v string)`

SetHeaderPosition sets HeaderPosition field to given value.

### HasHeaderPosition

`func (o *HeaderOptions) HasHeaderPosition() bool`

HasHeaderPosition returns a boolean if a field has been set.

### GetTopHeaderContentType2

`func (o *HeaderOptions) GetTopHeaderContentType2() string`

GetTopHeaderContentType2 returns the TopHeaderContentType2 field if non-nil, zero value otherwise.

### GetTopHeaderContentType2Ok

`func (o *HeaderOptions) GetTopHeaderContentType2Ok() (*string, bool)`

GetTopHeaderContentType2Ok returns a tuple with the TopHeaderContentType2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHeaderContentType2

`func (o *HeaderOptions) SetTopHeaderContentType2(v string)`

SetTopHeaderContentType2 sets TopHeaderContentType2 field to given value.

### HasTopHeaderContentType2

`func (o *HeaderOptions) HasTopHeaderContentType2() bool`

HasTopHeaderContentType2 returns a boolean if a field has been set.

### GetTopHeaderContentType1

`func (o *HeaderOptions) GetTopHeaderContentType1() string`

GetTopHeaderContentType1 returns the TopHeaderContentType1 field if non-nil, zero value otherwise.

### GetTopHeaderContentType1Ok

`func (o *HeaderOptions) GetTopHeaderContentType1Ok() (*string, bool)`

GetTopHeaderContentType1Ok returns a tuple with the TopHeaderContentType1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHeaderContentType1

`func (o *HeaderOptions) SetTopHeaderContentType1(v string)`

SetTopHeaderContentType1 sets TopHeaderContentType1 field to given value.

### HasTopHeaderContentType1

`func (o *HeaderOptions) HasTopHeaderContentType1() bool`

HasTopHeaderContentType1 returns a boolean if a field has been set.

### GetTopHeaderBackgroundColor

`func (o *HeaderOptions) GetTopHeaderBackgroundColor() string`

GetTopHeaderBackgroundColor returns the TopHeaderBackgroundColor field if non-nil, zero value otherwise.

### GetTopHeaderBackgroundColorOk

`func (o *HeaderOptions) GetTopHeaderBackgroundColorOk() (*string, bool)`

GetTopHeaderBackgroundColorOk returns a tuple with the TopHeaderBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHeaderBackgroundColor

`func (o *HeaderOptions) SetTopHeaderBackgroundColor(v string)`

SetTopHeaderBackgroundColor sets TopHeaderBackgroundColor field to given value.

### HasTopHeaderBackgroundColor

`func (o *HeaderOptions) HasTopHeaderBackgroundColor() bool`

HasTopHeaderBackgroundColor returns a boolean if a field has been set.

### SetTopHeaderBackgroundColorNil

`func (o *HeaderOptions) SetTopHeaderBackgroundColorNil(b bool)`

 SetTopHeaderBackgroundColorNil sets the value for TopHeaderBackgroundColor to be an explicit nil

### UnsetTopHeaderBackgroundColor
`func (o *HeaderOptions) UnsetTopHeaderBackgroundColor()`

UnsetTopHeaderBackgroundColor ensures that no value is present for TopHeaderBackgroundColor, not even an explicit nil
### GetTopHeaderContent1

`func (o *HeaderOptions) GetTopHeaderContent1() string`

GetTopHeaderContent1 returns the TopHeaderContent1 field if non-nil, zero value otherwise.

### GetTopHeaderContent1Ok

`func (o *HeaderOptions) GetTopHeaderContent1Ok() (*string, bool)`

GetTopHeaderContent1Ok returns a tuple with the TopHeaderContent1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHeaderContent1

`func (o *HeaderOptions) SetTopHeaderContent1(v string)`

SetTopHeaderContent1 sets TopHeaderContent1 field to given value.

### HasTopHeaderContent1

`func (o *HeaderOptions) HasTopHeaderContent1() bool`

HasTopHeaderContent1 returns a boolean if a field has been set.

### SetTopHeaderContent1Nil

`func (o *HeaderOptions) SetTopHeaderContent1Nil(b bool)`

 SetTopHeaderContent1Nil sets the value for TopHeaderContent1 to be an explicit nil

### UnsetTopHeaderContent1
`func (o *HeaderOptions) UnsetTopHeaderContent1()`

UnsetTopHeaderContent1 ensures that no value is present for TopHeaderContent1, not even an explicit nil
### GetTopHeaderContent2

`func (o *HeaderOptions) GetTopHeaderContent2() string`

GetTopHeaderContent2 returns the TopHeaderContent2 field if non-nil, zero value otherwise.

### GetTopHeaderContent2Ok

`func (o *HeaderOptions) GetTopHeaderContent2Ok() (*string, bool)`

GetTopHeaderContent2Ok returns a tuple with the TopHeaderContent2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHeaderContent2

`func (o *HeaderOptions) SetTopHeaderContent2(v string)`

SetTopHeaderContent2 sets TopHeaderContent2 field to given value.

### HasTopHeaderContent2

`func (o *HeaderOptions) HasTopHeaderContent2() bool`

HasTopHeaderContent2 returns a boolean if a field has been set.

### SetTopHeaderContent2Nil

`func (o *HeaderOptions) SetTopHeaderContent2Nil(b bool)`

 SetTopHeaderContent2Nil sets the value for TopHeaderContent2 to be an explicit nil

### UnsetTopHeaderContent2
`func (o *HeaderOptions) UnsetTopHeaderContent2()`

UnsetTopHeaderContent2 ensures that no value is present for TopHeaderContent2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


