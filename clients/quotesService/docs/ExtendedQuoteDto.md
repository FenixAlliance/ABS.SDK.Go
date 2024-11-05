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
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
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
**ForexRate** | Pointer to **float64** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**TotalTaxBase** | Pointer to **float64** |  | [optional] 
**TotalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalProfitInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalDetailAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWithholdingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**TotalInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxBaseAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDiscountsAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalSurchargesAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalDiscountsAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalSurchargesAmountInUsd** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxBaseAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**EffectiveTo** | Pointer to **NullableTime** |  | [optional] 
**EffectiveFrom** | Pointer to **NullableTime** |  | [optional] 
**QuoteStatus** | Pointer to **int32** |  | [optional] 
**FreightTerms** | Pointer to **int32** |  | [optional] 
**CostCalculationMethod** | Pointer to **int32** |  | [optional] 
**CustomDiscountsAmount** | Pointer to **float64** |  | [optional] 
**User** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**Individual** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 
**Organization** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 
**ReceiverTenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**Enrollment** | Pointer to [**TenantEnrolmentDto**](TenantEnrolmentDto.md) |  | [optional] 

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

### GetTotalAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

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

### GetTotalDetailAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalDetailAmountInUsd() float64`

GetTotalDetailAmountInUsd returns the TotalDetailAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDetailAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalDetailAmountInUsdOk() (*float64, bool)`

GetTotalDetailAmountInUsdOk returns a tuple with the TotalDetailAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalDetailAmountInUsd(v float64)`

SetTotalDetailAmountInUsd sets TotalDetailAmountInUsd field to given value.

### HasTotalDetailAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalDetailAmountInUsd() bool`

HasTotalDetailAmountInUsd returns a boolean if a field has been set.

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

### GetTotalWithholdingTaxesInUsd

`func (o *ExtendedQuoteDto) GetTotalWithholdingTaxesInUsd() float64`

GetTotalWithholdingTaxesInUsd returns the TotalWithholdingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithholdingTaxesInUsdOk

`func (o *ExtendedQuoteDto) GetTotalWithholdingTaxesInUsdOk() (*float64, bool)`

GetTotalWithholdingTaxesInUsdOk returns a tuple with the TotalWithholdingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithholdingTaxesInUsd

`func (o *ExtendedQuoteDto) SetTotalWithholdingTaxesInUsd(v float64)`

SetTotalWithholdingTaxesInUsd sets TotalWithholdingTaxesInUsd field to given value.

### HasTotalWithholdingTaxesInUsd

`func (o *ExtendedQuoteDto) HasTotalWithholdingTaxesInUsd() bool`

HasTotalWithholdingTaxesInUsd returns a boolean if a field has been set.

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

### GetCurrency

`func (o *ExtendedQuoteDto) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExtendedQuoteDto) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExtendedQuoteDto) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExtendedQuoteDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *ExtendedQuoteDto) GetTotalInUsd() Money`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *ExtendedQuoteDto) GetTotalInUsdOk() (*Money, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *ExtendedQuoteDto) SetTotalInUsd(v Money)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *ExtendedQuoteDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetTotalTaxAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalTaxAmountInUsd() Money`

GetTotalTaxAmountInUsd returns the TotalTaxAmountInUsd field if non-nil, zero value otherwise.

### GetTotalTaxAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalTaxAmountInUsdOk() (*Money, bool)`

GetTotalTaxAmountInUsdOk returns a tuple with the TotalTaxAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalTaxAmountInUsd(v Money)`

SetTotalTaxAmountInUsd sets TotalTaxAmountInUsd field to given value.

### HasTotalTaxAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalTaxAmountInUsd() bool`

HasTotalTaxAmountInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalTaxBaseAmountInUsd() Money`

GetTotalTaxBaseAmountInUsd returns the TotalTaxBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalTaxBaseAmountInUsdOk() (*Money, bool)`

GetTotalTaxBaseAmountInUsdOk returns a tuple with the TotalTaxBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalTaxBaseAmountInUsd(v Money)`

SetTotalTaxBaseAmountInUsd sets TotalTaxBaseAmountInUsd field to given value.

### HasTotalTaxBaseAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalTaxBaseAmountInUsd() bool`

HasTotalTaxBaseAmountInUsd returns a boolean if a field has been set.

### GetTotalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalDiscountsAmountInUsd() Money`

GetTotalDiscountsAmountInUsd returns the TotalDiscountsAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalDiscountsAmountInUsdOk() (*Money, bool)`

GetTotalDiscountsAmountInUsdOk returns a tuple with the TotalDiscountsAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalDiscountsAmountInUsd(v Money)`

SetTotalDiscountsAmountInUsd sets TotalDiscountsAmountInUsd field to given value.

### HasTotalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalDiscountsAmountInUsd() bool`

HasTotalDiscountsAmountInUsd returns a boolean if a field has been set.

### GetTotalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalSurchargesAmountInUsd() Money`

GetTotalSurchargesAmountInUsd returns the TotalSurchargesAmountInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalSurchargesAmountInUsdOk() (*Money, bool)`

GetTotalSurchargesAmountInUsdOk returns a tuple with the TotalSurchargesAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalSurchargesAmountInUsd(v Money)`

SetTotalSurchargesAmountInUsd sets TotalSurchargesAmountInUsd field to given value.

### HasTotalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalSurchargesAmountInUsd() bool`

HasTotalSurchargesAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsAmountInUsd() Money`

GetTotalGlobalDiscountsAmountInUsd returns the TotalGlobalDiscountsAmountInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsAmountInUsdOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountInUsdOk returns a tuple with the TotalGlobalDiscountsAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscountsAmountInUsd(v Money)`

