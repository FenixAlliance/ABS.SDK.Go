# DealUnitDto

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
**Ordered** | Pointer to **bool** |  | [optional] 
**DealUnitFeedId** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowId** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowStageId** | Pointer to **NullableString** |  | [optional] 
**BillingLocationId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
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
**LinesCount** | Pointer to **int32** |  | [optional] 
**CustomTotalAmount** | Pointer to **float64** |  | [optional] 
**CustomDetailAmount** | Pointer to **float64** |  | [optional] 
**CustomProfitAmount** | Pointer to **float64** |  | [optional] 
**CustomShippingCostAmount** | Pointer to **float64** |  | [optional] 
**CustomWithholdingTaxAmount** | Pointer to **float64** |  | [optional] 
**CustomSurchargesAmount** | Pointer to **float64** |  | [optional] 
**CustomDiscountsAmount** | Pointer to **float64** |  | [optional] 
**CustomShippingTaxAmount** | Pointer to **float64** |  | [optional] 

## Methods

### NewDealUnitDto

`func NewDealUnitDto() *DealUnitDto`

NewDealUnitDto instantiates a new DealUnitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitDtoWithDefaults

`func NewDealUnitDtoWithDefaults() *DealUnitDto`

NewDealUnitDtoWithDefaults instantiates a new DealUnitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DealUnitDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DealUnitDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *DealUnitDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DealUnitDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DealUnitDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *DealUnitDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *DealUnitDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *DealUnitDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *DealUnitDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetType

`func (o *DealUnitDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealUnitDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealUnitDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DealUnitDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DealUnitDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DealUnitDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *DealUnitDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DealUnitDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DealUnitDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DealUnitDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DealUnitDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DealUnitDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *DealUnitDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DealUnitDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DealUnitDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DealUnitDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DealUnitDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DealUnitDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *DealUnitDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCurrencyId

`func (o *DealUnitDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *DealUnitDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *DealUnitDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *DealUnitDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *DealUnitDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *DealUnitDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *DealUnitDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceListId

`func (o *DealUnitDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *DealUnitDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *DealUnitDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *DealUnitDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *DealUnitDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *DealUnitDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetEnrollmentId

`func (o *DealUnitDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *DealUnitDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *DealUnitDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *DealUnitDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *DealUnitDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *DealUnitDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetIndividualId

`func (o *DealUnitDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *DealUnitDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *DealUnitDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *DealUnitDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *DealUnitDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *DealUnitDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *DealUnitDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DealUnitDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DealUnitDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DealUnitDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *DealUnitDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DealUnitDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *DealUnitDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *DealUnitDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *DealUnitDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *DealUnitDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *DealUnitDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *DealUnitDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetFirstName

`func (o *DealUnitDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DealUnitDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DealUnitDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DealUnitDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *DealUnitDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *DealUnitDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *DealUnitDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DealUnitDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DealUnitDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DealUnitDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *DealUnitDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *DealUnitDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *DealUnitDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DealUnitDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DealUnitDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DealUnitDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *DealUnitDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *DealUnitDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetBillingEmail

`func (o *DealUnitDto) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *DealUnitDto) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *DealUnitDto) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *DealUnitDto) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *DealUnitDto) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *DealUnitDto) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddressLine1

`func (o *DealUnitDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DealUnitDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DealUnitDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DealUnitDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *DealUnitDto) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *DealUnitDto) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *DealUnitDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DealUnitDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DealUnitDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DealUnitDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *DealUnitDto) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *DealUnitDto) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *DealUnitDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DealUnitDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DealUnitDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *DealUnitDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *DealUnitDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *DealUnitDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *DealUnitDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *DealUnitDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *DealUnitDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *DealUnitDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *DealUnitDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *DealUnitDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *DealUnitDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *DealUnitDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *DealUnitDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *DealUnitDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *DealUnitDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *DealUnitDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *DealUnitDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *DealUnitDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *DealUnitDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *DealUnitDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *DealUnitDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *DealUnitDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCustomerNotes

`func (o *DealUnitDto) GetCustomerNotes() string`

GetCustomerNotes returns the CustomerNotes field if non-nil, zero value otherwise.

### GetCustomerNotesOk

`func (o *DealUnitDto) GetCustomerNotesOk() (*string, bool)`

GetCustomerNotesOk returns a tuple with the CustomerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotes

`func (o *DealUnitDto) SetCustomerNotes(v string)`

SetCustomerNotes sets CustomerNotes field to given value.

### HasCustomerNotes

`func (o *DealUnitDto) HasCustomerNotes() bool`

HasCustomerNotes returns a boolean if a field has been set.

### SetCustomerNotesNil

`func (o *DealUnitDto) SetCustomerNotesNil(b bool)`

 SetCustomerNotesNil sets the value for CustomerNotes to be an explicit nil

