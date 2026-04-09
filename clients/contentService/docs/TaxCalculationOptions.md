# TaxCalculationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaxes** | Pointer to **bool** |  | [optional] 
**RoundTaxesAtSubtotalLevel** | Pointer to **bool** |  | [optional] 
**DisplayPriceSuffix** | Pointer to **NullableString** |  | [optional] 
**DisplayPricePrefix** | Pointer to **NullableString** |  | [optional] 
**StandardRates** | Pointer to **[]string** |  | [optional] 
**ZeroRateRates** | Pointer to **[]string** |  | [optional] 
**ReducedRateRates** | Pointer to **[]string** |  | [optional] 
**AdditionalTaxClasses** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewTaxCalculationOptions

`func NewTaxCalculationOptions() *TaxCalculationOptions`

NewTaxCalculationOptions instantiates a new TaxCalculationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCalculationOptionsWithDefaults

`func NewTaxCalculationOptionsWithDefaults() *TaxCalculationOptions`

NewTaxCalculationOptionsWithDefaults instantiates a new TaxCalculationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaxes

`func (o *TaxCalculationOptions) GetEnableTaxes() bool`

GetEnableTaxes returns the EnableTaxes field if non-nil, zero value otherwise.

### GetEnableTaxesOk

`func (o *TaxCalculationOptions) GetEnableTaxesOk() (*bool, bool)`

GetEnableTaxesOk returns a tuple with the EnableTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaxes

`func (o *TaxCalculationOptions) SetEnableTaxes(v bool)`

SetEnableTaxes sets EnableTaxes field to given value.

### HasEnableTaxes

`func (o *TaxCalculationOptions) HasEnableTaxes() bool`

HasEnableTaxes returns a boolean if a field has been set.

### GetRoundTaxesAtSubtotalLevel

`func (o *TaxCalculationOptions) GetRoundTaxesAtSubtotalLevel() bool`

GetRoundTaxesAtSubtotalLevel returns the RoundTaxesAtSubtotalLevel field if non-nil, zero value otherwise.

### GetRoundTaxesAtSubtotalLevelOk

`func (o *TaxCalculationOptions) GetRoundTaxesAtSubtotalLevelOk() (*bool, bool)`

GetRoundTaxesAtSubtotalLevelOk returns a tuple with the RoundTaxesAtSubtotalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTaxesAtSubtotalLevel

`func (o *TaxCalculationOptions) SetRoundTaxesAtSubtotalLevel(v bool)`

SetRoundTaxesAtSubtotalLevel sets RoundTaxesAtSubtotalLevel field to given value.

### HasRoundTaxesAtSubtotalLevel

`func (o *TaxCalculationOptions) HasRoundTaxesAtSubtotalLevel() bool`

HasRoundTaxesAtSubtotalLevel returns a boolean if a field has been set.

### GetDisplayPriceSuffix

`func (o *TaxCalculationOptions) GetDisplayPriceSuffix() string`

GetDisplayPriceSuffix returns the DisplayPriceSuffix field if non-nil, zero value otherwise.

### GetDisplayPriceSuffixOk

`func (o *TaxCalculationOptions) GetDisplayPriceSuffixOk() (*string, bool)`

GetDisplayPriceSuffixOk returns a tuple with the DisplayPriceSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPriceSuffix

`func (o *TaxCalculationOptions) SetDisplayPriceSuffix(v string)`

SetDisplayPriceSuffix sets DisplayPriceSuffix field to given value.

### HasDisplayPriceSuffix

`func (o *TaxCalculationOptions) HasDisplayPriceSuffix() bool`

HasDisplayPriceSuffix returns a boolean if a field has been set.

### SetDisplayPriceSuffixNil

`func (o *TaxCalculationOptions) SetDisplayPriceSuffixNil(b bool)`

 SetDisplayPriceSuffixNil sets the value for DisplayPriceSuffix to be an explicit nil

### UnsetDisplayPriceSuffix
`func (o *TaxCalculationOptions) UnsetDisplayPriceSuffix()`

UnsetDisplayPriceSuffix ensures that no value is present for DisplayPriceSuffix, not even an explicit nil
### GetDisplayPricePrefix

`func (o *TaxCalculationOptions) GetDisplayPricePrefix() string`

GetDisplayPricePrefix returns the DisplayPricePrefix field if non-nil, zero value otherwise.

### GetDisplayPricePrefixOk

`func (o *TaxCalculationOptions) GetDisplayPricePrefixOk() (*string, bool)`

GetDisplayPricePrefixOk returns a tuple with the DisplayPricePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPricePrefix

`func (o *TaxCalculationOptions) SetDisplayPricePrefix(v string)`

SetDisplayPricePrefix sets DisplayPricePrefix field to given value.

### HasDisplayPricePrefix

`func (o *TaxCalculationOptions) HasDisplayPricePrefix() bool`

HasDisplayPricePrefix returns a boolean if a field has been set.

### SetDisplayPricePrefixNil

