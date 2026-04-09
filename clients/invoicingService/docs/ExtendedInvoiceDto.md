# ExtendedInvoiceDto

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
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**User** | Pointer to [**SimpleUserDto**](SimpleUserDto.md) |  | [optional] 
**ReceiverTenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**Individual** | Pointer to [**SimpleContactDto**](SimpleContactDto.md) |  | [optional] 
**Organization** | Pointer to [**SimpleContactDto**](SimpleContactDto.md) |  | [optional] 
**Enrollment** | Pointer to [**SimpleTenantEnrollmentDto**](SimpleTenantEnrollmentDto.md) |  | [optional] 
**InvoiceLines** | Pointer to [**[]InvoiceLineDto**](InvoiceLineDto.md) |  | [optional] 

## Methods

### NewExtendedInvoiceDto

`func NewExtendedInvoiceDto() *ExtendedInvoiceDto`

NewExtendedInvoiceDto instantiates a new ExtendedInvoiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedInvoiceDtoWithDefaults

`func NewExtendedInvoiceDtoWithDefaults() *ExtendedInvoiceDto`

NewExtendedInvoiceDtoWithDefaults instantiates a new ExtendedInvoiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedInvoiceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedInvoiceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedInvoiceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedInvoiceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedInvoiceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedInvoiceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedInvoiceDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedInvoiceDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedInvoiceDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedInvoiceDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedInvoiceDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedInvoiceDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *ExtendedInvoiceDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *ExtendedInvoiceDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *ExtendedInvoiceDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *ExtendedInvoiceDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetType

`func (o *ExtendedInvoiceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedInvoiceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedInvoiceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExtendedInvoiceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExtendedInvoiceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExtendedInvoiceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ExtendedInvoiceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExtendedInvoiceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExtendedInvoiceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExtendedInvoiceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ExtendedInvoiceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExtendedInvoiceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *ExtendedInvoiceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExtendedInvoiceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExtendedInvoiceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExtendedInvoiceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExtendedInvoiceDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExtendedInvoiceDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *ExtendedInvoiceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedInvoiceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedInvoiceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedInvoiceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedInvoiceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedInvoiceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *ExtendedInvoiceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedInvoiceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedInvoiceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedInvoiceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExtendedInvoiceDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExtendedInvoiceDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceListId

`func (o *ExtendedInvoiceDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ExtendedInvoiceDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ExtendedInvoiceDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ExtendedInvoiceDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ExtendedInvoiceDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ExtendedInvoiceDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetEnrollmentId

`func (o *ExtendedInvoiceDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ExtendedInvoiceDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ExtendedInvoiceDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ExtendedInvoiceDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ExtendedInvoiceDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ExtendedInvoiceDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *ExtendedInvoiceDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *ExtendedInvoiceDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *ExtendedInvoiceDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *ExtendedInvoiceDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *ExtendedInvoiceDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *ExtendedInvoiceDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *ExtendedInvoiceDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ExtendedInvoiceDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ExtendedInvoiceDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ExtendedInvoiceDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ExtendedInvoiceDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ExtendedInvoiceDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *ExtendedInvoiceDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *ExtendedInvoiceDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *ExtendedInvoiceDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *ExtendedInvoiceDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *ExtendedInvoiceDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *ExtendedInvoiceDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *ExtendedInvoiceDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExtendedInvoiceDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExtendedInvoiceDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ExtendedInvoiceDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ExtendedInvoiceDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ExtendedInvoiceDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ExtendedInvoiceDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExtendedInvoiceDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExtendedInvoiceDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ExtendedInvoiceDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ExtendedInvoiceDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ExtendedInvoiceDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *ExtendedInvoiceDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ExtendedInvoiceDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ExtendedInvoiceDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ExtendedInvoiceDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *ExtendedInvoiceDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *ExtendedInvoiceDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *ExtendedInvoiceDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *ExtendedInvoiceDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *ExtendedInvoiceDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *ExtendedInvoiceDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *ExtendedInvoiceDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *ExtendedInvoiceDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *ExtendedInvoiceDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ExtendedInvoiceDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ExtendedInvoiceDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ExtendedInvoiceDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *ExtendedInvoiceDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *ExtendedInvoiceDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *ExtendedInvoiceDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ExtendedInvoiceDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ExtendedInvoiceDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ExtendedInvoiceDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *ExtendedInvoiceDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *ExtendedInvoiceDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *ExtendedInvoiceDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ExtendedInvoiceDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ExtendedInvoiceDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ExtendedInvoiceDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ExtendedInvoiceDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ExtendedInvoiceDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *ExtendedInvoiceDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ExtendedInvoiceDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ExtendedInvoiceDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ExtendedInvoiceDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ExtendedInvoiceDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ExtendedInvoiceDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *ExtendedInvoiceDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ExtendedInvoiceDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ExtendedInvoiceDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ExtendedInvoiceDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ExtendedInvoiceDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ExtendedInvoiceDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *ExtendedInvoiceDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ExtendedInvoiceDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ExtendedInvoiceDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ExtendedInvoiceDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ExtendedInvoiceDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ExtendedInvoiceDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCustomerNotes

`func (o *ExtendedInvoiceDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *ExtendedInvoiceDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *ExtendedInvoiceDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *ExtendedInvoiceDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *ExtendedInvoiceDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *ExtendedInvoiceDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetTaxCalculationMethod

`func (o *ExtendedInvoiceDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *ExtendedInvoiceDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *ExtendedInvoiceDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *ExtendedInvoiceDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *ExtendedInvoiceDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *ExtendedInvoiceDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *ExtendedInvoiceDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *ExtendedInvoiceDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRate

`func (o *ExtendedInvoiceDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *ExtendedInvoiceDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *ExtendedInvoiceDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *ExtendedInvoiceDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ExtendedInvoiceDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExtendedInvoiceDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExtendedInvoiceDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExtendedInvoiceDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *ExtendedInvoiceDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *ExtendedInvoiceDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *ExtendedInvoiceDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *ExtendedInvoiceDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *ExtendedInvoiceDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *ExtendedInvoiceDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *ExtendedInvoiceDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *ExtendedInvoiceDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *ExtendedInvoiceDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ExtendedInvoiceDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *ExtendedInvoiceDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *ExtendedInvoiceDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *ExtendedInvoiceDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *ExtendedInvoiceDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *ExtendedInvoiceDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *ExtendedInvoiceDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *ExtendedInvoiceDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *ExtendedInvoiceDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *ExtendedInvoiceDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *ExtendedInvoiceDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *ExtendedInvoiceDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *ExtendedInvoiceDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *ExtendedInvoiceDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *ExtendedInvoiceDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *ExtendedInvoiceDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *ExtendedInvoiceDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *ExtendedInvoiceDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *ExtendedInvoiceDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *ExtendedInvoiceDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *ExtendedInvoiceDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *ExtendedInvoiceDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *ExtendedInvoiceDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *ExtendedInvoiceDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *ExtendedInvoiceDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *ExtendedInvoiceDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *ExtendedInvoiceDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *ExtendedInvoiceDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *ExtendedInvoiceDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *ExtendedInvoiceDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *ExtendedInvoiceDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotal

`func (o *ExtendedInvoiceDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExtendedInvoiceDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExtendedInvoiceDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExtendedInvoiceDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *ExtendedInvoiceDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *ExtendedInvoiceDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *ExtendedInvoiceDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *ExtendedInvoiceDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *ExtendedInvoiceDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *ExtendedInvoiceDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetTotalDetailInUsd

`func (o *ExtendedInvoiceDto) GetTotalDetailInUsd() float64`

GetTotalDetailInUsd returns the TotalDetailInUsd field if non-nil, zero value otherwise.

### GetTotalDetailInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalDetailInUsdOk() (*float64, bool)`

GetTotalDetailInUsdOk returns a tuple with the TotalDetailInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailInUsd

`func (o *ExtendedInvoiceDto) SetTotalDetailInUsd(v float64)`

SetTotalDetailInUsd sets TotalDetailInUsd field to given value.

### HasTotalDetailInUsd

`func (o *ExtendedInvoiceDto) HasTotalDetailInUsd() bool`

HasTotalDetailInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *ExtendedInvoiceDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *ExtendedInvoiceDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *ExtendedInvoiceDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *ExtendedInvoiceDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *ExtendedInvoiceDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *ExtendedInvoiceDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *ExtendedInvoiceDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *ExtendedInvoiceDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *ExtendedInvoiceDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *ExtendedInvoiceDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *ExtendedInvoiceDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *ExtendedInvoiceDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *ExtendedInvoiceDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *ExtendedInvoiceDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *ExtendedInvoiceDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *ExtendedInvoiceDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *ExtendedInvoiceDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *ExtendedInvoiceDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *ExtendedInvoiceDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *ExtendedInvoiceDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *ExtendedInvoiceDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *ExtendedInvoiceDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *ExtendedInvoiceDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *ExtendedInvoiceDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *ExtendedInvoiceDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *ExtendedInvoiceDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *ExtendedInvoiceDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *ExtendedInvoiceDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *ExtendedInvoiceDto) GetTotalInUsd() float64`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *ExtendedInvoiceDto) GetTotalInUsdOk() (*float64, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *ExtendedInvoiceDto) SetTotalInUsd(v float64)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *ExtendedInvoiceDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetPaid

`func (o *ExtendedInvoiceDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *ExtendedInvoiceDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *ExtendedInvoiceDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *ExtendedInvoiceDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetNumber

`func (o *ExtendedInvoiceDto) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ExtendedInvoiceDto) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ExtendedInvoiceDto) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ExtendedInvoiceDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNotes

`func (o *ExtendedInvoiceDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ExtendedInvoiceDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ExtendedInvoiceDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ExtendedInvoiceDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ExtendedInvoiceDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ExtendedInvoiceDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOrderId

`func (o *ExtendedInvoiceDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ExtendedInvoiceDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ExtendedInvoiceDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ExtendedInvoiceDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ExtendedInvoiceDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ExtendedInvoiceDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetEnumeration

`func (o *ExtendedInvoiceDto) GetEnumeration() string`

GetEnumeration returns the Enumeration field if non-nil, zero value otherwise.

### GetEnumerationOk

`func (o *ExtendedInvoiceDto) GetEnumerationOk() (*string, bool)`

GetEnumerationOk returns a tuple with the Enumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumeration

`func (o *ExtendedInvoiceDto) SetEnumeration(v string)`

SetEnumeration sets Enumeration field to given value.

### HasEnumeration

`func (o *ExtendedInvoiceDto) HasEnumeration() bool`

HasEnumeration returns a boolean if a field has been set.

### SetEnumerationNil

`func (o *ExtendedInvoiceDto) SetEnumerationNil(b bool)`

 SetEnumerationNil sets the value for Enumeration to be an explicit nil

### UnsetEnumeration
`func (o *ExtendedInvoiceDto) UnsetEnumeration()`

UnsetEnumeration ensures that no value is present for Enumeration, not even an explicit nil
### GetPaymentModeId

`func (o *ExtendedInvoiceDto) GetPaymentModeId() string`

GetPaymentModeId returns the PaymentModeId field if non-nil, zero value otherwise.

### GetPaymentModeIdOk

`func (o *ExtendedInvoiceDto) GetPaymentModeIdOk() (*string, bool)`

GetPaymentModeIdOk returns a tuple with the PaymentModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeId

`func (o *ExtendedInvoiceDto) SetPaymentModeId(v string)`

SetPaymentModeId sets PaymentModeId field to given value.

### HasPaymentModeId

`func (o *ExtendedInvoiceDto) HasPaymentModeId() bool`

HasPaymentModeId returns a boolean if a field has been set.

### SetPaymentModeIdNil

`func (o *ExtendedInvoiceDto) SetPaymentModeIdNil(b bool)`

 SetPaymentModeIdNil sets the value for PaymentModeId to be an explicit nil

### UnsetPaymentModeId
`func (o *ExtendedInvoiceDto) UnsetPaymentModeId()`

UnsetPaymentModeId ensures that no value is present for PaymentModeId, not even an explicit nil
### GetEnumerationRangeId

`func (o *ExtendedInvoiceDto) GetEnumerationRangeId() string`

GetEnumerationRangeId returns the EnumerationRangeId field if non-nil, zero value otherwise.

### GetEnumerationRangeIdOk

`func (o *ExtendedInvoiceDto) GetEnumerationRangeIdOk() (*string, bool)`

GetEnumerationRangeIdOk returns a tuple with the EnumerationRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumerationRangeId

`func (o *ExtendedInvoiceDto) SetEnumerationRangeId(v string)`

SetEnumerationRangeId sets EnumerationRangeId field to given value.

### HasEnumerationRangeId

`func (o *ExtendedInvoiceDto) HasEnumerationRangeId() bool`

HasEnumerationRangeId returns a boolean if a field has been set.

### SetEnumerationRangeIdNil

`func (o *ExtendedInvoiceDto) SetEnumerationRangeIdNil(b bool)`

 SetEnumerationRangeIdNil sets the value for EnumerationRangeId to be an explicit nil

### UnsetEnumerationRangeId
`func (o *ExtendedInvoiceDto) UnsetEnumerationRangeId()`

UnsetEnumerationRangeId ensures that no value is present for EnumerationRangeId, not even an explicit nil
### GetEmisorBillingProfileId

`func (o *ExtendedInvoiceDto) GetEmisorBillingProfileId() string`

GetEmisorBillingProfileId returns the EmisorBillingProfileId field if non-nil, zero value otherwise.

### GetEmisorBillingProfileIdOk

`func (o *ExtendedInvoiceDto) GetEmisorBillingProfileIdOk() (*string, bool)`

GetEmisorBillingProfileIdOk returns a tuple with the EmisorBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorBillingProfileId

`func (o *ExtendedInvoiceDto) SetEmisorBillingProfileId(v string)`

SetEmisorBillingProfileId sets EmisorBillingProfileId field to given value.

### HasEmisorBillingProfileId

`func (o *ExtendedInvoiceDto) HasEmisorBillingProfileId() bool`

HasEmisorBillingProfileId returns a boolean if a field has been set.

### SetEmisorBillingProfileIdNil

`func (o *ExtendedInvoiceDto) SetEmisorBillingProfileIdNil(b bool)`

 SetEmisorBillingProfileIdNil sets the value for EmisorBillingProfileId to be an explicit nil

### UnsetEmisorBillingProfileId
`func (o *ExtendedInvoiceDto) UnsetEmisorBillingProfileId()`

UnsetEmisorBillingProfileId ensures that no value is present for EmisorBillingProfileId, not even an explicit nil
### GetReceiverBillingProfileId

`func (o *ExtendedInvoiceDto) GetReceiverBillingProfileId() string`

GetReceiverBillingProfileId returns the ReceiverBillingProfileId field if non-nil, zero value otherwise.

### GetReceiverBillingProfileIdOk

`func (o *ExtendedInvoiceDto) GetReceiverBillingProfileIdOk() (*string, bool)`

GetReceiverBillingProfileIdOk returns a tuple with the ReceiverBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBillingProfileId

`func (o *ExtendedInvoiceDto) SetReceiverBillingProfileId(v string)`

SetReceiverBillingProfileId sets ReceiverBillingProfileId field to given value.

### HasReceiverBillingProfileId

`func (o *ExtendedInvoiceDto) HasReceiverBillingProfileId() bool`

HasReceiverBillingProfileId returns a boolean if a field has been set.

### SetReceiverBillingProfileIdNil

`func (o *ExtendedInvoiceDto) SetReceiverBillingProfileIdNil(b bool)`

 SetReceiverBillingProfileIdNil sets the value for ReceiverBillingProfileId to be an explicit nil

### UnsetReceiverBillingProfileId
`func (o *ExtendedInvoiceDto) UnsetReceiverBillingProfileId()`

UnsetReceiverBillingProfileId ensures that no value is present for ReceiverBillingProfileId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *ExtendedInvoiceDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *ExtendedInvoiceDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *ExtendedInvoiceDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *ExtendedInvoiceDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *ExtendedInvoiceDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *ExtendedInvoiceDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *ExtendedInvoiceDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *ExtendedInvoiceDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *ExtendedInvoiceDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *ExtendedInvoiceDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *ExtendedInvoiceDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *ExtendedInvoiceDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetPaymentDue

`func (o *ExtendedInvoiceDto) GetPaymentDue() time.Time`

GetPaymentDue returns the PaymentDue field if non-nil, zero value otherwise.

### GetPaymentDueOk

`func (o *ExtendedInvoiceDto) GetPaymentDueOk() (*time.Time, bool)`

GetPaymentDueOk returns a tuple with the PaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDue

`func (o *ExtendedInvoiceDto) SetPaymentDue(v time.Time)`

SetPaymentDue sets PaymentDue field to given value.

### HasPaymentDue

`func (o *ExtendedInvoiceDto) HasPaymentDue() bool`

HasPaymentDue returns a boolean if a field has been set.

### SetPaymentDueNil

`func (o *ExtendedInvoiceDto) SetPaymentDueNil(b bool)`

 SetPaymentDueNil sets the value for PaymentDue to be an explicit nil

### UnsetPaymentDue
`func (o *ExtendedInvoiceDto) UnsetPaymentDue()`

UnsetPaymentDue ensures that no value is present for PaymentDue, not even an explicit nil
### GetInvoiceType

`func (o *ExtendedInvoiceDto) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *ExtendedInvoiceDto) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *ExtendedInvoiceDto) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *ExtendedInvoiceDto) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetDocumentType

`func (o *ExtendedInvoiceDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ExtendedInvoiceDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ExtendedInvoiceDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *ExtendedInvoiceDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *ExtendedInvoiceDto) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *ExtendedInvoiceDto) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *ExtendedInvoiceDto) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *ExtendedInvoiceDto) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetTenant

`func (o *ExtendedInvoiceDto) GetTenant() TenantDto`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExtendedInvoiceDto) GetTenantOk() (*TenantDto, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExtendedInvoiceDto) SetTenant(v TenantDto)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ExtendedInvoiceDto) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *ExtendedInvoiceDto) GetUser() SimpleUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExtendedInvoiceDto) GetUserOk() (*SimpleUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExtendedInvoiceDto) SetUser(v SimpleUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExtendedInvoiceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReceiverTenant

`func (o *ExtendedInvoiceDto) GetReceiverTenant() TenantDto`

GetReceiverTenant returns the ReceiverTenant field if non-nil, zero value otherwise.

### GetReceiverTenantOk

`func (o *ExtendedInvoiceDto) GetReceiverTenantOk() (*TenantDto, bool)`

GetReceiverTenantOk returns a tuple with the ReceiverTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenant

`func (o *ExtendedInvoiceDto) SetReceiverTenant(v TenantDto)`

SetReceiverTenant sets ReceiverTenant field to given value.

### HasReceiverTenant

`func (o *ExtendedInvoiceDto) HasReceiverTenant() bool`

HasReceiverTenant returns a boolean if a field has been set.

### GetIndividual

`func (o *ExtendedInvoiceDto) GetIndividual() SimpleContactDto`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *ExtendedInvoiceDto) GetIndividualOk() (*SimpleContactDto, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *ExtendedInvoiceDto) SetIndividual(v SimpleContactDto)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *ExtendedInvoiceDto) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *ExtendedInvoiceDto) GetOrganization() SimpleContactDto`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ExtendedInvoiceDto) GetOrganizationOk() (*SimpleContactDto, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ExtendedInvoiceDto) SetOrganization(v SimpleContactDto)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ExtendedInvoiceDto) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEnrollment

