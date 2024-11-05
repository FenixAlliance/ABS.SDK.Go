# InvoiceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Closed** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**PaymentTermId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
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
**Paid** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**CustomerNotes** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**Enumeration** | Pointer to **NullableString** |  | [optional] 
**PaymentModeId** | Pointer to **NullableString** |  | [optional] 
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
**EnumerationRangeId** | Pointer to **NullableString** |  | [optional] 
**EmisorBillingProfileId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBillingProfileId** | Pointer to **NullableString** |  | [optional] 
**EmisorWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**PaymentDue** | Pointer to **NullableTime** |  | [optional] 
**InvoiceType** | Pointer to **int32** |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**InvoiceStatus** | Pointer to **int32** |  | [optional] 
**ValidFrom** | Pointer to **NullableTime** |  | [optional] 
**ValidTo** | Pointer to **NullableTime** |  | [optional] 
**InvoiceReferences** | Pointer to [**[]InvoiceReferenceDto**](InvoiceReferenceDto.md) |  | [optional] 
**InvoiceItemRecords** | Pointer to [**[]InvoiceItemRecordDto**](InvoiceItemRecordDto.md) |  | [optional] 
**InvoiceAdjustments** | Pointer to [**[]InvoiceAdjustmentDto**](InvoiceAdjustmentDto.md) |  | [optional] 

## Methods

### NewInvoiceCreateDto

`func NewInvoiceCreateDto() *InvoiceCreateDto`

NewInvoiceCreateDto instantiates a new InvoiceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreateDtoWithDefaults

`func NewInvoiceCreateDtoWithDefaults() *InvoiceCreateDto`

NewInvoiceCreateDtoWithDefaults instantiates a new InvoiceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetClosed

`func (o *InvoiceCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *InvoiceCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *InvoiceCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *InvoiceCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetTitle

`func (o *InvoiceCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvoiceCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvoiceCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *InvoiceCreateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InvoiceCreateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InvoiceCreateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *InvoiceCreateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *InvoiceCreateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InvoiceCreateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *InvoiceCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetPriceListId

`func (o *InvoiceCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *InvoiceCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *InvoiceCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *InvoiceCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *InvoiceCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *InvoiceCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDescription

`func (o *InvoiceCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *InvoiceCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *InvoiceCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *InvoiceCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *InvoiceCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *InvoiceCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *InvoiceCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetPaymentTermId

`func (o *InvoiceCreateDto) GetPaymentTermId() string`

GetPaymentTermId returns the PaymentTermId field if non-nil, zero value otherwise.

### GetPaymentTermIdOk

`func (o *InvoiceCreateDto) GetPaymentTermIdOk() (*string, bool)`

GetPaymentTermIdOk returns a tuple with the PaymentTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermId

`func (o *InvoiceCreateDto) SetPaymentTermId(v string)`

SetPaymentTermId sets PaymentTermId field to given value.

### HasPaymentTermId

`func (o *InvoiceCreateDto) HasPaymentTermId() bool`

HasPaymentTermId returns a boolean if a field has been set.

### SetPaymentTermIdNil

`func (o *InvoiceCreateDto) SetPaymentTermIdNil(b bool)`

 SetPaymentTermIdNil sets the value for PaymentTermId to be an explicit nil

### UnsetPaymentTermId
`func (o *InvoiceCreateDto) UnsetPaymentTermId()`

UnsetPaymentTermId ensures that no value is present for PaymentTermId, not even an explicit nil
### GetOrganizationId

`func (o *InvoiceCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InvoiceCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InvoiceCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InvoiceCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *InvoiceCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *InvoiceCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetCurrencyId

`func (o *InvoiceCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InvoiceCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InvoiceCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InvoiceCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *InvoiceCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *InvoiceCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetForexRate

`func (o *InvoiceCreateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *InvoiceCreateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *InvoiceCreateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *InvoiceCreateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetFirstName

`func (o *InvoiceCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InvoiceCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InvoiceCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InvoiceCreateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *InvoiceCreateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *InvoiceCreateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *InvoiceCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InvoiceCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InvoiceCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InvoiceCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *InvoiceCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *InvoiceCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *InvoiceCreateDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceCreateDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceCreateDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvoiceCreateDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InvoiceCreateDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InvoiceCreateDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *InvoiceCreateDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *InvoiceCreateDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *InvoiceCreateDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *InvoiceCreateDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *InvoiceCreateDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *InvoiceCreateDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *InvoiceCreateDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *InvoiceCreateDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *InvoiceCreateDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *InvoiceCreateDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *InvoiceCreateDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *InvoiceCreateDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *InvoiceCreateDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *InvoiceCreateDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *InvoiceCreateDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *InvoiceCreateDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *InvoiceCreateDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *InvoiceCreateDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *InvoiceCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *InvoiceCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *InvoiceCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *InvoiceCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *InvoiceCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *InvoiceCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *InvoiceCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *InvoiceCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *InvoiceCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *InvoiceCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *InvoiceCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *InvoiceCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *InvoiceCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *InvoiceCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *InvoiceCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *InvoiceCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *InvoiceCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *InvoiceCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *InvoiceCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *InvoiceCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *InvoiceCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *InvoiceCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *InvoiceCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *InvoiceCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetPaid

`func (o *InvoiceCreateDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *InvoiceCreateDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *InvoiceCreateDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *InvoiceCreateDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetNumber

`func (o *InvoiceCreateDto) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceCreateDto) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceCreateDto) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceCreateDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InvoiceCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InvoiceCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCustomerNotes

`func (o *InvoiceCreateDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *InvoiceCreateDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *InvoiceCreateDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *InvoiceCreateDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *InvoiceCreateDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *InvoiceCreateDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetOrderId

`func (o *InvoiceCreateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InvoiceCreateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InvoiceCreateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *InvoiceCreateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *InvoiceCreateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *InvoiceCreateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetEnumeration

`func (o *InvoiceCreateDto) GetEnumeration() string`

GetEnumeration returns the Enumeration field if non-nil, zero value otherwise.

### GetEnumerationOk

`func (o *InvoiceCreateDto) GetEnumerationOk() (*string, bool)`

GetEnumerationOk returns a tuple with the Enumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumeration

`func (o *InvoiceCreateDto) SetEnumeration(v string)`

SetEnumeration sets Enumeration field to given value.

### HasEnumeration

`func (o *InvoiceCreateDto) HasEnumeration() bool`

HasEnumeration returns a boolean if a field has been set.

### SetEnumerationNil

`func (o *InvoiceCreateDto) SetEnumerationNil(b bool)`

 SetEnumerationNil sets the value for Enumeration to be an explicit nil

### UnsetEnumeration
`func (o *InvoiceCreateDto) UnsetEnumeration()`

UnsetEnumeration ensures that no value is present for Enumeration, not even an explicit nil
### GetPaymentModeId

`func (o *InvoiceCreateDto) GetPaymentModeId() string`

GetPaymentModeId returns the PaymentModeId field if non-nil, zero value otherwise.

### GetPaymentModeIdOk

`func (o *InvoiceCreateDto) GetPaymentModeIdOk() (*string, bool)`

GetPaymentModeIdOk returns a tuple with the PaymentModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeId

`func (o *InvoiceCreateDto) SetPaymentModeId(v string)`

SetPaymentModeId sets PaymentModeId field to given value.

### HasPaymentModeId

`func (o *InvoiceCreateDto) HasPaymentModeId() bool`

HasPaymentModeId returns a boolean if a field has been set.

### SetPaymentModeIdNil

`func (o *InvoiceCreateDto) SetPaymentModeIdNil(b bool)`

 SetPaymentModeIdNil sets the value for PaymentModeId to be an explicit nil

### UnsetPaymentModeId
`func (o *InvoiceCreateDto) UnsetPaymentModeId()`

UnsetPaymentModeId ensures that no value is present for PaymentModeId, not even an explicit nil
### GetReceiverTenantId

`func (o *InvoiceCreateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *InvoiceCreateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *InvoiceCreateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *InvoiceCreateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *InvoiceCreateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *InvoiceCreateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetEnumerationRangeId

`func (o *InvoiceCreateDto) GetEnumerationRangeId() string`

GetEnumerationRangeId returns the EnumerationRangeId field if non-nil, zero value otherwise.

### GetEnumerationRangeIdOk

`func (o *InvoiceCreateDto) GetEnumerationRangeIdOk() (*string, bool)`

GetEnumerationRangeIdOk returns a tuple with the EnumerationRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumerationRangeId

`func (o *InvoiceCreateDto) SetEnumerationRangeId(v string)`

SetEnumerationRangeId sets EnumerationRangeId field to given value.

### HasEnumerationRangeId

`func (o *InvoiceCreateDto) HasEnumerationRangeId() bool`

HasEnumerationRangeId returns a boolean if a field has been set.

### SetEnumerationRangeIdNil

`func (o *InvoiceCreateDto) SetEnumerationRangeIdNil(b bool)`

 SetEnumerationRangeIdNil sets the value for EnumerationRangeId to be an explicit nil

### UnsetEnumerationRangeId
`func (o *InvoiceCreateDto) UnsetEnumerationRangeId()`

UnsetEnumerationRangeId ensures that no value is present for EnumerationRangeId, not even an explicit nil
### GetEmisorBillingProfileId

`func (o *InvoiceCreateDto) GetEmisorBillingProfileId() string`

GetEmisorBillingProfileId returns the EmisorBillingProfileId field if non-nil, zero value otherwise.

### GetEmisorBillingProfileIdOk

`func (o *InvoiceCreateDto) GetEmisorBillingProfileIdOk() (*string, bool)`

GetEmisorBillingProfileIdOk returns a tuple with the EmisorBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorBillingProfileId

`func (o *InvoiceCreateDto) SetEmisorBillingProfileId(v string)`

SetEmisorBillingProfileId sets EmisorBillingProfileId field to given value.

### HasEmisorBillingProfileId

`func (o *InvoiceCreateDto) HasEmisorBillingProfileId() bool`

HasEmisorBillingProfileId returns a boolean if a field has been set.

### SetEmisorBillingProfileIdNil

`func (o *InvoiceCreateDto) SetEmisorBillingProfileIdNil(b bool)`

 SetEmisorBillingProfileIdNil sets the value for EmisorBillingProfileId to be an explicit nil

### UnsetEmisorBillingProfileId
`func (o *InvoiceCreateDto) UnsetEmisorBillingProfileId()`

UnsetEmisorBillingProfileId ensures that no value is present for EmisorBillingProfileId, not even an explicit nil
### GetReceiverBillingProfileId

`func (o *InvoiceCreateDto) GetReceiverBillingProfileId() string`

GetReceiverBillingProfileId returns the ReceiverBillingProfileId field if non-nil, zero value otherwise.

### GetReceiverBillingProfileIdOk

`func (o *InvoiceCreateDto) GetReceiverBillingProfileIdOk() (*string, bool)`

GetReceiverBillingProfileIdOk returns a tuple with the ReceiverBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBillingProfileId

`func (o *InvoiceCreateDto) SetReceiverBillingProfileId(v string)`

SetReceiverBillingProfileId sets ReceiverBillingProfileId field to given value.

### HasReceiverBillingProfileId

`func (o *InvoiceCreateDto) HasReceiverBillingProfileId() bool`

HasReceiverBillingProfileId returns a boolean if a field has been set.

### SetReceiverBillingProfileIdNil

`func (o *InvoiceCreateDto) SetReceiverBillingProfileIdNil(b bool)`

 SetReceiverBillingProfileIdNil sets the value for ReceiverBillingProfileId to be an explicit nil

### UnsetReceiverBillingProfileId
`func (o *InvoiceCreateDto) UnsetReceiverBillingProfileId()`

UnsetReceiverBillingProfileId ensures that no value is present for ReceiverBillingProfileId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *InvoiceCreateDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *InvoiceCreateDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *InvoiceCreateDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *InvoiceCreateDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *InvoiceCreateDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *InvoiceCreateDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *InvoiceCreateDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *InvoiceCreateDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *InvoiceCreateDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *InvoiceCreateDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *InvoiceCreateDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *InvoiceCreateDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetPaymentDue

`func (o *InvoiceCreateDto) GetPaymentDue() time.Time`

GetPaymentDue returns the PaymentDue field if non-nil, zero value otherwise.

### GetPaymentDueOk

`func (o *InvoiceCreateDto) GetPaymentDueOk() (*time.Time, bool)`

GetPaymentDueOk returns a tuple with the PaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDue

`func (o *InvoiceCreateDto) SetPaymentDue(v time.Time)`

SetPaymentDue sets PaymentDue field to given value.

### HasPaymentDue

`func (o *InvoiceCreateDto) HasPaymentDue() bool`

HasPaymentDue returns a boolean if a field has been set.

### SetPaymentDueNil

`func (o *InvoiceCreateDto) SetPaymentDueNil(b bool)`

 SetPaymentDueNil sets the value for PaymentDue to be an explicit nil

### UnsetPaymentDue
`func (o *InvoiceCreateDto) UnsetPaymentDue()`

UnsetPaymentDue ensures that no value is present for PaymentDue, not even an explicit nil
### GetInvoiceType

`func (o *InvoiceCreateDto) GetInvoiceType() int32`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *InvoiceCreateDto) GetInvoiceTypeOk() (*int32, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *InvoiceCreateDto) SetInvoiceType(v int32)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *InvoiceCreateDto) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetDocumentType

`func (o *InvoiceCreateDto) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *InvoiceCreateDto) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *InvoiceCreateDto) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *InvoiceCreateDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceCreateDto) GetInvoiceStatus() int32`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceCreateDto) GetInvoiceStatusOk() (*int32, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceCreateDto) SetInvoiceStatus(v int32)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceCreateDto) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetValidFrom

`func (o *InvoiceCreateDto) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *InvoiceCreateDto) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *InvoiceCreateDto) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *InvoiceCreateDto) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFromNil

