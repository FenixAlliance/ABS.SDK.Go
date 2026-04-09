# BankTransactionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankTransactionCreateDto

`func NewBankTransactionCreateDto() *BankTransactionCreateDto`

NewBankTransactionCreateDto instantiates a new BankTransactionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionCreateDtoWithDefaults

`func NewBankTransactionCreateDtoWithDefaults() *BankTransactionCreateDto`

NewBankTransactionCreateDtoWithDefaults instantiates a new BankTransactionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransactionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransactionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransactionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankTransactionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BankTransactionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankTransactionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankTransactionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankTransactionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBankProfileId

`func (o *BankTransactionCreateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankTransactionCreateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankTransactionCreateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankTransactionCreateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankTransactionCreateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankTransactionCreateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankTransactionCreateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankTransactionCreateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankTransactionCreateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankTransactionCreateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankTransactionCreateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankTransactionCreateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


