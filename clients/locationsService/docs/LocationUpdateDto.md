# LocationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Fax** | Pointer to **NullableString** |  | [optional] 
**Address1** | Pointer to **NullableString** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**Address3** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**IsRoutable** | Pointer to **bool** |  | [optional] 
**IsGlobalPrimary** | Pointer to **bool** |  | [optional] 
**IsCountryPrimary** | Pointer to **bool** |  | [optional] 
**CanGenerateLabels** | Pointer to **bool** |  | [optional] 
**IsDefaultSenderAddress** | Pointer to **bool** |  | [optional] 
**IsDefaultReturnAddress** | Pointer to **bool** |  | [optional] 
**IsDefaultSuppingLocation** | Pointer to **bool** |  | [optional] 

## Methods

### NewLocationUpdateDto

`func NewLocationUpdateDto() *LocationUpdateDto`

NewLocationUpdateDto instantiates a new LocationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationUpdateDtoWithDefaults

`func NewLocationUpdateDtoWithDefaults() *LocationUpdateDto`

NewLocationUpdateDtoWithDefaults instantiates a new LocationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LocationUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LocationUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LocationUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LocationUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LocationUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LocationUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetEmail

`func (o *LocationUpdateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationUpdateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationUpdateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationUpdateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocationUpdateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocationUpdateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *LocationUpdateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LocationUpdateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LocationUpdateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LocationUpdateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *LocationUpdateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *LocationUpdateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *LocationUpdateDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *LocationUpdateDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *LocationUpdateDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *LocationUpdateDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *LocationUpdateDto) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *LocationUpdateDto) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetAddress1

`func (o *LocationUpdateDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationUpdateDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationUpdateDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationUpdateDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *LocationUpdateDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *LocationUpdateDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *LocationUpdateDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationUpdateDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationUpdateDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationUpdateDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *LocationUpdateDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *LocationUpdateDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *LocationUpdateDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *LocationUpdateDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *LocationUpdateDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *LocationUpdateDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *LocationUpdateDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *LocationUpdateDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetUnit

`func (o *LocationUpdateDto) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LocationUpdateDto) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LocationUpdateDto) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LocationUpdateDto) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *LocationUpdateDto) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *LocationUpdateDto) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetCityId

`func (o *LocationUpdateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *LocationUpdateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *LocationUpdateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *LocationUpdateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *LocationUpdateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *LocationUpdateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *LocationUpdateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *LocationUpdateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *LocationUpdateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *LocationUpdateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *LocationUpdateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *LocationUpdateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetPostalCode

`func (o *LocationUpdateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LocationUpdateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LocationUpdateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LocationUpdateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *LocationUpdateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *LocationUpdateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *LocationUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *LocationUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *LocationUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *LocationUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *LocationUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *LocationUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTenantId

`func (o *LocationUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LocationUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LocationUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LocationUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LocationUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LocationUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLongitude

`func (o *LocationUpdateDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationUpdateDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationUpdateDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocationUpdateDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationUpdateDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationUpdateDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationUpdateDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocationUpdateDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetIsRoutable

`func (o *LocationUpdateDto) GetIsRoutable() bool`

GetIsRoutable returns the IsRoutable field if non-nil, zero value otherwise.

### GetIsRoutableOk

`func (o *LocationUpdateDto) GetIsRoutableOk() (*bool, bool)`

GetIsRoutableOk returns a tuple with the IsRoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoutable

`func (o *LocationUpdateDto) SetIsRoutable(v bool)`

SetIsRoutable sets IsRoutable field to given value.

### HasIsRoutable

`func (o *LocationUpdateDto) HasIsRoutable() bool`

HasIsRoutable returns a boolean if a field has been set.

### GetIsGlobalPrimary

`func (o *LocationUpdateDto) GetIsGlobalPrimary() bool`

GetIsGlobalPrimary returns the IsGlobalPrimary field if non-nil, zero value otherwise.

### GetIsGlobalPrimaryOk

`func (o *LocationUpdateDto) GetIsGlobalPrimaryOk() (*bool, bool)`

GetIsGlobalPrimaryOk returns a tuple with the IsGlobalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalPrimary

`func (o *LocationUpdateDto) SetIsGlobalPrimary(v bool)`

SetIsGlobalPrimary sets IsGlobalPrimary field to given value.

### HasIsGlobalPrimary

`func (o *LocationUpdateDto) HasIsGlobalPrimary() bool`

HasIsGlobalPrimary returns a boolean if a field has been set.

### GetIsCountryPrimary

`func (o *LocationUpdateDto) GetIsCountryPrimary() bool`

GetIsCountryPrimary returns the IsCountryPrimary field if non-nil, zero value otherwise.

### GetIsCountryPrimaryOk

`func (o *LocationUpdateDto) GetIsCountryPrimaryOk() (*bool, bool)`

GetIsCountryPrimaryOk returns a tuple with the IsCountryPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCountryPrimary

`func (o *LocationUpdateDto) SetIsCountryPrimary(v bool)`

SetIsCountryPrimary sets IsCountryPrimary field to given value.

### HasIsCountryPrimary

`func (o *LocationUpdateDto) HasIsCountryPrimary() bool`

HasIsCountryPrimary returns a boolean if a field has been set.

### GetCanGenerateLabels

`func (o *LocationUpdateDto) GetCanGenerateLabels() bool`

GetCanGenerateLabels returns the CanGenerateLabels field if non-nil, zero value otherwise.

### GetCanGenerateLabelsOk

`func (o *LocationUpdateDto) GetCanGenerateLabelsOk() (*bool, bool)`

GetCanGenerateLabelsOk returns a tuple with the CanGenerateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGenerateLabels

`func (o *LocationUpdateDto) SetCanGenerateLabels(v bool)`

SetCanGenerateLabels sets CanGenerateLabels field to given value.

### HasCanGenerateLabels

`func (o *LocationUpdateDto) HasCanGenerateLabels() bool`

HasCanGenerateLabels returns a boolean if a field has been set.

### GetIsDefaultSenderAddress

`func (o *LocationUpdateDto) GetIsDefaultSenderAddress() bool`

GetIsDefaultSenderAddress returns the IsDefaultSenderAddress field if non-nil, zero value otherwise.

### GetIsDefaultSenderAddressOk

`func (o *LocationUpdateDto) GetIsDefaultSenderAddressOk() (*bool, bool)`

GetIsDefaultSenderAddressOk returns a tuple with the IsDefaultSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSenderAddress

`func (o *LocationUpdateDto) SetIsDefaultSenderAddress(v bool)`

SetIsDefaultSenderAddress sets IsDefaultSenderAddress field to given value.

### HasIsDefaultSenderAddress

`func (o *LocationUpdateDto) HasIsDefaultSenderAddress() bool`

HasIsDefaultSenderAddress returns a boolean if a field has been set.

### GetIsDefaultReturnAddress

`func (o *LocationUpdateDto) GetIsDefaultReturnAddress() bool`

GetIsDefaultReturnAddress returns the IsDefaultReturnAddress field if non-nil, zero value otherwise.

### GetIsDefaultReturnAddressOk

`func (o *LocationUpdateDto) GetIsDefaultReturnAddressOk() (*bool, bool)`

GetIsDefaultReturnAddressOk returns a tuple with the IsDefaultReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultReturnAddress

`func (o *LocationUpdateDto) SetIsDefaultReturnAddress(v bool)`

SetIsDefaultReturnAddress sets IsDefaultReturnAddress field to given value.

### HasIsDefaultReturnAddress

`func (o *LocationUpdateDto) HasIsDefaultReturnAddress() bool`

HasIsDefaultReturnAddress returns a boolean if a field has been set.

### GetIsDefaultSuppingLocation

`func (o *LocationUpdateDto) GetIsDefaultSuppingLocation() bool`

GetIsDefaultSuppingLocation returns the IsDefaultSuppingLocation field if non-nil, zero value otherwise.

### GetIsDefaultSuppingLocationOk

`func (o *LocationUpdateDto) GetIsDefaultSuppingLocationOk() (*bool, bool)`

GetIsDefaultSuppingLocationOk returns a tuple with the IsDefaultSuppingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSuppingLocation

`func (o *LocationUpdateDto) SetIsDefaultSuppingLocation(v bool)`

SetIsDefaultSuppingLocation sets IsDefaultSuppingLocation field to given value.

### HasIsDefaultSuppingLocation

`func (o *LocationUpdateDto) HasIsDefaultSuppingLocation() bool`

HasIsDefaultSuppingLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


