# FooterOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfColumns** | Pointer to **int32** |  | [optional] 
**EnableWidgets** | Pointer to **bool** |  | [optional] 
**EnableCopyrightBar** | Pointer to **bool** |  | [optional] 
**CenterWidgetsContent** | Pointer to **bool** |  | [optional] 
**CenterCopyrightContent** | Pointer to **bool** |  | [optional] 
**EnableVerticalWidgetDividerLine** | Pointer to **bool** |  | [optional] 
**VerticalWidgetDividerLineSize** | Pointer to **int32** |  | [optional] 
**CopyrightText** | Pointer to **NullableString** |  | [optional] 
**CopyrightBackgroundColor** | Pointer to **NullableString** |  | [optional] 
**BorderSize** | Pointer to **int32** |  | [optional] 
**BorderColor** | Pointer to **NullableString** |  | [optional] 
**WidgetDividerColor** | Pointer to **NullableString** |  | [optional] 
**WidgetDivider** | Pointer to **string** |  | [optional] 
**CopyrightPadding** | Pointer to [**Padding**](Padding.md) |  | [optional] 
**WidgetsAreaPadding** | Pointer to [**Padding**](Padding.md) |  | [optional] 
**FooterAreaPadding** | Pointer to [**Padding**](Padding.md) |  | [optional] 
**FooterBackground** | Pointer to [**Background**](Background.md) |  | [optional] 
**CopyrightBackground** | Pointer to [**Background**](Background.md) |  | [optional] 
**HeadingsTypography** | Pointer to [**Typography**](Typography.md) |  | [optional] 
**WidgetsTypography** | Pointer to [**Typography**](Typography.md) |  | [optional] 
**CopyrightTypography** | Pointer to [**Typography**](Typography.md) |  | [optional] 

## Methods

### NewFooterOptions

`func NewFooterOptions() *FooterOptions`

NewFooterOptions instantiates a new FooterOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFooterOptionsWithDefaults

`func NewFooterOptionsWithDefaults() *FooterOptions`

NewFooterOptionsWithDefaults instantiates a new FooterOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfColumns

`func (o *FooterOptions) GetNumberOfColumns() int32`

GetNumberOfColumns returns the NumberOfColumns field if non-nil, zero value otherwise.

### GetNumberOfColumnsOk

`func (o *FooterOptions) GetNumberOfColumnsOk() (*int32, bool)`

GetNumberOfColumnsOk returns a tuple with the NumberOfColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfColumns

`func (o *FooterOptions) SetNumberOfColumns(v int32)`

SetNumberOfColumns sets NumberOfColumns field to given value.

### HasNumberOfColumns

`func (o *FooterOptions) HasNumberOfColumns() bool`

HasNumberOfColumns returns a boolean if a field has been set.

### GetEnableWidgets

`func (o *FooterOptions) GetEnableWidgets() bool`

GetEnableWidgets returns the EnableWidgets field if non-nil, zero value otherwise.

### GetEnableWidgetsOk

`func (o *FooterOptions) GetEnableWidgetsOk() (*bool, bool)`

GetEnableWidgetsOk returns a tuple with the EnableWidgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWidgets

`func (o *FooterOptions) SetEnableWidgets(v bool)`

SetEnableWidgets sets EnableWidgets field to given value.

### HasEnableWidgets

`func (o *FooterOptions) HasEnableWidgets() bool`

HasEnableWidgets returns a boolean if a field has been set.

### GetEnableCopyrightBar

`func (o *FooterOptions) GetEnableCopyrightBar() bool`

GetEnableCopyrightBar returns the EnableCopyrightBar field if non-nil, zero value otherwise.

### GetEnableCopyrightBarOk

`func (o *FooterOptions) GetEnableCopyrightBarOk() (*bool, bool)`

GetEnableCopyrightBarOk returns a tuple with the EnableCopyrightBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyrightBar

`func (o *FooterOptions) SetEnableCopyrightBar(v bool)`

SetEnableCopyrightBar sets EnableCopyrightBar field to given value.

### HasEnableCopyrightBar

`func (o *FooterOptions) HasEnableCopyrightBar() bool`

HasEnableCopyrightBar returns a boolean if a field has been set.

### GetCenterWidgetsContent

`func (o *FooterOptions) GetCenterWidgetsContent() bool`

GetCenterWidgetsContent returns the CenterWidgetsContent field if non-nil, zero value otherwise.

### GetCenterWidgetsContentOk

`func (o *FooterOptions) GetCenterWidgetsContentOk() (*bool, bool)`

GetCenterWidgetsContentOk returns a tuple with the CenterWidgetsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterWidgetsContent

