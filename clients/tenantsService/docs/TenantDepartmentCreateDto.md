# TenantDepartmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantDepartmentCreateDto

`func NewTenantDepartmentCreateDto() *TenantDepartmentCreateDto`

NewTenantDepartmentCreateDto instantiates a new TenantDepartmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDepartmentCreateDtoWithDefaults

`func NewTenantDepartmentCreateDtoWithDefaults() *TenantDepartmentCreateDto`

NewTenantDepartmentCreateDtoWithDefaults instantiates a new TenantDepartmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantDepartmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantDepartmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantDepartmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantDepartmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantDepartmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantDepartmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantDepartmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantDepartmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TenantDepartmentCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantDepartmentCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantDepartmentCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantDepartmentCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantDepartmentCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantDepartmentCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantDepartmentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantDepartmentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantDepartmentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantDepartmentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantDepartmentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantDepartmentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantDepartmentCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantDepartmentCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantDepartmentCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantDepartmentCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOrganizationProfileId

`func (o *TenantDepartmentCreateDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantDepartmentCreateDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantDepartmentCreateDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantDepartmentCreateDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantDepartmentCreateDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantDepartmentCreateDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil
### GetParentDepartmentId

`func (o *TenantDepartmentCreateDto) GetParentDepartmentId() string`

GetParentDepartmentId returns the ParentDepartmentId field if non-nil, zero value otherwise.

### GetParentDepartmentIdOk

`func (o *TenantDepartmentCreateDto) GetParentDepartmentIdOk() (*string, bool)`

GetParentDepartmentIdOk returns a tuple with the ParentDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDepartmentId

`func (o *TenantDepartmentCreateDto) SetParentDepartmentId(v string)`

SetParentDepartmentId sets ParentDepartmentId field to given value.

### HasParentDepartmentId

`func (o *TenantDepartmentCreateDto) HasParentDepartmentId() bool`

HasParentDepartmentId returns a boolean if a field has been set.

### SetParentDepartmentIdNil

`func (o *TenantDepartmentCreateDto) SetParentDepartmentIdNil(b bool)`

 SetParentDepartmentIdNil sets the value for ParentDepartmentId to be an explicit nil

### UnsetParentDepartmentId
`func (o *TenantDepartmentCreateDto) UnsetParentDepartmentId()`

UnsetParentDepartmentId ensures that no value is present for ParentDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


