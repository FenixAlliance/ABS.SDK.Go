# QuoteCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**PaymentTermId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**AddressLine1** | Pointer to **NullableString** |  | [optional] 
**AddressLine2** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDetail** | Pointer to **float64** |  | [optional] 
**TotalDetailCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalProfit** | Pointer to **float64** |  | [optional] 
**TotalProfitCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingCost** | Pointer to **float64** |  | [optional] 
**TotalShippingCostCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingTax** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWithheldTax** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxBase** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**TotalTaxesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**TotalCurrencyId** | Pointer to **NullableString** |  | [optional] 
**CostCalculationMethod** | Pointer to **string** |  | [optional] 
**TaxCalculationMethod** | Pointer to **string** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**DealUnitId** | Pointer to **NullableString** |  | [optional] 
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
**EffectiveTo** | Pointer to **NullableTime** |  | [optional] 
**EffectiveFrom** | Pointer to **NullableTime** |  | [optional] 
**QuoteStatus** | Pointer to **string** |  | [optional] 
**QuoteLines** | Pointer to [**[]QuoteLineCreateDto**](QuoteLineCreateDto.md) |  | [optional] 

## Methods

### NewQuoteCreateDto

`func NewQuoteCreateDto() *QuoteCreateDto`

NewQuoteCreateDto instantiates a new QuoteCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateDtoWithDefaults

`func NewQuoteCreateDtoWithDefaults() *QuoteCreateDto`

NewQuoteCreateDtoWithDefaults instantiates a new QuoteCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuoteCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *QuoteCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QuoteCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QuoteCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QuoteCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetClosed

