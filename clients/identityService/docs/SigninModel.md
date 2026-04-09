# SigninModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigninModel

`func NewSigninModel() *SigninModel`

NewSigninModel instantiates a new SigninModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigninModelWithDefaults

`func NewSigninModelWithDefaults() *SigninModel`

NewSigninModelWithDefaults instantiates a new SigninModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SigninModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SigninModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SigninModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SigninModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SigninModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SigninModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *SigninModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SigninModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SigninModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SigninModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SigninModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SigninModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


