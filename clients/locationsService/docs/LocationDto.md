# LocationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Business** | Pointer to **NullableString** |  | [optional] 
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

### NewLocationDto

`func NewLocationDto() *LocationDto`

NewLocationDto instantiates a new LocationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationDtoWithDefaults

`func NewLocationDtoWithDefaults() *LocationDto`

NewLocationDtoWithDefaults instantiates a new LocationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LocationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LocationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LocationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LocationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LocationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LocationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LocationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LocationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *LocationDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LocationDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LocationDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LocationDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LocationDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LocationDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetBusiness

`func (o *LocationDto) GetBusiness() string`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *LocationDto) GetBusinessOk() (*string, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *LocationDto) SetBusiness(v string)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *LocationDto) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### SetBusinessNil

`func (o *LocationDto) SetBusinessNil(b bool)`

 SetBusinessNil sets the value for Business to be an explicit nil

### UnsetBusiness
`func (o *LocationDto) UnsetBusiness()`

UnsetBusiness ensures that no value is present for Business, not even an explicit nil
### GetEmail

`func (o *LocationDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocationDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocationDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *LocationDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LocationDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LocationDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LocationDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *LocationDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *LocationDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *LocationDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *LocationDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *LocationDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *LocationDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *LocationDto) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *LocationDto) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetAddress1

`func (o *LocationDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *LocationDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *LocationDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *LocationDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *LocationDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *LocationDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *LocationDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *LocationDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *LocationDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *LocationDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *LocationDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *LocationDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetUnit

`func (o *LocationDto) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LocationDto) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LocationDto) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LocationDto) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *LocationDto) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *LocationDto) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetCityId

`func (o *LocationDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *LocationDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *LocationDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *LocationDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *LocationDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *LocationDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *LocationDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *LocationDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *LocationDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *LocationDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *LocationDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *LocationDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetPostalCode

`func (o *LocationDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LocationDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LocationDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LocationDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *LocationDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *LocationDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryId

`func (o *LocationDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *LocationDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *LocationDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *LocationDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *LocationDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *LocationDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetLongitude

`func (o *LocationDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocationDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocationDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetIsRoutable

`func (o *LocationDto) GetIsRoutable() bool`

GetIsRoutable returns the IsRoutable field if non-nil, zero value otherwise.

### GetIsRoutableOk

`func (o *LocationDto) GetIsRoutableOk() (*bool, bool)`

GetIsRoutableOk returns a tuple with the IsRoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoutable

`func (o *LocationDto) SetIsRoutable(v bool)`

SetIsRoutable sets IsRoutable field to given value.

### HasIsRoutable

`func (o *LocationDto) HasIsRoutable() bool`

HasIsRoutable returns a boolean if a field has been set.

### GetIsGlobalPrimary

`func (o *LocationDto) GetIsGlobalPrimary() bool`

GetIsGlobalPrimary returns the IsGlobalPrimary field if non-nil, zero value otherwise.

### GetIsGlobalPrimaryOk

`func (o *LocationDto) GetIsGlobalPrimaryOk() (*bool, bool)`

GetIsGlobalPrimaryOk returns a tuple with the IsGlobalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalPrimary

`func (o *LocationDto) SetIsGlobalPrimary(v bool)`

SetIsGlobalPrimary sets IsGlobalPrimary field to given value.

### HasIsGlobalPrimary

`func (o *LocationDto) HasIsGlobalPrimary() bool`

HasIsGlobalPrimary returns a boolean if a field has been set.

### GetIsCountryPrimary

`func (o *LocationDto) GetIsCountryPrimary() bool`

GetIsCountryPrimary returns the IsCountryPrimary field if non-nil, zero value otherwise.

### GetIsCountryPrimaryOk

`func (o *LocationDto) GetIsCountryPrimaryOk() (*bool, bool)`

GetIsCountryPrimaryOk returns a tuple with the IsCountryPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCountryPrimary

`func (o *LocationDto) SetIsCountryPrimary(v bool)`

SetIsCountryPrimary sets IsCountryPrimary field to given value.

### HasIsCountryPrimary

`func (o *LocationDto) HasIsCountryPrimary() bool`

HasIsCountryPrimary returns a boolean if a field has been set.

### GetCanGenerateLabels

`func (o *LocationDto) GetCanGenerateLabels() bool`

GetCanGenerateLabels returns the CanGenerateLabels field if non-nil, zero value otherwise.

### GetCanGenerateLabelsOk

`func (o *LocationDto) GetCanGenerateLabelsOk() (*bool, bool)`

GetCanGenerateLabelsOk returns a tuple with the CanGenerateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGenerateLabels

`func (o *LocationDto) SetCanGenerateLabels(v bool)`

SetCanGenerateLabels sets CanGenerateLabels field to given value.

### HasCanGenerateLabels

`func (o *LocationDto) HasCanGenerateLabels() bool`

HasCanGenerateLabels returns a boolean if a field has been set.

### GetIsDefaultSenderAddress

`func (o *LocationDto) GetIsDefaultSenderAddress() bool`

GetIsDefaultSenderAddress returns the IsDefaultSenderAddress field if non-nil, zero value otherwise.

### GetIsDefaultSenderAddressOk

`func (o *LocationDto) GetIsDefaultSenderAddressOk() (*bool, bool)`

GetIsDefaultSenderAddressOk returns a tuple with the IsDefaultSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSenderAddress

`func (o *LocationDto) SetIsDefaultSenderAddress(v bool)`

SetIsDefaultSenderAddress sets IsDefaultSenderAddress field to given value.

### HasIsDefaultSenderAddress

`func (o *LocationDto) HasIsDefaultSenderAddress() bool`

HasIsDefaultSenderAddress returns a boolean if a field has been set.

### GetIsDefaultReturnAddress

`func (o *LocationDto) GetIsDefaultReturnAddress() bool`

GetIsDefaultReturnAddress returns the IsDefaultReturnAddress field if non-nil, zero value otherwise.

### GetIsDefaultReturnAddressOk

`func (o *LocationDto) GetIsDefaultReturnAddressOk() (*bool, bool)`

GetIsDefaultReturnAddressOk returns a tuple with the IsDefaultReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultReturnAddress

`func (o *LocationDto) SetIsDefaultReturnAddress(v bool)`

SetIsDefaultReturnAddress sets IsDefaultReturnAddress field to given value.

### HasIsDefaultReturnAddress

`func (o *LocationDto) HasIsDefaultReturnAddress() bool`

HasIsDefaultReturnAddress returns a boolean if a field has been set.

### GetIsDefaultSuppingLocation

`func (o *LocationDto) GetIsDefaultSuppingLocation() bool`

GetIsDefaultSuppingLocation returns the IsDefaultSuppingLocation field if non-nil, zero value otherwise.

### GetIsDefaultSuppingLocationOk

`func (o *LocationDto) GetIsDefaultSuppingLocationOk() (*bool, bool)`

GetIsDefaultSuppingLocationOk returns a tuple with the IsDefaultSuppingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSuppingLocation

`func (o *LocationDto) SetIsDefaultSuppingLocation(v bool)`

SetIsDefaultSuppingLocation sets IsDefaultSuppingLocation field to given value.

### HasIsDefaultSuppingLocation

`func (o *LocationDto) HasIsDefaultSuppingLocation() bool`

HasIsDefaultSuppingLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


