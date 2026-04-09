# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**SupportedFeatures** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *PaymentMethod) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PaymentMethod) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PaymentMethod) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PaymentMethod) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPriority

`func (o *PaymentMethod) GetPriority() bool`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PaymentMethod) GetPriorityOk() (*bool, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PaymentMethod) SetPriority(v bool)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PaymentMethod) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PaymentMethod) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaymentMethod) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *PaymentMethod) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PaymentMethod) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PaymentMethod) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *PaymentMethod) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *PaymentMethod) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *PaymentMethod) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetSupportedFeatures

`func (o *PaymentMethod) GetSupportedFeatures() []string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PaymentMethod) GetSupportedFeaturesOk() (*[]string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PaymentMethod) SetSupportedFeatures(v []string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PaymentMethod) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### SetSupportedFeaturesNil

`func (o *PaymentMethod) SetSupportedFeaturesNil(b bool)`

 SetSupportedFeaturesNil sets the value for SupportedFeatures to be an explicit nil

### UnsetSupportedFeatures
`func (o *PaymentMethod) UnsetSupportedFeatures()`

UnsetSupportedFeatures ensures that no value is present for SupportedFeatures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


