# CostCentreDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostCentreType** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CostCentresGroupId** | Pointer to **NullableString** |  | [optional] 
**ParentCostCentreId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostCentreDto

`func NewCostCentreDto() *CostCentreDto`

NewCostCentreDto instantiates a new CostCentreDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostCentreDtoWithDefaults

`func NewCostCentreDtoWithDefaults() *CostCentreDto`

NewCostCentreDtoWithDefaults instantiates a new CostCentreDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostCentreDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostCentreDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostCentreDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CostCentreDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CostCentreDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CostCentreDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CostCentreDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CostCentreDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CostCentreDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CostCentreDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CostCentreDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CostCentreDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CostCentreDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostCentreDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostCentreDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostCentreDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostCentreDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostCentreDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisabled

`func (o *CostCentreDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CostCentreDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CostCentreDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CostCentreDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *CostCentreDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CostCentreDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CostCentreDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CostCentreDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CostCentreDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CostCentreDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostCentreType

`func (o *CostCentreDto) GetCostCentreType() string`

GetCostCentreType returns the CostCentreType field if non-nil, zero value otherwise.

### GetCostCentreTypeOk

`func (o *CostCentreDto) GetCostCentreTypeOk() (*string, bool)`

GetCostCentreTypeOk returns a tuple with the CostCentreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentreType

`func (o *CostCentreDto) SetCostCentreType(v string)`

SetCostCentreType sets CostCentreType field to given value.

### HasCostCentreType

`func (o *CostCentreDto) HasCostCentreType() bool`

HasCostCentreType returns a boolean if a field has been set.

### GetTenantId

`func (o *CostCentreDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostCentreDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostCentreDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CostCentreDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CostCentreDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CostCentreDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCostCentresGroupId

`func (o *CostCentreDto) GetCostCentresGroupId() string`

GetCostCentresGroupId returns the CostCentresGroupId field if non-nil, zero value otherwise.

### GetCostCentresGroupIdOk

`func (o *CostCentreDto) GetCostCentresGroupIdOk() (*string, bool)`

GetCostCentresGroupIdOk returns a tuple with the CostCentresGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentresGroupId

`func (o *CostCentreDto) SetCostCentresGroupId(v string)`

SetCostCentresGroupId sets CostCentresGroupId field to given value.

### HasCostCentresGroupId

`func (o *CostCentreDto) HasCostCentresGroupId() bool`

HasCostCentresGroupId returns a boolean if a field has been set.

### SetCostCentresGroupIdNil

`func (o *CostCentreDto) SetCostCentresGroupIdNil(b bool)`

 SetCostCentresGroupIdNil sets the value for CostCentresGroupId to be an explicit nil

### UnsetCostCentresGroupId
`func (o *CostCentreDto) UnsetCostCentresGroupId()`

UnsetCostCentresGroupId ensures that no value is present for CostCentresGroupId, not even an explicit nil
### GetParentCostCentreId

`func (o *CostCentreDto) GetParentCostCentreId() string`

GetParentCostCentreId returns the ParentCostCentreId field if non-nil, zero value otherwise.

### GetParentCostCentreIdOk

`func (o *CostCentreDto) GetParentCostCentreIdOk() (*string, bool)`

GetParentCostCentreIdOk returns a tuple with the ParentCostCentreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCostCentreId

`func (o *CostCentreDto) SetParentCostCentreId(v string)`

SetParentCostCentreId sets ParentCostCentreId field to given value.

### HasParentCostCentreId

`func (o *CostCentreDto) HasParentCostCentreId() bool`

HasParentCostCentreId returns a boolean if a field has been set.

### SetParentCostCentreIdNil

`func (o *CostCentreDto) SetParentCostCentreIdNil(b bool)`

 SetParentCostCentreIdNil sets the value for ParentCostCentreId to be an explicit nil

### UnsetParentCostCentreId
`func (o *CostCentreDto) UnsetParentCostCentreId()`

UnsetParentCostCentreId ensures that no value is present for ParentCostCentreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


