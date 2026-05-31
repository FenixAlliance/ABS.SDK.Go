# PortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Address1** | Pointer to **NullableString** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**Address3** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**CustomCity** | Pointer to **NullableString** |  | [optional] 
**CustomState** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Fax** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CountryStateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**UnLocode** | Pointer to **NullableString** |  | [optional] 
**IataCode** | Pointer to **NullableString** |  | [optional] 
**PortType** | Pointer to **NullableString** |  | [optional] 
**PortAuthority** | Pointer to **NullableString** |  | [optional] 
**HasCustomsFacility** | Pointer to **bool** |  | [optional] 
**IsFreeTradezone** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**ParentPortId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPortDto

`func NewPortDto() *PortDto`

NewPortDto instantiates a new PortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDtoWithDefaults

`func NewPortDtoWithDefaults() *PortDto`

NewPortDtoWithDefaults instantiates a new PortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PortDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PortDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PortDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PortDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PortDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PortDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PortDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PortDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *PortDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PortDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PortDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PortDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PortDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PortDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCompany

`func (o *PortDto) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PortDto) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PortDto) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PortDto) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PortDto) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PortDto) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetEmail

`func (o *PortDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PortDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PortDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PortDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PortDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PortDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress1

`func (o *PortDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *PortDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *PortDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *PortDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *PortDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *PortDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *PortDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *PortDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *PortDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *PortDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *PortDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *PortDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *PortDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *PortDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *PortDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *PortDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *PortDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *PortDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetUnit

`func (o *PortDto) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PortDto) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PortDto) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PortDto) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *PortDto) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *PortDto) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetCustomCity

`func (o *PortDto) GetCustomCity() string`

GetCustomCity returns the CustomCity field if non-nil, zero value otherwise.

### GetCustomCityOk

`func (o *PortDto) GetCustomCityOk() (*string, bool)`

GetCustomCityOk returns a tuple with the CustomCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCity

`func (o *PortDto) SetCustomCity(v string)`

SetCustomCity sets CustomCity field to given value.

### HasCustomCity

`func (o *PortDto) HasCustomCity() bool`

HasCustomCity returns a boolean if a field has been set.

### SetCustomCityNil

`func (o *PortDto) SetCustomCityNil(b bool)`

 SetCustomCityNil sets the value for CustomCity to be an explicit nil

### UnsetCustomCity
`func (o *PortDto) UnsetCustomCity()`

UnsetCustomCity ensures that no value is present for CustomCity, not even an explicit nil
### GetCustomState

`func (o *PortDto) GetCustomState() string`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *PortDto) GetCustomStateOk() (*string, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *PortDto) SetCustomState(v string)`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *PortDto) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *PortDto) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *PortDto) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetPostalCode

`func (o *PortDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PortDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PortDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PortDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *PortDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *PortDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPhone

`func (o *PortDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PortDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PortDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PortDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PortDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PortDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *PortDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *PortDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *PortDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *PortDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *PortDto) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *PortDto) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetLongitude

`func (o *PortDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PortDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PortDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PortDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *PortDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PortDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PortDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PortDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetCountryId

`func (o *PortDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *PortDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *PortDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *PortDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *PortDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *PortDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryStateId

`func (o *PortDto) GetCountryStateId() string`

GetCountryStateId returns the CountryStateId field if non-nil, zero value otherwise.

### GetCountryStateIdOk

`func (o *PortDto) GetCountryStateIdOk() (*string, bool)`

GetCountryStateIdOk returns a tuple with the CountryStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryStateId

`func (o *PortDto) SetCountryStateId(v string)`

SetCountryStateId sets CountryStateId field to given value.

### HasCountryStateId

`func (o *PortDto) HasCountryStateId() bool`

HasCountryStateId returns a boolean if a field has been set.

### SetCountryStateIdNil

`func (o *PortDto) SetCountryStateIdNil(b bool)`

 SetCountryStateIdNil sets the value for CountryStateId to be an explicit nil

### UnsetCountryStateId
`func (o *PortDto) UnsetCountryStateId()`

UnsetCountryStateId ensures that no value is present for CountryStateId, not even an explicit nil
### GetCityId

`func (o *PortDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *PortDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *PortDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *PortDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *PortDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *PortDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetTenantId

`func (o *PortDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PortDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PortDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PortDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PortDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PortDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PortDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PortDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PortDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PortDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PortDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PortDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetUnLocode

`func (o *PortDto) GetUnLocode() string`

GetUnLocode returns the UnLocode field if non-nil, zero value otherwise.

### GetUnLocodeOk

`func (o *PortDto) GetUnLocodeOk() (*string, bool)`

GetUnLocodeOk returns a tuple with the UnLocode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnLocode

`func (o *PortDto) SetUnLocode(v string)`

SetUnLocode sets UnLocode field to given value.

### HasUnLocode

`func (o *PortDto) HasUnLocode() bool`

HasUnLocode returns a boolean if a field has been set.

### SetUnLocodeNil

`func (o *PortDto) SetUnLocodeNil(b bool)`

 SetUnLocodeNil sets the value for UnLocode to be an explicit nil

### UnsetUnLocode
`func (o *PortDto) UnsetUnLocode()`

UnsetUnLocode ensures that no value is present for UnLocode, not even an explicit nil
### GetIataCode

`func (o *PortDto) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *PortDto) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *PortDto) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *PortDto) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### SetIataCodeNil

