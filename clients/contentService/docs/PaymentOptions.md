# PaymentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) |  | [optional] 

## Methods

### NewPaymentOptions

`func NewPaymentOptions() *PaymentOptions`

NewPaymentOptions instantiates a new PaymentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionsWithDefaults

`func NewPaymentOptionsWithDefaults() *PaymentOptions`

NewPaymentOptionsWithDefaults instantiates a new PaymentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *PaymentOptions) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentOptions) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentOptions) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentOptions) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### SetPaymentMethodsNil

`func (o *PaymentOptions) SetPaymentMethodsNil(b bool)`

 SetPaymentMethodsNil sets the value for PaymentMethods to be an explicit nil

### UnsetPaymentMethods
`func (o *PaymentOptions) UnsetPaymentMethods()`

UnsetPaymentMethods ensures that no value is present for PaymentMethods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


