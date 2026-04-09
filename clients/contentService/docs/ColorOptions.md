# ColorOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimarySkin** | Pointer to **string** |  | [optional] 
**PrimaryColor** | Pointer to **NullableString** |  | [optional] 
**SecondaryColor** | Pointer to **NullableString** |  | [optional] 
**ColorScheme** | Pointer to [**ColorScheme**](ColorScheme.md) |  | [optional] 

## Methods

### NewColorOptions

`func NewColorOptions() *ColorOptions`

NewColorOptions instantiates a new ColorOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorOptionsWithDefaults

`func NewColorOptionsWithDefaults() *ColorOptions`

NewColorOptionsWithDefaults instantiates a new ColorOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimarySkin

`func (o *ColorOptions) GetPrimarySkin() string`

GetPrimarySkin returns the PrimarySkin field if non-nil, zero value otherwise.

### GetPrimarySkinOk

`func (o *ColorOptions) GetPrimarySkinOk() (*string, bool)`

GetPrimarySkinOk returns a tuple with the PrimarySkin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySkin

`func (o *ColorOptions) SetPrimarySkin(v string)`

SetPrimarySkin sets PrimarySkin field to given value.

### HasPrimarySkin

`func (o *ColorOptions) HasPrimarySkin() bool`

HasPrimarySkin returns a boolean if a field has been set.

### GetPrimaryColor

`func (o *ColorOptions) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *ColorOptions) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *ColorOptions) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *ColorOptions) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### SetPrimaryColorNil

`func (o *ColorOptions) SetPrimaryColorNil(b bool)`

 SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil

### UnsetPrimaryColor
`func (o *ColorOptions) UnsetPrimaryColor()`

UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
### GetSecondaryColor

`func (o *ColorOptions) GetSecondaryColor() string`

GetSecondaryColor returns the SecondaryColor field if non-nil, zero value otherwise.

### GetSecondaryColorOk

`func (o *ColorOptions) GetSecondaryColorOk() (*string, bool)`

GetSecondaryColorOk returns a tuple with the SecondaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColor

`func (o *ColorOptions) SetSecondaryColor(v string)`

SetSecondaryColor sets SecondaryColor field to given value.

### HasSecondaryColor

`func (o *ColorOptions) HasSecondaryColor() bool`

HasSecondaryColor returns a boolean if a field has been set.

### SetSecondaryColorNil

`func (o *ColorOptions) SetSecondaryColorNil(b bool)`

 SetSecondaryColorNil sets the value for SecondaryColor to be an explicit nil

### UnsetSecondaryColor
`func (o *ColorOptions) UnsetSecondaryColor()`

UnsetSecondaryColor ensures that no value is present for SecondaryColor, not even an explicit nil
### GetColorScheme

`func (o *ColorOptions) GetColorScheme() ColorScheme`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *ColorOptions) GetColorSchemeOk() (*ColorScheme, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *ColorOptions) SetColorScheme(v ColorScheme)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *ColorOptions) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


