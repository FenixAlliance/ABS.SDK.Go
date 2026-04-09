# CostCentreGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ParentCostCentresGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreGroupDto

`func NewCostCentreGroupDto() *CostCentreGroupDto`

NewCostCentreGroupDto instantiates a new CostCentreGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreGroupDtoWithDefaults

`func NewCostCentreGroupDtoWithDefaults() *CostCentreGroupDto`

NewCostCentreGroupDtoWithDefaults instantiates a new CostCentreGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreGroupDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreGroupDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreGroupDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreGroupDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CostCentreGroupDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CostCentreGroupDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CostCentreGroupDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreGroupDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreGroupDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreGroupDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CostCentreGroupDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CostCentreGroupDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CostCentreGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreGroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreGroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CostCentreGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CostCentreGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CostCentreGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CostCentreGroupDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CostCentreGroupDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CostCentreGroupDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *CostCentreGroupDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CostCentreGroupDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CostCentreGroupDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CostCentreGroupDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTenantId

`func (o *CostCentreGroupDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreGroupDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreGroupDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreGroupDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreGroupDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreGroupDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetParentCostCentresGroupId

`func (o *CostCentreGroupDto) GetParentCostCentresGroupId() string`

GetParentCostCentresGroupId returns the ParentCostCentresGroupId field if non-nil, zero value otherwise.

### GetParentCostCentresGroupIdOk

`func (o *CostCentreGroupDto) GetParentCostCentresGroupIdOk() (*string, bool)`

GetParentCostCentresGroupIdOk returns a tuple with the ParentCostCentresGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCostCentresGroupId

`func (o *CostCentreGroupDto) SetParentCostCentresGroupId(v string)`

SetParentCostCentresGroupId sets ParentCostCentresGroupId field to given value.

### HasParentCostCentresGroupId

`func (o *CostCentreGroupDto) HasParentCostCentresGroupId() bool`

HasParentCostCentresGroupId returns a boolean if a field has been set.

### SetParentCostCentresGroupIdNil

`func (o *CostCentreGroupDto) SetParentCostCentresGroupIdNil(b bool)`

 SetParentCostCentresGroupIdNil sets the value for ParentCostCentresGroupId to be an explicit nil

### UnsetParentCostCentresGroupId
`func (o *CostCentreGroupDto) UnsetParentCostCentresGroupId()`

UnsetParentCostCentresGroupId ensures that no value is present for ParentCostCentresGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


