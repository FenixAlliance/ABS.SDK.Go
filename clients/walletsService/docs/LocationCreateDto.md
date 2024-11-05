# LocationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewLocationCreateDto

`func NewLocationCreateDto() *LocationCreateDto`

NewLocationCreateDto instantiates a new LocationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCreateDtoWithDefaults

`func NewLocationCreateDtoWithDefaults() *LocationCreateDto`

NewLocationCreateDtoWithDefaults instantiates a new LocationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LocationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LocationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *LocationCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LocationCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LocationCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LocationCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LocationCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LocationCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetEmail

`func (o *LocationCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationCreateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocationCreateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocationCreateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *LocationCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LocationCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LocationCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LocationCreateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *LocationCreateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *LocationCreateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *LocationCreateDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *LocationCreateDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *LocationCreateDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *LocationCreateDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *LocationCreateDto) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *LocationCreateDto) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetAddress1

`func (o *LocationCreateDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationCreateDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationCreateDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationCreateDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *LocationCreateDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *LocationCreateDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *LocationCreateDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationCreateDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationCreateDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationCreateDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *LocationCreateDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *LocationCreateDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *LocationCreateDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *LocationCreateDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *LocationCreateDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *LocationCreateDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *LocationCreateDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *LocationCreateDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetUnit

`func (o *LocationCreateDto) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LocationCreateDto) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LocationCreateDto) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LocationCreateDto) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *LocationCreateDto) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *LocationCreateDto) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetCityId

`func (o *LocationCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *LocationCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *LocationCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *LocationCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *LocationCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *LocationCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *LocationCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *LocationCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *LocationCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *LocationCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *LocationCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *LocationCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetPostalCode

`func (o *LocationCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LocationCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LocationCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LocationCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *LocationCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *LocationCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *LocationCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *LocationCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *LocationCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *LocationCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *LocationCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *LocationCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTenantId

`func (o *LocationCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LocationCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LocationCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LocationCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LocationCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LocationCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLongitude

`func (o *LocationCreateDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationCreateDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationCreateDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocationCreateDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationCreateDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationCreateDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationCreateDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocationCreateDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetIsRoutable

`func (o *LocationCreateDto) GetIsRoutable() bool`

GetIsRoutable returns the IsRoutable field if non-nil, zero value otherwise.

### GetIsRoutableOk

`func (o *LocationCreateDto) GetIsRoutableOk() (*bool, bool)`

GetIsRoutableOk returns a tuple with the IsRoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoutable

`func (o *LocationCreateDto) SetIsRoutable(v bool)`

SetIsRoutable sets IsRoutable field to given value.

### HasIsRoutable

`func (o *LocationCreateDto) HasIsRoutable() bool`

HasIsRoutable returns a boolean if a field has been set.

### GetIsGlobalPrimary

`func (o *LocationCreateDto) GetIsGlobalPrimary() bool`

GetIsGlobalPrimary returns the IsGlobalPrimary field if non-nil, zero value otherwise.

### GetIsGlobalPrimaryOk

`func (o *LocationCreateDto) GetIsGlobalPrimaryOk() (*bool, bool)`

GetIsGlobalPrimaryOk returns a tuple with the IsGlobalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalPrimary

`func (o *LocationCreateDto) SetIsGlobalPrimary(v bool)`

SetIsGlobalPrimary sets IsGlobalPrimary field to given value.

### HasIsGlobalPrimary

`func (o *LocationCreateDto) HasIsGlobalPrimary() bool`

HasIsGlobalPrimary returns a boolean if a field has been set.

### GetIsCountryPrimary

`func (o *LocationCreateDto) GetIsCountryPrimary() bool`

GetIsCountryPrimary returns the IsCountryPrimary field if non-nil, zero value otherwise.

### GetIsCountryPrimaryOk

`func (o *LocationCreateDto) GetIsCountryPrimaryOk() (*bool, bool)`

GetIsCountryPrimaryOk returns a tuple with the IsCountryPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCountryPrimary

`func (o *LocationCreateDto) SetIsCountryPrimary(v bool)`

SetIsCountryPrimary sets IsCountryPrimary field to given value.

### HasIsCountryPrimary

`func (o *LocationCreateDto) HasIsCountryPrimary() bool`

HasIsCountryPrimary returns a boolean if a field has been set.

### GetCanGenerateLabels

`func (o *LocationCreateDto) GetCanGenerateLabels() bool`

GetCanGenerateLabels returns the CanGenerateLabels field if non-nil, zero value otherwise.

### GetCanGenerateLabelsOk

`func (o *LocationCreateDto) GetCanGenerateLabelsOk() (*bool, bool)`

GetCanGenerateLabelsOk returns a tuple with the CanGenerateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGenerateLabels

`func (o *LocationCreateDto) SetCanGenerateLabels(v bool)`

SetCanGenerateLabels sets CanGenerateLabels field to given value.

### HasCanGenerateLabels

`func (o *LocationCreateDto) HasCanGenerateLabels() bool`

HasCanGenerateLabels returns a boolean if a field has been set.

### GetIsDefaultSenderAddress

`func (o *LocationCreateDto) GetIsDefaultSenderAddress() bool`

GetIsDefaultSenderAddress returns the IsDefaultSenderAddress field if non-nil, zero value otherwise.

### GetIsDefaultSenderAddressOk

`func (o *LocationCreateDto) GetIsDefaultSenderAddressOk() (*bool, bool)`

GetIsDefaultSenderAddressOk returns a tuple with the IsDefaultSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSenderAddress

`func (o *LocationCreateDto) SetIsDefaultSenderAddress(v bool)`

SetIsDefaultSenderAddress sets IsDefaultSenderAddress field to given value.

### HasIsDefaultSenderAddress

`func (o *LocationCreateDto) HasIsDefaultSenderAddress() bool`

HasIsDefaultSenderAddress returns a boolean if a field has been set.

### GetIsDefaultReturnAddress

`func (o *LocationCreateDto) GetIsDefaultReturnAddress() bool`

GetIsDefaultReturnAddress returns the IsDefaultReturnAddress field if non-nil, zero value otherwise.

### GetIsDefaultReturnAddressOk

`func (o *LocationCreateDto) GetIsDefaultReturnAddressOk() (*bool, bool)`

GetIsDefaultReturnAddressOk returns a tuple with the IsDefaultReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultReturnAddress

`func (o *LocationCreateDto) SetIsDefaultReturnAddress(v bool)`

SetIsDefaultReturnAddress sets IsDefaultReturnAddress field to given value.

### HasIsDefaultReturnAddress

`func (o *LocationCreateDto) HasIsDefaultReturnAddress() bool`

HasIsDefaultReturnAddress returns a boolean if a field has been set.

### GetIsDefaultSuppingLocation

`func (o *LocationCreateDto) GetIsDefaultSuppingLocation() bool`

GetIsDefaultSuppingLocation returns the IsDefaultSuppingLocation field if non-nil, zero value otherwise.

### GetIsDefaultSuppingLocationOk

`func (o *LocationCreateDto) GetIsDefaultSuppingLocationOk() (*bool, bool)`

GetIsDefaultSuppingLocationOk returns a tuple with the IsDefaultSuppingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSuppingLocation

`func (o *LocationCreateDto) SetIsDefaultSuppingLocation(v bool)`

SetIsDefaultSuppingLocation sets IsDefaultSuppingLocation field to given value.

### HasIsDefaultSuppingLocation

`func (o *LocationCreateDto) HasIsDefaultSuppingLocation() bool`

HasIsDefaultSuppingLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


