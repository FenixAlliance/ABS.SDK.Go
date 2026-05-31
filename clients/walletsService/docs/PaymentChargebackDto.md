# PaymentChargebackDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**RequestDate** | Pointer to **time.Time** |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 
**BankProfileName** | Pointer to **NullableString** |  | [optional] 
**TotalFees** | Pointer to **float64** |  | [optional] 

## Methods

### NewPaymentChargebackDto

`func NewPaymentChargebackDto() *PaymentChargebackDto`

NewPaymentChargebackDto instantiates a new PaymentChargebackDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChargebackDtoWithDefaults

`func NewPaymentChargebackDtoWithDefaults() *PaymentChargebackDto`

NewPaymentChargebackDtoWithDefaults instantiates a new PaymentChargebackDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentChargebackDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentChargebackDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentChargebackDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentChargebackDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentChargebackDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentChargebackDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentChargebackDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentChargebackDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentChargebackDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentChargebackDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentChargebackDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentChargebackDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequestDate

`func (o *PaymentChargebackDto) GetRequestDate() time.Time`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *PaymentChargebackDto) GetRequestDateOk() (*time.Time, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *PaymentChargebackDto) SetRequestDate(v time.Time)`

SetRequestDate sets RequestDate field to given value.

### HasRequestDate

`func (o *PaymentChargebackDto) HasRequestDate() bool`

HasRequestDate returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentChargebackDto) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentChargebackDto) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentChargebackDto) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentChargebackDto) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *PaymentChargebackDto) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PaymentChargebackDto) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetBankProfileId

`func (o *PaymentChargebackDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *PaymentChargebackDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *PaymentChargebackDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *PaymentChargebackDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *PaymentChargebackDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *PaymentChargebackDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankProfileName

`func (o *PaymentChargebackDto) GetBankProfileName() string`

GetBankProfileName returns the BankProfileName field if non-nil, zero value otherwise.

### GetBankProfileNameOk

`func (o *PaymentChargebackDto) GetBankProfileNameOk() (*string, bool)`

GetBankProfileNameOk returns a tuple with the BankProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileName

`func (o *PaymentChargebackDto) SetBankProfileName(v string)`

SetBankProfileName sets BankProfileName field to given value.

### HasBankProfileName

`func (o *PaymentChargebackDto) HasBankProfileName() bool`

HasBankProfileName returns a boolean if a field has been set.

### SetBankProfileNameNil

`func (o *PaymentChargebackDto) SetBankProfileNameNil(b bool)`

 SetBankProfileNameNil sets the value for BankProfileName to be an explicit nil

### UnsetBankProfileName
`func (o *PaymentChargebackDto) UnsetBankProfileName()`

UnsetBankProfileName ensures that no value is present for BankProfileName, not even an explicit nil
### GetTotalFees

`func (o *PaymentChargebackDto) GetTotalFees() float64`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *PaymentChargebackDto) GetTotalFeesOk() (*float64, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *PaymentChargebackDto) SetTotalFees(v float64)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *PaymentChargebackDto) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


