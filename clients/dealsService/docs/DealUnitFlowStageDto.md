# DealUnitFlowStageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessProcessStageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDealUnitFlowStageDto

`func NewDealUnitFlowStageDto() *DealUnitFlowStageDto`

NewDealUnitFlowStageDto instantiates a new DealUnitFlowStageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitFlowStageDtoWithDefaults

`func NewDealUnitFlowStageDtoWithDefaults() *DealUnitFlowStageDto`

NewDealUnitFlowStageDtoWithDefaults instantiates a new DealUnitFlowStageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitFlowStageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitFlowStageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitFlowStageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitFlowStageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DealUnitFlowStageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DealUnitFlowStageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *DealUnitFlowStageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitFlowStageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitFlowStageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitFlowStageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DealUnitFlowStageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DealUnitFlowStageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetOrder

`func (o *DealUnitFlowStageDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DealUnitFlowStageDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DealUnitFlowStageDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DealUnitFlowStageDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *DealUnitFlowStageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DealUnitFlowStageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DealUnitFlowStageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DealUnitFlowStageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DealUnitFlowStageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DealUnitFlowStageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDealUnitFlowId

`func (o *DealUnitFlowStageDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitFlowStageDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitFlowStageDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitFlowStageDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitFlowStageDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitFlowStageDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetTenantId

`func (o *DealUnitFlowStageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitFlowStageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitFlowStageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitFlowStageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitFlowStageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitFlowStageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *DealUnitFlowStageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitFlowStageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitFlowStageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitFlowStageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitFlowStageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitFlowStageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnrolmentId

`func (o *DealUnitFlowStageDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DealUnitFlowStageDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DealUnitFlowStageDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DealUnitFlowStageDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DealUnitFlowStageDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DealUnitFlowStageDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetParentBusinessProcessStageId

`func (o *DealUnitFlowStageDto) GetParentBusinessProcessStageId() string`

GetParentBusinessProcessStageId returns the ParentBusinessProcessStageId field if non-nil, zero value otherwise.

### GetParentBusinessProcessStageIdOk

`func (o *DealUnitFlowStageDto) GetParentBusinessProcessStageIdOk() (*string, bool)`

GetParentBusinessProcessStageIdOk returns a tuple with the ParentBusinessProcessStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessProcessStageId

`func (o *DealUnitFlowStageDto) SetParentBusinessProcessStageId(v string)`

SetParentBusinessProcessStageId sets ParentBusinessProcessStageId field to given value.

### HasParentBusinessProcessStageId

`func (o *DealUnitFlowStageDto) HasParentBusinessProcessStageId() bool`

HasParentBusinessProcessStageId returns a boolean if a field has been set.

### SetParentBusinessProcessStageIdNil

`func (o *DealUnitFlowStageDto) SetParentBusinessProcessStageIdNil(b bool)`

 SetParentBusinessProcessStageIdNil sets the value for ParentBusinessProcessStageId to be an explicit nil

### UnsetParentBusinessProcessStageId
`func (o *DealUnitFlowStageDto) UnsetParentBusinessProcessStageId()`

UnsetParentBusinessProcessStageId ensures that no value is present for ParentBusinessProcessStageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


