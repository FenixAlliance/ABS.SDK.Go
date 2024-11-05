# SecurityPermissionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsSystemPermission** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityPermissionDto

`func NewSecurityPermissionDto() *SecurityPermissionDto`

NewSecurityPermissionDto instantiates a new SecurityPermissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPermissionDtoWithDefaults

`func NewSecurityPermissionDtoWithDefaults() *SecurityPermissionDto`

NewSecurityPermissionDtoWithDefaults instantiates a new SecurityPermissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityPermissionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityPermissionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityPermissionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityPermissionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SecurityPermissionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SecurityPermissionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SecurityPermissionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SecurityPermissionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SecurityPermissionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SecurityPermissionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SecurityPermissionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SecurityPermissionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *SecurityPermissionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityPermissionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityPermissionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityPermissionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityPermissionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityPermissionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *SecurityPermissionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SecurityPermissionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SecurityPermissionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SecurityPermissionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SecurityPermissionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SecurityPermissionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *SecurityPermissionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityPermissionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityPermissionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityPermissionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityPermissionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityPermissionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsSystemPermission

`func (o *SecurityPermissionDto) GetIsSystemPermission() bool`

GetIsSystemPermission returns the IsSystemPermission field if non-nil, zero value otherwise.

### GetIsSystemPermissionOk

`func (o *SecurityPermissionDto) GetIsSystemPermissionOk() (*bool, bool)`

GetIsSystemPermissionOk returns a tuple with the IsSystemPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemPermission

`func (o *SecurityPermissionDto) SetIsSystemPermission(v bool)`

SetIsSystemPermission sets IsSystemPermission field to given value.

### HasIsSystemPermission

`func (o *SecurityPermissionDto) HasIsSystemPermission() bool`

HasIsSystemPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