`func (o *QuoteCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *QuoteCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *QuoteCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *QuoteCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetTitle

`func (o *QuoteCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuoteCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuoteCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuoteCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QuoteCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QuoteCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriceListId

`func (o *QuoteCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *QuoteCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *QuoteCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *QuoteCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *QuoteCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *QuoteCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDescription

`func (o *QuoteCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuoteCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuoteCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIndividualId

`func (o *QuoteCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *QuoteCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *QuoteCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *QuoteCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *QuoteCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *QuoteCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetPaymentTermId

`func (o *QuoteCreateDto) GetPaymentTermId() string`

GetPaymentTermId returns the PaymentTermId field if non-nil, zero value otherwise.

### GetPaymentTermIdOk

`func (o *QuoteCreateDto) GetPaymentTermIdOk() (*string, bool)`

GetPaymentTermIdOk returns a tuple with the PaymentTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermId

`func (o *QuoteCreateDto) SetPaymentTermId(v string)`

SetPaymentTermId sets PaymentTermId field to given value.

### HasPaymentTermId

`func (o *QuoteCreateDto) HasPaymentTermId() bool`

HasPaymentTermId returns a boolean if a field has been set.

### SetPaymentTermIdNil

`func (o *QuoteCreateDto) SetPaymentTermIdNil(b bool)`

 SetPaymentTermIdNil sets the value for PaymentTermId to be an explicit nil

### UnsetPaymentTermId
`func (o *QuoteCreateDto) UnsetPaymentTermId()`

UnsetPaymentTermId ensures that no value is present for PaymentTermId, not even an explicit nil
### GetOrganizationId

`func (o *QuoteCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *QuoteCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *QuoteCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *QuoteCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *QuoteCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *QuoteCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetFirstName

`func (o *QuoteCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuoteCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuoteCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuoteCreateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *QuoteCreateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *QuoteCreateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *QuoteCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuoteCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuoteCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuoteCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *QuoteCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *QuoteCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *QuoteCreateDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuoteCreateDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuoteCreateDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuoteCreateDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *QuoteCreateDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *QuoteCreateDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *QuoteCreateDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *QuoteCreateDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *QuoteCreateDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *QuoteCreateDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *QuoteCreateDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *QuoteCreateDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *QuoteCreateDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *QuoteCreateDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *QuoteCreateDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *QuoteCreateDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *QuoteCreateDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *QuoteCreateDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *QuoteCreateDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *QuoteCreateDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *QuoteCreateDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *QuoteCreateDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *QuoteCreateDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *QuoteCreateDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *QuoteCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuoteCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuoteCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuoteCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *QuoteCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *QuoteCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *QuoteCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *QuoteCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *QuoteCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *QuoteCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *QuoteCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *QuoteCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *QuoteCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *QuoteCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *QuoteCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *QuoteCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *QuoteCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *QuoteCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *QuoteCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *QuoteCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *QuoteCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *QuoteCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *QuoteCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *QuoteCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetForexRate

`func (o *QuoteCreateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *QuoteCreateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *QuoteCreateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *QuoteCreateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuoteCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuoteCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuoteCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuoteCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *QuoteCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *QuoteCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *QuoteCreateDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *QuoteCreateDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *QuoteCreateDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *QuoteCreateDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *QuoteCreateDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *QuoteCreateDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *QuoteCreateDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *QuoteCreateDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *QuoteCreateDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *QuoteCreateDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *QuoteCreateDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *QuoteCreateDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *QuoteCreateDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *QuoteCreateDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *QuoteCreateDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *QuoteCreateDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *QuoteCreateDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *QuoteCreateDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *QuoteCreateDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *QuoteCreateDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *QuoteCreateDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *QuoteCreateDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *QuoteCreateDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *QuoteCreateDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *QuoteCreateDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *QuoteCreateDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *QuoteCreateDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *QuoteCreateDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *QuoteCreateDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *QuoteCreateDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *QuoteCreateDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *QuoteCreateDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *QuoteCreateDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *QuoteCreateDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *QuoteCreateDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *QuoteCreateDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *QuoteCreateDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *QuoteCreateDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *QuoteCreateDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *QuoteCreateDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *QuoteCreateDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *QuoteCreateDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *QuoteCreateDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *QuoteCreateDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *QuoteCreateDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *QuoteCreateDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *QuoteCreateDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *QuoteCreateDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *QuoteCreateDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *QuoteCreateDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *QuoteCreateDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *QuoteCreateDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *QuoteCreateDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *QuoteCreateDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *QuoteCreateDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *QuoteCreateDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *QuoteCreateDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *QuoteCreateDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *QuoteCreateDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *QuoteCreateDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *QuoteCreateDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *QuoteCreateDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *QuoteCreateDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *QuoteCreateDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *QuoteCreateDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *QuoteCreateDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *QuoteCreateDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *QuoteCreateDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *QuoteCreateDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *QuoteCreateDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *QuoteCreateDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *QuoteCreateDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *QuoteCreateDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *QuoteCreateDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *QuoteCreateDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *QuoteCreateDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *QuoteCreateDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *QuoteCreateDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *QuoteCreateDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *QuoteCreateDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *QuoteCreateDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *QuoteCreateDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *QuoteCreateDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *QuoteCreateDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *QuoteCreateDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *QuoteCreateDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *QuoteCreateDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *QuoteCreateDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotal

`func (o *QuoteCreateDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuoteCreateDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuoteCreateDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QuoteCreateDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *QuoteCreateDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *QuoteCreateDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *QuoteCreateDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *QuoteCreateDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *QuoteCreateDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *QuoteCreateDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetCostCalculationMethod

`func (o *QuoteCreateDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *QuoteCreateDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *QuoteCreateDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *QuoteCreateDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetTaxCalculationMethod

`func (o *QuoteCreateDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *QuoteCreateDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *QuoteCreateDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *QuoteCreateDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCartId

`func (o *QuoteCreateDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *QuoteCreateDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *QuoteCreateDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *QuoteCreateDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *QuoteCreateDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *QuoteCreateDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetDealUnitId

`func (o *QuoteCreateDto) GetDealUnitId() string`

GetDealUnitId returns the DealUnitId field if non-nil, zero value otherwise.

### GetDealUnitIdOk

`func (o *QuoteCreateDto) GetDealUnitIdOk() (*string, bool)`

GetDealUnitIdOk returns a tuple with the DealUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitId

`func (o *QuoteCreateDto) SetDealUnitId(v string)`

SetDealUnitId sets DealUnitId field to given value.

### HasDealUnitId

`func (o *QuoteCreateDto) HasDealUnitId() bool`

HasDealUnitId returns a boolean if a field has been set.

### SetDealUnitIdNil

`func (o *QuoteCreateDto) SetDealUnitIdNil(b bool)`

 SetDealUnitIdNil sets the value for DealUnitId to be an explicit nil

### UnsetDealUnitId
`func (o *QuoteCreateDto) UnsetDealUnitId()`

UnsetDealUnitId ensures that no value is present for DealUnitId, not even an explicit nil
### GetReceiverTenantId

`func (o *QuoteCreateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *QuoteCreateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *QuoteCreateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *QuoteCreateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *QuoteCreateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *QuoteCreateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetEffectiveTo

`func (o *QuoteCreateDto) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *QuoteCreateDto) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *QuoteCreateDto) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *QuoteCreateDto) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *QuoteCreateDto) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *QuoteCreateDto) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetEffectiveFrom

`func (o *QuoteCreateDto) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *QuoteCreateDto) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *QuoteCreateDto) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *QuoteCreateDto) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### SetEffectiveFromNil

`func (o *QuoteCreateDto) SetEffectiveFromNil(b bool)`

 SetEffectiveFromNil sets the value for EffectiveFrom to be an explicit nil

### UnsetEffectiveFrom
`func (o *QuoteCreateDto) UnsetEffectiveFrom()`

UnsetEffectiveFrom ensures that no value is present for EffectiveFrom, not even an explicit nil
### GetQuoteStatus

`func (o *QuoteCreateDto) GetQuoteStatus() string`

GetQuoteStatus returns the QuoteStatus field if non-nil, zero value otherwise.

### GetQuoteStatusOk

`func (o *QuoteCreateDto) GetQuoteStatusOk() (*string, bool)`

GetQuoteStatusOk returns a tuple with the QuoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStatus

`func (o *QuoteCreateDto) SetQuoteStatus(v string)`

SetQuoteStatus sets QuoteStatus field to given value.

### HasQuoteStatus

`func (o *QuoteCreateDto) HasQuoteStatus() bool`

HasQuoteStatus returns a boolean if a field has been set.

### GetQuoteLines

`func (o *QuoteCreateDto) GetQuoteLines() []QuoteLineCreateDto`

GetQuoteLines returns the QuoteLines field if non-nil, zero value otherwise.

### GetQuoteLinesOk

`func (o *QuoteCreateDto) GetQuoteLinesOk() (*[]QuoteLineCreateDto, bool)`

GetQuoteLinesOk returns a tuple with the QuoteLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLines

`func (o *QuoteCreateDto) SetQuoteLines(v []QuoteLineCreateDto)`

SetQuoteLines sets QuoteLines field to given value.

### HasQuoteLines

`func (o *QuoteCreateDto) HasQuoteLines() bool`

HasQuoteLines returns a boolean if a field has been set.

### SetQuoteLinesNil

`func (o *QuoteCreateDto) SetQuoteLinesNil(b bool)`

 SetQuoteLinesNil sets the value for QuoteLines to be an explicit nil

### UnsetQuoteLines
`func (o *QuoteCreateDto) UnsetQuoteLines()`

UnsetQuoteLines ensures that no value is present for QuoteLines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


