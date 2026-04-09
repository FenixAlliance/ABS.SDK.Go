# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**Password** | **NullableString** |  | 

## Methods

### NewRegisterRequest

`func NewRegisterRequest(email NullableString, password NullableString, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *RegisterRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RegisterRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *RegisterRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *RegisterRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisterRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


