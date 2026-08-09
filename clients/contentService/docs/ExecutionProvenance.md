# ExecutionProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiation** | Pointer to **string** |  | [optional] 
**OnBehalfOfActorId** | Pointer to **map[string]interface{}** |  | [optional] 
**OnBehalfOfActorKind** | Pointer to **NullableString** |  | [optional] 
**CausationId** | Pointer to **NullableString** |  | [optional] 
**OriginatingWorkflowInstanceId** | Pointer to **NullableString** |  | [optional] 
**EventDepth** | Pointer to **int32** |  | [optional] 

## Methods

### NewExecutionProvenance

`func NewExecutionProvenance() *ExecutionProvenance`

NewExecutionProvenance instantiates a new ExecutionProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionProvenanceWithDefaults

`func NewExecutionProvenanceWithDefaults() *ExecutionProvenance`

NewExecutionProvenanceWithDefaults instantiates a new ExecutionProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiation

`func (o *ExecutionProvenance) GetInitiation() string`

GetInitiation returns the Initiation field if non-nil, zero value otherwise.

### GetInitiationOk

`func (o *ExecutionProvenance) GetInitiationOk() (*string, bool)`

GetInitiationOk returns a tuple with the Initiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiation

`func (o *ExecutionProvenance) SetInitiation(v string)`

SetInitiation sets Initiation field to given value.

### HasInitiation

`func (o *ExecutionProvenance) HasInitiation() bool`

HasInitiation returns a boolean if a field has been set.

### GetOnBehalfOfActorId

`func (o *ExecutionProvenance) GetOnBehalfOfActorId() map[string]interface{}`

GetOnBehalfOfActorId returns the OnBehalfOfActorId field if non-nil, zero value otherwise.

### GetOnBehalfOfActorIdOk

`func (o *ExecutionProvenance) GetOnBehalfOfActorIdOk() (*map[string]interface{}, bool)`

GetOnBehalfOfActorIdOk returns a tuple with the OnBehalfOfActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOfActorId

`func (o *ExecutionProvenance) SetOnBehalfOfActorId(v map[string]interface{})`

SetOnBehalfOfActorId sets OnBehalfOfActorId field to given value.

### HasOnBehalfOfActorId

`func (o *ExecutionProvenance) HasOnBehalfOfActorId() bool`

HasOnBehalfOfActorId returns a boolean if a field has been set.

### GetOnBehalfOfActorKind

`func (o *ExecutionProvenance) GetOnBehalfOfActorKind() string`

GetOnBehalfOfActorKind returns the OnBehalfOfActorKind field if non-nil, zero value otherwise.

### GetOnBehalfOfActorKindOk

`func (o *ExecutionProvenance) GetOnBehalfOfActorKindOk() (*string, bool)`

GetOnBehalfOfActorKindOk returns a tuple with the OnBehalfOfActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOfActorKind

`func (o *ExecutionProvenance) SetOnBehalfOfActorKind(v string)`

SetOnBehalfOfActorKind sets OnBehalfOfActorKind field to given value.

### HasOnBehalfOfActorKind

`func (o *ExecutionProvenance) HasOnBehalfOfActorKind() bool`

HasOnBehalfOfActorKind returns a boolean if a field has been set.

### SetOnBehalfOfActorKindNil

`func (o *ExecutionProvenance) SetOnBehalfOfActorKindNil(b bool)`

 SetOnBehalfOfActorKindNil sets the value for OnBehalfOfActorKind to be an explicit nil

### UnsetOnBehalfOfActorKind
`func (o *ExecutionProvenance) UnsetOnBehalfOfActorKind()`

UnsetOnBehalfOfActorKind ensures that no value is present for OnBehalfOfActorKind, not even an explicit nil
### GetCausationId

`func (o *ExecutionProvenance) GetCausationId() string`

GetCausationId returns the CausationId field if non-nil, zero value otherwise.

### GetCausationIdOk

`func (o *ExecutionProvenance) GetCausationIdOk() (*string, bool)`

GetCausationIdOk returns a tuple with the CausationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausationId

`func (o *ExecutionProvenance) SetCausationId(v string)`

SetCausationId sets CausationId field to given value.

### HasCausationId

`func (o *ExecutionProvenance) HasCausationId() bool`

HasCausationId returns a boolean if a field has been set.

### SetCausationIdNil

`func (o *ExecutionProvenance) SetCausationIdNil(b bool)`

 SetCausationIdNil sets the value for CausationId to be an explicit nil

### UnsetCausationId
`func (o *ExecutionProvenance) UnsetCausationId()`

UnsetCausationId ensures that no value is present for CausationId, not even an explicit nil
### GetOriginatingWorkflowInstanceId

`func (o *ExecutionProvenance) GetOriginatingWorkflowInstanceId() string`

GetOriginatingWorkflowInstanceId returns the OriginatingWorkflowInstanceId field if non-nil, zero value otherwise.

### GetOriginatingWorkflowInstanceIdOk

`func (o *ExecutionProvenance) GetOriginatingWorkflowInstanceIdOk() (*string, bool)`

GetOriginatingWorkflowInstanceIdOk returns a tuple with the OriginatingWorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingWorkflowInstanceId

`func (o *ExecutionProvenance) SetOriginatingWorkflowInstanceId(v string)`

SetOriginatingWorkflowInstanceId sets OriginatingWorkflowInstanceId field to given value.

### HasOriginatingWorkflowInstanceId

`func (o *ExecutionProvenance) HasOriginatingWorkflowInstanceId() bool`

HasOriginatingWorkflowInstanceId returns a boolean if a field has been set.

### SetOriginatingWorkflowInstanceIdNil

`func (o *ExecutionProvenance) SetOriginatingWorkflowInstanceIdNil(b bool)`

 SetOriginatingWorkflowInstanceIdNil sets the value for OriginatingWorkflowInstanceId to be an explicit nil

### UnsetOriginatingWorkflowInstanceId
`func (o *ExecutionProvenance) UnsetOriginatingWorkflowInstanceId()`

UnsetOriginatingWorkflowInstanceId ensures that no value is present for OriginatingWorkflowInstanceId, not even an explicit nil
### GetEventDepth

`func (o *ExecutionProvenance) GetEventDepth() int32`

GetEventDepth returns the EventDepth field if non-nil, zero value otherwise.

### GetEventDepthOk

`func (o *ExecutionProvenance) GetEventDepthOk() (*int32, bool)`

GetEventDepthOk returns a tuple with the EventDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDepth

`func (o *ExecutionProvenance) SetEventDepth(v int32)`

SetEventDepth sets EventDepth field to given value.

### HasEventDepth

`func (o *ExecutionProvenance) HasEventDepth() bool`

HasEventDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


