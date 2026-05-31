# PaymentRefundDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**RefundRequestId** | Pointer to **NullableString** |  | [optional] 
**TotalFees** | Pointer to **float64** |  | [optional] 

## Methods

### NewPaymentRefundDto

`func NewPaymentRefundDto() *PaymentRefundDto`

NewPaymentRefundDto instantiates a new PaymentRefundDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRefundDtoWithDefaults

`func NewPaymentRefundDtoWithDefaults() *PaymentRefundDto`

NewPaymentRefundDtoWithDefaults instantiates a new PaymentRefundDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRefundDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRefundDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRefundDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRefundDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentRefundDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentRefundDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentRefundDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentRefundDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentRefundDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentRefundDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentRefundDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentRefundDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPaymentId

`func (o *PaymentRefundDto) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentRefundDto) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentRefundDto) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentRefundDto) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *PaymentRefundDto) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PaymentRefundDto) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetWalletAccountId

`func (o *PaymentRefundDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *PaymentRefundDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *PaymentRefundDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *PaymentRefundDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *PaymentRefundDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *PaymentRefundDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetRefundRequestId

`func (o *PaymentRefundDto) GetRefundRequestId() string`

GetRefundRequestId returns the RefundRequestId field if non-nil, zero value otherwise.

### GetRefundRequestIdOk

`func (o *PaymentRefundDto) GetRefundRequestIdOk() (*string, bool)`

GetRefundRequestIdOk returns a tuple with the RefundRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundRequestId

`func (o *PaymentRefundDto) SetRefundRequestId(v string)`

SetRefundRequestId sets RefundRequestId field to given value.

### HasRefundRequestId

`func (o *PaymentRefundDto) HasRefundRequestId() bool`

HasRefundRequestId returns a boolean if a field has been set.

### SetRefundRequestIdNil

`func (o *PaymentRefundDto) SetRefundRequestIdNil(b bool)`

 SetRefundRequestIdNil sets the value for RefundRequestId to be an explicit nil

### UnsetRefundRequestId
`func (o *PaymentRefundDto) UnsetRefundRequestId()`

UnsetRefundRequestId ensures that no value is present for RefundRequestId, not even an explicit nil
### GetTotalFees

`func (o *PaymentRefundDto) GetTotalFees() float64`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *PaymentRefundDto) GetTotalFeesOk() (*float64, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *PaymentRefundDto) SetTotalFees(v float64)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *PaymentRefundDto) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


