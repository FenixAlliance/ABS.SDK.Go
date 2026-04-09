# TwoFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedKey** | **NullableString** |  | 
**RecoveryCodesLeft** | **int32** |  | 
**RecoveryCodes** | Pointer to **[]string** |  | [optional] 
**IsTwoFactorEnabled** | **bool** |  | 
**IsMachineRemembered** | **bool** |  | 

## Methods

### NewTwoFactorResponse

`func NewTwoFactorResponse(sharedKey NullableString, recoveryCodesLeft int32, isTwoFactorEnabled bool, isMachineRemembered bool, ) *TwoFactorResponse`

NewTwoFactorResponse instantiates a new TwoFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorResponseWithDefaults

`func NewTwoFactorResponseWithDefaults() *TwoFactorResponse`

NewTwoFactorResponseWithDefaults instantiates a new TwoFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedKey

`func (o *TwoFactorResponse) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *TwoFactorResponse) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *TwoFactorResponse) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.


### SetSharedKeyNil

`func (o *TwoFactorResponse) SetSharedKeyNil(b bool)`

 SetSharedKeyNil sets the value for SharedKey to be an explicit nil

### UnsetSharedKey
`func (o *TwoFactorResponse) UnsetSharedKey()`

UnsetSharedKey ensures that no value is present for SharedKey, not even an explicit nil
### GetRecoveryCodesLeft

`func (o *TwoFactorResponse) GetRecoveryCodesLeft() int32`

GetRecoveryCodesLeft returns the RecoveryCodesLeft field if non-nil, zero value otherwise.

### GetRecoveryCodesLeftOk

`func (o *TwoFactorResponse) GetRecoveryCodesLeftOk() (*int32, bool)`

GetRecoveryCodesLeftOk returns a tuple with the RecoveryCodesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodesLeft

`func (o *TwoFactorResponse) SetRecoveryCodesLeft(v int32)`

SetRecoveryCodesLeft sets RecoveryCodesLeft field to given value.


### GetRecoveryCodes

`func (o *TwoFactorResponse) GetRecoveryCodes() []string`

GetRecoveryCodes returns the RecoveryCodes field if non-nil, zero value otherwise.

### GetRecoveryCodesOk

`func (o *TwoFactorResponse) GetRecoveryCodesOk() (*[]string, bool)`

GetRecoveryCodesOk returns a tuple with the RecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodes

`func (o *TwoFactorResponse) SetRecoveryCodes(v []string)`

SetRecoveryCodes sets RecoveryCodes field to given value.

### HasRecoveryCodes

`func (o *TwoFactorResponse) HasRecoveryCodes() bool`

HasRecoveryCodes returns a boolean if a field has been set.

### SetRecoveryCodesNil

`func (o *TwoFactorResponse) SetRecoveryCodesNil(b bool)`

 SetRecoveryCodesNil sets the value for RecoveryCodes to be an explicit nil

### UnsetRecoveryCodes
`func (o *TwoFactorResponse) UnsetRecoveryCodes()`

UnsetRecoveryCodes ensures that no value is present for RecoveryCodes, not even an explicit nil
### GetIsTwoFactorEnabled

`func (o *TwoFactorResponse) GetIsTwoFactorEnabled() bool`

GetIsTwoFactorEnabled returns the IsTwoFactorEnabled field if non-nil, zero value otherwise.

### GetIsTwoFactorEnabledOk

`func (o *TwoFactorResponse) GetIsTwoFactorEnabledOk() (*bool, bool)`

GetIsTwoFactorEnabledOk returns a tuple with the IsTwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTwoFactorEnabled

`func (o *TwoFactorResponse) SetIsTwoFactorEnabled(v bool)`

SetIsTwoFactorEnabled sets IsTwoFactorEnabled field to given value.


### GetIsMachineRemembered

`func (o *TwoFactorResponse) GetIsMachineRemembered() bool`

GetIsMachineRemembered returns the IsMachineRemembered field if non-nil, zero value otherwise.

### GetIsMachineRememberedOk

`func (o *TwoFactorResponse) GetIsMachineRememberedOk() (*bool, bool)`

GetIsMachineRememberedOk returns a tuple with the IsMachineRemembered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachineRemembered

`func (o *TwoFactorResponse) SetIsMachineRemembered(v bool)`

SetIsMachineRemembered sets IsMachineRemembered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


