# ResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**ResetCode** | **NullableString** |  | 
**NewPassword** | **NullableString** |  | 

## Methods

### NewResetPasswordRequest

`func NewResetPasswordRequest(email NullableString, resetCode NullableString, newPassword NullableString, ) *ResetPasswordRequest`

NewResetPasswordRequest instantiates a new ResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordRequestWithDefaults

`func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest`

NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ResetPasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ResetPasswordRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ResetPasswordRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetResetCode

`func (o *ResetPasswordRequest) GetResetCode() string`

GetResetCode returns the ResetCode field if non-nil, zero value otherwise.

### GetResetCodeOk

`func (o *ResetPasswordRequest) GetResetCodeOk() (*string, bool)`

GetResetCodeOk returns a tuple with the ResetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCode

`func (o *ResetPasswordRequest) SetResetCode(v string)`

SetResetCode sets ResetCode field to given value.


### SetResetCodeNil

`func (o *ResetPasswordRequest) SetResetCodeNil(b bool)`

 SetResetCodeNil sets the value for ResetCode to be an explicit nil

### UnsetResetCode
`func (o *ResetPasswordRequest) UnsetResetCode()`

UnsetResetCode ensures that no value is present for ResetCode, not even an explicit nil
### GetNewPassword

`func (o *ResetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### SetNewPasswordNil

`func (o *ResetPasswordRequest) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *ResetPasswordRequest) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


