# TenantUnitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**BusinessUnitQualifiedName** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessUnitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantUnitDto

`func NewTenantUnitDto() *TenantUnitDto`

NewTenantUnitDto instantiates a new TenantUnitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUnitDtoWithDefaults

`func NewTenantUnitDtoWithDefaults() *TenantUnitDto`

NewTenantUnitDtoWithDefaults instantiates a new TenantUnitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantUnitDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantUnitDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantUnitDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantUnitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantUnitDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantUnitDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantUnitDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantUnitDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantUnitDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantUnitDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantUnitDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantUnitDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *TenantUnitDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantUnitDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantUnitDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantUnitDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantUnitDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantUnitDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantUnitDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantUnitDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantUnitDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantUnitDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantUnitDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantUnitDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetName

`func (o *TenantUnitDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantUnitDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantUnitDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantUnitDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantUnitDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantUnitDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantUnitDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantUnitDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantUnitDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantUnitDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantUnitDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantUnitDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantUnitDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantUnitDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantUnitDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantUnitDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetBusinessUnitQualifiedName

`func (o *TenantUnitDto) GetBusinessUnitQualifiedName() string`

GetBusinessUnitQualifiedName returns the BusinessUnitQualifiedName field if non-nil, zero value otherwise.

### GetBusinessUnitQualifiedNameOk

`func (o *TenantUnitDto) GetBusinessUnitQualifiedNameOk() (*string, bool)`

GetBusinessUnitQualifiedNameOk returns a tuple with the BusinessUnitQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitQualifiedName

`func (o *TenantUnitDto) SetBusinessUnitQualifiedName(v string)`

SetBusinessUnitQualifiedName sets BusinessUnitQualifiedName field to given value.

### HasBusinessUnitQualifiedName

`func (o *TenantUnitDto) HasBusinessUnitQualifiedName() bool`

HasBusinessUnitQualifiedName returns a boolean if a field has been set.

### SetBusinessUnitQualifiedNameNil

`func (o *TenantUnitDto) SetBusinessUnitQualifiedNameNil(b bool)`

 SetBusinessUnitQualifiedNameNil sets the value for BusinessUnitQualifiedName to be an explicit nil

### UnsetBusinessUnitQualifiedName
`func (o *TenantUnitDto) UnsetBusinessUnitQualifiedName()`

UnsetBusinessUnitQualifiedName ensures that no value is present for BusinessUnitQualifiedName, not even an explicit nil
### GetCountryId

`func (o *TenantUnitDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TenantUnitDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TenantUnitDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TenantUnitDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TenantUnitDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TenantUnitDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetOrganizationProfileId

`func (o *TenantUnitDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantUnitDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantUnitDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantUnitDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantUnitDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantUnitDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil
### GetParentBusinessUnitId

`func (o *TenantUnitDto) GetParentBusinessUnitId() string`

GetParentBusinessUnitId returns the ParentBusinessUnitId field if non-nil, zero value otherwise.

### GetParentBusinessUnitIdOk

`func (o *TenantUnitDto) GetParentBusinessUnitIdOk() (*string, bool)`

GetParentBusinessUnitIdOk returns a tuple with the ParentBusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessUnitId

`func (o *TenantUnitDto) SetParentBusinessUnitId(v string)`

SetParentBusinessUnitId sets ParentBusinessUnitId field to given value.

### HasParentBusinessUnitId

`func (o *TenantUnitDto) HasParentBusinessUnitId() bool`

HasParentBusinessUnitId returns a boolean if a field has been set.

### SetParentBusinessUnitIdNil

`func (o *TenantUnitDto) SetParentBusinessUnitIdNil(b bool)`

 SetParentBusinessUnitIdNil sets the value for ParentBusinessUnitId to be an explicit nil

### UnsetParentBusinessUnitId
`func (o *TenantUnitDto) UnsetParentBusinessUnitId()`

UnsetParentBusinessUnitId ensures that no value is present for ParentBusinessUnitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


