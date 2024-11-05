# DealUnitUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Closed** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
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
**BillingLocationId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
**ShippingMethodId** | Pointer to **NullableString** |  | [optional] 
**Ordered** | Pointer to **bool** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**DealUnitFeedId** | Pointer to **NullableString** |  | [optional] 
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

### NewDealUnitUpdateDto

`func NewDealUnitUpdateDto() *DealUnitUpdateDto`

NewDealUnitUpdateDto instantiates a new DealUnitUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitUpdateDtoWithDefaults

`func NewDealUnitUpdateDtoWithDefaults() *DealUnitUpdateDto`

NewDealUnitUpdateDtoWithDefaults instantiates a new DealUnitUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosed

`func (o *DealUnitUpdateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *DealUnitUpdateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *DealUnitUpdateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *DealUnitUpdateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetTitle

`func (o *DealUnitUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DealUnitUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DealUnitUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DealUnitUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DealUnitUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DealUnitUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *DealUnitUpdateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DealUnitUpdateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DealUnitUpdateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DealUnitUpdateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DealUnitUpdateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DealUnitUpdateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *DealUnitUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetForexRate

`func (o *DealUnitUpdateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *DealUnitUpdateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *DealUnitUpdateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *DealUnitUpdateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *DealUnitUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *DealUnitUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *DealUnitUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *DealUnitUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *DealUnitUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *DealUnitUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPriceListId

`func (o *DealUnitUpdateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *DealUnitUpdateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *DealUnitUpdateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *DealUnitUpdateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *DealUnitUpdateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *DealUnitUpdateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetDescription

`func (o *DealUnitUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnrollmentId

`func (o *DealUnitUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *DealUnitUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *DealUnitUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *DealUnitUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *DealUnitUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *DealUnitUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *DealUnitUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *DealUnitUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *DealUnitUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *DealUnitUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *DealUnitUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *DealUnitUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetPaymentTermId

`func (o *DealUnitUpdateDto) GetPaymentTermId() string`

GetPaymentTermId returns the PaymentTermId field if non-nil, zero value otherwise.

### GetPaymentTermIdOk

`func (o *DealUnitUpdateDto) GetPaymentTermIdOk() (*string, bool)`

GetPaymentTermIdOk returns a tuple with the PaymentTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermId

`func (o *DealUnitUpdateDto) SetPaymentTermId(v string)`

SetPaymentTermId sets PaymentTermId field to given value.

### HasPaymentTermId

`func (o *DealUnitUpdateDto) HasPaymentTermId() bool`

HasPaymentTermId returns a boolean if a field has been set.

### SetPaymentTermIdNil

`func (o *DealUnitUpdateDto) SetPaymentTermIdNil(b bool)`

 SetPaymentTermIdNil sets the value for PaymentTermId to be an explicit nil

### UnsetPaymentTermId
`func (o *DealUnitUpdateDto) UnsetPaymentTermId()`

UnsetPaymentTermId ensures that no value is present for PaymentTermId, not even an explicit nil
### GetOrganizationId

`func (o *DealUnitUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DealUnitUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DealUnitUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DealUnitUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *DealUnitUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DealUnitUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *DealUnitUpdateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *DealUnitUpdateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *DealUnitUpdateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *DealUnitUpdateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *DealUnitUpdateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *DealUnitUpdateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *DealUnitUpdateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DealUnitUpdateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DealUnitUpdateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DealUnitUpdateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *DealUnitUpdateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *DealUnitUpdateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *DealUnitUpdateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DealUnitUpdateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DealUnitUpdateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DealUnitUpdateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *DealUnitUpdateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *DealUnitUpdateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *DealUnitUpdateDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DealUnitUpdateDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DealUnitUpdateDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DealUnitUpdateDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *DealUnitUpdateDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *DealUnitUpdateDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *DealUnitUpdateDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *DealUnitUpdateDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *DealUnitUpdateDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *DealUnitUpdateDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *DealUnitUpdateDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *DealUnitUpdateDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *DealUnitUpdateDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DealUnitUpdateDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DealUnitUpdateDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DealUnitUpdateDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *DealUnitUpdateDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *DealUnitUpdateDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *DealUnitUpdateDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DealUnitUpdateDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DealUnitUpdateDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DealUnitUpdateDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *DealUnitUpdateDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *DealUnitUpdateDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *DealUnitUpdateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DealUnitUpdateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DealUnitUpdateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *DealUnitUpdateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *DealUnitUpdateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *DealUnitUpdateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *DealUnitUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *DealUnitUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *DealUnitUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *DealUnitUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *DealUnitUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *DealUnitUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *DealUnitUpdateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *DealUnitUpdateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *DealUnitUpdateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *DealUnitUpdateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *DealUnitUpdateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *DealUnitUpdateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *DealUnitUpdateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *DealUnitUpdateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *DealUnitUpdateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *DealUnitUpdateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *DealUnitUpdateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *DealUnitUpdateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetBillingLocationId

`func (o *DealUnitUpdateDto) GetBillingLocationId() string`

GetBillingLocationId returns the BillingLocationId field if non-nil, zero value otherwise.

### GetBillingLocationIdOk

`func (o *DealUnitUpdateDto) GetBillingLocationIdOk() (*string, bool)`

GetBillingLocationIdOk returns a tuple with the BillingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLocationId

`func (o *DealUnitUpdateDto) SetBillingLocationId(v string)`

SetBillingLocationId sets BillingLocationId field to given value.

### HasBillingLocationId

`func (o *DealUnitUpdateDto) HasBillingLocationId() bool`

HasBillingLocationId returns a boolean if a field has been set.

### SetBillingLocationIdNil

`func (o *DealUnitUpdateDto) SetBillingLocationIdNil(b bool)`

 SetBillingLocationIdNil sets the value for BillingLocationId to be an explicit nil

### UnsetBillingLocationId
`func (o *DealUnitUpdateDto) UnsetBillingLocationId()`

UnsetBillingLocationId ensures that no value is present for BillingLocationId, not even an explicit nil
### GetShippingLocationId

`func (o *DealUnitUpdateDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *DealUnitUpdateDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *DealUnitUpdateDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *DealUnitUpdateDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *DealUnitUpdateDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *DealUnitUpdateDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetShippingMethodId

`func (o *DealUnitUpdateDto) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *DealUnitUpdateDto) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *DealUnitUpdateDto) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *DealUnitUpdateDto) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### SetShippingMethodIdNil

`func (o *DealUnitUpdateDto) SetShippingMethodIdNil(b bool)`

 SetShippingMethodIdNil sets the value for ShippingMethodId to be an explicit nil

### UnsetShippingMethodId
`func (o *DealUnitUpdateDto) UnsetShippingMethodId()`

UnsetShippingMethodId ensures that no value is present for ShippingMethodId, not even an explicit nil
### GetOrdered

`func (o *DealUnitUpdateDto) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *DealUnitUpdateDto) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *DealUnitUpdateDto) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.

### HasOrdered

`func (o *DealUnitUpdateDto) HasOrdered() bool`

HasOrdered returns a boolean if a field has been set.

### GetCartId

`func (o *DealUnitUpdateDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *DealUnitUpdateDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *DealUnitUpdateDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *DealUnitUpdateDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *DealUnitUpdateDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *DealUnitUpdateDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetDealUnitFeedId

`func (o *DealUnitUpdateDto) GetDealUnitFeedId() string`

GetDealUnitFeedId returns the DealUnitFeedId field if non-nil, zero value otherwise.

### GetDealUnitFeedIdOk

`func (o *DealUnitUpdateDto) GetDealUnitFeedIdOk() (*string, bool)`

GetDealUnitFeedIdOk returns a tuple with the DealUnitFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFeedId

`func (o *DealUnitUpdateDto) SetDealUnitFeedId(v string)`

SetDealUnitFeedId sets DealUnitFeedId field to given value.

### HasDealUnitFeedId

`func (o *DealUnitUpdateDto) HasDealUnitFeedId() bool`

HasDealUnitFeedId returns a boolean if a field has been set.

### SetDealUnitFeedIdNil

`func (o *DealUnitUpdateDto) SetDealUnitFeedIdNil(b bool)`

 SetDealUnitFeedIdNil sets the value for DealUnitFeedId to be an explicit nil

### UnsetDealUnitFeedId
`func (o *DealUnitUpdateDto) UnsetDealUnitFeedId()`

UnsetDealUnitFeedId ensures that no value is present for DealUnitFeedId, not even an explicit nil
### GetDealUnitFlowId

`func (o *DealUnitUpdateDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitUpdateDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitUpdateDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitUpdateDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitUpdateDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitUpdateDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetDealUnitFlowStageId

`func (o *DealUnitUpdateDto) GetDealUnitFlowStageId() string`

GetDealUnitFlowStageId returns the DealUnitFlowStageId field if non-nil, zero value otherwise.

### GetDealUnitFlowStageIdOk

`func (o *DealUnitUpdateDto) GetDealUnitFlowStageIdOk() (*string, bool)`

GetDealUnitFlowStageIdOk returns a tuple with the DealUnitFlowStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowStageId

`func (o *DealUnitUpdateDto) SetDealUnitFlowStageId(v string)`

SetDealUnitFlowStageId sets DealUnitFlowStageId field to given value.

### HasDealUnitFlowStageId

`func (o *DealUnitUpdateDto) HasDealUnitFlowStageId() bool`

HasDealUnitFlowStageId returns a boolean if a field has been set.

### SetDealUnitFlowStageIdNil

`func (o *DealUnitUpdateDto) SetDealUnitFlowStageIdNil(b bool)`

 SetDealUnitFlowStageIdNil sets the value for DealUnitFlowStageId to be an explicit nil

### UnsetDealUnitFlowStageId
`func (o *DealUnitUpdateDto) UnsetDealUnitFlowStageId()`

UnsetDealUnitFlowStageId ensures that no value is present for DealUnitFlowStageId, not even an explicit nil
### GetPartnerCreated

`func (o *DealUnitUpdateDto) GetPartnerCreated() bool`

GetPartnerCreated returns the PartnerCreated field if non-nil, zero value otherwise.

### GetPartnerCreatedOk

`func (o *DealUnitUpdateDto) GetPartnerCreatedOk() (*bool, bool)`

GetPartnerCreatedOk returns a tuple with the PartnerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCreated

`func (o *DealUnitUpdateDto) SetPartnerCreated(v bool)`

SetPartnerCreated sets PartnerCreated field to given value.

### HasPartnerCreated

`func (o *DealUnitUpdateDto) HasPartnerCreated() bool`

HasPartnerCreated returns a boolean if a field has been set.

### GetPartnerCollaboration

`func (o *DealUnitUpdateDto) GetPartnerCollaboration() bool`

GetPartnerCollaboration returns the PartnerCollaboration field if non-nil, zero value otherwise.

### GetPartnerCollaborationOk

`func (o *DealUnitUpdateDto) GetPartnerCollaborationOk() (*bool, bool)`

GetPartnerCollaborationOk returns a tuple with the PartnerCollaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCollaboration

`func (o *DealUnitUpdateDto) SetPartnerCollaboration(v bool)`

SetPartnerCollaboration sets PartnerCollaboration field to given value.

### HasPartnerCollaboration

`func (o *DealUnitUpdateDto) HasPartnerCollaboration() bool`

HasPartnerCollaboration returns a boolean if a field has been set.

### GetProposedSolution

`func (o *DealUnitUpdateDto) GetProposedSolution() string`

GetProposedSolution returns the ProposedSolution field if non-nil, zero value otherwise.

### GetProposedSolutionOk

`func (o *DealUnitUpdateDto) GetProposedSolutionOk() (*string, bool)`

GetProposedSolutionOk returns a tuple with the ProposedSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedSolution

`func (o *DealUnitUpdateDto) SetProposedSolution(v string)`

SetProposedSolution sets ProposedSolution field to given value.

### HasProposedSolution

`func (o *DealUnitUpdateDto) HasProposedSolution() bool`

HasProposedSolution returns a boolean if a field has been set.

### SetProposedSolutionNil

`func (o *DealUnitUpdateDto) SetProposedSolutionNil(b bool)`

 SetProposedSolutionNil sets the value for ProposedSolution to be an explicit nil

### UnsetProposedSolution
`func (o *DealUnitUpdateDto) UnsetProposedSolution()`

UnsetProposedSolution ensures that no value is present for ProposedSolution, not even an explicit nil
### GetCurrentSituation

`func (o *DealUnitUpdateDto) GetCurrentSituation() string`

GetCurrentSituation returns the CurrentSituation field if non-nil, zero value otherwise.

### GetCurrentSituationOk

`func (o *DealUnitUpdateDto) GetCurrentSituationOk() (*string, bool)`

GetCurrentSituationOk returns a tuple with the CurrentSituation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSituation

`func (o *DealUnitUpdateDto) SetCurrentSituation(v string)`

SetCurrentSituation sets CurrentSituation field to given value.

### HasCurrentSituation

`func (o *DealUnitUpdateDto) HasCurrentSituation() bool`

HasCurrentSituation returns a boolean if a field has been set.

### SetCurrentSituationNil

`func (o *DealUnitUpdateDto) SetCurrentSituationNil(b bool)`

 SetCurrentSituationNil sets the value for CurrentSituation to be an explicit nil

### UnsetCurrentSituation
`func (o *DealUnitUpdateDto) UnsetCurrentSituation()`

UnsetCurrentSituation ensures that no value is present for CurrentSituation, not even an explicit nil
### GetCustomerNeed

`func (o *DealUnitUpdateDto) GetCustomerNeed() string`

GetCustomerNeed returns the CustomerNeed field if non-nil, zero value otherwise.

### GetCustomerNeedOk

`func (o *DealUnitUpdateDto) GetCustomerNeedOk() (*string, bool)`

GetCustomerNeedOk returns a tuple with the CustomerNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNeed

`func (o *DealUnitUpdateDto) SetCustomerNeed(v string)`

SetCustomerNeed sets CustomerNeed field to given value.

### HasCustomerNeed

`func (o *DealUnitUpdateDto) HasCustomerNeed() bool`

HasCustomerNeed returns a boolean if a field has been set.

### SetCustomerNeedNil

`func (o *DealUnitUpdateDto) SetCustomerNeedNil(b bool)`

 SetCustomerNeedNil sets the value for CustomerNeed to be an explicit nil

### UnsetCustomerNeed
`func (o *DealUnitUpdateDto) UnsetCustomerNeed()`

UnsetCustomerNeed ensures that no value is present for CustomerNeed, not even an explicit nil
### GetWonDate

`func (o *DealUnitUpdateDto) GetWonDate() time.Time`

GetWonDate returns the WonDate field if non-nil, zero value otherwise.

### GetWonDateOk

`func (o *DealUnitUpdateDto) GetWonDateOk() (*time.Time, bool)`

GetWonDateOk returns a tuple with the WonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWonDate

`func (o *DealUnitUpdateDto) SetWonDate(v time.Time)`

SetWonDate sets WonDate field to given value.

### HasWonDate

`func (o *DealUnitUpdateDto) HasWonDate() bool`

HasWonDate returns a boolean if a field has been set.

### GetLostDate

`func (o *DealUnitUpdateDto) GetLostDate() time.Time`

GetLostDate returns the LostDate field if non-nil, zero value otherwise.

### GetLostDateOk

`func (o *DealUnitUpdateDto) GetLostDateOk() (*time.Time, bool)`

GetLostDateOk returns a tuple with the LostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostDate

`func (o *DealUnitUpdateDto) SetLostDate(v time.Time)`

SetLostDate sets LostDate field to given value.

### HasLostDate

`func (o *DealUnitUpdateDto) HasLostDate() bool`

HasLostDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *DealUnitUpdateDto) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *DealUnitUpdateDto) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *DealUnitUpdateDto) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *DealUnitUpdateDto) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *DealUnitUpdateDto) GetDeliveredDate() time.Time`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *DealUnitUpdateDto) GetDeliveredDateOk() (*time.Time, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *DealUnitUpdateDto) SetDeliveredDate(v time.Time)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *DealUnitUpdateDto) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### GetClosedTimestamp

`func (o *DealUnitUpdateDto) GetClosedTimestamp() time.Time`

GetClosedTimestamp returns the ClosedTimestamp field if non-nil, zero value otherwise.

### GetClosedTimestampOk

`func (o *DealUnitUpdateDto) GetClosedTimestampOk() (*time.Time, bool)`

GetClosedTimestampOk returns a tuple with the ClosedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTimestamp

`func (o *DealUnitUpdateDto) SetClosedTimestamp(v time.Time)`

SetClosedTimestamp sets ClosedTimestamp field to given value.

### HasClosedTimestamp

`func (o *DealUnitUpdateDto) HasClosedTimestamp() bool`

HasClosedTimestamp returns a boolean if a field has been set.

### GetExpectedCloseDate

`func (o *DealUnitUpdateDto) GetExpectedCloseDate() time.Time`

GetExpectedCloseDate returns the ExpectedCloseDate field if non-nil, zero value otherwise.

### GetExpectedCloseDateOk

`func (o *DealUnitUpdateDto) GetExpectedCloseDateOk() (*time.Time, bool)`

GetExpectedCloseDateOk returns a tuple with the ExpectedCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCloseDate

`func (o *DealUnitUpdateDto) SetExpectedCloseDate(v time.Time)`

SetExpectedCloseDate sets ExpectedCloseDate field to given value.

### HasExpectedCloseDate

`func (o *DealUnitUpdateDto) HasExpectedCloseDate() bool`

HasExpectedCloseDate returns a boolean if a field has been set.

### GetDealUnitStatus

`func (o *DealUnitUpdateDto) GetDealUnitStatus() int32`

GetDealUnitStatus returns the DealUnitStatus field if non-nil, zero value otherwise.

### GetDealUnitStatusOk

`func (o *DealUnitUpdateDto) GetDealUnitStatusOk() (*int32, bool)`

GetDealUnitStatusOk returns a tuple with the DealUnitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitStatus

`func (o *DealUnitUpdateDto) SetDealUnitStatus(v int32)`

SetDealUnitStatus sets DealUnitStatus field to given value.

### HasDealUnitStatus

`func (o *DealUnitUpdateDto) HasDealUnitStatus() bool`

HasDealUnitStatus returns a boolean if a field has been set.

### GetDealUnitPurchaseProcess

`func (o *DealUnitUpdateDto) GetDealUnitPurchaseProcess() int32`

GetDealUnitPurchaseProcess returns the DealUnitPurchaseProcess field if non-nil, zero value otherwise.

### GetDealUnitPurchaseProcessOk

`func (o *DealUnitUpdateDto) GetDealUnitPurchaseProcessOk() (*int32, bool)`

GetDealUnitPurchaseProcessOk returns a tuple with the DealUnitPurchaseProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitPurchaseProcess

`func (o *DealUnitUpdateDto) SetDealUnitPurchaseProcess(v int32)`

SetDealUnitPurchaseProcess sets DealUnitPurchaseProcess field to given value.

### HasDealUnitPurchaseProcess

`func (o *DealUnitUpdateDto) HasDealUnitPurchaseProcess() bool`

HasDealUnitPurchaseProcess returns a boolean if a field has been set.

### GetDealUnitForecastCategory

`func (o *DealUnitUpdateDto) GetDealUnitForecastCategory() int32`

GetDealUnitForecastCategory returns the DealUnitForecastCategory field if non-nil, zero value otherwise.

### GetDealUnitForecastCategoryOk

`func (o *DealUnitUpdateDto) GetDealUnitForecastCategoryOk() (*int32, bool)`

GetDealUnitForecastCategoryOk returns a tuple with the DealUnitForecastCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitForecastCategory

`func (o *DealUnitUpdateDto) SetDealUnitForecastCategory(v int32)`

SetDealUnitForecastCategory sets DealUnitForecastCategory field to given value.

### HasDealUnitForecastCategory

`func (o *DealUnitUpdateDto) HasDealUnitForecastCategory() bool`

HasDealUnitForecastCategory returns a boolean if a field has been set.

### GetDealUnitAmountsCalculation

`func (o *DealUnitUpdateDto) GetDealUnitAmountsCalculation() int32`

GetDealUnitAmountsCalculation returns the DealUnitAmountsCalculation field if non-nil, zero value otherwise.

### GetDealUnitAmountsCalculationOk

`func (o *DealUnitUpdateDto) GetDealUnitAmountsCalculationOk() (*int32, bool)`

GetDealUnitAmountsCalculationOk returns a tuple with the DealUnitAmountsCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitAmountsCalculation

`func (o *DealUnitUpdateDto) SetDealUnitAmountsCalculation(v int32)`

SetDealUnitAmountsCalculation sets DealUnitAmountsCalculation field to given value.

### HasDealUnitAmountsCalculation

`func (o *DealUnitUpdateDto) HasDealUnitAmountsCalculation() bool`

HasDealUnitAmountsCalculation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