`func (o *InvoiceCreateDto) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *InvoiceCreateDto) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetValidTo

`func (o *InvoiceCreateDto) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *InvoiceCreateDto) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *InvoiceCreateDto) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *InvoiceCreateDto) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidToNil

`func (o *InvoiceCreateDto) SetValidToNil(b bool)`

 SetValidToNil sets the value for ValidTo to be an explicit nil

### UnsetValidTo
`func (o *InvoiceCreateDto) UnsetValidTo()`

UnsetValidTo ensures that no value is present for ValidTo, not even an explicit nil
### GetInvoiceReferences

`func (o *InvoiceCreateDto) GetInvoiceReferences() []InvoiceReferenceDto`

GetInvoiceReferences returns the InvoiceReferences field if non-nil, zero value otherwise.

### GetInvoiceReferencesOk

`func (o *InvoiceCreateDto) GetInvoiceReferencesOk() (*[]InvoiceReferenceDto, bool)`

GetInvoiceReferencesOk returns a tuple with the InvoiceReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReferences

`func (o *InvoiceCreateDto) SetInvoiceReferences(v []InvoiceReferenceDto)`

SetInvoiceReferences sets InvoiceReferences field to given value.

### HasInvoiceReferences

`func (o *InvoiceCreateDto) HasInvoiceReferences() bool`

