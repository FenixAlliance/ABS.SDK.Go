# DealUnitCreateDto

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
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
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
**DealUnitFlowId** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowStageId** | Pointer to **NullableString** |  | [optional] 
**PartnerCreated** | Pointer to **bool** |  | [optional] 
**PartnerCollaboration** | Pointer to **bool** |  | [optional] 
**ProposedSolution** | Pointer to **NullableString** |  | [optional] 
**CurrentSituation** | Pointer to **NullableString** |  | [optional] 
**CustomerNeed** | Pointer to **NullableString** |  | [optional] 
**WonDate** | Pointer to **time.Time** |  | [optional] 
**LostDate** | Pointer to **time.Time** |  | [optional] 
**ExpiryDate** | Pointer to **time.Time** |  | [optional] 
**DeliveredDate** | Pointer to **time.Time** |  | [optional] 
**ClosedTimestamp** | Pointer to **time.Time** |  | [optional] 
**ExpectedCloseDate** | Pointer to **time.Time** |  | [optional] 
**DealUnitStatus** | Pointer to **string** |  | [optional] 
**DealUnitPurchaseProcess** | Pointer to **string** |  | [optional] 
**DealUnitForecastCategory** | Pointer to **string** |  | [optional] 
**DealUnitAmountsCalculation** | Pointer to **string** |  | [optional] 
**DealUnitLines** | Pointer to [**[]DealUnitLineCreateDto**](DealUnitLineCreateDto.md) |  | [optional] 

## Methods

### NewDealUnitCreateDto

`func NewDealUnitCreateDto() *DealUnitCreateDto`

NewDealUnitCreateDto instantiates a new DealUnitCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitCreateDtoWithDefaults

`func NewDealUnitCreateDtoWithDefaults() *DealUnitCreateDto`

NewDealUnitCreateDtoWithDefaults instantiates a new DealUnitCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *DealUnitCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetClosed

`func (o *DealUnitCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *DealUnitCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *DealUnitCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *DealUnitCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetTitle

`func (o *DealUnitCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DealUnitCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DealUnitCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DealUnitCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DealUnitCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DealUnitCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriceListId

`func (o *DealUnitCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *DealUnitCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *DealUnitCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *DealUnitCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *DealUnitCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *DealUnitCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDescription

`func (o *DealUnitCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIndividualId

`func (o *DealUnitCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *DealUnitCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *DealUnitCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *DealUnitCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *DealUnitCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *DealUnitCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetPaymentTermId

`func (o *DealUnitCreateDto) GetPaymentTermId() string`

GetPaymentTermId returns the PaymentTermId field if non-nil, zero value otherwise.

### GetPaymentTermIdOk

`func (o *DealUnitCreateDto) GetPaymentTermIdOk() (*string, bool)`

GetPaymentTermIdOk returns a tuple with the PaymentTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermId

`func (o *DealUnitCreateDto) SetPaymentTermId(v string)`

SetPaymentTermId sets PaymentTermId field to given value.

### HasPaymentTermId

`func (o *DealUnitCreateDto) HasPaymentTermId() bool`

HasPaymentTermId returns a boolean if a field has been set.

### SetPaymentTermIdNil

`func (o *DealUnitCreateDto) SetPaymentTermIdNil(b bool)`

 SetPaymentTermIdNil sets the value for PaymentTermId to be an explicit nil

### UnsetPaymentTermId
`func (o *DealUnitCreateDto) UnsetPaymentTermId()`

UnsetPaymentTermId ensures that no value is present for PaymentTermId, not even an explicit nil
### GetOrganizationId

`func (o *DealUnitCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DealUnitCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DealUnitCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DealUnitCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *DealUnitCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DealUnitCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *DealUnitCreateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *DealUnitCreateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *DealUnitCreateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *DealUnitCreateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *DealUnitCreateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *DealUnitCreateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *DealUnitCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DealUnitCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DealUnitCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DealUnitCreateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *DealUnitCreateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *DealUnitCreateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *DealUnitCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DealUnitCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DealUnitCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DealUnitCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *DealUnitCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *DealUnitCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *DealUnitCreateDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DealUnitCreateDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DealUnitCreateDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DealUnitCreateDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *DealUnitCreateDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *DealUnitCreateDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *DealUnitCreateDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *DealUnitCreateDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *DealUnitCreateDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *DealUnitCreateDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *DealUnitCreateDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *DealUnitCreateDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *DealUnitCreateDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DealUnitCreateDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DealUnitCreateDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DealUnitCreateDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *DealUnitCreateDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *DealUnitCreateDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *DealUnitCreateDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DealUnitCreateDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DealUnitCreateDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DealUnitCreateDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *DealUnitCreateDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *DealUnitCreateDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *DealUnitCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DealUnitCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DealUnitCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *DealUnitCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *DealUnitCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *DealUnitCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *DealUnitCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *DealUnitCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *DealUnitCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *DealUnitCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *DealUnitCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *DealUnitCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *DealUnitCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *DealUnitCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *DealUnitCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *DealUnitCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *DealUnitCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *DealUnitCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *DealUnitCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *DealUnitCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *DealUnitCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *DealUnitCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *DealUnitCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *DealUnitCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetForexRate

`func (o *DealUnitCreateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *DealUnitCreateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *DealUnitCreateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *DealUnitCreateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *DealUnitCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *DealUnitCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *DealUnitCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *DealUnitCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *DealUnitCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *DealUnitCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *DealUnitCreateDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *DealUnitCreateDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *DealUnitCreateDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *DealUnitCreateDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *DealUnitCreateDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *DealUnitCreateDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *DealUnitCreateDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *DealUnitCreateDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *DealUnitCreateDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *DealUnitCreateDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *DealUnitCreateDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *DealUnitCreateDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *DealUnitCreateDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *DealUnitCreateDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *DealUnitCreateDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *DealUnitCreateDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *DealUnitCreateDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *DealUnitCreateDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *DealUnitCreateDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *DealUnitCreateDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *DealUnitCreateDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *DealUnitCreateDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *DealUnitCreateDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *DealUnitCreateDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *DealUnitCreateDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *DealUnitCreateDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *DealUnitCreateDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *DealUnitCreateDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *DealUnitCreateDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *DealUnitCreateDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *DealUnitCreateDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *DealUnitCreateDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *DealUnitCreateDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *DealUnitCreateDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *DealUnitCreateDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *DealUnitCreateDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *DealUnitCreateDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *DealUnitCreateDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *DealUnitCreateDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *DealUnitCreateDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *DealUnitCreateDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *DealUnitCreateDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *DealUnitCreateDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *DealUnitCreateDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *DealUnitCreateDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *DealUnitCreateDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *DealUnitCreateDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *DealUnitCreateDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *DealUnitCreateDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *DealUnitCreateDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *DealUnitCreateDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *DealUnitCreateDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *DealUnitCreateDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *DealUnitCreateDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *DealUnitCreateDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *DealUnitCreateDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *DealUnitCreateDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *DealUnitCreateDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *DealUnitCreateDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *DealUnitCreateDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *DealUnitCreateDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *DealUnitCreateDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *DealUnitCreateDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *DealUnitCreateDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *DealUnitCreateDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *DealUnitCreateDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *DealUnitCreateDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *DealUnitCreateDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *DealUnitCreateDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *DealUnitCreateDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *DealUnitCreateDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *DealUnitCreateDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *DealUnitCreateDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *DealUnitCreateDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *DealUnitCreateDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *DealUnitCreateDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *DealUnitCreateDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotal

`func (o *DealUnitCreateDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DealUnitCreateDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DealUnitCreateDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DealUnitCreateDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *DealUnitCreateDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *DealUnitCreateDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *DealUnitCreateDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *DealUnitCreateDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *DealUnitCreateDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *DealUnitCreateDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetCostCalculationMethod

`func (o *DealUnitCreateDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *DealUnitCreateDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *DealUnitCreateDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *DealUnitCreateDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetTaxCalculationMethod

`func (o *DealUnitCreateDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *DealUnitCreateDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *DealUnitCreateDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *DealUnitCreateDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetDealUnitFlowId

`func (o *DealUnitCreateDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitCreateDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitCreateDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitCreateDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitCreateDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitCreateDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetDealUnitFlowStageId

`func (o *DealUnitCreateDto) GetDealUnitFlowStageId() string`

GetDealUnitFlowStageId returns the DealUnitFlowStageId field if non-nil, zero value otherwise.

### GetDealUnitFlowStageIdOk

`func (o *DealUnitCreateDto) GetDealUnitFlowStageIdOk() (*string, bool)`

GetDealUnitFlowStageIdOk returns a tuple with the DealUnitFlowStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowStageId

`func (o *DealUnitCreateDto) SetDealUnitFlowStageId(v string)`

SetDealUnitFlowStageId sets DealUnitFlowStageId field to given value.

### HasDealUnitFlowStageId

`func (o *DealUnitCreateDto) HasDealUnitFlowStageId() bool`

HasDealUnitFlowStageId returns a boolean if a field has been set.

### SetDealUnitFlowStageIdNil

`func (o *DealUnitCreateDto) SetDealUnitFlowStageIdNil(b bool)`

 SetDealUnitFlowStageIdNil sets the value for DealUnitFlowStageId to be an explicit nil

### UnsetDealUnitFlowStageId
`func (o *DealUnitCreateDto) UnsetDealUnitFlowStageId()`

UnsetDealUnitFlowStageId ensures that no value is present for DealUnitFlowStageId, not even an explicit nil
### GetPartnerCreated

`func (o *DealUnitCreateDto) GetPartnerCreated() bool`

GetPartnerCreated returns the PartnerCreated field if non-nil, zero value otherwise.

### GetPartnerCreatedOk

`func (o *DealUnitCreateDto) GetPartnerCreatedOk() (*bool, bool)`

GetPartnerCreatedOk returns a tuple with the PartnerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCreated

`func (o *DealUnitCreateDto) SetPartnerCreated(v bool)`

SetPartnerCreated sets PartnerCreated field to given value.

### HasPartnerCreated

`func (o *DealUnitCreateDto) HasPartnerCreated() bool`

HasPartnerCreated returns a boolean if a field has been set.

### GetPartnerCollaboration

`func (o *DealUnitCreateDto) GetPartnerCollaboration() bool`

GetPartnerCollaboration returns the PartnerCollaboration field if non-nil, zero value otherwise.

### GetPartnerCollaborationOk

`func (o *DealUnitCreateDto) GetPartnerCollaborationOk() (*bool, bool)`

GetPartnerCollaborationOk returns a tuple with the PartnerCollaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCollaboration

`func (o *DealUnitCreateDto) SetPartnerCollaboration(v bool)`

SetPartnerCollaboration sets PartnerCollaboration field to given value.

### HasPartnerCollaboration

`func (o *DealUnitCreateDto) HasPartnerCollaboration() bool`

HasPartnerCollaboration returns a boolean if a field has been set.

### GetProposedSolution

`func (o *DealUnitCreateDto) GetProposedSolution() string`

GetProposedSolution returns the ProposedSolution field if non-nil, zero value otherwise.

### GetProposedSolutionOk

`func (o *DealUnitCreateDto) GetProposedSolutionOk() (*string, bool)`

GetProposedSolutionOk returns a tuple with the ProposedSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedSolution

`func (o *DealUnitCreateDto) SetProposedSolution(v string)`

SetProposedSolution sets ProposedSolution field to given value.

### HasProposedSolution

`func (o *DealUnitCreateDto) HasProposedSolution() bool`

HasProposedSolution returns a boolean if a field has been set.

### SetProposedSolutionNil

`func (o *DealUnitCreateDto) SetProposedSolutionNil(b bool)`

 SetProposedSolutionNil sets the value for ProposedSolution to be an explicit nil

### UnsetProposedSolution
`func (o *DealUnitCreateDto) UnsetProposedSolution()`

UnsetProposedSolution ensures that no value is present for ProposedSolution, not even an explicit nil
### GetCurrentSituation

`func (o *DealUnitCreateDto) GetCurrentSituation() string`

GetCurrentSituation returns the CurrentSituation field if non-nil, zero value otherwise.

### GetCurrentSituationOk

`func (o *DealUnitCreateDto) GetCurrentSituationOk() (*string, bool)`

GetCurrentSituationOk returns a tuple with the CurrentSituation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSituation

`func (o *DealUnitCreateDto) SetCurrentSituation(v string)`

SetCurrentSituation sets CurrentSituation field to given value.

### HasCurrentSituation

`func (o *DealUnitCreateDto) HasCurrentSituation() bool`

HasCurrentSituation returns a boolean if a field has been set.

### SetCurrentSituationNil

`func (o *DealUnitCreateDto) SetCurrentSituationNil(b bool)`

 SetCurrentSituationNil sets the value for CurrentSituation to be an explicit nil

### UnsetCurrentSituation
`func (o *DealUnitCreateDto) UnsetCurrentSituation()`

UnsetCurrentSituation ensures that no value is present for CurrentSituation, not even an explicit nil
### GetCustomerNeed

`func (o *DealUnitCreateDto) GetCustomerNeed() string`

GetCustomerNeed returns the CustomerNeed field if non-nil, zero value otherwise.

### GetCustomerNeedOk

`func (o *DealUnitCreateDto) GetCustomerNeedOk() (*string, bool)`

GetCustomerNeedOk returns a tuple with the CustomerNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNeed

`func (o *DealUnitCreateDto) SetCustomerNeed(v string)`

SetCustomerNeed sets CustomerNeed field to given value.

### HasCustomerNeed

`func (o *DealUnitCreateDto) HasCustomerNeed() bool`

HasCustomerNeed returns a boolean if a field has been set.

### SetCustomerNeedNil

`func (o *DealUnitCreateDto) SetCustomerNeedNil(b bool)`

 SetCustomerNeedNil sets the value for CustomerNeed to be an explicit nil

### UnsetCustomerNeed
`func (o *DealUnitCreateDto) UnsetCustomerNeed()`

UnsetCustomerNeed ensures that no value is present for CustomerNeed, not even an explicit nil
### GetWonDate

`func (o *DealUnitCreateDto) GetWonDate() time.Time`

GetWonDate returns the WonDate field if non-nil, zero value otherwise.

### GetWonDateOk

`func (o *DealUnitCreateDto) GetWonDateOk() (*time.Time, bool)`

GetWonDateOk returns a tuple with the WonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWonDate

`func (o *DealUnitCreateDto) SetWonDate(v time.Time)`

SetWonDate sets WonDate field to given value.

### HasWonDate

`func (o *DealUnitCreateDto) HasWonDate() bool`

HasWonDate returns a boolean if a field has been set.

### GetLostDate

`func (o *DealUnitCreateDto) GetLostDate() time.Time`

GetLostDate returns the LostDate field if non-nil, zero value otherwise.

### GetLostDateOk

`func (o *DealUnitCreateDto) GetLostDateOk() (*time.Time, bool)`

GetLostDateOk returns a tuple with the LostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostDate

`func (o *DealUnitCreateDto) SetLostDate(v time.Time)`

SetLostDate sets LostDate field to given value.

### HasLostDate

`func (o *DealUnitCreateDto) HasLostDate() bool`

HasLostDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *DealUnitCreateDto) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *DealUnitCreateDto) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *DealUnitCreateDto) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *DealUnitCreateDto) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *DealUnitCreateDto) GetDeliveredDate() time.Time`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *DealUnitCreateDto) GetDeliveredDateOk() (*time.Time, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *DealUnitCreateDto) SetDeliveredDate(v time.Time)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *DealUnitCreateDto) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### GetClosedTimestamp

`func (o *DealUnitCreateDto) GetClosedTimestamp() time.Time`

GetClosedTimestamp returns the ClosedTimestamp field if non-nil, zero value otherwise.

### GetClosedTimestampOk

`func (o *DealUnitCreateDto) GetClosedTimestampOk() (*time.Time, bool)`

GetClosedTimestampOk returns a tuple with the ClosedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTimestamp

`func (o *DealUnitCreateDto) SetClosedTimestamp(v time.Time)`

SetClosedTimestamp sets ClosedTimestamp field to given value.

### HasClosedTimestamp

`func (o *DealUnitCreateDto) HasClosedTimestamp() bool`

HasClosedTimestamp returns a boolean if a field has been set.

### GetExpectedCloseDate

`func (o *DealUnitCreateDto) GetExpectedCloseDate() time.Time`

GetExpectedCloseDate returns the ExpectedCloseDate field if non-nil, zero value otherwise.

### GetExpectedCloseDateOk

`func (o *DealUnitCreateDto) GetExpectedCloseDateOk() (*time.Time, bool)`

GetExpectedCloseDateOk returns a tuple with the ExpectedCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCloseDate

`func (o *DealUnitCreateDto) SetExpectedCloseDate(v time.Time)`

SetExpectedCloseDate sets ExpectedCloseDate field to given value.

### HasExpectedCloseDate

`func (o *DealUnitCreateDto) HasExpectedCloseDate() bool`

HasExpectedCloseDate returns a boolean if a field has been set.

### GetDealUnitStatus

`func (o *DealUnitCreateDto) GetDealUnitStatus() string`

GetDealUnitStatus returns the DealUnitStatus field if non-nil, zero value otherwise.

### GetDealUnitStatusOk

`func (o *DealUnitCreateDto) GetDealUnitStatusOk() (*string, bool)`

GetDealUnitStatusOk returns a tuple with the DealUnitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitStatus

`func (o *DealUnitCreateDto) SetDealUnitStatus(v string)`

SetDealUnitStatus sets DealUnitStatus field to given value.

### HasDealUnitStatus

`func (o *DealUnitCreateDto) HasDealUnitStatus() bool`

HasDealUnitStatus returns a boolean if a field has been set.

### GetDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) GetDealUnitPurchaseProcess() string`

GetDealUnitPurchaseProcess returns the DealUnitPurchaseProcess field if non-nil, zero value otherwise.

### GetDealUnitPurchaseProcessOk

`func (o *DealUnitCreateDto) GetDealUnitPurchaseProcessOk() (*string, bool)`

GetDealUnitPurchaseProcessOk returns a tuple with the DealUnitPurchaseProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) SetDealUnitPurchaseProcess(v string)`

SetDealUnitPurchaseProcess sets DealUnitPurchaseProcess field to given value.

### HasDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) HasDealUnitPurchaseProcess() bool`

HasDealUnitPurchaseProcess returns a boolean if a field has been set.

### GetDealUnitForecastCategory

`func (o *DealUnitCreateDto) GetDealUnitForecastCategory() string`

GetDealUnitForecastCategory returns the DealUnitForecastCategory field if non-nil, zero value otherwise.

### GetDealUnitForecastCategoryOk

`func (o *DealUnitCreateDto) GetDealUnitForecastCategoryOk() (*string, bool)`

GetDealUnitForecastCategoryOk returns a tuple with the DealUnitForecastCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitForecastCategory

`func (o *DealUnitCreateDto) SetDealUnitForecastCategory(v string)`

SetDealUnitForecastCategory sets DealUnitForecastCategory field to given value.

### HasDealUnitForecastCategory

`func (o *DealUnitCreateDto) HasDealUnitForecastCategory() bool`

HasDealUnitForecastCategory returns a boolean if a field has been set.

### GetDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) GetDealUnitAmountsCalculation() string`

GetDealUnitAmountsCalculation returns the DealUnitAmountsCalculation field if non-nil, zero value otherwise.

### GetDealUnitAmountsCalculationOk

`func (o *DealUnitCreateDto) GetDealUnitAmountsCalculationOk() (*string, bool)`

GetDealUnitAmountsCalculationOk returns a tuple with the DealUnitAmountsCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) SetDealUnitAmountsCalculation(v string)`

SetDealUnitAmountsCalculation sets DealUnitAmountsCalculation field to given value.

### HasDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) HasDealUnitAmountsCalculation() bool`

HasDealUnitAmountsCalculation returns a boolean if a field has been set.

### GetDealUnitLines

`func (o *DealUnitCreateDto) GetDealUnitLines() []DealUnitLineCreateDto`

GetDealUnitLines returns the DealUnitLines field if non-nil, zero value otherwise.

### GetDealUnitLinesOk

`func (o *DealUnitCreateDto) GetDealUnitLinesOk() (*[]DealUnitLineCreateDto, bool)`

GetDealUnitLinesOk returns a tuple with the DealUnitLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitLines

`func (o *DealUnitCreateDto) SetDealUnitLines(v []DealUnitLineCreateDto)`

SetDealUnitLines sets DealUnitLines field to given value.

### HasDealUnitLines

`func (o *DealUnitCreateDto) HasDealUnitLines() bool`

HasDealUnitLines returns a boolean if a field has been set.

### SetDealUnitLinesNil

`func (o *DealUnitCreateDto) SetDealUnitLinesNil(b bool)`

 SetDealUnitLinesNil sets the value for DealUnitLines to be an explicit nil

### UnsetDealUnitLines
`func (o *DealUnitCreateDto) UnsetDealUnitLines()`

UnsetDealUnitLines ensures that no value is present for DealUnitLines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