`func (o *ExtendedInvoiceDto) GetEnrollment() SimpleTenantEnrollmentDto`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ExtendedInvoiceDto) GetEnrollmentOk() (*SimpleTenantEnrollmentDto, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ExtendedInvoiceDto) SetEnrollment(v SimpleTenantEnrollmentDto)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *ExtendedInvoiceDto) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetInvoiceLines

`func (o *ExtendedInvoiceDto) GetInvoiceLines() []InvoiceLineDto`

GetInvoiceLines returns the InvoiceLines field if non-nil, zero value otherwise.

### GetInvoiceLinesOk

`func (o *ExtendedInvoiceDto) GetInvoiceLinesOk() (*[]InvoiceLineDto, bool)`

GetInvoiceLinesOk returns a tuple with the InvoiceLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLines

`func (o *ExtendedInvoiceDto) SetInvoiceLines(v []InvoiceLineDto)`

SetInvoiceLines sets InvoiceLines field to given value.

### HasInvoiceLines

`func (o *ExtendedInvoiceDto) HasInvoiceLines() bool`

HasInvoiceLines returns a boolean if a field has been set.

### SetInvoiceLinesNil

`func (o *ExtendedInvoiceDto) SetInvoiceLinesNil(b bool)`

 SetInvoiceLinesNil sets the value for InvoiceLines to be an explicit nil

### UnsetInvoiceLines
`func (o *ExtendedInvoiceDto) UnsetInvoiceLines()`

UnsetInvoiceLines ensures that no value is present for InvoiceLines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


