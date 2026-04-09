# TenantUnitCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**BusinessUnitQualifiedName** | Pointer to **NullableString** |  | [optional] 
**CountryID** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessUnitID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantUnitCreateDto

`func NewTenantUnitCreateDto() *TenantUnitCreateDto`

NewTenantUnitCreateDto instantiates a new TenantUnitCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUnitCreateDtoWithDefaults

`func NewTenantUnitCreateDtoWithDefaults() *TenantUnitCreateDto`

NewTenantUnitCreateDtoWithDefaults instantiates a new TenantUnitCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantUnitCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantUnitCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantUnitCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantUnitCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantUnitCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantUnitCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantUnitCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantUnitCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessID

`func (o *TenantUnitCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantUnitCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantUnitCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *TenantUnitCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *TenantUnitCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *TenantUnitCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantUnitCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantUnitCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantUnitCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantUnitCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantUnitCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantUnitCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetName

`func (o *TenantUnitCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantUnitCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantUnitCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantUnitCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantUnitCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantUnitCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantUnitCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantUnitCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantUnitCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantUnitCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantUnitCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantUnitCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantUnitCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantUnitCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantUnitCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantUnitCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetBusinessUnitQualifiedName

`func (o *TenantUnitCreateDto) GetBusinessUnitQualifiedName() string`

GetBusinessUnitQualifiedName returns the BusinessUnitQualifiedName field if non-nil, zero value otherwise.

### GetBusinessUnitQualifiedNameOk

`func (o *TenantUnitCreateDto) GetBusinessUnitQualifiedNameOk() (*string, bool)`

GetBusinessUnitQualifiedNameOk returns a tuple with the BusinessUnitQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitQualifiedName

`func (o *TenantUnitCreateDto) SetBusinessUnitQualifiedName(v string)`

SetBusinessUnitQualifiedName sets BusinessUnitQualifiedName field to given value.

### HasBusinessUnitQualifiedName

`func (o *TenantUnitCreateDto) HasBusinessUnitQualifiedName() bool`

HasBusinessUnitQualifiedName returns a boolean if a field has been set.

### SetBusinessUnitQualifiedNameNil

`func (o *TenantUnitCreateDto) SetBusinessUnitQualifiedNameNil(b bool)`

 SetBusinessUnitQualifiedNameNil sets the value for BusinessUnitQualifiedName to be an explicit nil

### UnsetBusinessUnitQualifiedName
`func (o *TenantUnitCreateDto) UnsetBusinessUnitQualifiedName()`

UnsetBusinessUnitQualifiedName ensures that no value is present for BusinessUnitQualifiedName, not even an explicit nil
### GetCountryID

`func (o *TenantUnitCreateDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *TenantUnitCreateDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *TenantUnitCreateDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *TenantUnitCreateDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *TenantUnitCreateDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *TenantUnitCreateDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetOrganizationProfileID

`func (o *TenantUnitCreateDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantUnitCreateDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantUnitCreateDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantUnitCreateDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantUnitCreateDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantUnitCreateDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil
### GetParentBusinessUnitID

`func (o *TenantUnitCreateDto) GetParentBusinessUnitID() string`

GetParentBusinessUnitID returns the ParentBusinessUnitID field if non-nil, zero value otherwise.

### GetParentBusinessUnitIDOk

`func (o *TenantUnitCreateDto) GetParentBusinessUnitIDOk() (*string, bool)`

GetParentBusinessUnitIDOk returns a tuple with the ParentBusinessUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessUnitID

`func (o *TenantUnitCreateDto) SetParentBusinessUnitID(v string)`

SetParentBusinessUnitID sets ParentBusinessUnitID field to given value.

### HasParentBusinessUnitID

`func (o *TenantUnitCreateDto) HasParentBusinessUnitID() bool`

HasParentBusinessUnitID returns a boolean if a field has been set.

### SetParentBusinessUnitIDNil

`func (o *TenantUnitCreateDto) SetParentBusinessUnitIDNil(b bool)`

 SetParentBusinessUnitIDNil sets the value for ParentBusinessUnitID to be an explicit nil

### UnsetParentBusinessUnitID
`func (o *TenantUnitCreateDto) UnsetParentBusinessUnitID()`

UnsetParentBusinessUnitID ensures that no value is present for ParentBusinessUnitID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


