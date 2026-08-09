# PatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**From** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchOperation

`func NewPatchOperation() *PatchOperation`

NewPatchOperation instantiates a new PatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationWithDefaults

`func NewPatchOperationWithDefaults() *PatchOperation`

NewPatchOperationWithDefaults instantiates a new PatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchOperation) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *PatchOperation) HasOp() bool`

HasOp returns a boolean if a field has been set.

### SetOpNil

`func (o *PatchOperation) SetOpNil(b bool)`

 SetOpNil sets the value for Op to be an explicit nil

### UnsetOp
`func (o *PatchOperation) UnsetOp()`

UnsetOp ensures that no value is present for Op, not even an explicit nil
### GetPath

`func (o *PatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchOperation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchOperation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PatchOperation) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PatchOperation) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetFrom

`func (o *PatchOperation) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatchOperation) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatchOperation) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatchOperation) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *PatchOperation) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *PatchOperation) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetValue

`func (o *PatchOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchOperation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