### UnsetCustomerNotes
`func (o *DealUnitDto) UnsetCustomerNotes()`

UnsetCustomerNotes ensures that no value is present for CustomerNotes, not even an explicit nil
### GetForexRate

`func (o *DealUnitDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *DealUnitDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *DealUnitDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *DealUnitDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotal

`func (o *DealUnitDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DealUnitDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DealUnitDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DealUnitDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *DealUnitDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *DealUnitDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *DealUnitDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *DealUnitDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxBase

`func (o *DealUnitDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *DealUnitDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *DealUnitDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *DealUnitDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *DealUnitDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *DealUnitDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *DealUnitDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *DealUnitDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalSurcharges

`func (o *DealUnitDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *DealUnitDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *DealUnitDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *DealUnitDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalDiscounts

`func (o *DealUnitDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *DealUnitDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *DealUnitDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *DealUnitDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalSurcharges

`func (o *DealUnitDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *DealUnitDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *DealUnitDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *DealUnitDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *DealUnitDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *DealUnitDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *DealUnitDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *DealUnitDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalAmountInUsd

`func (o *DealUnitDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *DealUnitDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *DealUnitDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *DealUnitDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *DealUnitDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *DealUnitDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *DealUnitDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *DealUnitDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *DealUnitDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *DealUnitDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *DealUnitDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *DealUnitDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *DealUnitDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *DealUnitDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *DealUnitDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *DealUnitDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *DealUnitDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *DealUnitDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *DealUnitDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *DealUnitDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalDetailAmountInUsd

`func (o *DealUnitDto) GetTotalDetailAmountInUsd() float64`

GetTotalDetailAmountInUsd returns the TotalDetailAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDetailAmountInUsdOk

`func (o *DealUnitDto) GetTotalDetailAmountInUsdOk() (*float64, bool)`

GetTotalDetailAmountInUsdOk returns a tuple with the TotalDetailAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmountInUsd

`func (o *DealUnitDto) SetTotalDetailAmountInUsd(v float64)`

SetTotalDetailAmountInUsd sets TotalDetailAmountInUsd field to given value.

### HasTotalDetailAmountInUsd

`func (o *DealUnitDto) HasTotalDetailAmountInUsd() bool`

HasTotalDetailAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *DealUnitDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *DealUnitDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *DealUnitDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *DealUnitDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *DealUnitDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *DealUnitDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *DealUnitDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *DealUnitDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalWithholdingTaxesInUsd

`func (o *DealUnitDto) GetTotalWithholdingTaxesInUsd() float64`

GetTotalWithholdingTaxesInUsd returns the TotalWithholdingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithholdingTaxesInUsdOk

`func (o *DealUnitDto) GetTotalWithholdingTaxesInUsdOk() (*float64, bool)`

GetTotalWithholdingTaxesInUsdOk returns a tuple with the TotalWithholdingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithholdingTaxesInUsd

`func (o *DealUnitDto) SetTotalWithholdingTaxesInUsd(v float64)`

SetTotalWithholdingTaxesInUsd sets TotalWithholdingTaxesInUsd field to given value.

### HasTotalWithholdingTaxesInUsd

`func (o *DealUnitDto) HasTotalWithholdingTaxesInUsd() bool`

HasTotalWithholdingTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *DealUnitDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *DealUnitDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *DealUnitDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *DealUnitDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *DealUnitDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *DealUnitDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *DealUnitDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *DealUnitDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetCurrency

`func (o *DealUnitDto) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DealUnitDto) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DealUnitDto) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DealUnitDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *DealUnitDto) GetTotalInUsd() Money`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *DealUnitDto) GetTotalInUsdOk() (*Money, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *DealUnitDto) SetTotalInUsd(v Money)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *DealUnitDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetTotalTaxAmountInUsd

`func (o *DealUnitDto) GetTotalTaxAmountInUsd() Money`

GetTotalTaxAmountInUsd returns the TotalTaxAmountInUsd field if non-nil, zero value otherwise.

### GetTotalTaxAmountInUsdOk

`func (o *DealUnitDto) GetTotalTaxAmountInUsdOk() (*Money, bool)`

GetTotalTaxAmountInUsdOk returns a tuple with the TotalTaxAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountInUsd

`func (o *DealUnitDto) SetTotalTaxAmountInUsd(v Money)`

SetTotalTaxAmountInUsd sets TotalTaxAmountInUsd field to given value.

### HasTotalTaxAmountInUsd

`func (o *DealUnitDto) HasTotalTaxAmountInUsd() bool`

HasTotalTaxAmountInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseAmountInUsd

`func (o *DealUnitDto) GetTotalTaxBaseAmountInUsd() Money`

GetTotalTaxBaseAmountInUsd returns the TotalTaxBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountInUsdOk

`func (o *DealUnitDto) GetTotalTaxBaseAmountInUsdOk() (*Money, bool)`

GetTotalTaxBaseAmountInUsdOk returns a tuple with the TotalTaxBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmountInUsd

`func (o *DealUnitDto) SetTotalTaxBaseAmountInUsd(v Money)`

SetTotalTaxBaseAmountInUsd sets TotalTaxBaseAmountInUsd field to given value.

### HasTotalTaxBaseAmountInUsd

`func (o *DealUnitDto) HasTotalTaxBaseAmountInUsd() bool`

HasTotalTaxBaseAmountInUsd returns a boolean if a field has been set.

### GetTotalDiscountsAmountInUsd

`func (o *DealUnitDto) GetTotalDiscountsAmountInUsd() Money`

GetTotalDiscountsAmountInUsd returns the TotalDiscountsAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountInUsdOk

`func (o *DealUnitDto) GetTotalDiscountsAmountInUsdOk() (*Money, bool)`

GetTotalDiscountsAmountInUsdOk returns a tuple with the TotalDiscountsAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmountInUsd

`func (o *DealUnitDto) SetTotalDiscountsAmountInUsd(v Money)`

SetTotalDiscountsAmountInUsd sets TotalDiscountsAmountInUsd field to given value.

### HasTotalDiscountsAmountInUsd

`func (o *DealUnitDto) HasTotalDiscountsAmountInUsd() bool`

HasTotalDiscountsAmountInUsd returns a boolean if a field has been set.

### GetTotalSurchargesAmountInUsd

`func (o *DealUnitDto) GetTotalSurchargesAmountInUsd() Money`

GetTotalSurchargesAmountInUsd returns the TotalSurchargesAmountInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountInUsdOk

`func (o *DealUnitDto) GetTotalSurchargesAmountInUsdOk() (*Money, bool)`

GetTotalSurchargesAmountInUsdOk returns a tuple with the TotalSurchargesAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmountInUsd

`func (o *DealUnitDto) SetTotalSurchargesAmountInUsd(v Money)`

SetTotalSurchargesAmountInUsd sets TotalSurchargesAmountInUsd field to given value.

### HasTotalSurchargesAmountInUsd

`func (o *DealUnitDto) HasTotalSurchargesAmountInUsd() bool`

HasTotalSurchargesAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsAmountInUsd

`func (o *DealUnitDto) GetTotalGlobalDiscountsAmountInUsd() Money`

GetTotalGlobalDiscountsAmountInUsd returns the TotalGlobalDiscountsAmountInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountInUsdOk

`func (o *DealUnitDto) GetTotalGlobalDiscountsAmountInUsdOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountInUsdOk returns a tuple with the TotalGlobalDiscountsAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmountInUsd

