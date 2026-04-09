# CostCentreCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostCentreType** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CostCentresGroupId** | Pointer to **NullableString** |  | [optional] 
**ParentCostCentreId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreCreateDto

`func NewCostCentreCreateDto() *CostCentreCreateDto`

NewCostCentreCreateDto instantiates a new CostCentreCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreCreateDtoWithDefaults

`func NewCostCentreCreateDtoWithDefaults() *CostCentreCreateDto`

NewCostCentreCreateDtoWithDefaults instantiates a new CostCentreCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CostCentreCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CostCentreCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisabled

`func (o *CostCentreCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CostCentreCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CostCentreCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CostCentreCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *CostCentreCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CostCentreCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CostCentreCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CostCentreCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CostCentreCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CostCentreCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostCentreType

`func (o *CostCentreCreateDto) GetCostCentreType() string`

GetCostCentreType returns the CostCentreType field if non-nil, zero value otherwise.

### GetCostCentreTypeOk

`func (o *CostCentreCreateDto) GetCostCentreTypeOk() (*string, bool)`

GetCostCentreTypeOk returns a tuple with the CostCentreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreType

`func (o *CostCentreCreateDto) SetCostCentreType(v string)`

SetCostCentreType sets CostCentreType field to given value.

### HasCostCentreType

`func (o *CostCentreCreateDto) HasCostCentreType() bool`

HasCostCentreType returns a boolean if a field has been set.

### GetTenantId

`func (o *CostCentreCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCostCentresGroupId

`func (o *CostCentreCreateDto) GetCostCentresGroupId() string`

GetCostCentresGroupId returns the CostCentresGroupId field if non-nil, zero value otherwise.

### GetCostCentresGroupIdOk

`func (o *CostCentreCreateDto) GetCostCentresGroupIdOk() (*string, bool)`

GetCostCentresGroupIdOk returns a tuple with the CostCentresGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentresGroupId

`func (o *CostCentreCreateDto) SetCostCentresGroupId(v string)`

SetCostCentresGroupId sets CostCentresGroupId field to given value.

### HasCostCentresGroupId

`func (o *CostCentreCreateDto) HasCostCentresGroupId() bool`

HasCostCentresGroupId returns a boolean if a field has been set.

### SetCostCentresGroupIdNil

`func (o *CostCentreCreateDto) SetCostCentresGroupIdNil(b bool)`

 SetCostCentresGroupIdNil sets the value for CostCentresGroupId to be an explicit nil

### UnsetCostCentresGroupId
`func (o *CostCentreCreateDto) UnsetCostCentresGroupId()`

UnsetCostCentresGroupId ensures that no value is present for CostCentresGroupId, not even an explicit nil
### GetParentCostCentreId

`func (o *CostCentreCreateDto) GetParentCostCentreId() string`

GetParentCostCentreId returns the ParentCostCentreId field if non-nil, zero value otherwise.

### GetParentCostCentreIdOk

`func (o *CostCentreCreateDto) GetParentCostCentreIdOk() (*string, bool)`

GetParentCostCentreIdOk returns a tuple with the ParentCostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCostCentreId

`func (o *CostCentreCreateDto) SetParentCostCentreId(v string)`

SetParentCostCentreId sets ParentCostCentreId field to given value.

### HasParentCostCentreId

`func (o *CostCentreCreateDto) HasParentCostCentreId() bool`

HasParentCostCentreId returns a boolean if a field has been set.

### SetParentCostCentreIdNil

`func (o *CostCentreCreateDto) SetParentCostCentreIdNil(b bool)`

 SetParentCostCentreIdNil sets the value for ParentCostCentreId to be an explicit nil

### UnsetParentCostCentreId
`func (o *CostCentreCreateDto) UnsetParentCostCentreId()`

UnsetParentCostCentreId ensures that no value is present for ParentCostCentreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


