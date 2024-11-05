# DealUnitFlowStageUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowId** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessProcessStageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDealUnitFlowStageUpdateDto

`func NewDealUnitFlowStageUpdateDto() *DealUnitFlowStageUpdateDto`

NewDealUnitFlowStageUpdateDto instantiates a new DealUnitFlowStageUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitFlowStageUpdateDtoWithDefaults

`func NewDealUnitFlowStageUpdateDtoWithDefaults() *DealUnitFlowStageUpdateDto`

NewDealUnitFlowStageUpdateDtoWithDefaults instantiates a new DealUnitFlowStageUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *DealUnitFlowStageUpdateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DealUnitFlowStageUpdateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DealUnitFlowStageUpdateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DealUnitFlowStageUpdateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *DealUnitFlowStageUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DealUnitFlowStageUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DealUnitFlowStageUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DealUnitFlowStageUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DealUnitFlowStageUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DealUnitFlowStageUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *DealUnitFlowStageUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitFlowStageUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitFlowStageUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitFlowStageUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitFlowStageUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitFlowStageUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnrolmentId

`func (o *DealUnitFlowStageUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DealUnitFlowStageUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DealUnitFlowStageUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DealUnitFlowStageUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DealUnitFlowStageUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DealUnitFlowStageUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetDealUnitFlowId

`func (o *DealUnitFlowStageUpdateDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitFlowStageUpdateDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitFlowStageUpdateDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitFlowStageUpdateDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitFlowStageUpdateDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitFlowStageUpdateDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetParentBusinessProcessStageId

`func (o *DealUnitFlowStageUpdateDto) GetParentBusinessProcessStageId() string`

GetParentBusinessProcessStageId returns the ParentBusinessProcessStageId field if non-nil, zero value otherwise.

### GetParentBusinessProcessStageIdOk

`func (o *DealUnitFlowStageUpdateDto) GetParentBusinessProcessStageIdOk() (*string, bool)`

GetParentBusinessProcessStageIdOk returns a tuple with the ParentBusinessProcessStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessProcessStageId

`func (o *DealUnitFlowStageUpdateDto) SetParentBusinessProcessStageId(v string)`

SetParentBusinessProcessStageId sets ParentBusinessProcessStageId field to given value.

### HasParentBusinessProcessStageId

`func (o *DealUnitFlowStageUpdateDto) HasParentBusinessProcessStageId() bool`

HasParentBusinessProcessStageId returns a boolean if a field has been set.

### SetParentBusinessProcessStageIdNil

`func (o *DealUnitFlowStageUpdateDto) SetParentBusinessProcessStageIdNil(b bool)`

 SetParentBusinessProcessStageIdNil sets the value for ParentBusinessProcessStageId to be an explicit nil

### UnsetParentBusinessProcessStageId
`func (o *DealUnitFlowStageUpdateDto) UnsetParentBusinessProcessStageId()`

UnsetParentBusinessProcessStageId ensures that no value is present for ParentBusinessProcessStageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


