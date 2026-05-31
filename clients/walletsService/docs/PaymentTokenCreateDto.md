# PaymentTokenCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Mask** | **string** |  | 
**TokenType** | Pointer to **NullableString** |  | [optional] 
**CardFranchise** | Pointer to **NullableString** |  | [optional] 
**CardExpirationMonth** | **string** |  | 
**CardExpirationYear** | **string** |  | 
**ValidUntil** | Pointer to **NullableTime** |  | [optional] 
**PaymentGatewayId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTokenCreateDto

`func NewPaymentTokenCreateDto(mask string, cardExpirationMonth string, cardExpirationYear string, ) *PaymentTokenCreateDto`

NewPaymentTokenCreateDto instantiates a new PaymentTokenCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTokenCreateDtoWithDefaults

`func NewPaymentTokenCreateDtoWithDefaults() *PaymentTokenCreateDto`

NewPaymentTokenCreateDtoWithDefaults instantiates a new PaymentTokenCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentTokenCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTokenCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTokenCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTokenCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PaymentTokenCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentTokenCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentTokenCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentTokenCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMask

`func (o *PaymentTokenCreateDto) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *PaymentTokenCreateDto) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *PaymentTokenCreateDto) SetMask(v string)`

SetMask sets Mask field to given value.


### GetTokenType

`func (o *PaymentTokenCreateDto) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *PaymentTokenCreateDto) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *PaymentTokenCreateDto) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *PaymentTokenCreateDto) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *PaymentTokenCreateDto) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *PaymentTokenCreateDto) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetCardFranchise

`func (o *PaymentTokenCreateDto) GetCardFranchise() string`

GetCardFranchise returns the CardFranchise field if non-nil, zero value otherwise.

### GetCardFranchiseOk

`func (o *PaymentTokenCreateDto) GetCardFranchiseOk() (*string, bool)`

GetCardFranchiseOk returns a tuple with the CardFranchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFranchise

`func (o *PaymentTokenCreateDto) SetCardFranchise(v string)`

SetCardFranchise sets CardFranchise field to given value.

### HasCardFranchise

`func (o *PaymentTokenCreateDto) HasCardFranchise() bool`

HasCardFranchise returns a boolean if a field has been set.

### SetCardFranchiseNil

`func (o *PaymentTokenCreateDto) SetCardFranchiseNil(b bool)`

 SetCardFranchiseNil sets the value for CardFranchise to be an explicit nil

### UnsetCardFranchise
`func (o *PaymentTokenCreateDto) UnsetCardFranchise()`

UnsetCardFranchise ensures that no value is present for CardFranchise, not even an explicit nil
### GetCardExpirationMonth

`func (o *PaymentTokenCreateDto) GetCardExpirationMonth() string`

GetCardExpirationMonth returns the CardExpirationMonth field if non-nil, zero value otherwise.

### GetCardExpirationMonthOk

`func (o *PaymentTokenCreateDto) GetCardExpirationMonthOk() (*string, bool)`

GetCardExpirationMonthOk returns a tuple with the CardExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationMonth

`func (o *PaymentTokenCreateDto) SetCardExpirationMonth(v string)`

SetCardExpirationMonth sets CardExpirationMonth field to given value.


### GetCardExpirationYear

`func (o *PaymentTokenCreateDto) GetCardExpirationYear() string`

GetCardExpirationYear returns the CardExpirationYear field if non-nil, zero value otherwise.

### GetCardExpirationYearOk

`func (o *PaymentTokenCreateDto) GetCardExpirationYearOk() (*string, bool)`

GetCardExpirationYearOk returns a tuple with the CardExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationYear

`func (o *PaymentTokenCreateDto) SetCardExpirationYear(v string)`

SetCardExpirationYear sets CardExpirationYear field to given value.


### GetValidUntil

`func (o *PaymentTokenCreateDto) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *PaymentTokenCreateDto) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *PaymentTokenCreateDto) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *PaymentTokenCreateDto) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *PaymentTokenCreateDto) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *PaymentTokenCreateDto) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetPaymentGatewayId

`func (o *PaymentTokenCreateDto) GetPaymentGatewayId() string`

GetPaymentGatewayId returns the PaymentGatewayId field if non-nil, zero value otherwise.

### GetPaymentGatewayIdOk

`func (o *PaymentTokenCreateDto) GetPaymentGatewayIdOk() (*string, bool)`

GetPaymentGatewayIdOk returns a tuple with the PaymentGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayId

`func (o *PaymentTokenCreateDto) SetPaymentGatewayId(v string)`

SetPaymentGatewayId sets PaymentGatewayId field to given value.

### HasPaymentGatewayId

`func (o *PaymentTokenCreateDto) HasPaymentGatewayId() bool`

HasPaymentGatewayId returns a boolean if a field has been set.

### SetPaymentGatewayIdNil

`func (o *PaymentTokenCreateDto) SetPaymentGatewayIdNil(b bool)`

 SetPaymentGatewayIdNil sets the value for PaymentGatewayId to be an explicit nil

### UnsetPaymentGatewayId
`func (o *PaymentTokenCreateDto) UnsetPaymentGatewayId()`

UnsetPaymentGatewayId ensures that no value is present for PaymentGatewayId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


