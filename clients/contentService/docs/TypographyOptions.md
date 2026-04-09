# TypographyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyTypography** | Pointer to [**Typography**](Typography.md) |  | [optional] 
**HeadersTypography** | Pointer to [**Typography**](Typography.md) |  | [optional] 
**CustomFonts** | Pointer to [**[]CustomFont**](CustomFont.md) |  | [optional] 

## Methods

### NewTypographyOptions

`func NewTypographyOptions() *TypographyOptions`

NewTypographyOptions instantiates a new TypographyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypographyOptionsWithDefaults

`func NewTypographyOptionsWithDefaults() *TypographyOptions`

NewTypographyOptionsWithDefaults instantiates a new TypographyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyTypography

`func (o *TypographyOptions) GetBodyTypography() Typography`

GetBodyTypography returns the BodyTypography field if non-nil, zero value otherwise.

### GetBodyTypographyOk

`func (o *TypographyOptions) GetBodyTypographyOk() (*Typography, bool)`

GetBodyTypographyOk returns a tuple with the BodyTypography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTypography

`func (o *TypographyOptions) SetBodyTypography(v Typography)`

SetBodyTypography sets BodyTypography field to given value.

### HasBodyTypography

`func (o *TypographyOptions) HasBodyTypography() bool`

HasBodyTypography returns a boolean if a field has been set.

### GetHeadersTypography

`func (o *TypographyOptions) GetHeadersTypography() Typography`

GetHeadersTypography returns the HeadersTypography field if non-nil, zero value otherwise.

### GetHeadersTypographyOk

`func (o *TypographyOptions) GetHeadersTypographyOk() (*Typography, bool)`

GetHeadersTypographyOk returns a tuple with the HeadersTypography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersTypography

`func (o *TypographyOptions) SetHeadersTypography(v Typography)`

SetHeadersTypography sets HeadersTypography field to given value.

### HasHeadersTypography

`func (o *TypographyOptions) HasHeadersTypography() bool`

HasHeadersTypography returns a boolean if a field has been set.

### GetCustomFonts

`func (o *TypographyOptions) GetCustomFonts() []CustomFont`

GetCustomFonts returns the CustomFonts field if non-nil, zero value otherwise.

### GetCustomFontsOk

`func (o *TypographyOptions) GetCustomFontsOk() (*[]CustomFont, bool)`

GetCustomFontsOk returns a tuple with the CustomFonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFonts

`func (o *TypographyOptions) SetCustomFonts(v []CustomFont)`

SetCustomFonts sets CustomFonts field to given value.

### HasCustomFonts

`func (o *TypographyOptions) HasCustomFonts() bool`

HasCustomFonts returns a boolean if a field has been set.

### SetCustomFontsNil

`func (o *TypographyOptions) SetCustomFontsNil(b bool)`

 SetCustomFontsNil sets the value for CustomFonts to be an explicit nil

### UnsetCustomFonts
`func (o *TypographyOptions) UnsetCustomFonts()`

UnsetCustomFonts ensures that no value is present for CustomFonts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


