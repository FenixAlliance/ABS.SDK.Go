# CostCentreUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostCentreType** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CostCentresGroupId** | Pointer to **NullableString** |  | [optional] 
**ParentCostCentreId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreUpdateDto

`func NewCostCentreUpdateDto() *CostCentreUpdateDto`

NewCostCentreUpdateDto instantiates a new CostCentreUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreUpdateDtoWithDefaults

`func NewCostCentreUpdateDtoWithDefaults() *CostCentreUpdateDto`

NewCostCentreUpdateDtoWithDefaults instantiates a new CostCentreUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CostCentreUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisabled

`func (o *CostCentreUpdateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CostCentreUpdateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CostCentreUpdateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CostCentreUpdateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *CostCentreUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CostCentreUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CostCentreUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CostCentreUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CostCentreUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CostCentreUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostCentreType

`func (o *CostCentreUpdateDto) GetCostCentreType() string`

GetCostCentreType returns the CostCentreType field if non-nil, zero value otherwise.

### GetCostCentreTypeOk

`func (o *CostCentreUpdateDto) GetCostCentreTypeOk() (*string, bool)`

GetCostCentreTypeOk returns a tuple with the CostCentreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreType

`func (o *CostCentreUpdateDto) SetCostCentreType(v string)`

SetCostCentreType sets CostCentreType field to given value.

### HasCostCentreType

`func (o *CostCentreUpdateDto) HasCostCentreType() bool`

HasCostCentreType returns a boolean if a field has been set.

### GetTenantId

`func (o *CostCentreUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCostCentresGroupId

`func (o *CostCentreUpdateDto) GetCostCentresGroupId() string`

GetCostCentresGroupId returns the CostCentresGroupId field if non-nil, zero value otherwise.

### GetCostCentresGroupIdOk

`func (o *CostCentreUpdateDto) GetCostCentresGroupIdOk() (*string, bool)`

GetCostCentresGroupIdOk returns a tuple with the CostCentresGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentresGroupId

`func (o *CostCentreUpdateDto) SetCostCentresGroupId(v string)`

SetCostCentresGroupId sets CostCentresGroupId field to given value.

### HasCostCentresGroupId

`func (o *CostCentreUpdateDto) HasCostCentresGroupId() bool`

HasCostCentresGroupId returns a boolean if a field has been set.

### SetCostCentresGroupIdNil

`func (o *CostCentreUpdateDto) SetCostCentresGroupIdNil(b bool)`

 SetCostCentresGroupIdNil sets the value for CostCentresGroupId to be an explicit nil

### UnsetCostCentresGroupId
`func (o *CostCentreUpdateDto) UnsetCostCentresGroupId()`

UnsetCostCentresGroupId ensures that no value is present for CostCentresGroupId, not even an explicit nil
### GetParentCostCentreId

`func (o *CostCentreUpdateDto) GetParentCostCentreId() string`

GetParentCostCentreId returns the ParentCostCentreId field if non-nil, zero value otherwise.

### GetParentCostCentreIdOk

`func (o *CostCentreUpdateDto) GetParentCostCentreIdOk() (*string, bool)`

GetParentCostCentreIdOk returns a tuple with the ParentCostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCostCentreId

`func (o *CostCentreUpdateDto) SetParentCostCentreId(v string)`

SetParentCostCentreId sets ParentCostCentreId field to given value.

### HasParentCostCentreId

`func (o *CostCentreUpdateDto) HasParentCostCentreId() bool`

HasParentCostCentreId returns a boolean if a field has been set.

### SetParentCostCentreIdNil

`func (o *CostCentreUpdateDto) SetParentCostCentreIdNil(b bool)`

 SetParentCostCentreIdNil sets the value for ParentCostCentreId to be an explicit nil

### UnsetParentCostCentreId
`func (o *CostCentreUpdateDto) UnsetParentCostCentreId()`

UnsetParentCostCentreId ensures that no value is present for ParentCostCentreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


