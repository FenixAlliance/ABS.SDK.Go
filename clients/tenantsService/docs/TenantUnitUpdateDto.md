# TenantUnitUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessUnitId** | Pointer to **NullableString** |  | [optional] 

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

### GetCountryId

`func (o *TenantUnitUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TenantUnitUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TenantUnitUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *TenantUnitUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *TenantUnitUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *TenantUnitUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetOrganizationProfileId

`func (o *TenantUnitUpdateDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantUnitUpdateDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantUnitUpdateDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantUnitUpdateDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantUnitUpdateDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantUnitUpdateDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil
### GetParentBusinessUnitId

`func (o *TenantUnitUpdateDto) GetParentBusinessUnitId() string`

GetParentBusinessUnitId returns the ParentBusinessUnitId field if non-nil, zero value otherwise.

### GetParentBusinessUnitIdOk

`func (o *TenantUnitUpdateDto) GetParentBusinessUnitIdOk() (*string, bool)`

GetParentBusinessUnitIdOk returns a tuple with the ParentBusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessUnitId

`func (o *TenantUnitUpdateDto) SetParentBusinessUnitId(v string)`

SetParentBusinessUnitId sets ParentBusinessUnitId field to given value.

### HasParentBusinessUnitId

`func (o *TenantUnitUpdateDto) HasParentBusinessUnitId() bool`

HasParentBusinessUnitId returns a boolean if a field has been set.

### SetParentBusinessUnitIdNil

`func (o *TenantUnitUpdateDto) SetParentBusinessUnitIdNil(b bool)`

 SetParentBusinessUnitIdNil sets the value for ParentBusinessUnitId to be an explicit nil

### UnsetParentBusinessUnitId
`func (o *TenantUnitUpdateDto) UnsetParentBusinessUnitId()`

UnsetParentBusinessUnitId ensures that no value is present for ParentBusinessUnitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


