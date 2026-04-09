# SlidingBarOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **string** |  | [optional] 
**ContentPadding** | Pointer to [**Padding**](Padding.md) |  | [optional] 
**ContentAlignment** | Pointer to **string** |  | [optional] 
**ColumnsCount** | Pointer to **int32** |  | [optional] 
**EnableSticky** | Pointer to **bool** |  | [optional] 
**OpenOnPageLoad** | Pointer to **bool** |  | [optional] 
**EnableOnMobile** | Pointer to **bool** |  | [optional] 
**EnableOnDesktop** | Pointer to **bool** |  | [optional] 

## Methods

### NewSlidingBarOptions

`func NewSlidingBarOptions() *SlidingBarOptions`

NewSlidingBarOptions instantiates a new SlidingBarOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlidingBarOptionsWithDefaults

`func NewSlidingBarOptionsWithDefaults() *SlidingBarOptions`

NewSlidingBarOptionsWithDefaults instantiates a new SlidingBarOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *SlidingBarOptions) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SlidingBarOptions) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SlidingBarOptions) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SlidingBarOptions) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetContentPadding

`func (o *SlidingBarOptions) GetContentPadding() Padding`

GetContentPadding returns the ContentPadding field if non-nil, zero value otherwise.

### GetContentPaddingOk

`func (o *SlidingBarOptions) GetContentPaddingOk() (*Padding, bool)`

GetContentPaddingOk returns a tuple with the ContentPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPadding

`func (o *SlidingBarOptions) SetContentPadding(v Padding)`

SetContentPadding sets ContentPadding field to given value.

### HasContentPadding

`func (o *SlidingBarOptions) HasContentPadding() bool`

HasContentPadding returns a boolean if a field has been set.

### GetContentAlignment

`func (o *SlidingBarOptions) GetContentAlignment() string`

GetContentAlignment returns the ContentAlignment field if non-nil, zero value otherwise.

### GetContentAlignmentOk

`func (o *SlidingBarOptions) GetContentAlignmentOk() (*string, bool)`

GetContentAlignmentOk returns a tuple with the ContentAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAlignment

`func (o *SlidingBarOptions) SetContentAlignment(v string)`

SetContentAlignment sets ContentAlignment field to given value.

### HasContentAlignment

`func (o *SlidingBarOptions) HasContentAlignment() bool`

HasContentAlignment returns a boolean if a field has been set.

### GetColumnsCount

`func (o *SlidingBarOptions) GetColumnsCount() int32`

GetColumnsCount returns the ColumnsCount field if non-nil, zero value otherwise.

### GetColumnsCountOk

`func (o *SlidingBarOptions) GetColumnsCountOk() (*int32, bool)`

GetColumnsCountOk returns a tuple with the ColumnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsCount

`func (o *SlidingBarOptions) SetColumnsCount(v int32)`

SetColumnsCount sets ColumnsCount field to given value.

### HasColumnsCount

`func (o *SlidingBarOptions) HasColumnsCount() bool`

HasColumnsCount returns a boolean if a field has been set.

### GetEnableSticky

`func (o *SlidingBarOptions) GetEnableSticky() bool`

GetEnableSticky returns the EnableSticky field if non-nil, zero value otherwise.

### GetEnableStickyOk

`func (o *SlidingBarOptions) GetEnableStickyOk() (*bool, bool)`

GetEnableStickyOk returns a tuple with the EnableSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSticky

`func (o *SlidingBarOptions) SetEnableSticky(v bool)`

SetEnableSticky sets EnableSticky field to given value.

### HasEnableSticky

`func (o *SlidingBarOptions) HasEnableSticky() bool`

HasEnableSticky returns a boolean if a field has been set.

### GetOpenOnPageLoad

`func (o *SlidingBarOptions) GetOpenOnPageLoad() bool`

GetOpenOnPageLoad returns the OpenOnPageLoad field if non-nil, zero value otherwise.

### GetOpenOnPageLoadOk

`func (o *SlidingBarOptions) GetOpenOnPageLoadOk() (*bool, bool)`

GetOpenOnPageLoadOk returns a tuple with the OpenOnPageLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOnPageLoad

`func (o *SlidingBarOptions) SetOpenOnPageLoad(v bool)`

SetOpenOnPageLoad sets OpenOnPageLoad field to given value.

### HasOpenOnPageLoad

`func (o *SlidingBarOptions) HasOpenOnPageLoad() bool`

HasOpenOnPageLoad returns a boolean if a field has been set.

### GetEnableOnMobile

`func (o *SlidingBarOptions) GetEnableOnMobile() bool`

GetEnableOnMobile returns the EnableOnMobile field if non-nil, zero value otherwise.

### GetEnableOnMobileOk

`func (o *SlidingBarOptions) GetEnableOnMobileOk() (*bool, bool)`

GetEnableOnMobileOk returns a tuple with the EnableOnMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnMobile

`func (o *SlidingBarOptions) SetEnableOnMobile(v bool)`

SetEnableOnMobile sets EnableOnMobile field to given value.

### HasEnableOnMobile

`func (o *SlidingBarOptions) HasEnableOnMobile() bool`

HasEnableOnMobile returns a boolean if a field has been set.

### GetEnableOnDesktop

`func (o *SlidingBarOptions) GetEnableOnDesktop() bool`

GetEnableOnDesktop returns the EnableOnDesktop field if non-nil, zero value otherwise.

### GetEnableOnDesktopOk

`func (o *SlidingBarOptions) GetEnableOnDesktopOk() (*bool, bool)`

GetEnableOnDesktopOk returns a tuple with the EnableOnDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnDesktop

`func (o *SlidingBarOptions) SetEnableOnDesktop(v bool)`

SetEnableOnDesktop sets EnableOnDesktop field to given value.

### HasEnableOnDesktop

`func (o *SlidingBarOptions) HasEnableOnDesktop() bool`

HasEnableOnDesktop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


