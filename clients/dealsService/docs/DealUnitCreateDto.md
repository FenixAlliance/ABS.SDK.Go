# DealUnitCreateDto

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
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
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
**DealUnitStatus** | Pointer to **int32** |  | [optional] 
**DealUnitPurchaseProcess** | Pointer to **int32** |  | [optional] 
**DealUnitForecastCategory** | Pointer to **int32** |  | [optional] 
**DealUnitAmountsCalculation** | Pointer to **int32** |  | [optional] 

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
### GetUserId

`func (o *DealUnitCreateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DealUnitCreateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DealUnitCreateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DealUnitCreateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DealUnitCreateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DealUnitCreateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *DealUnitCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
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
### GetEnrollmentId

`func (o *DealUnitCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *DealUnitCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *DealUnitCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *DealUnitCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *DealUnitCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *DealUnitCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
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

`func (o *DealUnitCreateDto) GetDealUnitStatus() int32`

GetDealUnitStatus returns the DealUnitStatus field if non-nil, zero value otherwise.

### GetDealUnitStatusOk

`func (o *DealUnitCreateDto) GetDealUnitStatusOk() (*int32, bool)`

GetDealUnitStatusOk returns a tuple with the DealUnitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitStatus

`func (o *DealUnitCreateDto) SetDealUnitStatus(v int32)`

SetDealUnitStatus sets DealUnitStatus field to given value.

### HasDealUnitStatus

`func (o *DealUnitCreateDto) HasDealUnitStatus() bool`

HasDealUnitStatus returns a boolean if a field has been set.

### GetDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) GetDealUnitPurchaseProcess() int32`

GetDealUnitPurchaseProcess returns the DealUnitPurchaseProcess field if non-nil, zero value otherwise.

### GetDealUnitPurchaseProcessOk

`func (o *DealUnitCreateDto) GetDealUnitPurchaseProcessOk() (*int32, bool)`

GetDealUnitPurchaseProcessOk returns a tuple with the DealUnitPurchaseProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) SetDealUnitPurchaseProcess(v int32)`

SetDealUnitPurchaseProcess sets DealUnitPurchaseProcess field to given value.

### HasDealUnitPurchaseProcess

`func (o *DealUnitCreateDto) HasDealUnitPurchaseProcess() bool`

HasDealUnitPurchaseProcess returns a boolean if a field has been set.

### GetDealUnitForecastCategory

`func (o *DealUnitCreateDto) GetDealUnitForecastCategory() int32`

GetDealUnitForecastCategory returns the DealUnitForecastCategory field if non-nil, zero value otherwise.

### GetDealUnitForecastCategoryOk

`func (o *DealUnitCreateDto) GetDealUnitForecastCategoryOk() (*int32, bool)`

GetDealUnitForecastCategoryOk returns a tuple with the DealUnitForecastCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitForecastCategory

`func (o *DealUnitCreateDto) SetDealUnitForecastCategory(v int32)`

SetDealUnitForecastCategory sets DealUnitForecastCategory field to given value.

### HasDealUnitForecastCategory

`func (o *DealUnitCreateDto) HasDealUnitForecastCategory() bool`

HasDealUnitForecastCategory returns a boolean if a field has been set.

### GetDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) GetDealUnitAmountsCalculation() int32`

GetDealUnitAmountsCalculation returns the DealUnitAmountsCalculation field if non-nil, zero value otherwise.

### GetDealUnitAmountsCalculationOk

`func (o *DealUnitCreateDto) GetDealUnitAmountsCalculationOk() (*int32, bool)`

GetDealUnitAmountsCalculationOk returns a tuple with the DealUnitAmountsCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) SetDealUnitAmountsCalculation(v int32)`

SetDealUnitAmountsCalculation sets DealUnitAmountsCalculation field to given value.

### HasDealUnitAmountsCalculation

`func (o *DealUnitCreateDto) HasDealUnitAmountsCalculation() bool`

HasDealUnitAmountsCalculation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


