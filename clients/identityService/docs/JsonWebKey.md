# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **NullableString** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Use** | Pointer to **NullableString** |  | [optional] 
**Kty** | Pointer to **NullableString** |  | [optional] 
**E** | Pointer to **NullableString** |  | [optional] 
**N** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey() *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *JsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *JsonWebKey) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *JsonWebKey) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetNbf

`func (o *JsonWebKey) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *JsonWebKey) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *JsonWebKey) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *JsonWebKey) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetUse

`func (o *JsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *JsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### SetUseNil

`func (o *JsonWebKey) SetUseNil(b bool)`

 SetUseNil sets the value for Use to be an explicit nil

### UnsetUse
`func (o *JsonWebKey) UnsetUse()`

UnsetUse ensures that no value is present for Use, not even an explicit nil
### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *JsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### SetKtyNil

`func (o *JsonWebKey) SetKtyNil(b bool)`

 SetKtyNil sets the value for Kty to be an explicit nil

### UnsetKty
`func (o *JsonWebKey) UnsetKty()`

UnsetKty ensures that no value is present for Kty, not even an explicit nil
### GetE

`func (o *JsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *JsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### SetENil

`func (o *JsonWebKey) SetENil(b bool)`

 SetENil sets the value for E to be an explicit nil

### UnsetE
`func (o *JsonWebKey) UnsetE()`

UnsetE ensures that no value is present for E, not even an explicit nil
### GetN

`func (o *JsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *JsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *JsonWebKey) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *JsonWebKey) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