`func (o *DealUnitDto) SetTotalGlobalDiscountsAmountInUsd(v Money)`

SetTotalGlobalDiscountsAmountInUsd sets TotalGlobalDiscountsAmountInUsd field to given value.

### HasTotalGlobalDiscountsAmountInUsd

`func (o *DealUnitDto) HasTotalGlobalDiscountsAmountInUsd() bool`

HasTotalGlobalDiscountsAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesAmountInUsd

`func (o *DealUnitDto) GetTotalGlobalSurchargesAmountInUsd() Money`

GetTotalGlobalSurchargesAmountInUsd returns the TotalGlobalSurchargesAmountInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountInUsdOk

`func (o *DealUnitDto) GetTotalGlobalSurchargesAmountInUsdOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountInUsdOk returns a tuple with the TotalGlobalSurchargesAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmountInUsd

`func (o *DealUnitDto) SetTotalGlobalSurchargesAmountInUsd(v Money)`

SetTotalGlobalSurchargesAmountInUsd sets TotalGlobalSurchargesAmountInUsd field to given value.

### HasTotalGlobalSurchargesAmountInUsd

`func (o *DealUnitDto) HasTotalGlobalSurchargesAmountInUsd() bool`

HasTotalGlobalSurchargesAmountInUsd returns a boolean if a field has been set.

### GetTotalAmount

`func (o *DealUnitDto) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *DealUnitDto) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *DealUnitDto) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *DealUnitDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalTaxAmount

`func (o *DealUnitDto) GetTotalTaxAmount() Money`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *DealUnitDto) GetTotalTaxAmountOk() (*Money, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *DealUnitDto) SetTotalTaxAmount(v Money)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *DealUnitDto) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### GetTotalTaxBaseAmount

`func (o *DealUnitDto) GetTotalTaxBaseAmount() Money`

GetTotalTaxBaseAmount returns the TotalTaxBaseAmount field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountOk

