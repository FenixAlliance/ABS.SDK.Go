# PaymentTokenDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 
**Mask** | Pointer to **NullableString** |  | [optional] 
**TokenType** | Pointer to **NullableString** |  | [optional] 
**CardFranchise** | Pointer to **NullableString** |  | [optional] 
**CardExpirationMonth** | Pointer to **NullableString** |  | [optional] 
**CardExpirationYear** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ValidUntil** | Pointer to **NullableTime** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**PaymentGatewayId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTokenDto

`func NewPaymentTokenDto() *PaymentTokenDto`

NewPaymentTokenDto instantiates a new PaymentTokenDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTokenDtoWithDefaults

`func NewPaymentTokenDtoWithDefaults() *PaymentTokenDto`

NewPaymentTokenDtoWithDefaults instantiates a new PaymentTokenDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentTokenDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTokenDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTokenDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTokenDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentTokenDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentTokenDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentTokenDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentTokenDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentTokenDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentTokenDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentTokenDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentTokenDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTest

`func (o *PaymentTokenDto) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaymentTokenDto) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaymentTokenDto) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaymentTokenDto) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetMask

`func (o *PaymentTokenDto) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *PaymentTokenDto) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *PaymentTokenDto) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *PaymentTokenDto) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *PaymentTokenDto) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *PaymentTokenDto) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetTokenType

`func (o *PaymentTokenDto) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *PaymentTokenDto) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *PaymentTokenDto) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *PaymentTokenDto) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *PaymentTokenDto) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *PaymentTokenDto) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetCardFranchise

`func (o *PaymentTokenDto) GetCardFranchise() string`

GetCardFranchise returns the CardFranchise field if non-nil, zero value otherwise.

### GetCardFranchiseOk

`func (o *PaymentTokenDto) GetCardFranchiseOk() (*string, bool)`

GetCardFranchiseOk returns a tuple with the CardFranchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFranchise

`func (o *PaymentTokenDto) SetCardFranchise(v string)`

SetCardFranchise sets CardFranchise field to given value.

### HasCardFranchise

`func (o *PaymentTokenDto) HasCardFranchise() bool`

HasCardFranchise returns a boolean if a field has been set.

### SetCardFranchiseNil

`func (o *PaymentTokenDto) SetCardFranchiseNil(b bool)`

 SetCardFranchiseNil sets the value for CardFranchise to be an explicit nil

### UnsetCardFranchise
`func (o *PaymentTokenDto) UnsetCardFranchise()`

UnsetCardFranchise ensures that no value is present for CardFranchise, not even an explicit nil
### GetCardExpirationMonth

`func (o *PaymentTokenDto) GetCardExpirationMonth() string`

GetCardExpirationMonth returns the CardExpirationMonth field if non-nil, zero value otherwise.

### GetCardExpirationMonthOk

`func (o *PaymentTokenDto) GetCardExpirationMonthOk() (*string, bool)`

GetCardExpirationMonthOk returns a tuple with the CardExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationMonth

`func (o *PaymentTokenDto) SetCardExpirationMonth(v string)`

SetCardExpirationMonth sets CardExpirationMonth field to given value.

### HasCardExpirationMonth

`func (o *PaymentTokenDto) HasCardExpirationMonth() bool`

HasCardExpirationMonth returns a boolean if a field has been set.

### SetCardExpirationMonthNil

`func (o *PaymentTokenDto) SetCardExpirationMonthNil(b bool)`

 SetCardExpirationMonthNil sets the value for CardExpirationMonth to be an explicit nil

### UnsetCardExpirationMonth
`func (o *PaymentTokenDto) UnsetCardExpirationMonth()`

UnsetCardExpirationMonth ensures that no value is present for CardExpirationMonth, not even an explicit nil
### GetCardExpirationYear

`func (o *PaymentTokenDto) GetCardExpirationYear() string`

GetCardExpirationYear returns the CardExpirationYear field if non-nil, zero value otherwise.

### GetCardExpirationYearOk

`func (o *PaymentTokenDto) GetCardExpirationYearOk() (*string, bool)`

GetCardExpirationYearOk returns a tuple with the CardExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationYear

`func (o *PaymentTokenDto) SetCardExpirationYear(v string)`

SetCardExpirationYear sets CardExpirationYear field to given value.

### HasCardExpirationYear

`func (o *PaymentTokenDto) HasCardExpirationYear() bool`

HasCardExpirationYear returns a boolean if a field has been set.

### SetCardExpirationYearNil

`func (o *PaymentTokenDto) SetCardExpirationYearNil(b bool)`

 SetCardExpirationYearNil sets the value for CardExpirationYear to be an explicit nil

### UnsetCardExpirationYear
`func (o *PaymentTokenDto) UnsetCardExpirationYear()`

UnsetCardExpirationYear ensures that no value is present for CardExpirationYear, not even an explicit nil
### GetStatus

`func (o *PaymentTokenDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentTokenDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentTokenDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentTokenDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentTokenDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentTokenDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetValidUntil

`func (o *PaymentTokenDto) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *PaymentTokenDto) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *PaymentTokenDto) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *PaymentTokenDto) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *PaymentTokenDto) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *PaymentTokenDto) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetWalletAccountId

`func (o *PaymentTokenDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *PaymentTokenDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *PaymentTokenDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *PaymentTokenDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *PaymentTokenDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *PaymentTokenDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetPaymentGatewayId

`func (o *PaymentTokenDto) GetPaymentGatewayId() string`

GetPaymentGatewayId returns the PaymentGatewayId field if non-nil, zero value otherwise.

### GetPaymentGatewayIdOk

`func (o *PaymentTokenDto) GetPaymentGatewayIdOk() (*string, bool)`

GetPaymentGatewayIdOk returns a tuple with the PaymentGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayId

`func (o *PaymentTokenDto) SetPaymentGatewayId(v string)`

SetPaymentGatewayId sets PaymentGatewayId field to given value.

### HasPaymentGatewayId

`func (o *PaymentTokenDto) HasPaymentGatewayId() bool`

HasPaymentGatewayId returns a boolean if a field has been set.

### SetPaymentGatewayIdNil

`func (o *PaymentTokenDto) SetPaymentGatewayIdNil(b bool)`

 SetPaymentGatewayIdNil sets the value for PaymentGatewayId to be an explicit nil

### UnsetPaymentGatewayId
`func (o *PaymentTokenDto) UnsetPaymentGatewayId()`

UnsetPaymentGatewayId ensures that no value is present for PaymentGatewayId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