`func (o *PortDto) SetIataCodeNil(b bool)`

 SetIataCodeNil sets the value for IataCode to be an explicit nil

### UnsetIataCode
`func (o *PortDto) UnsetIataCode()`

UnsetIataCode ensures that no value is present for IataCode, not even an explicit nil
### GetPortType

`func (o *PortDto) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *PortDto) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *PortDto) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *PortDto) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *PortDto) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *PortDto) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetPortAuthority

`func (o *PortDto) GetPortAuthority() string`

GetPortAuthority returns the PortAuthority field if non-nil, zero value otherwise.

### GetPortAuthorityOk

`func (o *PortDto) GetPortAuthorityOk() (*string, bool)`

GetPortAuthorityOk returns a tuple with the PortAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAuthority

`func (o *PortDto) SetPortAuthority(v string)`

SetPortAuthority sets PortAuthority field to given value.

### HasPortAuthority

`func (o *PortDto) HasPortAuthority() bool`

HasPortAuthority returns a boolean if a field has been set.

### SetPortAuthorityNil

`func (o *PortDto) SetPortAuthorityNil(b bool)`

 SetPortAuthorityNil sets the value for PortAuthority to be an explicit nil

### UnsetPortAuthority
`func (o *PortDto) UnsetPortAuthority()`

UnsetPortAuthority ensures that no value is present for PortAuthority, not even an explicit nil
### GetHasCustomsFacility

`func (o *PortDto) GetHasCustomsFacility() bool`

GetHasCustomsFacility returns the HasCustomsFacility field if non-nil, zero value otherwise.

### GetHasCustomsFacilityOk

`func (o *PortDto) GetHasCustomsFacilityOk() (*bool, bool)`

GetHasCustomsFacilityOk returns a tuple with the HasCustomsFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomsFacility

`func (o *PortDto) SetHasCustomsFacility(v bool)`

SetHasCustomsFacility sets HasCustomsFacility field to given value.

### HasHasCustomsFacility

`func (o *PortDto) HasHasCustomsFacility() bool`

HasHasCustomsFacility returns a boolean if a field has been set.

### GetIsFreeTradezone

`func (o *PortDto) GetIsFreeTradezone() bool`

GetIsFreeTradezone returns the IsFreeTradezone field if non-nil, zero value otherwise.

### GetIsFreeTradezoneOk

`func (o *PortDto) GetIsFreeTradezoneOk() (*bool, bool)`

GetIsFreeTradezoneOk returns a tuple with the IsFreeTradezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeTradezone

`func (o *PortDto) SetIsFreeTradezone(v bool)`

SetIsFreeTradezone sets IsFreeTradezone field to given value.

### HasIsFreeTradezone

`func (o *PortDto) HasIsFreeTradezone() bool`

HasIsFreeTradezone returns a boolean if a field has been set.

### GetIsActive

`func (o *PortDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PortDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PortDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PortDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentPortId

`func (o *PortDto) GetParentPortId() string`

GetParentPortId returns the ParentPortId field if non-nil, zero value otherwise.

### GetParentPortIdOk

`func (o *PortDto) GetParentPortIdOk() (*string, bool)`

GetParentPortIdOk returns a tuple with the ParentPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPortId

`func (o *PortDto) SetParentPortId(v string)`

SetParentPortId sets ParentPortId field to given value.

### HasParentPortId

`func (o *PortDto) HasParentPortId() bool`

HasParentPortId returns a boolean if a field has been set.

### SetParentPortIdNil

`func (o *PortDto) SetParentPortIdNil(b bool)`

 SetParentPortIdNil sets the value for ParentPortId to be an explicit nil

### UnsetParentPortId
`func (o *PortDto) UnsetParentPortId()`

UnsetParentPortId ensures that no value is present for ParentPortId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