`func (o *DealUnitDto) GetTotalTaxBaseAmountOk() (*Money, bool)`

GetTotalTaxBaseAmountOk returns a tuple with the TotalTaxBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmount

`func (o *DealUnitDto) SetTotalTaxBaseAmount(v Money)`

SetTotalTaxBaseAmount sets TotalTaxBaseAmount field to given value.

### HasTotalTaxBaseAmount

`func (o *DealUnitDto) HasTotalTaxBaseAmount() bool`

HasTotalTaxBaseAmount returns a boolean if a field has been set.

### GetTotalDiscountsAmount

`func (o *DealUnitDto) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *DealUnitDto) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *DealUnitDto) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *DealUnitDto) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurchargesAmount

`func (o *DealUnitDto) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *DealUnitDto) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *DealUnitDto) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *DealUnitDto) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalGlobalDiscountsAmount

`func (o *DealUnitDto) GetTotalGlobalDiscountsAmount() Money`

GetTotalGlobalDiscountsAmount returns the TotalGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountOk

`func (o *DealUnitDto) GetTotalGlobalDiscountsAmountOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountOk returns a tuple with the TotalGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmount

`func (o *DealUnitDto) SetTotalGlobalDiscountsAmount(v Money)`

SetTotalGlobalDiscountsAmount sets TotalGlobalDiscountsAmount field to given value.

### HasTotalGlobalDiscountsAmount

`func (o *DealUnitDto) HasTotalGlobalDiscountsAmount() bool`

HasTotalGlobalDiscountsAmount returns a boolean if a field has been set.

### GetTotalGlobalSurchargesAmount

`func (o *DealUnitDto) GetTotalGlobalSurchargesAmount() Money`

GetTotalGlobalSurchargesAmount returns the TotalGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountOk

`func (o *DealUnitDto) GetTotalGlobalSurchargesAmountOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountOk returns a tuple with the TotalGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmount

`func (o *DealUnitDto) SetTotalGlobalSurchargesAmount(v Money)`

SetTotalGlobalSurchargesAmount sets TotalGlobalSurchargesAmount field to given value.

### HasTotalGlobalSurchargesAmount

`func (o *DealUnitDto) HasTotalGlobalSurchargesAmount() bool`

HasTotalGlobalSurchargesAmount returns a boolean if a field has been set.

### GetOrdered