`func (o *TaxCalculationOptions) SetDisplayPricePrefixNil(b bool)`

 SetDisplayPricePrefixNil sets the value for DisplayPricePrefix to be an explicit nil

### UnsetDisplayPricePrefix
`func (o *TaxCalculationOptions) UnsetDisplayPricePrefix()`

UnsetDisplayPricePrefix ensures that no value is present for DisplayPricePrefix, not even an explicit nil
### GetStandardRates

`func (o *TaxCalculationOptions) GetStandardRates() []string`

GetStandardRates returns the StandardRates field if non-nil, zero value otherwise.

### GetStandardRatesOk

`func (o *TaxCalculationOptions) GetStandardRatesOk() (*[]string, bool)`

GetStandardRatesOk returns a tuple with the StandardRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardRates

`func (o *TaxCalculationOptions) SetStandardRates(v []string)`

SetStandardRates sets StandardRates field to given value.

### HasStandardRates

`func (o *TaxCalculationOptions) HasStandardRates() bool`

HasStandardRates returns a boolean if a field has been set.

### SetStandardRatesNil

`func (o *TaxCalculationOptions) SetStandardRatesNil(b bool)`

 SetStandardRatesNil sets the value for StandardRates to be an explicit nil

### UnsetStandardRates
`func (o *TaxCalculationOptions) UnsetStandardRates()`

UnsetStandardRates ensures that no value is present for StandardRates, not even an explicit nil
### GetZeroRateRates

`func (o *TaxCalculationOptions) GetZeroRateRates() []string`

GetZeroRateRates returns the ZeroRateRates field if non-nil, zero value otherwise.

### GetZeroRateRatesOk

`func (o *TaxCalculationOptions) GetZeroRateRatesOk() (*[]string, bool)`

GetZeroRateRatesOk returns a tuple with the ZeroRateRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroRateRates

`func (o *TaxCalculationOptions) SetZeroRateRates(v []string)`

SetZeroRateRates sets ZeroRateRates field to given value.

### HasZeroRateRates

`func (o *TaxCalculationOptions) HasZeroRateRates() bool`

HasZeroRateRates returns a boolean if a field has been set.

### SetZeroRateRatesNil

`func (o *TaxCalculationOptions) SetZeroRateRatesNil(b bool)`

 SetZeroRateRatesNil sets the value for ZeroRateRates to be an explicit nil

### UnsetZeroRateRates
`func (o *TaxCalculationOptions) UnsetZeroRateRates()`

UnsetZeroRateRates ensures that no value is present for ZeroRateRates, not even an explicit nil
### GetReducedRateRates

`func (o *TaxCalculationOptions) GetReducedRateRates() []string`

GetReducedRateRates returns the ReducedRateRates field if non-nil, zero value otherwise.

### GetReducedRateRatesOk

`func (o *TaxCalculationOptions) GetReducedRateRatesOk() (*[]string, bool)`

GetReducedRateRatesOk returns a tuple with the ReducedRateRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedRateRates

`func (o *TaxCalculationOptions) SetReducedRateRates(v []string)`

SetReducedRateRates sets ReducedRateRates field to given value.

### HasReducedRateRates

`func (o *TaxCalculationOptions) HasReducedRateRates() bool`

HasReducedRateRates returns a boolean if a field has been set.

### SetReducedRateRatesNil

`func (o *TaxCalculationOptions) SetReducedRateRatesNil(b bool)`

 SetReducedRateRatesNil sets the value for ReducedRateRates to be an explicit nil

### UnsetReducedRateRates
`func (o *TaxCalculationOptions) UnsetReducedRateRates()`

UnsetReducedRateRates ensures that no value is present for ReducedRateRates, not even an explicit nil
### GetAdditionalTaxClasses

`func (o *TaxCalculationOptions) GetAdditionalTaxClasses() map[string][]string`

GetAdditionalTaxClasses returns the AdditionalTaxClasses field if non-nil, zero value otherwise.

### GetAdditionalTaxClassesOk

`func (o *TaxCalculationOptions) GetAdditionalTaxClassesOk() (*map[string][]string, bool)`

GetAdditionalTaxClassesOk returns a tuple with the AdditionalTaxClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTaxClasses

`func (o *TaxCalculationOptions) SetAdditionalTaxClasses(v map[string][]string)`

SetAdditionalTaxClasses sets AdditionalTaxClasses field to given value.

### HasAdditionalTaxClasses

`func (o *TaxCalculationOptions) HasAdditionalTaxClasses() bool`

HasAdditionalTaxClasses returns a boolean if a field has been set.

### SetAdditionalTaxClassesNil

`func (o *TaxCalculationOptions) SetAdditionalTaxClassesNil(b bool)`

 SetAdditionalTaxClassesNil sets the value for AdditionalTaxClasses to be an explicit nil

### UnsetAdditionalTaxClasses
`func (o *TaxCalculationOptions) UnsetAdditionalTaxClasses()`

UnsetAdditionalTaxClasses ensures that no value is present for AdditionalTaxClasses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


