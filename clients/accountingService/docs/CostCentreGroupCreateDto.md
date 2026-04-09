# CostCentreGroupCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ParentCostCentresGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreGroupCreateDto

`func NewCostCentreGroupCreateDto() *CostCentreGroupCreateDto`

NewCostCentreGroupCreateDto instantiates a new CostCentreGroupCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreGroupCreateDtoWithDefaults

`func NewCostCentreGroupCreateDtoWithDefaults() *CostCentreGroupCreateDto`

NewCostCentreGroupCreateDtoWithDefaults instantiates a new CostCentreGroupCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreGroupCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreGroupCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreGroupCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreGroupCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CostCentreGroupCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreGroupCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreGroupCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreGroupCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CostCentreGroupCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreGroupCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreGroupCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreGroupCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreGroupCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreGroupCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CostCentreGroupCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CostCentreGroupCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CostCentreGroupCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CostCentreGroupCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CostCentreGroupCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CostCentreGroupCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *CostCentreGroupCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CostCentreGroupCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CostCentreGroupCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CostCentreGroupCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTenantId

`func (o *CostCentreGroupCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreGroupCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreGroupCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreGroupCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreGroupCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreGroupCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetParentCostCentresGroupId

`func (o *CostCentreGroupCreateDto) GetParentCostCentresGroupId() string`

GetParentCostCentresGroupId returns the ParentCostCentresGroupId field if non-nil, zero value otherwise.

### GetParentCostCentresGroupIdOk

`func (o *CostCentreGroupCreateDto) GetParentCostCentresGroupIdOk() (*string, bool)`

GetParentCostCentresGroupIdOk returns a tuple with the ParentCostCentresGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCostCentresGroupId

`func (o *CostCentreGroupCreateDto) SetParentCostCentresGroupId(v string)`

SetParentCostCentresGroupId sets ParentCostCentresGroupId field to given value.

### HasParentCostCentresGroupId

`func (o *CostCentreGroupCreateDto) HasParentCostCentresGroupId() bool`

HasParentCostCentresGroupId returns a boolean if a field has been set.

### SetParentCostCentresGroupIdNil

`func (o *CostCentreGroupCreateDto) SetParentCostCentresGroupIdNil(b bool)`

 SetParentCostCentresGroupIdNil sets the value for ParentCostCentresGroupId to be an explicit nil

### UnsetParentCostCentresGroupId
`func (o *CostCentreGroupCreateDto) UnsetParentCostCentresGroupId()`

UnsetParentCostCentresGroupId ensures that no value is present for ParentCostCentresGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


