# TenantDepartmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentDepartmentID** | Pointer to **NullableString** |  | [optional] 

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

### GetOrganizationProfileID

`func (o *TenantDepartmentCreateDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantDepartmentCreateDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantDepartmentCreateDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantDepartmentCreateDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantDepartmentCreateDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantDepartmentCreateDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil
### GetParentDepartmentID

`func (o *TenantDepartmentCreateDto) GetParentDepartmentID() string`

GetParentDepartmentID returns the ParentDepartmentID field if non-nil, zero value otherwise.

### GetParentDepartmentIDOk

`func (o *TenantDepartmentCreateDto) GetParentDepartmentIDOk() (*string, bool)`

GetParentDepartmentIDOk returns a tuple with the ParentDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDepartmentID

`func (o *TenantDepartmentCreateDto) SetParentDepartmentID(v string)`

SetParentDepartmentID sets ParentDepartmentID field to given value.

### HasParentDepartmentID

`func (o *TenantDepartmentCreateDto) HasParentDepartmentID() bool`

HasParentDepartmentID returns a boolean if a field has been set.

### SetParentDepartmentIDNil

`func (o *TenantDepartmentCreateDto) SetParentDepartmentIDNil(b bool)`

 SetParentDepartmentIDNil sets the value for ParentDepartmentID to be an explicit nil

### UnsetParentDepartmentID
`func (o *TenantDepartmentCreateDto) UnsetParentDepartmentID()`

UnsetParentDepartmentID ensures that no value is present for ParentDepartmentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


