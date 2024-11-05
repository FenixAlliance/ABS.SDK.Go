# DealUnitFlowStageCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DealUnitFlowId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessProcessStageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDealUnitFlowStageCreateDto

`func NewDealUnitFlowStageCreateDto() *DealUnitFlowStageCreateDto`

NewDealUnitFlowStageCreateDto instantiates a new DealUnitFlowStageCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitFlowStageCreateDtoWithDefaults

`func NewDealUnitFlowStageCreateDtoWithDefaults() *DealUnitFlowStageCreateDto`

NewDealUnitFlowStageCreateDtoWithDefaults instantiates a new DealUnitFlowStageCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitFlowStageCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitFlowStageCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitFlowStageCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitFlowStageCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *DealUnitFlowStageCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitFlowStageCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitFlowStageCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitFlowStageCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOrder

`func (o *DealUnitFlowStageCreateDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DealUnitFlowStageCreateDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DealUnitFlowStageCreateDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DealUnitFlowStageCreateDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *DealUnitFlowStageCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DealUnitFlowStageCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DealUnitFlowStageCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DealUnitFlowStageCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DealUnitFlowStageCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DealUnitFlowStageCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDealUnitFlowId

`func (o *DealUnitFlowStageCreateDto) GetDealUnitFlowId() string`

GetDealUnitFlowId returns the DealUnitFlowId field if non-nil, zero value otherwise.

### GetDealUnitFlowIdOk

`func (o *DealUnitFlowStageCreateDto) GetDealUnitFlowIdOk() (*string, bool)`

GetDealUnitFlowIdOk returns a tuple with the DealUnitFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealUnitFlowId

`func (o *DealUnitFlowStageCreateDto) SetDealUnitFlowId(v string)`

SetDealUnitFlowId sets DealUnitFlowId field to given value.

### HasDealUnitFlowId

`func (o *DealUnitFlowStageCreateDto) HasDealUnitFlowId() bool`

HasDealUnitFlowId returns a boolean if a field has been set.

### SetDealUnitFlowIdNil

`func (o *DealUnitFlowStageCreateDto) SetDealUnitFlowIdNil(b bool)`

 SetDealUnitFlowIdNil sets the value for DealUnitFlowId to be an explicit nil

### UnsetDealUnitFlowId
`func (o *DealUnitFlowStageCreateDto) UnsetDealUnitFlowId()`

UnsetDealUnitFlowId ensures that no value is present for DealUnitFlowId, not even an explicit nil
### GetTenantId

`func (o *DealUnitFlowStageCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitFlowStageCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitFlowStageCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitFlowStageCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitFlowStageCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitFlowStageCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *DealUnitFlowStageCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitFlowStageCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitFlowStageCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitFlowStageCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitFlowStageCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitFlowStageCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnrolmentId

`func (o *DealUnitFlowStageCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DealUnitFlowStageCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DealUnitFlowStageCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DealUnitFlowStageCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DealUnitFlowStageCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DealUnitFlowStageCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetParentBusinessProcessStageId

`func (o *DealUnitFlowStageCreateDto) GetParentBusinessProcessStageId() string`

GetParentBusinessProcessStageId returns the ParentBusinessProcessStageId field if non-nil, zero value otherwise.

### GetParentBusinessProcessStageIdOk

`func (o *DealUnitFlowStageCreateDto) GetParentBusinessProcessStageIdOk() (*string, bool)`

GetParentBusinessProcessStageIdOk returns a tuple with the ParentBusinessProcessStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessProcessStageId

`func (o *DealUnitFlowStageCreateDto) SetParentBusinessProcessStageId(v string)`

SetParentBusinessProcessStageId sets ParentBusinessProcessStageId field to given value.

### HasParentBusinessProcessStageId

`func (o *DealUnitFlowStageCreateDto) HasParentBusinessProcessStageId() bool`

HasParentBusinessProcessStageId returns a boolean if a field has been set.

### SetParentBusinessProcessStageIdNil

`func (o *DealUnitFlowStageCreateDto) SetParentBusinessProcessStageIdNil(b bool)`

 SetParentBusinessProcessStageIdNil sets the value for ParentBusinessProcessStageId to be an explicit nil

### UnsetParentBusinessProcessStageId
`func (o *DealUnitFlowStageCreateDto) UnsetParentBusinessProcessStageId()`

UnsetParentBusinessProcessStageId ensures that no value is present for ParentBusinessProcessStageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


