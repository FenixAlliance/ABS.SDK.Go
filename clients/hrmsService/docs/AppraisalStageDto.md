# AppraisalStageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**AppraisalWorkflowId** | Pointer to **NullableString** |  | [optional] 
**StageOrder** | Pointer to **int32** |  | [optional] 

## Methods

### NewAppraisalStageDto

`func NewAppraisalStageDto() *AppraisalStageDto`

NewAppraisalStageDto instantiates a new AppraisalStageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppraisalStageDtoWithDefaults

`func NewAppraisalStageDtoWithDefaults() *AppraisalStageDto`

NewAppraisalStageDtoWithDefaults instantiates a new AppraisalStageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppraisalStageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppraisalStageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppraisalStageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppraisalStageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppraisalStageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppraisalStageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AppraisalStageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppraisalStageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppraisalStageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppraisalStageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AppraisalStageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AppraisalStageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *AppraisalStageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppraisalStageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppraisalStageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppraisalStageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppraisalStageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppraisalStageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AppraisalStageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppraisalStageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppraisalStageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppraisalStageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppraisalStageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppraisalStageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTenantId

`func (o *AppraisalStageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AppraisalStageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AppraisalStageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AppraisalStageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AppraisalStageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AppraisalStageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAppraisalWorkflowId

`func (o *AppraisalStageDto) GetAppraisalWorkflowId() string`

GetAppraisalWorkflowId returns the AppraisalWorkflowId field if non-nil, zero value otherwise.

### GetAppraisalWorkflowIdOk

`func (o *AppraisalStageDto) GetAppraisalWorkflowIdOk() (*string, bool)`

GetAppraisalWorkflowIdOk returns a tuple with the AppraisalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalWorkflowId

`func (o *AppraisalStageDto) SetAppraisalWorkflowId(v string)`

SetAppraisalWorkflowId sets AppraisalWorkflowId field to given value.

### HasAppraisalWorkflowId

`func (o *AppraisalStageDto) HasAppraisalWorkflowId() bool`

HasAppraisalWorkflowId returns a boolean if a field has been set.

### SetAppraisalWorkflowIdNil

`func (o *AppraisalStageDto) SetAppraisalWorkflowIdNil(b bool)`

 SetAppraisalWorkflowIdNil sets the value for AppraisalWorkflowId to be an explicit nil

### UnsetAppraisalWorkflowId
`func (o *AppraisalStageDto) UnsetAppraisalWorkflowId()`

UnsetAppraisalWorkflowId ensures that no value is present for AppraisalWorkflowId, not even an explicit nil
### GetStageOrder

`func (o *AppraisalStageDto) GetStageOrder() int32`

GetStageOrder returns the StageOrder field if non-nil, zero value otherwise.

### GetStageOrderOk

`func (o *AppraisalStageDto) GetStageOrderOk() (*int32, bool)`

GetStageOrderOk returns a tuple with the StageOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageOrder

`func (o *AppraisalStageDto) SetStageOrder(v int32)`

SetStageOrder sets StageOrder field to given value.

### HasStageOrder

`func (o *AppraisalStageDto) HasStageOrder() bool`

HasStageOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


