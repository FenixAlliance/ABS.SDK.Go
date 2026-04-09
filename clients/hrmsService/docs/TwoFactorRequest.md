# TwoFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **NullableBool** |  | [optional] 
**TwoFactorCode** | Pointer to **NullableString** |  | [optional] 
**ResetSharedKey** | Pointer to **bool** |  | [optional] 
**ResetRecoveryCodes** | Pointer to **bool** |  | [optional] 
**ForgetMachine** | Pointer to **bool** |  | [optional] 

## Methods

### NewTwoFactorRequest

`func NewTwoFactorRequest() *TwoFactorRequest`

NewTwoFactorRequest instantiates a new TwoFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorRequestWithDefaults

`func NewTwoFactorRequestWithDefaults() *TwoFactorRequest`

NewTwoFactorRequestWithDefaults instantiates a new TwoFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *TwoFactorRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *TwoFactorRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *TwoFactorRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *TwoFactorRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *TwoFactorRequest) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *TwoFactorRequest) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetTwoFactorCode

`func (o *TwoFactorRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *TwoFactorRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *TwoFactorRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *TwoFactorRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *TwoFactorRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *TwoFactorRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil
### GetResetSharedKey

`func (o *TwoFactorRequest) GetResetSharedKey() bool`

GetResetSharedKey returns the ResetSharedKey field if non-nil, zero value otherwise.

### GetResetSharedKeyOk

`func (o *TwoFactorRequest) GetResetSharedKeyOk() (*bool, bool)`

GetResetSharedKeyOk returns a tuple with the ResetSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetSharedKey

`func (o *TwoFactorRequest) SetResetSharedKey(v bool)`

SetResetSharedKey sets ResetSharedKey field to given value.

### HasResetSharedKey

`func (o *TwoFactorRequest) HasResetSharedKey() bool`

HasResetSharedKey returns a boolean if a field has been set.

### GetResetRecoveryCodes

`func (o *TwoFactorRequest) GetResetRecoveryCodes() bool`

GetResetRecoveryCodes returns the ResetRecoveryCodes field if non-nil, zero value otherwise.

### GetResetRecoveryCodesOk

`func (o *TwoFactorRequest) GetResetRecoveryCodesOk() (*bool, bool)`

GetResetRecoveryCodesOk returns a tuple with the ResetRecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetRecoveryCodes

`func (o *TwoFactorRequest) SetResetRecoveryCodes(v bool)`

SetResetRecoveryCodes sets ResetRecoveryCodes field to given value.

### HasResetRecoveryCodes

`func (o *TwoFactorRequest) HasResetRecoveryCodes() bool`

HasResetRecoveryCodes returns a boolean if a field has been set.

### GetForgetMachine

`func (o *TwoFactorRequest) GetForgetMachine() bool`

GetForgetMachine returns the ForgetMachine field if non-nil, zero value otherwise.

### GetForgetMachineOk

`func (o *TwoFactorRequest) GetForgetMachineOk() (*bool, bool)`

GetForgetMachineOk returns a tuple with the ForgetMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetMachine

`func (o *TwoFactorRequest) SetForgetMachine(v bool)`

SetForgetMachine sets ForgetMachine field to given value.

### HasForgetMachine

`func (o *TwoFactorRequest) HasForgetMachine() bool`

HasForgetMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


