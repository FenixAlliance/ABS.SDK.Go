# StoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**FooterLogo** | Pointer to **NullableString** |  | [optional] 
**Tagline** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AddressLine1** | Pointer to **NullableString** |  | [optional] 
**AddressLine2** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**NumberOfDecimals** | Pointer to **int32** |  | [optional] 
**DecimalSeparator** | Pointer to **NullableString** |  | [optional] 
**SellToAllCountries** | Pointer to **bool** |  | [optional] 
**CartOptions** | Pointer to [**CartOptions**](CartOptions.md) |  | [optional] 
**EmailOptions** | Pointer to [**EmailOptions**](EmailOptions.md) |  | [optional] 
**CouponsOptions** | Pointer to [**CouponsOptions**](CouponsOptions.md) |  | [optional] 
**PaymentOptions** | Pointer to [**PaymentOptions**](PaymentOptions.md) |  | [optional] 
**ProductOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ReviewsOptions** | Pointer to [**ReviewsOptions**](ReviewsOptions.md) |  | [optional] 
**AdvancedOptions** | Pointer to [**AdvancedOptions**](AdvancedOptions.md) |  | [optional] 
**ServicesOptions** | Pointer to [**ServicesOptions**](ServicesOptions.md) |  | [optional] 
**InventoryOptions** | Pointer to [**InventoryOptions**](InventoryOptions.md) |  | [optional] 
**IntegrationOptions** | Pointer to [**IntegrationOptions**](IntegrationOptions.md) |  | [optional] 
**MeasurementOptions** | Pointer to [**MeasurementOptions**](MeasurementOptions.md) |  | [optional] 
**DownloadablesOptions** | Pointer to [**DownloadablesOptions**](DownloadablesOptions.md) |  | [optional] 
**SubscriptionsOptions** | Pointer to [**SubscriptionsOptions**](SubscriptionsOptions.md) |  | [optional] 
**TaxCalculationOptions** | Pointer to [**TaxCalculationOptions**](TaxCalculationOptions.md) |  | [optional] 
**RecommendationOptions** | Pointer to [**RecommendationOptions**](RecommendationOptions.md) |  | [optional] 
**PriceCalculationOptions** | Pointer to [**PriceCalculationOptions**](PriceCalculationOptions.md) |  | [optional] 
**IdentityAndPrivacyOptions** | Pointer to [**IdentityAndPrivacyOptions**](IdentityAndPrivacyOptions.md) |  | [optional] 
**IncludedSellingLocations** | Pointer to **[]string** |  | [optional] 
**ExcludedSellingLocations** | Pointer to **[]string** |  | [optional] 
**IncludedShippingLocations** | Pointer to **[]string** |  | [optional] 
**ExcludedShippingLocations** | Pointer to **[]string** |  | [optional] 
**CurrencyPosition** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreOptions

`func NewStoreOptions() *StoreOptions`

NewStoreOptions instantiates a new StoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreOptionsWithDefaults

`func NewStoreOptionsWithDefaults() *StoreOptions`

NewStoreOptionsWithDefaults instantiates a new StoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *StoreOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *StoreOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *StoreOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *StoreOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *StoreOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StoreOptions) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StoreOptions) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogo

`func (o *StoreOptions) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *StoreOptions) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *StoreOptions) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *StoreOptions) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *StoreOptions) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *StoreOptions) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetFooterLogo

`func (o *StoreOptions) GetFooterLogo() string`

GetFooterLogo returns the FooterLogo field if non-nil, zero value otherwise.

### GetFooterLogoOk

`func (o *StoreOptions) GetFooterLogoOk() (*string, bool)`

GetFooterLogoOk returns a tuple with the FooterLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLogo

`func (o *StoreOptions) SetFooterLogo(v string)`

SetFooterLogo sets FooterLogo field to given value.

### HasFooterLogo

`func (o *StoreOptions) HasFooterLogo() bool`

HasFooterLogo returns a boolean if a field has been set.

### SetFooterLogoNil

`func (o *StoreOptions) SetFooterLogoNil(b bool)`

 SetFooterLogoNil sets the value for FooterLogo to be an explicit nil

### UnsetFooterLogo
`func (o *StoreOptions) UnsetFooterLogo()`

UnsetFooterLogo ensures that no value is present for FooterLogo, not even an explicit nil
### GetTagline

`func (o *StoreOptions) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *StoreOptions) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *StoreOptions) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *StoreOptions) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### SetTaglineNil

