# InfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**IsEmailConfirmed** | **bool** |  | 

## Methods

### NewInfoResponse

`func NewInfoResponse(email NullableString, isEmailConfirmed bool, ) *InfoResponse`

NewInfoResponse instantiates a new InfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseWithDefaults

`func NewInfoResponseWithDefaults() *InfoResponse`

NewInfoResponseWithDefaults instantiates a new InfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *InfoResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *InfoResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsEmailConfirmed

`func (o *InfoResponse) GetIsEmailConfirmed() bool`

GetIsEmailConfirmed returns the IsEmailConfirmed field if non-nil, zero value otherwise.

### GetIsEmailConfirmedOk

`func (o *InfoResponse) GetIsEmailConfirmedOk() (*bool, bool)`

GetIsEmailConfirmedOk returns a tuple with the IsEmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailConfirmed

`func (o *InfoResponse) SetIsEmailConfirmed(v bool)`

SetIsEmailConfirmed sets IsEmailConfirmed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


