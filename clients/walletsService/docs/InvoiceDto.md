# InvoiceDto

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
**Paid** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **int64** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**Enumeration** | Pointer to **NullableString** |  | [optional] 
**PaymentModeId** | Pointer to **NullableString** |  | [optional] 
**EnumerationRangeId** | Pointer to **NullableString** |  | [optional] 
**EmisorBillingProfileId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBillingProfileId** | Pointer to **NullableString** |  | [optional] 
**EmisorWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**PaymentDue** | Pointer to **NullableTime** |  | [optional] 
**InvoiceType** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**InvoiceStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceDto

`func NewInvoiceDto() *InvoiceDto`

NewInvoiceDto instantiates a new InvoiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDtoWithDefaults

`func NewInvoiceDtoWithDefaults() *InvoiceDto`

NewInvoiceDtoWithDefaults instantiates a new InvoiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InvoiceDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InvoiceDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InvoiceDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *InvoiceDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *InvoiceDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *InvoiceDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *InvoiceDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetType

`func (o *InvoiceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InvoiceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InvoiceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InvoiceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvoiceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvoiceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *InvoiceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InvoiceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InvoiceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *InvoiceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *InvoiceDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InvoiceDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *InvoiceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *InvoiceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceListId

`func (o *InvoiceDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *InvoiceDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *InvoiceDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *InvoiceDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *InvoiceDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *InvoiceDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *InvoiceDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *InvoiceDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *InvoiceDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *InvoiceDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *InvoiceDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *InvoiceDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *InvoiceDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InvoiceDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InvoiceDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InvoiceDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *InvoiceDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *InvoiceDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *InvoiceDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *InvoiceDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *InvoiceDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *InvoiceDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *InvoiceDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *InvoiceDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *InvoiceDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InvoiceDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InvoiceDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InvoiceDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *InvoiceDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *InvoiceDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *InvoiceDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InvoiceDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InvoiceDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InvoiceDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *InvoiceDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *InvoiceDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *InvoiceDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvoiceDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InvoiceDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InvoiceDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *InvoiceDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *InvoiceDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *InvoiceDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *InvoiceDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *InvoiceDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *InvoiceDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *InvoiceDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *InvoiceDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *InvoiceDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *InvoiceDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *InvoiceDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *InvoiceDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *InvoiceDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *InvoiceDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *InvoiceDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *InvoiceDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *InvoiceDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *InvoiceDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *InvoiceDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *InvoiceDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *InvoiceDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *InvoiceDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *InvoiceDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *InvoiceDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *InvoiceDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *InvoiceDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *InvoiceDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *InvoiceDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *InvoiceDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *InvoiceDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *InvoiceDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *InvoiceDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *InvoiceDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *InvoiceDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *InvoiceDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *InvoiceDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *InvoiceDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *InvoiceDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *InvoiceDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *InvoiceDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *InvoiceDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *InvoiceDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCustomerNotes

`func (o *InvoiceDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *InvoiceDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *InvoiceDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *InvoiceDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *InvoiceDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *InvoiceDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetTaxCalculationMethod

`func (o *InvoiceDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *InvoiceDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *InvoiceDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *InvoiceDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *InvoiceDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *InvoiceDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *InvoiceDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *InvoiceDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRate

`func (o *InvoiceDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *InvoiceDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *InvoiceDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *InvoiceDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *InvoiceDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *InvoiceDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *InvoiceDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *InvoiceDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *InvoiceDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *InvoiceDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *InvoiceDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *InvoiceDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *InvoiceDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *InvoiceDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *InvoiceDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *InvoiceDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *InvoiceDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *InvoiceDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *InvoiceDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *InvoiceDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *InvoiceDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *InvoiceDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *InvoiceDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *InvoiceDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *InvoiceDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *InvoiceDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *InvoiceDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *InvoiceDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *InvoiceDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *InvoiceDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *InvoiceDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *InvoiceDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *InvoiceDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *InvoiceDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *InvoiceDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *InvoiceDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *InvoiceDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *InvoiceDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *InvoiceDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *InvoiceDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *InvoiceDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *InvoiceDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *InvoiceDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *InvoiceDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *InvoiceDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *InvoiceDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *InvoiceDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *InvoiceDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *InvoiceDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *InvoiceDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *InvoiceDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *InvoiceDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *InvoiceDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *InvoiceDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *InvoiceDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *InvoiceDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *InvoiceDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *InvoiceDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *InvoiceDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *InvoiceDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *InvoiceDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *InvoiceDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *InvoiceDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *InvoiceDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *InvoiceDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *InvoiceDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *InvoiceDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *InvoiceDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *InvoiceDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *InvoiceDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *InvoiceDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *InvoiceDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *InvoiceDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *InvoiceDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *InvoiceDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *InvoiceDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *InvoiceDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *InvoiceDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *InvoiceDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *InvoiceDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *InvoiceDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *InvoiceDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *InvoiceDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *InvoiceDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *InvoiceDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *InvoiceDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *InvoiceDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *InvoiceDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *InvoiceDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *InvoiceDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *InvoiceDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *InvoiceDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *InvoiceDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *InvoiceDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *InvoiceDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *InvoiceDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *InvoiceDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *InvoiceDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *InvoiceDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *InvoiceDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *InvoiceDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *InvoiceDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *InvoiceDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *InvoiceDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *InvoiceDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *InvoiceDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *InvoiceDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *InvoiceDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *InvoiceDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *InvoiceDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *InvoiceDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *InvoiceDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *InvoiceDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *InvoiceDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *InvoiceDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotal

`func (o *InvoiceDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InvoiceDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *InvoiceDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *InvoiceDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *InvoiceDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *InvoiceDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *InvoiceDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *InvoiceDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetTotalDetailInUsd

`func (o *InvoiceDto) GetTotalDetailInUsd() float64`

GetTotalDetailInUsd returns the TotalDetailInUsd field if non-nil, zero value otherwise.

### GetTotalDetailInUsdOk

`func (o *InvoiceDto) GetTotalDetailInUsdOk() (*float64, bool)`

GetTotalDetailInUsdOk returns a tuple with the TotalDetailInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailInUsd

`func (o *InvoiceDto) SetTotalDetailInUsd(v float64)`

SetTotalDetailInUsd sets TotalDetailInUsd field to given value.

### HasTotalDetailInUsd

`func (o *InvoiceDto) HasTotalDetailInUsd() bool`

HasTotalDetailInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *InvoiceDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *InvoiceDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *InvoiceDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *InvoiceDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *InvoiceDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *InvoiceDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *InvoiceDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *InvoiceDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *InvoiceDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *InvoiceDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *InvoiceDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *InvoiceDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *InvoiceDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *InvoiceDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *InvoiceDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *InvoiceDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *InvoiceDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *InvoiceDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *InvoiceDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *InvoiceDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *InvoiceDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *InvoiceDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *InvoiceDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *InvoiceDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *InvoiceDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *InvoiceDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *InvoiceDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *InvoiceDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *InvoiceDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *InvoiceDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *InvoiceDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *InvoiceDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *InvoiceDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *InvoiceDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *InvoiceDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *InvoiceDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *InvoiceDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *InvoiceDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *InvoiceDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *InvoiceDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *InvoiceDto) GetTotalInUsd() float64`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *InvoiceDto) GetTotalInUsdOk() (*float64, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *InvoiceDto) SetTotalInUsd(v float64)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *InvoiceDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetPaid

`func (o *InvoiceDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *InvoiceDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *InvoiceDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *InvoiceDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetNumber

`func (o *InvoiceDto) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceDto) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceDto) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InvoiceDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InvoiceDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOrderId

`func (o *InvoiceDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InvoiceDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InvoiceDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *InvoiceDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *InvoiceDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *InvoiceDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetEnumeration

`func (o *InvoiceDto) GetEnumeration() string`

GetEnumeration returns the Enumeration field if non-nil, zero value otherwise.

### GetEnumerationOk

`func (o *InvoiceDto) GetEnumerationOk() (*string, bool)`

GetEnumerationOk returns a tuple with the Enumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumeration

`func (o *InvoiceDto) SetEnumeration(v string)`

SetEnumeration sets Enumeration field to given value.

### HasEnumeration

`func (o *InvoiceDto) HasEnumeration() bool`

HasEnumeration returns a boolean if a field has been set.

### SetEnumerationNil

`func (o *InvoiceDto) SetEnumerationNil(b bool)`

 SetEnumerationNil sets the value for Enumeration to be an explicit nil

### UnsetEnumeration
`func (o *InvoiceDto) UnsetEnumeration()`

UnsetEnumeration ensures that no value is present for Enumeration, not even an explicit nil
### GetPaymentModeId

`func (o *InvoiceDto) GetPaymentModeId() string`

GetPaymentModeId returns the PaymentModeId field if non-nil, zero value otherwise.

### GetPaymentModeIdOk

`func (o *InvoiceDto) GetPaymentModeIdOk() (*string, bool)`

GetPaymentModeIdOk returns a tuple with the PaymentModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeId

`func (o *InvoiceDto) SetPaymentModeId(v string)`

SetPaymentModeId sets PaymentModeId field to given value.

### HasPaymentModeId

`func (o *InvoiceDto) HasPaymentModeId() bool`

HasPaymentModeId returns a boolean if a field has been set.

### SetPaymentModeIdNil

`func (o *InvoiceDto) SetPaymentModeIdNil(b bool)`

 SetPaymentModeIdNil sets the value for PaymentModeId to be an explicit nil

### UnsetPaymentModeId
`func (o *InvoiceDto) UnsetPaymentModeId()`

UnsetPaymentModeId ensures that no value is present for PaymentModeId, not even an explicit nil
### GetEnumerationRangeId

`func (o *InvoiceDto) GetEnumerationRangeId() string`

GetEnumerationRangeId returns the EnumerationRangeId field if non-nil, zero value otherwise.

### GetEnumerationRangeIdOk

`func (o *InvoiceDto) GetEnumerationRangeIdOk() (*string, bool)`

GetEnumerationRangeIdOk returns a tuple with the EnumerationRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumerationRangeId

`func (o *InvoiceDto) SetEnumerationRangeId(v string)`

SetEnumerationRangeId sets EnumerationRangeId field to given value.

### HasEnumerationRangeId

`func (o *InvoiceDto) HasEnumerationRangeId() bool`

HasEnumerationRangeId returns a boolean if a field has been set.

### SetEnumerationRangeIdNil

`func (o *InvoiceDto) SetEnumerationRangeIdNil(b bool)`

 SetEnumerationRangeIdNil sets the value for EnumerationRangeId to be an explicit nil

### UnsetEnumerationRangeId
`func (o *InvoiceDto) UnsetEnumerationRangeId()`

UnsetEnumerationRangeId ensures that no value is present for EnumerationRangeId, not even an explicit nil
### GetEmisorBillingProfileId

`func (o *InvoiceDto) GetEmisorBillingProfileId() string`

GetEmisorBillingProfileId returns the EmisorBillingProfileId field if non-nil, zero value otherwise.

### GetEmisorBillingProfileIdOk

`func (o *InvoiceDto) GetEmisorBillingProfileIdOk() (*string, bool)`

GetEmisorBillingProfileIdOk returns a tuple with the EmisorBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorBillingProfileId

`func (o *InvoiceDto) SetEmisorBillingProfileId(v string)`

SetEmisorBillingProfileId sets EmisorBillingProfileId field to given value.

### HasEmisorBillingProfileId

`func (o *InvoiceDto) HasEmisorBillingProfileId() bool`

HasEmisorBillingProfileId returns a boolean if a field has been set.

### SetEmisorBillingProfileIdNil

`func (o *InvoiceDto) SetEmisorBillingProfileIdNil(b bool)`

 SetEmisorBillingProfileIdNil sets the value for EmisorBillingProfileId to be an explicit nil

### UnsetEmisorBillingProfileId
`func (o *InvoiceDto) UnsetEmisorBillingProfileId()`

UnsetEmisorBillingProfileId ensures that no value is present for EmisorBillingProfileId, not even an explicit nil
### GetReceiverBillingProfileId

`func (o *InvoiceDto) GetReceiverBillingProfileId() string`

GetReceiverBillingProfileId returns the ReceiverBillingProfileId field if non-nil, zero value otherwise.

### GetReceiverBillingProfileIdOk

`func (o *InvoiceDto) GetReceiverBillingProfileIdOk() (*string, bool)`

GetReceiverBillingProfileIdOk returns a tuple with the ReceiverBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBillingProfileId

`func (o *InvoiceDto) SetReceiverBillingProfileId(v string)`

SetReceiverBillingProfileId sets ReceiverBillingProfileId field to given value.

### HasReceiverBillingProfileId

`func (o *InvoiceDto) HasReceiverBillingProfileId() bool`

HasReceiverBillingProfileId returns a boolean if a field has been set.

### SetReceiverBillingProfileIdNil

`func (o *InvoiceDto) SetReceiverBillingProfileIdNil(b bool)`

 SetReceiverBillingProfileIdNil sets the value for ReceiverBillingProfileId to be an explicit nil

### UnsetReceiverBillingProfileId
`func (o *InvoiceDto) UnsetReceiverBillingProfileId()`

UnsetReceiverBillingProfileId ensures that no value is present for ReceiverBillingProfileId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *InvoiceDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *InvoiceDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *InvoiceDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *InvoiceDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *InvoiceDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *InvoiceDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *InvoiceDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *InvoiceDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *InvoiceDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *InvoiceDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *InvoiceDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *InvoiceDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetPaymentDue

`func (o *InvoiceDto) GetPaymentDue() time.Time`

GetPaymentDue returns the PaymentDue field if non-nil, zero value otherwise.

### GetPaymentDueOk

`func (o *InvoiceDto) GetPaymentDueOk() (*time.Time, bool)`

GetPaymentDueOk returns a tuple with the PaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDue

`func (o *InvoiceDto) SetPaymentDue(v time.Time)`

SetPaymentDue sets PaymentDue field to given value.

### HasPaymentDue

`func (o *InvoiceDto) HasPaymentDue() bool`

HasPaymentDue returns a boolean if a field has been set.

### SetPaymentDueNil

`func (o *InvoiceDto) SetPaymentDueNil(b bool)`

 SetPaymentDueNil sets the value for PaymentDue to be an explicit nil

### UnsetPaymentDue
`func (o *InvoiceDto) UnsetPaymentDue()`

UnsetPaymentDue ensures that no value is present for PaymentDue, not even an explicit nil
### GetInvoiceType

`func (o *InvoiceDto) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *InvoiceDto) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *InvoiceDto) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *InvoiceDto) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetDocumentType

`func (o *InvoiceDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *InvoiceDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *InvoiceDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *InvoiceDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceDto) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceDto) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceDto) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceDto) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