`func (o *DealUnitDto) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *DealUnitDto) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *DealUnitDto) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.

### HasOrdered

`func (o *DealUnitDto) HasOrdered() bool`

HasOrdered returns a boolean if a field has been set.

### GetDealUnitFeedId

`func (o *DealUnitDto) GetDealUnitFeedId() string`

GetDealUnitFeedId returns the DealUnitFeedId field if non-nil, zero value otherwise.

### GetDealUnitFeedIdOk

`func (o *DealUnitDto) GetDealUnitFeedIdOk() (*string, bool)`

GetDealUnitFeedIdOk returns a tuple with the DealUnitFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFeedId

`func (o *DealUnitDto) SetDealUnitFeedId(v string)`

SetDealUnitFeedId sets DealUnitFeedId field to given value.

### HasDealUnitFeedId

`func (o *DealUnitDto) HasDealUnitFeedId() bool`

HasDealUnitFeedId returns a boolean if a field has been set.

### SetDealUnitFeedIdNil

`func (o *DealUnitDto) SetDealUnitFeedIdNil(b bool)`

 SetDealUnitFeedIdNil sets the value for DealUnitFeedId to be an explicit nil

### UnsetDealUnitFeedId
`func (o *DealUnitDto) UnsetDealUnitFeedId()`

UnsetDealUnitFeedId ensures that no value is present for DealUnitFeedId, not even an explicit nil
### GetDealUnitFlowId

`func (o *DealUnitDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetDealUnitFlowStageId

`func (o *DealUnitDto) GetDealUnitFlowStageId() string`

GetDealUnitFlowStageId returns the DealUnitFlowStageId field if non-nil, zero value otherwise.

### GetDealUnitFlowStageIdOk

`func (o *DealUnitDto) GetDealUnitFlowStageIdOk() (*string, bool)`

GetDealUnitFlowStageIdOk returns a tuple with the DealUnitFlowStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowStageId

`func (o *DealUnitDto) SetDealUnitFlowStageId(v string)`

SetDealUnitFlowStageId sets DealUnitFlowStageId field to given value.

### HasDealUnitFlowStageId

`func (o *DealUnitDto) HasDealUnitFlowStageId() bool`

HasDealUnitFlowStageId returns a boolean if a field has been set.

### SetDealUnitFlowStageIdNil

`func (o *DealUnitDto) SetDealUnitFlowStageIdNil(b bool)`

 SetDealUnitFlowStageIdNil sets the value for DealUnitFlowStageId to be an explicit nil

### UnsetDealUnitFlowStageId
`func (o *DealUnitDto) UnsetDealUnitFlowStageId()`

UnsetDealUnitFlowStageId ensures that no value is present for DealUnitFlowStageId, not even an explicit nil
### GetBillingLocationId

`func (o *DealUnitDto) GetBillingLocationId() string`

GetBillingLocationId returns the BillingLocationId field if non-nil, zero value otherwise.

### GetBillingLocationIdOk

`func (o *DealUnitDto) GetBillingLocationIdOk() (*string, bool)`

GetBillingLocationIdOk returns a tuple with the BillingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLocationId

`func (o *DealUnitDto) SetBillingLocationId(v string)`

SetBillingLocationId sets BillingLocationId field to given value.

### HasBillingLocationId

`func (o *DealUnitDto) HasBillingLocationId() bool`

HasBillingLocationId returns a boolean if a field has been set.

### SetBillingLocationIdNil

`func (o *DealUnitDto) SetBillingLocationIdNil(b bool)`

 SetBillingLocationIdNil sets the value for BillingLocationId to be an explicit nil

### UnsetBillingLocationId
`func (o *DealUnitDto) UnsetBillingLocationId()`

UnsetBillingLocationId ensures that no value is present for BillingLocationId, not even an explicit nil
### GetShippingLocationId

`func (o *DealUnitDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *DealUnitDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *DealUnitDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *DealUnitDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *DealUnitDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *DealUnitDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetPartnerCreated

`func (o *DealUnitDto) GetPartnerCreated() bool`

GetPartnerCreated returns the PartnerCreated field if non-nil, zero value otherwise.

### GetPartnerCreatedOk

`func (o *DealUnitDto) GetPartnerCreatedOk() (*bool, bool)`

GetPartnerCreatedOk returns a tuple with the PartnerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCreated

`func (o *DealUnitDto) SetPartnerCreated(v bool)`

SetPartnerCreated sets PartnerCreated field to given value.

### HasPartnerCreated

`func (o *DealUnitDto) HasPartnerCreated() bool`

HasPartnerCreated returns a boolean if a field has been set.

### GetPartnerCollaboration

`func (o *DealUnitDto) GetPartnerCollaboration() bool`

GetPartnerCollaboration returns the PartnerCollaboration field if non-nil, zero value otherwise.

### GetPartnerCollaborationOk

`func (o *DealUnitDto) GetPartnerCollaborationOk() (*bool, bool)`

GetPartnerCollaborationOk returns a tuple with the PartnerCollaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCollaboration

`func (o *DealUnitDto) SetPartnerCollaboration(v bool)`

SetPartnerCollaboration sets PartnerCollaboration field to given value.

### HasPartnerCollaboration

`func (o *DealUnitDto) HasPartnerCollaboration() bool`

HasPartnerCollaboration returns a boolean if a field has been set.

### GetProposedSolution

`func (o *DealUnitDto) GetProposedSolution() string`

GetProposedSolution returns the ProposedSolution field if non-nil, zero value otherwise.

### GetProposedSolutionOk

`func (o *DealUnitDto) GetProposedSolutionOk() (*string, bool)`

GetProposedSolutionOk returns a tuple with the ProposedSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedSolution

`func (o *DealUnitDto) SetProposedSolution(v string)`

SetProposedSolution sets ProposedSolution field to given value.

### HasProposedSolution

`func (o *DealUnitDto) HasProposedSolution() bool`

HasProposedSolution returns a boolean if a field has been set.

### SetProposedSolutionNil

`func (o *DealUnitDto) SetProposedSolutionNil(b bool)`

 SetProposedSolutionNil sets the value for ProposedSolution to be an explicit nil

### UnsetProposedSolution
`func (o *DealUnitDto) UnsetProposedSolution()`

UnsetProposedSolution ensures that no value is present for ProposedSolution, not even an explicit nil
### GetCurrentSituation

`func (o *DealUnitDto) GetCurrentSituation() string`

GetCurrentSituation returns the CurrentSituation field if non-nil, zero value otherwise.

### GetCurrentSituationOk

`func (o *DealUnitDto) GetCurrentSituationOk() (*string, bool)`

GetCurrentSituationOk returns a tuple with the CurrentSituation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSituation

`func (o *DealUnitDto) SetCurrentSituation(v string)`

SetCurrentSituation sets CurrentSituation field to given value.

### HasCurrentSituation

`func (o *DealUnitDto) HasCurrentSituation() bool`

HasCurrentSituation returns a boolean if a field has been set.

### SetCurrentSituationNil

`func (o *DealUnitDto) SetCurrentSituationNil(b bool)`

 SetCurrentSituationNil sets the value for CurrentSituation to be an explicit nil

### UnsetCurrentSituation
`func (o *DealUnitDto) UnsetCurrentSituation()`

UnsetCurrentSituation ensures that no value is present for CurrentSituation, not even an explicit nil
### GetCustomerNeed

`func (o *DealUnitDto) GetCustomerNeed() string`

GetCustomerNeed returns the CustomerNeed field if non-nil, zero value otherwise.

### GetCustomerNeedOk

`func (o *DealUnitDto) GetCustomerNeedOk() (*string, bool)`

GetCustomerNeedOk returns a tuple with the CustomerNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNeed

`func (o *DealUnitDto) SetCustomerNeed(v string)`

SetCustomerNeed sets CustomerNeed field to given value.

### HasCustomerNeed

`func (o *DealUnitDto) HasCustomerNeed() bool`

HasCustomerNeed returns a boolean if a field has been set.

### SetCustomerNeedNil

`func (o *DealUnitDto) SetCustomerNeedNil(b bool)`

 SetCustomerNeedNil sets the value for CustomerNeed to be an explicit nil

### UnsetCustomerNeed
`func (o *DealUnitDto) UnsetCustomerNeed()`

UnsetCustomerNeed ensures that no value is present for CustomerNeed, not even an explicit nil
### GetWonDate

`func (o *DealUnitDto) GetWonDate() time.Time`

GetWonDate returns the WonDate field if non-nil, zero value otherwise.

### GetWonDateOk

`func (o *DealUnitDto) GetWonDateOk() (*time.Time, bool)`

GetWonDateOk returns a tuple with the WonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWonDate

`func (o *DealUnitDto) SetWonDate(v time.Time)`

SetWonDate sets WonDate field to given value.

### HasWonDate

`func (o *DealUnitDto) HasWonDate() bool`

HasWonDate returns a boolean if a field has been set.

### GetLostDate

`func (o *DealUnitDto) GetLostDate() time.Time`

GetLostDate returns the LostDate field if non-nil, zero value otherwise.

### GetLostDateOk

`func (o *DealUnitDto) GetLostDateOk() (*time.Time, bool)`

GetLostDateOk returns a tuple with the LostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostDate

`func (o *DealUnitDto) SetLostDate(v time.Time)`

SetLostDate sets LostDate field to given value.

### HasLostDate

`func (o *DealUnitDto) HasLostDate() bool`

HasLostDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *DealUnitDto) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *DealUnitDto) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *DealUnitDto) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *DealUnitDto) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *DealUnitDto) GetDeliveredDate() time.Time`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *DealUnitDto) GetDeliveredDateOk() (*time.Time, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *DealUnitDto) SetDeliveredDate(v time.Time)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *DealUnitDto) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### GetClosedTimestamp

