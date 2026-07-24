# LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**Password** | **NullableString** |  | 
**TwoFactorCode** | Pointer to **NullableString** |  | [optional] 
**TwoFactorRecoveryCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoginRequest

`func NewLoginRequest(email NullableString, password NullableString, ) *LoginRequest`

NewLoginRequest instantiates a new LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRequestWithDefaults

`func NewLoginRequestWithDefaults() *LoginRequest`

NewLoginRequestWithDefaults instantiates a new LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *LoginRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LoginRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *LoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *LoginRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LoginRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTwoFactorCode

`func (o *LoginRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *LoginRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *LoginRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *LoginRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *LoginRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *LoginRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil
### GetTwoFactorRecoveryCode

`func (o *LoginRequest) GetTwoFactorRecoveryCode() string`

GetTwoFactorRecoveryCode returns the TwoFactorRecoveryCode field if non-nil, zero value otherwise.

### GetTwoFactorRecoveryCodeOk

`func (o *LoginRequest) GetTwoFactorRecoveryCodeOk() (*string, bool)`

GetTwoFactorRecoveryCodeOk returns a tuple with the TwoFactorRecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRecoveryCode

`func (o *LoginRequest) SetTwoFactorRecoveryCode(v string)`

SetTwoFactorRecoveryCode sets TwoFactorRecoveryCode field to given value.

### HasTwoFactorRecoveryCode

`func (o *LoginRequest) HasTwoFactorRecoveryCode() bool`

HasTwoFactorRecoveryCode returns a boolean if a field has been set.

### SetTwoFactorRecoveryCodeNil

`func (o *LoginRequest) SetTwoFactorRecoveryCodeNil(b bool)`

 SetTwoFactorRecoveryCodeNil sets the value for TwoFactorRecoveryCode to be an explicit nil

### UnsetTwoFactorRecoveryCode
`func (o *LoginRequest) UnsetTwoFactorRecoveryCode()`

UnsetTwoFactorRecoveryCode ensures that no value is present for TwoFactorRecoveryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


