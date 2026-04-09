# TenantUnitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
### GetBusinessID

`func (o *TenantUnitDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantUnitDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantUnitDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *TenantUnitDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *TenantUnitDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *TenantUnitDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantUnitDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantUnitDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantUnitDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantUnitDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantUnitDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantUnitDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
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
### GetCountryID

`func (o *TenantUnitDto) GetCountryID() string`

GetCountryID returns the CountryID field if non-nil, zero value otherwise.

### GetCountryIDOk

`func (o *TenantUnitDto) GetCountryIDOk() (*string, bool)`

GetCountryIDOk returns a tuple with the CountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryID

`func (o *TenantUnitDto) SetCountryID(v string)`

SetCountryID sets CountryID field to given value.

### HasCountryID

`func (o *TenantUnitDto) HasCountryID() bool`

HasCountryID returns a boolean if a field has been set.

### SetCountryIDNil

`func (o *TenantUnitDto) SetCountryIDNil(b bool)`

 SetCountryIDNil sets the value for CountryID to be an explicit nil

### UnsetCountryID
`func (o *TenantUnitDto) UnsetCountryID()`

UnsetCountryID ensures that no value is present for CountryID, not even an explicit nil
### GetOrganizationProfileID

`func (o *TenantUnitDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantUnitDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantUnitDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantUnitDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantUnitDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantUnitDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil
### GetParentBusinessUnitID

`func (o *TenantUnitDto) GetParentBusinessUnitID() string`

GetParentBusinessUnitID returns the ParentBusinessUnitID field if non-nil, zero value otherwise.

### GetParentBusinessUnitIDOk

`func (o *TenantUnitDto) GetParentBusinessUnitIDOk() (*string, bool)`

GetParentBusinessUnitIDOk returns a tuple with the ParentBusinessUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessUnitID

`func (o *TenantUnitDto) SetParentBusinessUnitID(v string)`

SetParentBusinessUnitID sets ParentBusinessUnitID field to given value.

### HasParentBusinessUnitID

`func (o *TenantUnitDto) HasParentBusinessUnitID() bool`

HasParentBusinessUnitID returns a boolean if a field has been set.

### SetParentBusinessUnitIDNil

`func (o *TenantUnitDto) SetParentBusinessUnitIDNil(b bool)`

 SetParentBusinessUnitIDNil sets the value for ParentBusinessUnitID to be an explicit nil

### UnsetParentBusinessUnitID
`func (o *TenantUnitDto) UnsetParentBusinessUnitID()`

UnsetParentBusinessUnitID ensures that no value is present for ParentBusinessUnitID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