`func (o *DealUnitDto) GetClosedTimestamp() time.Time`

GetClosedTimestamp returns the ClosedTimestamp field if non-nil, zero value otherwise.

### GetClosedTimestampOk

`func (o *DealUnitDto) GetClosedTimestampOk() (*time.Time, bool)`

GetClosedTimestampOk returns a tuple with the ClosedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTimestamp

`func (o *DealUnitDto) SetClosedTimestamp(v time.Time)`

SetClosedTimestamp sets ClosedTimestamp field to given value.

### HasClosedTimestamp

`func (o *DealUnitDto) HasClosedTimestamp() bool`

HasClosedTimestamp returns a boolean if a field has been set.

### GetExpectedCloseDate

`func (o *DealUnitDto) GetExpectedCloseDate() time.Time`

GetExpectedCloseDate returns the ExpectedCloseDate field if non-nil, zero value otherwise.

### GetExpectedCloseDateOk

`func (o *DealUnitDto) GetExpectedCloseDateOk() (*time.Time, bool)`

GetExpectedCloseDateOk returns a tuple with the ExpectedCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCloseDate

`func (o *DealUnitDto) SetExpectedCloseDate(v time.Time)`

SetExpectedCloseDate sets ExpectedCloseDate field to given value.

### HasExpectedCloseDate

`func (o *DealUnitDto) HasExpectedCloseDate() bool`

HasExpectedCloseDate returns a boolean if a field has been set.

### GetDealUnitStatus

`func (o *DealUnitDto) GetDealUnitStatus() int32`

GetDealUnitStatus returns the DealUnitStatus field if non-nil, zero value otherwise.

### GetDealUnitStatusOk

`func (o *DealUnitDto) GetDealUnitStatusOk() (*int32, bool)`

GetDealUnitStatusOk returns a tuple with the DealUnitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitStatus

`func (o *DealUnitDto) SetDealUnitStatus(v int32)`

SetDealUnitStatus sets DealUnitStatus field to given value.

### HasDealUnitStatus

