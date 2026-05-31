# PaymentTokenUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mask** | Pointer to **NullableString** |  | [optional] 
**TokenType** | Pointer to **NullableString** |  | [optional] 
**CardFranchise** | Pointer to **NullableString** |  | [optional] 
**CardExpirationMonth** | Pointer to **NullableString** |  | [optional] 
**CardExpirationYear** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ValidUntil** | Pointer to **NullableTime** |  | [optional] 
**PaymentGatewayId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTokenUpdateDto

`func NewPaymentTokenUpdateDto() *PaymentTokenUpdateDto`

NewPaymentTokenUpdateDto instantiates a new PaymentTokenUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTokenUpdateDtoWithDefaults

`func NewPaymentTokenUpdateDtoWithDefaults() *PaymentTokenUpdateDto`

NewPaymentTokenUpdateDtoWithDefaults instantiates a new PaymentTokenUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMask

`func (o *PaymentTokenUpdateDto) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *PaymentTokenUpdateDto) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *PaymentTokenUpdateDto) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *PaymentTokenUpdateDto) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *PaymentTokenUpdateDto) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *PaymentTokenUpdateDto) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetTokenType

`func (o *PaymentTokenUpdateDto) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *PaymentTokenUpdateDto) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *PaymentTokenUpdateDto) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *PaymentTokenUpdateDto) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *PaymentTokenUpdateDto) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *PaymentTokenUpdateDto) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetCardFranchise

`func (o *PaymentTokenUpdateDto) GetCardFranchise() string`

GetCardFranchise returns the CardFranchise field if non-nil, zero value otherwise.

### GetCardFranchiseOk

`func (o *PaymentTokenUpdateDto) GetCardFranchiseOk() (*string, bool)`

GetCardFranchiseOk returns a tuple with the CardFranchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFranchise

`func (o *PaymentTokenUpdateDto) SetCardFranchise(v string)`

SetCardFranchise sets CardFranchise field to given value.

### HasCardFranchise

`func (o *PaymentTokenUpdateDto) HasCardFranchise() bool`

HasCardFranchise returns a boolean if a field has been set.

### SetCardFranchiseNil

`func (o *PaymentTokenUpdateDto) SetCardFranchiseNil(b bool)`

 SetCardFranchiseNil sets the value for CardFranchise to be an explicit nil

### UnsetCardFranchise
`func (o *PaymentTokenUpdateDto) UnsetCardFranchise()`

UnsetCardFranchise ensures that no value is present for CardFranchise, not even an explicit nil
### GetCardExpirationMonth

`func (o *PaymentTokenUpdateDto) GetCardExpirationMonth() string`

GetCardExpirationMonth returns the CardExpirationMonth field if non-nil, zero value otherwise.

### GetCardExpirationMonthOk

`func (o *PaymentTokenUpdateDto) GetCardExpirationMonthOk() (*string, bool)`

GetCardExpirationMonthOk returns a tuple with the CardExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationMonth

`func (o *PaymentTokenUpdateDto) SetCardExpirationMonth(v string)`

SetCardExpirationMonth sets CardExpirationMonth field to given value.

### HasCardExpirationMonth

`func (o *PaymentTokenUpdateDto) HasCardExpirationMonth() bool`

HasCardExpirationMonth returns a boolean if a field has been set.

### SetCardExpirationMonthNil

`func (o *PaymentTokenUpdateDto) SetCardExpirationMonthNil(b bool)`

 SetCardExpirationMonthNil sets the value for CardExpirationMonth to be an explicit nil

### UnsetCardExpirationMonth
`func (o *PaymentTokenUpdateDto) UnsetCardExpirationMonth()`

UnsetCardExpirationMonth ensures that no value is present for CardExpirationMonth, not even an explicit nil
### GetCardExpirationYear

`func (o *PaymentTokenUpdateDto) GetCardExpirationYear() string`

GetCardExpirationYear returns the CardExpirationYear field if non-nil, zero value otherwise.

### GetCardExpirationYearOk

`func (o *PaymentTokenUpdateDto) GetCardExpirationYearOk() (*string, bool)`

GetCardExpirationYearOk returns a tuple with the CardExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationYear

`func (o *PaymentTokenUpdateDto) SetCardExpirationYear(v string)`

SetCardExpirationYear sets CardExpirationYear field to given value.

### HasCardExpirationYear

`func (o *PaymentTokenUpdateDto) HasCardExpirationYear() bool`

HasCardExpirationYear returns a boolean if a field has been set.

### SetCardExpirationYearNil

`func (o *PaymentTokenUpdateDto) SetCardExpirationYearNil(b bool)`

 SetCardExpirationYearNil sets the value for CardExpirationYear to be an explicit nil

### UnsetCardExpirationYear
`func (o *PaymentTokenUpdateDto) UnsetCardExpirationYear()`

UnsetCardExpirationYear ensures that no value is present for CardExpirationYear, not even an explicit nil
### GetStatus

`func (o *PaymentTokenUpdateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentTokenUpdateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentTokenUpdateDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentTokenUpdateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentTokenUpdateDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentTokenUpdateDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetValidUntil

`func (o *PaymentTokenUpdateDto) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *PaymentTokenUpdateDto) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *PaymentTokenUpdateDto) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *PaymentTokenUpdateDto) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *PaymentTokenUpdateDto) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *PaymentTokenUpdateDto) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetPaymentGatewayId

`func (o *PaymentTokenUpdateDto) GetPaymentGatewayId() string`

GetPaymentGatewayId returns the PaymentGatewayId field if non-nil, zero value otherwise.

### GetPaymentGatewayIdOk

`func (o *PaymentTokenUpdateDto) GetPaymentGatewayIdOk() (*string, bool)`

GetPaymentGatewayIdOk returns a tuple with the PaymentGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayId

`func (o *PaymentTokenUpdateDto) SetPaymentGatewayId(v string)`

SetPaymentGatewayId sets PaymentGatewayId field to given value.

### HasPaymentGatewayId

`func (o *PaymentTokenUpdateDto) HasPaymentGatewayId() bool`

HasPaymentGatewayId returns a boolean if a field has been set.

### SetPaymentGatewayIdNil

`func (o *PaymentTokenUpdateDto) SetPaymentGatewayIdNil(b bool)`

 SetPaymentGatewayIdNil sets the value for PaymentGatewayId to be an explicit nil

### UnsetPaymentGatewayId
`func (o *PaymentTokenUpdateDto) UnsetPaymentGatewayId()`

UnsetPaymentGatewayId ensures that no value is present for PaymentGatewayId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