`func (o *StoreOptions) SetTaglineNil(b bool)`

 SetTaglineNil sets the value for Tagline to be an explicit nil

### UnsetTagline
`func (o *StoreOptions) UnsetTagline()`

UnsetTagline ensures that no value is present for Tagline, not even an explicit nil
### GetDescription

`func (o *StoreOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoreOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StoreOptions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StoreOptions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAddressLine1

`func (o *StoreOptions) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *StoreOptions) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *StoreOptions) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *StoreOptions) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *StoreOptions) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *StoreOptions) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *StoreOptions) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *StoreOptions) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *StoreOptions) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *StoreOptions) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *StoreOptions) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *StoreOptions) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCityId

`func (o *StoreOptions) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *StoreOptions) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *StoreOptions) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *StoreOptions) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *StoreOptions) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *StoreOptions) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *StoreOptions) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *StoreOptions) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *StoreOptions) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *StoreOptions) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *StoreOptions) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *StoreOptions) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCountryId

`func (o *StoreOptions) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *StoreOptions) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *StoreOptions) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *StoreOptions) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *StoreOptions) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *StoreOptions) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCurrencyId

`func (o *StoreOptions) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *StoreOptions) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *StoreOptions) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *StoreOptions) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *StoreOptions) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *StoreOptions) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPostalCode

`func (o *StoreOptions) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *StoreOptions) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *StoreOptions) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *StoreOptions) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *StoreOptions) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *StoreOptions) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetNumberOfDecimals

