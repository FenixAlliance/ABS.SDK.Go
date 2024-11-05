# DealUnitFlowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessProcessId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**TenantEnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDealUnitFlowDto

`func NewDealUnitFlowDto() *DealUnitFlowDto`

NewDealUnitFlowDto instantiates a new DealUnitFlowDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUnitFlowDtoWithDefaults

`func NewDealUnitFlowDtoWithDefaults() *DealUnitFlowDto`

NewDealUnitFlowDtoWithDefaults instantiates a new DealUnitFlowDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUnitFlowDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUnitFlowDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUnitFlowDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DealUnitFlowDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DealUnitFlowDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DealUnitFlowDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *DealUnitFlowDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DealUnitFlowDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DealUnitFlowDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DealUnitFlowDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DealUnitFlowDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DealUnitFlowDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *DealUnitFlowDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DealUnitFlowDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DealUnitFlowDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DealUnitFlowDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DealUnitFlowDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DealUnitFlowDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *DealUnitFlowDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUnitFlowDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUnitFlowDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUnitFlowDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealUnitFlowDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealUnitFlowDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentBusinessProcessId

`func (o *DealUnitFlowDto) GetParentBusinessProcessId() string`

GetParentBusinessProcessId returns the ParentBusinessProcessId field if non-nil, zero value otherwise.

### GetParentBusinessProcessIdOk

`func (o *DealUnitFlowDto) GetParentBusinessProcessIdOk() (*string, bool)`

GetParentBusinessProcessIdOk returns a tuple with the ParentBusinessProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessProcessId

`func (o *DealUnitFlowDto) SetParentBusinessProcessId(v string)`

SetParentBusinessProcessId sets ParentBusinessProcessId field to given value.

### HasParentBusinessProcessId

`func (o *DealUnitFlowDto) HasParentBusinessProcessId() bool`

HasParentBusinessProcessId returns a boolean if a field has been set.

### SetParentBusinessProcessIdNil

`func (o *DealUnitFlowDto) SetParentBusinessProcessIdNil(b bool)`

 SetParentBusinessProcessIdNil sets the value for ParentBusinessProcessId to be an explicit nil

### UnsetParentBusinessProcessId
`func (o *DealUnitFlowDto) UnsetParentBusinessProcessId()`

UnsetParentBusinessProcessId ensures that no value is present for ParentBusinessProcessId, not even an explicit nil
### GetTenantId

`func (o *DealUnitFlowDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DealUnitFlowDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DealUnitFlowDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DealUnitFlowDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DealUnitFlowDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DealUnitFlowDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantEnrolmentId

`func (o *DealUnitFlowDto) GetTenantEnrolmentId() string`

GetTenantEnrolmentId returns the TenantEnrolmentId field if non-nil, zero value otherwise.

### GetTenantEnrolmentIdOk

`func (o *DealUnitFlowDto) GetTenantEnrolmentIdOk() (*string, bool)`

GetTenantEnrolmentIdOk returns a tuple with the TenantEnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantEnrolmentId

`func (o *DealUnitFlowDto) SetTenantEnrolmentId(v string)`

SetTenantEnrolmentId sets TenantEnrolmentId field to given value.

### HasTenantEnrolmentId

`func (o *DealUnitFlowDto) HasTenantEnrolmentId() bool`

HasTenantEnrolmentId returns a boolean if a field has been set.

### SetTenantEnrolmentIdNil

`func (o *DealUnitFlowDto) SetTenantEnrolmentIdNil(b bool)`

 SetTenantEnrolmentIdNil sets the value for TenantEnrolmentId to be an explicit nil

### UnsetTenantEnrolmentId
`func (o *DealUnitFlowDto) UnsetTenantEnrolmentId()`

UnsetTenantEnrolmentId ensures that no value is present for TenantEnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


