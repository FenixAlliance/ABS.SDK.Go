# BillingProfileCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TaxId** | **string** |  | 
**Phone** | **string** |  | 
**Email** | **string** |  | 
**Address** | **string** |  | 
**Address1** | Pointer to **NullableString** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | **string** |  | 
**BusinessName** | **string** |  | 
**CommercialName** | **string** |  | 
**Ticker** | Pointer to **NullableString** |  | [optional] 
**Duns** | Pointer to **NullableString** |  | [optional] 
**IsPublicCompany** | Pointer to **bool** |  | [optional] 
**IsFactaCustomer** | Pointer to **bool** |  | [optional] 
**CountryId** | **string** |  | 
**StateId** | **string** |  | 
**CityId** | **string** |  | 
**FiscalIdentificationTypeId** | **string** |  | 
**FiscalAuthorityId** | **string** |  | 
**FiscalRegimeId** | **string** |  | 

## Methods

### NewBillingProfileCreateDto

`func NewBillingProfileCreateDto(taxId string, phone string, email string, address string, postalCode string, businessName string, commercialName string, countryId string, stateId string, cityId string, fiscalIdentificationTypeId string, fiscalAuthorityId string, fiscalRegimeId string, ) *BillingProfileCreateDto`

NewBillingProfileCreateDto instantiates a new BillingProfileCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileCreateDtoWithDefaults

`func NewBillingProfileCreateDtoWithDefaults() *BillingProfileCreateDto`

NewBillingProfileCreateDtoWithDefaults instantiates a new BillingProfileCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingProfileCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingProfileCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingProfileCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingProfileCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BillingProfileCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillingProfileCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillingProfileCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillingProfileCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetContactId

`func (o *BillingProfileCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *BillingProfileCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *BillingProfileCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *BillingProfileCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *BillingProfileCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *BillingProfileCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTaxId

`func (o *BillingProfileCreateDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingProfileCreateDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingProfileCreateDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### GetPhone

`func (o *BillingProfileCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingProfileCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingProfileCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *BillingProfileCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingProfileCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingProfileCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddress

`func (o *BillingProfileCreateDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingProfileCreateDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingProfileCreateDto) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddress1

`func (o *BillingProfileCreateDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *BillingProfileCreateDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *BillingProfileCreateDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *BillingProfileCreateDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *BillingProfileCreateDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *BillingProfileCreateDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *BillingProfileCreateDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *BillingProfileCreateDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *BillingProfileCreateDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *BillingProfileCreateDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *BillingProfileCreateDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *BillingProfileCreateDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetPostalCode

`func (o *BillingProfileCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingProfileCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingProfileCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetBusinessName

`func (o *BillingProfileCreateDto) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *BillingProfileCreateDto) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *BillingProfileCreateDto) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetCommercialName

`func (o *BillingProfileCreateDto) GetCommercialName() string`

GetCommercialName returns the CommercialName field if non-nil, zero value otherwise.

### GetCommercialNameOk

`func (o *BillingProfileCreateDto) GetCommercialNameOk() (*string, bool)`

GetCommercialNameOk returns a tuple with the CommercialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialName

`func (o *BillingProfileCreateDto) SetCommercialName(v string)`

SetCommercialName sets CommercialName field to given value.


### GetTicker

`func (o *BillingProfileCreateDto) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *BillingProfileCreateDto) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *BillingProfileCreateDto) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *BillingProfileCreateDto) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *BillingProfileCreateDto) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *BillingProfileCreateDto) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetDuns

`func (o *BillingProfileCreateDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *BillingProfileCreateDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *BillingProfileCreateDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *BillingProfileCreateDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *BillingProfileCreateDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *BillingProfileCreateDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetIsPublicCompany

`func (o *BillingProfileCreateDto) GetIsPublicCompany() bool`

GetIsPublicCompany returns the IsPublicCompany field if non-nil, zero value otherwise.

### GetIsPublicCompanyOk

`func (o *BillingProfileCreateDto) GetIsPublicCompanyOk() (*bool, bool)`

GetIsPublicCompanyOk returns a tuple with the IsPublicCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicCompany

`func (o *BillingProfileCreateDto) SetIsPublicCompany(v bool)`

SetIsPublicCompany sets IsPublicCompany field to given value.

### HasIsPublicCompany

`func (o *BillingProfileCreateDto) HasIsPublicCompany() bool`

HasIsPublicCompany returns a boolean if a field has been set.

### GetIsFactaCustomer

`func (o *BillingProfileCreateDto) GetIsFactaCustomer() bool`

GetIsFactaCustomer returns the IsFactaCustomer field if non-nil, zero value otherwise.

### GetIsFactaCustomerOk

`func (o *BillingProfileCreateDto) GetIsFactaCustomerOk() (*bool, bool)`

GetIsFactaCustomerOk returns a tuple with the IsFactaCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactaCustomer

`func (o *BillingProfileCreateDto) SetIsFactaCustomer(v bool)`

SetIsFactaCustomer sets IsFactaCustomer field to given value.

### HasIsFactaCustomer

`func (o *BillingProfileCreateDto) HasIsFactaCustomer() bool`

HasIsFactaCustomer returns a boolean if a field has been set.

### GetCountryId

`func (o *BillingProfileCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *BillingProfileCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *BillingProfileCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.


### GetStateId

`func (o *BillingProfileCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *BillingProfileCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *BillingProfileCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.


### GetCityId

`func (o *BillingProfileCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *BillingProfileCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *BillingProfileCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.


### GetFiscalIdentificationTypeId

`func (o *BillingProfileCreateDto) GetFiscalIdentificationTypeId() string`

GetFiscalIdentificationTypeId returns the FiscalIdentificationTypeId field if non-nil, zero value otherwise.

### GetFiscalIdentificationTypeIdOk

`func (o *BillingProfileCreateDto) GetFiscalIdentificationTypeIdOk() (*string, bool)`

GetFiscalIdentificationTypeIdOk returns a tuple with the FiscalIdentificationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalIdentificationTypeId

`func (o *BillingProfileCreateDto) SetFiscalIdentificationTypeId(v string)`

SetFiscalIdentificationTypeId sets FiscalIdentificationTypeId field to given value.


### GetFiscalAuthorityId

`func (o *BillingProfileCreateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *BillingProfileCreateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *BillingProfileCreateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.


### GetFiscalRegimeId

`func (o *BillingProfileCreateDto) GetFiscalRegimeId() string`

GetFiscalRegimeId returns the FiscalRegimeId field if non-nil, zero value otherwise.

### GetFiscalRegimeIdOk

`func (o *BillingProfileCreateDto) GetFiscalRegimeIdOk() (*string, bool)`

GetFiscalRegimeIdOk returns a tuple with the FiscalRegimeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalRegimeId

`func (o *BillingProfileCreateDto) SetFiscalRegimeId(v string)`

SetFiscalRegimeId sets FiscalRegimeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