`func (o *StoreOptions) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *StoreOptions) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *StoreOptions) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *StoreOptions) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### GetDecimalSeparator

`func (o *StoreOptions) GetDecimalSeparator() string`

GetDecimalSeparator returns the DecimalSeparator field if non-nil, zero value otherwise.

### GetDecimalSeparatorOk

`func (o *StoreOptions) GetDecimalSeparatorOk() (*string, bool)`

GetDecimalSeparatorOk returns a tuple with the DecimalSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalSeparator

`func (o *StoreOptions) SetDecimalSeparator(v string)`

SetDecimalSeparator sets DecimalSeparator field to given value.

### HasDecimalSeparator

`func (o *StoreOptions) HasDecimalSeparator() bool`

HasDecimalSeparator returns a boolean if a field has been set.

### SetDecimalSeparatorNil

`func (o *StoreOptions) SetDecimalSeparatorNil(b bool)`

 SetDecimalSeparatorNil sets the value for DecimalSeparator to be an explicit nil

### UnsetDecimalSeparator
`func (o *StoreOptions) UnsetDecimalSeparator()`

UnsetDecimalSeparator ensures that no value is present for DecimalSeparator, not even an explicit nil
### GetSellToAllCountries

`func (o *StoreOptions) GetSellToAllCountries() bool`

GetSellToAllCountries returns the SellToAllCountries field if non-nil, zero value otherwise.

### GetSellToAllCountriesOk

`func (o *StoreOptions) GetSellToAllCountriesOk() (*bool, bool)`

GetSellToAllCountriesOk returns a tuple with the SellToAllCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellToAllCountries

`func (o *StoreOptions) SetSellToAllCountries(v bool)`

SetSellToAllCountries sets SellToAllCountries field to given value.

### HasSellToAllCountries

`func (o *StoreOptions) HasSellToAllCountries() bool`

HasSellToAllCountries returns a boolean if a field has been set.

### GetCartOptions

`func (o *StoreOptions) GetCartOptions() CartOptions`

GetCartOptions returns the CartOptions field if non-nil, zero value otherwise.

### GetCartOptionsOk

`func (o *StoreOptions) GetCartOptionsOk() (*CartOptions, bool)`

GetCartOptionsOk returns a tuple with the CartOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartOptions

`func (o *StoreOptions) SetCartOptions(v CartOptions)`

SetCartOptions sets CartOptions field to given value.

### HasCartOptions

`func (o *StoreOptions) HasCartOptions() bool`

HasCartOptions returns a boolean if a field has been set.

### GetEmailOptions

`func (o *StoreOptions) GetEmailOptions() EmailOptions`

GetEmailOptions returns the EmailOptions field if non-nil, zero value otherwise.

### GetEmailOptionsOk

`func (o *StoreOptions) GetEmailOptionsOk() (*EmailOptions, bool)`

GetEmailOptionsOk returns a tuple with the EmailOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOptions

`func (o *StoreOptions) SetEmailOptions(v EmailOptions)`

SetEmailOptions sets EmailOptions field to given value.

### HasEmailOptions

`func (o *StoreOptions) HasEmailOptions() bool`

HasEmailOptions returns a boolean if a field has been set.

### GetCouponsOptions

`func (o *StoreOptions) GetCouponsOptions() CouponsOptions`

GetCouponsOptions returns the CouponsOptions field if non-nil, zero value otherwise.

### GetCouponsOptionsOk

`func (o *StoreOptions) GetCouponsOptionsOk() (*CouponsOptions, bool)`

GetCouponsOptionsOk returns a tuple with the CouponsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsOptions

`func (o *StoreOptions) SetCouponsOptions(v CouponsOptions)`

SetCouponsOptions sets CouponsOptions field to given value.

### HasCouponsOptions

`func (o *StoreOptions) HasCouponsOptions() bool`

HasCouponsOptions returns a boolean if a field has been set.

### GetPaymentOptions

`func (o *StoreOptions) GetPaymentOptions() PaymentOptions`

GetPaymentOptions returns the PaymentOptions field if non-nil, zero value otherwise.

### GetPaymentOptionsOk

`func (o *StoreOptions) GetPaymentOptionsOk() (*PaymentOptions, bool)`

GetPaymentOptionsOk returns a tuple with the PaymentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptions

`func (o *StoreOptions) SetPaymentOptions(v PaymentOptions)`

SetPaymentOptions sets PaymentOptions field to given value.

### HasPaymentOptions

`func (o *StoreOptions) HasPaymentOptions() bool`

HasPaymentOptions returns a boolean if a field has been set.

### GetProductOptions

`func (o *StoreOptions) GetProductOptions() map[string]interface{}`

GetProductOptions returns the ProductOptions field if non-nil, zero value otherwise.

### GetProductOptionsOk

`func (o *StoreOptions) GetProductOptionsOk() (*map[string]interface{}, bool)`

GetProductOptionsOk returns a tuple with the ProductOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptions

`func (o *StoreOptions) SetProductOptions(v map[string]interface{})`

SetProductOptions sets ProductOptions field to given value.

### HasProductOptions

`func (o *StoreOptions) HasProductOptions() bool`

HasProductOptions returns a boolean if a field has been set.

### GetReviewsOptions

`func (o *StoreOptions) GetReviewsOptions() ReviewsOptions`

GetReviewsOptions returns the ReviewsOptions field if non-nil, zero value otherwise.

### GetReviewsOptionsOk

`func (o *StoreOptions) GetReviewsOptionsOk() (*ReviewsOptions, bool)`

GetReviewsOptionsOk returns a tuple with the ReviewsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewsOptions

`func (o *StoreOptions) SetReviewsOptions(v ReviewsOptions)`

SetReviewsOptions sets ReviewsOptions field to given value.

### HasReviewsOptions

`func (o *StoreOptions) HasReviewsOptions() bool`

HasReviewsOptions returns a boolean if a field has been set.

### GetAdvancedOptions

`func (o *StoreOptions) GetAdvancedOptions() AdvancedOptions`

GetAdvancedOptions returns the AdvancedOptions field if non-nil, zero value otherwise.

### GetAdvancedOptionsOk

`func (o *StoreOptions) GetAdvancedOptionsOk() (*AdvancedOptions, bool)`

GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOptions

`func (o *StoreOptions) SetAdvancedOptions(v AdvancedOptions)`

SetAdvancedOptions sets AdvancedOptions field to given value.

### HasAdvancedOptions

`func (o *StoreOptions) HasAdvancedOptions() bool`

HasAdvancedOptions returns a boolean if a field has been set.

### GetServicesOptions

`func (o *StoreOptions) GetServicesOptions() ServicesOptions`

GetServicesOptions returns the ServicesOptions field if non-nil, zero value otherwise.

### GetServicesOptionsOk

`func (o *StoreOptions) GetServicesOptionsOk() (*ServicesOptions, bool)`

GetServicesOptionsOk returns a tuple with the ServicesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesOptions

`func (o *StoreOptions) SetServicesOptions(v ServicesOptions)`

SetServicesOptions sets ServicesOptions field to given value.

### HasServicesOptions

`func (o *StoreOptions) HasServicesOptions() bool`

HasServicesOptions returns a boolean if a field has been set.

### GetInventoryOptions

`func (o *StoreOptions) GetInventoryOptions() InventoryOptions`

GetInventoryOptions returns the InventoryOptions field if non-nil, zero value otherwise.

### GetInventoryOptionsOk

`func (o *StoreOptions) GetInventoryOptionsOk() (*InventoryOptions, bool)`

GetInventoryOptionsOk returns a tuple with the InventoryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryOptions

`func (o *StoreOptions) SetInventoryOptions(v InventoryOptions)`

SetInventoryOptions sets InventoryOptions field to given value.

### HasInventoryOptions

`func (o *StoreOptions) HasInventoryOptions() bool`

HasInventoryOptions returns a boolean if a field has been set.

### GetIntegrationOptions

`func (o *StoreOptions) GetIntegrationOptions() IntegrationOptions`

GetIntegrationOptions returns the IntegrationOptions field if non-nil, zero value otherwise.

### GetIntegrationOptionsOk

`func (o *StoreOptions) GetIntegrationOptionsOk() (*IntegrationOptions, bool)`

GetIntegrationOptionsOk returns a tuple with the IntegrationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationOptions

`func (o *StoreOptions) SetIntegrationOptions(v IntegrationOptions)`

SetIntegrationOptions sets IntegrationOptions field to given value.

### HasIntegrationOptions

`func (o *StoreOptions) HasIntegrationOptions() bool`

HasIntegrationOptions returns a boolean if a field has been set.

### GetMeasurementOptions

`func (o *StoreOptions) GetMeasurementOptions() MeasurementOptions`

GetMeasurementOptions returns the MeasurementOptions field if non-nil, zero value otherwise.

### GetMeasurementOptionsOk

`func (o *StoreOptions) GetMeasurementOptionsOk() (*MeasurementOptions, bool)`

GetMeasurementOptionsOk returns a tuple with the MeasurementOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementOptions

`func (o *StoreOptions) SetMeasurementOptions(v MeasurementOptions)`

SetMeasurementOptions sets MeasurementOptions field to given value.

### HasMeasurementOptions

`func (o *StoreOptions) HasMeasurementOptions() bool`

HasMeasurementOptions returns a boolean if a field has been set.

### GetDownloadablesOptions

`func (o *StoreOptions) GetDownloadablesOptions() DownloadablesOptions`

GetDownloadablesOptions returns the DownloadablesOptions field if non-nil, zero value otherwise.

### GetDownloadablesOptionsOk

`func (o *StoreOptions) GetDownloadablesOptionsOk() (*DownloadablesOptions, bool)`

GetDownloadablesOptionsOk returns a tuple with the DownloadablesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadablesOptions

`func (o *StoreOptions) SetDownloadablesOptions(v DownloadablesOptions)`

SetDownloadablesOptions sets DownloadablesOptions field to given value.

### HasDownloadablesOptions

`func (o *StoreOptions) HasDownloadablesOptions() bool`

HasDownloadablesOptions returns a boolean if a field has been set.

### GetSubscriptionsOptions

`func (o *StoreOptions) GetSubscriptionsOptions() SubscriptionsOptions`

GetSubscriptionsOptions returns the SubscriptionsOptions field if non-nil, zero value otherwise.

### GetSubscriptionsOptionsOk

`func (o *StoreOptions) GetSubscriptionsOptionsOk() (*SubscriptionsOptions, bool)`

GetSubscriptionsOptionsOk returns a tuple with the SubscriptionsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsOptions

`func (o *StoreOptions) SetSubscriptionsOptions(v SubscriptionsOptions)`

SetSubscriptionsOptions sets SubscriptionsOptions field to given value.

### HasSubscriptionsOptions

`func (o *StoreOptions) HasSubscriptionsOptions() bool`

HasSubscriptionsOptions returns a boolean if a field has been set.

### GetTaxCalculationOptions

`func (o *StoreOptions) GetTaxCalculationOptions() TaxCalculationOptions`

GetTaxCalculationOptions returns the TaxCalculationOptions field if non-nil, zero value otherwise.

### GetTaxCalculationOptionsOk

`func (o *StoreOptions) GetTaxCalculationOptionsOk() (*TaxCalculationOptions, bool)`

GetTaxCalculationOptionsOk returns a tuple with the TaxCalculationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationOptions

`func (o *StoreOptions) SetTaxCalculationOptions(v TaxCalculationOptions)`

SetTaxCalculationOptions sets TaxCalculationOptions field to given value.

### HasTaxCalculationOptions

`func (o *StoreOptions) HasTaxCalculationOptions() bool`

HasTaxCalculationOptions returns a boolean if a field has been set.

### GetRecommendationOptions

`func (o *StoreOptions) GetRecommendationOptions() RecommendationOptions`

GetRecommendationOptions returns the RecommendationOptions field if non-nil, zero value otherwise.

### GetRecommendationOptionsOk

`func (o *StoreOptions) GetRecommendationOptionsOk() (*RecommendationOptions, bool)`

GetRecommendationOptionsOk returns a tuple with the RecommendationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationOptions

`func (o *StoreOptions) SetRecommendationOptions(v RecommendationOptions)`

SetRecommendationOptions sets RecommendationOptions field to given value.

### HasRecommendationOptions

`func (o *StoreOptions) HasRecommendationOptions() bool`

HasRecommendationOptions returns a boolean if a field has been set.

### GetPriceCalculationOptions

`func (o *StoreOptions) GetPriceCalculationOptions() PriceCalculationOptions`

GetPriceCalculationOptions returns the PriceCalculationOptions field if non-nil, zero value otherwise.

### GetPriceCalculationOptionsOk

`func (o *StoreOptions) GetPriceCalculationOptionsOk() (*PriceCalculationOptions, bool)`

GetPriceCalculationOptionsOk returns a tuple with the PriceCalculationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCalculationOptions

`func (o *StoreOptions) SetPriceCalculationOptions(v PriceCalculationOptions)`

SetPriceCalculationOptions sets PriceCalculationOptions field to given value.

### HasPriceCalculationOptions

`func (o *StoreOptions) HasPriceCalculationOptions() bool`

HasPriceCalculationOptions returns a boolean if a field has been set.

### GetIdentityAndPrivacyOptions

`func (o *StoreOptions) GetIdentityAndPrivacyOptions() IdentityAndPrivacyOptions`

GetIdentityAndPrivacyOptions returns the IdentityAndPrivacyOptions field if non-nil, zero value otherwise.

### GetIdentityAndPrivacyOptionsOk

`func (o *StoreOptions) GetIdentityAndPrivacyOptionsOk() (*IdentityAndPrivacyOptions, bool)`

GetIdentityAndPrivacyOptionsOk returns a tuple with the IdentityAndPrivacyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAndPrivacyOptions

`func (o *StoreOptions) SetIdentityAndPrivacyOptions(v IdentityAndPrivacyOptions)`

SetIdentityAndPrivacyOptions sets IdentityAndPrivacyOptions field to given value.

### HasIdentityAndPrivacyOptions

`func (o *StoreOptions) HasIdentityAndPrivacyOptions() bool`

HasIdentityAndPrivacyOptions returns a boolean if a field has been set.

### GetIncludedSellingLocations

`func (o *StoreOptions) GetIncludedSellingLocations() []string`

GetIncludedSellingLocations returns the IncludedSellingLocations field if non-nil, zero value otherwise.

### GetIncludedSellingLocationsOk

`func (o *StoreOptions) GetIncludedSellingLocationsOk() (*[]string, bool)`

GetIncludedSellingLocationsOk returns a tuple with the IncludedSellingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSellingLocations

`func (o *StoreOptions) SetIncludedSellingLocations(v []string)`

SetIncludedSellingLocations sets IncludedSellingLocations field to given value.

### HasIncludedSellingLocations

`func (o *StoreOptions) HasIncludedSellingLocations() bool`

HasIncludedSellingLocations returns a boolean if a field has been set.

### SetIncludedSellingLocationsNil

`func (o *StoreOptions) SetIncludedSellingLocationsNil(b bool)`

 SetIncludedSellingLocationsNil sets the value for IncludedSellingLocations to be an explicit nil

### UnsetIncludedSellingLocations
`func (o *StoreOptions) UnsetIncludedSellingLocations()`

UnsetIncludedSellingLocations ensures that no value is present for IncludedSellingLocations, not even an explicit nil
### GetExcludedSellingLocations

`func (o *StoreOptions) GetExcludedSellingLocations() []string`

GetExcludedSellingLocations returns the ExcludedSellingLocations field if non-nil, zero value otherwise.

### GetExcludedSellingLocationsOk

`func (o *StoreOptions) GetExcludedSellingLocationsOk() (*[]string, bool)`

GetExcludedSellingLocationsOk returns a tuple with the ExcludedSellingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSellingLocations

`func (o *StoreOptions) SetExcludedSellingLocations(v []string)`

SetExcludedSellingLocations sets ExcludedSellingLocations field to given value.

### HasExcludedSellingLocations

`func (o *StoreOptions) HasExcludedSellingLocations() bool`

HasExcludedSellingLocations returns a boolean if a field has been set.

### SetExcludedSellingLocationsNil

`func (o *StoreOptions) SetExcludedSellingLocationsNil(b bool)`

 SetExcludedSellingLocationsNil sets the value for ExcludedSellingLocations to be an explicit nil

### UnsetExcludedSellingLocations
`func (o *StoreOptions) UnsetExcludedSellingLocations()`

UnsetExcludedSellingLocations ensures that no value is present for ExcludedSellingLocations, not even an explicit nil
### GetIncludedShippingLocations

`func (o *StoreOptions) GetIncludedShippingLocations() []string`

GetIncludedShippingLocations returns the IncludedShippingLocations field if non-nil, zero value otherwise.

### GetIncludedShippingLocationsOk

`func (o *StoreOptions) GetIncludedShippingLocationsOk() (*[]string, bool)`

GetIncludedShippingLocationsOk returns a tuple with the IncludedShippingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedShippingLocations

`func (o *StoreOptions) SetIncludedShippingLocations(v []string)`

SetIncludedShippingLocations sets IncludedShippingLocations field to given value.

### HasIncludedShippingLocations

`func (o *StoreOptions) HasIncludedShippingLocations() bool`

HasIncludedShippingLocations returns a boolean if a field has been set.

### SetIncludedShippingLocationsNil

`func (o *StoreOptions) SetIncludedShippingLocationsNil(b bool)`

 SetIncludedShippingLocationsNil sets the value for IncludedShippingLocations to be an explicit nil

### UnsetIncludedShippingLocations
`func (o *StoreOptions) UnsetIncludedShippingLocations()`

UnsetIncludedShippingLocations ensures that no value is present for IncludedShippingLocations, not even an explicit nil
### GetExcludedShippingLocations

`func (o *StoreOptions) GetExcludedShippingLocations() []string`

GetExcludedShippingLocations returns the ExcludedShippingLocations field if non-nil, zero value otherwise.

### GetExcludedShippingLocationsOk

`func (o *StoreOptions) GetExcludedShippingLocationsOk() (*[]string, bool)`

GetExcludedShippingLocationsOk returns a tuple with the ExcludedShippingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedShippingLocations

`func (o *StoreOptions) SetExcludedShippingLocations(v []string)`

SetExcludedShippingLocations sets ExcludedShippingLocations field to given value.

### HasExcludedShippingLocations

`func (o *StoreOptions) HasExcludedShippingLocations() bool`

HasExcludedShippingLocations returns a boolean if a field has been set.

### SetExcludedShippingLocationsNil

`func (o *StoreOptions) SetExcludedShippingLocationsNil(b bool)`

 SetExcludedShippingLocationsNil sets the value for ExcludedShippingLocations to be an explicit nil

### UnsetExcludedShippingLocations
`func (o *StoreOptions) UnsetExcludedShippingLocations()`

UnsetExcludedShippingLocations ensures that no value is present for ExcludedShippingLocations, not even an explicit nil
### GetCurrencyPosition

`func (o *StoreOptions) GetCurrencyPosition() string`

GetCurrencyPosition returns the CurrencyPosition field if non-nil, zero value otherwise.

### GetCurrencyPositionOk

`func (o *StoreOptions) GetCurrencyPositionOk() (*string, bool)`

GetCurrencyPositionOk returns a tuple with the CurrencyPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPosition

`func (o *StoreOptions) SetCurrencyPosition(v string)`

SetCurrencyPosition sets CurrencyPosition field to given value.

### HasCurrencyPosition

`func (o *StoreOptions) HasCurrencyPosition() bool`

HasCurrencyPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


