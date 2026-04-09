# OrderCreateDto

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
**QuoteId** | Pointer to **NullableString** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**ParentOrderId** | Pointer to **NullableString** |  | [optional] 
**ShippingMethodId** | Pointer to **NullableString** |  | [optional] 
**BillingLocationId** | Pointer to **NullableString** |  | [optional] 
**CustomerNotes** | Pointer to **NullableString** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**QuoteStatus** | Pointer to **string** |  | [optional] 
**FreightTerms** | Pointer to **string** |  | [optional] 
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
**QualifiedIdentifier** | Pointer to **NullableString** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxInUsd** | Pointer to **float64** |  | [optional] 
**TotalAmountInUsd** | Pointer to **float64** |  | [optional] 
**EffectiveTo** | Pointer to **NullableTime** |  | [optional] 
**EffectiveFrom** | Pointer to **NullableTime** |  | [optional] 
**OrderLines** | Pointer to [**[]OrderLineCreateDto**](OrderLineCreateDto.md) |  | [optional] 

## Methods

### NewOrderCreateDto

`func NewOrderCreateDto() *OrderCreateDto`

NewOrderCreateDto instantiates a new OrderCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateDtoWithDefaults

`func NewOrderCreateDtoWithDefaults() *OrderCreateDto`

NewOrderCreateDtoWithDefaults instantiates a new OrderCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *OrderCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OrderCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetClosed

`func (o *OrderCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *OrderCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *OrderCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *OrderCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetTitle

`func (o *OrderCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrderCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrderCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OrderCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OrderCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OrderCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriceListId

`func (o *OrderCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *OrderCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *OrderCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *OrderCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *OrderCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *OrderCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDescription

`func (o *OrderCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrderCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrderCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIndividualId

`func (o *OrderCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *OrderCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *OrderCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *OrderCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *OrderCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *OrderCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetPaymentTermId

`func (o *OrderCreateDto) GetPaymentTermId() string`

GetPaymentTermId returns the PaymentTermId field if non-nil, zero value otherwise.

### GetPaymentTermIdOk

`func (o *OrderCreateDto) GetPaymentTermIdOk() (*string, bool)`

GetPaymentTermIdOk returns a tuple with the PaymentTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermId

`func (o *OrderCreateDto) SetPaymentTermId(v string)`

SetPaymentTermId sets PaymentTermId field to given value.

### HasPaymentTermId

`func (o *OrderCreateDto) HasPaymentTermId() bool`

HasPaymentTermId returns a boolean if a field has been set.

### SetPaymentTermIdNil

`func (o *OrderCreateDto) SetPaymentTermIdNil(b bool)`

 SetPaymentTermIdNil sets the value for PaymentTermId to be an explicit nil

### UnsetPaymentTermId
`func (o *OrderCreateDto) UnsetPaymentTermId()`

UnsetPaymentTermId ensures that no value is present for PaymentTermId, not even an explicit nil
### GetOrganizationId

`func (o *OrderCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrderCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrderCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrderCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *OrderCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *OrderCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetFirstName

`func (o *OrderCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrderCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrderCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrderCreateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *OrderCreateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *OrderCreateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *OrderCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrderCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrderCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrderCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *OrderCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *OrderCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *OrderCreateDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrderCreateDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrderCreateDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *OrderCreateDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrderCreateDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrderCreateDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrderCreateDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *OrderCreateDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *OrderCreateDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *OrderCreateDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrderCreateDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrderCreateDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrderCreateDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrderCreateDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrderCreateDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *OrderCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *OrderCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *OrderCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *OrderCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *OrderCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *OrderCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *OrderCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *OrderCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *OrderCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *OrderCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *OrderCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *OrderCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *OrderCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *OrderCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *OrderCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *OrderCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *OrderCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *OrderCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *OrderCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *OrderCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *OrderCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetForexRate

`func (o *OrderCreateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *OrderCreateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *OrderCreateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *OrderCreateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *OrderCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *OrderCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *OrderCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *OrderCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *OrderCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *OrderCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *OrderCreateDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *OrderCreateDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *OrderCreateDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *OrderCreateDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *OrderCreateDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *OrderCreateDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *OrderCreateDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *OrderCreateDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *OrderCreateDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *OrderCreateDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *OrderCreateDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *OrderCreateDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *OrderCreateDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *OrderCreateDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *OrderCreateDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *OrderCreateDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *OrderCreateDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *OrderCreateDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *OrderCreateDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *OrderCreateDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *OrderCreateDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *OrderCreateDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *OrderCreateDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *OrderCreateDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *OrderCreateDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *OrderCreateDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *OrderCreateDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *OrderCreateDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *OrderCreateDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *OrderCreateDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *OrderCreateDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *OrderCreateDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *OrderCreateDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *OrderCreateDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *OrderCreateDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *OrderCreateDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *OrderCreateDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *OrderCreateDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *OrderCreateDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *OrderCreateDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *OrderCreateDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *OrderCreateDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *OrderCreateDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *OrderCreateDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *OrderCreateDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *OrderCreateDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *OrderCreateDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *OrderCreateDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *OrderCreateDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *OrderCreateDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *OrderCreateDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *OrderCreateDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *OrderCreateDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *OrderCreateDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *OrderCreateDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *OrderCreateDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *OrderCreateDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *OrderCreateDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *OrderCreateDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *OrderCreateDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *OrderCreateDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *OrderCreateDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *OrderCreateDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *OrderCreateDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *OrderCreateDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *OrderCreateDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *OrderCreateDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *OrderCreateDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *OrderCreateDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *OrderCreateDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *OrderCreateDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *OrderCreateDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *OrderCreateDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *OrderCreateDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *OrderCreateDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *OrderCreateDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *OrderCreateDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *OrderCreateDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *OrderCreateDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *OrderCreateDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *OrderCreateDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *OrderCreateDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *OrderCreateDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *OrderCreateDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *OrderCreateDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *OrderCreateDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *OrderCreateDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *OrderCreateDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *OrderCreateDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *OrderCreateDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *OrderCreateDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *OrderCreateDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *OrderCreateDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *OrderCreateDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *OrderCreateDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *OrderCreateDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *OrderCreateDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *OrderCreateDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *OrderCreateDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *OrderCreateDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *OrderCreateDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *OrderCreateDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *OrderCreateDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *OrderCreateDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *OrderCreateDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *OrderCreateDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *OrderCreateDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *OrderCreateDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *OrderCreateDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *OrderCreateDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotal

`func (o *OrderCreateDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderCreateDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderCreateDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderCreateDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *OrderCreateDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *OrderCreateDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *OrderCreateDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *OrderCreateDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *OrderCreateDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *OrderCreateDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetCostCalculationMethod

`func (o *OrderCreateDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *OrderCreateDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *OrderCreateDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *OrderCreateDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetTaxCalculationMethod

`func (o *OrderCreateDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *OrderCreateDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *OrderCreateDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *OrderCreateDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCartId

`func (o *OrderCreateDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *OrderCreateDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *OrderCreateDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *OrderCreateDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *OrderCreateDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *OrderCreateDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetQuoteId

`func (o *OrderCreateDto) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *OrderCreateDto) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *OrderCreateDto) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *OrderCreateDto) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### SetQuoteIdNil

`func (o *OrderCreateDto) SetQuoteIdNil(b bool)`

 SetQuoteIdNil sets the value for QuoteId to be an explicit nil

### UnsetQuoteId
`func (o *OrderCreateDto) UnsetQuoteId()`

UnsetQuoteId ensures that no value is present for QuoteId, not even an explicit nil
### GetWalletId

`func (o *OrderCreateDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *OrderCreateDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *OrderCreateDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *OrderCreateDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *OrderCreateDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *OrderCreateDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetParentOrderId

`func (o *OrderCreateDto) GetParentOrderId() string`

GetParentOrderId returns the ParentOrderId field if non-nil, zero value otherwise.

### GetParentOrderIdOk

`func (o *OrderCreateDto) GetParentOrderIdOk() (*string, bool)`

GetParentOrderIdOk returns a tuple with the ParentOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrderId

`func (o *OrderCreateDto) SetParentOrderId(v string)`

SetParentOrderId sets ParentOrderId field to given value.

### HasParentOrderId

`func (o *OrderCreateDto) HasParentOrderId() bool`

HasParentOrderId returns a boolean if a field has been set.

### SetParentOrderIdNil

`func (o *OrderCreateDto) SetParentOrderIdNil(b bool)`

 SetParentOrderIdNil sets the value for ParentOrderId to be an explicit nil

### UnsetParentOrderId
`func (o *OrderCreateDto) UnsetParentOrderId()`

UnsetParentOrderId ensures that no value is present for ParentOrderId, not even an explicit nil
### GetShippingMethodId

`func (o *OrderCreateDto) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *OrderCreateDto) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *OrderCreateDto) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *OrderCreateDto) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### SetShippingMethodIdNil

`func (o *OrderCreateDto) SetShippingMethodIdNil(b bool)`

 SetShippingMethodIdNil sets the value for ShippingMethodId to be an explicit nil

### UnsetShippingMethodId
`func (o *OrderCreateDto) UnsetShippingMethodId()`

UnsetShippingMethodId ensures that no value is present for ShippingMethodId, not even an explicit nil
### GetBillingLocationId

`func (o *OrderCreateDto) GetBillingLocationId() string`

GetBillingLocationId returns the BillingLocationId field if non-nil, zero value otherwise.

### GetBillingLocationIdOk

`func (o *OrderCreateDto) GetBillingLocationIdOk() (*string, bool)`

GetBillingLocationIdOk returns a tuple with the BillingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLocationId

`func (o *OrderCreateDto) SetBillingLocationId(v string)`

SetBillingLocationId sets BillingLocationId field to given value.

### HasBillingLocationId

`func (o *OrderCreateDto) HasBillingLocationId() bool`

HasBillingLocationId returns a boolean if a field has been set.

### SetBillingLocationIdNil

`func (o *OrderCreateDto) SetBillingLocationIdNil(b bool)`

 SetBillingLocationIdNil sets the value for BillingLocationId to be an explicit nil

### UnsetBillingLocationId
`func (o *OrderCreateDto) UnsetBillingLocationId()`

UnsetBillingLocationId ensures that no value is present for BillingLocationId, not even an explicit nil
### GetCustomerNotes

`func (o *OrderCreateDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *OrderCreateDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *OrderCreateDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *OrderCreateDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *OrderCreateDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *OrderCreateDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetOrderStatus

`func (o *OrderCreateDto) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderCreateDto) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderCreateDto) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderCreateDto) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetQuoteStatus

`func (o *OrderCreateDto) GetQuoteStatus() string`

GetQuoteStatus returns the QuoteStatus field if non-nil, zero value otherwise.

### GetQuoteStatusOk

`func (o *OrderCreateDto) GetQuoteStatusOk() (*string, bool)`

GetQuoteStatusOk returns a tuple with the QuoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStatus

`func (o *OrderCreateDto) SetQuoteStatus(v string)`

SetQuoteStatus sets QuoteStatus field to given value.

### HasQuoteStatus

`func (o *OrderCreateDto) HasQuoteStatus() bool`

HasQuoteStatus returns a boolean if a field has been set.

### GetFreightTerms

`func (o *OrderCreateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *OrderCreateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *OrderCreateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *OrderCreateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### GetReceiverTenantId

`func (o *OrderCreateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *OrderCreateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *OrderCreateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *OrderCreateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *OrderCreateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *OrderCreateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetShippingLocationId

`func (o *OrderCreateDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *OrderCreateDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *OrderCreateDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *OrderCreateDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *OrderCreateDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *OrderCreateDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetQualifiedIdentifier

`func (o *OrderCreateDto) GetQualifiedIdentifier() string`

GetQualifiedIdentifier returns the QualifiedIdentifier field if non-nil, zero value otherwise.

### GetQualifiedIdentifierOk

`func (o *OrderCreateDto) GetQualifiedIdentifierOk() (*string, bool)`

GetQualifiedIdentifierOk returns a tuple with the QualifiedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedIdentifier

`func (o *OrderCreateDto) SetQualifiedIdentifier(v string)`

SetQualifiedIdentifier sets QualifiedIdentifier field to given value.

### HasQualifiedIdentifier

`func (o *OrderCreateDto) HasQualifiedIdentifier() bool`

HasQualifiedIdentifier returns a boolean if a field has been set.

### SetQualifiedIdentifierNil

`func (o *OrderCreateDto) SetQualifiedIdentifierNil(b bool)`

 SetQualifiedIdentifierNil sets the value for QualifiedIdentifier to be an explicit nil

### UnsetQualifiedIdentifier
`func (o *OrderCreateDto) UnsetQualifiedIdentifier()`

UnsetQualifiedIdentifier ensures that no value is present for QualifiedIdentifier, not even an explicit nil
### GetTotalTaxesInUsd

`func (o *OrderCreateDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *OrderCreateDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *OrderCreateDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *OrderCreateDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *OrderCreateDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *OrderCreateDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *OrderCreateDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *OrderCreateDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *OrderCreateDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *OrderCreateDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *OrderCreateDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *OrderCreateDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *OrderCreateDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *OrderCreateDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *OrderCreateDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *OrderCreateDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxInUsd

`func (o *OrderCreateDto) GetTotalShippingTaxInUsd() float64`

GetTotalShippingTaxInUsd returns the TotalShippingTaxInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxInUsdOk

`func (o *OrderCreateDto) GetTotalShippingTaxInUsdOk() (*float64, bool)`

GetTotalShippingTaxInUsdOk returns a tuple with the TotalShippingTaxInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxInUsd

`func (o *OrderCreateDto) SetTotalShippingTaxInUsd(v float64)`

SetTotalShippingTaxInUsd sets TotalShippingTaxInUsd field to given value.

### HasTotalShippingTaxInUsd

`func (o *OrderCreateDto) HasTotalShippingTaxInUsd() bool`

HasTotalShippingTaxInUsd returns a boolean if a field has been set.

### GetTotalAmountInUsd

`func (o *OrderCreateDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *OrderCreateDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *OrderCreateDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *OrderCreateDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *OrderCreateDto) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *OrderCreateDto) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *OrderCreateDto) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *OrderCreateDto) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *OrderCreateDto) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *OrderCreateDto) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetEffectiveFrom

`func (o *OrderCreateDto) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *OrderCreateDto) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *OrderCreateDto) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *OrderCreateDto) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### SetEffectiveFromNil

`func (o *OrderCreateDto) SetEffectiveFromNil(b bool)`

 SetEffectiveFromNil sets the value for EffectiveFrom to be an explicit nil

### UnsetEffectiveFrom
`func (o *OrderCreateDto) UnsetEffectiveFrom()`

UnsetEffectiveFrom ensures that no value is present for EffectiveFrom, not even an explicit nil
### GetOrderLines

`func (o *OrderCreateDto) GetOrderLines() []OrderLineCreateDto`

GetOrderLines returns the OrderLines field if non-nil, zero value otherwise.

### GetOrderLinesOk

`func (o *OrderCreateDto) GetOrderLinesOk() (*[]OrderLineCreateDto, bool)`

GetOrderLinesOk returns a tuple with the OrderLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLines

`func (o *OrderCreateDto) SetOrderLines(v []OrderLineCreateDto)`

SetOrderLines sets OrderLines field to given value.

### HasOrderLines

`func (o *OrderCreateDto) HasOrderLines() bool`

HasOrderLines returns a boolean if a field has been set.

### SetOrderLinesNil

`func (o *OrderCreateDto) SetOrderLinesNil(b bool)`

 SetOrderLinesNil sets the value for OrderLines to be an explicit nil

### UnsetOrderLines
`func (o *OrderCreateDto) UnsetOrderLines()`

UnsetOrderLines ensures that no value is present for OrderLines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


