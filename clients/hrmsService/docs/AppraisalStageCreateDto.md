# AppraisalStageCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AppraisalWorkflowId** | **string** |  | 
**StageOrder** | **int32** |  | 

## Methods

### NewAppraisalStageCreateDto

`func NewAppraisalStageCreateDto(name string, appraisalWorkflowId string, stageOrder int32, ) *AppraisalStageCreateDto`

NewAppraisalStageCreateDto instantiates a new AppraisalStageCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppraisalStageCreateDtoWithDefaults

`func NewAppraisalStageCreateDtoWithDefaults() *AppraisalStageCreateDto`

NewAppraisalStageCreateDtoWithDefaults instantiates a new AppraisalStageCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppraisalStageCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppraisalStageCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppraisalStageCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppraisalStageCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AppraisalStageCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppraisalStageCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppraisalStageCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppraisalStageCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *AppraisalStageCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppraisalStageCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppraisalStageCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppraisalStageCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppraisalStageCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppraisalStageCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppraisalStageCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppraisalStageCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppraisalStageCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAppraisalWorkflowId

`func (o *AppraisalStageCreateDto) GetAppraisalWorkflowId() string`

GetAppraisalWorkflowId returns the AppraisalWorkflowId field if non-nil, zero value otherwise.

### GetAppraisalWorkflowIdOk

`func (o *AppraisalStageCreateDto) GetAppraisalWorkflowIdOk() (*string, bool)`

GetAppraisalWorkflowIdOk returns a tuple with the AppraisalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppraisalWorkflowId

`func (o *AppraisalStageCreateDto) SetAppraisalWorkflowId(v string)`

SetAppraisalWorkflowId sets AppraisalWorkflowId field to given value.


### GetStageOrder

`func (o *AppraisalStageCreateDto) GetStageOrder() int32`

GetStageOrder returns the StageOrder field if non-nil, zero value otherwise.

### GetStageOrderOk

`func (o *AppraisalStageCreateDto) GetStageOrderOk() (*int32, bool)`

GetStageOrderOk returns a tuple with the StageOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageOrder

`func (o *AppraisalStageCreateDto) SetStageOrder(v int32)`

SetStageOrder sets StageOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


