# TenantDepartmentDto

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
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentDepartmentID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantDepartmentDto

`func NewTenantDepartmentDto() *TenantDepartmentDto`

NewTenantDepartmentDto instantiates a new TenantDepartmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDepartmentDtoWithDefaults

`func NewTenantDepartmentDtoWithDefaults() *TenantDepartmentDto`

NewTenantDepartmentDtoWithDefaults instantiates a new TenantDepartmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantDepartmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantDepartmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantDepartmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantDepartmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantDepartmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantDepartmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantDepartmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantDepartmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantDepartmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantDepartmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantDepartmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantDepartmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBusinessID

`func (o *TenantDepartmentDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantDepartmentDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantDepartmentDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *TenantDepartmentDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *TenantDepartmentDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *TenantDepartmentDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantDepartmentDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantDepartmentDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantDepartmentDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantDepartmentDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantDepartmentDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantDepartmentDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetName

`func (o *TenantDepartmentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantDepartmentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantDepartmentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantDepartmentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantDepartmentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantDepartmentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantDepartmentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantDepartmentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantDepartmentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantDepartmentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantDepartmentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantDepartmentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantDepartmentDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantDepartmentDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantDepartmentDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantDepartmentDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOrganizationProfileID

`func (o *TenantDepartmentDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantDepartmentDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantDepartmentDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantDepartmentDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantDepartmentDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantDepartmentDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil
### GetParentDepartmentID

`func (o *TenantDepartmentDto) GetParentDepartmentID() string`

GetParentDepartmentID returns the ParentDepartmentID field if non-nil, zero value otherwise.

### GetParentDepartmentIDOk

`func (o *TenantDepartmentDto) GetParentDepartmentIDOk() (*string, bool)`

GetParentDepartmentIDOk returns a tuple with the ParentDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDepartmentID

`func (o *TenantDepartmentDto) SetParentDepartmentID(v string)`

SetParentDepartmentID sets ParentDepartmentID field to given value.

### HasParentDepartmentID

`func (o *TenantDepartmentDto) HasParentDepartmentID() bool`

HasParentDepartmentID returns a boolean if a field has been set.

### SetParentDepartmentIDNil

`func (o *TenantDepartmentDto) SetParentDepartmentIDNil(b bool)`

 SetParentDepartmentIDNil sets the value for ParentDepartmentID to be an explicit nil

### UnsetParentDepartmentID
`func (o *TenantDepartmentDto) UnsetParentDepartmentID()`

UnsetParentDepartmentID ensures that no value is present for ParentDepartmentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


