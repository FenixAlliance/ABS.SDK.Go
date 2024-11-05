# PaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | Pointer to **bool** |  | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**Bank** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to **interface{}** |  | [optional] 
**Response** | Pointer to **NullableString** |  | [optional] 
**AuthCode** | Pointer to **NullableString** |  | [optional] 
**PaymentID** | Pointer to **NullableString** |  | [optional] 
**Franchise** | Pointer to **NullableString** |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**PaymentStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaymentResponse

`func NewPaymentResponse() *PaymentResponse`

NewPaymentResponse instantiates a new PaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseWithDefaults

`func NewPaymentResponseWithDefaults() *PaymentResponse`

NewPaymentResponseWithDefaults instantiates a new PaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *PaymentResponse) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaymentResponse) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaymentResponse) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaymentResponse) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetIp

`func (o *PaymentResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PaymentResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PaymentResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PaymentResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *PaymentResponse) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *PaymentResponse) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetBank

`func (o *PaymentResponse) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *PaymentResponse) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *PaymentResponse) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *PaymentResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.

### SetBankNil

`func (o *PaymentResponse) SetBankNil(b bool)`

 SetBankNil sets the value for Bank to be an explicit nil

### UnsetBank
`func (o *PaymentResponse) UnsetBank()`

UnsetBank ensures that no value is present for Bank, not even an explicit nil
### GetStatus

`func (o *PaymentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrors

`func (o *PaymentResponse) GetErrors() interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PaymentResponse) GetErrorsOk() (*interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PaymentResponse) SetErrors(v interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PaymentResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *PaymentResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *PaymentResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetResponse

`func (o *PaymentResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PaymentResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PaymentResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PaymentResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *PaymentResponse) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *PaymentResponse) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetAuthCode

`func (o *PaymentResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaymentResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *PaymentResponse) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *PaymentResponse) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetPaymentID

`func (o *PaymentResponse) GetPaymentID() string`

GetPaymentID returns the PaymentID field if non-nil, zero value otherwise.

### GetPaymentIDOk

`func (o *PaymentResponse) GetPaymentIDOk() (*string, bool)`

GetPaymentIDOk returns a tuple with the PaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentID

`func (o *PaymentResponse) SetPaymentID(v string)`

SetPaymentID sets PaymentID field to given value.

### HasPaymentID

`func (o *PaymentResponse) HasPaymentID() bool`

HasPaymentID returns a boolean if a field has been set.

### SetPaymentIDNil

`func (o *PaymentResponse) SetPaymentIDNil(b bool)`

 SetPaymentIDNil sets the value for PaymentID to be an explicit nil

### UnsetPaymentID
`func (o *PaymentResponse) UnsetPaymentID()`

UnsetPaymentID ensures that no value is present for PaymentID, not even an explicit nil
### GetFranchise

`func (o *PaymentResponse) GetFranchise() string`

GetFranchise returns the Franchise field if non-nil, zero value otherwise.

### GetFranchiseOk

`func (o *PaymentResponse) GetFranchiseOk() (*string, bool)`

GetFranchiseOk returns a tuple with the Franchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFranchise

`func (o *PaymentResponse) SetFranchise(v string)`

SetFranchise sets Franchise field to given value.

### HasFranchise

`func (o *PaymentResponse) HasFranchise() bool`

HasFranchise returns a boolean if a field has been set.

### SetFranchiseNil

`func (o *PaymentResponse) SetFranchiseNil(b bool)`

 SetFranchiseNil sets the value for Franchise to be an explicit nil

### UnsetFranchise
`func (o *PaymentResponse) UnsetFranchise()`

UnsetFranchise ensures that no value is present for Franchise, not even an explicit nil
### GetSignature

`func (o *PaymentResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PaymentResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *PaymentResponse) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *PaymentResponse) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetPaymentStatus

`func (o *PaymentResponse) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentResponse) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentResponse) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PaymentResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