SetTotalGlobalDiscountsAmountInUsd sets TotalGlobalDiscountsAmountInUsd field to given value.

### HasTotalGlobalDiscountsAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalGlobalDiscountsAmountInUsd() bool`

HasTotalGlobalDiscountsAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesAmountInUsd() Money`

GetTotalGlobalSurchargesAmountInUsd returns the TotalGlobalSurchargesAmountInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountInUsdOk

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesAmountInUsdOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountInUsdOk returns a tuple with the TotalGlobalSurchargesAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) SetTotalGlobalSurchargesAmountInUsd(v Money)`

SetTotalGlobalSurchargesAmountInUsd sets TotalGlobalSurchargesAmountInUsd field to given value.

### HasTotalGlobalSurchargesAmountInUsd

`func (o *ExtendedQuoteDto) HasTotalGlobalSurchargesAmountInUsd() bool`

HasTotalGlobalSurchargesAmountInUsd returns a boolean if a field has been set.

### GetTotalAmount

`func (o *ExtendedQuoteDto) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ExtendedQuoteDto) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ExtendedQuoteDto) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ExtendedQuoteDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalTaxAmount

`func (o *ExtendedQuoteDto) GetTotalTaxAmount() Money`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *ExtendedQuoteDto) GetTotalTaxAmountOk() (*Money, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *ExtendedQuoteDto) SetTotalTaxAmount(v Money)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *ExtendedQuoteDto) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### GetTotalTaxBaseAmount

`func (o *ExtendedQuoteDto) GetTotalTaxBaseAmount() Money`

GetTotalTaxBaseAmount returns the TotalTaxBaseAmount field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountOk

`func (o *ExtendedQuoteDto) GetTotalTaxBaseAmountOk() (*Money, bool)`

GetTotalTaxBaseAmountOk returns a tuple with the TotalTaxBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmount

`func (o *ExtendedQuoteDto) SetTotalTaxBaseAmount(v Money)`

SetTotalTaxBaseAmount sets TotalTaxBaseAmount field to given value.

### HasTotalTaxBaseAmount

`func (o *ExtendedQuoteDto) HasTotalTaxBaseAmount() bool`

HasTotalTaxBaseAmount returns a boolean if a field has been set.

### GetTotalDiscountsAmount

`func (o *ExtendedQuoteDto) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *ExtendedQuoteDto) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *ExtendedQuoteDto) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *ExtendedQuoteDto) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurchargesAmount

`func (o *ExtendedQuoteDto) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *ExtendedQuoteDto) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *ExtendedQuoteDto) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *ExtendedQuoteDto) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalGlobalDiscountsAmount

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsAmount() Money`

GetTotalGlobalDiscountsAmount returns the TotalGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountOk

`func (o *ExtendedQuoteDto) GetTotalGlobalDiscountsAmountOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountOk returns a tuple with the TotalGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmount

`func (o *ExtendedQuoteDto) SetTotalGlobalDiscountsAmount(v Money)`

SetTotalGlobalDiscountsAmount sets TotalGlobalDiscountsAmount field to given value.

### HasTotalGlobalDiscountsAmount

`func (o *ExtendedQuoteDto) HasTotalGlobalDiscountsAmount() bool`

HasTotalGlobalDiscountsAmount returns a boolean if a field has been set.

### GetTotalGlobalSurchargesAmount

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesAmount() Money`

GetTotalGlobalSurchargesAmount returns the TotalGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountOk

`func (o *ExtendedQuoteDto) GetTotalGlobalSurchargesAmountOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountOk returns a tuple with the TotalGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmount

`func (o *ExtendedQuoteDto) SetTotalGlobalSurchargesAmount(v Money)`

SetTotalGlobalSurchargesAmount sets TotalGlobalSurchargesAmount field to given value.

### HasTotalGlobalSurchargesAmount

`func (o *ExtendedQuoteDto) HasTotalGlobalSurchargesAmount() bool`

HasTotalGlobalSurchargesAmount returns a boolean if a field has been set.

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

`func (o *ExtendedQuoteDto) GetQuoteStatus() int32`

GetQuoteStatus returns the QuoteStatus field if non-nil, zero value otherwise.

### GetQuoteStatusOk

`func (o *ExtendedQuoteDto) GetQuoteStatusOk() (*int32, bool)`

GetQuoteStatusOk returns a tuple with the QuoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStatus

`func (o *ExtendedQuoteDto) SetQuoteStatus(v int32)`

SetQuoteStatus sets QuoteStatus field to given value.

### HasQuoteStatus

`func (o *ExtendedQuoteDto) HasQuoteStatus() bool`

HasQuoteStatus returns a boolean if a field has been set.

### GetFreightTerms

`func (o *ExtendedQuoteDto) GetFreightTerms() int32`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *ExtendedQuoteDto) GetFreightTermsOk() (*int32, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *ExtendedQuoteDto) SetFreightTerms(v int32)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *ExtendedQuoteDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *ExtendedQuoteDto) GetCostCalculationMethod() int32`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *ExtendedQuoteDto) GetCostCalculationMethodOk() (*int32, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *ExtendedQuoteDto) SetCostCalculationMethod(v int32)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *ExtendedQuoteDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

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

`func (o *ExtendedQuoteDto) GetEnrollment() TenantEnrolmentDto`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ExtendedQuoteDto) GetEnrollmentOk() (*TenantEnrolmentDto, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ExtendedQuoteDto) SetEnrollment(v TenantEnrolmentDto)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *ExtendedQuoteDto) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