`func (o *DealUnitDto) HasDealUnitStatus() bool`

HasDealUnitStatus returns a boolean if a field has been set.

### GetDealUnitPurchaseProcess

`func (o *DealUnitDto) GetDealUnitPurchaseProcess() int32`

GetDealUnitPurchaseProcess returns the DealUnitPurchaseProcess field if non-nil, zero value otherwise.

### GetDealUnitPurchaseProcessOk

`func (o *DealUnitDto) GetDealUnitPurchaseProcessOk() (*int32, bool)`

GetDealUnitPurchaseProcessOk returns a tuple with the DealUnitPurchaseProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitPurchaseProcess

`func (o *DealUnitDto) SetDealUnitPurchaseProcess(v int32)`

SetDealUnitPurchaseProcess sets DealUnitPurchaseProcess field to given value.

### HasDealUnitPurchaseProcess

`func (o *DealUnitDto) HasDealUnitPurchaseProcess() bool`

HasDealUnitPurchaseProcess returns a boolean if a field has been set.

### GetDealUnitForecastCategory

`func (o *DealUnitDto) GetDealUnitForecastCategory() int32`

GetDealUnitForecastCategory returns the DealUnitForecastCategory field if non-nil, zero value otherwise.

### GetDealUnitForecastCategoryOk

`func (o *DealUnitDto) GetDealUnitForecastCategoryOk() (*int32, bool)`

GetDealUnitForecastCategoryOk returns a tuple with the DealUnitForecastCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitForecastCategory

`func (o *DealUnitDto) SetDealUnitForecastCategory(v int32)`

SetDealUnitForecastCategory sets DealUnitForecastCategory field to given value.

### HasDealUnitForecastCategory

`func (o *DealUnitDto) HasDealUnitForecastCategory() bool`

HasDealUnitForecastCategory returns a boolean if a field has been set.

### GetDealUnitAmountsCalculation

`func (o *DealUnitDto) GetDealUnitAmountsCalculation() int32`

GetDealUnitAmountsCalculation returns the DealUnitAmountsCalculation field if non-nil, zero value otherwise.

### GetDealUnitAmountsCalculationOk

`func (o *DealUnitDto) GetDealUnitAmountsCalculationOk() (*int32, bool)`

GetDealUnitAmountsCalculationOk returns a tuple with the DealUnitAmountsCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitAmountsCalculation

`func (o *DealUnitDto) SetDealUnitAmountsCalculation(v int32)`

SetDealUnitAmountsCalculation sets DealUnitAmountsCalculation field to given value.

### HasDealUnitAmountsCalculation

`func (o *DealUnitDto) HasDealUnitAmountsCalculation() bool`

HasDealUnitAmountsCalculation returns a boolean if a field has been set.

### GetLinesCount

`func (o *DealUnitDto) GetLinesCount() int32`

GetLinesCount returns the LinesCount field if non-nil, zero value otherwise.

### GetLinesCountOk

`func (o *DealUnitDto) GetLinesCountOk() (*int32, bool)`

GetLinesCountOk returns a tuple with the LinesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesCount

`func (o *DealUnitDto) SetLinesCount(v int32)`

SetLinesCount sets LinesCount field to given value.

### HasLinesCount

`func (o *DealUnitDto) HasLinesCount() bool`

HasLinesCount returns a boolean if a field has been set.

### GetCustomTotalAmount

`func (o *DealUnitDto) GetCustomTotalAmount() float64`

GetCustomTotalAmount returns the CustomTotalAmount field if non-nil, zero value otherwise.

### GetCustomTotalAmountOk

`func (o *DealUnitDto) GetCustomTotalAmountOk() (*float64, bool)`

GetCustomTotalAmountOk returns a tuple with the CustomTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTotalAmount

`func (o *DealUnitDto) SetCustomTotalAmount(v float64)`

SetCustomTotalAmount sets CustomTotalAmount field to given value.

### HasCustomTotalAmount

`func (o *DealUnitDto) HasCustomTotalAmount() bool`

HasCustomTotalAmount returns a boolean if a field has been set.

### GetCustomDetailAmount

`func (o *DealUnitDto) GetCustomDetailAmount() float64`

GetCustomDetailAmount returns the CustomDetailAmount field if non-nil, zero value otherwise.

### GetCustomDetailAmountOk

`func (o *DealUnitDto) GetCustomDetailAmountOk() (*float64, bool)`

GetCustomDetailAmountOk returns a tuple with the CustomDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDetailAmount

`func (o *DealUnitDto) SetCustomDetailAmount(v float64)`

SetCustomDetailAmount sets CustomDetailAmount field to given value.

### HasCustomDetailAmount

`func (o *DealUnitDto) HasCustomDetailAmount() bool`

