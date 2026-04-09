# ReceiptDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**AccountHolderId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ReceiptType** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReceiptDto

`func NewReceiptDto() *ReceiptDto`

NewReceiptDto instantiates a new ReceiptDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptDtoWithDefaults

`func NewReceiptDtoWithDefaults() *ReceiptDto`

NewReceiptDtoWithDefaults instantiates a new ReceiptDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceiptDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceiptDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceiptDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReceiptDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceiptDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceiptDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ReceiptDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReceiptDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReceiptDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReceiptDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ReceiptDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ReceiptDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPaymentId

`func (o *ReceiptDto) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ReceiptDto) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ReceiptDto) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ReceiptDto) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *ReceiptDto) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *ReceiptDto) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetTenantId

`func (o *ReceiptDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReceiptDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReceiptDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ReceiptDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ReceiptDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ReceiptDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetClosed

`func (o *ReceiptDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *ReceiptDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *ReceiptDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *ReceiptDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetAccountHolderId

`func (o *ReceiptDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *ReceiptDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *ReceiptDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *ReceiptDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *ReceiptDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *ReceiptDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetContactId

`func (o *ReceiptDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ReceiptDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ReceiptDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ReceiptDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ReceiptDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ReceiptDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetEnrollmentId

`func (o *ReceiptDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ReceiptDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ReceiptDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ReceiptDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ReceiptDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ReceiptDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetReceiptType

`func (o *ReceiptDto) GetReceiptType() string`

GetReceiptType returns the ReceiptType field if non-nil, zero value otherwise.

### GetReceiptTypeOk

`func (o *ReceiptDto) GetReceiptTypeOk() (*string, bool)`

GetReceiptTypeOk returns a tuple with the ReceiptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptType

`func (o *ReceiptDto) SetReceiptType(v string)`

SetReceiptType sets ReceiptType field to given value.

### HasReceiptType

`func (o *ReceiptDto) HasReceiptType() bool`

HasReceiptType returns a boolean if a field has been set.

### GetOrderId

`func (o *ReceiptDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ReceiptDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ReceiptDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ReceiptDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ReceiptDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ReceiptDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetInvoiceId

`func (o *ReceiptDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ReceiptDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ReceiptDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ReceiptDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *ReceiptDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ReceiptDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


