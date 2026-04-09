# TenantUnitUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**BusinessUnitQualifiedName** | Pointer to **NullableString** |  | [optional] 
**CountryID** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessUnitID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantUnitUpdateDto

`func NewTenantUnitUpdateDto() *TenantUnitUpdateDto`

NewTenantUnitUpdateDto instantiates a new TenantUnitUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUnitUpdateDtoWithDefaults

`func NewTenantUnitUpdateDtoWithDefaults() *TenantUnitUpdateDto`

NewTenantUnitUpdateDtoWithDefaults instantiates a new TenantUnitUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TenantUnitUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantUnitUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantUnitUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantUnitUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantUnitUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantUnitUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantUnitUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantUnitUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantUnitUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantUnitUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantUnitUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantUnitUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantUnitUpdateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantUnitUpdateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantUnitUpdateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantUnitUpdateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetBusinessUnitQualifiedName

`func (o *TenantUnitUpdateDto) GetBusinessUnitQualifiedName() string`

GetBusinessUnitQualifiedName returns the BusinessUnitQualifiedName field if non-nil, zero value otherwise.

### GetBusinessUnitQualifiedNameOk

`func (o *TenantUnitUpdateDto) GetBusinessUnitQualifiedNameOk() (*string, bool)`

GetBusinessUnitQualifiedNameOk returns a tuple with the BusinessUnitQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitQualifiedName

`func (o *TenantUnitUpdateDto) SetBusinessUnitQualifiedName(v string)`

SetBusinessUnitQualifiedName sets BusinessUnitQualifiedName field to given value.

### HasBusinessUnitQualifiedName

`func (o *TenantUnitUpdateDto) HasBusinessUnitQualifiedName() bool`

HasBusinessUnitQualifiedName returns a boolean if a field has been set.

### SetBusinessUnitQualifiedNameNil

`func (o *TenantUnitUpdateDto) SetBusinessUnitQualifiedNameNil(b bool)`

 SetBusinessUnitQualifiedNameNil sets the value for BusinessUnitQualifiedName to be an explicit nil

### UnsetBusinessUnitQualifiedName
`func (o *TenantUnitUpdateDto) UnsetBusinessUnitQualifiedName()`

UnsetBusinessUnitQualifiedName ensures that no value is present for BusinessUnitQualifiedName, not even an explicit nil
### GetCountryID

`func (o *TenantUnitUpdateDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *TenantUnitUpdateDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *TenantUnitUpdateDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *TenantUnitUpdateDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *TenantUnitUpdateDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *TenantUnitUpdateDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetOrganizationProfileID

`func (o *TenantUnitUpdateDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantUnitUpdateDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantUnitUpdateDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantUnitUpdateDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantUnitUpdateDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantUnitUpdateDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil
### GetParentBusinessUnitID

`func (o *TenantUnitUpdateDto) GetParentBusinessUnitID() string`

GetParentBusinessUnitID returns the ParentBusinessUnitID field if non-nil, zero value otherwise.

### GetParentBusinessUnitIDOk

`func (o *TenantUnitUpdateDto) GetParentBusinessUnitIDOk() (*string, bool)`

GetParentBusinessUnitIDOk returns a tuple with the ParentBusinessUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessUnitID

`func (o *TenantUnitUpdateDto) SetParentBusinessUnitID(v string)`

SetParentBusinessUnitID sets ParentBusinessUnitID field to given value.

### HasParentBusinessUnitID

`func (o *TenantUnitUpdateDto) HasParentBusinessUnitID() bool`

HasParentBusinessUnitID returns a boolean if a field has been set.

### SetParentBusinessUnitIDNil

`func (o *TenantUnitUpdateDto) SetParentBusinessUnitIDNil(b bool)`

 SetParentBusinessUnitIDNil sets the value for ParentBusinessUnitID to be an explicit nil

### UnsetParentBusinessUnitID
`func (o *TenantUnitUpdateDto) UnsetParentBusinessUnitID()`

UnsetParentBusinessUnitID ensures that no value is present for ParentBusinessUnitID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


