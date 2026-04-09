# JsonWebToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**JsonWebTokenHeader**](JsonWebTokenHeader.md) |  | [optional] 
**Payload** | Pointer to [**JsonWebTokenPayload**](JsonWebTokenPayload.md) |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**TokenType** | Pointer to **NullableString** |  | [optional] 
**ExpiresIn** | Pointer to **int64** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJsonWebToken

`func NewJsonWebToken() *JsonWebToken`

NewJsonWebToken instantiates a new JsonWebToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebTokenWithDefaults

`func NewJsonWebTokenWithDefaults() *JsonWebToken`

NewJsonWebTokenWithDefaults instantiates a new JsonWebToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *JsonWebToken) GetHeader() JsonWebTokenHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *JsonWebToken) GetHeaderOk() (*JsonWebTokenHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *JsonWebToken) SetHeader(v JsonWebTokenHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *JsonWebToken) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetPayload

`func (o *JsonWebToken) GetPayload() JsonWebTokenPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *JsonWebToken) GetPayloadOk() (*JsonWebTokenPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *JsonWebToken) SetPayload(v JsonWebTokenPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *JsonWebToken) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSignature

`func (o *JsonWebToken) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *JsonWebToken) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *JsonWebToken) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *JsonWebToken) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *JsonWebToken) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *JsonWebToken) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetTokenType

`func (o *JsonWebToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *JsonWebToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *JsonWebToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *JsonWebToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *JsonWebToken) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *JsonWebToken) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetExpiresIn

`func (o *JsonWebToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *JsonWebToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *JsonWebToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *JsonWebToken) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetAccessToken

`func (o *JsonWebToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *JsonWebToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *JsonWebToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *JsonWebToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *JsonWebToken) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *JsonWebToken) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