HasCustomDetailAmount returns a boolean if a field has been set.

### GetCustomProfitAmount

`func (o *DealUnitDto) GetCustomProfitAmount() float64`

GetCustomProfitAmount returns the CustomProfitAmount field if non-nil, zero value otherwise.

### GetCustomProfitAmountOk

`func (o *DealUnitDto) GetCustomProfitAmountOk() (*float64, bool)`

GetCustomProfitAmountOk returns a tuple with the CustomProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfitAmount

`func (o *DealUnitDto) SetCustomProfitAmount(v float64)`

SetCustomProfitAmount sets CustomProfitAmount field to given value.

### HasCustomProfitAmount

`func (o *DealUnitDto) HasCustomProfitAmount() bool`

HasCustomProfitAmount returns a boolean if a field has been set.

### GetCustomShippingCostAmount

`func (o *DealUnitDto) GetCustomShippingCostAmount() float64`

GetCustomShippingCostAmount returns the CustomShippingCostAmount field if non-nil, zero value otherwise.

### GetCustomShippingCostAmountOk

`func (o *DealUnitDto) GetCustomShippingCostAmountOk() (*float64, bool)`

GetCustomShippingCostAmountOk returns a tuple with the CustomShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingCostAmount

`func (o *DealUnitDto) SetCustomShippingCostAmount(v float64)`

SetCustomShippingCostAmount sets CustomShippingCostAmount field to given value.

### HasCustomShippingCostAmount

`func (o *DealUnitDto) HasCustomShippingCostAmount() bool`

HasCustomShippingCostAmount returns a boolean if a field has been set.

### GetCustomWithholdingTaxAmount

`func (o *DealUnitDto) GetCustomWithholdingTaxAmount() float64`

GetCustomWithholdingTaxAmount returns the CustomWithholdingTaxAmount field if non-nil, zero value otherwise.

### GetCustomWithholdingTaxAmountOk

`func (o *DealUnitDto) GetCustomWithholdingTaxAmountOk() (*float64, bool)`

GetCustomWithholdingTaxAmountOk returns a tuple with the CustomWithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWithholdingTaxAmount

`func (o *DealUnitDto) SetCustomWithholdingTaxAmount(v float64)`

SetCustomWithholdingTaxAmount sets CustomWithholdingTaxAmount field to given value.

### HasCustomWithholdingTaxAmount

`func (o *DealUnitDto) HasCustomWithholdingTaxAmount() bool`

HasCustomWithholdingTaxAmount returns a boolean if a field has been set.

### GetCustomSurchargesAmount

`func (o *DealUnitDto) GetCustomSurchargesAmount() float64`

GetCustomSurchargesAmount returns the CustomSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomSurchargesAmountOk

`func (o *DealUnitDto) GetCustomSurchargesAmountOk() (*float64, bool)`

GetCustomSurchargesAmountOk returns a tuple with the CustomSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargesAmount

`func (o *DealUnitDto) SetCustomSurchargesAmount(v float64)`

SetCustomSurchargesAmount sets CustomSurchargesAmount field to given value.

### HasCustomSurchargesAmount

`func (o *DealUnitDto) HasCustomSurchargesAmount() bool`

HasCustomSurchargesAmount returns a boolean if a field has been set.

### GetCustomDiscountsAmount

`func (o *DealUnitDto) GetCustomDiscountsAmount() float64`

GetCustomDiscountsAmount returns the CustomDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomDiscountsAmountOk

`func (o *DealUnitDto) GetCustomDiscountsAmountOk() (*float64, bool)`

GetCustomDiscountsAmountOk returns a tuple with the CustomDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDiscountsAmount

`func (o *DealUnitDto) SetCustomDiscountsAmount(v float64)`

SetCustomDiscountsAmount sets CustomDiscountsAmount field to given value.

### HasCustomDiscountsAmount

`func (o *DealUnitDto) HasCustomDiscountsAmount() bool`

HasCustomDiscountsAmount returns a boolean if a field has been set.

### GetCustomShippingTaxAmount

`func (o *DealUnitDto) GetCustomShippingTaxAmount() float64`

GetCustomShippingTaxAmount returns the CustomShippingTaxAmount field if non-nil, zero value otherwise.

### GetCustomShippingTaxAmountOk

`func (o *DealUnitDto) GetCustomShippingTaxAmountOk() (*float64, bool)`

GetCustomShippingTaxAmountOk returns a tuple with the CustomShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingTaxAmount

`func (o *DealUnitDto) SetCustomShippingTaxAmount(v float64)`

SetCustomShippingTaxAmount sets CustomShippingTaxAmount field to given value.

### HasCustomShippingTaxAmount

`func (o *DealUnitDto) HasCustomShippingTaxAmount() bool`

HasCustomShippingTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


