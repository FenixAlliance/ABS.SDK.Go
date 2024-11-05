# DealUnitFlowCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessProcessId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**TenantEnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDealUnitFlowCreateDto

`func NewDealUnitFlowCreateDto() *DealUnitFlowCreateDto`

NewDealUnitFlowCreateDto instantiates a new DealUnitFlowCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitFlowCreateDtoWithDefaults

`func NewDealUnitFlowCreateDtoWithDefaults() *DealUnitFlowCreateDto`

NewDealUnitFlowCreateDtoWithDefaults instantiates a new DealUnitFlowCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitFlowCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitFlowCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitFlowCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitFlowCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *DealUnitFlowCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitFlowCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitFlowCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitFlowCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *DealUnitFlowCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DealUnitFlowCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DealUnitFlowCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DealUnitFlowCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DealUnitFlowCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DealUnitFlowCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *DealUnitFlowCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitFlowCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitFlowCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitFlowCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitFlowCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitFlowCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentBusinessProcessId

`func (o *DealUnitFlowCreateDto) GetParentBusinessProcessId() string`

GetParentBusinessProcessId returns the ParentBusinessProcessId field if non-nil, zero value otherwise.

### GetParentBusinessProcessIdOk

`func (o *DealUnitFlowCreateDto) GetParentBusinessProcessIdOk() (*string, bool)`

GetParentBusinessProcessIdOk returns a tuple with the ParentBusinessProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessProcessId

`func (o *DealUnitFlowCreateDto) SetParentBusinessProcessId(v string)`

SetParentBusinessProcessId sets ParentBusinessProcessId field to given value.

### HasParentBusinessProcessId

`func (o *DealUnitFlowCreateDto) HasParentBusinessProcessId() bool`

HasParentBusinessProcessId returns a boolean if a field has been set.

### SetParentBusinessProcessIdNil

`func (o *DealUnitFlowCreateDto) SetParentBusinessProcessIdNil(b bool)`

 SetParentBusinessProcessIdNil sets the value for ParentBusinessProcessId to be an explicit nil

### UnsetParentBusinessProcessId
`func (o *DealUnitFlowCreateDto) UnsetParentBusinessProcessId()`

UnsetParentBusinessProcessId ensures that no value is present for ParentBusinessProcessId, not even an explicit nil
### GetTenantId

`func (o *DealUnitFlowCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitFlowCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitFlowCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitFlowCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitFlowCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitFlowCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantEnrolmentId

`func (o *DealUnitFlowCreateDto) GetTenantEnrolmentId() string`

GetTenantEnrolmentId returns the TenantEnrolmentId field if non-nil, zero value otherwise.

### GetTenantEnrolmentIdOk

`func (o *DealUnitFlowCreateDto) GetTenantEnrolmentIdOk() (*string, bool)`

GetTenantEnrolmentIdOk returns a tuple with the TenantEnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantEnrolmentId

`func (o *DealUnitFlowCreateDto) SetTenantEnrolmentId(v string)`

SetTenantEnrolmentId sets TenantEnrolmentId field to given value.

### HasTenantEnrolmentId

`func (o *DealUnitFlowCreateDto) HasTenantEnrolmentId() bool`

HasTenantEnrolmentId returns a boolean if a field has been set.

### SetTenantEnrolmentIdNil

`func (o *DealUnitFlowCreateDto) SetTenantEnrolmentIdNil(b bool)`

 SetTenantEnrolmentIdNil sets the value for TenantEnrolmentId to be an explicit nil

### UnsetTenantEnrolmentId
`func (o *DealUnitFlowCreateDto) UnsetTenantEnrolmentId()`

UnsetTenantEnrolmentId ensures that no value is present for TenantEnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