HasInvoiceReferences returns a boolean if a field has been set.

### SetInvoiceReferencesNil

`func (o *InvoiceCreateDto) SetInvoiceReferencesNil(b bool)`

 SetInvoiceReferencesNil sets the value for InvoiceReferences to be an explicit nil

### UnsetInvoiceReferences
`func (o *InvoiceCreateDto) UnsetInvoiceReferences()`

UnsetInvoiceReferences ensures that no value is present for InvoiceReferences, not even an explicit nil
### GetInvoiceItemRecords

`func (o *InvoiceCreateDto) GetInvoiceItemRecords() []InvoiceItemRecordDto`

GetInvoiceItemRecords returns the InvoiceItemRecords field if non-nil, zero value otherwise.

### GetInvoiceItemRecordsOk

`func (o *InvoiceCreateDto) GetInvoiceItemRecordsOk() (*[]InvoiceItemRecordDto, bool)`

GetInvoiceItemRecordsOk returns a tuple with the InvoiceItemRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemRecords

`func (o *InvoiceCreateDto) SetInvoiceItemRecords(v []InvoiceItemRecordDto)`

SetInvoiceItemRecords sets InvoiceItemRecords field to given value.

### HasInvoiceItemRecords

`func (o *InvoiceCreateDto) HasInvoiceItemRecords() bool`

HasInvoiceItemRecords returns a boolean if a field has been set.

### SetInvoiceItemRecordsNil

`func (o *InvoiceCreateDto) SetInvoiceItemRecordsNil(b bool)`

 SetInvoiceItemRecordsNil sets the value for InvoiceItemRecords to be an explicit nil

### UnsetInvoiceItemRecords
`func (o *InvoiceCreateDto) UnsetInvoiceItemRecords()`

UnsetInvoiceItemRecords ensures that no value is present for InvoiceItemRecords, not even an explicit nil
### GetInvoiceAdjustments

`func (o *InvoiceCreateDto) GetInvoiceAdjustments() []InvoiceAdjustmentDto`

GetInvoiceAdjustments returns the InvoiceAdjustments field if non-nil, zero value otherwise.

### GetInvoiceAdjustmentsOk

`func (o *InvoiceCreateDto) GetInvoiceAdjustmentsOk() (*[]InvoiceAdjustmentDto, bool)`

GetInvoiceAdjustmentsOk returns a tuple with the InvoiceAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAdjustments

`func (o *InvoiceCreateDto) SetInvoiceAdjustments(v []InvoiceAdjustmentDto)`

SetInvoiceAdjustments sets InvoiceAdjustments field to given value.

### HasInvoiceAdjustments

`func (o *InvoiceCreateDto) HasInvoiceAdjustments() bool`

HasInvoiceAdjustments returns a boolean if a field has been set.

### SetInvoiceAdjustmentsNil

`func (o *InvoiceCreateDto) SetInvoiceAdjustmentsNil(b bool)`

 SetInvoiceAdjustmentsNil sets the value for InvoiceAdjustments to be an explicit nil

### UnsetInvoiceAdjustments
`func (o *InvoiceCreateDto) UnsetInvoiceAdjustments()`

UnsetInvoiceAdjustments ensures that no value is present for InvoiceAdjustments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


