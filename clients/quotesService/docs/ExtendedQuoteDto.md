# ExtendedQuoteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
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
**CustomerNotes** | Pointer to **NullableString** |  | [optional] 
**TaxCalculationMethod** | Pointer to **string** |  | [optional] 
**CostCalculationMethod** | Pointer to **string** |  | [optional] 
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
**TotalTaxBase** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**TotalTaxesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingCost** | Pointer to **float64** |  | [optional] 
**TotalShippingCostCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingTax** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWithheldTax** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**TotalCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDetailInUsd** | Pointer to **float64** |  | [optional] 
**TotalProfitInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalInUsd** | Pointer to **float64** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**DealUnitId** | Pointer to **NullableString** |  | [optional] 
**EffectiveTo** | Pointer to **NullableTime** |  | [optional] 
**EffectiveFrom** | Pointer to **NullableTime** |  | [optional] 
**QuoteStatus** | Pointer to **string** |  | [optional] 
**FreightTerms** | Pointer to **string** |  | [optional] 
**CustomDiscountsAmount** | Pointer to **float64** |  | [optional] 
**User** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**Individual** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 
**Organization** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 
**ReceiverTenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**Enrollment** | Pointer to [**TenantEnrollmentDto**](TenantEnrollmentDto.md) |  | [optional] 

## Methods

### NewExtendedQuoteDto

`func NewExtendedQuoteDto() *ExtendedQuoteDto`

NewExtendedQuoteDto instantiates a new ExtendedQuoteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedQuoteDtoWithDefaults

`func NewExtendedQuoteDtoWithDefaults() *ExtendedQuoteDto`

NewExtendedQuoteDtoWithDefaults instantiates a new ExtendedQuoteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedQuoteDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedQuoteDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedQuoteDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedQuoteDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedQuoteDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedQuoteDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedQuoteDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedQuoteDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedQuoteDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedQuoteDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedQuoteDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedQuoteDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *ExtendedQuoteDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *ExtendedQuoteDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *ExtendedQuoteDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *ExtendedQuoteDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetType

`func (o *ExtendedQuoteDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedQuoteDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedQuoteDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExtendedQuoteDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExtendedQuoteDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExtendedQuoteDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ExtendedQuoteDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExtendedQuoteDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExtendedQuoteDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExtendedQuoteDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ExtendedQuoteDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExtendedQuoteDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *ExtendedQuoteDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExtendedQuoteDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExtendedQuoteDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExtendedQuoteDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExtendedQuoteDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExtendedQuoteDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *ExtendedQuoteDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedQuoteDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedQuoteDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedQuoteDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedQuoteDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedQuoteDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *ExtendedQuoteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedQuoteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedQuoteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedQuoteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExtendedQuoteDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExtendedQuoteDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceListId

`func (o *ExtendedQuoteDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ExtendedQuoteDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ExtendedQuoteDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ExtendedQuoteDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ExtendedQuoteDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ExtendedQuoteDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetEnrollmentId

`func (o *ExtendedQuoteDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ExtendedQuoteDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ExtendedQuoteDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ExtendedQuoteDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ExtendedQuoteDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ExtendedQuoteDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *ExtendedQuoteDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *ExtendedQuoteDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *ExtendedQuoteDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *ExtendedQuoteDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *ExtendedQuoteDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *ExtendedQuoteDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *ExtendedQuoteDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ExtendedQuoteDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ExtendedQuoteDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ExtendedQuoteDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ExtendedQuoteDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ExtendedQuoteDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *ExtendedQuoteDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *ExtendedQuoteDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *ExtendedQuoteDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *ExtendedQuoteDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *ExtendedQuoteDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *ExtendedQuoteDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *ExtendedQuoteDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExtendedQuoteDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExtendedQuoteDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ExtendedQuoteDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ExtendedQuoteDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ExtendedQuoteDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ExtendedQuoteDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExtendedQuoteDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExtendedQuoteDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ExtendedQuoteDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ExtendedQuoteDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ExtendedQuoteDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *ExtendedQuoteDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ExtendedQuoteDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ExtendedQuoteDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ExtendedQuoteDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *ExtendedQuoteDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *ExtendedQuoteDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *ExtendedQuoteDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *ExtendedQuoteDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *ExtendedQuoteDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *ExtendedQuoteDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *ExtendedQuoteDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *ExtendedQuoteDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *ExtendedQuoteDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ExtendedQuoteDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ExtendedQuoteDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ExtendedQuoteDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *ExtendedQuoteDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *ExtendedQuoteDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *ExtendedQuoteDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ExtendedQuoteDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ExtendedQuoteDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ExtendedQuoteDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *ExtendedQuoteDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *ExtendedQuoteDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *ExtendedQuoteDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ExtendedQuoteDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ExtendedQuoteDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ExtendedQuoteDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ExtendedQuoteDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ExtendedQuoteDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *ExtendedQuoteDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ExtendedQuoteDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ExtendedQuoteDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ExtendedQuoteDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ExtendedQuoteDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ExtendedQuoteDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *ExtendedQuoteDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ExtendedQuoteDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ExtendedQuoteDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ExtendedQuoteDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ExtendedQuoteDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ExtendedQuoteDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *ExtendedQuoteDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ExtendedQuoteDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ExtendedQuoteDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ExtendedQuoteDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ExtendedQuoteDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ExtendedQuoteDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCustomerNotes

`func (o *ExtendedQuoteDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *ExtendedQuoteDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *ExtendedQuoteDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *ExtendedQuoteDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *ExtendedQuoteDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *ExtendedQuoteDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetTaxCalculationMethod

`func (o *ExtendedQuoteDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *ExtendedQuoteDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *ExtendedQuoteDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *ExtendedQuoteDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *ExtendedQuoteDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *ExtendedQuoteDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *ExtendedQuoteDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *ExtendedQuoteDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRate

`func (o *ExtendedQuoteDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *ExtendedQuoteDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *ExtendedQuoteDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *ExtendedQuoteDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ExtendedQuoteDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExtendedQuoteDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExtendedQuoteDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExtendedQuoteDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExtendedQuoteDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExtendedQuoteDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *ExtendedQuoteDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *ExtendedQuoteDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *ExtendedQuoteDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *ExtendedQuoteDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *ExtendedQuoteDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *ExtendedQuoteDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *ExtendedQuoteDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *ExtendedQuoteDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *ExtendedQuoteDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *ExtendedQuoteDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *ExtendedQuoteDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *ExtendedQuoteDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *ExtendedQuoteDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *ExtendedQuoteDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *ExtendedQuoteDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ExtendedQuoteDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *ExtendedQuoteDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *ExtendedQuoteDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *ExtendedQuoteDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *ExtendedQuoteDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *ExtendedQuoteDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *ExtendedQuoteDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *ExtendedQuoteDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *ExtendedQuoteDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *ExtendedQuoteDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *ExtendedQuoteDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *ExtendedQuoteDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *ExtendedQuoteDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *ExtendedQuoteDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *ExtendedQuoteDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *ExtendedQuoteDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *ExtendedQuoteDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *ExtendedQuoteDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *ExtendedQuoteDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *ExtendedQuoteDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *ExtendedQuoteDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *ExtendedQuoteDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *ExtendedQuoteDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *ExtendedQuoteDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *ExtendedQuoteDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *ExtendedQuoteDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *ExtendedQuoteDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *ExtendedQuoteDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *ExtendedQuoteDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *ExtendedQuoteDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *ExtendedQuoteDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *ExtendedQuoteDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *ExtendedQuoteDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *ExtendedQuoteDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *ExtendedQuoteDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *ExtendedQuoteDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *ExtendedQuoteDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *ExtendedQuoteDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *ExtendedQuoteDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *ExtendedQuoteDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *ExtendedQuoteDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *ExtendedQuoteDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *ExtendedQuoteDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *ExtendedQuoteDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *ExtendedQuoteDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *ExtendedQuoteDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *ExtendedQuoteDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *ExtendedQuoteDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotal

`func (o *ExtendedQuoteDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExtendedQuoteDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExtendedQuoteDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExtendedQuoteDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *ExtendedQuoteDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *ExtendedQuoteDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *ExtendedQuoteDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *ExtendedQuoteDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *ExtendedQuoteDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *ExtendedQuoteDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetTotalDetailInUsd

`func (o *ExtendedQuoteDto) GetTotalDetailInUsd() float64`

GetTotalDetailInUsd returns the TotalDetailInUsd field if non-nil, zero value otherwise.

### GetTotalDetailInUsdOk

`func (o *ExtendedQuoteDto) GetTotalDetailInUsdOk() (*float64, bool)`

GetTotalDetailInUsdOk returns a tuple with the TotalDetailInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailInUsd

`func (o *ExtendedQuoteDto) SetTotalDetailInUsd(v float64)`

SetTotalDetailInUsd sets TotalDetailInUsd field to given value.

### HasTotalDetailInUsd

`func (o *ExtendedQuoteDto) HasTotalDetailInUsd() bool`

HasTotalDetailInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *ExtendedQuoteDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *ExtendedQuoteDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *ExtendedQuoteDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *ExtendedQuoteDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *ExtendedQuoteDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *ExtendedQuoteDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *ExtendedQuoteDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *ExtendedQuoteDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *ExtendedQuoteDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *ExtendedQuoteDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *ExtendedQuoteDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *ExtendedQuoteDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *ExtendedQuoteDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *ExtendedQuoteDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *ExtendedQuoteDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *ExtendedQuoteDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *ExtendedQuoteDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *ExtendedQuoteDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *ExtendedQuoteDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *ExtendedQuoteDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *ExtendedQuoteDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *ExtendedQuoteDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *ExtendedQuoteDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *ExtendedQuoteDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *ExtendedQuoteDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *ExtendedQuoteDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *ExtendedQuoteDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *ExtendedQuoteDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *ExtendedQuoteDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *ExtendedQuoteDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *ExtendedQuoteDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *ExtendedQuoteDto) GetTotalInUsd() float64`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *ExtendedQuoteDto) GetTotalInUsdOk() (*float64, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *ExtendedQuoteDto) SetTotalInUsd(v float64)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *ExtendedQuoteDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetCartId

`func (o *ExtendedQuoteDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ExtendedQuoteDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ExtendedQuoteDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ExtendedQuoteDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ExtendedQuoteDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ExtendedQuoteDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetDealUnitId

`func (o *ExtendedQuoteDto) GetDealUnitId() string`

GetDealUnitId returns the DealUnitId field if non-nil, zero value otherwise.

### GetDealUnitIdOk

`func (o *ExtendedQuoteDto) GetDealUnitIdOk() (*string, bool)`

GetDealUnitIdOk returns a tuple with the DealUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitId

`func (o *ExtendedQuoteDto) SetDealUnitId(v string)`

SetDealUnitId sets DealUnitId field to given value.

### HasDealUnitId

`func (o *ExtendedQuoteDto) HasDealUnitId() bool`

HasDealUnitId returns a boolean if a field has been set.

### SetDealUnitIdNil

`func (o *ExtendedQuoteDto) SetDealUnitIdNil(b bool)`

 SetDealUnitIdNil sets the value for DealUnitId to be an explicit nil

### UnsetDealUnitId
`func (o *ExtendedQuoteDto) UnsetDealUnitId()`

UnsetDealUnitId ensures that no value is present for DealUnitId, not even an explicit nil
### GetEffectiveTo

`func (o *ExtendedQuoteDto) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *ExtendedQuoteDto) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *ExtendedQuoteDto) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *ExtendedQuoteDto) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *ExtendedQuoteDto) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *ExtendedQuoteDto) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetEffectiveFrom

`func (o *ExtendedQuoteDto) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ExtendedQuoteDto) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ExtendedQuoteDto) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ExtendedQuoteDto) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### SetEffectiveFromNil

`func (o *ExtendedQuoteDto) SetEffectiveFromNil(b bool)`

 SetEffectiveFromNil sets the value for EffectiveFrom to be an explicit nil

### UnsetEffectiveFrom
`func (o *ExtendedQuoteDto) UnsetEffectiveFrom()`

UnsetEffectiveFrom ensures that no value is present for EffectiveFrom, not even an explicit nil
### GetQuoteStatus

`func (o *ExtendedQuoteDto) GetQuoteStatus() string`

GetQuoteStatus returns the QuoteStatus field if non-nil, zero value otherwise.

### GetQuoteStatusOk

`func (o *ExtendedQuoteDto) GetQuoteStatusOk() (*string, bool)`

GetQuoteStatusOk returns a tuple with the QuoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStatus

`func (o *ExtendedQuoteDto) SetQuoteStatus(v string)`

SetQuoteStatus sets QuoteStatus field to given value.

### HasQuoteStatus

`func (o *ExtendedQuoteDto) HasQuoteStatus() bool`

HasQuoteStatus returns a boolean if a field has been set.

### GetFreightTerms

`func (o *ExtendedQuoteDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *ExtendedQuoteDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *ExtendedQuoteDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *ExtendedQuoteDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### GetCustomDiscountsAmount

`func (o *ExtendedQuoteDto) GetCustomDiscountsAmount() float64`

GetCustomDiscountsAmount returns the CustomDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomDiscountsAmountOk

`func (o *ExtendedQuoteDto) GetCustomDiscountsAmountOk() (*float64, bool)`

GetCustomDiscountsAmountOk returns a tuple with the CustomDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDiscountsAmount

`func (o *ExtendedQuoteDto) SetCustomDiscountsAmount(v float64)`

SetCustomDiscountsAmount sets CustomDiscountsAmount field to given value.

### HasCustomDiscountsAmount

`func (o *ExtendedQuoteDto) HasCustomDiscountsAmount() bool`

HasCustomDiscountsAmount returns a boolean if a field has been set.

### GetUser

`func (o *ExtendedQuoteDto) GetUser() UserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExtendedQuoteDto) GetUserOk() (*UserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExtendedQuoteDto) SetUser(v UserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExtendedQuoteDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTenant

`func (o *ExtendedQuoteDto) GetTenant() TenantDto`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExtendedQuoteDto) GetTenantOk() (*TenantDto, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExtendedQuoteDto) SetTenant(v TenantDto)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ExtendedQuoteDto) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetIndividual

`func (o *ExtendedQuoteDto) GetIndividual() ContactDto`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *ExtendedQuoteDto) GetIndividualOk() (*ContactDto, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *ExtendedQuoteDto) SetIndividual(v ContactDto)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *ExtendedQuoteDto) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *ExtendedQuoteDto) GetOrganization() ContactDto`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ExtendedQuoteDto) GetOrganizationOk() (*ContactDto, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ExtendedQuoteDto) SetOrganization(v ContactDto)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ExtendedQuoteDto) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetReceiverTenant

`func (o *ExtendedQuoteDto) GetReceiverTenant() TenantDto`

GetReceiverTenant returns the ReceiverTenant field if non-nil, zero value otherwise.

### GetReceiverTenantOk

`func (o *ExtendedQuoteDto) GetReceiverTenantOk() (*TenantDto, bool)`

GetReceiverTenantOk returns a tuple with the ReceiverTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenant

`func (o *ExtendedQuoteDto) SetReceiverTenant(v TenantDto)`

SetReceiverTenant sets ReceiverTenant field to given value.

### HasReceiverTenant

`func (o *ExtendedQuoteDto) HasReceiverTenant() bool`

HasReceiverTenant returns a boolean if a field has been set.

### GetEnrollment

`func (o *ExtendedQuoteDto) GetEnrollment() TenantEnrollmentDto`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ExtendedQuoteDto) GetEnrollmentOk() (*TenantEnrollmentDto, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ExtendedQuoteDto) SetEnrollment(v TenantEnrollmentDto)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *ExtendedQuoteDto) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