`func (o *FooterOptions) SetCenterWidgetsContent(v bool)`

SetCenterWidgetsContent sets CenterWidgetsContent field to given value.

### HasCenterWidgetsContent

`func (o *FooterOptions) HasCenterWidgetsContent() bool`

HasCenterWidgetsContent returns a boolean if a field has been set.

### GetCenterCopyrightContent

`func (o *FooterOptions) GetCenterCopyrightContent() bool`

GetCenterCopyrightContent returns the CenterCopyrightContent field if non-nil, zero value otherwise.

### GetCenterCopyrightContentOk

`func (o *FooterOptions) GetCenterCopyrightContentOk() (*bool, bool)`

GetCenterCopyrightContentOk returns a tuple with the CenterCopyrightContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterCopyrightContent

`func (o *FooterOptions) SetCenterCopyrightContent(v bool)`

SetCenterCopyrightContent sets CenterCopyrightContent field to given value.

### HasCenterCopyrightContent

`func (o *FooterOptions) HasCenterCopyrightContent() bool`

HasCenterCopyrightContent returns a boolean if a field has been set.

### GetEnableVerticalWidgetDividerLine

`func (o *FooterOptions) GetEnableVerticalWidgetDividerLine() bool`

GetEnableVerticalWidgetDividerLine returns the EnableVerticalWidgetDividerLine field if non-nil, zero value otherwise.

### GetEnableVerticalWidgetDividerLineOk

`func (o *FooterOptions) GetEnableVerticalWidgetDividerLineOk() (*bool, bool)`

GetEnableVerticalWidgetDividerLineOk returns a tuple with the EnableVerticalWidgetDividerLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVerticalWidgetDividerLine

`func (o *FooterOptions) SetEnableVerticalWidgetDividerLine(v bool)`

SetEnableVerticalWidgetDividerLine sets EnableVerticalWidgetDividerLine field to given value.

### HasEnableVerticalWidgetDividerLine

`func (o *FooterOptions) HasEnableVerticalWidgetDividerLine() bool`

HasEnableVerticalWidgetDividerLine returns a boolean if a field has been set.

### GetVerticalWidgetDividerLineSize

`func (o *FooterOptions) GetVerticalWidgetDividerLineSize() int32`

GetVerticalWidgetDividerLineSize returns the VerticalWidgetDividerLineSize field if non-nil, zero value otherwise.

### GetVerticalWidgetDividerLineSizeOk

`func (o *FooterOptions) GetVerticalWidgetDividerLineSizeOk() (*int32, bool)`

GetVerticalWidgetDividerLineSizeOk returns a tuple with the VerticalWidgetDividerLineSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalWidgetDividerLineSize

`func (o *FooterOptions) SetVerticalWidgetDividerLineSize(v int32)`

SetVerticalWidgetDividerLineSize sets VerticalWidgetDividerLineSize field to given value.

### HasVerticalWidgetDividerLineSize

`func (o *FooterOptions) HasVerticalWidgetDividerLineSize() bool`

HasVerticalWidgetDividerLineSize returns a boolean if a field has been set.

### GetCopyrightText

`func (o *FooterOptions) GetCopyrightText() string`

GetCopyrightText returns the CopyrightText field if non-nil, zero value otherwise.

### GetCopyrightTextOk

`func (o *FooterOptions) GetCopyrightTextOk() (*string, bool)`

GetCopyrightTextOk returns a tuple with the CopyrightText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightText

`func (o *FooterOptions) SetCopyrightText(v string)`

SetCopyrightText sets CopyrightText field to given value.

### HasCopyrightText

`func (o *FooterOptions) HasCopyrightText() bool`

HasCopyrightText returns a boolean if a field has been set.

### SetCopyrightTextNil

`func (o *FooterOptions) SetCopyrightTextNil(b bool)`

 SetCopyrightTextNil sets the value for CopyrightText to be an explicit nil

### UnsetCopyrightText
`func (o *FooterOptions) UnsetCopyrightText()`

UnsetCopyrightText ensures that no value is present for CopyrightText, not even an explicit nil
### GetCopyrightBackgroundColor

`func (o *FooterOptions) GetCopyrightBackgroundColor() string`

GetCopyrightBackgroundColor returns the CopyrightBackgroundColor field if non-nil, zero value otherwise.

### GetCopyrightBackgroundColorOk

`func (o *FooterOptions) GetCopyrightBackgroundColorOk() (*string, bool)`

GetCopyrightBackgroundColorOk returns a tuple with the CopyrightBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightBackgroundColor

`func (o *FooterOptions) SetCopyrightBackgroundColor(v string)`

SetCopyrightBackgroundColor sets CopyrightBackgroundColor field to given value.

### HasCopyrightBackgroundColor

`func (o *FooterOptions) HasCopyrightBackgroundColor() bool`

HasCopyrightBackgroundColor returns a boolean if a field has been set.

### SetCopyrightBackgroundColorNil

`func (o *FooterOptions) SetCopyrightBackgroundColorNil(b bool)`

 SetCopyrightBackgroundColorNil sets the value for CopyrightBackgroundColor to be an explicit nil

### UnsetCopyrightBackgroundColor
`func (o *FooterOptions) UnsetCopyrightBackgroundColor()`

UnsetCopyrightBackgroundColor ensures that no value is present for CopyrightBackgroundColor, not even an explicit nil
### GetBorderSize

`func (o *FooterOptions) GetBorderSize() int32`

GetBorderSize returns the BorderSize field if non-nil, zero value otherwise.

### GetBorderSizeOk

`func (o *FooterOptions) GetBorderSizeOk() (*int32, bool)`

GetBorderSizeOk returns a tuple with the BorderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderSize

`func (o *FooterOptions) SetBorderSize(v int32)`

SetBorderSize sets BorderSize field to given value.

### HasBorderSize

`func (o *FooterOptions) HasBorderSize() bool`

HasBorderSize returns a boolean if a field has been set.

### GetBorderColor

`func (o *FooterOptions) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *FooterOptions) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *FooterOptions) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *FooterOptions) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### SetBorderColorNil

`func (o *FooterOptions) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *FooterOptions) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetWidgetDividerColor

`func (o *FooterOptions) GetWidgetDividerColor() string`

GetWidgetDividerColor returns the WidgetDividerColor field if non-nil, zero value otherwise.

### GetWidgetDividerColorOk

`func (o *FooterOptions) GetWidgetDividerColorOk() (*string, bool)`

GetWidgetDividerColorOk returns a tuple with the WidgetDividerColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetDividerColor

`func (o *FooterOptions) SetWidgetDividerColor(v string)`

SetWidgetDividerColor sets WidgetDividerColor field to given value.

### HasWidgetDividerColor

`func (o *FooterOptions) HasWidgetDividerColor() bool`

HasWidgetDividerColor returns a boolean if a field has been set.

### SetWidgetDividerColorNil

`func (o *FooterOptions) SetWidgetDividerColorNil(b bool)`

 SetWidgetDividerColorNil sets the value for WidgetDividerColor to be an explicit nil

### UnsetWidgetDividerColor
`func (o *FooterOptions) UnsetWidgetDividerColor()`

UnsetWidgetDividerColor ensures that no value is present for WidgetDividerColor, not even an explicit nil
### GetWidgetDivider

`func (o *FooterOptions) GetWidgetDivider() string`

GetWidgetDivider returns the WidgetDivider field if non-nil, zero value otherwise.

### GetWidgetDividerOk

`func (o *FooterOptions) GetWidgetDividerOk() (*string, bool)`

GetWidgetDividerOk returns a tuple with the WidgetDivider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetDivider

`func (o *FooterOptions) SetWidgetDivider(v string)`

SetWidgetDivider sets WidgetDivider field to given value.

### HasWidgetDivider

`func (o *FooterOptions) HasWidgetDivider() bool`

HasWidgetDivider returns a boolean if a field has been set.

### GetCopyrightPadding

`func (o *FooterOptions) GetCopyrightPadding() Padding`

GetCopyrightPadding returns the CopyrightPadding field if non-nil, zero value otherwise.

### GetCopyrightPaddingOk

`func (o *FooterOptions) GetCopyrightPaddingOk() (*Padding, bool)`

GetCopyrightPaddingOk returns a tuple with the CopyrightPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightPadding

`func (o *FooterOptions) SetCopyrightPadding(v Padding)`

SetCopyrightPadding sets CopyrightPadding field to given value.

### HasCopyrightPadding

`func (o *FooterOptions) HasCopyrightPadding() bool`

HasCopyrightPadding returns a boolean if a field has been set.

### GetWidgetsAreaPadding

`func (o *FooterOptions) GetWidgetsAreaPadding() Padding`

GetWidgetsAreaPadding returns the WidgetsAreaPadding field if non-nil, zero value otherwise.

### GetWidgetsAreaPaddingOk

`func (o *FooterOptions) GetWidgetsAreaPaddingOk() (*Padding, bool)`

GetWidgetsAreaPaddingOk returns a tuple with the WidgetsAreaPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetsAreaPadding

`func (o *FooterOptions) SetWidgetsAreaPadding(v Padding)`

SetWidgetsAreaPadding sets WidgetsAreaPadding field to given value.

### HasWidgetsAreaPadding

`func (o *FooterOptions) HasWidgetsAreaPadding() bool`

HasWidgetsAreaPadding returns a boolean if a field has been set.

### GetFooterAreaPadding

`func (o *FooterOptions) GetFooterAreaPadding() Padding`

GetFooterAreaPadding returns the FooterAreaPadding field if non-nil, zero value otherwise.

### GetFooterAreaPaddingOk

`func (o *FooterOptions) GetFooterAreaPaddingOk() (*Padding, bool)`

GetFooterAreaPaddingOk returns a tuple with the FooterAreaPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterAreaPadding

`func (o *FooterOptions) SetFooterAreaPadding(v Padding)`

SetFooterAreaPadding sets FooterAreaPadding field to given value.

### HasFooterAreaPadding

`func (o *FooterOptions) HasFooterAreaPadding() bool`

HasFooterAreaPadding returns a boolean if a field has been set.

### GetFooterBackground

`func (o *FooterOptions) GetFooterBackground() Background`

GetFooterBackground returns the FooterBackground field if non-nil, zero value otherwise.

### GetFooterBackgroundOk

`func (o *FooterOptions) GetFooterBackgroundOk() (*Background, bool)`

GetFooterBackgroundOk returns a tuple with the FooterBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterBackground

`func (o *FooterOptions) SetFooterBackground(v Background)`

SetFooterBackground sets FooterBackground field to given value.

### HasFooterBackground

`func (o *FooterOptions) HasFooterBackground() bool`

HasFooterBackground returns a boolean if a field has been set.

### GetCopyrightBackground

`func (o *FooterOptions) GetCopyrightBackground() Background`

GetCopyrightBackground returns the CopyrightBackground field if non-nil, zero value otherwise.

### GetCopyrightBackgroundOk

`func (o *FooterOptions) GetCopyrightBackgroundOk() (*Background, bool)`

GetCopyrightBackgroundOk returns a tuple with the CopyrightBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightBackground

`func (o *FooterOptions) SetCopyrightBackground(v Background)`

SetCopyrightBackground sets CopyrightBackground field to given value.

### HasCopyrightBackground

`func (o *FooterOptions) HasCopyrightBackground() bool`

HasCopyrightBackground returns a boolean if a field has been set.

### GetHeadingsTypography

`func (o *FooterOptions) GetHeadingsTypography() Typography`

GetHeadingsTypography returns the HeadingsTypography field if non-nil, zero value otherwise.

### GetHeadingsTypographyOk

`func (o *FooterOptions) GetHeadingsTypographyOk() (*Typography, bool)`

GetHeadingsTypographyOk returns a tuple with the HeadingsTypography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadingsTypography

`func (o *FooterOptions) SetHeadingsTypography(v Typography)`

SetHeadingsTypography sets HeadingsTypography field to given value.

### HasHeadingsTypography

`func (o *FooterOptions) HasHeadingsTypography() bool`

HasHeadingsTypography returns a boolean if a field has been set.

### GetWidgetsTypography

`func (o *FooterOptions) GetWidgetsTypography() Typography`

GetWidgetsTypography returns the WidgetsTypography field if non-nil, zero value otherwise.

### GetWidgetsTypographyOk

`func (o *FooterOptions) GetWidgetsTypographyOk() (*Typography, bool)`

GetWidgetsTypographyOk returns a tuple with the WidgetsTypography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetsTypography

`func (o *FooterOptions) SetWidgetsTypography(v Typography)`

SetWidgetsTypography sets WidgetsTypography field to given value.

### HasWidgetsTypography

`func (o *FooterOptions) HasWidgetsTypography() bool`

HasWidgetsTypography returns a boolean if a field has been set.

### GetCopyrightTypography

`func (o *FooterOptions) GetCopyrightTypography() Typography`

GetCopyrightTypography returns the CopyrightTypography field if non-nil, zero value otherwise.

### GetCopyrightTypographyOk

`func (o *FooterOptions) GetCopyrightTypographyOk() (*Typography, bool)`

GetCopyrightTypographyOk returns a tuple with the CopyrightTypography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTypography

`func (o *FooterOptions) SetCopyrightTypography(v Typography)`

SetCopyrightTypography sets CopyrightTypography field to given value.

### HasCopyrightTypography

`func (o *FooterOptions) HasCopyrightTypography() bool`

HasCopyrightTypography returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


